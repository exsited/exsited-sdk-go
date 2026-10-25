package config

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"math"
	"net/http"
	"net/netip"
	"net/url"
	"os"
	"reflect"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	"github.com/exsited/exsited-sdk-go/v4/redact"
)

// fallbackVersion is reported when the build carries no module version. Bump it on every release.
const fallbackVersion = "4.0.0"

var sdkVersion = readModuleVersion()

// SDKVersion is the SDK's own version, reported in the default User-Agent header.
func SDKVersion() string { return sdkVersion }

// Version is SDKVersion under a shorter name: the same value.
func Version() string { return sdkVersion }

// DefaultUserAgent is the User-Agent this SDK sends unless SdkConfig.UserAgent replaces it.
func DefaultUserAgent() string { return "exsited-go-sdk/" + sdkVersion }

func readModuleVersion() string {
	const modulePath = "github.com/exsited/exsited-sdk-go/v4"
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return fallbackVersion
	}
	candidates := append([]*debug.Module{&info.Main}, info.Deps...)
	for _, m := range candidates {
		if m == nil || m.Path != modulePath {
			continue
		}
		if m.Replace != nil {
			break // a replace directive's version is a placeholder, not a release
		}
		v := strings.TrimPrefix(m.Version, "v")
		if v == "" || v == "(devel)" || strings.HasPrefix(v, "0.0.0") {
			break // source checkout / untagged pseudo-version
		}
		return v
	}
	return fallbackVersion
}

// APIVersion is the resource version a path renders with; this SDK serves v4 only.
type APIVersion string

// V4 is the one resource version this SDK serves.
const V4 APIVersion = "v4"

// TokenPath is where a client of this version mints its token: /api/v4/oauth2/token.
func (v APIVersion) TokenPath() string {
	return "/api/v4/oauth2/token"
}

const (
	DefaultAPIVersion             = V4
	DefaultTimeoutSeconds         = 30.0
	DefaultConnectTimeoutSeconds  = 10.0
	DefaultMaxRetries             = 1
	DefaultRetryBackoffSeconds    = 0.5
	DefaultRetryBackoffMaxSeconds = 20.0
	DefaultPoolMaxsize            = 10
	DefaultLogLevel               = "WARNING"
	DefaultPrintFormat            = PrintFormatPretty

	// DefaultMaxResponseBytes caps how much of one response the SDK buffers in memory: 64 MiB.
	DefaultMaxResponseBytes = 64 << 20

	// PrintFormatPretty is the multi-line console print format.
	PrintFormatPretty = "pretty"
	// PrintFormatCompact is the single-line console print format.
	PrintFormatCompact = "compact"

	// ResponseFormatMap presents ApiResponse.Data as the decoded body: a map, a list or bytes.
	ResponseFormatMap = "map"
	// ResponseFormatJSON presents ApiResponse.Data as the server's own body text.
	ResponseFormatJSON = "json"
	// ResponseFormatDTO presents ApiResponse.Data as the endpoint's typed view.
	ResponseFormatDTO = "dto"

	DefaultResponseFormat = ResponseFormatMap

	// ResponseFormatEnvVar is the lowest-priority tier deciding ApiResponse.Data.
	ResponseFormatEnvVar = "EXSITED_RESPONSE_FORMAT"

	maxHTTPCodes = 500
)

// ResponseFormat is a response format name; the constants above are its values.
type ResponseFormat = string

// FormatMap, FormatJSON and FormatDto are aliases of the ResponseFormat constants.
const (
	FormatMap  = ResponseFormatMap
	FormatJSON = ResponseFormatJSON
	FormatDto  = ResponseFormatDTO
)

// Unexported, since an exported slice is writable; DefaultRetryableCodes and DefaultUnsafeRetryableCodes return copies.
var (
	defaultRetryableHTTPCodes = []int{429, 500, 502, 503, 504}
	// Empty: neither 429 nor 503 proves an unsafe request was refused rather than run.
	defaultUnsafeRetryableHTTPCodes = []int{}

	// Severity order, not alphabetical: the refusal message names them in this order.
	logLevelNames     = []string{"DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"}
	validLogLevels    = map[string]bool{"DEBUG": true, "INFO": true, "WARNING": true, "ERROR": true, "CRITICAL": true}
	validPrintFormats = map[string]bool{PrintFormatPretty: true, PrintFormatCompact: true}

	// The three response formats, in the order every refusal names them.
	responseFormatNames  = []string{ResponseFormatMap, ResponseFormatJSON, ResponseFormatDTO}
	validResponseFormats = map[string]bool{
		ResponseFormatMap: true, ResponseFormatJSON: true, ResponseFormatDTO: true,
	}

	// The proxy schemes net/http can dial through; any other fails at the first request.
	proxySchemes = map[string]bool{"http": true, "https": true, "socks5": true, "socks5h": true}

	// requiredAuthFields are the credentials keys AuthConfig.Validate refuses blank.
	requiredAuthFields = []string{"client_name", "base_url", "client_id", "client_secret", "redirect_uri"}
)

type SdkConfig struct {
	Auth AuthConfig

	// APIVersion is the API version segment; this SDK serves "v4" only.
	APIVersion APIVersion
	// TimeoutSeconds bounds one attempt, headers and body read alike, not the call as a whole.
	TimeoutSeconds        float64
	ConnectTimeoutSeconds float64
	// MaxRetries is total attempts including the first, so 1 disables retrying.
	MaxRetries int
	// RetryableHTTPCodes are statuses worth another attempt for a safe method.
	RetryableHTTPCodes []int
	// UnsafeRetryableHTTPCodes are statuses to retry even for an unsafe method.
	UnsafeRetryableHTTPCodes []int
	// RetryBackoffSeconds is the base backoff before the first retry; nil means 0.5, so Ptr(0.0) retries at once.
	RetryBackoffSeconds *float64
	// RetryBackoffMaxSeconds is the ceiling on any single backoff.
	RetryBackoffMaxSeconds float64

	// VerifySsl controls TLS certificate verification; nil means true, and Ptr(false) emits a warning.
	VerifySsl *bool
	// CABundle is a PEM file to trust instead of disabling verification; refused beside VerifySsl=false.
	CABundle string
	ProxyURL string
	// PoolMaxsize bounds idle connections per host.
	PoolMaxsize int
	// MaxResponseBytes caps how much of one response body the transport buffers; 0 takes the default.
	MaxResponseBytes int

	// DefaultHeaders are sent with every request; a Content-Type here labels every encoded body.
	DefaultHeaders map[string]string
	// RaiseOnError also returns an *SdkError for a failed response, judged by ApiResponse.Succeeded.
	RaiseOnError bool
	// ResponseFormat is what ApiResponse.Data follows: "map", "json" or "dto".
	ResponseFormat string
	UserAgent      string

	LogLevel  string
	LogBodies bool
	// LogFile also appends the client's log lines to this path; the handle is released by Close.
	LogFile string

	// The six print switches, off unless set, print the resolved URL, query values, body and a curl line.
	PrintApiInfo        bool
	PrintRequestPayload bool
	PrintResponseBody   bool
	PrintHeaders        bool
	PrintCurl           bool
	PrintOnlyOnError    bool
	// PrintFormat is "pretty" or "compact".
	PrintFormat string
	PrintColor  bool
	// PrintMaxChars truncates the rendered exchange block; 0 = unlimited.
	PrintMaxChars int

	// PrintRequests is an alias of PrintApiInfo; either one turns the block on.
	PrintRequests bool
	// DiagnosticWriter receives everything the SDK prints; nil keeps stdout for exchanges and stderr for warnings.
	DiagnosticWriter io.Writer
	// HTTPClient is sent through instead of a built one, with redirect-following removed from a copy of it.
	// TLS, proxy and pool settings cannot reach a client the SDK did not build, so they are refused beside it.
	HTTPClient *http.Client
}

func (c SdkConfig) BaseURL() string { return c.Auth.BaseURL }

func (c SdkConfig) ClientName() string { return c.Auth.ClientName }

func (c SdkConfig) ClientID() string { return c.Auth.ClientID }

// ClientSecret is the OAuth client secret, from Auth. Redacted by String, not here.
func (c SdkConfig) ClientSecret() string { return c.Auth.ClientSecret }

func (c SdkConfig) RedirectURI() string { return c.Auth.RedirectURI }

func (c SdkConfig) GrantType() string { return c.Auth.GrantType }

// TokenCachePath is where tokens are cached, from Auth. Empty means the platform default.
func (c SdkConfig) TokenCachePath() string { return c.Auth.TokenCachePath }

// TokenURL is the absolute token endpoint for this config.
func (c SdkConfig) TokenURL() string { return c.Auth.BaseURL + c.APIVersion.TokenPath() }

// Timeout is the connect and read timeouts, in that order.
func (c SdkConfig) Timeout() (connect, read float64) {
	return c.ConnectTimeoutSeconds, c.TimeoutSeconds
}

// DefaultHeadersDict is DefaultHeaders as a fresh map the caller may mutate freely.
func (c SdkConfig) DefaultHeadersDict() map[string]string {
	out := make(map[string]string, len(c.DefaultHeaders))
	maps.Copy(out, c.DefaultHeaders)
	return out
}

// VerifySslEnabled reads the tri-state VerifySsl with its default: nil means true.
func (c SdkConfig) VerifySslEnabled() bool { return c.VerifySsl == nil || *c.VerifySsl }

// ResponseFormatValue resolves the format: this config's, then EXSITED_RESPONSE_FORMAT, then the default.
func (c SdkConfig) ResponseFormatValue() string {
	if c.ResponseFormat != "" {
		return c.ResponseFormat
	}
	if fromEnvironment, set := environmentResponseFormat(); set {
		return fromEnvironment
	}
	return DefaultResponseFormat
}

// environmentResponseFormat reads EXSITED_RESPONSE_FORMAT trimmed and lower-cased; blank reads as unset.
func environmentResponseFormat() (string, bool) {
	value := strings.ToLower(strings.TrimSpace(os.Getenv(ResponseFormatEnvVar)))
	return value, value != ""
}

// RetryBackoffSecondsValue reads RetryBackoffSeconds with its default: nil means 0.5.
func (c SdkConfig) RetryBackoffSecondsValue() float64 {
	if c.RetryBackoffSeconds == nil {
		return DefaultRetryBackoffSeconds
	}
	return *c.RetryBackoffSeconds
}

// PrintsApiInfo reports whether the exchange block's info rows are on, by either spelling.
func (c SdkConfig) PrintsApiInfo() bool { return c.PrintApiInfo || c.PrintRequests }

// String renders every field in declaration order.
func (c SdkConfig) String() string {
	return fmt.Sprintf(
		"SdkConfig(auth=%s, apiVersion=%q, timeoutSeconds=%v, connectTimeoutSeconds=%v, maxRetries=%d, "+
			"retryableHttpCodes=%v, unsafeRetryableHttpCodes=%v, retryBackoffSeconds=%v, "+
			"retryBackoffMaxSeconds=%v, verifySsl=%v, caBundle=%s, proxyUrl=%s, poolMaxsize=%d, "+
			"maxResponseBytes=%d, "+
			"defaultHeaders=%v, raiseOnError=%v, responseFormat=%q, userAgent=%s, logLevel=%q, logBodies=%v, logFile=%s, "+
			"printApiInfo=%v, printRequestPayload=%v, printResponseBody=%v, printHeaders=%v, printCurl=%v, "+
			"printOnlyOnError=%v, printFormat=%q, printColor=%v, printMaxChars=%d, "+
			"printRequests=%v, diagnosticWriter=%s, httpClient=%s)",
		c.Auth, redact.RedactValue(string(c.APIVersion)), c.TimeoutSeconds, c.ConnectTimeoutSeconds,
		c.MaxRetries,
		c.RetryableHTTPCodes, c.UnsafeRetryableHTTPCodes, c.RetryBackoffSecondsValue(),
		c.RetryBackoffMaxSeconds, c.VerifySslEnabled(),
		// Described, not masked: these field names are not secret names, so key masking would miss them.
		redact.DescribePath(c.CABundle), redact.Describe(c.ProxyURL), c.PoolMaxsize,
		c.MaxResponseBytes,
		redact.RedactMapping(c.DefaultHeaders), c.RaiseOnError,
		// The resolved format, not the raw field, so a pasted config explains itself.
		redact.RedactValue(c.ResponseFormatValue()), redact.Describe(c.UserAgent),
		redact.RedactValue(c.LogLevel), c.LogBodies,
		redact.DescribePath(c.LogFile),
		c.PrintApiInfo, c.PrintRequestPayload, c.PrintResponseBody, c.PrintHeaders, c.PrintCurl,
		c.PrintOnlyOnError, redact.RedactValue(c.PrintFormat), c.PrintColor, c.PrintMaxChars,
		c.PrintRequests, describeHandle(c.DiagnosticWriter != nil), describeHandle(c.HTTPClient != nil))
}

func describeHandle(set bool) string {
	if set {
		return "<set>"
	}
	return "<unset>"
}

func (c SdkConfig) Normalised() SdkConfig {
	out := c
	out.Auth = normalisedAuth(c.Auth)
	out.APIVersion = normaliseAPIVersion(c.APIVersion)
	if out.TimeoutSeconds == 0 {
		out.TimeoutSeconds = DefaultTimeoutSeconds
	}
	if out.ConnectTimeoutSeconds == 0 {
		out.ConnectTimeoutSeconds = DefaultConnectTimeoutSeconds
	}
	if out.MaxRetries == 0 {
		out.MaxRetries = DefaultMaxRetries
	}
	// Copied unconditionally, so a caller mutating what they passed cannot reach a built client.
	out.RetryableHTTPCodes = copyCodes(c.RetryableHTTPCodes, defaultRetryableHTTPCodes)
	out.UnsafeRetryableHTTPCodes = copyCodes(c.UnsafeRetryableHTTPCodes, defaultUnsafeRetryableHTTPCodes)
	if c.DefaultHeaders != nil {
		headers := make(map[string]string, len(c.DefaultHeaders))
		maps.Copy(headers, c.DefaultHeaders)
		out.DefaultHeaders = headers
	}
	if out.RetryBackoffSeconds == nil {
		backoff := DefaultRetryBackoffSeconds
		out.RetryBackoffSeconds = &backoff
	} else {
		backoff := *c.RetryBackoffSeconds
		out.RetryBackoffSeconds = &backoff
	}
	if out.RetryBackoffMaxSeconds == 0 {
		out.RetryBackoffMaxSeconds = DefaultRetryBackoffMaxSeconds
	}
	// Trimmed here so validate and the transport agree: a blank proxyUrl would disable ProxyFromEnvironment.
	out.CABundle = normaliseOptionalPath(c.CABundle)
	out.ProxyURL = normaliseProxyURL(c.ProxyURL)
	if out.VerifySsl == nil {
		verify := true
		out.VerifySsl = &verify
	} else {
		verify := *c.VerifySsl
		out.VerifySsl = &verify
	}
	if out.PoolMaxsize == 0 {
		out.PoolMaxsize = DefaultPoolMaxsize
	}
	if out.MaxResponseBytes == 0 {
		out.MaxResponseBytes = DefaultMaxResponseBytes
	}
	if out.LogLevel == "" {
		out.LogLevel = DefaultLogLevel
	} else {
		out.LogLevel = strings.ToUpper(strings.TrimSpace(c.LogLevel))
	}
	if out.PrintFormat == "" {
		out.PrintFormat = DefaultPrintFormat
	} else {
		out.PrintFormat = strings.ToLower(strings.TrimSpace(c.PrintFormat))
	}
	// Not defaulted, unlike printFormat: blank must survive for the environment and default to decide.
	out.ResponseFormat = strings.ToLower(strings.TrimSpace(c.ResponseFormat))
	if out.UserAgent == "" {
		out.UserAgent = DefaultUserAgent()
	}
	return out
}

// normalisedAuth applies what AuthConfig.Validate normalises, a trimmed BaseURL and the default grant, to a copy.
func normalisedAuth(auth AuthConfig) AuthConfig {
	out := auth
	_ = out.Validate()
	return out
}

// Validate reports every problem at once, judging the config as NewExsitedClient would see it.
func (c SdkConfig) Validate() error {
	return c.Normalised().validate()
}

// validate is the body of Validate; it requires an already-Normalised receiver.
func (c SdkConfig) validate() error {
	var problems []string
	// details carries the described values on CauseText, which Explain renders and Error does not.
	var details []string

	problems = append(problems, headerProblems(c.DefaultHeaders)...)

	// AuthConfig has no constructor, so this is the only place its fields are checked.
	if err := c.Auth.Validate(); err != nil {
		problems = append(problems, err.Error())
		var authError *exerrors.SdkError
		if errors.As(err, &authError) && authError.CauseText != "" {
			details = append(details, authError.CauseText)
		}
	}
	if c.APIVersion != V4 {
		problems = append(problems, fmt.Sprintf(
			"apiVersion: expected 'v4' or '4', got '%s' - this SDK serves v4 only", c.APIVersion))
	}
	if c.MaxRetries < 1 {
		problems = append(problems, fmt.Sprintf("maxRetries: must be at least 1 (1 disables retry), got %d", c.MaxRetries))
	}
	if !finitePositive(c.TimeoutSeconds) {
		problems = append(problems, "timeoutSeconds: must be a finite number > 0")
	}
	if !finitePositive(c.ConnectTimeoutSeconds) {
		problems = append(problems, "connectTimeoutSeconds: must be a finite number > 0")
	}
	if IsNaNOrInf(c.RetryBackoffSecondsValue()) || c.RetryBackoffSecondsValue() < 0 {
		problems = append(problems, "retryBackoffSeconds: must be a finite number >= 0")
	}
	if IsNaNOrInf(c.RetryBackoffMaxSeconds) || c.RetryBackoffMaxSeconds < c.RetryBackoffSecondsValue() {
		problems = append(problems, "retryBackoffMaxSeconds: must be finite and >= retryBackoffSeconds")
	}
	if c.PoolMaxsize < 1 {
		problems = append(problems, "poolMaxsize: must be at least 1")
	}
	if c.MaxResponseBytes < 1 {
		// Only a negative value reaches here; 0 already became the default.
		problems = append(problems, fmt.Sprintf(
			"maxResponseBytes: must be at least 1 byte (0 takes the %d-byte default), got %d",
			DefaultMaxResponseBytes, c.MaxResponseBytes))
	}
	if c.ProxyURL != "" {
		// Refused at construction: a proxy that fails to parse and silently goes direct is worse.
		if p, err := url.Parse(c.ProxyURL); err != nil || p.Scheme == "" || p.Host == "" {
			problems = append(problems, "proxyUrl: is not a URL - expected something like "+
				"http://proxy.internal:8080 (scheme and host required)")
		} else if !proxySchemes[strings.ToLower(p.Scheme)] {
			problems = append(problems, fmt.Sprintf(
				"proxyUrl: scheme '%s' is not one net/http can proxy through; expected http, https, "+
					"socks5 or socks5h", redact.NeutraliseControls(p.Scheme)))
		}
	}
	validateHTTPCodes(&problems, "retryableHttpCodes", c.RetryableHTTPCodes)
	validateHTTPCodes(&problems, "unsafeRetryableHttpCodes", c.UnsafeRetryableHTTPCodes)
	if !validLogLevels[c.LogLevel] {
		problems = append(problems, fmt.Sprintf("logLevel: expected one of %s, got '%s'",
			strings.Join(logLevelNames, ", "), c.LogLevel))
	}
	if !validPrintFormats[c.PrintFormat] {
		problems = append(problems, fmt.Sprintf("printFormat: expected one of %s, %s, got '%s'", PrintFormatPretty, PrintFormatCompact, c.PrintFormat))
	}
	validateResponseFormat(&problems, c.ResponseFormat)
	if c.PrintMaxChars < 0 {
		problems = append(problems, "printMaxChars: must be >= 0 (0 = unlimited)")
	}
	if strings.TrimSpace(c.UserAgent) == "" {
		problems = append(problems, "userAgent: must not be blank")
	}
	if c.CABundle != "" {
		if !c.VerifySslEnabled() {
			problems = append(problems, "caBundle and verifySsl=false are contradictory: one supplies a "+
				"trust store, the other disables verification")
		}
		if info, err := os.Stat(c.CABundle); err != nil || info.IsDir() {
			problems = append(problems, fmt.Sprintf(
				"caBundle: %s does not name a readable file", redact.DescribePath(c.CABundle)))
			details = append(details, "caBundle: "+c.CABundle)
		}
	}
	if c.HTTPClient != nil {
		if named := transportSettingsNamed(c); len(named) > 0 {
			problems = append(problems, fmt.Sprintf(
				"httpClient: cannot be combined with %s - the SDK cannot apply those to a client it did "+
					"not build, and ignoring them would leave a setting that reads as configuration and "+
					"does nothing; configure them on your own http.Transport, or drop HTTPClient",
				strings.Join(named, ", ")))
		}
	}

	if len(problems) > 0 {
		refusal := exerrors.NewConfigError(strings.Join(problems, "; "))
		refusal.CauseText = strings.Join(details, "; ")
		return refusal
	}
	return nil
}

// transportSettingsNamed lists the transport settings a supplied HTTPClient would silently ignore.
func transportSettingsNamed(c SdkConfig) []string {
	var named []string
	if !c.VerifySslEnabled() {
		named = append(named, "verifySsl")
	}
	if c.CABundle != "" {
		named = append(named, "caBundle")
	}
	if c.ProxyURL != "" {
		named = append(named, "proxyUrl")
	}
	if c.PoolMaxsize != 0 && c.PoolMaxsize != DefaultPoolMaxsize {
		named = append(named, "poolMaxsize")
	}
	return named
}

// NewSdkConfig validates a config and returns it normalised, with the response format resolved once.
func NewSdkConfig(config SdkConfig) (*SdkConfig, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}
	resolved := config.Normalised()
	resolved.ResponseFormat = resolved.ResponseFormatValue()
	return &resolved, nil
}

func validateResponseFormat(problems *[]string, configured string) {
	if configured != "" {
		if !validResponseFormats[configured] {
			*problems = append(*problems, fmt.Sprintf("responseFormat: expected one of %s, got '%s'",
				strings.Join(responseFormatNames, ", "), configured))
		}
		return
	}
	fromEnvironment, set := environmentResponseFormat()
	if set && !validResponseFormats[fromEnvironment] {
		*problems = append(*problems, fmt.Sprintf(
			"%s: expected one of %s, got '%s' - refused here rather than quietly falling back to '%s', "+
				"which would leave every response in a format nobody asked for",
			ResponseFormatEnvVar, strings.Join(responseFormatNames, ", "), fromEnvironment,
			DefaultResponseFormat))
	}
}

// ParseResponseFormat reads a format name, trimmed and case-insensitive, refusing anything but the three.
func ParseResponseFormat(value string) (string, error) {
	format := strings.ToLower(strings.TrimSpace(value))
	if !validResponseFormats[format] {
		return DefaultResponseFormat, exerrors.NewUsageError(fmt.Sprintf(
			"responseFormat: expected one of %s, got '%s'", strings.Join(responseFormatNames, ", "), value))
	}
	return format, nil
}

func validateHTTPCodes(problems *[]string, name string, codes []int) {
	*problems = append(*problems, HTTPCodeSetProblems(name, codes)...)
}

func HTTPCodeSetProblems(name string, codes []int) []string {
	if len(codes) > maxHTTPCodes {
		return []string{fmt.Sprintf(
			"%s: at most %d status codes (100-599 is every code there is), got more than that",
			name, maxHTTPCodes)}
	}
	var problems []string
	for _, code := range codes {
		if code < 100 || code > 599 {
			problems = append(problems, fmt.Sprintf("%s: %d is not an HTTP status code", name, code))
		}
	}
	return problems
}

// headerProblems refuses a DefaultHeaders entry the transport could not send.
func headerProblems(headers map[string]string) []string {
	var problems []string
	for name, value := range headers {
		problem := HeaderNameProblem(name)
		if problem == "" {
			problem = HeaderValueProblem(value)
		}
		if problem != "" {
			problems = append(problems, fmt.Sprintf(
				"defaultHeaders: '%s' cannot be sent as a header - %s. The value is not shown: "+
					"it may be a credential", redact.NeutraliseControls(name), problem))
		}
	}
	sort.Strings(problems)
	return problems
}

// The one header grammar this SDK applies: RFC 9110's, narrowed to printable ASCII.

// HeaderNameProblem names what makes a header name unsendable, or "" when it is sendable.
func HeaderNameProblem(name string) string {
	if name == "" {
		return "a blank name"
	}
	if !IsHTTPToken(name) {
		return "a name outside the RFC 9110 token set, which is letters, digits and " +
			"!#$%&'*+-.^_`|~ - a space, a colon and a non-ASCII byte are all refused there"
	}
	return ""
}

// HeaderValueProblem names what makes a header value unsendable, or "" when it is sendable.
func HeaderValueProblem(value string) string {
	if value == "" {
		return ""
	}
	if value[0] == ' ' || value[0] == '\t' {
		return "a value opening with a space or a tab, which several HTTP clients trim, so the " +
			"value on the wire would not be the one written"
	}
	for index := 0; index < len(value); index++ {
		byteValue := value[index]
		switch {
		case byteValue == '\t':
			continue // legal inside a field value; only a leading or trailing one is refused
		case byteValue == '\r':
			return fmt.Sprintf("a carriage return at byte %d, which would end the header line "+
				"and let whatever follows be read as a header of its own", index)
		case byteValue == '\n':
			return fmt.Sprintf("a newline at byte %d, which would end the header line and let "+
				"whatever follows be read as a header of its own", index)
		case byteValue == 0x00:
			return fmt.Sprintf("a NUL at byte %d", index)
		case byteValue < 0x20 || byteValue == 0x7f:
			return fmt.Sprintf("a control character at byte %d", index)
		case byteValue >= 0x80:
			return fmt.Sprintf("a non-ASCII byte at byte %d; a header value has no charset on "+
				"the wire, so encode it yourself - base64, or RFC 8187 percent-encoding", index)
		}
	}
	// Last, so a trailing CR or NUL is named as a smuggling primitive rather than as whitespace.
	if last := value[len(value)-1]; last == ' ' || last == '\t' {
		return "a value ending in a space or a tab, which several HTTP clients trim, so the " +
			"value on the wire would not be the one written"
	}
	return ""
}

// IsHTTPToken reports whether a value is an RFC 9110 token, the grammar of header names and methods.
func IsHTTPToken(value string) bool {
	for index := 0; index < len(value); index++ {
		c := value[index]
		allowed := (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') ||
			strings.IndexByte("!#$%&'*+-.^_`|~", c) >= 0
		if !allowed {
			return false
		}
	}
	return true
}

func normaliseAPIVersion(raw APIVersion) APIVersion {
	if raw == "" {
		return DefaultAPIVersion
	}
	value := strings.ToLower(strings.TrimSpace(string(raw)))
	if value == "" {
		return ""
	}
	if !strings.HasPrefix(value, "v") {
		value = "v" + value
	}
	return APIVersion(value)
}

// normaliseOptionalPath keeps a whitespace-only path verbatim, so it fails the file check.
func normaliseOptionalPath(raw string) string {
	if trimmed := strings.TrimSpace(raw); trimmed != "" || raw == "" {
		return trimmed
	}
	return raw
}

// DeepCopy returns a copy that shares no mutable state with the receiver; the two handles stay shared.
func (c SdkConfig) DeepCopy() SdkConfig {
	out := c
	if c.RetryableHTTPCodes != nil {
		out.RetryableHTTPCodes = append(make([]int, 0, len(c.RetryableHTTPCodes)), c.RetryableHTTPCodes...)
	}
	if c.UnsafeRetryableHTTPCodes != nil {
		out.UnsafeRetryableHTTPCodes = append(make([]int, 0, len(c.UnsafeRetryableHTTPCodes)), c.UnsafeRetryableHTTPCodes...)
	}
	if c.DefaultHeaders != nil {
		headers := make(map[string]string, len(c.DefaultHeaders))
		maps.Copy(headers, c.DefaultHeaders)
		out.DefaultHeaders = headers
	}
	if c.RetryBackoffSeconds != nil {
		backoff := *c.RetryBackoffSeconds
		out.RetryBackoffSeconds = &backoff
	}
	if c.VerifySsl != nil {
		verify := *c.VerifySsl
		out.VerifySsl = &verify
	}
	return out
}

// WithOverrides derives a new config, changing the snake_case-named fields in changes.
func (c SdkConfig) WithOverrides(changes map[string]any) (SdkConfig, error) {
	return applyOverrides(c, changes)
}

func SdkConfigFromCredentials(credentials []any, instance any, overrides map[string]any) (SdkConfig, error) {
	blocks := make([]any, len(credentials))
	copy(blocks, credentials)
	chosen, err := SelectCredentials(blocks, instance)
	if err != nil {
		return SdkConfig{}, err
	}

	switch typed := chosen.(type) {
	case SdkConfig:
		return applyOverrides(typed, overrides)
	case *SdkConfig:
		if typed == nil {
			return SdkConfig{}, nilEntry(instance, "*config.SdkConfig")
		}
		return applyOverrides(*typed, overrides)
	case AuthConfig:
		return applyOverrides(SdkConfig{Auth: typed}, overrides)
	case *AuthConfig:
		if typed == nil {
			return SdkConfig{}, nilEntry(instance, "*config.AuthConfig")
		}
		return applyOverrides(SdkConfig{Auth: *typed}, overrides)
	case map[string]any:
		return fromMapping(typed, instance, overrides)
	default:
		return SdkConfig{}, exerrors.NewConfigError(fmt.Sprintf(
			"credentials[%s]: expected an SdkConfig, an AuthConfig or a mapping of config "+
				"fields, got %T", selectorLabel(instance), chosen))
	}
}

func nilEntry(instance any, what string) error {
	return exerrors.NewConfigError(fmt.Sprintf(
		"credentials[%s]: the entry is a nil %s", selectorLabel(instance), what))
}

// selectorLabel names the selector for a message; an absent one reads as 0, the entry it chose.
func selectorLabel(instance any) string {
	if instance == nil {
		return "0"
	}
	if name, ok := instance.(string); ok {
		return fmt.Sprintf("%q", name)
	}
	return fmt.Sprintf("%v", instance)
}

func fromMapping(block map[string]any, instance any, overrides map[string]any) (SdkConfig, error) {
	remaining := make(map[string]any, len(block))
	maps.Copy(remaining, block)

	// "auth" is a field of the config but not of AuthConfig, so it comes out first.
	suppliedAuth, hasAuth := remaining["auth"]
	delete(remaining, "auth")

	authNames := authFieldNames()
	authValues := map[string]any{}
	for key := range remaining {
		if authNames[key] {
			authValues[key] = remaining[key]
			delete(remaining, key)
		}
	}

	settings := make(map[string]any, len(remaining)+len(overrides))
	maps.Copy(settings, remaining)
	maps.Copy(settings, overrides)

	if hasAuth && len(authValues) > 0 {
		named := make([]string, 0, len(authValues))
		for key := range authValues {
			named = append(named, key)
		}
		sort.Strings(named)
		return SdkConfig{}, exerrors.NewConfigError(
			"auth: an AuthConfig was supplied alongside the credential field(s) " +
				strings.Join(named, ", ") + "; pass one or the other")
	}

	var config SdkConfig
	if !hasAuth {
		// Checked here, since a missing required key would otherwise fail deeper, naming no field.
		if absent := absentRequiredAuthFields(authValues); len(absent) > 0 {
			return SdkConfig{}, exerrors.NewConfigError(
				"credentials: missing required field(s): " + strings.Join(absent, ", "))
		}
		auth, err := authFromMapping(authValues)
		if err != nil {
			return SdkConfig{}, err
		}
		config.Auth = auth
	} else {
		switch typed := suppliedAuth.(type) {
		case AuthConfig:
			config.Auth = typed
		case *AuthConfig:
			if typed == nil {
				return SdkConfig{}, nilEntry(instance, "*config.AuthConfig")
			}
			config.Auth = *typed
		default:
			return SdkConfig{}, exerrors.NewConfigError(fmt.Sprintf(
				"credentials[%s]: auth must be an AuthConfig, got %T",
				selectorLabel(instance), suppliedAuth))
		}
	}
	return applyOverrides(config, settings)
}

func authFieldNames() map[string]bool {
	names := map[string]bool{}
	authType := reflect.TypeFor[AuthConfig]()
	for index := 0; index < authType.NumField(); index++ {
		names[snakeCase(authType.Field(index).Name)] = true
	}
	return names
}

func absentRequiredAuthFields(supplied map[string]any) []string {
	var absent []string
	for _, name := range requiredAuthFields {
		if _, found := supplied[name]; !found {
			absent = append(absent, name)
		}
	}
	return absent
}

func authFromMapping(values map[string]any) (AuthConfig, error) {
	var auth AuthConfig
	if err := assignFields(reflect.ValueOf(&auth).Elem(), values, "credentials"); err != nil {
		return AuthConfig{}, err
	}
	return auth, nil
}

func applyOverrides(config SdkConfig, overrides map[string]any) (SdkConfig, error) {
	if len(overrides) > 0 {
		if err := assignFields(reflect.ValueOf(&config).Elem(), overrides, "overrides"); err != nil {
			return SdkConfig{}, err
		}
	}
	if err := config.Validate(); err != nil {
		return SdkConfig{}, err
	}
	return config.Normalised(), nil
}

func assignFields(target reflect.Value, values map[string]any, where string) error {
	targetType := target.Type()
	byName := map[string]int{}
	for index := 0; index < targetType.NumField(); index++ {
		byName[snakeCase(targetType.Field(index).Name)] = index
	}
	names := make([]string, 0, len(values))
	for key := range values {
		names = append(names, key)
	}
	sort.Strings(names)
	unknown := make([]string, 0)
	for _, key := range names {
		index, found := byName[key]
		if !found {
			unknown = append(unknown, key)
			continue
		}
		if err := assignOne(target.Field(index), values[key], key); err != nil {
			return err
		}
	}
	if len(unknown) > 0 {
		known := make([]string, 0, len(byName))
		for name := range byName {
			known = append(known, name)
		}
		sort.Strings(known)
		return exerrors.NewConfigError("unknown config field(s): " + strings.Join(unknown, ", ") +
			" (in " + where + "); known: " + strings.Join(known, ", "))
	}
	return nil
}

// assignOne converts one loader-shaped value onto one field, refusing by name rather than panicking.
func assignOne(field reflect.Value, value any, name string) error {
	if value == nil {
		return nil // absent and explicitly null read alike: leave the zero value
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(field.Type()) {
		field.Set(incoming)
		return nil
	}
	switch field.Kind() {
	case reflect.Float64:
		if number, ok := asFloat(value); ok {
			field.SetFloat(number)
			return nil
		}
	case reflect.Int:
		if number, ok := asFloat(value); ok && number == float64(int(number)) {
			field.SetInt(int64(number))
			return nil
		}
	case reflect.String:
		if text, ok := value.(string); ok {
			field.SetString(text)
			return nil
		}
		// A non-string scalar renders as text, so {"default_headers": {"X-Retry-Count": 3}} works.
		if text, ok := scalarText(value); ok {
			field.SetString(text)
			return nil
		}
	case reflect.Bool:
		if flag, ok := value.(bool); ok {
			field.SetBool(flag)
			return nil
		}
	case reflect.Pointer:
		// A loader supplies the bare value for a *bool / *float64 field; this wraps it.
		holder := reflect.New(field.Type().Elem())
		if err := assignOne(holder.Elem(), value, name); err != nil {
			return err
		}
		field.Set(holder)
		return nil
	case reflect.Slice:
		if items, ok := value.([]any); ok {
			out := reflect.MakeSlice(field.Type(), 0, len(items))
			for _, item := range items {
				element := reflect.New(field.Type().Elem()).Elem()
				if err := assignOne(element, item, name); err != nil {
					return err
				}
				out = reflect.Append(out, element)
			}
			field.Set(out)
			return nil
		}
	case reflect.Map:
		if entries, ok := value.(map[string]any); ok {
			out := reflect.MakeMapWithSize(field.Type(), len(entries))
			for key, entry := range entries {
				element := reflect.New(field.Type().Elem()).Elem()
				if err := assignOne(element, entry, name); err != nil {
					return err
				}
				out.SetMapIndex(reflect.ValueOf(key), element)
			}
			field.Set(out)
			return nil
		}
	}
	return exerrors.NewConfigError(fmt.Sprintf(
		"%s: expected %s, got %T", name, field.Type(), value))
}

func scalarText(value any) (string, bool) {
	switch v := value.(type) {
	case bool:
		if v {
			return "True", true
		}
		return "False", true
	case int:
		return strconv.Itoa(v), true
	case int64:
		return strconv.FormatInt(v, 10), true
	case float64:
		return exerrors.FormatSeconds(v), true
	case float32:
		return exerrors.FormatSeconds(float64(v)), true
	}
	return "", false
}

func asFloat(value any) (float64, bool) {
	switch number := value.(type) {
	case float64:
		return number, true
	case float32:
		return float64(number), true
	case int:
		return float64(number), true
	case int64:
		return float64(number), true
	}
	return 0, false
}

// snakeCase renders a Go field name as a credentials key, a run of capitals as one word: api_version.
func snakeCase(name string) string {
	runes := []rune(name)
	var out strings.Builder
	for index, r := range runes {
		if r >= 'A' && r <= 'Z' && index > 0 {
			previousLower := runes[index-1] >= 'a' && runes[index-1] <= 'z'
			nextLower := index+1 < len(runes) && runes[index+1] >= 'a' && runes[index+1] <= 'z'
			if previousLower || nextLower {
				out.WriteByte('_')
			}
		}
		out.WriteRune(lowered(r))
	}
	return out.String()
}

func lowered(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func IsResponseFormat(value string) bool { return validResponseFormats[value] }

// ResponseFormatNames returns a fresh copy of the three legal response formats, in refusal order.
func ResponseFormatNames() []string {
	return append(make([]string, 0, len(responseFormatNames)), responseFormatNames...)
}

func DefaultRetryableCodes() []int { return copyCodes(defaultRetryableHTTPCodes, nil) }

func DefaultUnsafeRetryableCodes() []int { return copyCodes(defaultUnsafeRetryableHTTPCodes, nil) }

// copyCodes substitutes fallback only for nil: an empty slice means retry nothing.
func copyCodes(codes, fallback []int) []int {
	if codes == nil {
		codes = fallback
	}
	return append(make([]int, 0, len(codes)), codes...)
}

// normaliseProxyURL gives a schemeless value an http:// scheme, so "proxy.internal:8080" works.
func normaliseProxyURL(raw string) string {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" || strings.Contains(trimmed, "://") {
		return trimmed
	}
	if p, err := url.Parse(trimmed); err == nil && p.Host != "" {
		return trimmed
	}
	return "http://" + trimmed
}

func finitePositive(value float64) bool { return !IsNaNOrInf(value) && value > 0 }

func IsNaNOrInf(value float64) bool { return math.IsNaN(value) || math.IsInf(value, 0) }

// portProblem describes what is wrong with the URL's port, or "" when nothing is; no port at all is fine.
func portProblem(parsed *url.URL) string {
	port := parsed.Port()
	if port == "" {
		// url.Parse treats a bare trailing colon as a valid empty port, so the host keeps it.
		if strings.HasSuffix(parsed.Host, ":") {
			return "port is empty - drop the trailing colon or give a number from 1 to 65535"
		}
		return ""
	}
	number, err := strconv.Atoi(port)
	if err != nil || number < 1 || number > 65535 {
		return "port must be a number from 1 to 65535, got " + port
	}
	return ""
}

// joinProblems renders a batched configuration report.
func joinProblems(instance string, problems []string) string {
	return "instance " + instance + ": " + strings.Join(problems, "; ")
}

// loopbackHostnames are the names for which http:// reaches nothing but this machine.
var loopbackHostnames = map[string]bool{
	"localhost":             true,
	"localhost.localdomain": true,
}

// IsLoopbackHost reports whether a host name or address refers only to this machine.
func IsLoopbackHost(hostname string) bool {
	if hostname == "" {
		return false
	}
	if loopbackHostnames[strings.ToLower(hostname)] {
		return true
	}
	// url.URL.Hostname has already stripped the brackets from an IPv6 literal.
	if address, err := netip.ParseAddr(hostname); err == nil {
		return address.IsLoopback()
	}
	return false
}

// PlaintextTransportWarning is the warning an http:// base URL to a remote host earns, or "" for none.
func PlaintextTransportWarning(baseURL string) string {
	parsed, err := url.Parse(strings.TrimSpace(baseURL))
	if err != nil || parsed.Scheme != "http" || IsLoopbackHost(parsed.Hostname()) {
		return ""
	}
	// The host is not named.
	return "baseUrl uses http://, so every request to a host that is not this machine crosses the " +
		"network in cleartext - the bearer token, and the credentials in the token request itself. " +
		"Use https:// unless this is a local development instance"
}
