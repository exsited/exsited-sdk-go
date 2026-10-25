package http

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	excfg "github.com/exsited/exsited-sdk-go/v4/config"
	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	"github.com/exsited/exsited-sdk-go/v4/redact"
)

const contentTypeHeader = "Content-Type"

type HTTPTransport struct {
	config     excfg.SdkConfig
	retry      *RetryPolicy
	client     *http.Client
	onProgress func(string)
}

type TransportOption func(*transportOptions)

type transportOptions struct {
	retry      *RetryPolicy
	client     *http.Client
	onProgress func(string)
}

// WithRetryPolicy replaces the policy the config would derive; nil leaves the config's own.
func WithRetryPolicy(retry *RetryPolicy) TransportOption {
	return func(o *transportOptions) { o.retry = retry }
}

// WithHTTPClient injects the net/http client to send through, over SdkConfig.HTTPClient; only its CheckRedirect is replaced.
func WithHTTPClient(client *http.Client) TransportOption {
	return func(o *transportOptions) { o.client = client }
}

// WithProgress wires the per-attempt progress hook.
func WithProgress(onProgress func(string)) TransportOption {
	return func(o *transportOptions) { o.onProgress = onProgress }
}

// NewHTTPTransport builds a transport for a config, validating it first.
func NewHTTPTransport(config excfg.SdkConfig, options ...TransportOption) (*HTTPTransport, error) {
	c := config.Normalised()
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return newHTTPTransport(c, options...)
}

func newHTTPTransport(config excfg.SdkConfig, options ...TransportOption) (*HTTPTransport, error) {
	applied := &transportOptions{}
	for _, option := range options {
		option(applied)
	}
	retry := applied.retry
	if retry == nil {
		derived, err := NewRetryPolicyFromConfig(config)
		if err != nil {
			panic(err) // config was validated upstream, so this is a programmer error
		}
		retry = derived
	}
	frozenRetry, err := retry.frozen()
	if err != nil {
		return nil, err
	}
	retry = frozenRetry
	client := applied.client
	if client == nil {
		client = config.HTTPClient
	}
	if client == nil {
		built, err := buildHTTPClient(config)
		if err != nil {
			return nil, err
		}
		client = built
	} else {
		client = withoutRedirects(client)
	}
	onProgress := applied.onProgress
	if onProgress == nil {
		onProgress = func(string) {}
	}
	return &HTTPTransport{config: config, retry: retry, client: client, onProgress: onProgress}, nil
}

// withoutRedirects copies a caller's client so it cannot follow a redirect, sharing its Transport.
func withoutRedirects(client *http.Client) *http.Client {
	copied := *client
	copied.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	return &copied
}

func buildHTTPClient(config excfg.SdkConfig) (*http.Client, error) {
	tlsConfig := &tls.Config{InsecureSkipVerify: !config.VerifySslEnabled()} // #nosec G402 -- opt-in, and off only when the caller sets VerifySsl=false
	if config.CABundle != "" {
		// Neither refusal below renders the path or the OS's own words; both ride CauseText for Explain.
		pem, err := readCappedFile(config.CABundle, MaxCABundleBytes)
		if err != nil {
			refusal := exerrors.NewConfigError(fmt.Sprintf(
				"caBundle: %s could not be read", redact.DescribePath(config.CABundle)))
			refusal.Cause = err
			refusal.CauseText = fmt.Sprintf("%s: %v", config.CABundle, err)
			return nil, refusal
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			refusal := exerrors.NewConfigError(fmt.Sprintf(
				"caBundle: %s holds no PEM certificates", redact.DescribePath(config.CABundle)))
			refusal.CauseText = config.CABundle
			return nil, refusal
		}
		tlsConfig.RootCAs = pool
	}
	dialer := &net.Dialer{Timeout: durationSeconds(config.ConnectTimeoutSeconds)}
	transport := &http.Transport{
		DialContext:         dialer.DialContext,
		MaxIdleConns:        config.PoolMaxsize * 2,
		MaxIdleConnsPerHost: config.PoolMaxsize,
		ForceAttemptHTTP2:   true,
		TLSClientConfig:     tlsConfig,
	}
	if config.ProxyURL != "" {
		// A backstop: validation already refused an unparseable value, so fail loudly rather than go direct.
		p, err := url.Parse(config.ProxyURL)
		if err != nil {
			// Unwrapped: *url.Error prepends the whole input, and a proxy URL often carries user:password@.
			var parseErr *url.Error
			if errors.As(err, &parseErr) {
				err = parseErr.Err
			}
			return nil, exerrors.NewConfigError(fmt.Sprintf("proxyUrl: is not a URL (%v)", err))
		}
		transport.Proxy = http.ProxyURL(p)
	} else {
		transport.Proxy = http.ProxyFromEnvironment
	}
	return &http.Client{
		Transport:     transport,
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
	}, nil
}

// MaxCABundleBytes caps how much of a CABundle PEM file is read into memory: 4 MiB.
const MaxCABundleBytes = 4 << 20

func ReadCappedFile(path string, limitBytes int) ([]byte, error) {
	return readCappedFile(path, limitBytes)
}

func readCappedFile(path string, limitBytes int) ([]byte, error) {
	handle, err := os.Open(path) //#nosec G304 -- the path is the caller's own configured path
	if err != nil {
		return nil, err
	}
	defer handle.Close() //nolint:errcheck // read-only handle
	content, err := io.ReadAll(io.LimitReader(handle, int64(limitBytes)+1))
	if err != nil {
		return nil, err
	}
	if len(content) > limitBytes {
		return nil, fmt.Errorf("the file is larger than the %d-byte cap this SDK reads it under",
			limitBytes)
	}
	return content, nil
}

type rawResponse struct {
	status      int
	headers     map[string]string
	contentType string
	retryAfter  string
	body        []byte
}

// SendHTTPRequest sends one request, retried per the policy and decoded by content type.
func (t *HTTPTransport) SendHTTPRequest(ctx context.Context, request TransportRequest) (*ApiResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	method, err := NormaliseMethod(request.HTTPMethod)
	if err != nil {
		return nil, err
	}
	// Preparation runs before the attempt timer, so the operation deadline bounds it; read here and after assembly.
	if err := refuseIfBudgetSpent(ctx, "before the request body was assembled"); err != nil {
		return nil, err
	}

	request, err = request.Materialised()
	if err != nil {
		return nil, err
	}
	if err := CheckRequestHeaders(request.Headers); err != nil {
		return nil, err
	}
	fullURL, err := buildTransportURL(request.URL, request.QueryParams)
	if err != nil {
		return nil, err
	}
	if err := refuseIfBudgetSpent(ctx, "while the request body was being assembled"); err != nil {
		return nil, err
	}

	started := time.Now()
	attempt := 0
	requestsSent := 0
	var lastError error

	for attempt < t.retry.MaxAttempts {
		attempt++
		requestsSent++
		// Named by endpoint template, not resolved URL: a progress hook reaches stdout and the log alike.
		t.onProgress(fmt.Sprintf("%s %s sending (attempt %d/%d, connect %gs, read %gs)",
			method, exerrors.EndpointLabel(request.EndpointTemplate, fullURL),
			attempt, t.retry.MaxAttempts,
			t.config.ConnectTimeoutSeconds, t.config.TimeoutSeconds))

		raw, sendError := t.doOnce(ctx, method, fullURL, request)
		if sendError != nil {
			lastError = sendError
			if errors.Is(sendError, errResponseTooLarge) {
				// Returned here: the shared exit below would relabel the cap as a network failure.
				info := t.requestInfo(method, request.URL, request, requestsSent,
					t.retry.MaxAttempts, started)
				return nil, &exerrors.SdkError{
					// The refusal's own sentence only; Error() renders the method and endpoint off the record.
					Message:     sendError.Error(),
					ErrorType:   ErrorTypeUnknown,
					RequestInfo: info,
					Cause:       sendError,
				}
			}
			if attempt >= t.retry.MaxAttempts || !t.retry.ShouldRetryAfterException(sendError, method) {
				break
			}
			delay := t.retry.ComputeBackoffSeconds(attempt, "")
			// Reported before the sleep: an unannounced wait reads as a hung process.
			t.onProgress(fmt.Sprintf("%T; retrying in %.1fs (attempt %d/%d)",
				sendError, delay, attempt+1, t.retry.MaxAttempts))
			if delayError := t.delay(ctx, delay); delayError != nil {
				lastError = delayError
				break
			}
			continue
		}

		if attempt < t.retry.MaxAttempts && t.retry.ShouldRetryStatusCode(raw.status, method) {
			delay := t.retry.ComputeBackoffSeconds(attempt, raw.retryAfter)
			t.onProgress(fmt.Sprintf("%d; retrying in %.1fs (attempt %d/%d)",
				raw.status, delay, attempt+1, t.retry.MaxAttempts))
			if delayError := t.delay(ctx, delay); delayError != nil {
				lastError = delayError
				break
			}
			continue
		}

		decoded := DecodeResponseBody(raw.status, raw.body, raw.contentType)
		response, buildErr := NewApiResponse(raw.status, decoded, raw.headers, raw.contentType,
			t.requestInfo(method, request.URL, request, requestsSent, t.retry.MaxAttempts, started),
			ErrorTypeNone)
		if buildErr != nil {
			return nil, buildErr // unreachable: ClassifyErrorType only yields members
		}

		response.present(raw.body, raw.contentType, request.ResponseFormat, request.DTOShape)
		return response, nil
	}

	// Canceled is tested first: it is the caller withdrawing, where DeadlineExceeded is our own budget.
	verb := "failed"
	errorType := ErrorTypeNetwork
	switch {
	case errors.Is(lastError, context.Canceled):
		verb = "cancelled"
		errorType = ErrorTypeUnknown
	case isTimeoutError(lastError):
		verb = "timed out"
		errorType = ErrorTypeTimeout
	}
	// The failure's type, not its text, which names the host and port; the text rides CauseText.
	failureType := ""
	causeText := ""
	if lastError != nil {
		failureType = fmt.Sprintf("%T", lastError)
		causeText = redact.RedactText(lastError.Error())
	}
	info := t.requestInfo(method, request.URL, request, requestsSent, t.retry.MaxAttempts, started)
	return nil, &exerrors.SdkError{
		// What went wrong, and nothing about which call it was: Error() renders that off the record.
		Message:     fmt.Sprintf("%s after %d attempt(s): %s", verb, requestsSent, failureType),
		ErrorType:   errorType,
		RequestInfo: info,
		Cause:       lastError,
		CauseText:   causeText,
	}
}

// refuseIfBudgetSpent folds a request's preparation into the operation deadline the caller set.
func refuseIfBudgetSpent(ctx context.Context, when string) error {
	cause := ctx.Err()
	if cause == nil {
		return nil
	}
	verb, errorType := "cancelled", ErrorTypeUnknown
	if errors.Is(cause, context.DeadlineExceeded) {
		verb, errorType = "timed out", ErrorTypeTimeout
	}
	return &exerrors.SdkError{
		Message:   fmt.Sprintf("%s %s, so nothing was sent", verb, when),
		ErrorType: errorType,
		Cause:     cause,
	}
}

// Config returns a deep copy of the configuration this transport was built with.
func (t *HTTPTransport) Config() excfg.SdkConfig { return t.config.DeepCopy() }

// MaxAttempts is the attempt ceiling this transport's retry policy grants one request.
func (t *HTTPTransport) MaxAttempts() int { return t.retry.MaxAttempts }

// Close releases pooled connections.
func (t *HTTPTransport) Close() {
	t.discardPooledConnections()
}

// CloseTransport releases pooled connections - an alias for Close.
func (t *HTTPTransport) CloseTransport() {
	t.Close()
}

// Verify reports the effective TLS trust decision: the CA bundle path if configured, else verifySsl.
func (t *HTTPTransport) Verify() any {
	if t.config.CABundle != "" {
		return t.config.CABundle
	}
	return t.config.VerifySslEnabled()
}

func (t *HTTPTransport) discardPooledConnections() {
	if tr, ok := t.client.Transport.(*http.Transport); ok {
		tr.CloseIdleConnections()
	}
}

func (t *HTTPTransport) doOnce(ctx context.Context, method, fullURL string, request TransportRequest) (*rawResponse, error) {
	requestContext, cancel := context.WithTimeout(ctx, durationSeconds(t.config.TimeoutSeconds))
	defer cancel()

	httpRequest, err := buildRequest(requestContext, method, fullURL, request)
	if err != nil {
		return nil, err
	}
	httpResponse, err := t.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	body, err := readCappedBody(httpResponse.Body, httpResponse.ContentLength, t.config.MaxResponseBytes)
	if err != nil {
		return nil, err
	}
	return &rawResponse{
		status:      httpResponse.StatusCode,
		headers:     flattenHeaders(httpResponse.Header),
		contentType: httpResponse.Header.Get("Content-Type"),
		retryAfter:  httpResponse.Header.Get("Retry-After"),
		body:        body,
	}, nil
}

// errResponseTooLarge marks the one doOnce failure that is not a transport failure, so it is not retried.
var errResponseTooLarge = errors.New("exsited: response body over the buffered-response cap")

// readCappedBody buffers a response body up to limitBytes, refusing a longer one.
func readCappedBody(source io.Reader, declaredBytes int64, limitBytes int) ([]byte, error) {
	if declaredBytes > int64(limitBytes) {
		return nil, refuseOversizedBody(limitBytes, declaredBytes)
	}
	body, err := io.ReadAll(io.LimitReader(source, int64(limitBytes)+1))
	if err != nil {
		return nil, err
	}
	if len(body) > limitBytes {
		return nil, refuseOversizedBody(limitBytes, declaredBytes)
	}
	return body, nil
}

// refuseOversizedBody wraps errResponseTooLarge so the attempt loop leaves without retrying.
func refuseOversizedBody(limitBytes int, declaredBytes int64) error {
	declared := ""
	if declaredBytes >= 0 {
		declared = fmt.Sprintf(" (Content-Length %d)", declaredBytes)
	}
	return fmt.Errorf("%w: the response body is larger than the %d-byte cap this client "+
		"buffers%s; raise SdkConfig.MaxResponseBytes to accept it",
		errResponseTooLarge, limitBytes, declared)
}

func (t *HTTPTransport) delay(ctx context.Context, seconds float64) error {
	if seconds <= 0 {
		return nil
	}
	timer := time.NewTimer(durationSeconds(seconds))
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func (t *HTTPTransport) requestInfo(method, url string, request TransportRequest, attempt, maxAttempts int, started time.Time) *RequestInfo {
	// Rounded where it is measured, so every call site prints the same digits.
	elapsed := exerrors.RoundSeconds(time.Since(started).Seconds())
	info := NewRequestInfo(t.config.Auth.ClientName, string(t.config.APIVersion), method, url, request.QueryParams,
		request.RequestInfoBody, request.RequestInfoHeaders, attempt, maxAttempts, elapsed)
	// Set directly: a template holds no values, so the redacting constructor has nothing to do.
	info.EndpointTemplate = request.EndpointTemplate
	return info
}

// NormaliseMethod upper-cases and trims, and refuses anything outside the RFC 9110 token set.
func NormaliseMethod(httpMethod string) (string, error) {
	if strings.TrimSpace(httpMethod) == "" {
		return "", exerrors.NewUsageError(fmt.Sprintf(
			"http_method: expected a non-blank string like 'GET', got '%s'", httpMethod))
	}
	method := strings.ToUpper(strings.TrimSpace(httpMethod))
	if !excfg.IsHTTPToken(method) {
		// Neutralised and redacted, not echoed: a refused verb may carry CR/LF or control characters.
		return "", exerrors.NewUsageError(redact.RedactText(fmt.Sprintf(
			"httpMethod must be a single verb with no spaces or control characters, got '%s'",
			redact.NeutraliseControls(method))))
	}
	return method, nil
}

func buildTransportURL(rawURL string, queryParams map[string]any) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		// Described, not echoed: an unparseable URL may carry userinfo that key masking cannot reach.
		refusal := exerrors.NewUsageError(fmt.Sprintf(
			"url: must be an absolute http:// or https:// URL, got %s", redact.Describe(rawURL)))
		refusal.CauseText = redact.RedactText(rawURL)
		return "", refusal
	}
	// Ahead of the early return: a fragment is dropped on the wire whether or not there is a query.
	if err := rejectAFragment("url", rawURL); err != nil {
		return "", err
	}
	if len(queryParams) == 0 {
		return rawURL, nil
	}
	keys := make([]string, 0, len(queryParams))
	for k := range queryParams {
		keys = append(keys, k)
	}
	sort.Strings(keys) // stable order (query order is not significant to the server)

	var b strings.Builder
	b.WriteString(rawURL)
	first := !strings.Contains(rawURL, "?")
	for _, key := range keys {
		for _, single := range ExpandValue(queryParams[key]) {
			if first {
				b.WriteByte('?')
			} else {
				b.WriteByte('&')
			}
			b.WriteString(url.QueryEscape(key))
			b.WriteByte('=')
			b.WriteString(url.QueryEscape(single))
			first = false
		}
	}
	return b.String(), nil
}

// ExpandValue renders one query value as the strings the wire carries, a slice as one per element.
func ExpandValue(value any) []string {
	switch v := value.(type) {
	case string:
		return []string{v}
	case []any:
		out := make([]string, 0, len(v))
		for _, item := range v {
			out = append(out, stringifyValue(item))
		}
		return out
	default:
		return []string{stringifyValue(value)}
	}
}

func stringifyValue(value any) string {
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}

// CheckRequestHeaders refuses, before any I/O, a header the wire cannot carry.
func CheckRequestHeaders(headers map[string]string) error {
	names := make([]string, 0, len(headers))
	for name := range headers {
		names = append(names, name)
	}
	sort.Strings(names) // a Go map iterates randomly; the header named must not change per run
	for _, name := range names {
		problem := excfg.HeaderNameProblem(name)
		if problem == "" {
			problem = excfg.HeaderValueProblem(headers[name])
		}
		if problem != "" {
			return exerrors.NewUsageError(fmt.Sprintf(
				"header '%s': not something the transport can send - %s. The value is not "+
					"shown: it may be a credential", redact.NeutraliseControls(name), problem))
		}
	}
	return nil
}

func buildRequest(ctx context.Context, method, fullURL string, request TransportRequest) (*http.Request, error) {
	assembled := request.wire
	if assembled == nil {
		// Defence only: every send path materialises first, where the body is assembled and capped.
		built, err := assembleWireBody(request, nil)
		if err != nil {
			return nil, err
		}
		assembled = built
	}

	var body io.Reader
	if assembled != nil {
		// A fresh reader over the same bytes, starting at their start on every attempt.
		body = bytes.NewReader(assembled.body)
	}
	httpRequest, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		// net/http's refusal names the whole URL it was handed, so the text rides CauseText.
		refusal := exerrors.NewUsageError(
			"the request could not be built for this method and URL")
		refusal.Cause = err
		refusal.CauseText = redact.RedactText(err.Error())
		return nil, refusal
	}
	applyHeaders(httpRequest, request.Headers)
	switch {
	case assembled == nil:
		// No body, so nothing to label.
	case assembled.kind == wireMultipart:
		// Unconditional: the boundary is this SDK's, so a caller's Content-Type is overwritten, not rejected.
		httpRequest.Header.Set(contentTypeHeader, assembled.contentType)
	case httpRequest.Header.Get(contentTypeHeader) == "":
		// Set only when the caller has not, so a vendor type survives unchecked.
		httpRequest.Header.Set(contentTypeHeader, assembled.contentType)
	}
	return httpRequest, nil
}

// MaxMultipartBytes caps the multipart body this SDK assembles in memory: 64 MiB, judged before any copy.
const MaxMultipartBytes = 64 << 20

// The boundary's shape is fixed at 52 bytes, so the framing arithmetic knows its length in advance.
const (
	multipartBoundaryPrefix = "exsitedGoSdkBoundary"
	multipartBoundaryDigits = 32
	multipartBoundaryLength = len(multipartBoundaryPrefix) + multipartBoundaryDigits
)

// multipartPart is one resolved part of a multipart body, ready to write and already judged.
type multipartPart struct {
	// what names the part in a refusal - "files part 'file'", "formData field 'name'".
	what    string
	header  textproto.MIMEHeader
	content []byte
}

// assembleMultipartBody renders Files and FormData into the exact bytes of a multipart body, once.
func assembleMultipartBody(request TransportRequest) (*wireBody, error) {
	parts, err := resolveMultipartParts(request)
	if err != nil {
		return nil, err
	}
	predicted, err := refuseOversizedMultipart(parts)
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	// One allocation of exactly the document's size, where bytes.Buffer would double its way there.
	buffer.Grow(predicted)
	writer := multipart.NewWriter(&buffer)
	boundary, err := multipartBoundary(parts)
	if err != nil {
		return nil, err
	}
	if err := writer.SetBoundary(boundary); err != nil {
		// Unreachable: the derived boundary is 52 bytes of letters and digits.
		return nil, exerrors.NewUsageError("files: could not set the multipart boundary: " + err.Error())
	}
	for _, part := range parts {
		written, err := writer.CreatePart(part.header)
		if err != nil {
			return nil, exerrors.NewUsageError(part.what + ": could not create part: " + err.Error())
		}
		if _, err := written.Write(part.content); err != nil {
			return nil, exerrors.NewUsageError(part.what + ": could not be written: " + err.Error())
		}
	}
	if err := writer.Close(); err != nil {
		return nil, exerrors.NewUsageError("files: could not finalise multipart body: " + err.Error())
	}
	// The produced document, not the predicted one, so the ceiling holds for the bytes on the wire.
	if err := refuseOversizedWireBody("files", "the multipart body", buffer.Len()); err != nil {
		return nil, err
	}
	return &wireBody{kind: wireMultipart, contentType: writer.FormDataContentType(),
		body: buffer.Bytes()}, nil
}

// resolveMultipartParts judges every part in write order: files, then formData, each sorted by name.
func resolveMultipartParts(request TransportRequest) ([]multipartPart, error) {
	parts := make([]multipartPart, 0, len(request.Files)+len(request.FormData))
	for _, field := range sortedMapKeys(request.Files) {
		filename, content, contentType, err := resolveFilePart(field, request.Files[field])
		if err != nil {
			return nil, err
		}
		// CreatePart rather than CreateFormFile, which hardcodes application/octet-stream.
		header := textproto.MIMEHeader{}
		header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeMultipartParam(field), escapeMultipartParam(filename)))
		if contentType != "" {
			header.Set(contentTypeHeader, contentType)
		}
		parts = append(parts, multipartPart{
			what:    fmt.Sprintf("files: '%s'", field),
			header:  header,
			content: content,
		})
	}
	// Not writer.WriteField: its quote escaper writes CR and LF in a field name verbatim.
	for _, name := range sortedMapKeys(request.FormData) {
		if err := refuseMultipartHeaderBytes(formDataLabel(name), name); err != nil {
			return nil, err
		}
		header := textproto.MIMEHeader{}
		header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"`,
			escapeMultipartParam(name)))
		parts = append(parts, multipartPart{
			what:    "formData",
			header:  header,
			content: []byte(stringifyValue(request.FormData[name])),
		})
	}
	return parts, nil
}

// sortedMapKeys sorts, so a body is one document rather than one per range over the mapping.
func sortedMapKeys(source map[string]any) []string {
	keys := make([]string, 0, len(source))
	for key := range source {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// multipartBoundary derives the boundary from the parts, re-salting until it appears in none of them.
func multipartBoundary(parts []multipartPart) (string, error) {
	sum := multipartPartsDigest(parts)
	for salt := 0; salt < maxMultipartBoundarySalts; salt++ {
		candidate := multipartBoundaryCandidate(sum, salt)
		if !multipartPartsContain(parts, candidate) {
			return candidate, nil
		}
	}
	// Unreachable short of a preimage break: each candidate is 128 bits derived from the whole body.
	return "", exerrors.NewUsageError(fmt.Sprintf(
		"files: could not derive a multipart boundary absent from the body after %d attempts",
		maxMultipartBoundarySalts))
}

const maxMultipartBoundarySalts = 64

// multipartPartsDigest mixes lengths in, so two different bodies cannot stream the same bytes in.
func multipartPartsDigest(parts []multipartPart) []byte {
	digest := sha256.New()
	for _, part := range parts {
		for _, name := range sortedHeaderNames(part.header) {
			fmt.Fprintf(digest, "%s\x00%s\x00", name, strings.Join(part.header[name], "\x00"))
		}
		fmt.Fprintf(digest, "%d\x00", len(part.content))
		digest.Write(part.content)
	}
	return digest.Sum(nil)
}

// multipartBoundaryCandidate is the salt'th boundary for parts digesting to sum; salt 0 is the digest.
func multipartBoundaryCandidate(sum []byte, salt int) string {
	digits := hex.EncodeToString(sum)
	if salt > 0 {
		resalted := sha256.New()
		resalted.Write(sum)
		fmt.Fprintf(resalted, "\x00salt\x00%d", salt)
		digits = hex.EncodeToString(resalted.Sum(nil))
	}
	return multipartBoundaryPrefix + digits[:multipartBoundaryDigits]
}

// multipartPartsContain searches for the bare boundary, which a lenient receiver would also match.
func multipartPartsContain(parts []multipartPart, candidate string) bool {
	needle := []byte(candidate)
	for _, part := range parts {
		for name, values := range part.header {
			if strings.Contains(name, candidate) {
				return true
			}
			for _, value := range values {
				if strings.Contains(value, candidate) {
					return true
				}
			}
		}
		if bytes.Contains(part.content, needle) {
			return true
		}
	}
	return false
}

func sortedHeaderNames(header textproto.MIMEHeader) []string {
	names := make([]string, 0, len(header))
	for name := range header {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// multipartWireSize is the exact byte length assembleMultipartBody will write, per mime/multipart's framing.
func multipartWireSize(parts []multipartPart, limit int) (int, bool) {
	total := multipartBoundaryLength + 8 // the closing delimiter line
	for index, part := range parts {
		if index == 0 {
			total += multipartBoundaryLength + 4
		} else {
			total += multipartBoundaryLength + 6
		}
		for name, values := range part.header {
			for _, value := range values {
				total += len(name) + 2 + len(value) + 2 // "name: value" CRLF
			}
		}
		total += 2 // the blank line that ends the header block
		total += len(part.content)
		if total > limit {
			return 0, false
		}
	}
	return total, true
}

// refuseOversizedMultipart refuses a body over MaxMultipartBytes before a byte is copied.
func refuseOversizedMultipart(parts []multipartPart) (int, error) {
	size, within := multipartWireSize(parts, MaxMultipartBytes)
	if !within {
		return 0, exerrors.NewUsageError(fmt.Sprintf(
			"files: the multipart body is larger than the %d-byte cap this SDK assembles in "+
				"memory; send a body this size with a request of your own, streaming it from "+
				"disk rather than buffering it whole", MaxMultipartBytes))
	}
	return size, nil
}

// CheckFiles refuses, before any I/O and the token grant, files entries that cannot become parts.
func CheckFiles(files map[string]any) error {
	fields := make([]string, 0, len(files))
	for field := range files {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	for _, field := range fields {
		if _, _, _, err := resolveFilePart(field, files[field]); err != nil {
			return err
		}
	}
	return nil
}

// CheckFormDataNames refuses, before the token grant, a formData name that cannot label a part.
func CheckFormDataNames(formData map[string]any) error {
	names := make([]string, 0, len(formData))
	for name := range formData {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		if err := refuseMultipartHeaderBytes(formDataLabel(name), name); err != nil {
			return err
		}
	}
	return nil
}

// formDataLabel names a field by its length, not its text, which would carry the forged header.
func formDataLabel(name string) string {
	return fmt.Sprintf("a formData name of %d byte(s)", len(name))
}

func resolveFilePart(field string, value any) (filename string, content []byte, contentType string, err error) {
	resolved := false
	switch v := value.(type) {
	case []byte, string:
		if payload, ok := partContent(v); ok {
			filename, content, resolved = field, payload, true
		}
	case []any:
		// The triple lets the caller set the part's content type, as an .xlsx import needs.
		if len(v) > 3 {
			return "", nil, "", exerrors.NewUsageError(fmt.Sprintf(
				"files: '%s' has %d elements; the tuple form is [filename, contents] or "+
					"[filename, contents, contentType]", field, len(v)))
		}
		if len(v) >= 2 {
			if name, ok := v[0].(string); ok {
				if bytesValue, contentOK := partContent(v[1]); contentOK {
					filename, content, resolved = name, bytesValue, true
				}
			}
		}
		if len(v) == 3 && resolved {
			ct, ok := v[2].(string)
			if !ok {
				return "", nil, "", exerrors.NewUsageError(fmt.Sprintf(
					"files: '%s' third element must be the part's content type as a string, got %T",
					field, v[2]))
			}
			contentType = ct
		}
	}
	if !resolved {
		return "", nil, "", exerrors.NewUsageError(fmt.Sprintf(
			"files: '%s' must be the contents - []byte or a string of text - or [filename, contents], "+
				"or [filename, contents, contentType]. Text is the file's CONTENTS, never a path: "+
				"nothing here opens a file. Name the file so the receiver can parse it: "+
				"[]any{\"notes.txt\", text, \"text/plain\"}.", field))
	}
	// All three reach a part header, so all three are held to the same rule.
	if err := refuseMultipartHeaderBytes(fmt.Sprintf("a files part name of %d byte(s)", len(field)), field); err != nil {
		return "", nil, "", err
	}
	for _, held := range []struct{ what, value string }{
		{fmt.Sprintf("the filename on part '%s'", field), filename},
		{fmt.Sprintf("the content type on part '%s'", field), contentType},
	} {
		if err := refuseMultipartHeaderBytes(held.what, held.value); err != nil {
			return "", nil, "", err
		}
	}
	return filename, content, contentType, nil
}

func partContent(value any) ([]byte, bool) {
	switch v := value.(type) {
	case []byte:
		return v, v != nil
	case string:
		return []byte(v), true
	}
	return nil, false
}

// multipartParamEscaper percent-escapes the two bytes that can end a quoted parameter early.
var multipartParamEscaper = strings.NewReplacer(`"`, "%22", `\`, "%5C")

func escapeMultipartParam(value string) string {
	return multipartParamEscaper.Replace(value)
}

// refuseMultipartHeaderBytes refuses CR, LF or NUL in a part name, upload filename or content type.
func refuseMultipartHeaderBytes(what, value string) error {
	for index := 0; index < len(value); index++ {
		name := ""
		switch value[index] {
		case '\r':
			name = "a carriage return"
		case '\n':
			name = "a newline"
		case 0x00:
			name = "a NUL"
		}
		if name != "" {
			return exerrors.NewUsageError(fmt.Sprintf(
				"%s carries %s at byte %d, which would forge a header on the multipart part this "+
					"SDK sends. The value is not shown: it may be attacker-chosen. Strip CR, LF and "+
					"NUL from an upload filename before passing it on", what, name, index))
		}
	}
	return nil
}

func applyHeaders(httpRequest *http.Request, headers map[string]string) {
	for k, v := range headers {
		httpRequest.Header.Set(k, v)
	}
}

func flattenHeaders(header http.Header) map[string]string {
	out := make(map[string]string, len(header))
	for k, v := range header {
		if len(v) > 0 {
			// Joined, never dropped, so a second Set-Cookie survives.
			out[k] = strings.Join(v, ", ")
		}
	}
	return out
}

func durationSeconds(seconds float64) time.Duration {
	return time.Duration(seconds * float64(time.Second))
}

// IdempotencyHeader is the header v4 uses to collapse a repeated create into one.
const IdempotencyHeader = "X-Idempotency-Key"

// idempotencyKeyedMethods are the methods the SDK generates a key for: POST, the one not already idempotent.
var idempotencyKeyedMethods = map[string]bool{http.MethodPost: true}

// WithIdempotencyKey copies the headers and adds a random key when one is due; a caller's own key always wins.
func WithIdempotencyKey(method string, headers map[string]string) map[string]string {
	merged := make(map[string]string, len(headers)+1)
	for name, value := range headers {
		merged[name] = value
	}
	if !idempotencyKeyedMethods[strings.ToUpper(strings.TrimSpace(method))] {
		return merged
	}
	for name := range merged {
		if strings.EqualFold(name, IdempotencyHeader) {
			return merged
		}
	}
	buffer := make([]byte, 16)
	if _, err := rand.Read(buffer); err != nil {
		// Without a random key a replay would collapse unrelated requests, which is worse than none.
		return merged
	}
	merged[IdempotencyHeader] = hex.EncodeToString(buffer)
	return merged
}

// tokenFailureMarkers is what a 401 body says when the access token, not the request, is the problem.
var tokenFailureMarkers = map[string]bool{
	"unauthorized":         true,
	"invalid_access_token": true,
	"invalid_token":        true,
	"expired_token":        true,
	"invalid_request":      true,
}

// tokenFailureKeys are read independently: a first-non-nil chain never reaches "code", where v4 puts the word.
var tokenFailureKeys = []string{"code", "error", "message"}

// MeansExpiredToken reports whether a 401 body is a stale access token rather than a business rejection.
//
// A body that is not a JSON object says nothing about the token, so it is guessed dead only for a safe method.
func MeansExpiredToken(body any, method string) bool {
	record, ok := body.(map[string]any)
	if !ok {
		return IsSafeMethod(method)
	}
	for _, key := range tokenFailureKeys {
		if namesATokenFailure(record[key]) {
			return true
		}
	}
	return false
}

// namesATokenFailure matches the whole trimmed value or any word of it, never a substring.
func namesATokenFailure(value any) bool {
	text, ok := value.(string)
	if !ok {
		return false
	}
	text = strings.ToLower(strings.TrimSpace(text))
	if text == "" {
		return false
	}
	if tokenFailureMarkers[text] {
		return true
	}
	for _, word := range strings.Fields(text) {
		if tokenFailureMarkers[word] {
			return true
		}
	}
	return false
}

// RawExchange is one completed exchange as the token provider reads it.
type RawExchange struct {
	Status      int
	Headers     map[string]string
	ContentType string
	Text        string
	Decoded     any
	Attempts    int
	// Elapsed is in milliseconds.
	Elapsed int64
	// Body is the body as it arrived, before decoding.
	Body []byte

	response *ApiResponse
}

// SendOnce sends one request through SendHTTPRequest, with the SDK's own headers and an optional bearer token.
func (t *HTTPTransport) SendOnce(
	ctx context.Context, method, absolute string, queryParams map[string]any,
	body any, formData map[string]any, headers map[string]string, token string,
) (*RawExchange, error) {
	merged := mergeHeaders(t.config, headers)
	if token != "" {
		merged["Authorization"] = "Bearer " + token
	}
	reported := body
	if reported == nil && formData != nil {
		reported = formData
	}
	// Assembled before the record is redacted, so an oversized body is refused before it is scanned.
	request, err := TransportRequest{
		HTTPMethod:  method,
		URL:         absolute,
		QueryParams: queryParams,
		JSONBody:    body,
		FormData:    formData,
		Headers:     merged,
	}.Materialised()
	if err != nil {
		return nil, err
	}
	request.RequestInfoHeaders = redact.RedactMapping(merged)
	request.RequestInfoBody = redact.RedactValue(reported)
	response, err := t.SendHTTPRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	return exchangeOf(response), nil
}

// mergeHeaders applies the SDK's Accept and User-Agent, then DefaultHeaders, then the caller's, case-insensitively.
func mergeHeaders(config excfg.SdkConfig, headers map[string]string) map[string]string {
	canonical := map[string]string{}
	values := map[string]string{}
	put := func(name, value string) {
		lower := strings.ToLower(name)
		canonical[lower] = name
		values[lower] = value
	}
	put("Accept", "application/json")
	put("User-Agent", config.UserAgent)
	for _, name := range sortedHeaderKeys(config.DefaultHeaders) {
		put(name, config.DefaultHeaders[name])
	}
	for _, name := range sortedHeaderKeys(headers) {
		put(name, headers[name])
	}
	out := make(map[string]string, len(values))
	for lower, value := range values {
		out[canonical[lower]] = value
	}
	return out
}

func sortedHeaderKeys(headers map[string]string) []string {
	names := make([]string, 0, len(headers))
	for name := range headers {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func exchangeOf(response *ApiResponse) *RawExchange {
	attempts, elapsed := 0, int64(0)
	if response.RequestInfo != nil {
		attempts = response.RequestInfo.AttemptCount
		elapsed = int64(response.RequestInfo.ElapsedSeconds * 1000)
	}
	return &RawExchange{
		Status:      response.StatusCode,
		Headers:     response.Headers,
		ContentType: response.ContentType,
		Text:        response.JSON(),
		Decoded:     response.Map,
		Attempts:    attempts,
		Elapsed:     elapsed,
		Body:        response.rawBody,
		response:    response,
	}
}

// NewApiResponseFromExchange is the response an exchange came from, re-presented under format and shape.
func NewApiResponseFromExchange(
	exchange *RawExchange, info *RequestInfo, format string, shape ...DTOShape,
) *ApiResponse {
	response := exchange.response
	if response == nil {
		response = &ApiResponse{
			StatusCode:  exchange.Status,
			Headers:     exchange.Headers,
			ContentType: exchange.ContentType,
			Map:         exchange.Decoded,
		}
		if response.Headers == nil {
			response.Headers = map[string]string{}
		}
		response.ErrorType = classifyErrorType(exchange.Status, response.BodyDeniesSuccess())
	}
	if info != nil {
		response.RequestInfo = info
	}
	var declared DTOShape
	if len(shape) > 0 {
		declared = shape[0]
	}
	response.present(exchange.Body, exchange.ContentType, format, declared)
	return response
}
