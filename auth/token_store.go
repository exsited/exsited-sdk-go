package auth

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode/utf8"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
)

const (
	defaultTokenTTLSeconds = 3600
	// Refreshed this long before expiry, so a token never expires mid-flight.
	expirySkewSeconds = 60
	// The longest lifetime believed from the server. expiresAt is an absolute
	// stamp from whichever machine minted the token, so a clock running ahead
	// could make an expired one look fresh; bounding the window lets it reach the
	// API and be answered 401, which refresh-and-replay already handles.
	maxTokenTTLSeconds = 24 * 60 * 60

	defaultTokenCacheFilename = "exsited_go_token.json"
	// How old a lock file's mtime must be before age alone breaks it. Age is the
	// backstop: a stamp naming this host is judged by whether its process still
	// runs. Age is left for a stamp from another machine, whose pid means nothing
	// here, and for a zombie, which accepts signal 0 and so reads as alive for ever.
	staleLockMillis = 30_000
	// How long a caller waits for someone else's lock before the call is refused.
	// Far below staleLockMillis, so a caller never blocks waiting for the age gate.
	lockWaitMillis = 10_000
	lockPollMillis = 25
	// The stamp field naming the machine that wrote a lock. Keyed rather than
	// positional, so a stamp without one reads as carrying no host rather than
	// having another field misread as one.
	lockHostField = "host="
)

// MaxTokenCacheBytes is the most of the cache file this SDK will read, and the
// size above which it refuses the file instead.
//
// The file sits at a well-known per-user path any process of this user can append
// to, so the read is bounded. One entry carrying two 4 KB credentials is about
// 9 KB and an ordinary one is under 1 KB, so a megabyte holds far more credential
// sets than a machine has.
const MaxTokenCacheBytes = 1024 * 1024

// The two failures this store reports rather than swallowing. Match with
// errors.Is: both arrive wrapped in an *errors.SdkError, whose Unwrap answers the
// sentinel.
var (
	// ErrTokenCache means the cache file is present and cannot be used.
	//
	// Distinct from an absent cache, which costs one authentication. Both
	// mutations rebuild the whole file from what the read returned, so answering
	// "empty" for a file that could not be parsed would make the next write delete
	// every other namespace's entry, refresh tokens included.
	ErrTokenCache = errors.New("exsited: the token cache is present and cannot be used")
	// ErrTokenLock means there was no cross-process lock to take, so the work the
	// lock guards was refused rather than run without exclusion: two workers
	// writing the shared cache at once overwrite each other's entry.
	ErrTokenLock = errors.New("exsited: the token cache could not be locked across processes")
)

type TokenRecord struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    int64  `json:"expiresAt"`
}

func (t TokenRecord) NeedsRefresh() bool {
	return time.Now().UnixMilli() >= t.ExpiresAt-expirySkewSeconds*1000
}

// TokenStore is where the access and refresh tokens live between calls.
//
// Read answers false for an absent token, which costs one authentication, and an
// error only when a cache that is there cannot be used. Write refuses rather than
// rebuilding a shared file from a snapshot it could not take. WithLock refuses
// rather than running the work with no exclusion between processes.
//
// Clear removes one namespace's entry and leaves every other one alone. Write
// refuses an empty credential, so Clear is how a refused token leaves the cache;
// see TokenProvider.InvalidateAccessToken.
type TokenStore interface {
	Read(namespace string) (TokenRecord, bool, error)
	Write(namespace string, record TokenRecord) error
	Clear(namespace string) error
	WithLock(namespace string, work func() error) error
}

// MemoryTokenStore keeps the token for the life of the process.
type MemoryTokenStore struct {
	mutex sync.Mutex
	// renewal serialises the store side of a renewal, while mutex guards each
	// individual read and write. They are separate because WithLock holds one
	// across a callback that calls Write, and a sync.Mutex is not reentrant.
	renewal sync.Mutex
	records map[string]TokenRecord
}

func NewMemoryTokenStore() *MemoryTokenStore {
	return &MemoryTokenStore{records: map[string]TokenRecord{}}
}

// Read answers the record for namespace. The error is always nil: memory cannot
// be half-parsed, and the signature exists so the two stores are interchangeable.
func (s *MemoryTokenStore) Read(namespace string) (TokenRecord, bool, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	record, ok := s.records[namespace]
	return record, ok, nil
}

// Write stores the record. The error is always nil, for the reason Read gives.
func (s *MemoryTokenStore) Write(namespace string, record TokenRecord) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.records[namespace] = record
	return nil
}

// Clear forgets this namespace's record. The error is always nil, for the reason
// Read gives; clearing an absent entry is a no-op.
func (s *MemoryTokenStore) Clear(namespace string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.records, namespace)
	return nil
}

func (s *MemoryTokenStore) WithLock(_ string, work func() error) error {
	// Not shared with another process, so a mutex is the whole lock and cannot fail.
	s.renewal.Lock()
	defer s.renewal.Unlock()
	return work()
}

// FileTokenStore writes the token to a file guarded by a lock, so concurrent
// processes do not lose each other's cache entries.
//
// The lock is a sidecar file created with O_CREATE|O_EXCL, which POSIX and
// Windows both make atomic. A file rather than a directory, because a file
// carries a stamp, and the stamp is what makes recovery better than a guess
// about age.
//
// Nothing releases that file when the holder dies, so a leaked one is recovered
// two ways. A stamp naming this host is judged by whether its pid is still
// running. Everything else (another machine's pid, or a zombie, which accepts
// signal 0 for ever) falls back to the mtime gate at staleLockMillis.
//
// Failing to take the lock fails the call: WithLock returns an ErrTokenLock
// refusal and does not run the work. A caller who would rather have no
// coordination than no token sets InMemoryToken instead.
type FileTokenStore struct {
	// path is absolute and resolved once, at construction, so a later os.Chdir
	// cannot move the I/O away from the guard and the sidecar.
	path string
	// guard is per cache path rather than per store: see cachePathGuard.
	guard *cachePathGuard
	// heldStamp is the exact bytes this holder wrote into the sidecar, compared at
	// release so a sidecar somebody else created meanwhile is left alone. Written
	// and read under the guard's renewal mutex.
	heldStamp string
}

// cachePathGuard is the in-process half of the lock over one cache path.
//
// One instance per file, shared by every store pointed at that file, so two
// clients built from one set of credentials coordinate with each other.
type cachePathGuard struct {
	// mutex guards one individual read or write.
	mutex sync.Mutex
	// renewal is the in-process half of the cross-process lock. It is a different
	// mutex from the one above because WithLock holds it across a callback that
	// calls Write, and a sync.Mutex is not reentrant.
	renewal sync.Mutex
	// degradedReason is why cross-process locking is not in force, or nil while it
	// is. Written under renewal, read by CrossProcessLockingDegraded from any
	// goroutine, so it is atomic.
	degradedReason atomic.Pointer[string]
	// degradationAnnounced is whether the current unbroken run of refusals has
	// already been announced, so a permanently unlockable path does not write a
	// line per call. Re-armed after a success. Under renewal.
	degradationAnnounced bool
}

var (
	guardsByCachePath   = map[string]*cachePathGuard{}
	guardsRegistryMutex sync.Mutex
)

// guardForCachePath finds the one guard for a resolved path, creating it once.
func guardForCachePath(path string) *cachePathGuard {
	key := path
	if runtime.GOOS == "windows" {
		// Two spellings that differ only in case name one file here, and two
		// guards over one file are no guard at all.
		key = strings.ToLower(key)
	}
	guardsRegistryMutex.Lock()
	defer guardsRegistryMutex.Unlock()
	guard, found := guardsByCachePath[key]
	if !found {
		guard = &cachePathGuard{}
		guardsByCachePath[key] = guard
	}
	return guard
}

// resolveCachePath is the absolute real file a cache path names, worked out once.
//
// The directory is resolved through symlinks rather than the file, because the
// file usually does not exist yet on the first call, and resolving it would
// answer one path before it was written and another afterwards.
func resolveCachePath(path string) string {
	absolute, err := filepath.Abs(path)
	if err != nil {
		absolute = filepath.Clean(path)
	}
	directory, base := filepath.Split(absolute)
	if real, err := filepath.EvalSymlinks(filepath.Clean(directory)); err == nil {
		return filepath.Join(real, base)
	}
	return absolute
}

// lockSerial distinguishes two stores in one process stamping in the same
// nanosecond, which is what makes a stamp identify one holder rather than one
// instant.
var lockSerial uint64

// lockHostname names the machine writing a stamp, so a pid is only ever judged
// against the process table it was minted in. "" when the machine cannot name
// itself, or names itself with whitespace, which would split the stamp into
// fields that no longer line up; every pid rule then falls back to the age gate.
func lockHostname() string {
	name, err := os.Hostname()
	if err != nil || name == "" || strings.ContainsAny(name, " \t\r\n") {
		return ""
	}
	return name
}

func DefaultTokenCachePath() string {
	var directory string
	switch runtime.GOOS {
	case "windows":
		local := os.Getenv("LOCALAPPDATA")
		if local == "" {
			local = os.TempDir()
		}
		directory = filepath.Join(local, "exsited")
	case "darwin":
		home, _ := os.UserHomeDir()
		directory = filepath.Join(home, "Library", "Caches", "exsited")
	default:
		cache := os.Getenv("XDG_CACHE_HOME")
		if cache == "" {
			home, _ := os.UserHomeDir()
			cache = filepath.Join(home, ".cache")
		}
		directory = filepath.Join(cache, "exsited")
	}
	return filepath.Join(directory, defaultTokenCacheFilename)
}

// BuildCredentialsNamespace fingerprints a set of credentials, for keying a
// shared token file.
//
// Derived from every value that decides which principal the server mints for, so
// identical credentials share one cached token and any difference gets its own
// entry. RedirectURI is sent by the token request but selects no token, so it is
// not included. Each value is fed in with its length rather than joined on a
// separator, which would be forgeable.
//
// The digest keys entries already on disk, so treat it as frozen: one differing
// byte silently orphans every cached token.
func BuildCredentialsNamespace(baseURL, clientID, clientSecret, grantType string) string {
	// The leading "v4" is the fingerprint scheme label.
	values := []string{"v4", baseURL, clientID, clientSecret, "", grantType}
	material := ""
	for _, value := range values {
		// Counted in code points, not bytes; that is part of the frozen digest.
		material += strconv.Itoa(utf8.RuneCountInString(value)) + ":" + value
	}
	digest := sha256.Sum256([]byte(material))
	return hex.EncodeToString(digest[:])[:16]
}

// readAll is every namespace in the cache file, or an empty map only when the
// file is missing or empty.
//
// Anything else (unreadable, larger than MaxTokenCacheBytes, not a JSON object)
// is an ErrTokenCache refusal rather than an empty map, because Write rebuilds
// the whole file from what this returns.
//
// Entries are kept as raw JSON and decoded one at a time, so a namespace this
// build cannot make sense of is still written back byte for byte rather than
// dropped by a rewrite that only meant to add one entry.
func (s *FileTokenStore) readAll() (map[string]json.RawMessage, error) {
	contents, err := s.readWithinTheCeiling()
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]json.RawMessage{}, nil
		}
		return nil, err
	}
	if len(bytes.TrimSpace(contents)) == 0 {
		return map[string]json.RawMessage{}, nil
	}
	var all map[string]json.RawMessage
	if err := json.Unmarshal(contents, &all); err != nil {
		return nil, s.refuseTheCacheFile("does not hold a JSON object of credential namespaces", err)
	}
	if all == nil {
		return map[string]json.RawMessage{}, nil
	}
	return all, nil
}

// readWithinTheCeiling reads the cache file one byte past MaxTokenCacheBytes:
// enough to know the file is over the ceiling without holding the whole of it.
func (s *FileTokenStore) readWithinTheCeiling() ([]byte, error) {
	file, err := os.Open(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			// Passed through unwrapped, because the caller tells "absent", which is
			// ordinary, from "present and unusable", which is not.
			return nil, err
		}
		return nil, s.refuseTheCacheFile("could not be opened", err)
	}
	defer func() { _ = file.Close() }()
	contents, err := io.ReadAll(io.LimitReader(file, MaxTokenCacheBytes+1))
	if err != nil {
		return nil, s.refuseTheCacheFile("could not be read", err)
	}
	if len(contents) > MaxTokenCacheBytes {
		return nil, s.refuseTheCacheFile(fmt.Sprintf(
			"is larger than the %d-byte ceiling this SDK reads it to, so it was not loaded",
			MaxTokenCacheBytes), nil)
	}
	return contents, nil
}

// refuseTheCacheFile is the one refusal for a cache file that cannot be used.
func (s *FileTokenStore) refuseTheCacheFile(problem string, cause error) error {
	native := ""
	if cause != nil {
		native = ": " + cause.Error()
	}
	return &exerrors.SdkError{
		Kind: ErrTokenCache,
		Message: fmt.Sprintf(
			"the token cache %s. It is not read as empty: this file is keyed by "+
				"credentials namespace and is meant to be shared, so rewriting it from an "+
				"empty snapshot would discard every other namespace in it. Repair or move "+
				"the file, or set InMemoryToken to keep the token in memory. Path: %s%s",
			problem, s.path, native),
	}
}

// Read answers this namespace's record. Absent is (false, nil) and costs one
// authentication; a cache that is there and cannot be used is an error.
func (s *FileTokenStore) Read(namespace string) (TokenRecord, bool, error) {
	s.guard.mutex.Lock()
	defer s.guard.mutex.Unlock()
	all, err := s.readAll()
	if err != nil {
		return TokenRecord{}, false, err
	}
	raw, found := all[namespace]
	if !found {
		return TokenRecord{}, false, nil
	}
	var record TokenRecord
	if err := json.Unmarshal(raw, &record); err != nil {
		// One entry this build cannot read is not the file's problem: it stays on
		// disk untouched and this store re-authenticates. Refusing the whole file
		// for it would make one foreign entry lock every namespace out.
		return TokenRecord{}, false, nil
	}
	if err := validateTokenCredentials(record.AccessToken, record.RefreshToken); err != nil {
		// A cached credential that does not satisfy its grammar cannot be sent, so
		// it counts as none and a fresh token is minted.
		return TokenRecord{}, false, nil
	}
	return record, true, nil
}

// Write replaces this namespace's entry and leaves every other one alone.
//
// It refuses when the existing file could not be read, because the snapshot this
// write would rebuild the file from is then not the file's contents.
func (s *FileTokenStore) Write(namespace string, record TokenRecord) error {
	if err := validateTokenCredentials(record.AccessToken, record.RefreshToken); err != nil {
		return err
	}
	s.guard.mutex.Lock()
	defer s.guard.mutex.Unlock()
	all, err := s.readAll()
	if err != nil {
		return err
	}
	entry, err := json.Marshal(record)
	if err != nil {
		return &exerrors.SdkError{
			Kind:    ErrTokenCache,
			Message: "the token record could not be encoded: " + err.Error(),
		}
	}
	all[namespace] = entry
	encoded, err := json.Marshal(all)
	if err != nil {
		return &exerrors.SdkError{
			Kind:    ErrTokenCache,
			Message: "the token cache could not be encoded: " + err.Error(),
		}
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0o700); err != nil {
		return s.refuseTheCacheFile("could not be written: its directory does not exist", err)
	}
	// Written to a sibling and renamed, so a reader never sees half a file.
	temporary := s.path + "." + strconv.Itoa(os.Getpid()) + ".tmp"
	if err := os.WriteFile(temporary, encoded, 0o600); err != nil {
		return s.refuseTheCacheFile("could not be written", err)
	}
	if err := os.Rename(temporary, s.path); err != nil {
		_ = os.Remove(temporary)
		return s.refuseTheCacheFile("could not be replaced", err)
	}
	return nil
}

// Clear removes this namespace's entry and leaves every other one alone.
//
// It refuses when the existing file could not be read, for the reason Write gives.
// An absent entry is success, so a compare-and-clear that lost a race does not
// look like a broken cache. Other entries are written back as raw JSON, untouched.
func (s *FileTokenStore) Clear(namespace string) error {
	s.guard.mutex.Lock()
	defer s.guard.mutex.Unlock()
	all, err := s.readAll()
	if err != nil {
		return err
	}
	if _, found := all[namespace]; !found {
		return nil
	}
	delete(all, namespace)
	encoded, err := json.Marshal(all)
	if err != nil {
		return &exerrors.SdkError{
			Kind:    ErrTokenCache,
			Message: "the token cache could not be encoded: " + err.Error(),
		}
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0o700); err != nil {
		return s.refuseTheCacheFile("could not be written: its directory does not exist", err)
	}
	// Written to a sibling and renamed, so a reader never sees half a file.
	temporary := s.path + "." + strconv.Itoa(os.Getpid()) + ".clear.tmp"
	if err := os.WriteFile(temporary, encoded, 0o600); err != nil {
		return s.refuseTheCacheFile("could not be written", err)
	}
	if err := os.Rename(temporary, s.path); err != nil {
		_ = os.Remove(temporary)
		return s.refuseTheCacheFile("could not be replaced", err)
	}
	return nil
}

// WithLock runs work under the in-process mutex and the cross-process sidecar.
//
// It takes the guard's renewal mutex, not its plain mutex, which Read and Write
// take individually: work() calls Write, and a sync.Mutex is not reentrant.
//
// It fails closed: when there is no sidecar to take, work does not run and an
// ErrTokenLock refusal is returned.
//
// TokenProvider.renew wraps only the store write in this, not the token request,
// so concurrent processes can still mint more than one token; the lock prevents
// two of them writing the cache file at once.
func (s *FileTokenStore) WithLock(namespace string, work func() error) error {
	// Taken first and unconditionally, so even a refusal is serialised and
	// announced once rather than once per goroutine.
	s.guard.renewal.Lock()
	defer s.guard.renewal.Unlock()

	lockPath := s.lockPath(namespace)
	if refusal := s.acquireLockFile(lockPath); refusal != nil {
		return refusal
	}
	defer s.releaseLockFile(lockPath)
	return work()
}

func (s *FileTokenStore) lockPath(namespace string) string {
	return s.path + "." + namespace + ".lock"
}

// CrossProcessLockingDegraded is why the last attempt on this cache path could
// not take the cross-process lock, or "" while the lock is in force.
//
// The work is refused when this is set, so a caller does not have to poll it to
// stay safe; it says which condition it was, and lets a supervisor notice a path
// that never locks. Per cache path rather than per store.
func (s *FileTokenStore) CrossProcessLockingDegraded() string {
	if reason := s.guard.degradedReason.Load(); reason != nil {
		return *reason
	}
	return ""
}

// degrade records the condition, announces it once per unbroken run of refusals,
// and returns the refusal for the caller to return.
//
// Called with renewal held, which is what makes the plain read-modify-write of
// degradationAnnounced safe.
func (s *FileTokenStore) degrade(reason string) error {
	s.guard.degradedReason.Store(&reason)
	refusal := &exerrors.SdkError{
		Kind: ErrTokenLock,
		Message: fmt.Sprintf(
			"the token cache at %s could not be locked across processes, so the "+
				"operation was refused rather than run without exclusion: %s. Two "+
				"processes sharing this file would otherwise write it at the same time "+
				"and overwrite each other's entry, losing a refresh token that was still "+
				"valid. Set InMemoryToken to keep the token in memory, or point "+
				"TokenCachePath at a filesystem where locking works",
			s.path, reason),
	}
	if s.guard.degradationAnnounced {
		return refusal
	}
	// Once per unbroken run: a line per token refresh would flood the log, and a
	// condition that comes back after a recovery is announced again.
	s.guard.degradationAnnounced = true
	// stderr, so a diagnostic cannot corrupt piped output. No token is in it.
	fmt.Fprintf(os.Stderr, "exsited: %s\n", refusal.Message)
	return refusal
}

// acquireLockFile takes the sidecar, or returns the refusal saying why it could
// not. A non-nil answer means the work does not run.
func (s *FileTokenStore) acquireLockFile(lockPath string) error {
	deadline := time.Now().Add(lockWaitMillis * time.Millisecond)
	for {
		if err := os.MkdirAll(filepath.Dir(lockPath), 0o700); err != nil {
			return s.degrade(
				"the directory holding the token cache could not be created: " + err.Error())
		}
		file, err := os.OpenFile(lockPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)
		if err == nil {
			// Read back by breakStaleLock and releaseLockFile. Age comes from the
			// file's mtime, never from the stamp's timestamp, which a clock this
			// process does not control may have written.
			stamp := fmt.Sprintf("%d %s%s %s %d\n", os.Getpid(), lockHostField, lockHostname(),
				time.Now().UTC().Format(time.RFC3339), atomic.AddUint64(&lockSerial, 1))
			_, _ = file.WriteString(stamp)
			_ = file.Close()
			s.heldStamp = stamp
			// A success clears a passed degradation and re-arms its announcement.
			s.guard.degradedReason.Store(nil)
			s.guard.degradationAnnounced = false
			return nil
		}
		if !errors.Is(err, fs.ErrExist) {
			// An exclusive create can be refused without "already exists": Windows
			// answers a directory at the path with a permission denial. Anything
			// that exists is treated as somebody else's lock, so the stale rules
			// below can still clear it; a path that does not exist means the
			// filesystem cannot hold the lock, which no waiting changes.
			if _, statErr := os.Stat(lockPath); statErr != nil {
				return s.degrade("the lock file could not be created: " + err.Error())
			}
		}
		if breakStaleLock(lockPath) {
			continue
		}
		if !time.Now().Before(deadline) {
			return s.degrade(fmt.Sprintf("another holder kept the lock for longer than %s",
				lockWaitMillis*time.Millisecond))
		}
		time.Sleep(lockPollMillis * time.Millisecond)
	}
}

// breakStaleLock removes a leaked sidecar and reports whether it did. Exclusive
// creation, unlike an OS advisory lock, is not released when the holder dies, so
// one killed worker would otherwise stall every later token write.
//
// The prompt rule is a stamp naming this host whose pid is gone; a pid means
// nothing on another machine, and a cache path on a network mount puts two
// machines on one sidecar. Everything else waits for the age gate.
func breakStaleLock(lockPath string) bool {
	pid, host, stamped := readLockStamp(lockPath)
	// lockHostname answers "" when the machine cannot name itself, and two
	// unnameable hosts would compare equal.
	if stamped && host != "" && host == lockHostname() && !processAlive(pid) {
		return os.Remove(lockPath) == nil
	}
	info, err := os.Stat(lockPath)
	if err != nil || time.Since(info.ModTime()) < staleLockMillis*time.Millisecond {
		return false
	}
	return os.Remove(lockPath) == nil
}

// releaseLockFile removes the sidecar only while it still carries this holder's
// own stamp. After a takeover it belongs to whoever broke this holder's lock, and
// removing that one would let a third writer in behind them.
func (s *FileTokenStore) releaseLockFile(lockPath string) {
	held := s.heldStamp
	s.heldStamp = ""
	if held == "" {
		return
	}
	contents, err := os.ReadFile(lockPath)
	if err != nil || string(contents) != held {
		return
	}
	_ = os.Remove(lockPath)
}

// readLockStamp reads the pid and the host a sidecar records, and reports whether
// the stamp was readable and in this SDK's shape. The pid is the first field; the
// host is the lockHostField-prefixed one, and is "" for a stamp carrying none.
func readLockStamp(lockPath string) (pid int, host string, ok bool) {
	contents, err := os.ReadFile(lockPath)
	if err != nil {
		return 0, "", false
	}
	fields := strings.Fields(string(contents))
	if len(fields) == 0 {
		return 0, "", false
	}
	pid, err = strconv.Atoi(fields[0])
	if err != nil || pid <= 0 {
		return 0, "", false
	}
	for _, field := range fields[1:] {
		if after, found := strings.CutPrefix(field, lockHostField); found {
			host = after
			break
		}
	}
	return pid, host, true
}

// processAlive reports whether a pid names a running process, which is not the
// same question as whether a handle to it can be opened. processAliveOS holds the
// two platform halves.
func processAlive(pid int) bool {
	if pid <= 0 {
		return false
	}
	return processAliveOS(pid)
}

// NewFileTokenStore builds a store backed by the JSON file at path.
//
// The path is resolved to an absolute real path here and never re-interpreted,
// so a later os.Chdir cannot move the I/O away from the guard or the sidecar,
// and every store naming the one file shares one in-process guard.
func NewFileTokenStore(path string) *FileTokenStore {
	resolved := resolveCachePath(path)
	return &FileTokenStore{path: resolved, guard: guardForCachePath(resolved)}
}
