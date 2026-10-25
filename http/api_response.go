package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/exsited/exsited-sdk-go/v4/config"
	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	"github.com/exsited/exsited-sdk-go/v4/redact"
)

const (
	maxReportedMessages = 50
	maxMessageDepth     = 4
)

// DTOShape builds the typed view of one response body, or nil where there is no typed form.
type DTOShape func(decoded any) any

type ApiResponse struct {
	StatusCode int
	// Data is the body under the resolved response format, live and never a copy; nil when there is no such form.
	Data        any
	Headers     map[string]string
	ContentType string
	RequestInfo *RequestInfo
	ErrorType   ErrorType

	// Map is always the decoded body, whatever the response format says, and not a copy: a map or a
	// list for JSON, a string for text, []byte for anything else, nil for a 204.
	Map any

	// Unexported, so an ApiResponse stays buildable as a literal; each view falls back when unset.
	rawBody    []byte   // the body as it arrived: not a copy, the slice the transport read
	bodyIsText bool     // whether the content type made those bytes text
	dtoShape   DTOShape // the typed view this endpoint produces, or nil

	// views guards the five fields below and nothing else.
	views sync.Mutex
	// jsonHasText is whether the body has a text form at all, which keeps Data nil rather than "" under JSON.
	jsonBuilt   bool
	jsonText    string
	jsonHasText bool
	dtoBuilt    bool
	dtoView     any
}

// JSON is always the server's own body text, the bytes as they arrived, never re-serialised.
func (r *ApiResponse) JSON() string {
	text, _ := r.jsonView()
	return text
}

// jsonView is JSON() plus whether this body has a text form at all.
func (r *ApiResponse) jsonView() (text string, hasText bool) {
	r.views.Lock()
	defer r.views.Unlock()
	if !r.jsonBuilt {
		r.jsonText, r.jsonHasText = r.buildJSON()
		r.jsonBuilt = true
	}
	return r.jsonText, r.jsonHasText
}

// buildJSON renders the body text once; called under the views lock.
func (r *ApiResponse) buildJSON() (text string, hasText bool) {
	if text, isText := r.Map.(string); isText {
		return text, true
	}
	if !r.bodyIsText || len(r.rawBody) == 0 {
		return "", false
	}
	return string(r.rawBody), true
}

// DTO is the typed view of this body, built once from the shape the endpoint declared.
//
// nil on a call that failed, the v4 error envelope included, because an error body is not a record.
func (r *ApiResponse) DTO() any {
	r.views.Lock()
	defer r.views.Unlock()
	if !r.dtoBuilt {
		if r.dtoShape != nil && r.Succeeded() {
			r.dtoView = r.dtoShape(r.Map)
		}
		r.dtoBuilt = true
	}
	return r.dtoView
}

// present records the body this response was decoded from and moves Data to the resolved format.
func (r *ApiResponse) present(rawBody []byte, contentType, responseFormat string, shape DTOShape) {
	r.rawBody = rawBody
	r.bodyIsText = bodyIsTextual(contentType)
	r.dtoShape = shape
	r.jsonBuilt, r.jsonText, r.jsonHasText = false, "", false
	r.dtoBuilt, r.dtoView = false, nil
	r.Data = r.Map
	switch responseFormat {
	case config.ResponseFormatJSON:
		// nil, not "", for a body with no text form, which "" would misreport as an empty body.
		if text, hasText := r.jsonView(); hasText {
			r.Data = text
		} else {
			r.Data = nil
		}
	case config.ResponseFormatDTO:
		r.Data = r.DTO()
	}
}

// bodyIsTextual applies DecodeResponseBody's two rules, so JSON() and Map agree on text bodies.
func bodyIsTextual(contentType string) bool {
	lowered := strings.ToLower(strings.TrimSpace(contentType))
	return isJSONMediaType(lowered) || strings.HasPrefix(lowered, "text/")
}

const jsonMediaType = "application/json"

// isJSONMediaType reports whether a lower-cased Content-Type names JSON, "+json" suffixes included.
func isJSONMediaType(loweredContentType string) bool {
	mediaType := loweredContentType
	if semicolon := strings.IndexByte(mediaType, ';'); semicolon >= 0 {
		mediaType = mediaType[:semicolon]
	}
	mediaType = strings.TrimSpace(mediaType)
	return mediaType == jsonMediaType || strings.HasSuffix(mediaType, "+json")
}

// NewApiResponse builds a response, deriving the classification from the status when none is given.
func NewApiResponse(statusCode int, data any, headers map[string]string, contentType string,
	requestInfo *RequestInfo, errorType ErrorType) (*ApiResponse, error) {
	if !knownErrorType(errorType) {
		return nil, exerrors.NewUsageError(fmt.Sprintf(
			"unknown error_type '%s'; expected an ErrorType or one of: %s",
			errorType, strings.Join(errorTypeMembers, ", ")))
	}
	if headers == nil {
		headers = map[string]string{}
	}
	denied := bodyDeniesSuccess(data)
	if errorType == ErrorTypeNone && (statusCode >= 300 || denied) {
		errorType = classifyErrorType(statusCode, denied)
	}
	return &ApiResponse{
		StatusCode:  statusCode,
		Data:        data,
		Map:         data,
		Headers:     headers,
		ContentType: contentType,
		RequestInfo: requestInfo,
		ErrorType:   errorType,
	}, nil
}

// Succeeded reports whether the status is 2xx and the body agrees.
func (r *ApiResponse) Succeeded() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300 && !bodyDeniesSuccess(r.Map)
}

// BodyDeniesSuccess reports whether the body names itself an error, as every v4 error body does.
func (r *ApiResponse) BodyDeniesSuccess() bool { return bodyDeniesSuccess(r.Map) }

// bodyDeniesSuccess keys on the v4 "object" member: a 2xx whose body says "error" is a failure.
func bodyDeniesSuccess(data any) bool {
	body, isMapping := data.(map[string]any)
	if !isMapping {
		return false
	}
	name, isText := body["object"].(string)
	return isText && strings.EqualFold(strings.TrimSpace(name), "error")
}

// GetErrorMessages reads the error envelope, most specific first, capped at maxReportedMessages.
func (r *ApiResponse) GetErrorMessages() []string {
	var body any = r.Map
	if _, ok := body.(map[string]any); !ok {
		body = parseJSONBody(body)
	}
	m, ok := body.(map[string]any)
	if !ok {
		return nil
	}
	messages := readErrorEntries(m["errors"])
	if len(messages) == 0 {
		messages = readMessageValue(m["message"], 0)
	}
	if len(messages) == 0 {
		messages = readMessageValue(m["error"], 0)
		if len(messages) == 0 {
			// The v4 envelope's own machine-readable tier.
			messages = readMessageValue(m["code"], 0)
		}
		if len(messages) == 0 {
			messages = readMessageValue(m["errorCode"], 0)
		}
	}
	if len(messages) == 0 {
		// RFC 7807, reached only when an intermediary answers the error rather than the API.
		messages = readMessageValue(m["detail"], 0)
		if len(messages) == 0 {
			messages = readMessageValue(m["title"], 0)
		}
	}
	return capMessages(messages)
}

// ErrorMessages is an alias of GetErrorMessages.
func (r *ApiResponse) ErrorMessages() []string { return r.GetErrorMessages() }

// RaiseForStatus returns an *SdkError when the call failed, with body and headers redacted, else nil.
func (r *ApiResponse) RaiseForStatus() error {
	if r.Succeeded() {
		return nil
	}
	messages := r.GetErrorMessages()
	sdkError := exerrors.NewResponseError(classificationSummary(r.ErrorType, len(messages)),
		r.StatusCode, r.ErrorType,
		redact.RedactValue(r.Map), redact.RedactMapping(r.Headers), r.RequestInfo, messages)
	sdkError.ErrorCode = r.ErrorCode()
	sdkError.RequestID = r.RequestID()
	sdkError.DocURL = r.DocURL()
	sdkError.FieldErrors = r.FieldErrors()
	return sdkError
}

// Err is an alias of RaiseForStatus.
func (r *ApiResponse) Err() error { return r.RaiseForStatus() }

// classificationSummary is the classification and the message count, none of the messages' text.
func classificationSummary(errorType ErrorType, reported int) string {
	name := errorType
	if name == "" || errorType == ErrorTypeNone {
		name = "error"
	}
	if reported == 0 {
		return name + " response"
	}
	return fmt.Sprintf("%s response (%d reported)", name, reported)
}

// ErrorCode is the API's machine-readable code, or "" when the body has none.
func (r *ApiResponse) ErrorCode() string {
	if code := r.errorField("code"); code != "" {
		return code
	}
	return r.errorField("error")
}

// RequestID is the API's own id for this request.
func (r *ApiResponse) RequestID() string { return r.errorField("request_id") }

// DocURL is documentation for ErrorCode.
func (r *ApiResponse) DocURL() string { return r.errorField("doc_url") }

// FieldErrors are the per-field validation failures, structured; only entries naming a field.
func (r *ApiResponse) FieldErrors() []exerrors.FieldError {
	body, ok := r.Map.(map[string]any)
	if !ok {
		return nil
	}
	entries, ok := body["errors"].([]any)
	if !ok {
		return nil
	}
	var out []exerrors.FieldError
	for _, entry := range entries {
		record, isMap := entry.(map[string]any)
		if !isMap {
			continue
		}
		field, _ := record["field"].(string)
		if strings.TrimSpace(field) == "" {
			continue
		}
		code, _ := record["code"].(string)
		message, _ := record["message"].(string)
		docURL, _ := record["doc_url"].(string)
		out = append(out, exerrors.FieldError{Field: field, Code: code, Message: message, DocURL: docURL})
	}
	return out
}

// errorField reads one top-level string off a failed response, so a record's own "code" never reads as one.
func (r *ApiResponse) errorField(key string) string {
	if r.Succeeded() {
		return ""
	}
	body, ok := r.Map.(map[string]any)
	if !ok {
		return ""
	}
	value, _ := body[key].(string)
	if strings.TrimSpace(value) == "" {
		return ""
	}
	return value
}

func parseJSONBody(body any) any {
	var text string
	switch b := body.(type) {
	case []byte:
		text = string(b)
	case string:
		text = b
	default:
		return nil
	}
	// Strip, unprefix, strip again: "\uFEFF  {...}" would otherwise fail the '{' guard.
	text = strings.TrimSpace(text)
	text = strings.TrimLeft(text, "\uFEFF")
	text = strings.TrimSpace(text)
	if !strings.HasPrefix(text, "{") && !strings.HasPrefix(text, "[") {
		return nil
	}
	if v, ok := DecodeJSON([]byte(text)); ok {
		if m, ok := v.(map[string]any); ok {
			return m
		}
	}
	return nil
}

func readErrorEntries(entries any) []string {
	var messages []string
	switch e := entries.(type) {
	case map[string]any:
		// Sorted: a Go map ranges randomly and SdkError.Message must be stable across runs.
		fields := make([]string, 0, len(e))
		for field := range e {
			fields = append(fields, field)
		}
		sort.Strings(fields)
		for _, field := range fields {
			prefixField(field, readMessageValue(e[field], 0), &messages)
		}
		return messages
	case []any:
		for _, entry := range e {
			if m, ok := entry.(map[string]any); ok {
				text := m["message"]
				if text == nil {
					text = m["code"]
				}
				prefixField(fieldName(m["field"]), readMessageValue(text, 0), &messages)
			} else {
				messages = append(messages, readMessageValue(entry, 0)...)
			}
		}
		return messages
	case string:
		return readMessageValue(e, 0)
	}
	return messages
}

func readMessageValue(value any, depth int) []string {
	if value == nil || depth > maxMessageDepth {
		return nil
	}
	switch v := value.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			return nil
		}
		return []string{v}
	case map[string]any:
		text := v["message"]
		if text == nil {
			text = v["code"]
		}
		return readMessageValue(text, depth+1)
	case []any:
		var messages []string
		for _, item := range v {
			messages = append(messages, readMessageValue(item, depth+1)...)
		}
		return messages
	default:
		text := fmt.Sprint(v)
		if strings.TrimSpace(text) == "" {
			return nil
		}
		return []string{text}
	}
}

// prefixField tests the field for presence, not truthiness: a bulk import numbers rows from row 0.
func prefixField(field string, fieldMessages []string, out *[]string) {
	// Blankness is judged trimmed, but the prefix keeps the API's own spelling.
	blank := strings.TrimSpace(field) == ""
	for _, m := range fieldMessages {
		if blank {
			*out = append(*out, m)
		} else {
			*out = append(*out, field+": "+m)
		}
	}
}

func fieldName(v any) string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v)
}

func capMessages(messages []string) []string {
	if len(messages) <= maxReportedMessages {
		return messages
	}
	capped := append([]string{}, messages[:maxReportedMessages]...)
	return append(capped, fmt.Sprintf("+%d more problems", len(messages)-maxReportedMessages))
}

// DecodeResponseBody decodes JSON to maps and lists, text/* to a string, else bytes; 204 or empty is nil.
func DecodeResponseBody(status int, body []byte, contentType string) any {
	if status == 204 || len(body) == 0 {
		return nil
	}
	lowered := strings.ToLower(strings.TrimSpace(contentType))
	if isJSONMediaType(lowered) {
		if v, ok := DecodeJSON(body); ok {
			return v
		}
		return string(body)
	}
	if strings.HasPrefix(lowered, "text/") {
		return decodeTextBody(body, lowered)
	}
	return body
}

// Tested as a substring, so a malformed parameter list cannot make body decoding fail.
const charsetParameter = "charset="

// decodeTextBody tries the declared charset, then UTF-8, then ISO-8859-1, which cannot fail.
func decodeTextBody(body []byte, loweredContentType string) string {
	if strings.Contains(loweredContentType, charsetParameter) {
		if decode := resolveCodec(declaredCharset(loweredContentType)); decode != nil {
			return decode(body)
		}
	}
	if utf8.Valid(body) {
		return string(body)
	}
	return decodeLatin1(body)
}

// declaredCharset reads the value after "charset="; lower-cased already by the caller.
func declaredCharset(loweredContentType string) string {
	value := loweredContentType[strings.Index(loweredContentType, charsetParameter)+len(charsetParameter):]
	if semicolon := strings.IndexByte(value, ';'); semicolon >= 0 {
		value = value[:semicolon]
	}
	return strings.Trim(strings.TrimSpace(value), `"'`)
}

// resolveCodec covers the charsets this package decodes without a dependency; anything else is nil.
func resolveCodec(name string) func([]byte) string {
	switch name {
	case "utf-8", "utf8":
		return decodeUTF8Replacing
	case "us-ascii", "ascii":
		return decodeASCIIReplacing
	case "iso-8859-1", "iso8859-1", "iso_8859-1", "latin-1", "latin1", "latin", "l1", "cp819":
		return decodeLatin1
	}
	return nil
}

func decodeUTF8Replacing(body []byte) string {
	if utf8.Valid(body) {
		return string(body)
	}
	return strings.ToValidUTF8(string(body), string(utf8.RuneError))
}

func decodeASCIIReplacing(body []byte) string {
	runes := make([]rune, len(body))
	for i, b := range body {
		if b < 0x80 {
			runes[i] = rune(b)
		} else {
			runes[i] = utf8.RuneError
		}
	}
	return string(runes)
}

// decodeLatin1 maps each byte to the code point of the same value; it cannot fail.
func decodeLatin1(body []byte) string {
	runes := make([]rune, len(body))
	for i, b := range body {
		runes[i] = rune(b)
	}
	return string(runes)
}

// DecodeJSON decodes exactly one JSON document, numbers kept as json.Number, and reports whether it parsed.
func DecodeJSON(body []byte) (any, bool) {
	dec := json.NewDecoder(bytes.NewReader(stripBOM(body)))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return nil, false
	}
	if _, err := dec.Token(); !errors.Is(err, io.EOF) {
		return nil, false
	}
	return v, true
}

func stripBOM(b []byte) []byte {
	if len(b) >= 3 && b[0] == 0xEF && b[1] == 0xBB && b[2] == 0xBF {
		return b[3:]
	}
	return b
}

// NumberAsInt64 reads a decoded JSON number that is exactly an integer, refusing a fraction.
func NumberAsInt64(raw any) (int64, bool) {
	switch typed := raw.(type) {
	case json.Number:
		if value, err := typed.Int64(); err == nil {
			return value, true
		}
		// Not integer syntax. 1.0 and 1e3 are still integral values, read through the float route.
		asFloat, err := typed.Float64()
		if err != nil || asFloat != float64(int64(asFloat)) {
			return 0, false
		}
		return int64(asFloat), true
	case float64:
		if typed != float64(int64(typed)) {
			return 0, false
		}
		return int64(typed), true
	}
	return 0, false
}

// NumberAsFloat64 reads a decoded JSON number as a float64; JSON has one number type, so nothing is lost.
func NumberAsFloat64(raw any) (float64, bool) {
	switch typed := raw.(type) {
	case json.Number:
		value, err := typed.Float64()
		return value, err == nil
	case float64:
		return typed, true
	}
	return 0, false
}
