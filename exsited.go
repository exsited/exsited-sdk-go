package exsited

import (
	"context"
	"iter"
	nethttp "net/http"

	"github.com/exsited/exsited-sdk-go/v4/auth"
	"github.com/exsited/exsited-sdk-go/v4/client"
	"github.com/exsited/exsited-sdk-go/v4/config"
	"github.com/exsited/exsited-sdk-go/v4/dto"
	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	"github.com/exsited/exsited-sdk-go/v4/http"
	"github.com/exsited/exsited-sdk-go/v4/redact"
	"github.com/exsited/exsited-sdk-go/v4/resources/abafiles"
	"github.com/exsited/exsited-sdk-go/v4/resources/accountemails"
	"github.com/exsited/exsited-sdk-go/v4/resources/accountgroups"
	"github.com/exsited/exsited-sdk-go/v4/resources/accountingcodes"
	"github.com/exsited/exsited-sdk-go/v4/resources/accounts"
	"github.com/exsited/exsited-sdk-go/v4/resources/actionscripts"
	"github.com/exsited/exsited-sdk-go/v4/resources/additionalcodes"
	"github.com/exsited/exsited-sdk-go/v4/resources/alertresponsibilities"
	"github.com/exsited/exsited-sdk-go/v4/resources/alertresponsibility"
	"github.com/exsited/exsited-sdk-go/v4/resources/alerts"
	"github.com/exsited/exsited-sdk-go/v4/resources/applications"
	"github.com/exsited/exsited-sdk-go/v4/resources/approvalprofiles"
	"github.com/exsited/exsited-sdk-go/v4/resources/approvals"
	"github.com/exsited/exsited-sdk-go/v4/resources/approvalusergroups"
	"github.com/exsited/exsited-sdk-go/v4/resources/attributeattachments"
	"github.com/exsited/exsited-sdk-go/v4/resources/billingperiods"
	"github.com/exsited/exsited-sdk-go/v4/resources/bookings"
	"github.com/exsited/exsited-sdk-go/v4/resources/chargingperiods"
	"github.com/exsited/exsited-sdk-go/v4/resources/communicationprofiles"
	"github.com/exsited/exsited-sdk-go/v4/resources/communications"
	"github.com/exsited/exsited-sdk-go/v4/resources/countries"
	"github.com/exsited/exsited-sdk-go/v4/resources/creditnoteapplications"
	"github.com/exsited/exsited-sdk-go/v4/resources/creditnotes"
	"github.com/exsited/exsited-sdk-go/v4/resources/customattributes"
	"github.com/exsited/exsited-sdk-go/v4/resources/customcomponents"
	"github.com/exsited/exsited-sdk-go/v4/resources/customevents"
	"github.com/exsited/exsited-sdk-go/v4/resources/customfeatures"
	"github.com/exsited/exsited-sdk-go/v4/resources/customfields"
	"github.com/exsited/exsited-sdk-go/v4/resources/customforms"
	"github.com/exsited/exsited-sdk-go/v4/resources/dayofmonth"
	"github.com/exsited/exsited-sdk-go/v4/resources/departments"
	"github.com/exsited/exsited-sdk-go/v4/resources/designations"
	"github.com/exsited/exsited-sdk-go/v4/resources/documenttemplates"
	"github.com/exsited/exsited-sdk-go/v4/resources/emailtemplates"
	"github.com/exsited/exsited-sdk-go/v4/resources/employees"
	"github.com/exsited/exsited-sdk-go/v4/resources/events"
	"github.com/exsited/exsited-sdk-go/v4/resources/express"
	"github.com/exsited/exsited-sdk-go/v4/resources/externaldatabases"
	"github.com/exsited/exsited-sdk-go/v4/resources/fulfillmentprofiles"
	"github.com/exsited/exsited-sdk-go/v4/resources/fulfillmentterms"
	"github.com/exsited/exsited-sdk-go/v4/resources/giftcertificates"
	"github.com/exsited/exsited-sdk-go/v4/resources/integrationconnections"
	"github.com/exsited/exsited-sdk-go/v4/resources/inventoryadjustments"
	"github.com/exsited/exsited-sdk-go/v4/resources/invoices"
	"github.com/exsited/exsited-sdk-go/v4/resources/invoiceterms"
	"github.com/exsited/exsited-sdk-go/v4/resources/itemfulfillments"
	"github.com/exsited/exsited-sdk-go/v4/resources/itemreceipts"
	"github.com/exsited/exsited-sdk-go/v4/resources/items"
	"github.com/exsited/exsited-sdk-go/v4/resources/labours"
	"github.com/exsited/exsited-sdk-go/v4/resources/milestones"
	"github.com/exsited/exsited-sdk-go/v4/resources/noteattachments"
	"github.com/exsited/exsited-sdk-go/v4/resources/notegroups"
	"github.com/exsited/exsited-sdk-go/v4/resources/orders"
	"github.com/exsited/exsited-sdk-go/v4/resources/organisation"
	"github.com/exsited/exsited-sdk-go/v4/resources/organisations"
	"github.com/exsited/exsited-sdk-go/v4/resources/payments"
	"github.com/exsited/exsited-sdk-go/v4/resources/portals"
	"github.com/exsited/exsited-sdk-go/v4/resources/pricingperiods"
	"github.com/exsited/exsited-sdk-go/v4/resources/proformainvoices"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchasecreditnoteapplications"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchasecreditnotes"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchaseexpress"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchaseinvoices"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchaseorders"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchasepayments"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchaserefunds"
	"github.com/exsited/exsited-sdk-go/v4/resources/quoteorders"
	"github.com/exsited/exsited-sdk-go/v4/resources/receiveterms"
	"github.com/exsited/exsited-sdk-go/v4/resources/reconciliations"
	"github.com/exsited/exsited-sdk-go/v4/resources/refunds"
	"github.com/exsited/exsited-sdk-go/v4/resources/rentalassets"
	"github.com/exsited/exsited-sdk-go/v4/resources/reports"
	"github.com/exsited/exsited-sdk-go/v4/resources/returnmerchandiseauthorisations"
	"github.com/exsited/exsited-sdk-go/v4/resources/returnvendorauthorizations"
	"github.com/exsited/exsited-sdk-go/v4/resources/roles"
	"github.com/exsited/exsited-sdk-go/v4/resources/salesreps"
	"github.com/exsited/exsited-sdk-go/v4/resources/salutations"
	"github.com/exsited/exsited-sdk-go/v4/resources/settings"
	"github.com/exsited/exsited-sdk-go/v4/resources/settingsworkflows"
	"github.com/exsited/exsited-sdk-go/v4/resources/shippingpackages"
	"github.com/exsited/exsited-sdk-go/v4/resources/shippingrules"
	"github.com/exsited/exsited-sdk-go/v4/resources/systemevents"
	"github.com/exsited/exsited-sdk-go/v4/resources/tasks"
	"github.com/exsited/exsited-sdk-go/v4/resources/teams"
	"github.com/exsited/exsited-sdk-go/v4/resources/timeperiods"
	"github.com/exsited/exsited-sdk-go/v4/resources/timesheets"
	"github.com/exsited/exsited-sdk-go/v4/resources/uom"
	"github.com/exsited/exsited-sdk-go/v4/resources/usages"
	"github.com/exsited/exsited-sdk-go/v4/resources/webhooks"
	"github.com/exsited/exsited-sdk-go/v4/resources/workflows"
	"github.com/exsited/exsited-sdk-go/v4/resources/zones"
)

func SDKVersion() string { return config.SDKVersion() }

// Version is SDKVersion under a shorter name: the same value.
func Version() string { return config.Version() }

// Ptr returns a pointer to v, for the optional/tri-state fields, e.g. exsited.Ptr(true).
func Ptr[T any](v T) *T { return &v }

type (
	AuthConfig = config.AuthConfig
	SdkConfig  = config.SdkConfig
	// APIVersion is the resource version a path renders with; this SDK serves v4 only.
	APIVersion = config.APIVersion
	// ResponseFormat is a response format name.
	ResponseFormat = config.ResponseFormat
)

// SelectCredentials picks one credentials block out of a list by position or by client name.
func SelectCredentials(credentials []any, instance any) (any, error) {
	return config.SelectCredentials(credentials, instance)
}

func ReadClientName(block any) string { return config.ReadClientName(block) }

// LabelCredentials names every entry, position-labelling the unnamed ones.
func LabelCredentials(credentials []any) []string { return config.LabelCredentials(credentials) }

// SdkConfigFromCredentials builds one validated config; overrides are keyed by snake_case field names.
func SdkConfigFromCredentials(credentials []any, instance any, overrides map[string]any) (SdkConfig, error) {
	return config.SdkConfigFromCredentials(credentials, instance, overrides)
}

// NewSdkConfig validates a config and returns it normalised, the response format resolved.
func NewSdkConfig(sdkConfig SdkConfig) (*SdkConfig, error) { return config.NewSdkConfig(sdkConfig) }

const (
	// V4 is the one resource version this SDK serves.
	V4                 = config.V4
	PrintFormatPretty  = config.PrintFormatPretty
	PrintFormatCompact = config.PrintFormatCompact

	ResponseFormatMap     = config.ResponseFormatMap
	ResponseFormatJSON    = config.ResponseFormatJSON
	ResponseFormatDTO     = config.ResponseFormatDTO
	DefaultResponseFormat = config.DefaultResponseFormat
	// ResponseFormatEnvVar is the lowest-priority tier deciding ApiResponse.Data.
	ResponseFormatEnvVar = config.ResponseFormatEnvVar

	// FormatMap, FormatJSON and FormatDto are aliases of the ResponseFormat constants.
	FormatMap  = config.FormatMap
	FormatJSON = config.FormatJSON
	FormatDto  = config.FormatDto
)

func IsResponseFormat(value string) bool { return config.IsResponseFormat(value) }

// ResponseFormatNames returns a fresh copy of the three legal response formats.
func ResponseFormatNames() []string { return config.ResponseFormatNames() }

// ParseResponseFormat reads a response format by name, refusing any other.
func ParseResponseFormat(value string) (string, error) { return config.ParseResponseFormat(value) }

type (
	// SdkError is the one error type the SDK returns.
	SdkError  = exerrors.SdkError
	ErrorType = exerrors.ErrorType
	// RequestContext is the request record an SdkError carries; *RequestInfo implements it.
	RequestContext = exerrors.RequestContext
	// EndpointNamer is the optional half of RequestContext: it names the endpoint template.
	EndpointNamer            = exerrors.EndpointNamer
	InsecureTransportWarning = exerrors.InsecureTransportWarning
	// FieldError is one field-level failure from the v4 error envelope.
	FieldError = exerrors.FieldError
)

// The sentinels are variables because errors.Is compares identity.
var (
	// ErrConfig matches configuration errors (errors.Is(err, ErrConfig)).
	ErrConfig = exerrors.ErrConfig
	ErrUsage  = exerrors.ErrUsage
	// ErrAttribute matches lookups of an unknown name; it is also a usage error (errors.Is both).
	ErrAttribute = exerrors.ErrAttribute
	ErrAuth      = exerrors.ErrAuth
	// ErrTokenCache matches a token cache file that is present and cannot be used.
	ErrTokenCache = auth.ErrTokenCache
	// ErrTokenLock matches work refused because no cross-process lock over the token cache could be taken.
	ErrTokenLock = auth.ErrTokenLock
)

// Build errors with these constructors: a hand-built &SdkError{} matches no sentinel.

func NewAttributeError(message string) *SdkError { return exerrors.NewAttributeError(message) }

func NewConfigError(message string) *SdkError { return exerrors.NewConfigError(message) }

func NewUsageError(message string) *SdkError { return exerrors.NewUsageError(message) }

func NewAuthError(message string) *SdkError { return exerrors.NewAuthError(message) }

func NewAuthResponseError(message string, statusCode int, errorType ErrorType, rawResponse any,
	responseHeaders map[string]string, requestInfo RequestContext, messages []string) *SdkError {
	return exerrors.NewAuthResponseError(message, statusCode, errorType, rawResponse,
		responseHeaders, requestInfo, messages)
}

func NewInsecureTransportWarning(message string) *InsecureTransportWarning {
	return exerrors.NewInsecureTransportWarning(message)
}

const (
	ErrorTypeNone             = exerrors.ErrorTypeNone
	ErrorTypeRedirect         = exerrors.ErrorTypeRedirect
	ErrorTypeValidation       = exerrors.ErrorTypeValidation
	ErrorTypeAuthentication   = exerrors.ErrorTypeAuthentication
	ErrorTypeAuthorization    = exerrors.ErrorTypeAuthorization
	ErrorTypeNotFound         = exerrors.ErrorTypeNotFound
	ErrorTypeMethodNotAllowed = exerrors.ErrorTypeMethodNotAllowed
	ErrorTypeConflict         = exerrors.ErrorTypeConflict
	ErrorTypeRateLimit        = exerrors.ErrorTypeRateLimit
	ErrorTypeServer           = exerrors.ErrorTypeServer
	ErrorTypeNetwork          = exerrors.ErrorTypeNetwork
	ErrorTypeTimeout          = exerrors.ErrorTypeTimeout
	ErrorTypeUnknown          = exerrors.ErrorTypeUnknown
)

// ClassifyErrorType derives the classification from a status.
func ClassifyErrorType(statusCode int) ErrorType { return http.ClassifyErrorType(statusCode) }

const (
	Redacted = redact.Redacted
	// TooDeep replaces a structure nested deeper than MaxRedactionDepth.
	TooDeep = redact.TooDeep
	// Circular replaces a container that contains itself.
	Circular = redact.Circular
	// MaxRedactionDepth is how deep RedactValue walks before writing TooDeep.
	MaxRedactionDepth = redact.MaxRedactionDepth
)

func ShouldRedact(key string) bool { return redact.ShouldRedact(key) }

// RedactValue walks a value and masks every secret-named entry, returning a fresh copy.
func RedactValue(value any) any { return redact.RedactValue(value) }

// RedactMapping masks the secret-named entries of a header map, returning a fresh copy.
func RedactMapping(headers map[string]string) map[string]string { return redact.RedactMapping(headers) }

// RedactText masks the secret-named query parameters and userinfo inside one string.
func RedactText(value string) string { return redact.RedactText(value) }

// SensitiveKeyMarkers returns a fresh copy of the substrings that make a key name secret.
func SensitiveKeyMarkers() []string { return redact.SensitiveKeyMarkers() }

type (
	ApiResponse      = http.ApiResponse
	RequestInfo      = http.RequestInfo
	RequestOptions   = http.RequestOptions
	TransportRequest = http.TransportRequest
	// Requester is the subset of the client that resource modules call.
	Requester     = http.Requester
	SortDirection = http.SortDirection
	// DTOShape builds the typed view of one response body.
	DTOShape = http.DTOShape
	// DTOPayload is a typed request body: anything that can hand over the map the SDK encodes.
	DTOPayload = http.DTOPayload
	// QueryPair is one normalised query parameter, in wire order.
	QueryPair = http.QueryPair
)

func BuildURL(baseURL, path string) (string, error) { return http.BuildURL(baseURL, path) }

// SubstitutePathParams fills a path template's {placeholders} from a map.
func SubstitutePathParams(pathTemplate string, pathParams map[string]any) (string, error) {
	return http.SubstitutePathParams(pathTemplate, pathParams)
}

// FindPathParamNames lists the slots a path template declares.
func FindPathParamNames(pathTemplate string) []string { return http.FindPathParamNames(pathTemplate) }

func BuildQueryParams(params map[string]any) (map[string]any, error) {
	return http.BuildQueryParams(params)
}

// EncodeQueryParams renders normalised pairs as a query string.
func EncodeQueryParams(pairs []QueryPair) string { return http.EncodeQueryParams(pairs) }

// NewRequestInfo is the redacting constructor for RequestInfo.
func NewRequestInfo(clientName, apiVersion, httpMethod, url string, queryParams map[string]any,
	body any, headers map[string]string, attemptCount, maxAttempts int, elapsedSeconds float64) *RequestInfo {
	return http.NewRequestInfo(clientName, apiVersion, httpMethod, url, queryParams, body, headers,
		attemptCount, maxAttempts, elapsedSeconds)
}

// DefaultRetryableHTTPCodes is the default retryable set for a safe method; a fresh copy each call.
func DefaultRetryableHTTPCodes() []int { return config.DefaultRetryableCodes() }

// DefaultUnsafeRetryableHTTPCodes is empty, so no write is replayed unless a caller opts in.
func DefaultUnsafeRetryableHTTPCodes() []int { return config.DefaultUnsafeRetryableCodes() }

const (
	SortAsc  = http.SortAsc
	SortDesc = http.SortDesc
	// Ascending and Descending are aliases of SortAsc and SortDesc.
	Ascending  = http.Ascending
	Descending = http.Descending
	// IdempotencyHeader is the header v4 uses to collapse a repeated create.
	IdempotencyHeader = http.IdempotencyHeader
)

// FilterOperators are the comparisons a v4 filter name can carry; FilterParam refuses any other.
var FilterOperators = http.FilterOperators

// FilterParam builds a v4 filter parameter name: filter_<field> for equality, else filter_<field>_<op>.
func FilterParam(field string, operator ...string) (string, error) {
	return http.FilterParam(field, operator...)
}

// MustFilterParam is FilterParam for a literal known good at compile time.
func MustFilterParam(field string, operator ...string) string {
	return http.MustFilterParam(field, operator...)
}

// SortParam builds a `sort` value, "-" prefixed to descend.
func SortParam(field string, direction ...SortDirection) (string, error) {
	return http.SortParam(field, direction...)
}

// MustSortParam is SortParam for a literal known good at compile time.
func MustSortParam(field string, direction ...SortDirection) string {
	return http.MustSortParam(field, direction...)
}

// SortParams builds the `order_by` and `direction` pair the other v4 families take.
func SortParams(field string, descending bool) (map[string]any, error) {
	return http.SortParams(field, descending)
}

// MustSortParams is SortParams for a literal known good at compile time.
func MustSortParams(field string, descending bool) map[string]any {
	return http.MustSortParams(field, descending)
}

type (
	HTTPTransport   = http.HTTPTransport
	TransportOption = http.TransportOption
	RetryPolicy     = http.RetryPolicy
)

// MaxMultipartBytes caps the multipart body assembled in memory: 64 MiB.
const MaxMultipartBytes = http.MaxMultipartBytes

// MaxRequestBodyBytes caps the final wire body of any request, whichever kind it is: 64 MiB.
const MaxRequestBodyBytes = http.MaxRequestBodyBytes

const MaxCABundleBytes = http.MaxCABundleBytes

const MaxTokenCacheBytes = auth.MaxTokenCacheBytes

func NewHTTPTransport(sdkConfig SdkConfig, options ...TransportOption) (*HTTPTransport, error) {
	return http.NewHTTPTransport(sdkConfig, options...)
}

func WithHTTPClient(httpClient *nethttp.Client) TransportOption {
	return http.WithHTTPClient(httpClient)
}

// WithProgress wires the transport's per-attempt progress hook.
func WithProgress(onProgress func(string)) TransportOption { return http.WithProgress(onProgress) }

// WithTransportRetryPolicy is http.WithRetryPolicy, named apart from the client's WithRetryPolicy.
func WithTransportRetryPolicy(retry *RetryPolicy) TransportOption {
	return http.WithRetryPolicy(retry)
}

// NewRetryPolicy constructs a retry policy; a nil code list takes the matching default set.
func NewRetryPolicy(maxAttempts int, backoffSeconds, backoffMaxSeconds float64,
	retryable, unsafeRetryable []int) (*RetryPolicy, error) {
	return http.NewRetryPolicy(maxAttempts, backoffSeconds, backoffMaxSeconds, retryable, unsafeRetryable)
}

// SnapshotJSONValue deep-copies one JSON-shaped value at any depth; RequestOptions.Frozen uses it.
func SnapshotJSONValue(value any) any { return http.SnapshotJSONValue(value) }

func ReadCappedFile(path string, limitBytes int) ([]byte, error) {
	return http.ReadCappedFile(path, limitBytes)
}

type (
	TokenRecord      = auth.TokenRecord
	TokenStore       = auth.TokenStore
	MemoryTokenStore = auth.MemoryTokenStore
	// FileTokenStore is token storage in a JSON file, shared across processes under a lock file.
	FileTokenStore = auth.FileTokenStore
	TokenProvider  = auth.TokenProvider
)

func NewMemoryTokenStore() *MemoryTokenStore { return auth.NewMemoryTokenStore() }

// NewFileTokenStore builds token storage in a JSON file, keyed by credential fingerprint.
func NewFileTokenStore(path string) *FileTokenStore { return auth.NewFileTokenStore(path) }

func NewTokenProvider(sdkConfig *SdkConfig, store TokenStore, transport *HTTPTransport) *TokenProvider {
	return auth.NewTokenProvider(sdkConfig, store, transport)
}

func BuildCredentialsNamespace(baseURL, clientID, clientSecret, grantType string) string {
	return auth.BuildCredentialsNamespace(baseURL, clientID, clientSecret, grantType)
}

// DefaultTokenCachePath is the per-user cache path used when no TokenCachePath is named.
func DefaultTokenCachePath() string { return auth.DefaultTokenCachePath() }

type (
	ExsitedClient  = client.ExsitedClient
	ClientOption   = client.ClientOption
	ClientRegistry = client.ClientRegistry
	// BoundResource is a resource family with a client applied; type-assert Value to the resource.
	BoundResource   = client.BoundResource
	ResourceFactory = client.ResourceFactory
	PaginateOptions = client.PaginateOptions
	// PageOptions is an alias of PaginateOptions.
	PageOptions = client.PageOptions
)

const (
	// DefaultRecordKey is the envelope key a cursor-paginated listing uses.
	DefaultRecordKey = client.DefaultRecordKey
	// NoNextPage is what the offset shape puts in next_page to mean "no more".
	NoNextPage = client.NoNextPage
)

func NewClient(sdkConfig SdkConfig, options ...ClientOption) (*ExsitedClient, error) {
	return client.NewClient(sdkConfig, options...)
}

// NewExsitedClient is an alias of NewClient.
func NewExsitedClient(sdkConfig SdkConfig, options ...ClientOption) (*ExsitedClient, error) {
	return client.NewExsitedClient(sdkConfig, options...)
}

func WithTransport(transport *HTTPTransport) ClientOption { return client.WithTransport(transport) }

func WithTokenStore(store TokenStore) ClientOption { return client.WithTokenStore(store) }

// WithTokenProvider injects a built TokenProvider, so clients can share one provider and cache.
func WithTokenProvider(provider *TokenProvider) ClientOption {
	return client.WithTokenProvider(provider)
}

func WithRetryPolicy(retry *RetryPolicy) ClientOption { return client.WithRetryPolicy(retry) }

func WithLogger(log func(string)) ClientOption { return client.WithLogger(log) }

func IterPages(ctx context.Context, requester Requester, path string,
	options PaginateOptions) iter.Seq2[*ApiResponse, error] {
	return client.IterPages(ctx, requester, path, options)
}

// IterRecords walks every record of a listing across every page, each element untouched.
func IterRecords(ctx context.Context, requester Requester, path, recordKey string,
	options PaginateOptions) iter.Seq2[any, error] {
	return client.IterRecords(ctx, requester, path, recordKey, options)
}

// ReadNextCursor reads the cursor shape's next position.
func ReadNextCursor(block map[string]any) string { return client.ReadNextCursor(block) }

// ReadNextOffset reads the offset shape's next position.
func ReadNextOffset(block map[string]any, recordsOnPage int) (int64, bool) {
	return client.ReadNextOffset(block, recordsOnPage)
}

func NewClientRegistry(credentials ...any) (*ClientRegistry, error) {
	return client.NewClientRegistry(credentials...)
}

func RegisterResourceModule(resourceName string, factory ResourceFactory) error {
	return client.RegisterResourceModule(resourceName, factory)
}

func RegisteredResourceNames() []string { return client.RegisteredResourceNames() }

// ResolveResourceModule returns the factory registered under a name.
func ResolveResourceModule(resourceName string) (ResourceFactory, error) {
	return client.ResolveResourceModule(resourceName)
}

// NewBoundResource builds the family registered under a name over a requester.
func NewBoundResource(requester Requester, resourceName string) (*BoundResource, error) {
	return client.NewBoundResource(requester, resourceName)
}

// RegisterInstances registers credential blocks by client name for later BuildClient calls.
func RegisterInstances(credentials ...any) error { return client.RegisterInstances(credentials...) }

func BuildClient(instance any, overrides ...func(*SdkConfig)) (*ExsitedClient, error) {
	return client.BuildClient(instance, overrides...)
}

func ListInstanceNames() ([]string, error) { return client.ListInstanceNames() }

func GetRegisteredInstances() ([]SdkConfig, error) { return client.GetRegisteredInstances() }

// RenderExchange renders one exchange as the console would print it.
func RenderExchange(response *ApiResponse, body any, files map[string]any,
	sdkConfig SdkConfig, formEncoded bool) string {
	return client.RenderExchange(response, body, files, sdkConfig, formEncoded)
}

func BuildCurlCommand(response *ApiResponse, payload any, files map[string]any,
	sdkConfig SdkConfig, formEncoded bool) string {
	return client.BuildCurlCommand(response, payload, files, sdkConfig, formEncoded)
}

// RenderJSONValue renders one value the way the console prints it: indented, or one line when compact.
func RenderJSONValue(value any, compact bool) string { return client.RenderJSONValue(value, compact) }

// ResolveBoxGlyphs returns the rule and ellipsis this console can carry.
func ResolveBoxGlyphs() (rule, ellipsis string) { return client.ResolveBoxGlyphs() }

// EncodeForConsole spells every rune above ASCII as an escape when the console cannot carry it.
func EncodeForConsole(text string) string { return client.EncodeForConsole(text) }

// Neutralise escapes every control character in one value, newlines included.
func Neutralise(text string) string { return client.Neutralise(text) }

// NeutraliseKeepNewlines is Neutralise for a block printed as several lines.
func NeutraliseKeepNewlines(text string) string { return client.NeutraliseKeepNewlines(text) }

type (
	// Dto is embedded by the typed views.
	Dto = dto.Dto
	// Enveloped is implemented by a view whose body is wrapped in one key.
	Enveloped = dto.Enveloped
	// BesideEnvelope is implemented by a view with fields beside the envelope.
	BesideEnvelope = dto.BesideEnvelope
	// TopLevelFields is BesideEnvelope with the method named TopLevelFields; either is honoured.
	TopLevelFields = dto.TopLevelFields
)

// DTOFields is the wire names a DTO type declares, in declaration order.
func DTOFields(value any) []string { return dto.Fields(value) }

// DecodeDto reads a body into a typed view, converting nothing.
func DecodeDto(body any, target any) error { return dto.DecodeDto(body, target) }

// EncodeDto renders a typed payload as the map the SDK encodes.
func EncodeDto(source any) (map[string]any, error) { return dto.EncodeDto(source) }

// AbaFiles is the aba files endpoints bound to a client.
type AbaFiles = abafiles.AbaFiles

// NewAbaFiles binds the resource to a requester of your own.
func NewAbaFiles(requester Requester) *AbaFiles { return abafiles.New(requester) }

// The aba files path templates and body keys, for the raw-verb call.
const (
	ABAFileGeneratePath = abafiles.ABAFileGeneratePath
)

// AccountEmails is the account emails endpoints bound to a client.
type AccountEmails = accountemails.AccountEmails

// The account emails views, aliased so a type assertion works under either spelling.
type (
	AccountEmailsPagination          = accountemails.Pagination
	AccountCommunicationsEmailRecord = accountemails.AccountCommunicationsEmailRecord
	AccountCommunicationsEmailList   = accountemails.AccountCommunicationsEmailList
	AccountEmailsPaginationLinks     = accountemails.PaginationLinks
	AccountEmailsAttachment          = accountemails.Attachment
)

// NewAccountEmails binds the resource to a requester of your own.
func NewAccountEmails(requester Requester) *AccountEmails { return accountemails.New(requester) }

// AccountCommunicationsEmailListShape is the DTOShape reading a body as AccountCommunicationsEmailList, for the raw-verb call.
func AccountCommunicationsEmailListShape(decoded any) any {
	return accountemails.AccountCommunicationsEmailListShape(decoded)
}

// AccountCommunicationsEmailRecordShape is the DTOShape reading a body as AccountCommunicationsEmailRecord, for the raw-verb call.
func AccountCommunicationsEmailRecordShape(decoded any) any {
	return accountemails.AccountCommunicationsEmailRecordShape(decoded)
}

// The account emails path templates and body keys, for the raw-verb call.
const (
	AccountCommunicationsEmailPath = accountemails.AccountCommunicationsEmailPath

	AccountEmailsKey = accountemails.AccountEmailsKey
	EmailKey         = accountemails.EmailKey
)

// AccountingCodes is the accounting codes endpoints bound to a client.
type AccountingCodes = accountingcodes.AccountingCodes

// The accounting codes views, aliased so a type assertion works under either spelling.
type (
	AccountingCodesPagination      = accountingcodes.Pagination
	SettingsAccountingCodeRecord   = accountingcodes.SettingsAccountingCodeRecord
	SettingsAccountingCodeList     = accountingcodes.SettingsAccountingCodeList
	AccountingCodesPaginationLinks = accountingcodes.PaginationLinks
)

// NewAccountingCodes binds the resource to a requester of your own.
func NewAccountingCodes(requester Requester) *AccountingCodes { return accountingcodes.New(requester) }

// SettingsAccountingCodeListShape is the DTOShape reading a body as SettingsAccountingCodeList, for the raw-verb call.
func SettingsAccountingCodeListShape(decoded any) any {
	return accountingcodes.SettingsAccountingCodeListShape(decoded)
}

// SettingsAccountingCodeRecordShape is the DTOShape reading a body as SettingsAccountingCodeRecord, for the raw-verb call.
func SettingsAccountingCodeRecordShape(decoded any) any {
	return accountingcodes.SettingsAccountingCodeRecordShape(decoded)
}

// The accounting codes path templates and body keys, for the raw-verb call.
const (
	SettingsAccountingCodesPath = accountingcodes.SettingsAccountingCodesPath
	SettingsAccountingCodePath  = accountingcodes.SettingsAccountingCodePath

	AccountingCodesKey = accountingcodes.AccountingCodesKey
	AccountingCodeKey  = accountingcodes.AccountingCodeKey
)

// Accounts is the accounts endpoints bound to a client.
type Accounts = accounts.Accounts

// The accounts views, aliased so a type assertion works under either spelling.
type (
	AccountsReference               = accounts.Reference
	AccountsUser                    = accounts.User
	Currency                        = accounts.Currency
	TimeZone                        = accounts.TimeZone
	AccountsCustomForm              = accounts.CustomForm
	SocialLinks                     = accounts.SocialLinks
	Address                         = accounts.Address
	ContactName                     = accounts.ContactName
	ContactEmail                    = accounts.ContactEmail
	ContactPhone                    = accounts.ContactPhone
	Contact                         = accounts.Contact
	PaymentProcessor                = accounts.PaymentProcessor
	PaymentMethod                   = accounts.PaymentMethod
	BillingPreferences              = accounts.BillingPreferences
	AccountsNoteAttachment          = accounts.NoteAttachment
	Note                            = accounts.Note
	AccountsPagination              = accounts.Pagination
	AccountRecord                   = accounts.AccountRecord
	AccountList                     = accounts.AccountList
	AccountImageView                = accounts.AccountImageView
	NoteView                        = accounts.NoteView
	AccountsNoteList                = accounts.NoteList
	AccountsNoteFileList            = accounts.NoteFileList
	AccountCancelBody               = accounts.AccountCancelBody
	AccountInformationUpdate        = accounts.AccountInformationUpdate
	AccountSetPasswordBody          = accounts.AccountSetPasswordBody
	AccountImageUpload              = accounts.AccountImageUpload
	AddressWrite                    = accounts.AddressWrite
	AddressesCreate                 = accounts.AddressesCreate
	AddressUpdate                   = accounts.AddressUpdate
	ContactWrite                    = accounts.ContactWrite
	BillingPreferencesWrite         = accounts.BillingPreferencesWrite
	PaymentMethodUpdate             = accounts.PaymentMethodUpdate
	AccountsAccount                 = accounts.Account
	AccountAddressView              = accounts.AccountAddressView
	AccountContactView              = accounts.AccountContactView
	AccountCreditNoteCurrency       = accounts.AccountCreditNoteCurrency
	AccountCreditNoteList           = accounts.AccountCreditNoteList
	AccountCreditNoteRecord         = accounts.AccountCreditNoteRecord
	AccountCustomObjectList         = accounts.AccountCustomObjectList
	AccountCustomObjectRecord       = accounts.AccountCustomObjectRecord
	AccountInventoriesView          = accounts.AccountInventoriesView
	AccountInvoiceCurrency          = accounts.AccountInvoiceCurrency
	AccountInvoiceCustomForm        = accounts.AccountInvoiceCustomForm
	AccountInvoiceList              = accounts.AccountInvoiceList
	AccountInvoiceRecord            = accounts.AccountInvoiceRecord
	AccountNoteFileView             = accounts.AccountNoteFileView
	AccountPaymentMethodView        = accounts.AccountPaymentMethodView
	AccountQuotePartialInvoicesView = accounts.AccountQuotePartialInvoicesView
	AccountsCreatedBy               = accounts.CreatedBy
	CustomAttribute                 = accounts.CustomAttribute
	AccountsInvoice                 = accounts.Invoice
	LastPayment                     = accounts.LastPayment
	AccountsLastUpdatedBy           = accounts.LastUpdatedBy
	AccountsOrder                   = accounts.Order
	AccountsPayment                 = accounts.Payment
	AccountsAccountingCode          = accounts.AccountingCode
	AccountsCommunicationPreference = accounts.CommunicationPreference
	CustomObject                    = accounts.CustomObject
	AccountsPaginationLinks         = accounts.PaginationLinks
	AccountsWatcher                 = accounts.Watcher
)

// NewAccounts binds the resource to a requester of your own.
func NewAccounts(requester Requester) *Accounts { return accounts.New(requester) }

// AccountListShape is the DTOShape reading a body as AccountList, for the raw-verb call.
func AccountListShape(decoded any) any { return accounts.AccountListShape(decoded) }

// AccountRecordShape is the DTOShape reading a body as AccountRecord, for the raw-verb call.
func AccountRecordShape(decoded any) any { return accounts.AccountRecordShape(decoded) }

// AccountImageViewShape is the DTOShape reading a body as AccountImageView, for the raw-verb call.
func AccountImageViewShape(decoded any) any { return accounts.AccountImageViewShape(decoded) }

// NoteViewShape is the DTOShape reading a body as NoteView, for the raw-verb call.
func NoteViewShape(decoded any) any { return accounts.NoteViewShape(decoded) }

// AccountsNoteFileListShape is the DTOShape reading a body as AccountsNoteFileList, for the raw-verb call.
func AccountsNoteFileListShape(decoded any) any { return accounts.NoteFileListShape(decoded) }

// AccountsNoteListShape is the DTOShape reading a body as AccountsNoteList, for the raw-verb call.
func AccountsNoteListShape(decoded any) any { return accounts.NoteListShape(decoded) }

// AccountsQueryParams returns a fresh copy of the names the listing accepts.
func AccountsQueryParams() []string { return accounts.AccountsQueryParams() }

// AccountsAccountAddressViewShape is the DTOShape reading a body as AccountsAccountAddressView, for the raw-verb call.
func AccountsAccountAddressViewShape(decoded any) any {
	return accounts.AccountAddressViewShape(decoded)
}

// AccountsAccountContactViewShape is the DTOShape reading a body as AccountsAccountContactView, for the raw-verb call.
func AccountsAccountContactViewShape(decoded any) any {
	return accounts.AccountContactViewShape(decoded)
}

// AccountsAccountCreditNoteListShape is the DTOShape reading a body as AccountsAccountCreditNoteList, for the raw-verb call.
func AccountsAccountCreditNoteListShape(decoded any) any {
	return accounts.AccountCreditNoteListShape(decoded)
}

// AccountsAccountCustomObjectListShape is the DTOShape reading a body as AccountsAccountCustomObjectList, for the raw-verb call.
func AccountsAccountCustomObjectListShape(decoded any) any {
	return accounts.AccountCustomObjectListShape(decoded)
}

// AccountsAccountInventoriesViewShape is the DTOShape reading a body as AccountsAccountInventoriesView, for the raw-verb call.
func AccountsAccountInventoriesViewShape(decoded any) any {
	return accounts.AccountInventoriesViewShape(decoded)
}

// AccountsAccountInvoiceListShape is the DTOShape reading a body as AccountsAccountInvoiceList, for the raw-verb call.
func AccountsAccountInvoiceListShape(decoded any) any {
	return accounts.AccountInvoiceListShape(decoded)
}

// AccountsAccountNoteFileViewShape is the DTOShape reading a body as AccountsAccountNoteFileView, for the raw-verb call.
func AccountsAccountNoteFileViewShape(decoded any) any {
	return accounts.AccountNoteFileViewShape(decoded)
}

// AccountsAccountPaymentMethodViewShape is the DTOShape reading a body as AccountsAccountPaymentMethodView, for the raw-verb call.
func AccountsAccountPaymentMethodViewShape(decoded any) any {
	return accounts.AccountPaymentMethodViewShape(decoded)
}

// AccountsAccountQuotePartialInvoicesViewShape is the DTOShape reading a body as AccountsAccountQuotePartialInvoicesView, for the raw-verb call.
func AccountsAccountQuotePartialInvoicesViewShape(decoded any) any {
	return accounts.AccountQuotePartialInvoicesViewShape(decoded)
}

// The accounts path templates and body keys, for the raw-verb call.
const (
	AccountsPath                        = accounts.AccountsPath
	AccountPath                         = accounts.AccountPath
	AccountInformationPath              = accounts.AccountInformationPath
	AccountCancelPath                   = accounts.AccountCancelPath
	AccountSetPasswordPath              = accounts.AccountSetPasswordPath
	AccountContactsPath                 = accounts.AccountContactsPath
	AccountContactPath                  = accounts.AccountContactPath
	AccountAddressesPath                = accounts.AccountAddressesPath
	AccountAddressPath                  = accounts.AccountAddressPath
	AccountBillingPreferencesPath       = accounts.AccountBillingPreferencesPath
	AccountPaymentMethodsPath           = accounts.AccountPaymentMethodsPath
	AccountPaymentMethodPath            = accounts.AccountPaymentMethodPath
	AccountsAccountNotesPath            = accounts.AccountNotesPath
	AccountsAccountNotePath             = accounts.AccountNotePath
	AccountNoteFilesPath                = accounts.AccountNoteFilesPath
	AccountNoteFilePath                 = accounts.AccountNoteFilePath
	AccountImagePath                    = accounts.AccountImagePath
	AccountEmailsPath                   = accounts.AccountEmailsPath
	AccountEmailPath                    = accounts.AccountEmailPath
	AccountsAccountOrdersPath           = accounts.AccountOrdersPath
	AccountsAccountInvoicesPath         = accounts.AccountInvoicesPath
	AccountsAccountPaymentsPath         = accounts.AccountPaymentsPath
	AccountsAccountRefundsPath          = accounts.AccountRefundsPath
	AccountCreditNotesPath              = accounts.AccountCreditNotesPath
	AccountsAccountPurchaseOrdersPath   = accounts.AccountPurchaseOrdersPath
	AccountsAccountPurchaseInvoicesPath = accounts.AccountPurchaseInvoicesPath
	AccountsAccountPurchasePaymentsPath = accounts.AccountPurchasePaymentsPath
	AccountsAccountPurchaseRefundsPath  = accounts.AccountPurchaseRefundsPath
	AccountCustomObjectsPath            = accounts.AccountCustomObjectsPath
	AccountInventoriesPath              = accounts.AccountInventoriesPath
	AccountQuotePartialInvoicesPath     = accounts.AccountQuotePartialInvoicesPath

	AccountsKey           = accounts.AccountsKey
	AccountKey            = accounts.AccountKey
	AddressKey            = accounts.AddressKey
	BillingPreferencesKey = accounts.BillingPreferencesKey
	ContactKey            = accounts.ContactKey
	AccountsNoteKey       = accounts.NoteKey
	PaymentMethodKey      = accounts.PaymentMethodKey
)

// Alerts is the alerts endpoints bound to a client.
type Alerts = alerts.Alerts

// The alerts views, aliased so a type assertion works under either spelling.
type (
	AlertsPagination       = alerts.Pagination
	AlertsRecord           = alerts.AlertsRecord
	AlertsEmployeeByIDList = alerts.AlertsEmployeeByIDList
	AlertsLabourByIDList   = alerts.AlertsLabourByIDList
	AlertsPaginationLinks  = alerts.PaginationLinks
)

// NewAlerts binds the resource to a requester of your own.
func NewAlerts(requester Requester) *Alerts { return alerts.New(requester) }

// AlertsRecordShape is the DTOShape reading a body as AlertsRecord, for the raw-verb call.
func AlertsRecordShape(decoded any) any { return alerts.AlertsRecordShape(decoded) }

// AlertsAlertsEmployeeByIDListShape is the DTOShape reading a body as AlertsAlertsEmployeeByIDList, for the raw-verb call.
func AlertsAlertsEmployeeByIDListShape(decoded any) any {
	return alerts.AlertsEmployeeByIDListShape(decoded)
}

// AlertsAlertsLabourByIDListShape is the DTOShape reading a body as AlertsAlertsLabourByIDList, for the raw-verb call.
func AlertsAlertsLabourByIDListShape(decoded any) any {
	return alerts.AlertsLabourByIDListShape(decoded)
}

// The alerts path templates and body keys, for the raw-verb call.
const (
	AlertsPath             = alerts.AlertsPath
	AlertPath              = alerts.AlertPath
	AlertsAccountByIDPath  = alerts.AlertsAccountByIDPath
	AlertsEmployeeByIDPath = alerts.AlertsEmployeeByIDPath
	AlertsLabourByIDPath   = alerts.AlertsLabourByIDPath

	AlertsKey = alerts.AlertsKey
)

// Approvals is the approvals endpoints bound to a client.
type Approvals = approvals.Approvals

// The approvals views, aliased so a type assertion works under either spelling.
type (
	ApprovalsPagination      = approvals.Pagination
	ApprovalsAccountRecord   = approvals.ApprovalsAccountRecord
	ApprovalsAccountList     = approvals.ApprovalsAccountList
	ApprovalsSaleorderRecord = approvals.ApprovalsSaleorderRecord
	ApprovalsPaginationLinks = approvals.PaginationLinks
)

// NewApprovals binds the resource to a requester of your own.
func NewApprovals(requester Requester) *Approvals { return approvals.New(requester) }

// ApprovalsAccountListShape is the DTOShape reading a body as ApprovalsAccountList, for the raw-verb call.
func ApprovalsAccountListShape(decoded any) any { return approvals.ApprovalsAccountListShape(decoded) }

// ApprovalsSaleorderRecordShape is the DTOShape reading a body as ApprovalsSaleorderRecord, for the raw-verb call.
func ApprovalsSaleorderRecordShape(decoded any) any {
	return approvals.ApprovalsSaleorderRecordShape(decoded)
}

// The approvals path templates and body keys, for the raw-verb call.
const (
	ApprovalsAccountPath       = approvals.ApprovalsAccountPath
	ApprovalsSaleorderPath     = approvals.ApprovalsSaleorderPath
	ApprovalsAccountByUUIDPath = approvals.ApprovalsAccountByUUIDPath

	ApprovalsKey = approvals.ApprovalsKey
)

// Bookings is the bookings endpoints bound to a client.
type Bookings = bookings.Bookings

// The bookings views, aliased so a type assertion works under either spelling.
type (
	BookingsPagination               = bookings.Pagination
	BookingRecord                    = bookings.BookingRecord
	BookingList                      = bookings.BookingList
	BookingByIdRecord                = bookings.BookingByIdRecord
	BookingsContractProperties       = bookings.ContractProperties
	BookingsPaginationLinks          = bookings.PaginationLinks
	BookingsAccount                  = bookings.Account
	BookingsBillingAddress           = bookings.BillingAddress
	BookingsBillingPreferences       = bookings.BillingPreferences
	BookingsCommunicationPreference  = bookings.CommunicationPreference
	BookingsCurrency                 = bookings.Currency
	BookingsCustomAttribute          = bookings.CustomAttribute
	BookingsCustomForm               = bookings.CustomForm
	BookingsDiscountProfile          = bookings.DiscountProfile
	BookingsItem                     = bookings.Item
	BookingsItemAccountingCode       = bookings.ItemAccountingCode
	BookingsItemPriceSnapshot        = bookings.ItemPriceSnapshot
	BookingsItemPriceTax             = bookings.ItemPriceTax
	BookingsItemProperties           = bookings.ItemProperties
	BookingsItemSaleTaxConfiguration = bookings.ItemSaleTaxConfiguration
	BookingsKpis                     = bookings.Kpis
	BookingsLine                     = bookings.Line
	BookingsManager                  = bookings.Manager
	BookingsPreOrderStockDetails     = bookings.PreOrderStockDetails
	BookingsPricingRule              = bookings.PricingRule
	BookingsProformaProperties       = bookings.ProformaProperties
	BookingsReferralAccount          = bookings.ReferralAccount
	BookingsSalesRep                 = bookings.SalesRep
	BookingsShippingAddress          = bookings.ShippingAddress
	BookingsShippingProfile          = bookings.ShippingProfile
	BookingsTaxCode                  = bookings.TaxCode
	BookingsTimeZone                 = bookings.TimeZone
)

// NewBookings binds the resource to a requester of your own.
func NewBookings(requester Requester) *Bookings { return bookings.New(requester) }

// BookingListShape is the DTOShape reading a body as BookingList, for the raw-verb call.
func BookingListShape(decoded any) any { return bookings.BookingListShape(decoded) }

// BookingByIdRecordShape is the DTOShape reading a body as BookingByIdRecord, for the raw-verb call.
func BookingByIdRecordShape(decoded any) any { return bookings.BookingByIdRecordShape(decoded) }

// The bookings path templates and body keys, for the raw-verb call.
const (
	BookingPath     = bookings.BookingPath
	BookingByIDPath = bookings.BookingByIDPath

	BookingsKey      = bookings.BookingsKey
	BookingsOrderKey = bookings.OrderKey
)

// CreditNoteApplications is the credit note applications endpoints bound to a client.
type CreditNoteApplications = creditnoteapplications.CreditNoteApplications

// The credit note applications views, aliased so a type assertion works under either spelling.
type (
	CreditNoteApplicationsPagination      = creditnoteapplications.Pagination
	CreditNoteApplicationRecord           = creditnoteapplications.CreditNoteApplicationRecord
	CreditNoteApplicationList             = creditnoteapplications.CreditNoteApplicationList
	CreditNoteApplicationReadRecord       = creditnoteapplications.CreditNoteApplicationReadRecord
	InvoiceCreditNoteApplicationsRecord   = creditnoteapplications.InvoiceCreditNoteApplicationsRecord
	CreditNoteApplicationsPaginationLinks = creditnoteapplications.PaginationLinks
	CreditNoteApplicationsCreatedBy       = creditnoteapplications.CreatedBy
	CreditNoteApplicationsCreditNote      = creditnoteapplications.CreditNote
	CreditNoteApplicationsPayment         = creditnoteapplications.Payment
	CreditNoteApplicationsRefund          = creditnoteapplications.Refund
)

// NewCreditNoteApplications binds the resource to a requester of your own.
func NewCreditNoteApplications(requester Requester) *CreditNoteApplications {
	return creditnoteapplications.New(requester)
}

// CreditNoteApplicationListShape is the DTOShape reading a body as CreditNoteApplicationList, for the raw-verb call.
func CreditNoteApplicationListShape(decoded any) any {
	return creditnoteapplications.CreditNoteApplicationListShape(decoded)
}

// CreditNoteApplicationReadRecordShape is the DTOShape reading a body as CreditNoteApplicationReadRecord, for the raw-verb call.
func CreditNoteApplicationReadRecordShape(decoded any) any {
	return creditnoteapplications.CreditNoteApplicationReadRecordShape(decoded)
}

// InvoiceCreditNoteApplicationsRecordShape is the DTOShape reading a body as InvoiceCreditNoteApplicationsRecord, for the raw-verb call.
func InvoiceCreditNoteApplicationsRecordShape(decoded any) any {
	return creditnoteapplications.InvoiceCreditNoteApplicationsRecordShape(decoded)
}

// The credit note applications path templates and body keys, for the raw-verb call.
const (
	CreditNoteApplicationsPath        = creditnoteapplications.CreditNoteApplicationsPath
	CreditNoteApplicationPath         = creditnoteapplications.CreditNoteApplicationPath
	InvoiceCreditNoteApplicationsPath = creditnoteapplications.InvoiceCreditNoteApplicationsPath

	CreditNoteApplicationsKey = creditnoteapplications.CreditNoteApplicationsKey
	CreditNoteApplicationKey  = creditnoteapplications.CreditNoteApplicationKey
)

// CreditNotes is the credit notes endpoints bound to a client.
type CreditNotes = creditnotes.CreditNotes

// The credit notes views, aliased so a type assertion works under either spelling.
type (
	CreditNotesPagination      = creditnotes.Pagination
	CreditNoteRecord           = creditnotes.CreditNoteRecord
	CreditNoteList             = creditnotes.CreditNoteList
	CreditNoteReadRecord       = creditnotes.CreditNoteReadRecord
	CreditNoteUpdateRecord     = creditnotes.CreditNoteUpdateRecord
	CreditNotesPaginationLinks = creditnotes.PaginationLinks
	CreditNotesAccount         = creditnotes.Account
	CreditNotesCreatedBy       = creditnotes.CreatedBy
	CreditNotesCurrency        = creditnotes.Currency
	CreditNotesCustomAttribute = creditnotes.CustomAttribute
	CreditNotesInvoice         = creditnotes.Invoice
	CreditNotesLastUpdatedBy   = creditnotes.LastUpdatedBy
	CreditNotesPayment         = creditnotes.Payment
)

// NewCreditNotes binds the resource to a requester of your own.
func NewCreditNotes(requester Requester) *CreditNotes { return creditnotes.New(requester) }

// CreditNoteListShape is the DTOShape reading a body as CreditNoteList, for the raw-verb call.
func CreditNoteListShape(decoded any) any { return creditnotes.CreditNoteListShape(decoded) }

// CreditNoteReadRecordShape is the DTOShape reading a body as CreditNoteReadRecord, for the raw-verb call.
func CreditNoteReadRecordShape(decoded any) any {
	return creditnotes.CreditNoteReadRecordShape(decoded)
}

// CreditNotesCreditNoteUpdateRecordShape is the DTOShape reading a body as CreditNotesCreditNoteUpdateRecord, for the raw-verb call.
func CreditNotesCreditNoteUpdateRecordShape(decoded any) any {
	return creditnotes.CreditNoteUpdateRecordShape(decoded)
}

// The credit notes path templates and body keys, for the raw-verb call.
const (
	CreditNotesPath = creditnotes.CreditNotesPath
	CreditNotePath  = creditnotes.CreditNotePath

	CreditNotesKey = creditnotes.CreditNotesKey
	CreditNoteKey  = creditnotes.CreditNoteKey
)

// CustomAttributes is the custom attributes endpoints bound to a client.
type CustomAttributes = customattributes.CustomAttributes

// The custom attributes views, aliased so a type assertion works under either spelling.
type (
	CustomAttributesPagination      = customattributes.Pagination
	CustomAttributeRecord           = customattributes.CustomAttributeRecord
	CustomAttributeList             = customattributes.CustomAttributeList
	CustomAttributesRecord          = customattributes.CustomAttributesRecord
	CustomAttributesActivityLog     = customattributes.ActivityLog
	Option                          = customattributes.Option
	CustomAttributesPaginationLinks = customattributes.PaginationLinks
	UseIn                           = customattributes.UseIn
)

// NewCustomAttributes binds the resource to a requester of your own.
func NewCustomAttributes(requester Requester) *CustomAttributes {
	return customattributes.New(requester)
}

// CustomAttributeListShape is the DTOShape reading a body as CustomAttributeList, for the raw-verb call.
func CustomAttributeListShape(decoded any) any {
	return customattributes.CustomAttributeListShape(decoded)
}

// CustomAttributesRecordShape is the DTOShape reading a body as CustomAttributesRecord, for the raw-verb call.
func CustomAttributesRecordShape(decoded any) any {
	return customattributes.CustomAttributesRecordShape(decoded)
}

// CustomAttributeRecordShape is the DTOShape reading a body as CustomAttributeRecord, for the raw-verb call.
func CustomAttributeRecordShape(decoded any) any {
	return customattributes.CustomAttributeRecordShape(decoded)
}

// The custom attributes path templates and body keys, for the raw-verb call.
const (
	CustomAttributesPath = customattributes.CustomAttributesPath
	CustomAttributePath  = customattributes.CustomAttributePath

	CustomAttributesKey = customattributes.CustomAttributesKey
	CustomAttributeKey  = customattributes.CustomAttributeKey
)

// CustomComponents is the custom components endpoints bound to a client.
type CustomComponents = customcomponents.CustomComponents

// The custom components views, aliased so a type assertion works under either spelling.
type (
	CustomComponentsPagination            = customcomponents.Pagination
	CustomComponentsComponentRecord       = customcomponents.ComponentRecord
	CustomComponentsComponentList         = customcomponents.ComponentList
	ComponentByUuidRecord                 = customcomponents.ComponentByUuidRecord
	ComponentObjectRecord                 = customcomponents.ComponentObjectRecord
	ComponentObjectsRecord                = customcomponents.ComponentObjectsRecord
	ComponentObjectActivityRecord         = customcomponents.ComponentObjectActivityRecord
	ComponentObjectCustomObjectsRecord    = customcomponents.ComponentObjectCustomObjectsRecord
	Component                             = customcomponents.Component
	CustomComponentsPaginationLinks       = customcomponents.PaginationLinks
	CustomComponentsActivityLog           = customcomponents.ActivityLog
	CustomComponentsAttribute             = customcomponents.Attribute
	ComponentObject                       = customcomponents.ComponentObject
	CustomComponentsCustomForm            = customcomponents.CustomForm
	CustomComponentsCustomFormRow         = customcomponents.CustomFormRow
	CustomComponentsFirstColumnAttribute  = customcomponents.FirstColumnAttribute
	CustomComponentsOption                = customcomponents.Option
	Parent                                = customcomponents.Parent
	CustomComponentsSecondColumnAttribute = customcomponents.SecondColumnAttribute
	CustomComponentsTypeActual            = customcomponents.TypeActual
)

// NewCustomComponents binds the resource to a requester of your own.
func NewCustomComponents(requester Requester) *CustomComponents {
	return customcomponents.New(requester)
}

// CustomComponentsComponentListShape is the DTOShape reading a body as CustomComponentsComponentList, for the raw-verb call.
func CustomComponentsComponentListShape(decoded any) any {
	return customcomponents.ComponentListShape(decoded)
}

// CustomComponentsComponentRecordShape is the DTOShape reading a body as CustomComponentsComponentRecord, for the raw-verb call.
func CustomComponentsComponentRecordShape(decoded any) any {
	return customcomponents.ComponentRecordShape(decoded)
}

// ComponentByUuidRecordShape is the DTOShape reading a body as ComponentByUuidRecord, for the raw-verb call.
func ComponentByUuidRecordShape(decoded any) any {
	return customcomponents.ComponentByUuidRecordShape(decoded)
}

// ComponentObjectsRecordShape is the DTOShape reading a body as ComponentObjectsRecord, for the raw-verb call.
func ComponentObjectsRecordShape(decoded any) any {
	return customcomponents.ComponentObjectsRecordShape(decoded)
}

// ComponentObjectRecordShape is the DTOShape reading a body as ComponentObjectRecord, for the raw-verb call.
func ComponentObjectRecordShape(decoded any) any {
	return customcomponents.ComponentObjectRecordShape(decoded)
}

// ComponentObjectActivityRecordShape is the DTOShape reading a body as ComponentObjectActivityRecord, for the raw-verb call.
func ComponentObjectActivityRecordShape(decoded any) any {
	return customcomponents.ComponentObjectActivityRecordShape(decoded)
}

// ComponentObjectCustomObjectsRecordShape is the DTOShape reading a body as ComponentObjectCustomObjectsRecord, for the raw-verb call.
func ComponentObjectCustomObjectsRecordShape(decoded any) any {
	return customcomponents.ComponentObjectCustomObjectsRecordShape(decoded)
}

// The custom components path templates and body keys, for the raw-verb call.
const (
	CustomComponentsComponentPath     = customcomponents.ComponentPath
	ComponentByUUIDPath               = customcomponents.ComponentByUUIDPath
	CustomComponentsComponentByIDPath = customcomponents.ComponentByIDPath
	ComponentObjectsPath              = customcomponents.ComponentObjectsPath
	ComponentObjectPath               = customcomponents.ComponentObjectPath
	ComponentObjectActivityPath       = customcomponents.ComponentObjectActivityPath
	ComponentObjectCustomObjectsPath  = customcomponents.ComponentObjectCustomObjectsPath
	ComponentObjectPDFPath            = customcomponents.ComponentObjectPDFPath

	CustomComponentsKey          = customcomponents.CustomComponentsKey
	AccCusCompKey                = customcomponents.AccCusCompKey
	CustomComponentsComponentKey = customcomponents.ComponentKey
)

// CustomFeatures is the custom features endpoints bound to a client.
type CustomFeatures = customfeatures.CustomFeatures

// The custom features views, aliased so a type assertion works under either spelling.
type (
	CustomFeatureRecord      = customfeatures.CustomFeatureRecord
	Attribute                = customfeatures.Attribute
	Column1                  = customfeatures.Column1
	Column2                  = customfeatures.Column2
	Configuration            = customfeatures.Configuration
	CustomFeaturesCustomForm = customfeatures.CustomForm
	CustomFeaturesSettings   = customfeatures.Settings
)

// NewCustomFeatures binds the resource to a requester of your own.
func NewCustomFeatures(requester Requester) *CustomFeatures { return customfeatures.New(requester) }

// CustomFeatureRecordShape is the DTOShape reading a body as CustomFeatureRecord, for the raw-verb call.
func CustomFeatureRecordShape(decoded any) any {
	return customfeatures.CustomFeatureRecordShape(decoded)
}

// The custom features path templates and body keys, for the raw-verb call.
const (
	CustomFeaturePath               = customfeatures.CustomFeaturePath
	CustomFeaturesComponentByIDPath = customfeatures.ComponentByIDPath
	CustomFeatureByIDPath           = customfeatures.CustomFeatureByIDPath

	CustomFeaturesKey = customfeatures.CustomFeaturesKey
)

// Employees is the employees endpoints bound to a client.
type Employees = employees.Employees

// The employees views, aliased so a type assertion works under either spelling.
type (
	EmployeesPagination      = employees.Pagination
	EmployeesEmployeeRecord  = employees.EmployeeRecord
	EmployeesEmployeeList    = employees.EmployeeList
	EmployeesPaginationLinks = employees.PaginationLinks
	EmployeesContact         = employees.Contact
	EmployeesCreatedBy       = employees.CreatedBy
	EmployeesCustomForm      = employees.CustomForm
	EmployeesEmail           = employees.Email
	EmployeesLastUpdatedBy   = employees.LastUpdatedBy
	EmployeesName            = employees.Name
)

// NewEmployees binds the resource to a requester of your own.
func NewEmployees(requester Requester) *Employees { return employees.New(requester) }

// EmployeesEmployeeListShape is the DTOShape reading a body as EmployeesEmployeeList, for the raw-verb call.
func EmployeesEmployeeListShape(decoded any) any { return employees.EmployeeListShape(decoded) }

// EmployeesEmployeeRecordShape is the DTOShape reading a body as EmployeesEmployeeRecord, for the raw-verb call.
func EmployeesEmployeeRecordShape(decoded any) any { return employees.EmployeeRecordShape(decoded) }

// The employees path templates and body keys, for the raw-verb call.
const (
	EmployeesEmployeesPath = employees.EmployeesPath
	EmployeePath           = employees.EmployeePath

	EmployeesKey = employees.EmployeesKey
	EmployeeKey  = employees.EmployeeKey
)

// Events is the events endpoints bound to a client.
type Events = events.Events

// The events views, aliased so a type assertion works under either spelling.
type (
	EventsPagination      = events.Pagination
	EventRecord           = events.EventRecord
	EventList             = events.EventList
	EventsPaginationLinks = events.PaginationLinks
)

// NewEvents binds the resource to a requester of your own.
func NewEvents(requester Requester) *Events { return events.New(requester) }

// EventListShape is the DTOShape reading a body as EventList, for the raw-verb call.
func EventListShape(decoded any) any { return events.EventListShape(decoded) }

// EventRecordShape is the DTOShape reading a body as EventRecord, for the raw-verb call.
func EventRecordShape(decoded any) any { return events.EventRecordShape(decoded) }

// The events path templates and body keys, for the raw-verb call.
const (
	EventsPath       = events.EventsPath
	EventPath        = events.EventPath
	EventExecutePath = events.EventExecutePath
	EventRemovePath  = events.EventRemovePath

	EventsKey = events.EventsKey
	EventKey  = events.EventKey
)

// Express is the express endpoints bound to a client.
type Express = express.Express

// The express views, aliased so a type assertion works under either spelling.
type (
	ExpressAccount                  = express.Account
	ExpressActivityLog              = express.ActivityLog
	ExpressBillingPreferences       = express.BillingPreferences
	ExpressCurrency                 = express.Currency
	ExpressCustomForm               = express.CustomForm
	ExpressCreateRecord             = express.ExpressCreateRecord
	ExpressItem                     = express.Item
	ExpressItemAccountingCode       = express.ItemAccountingCode
	ExpressItemPriceSnapshot        = express.ItemPriceSnapshot
	ExpressItemPriceTax             = express.ItemPriceTax
	ExpressItemSaleTaxConfiguration = express.ItemSaleTaxConfiguration
	ExpressKpis                     = express.Kpis
	ExpressLine                     = express.Line
	ExpressPreOrderStockDetails     = express.PreOrderStockDetails
	ExpressPricingRule              = express.PricingRule
	ExpressProformaProperties       = express.ProformaProperties
	ExpressTaxCode                  = express.TaxCode
	ExpressTimeZone                 = express.TimeZone
)

// NewExpress binds the resource to a requester of your own.
func NewExpress(requester Requester) *Express { return express.New(requester) }

// ExpressExpressCreateRecordShape is the DTOShape reading a body as ExpressExpressCreateRecord, for the raw-verb call.
func ExpressExpressCreateRecordShape(decoded any) any {
	return express.ExpressCreateRecordShape(decoded)
}

// The express path templates and body keys, for the raw-verb call.
const (
	ExpressPath = express.ExpressPath
)

// ExternalDatabases is the external databases endpoints bound to a client.
type ExternalDatabases = externaldatabases.ExternalDatabases

// The external databases views, aliased so a type assertion works under either spelling.
type (
	ExternalDatabasesPagination               = externaldatabases.Pagination
	ExternalDatabasesRecord                   = externaldatabases.ExternalDatabasesRecord
	ExternalDatabaseCreateRecord              = externaldatabases.ExternalDatabaseCreateRecord
	ExternalDatabaseInsertRecordCreateRecord  = externaldatabases.ExternalDatabaseInsertRecordCreateRecord
	ExternalDatabaseSearchCreateRecord        = externaldatabases.ExternalDatabaseSearchCreateRecord
	ExternalDatabaseTableList                 = externaldatabases.ExternalDatabaseTableList
	ExternalDatabaseUpdateRecordsUpdateRecord = externaldatabases.ExternalDatabaseUpdateRecordsUpdateRecord
	ExternalDatabasesPaginationLinks          = externaldatabases.PaginationLinks
)

// NewExternalDatabases binds the resource to a requester of your own.
func NewExternalDatabases(requester Requester) *ExternalDatabases {
	return externaldatabases.New(requester)
}

// ExternalDatabasesRecordShape is the DTOShape reading a body as ExternalDatabasesRecord, for the raw-verb call.
func ExternalDatabasesRecordShape(decoded any) any {
	return externaldatabases.ExternalDatabasesRecordShape(decoded)
}

// ExternalDatabasesExternalDatabaseCreateRecordShape is the DTOShape reading a body as ExternalDatabasesExternalDatabaseCreateRecord, for the raw-verb call.
func ExternalDatabasesExternalDatabaseCreateRecordShape(decoded any) any {
	return externaldatabases.ExternalDatabaseCreateRecordShape(decoded)
}

// ExternalDatabasesExternalDatabaseInsertRecordCreateRecordShape is the DTOShape reading a body as ExternalDatabasesExternalDatabaseInsertRecordCreateRecord, for the raw-verb call.
func ExternalDatabasesExternalDatabaseInsertRecordCreateRecordShape(decoded any) any {
	return externaldatabases.ExternalDatabaseInsertRecordCreateRecordShape(decoded)
}

// ExternalDatabasesExternalDatabaseSearchCreateRecordShape is the DTOShape reading a body as ExternalDatabasesExternalDatabaseSearchCreateRecord, for the raw-verb call.
func ExternalDatabasesExternalDatabaseSearchCreateRecordShape(decoded any) any {
	return externaldatabases.ExternalDatabaseSearchCreateRecordShape(decoded)
}

// ExternalDatabasesExternalDatabaseTableListShape is the DTOShape reading a body as ExternalDatabasesExternalDatabaseTableList, for the raw-verb call.
func ExternalDatabasesExternalDatabaseTableListShape(decoded any) any {
	return externaldatabases.ExternalDatabaseTableListShape(decoded)
}

// ExternalDatabasesExternalDatabaseUpdateRecordsUpdateRecordShape is the DTOShape reading a body as ExternalDatabasesExternalDatabaseUpdateRecordsUpdateRecord, for the raw-verb call.
func ExternalDatabasesExternalDatabaseUpdateRecordsUpdateRecordShape(decoded any) any {
	return externaldatabases.ExternalDatabaseUpdateRecordsUpdateRecordShape(decoded)
}

// The external databases path templates and body keys, for the raw-verb call.
const (
	ExternalDatabasesPath             = externaldatabases.ExternalDatabasesPath
	ExternalDatabasePath              = externaldatabases.ExternalDatabasePath
	ExternalDatabaseInsertRecordsPath = externaldatabases.ExternalDatabaseInsertRecordsPath
	ExternalDatabaseSearchPath        = externaldatabases.ExternalDatabaseSearchPath
	ExternalDatabaseUpdateRecordsPath = externaldatabases.ExternalDatabaseUpdateRecordsPath
	ExternalDatabaseTablePath         = externaldatabases.ExternalDatabaseTablePath

	ExternalDatabasesKey = externaldatabases.ExternalDatabasesKey
)

// GiftCertificates is the gift certificates endpoints bound to a client.
type GiftCertificates = giftcertificates.GiftCertificates

// The gift certificates views, aliased so a type assertion works under either spelling.
type (
	GiftCertificatesPagination       = giftcertificates.Pagination
	GiftCertificateRecord            = giftcertificates.GiftCertificateRecord
	GiftCertificateList              = giftcertificates.GiftCertificateList
	GiftCertificateTransactionRecord = giftcertificates.GiftCertificateTransactionRecord
	GiftCertificateTransactionList   = giftcertificates.GiftCertificateTransactionList
	GiftCertificateAllocationsRecord = giftcertificates.GiftCertificateAllocationsRecord
	GiftCertificatesActivityLog      = giftcertificates.ActivityLog
	GiftCertificatesPaginationLinks  = giftcertificates.PaginationLinks
	GiftCertificatesAccount          = giftcertificates.Account
	GiftCertificatesAccountingCode   = giftcertificates.AccountingCode
	Allocation                       = giftcertificates.Allocation
	GiftCertificatesCurrency         = giftcertificates.Currency
	Transaction                      = giftcertificates.Transaction
)

// NewGiftCertificates binds the resource to a requester of your own.
func NewGiftCertificates(requester Requester) *GiftCertificates {
	return giftcertificates.New(requester)
}

// GiftCertificateListShape is the DTOShape reading a body as GiftCertificateList, for the raw-verb call.
func GiftCertificateListShape(decoded any) any {
	return giftcertificates.GiftCertificateListShape(decoded)
}

// GiftCertificateRecordShape is the DTOShape reading a body as GiftCertificateRecord, for the raw-verb call.
func GiftCertificateRecordShape(decoded any) any {
	return giftcertificates.GiftCertificateRecordShape(decoded)
}

// GiftCertificateAllocationsRecordShape is the DTOShape reading a body as GiftCertificateAllocationsRecord, for the raw-verb call.
func GiftCertificateAllocationsRecordShape(decoded any) any {
	return giftcertificates.GiftCertificateAllocationsRecordShape(decoded)
}

// GiftCertificateTransactionListShape is the DTOShape reading a body as GiftCertificateTransactionList, for the raw-verb call.
func GiftCertificateTransactionListShape(decoded any) any {
	return giftcertificates.GiftCertificateTransactionListShape(decoded)
}

// The gift certificates path templates and body keys, for the raw-verb call.
const (
	GiftCertificatesPath            = giftcertificates.GiftCertificatesPath
	GiftCertificatePath             = giftcertificates.GiftCertificatePath
	GiftCertificateAllocationsPath  = giftcertificates.GiftCertificateAllocationsPath
	GiftCertificateTransactionsPath = giftcertificates.GiftCertificateTransactionsPath

	GiftCertificatesKey = giftcertificates.GiftCertificatesKey
	GiftCertificateKey  = giftcertificates.GiftCertificateKey
)

// IntegrationConnections is the integration connections endpoints bound to a client.
type IntegrationConnections = integrationconnections.IntegrationConnections

// The integration connections views, aliased so a type assertion works under either spelling.
type (
	IntegrationConnectionsPagination      = integrationconnections.Pagination
	IntegrationConnectionsRecord          = integrationconnections.IntegrationConnectionsRecord
	IntegrationConnectionsPaginationLinks = integrationconnections.PaginationLinks
)

// NewIntegrationConnections binds the resource to a requester of your own.
func NewIntegrationConnections(requester Requester) *IntegrationConnections {
	return integrationconnections.New(requester)
}

// IntegrationConnectionsRecordShape is the DTOShape reading a body as IntegrationConnectionsRecord, for the raw-verb call.
func IntegrationConnectionsRecordShape(decoded any) any {
	return integrationconnections.IntegrationConnectionsRecordShape(decoded)
}

// The integration connections path templates and body keys, for the raw-verb call.
const (
	IntegrationConnectionsPath = integrationconnections.IntegrationConnectionsPath
	IntegrationConnectionPath  = integrationconnections.IntegrationConnectionPath

	IntegrationConnectionsKey = integrationconnections.IntegrationConnectionsKey
)

// InventoryAdjustments is the inventory adjustments endpoints bound to a client.
type InventoryAdjustments = inventoryadjustments.InventoryAdjustments

// The inventory adjustments views, aliased so a type assertion works under either spelling.
type (
	InventoryAdjustmentsPagination              = inventoryadjustments.Pagination
	InventoryAdjustmentRecord                   = inventoryadjustments.InventoryAdjustmentRecord
	InventoryAdjustmentsInventoryAdjustmentList = inventoryadjustments.InventoryAdjustmentList
	InventoryAdjustmentsPaginationLinks         = inventoryadjustments.PaginationLinks
)

// NewInventoryAdjustments binds the resource to a requester of your own.
func NewInventoryAdjustments(requester Requester) *InventoryAdjustments {
	return inventoryadjustments.New(requester)
}

// InventoryAdjustmentsInventoryAdjustmentListShape is the DTOShape reading a body as InventoryAdjustmentsInventoryAdjustmentList, for the raw-verb call.
func InventoryAdjustmentsInventoryAdjustmentListShape(decoded any) any {
	return inventoryadjustments.InventoryAdjustmentListShape(decoded)
}

// InventoryAdjustmentRecordShape is the DTOShape reading a body as InventoryAdjustmentRecord, for the raw-verb call.
func InventoryAdjustmentRecordShape(decoded any) any {
	return inventoryadjustments.InventoryAdjustmentRecordShape(decoded)
}

// The inventory adjustments path templates and body keys, for the raw-verb call.
const (
	InventoryAdjustmentsPath = inventoryadjustments.InventoryAdjustmentsPath
	InventoryAdjustmentPath  = inventoryadjustments.InventoryAdjustmentPath

	InventoryAdjustmentsKey                    = inventoryadjustments.InventoryAdjustmentsKey
	InventoryAdjustmentsInventoryAdjustmentKey = inventoryadjustments.InventoryAdjustmentKey
)

// Invoices is the invoices endpoints bound to a client.
type Invoices = invoices.Invoices

// The invoices views, aliased so a type assertion works under either spelling.
type (
	InvoicesPagination                          = invoices.Pagination
	InvoiceRecord                               = invoices.InvoiceRecord
	InvoiceList                                 = invoices.InvoiceList
	InvoiceNoteList                             = invoices.InvoiceNoteList
	InvoicesAccount                             = invoices.Account
	InvoicesAttachment                          = invoices.Attachment
	InvoicesCreatedBy                           = invoices.CreatedBy
	Fulfillment                                 = invoices.Fulfillment
	InvoicesInvoice                             = invoices.Invoice
	InvoiceItemFulfillmentList                  = invoices.InvoiceItemFulfillmentList
	InvoiceItemFulfillmentRecord                = invoices.InvoiceItemFulfillmentRecord
	InvoiceItemFulfillmentView                  = invoices.InvoiceItemFulfillmentView
	InvoiceNoteFileList                         = invoices.InvoiceNoteFileList
	InvoiceNoteFileRecord                       = invoices.InvoiceNoteFileRecord
	InvoiceNoteFileView                         = invoices.InvoiceNoteFileView
	InvoiceNoteView                             = invoices.InvoiceNoteView
	InvoicePDFView                              = invoices.InvoicePDFView
	InvoiceReturnMerchandiseAuthorisationList   = invoices.InvoiceReturnMerchandiseAuthorisationList
	InvoiceReturnMerchandiseAuthorisationRecord = invoices.InvoiceReturnMerchandiseAuthorisationRecord
	InvoiceReturnMerchandiseAuthorisationView   = invoices.InvoiceReturnMerchandiseAuthorisationView
	InvoicesItem                                = invoices.Item
	InvoicesLastUpdatedBy                       = invoices.LastUpdatedBy
	InvoicesOrder                               = invoices.Order
	InvoicesPDF                                 = invoices.PDF
	InvoicesPaginationLinks                     = invoices.PaginationLinks
	ReceiveRma                                  = invoices.ReceiveRma
	SaleInvoice                                 = invoices.SaleInvoice
	InvoicesWatcher                             = invoices.Watcher
	InvoicesAccountingCode                      = invoices.AccountingCode
	InvoicesActivityLog                         = invoices.ActivityLog
	InvoicesCurrency                            = invoices.Currency
	InvoicesCustomForm                          = invoices.CustomForm
	InvoiceNote                                 = invoices.InvoiceNote
	InvoicesLastPayment                         = invoices.LastPayment
	InvoicesLine                                = invoices.Line
	InvoicesTax                                 = invoices.Tax
)

// NewInvoices binds the resource to a requester of your own.
func NewInvoices(requester Requester) *Invoices { return invoices.New(requester) }

// InvoiceListShape is the DTOShape reading a body as InvoiceList, for the raw-verb call.
func InvoiceListShape(decoded any) any { return invoices.InvoiceListShape(decoded) }

// InvoiceRecordShape is the DTOShape reading a body as InvoiceRecord, for the raw-verb call.
func InvoiceRecordShape(decoded any) any { return invoices.InvoiceRecordShape(decoded) }

// InvoiceNoteListShape is the DTOShape reading a body as InvoiceNoteList, for the raw-verb call.
func InvoiceNoteListShape(decoded any) any { return invoices.InvoiceNoteListShape(decoded) }

// InvoicesInvoiceItemFulfillmentListShape is the DTOShape reading a body as InvoicesInvoiceItemFulfillmentList, for the raw-verb call.
func InvoicesInvoiceItemFulfillmentListShape(decoded any) any {
	return invoices.InvoiceItemFulfillmentListShape(decoded)
}

// InvoicesInvoiceItemFulfillmentViewShape is the DTOShape reading a body as InvoicesInvoiceItemFulfillmentView, for the raw-verb call.
func InvoicesInvoiceItemFulfillmentViewShape(decoded any) any {
	return invoices.InvoiceItemFulfillmentViewShape(decoded)
}

// InvoicesInvoiceNoteFileListShape is the DTOShape reading a body as InvoicesInvoiceNoteFileList, for the raw-verb call.
func InvoicesInvoiceNoteFileListShape(decoded any) any {
	return invoices.InvoiceNoteFileListShape(decoded)
}

// InvoicesInvoiceNoteFileViewShape is the DTOShape reading a body as InvoicesInvoiceNoteFileView, for the raw-verb call.
func InvoicesInvoiceNoteFileViewShape(decoded any) any {
	return invoices.InvoiceNoteFileViewShape(decoded)
}

// InvoicesInvoiceNoteViewShape is the DTOShape reading a body as InvoicesInvoiceNoteView, for the raw-verb call.
func InvoicesInvoiceNoteViewShape(decoded any) any { return invoices.InvoiceNoteViewShape(decoded) }

// InvoicesInvoicePDFViewShape is the DTOShape reading a body as InvoicesInvoicePDFView, for the raw-verb call.
func InvoicesInvoicePDFViewShape(decoded any) any { return invoices.InvoicePDFViewShape(decoded) }

// InvoicesInvoiceReturnMerchandiseAuthorisationListShape is the DTOShape reading a body as InvoicesInvoiceReturnMerchandiseAuthorisationList, for the raw-verb call.
func InvoicesInvoiceReturnMerchandiseAuthorisationListShape(decoded any) any {
	return invoices.InvoiceReturnMerchandiseAuthorisationListShape(decoded)
}

// InvoicesInvoiceReturnMerchandiseAuthorisationViewShape is the DTOShape reading a body as InvoicesInvoiceReturnMerchandiseAuthorisationView, for the raw-verb call.
func InvoicesInvoiceReturnMerchandiseAuthorisationViewShape(decoded any) any {
	return invoices.InvoiceReturnMerchandiseAuthorisationViewShape(decoded)
}

// The invoices path templates and body keys, for the raw-verb call.
const (
	InvoicesPath                                       = invoices.InvoicesPath
	InvoicePath                                        = invoices.InvoicePath
	InvoiceAmendPath                                   = invoices.InvoiceAmendPath
	InvoiceCancelPath                                  = invoices.InvoiceCancelPath
	InvoiceInformationPath                             = invoices.InvoiceInformationPath
	InvoicesInvoiceItemFulfillmentsPath                = invoices.InvoiceItemFulfillmentsPath
	InvoiceNotesPath                                   = invoices.InvoiceNotesPath
	InvoicePDFPath                                     = invoices.InvoicePDFPath
	InvoiceReactivatePath                              = invoices.InvoiceReactivatePath
	InvoicesInvoiceReturnMerchandiseAuthorisationsPath = invoices.InvoiceReturnMerchandiseAuthorisationsPath
	OrderInvoicesPath                                  = invoices.OrderInvoicesPath
	InvoicesInvoiceItemFulfillmentPath                 = invoices.InvoiceItemFulfillmentPath
	InvoiceNotePath                                    = invoices.InvoiceNotePath
	InvoicesInvoiceReturnMerchandiseAuthorisationPath  = invoices.InvoiceReturnMerchandiseAuthorisationPath
	InvoiceNoteFilesPath                               = invoices.InvoiceNoteFilesPath
	InvoiceNoteFilePath                                = invoices.InvoiceNoteFilePath

	InvoicesKey        = invoices.InvoicesKey
	InvoicesInvoiceKey = invoices.InvoiceKey
)

// ItemFulfillments is the item fulfillments endpoints bound to a client.
type ItemFulfillments = itemfulfillments.ItemFulfillments

// The item fulfillments views, aliased so a type assertion works under either spelling.
type (
	ItemFulfillmentsPagination      = itemfulfillments.Pagination
	ItemFulfillmentRecord           = itemfulfillments.ItemFulfillmentRecord
	ItemFulfillmentList             = itemfulfillments.ItemFulfillmentList
	ItemFulfillmentReadRecord       = itemfulfillments.ItemFulfillmentReadRecord
	ItemFulfillmentsRecord          = itemfulfillments.ItemFulfillmentsRecord
	ItemFulfillmentUpdatedRecord    = itemfulfillments.ItemFulfillmentUpdatedRecord
	ItemFulfillment                 = itemfulfillments.ItemFulfillment
	ItemFulfillmentNoteCreateRecord = itemfulfillments.ItemFulfillmentNoteCreateRecord
	ItemFulfillmentsPaginationLinks = itemfulfillments.PaginationLinks
	ItemFulfillmentsAccount         = itemfulfillments.Account
	ItemFulfillmentsCreatedBy       = itemfulfillments.CreatedBy
	ItemFulfillmentsFulfillment     = itemfulfillments.Fulfillment
	ItemFulfillmentsInvoice         = itemfulfillments.Invoice
	ItemFulfillmentsItem            = itemfulfillments.Item
	ItemFulfillmentsLastUpdatedBy   = itemfulfillments.LastUpdatedBy
)

// NewItemFulfillments binds the resource to a requester of your own.
func NewItemFulfillments(requester Requester) *ItemFulfillments {
	return itemfulfillments.New(requester)
}

// ItemFulfillmentListShape is the DTOShape reading a body as ItemFulfillmentList, for the raw-verb call.
func ItemFulfillmentListShape(decoded any) any {
	return itemfulfillments.ItemFulfillmentListShape(decoded)
}

// ItemFulfillmentsRecordShape is the DTOShape reading a body as ItemFulfillmentsRecord, for the raw-verb call.
func ItemFulfillmentsRecordShape(decoded any) any {
	return itemfulfillments.ItemFulfillmentsRecordShape(decoded)
}

// ItemFulfillmentReadRecordShape is the DTOShape reading a body as ItemFulfillmentReadRecord, for the raw-verb call.
func ItemFulfillmentReadRecordShape(decoded any) any {
	return itemfulfillments.ItemFulfillmentReadRecordShape(decoded)
}

// ItemFulfillmentUpdatedRecordShape is the DTOShape reading a body as ItemFulfillmentUpdatedRecord, for the raw-verb call.
func ItemFulfillmentUpdatedRecordShape(decoded any) any {
	return itemfulfillments.ItemFulfillmentUpdatedRecordShape(decoded)
}

// ItemFulfillmentsItemFulfillmentNoteCreateRecordShape is the DTOShape reading a body as ItemFulfillmentsItemFulfillmentNoteCreateRecord, for the raw-verb call.
func ItemFulfillmentsItemFulfillmentNoteCreateRecordShape(decoded any) any {
	return itemfulfillments.ItemFulfillmentNoteCreateRecordShape(decoded)
}

// The item fulfillments path templates and body keys, for the raw-verb call.
const (
	ItemFulfillmentsPath     = itemfulfillments.ItemFulfillmentsPath
	ItemFulfillmentPath      = itemfulfillments.ItemFulfillmentPath
	ItemFulfillmentNotesPath = itemfulfillments.ItemFulfillmentNotesPath

	ItemFulfillmentsKey        = itemfulfillments.ItemFulfillmentsKey
	ItemFulfillmentsInvoiceKey = itemfulfillments.InvoiceKey
	ItemFulfillmentKey         = itemfulfillments.ItemFulfillmentKey
)

// ItemReceipts is the item receipts endpoints bound to a client.
type ItemReceipts = itemreceipts.ItemReceipts

// The item receipts views, aliased so a type assertion works under either spelling.
type (
	ItemReceiptsPagination          = itemreceipts.Pagination
	ItemReceiptRecord               = itemreceipts.ItemReceiptRecord
	ItemReceiptList                 = itemreceipts.ItemReceiptList
	ItemReceiptsRecord              = itemreceipts.ItemReceiptsRecord
	ItemReceiptUpdatedRecord        = itemreceipts.ItemReceiptUpdatedRecord
	PurchaseOrderItemReceiptsRecord = itemreceipts.PurchaseOrderItemReceiptsRecord
	ItemReceiptsCreatedBy           = itemreceipts.CreatedBy
	ItemReceiptsItem                = itemreceipts.Item
	ItemReceiptsLastUpdatedBy       = itemreceipts.LastUpdatedBy
	ItemReceiptsLine                = itemreceipts.Line
	ItemReceiptsPaginationLinks     = itemreceipts.PaginationLinks
	ItemReceiptsPurchaseOrder       = itemreceipts.PurchaseOrder
	PurchaseOrderLine               = itemreceipts.PurchaseOrderLine
	ItemReceiptsUom                 = itemreceipts.Uom
	ItemReceiptsWarehouse           = itemreceipts.Warehouse
)

// NewItemReceipts binds the resource to a requester of your own.
func NewItemReceipts(requester Requester) *ItemReceipts { return itemreceipts.New(requester) }

// ItemReceiptListShape is the DTOShape reading a body as ItemReceiptList, for the raw-verb call.
func ItemReceiptListShape(decoded any) any { return itemreceipts.ItemReceiptListShape(decoded) }

// ItemReceiptsRecordShape is the DTOShape reading a body as ItemReceiptsRecord, for the raw-verb call.
func ItemReceiptsRecordShape(decoded any) any { return itemreceipts.ItemReceiptsRecordShape(decoded) }

// ItemReceiptUpdatedRecordShape is the DTOShape reading a body as ItemReceiptUpdatedRecord, for the raw-verb call.
func ItemReceiptUpdatedRecordShape(decoded any) any {
	return itemreceipts.ItemReceiptUpdatedRecordShape(decoded)
}

// ItemReceiptRecordShape is the DTOShape reading a body as ItemReceiptRecord, for the raw-verb call.
func ItemReceiptRecordShape(decoded any) any { return itemreceipts.ItemReceiptRecordShape(decoded) }

// PurchaseOrderItemReceiptsRecordShape is the DTOShape reading a body as PurchaseOrderItemReceiptsRecord, for the raw-verb call.
func PurchaseOrderItemReceiptsRecordShape(decoded any) any {
	return itemreceipts.PurchaseOrderItemReceiptsRecordShape(decoded)
}

// The item receipts path templates and body keys, for the raw-verb call.
const (
	ItemReceiptsPath              = itemreceipts.ItemReceiptsPath
	ItemReceiptPath               = itemreceipts.ItemReceiptPath
	PurchaseOrderItemReceiptsPath = itemreceipts.PurchaseOrderItemReceiptsPath

	ItemReceiptsKey              = itemreceipts.ItemReceiptsKey
	ItemReceiptKey               = itemreceipts.ItemReceiptKey
	ItemReceiptsPurchaseOrderKey = itemreceipts.PurchaseOrderKey
)

// Items is the items endpoints bound to a client.
type Items = items.Items

// The items views, aliased so a type assertion works under either spelling.
type (
	ItemsReference               = items.Reference
	ItemsUser                    = items.User
	ItemsCustomForm              = items.CustomForm
	ItemCode                     = items.ItemCode
	ItemCurrency                 = items.ItemCurrency
	BundleDetail                 = items.BundleDetail
	FamilyDetail                 = items.FamilyDetail
	VariationOption              = items.VariationOption
	PackageDetail                = items.PackageDetail
	TaxConfiguration             = items.TaxConfiguration
	Pricing                      = items.Pricing
	SaleCharge                   = items.SaleCharge
	SaleBlock                    = items.SaleBlock
	PurchaseBlock                = items.PurchaseBlock
	InventoryProperties          = items.InventoryProperties
	WarehouseInventory           = items.WarehouseInventory
	InventoriesBlock             = items.InventoriesBlock
	ItemsNoteAttachment          = items.NoteAttachment
	NoteRecord                   = items.NoteRecord
	ItemRecord                   = items.ItemRecord
	ItemList                     = items.ItemList
	ItemsPagination              = items.Pagination
	ItemSaleView                 = items.ItemSaleView
	ItemPurchaseView             = items.ItemPurchaseView
	ItemInventoriesView          = items.ItemInventoriesView
	InventoryAdjustment          = items.InventoryAdjustment
	ItemsInventoryAdjustmentList = items.InventoryAdjustmentList
	ItemsNoteList                = items.NoteList
	ItemsNoteFileList            = items.NoteFileList
	ItemCreateBody               = items.ItemCreateBody
	ItemInformationUpdate        = items.ItemInformationUpdate
	ItemSaleWrite                = items.ItemSaleWrite
	ItemPurchaseWrite            = items.ItemPurchaseWrite
	ItemInventoriesWrite         = items.ItemInventoriesWrite
	InventoryAddBody             = items.InventoryAddBody
	ItemVariationsWrite          = items.ItemVariationsWrite
	ItemBundleDetailsWrite       = items.ItemBundleDetailsWrite
	ItemFamilyDetailsWrite       = items.ItemFamilyDetailsWrite
	ItemPackageDetailsWrite      = items.ItemPackageDetailsWrite
	ItemImageUpload              = items.ItemImageUpload
	Height                       = items.Height
	ItemNoteFileView             = items.ItemNoteFileView
	Length                       = items.Length
	PreorderConfiguration        = items.PreorderConfiguration
	Supplier                     = items.Supplier
	SupplierPricing              = items.SupplierPricing
	SupplierTaxConfiguration     = items.SupplierTaxConfiguration
	ItemsTaxCode                 = items.TaxCode
	Weight                       = items.Weight
	Width                        = items.Width
	BillingPeriod                = items.BillingPeriod
	ItemsCustomAttribute         = items.CustomAttribute
	ItemsItem                    = items.Item
	ItemsLastUpdatedBy           = items.LastUpdatedBy
	ItemsPaginationLinks         = items.PaginationLinks
	PaymentProperty              = items.PaymentProperty
	ItemsPricingLevel            = items.PricingLevel
	PricingModule                = items.PricingModule
	Property                     = items.Property
	PurchaseProperty             = items.PurchaseProperty
	Value                        = items.Value
	ItemsWarehouse               = items.Warehouse
	ItemsWatcher                 = items.Watcher
)

// NewItems binds the resource to a requester of your own.
func NewItems(requester Requester) *Items { return items.New(requester) }

// ItemListShape is the DTOShape reading a body as ItemList, for the raw-verb call.
func ItemListShape(decoded any) any { return items.ItemListShape(decoded) }

// ItemRecordShape is the DTOShape reading a body as ItemRecord, for the raw-verb call.
func ItemRecordShape(decoded any) any { return items.ItemRecordShape(decoded) }

// ItemSaleViewShape is the DTOShape reading a body as ItemSaleView, for the raw-verb call.
func ItemSaleViewShape(decoded any) any { return items.ItemSaleViewShape(decoded) }

// ItemPurchaseViewShape is the DTOShape reading a body as ItemPurchaseView, for the raw-verb call.
func ItemPurchaseViewShape(decoded any) any { return items.ItemPurchaseViewShape(decoded) }

// ItemInventoriesViewShape is the DTOShape reading a body as ItemInventoriesView, for the raw-verb call.
func ItemInventoriesViewShape(decoded any) any { return items.ItemInventoriesViewShape(decoded) }

// ItemsInventoryAdjustmentListShape is the DTOShape reading a body as ItemsInventoryAdjustmentList, for the raw-verb call.
func ItemsInventoryAdjustmentListShape(decoded any) any {
	return items.InventoryAdjustmentListShape(decoded)
}

// ItemsNoteListShape is the DTOShape reading a body as ItemsNoteList, for the raw-verb call.
func ItemsNoteListShape(decoded any) any { return items.NoteListShape(decoded) }

// NoteRecordShape is the DTOShape reading a body as NoteRecord, for the raw-verb call.
func NoteRecordShape(decoded any) any { return items.NoteRecordShape(decoded) }

// ItemsNoteFileListShape is the DTOShape reading a body as ItemsNoteFileList, for the raw-verb call.
func ItemsNoteFileListShape(decoded any) any { return items.NoteFileListShape(decoded) }

// ItemsQueryParams returns a fresh copy of the names the listing accepts.
func ItemsQueryParams() []string { return items.ItemsQueryParams() }

// ItemsItemNoteFileViewShape is the DTOShape reading a body as ItemsItemNoteFileView, for the raw-verb call.
func ItemsItemNoteFileViewShape(decoded any) any { return items.ItemNoteFileViewShape(decoded) }

// The items path templates and body keys, for the raw-verb call.
const (
	ItemsPath                         = items.ItemsPath
	ItemPath                          = items.ItemPath
	ItemInformationPath               = items.ItemInformationPath
	ItemDeactivatePath                = items.ItemDeactivatePath
	ItemReactivatePath                = items.ItemReactivatePath
	ItemImagePath                     = items.ItemImagePath
	ItemSalePath                      = items.ItemSalePath
	ItemPurchasePath                  = items.ItemPurchasePath
	ItemInventoriesPath               = items.ItemInventoriesPath
	ItemInventoryAddPath              = items.ItemInventoryAddPath
	ItemsItemInventoryAdjustmentsPath = items.ItemInventoryAdjustmentsPath
	ItemWarehouseInventoriesPath      = items.ItemWarehouseInventoriesPath
	ItemBundleDetailsPath             = items.ItemBundleDetailsPath
	ItemVariationsPath                = items.ItemVariationsPath
	ItemFamilyDetailsPath             = items.ItemFamilyDetailsPath
	ItemPackageDetailsPath            = items.ItemPackageDetailsPath
	ItemPackageDetailPath             = items.ItemPackageDetailPath
	ItemNotesPath                     = items.ItemNotesPath
	ItemNotePath                      = items.ItemNotePath
	ItemNoteFilesPath                 = items.ItemNoteFilesPath
	ItemNoteFilePath                  = items.ItemNoteFilePath

	ItemsKey                    = items.ItemsKey
	ItemsInventoryAdjustmentKey = items.InventoryAdjustmentKey
	ItemKey                     = items.ItemKey
	ItemsNoteKey                = items.NoteKey
)

// Labours is the labours endpoints bound to a client.
type Labours = labours.Labours

// The labours views, aliased so a type assertion works under either spelling.
type (
	LaboursPagination               = labours.Pagination
	LaboursLabourRecord             = labours.LabourRecord
	LaboursLabourList               = labours.LabourList
	LabourProfileRecord             = labours.LabourProfileRecord
	LabourProfileList               = labours.LabourProfileList
	LabourAvailabilityRecord        = labours.LabourAvailabilityRecord
	LabourAvailabilityList          = labours.LabourAvailabilityList
	LaboursAccount                  = labours.Account
	Availability                    = labours.Availability
	LaboursBillingPreferences       = labours.BillingPreferences
	CommunicationPreference         = labours.CommunicationPreference
	LaboursCreatedBy                = labours.CreatedBy
	LaboursCurrency                 = labours.Currency
	LaboursCustomForm               = labours.CustomForm
	Hour                            = labours.Hour
	LaboursLabour                   = labours.Labour
	LabourOrderList                 = labours.LabourOrderList
	LabourOrderRecord               = labours.LabourOrderRecord
	LabourProfileReadRecord         = labours.LabourProfileReadRecord
	LaboursLastUpdatedBy            = labours.LastUpdatedBy
	LaboursPaginationLinks          = labours.PaginationLinks
	Price                           = labours.Price
	PricingLevel                    = labours.PricingLevel
	LaboursProformaProperties       = labours.ProformaProperties
	Session                         = labours.Session
	LaboursTax                      = labours.Tax
	LaboursTimeZone                 = labours.TimeZone
	WorkingHour                     = labours.WorkingHour
	LaboursFax                      = labours.Fax
	LaboursLabourContact            = labours.LabourContact
	LaboursLabourCustomAvailability = labours.LabourCustomAvailability
	LaboursLabourProfile            = labours.LabourProfile
	LaboursMobile                   = labours.Mobile
	LaboursPhone                    = labours.Phone
	LaboursScheduleList             = labours.ScheduleList
	LaboursSessionList              = labours.SessionList
)

// NewLabours binds the resource to a requester of your own.
func NewLabours(requester Requester) *Labours { return labours.New(requester) }

// LabourAvailabilityListShape is the DTOShape reading a body as LabourAvailabilityList, for the raw-verb call.
func LabourAvailabilityListShape(decoded any) any {
	return labours.LabourAvailabilityListShape(decoded)
}

// LabourProfileListShape is the DTOShape reading a body as LabourProfileList, for the raw-verb call.
func LabourProfileListShape(decoded any) any { return labours.LabourProfileListShape(decoded) }

// LaboursLabourListShape is the DTOShape reading a body as LaboursLabourList, for the raw-verb call.
func LaboursLabourListShape(decoded any) any { return labours.LabourListShape(decoded) }

// LaboursLabourRecordShape is the DTOShape reading a body as LaboursLabourRecord, for the raw-verb call.
func LaboursLabourRecordShape(decoded any) any { return labours.LabourRecordShape(decoded) }

// LaboursLabourOrderListShape is the DTOShape reading a body as LaboursLabourOrderList, for the raw-verb call.
func LaboursLabourOrderListShape(decoded any) any { return labours.LabourOrderListShape(decoded) }

// LaboursLabourProfileReadRecordShape is the DTOShape reading a body as LaboursLabourProfileReadRecord, for the raw-verb call.
func LaboursLabourProfileReadRecordShape(decoded any) any {
	return labours.LabourProfileReadRecordShape(decoded)
}

// The labours path templates and body keys, for the raw-verb call.
const (
	LabourAvailabilityPath         = labours.LabourAvailabilityPath
	LabourProfilesPath             = labours.LabourProfilesPath
	LaboursLaboursPath             = labours.LaboursPath
	LabourProfilePath              = labours.LabourProfilePath
	LaboursAvailableForBookingPath = labours.LaboursAvailableForBookingPath
	LabourPath                     = labours.LabourPath
	LabourCancelPath               = labours.LabourCancelPath
	LabourReactivatePath           = labours.LabourReactivatePath
	LabourOrdersPath               = labours.LabourOrdersPath

	LaboursKey = labours.LaboursKey
	LabourKey  = labours.LabourKey
)

// Milestones is the milestones endpoints bound to a client.
type Milestones = milestones.Milestones

// The milestones views, aliased so a type assertion works under either spelling.
type (
	MilestonesPagination       = milestones.Pagination
	MilestonesMilestonesRecord = milestones.MilestonesRecord
	MilestonesLabourByIDList   = milestones.MilestonesLabourByIDList
	MilestonesPaginationLinks  = milestones.PaginationLinks
)

// NewMilestones binds the resource to a requester of your own.
func NewMilestones(requester Requester) *Milestones { return milestones.New(requester) }

// MilestonesMilestonesRecordShape is the DTOShape reading a body as MilestonesMilestonesRecord, for the raw-verb call.
func MilestonesMilestonesRecordShape(decoded any) any {
	return milestones.MilestonesRecordShape(decoded)
}

// MilestonesMilestonesLabourByIDListShape is the DTOShape reading a body as MilestonesMilestonesLabourByIDList, for the raw-verb call.
func MilestonesMilestonesLabourByIDListShape(decoded any) any {
	return milestones.MilestonesLabourByIDListShape(decoded)
}

// The milestones path templates and body keys, for the raw-verb call.
const (
	MilestonesMilestonesPath  = milestones.MilestonesPath
	MilestonesAccountByIDPath = milestones.MilestonesAccountByIDPath
	MilestonesLabourByIDPath  = milestones.MilestonesLabourByIDPath

	MilestonesKey = milestones.MilestonesKey
)

// NoteGroups is the note groups endpoints bound to a client.
type NoteGroups = notegroups.NoteGroups

// The note groups views, aliased so a type assertion works under either spelling.
type (
	NoteGroupsPagination      = notegroups.Pagination
	NoteGroupRecord           = notegroups.NoteGroupRecord
	NoteGroupList             = notegroups.NoteGroupList
	NoteGroupNotesRecord      = notegroups.NoteGroupNotesRecord
	NoteGroupsPaginationLinks = notegroups.PaginationLinks
)

// NewNoteGroups binds the resource to a requester of your own.
func NewNoteGroups(requester Requester) *NoteGroups { return notegroups.New(requester) }

// NoteGroupListShape is the DTOShape reading a body as NoteGroupList, for the raw-verb call.
func NoteGroupListShape(decoded any) any { return notegroups.NoteGroupListShape(decoded) }

// NoteGroupRecordShape is the DTOShape reading a body as NoteGroupRecord, for the raw-verb call.
func NoteGroupRecordShape(decoded any) any { return notegroups.NoteGroupRecordShape(decoded) }

// NoteGroupNotesRecordShape is the DTOShape reading a body as NoteGroupNotesRecord, for the raw-verb call.
func NoteGroupNotesRecordShape(decoded any) any { return notegroups.NoteGroupNotesRecordShape(decoded) }

// The note groups path templates and body keys, for the raw-verb call.
const (
	NoteGroupsPath     = notegroups.NoteGroupsPath
	NoteGroupPath      = notegroups.NoteGroupPath
	NoteGroupNotesPath = notegroups.NoteGroupNotesPath

	NoteGroupsKey = notegroups.NoteGroupsKey
	NoteGroupKey  = notegroups.NoteGroupKey
)

// Orders is the orders endpoints bound to a client.
type Orders = orders.Orders

// The orders views, aliased so a type assertion works under either spelling.
type (
	OrdersPagination               = orders.Pagination
	OrdersOrderRecord              = orders.OrderRecord
	OrdersOrderList                = orders.OrderList
	OrderNoteRecord                = orders.OrderNoteRecord
	OrderNoteList                  = orders.OrderNoteList
	OrdersOrderReadRecord          = orders.OrderReadRecord
	OrderInformationRecord         = orders.OrderInformationRecord
	AccountOrdersRecord            = orders.AccountOrdersRecord
	OrderBillingPreferencesRecord  = orders.OrderBillingPreferencesRecord
	OrdersActivityLog              = orders.ActivityLog
	OrdersAttachment               = orders.Attachment
	OrdersContractProperties       = orders.ContractProperties
	OrdersCreatedBy                = orders.CreatedBy
	OrdersItem                     = orders.Item
	OrdersItemAccountingCode       = orders.ItemAccountingCode
	ItemCustomAttribute            = orders.ItemCustomAttribute
	OrdersItemPriceSnapshot        = orders.ItemPriceSnapshot
	OrdersItemPriceTax             = orders.ItemPriceTax
	ItemProperties                 = orders.ItemProperties
	OrdersItemSaleTaxConfiguration = orders.ItemSaleTaxConfiguration
	OrdersLastUpdatedBy            = orders.LastUpdatedBy
	OrdersLine                     = orders.Line
	OrdersOrder                    = orders.Order
	OrderDiscardDraftCreateRecord  = orders.OrderDiscardDraftCreateRecord
	OrderExtendTrialCreateRecord   = orders.OrderExtendTrialCreateRecord
	OrderLineReadRecord            = orders.OrderLineReadRecord
	OrderLinesReadRecord           = orders.OrderLinesReadRecord
	OrderNoteFileList              = orders.OrderNoteFileList
	OrderNoteFileReadRecord        = orders.OrderNoteFileReadRecord
	OrderNoteFileRecord            = orders.OrderNoteFileRecord
	OrderNoteReadRecord            = orders.OrderNoteReadRecord
	OrderPDFReadRecord             = orders.OrderPDFReadRecord
	OrderPaymentMethodReadRecord   = orders.OrderPaymentMethodReadRecord
	OrderPaymentMethodsReadRecord  = orders.OrderPaymentMethodsReadRecord
	OrdersPDF                      = orders.PDF
	OrdersPaginationLinks          = orders.PaginationLinks
	OrdersPaymentMethod            = orders.PaymentMethod
	OrdersPreOrderStockDetails     = orders.PreOrderStockDetails
	OrdersPricingRule              = orders.PricingRule
	Processor                      = orders.Processor
	OrdersTaxCode                  = orders.TaxCode
	OrdersWatcher                  = orders.Watcher
	OrdersAccount                  = orders.Account
	OrdersBillingAddress           = orders.BillingAddress
	OrdersBillingPreferences       = orders.BillingPreferences
	Booking                        = orders.Booking
	OrdersCommunicationPreference  = orders.CommunicationPreference
	OrdersCurrency                 = orders.Currency
	OrdersCustomAttribute          = orders.CustomAttribute
	OrdersCustomForm               = orders.CustomForm
	OrdersDiscountProfile          = orders.DiscountProfile
	OrdersKpis                     = orders.Kpis
	OrdersManager                  = orders.Manager
	OrdersProformaProperties       = orders.ProformaProperties
	OrdersReferralAccount          = orders.ReferralAccount
	OrdersSalesRep                 = orders.SalesRep
	OrdersShippingAddress          = orders.ShippingAddress
	OrdersShippingProfile          = orders.ShippingProfile
	OrdersTimeZone                 = orders.TimeZone
)

// NewOrders binds the resource to a requester of your own.
func NewOrders(requester Requester) *Orders { return orders.New(requester) }

// OrdersOrderListShape is the DTOShape reading a body as OrdersOrderList, for the raw-verb call.
func OrdersOrderListShape(decoded any) any { return orders.OrderListShape(decoded) }

// OrdersOrderReadRecordShape is the DTOShape reading a body as OrdersOrderReadRecord, for the raw-verb call.
func OrdersOrderReadRecordShape(decoded any) any { return orders.OrderReadRecordShape(decoded) }

// AccountOrdersRecordShape is the DTOShape reading a body as AccountOrdersRecord, for the raw-verb call.
func AccountOrdersRecordShape(decoded any) any { return orders.AccountOrdersRecordShape(decoded) }

// OrderBillingPreferencesRecordShape is the DTOShape reading a body as OrderBillingPreferencesRecord, for the raw-verb call.
func OrderBillingPreferencesRecordShape(decoded any) any {
	return orders.OrderBillingPreferencesRecordShape(decoded)
}

// OrderInformationRecordShape is the DTOShape reading a body as OrderInformationRecord, for the raw-verb call.
func OrderInformationRecordShape(decoded any) any { return orders.OrderInformationRecordShape(decoded) }

// OrderNoteListShape is the DTOShape reading a body as OrderNoteList, for the raw-verb call.
func OrderNoteListShape(decoded any) any { return orders.OrderNoteListShape(decoded) }

// OrdersOrderDiscardDraftCreateRecordShape is the DTOShape reading a body as OrdersOrderDiscardDraftCreateRecord, for the raw-verb call.
func OrdersOrderDiscardDraftCreateRecordShape(decoded any) any {
	return orders.OrderDiscardDraftCreateRecordShape(decoded)
}

// OrderExtendTrialCreateRecordShape is the DTOShape reading a body as OrderExtendTrialCreateRecord, for the raw-verb call.
func OrderExtendTrialCreateRecordShape(decoded any) any {
	return orders.OrderExtendTrialCreateRecordShape(decoded)
}

// OrdersOrderLineReadRecordShape is the DTOShape reading a body as OrdersOrderLineReadRecord, for the raw-verb call.
func OrdersOrderLineReadRecordShape(decoded any) any { return orders.OrderLineReadRecordShape(decoded) }

// OrdersOrderLinesReadRecordShape is the DTOShape reading a body as OrdersOrderLinesReadRecord, for the raw-verb call.
func OrdersOrderLinesReadRecordShape(decoded any) any {
	return orders.OrderLinesReadRecordShape(decoded)
}

// OrdersOrderNoteFileListShape is the DTOShape reading a body as OrdersOrderNoteFileList, for the raw-verb call.
func OrdersOrderNoteFileListShape(decoded any) any { return orders.OrderNoteFileListShape(decoded) }

// OrdersOrderNoteFileReadRecordShape is the DTOShape reading a body as OrdersOrderNoteFileReadRecord, for the raw-verb call.
func OrdersOrderNoteFileReadRecordShape(decoded any) any {
	return orders.OrderNoteFileReadRecordShape(decoded)
}

// OrdersOrderNoteReadRecordShape is the DTOShape reading a body as OrdersOrderNoteReadRecord, for the raw-verb call.
func OrdersOrderNoteReadRecordShape(decoded any) any { return orders.OrderNoteReadRecordShape(decoded) }

// OrdersOrderPDFReadRecordShape is the DTOShape reading a body as OrdersOrderPDFReadRecord, for the raw-verb call.
func OrdersOrderPDFReadRecordShape(decoded any) any { return orders.OrderPDFReadRecordShape(decoded) }

// OrdersOrderPaymentMethodReadRecordShape is the DTOShape reading a body as OrdersOrderPaymentMethodReadRecord, for the raw-verb call.
func OrdersOrderPaymentMethodReadRecordShape(decoded any) any {
	return orders.OrderPaymentMethodReadRecordShape(decoded)
}

// OrdersOrderPaymentMethodsReadRecordShape is the DTOShape reading a body as OrdersOrderPaymentMethodsReadRecord, for the raw-verb call.
func OrdersOrderPaymentMethodsReadRecordShape(decoded any) any {
	return orders.OrderPaymentMethodsReadRecordShape(decoded)
}

// The orders path templates and body keys, for the raw-verb call.
const (
	OrdersOrdersPath                = orders.OrdersPath
	OrdersOrderPath                 = orders.OrderPath
	OrderActivatePath               = orders.OrderActivatePath
	OrderActivateDraftPath          = orders.OrderActivateDraftPath
	OrderBillingPreferencesPath     = orders.OrderBillingPreferencesPath
	OrderCancelPath                 = orders.OrderCancelPath
	OrderChangePath                 = orders.OrderChangePath
	OrderConvertToOrderPath         = orders.OrderConvertToOrderPath
	OrderCustomObjectsPath          = orders.OrderCustomObjectsPath
	OrderCustomObjectsInstancesPath = orders.OrderCustomObjectsInstancesPath
	OrderDiscardDraftPath           = orders.OrderDiscardDraftPath
	OrderExtendTrialPath            = orders.OrderExtendTrialPath
	OrderInformationPath            = orders.OrderInformationPath
	OrderItemFulfillmentsPath       = orders.OrderItemFulfillmentsPath
	OrderLinesPath                  = orders.OrderLinesPath
	OrderNotesPath                  = orders.OrderNotesPath
	OrderPaymentMethodsPath         = orders.OrderPaymentMethodsPath
	OrderPDFPath                    = orders.OrderPDFPath
	OrderReactivatePath             = orders.OrderReactivatePath
	OrderRelinquishPath             = orders.OrderRelinquishPath
	OrderLinePath                   = orders.OrderLinePath
	OrderNotePath                   = orders.OrderNotePath
	OrderPaymentMethodPath          = orders.OrderPaymentMethodPath
	OrderLineBillingPreferencesPath = orders.OrderLineBillingPreferencesPath
	OrderNoteFilesPath              = orders.OrderNoteFilesPath
	OrderNoteFilePath               = orders.OrderNoteFilePath

	OrdersKey      = orders.OrdersKey
	OrdersOrderKey = orders.OrderKey
)

// Payments is the payments endpoints bound to a client.
type Payments = payments.Payments

// The payments views, aliased so a type assertion works under either spelling.
type (
	PaymentsPagination        = payments.Pagination
	PaymentRecord             = payments.PaymentRecord
	PaymentList               = payments.PaymentList
	PaymentNoteRecord         = payments.PaymentNoteRecord
	PaymentNoteList           = payments.PaymentNoteList
	PaymentsNoteRecord        = payments.PaymentsNoteRecord
	PaymentsNoteList          = payments.PaymentsNoteList
	PaymentReadRecord         = payments.PaymentReadRecord
	AccountPaymentsRecord     = payments.AccountPaymentsRecord
	PaymentNoteReadRecord     = payments.PaymentNoteReadRecord
	PaymentsAttachment        = payments.Attachment
	PaymentsCreatedBy         = payments.CreatedBy
	PaymentsPaginationLinks   = payments.PaginationLinks
	PaymentsPayment           = payments.Payment
	PaymentNoteFileReadRecord = payments.PaymentNoteFileReadRecord
	PaymentsWatcher           = payments.Watcher
	PaymentsAccount           = payments.Account
	CreditApplied             = payments.CreditApplied
	PaymentsCustomAttribute   = payments.CustomAttribute
	GiftCertificateApplied    = payments.GiftCertificateApplied
	PaymentsInvoice           = payments.Invoice
	PaymentsLastUpdatedBy     = payments.LastUpdatedBy
	PaymentApplied            = payments.PaymentApplied
)

// NewPayments binds the resource to a requester of your own.
func NewPayments(requester Requester) *Payments { return payments.New(requester) }

// PaymentListShape is the DTOShape reading a body as PaymentList, for the raw-verb call.
func PaymentListShape(decoded any) any { return payments.PaymentListShape(decoded) }

// PaymentsNoteListShape is the DTOShape reading a body as PaymentsNoteList, for the raw-verb call.
func PaymentsNoteListShape(decoded any) any { return payments.PaymentsNoteListShape(decoded) }

// PaymentReadRecordShape is the DTOShape reading a body as PaymentReadRecord, for the raw-verb call.
func PaymentReadRecordShape(decoded any) any { return payments.PaymentReadRecordShape(decoded) }

// AccountPaymentsRecordShape is the DTOShape reading a body as AccountPaymentsRecord, for the raw-verb call.
func AccountPaymentsRecordShape(decoded any) any { return payments.AccountPaymentsRecordShape(decoded) }

// PaymentNoteListShape is the DTOShape reading a body as PaymentNoteList, for the raw-verb call.
func PaymentNoteListShape(decoded any) any { return payments.PaymentNoteListShape(decoded) }

// PaymentNoteReadRecordShape is the DTOShape reading a body as PaymentNoteReadRecord, for the raw-verb call.
func PaymentNoteReadRecordShape(decoded any) any { return payments.PaymentNoteReadRecordShape(decoded) }

// PaymentsPaymentNoteFileReadRecordShape is the DTOShape reading a body as PaymentsPaymentNoteFileReadRecord, for the raw-verb call.
func PaymentsPaymentNoteFileReadRecordShape(decoded any) any {
	return payments.PaymentNoteFileReadRecordShape(decoded)
}

// The payments path templates and body keys, for the raw-verb call.
const (
	PaymentsPath        = payments.PaymentsPath
	PaymentsNotesPath   = payments.PaymentsNotesPath
	PaymentPath         = payments.PaymentPath
	InvoicePaymentsPath = payments.InvoicePaymentsPath
	OrderPaymentsPath   = payments.OrderPaymentsPath
	PaymentNotesPath    = payments.PaymentNotesPath
	PaymentNotePath     = payments.PaymentNotePath
	PaymentNoteFilePath = payments.PaymentNoteFilePath

	PaymentsKey = payments.PaymentsKey
	PaymentKey  = payments.PaymentKey
)

// Portals is the portals endpoints bound to a client.
type Portals = portals.Portals

// The portals views, aliased so a type assertion works under either spelling.
type (
	PortalsPagination                = portals.Pagination
	PortalLoginRecord                = portals.PortalLoginRecord
	PortalChangePasswordCreateRecord = portals.PortalChangePasswordCreateRecord
	Portal                           = portals.Portal
	PortalsUser                      = portals.User
)

// NewPortals binds the resource to a requester of your own.
func NewPortals(requester Requester) *Portals { return portals.New(requester) }

// PortalLoginRecordShape is the DTOShape reading a body as PortalLoginRecord, for the raw-verb call.
func PortalLoginRecordShape(decoded any) any { return portals.PortalLoginRecordShape(decoded) }

// PortalsPortalChangePasswordCreateRecordShape is the DTOShape reading a body as PortalsPortalChangePasswordCreateRecord, for the raw-verb call.
func PortalsPortalChangePasswordCreateRecordShape(decoded any) any {
	return portals.PortalChangePasswordCreateRecordShape(decoded)
}

// The portals path templates and body keys, for the raw-verb call.
const (
	PortalActivatePath       = portals.PortalActivatePath
	PortalChangePasswordPath = portals.PortalChangePasswordPath
	PortalLoginPath          = portals.PortalLoginPath
	PortalResetPath          = portals.PortalResetPath
	PortalResetLinkPath      = portals.PortalResetLinkPath
	PortalValidatePath       = portals.PortalValidatePath

	PortalKey = portals.PortalKey
)

// PurchaseCreditNoteApplications is the purchase credit note applications endpoints bound to a client.
type PurchaseCreditNoteApplications = purchasecreditnoteapplications.PurchaseCreditNoteApplications

// The purchase credit note applications views, aliased so a type assertion works under either spelling.
type (
	PurchaseCreditNoteApplicationsPagination         = purchasecreditnoteapplications.Pagination
	PurchaseCreditNoteApplicationRecord              = purchasecreditnoteapplications.PurchaseCreditNoteApplicationRecord
	PurchaseCreditNoteApplicationList                = purchasecreditnoteapplications.PurchaseCreditNoteApplicationList
	PurchaseCreditNoteApplicationReadRecord          = purchasecreditnoteapplications.PurchaseCreditNoteApplicationReadRecord
	PurchaseCreditNoteApplicationsPaginationLinks    = purchasecreditnoteapplications.PaginationLinks
	PurchaseCreditNoteApplicationsCreatedBy          = purchasecreditnoteapplications.CreatedBy
	PurchaseCreditNoteApplicationsPurchaseCreditNote = purchasecreditnoteapplications.PurchaseCreditNote
	PurchaseCreditNoteApplicationsPurchasePayment    = purchasecreditnoteapplications.PurchasePayment
)

// NewPurchaseCreditNoteApplications binds the resource to a requester of your own.
func NewPurchaseCreditNoteApplications(requester Requester) *PurchaseCreditNoteApplications {
	return purchasecreditnoteapplications.New(requester)
}

// PurchaseCreditNoteApplicationListShape is the DTOShape reading a body as PurchaseCreditNoteApplicationList, for the raw-verb call.
func PurchaseCreditNoteApplicationListShape(decoded any) any {
	return purchasecreditnoteapplications.PurchaseCreditNoteApplicationListShape(decoded)
}

// PurchaseCreditNoteApplicationReadRecordShape is the DTOShape reading a body as PurchaseCreditNoteApplicationReadRecord, for the raw-verb call.
func PurchaseCreditNoteApplicationReadRecordShape(decoded any) any {
	return purchasecreditnoteapplications.PurchaseCreditNoteApplicationReadRecordShape(decoded)
}

// The purchase credit note applications path templates and body keys, for the raw-verb call.
const (
	PurchaseCreditNoteApplicationsPath = purchasecreditnoteapplications.PurchaseCreditNoteApplicationsPath
	PurchaseCreditNoteApplicationPath  = purchasecreditnoteapplications.PurchaseCreditNoteApplicationPath

	PurchaseCreditNoteApplicationsKey = purchasecreditnoteapplications.PurchaseCreditNoteApplicationsKey
)

// PurchaseCreditNotes is the purchase credit notes endpoints bound to a client.
type PurchaseCreditNotes = purchasecreditnotes.PurchaseCreditNotes

// The purchase credit notes views, aliased so a type assertion works under either spelling.
type (
	PurchaseCreditNotesPagination       = purchasecreditnotes.Pagination
	PurchaseCreditNoteRecord            = purchasecreditnotes.PurchaseCreditNoteRecord
	PurchaseCreditNoteList              = purchasecreditnotes.PurchaseCreditNoteList
	PurchaseCreditNoteReadRecord        = purchasecreditnotes.PurchaseCreditNoteReadRecord
	PurchaseCreditNotesPaginationLinks  = purchasecreditnotes.PaginationLinks
	PurchaseCreditNoteUpdateRecord      = purchasecreditnotes.PurchaseCreditNoteUpdateRecord
	Status                              = purchasecreditnotes.Status
	PurchaseCreditNotesAccount          = purchasecreditnotes.Account
	PurchaseCreditNotesCreatedBy        = purchasecreditnotes.CreatedBy
	PurchaseCreditNotesCurrency         = purchasecreditnotes.Currency
	PurchaseCreditNotesLastUpdatedBy    = purchasecreditnotes.LastUpdatedBy
	PurchaseCreditNotesPaymentMethod    = purchasecreditnotes.PaymentMethod
	PurchaseCreditNotesPaymentProcessor = purchasecreditnotes.PaymentProcessor
	PurchaseCreditNoteRow               = purchasecreditnotes.PurchaseCreditNoteRow
	PurchaseCreditNotesPurchasePayment  = purchasecreditnotes.PurchasePayment
)

// NewPurchaseCreditNotes binds the resource to a requester of your own.
func NewPurchaseCreditNotes(requester Requester) *PurchaseCreditNotes {
	return purchasecreditnotes.New(requester)
}

// PurchaseCreditNoteListShape is the DTOShape reading a body as PurchaseCreditNoteList, for the raw-verb call.
func PurchaseCreditNoteListShape(decoded any) any {
	return purchasecreditnotes.PurchaseCreditNoteListShape(decoded)
}

// PurchaseCreditNoteReadRecordShape is the DTOShape reading a body as PurchaseCreditNoteReadRecord, for the raw-verb call.
func PurchaseCreditNoteReadRecordShape(decoded any) any {
	return purchasecreditnotes.PurchaseCreditNoteReadRecordShape(decoded)
}

// PurchaseCreditNotesPurchaseCreditNoteUpdateRecordShape is the DTOShape reading a body as PurchaseCreditNotesPurchaseCreditNoteUpdateRecord, for the raw-verb call.
func PurchaseCreditNotesPurchaseCreditNoteUpdateRecordShape(decoded any) any {
	return purchasecreditnotes.PurchaseCreditNoteUpdateRecordShape(decoded)
}

// The purchase credit notes path templates and body keys, for the raw-verb call.
const (
	PurchaseCreditNotesPath = purchasecreditnotes.PurchaseCreditNotesPath
	PurchaseCreditNotePath  = purchasecreditnotes.PurchaseCreditNotePath

	PurchaseCreditNotesKey = purchasecreditnotes.PurchaseCreditNotesKey
	PurchaseCreditNoteKey  = purchasecreditnotes.PurchaseCreditNoteKey
)

// PurchaseInvoices is the purchase invoices endpoints bound to a client.
type PurchaseInvoices = purchaseinvoices.PurchaseInvoices

// The purchase invoices views, aliased so a type assertion works under either spelling.
type (
	PurchaseInvoicesPagination                         = purchaseinvoices.Pagination
	PurchaseInvoiceRecord                              = purchaseinvoices.PurchaseInvoiceRecord
	PurchaseInvoiceList                                = purchaseinvoices.PurchaseInvoiceList
	PurchaseInvoicePurchasePaymentRecord               = purchaseinvoices.PurchaseInvoicePurchasePaymentRecord
	PurchaseInvoicePurchasePaymentList                 = purchaseinvoices.PurchaseInvoicePurchasePaymentList
	PurchaseInvoiceNoteRecord                          = purchaseinvoices.PurchaseInvoiceNoteRecord
	PurchaseInvoiceNoteList                            = purchaseinvoices.PurchaseInvoiceNoteList
	PurchaseInvoicesNoteRecord                         = purchaseinvoices.PurchaseInvoicesNoteRecord
	PurchaseInvoicesNoteList                           = purchaseinvoices.PurchaseInvoicesNoteList
	PurchaseInvoiceReadRecord                          = purchaseinvoices.PurchaseInvoiceReadRecord
	AccountPurchaseInvoicesRecord                      = purchaseinvoices.AccountPurchaseInvoicesRecord
	PurchaseInvoiceInformationRecord                   = purchaseinvoices.PurchaseInvoiceInformationRecord
	PurchaseInvoiceCancelRecord                        = purchaseinvoices.PurchaseInvoiceCancelRecord
	PurchaseInvoiceNotesRecord                         = purchaseinvoices.PurchaseInvoiceNotesRecord
	PurchaseInvoiceLinesRecord                         = purchaseinvoices.PurchaseInvoiceLinesRecord
	PurchaseInvoiceLineRecord                          = purchaseinvoices.PurchaseInvoiceLineRecord
	PurchaseInvoiceNoteReadRecord                      = purchaseinvoices.PurchaseInvoiceNoteReadRecord
	PurchaseInvoicesAccount                            = purchaseinvoices.Account
	PurchaseInvoicesAttachment                         = purchaseinvoices.Attachment
	PurchaseInvoicesCreatedBy                          = purchaseinvoices.CreatedBy
	PurchaseInvoicesCurrency                           = purchaseinvoices.Currency
	PurchaseInvoicesItem                               = purchaseinvoices.Item
	PurchaseInvoicesKpis                               = purchaseinvoices.Kpis
	PurchaseInvoicesLastUpdatedBy                      = purchaseinvoices.LastUpdatedBy
	PurchaseInvoicesLine                               = purchaseinvoices.Line
	PurchaseInvoicesPaginationLinks                    = purchaseinvoices.PaginationLinks
	PurchaseInvoicesPurchaseInvoice                    = purchaseinvoices.PurchaseInvoice
	PurchaseInvoiceNoteFileList                        = purchaseinvoices.PurchaseInvoiceNoteFileList
	PurchaseInvoiceNoteFileReadRecord                  = purchaseinvoices.PurchaseInvoiceNoteFileReadRecord
	PurchaseInvoiceNoteFileRecord                      = purchaseinvoices.PurchaseInvoiceNoteFileRecord
	PurchaseInvoicesWatcher                            = purchaseinvoices.Watcher
	AccountPurchaseInvoice                             = purchaseinvoices.AccountPurchaseInvoice
	PurchaseInvoicesCreditNoteApplication              = purchaseinvoices.CreditNoteApplication
	PurchaseInvoicesCustomAttribute                    = purchaseinvoices.CustomAttribute
	PurchaseInvoicesPaymentMethod                      = purchaseinvoices.PaymentMethod
	PurchaseInvoicesPaymentMethodApplication           = purchaseinvoices.PaymentMethodApplication
	PurchaseInvoicesPaymentProcessor                   = purchaseinvoices.PaymentProcessor
	PurchaseInvoicesPurchaseCreditNote                 = purchaseinvoices.PurchaseCreditNote
	PurchaseInvoicesPurchaseInvoiceApplication         = purchaseinvoices.PurchaseInvoiceApplication
	PurchaseInvoicePurchaseCreditNoteApplicationList   = purchaseinvoices.PurchaseInvoicePurchaseCreditNoteApplicationList
	PurchaseInvoicePurchaseCreditNoteApplicationRecord = purchaseinvoices.PurchaseInvoicePurchaseCreditNoteApplicationRecord
	PurchaseInvoicesPurchaseOrder                      = purchaseinvoices.PurchaseOrder
	PurchaseInvoicesPurchasePayment                    = purchaseinvoices.PurchasePayment
)

// NewPurchaseInvoices binds the resource to a requester of your own.
func NewPurchaseInvoices(requester Requester) *PurchaseInvoices {
	return purchaseinvoices.New(requester)
}

// PurchaseInvoiceListShape is the DTOShape reading a body as PurchaseInvoiceList, for the raw-verb call.
func PurchaseInvoiceListShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceListShape(decoded)
}

// PurchaseInvoicesNoteListShape is the DTOShape reading a body as PurchaseInvoicesNoteList, for the raw-verb call.
func PurchaseInvoicesNoteListShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoicesNoteListShape(decoded)
}

// PurchaseInvoiceReadRecordShape is the DTOShape reading a body as PurchaseInvoiceReadRecord, for the raw-verb call.
func PurchaseInvoiceReadRecordShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceReadRecordShape(decoded)
}

// AccountPurchaseInvoicesRecordShape is the DTOShape reading a body as AccountPurchaseInvoicesRecord, for the raw-verb call.
func AccountPurchaseInvoicesRecordShape(decoded any) any {
	return purchaseinvoices.AccountPurchaseInvoicesRecordShape(decoded)
}

// PurchaseInvoiceCancelRecordShape is the DTOShape reading a body as PurchaseInvoiceCancelRecord, for the raw-verb call.
func PurchaseInvoiceCancelRecordShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceCancelRecordShape(decoded)
}

// PurchaseInvoiceRecordShape is the DTOShape reading a body as PurchaseInvoiceRecord, for the raw-verb call.
func PurchaseInvoiceRecordShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceRecordShape(decoded)
}

// PurchaseInvoiceInformationRecordShape is the DTOShape reading a body as PurchaseInvoiceInformationRecord, for the raw-verb call.
func PurchaseInvoiceInformationRecordShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceInformationRecordShape(decoded)
}

// PurchaseInvoiceLinesRecordShape is the DTOShape reading a body as PurchaseInvoiceLinesRecord, for the raw-verb call.
func PurchaseInvoiceLinesRecordShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceLinesRecordShape(decoded)
}

// PurchaseInvoiceNoteListShape is the DTOShape reading a body as PurchaseInvoiceNoteList, for the raw-verb call.
func PurchaseInvoiceNoteListShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceNoteListShape(decoded)
}

// PurchaseInvoiceNotesRecordShape is the DTOShape reading a body as PurchaseInvoiceNotesRecord, for the raw-verb call.
func PurchaseInvoiceNotesRecordShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceNotesRecordShape(decoded)
}

// PurchaseInvoicePurchasePaymentListShape is the DTOShape reading a body as PurchaseInvoicePurchasePaymentList, for the raw-verb call.
func PurchaseInvoicePurchasePaymentListShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoicePurchasePaymentListShape(decoded)
}

// PurchaseInvoiceLineRecordShape is the DTOShape reading a body as PurchaseInvoiceLineRecord, for the raw-verb call.
func PurchaseInvoiceLineRecordShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceLineRecordShape(decoded)
}

// PurchaseInvoiceNoteReadRecordShape is the DTOShape reading a body as PurchaseInvoiceNoteReadRecord, for the raw-verb call.
func PurchaseInvoiceNoteReadRecordShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceNoteReadRecordShape(decoded)
}

// PurchaseInvoicesPurchaseInvoiceNoteFileListShape is the DTOShape reading a body as PurchaseInvoicesPurchaseInvoiceNoteFileList, for the raw-verb call.
func PurchaseInvoicesPurchaseInvoiceNoteFileListShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceNoteFileListShape(decoded)
}

// PurchaseInvoicesPurchaseInvoiceNoteFileReadRecordShape is the DTOShape reading a body as PurchaseInvoicesPurchaseInvoiceNoteFileReadRecord, for the raw-verb call.
func PurchaseInvoicesPurchaseInvoiceNoteFileReadRecordShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoiceNoteFileReadRecordShape(decoded)
}

// PurchaseInvoicesPurchaseInvoicePurchaseCreditNoteApplicationListShape is the DTOShape reading a body as PurchaseInvoicesPurchaseInvoicePurchaseCreditNoteApplicationList, for the raw-verb call.
func PurchaseInvoicesPurchaseInvoicePurchaseCreditNoteApplicationListShape(decoded any) any {
	return purchaseinvoices.PurchaseInvoicePurchaseCreditNoteApplicationListShape(decoded)
}

// The purchase invoices path templates and body keys, for the raw-verb call.
const (
	PurchaseInvoicesPath                                              = purchaseinvoices.PurchaseInvoicesPath
	PurchaseInvoicesNotesPath                                         = purchaseinvoices.PurchaseInvoicesNotesPath
	PurchaseInvoicePath                                               = purchaseinvoices.PurchaseInvoicePath
	PurchaseInvoiceAmendPath                                          = purchaseinvoices.PurchaseInvoiceAmendPath
	PurchaseInvoiceCancelPath                                         = purchaseinvoices.PurchaseInvoiceCancelPath
	PurchaseInvoiceInformationPath                                    = purchaseinvoices.PurchaseInvoiceInformationPath
	PurchaseInvoiceLinesPath                                          = purchaseinvoices.PurchaseInvoiceLinesPath
	PurchaseInvoiceNotesPath                                          = purchaseinvoices.PurchaseInvoiceNotesPath
	PurchaseInvoicePDFPath                                            = purchaseinvoices.PurchaseInvoicePDFPath
	PurchaseInvoicesPurchaseInvoicePurchaseCreditNoteApplicationsPath = purchaseinvoices.PurchaseInvoicePurchaseCreditNoteApplicationsPath
	PurchaseInvoicesPurchaseInvoicePurchasePaymentsPath               = purchaseinvoices.PurchaseInvoicePurchasePaymentsPath
	PurchaseInvoiceReactivatePath                                     = purchaseinvoices.PurchaseInvoiceReactivatePath
	PurchaseInvoiceLinePath                                           = purchaseinvoices.PurchaseInvoiceLinePath
	PurchaseInvoiceNotePath                                           = purchaseinvoices.PurchaseInvoiceNotePath
	PurchaseInvoiceNoteFilesPath                                      = purchaseinvoices.PurchaseInvoiceNoteFilesPath
	PurchaseInvoiceNoteFilePath                                       = purchaseinvoices.PurchaseInvoiceNoteFilePath

	PurchaseInvoicesKey = purchaseinvoices.PurchaseInvoicesKey
	PurchaseInvoiceKey  = purchaseinvoices.PurchaseInvoiceKey
	WatcherListKey      = purchaseinvoices.WatcherListKey
)

// PurchaseOrders is the purchase orders endpoints bound to a client.
type PurchaseOrders = purchaseorders.PurchaseOrders

// The purchase orders views, aliased so a type assertion works under either spelling.
type (
	PurchaseOrdersPagination                 = purchaseorders.Pagination
	PurchaseOrderLineRecord                  = purchaseorders.PurchaseOrderLineRecord
	PurchaseOrderLineList                    = purchaseorders.PurchaseOrderLineList
	PurchaseOrderRecord                      = purchaseorders.PurchaseOrderRecord
	PurchaseOrderList                        = purchaseorders.PurchaseOrderList
	PurchaseOrderPurchaseInvoiceRecord       = purchaseorders.PurchaseOrderPurchaseInvoiceRecord
	PurchaseOrderPurchaseInvoiceList         = purchaseorders.PurchaseOrderPurchaseInvoiceList
	PurchaseOrderPurchasePaymentRecord       = purchaseorders.PurchaseOrderPurchasePaymentRecord
	PurchaseOrderPurchasePaymentList         = purchaseorders.PurchaseOrderPurchasePaymentList
	PurchaseOrdersNoteRecord                 = purchaseorders.PurchaseOrdersNoteRecord
	PurchaseOrdersNoteList                   = purchaseorders.PurchaseOrdersNoteList
	PurchaseOrderReadRecord                  = purchaseorders.PurchaseOrderReadRecord
	PurchaseOrderUpdatedRecord               = purchaseorders.PurchaseOrderUpdatedRecord
	PurchaseOrdersRecord                     = purchaseorders.PurchaseOrdersRecord
	PurchaseOrderInformationRecord           = purchaseorders.PurchaseOrderInformationRecord
	AccountPurchaseOrdersRecord              = purchaseorders.AccountPurchaseOrdersRecord
	PurchaseOrderNotesRecord                 = purchaseorders.PurchaseOrderNotesRecord
	PurchaseOrderNoteRecord                  = purchaseorders.PurchaseOrderNoteRecord
	PurchaseOrderLineReadRecord              = purchaseorders.PurchaseOrderLineReadRecord
	PurchaseOrdersAccount                    = purchaseorders.Account
	PurchaseOrdersAttachment                 = purchaseorders.Attachment
	PurchaseOrdersCreatedBy                  = purchaseorders.CreatedBy
	PurchaseOrdersCurrency                   = purchaseorders.Currency
	PurchaseOrdersData                       = purchaseorders.Data
	PurchaseOrdersItem                       = purchaseorders.Item
	PurchaseOrdersItemAccountingCode         = purchaseorders.ItemAccountingCode
	PurchaseOrdersItemPriceSnapshot          = purchaseorders.ItemPriceSnapshot
	PurchaseOrdersItemPriceTax               = purchaseorders.ItemPriceTax
	ItemPurchaseAccountingCode               = purchaseorders.ItemPurchaseAccountingCode
	ItemPurchaseTaxConfiguration             = purchaseorders.ItemPurchaseTaxConfiguration
	PurchaseOrdersKpis                       = purchaseorders.Kpis
	PurchaseOrdersLastUpdatedBy              = purchaseorders.LastUpdatedBy
	PurchaseOrdersLine                       = purchaseorders.Line
	PurchaseOrdersPaginationLinks            = purchaseorders.PaginationLinks
	PurchaseOrdersPricingRule                = purchaseorders.PricingRule
	PurchaseOrdersPurchaseOrder              = purchaseorders.PurchaseOrder
	PurchaseOrderNoteFileList                = purchaseorders.PurchaseOrderNoteFileList
	PurchaseOrderNoteFileReadRecord          = purchaseorders.PurchaseOrderNoteFileReadRecord
	PurchaseOrderNoteFileRecord              = purchaseorders.PurchaseOrderNoteFileRecord
	PurchaseOrderPurchaseInvoiceCreateRecord = purchaseorders.PurchaseOrderPurchaseInvoiceCreateRecord
	PurchaseOrdersTax                        = purchaseorders.Tax
	PurchaseOrdersTaxCode                    = purchaseorders.TaxCode
	PurchaseOrdersUom                        = purchaseorders.Uom
	PurchaseOrdersWarehouse                  = purchaseorders.Warehouse
	PurchaseOrdersWatcher                    = purchaseorders.Watcher
	PurchaseOrdersCreditNoteApplication      = purchaseorders.CreditNoteApplication
	PurchaseOrdersCustomAttribute            = purchaseorders.CustomAttribute
	PurchaseOrdersCustomForm                 = purchaseorders.CustomForm
	PurchaseOrdersPaymentMethod              = purchaseorders.PaymentMethod
	PurchaseOrdersPaymentMethodApplication   = purchaseorders.PaymentMethodApplication
	PurchaseOrdersPaymentProcessor           = purchaseorders.PaymentProcessor
	PurchaseOrdersPurchaseCreditNote         = purchaseorders.PurchaseCreditNote
	PurchaseOrdersPurchaseInvoice            = purchaseorders.PurchaseInvoice
	PurchaseOrdersPurchaseInvoiceApplication = purchaseorders.PurchaseInvoiceApplication
	PurchaseOrderKpis                        = purchaseorders.PurchaseOrderKpis
	PurchaseOrderUpdatedKpis                 = purchaseorders.PurchaseOrderUpdatedKpis
	PurchaseOrdersPurchaseOrdersKpis         = purchaseorders.PurchaseOrdersKpis
)

// NewPurchaseOrders binds the resource to a requester of your own.
func NewPurchaseOrders(requester Requester) *PurchaseOrders { return purchaseorders.New(requester) }

// PurchaseOrderListShape is the DTOShape reading a body as PurchaseOrderList, for the raw-verb call.
func PurchaseOrderListShape(decoded any) any { return purchaseorders.PurchaseOrderListShape(decoded) }

// PurchaseOrdersRecordShape is the DTOShape reading a body as PurchaseOrdersRecord, for the raw-verb call.
func PurchaseOrdersRecordShape(decoded any) any {
	return purchaseorders.PurchaseOrdersRecordShape(decoded)
}

// PurchaseOrdersNoteListShape is the DTOShape reading a body as PurchaseOrdersNoteList, for the raw-verb call.
func PurchaseOrdersNoteListShape(decoded any) any {
	return purchaseorders.PurchaseOrdersNoteListShape(decoded)
}

// PurchaseOrderUpdatedRecordShape is the DTOShape reading a body as PurchaseOrderUpdatedRecord, for the raw-verb call.
func PurchaseOrderUpdatedRecordShape(decoded any) any {
	return purchaseorders.PurchaseOrderUpdatedRecordShape(decoded)
}

// PurchaseOrderReadRecordShape is the DTOShape reading a body as PurchaseOrderReadRecord, for the raw-verb call.
func PurchaseOrderReadRecordShape(decoded any) any {
	return purchaseorders.PurchaseOrderReadRecordShape(decoded)
}

// AccountPurchaseOrdersRecordShape is the DTOShape reading a body as AccountPurchaseOrdersRecord, for the raw-verb call.
func AccountPurchaseOrdersRecordShape(decoded any) any {
	return purchaseorders.AccountPurchaseOrdersRecordShape(decoded)
}

// PurchaseOrderInformationRecordShape is the DTOShape reading a body as PurchaseOrderInformationRecord, for the raw-verb call.
func PurchaseOrderInformationRecordShape(decoded any) any {
	return purchaseorders.PurchaseOrderInformationRecordShape(decoded)
}

// PurchaseOrderLineListShape is the DTOShape reading a body as PurchaseOrderLineList, for the raw-verb call.
func PurchaseOrderLineListShape(decoded any) any {
	return purchaseorders.PurchaseOrderLineListShape(decoded)
}

// PurchaseOrderNotesRecordShape is the DTOShape reading a body as PurchaseOrderNotesRecord, for the raw-verb call.
func PurchaseOrderNotesRecordShape(decoded any) any {
	return purchaseorders.PurchaseOrderNotesRecordShape(decoded)
}

// PurchaseOrderPurchaseInvoiceListShape is the DTOShape reading a body as PurchaseOrderPurchaseInvoiceList, for the raw-verb call.
func PurchaseOrderPurchaseInvoiceListShape(decoded any) any {
	return purchaseorders.PurchaseOrderPurchaseInvoiceListShape(decoded)
}

// PurchaseOrderPurchasePaymentListShape is the DTOShape reading a body as PurchaseOrderPurchasePaymentList, for the raw-verb call.
func PurchaseOrderPurchasePaymentListShape(decoded any) any {
	return purchaseorders.PurchaseOrderPurchasePaymentListShape(decoded)
}

// PurchaseOrderLineReadRecordShape is the DTOShape reading a body as PurchaseOrderLineReadRecord, for the raw-verb call.
func PurchaseOrderLineReadRecordShape(decoded any) any {
	return purchaseorders.PurchaseOrderLineReadRecordShape(decoded)
}

// PurchaseOrderNoteRecordShape is the DTOShape reading a body as PurchaseOrderNoteRecord, for the raw-verb call.
func PurchaseOrderNoteRecordShape(decoded any) any {
	return purchaseorders.PurchaseOrderNoteRecordShape(decoded)
}

// PurchaseOrdersPurchaseOrderNoteFileListShape is the DTOShape reading a body as PurchaseOrdersPurchaseOrderNoteFileList, for the raw-verb call.
func PurchaseOrdersPurchaseOrderNoteFileListShape(decoded any) any {
	return purchaseorders.PurchaseOrderNoteFileListShape(decoded)
}

// PurchaseOrdersPurchaseOrderNoteFileReadRecordShape is the DTOShape reading a body as PurchaseOrdersPurchaseOrderNoteFileReadRecord, for the raw-verb call.
func PurchaseOrdersPurchaseOrderNoteFileReadRecordShape(decoded any) any {
	return purchaseorders.PurchaseOrderNoteFileReadRecordShape(decoded)
}

// PurchaseOrdersPurchaseOrderPurchaseInvoiceCreateRecordShape is the DTOShape reading a body as PurchaseOrdersPurchaseOrderPurchaseInvoiceCreateRecord, for the raw-verb call.
func PurchaseOrdersPurchaseOrderPurchaseInvoiceCreateRecordShape(decoded any) any {
	return purchaseorders.PurchaseOrderPurchaseInvoiceCreateRecordShape(decoded)
}

// The purchase orders path templates and body keys, for the raw-verb call.
const (
	PurchaseOrdersPath                              = purchaseorders.PurchaseOrdersPath
	PurchaseOrdersNotesPath                         = purchaseorders.PurchaseOrdersNotesPath
	PurchaseOrderPath                               = purchaseorders.PurchaseOrderPath
	PurchaseOrderCancelPath                         = purchaseorders.PurchaseOrderCancelPath
	PurchaseOrderChangePath                         = purchaseorders.PurchaseOrderChangePath
	PurchaseOrderInformationPath                    = purchaseorders.PurchaseOrderInformationPath
	PurchaseOrderLinesPath                          = purchaseorders.PurchaseOrderLinesPath
	PurchaseOrderNotesPath                          = purchaseorders.PurchaseOrderNotesPath
	PurchaseOrdersPurchaseOrderPurchaseInvoicesPath = purchaseorders.PurchaseOrderPurchaseInvoicesPath
	PurchaseOrdersPurchaseOrderPurchasePaymentsPath = purchaseorders.PurchaseOrderPurchasePaymentsPath
	PurchaseOrderReactivatePath                     = purchaseorders.PurchaseOrderReactivatePath
	PurchaseOrderLinePath                           = purchaseorders.PurchaseOrderLinePath
	PurchaseOrderNotePath                           = purchaseorders.PurchaseOrderNotePath
	PurchaseOrderNoteFilesPath                      = purchaseorders.PurchaseOrderNoteFilesPath
	PurchaseOrderNoteFilePath                       = purchaseorders.PurchaseOrderNoteFilePath

	PurchaseOrdersKey              = purchaseorders.PurchaseOrdersKey
	PurchaseOrdersPurchaseOrderKey = purchaseorders.PurchaseOrderKey
)

// PurchasePayments is the purchase payments endpoints bound to a client.
type PurchasePayments = purchasepayments.PurchasePayments

// The purchase payments views, aliased so a type assertion works under either spelling.
type (
	PurchasePaymentsPagination        = purchasepayments.Pagination
	PurchasePaymentRecord             = purchasepayments.PurchasePaymentRecord
	PurchasePaymentList               = purchasepayments.PurchasePaymentList
	PurchasePaymentReadRecord         = purchasepayments.PurchasePaymentReadRecord
	PurchasePaymentNoteRecord         = purchasepayments.PurchasePaymentNoteRecord
	PurchasePaymentNoteList           = purchasepayments.PurchasePaymentNoteList
	PurchasePaymentUpdatedRecord      = purchasepayments.PurchasePaymentUpdatedRecord
	PurchasePaymentsRecord            = purchasepayments.PurchasePaymentsRecord
	AccountPurchasePaymentsRecord     = purchasepayments.AccountPurchasePaymentsRecord
	PurchasePaymentNotesRecord        = purchasepayments.PurchasePaymentNotesRecord
	PurchasePaymentNoteReadRecord     = purchasepayments.PurchasePaymentNoteReadRecord
	PurchasePaymentsAccount           = purchasepayments.Account
	PurchasePaymentsActivityLog       = purchasepayments.ActivityLog
	PurchasePaymentsAttachment        = purchasepayments.Attachment
	PurchasePaymentsCreatedBy         = purchasepayments.CreatedBy
	CreditNoteApplication             = purchasepayments.CreditNoteApplication
	PurchasePaymentsCurrency          = purchasepayments.Currency
	PurchasePaymentsLastUpdatedBy     = purchasepayments.LastUpdatedBy
	PurchasePaymentsPaginationLinks   = purchasepayments.PaginationLinks
	PurchasePaymentsPaymentMethod     = purchasepayments.PaymentMethod
	PaymentMethodApplication          = purchasepayments.PaymentMethodApplication
	PurchasePaymentsPaymentProcessor  = purchasepayments.PaymentProcessor
	PurchaseCreditNote                = purchasepayments.PurchaseCreditNote
	PurchasePaymentsPurchaseInvoice   = purchasepayments.PurchaseInvoice
	PurchaseInvoiceApplication        = purchasepayments.PurchaseInvoiceApplication
	PurchasePaymentsPurchasePayment   = purchasepayments.PurchasePayment
	PurchasePaymentNoteFileList       = purchasepayments.PurchasePaymentNoteFileList
	PurchasePaymentNoteFileReadRecord = purchasepayments.PurchasePaymentNoteFileReadRecord
	PurchasePaymentNoteFileRecord     = purchasepayments.PurchasePaymentNoteFileRecord
	PurchasePaymentsWatcher           = purchasepayments.Watcher
)

// NewPurchasePayments binds the resource to a requester of your own.
func NewPurchasePayments(requester Requester) *PurchasePayments {
	return purchasepayments.New(requester)
}

// PurchasePaymentListShape is the DTOShape reading a body as PurchasePaymentList, for the raw-verb call.
func PurchasePaymentListShape(decoded any) any {
	return purchasepayments.PurchasePaymentListShape(decoded)
}

// PurchasePaymentsRecordShape is the DTOShape reading a body as PurchasePaymentsRecord, for the raw-verb call.
func PurchasePaymentsRecordShape(decoded any) any {
	return purchasepayments.PurchasePaymentsRecordShape(decoded)
}

// PurchasePaymentReadRecordShape is the DTOShape reading a body as PurchasePaymentReadRecord, for the raw-verb call.
func PurchasePaymentReadRecordShape(decoded any) any {
	return purchasepayments.PurchasePaymentReadRecordShape(decoded)
}

// PurchasePaymentUpdatedRecordShape is the DTOShape reading a body as PurchasePaymentUpdatedRecord, for the raw-verb call.
func PurchasePaymentUpdatedRecordShape(decoded any) any {
	return purchasepayments.PurchasePaymentUpdatedRecordShape(decoded)
}

// AccountPurchasePaymentsRecordShape is the DTOShape reading a body as AccountPurchasePaymentsRecord, for the raw-verb call.
func AccountPurchasePaymentsRecordShape(decoded any) any {
	return purchasepayments.AccountPurchasePaymentsRecordShape(decoded)
}

// PurchasePaymentNoteListShape is the DTOShape reading a body as PurchasePaymentNoteList, for the raw-verb call.
func PurchasePaymentNoteListShape(decoded any) any {
	return purchasepayments.PurchasePaymentNoteListShape(decoded)
}

// PurchasePaymentNotesRecordShape is the DTOShape reading a body as PurchasePaymentNotesRecord, for the raw-verb call.
func PurchasePaymentNotesRecordShape(decoded any) any {
	return purchasepayments.PurchasePaymentNotesRecordShape(decoded)
}

// PurchasePaymentNoteReadRecordShape is the DTOShape reading a body as PurchasePaymentNoteReadRecord, for the raw-verb call.
func PurchasePaymentNoteReadRecordShape(decoded any) any {
	return purchasepayments.PurchasePaymentNoteReadRecordShape(decoded)
}

// PurchasePaymentsPurchasePaymentNoteFileListShape is the DTOShape reading a body as PurchasePaymentsPurchasePaymentNoteFileList, for the raw-verb call.
func PurchasePaymentsPurchasePaymentNoteFileListShape(decoded any) any {
	return purchasepayments.PurchasePaymentNoteFileListShape(decoded)
}

// PurchasePaymentsPurchasePaymentNoteFileReadRecordShape is the DTOShape reading a body as PurchasePaymentsPurchasePaymentNoteFileReadRecord, for the raw-verb call.
func PurchasePaymentsPurchasePaymentNoteFileReadRecordShape(decoded any) any {
	return purchasepayments.PurchasePaymentNoteFileReadRecordShape(decoded)
}

// The purchase payments path templates and body keys, for the raw-verb call.
const (
	PurchasePaymentsPath                                = purchasepayments.PurchasePaymentsPath
	PurchasePaymentPath                                 = purchasepayments.PurchasePaymentPath
	PurchasePaymentsPurchaseInvoicePurchasePaymentsPath = purchasepayments.PurchaseInvoicePurchasePaymentsPath
	PurchasePaymentNotesPath                            = purchasepayments.PurchasePaymentNotesPath
	PurchasePaymentNotePath                             = purchasepayments.PurchasePaymentNotePath
	PurchasePaymentNoteFilesPath                        = purchasepayments.PurchasePaymentNoteFilesPath
	PurchasePaymentNoteFilePath                         = purchasepayments.PurchasePaymentNoteFilePath

	PurchasePaymentsKey = purchasepayments.PurchasePaymentsKey
	FilesKey            = purchasepayments.FilesKey
	PurchasePaymentKey  = purchasepayments.PurchasePaymentKey
)

// PurchaseRefunds is the purchase refunds endpoints bound to a client.
type PurchaseRefunds = purchaserefunds.PurchaseRefunds

// The purchase refunds views, aliased so a type assertion works under either spelling.
type (
	PurchaseRefundsPagination      = purchaserefunds.Pagination
	PurchaseRefundsRecord          = purchaserefunds.PurchaseRefundsRecord
	PurchaseRefundRecord           = purchaserefunds.PurchaseRefundRecord
	AccountPurchaseRefundsRecord   = purchaserefunds.AccountPurchaseRefundsRecord
	PurchaseRefundsAccount         = purchaserefunds.Account
	PurchaseRefundsCreatedBy       = purchaserefunds.CreatedBy
	PurchaseRefundsCreditNote      = purchaserefunds.CreditNote
	AccountPurchaseRefund          = purchaserefunds.AccountPurchaseRefund
	PurchaseRefundsCustomAttribute = purchaserefunds.CustomAttribute
	PurchaseRefundsPaginationLinks = purchaserefunds.PaginationLinks
	PurchaseRefundsPurchaseRefund  = purchaserefunds.PurchaseRefund
)

// NewPurchaseRefunds binds the resource to a requester of your own.
func NewPurchaseRefunds(requester Requester) *PurchaseRefunds { return purchaserefunds.New(requester) }

// PurchaseRefundsRecordShape is the DTOShape reading a body as PurchaseRefundsRecord, for the raw-verb call.
func PurchaseRefundsRecordShape(decoded any) any {
	return purchaserefunds.PurchaseRefundsRecordShape(decoded)
}

// PurchaseRefundRecordShape is the DTOShape reading a body as PurchaseRefundRecord, for the raw-verb call.
func PurchaseRefundRecordShape(decoded any) any {
	return purchaserefunds.PurchaseRefundRecordShape(decoded)
}

// AccountPurchaseRefundsRecordShape is the DTOShape reading a body as AccountPurchaseRefundsRecord, for the raw-verb call.
func AccountPurchaseRefundsRecordShape(decoded any) any {
	return purchaserefunds.AccountPurchaseRefundsRecordShape(decoded)
}

// The purchase refunds path templates and body keys, for the raw-verb call.
const (
	PurchaseRefundsPath = purchaserefunds.PurchaseRefundsPath
	PurchaseRefundPath  = purchaserefunds.PurchaseRefundPath

	PurchaseRefundsKey = purchaserefunds.PurchaseRefundsKey
	PurchaseRefundKey  = purchaserefunds.PurchaseRefundKey
)

// Refunds is the refunds endpoints bound to a client.
type Refunds = refunds.Refunds

// The refunds views, aliased so a type assertion works under either spelling.
type (
	RefundsPagination      = refunds.Pagination
	RefundRecord           = refunds.RefundRecord
	RefundList             = refunds.RefundList
	RefundUpdatedRecord    = refunds.RefundUpdatedRecord
	AccountRefundsRecord   = refunds.AccountRefundsRecord
	RefundsAccount         = refunds.Account
	RefundsCreditNote      = refunds.CreditNote
	RefundsPaginationLinks = refunds.PaginationLinks
	AccountRefund          = refunds.AccountRefund
)

// NewRefunds binds the resource to a requester of your own.
func NewRefunds(requester Requester) *Refunds { return refunds.New(requester) }

// RefundListShape is the DTOShape reading a body as RefundList, for the raw-verb call.
func RefundListShape(decoded any) any { return refunds.RefundListShape(decoded) }

// RefundRecordShape is the DTOShape reading a body as RefundRecord, for the raw-verb call.
func RefundRecordShape(decoded any) any { return refunds.RefundRecordShape(decoded) }

// RefundUpdatedRecordShape is the DTOShape reading a body as RefundUpdatedRecord, for the raw-verb call.
func RefundUpdatedRecordShape(decoded any) any { return refunds.RefundUpdatedRecordShape(decoded) }

// AccountRefundsRecordShape is the DTOShape reading a body as AccountRefundsRecord, for the raw-verb call.
func AccountRefundsRecordShape(decoded any) any { return refunds.AccountRefundsRecordShape(decoded) }

// The refunds path templates and body keys, for the raw-verb call.
const (
	RefundsPath           = refunds.RefundsPath
	RefundPath            = refunds.RefundPath
	CreditNoteRefundsPath = refunds.CreditNoteRefundsPath

	RefundsKey = refunds.RefundsKey
	RefundKey  = refunds.RefundKey
)

// RentalAssets is the rental assets endpoints bound to a client.
type RentalAssets = rentalassets.RentalAssets

// The rental assets views, aliased so a type assertion works under either spelling.
type (
	RentalAssetsPagination                = rentalassets.Pagination
	RentalAssetRecord                     = rentalassets.RentalAssetRecord
	RentalAssetList                       = rentalassets.RentalAssetList
	RentalAssetProfileRecord              = rentalassets.RentalAssetProfileRecord
	RentalAssetProfileList                = rentalassets.RentalAssetProfileList
	RentalAssetAvailabilityRecord         = rentalassets.RentalAssetAvailabilityRecord
	RentalAssetAvailabilityList           = rentalassets.RentalAssetAvailabilityList
	RentalAssetsCreatedBy                 = rentalassets.CreatedBy
	RentalAssetsCurrency                  = rentalassets.Currency
	RentalAssetsLastUpdatedBy             = rentalassets.LastUpdatedBy
	RentalAssetsPaginationLinks           = rentalassets.PaginationLinks
	RentalAsset                           = rentalassets.RentalAsset
	RentalAssetProfileReadRecord          = rentalassets.RentalAssetProfileReadRecord
	RentalAssetProfileSession             = rentalassets.RentalAssetProfileSession
	RentalAssetProfileSessionAvailability = rentalassets.RentalAssetProfileSessionAvailability
	RentalAssetProfileSessionPrice        = rentalassets.RentalAssetProfileSessionPrice
	RentalAssetProfileSessionWorkingHour  = rentalassets.RentalAssetProfileSessionWorkingHour
	ScheduleList                          = rentalassets.ScheduleList
	RentalAssetsCustomForm                = rentalassets.CustomForm
	RentalAssetCustomAvailability         = rentalassets.RentalAssetCustomAvailability
	RentalAssetProfile                    = rentalassets.RentalAssetProfile
	RentalAssetsSessionList               = rentalassets.SessionList
)

// NewRentalAssets binds the resource to a requester of your own.
func NewRentalAssets(requester Requester) *RentalAssets { return rentalassets.New(requester) }

// RentalAssetAvailabilityListShape is the DTOShape reading a body as RentalAssetAvailabilityList, for the raw-verb call.
func RentalAssetAvailabilityListShape(decoded any) any {
	return rentalassets.RentalAssetAvailabilityListShape(decoded)
}

// RentalAssetProfileListShape is the DTOShape reading a body as RentalAssetProfileList, for the raw-verb call.
func RentalAssetProfileListShape(decoded any) any {
	return rentalassets.RentalAssetProfileListShape(decoded)
}

// RentalAssetListShape is the DTOShape reading a body as RentalAssetList, for the raw-verb call.
func RentalAssetListShape(decoded any) any { return rentalassets.RentalAssetListShape(decoded) }

// RentalAssetRecordShape is the DTOShape reading a body as RentalAssetRecord, for the raw-verb call.
func RentalAssetRecordShape(decoded any) any { return rentalassets.RentalAssetRecordShape(decoded) }

// RentalAssetsRentalAssetProfileReadRecordShape is the DTOShape reading a body as RentalAssetsRentalAssetProfileReadRecord, for the raw-verb call.
func RentalAssetsRentalAssetProfileReadRecordShape(decoded any) any {
	return rentalassets.RentalAssetProfileReadRecordShape(decoded)
}

// The rental assets path templates and body keys, for the raw-verb call.
const (
	RentalAssetAvailabilityPath = rentalassets.RentalAssetAvailabilityPath
	RentalAssetProfilesPath     = rentalassets.RentalAssetProfilesPath
	RentalAssetsPath            = rentalassets.RentalAssetsPath
	RentalAssetProfilePath      = rentalassets.RentalAssetProfilePath
	RentalAssetPath             = rentalassets.RentalAssetPath
	RentalAssetCancelPath       = rentalassets.RentalAssetCancelPath
	RentalAssetReactivatePath   = rentalassets.RentalAssetReactivatePath

	RentalAssetsKey = rentalassets.RentalAssetsKey
	RentalAssetKey  = rentalassets.RentalAssetKey
)

// Reports is the reports endpoints bound to a client.
type Reports = reports.Reports

// The reports views, aliased so a type assertion works under either spelling.
type (
	ReportsPagination                                                     = reports.Pagination
	ReportsAccountsAccountOverviewRecord                                  = reports.ReportsAccountsAccountOverviewRecord
	Amended                                                               = reports.Amended
	Cancelled                                                             = reports.Cancelled
	Changed                                                               = reports.Changed
	Deleted                                                               = reports.Deleted
	InvoiceOverview                                                       = reports.InvoiceOverview
	InvoiceOverviewNew                                                    = reports.InvoiceOverviewNew
	ReportsInvoices                                                       = reports.Invoices
	NetGain                                                               = reports.NetGain
	NetPurchases                                                          = reports.NetPurchases
	OrderOverview                                                         = reports.OrderOverview
	OrderOverviewNew                                                      = reports.OrderOverviewNew
	ReportsOrders                                                         = reports.Orders
	ReportsPayment                                                        = reports.Payment
	PaymentAndRefundOverview                                              = reports.PaymentAndRefundOverview
	PaymentNew                                                            = reports.PaymentNew
	PurchaseInvoiceOverview                                               = reports.PurchaseInvoiceOverview
	PurchaseInvoiceOverviewNew                                            = reports.PurchaseInvoiceOverviewNew
	ReportsPurchaseInvoices                                               = reports.PurchaseInvoices
	PurchaseOrderOverview                                                 = reports.PurchaseOrderOverview
	PurchaseOrderOverviewNew                                              = reports.PurchaseOrderOverviewNew
	ReportsPurchaseOrders                                                 = reports.PurchaseOrders
	ReportsPurchasePayment                                                = reports.PurchasePayment
	PurchasePaymentAndRefundOverview                                      = reports.PurchasePaymentAndRefundOverview
	PurchasePaymentNew                                                    = reports.PurchasePaymentNew
	PurchaseRefund                                                        = reports.PurchaseRefund
	PurchaseRefundNew                                                     = reports.PurchaseRefundNew
	Reactivated                                                           = reports.Reactivated
	Refund                                                                = reports.Refund
	RefundNew                                                             = reports.RefundNew
	ReportsInvoicesInvoiceOverviewReadRecord                              = reports.ReportsInvoicesInvoiceOverviewReadRecord
	ReportsOrdersOrderOverviewReadRecord                                  = reports.ReportsOrdersOrderOverviewReadRecord
	ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecord              = reports.ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecord
	ReportsPurchaseOrdersPurchaseOrderOverviewReadRecord                  = reports.ReportsPurchaseOrdersPurchaseOrderOverviewReadRecord
	ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecord = reports.ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecord
	ReportsTransactionsPaymentAndRefundOverviewReadRecord                 = reports.ReportsTransactionsPaymentAndRefundOverviewReadRecord
	Voided                                                                = reports.Voided
	AccountOverview                                                       = reports.AccountOverview
	ReportsAccounts                                                       = reports.Accounts
)

// NewReports binds the resource to a requester of your own.
func NewReports(requester Requester) *Reports { return reports.New(requester) }

// ReportsAccountsAccountOverviewRecordShape is the DTOShape reading a body as ReportsAccountsAccountOverviewRecord, for the raw-verb call.
func ReportsAccountsAccountOverviewRecordShape(decoded any) any {
	return reports.ReportsAccountsAccountOverviewRecordShape(decoded)
}

// ReportsReportsInvoicesInvoiceOverviewReadRecordShape is the DTOShape reading a body as ReportsReportsInvoicesInvoiceOverviewReadRecord, for the raw-verb call.
func ReportsReportsInvoicesInvoiceOverviewReadRecordShape(decoded any) any {
	return reports.ReportsInvoicesInvoiceOverviewReadRecordShape(decoded)
}

// ReportsReportsOrdersOrderOverviewReadRecordShape is the DTOShape reading a body as ReportsReportsOrdersOrderOverviewReadRecord, for the raw-verb call.
func ReportsReportsOrdersOrderOverviewReadRecordShape(decoded any) any {
	return reports.ReportsOrdersOrderOverviewReadRecordShape(decoded)
}

// ReportsReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecordShape is the DTOShape reading a body as ReportsReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecord, for the raw-verb call.
func ReportsReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecordShape(decoded any) any {
	return reports.ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecordShape(decoded)
}

// ReportsReportsPurchaseOrdersPurchaseOrderOverviewReadRecordShape is the DTOShape reading a body as ReportsReportsPurchaseOrdersPurchaseOrderOverviewReadRecord, for the raw-verb call.
func ReportsReportsPurchaseOrdersPurchaseOrderOverviewReadRecordShape(decoded any) any {
	return reports.ReportsPurchaseOrdersPurchaseOrderOverviewReadRecordShape(decoded)
}

// ReportsReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecordShape is the DTOShape reading a body as ReportsReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecord, for the raw-verb call.
func ReportsReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecordShape(decoded any) any {
	return reports.ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecordShape(decoded)
}

// ReportsReportsTransactionsPaymentAndRefundOverviewReadRecordShape is the DTOShape reading a body as ReportsReportsTransactionsPaymentAndRefundOverviewReadRecord, for the raw-verb call.
func ReportsReportsTransactionsPaymentAndRefundOverviewReadRecordShape(decoded any) any {
	return reports.ReportsTransactionsPaymentAndRefundOverviewReadRecordShape(decoded)
}

// The reports path templates and body keys, for the raw-verb call.
const (
	ReportsAccountsAccountOverviewPath                              = reports.ReportsAccountsAccountOverviewPath
	ReportsInvoicesInvoiceOverviewPath                              = reports.ReportsInvoicesInvoiceOverviewPath
	ReportsOrdersOrderOverviewPath                                  = reports.ReportsOrdersOrderOverviewPath
	ReportsPurchaseInvoicesPurchaseInvoiceOverviewPath              = reports.ReportsPurchaseInvoicesPurchaseInvoiceOverviewPath
	ReportsPurchaseOrdersPurchaseOrderOverviewPath                  = reports.ReportsPurchaseOrdersPurchaseOrderOverviewPath
	ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewPath = reports.ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewPath
	ReportsTransactionsPaymentAndRefundOverviewPath                 = reports.ReportsTransactionsPaymentAndRefundOverviewPath

	ReportKey = reports.ReportKey
)

// ReturnMerchandiseAuthorisations is the return merchandise authorisations endpoints bound to a client.
type ReturnMerchandiseAuthorisations = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisations

// The return merchandise authorisations views, aliased so a type assertion works under either spelling.
type (
	ReturnMerchandiseAuthorisationsPagination          = returnmerchandiseauthorisations.Pagination
	ReturnMerchandiseAuthorisationRecord               = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationRecord
	ReturnMerchandiseAuthorisationList                 = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationList
	InvoiceReturnMerchandiseAuthorisationReceiveRecord = returnmerchandiseauthorisations.InvoiceReturnMerchandiseAuthorisationReceiveRecord
	InvoiceReturnMerchandiseAuthorisationReceiveList   = returnmerchandiseauthorisations.InvoiceReturnMerchandiseAuthorisationReceiveList
	ReturnMerchandiseAuthorisationReadRecord           = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationReadRecord
	ReturnMerchandiseAuthorisationsRecord              = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationsRecord
	ReturnMerchandiseAuthorisationsCreatedBy           = returnmerchandiseauthorisations.CreatedBy
	ReturnMerchandiseAuthorisationsPaginationLinks     = returnmerchandiseauthorisations.PaginationLinks
	ReturnMerchandiseAuthorisationsAccount             = returnmerchandiseauthorisations.Account
	ReturnMerchandiseAuthorisationsReceiveRma          = returnmerchandiseauthorisations.ReceiveRma
	ReturnMerchandiseAuthorisationsSaleInvoice         = returnmerchandiseauthorisations.SaleInvoice
)

// NewReturnMerchandiseAuthorisations binds the resource to a requester of your own.
func NewReturnMerchandiseAuthorisations(requester Requester) *ReturnMerchandiseAuthorisations {
	return returnmerchandiseauthorisations.New(requester)
}

// ReturnMerchandiseAuthorisationListShape is the DTOShape reading a body as ReturnMerchandiseAuthorisationList, for the raw-verb call.
func ReturnMerchandiseAuthorisationListShape(decoded any) any {
	return returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationListShape(decoded)
}

// ReturnMerchandiseAuthorisationsRecordShape is the DTOShape reading a body as ReturnMerchandiseAuthorisationsRecord, for the raw-verb call.
func ReturnMerchandiseAuthorisationsRecordShape(decoded any) any {
	return returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationsRecordShape(decoded)
}

// ReturnMerchandiseAuthorisationReadRecordShape is the DTOShape reading a body as ReturnMerchandiseAuthorisationReadRecord, for the raw-verb call.
func ReturnMerchandiseAuthorisationReadRecordShape(decoded any) any {
	return returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationReadRecordShape(decoded)
}

// InvoiceReturnMerchandiseAuthorisationReceiveListShape is the DTOShape reading a body as InvoiceReturnMerchandiseAuthorisationReceiveList, for the raw-verb call.
func InvoiceReturnMerchandiseAuthorisationReceiveListShape(decoded any) any {
	return returnmerchandiseauthorisations.InvoiceReturnMerchandiseAuthorisationReceiveListShape(decoded)
}

// InvoiceReturnMerchandiseAuthorisationReceiveRecordShape is the DTOShape reading a body as InvoiceReturnMerchandiseAuthorisationReceiveRecord, for the raw-verb call.
func InvoiceReturnMerchandiseAuthorisationReceiveRecordShape(decoded any) any {
	return returnmerchandiseauthorisations.InvoiceReturnMerchandiseAuthorisationReceiveRecordShape(decoded)
}

// The return merchandise authorisations path templates and body keys, for the raw-verb call.
const (
	ReturnMerchandiseAuthorisationsPath                    = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationsPath
	ReturnMerchandiseAuthorisationPath                     = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationPath
	InvoiceReturnMerchandiseAuthorisationReceivePath       = returnmerchandiseauthorisations.InvoiceReturnMerchandiseAuthorisationReceivePath
	InvoiceReturnMerchandiseAuthorisationReceiveByUUIDPath = returnmerchandiseauthorisations.InvoiceReturnMerchandiseAuthorisationReceiveByUUIDPath

	ReturnMerchandiseAuthorisationsKey        = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationsKey
	ReturnMerchandiseAuthorisationsInvoiceKey = returnmerchandiseauthorisations.InvoiceKey
)

// Settings is the settings endpoints bound to a client.
type Settings = settings.Settings

// The settings views, aliased so a type assertion works under either spelling.
type (
	SettingsPagination                         = settings.Pagination
	SettingsDiscountProfileRecord              = settings.SettingsDiscountProfileRecord
	SettingsDiscountProfileList                = settings.SettingsDiscountProfileList
	SettingsPaymentProcessorRecord             = settings.SettingsPaymentProcessorRecord
	SettingsPaymentProcessorList               = settings.SettingsPaymentProcessorList
	SettingsShippingProfileRecord              = settings.SettingsShippingProfileRecord
	SettingsShippingProfileList                = settings.SettingsShippingProfileList
	SettingsPricingLevelRecord                 = settings.SettingsPricingLevelRecord
	SettingsPricingLevelList                   = settings.SettingsPricingLevelList
	SettingsTimeZoneRecord                     = settings.SettingsTimeZoneRecord
	SettingsTimeZoneList                       = settings.SettingsTimeZoneList
	SettingsWarehouseRecord                    = settings.SettingsWarehouseRecord
	SettingsWarehouseList                      = settings.SettingsWarehouseList
	SettingsCurrencyRecord                     = settings.SettingsCurrencyRecord
	SettingsCurrencyList                       = settings.SettingsCurrencyList
	SettingsItemGroupRecord                    = settings.SettingsItemGroupRecord
	SettingsItemGroupList                      = settings.SettingsItemGroupList
	CustomObjectRecord                         = settings.CustomObjectRecord
	CustomObjectList                           = settings.CustomObjectList
	SettingsTaxRecord                          = settings.SettingsTaxRecord
	SettingsTaxList                            = settings.SettingsTaxList
	SettingsVariationRecord                    = settings.SettingsVariationRecord
	SettingsVariationList                      = settings.SettingsVariationList
	SettingsPaymentTermRecord                  = settings.SettingsPaymentTermRecord
	SettingsPaymentTermList                    = settings.SettingsPaymentTermList
	SettingsItemGroupsRecord                   = settings.SettingsItemGroupsRecord
	SettingsCustomLogicScriptRecord            = settings.SettingsCustomLogicScriptRecord
	SettingsCurrencyReadRecord                 = settings.SettingsCurrencyReadRecord
	SettingsCustomObjectCustomizeRecord        = settings.SettingsCustomObjectCustomizeRecord
	SettingsCustomObjectsRecord                = settings.SettingsCustomObjectsRecord
	SettingsVariationsRecord                   = settings.SettingsVariationsRecord
	SettingsTimeZoneReadRecord                 = settings.SettingsTimeZoneReadRecord
	SettingsCustomLogicScriptsRecord           = settings.SettingsCustomLogicScriptsRecord
	SettingsRecord                             = settings.SettingsRecord
	SettingsFieldSettingsItemRecord            = settings.SettingsFieldSettingsItemRecord
	SettingsAccount                            = settings.Account
	AccountGrouping                            = settings.AccountGrouping
	AccountManagement                          = settings.AccountManagement
	AccountingCode                             = settings.AccountingCode
	AdditionalContact1                         = settings.AdditionalContact1
	AdditionalContact2                         = settings.AdditionalContact2
	AdditionalContact3                         = settings.AdditionalContact3
	AvailableChargeTypes                       = settings.AvailableChargeTypes
	AvailablePricingMethods                    = settings.AvailablePricingMethods
	BillingContact                             = settings.BillingContact
	CreationOfSubsidiaryAccount                = settings.CreationOfSubsidiaryAccount
	SettingsData                               = settings.Data
	InventoryAdjustmentRules                   = settings.InventoryAdjustmentRules
	SettingsItem                               = settings.Item
	ItemGrouping                               = settings.ItemGrouping
	ItemManagement                             = settings.ItemManagement
	SettingsPaginationLinks                    = settings.PaginationLinks
	PreOrder                                   = settings.PreOrder
	PricingPeriodCanVaryFromChargingPeriod     = settings.PricingPeriodCanVaryFromChargingPeriod
	Product                                    = settings.Product
	PurchasePriceEnteredIsInclusiveOfTax       = settings.PurchasePriceEnteredIsInclusiveOfTax
	ReferralTrackingAtAccountLevel             = settings.ReferralTrackingAtAccountLevel
	SalePriceEnteredIsInclusiveOfTax           = settings.SalePriceEnteredIsInclusiveOfTax
	SalesRepTrackingAtAccountLevel             = settings.SalesRepTrackingAtAccountLevel
	SettingsGeneralSettingsAccountUpdateRecord = settings.SettingsGeneralSettingsAccountUpdateRecord
	SettingsGeneralSettingsItemUpdateRecord    = settings.SettingsGeneralSettingsItemUpdateRecord
	SettingsItemGroupReadRecord                = settings.SettingsItemGroupReadRecord
	ShippingContact                            = settings.ShippingContact
	SettingsTax                                = settings.Tax
	UseVolumePricing                           = settings.UseVolumePricing
	SettingsAttribute                          = settings.Attribute
	SettingsConfiguration                      = settings.Configuration
	Country                                    = settings.Country
	CountryTaxByState                          = settings.CountryTaxByState
	SettingsCurrency                           = settings.Currency
	Duration                                   = settings.Duration
	Resource                                   = settings.Resource
	SettingsTaxCode                            = settings.TaxCode
	TaxRule                                    = settings.TaxRule
)

// NewSettings binds the resource to a requester of your own.
func NewSettings(requester Requester) *Settings { return settings.New(requester) }

// CustomObjectListShape is the DTOShape reading a body as CustomObjectList, for the raw-verb call.
func CustomObjectListShape(decoded any) any { return settings.CustomObjectListShape(decoded) }

// SettingsRecordShape is the DTOShape reading a body as SettingsRecord, for the raw-verb call.
func SettingsRecordShape(decoded any) any { return settings.SettingsRecordShape(decoded) }

// SettingsCustomObjectsRecordShape is the DTOShape reading a body as SettingsCustomObjectsRecord, for the raw-verb call.
func SettingsCustomObjectsRecordShape(decoded any) any {
	return settings.SettingsCustomObjectsRecordShape(decoded)
}

// SettingsCurrencyListShape is the DTOShape reading a body as SettingsCurrencyList, for the raw-verb call.
func SettingsCurrencyListShape(decoded any) any { return settings.SettingsCurrencyListShape(decoded) }

// SettingsCustomLogicScriptsRecordShape is the DTOShape reading a body as SettingsCustomLogicScriptsRecord, for the raw-verb call.
func SettingsCustomLogicScriptsRecordShape(decoded any) any {
	return settings.SettingsCustomLogicScriptsRecordShape(decoded)
}

// SettingsDiscountProfileListShape is the DTOShape reading a body as SettingsDiscountProfileList, for the raw-verb call.
func SettingsDiscountProfileListShape(decoded any) any {
	return settings.SettingsDiscountProfileListShape(decoded)
}

// SettingsItemGroupListShape is the DTOShape reading a body as SettingsItemGroupList, for the raw-verb call.
func SettingsItemGroupListShape(decoded any) any { return settings.SettingsItemGroupListShape(decoded) }

// SettingsItemGroupsRecordShape is the DTOShape reading a body as SettingsItemGroupsRecord, for the raw-verb call.
func SettingsItemGroupsRecordShape(decoded any) any {
	return settings.SettingsItemGroupsRecordShape(decoded)
}

// SettingsPaymentProcessorListShape is the DTOShape reading a body as SettingsPaymentProcessorList, for the raw-verb call.
func SettingsPaymentProcessorListShape(decoded any) any {
	return settings.SettingsPaymentProcessorListShape(decoded)
}

// SettingsPaymentTermListShape is the DTOShape reading a body as SettingsPaymentTermList, for the raw-verb call.
func SettingsPaymentTermListShape(decoded any) any {
	return settings.SettingsPaymentTermListShape(decoded)
}

// SettingsPricingLevelListShape is the DTOShape reading a body as SettingsPricingLevelList, for the raw-verb call.
func SettingsPricingLevelListShape(decoded any) any {
	return settings.SettingsPricingLevelListShape(decoded)
}

// SettingsShippingProfileListShape is the DTOShape reading a body as SettingsShippingProfileList, for the raw-verb call.
func SettingsShippingProfileListShape(decoded any) any {
	return settings.SettingsShippingProfileListShape(decoded)
}

// SettingsTaxListShape is the DTOShape reading a body as SettingsTaxList, for the raw-verb call.
func SettingsTaxListShape(decoded any) any { return settings.SettingsTaxListShape(decoded) }

// SettingsTimeZoneListShape is the DTOShape reading a body as SettingsTimeZoneList, for the raw-verb call.
func SettingsTimeZoneListShape(decoded any) any { return settings.SettingsTimeZoneListShape(decoded) }

// SettingsVariationListShape is the DTOShape reading a body as SettingsVariationList, for the raw-verb call.
func SettingsVariationListShape(decoded any) any { return settings.SettingsVariationListShape(decoded) }

// SettingsVariationsRecordShape is the DTOShape reading a body as SettingsVariationsRecord, for the raw-verb call.
func SettingsVariationsRecordShape(decoded any) any {
	return settings.SettingsVariationsRecordShape(decoded)
}

// SettingsWarehouseListShape is the DTOShape reading a body as SettingsWarehouseList, for the raw-verb call.
func SettingsWarehouseListShape(decoded any) any { return settings.SettingsWarehouseListShape(decoded) }

// SettingsCurrencyReadRecordShape is the DTOShape reading a body as SettingsCurrencyReadRecord, for the raw-verb call.
func SettingsCurrencyReadRecordShape(decoded any) any {
	return settings.SettingsCurrencyReadRecordShape(decoded)
}

// SettingsCustomLogicScriptRecordShape is the DTOShape reading a body as SettingsCustomLogicScriptRecord, for the raw-verb call.
func SettingsCustomLogicScriptRecordShape(decoded any) any {
	return settings.SettingsCustomLogicScriptRecordShape(decoded)
}

// SettingsFieldSettingsItemRecordShape is the DTOShape reading a body as SettingsFieldSettingsItemRecord, for the raw-verb call.
func SettingsFieldSettingsItemRecordShape(decoded any) any {
	return settings.SettingsFieldSettingsItemRecordShape(decoded)
}

// SettingsTimeZoneReadRecordShape is the DTOShape reading a body as SettingsTimeZoneReadRecord, for the raw-verb call.
func SettingsTimeZoneReadRecordShape(decoded any) any {
	return settings.SettingsTimeZoneReadRecordShape(decoded)
}

// SettingsCustomObjectCustomizeRecordShape is the DTOShape reading a body as SettingsCustomObjectCustomizeRecord, for the raw-verb call.
func SettingsCustomObjectCustomizeRecordShape(decoded any) any {
	return settings.SettingsCustomObjectCustomizeRecordShape(decoded)
}

// SettingsSettingsGeneralSettingsAccountUpdateRecordShape is the DTOShape reading a body as SettingsSettingsGeneralSettingsAccountUpdateRecord, for the raw-verb call.
func SettingsSettingsGeneralSettingsAccountUpdateRecordShape(decoded any) any {
	return settings.SettingsGeneralSettingsAccountUpdateRecordShape(decoded)
}

// SettingsSettingsGeneralSettingsItemUpdateRecordShape is the DTOShape reading a body as SettingsSettingsGeneralSettingsItemUpdateRecord, for the raw-verb call.
func SettingsSettingsGeneralSettingsItemUpdateRecordShape(decoded any) any {
	return settings.SettingsGeneralSettingsItemUpdateRecordShape(decoded)
}

// SettingsSettingsItemGroupReadRecordShape is the DTOShape reading a body as SettingsSettingsItemGroupReadRecord, for the raw-verb call.
func SettingsSettingsItemGroupReadRecordShape(decoded any) any {
	return settings.SettingsItemGroupReadRecordShape(decoded)
}

// The settings path templates and body keys, for the raw-verb call.
const (
	CustomObjectsPath                       = settings.CustomObjectsPath
	SettingsPath                            = settings.SettingsPath
	CustomObjectPath                        = settings.CustomObjectPath
	SettingsCurrenciesPath                  = settings.SettingsCurrenciesPath
	SettingsCustomLogicScriptsPath          = settings.SettingsCustomLogicScriptsPath
	SettingsCustomObjectsPath               = settings.SettingsCustomObjectsPath
	SettingsDiscountProfilesPath            = settings.SettingsDiscountProfilesPath
	SettingsItemGroupsPath                  = settings.SettingsItemGroupsPath
	SettingsPaymentProcessorsPath           = settings.SettingsPaymentProcessorsPath
	SettingsPaymentTermsPath                = settings.SettingsPaymentTermsPath
	SettingsPricingLevelsPath               = settings.SettingsPricingLevelsPath
	SettingsShippingProfilesPath            = settings.SettingsShippingProfilesPath
	SettingsTaxesPath                       = settings.SettingsTaxesPath
	SettingsTimeZonesPath                   = settings.SettingsTimeZonesPath
	SettingsVariationsPath                  = settings.SettingsVariationsPath
	SettingsWarehousesPath                  = settings.SettingsWarehousesPath
	SettingsCurrencyPath                    = settings.SettingsCurrencyPath
	SettingsCustomLogicScriptPath           = settings.SettingsCustomLogicScriptPath
	SettingsCustomObjectPath                = settings.SettingsCustomObjectPath
	SettingsFieldSettingsAccountPath        = settings.SettingsFieldSettingsAccountPath
	SettingsFieldSettingsItemPath           = settings.SettingsFieldSettingsItemPath
	SettingsGeneralSettingsAccountPath      = settings.SettingsGeneralSettingsAccountPath
	SettingsGeneralSettingsItemPath         = settings.SettingsGeneralSettingsItemPath
	SettingsItemGroupPath                   = settings.SettingsItemGroupPath
	SettingsTimeZonePath                    = settings.SettingsTimeZonePath
	SettingsVariationPath                   = settings.SettingsVariationPath
	SettingsCustomObjectCustomizePath       = settings.SettingsCustomObjectCustomizePath
	SettingsItemGroupBillingPreferencesPath = settings.SettingsItemGroupBillingPreferencesPath

	SettingsKey          = settings.SettingsKey
	CurrencyKey          = settings.CurrencyKey
	CustomLogicScriptKey = settings.CustomLogicScriptKey
	CustomObjectKey      = settings.CustomObjectKey
	ItemGroupKey         = settings.ItemGroupKey
	TimeZoneKey          = settings.TimeZoneKey
	VariationKey         = settings.VariationKey
)

// Tasks is the tasks endpoints bound to a client.
type Tasks = tasks.Tasks

// The tasks views, aliased so a type assertion works under either spelling.
type (
	TasksPagination               = tasks.Pagination
	TaskRecord                    = tasks.TaskRecord
	TaskList                      = tasks.TaskList
	TasksActivityLog              = tasks.ActivityLog
	Assignees                     = tasks.Assignees
	Checklist                     = tasks.Checklist
	CompletedBy                   = tasks.CompletedBy
	TasksCreatedBy                = tasks.CreatedBy
	TasksLastUpdatedBy            = tasks.LastUpdatedBy
	TasksPaginationLinks          = tasks.PaginationLinks
	SelectedEmployee              = tasks.SelectedEmployee
	SelectedLabour                = tasks.SelectedLabour
	TaskChecklistReadRecord       = tasks.TaskChecklistReadRecord
	TaskNoteList                  = tasks.TaskNoteList
	TaskReadRecord                = tasks.TaskReadRecord
	TasksEmail                    = tasks.Email
	TasksName                     = tasks.Name
	TaskChecklistByUUIDReadRecord = tasks.TaskChecklistByUUIDReadRecord
)

// NewTasks binds the resource to a requester of your own.
func NewTasks(requester Requester) *Tasks { return tasks.New(requester) }

// TaskListShape is the DTOShape reading a body as TaskList, for the raw-verb call.
func TaskListShape(decoded any) any { return tasks.TaskListShape(decoded) }

// TasksTaskChecklistReadRecordShape is the DTOShape reading a body as TasksTaskChecklistReadRecord, for the raw-verb call.
func TasksTaskChecklistReadRecordShape(decoded any) any {
	return tasks.TaskChecklistReadRecordShape(decoded)
}

// TasksTaskNoteListShape is the DTOShape reading a body as TasksTaskNoteList, for the raw-verb call.
func TasksTaskNoteListShape(decoded any) any { return tasks.TaskNoteListShape(decoded) }

// TasksTaskReadRecordShape is the DTOShape reading a body as TasksTaskReadRecord, for the raw-verb call.
func TasksTaskReadRecordShape(decoded any) any { return tasks.TaskReadRecordShape(decoded) }

// TasksTaskChecklistByUUIDReadRecordShape is the DTOShape reading a body as TasksTaskChecklistByUUIDReadRecord, for the raw-verb call.
func TasksTaskChecklistByUUIDReadRecordShape(decoded any) any {
	return tasks.TaskChecklistByUUIDReadRecordShape(decoded)
}

// The tasks path templates and body keys, for the raw-verb call.
const (
	TasksPath               = tasks.TasksPath
	TaskPath                = tasks.TaskPath
	TaskChecklistPath       = tasks.TaskChecklistPath
	TaskNotesPath           = tasks.TaskNotesPath
	TaskChecklistByUUIDPath = tasks.TaskChecklistByUUIDPath
	TaskNotePath            = tasks.TaskNotePath

	TasksKey = tasks.TasksKey
)

// Timesheets is the timesheets endpoints bound to a client.
type Timesheets = timesheets.Timesheets

// The timesheets views, aliased so a type assertion works under either spelling.
type (
	TimesheetsPagination      = timesheets.Pagination
	TimesheetRecord           = timesheets.TimesheetRecord
	TimesheetList             = timesheets.TimesheetList
	TimesheetsActivityLog     = timesheets.ActivityLog
	TimesheetsCreatedBy       = timesheets.CreatedBy
	TimesheetsLabour          = timesheets.Labour
	TimesheetsLastUpdatedBy   = timesheets.LastUpdatedBy
	TimesheetsPaginationLinks = timesheets.PaginationLinks
	SaleOrder                 = timesheets.SaleOrder
	SaleOrderItem             = timesheets.SaleOrderItem
	TimesheetCreateRecord     = timesheets.TimesheetCreateRecord
)

// NewTimesheets binds the resource to a requester of your own.
func NewTimesheets(requester Requester) *Timesheets { return timesheets.New(requester) }

// TimesheetListShape is the DTOShape reading a body as TimesheetList, for the raw-verb call.
func TimesheetListShape(decoded any) any { return timesheets.TimesheetListShape(decoded) }

// TimesheetsTimesheetCreateRecordShape is the DTOShape reading a body as TimesheetsTimesheetCreateRecord, for the raw-verb call.
func TimesheetsTimesheetCreateRecordShape(decoded any) any {
	return timesheets.TimesheetCreateRecordShape(decoded)
}

// The timesheets path templates and body keys, for the raw-verb call.
const (
	TimesheetsPath = timesheets.TimesheetsPath
	TimesheetPath  = timesheets.TimesheetPath

	TimesheetsKey = timesheets.TimesheetsKey
)

// Usages is the usages endpoints bound to a client.
type Usages = usages.Usages

// The usages views, aliased so a type assertion works under either spelling.
type (
	UsagesPagination      = usages.Pagination
	UsageRecord           = usages.UsageRecord
	UsageList             = usages.UsageList
	UsagesRecord          = usages.UsagesRecord
	UsageUpdatedRecord    = usages.UsageUpdatedRecord
	UsagesPaginationLinks = usages.PaginationLinks
	ChargeItem            = usages.ChargeItem
	UsagesCreatedBy       = usages.CreatedBy
	UsagesCustomAttribute = usages.CustomAttribute
	UsagesLastUpdatedBy   = usages.LastUpdatedBy
)

// NewUsages binds the resource to a requester of your own.
func NewUsages(requester Requester) *Usages { return usages.New(requester) }

// UsageListShape is the DTOShape reading a body as UsageList, for the raw-verb call.
func UsageListShape(decoded any) any { return usages.UsageListShape(decoded) }

// UsagesRecordShape is the DTOShape reading a body as UsagesRecord, for the raw-verb call.
func UsagesRecordShape(decoded any) any { return usages.UsagesRecordShape(decoded) }

// UsageRecordShape is the DTOShape reading a body as UsageRecord, for the raw-verb call.
func UsageRecordShape(decoded any) any { return usages.UsageRecordShape(decoded) }

// UsageUpdatedRecordShape is the DTOShape reading a body as UsageUpdatedRecord, for the raw-verb call.
func UsageUpdatedRecordShape(decoded any) any { return usages.UsageUpdatedRecordShape(decoded) }

// The usages path templates and body keys, for the raw-verb call.
const (
	UsagesPath = usages.UsagesPath
	UsagePath  = usages.UsagePath

	UsagesKey = usages.UsagesKey
	UsageKey  = usages.UsageKey
)

// Workflows is the workflows endpoints bound to a client.
type Workflows = workflows.Workflows

// The workflows views, aliased so a type assertion works under either spelling.
type (
	ActivityDefinition       = workflows.ActivityDefinition
	WorkflowsPagination      = workflows.Pagination
	WorkflowsPaginationLinks = workflows.PaginationLinks
	Workflow                 = workflows.Workflow
	WorkflowMilestoneList    = workflows.WorkflowMilestoneList
	WorkflowMilestoneRecord  = workflows.WorkflowMilestoneRecord
	WorkflowVersion          = workflows.WorkflowVersion
)

// NewWorkflows binds the resource to a requester of your own.
func NewWorkflows(requester Requester) *Workflows { return workflows.New(requester) }

// WorkflowsWorkflowMilestoneListShape is the DTOShape reading a body as WorkflowsWorkflowMilestoneList, for the raw-verb call.
func WorkflowsWorkflowMilestoneListShape(decoded any) any {
	return workflows.WorkflowMilestoneListShape(decoded)
}

// The workflows path templates and body keys, for the raw-verb call.
const (
	MilestonePath                              = workflows.MilestonePath
	WorkflowComponentByIDPath                  = workflows.WorkflowComponentByIDPath
	WorkflowMilestonesPath                     = workflows.WorkflowMilestonesPath
	WorkflowComponentCcMilestonePath           = workflows.WorkflowComponentCcMilestonePath
	WorkflowComponentHistoryPath               = workflows.WorkflowComponentHistoryPath
	WorkflowComponentCcInstanceCcMilestonePath = workflows.WorkflowComponentCcInstanceCcMilestonePath

	AccountsFileKey                                                  = accounts.FileKey
	CustomComponentObjectKey                                         = customcomponents.CustomComponentObjectKey
	CustomFeatureKey                                                 = customfeatures.CustomFeatureKey
	ExpressKey                                                       = express.ExpressKey
	ExternalDatabaseKey                                              = externaldatabases.ExternalDatabaseKey
	InvoicesFileKey                                                  = invoices.FileKey
	InvoicesItemFulfillmentKey                                       = invoices.ItemFulfillmentKey
	InvoicesNoteKey                                                  = invoices.NoteKey
	InvoicesReturnMerchandiseAuthorisationKey                        = invoices.ReturnMerchandiseAuthorisationKey
	ItemFulfillmentsNoteKey                                          = itemfulfillments.NoteKey
	ItemsFileKey                                                     = items.FileKey
	LabourProfileKey                                                 = labours.LabourProfileKey
	OrdersFileKey                                                    = orders.FileKey
	OrdersNoteKey                                                    = orders.NoteKey
	OrderLineKey                                                     = orders.OrderLineKey
	PaymentsFileKey                                                  = payments.FileKey
	PaymentsNoteKey                                                  = payments.NoteKey
	PurchaseInvoicesFileKey                                          = purchaseinvoices.FileKey
	PurchaseInvoicesNoteKey                                          = purchaseinvoices.NoteKey
	PurchaseOrdersFileKey                                            = purchaseorders.FileKey
	PurchaseOrdersNoteKey                                            = purchaseorders.NoteKey
	PurchaseOrdersPurchaseInvoiceKey                                 = purchaseorders.PurchaseInvoiceKey
	PurchaseOrderLineKey                                             = purchaseorders.PurchaseOrderLineKey
	PurchasePaymentsFileKey                                          = purchasepayments.FileKey
	PurchasePaymentsNoteKey                                          = purchasepayments.NoteKey
	RentalAssetProfileKey                                            = rentalassets.RentalAssetProfileKey
	ReceiveRmaKey                                                    = returnmerchandiseauthorisations.ReceiveRmaKey
	ReturnMerchandiseAuthorisationsReturnMerchandiseAuthorisationKey = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationKey
	GeneralSettingsKey                                               = settings.GeneralSettingsKey
	SettingsEnvelopeKey                                              = settings.SettingsEnvelopeKey
	WorkTaskKey                                                      = tasks.WorkTaskKey
	TimesheetKey                                                     = timesheets.TimesheetKey
)

// AccountGroups is the account groups endpoints bound to a client.
type AccountGroups = accountgroups.AccountGroups

// The account groups views, aliased so a type assertion works under either spelling.
type (
	SettingsAccountGroupList      = accountgroups.SettingsAccountGroupList
	SettingsAccountGroupRecord    = accountgroups.SettingsAccountGroupRecord
	SettingsAccountGroupStateRule = accountgroups.SettingsAccountGroupStateRule
	AccountGroupsPagination       = accountgroups.Pagination
	AccountGroupsPaginationLinks  = accountgroups.PaginationLinks
)

// NewAccountGroups binds the resource to a requester of your own.
func NewAccountGroups(requester Requester) *AccountGroups { return accountgroups.New(requester) }

// SettingsAccountGroupListShape is the DTOShape reading a body as SettingsAccountGroupList, for the raw-verb call.
func SettingsAccountGroupListShape(decoded any) any {
	return accountgroups.SettingsAccountGroupListShape(decoded)
}

// SettingsAccountGroupRecordShape is the DTOShape reading a body as SettingsAccountGroupRecord, for the raw-verb call.
func SettingsAccountGroupRecordShape(decoded any) any {
	return accountgroups.SettingsAccountGroupRecordShape(decoded)
}

// The account groups path templates and body keys, for the raw-verb call.
const (
	SettingsAccountGroupsPath = accountgroups.SettingsAccountGroupsPath
	SettingsAccountGroupPath  = accountgroups.SettingsAccountGroupPath

	SettingsAccountGroupsKey = accountgroups.SettingsAccountGroupsKey
	AccountGroupKey          = accountgroups.AccountGroupKey
)

// More accounts views, aliased so a type assertion works under either spelling.
type (
	AccountKpisRecord             = accounts.AccountKpisRecord
	AccountFinancialSummaryRecord = accounts.AccountFinancialSummaryRecord
	AccountFinancialSummaryDetail = accounts.AccountFinancialSummaryDetail
	AccountCustomObjectView       = accounts.AccountCustomObjectView
	AccountComponentObjectRecord  = accounts.AccountComponentObjectRecord
)

// AccountKpisRecordShape is the DTOShape reading a body as AccountKpisRecord, for the raw-verb call.
func AccountKpisRecordShape(decoded any) any { return accounts.AccountKpisRecordShape(decoded) }

// AccountFinancialSummaryRecordShape is the DTOShape reading a body as AccountFinancialSummaryRecord, for the raw-verb call.
func AccountFinancialSummaryRecordShape(decoded any) any {
	return accounts.AccountFinancialSummaryRecordShape(decoded)
}

// AccountCustomObjectViewShape is the DTOShape reading a body as AccountCustomObjectView, for the raw-verb call.
func AccountCustomObjectViewShape(decoded any) any {
	return accounts.AccountCustomObjectViewShape(decoded)
}

// AccountComponentObjectRecordShape is the DTOShape reading a body as AccountComponentObjectRecord, for the raw-verb call.
func AccountComponentObjectRecordShape(decoded any) any {
	return accounts.AccountComponentObjectRecordShape(decoded)
}

// More accounts path templates and body keys, for the raw-verb call.
const (
	AccountReactivatePath             = accounts.AccountReactivatePath
	AccountCustomObjectPath           = accounts.AccountCustomObjectPath
	AccountCustomObjectInstancePath   = accounts.AccountCustomObjectInstancePath
	AccountPurgePath                  = accounts.AccountPurgePath
	AccountKPIsPath                   = accounts.AccountKPIsPath
	AccountFinancialSummaryPath       = accounts.AccountFinancialSummaryPath
	AccountCustomObjectInstanceOfPath = accounts.AccountCustomObjectInstanceOfPath
	AccountComponentObjectsPath       = accounts.AccountComponentObjectsPath
	AccountComponentObjectPath        = accounts.AccountComponentObjectPath

	AccountKPIsKey                   = accounts.AccountKPIsKey
	AccountsCustomObjectKey          = accounts.CustomObjectKey
	AccountsCustomComponentObjectKey = accounts.CustomComponentObjectKey
)

// ActionScripts is the action scripts endpoints bound to a client.
type ActionScripts = actionscripts.ActionScripts

// The action scripts views, aliased so a type assertion works under either spelling.
type (
	ActionScriptList             = actionscripts.ActionScriptList
	ActionScriptRecord           = actionscripts.ActionScriptRecord
	ActionScriptCreatedBy        = actionscripts.ActionScriptCreatedBy
	ActionScriptLastUpdatedBy    = actionscripts.ActionScriptLastUpdatedBy
	ActionScriptsPagination      = actionscripts.Pagination
	ActionScriptsPaginationLinks = actionscripts.PaginationLinks
)

// NewActionScripts binds the resource to a requester of your own.
func NewActionScripts(requester Requester) *ActionScripts { return actionscripts.New(requester) }

// ActionScriptListShape is the DTOShape reading a body as ActionScriptList, for the raw-verb call.
func ActionScriptListShape(decoded any) any { return actionscripts.ActionScriptListShape(decoded) }

// ActionScriptRecordShape is the DTOShape reading a body as ActionScriptRecord, for the raw-verb call.
func ActionScriptRecordShape(decoded any) any { return actionscripts.ActionScriptRecordShape(decoded) }

// The action scripts path templates and body keys, for the raw-verb call.
const (
	ActionScriptsPath = actionscripts.ActionScriptsPath
	ActionScriptPath  = actionscripts.ActionScriptPath

	ActionScriptsKey = actionscripts.ActionScriptsKey
	ActionScriptKey  = actionscripts.ActionScriptKey
)

// AdditionalCodes is the additional codes endpoints bound to a client.
type AdditionalCodes = additionalcodes.AdditionalCodes

// The additional codes views, aliased so a type assertion works under either spelling.
type (
	SettingsAdditionalCodeList     = additionalcodes.SettingsAdditionalCodeList
	SettingsAdditionalCodeRecord   = additionalcodes.SettingsAdditionalCodeRecord
	AdditionalCodesPagination      = additionalcodes.Pagination
	AdditionalCodesPaginationLinks = additionalcodes.PaginationLinks
)

// NewAdditionalCodes binds the resource to a requester of your own.
func NewAdditionalCodes(requester Requester) *AdditionalCodes { return additionalcodes.New(requester) }

// SettingsAdditionalCodeListShape is the DTOShape reading a body as SettingsAdditionalCodeList, for the raw-verb call.
func SettingsAdditionalCodeListShape(decoded any) any {
	return additionalcodes.SettingsAdditionalCodeListShape(decoded)
}

// SettingsAdditionalCodeRecordShape is the DTOShape reading a body as SettingsAdditionalCodeRecord, for the raw-verb call.
func SettingsAdditionalCodeRecordShape(decoded any) any {
	return additionalcodes.SettingsAdditionalCodeRecordShape(decoded)
}

// The additional codes path templates and body keys, for the raw-verb call.
const (
	SettingsAdditionalCodesPath = additionalcodes.SettingsAdditionalCodesPath
	SettingsAdditionalCodePath  = additionalcodes.SettingsAdditionalCodePath

	SettingsAdditionalCodesKey = additionalcodes.SettingsAdditionalCodesKey
	AdditionalCodeKey          = additionalcodes.AdditionalCodeKey
)

// AlertResponsibilities is the alert responsibilities endpoints bound to a client.
type AlertResponsibilities = alertresponsibilities.AlertResponsibilities

// NewAlertResponsibilities binds the resource to a requester of your own.
func NewAlertResponsibilities(requester Requester) *AlertResponsibilities {
	return alertresponsibilities.New(requester)
}

// The alert responsibilities path templates and body keys, for the raw-verb call.
const (
	SettingsAlertResponsibilitiesPath = alertresponsibilities.SettingsAlertResponsibilitiesPath
	SettingsAlertResponsibilityPath   = alertresponsibilities.SettingsAlertResponsibilityPath

	SettingsAlertResponsibilitiesKey            = alertresponsibilities.SettingsAlertResponsibilitiesKey
	AlertResponsibilitiesAlertResponsibilityKey = alertresponsibilities.AlertResponsibilityKey
)

// AlertResponsibility is the alert responsibility endpoints bound to a client.
type AlertResponsibility = alertresponsibility.AlertResponsibility

// The alert responsibility views, aliased so a type assertion works under either spelling.
type (
	AlertResponsibilityList            = alertresponsibility.AlertResponsibilityList
	AlertResponsibilityRecord          = alertresponsibility.AlertResponsibilityRecord
	AlertResponsibilityPagination      = alertresponsibility.Pagination
	AlertResponsibilityPaginationLinks = alertresponsibility.PaginationLinks
)

// NewAlertResponsibility binds the resource to a requester of your own.
func NewAlertResponsibility(requester Requester) *AlertResponsibility {
	return alertresponsibility.New(requester)
}

// AlertResponsibilityListShape is the DTOShape reading a body as AlertResponsibilityList, for the raw-verb call.
func AlertResponsibilityListShape(decoded any) any {
	return alertresponsibility.AlertResponsibilityListShape(decoded)
}

// AlertResponsibilityRecordShape is the DTOShape reading a body as AlertResponsibilityRecord, for the raw-verb call.
func AlertResponsibilityRecordShape(decoded any) any {
	return alertresponsibility.AlertResponsibilityRecordShape(decoded)
}

// The alert responsibility path templates and body keys, for the raw-verb call.
const (
	AlertResponsibilitiesPath = alertresponsibility.AlertResponsibilitiesPath
	AlertResponsibilityPath   = alertresponsibility.AlertResponsibilityPath

	AlertResponsibilitiesKey                  = alertresponsibility.AlertResponsibilitiesKey
	AlertResponsibilityAlertResponsibilityKey = alertresponsibility.AlertResponsibilityKey
)

// Applications is the applications endpoints bound to a client.
type Applications = applications.Applications

// The applications views, aliased so a type assertion works under either spelling.
type (
	SettingsApplicationList     = applications.SettingsApplicationList
	SettingsApplicationRecord   = applications.SettingsApplicationRecord
	ApplicationsPagination      = applications.Pagination
	ApplicationsPaginationLinks = applications.PaginationLinks
)

// NewApplications binds the resource to a requester of your own.
func NewApplications(requester Requester) *Applications { return applications.New(requester) }

// SettingsApplicationListShape is the DTOShape reading a body as SettingsApplicationList, for the raw-verb call.
func SettingsApplicationListShape(decoded any) any {
	return applications.SettingsApplicationListShape(decoded)
}

// SettingsApplicationRecordShape is the DTOShape reading a body as SettingsApplicationRecord, for the raw-verb call.
func SettingsApplicationRecordShape(decoded any) any {
	return applications.SettingsApplicationRecordShape(decoded)
}

// The applications path templates and body keys, for the raw-verb call.
const (
	SettingsApplicationsPath = applications.SettingsApplicationsPath
	SettingsApplicationPath  = applications.SettingsApplicationPath

	SettingsApplicationsKey = applications.SettingsApplicationsKey
	ApplicationKey          = applications.ApplicationKey
)

// ApprovalProfiles is the approval profiles endpoints bound to a client.
type ApprovalProfiles = approvalprofiles.ApprovalProfiles

// The approval profiles views, aliased so a type assertion works under either spelling.
type (
	SettingsApprovalProfileList     = approvalprofiles.SettingsApprovalProfileList
	SettingsApprovalProfileRecord   = approvalprofiles.SettingsApprovalProfileRecord
	ApprovalProfilesPagination      = approvalprofiles.Pagination
	ApprovalProfilesPaginationLinks = approvalprofiles.PaginationLinks
)

// NewApprovalProfiles binds the resource to a requester of your own.
func NewApprovalProfiles(requester Requester) *ApprovalProfiles {
	return approvalprofiles.New(requester)
}

// SettingsApprovalProfileListShape is the DTOShape reading a body as SettingsApprovalProfileList, for the raw-verb call.
func SettingsApprovalProfileListShape(decoded any) any {
	return approvalprofiles.SettingsApprovalProfileListShape(decoded)
}

// SettingsApprovalProfileRecordShape is the DTOShape reading a body as SettingsApprovalProfileRecord, for the raw-verb call.
func SettingsApprovalProfileRecordShape(decoded any) any {
	return approvalprofiles.SettingsApprovalProfileRecordShape(decoded)
}

// The approval profiles path templates and body keys, for the raw-verb call.
const (
	SettingsApprovalProfilesPath = approvalprofiles.SettingsApprovalProfilesPath
	SettingsApprovalProfilePath  = approvalprofiles.SettingsApprovalProfilePath

	SettingsApprovalProfilesKey = approvalprofiles.SettingsApprovalProfilesKey
	ApprovalProfileKey          = approvalprofiles.ApprovalProfileKey
)

// More approvals path templates and body keys, for the raw-verb call.
const (
	ApprovalsApprovePath = approvals.ApprovalsApprovePath
	ApprovalsRejectPath  = approvals.ApprovalsRejectPath
)

// ApprovalUserGroups is the approval user groups endpoints bound to a client.
type ApprovalUserGroups = approvalusergroups.ApprovalUserGroups

// The approval user groups views, aliased so a type assertion works under either spelling.
type (
	SettingsApprovalUserGroupList     = approvalusergroups.SettingsApprovalUserGroupList
	SettingsApprovalUserGroupRecord   = approvalusergroups.SettingsApprovalUserGroupRecord
	ApprovalUserGroupsPagination      = approvalusergroups.Pagination
	ApprovalUserGroupsPaginationLinks = approvalusergroups.PaginationLinks
)

// NewApprovalUserGroups binds the resource to a requester of your own.
func NewApprovalUserGroups(requester Requester) *ApprovalUserGroups {
	return approvalusergroups.New(requester)
}

// SettingsApprovalUserGroupListShape is the DTOShape reading a body as SettingsApprovalUserGroupList, for the raw-verb call.
func SettingsApprovalUserGroupListShape(decoded any) any {
	return approvalusergroups.SettingsApprovalUserGroupListShape(decoded)
}

// SettingsApprovalUserGroupRecordShape is the DTOShape reading a body as SettingsApprovalUserGroupRecord, for the raw-verb call.
func SettingsApprovalUserGroupRecordShape(decoded any) any {
	return approvalusergroups.SettingsApprovalUserGroupRecordShape(decoded)
}

// The approval user groups path templates and body keys, for the raw-verb call.
const (
	SettingsApprovalUserGroupsPath = approvalusergroups.SettingsApprovalUserGroupsPath
	SettingsApprovalUserGroupPath  = approvalusergroups.SettingsApprovalUserGroupPath

	SettingsApprovalUserGroupsKey = approvalusergroups.SettingsApprovalUserGroupsKey
	ApprovalUserGroupKey          = approvalusergroups.ApprovalUserGroupKey
)

// AttributeAttachments is the attribute attachments endpoints bound to a client.
type AttributeAttachments = attributeattachments.AttributeAttachments

// NewAttributeAttachments binds the resource to a requester of your own.
func NewAttributeAttachments(requester Requester) *AttributeAttachments {
	return attributeattachments.New(requester)
}

// The attribute attachments path templates and body keys, for the raw-verb call.
const (
	AttributeAttachmentPath = attributeattachments.AttributeAttachmentPath
)

// BillingPeriods is the billing periods endpoints bound to a client.
type BillingPeriods = billingperiods.BillingPeriods

// The billing periods views, aliased so a type assertion works under either spelling.
type (
	SettingsBillingPeriodList     = billingperiods.SettingsBillingPeriodList
	SettingsBillingPeriodRecord   = billingperiods.SettingsBillingPeriodRecord
	BillingPeriodsPagination      = billingperiods.Pagination
	BillingPeriodsPaginationLinks = billingperiods.PaginationLinks
)

// NewBillingPeriods binds the resource to a requester of your own.
func NewBillingPeriods(requester Requester) *BillingPeriods { return billingperiods.New(requester) }

// SettingsBillingPeriodListShape is the DTOShape reading a body as SettingsBillingPeriodList, for the raw-verb call.
func SettingsBillingPeriodListShape(decoded any) any {
	return billingperiods.SettingsBillingPeriodListShape(decoded)
}

// SettingsBillingPeriodRecordShape is the DTOShape reading a body as SettingsBillingPeriodRecord, for the raw-verb call.
func SettingsBillingPeriodRecordShape(decoded any) any {
	return billingperiods.SettingsBillingPeriodRecordShape(decoded)
}

// The billing periods path templates and body keys, for the raw-verb call.
const (
	SettingsBillingPeriodsPath = billingperiods.SettingsBillingPeriodsPath
	SettingsBillingPeriodPath  = billingperiods.SettingsBillingPeriodPath

	SettingsBillingPeriodsKey = billingperiods.SettingsBillingPeriodsKey
	BillingPeriodKey          = billingperiods.BillingPeriodKey
)

// More bookings views, aliased so a type assertion works under either spelling.
type (
	BookingNoteList                                   = bookings.BookingNoteList
	BookingNote                                       = bookings.BookingNote
	BookingNoteAttachment                             = bookings.BookingNoteAttachment
	BookingNoteCreatedBy                              = bookings.BookingNoteCreatedBy
	BookingNoteLastUpdatedBy                          = bookings.BookingNoteLastUpdatedBy
	BookingNoteWatcher                                = bookings.BookingNoteWatcher
	BookingBillingPreferencesView                     = bookings.BookingBillingPreferencesView
	BookingBillingPreferencesDetail                   = bookings.BookingBillingPreferencesDetail
	BookingBillingPreferencesDetailProformaProperties = bookings.BookingBillingPreferencesDetailProformaProperties
	BookingLineRecord                                 = bookings.BookingLineRecord
	BookingLineItem                                   = bookings.BookingLineItem
	BookingLineItemAccountingCode                     = bookings.BookingLineItemAccountingCode
	BookingLineItemPriceSnapshot                      = bookings.BookingLineItemPriceSnapshot
	BookingLineItemPriceSnapshotPricingRule           = bookings.BookingLineItemPriceSnapshotPricingRule
	BookingLineItemPriceTax                           = bookings.BookingLineItemPriceTax
	BookingLineItemProperties                         = bookings.BookingLineItemProperties
	BookingLineItemSaleTaxConfiguration               = bookings.BookingLineItemSaleTaxConfiguration
	BookingLineItemSaleTaxConfigurationTaxCode        = bookings.BookingLineItemSaleTaxConfigurationTaxCode
	BookingLinePreOrderStockDetails                   = bookings.BookingLinePreOrderStockDetails
	BookingCustomObjectList                           = bookings.BookingCustomObjectList
	BookingCustomObjectEntry                          = bookings.BookingCustomObjectEntry
	BookingCustomObjectRecord                         = bookings.BookingCustomObjectRecord
	BookingInvoiceList                                = bookings.BookingInvoiceList
	BookingInvoice                                    = bookings.BookingInvoice
	BookingInvoiceAccount                             = bookings.BookingInvoiceAccount
	BookingInvoiceCreatedBy                           = bookings.BookingInvoiceCreatedBy
	BookingInvoiceCurrency                            = bookings.BookingInvoiceCurrency
	BookingInvoiceCustomForm                          = bookings.BookingInvoiceCustomForm
	BookingInvoiceOrder                               = bookings.BookingInvoiceOrder
	BookingItemFulfillmentRecord                      = bookings.BookingItemFulfillmentRecord
	BookingItemFulfillmentAccount                     = bookings.BookingItemFulfillmentAccount
	BookingItemFulfillmentCreatedBy                   = bookings.BookingItemFulfillmentCreatedBy
	BookingItemFulfillmentFulfillment                 = bookings.BookingItemFulfillmentFulfillment
	BookingItemFulfillmentFulfillmentItem             = bookings.BookingItemFulfillmentFulfillmentItem
	BookingItemFulfillmentInvoice                     = bookings.BookingItemFulfillmentInvoice
	BookingItemFulfillmentLastUpdatedBy               = bookings.BookingItemFulfillmentLastUpdatedBy
	BookingKpisRecord                                 = bookings.BookingKpisRecord
	BookingPaymentMethodsRecord                       = bookings.BookingPaymentMethodsRecord
	BookingPaymentMethodsPaymentMethod                = bookings.BookingPaymentMethodsPaymentMethod
	BookingPaymentMethodsPaymentMethodProcessor       = bookings.BookingPaymentMethodsPaymentMethodProcessor
	BookingPaymentMethodRecord                        = bookings.BookingPaymentMethodRecord
	BookingPaymentMethodPaymentMethod                 = bookings.BookingPaymentMethodPaymentMethod
	BookingPaymentMethodPaymentMethodProcessor        = bookings.BookingPaymentMethodPaymentMethodProcessor
	BookingPaymentList                                = bookings.BookingPaymentList
	BookingPayment                                    = bookings.BookingPayment
	BookingPaymentAccount                             = bookings.BookingPaymentAccount
	BookingPaymentCreatedBy                           = bookings.BookingPaymentCreatedBy
	BookingPaymentInvoice                             = bookings.BookingPaymentInvoice
	BookingPaymentPaymentApplied                      = bookings.BookingPaymentPaymentApplied
	BookingPaymentView                                = bookings.BookingPaymentView
	BookingPaymentViewAccount                         = bookings.BookingPaymentViewAccount
	BookingPaymentViewCreatedBy                       = bookings.BookingPaymentViewCreatedBy
	BookingPaymentViewInvoice                         = bookings.BookingPaymentViewInvoice
	BookingPaymentViewPaymentApplied                  = bookings.BookingPaymentViewPaymentApplied
	BookingPdfRecord                                  = bookings.BookingPdfRecord
	BookingPdfPdf                                     = bookings.BookingPdfPdf
	BookingQuotePartialInvoicesRecord                 = bookings.BookingQuotePartialInvoicesRecord
)

// BookingNoteListShape is the DTOShape reading a body as BookingNoteList, for the raw-verb call.
func BookingNoteListShape(decoded any) any { return bookings.BookingNoteListShape(decoded) }

// BookingBillingPreferencesViewShape is the DTOShape reading a body as BookingBillingPreferencesView, for the raw-verb call.
func BookingBillingPreferencesViewShape(decoded any) any {
	return bookings.BookingBillingPreferencesViewShape(decoded)
}

// BookingLineRecordShape is the DTOShape reading a body as BookingLineRecord, for the raw-verb call.
func BookingLineRecordShape(decoded any) any { return bookings.BookingLineRecordShape(decoded) }

// BookingCustomObjectListShape is the DTOShape reading a body as BookingCustomObjectList, for the raw-verb call.
func BookingCustomObjectListShape(decoded any) any {
	return bookings.BookingCustomObjectListShape(decoded)
}

// BookingCustomObjectRecordShape is the DTOShape reading a body as BookingCustomObjectRecord, for the raw-verb call.
func BookingCustomObjectRecordShape(decoded any) any {
	return bookings.BookingCustomObjectRecordShape(decoded)
}

// BookingInvoiceListShape is the DTOShape reading a body as BookingInvoiceList, for the raw-verb call.
func BookingInvoiceListShape(decoded any) any { return bookings.BookingInvoiceListShape(decoded) }

// BookingItemFulfillmentRecordShape is the DTOShape reading a body as BookingItemFulfillmentRecord, for the raw-verb call.
func BookingItemFulfillmentRecordShape(decoded any) any {
	return bookings.BookingItemFulfillmentRecordShape(decoded)
}

// BookingKpisRecordShape is the DTOShape reading a body as BookingKpisRecord, for the raw-verb call.
func BookingKpisRecordShape(decoded any) any { return bookings.BookingKpisRecordShape(decoded) }

// BookingPaymentMethodsRecordShape is the DTOShape reading a body as BookingPaymentMethodsRecord, for the raw-verb call.
func BookingPaymentMethodsRecordShape(decoded any) any {
	return bookings.BookingPaymentMethodsRecordShape(decoded)
}

// BookingPaymentMethodRecordShape is the DTOShape reading a body as BookingPaymentMethodRecord, for the raw-verb call.
func BookingPaymentMethodRecordShape(decoded any) any {
	return bookings.BookingPaymentMethodRecordShape(decoded)
}

// BookingPaymentListShape is the DTOShape reading a body as BookingPaymentList, for the raw-verb call.
func BookingPaymentListShape(decoded any) any { return bookings.BookingPaymentListShape(decoded) }

// BookingPaymentViewShape is the DTOShape reading a body as BookingPaymentView, for the raw-verb call.
func BookingPaymentViewShape(decoded any) any { return bookings.BookingPaymentViewShape(decoded) }

// BookingPdfRecordShape is the DTOShape reading a body as BookingPdfRecord, for the raw-verb call.
func BookingPdfRecordShape(decoded any) any { return bookings.BookingPdfRecordShape(decoded) }

// BookingQuotePartialInvoicesRecordShape is the DTOShape reading a body as BookingQuotePartialInvoicesRecord, for the raw-verb call.
func BookingQuotePartialInvoicesRecordShape(decoded any) any {
	return bookings.BookingQuotePartialInvoicesRecordShape(decoded)
}

// More bookings path templates and body keys, for the raw-verb call.
const (
	BookingPreviewPath                       = bookings.BookingPreviewPath
	BookingNotesPath                         = bookings.BookingNotesPath
	BookingBillingPreferencesPath            = bookings.BookingBillingPreferencesPath
	BookingChangePath                        = bookings.BookingChangePath
	BookingChangePreviewPath                 = bookings.BookingChangePreviewPath
	BookingInformationPath                   = bookings.BookingInformationPath
	BookingLinesPath                         = bookings.BookingLinesPath
	BookingLinePath                          = bookings.BookingLinePath
	BookingLineBillingPreferencesPath        = bookings.BookingLineBillingPreferencesPath
	BookingCustomObjectsPath                 = bookings.BookingCustomObjectsPath
	BookingCustomObjectPath                  = bookings.BookingCustomObjectPath
	BookingCustomObjectInstancePath          = bookings.BookingCustomObjectInstancePath
	BookingCustomObjectsInstancesPath        = bookings.BookingCustomObjectsInstancesPath
	BookingCustomObjectsInstancePath         = bookings.BookingCustomObjectsInstancePath
	BookingCustomObjectsInstanceInstancePath = bookings.BookingCustomObjectsInstanceInstancePath
	BookingInvoicesPath                      = bookings.BookingInvoicesPath
	BookingItemFulfillmentsPath              = bookings.BookingItemFulfillmentsPath
	BookingItemFulfillmentPath               = bookings.BookingItemFulfillmentPath
	BookingKPIsPath                          = bookings.BookingKPIsPath
	BookingPaymentMethodsPath                = bookings.BookingPaymentMethodsPath
	BookingPaymentMethodPath                 = bookings.BookingPaymentMethodPath
	BookingPaymentsPath                      = bookings.BookingPaymentsPath
	BookingPaymentPath                       = bookings.BookingPaymentPath
	BookingPDFPath                           = bookings.BookingPDFPath
	BookingQuotePartialInvoicesPath          = bookings.BookingQuotePartialInvoicesPath
	BookingShippingPDFPath                   = bookings.BookingShippingPDFPath

	BookingsOrderLineKey       = bookings.OrderLineKey
	BookingsCustomObjectKey    = bookings.CustomObjectKey
	BookingsItemFulfillmentKey = bookings.ItemFulfillmentKey
	BookingsOrderKPIsKey       = bookings.OrderKPIsKey
	BookingsPaymentKey         = bookings.PaymentKey
)

// ChargingPeriods is the charging periods endpoints bound to a client.
type ChargingPeriods = chargingperiods.ChargingPeriods

// The charging periods views, aliased so a type assertion works under either spelling.
type (
	SettingsChargingPeriodList       = chargingperiods.SettingsChargingPeriodList
	SettingsChargingPeriodRecord     = chargingperiods.SettingsChargingPeriodRecord
	SettingsChargingPeriodTimePeriod = chargingperiods.SettingsChargingPeriodTimePeriod
	ChargingPeriodsPagination        = chargingperiods.Pagination
	ChargingPeriodsPaginationLinks   = chargingperiods.PaginationLinks
)

// NewChargingPeriods binds the resource to a requester of your own.
func NewChargingPeriods(requester Requester) *ChargingPeriods { return chargingperiods.New(requester) }

// SettingsChargingPeriodListShape is the DTOShape reading a body as SettingsChargingPeriodList, for the raw-verb call.
func SettingsChargingPeriodListShape(decoded any) any {
	return chargingperiods.SettingsChargingPeriodListShape(decoded)
}

// SettingsChargingPeriodRecordShape is the DTOShape reading a body as SettingsChargingPeriodRecord, for the raw-verb call.
func SettingsChargingPeriodRecordShape(decoded any) any {
	return chargingperiods.SettingsChargingPeriodRecordShape(decoded)
}

// The charging periods path templates and body keys, for the raw-verb call.
const (
	SettingsChargingPeriodsPath = chargingperiods.SettingsChargingPeriodsPath
	SettingsChargingPeriodPath  = chargingperiods.SettingsChargingPeriodPath

	SettingsChargingPeriodsKey = chargingperiods.SettingsChargingPeriodsKey
	ChargingPeriodKey          = chargingperiods.ChargingPeriodKey
)

// CommunicationProfiles is the communication profiles endpoints bound to a client.
type CommunicationProfiles = communicationprofiles.CommunicationProfiles

// The communication profiles views, aliased so a type assertion works under either spelling.
type (
	SettingsCommunicationProfileList     = communicationprofiles.SettingsCommunicationProfileList
	SettingsCommunicationProfileRecord   = communicationprofiles.SettingsCommunicationProfileRecord
	CommunicationProfilesPagination      = communicationprofiles.Pagination
	CommunicationProfilesPaginationLinks = communicationprofiles.PaginationLinks
)

// NewCommunicationProfiles binds the resource to a requester of your own.
func NewCommunicationProfiles(requester Requester) *CommunicationProfiles {
	return communicationprofiles.New(requester)
}

// SettingsCommunicationProfileListShape is the DTOShape reading a body as SettingsCommunicationProfileList, for the raw-verb call.
func SettingsCommunicationProfileListShape(decoded any) any {
	return communicationprofiles.SettingsCommunicationProfileListShape(decoded)
}

// SettingsCommunicationProfileRecordShape is the DTOShape reading a body as SettingsCommunicationProfileRecord, for the raw-verb call.
func SettingsCommunicationProfileRecordShape(decoded any) any {
	return communicationprofiles.SettingsCommunicationProfileRecordShape(decoded)
}

// The communication profiles path templates and body keys, for the raw-verb call.
const (
	SettingsCommunicationProfilesPath = communicationprofiles.SettingsCommunicationProfilesPath
	SettingsCommunicationProfilePath  = communicationprofiles.SettingsCommunicationProfilePath

	SettingsCommunicationProfilesKey = communicationprofiles.SettingsCommunicationProfilesKey
	CommunicationProfileKey          = communicationprofiles.CommunicationProfileKey
)

// Communications is the communications endpoints bound to a client.
type Communications = communications.Communications

// NewCommunications binds the resource to a requester of your own.
func NewCommunications(requester Requester) *Communications { return communications.New(requester) }

// The communications path templates and body keys, for the raw-verb call.
const (
	CommunicationsSendEmailPath = communications.CommunicationsSendEmailPath

	CommunicationsEmailKey = communications.EmailKey
)

// Countries is the countries endpoints bound to a client.
type Countries = countries.Countries

// The countries views, aliased so a type assertion works under either spelling.
type (
	SettingsCountryList      = countries.SettingsCountryList
	SettingsCountryRecord    = countries.SettingsCountryRecord
	CountriesPagination      = countries.Pagination
	CountriesPaginationLinks = countries.PaginationLinks
)

// NewCountries binds the resource to a requester of your own.
func NewCountries(requester Requester) *Countries { return countries.New(requester) }

// SettingsCountryListShape is the DTOShape reading a body as SettingsCountryList, for the raw-verb call.
func SettingsCountryListShape(decoded any) any { return countries.SettingsCountryListShape(decoded) }

// SettingsCountryRecordShape is the DTOShape reading a body as SettingsCountryRecord, for the raw-verb call.
func SettingsCountryRecordShape(decoded any) any {
	return countries.SettingsCountryRecordShape(decoded)
}

// The countries path templates and body keys, for the raw-verb call.
const (
	SettingsCountriesPath = countries.SettingsCountriesPath
	SettingsCountryPath   = countries.SettingsCountryPath

	SettingsCountriesKey = countries.SettingsCountriesKey
	CountryKey           = countries.CountryKey
)

// More credit notes views, aliased so a type assertion works under either spelling.
type (
	CreditNoteItemsRecord        = creditnotes.CreditNoteItemsRecord
	CreditNoteItem               = creditnotes.CreditNoteItem
	CreditNoteApplicationsRecord = creditnotes.CreditNoteApplicationsRecord
	CreditNoteNoteList           = creditnotes.CreditNoteNoteList
	CreditNoteNoteRecord         = creditnotes.CreditNoteNoteRecord
	CreditNoteNoteCreatedBy      = creditnotes.CreditNoteNoteCreatedBy
	CreditNoteNoteDetail         = creditnotes.CreditNoteNoteDetail
)

// CreditNoteItemsRecordShape is the DTOShape reading a body as CreditNoteItemsRecord, for the raw-verb call.
func CreditNoteItemsRecordShape(decoded any) any {
	return creditnotes.CreditNoteItemsRecordShape(decoded)
}

// CreditNoteApplicationsRecordShape is the DTOShape reading a body as CreditNoteApplicationsRecord, for the raw-verb call.
func CreditNoteApplicationsRecordShape(decoded any) any {
	return creditnotes.CreditNoteApplicationsRecordShape(decoded)
}

// CreditNoteNoteListShape is the DTOShape reading a body as CreditNoteNoteList, for the raw-verb call.
func CreditNoteNoteListShape(decoded any) any { return creditnotes.CreditNoteNoteListShape(decoded) }

// CreditNoteNoteRecordShape is the DTOShape reading a body as CreditNoteNoteRecord, for the raw-verb call.
func CreditNoteNoteRecordShape(decoded any) any {
	return creditnotes.CreditNoteNoteRecordShape(decoded)
}

// More credit notes path templates and body keys, for the raw-verb call.
const (
	CreditNoteApplyPath                  = creditnotes.CreditNoteApplyPath
	CreditNoteItemsPath                  = creditnotes.CreditNoteItemsPath
	CreditNoteCreditNoteApplicationsPath = creditnotes.CreditNoteCreditNoteApplicationsPath
	CreditNoteNotesPath                  = creditnotes.CreditNoteNotesPath
	CreditNoteNotePath                   = creditnotes.CreditNoteNotePath
	CreditNoteNoteFilesPath              = creditnotes.CreditNoteNoteFilesPath
	CreditNoteNoteFilePath               = creditnotes.CreditNoteNoteFilePath

	CreditNotesNoteKey = creditnotes.NoteKey
)

// More custom components views, aliased so a type assertion works under either spelling.
type (
	ComponentObjectsDataDetail                     = customcomponents.ComponentObjectsDataDetail
	ComponentObjectsDataCustomForm                 = customcomponents.ComponentObjectsDataCustomForm
	ComponentObjectsDataParent                     = customcomponents.ComponentObjectsDataParent
	ComponentObjectRecordView                      = customcomponents.ComponentObjectRecordView
	ComponentObjectCustomComponentObjectDetail     = customcomponents.ComponentObjectCustomComponentObjectDetail
	ComponentObjectCustomComponentObjectCustomForm = customcomponents.ComponentObjectCustomComponentObjectCustomForm
	ComponentObjectCustomComponentObjectParent     = customcomponents.ComponentObjectCustomComponentObjectParent
	ComponentObjectRecordActivityLog               = customcomponents.ComponentObjectRecordActivityLog
	ComponentObjectActivityEntryRecord             = customcomponents.ComponentObjectActivityEntryRecord
	ComponentObjectNotesRecord                     = customcomponents.ComponentObjectNotesRecord
	ComponentObjectNoteRecord                      = customcomponents.ComponentObjectNoteRecord
	ComponentObjectNoteAttachment                  = customcomponents.ComponentObjectNoteAttachment
	ComponentObjectNoteCreatedBy                   = customcomponents.ComponentObjectNoteCreatedBy
	ComponentObjectNoteCustomComponentObject       = customcomponents.ComponentObjectNoteCustomComponentObject
	ComponentObjectNoteLastUpdatedBy               = customcomponents.ComponentObjectNoteLastUpdatedBy
	ComponentObjectNoteFilesRecord                 = customcomponents.ComponentObjectNoteFilesRecord
	ComponentObjectNoteFileRecord                  = customcomponents.ComponentObjectNoteFileRecord
)

// ComponentObjectRecordViewShape is the DTOShape reading a body as ComponentObjectRecordView, for the raw-verb call.
func ComponentObjectRecordViewShape(decoded any) any {
	return customcomponents.ComponentObjectRecordViewShape(decoded)
}

// ComponentObjectActivityEntryRecordShape is the DTOShape reading a body as ComponentObjectActivityEntryRecord, for the raw-verb call.
func ComponentObjectActivityEntryRecordShape(decoded any) any {
	return customcomponents.ComponentObjectActivityEntryRecordShape(decoded)
}

// ComponentObjectNotesRecordShape is the DTOShape reading a body as ComponentObjectNotesRecord, for the raw-verb call.
func ComponentObjectNotesRecordShape(decoded any) any {
	return customcomponents.ComponentObjectNotesRecordShape(decoded)
}

// ComponentObjectNoteRecordShape is the DTOShape reading a body as ComponentObjectNoteRecord, for the raw-verb call.
func ComponentObjectNoteRecordShape(decoded any) any {
	return customcomponents.ComponentObjectNoteRecordShape(decoded)
}

// ComponentObjectNoteFilesRecordShape is the DTOShape reading a body as ComponentObjectNoteFilesRecord, for the raw-verb call.
func ComponentObjectNoteFilesRecordShape(decoded any) any {
	return customcomponents.ComponentObjectNoteFilesRecordShape(decoded)
}

// ComponentObjectNoteFileRecordShape is the DTOShape reading a body as ComponentObjectNoteFileRecord, for the raw-verb call.
func ComponentObjectNoteFileRecordShape(decoded any) any {
	return customcomponents.ComponentObjectNoteFileRecordShape(decoded)
}

// More custom components path templates and body keys, for the raw-verb call.
const (
	ComponentObjectActivityEntryPath        = customcomponents.ComponentObjectActivityEntryPath
	ComponentObjectMilestonesPath           = customcomponents.ComponentObjectMilestonesPath
	ComponentObjectMilestonePath            = customcomponents.ComponentObjectMilestonePath
	ComponentObjectHistoryPath              = customcomponents.ComponentObjectHistoryPath
	ComponentObjectHistoryEntryPath         = customcomponents.ComponentObjectHistoryEntryPath
	ComponentObjectNotesPath                = customcomponents.ComponentObjectNotesPath
	ComponentObjectNotePath                 = customcomponents.ComponentObjectNotePath
	ComponentObjectNoteFilesPath            = customcomponents.ComponentObjectNoteFilesPath
	ComponentObjectNoteFilePath             = customcomponents.ComponentObjectNoteFilePath
	ComponentObjectCustomObjectPath         = customcomponents.ComponentObjectCustomObjectPath
	ComponentObjectCustomObjectInstancePath = customcomponents.ComponentObjectCustomObjectInstancePath
	ComponentObjectLinkDomainsPath          = customcomponents.ComponentObjectLinkDomainsPath
	ComponentObjectUnlinkDomainsPath        = customcomponents.ComponentObjectUnlinkDomainsPath
	ComponentObjectChangeStatusPath         = customcomponents.ComponentObjectChangeStatusPath

	CustomComponentObjectActivityKey = customcomponents.CustomComponentObjectActivityKey
	CustomComponentsNoteKey          = customcomponents.NoteKey
	CustomComponentsFileKey          = customcomponents.FileKey
)

// CustomEvents is the custom events endpoints bound to a client.
type CustomEvents = customevents.CustomEvents

// The custom events views, aliased so a type assertion works under either spelling.
type (
	SettingsCustomEventList     = customevents.SettingsCustomEventList
	SettingsCustomEventRecord   = customevents.SettingsCustomEventRecord
	CustomEventsPagination      = customevents.Pagination
	CustomEventsPaginationLinks = customevents.PaginationLinks
)

// NewCustomEvents binds the resource to a requester of your own.
func NewCustomEvents(requester Requester) *CustomEvents { return customevents.New(requester) }

// SettingsCustomEventListShape is the DTOShape reading a body as SettingsCustomEventList, for the raw-verb call.
func SettingsCustomEventListShape(decoded any) any {
	return customevents.SettingsCustomEventListShape(decoded)
}

// SettingsCustomEventRecordShape is the DTOShape reading a body as SettingsCustomEventRecord, for the raw-verb call.
func SettingsCustomEventRecordShape(decoded any) any {
	return customevents.SettingsCustomEventRecordShape(decoded)
}

// The custom events path templates and body keys, for the raw-verb call.
const (
	SettingsCustomEventsPath = customevents.SettingsCustomEventsPath
	SettingsCustomEventPath  = customevents.SettingsCustomEventPath

	SettingsCustomEventsKey = customevents.SettingsCustomEventsKey
	CustomEventKey          = customevents.CustomEventKey
)

// More custom features views, aliased so a type assertion works under either spelling.
type (
	CustomFeatureList                                          = customfeatures.CustomFeatureList
	CustomFeatureDetail                                        = customfeatures.CustomFeatureDetail
	CustomFeatureDetailCustomForm                              = customfeatures.CustomFeatureDetailCustomForm
	CustomFeatureDetailCustomFormConfiguration                 = customfeatures.CustomFeatureDetailCustomFormConfiguration
	CustomFeatureDetailCustomFormConfigurationColumn1          = customfeatures.CustomFeatureDetailCustomFormConfigurationColumn1
	CustomFeatureDetailCustomFormConfigurationColumn1Attribute = customfeatures.CustomFeatureDetailCustomFormConfigurationColumn1Attribute
	CustomFeatureDetailCustomFormConfigurationColumn2          = customfeatures.CustomFeatureDetailCustomFormConfigurationColumn2
	CustomFeatureDetailCustomFormConfigurationColumn2Attribute = customfeatures.CustomFeatureDetailCustomFormConfigurationColumn2Attribute
	CustomFeatureDetailCustomFormConfigurationSettings         = customfeatures.CustomFeatureDetailCustomFormConfigurationSettings
	CustomFeaturesPagination                                   = customfeatures.Pagination
	CustomFeaturesPaginationLinks                              = customfeatures.PaginationLinks
)

// CustomFeatureListShape is the DTOShape reading a body as CustomFeatureList, for the raw-verb call.
func CustomFeatureListShape(decoded any) any { return customfeatures.CustomFeatureListShape(decoded) }

// CustomFields is the custom fields endpoints bound to a client.
type CustomFields = customfields.CustomFields

// The custom fields views, aliased so a type assertion works under either spelling.
type (
	SettingsCustomFieldList     = customfields.SettingsCustomFieldList
	SettingsCustomFieldRecord   = customfields.SettingsCustomFieldRecord
	CustomFieldsPagination      = customfields.Pagination
	CustomFieldsPaginationLinks = customfields.PaginationLinks
)

// NewCustomFields binds the resource to a requester of your own.
func NewCustomFields(requester Requester) *CustomFields { return customfields.New(requester) }

// SettingsCustomFieldListShape is the DTOShape reading a body as SettingsCustomFieldList, for the raw-verb call.
func SettingsCustomFieldListShape(decoded any) any {
	return customfields.SettingsCustomFieldListShape(decoded)
}

// SettingsCustomFieldRecordShape is the DTOShape reading a body as SettingsCustomFieldRecord, for the raw-verb call.
func SettingsCustomFieldRecordShape(decoded any) any {
	return customfields.SettingsCustomFieldRecordShape(decoded)
}

// The custom fields path templates and body keys, for the raw-verb call.
const (
	SettingsCustomFieldsPath = customfields.SettingsCustomFieldsPath
	SettingsCustomFieldPath  = customfields.SettingsCustomFieldPath

	SettingsCustomFieldsKey = customfields.SettingsCustomFieldsKey
	CustomFieldKey          = customfields.CustomFieldKey
)

// CustomForms is the custom forms endpoints bound to a client.
type CustomForms = customforms.CustomForms

// The custom forms views, aliased so a type assertion works under either spelling.
type (
	SettingsCustomFormList     = customforms.SettingsCustomFormList
	SettingsCustomFormRecord   = customforms.SettingsCustomFormRecord
	CustomFormsPagination      = customforms.Pagination
	CustomFormsPaginationLinks = customforms.PaginationLinks
)

// NewCustomForms binds the resource to a requester of your own.
func NewCustomForms(requester Requester) *CustomForms { return customforms.New(requester) }

// SettingsCustomFormListShape is the DTOShape reading a body as SettingsCustomFormList, for the raw-verb call.
func SettingsCustomFormListShape(decoded any) any {
	return customforms.SettingsCustomFormListShape(decoded)
}

// SettingsCustomFormRecordShape is the DTOShape reading a body as SettingsCustomFormRecord, for the raw-verb call.
func SettingsCustomFormRecordShape(decoded any) any {
	return customforms.SettingsCustomFormRecordShape(decoded)
}

// The custom forms path templates and body keys, for the raw-verb call.
const (
	SettingsCustomFormsPath = customforms.SettingsCustomFormsPath
	SettingsCustomFormPath  = customforms.SettingsCustomFormPath

	SettingsCustomFormsKey = customforms.SettingsCustomFormsKey
	CustomFormKey          = customforms.CustomFormKey
)

// DayOfMonth is the day of month endpoints bound to a client.
type DayOfMonth = dayofmonth.DayOfMonth

// The day of month views, aliased so a type assertion works under either spelling.
type (
	SettingsDayOfMonthList    = dayofmonth.SettingsDayOfMonthList
	SettingsDayOfMonthRecord  = dayofmonth.SettingsDayOfMonthRecord
	DayOfMonthPagination      = dayofmonth.Pagination
	DayOfMonthPaginationLinks = dayofmonth.PaginationLinks
)

// NewDayOfMonth binds the resource to a requester of your own.
func NewDayOfMonth(requester Requester) *DayOfMonth { return dayofmonth.New(requester) }

// SettingsDayOfMonthListShape is the DTOShape reading a body as SettingsDayOfMonthList, for the raw-verb call.
func SettingsDayOfMonthListShape(decoded any) any {
	return dayofmonth.SettingsDayOfMonthListShape(decoded)
}

// SettingsDayOfMonthRecordShape is the DTOShape reading a body as SettingsDayOfMonthRecord, for the raw-verb call.
func SettingsDayOfMonthRecordShape(decoded any) any {
	return dayofmonth.SettingsDayOfMonthRecordShape(decoded)
}

// The day of month path templates and body keys, for the raw-verb call.
const (
	SettingsDaysOfMonthPath = dayofmonth.SettingsDaysOfMonthPath
	SettingsDayOfMonthPath  = dayofmonth.SettingsDayOfMonthPath

	SettingsDaysOfMonthKey = dayofmonth.SettingsDaysOfMonthKey
	DayOfMonthKey          = dayofmonth.DayOfMonthKey
)

// Departments is the departments endpoints bound to a client.
type Departments = departments.Departments

// The departments views, aliased so a type assertion works under either spelling.
type (
	SettingsDepartmentList     = departments.SettingsDepartmentList
	SettingsDepartmentRecord   = departments.SettingsDepartmentRecord
	SettingsDepartmentDivision = departments.SettingsDepartmentDivision
	DepartmentsPagination      = departments.Pagination
	DepartmentsPaginationLinks = departments.PaginationLinks
)

// NewDepartments binds the resource to a requester of your own.
func NewDepartments(requester Requester) *Departments { return departments.New(requester) }

// SettingsDepartmentListShape is the DTOShape reading a body as SettingsDepartmentList, for the raw-verb call.
func SettingsDepartmentListShape(decoded any) any {
	return departments.SettingsDepartmentListShape(decoded)
}

// SettingsDepartmentRecordShape is the DTOShape reading a body as SettingsDepartmentRecord, for the raw-verb call.
func SettingsDepartmentRecordShape(decoded any) any {
	return departments.SettingsDepartmentRecordShape(decoded)
}

// The departments path templates and body keys, for the raw-verb call.
const (
	SettingsDepartmentsPath = departments.SettingsDepartmentsPath
	SettingsDepartmentPath  = departments.SettingsDepartmentPath

	SettingsDepartmentsKey = departments.SettingsDepartmentsKey
	DepartmentKey          = departments.DepartmentKey
)

// Designations is the designations endpoints bound to a client.
type Designations = designations.Designations

// The designations views, aliased so a type assertion works under either spelling.
type (
	SettingsDesignationList     = designations.SettingsDesignationList
	SettingsDesignationRecord   = designations.SettingsDesignationRecord
	DesignationsPagination      = designations.Pagination
	DesignationsPaginationLinks = designations.PaginationLinks
)

// NewDesignations binds the resource to a requester of your own.
func NewDesignations(requester Requester) *Designations { return designations.New(requester) }

// SettingsDesignationListShape is the DTOShape reading a body as SettingsDesignationList, for the raw-verb call.
func SettingsDesignationListShape(decoded any) any {
	return designations.SettingsDesignationListShape(decoded)
}

// SettingsDesignationRecordShape is the DTOShape reading a body as SettingsDesignationRecord, for the raw-verb call.
func SettingsDesignationRecordShape(decoded any) any {
	return designations.SettingsDesignationRecordShape(decoded)
}

// The designations path templates and body keys, for the raw-verb call.
const (
	SettingsDesignationsPath = designations.SettingsDesignationsPath
	SettingsDesignationPath  = designations.SettingsDesignationPath

	SettingsDesignationsKey = designations.SettingsDesignationsKey
	DesignationKey          = designations.DesignationKey
)

// DocumentTemplates is the document templates endpoints bound to a client.
type DocumentTemplates = documenttemplates.DocumentTemplates

// The document templates views, aliased so a type assertion works under either spelling.
type (
	SettingsDocumentTemplateList     = documenttemplates.SettingsDocumentTemplateList
	SettingsDocumentTemplateRecord   = documenttemplates.SettingsDocumentTemplateRecord
	DocumentTemplatesPagination      = documenttemplates.Pagination
	DocumentTemplatesPaginationLinks = documenttemplates.PaginationLinks
)

// NewDocumentTemplates binds the resource to a requester of your own.
func NewDocumentTemplates(requester Requester) *DocumentTemplates {
	return documenttemplates.New(requester)
}

// SettingsDocumentTemplateListShape is the DTOShape reading a body as SettingsDocumentTemplateList, for the raw-verb call.
func SettingsDocumentTemplateListShape(decoded any) any {
	return documenttemplates.SettingsDocumentTemplateListShape(decoded)
}

// SettingsDocumentTemplateRecordShape is the DTOShape reading a body as SettingsDocumentTemplateRecord, for the raw-verb call.
func SettingsDocumentTemplateRecordShape(decoded any) any {
	return documenttemplates.SettingsDocumentTemplateRecordShape(decoded)
}

// The document templates path templates and body keys, for the raw-verb call.
const (
	SettingsDocumentTemplatesPath = documenttemplates.SettingsDocumentTemplatesPath
	SettingsDocumentTemplatePath  = documenttemplates.SettingsDocumentTemplatePath

	SettingsDocumentTemplatesKey = documenttemplates.SettingsDocumentTemplatesKey
	DocumentTemplateKey          = documenttemplates.DocumentTemplateKey
)

// EmailTemplates is the email templates endpoints bound to a client.
type EmailTemplates = emailtemplates.EmailTemplates

// The email templates views, aliased so a type assertion works under either spelling.
type (
	SettingsEmailTemplateList     = emailtemplates.SettingsEmailTemplateList
	SettingsEmailTemplateRecord   = emailtemplates.SettingsEmailTemplateRecord
	EmailTemplatesPagination      = emailtemplates.Pagination
	EmailTemplatesPaginationLinks = emailtemplates.PaginationLinks
)

// NewEmailTemplates binds the resource to a requester of your own.
func NewEmailTemplates(requester Requester) *EmailTemplates { return emailtemplates.New(requester) }

// SettingsEmailTemplateListShape is the DTOShape reading a body as SettingsEmailTemplateList, for the raw-verb call.
func SettingsEmailTemplateListShape(decoded any) any {
	return emailtemplates.SettingsEmailTemplateListShape(decoded)
}

// SettingsEmailTemplateRecordShape is the DTOShape reading a body as SettingsEmailTemplateRecord, for the raw-verb call.
func SettingsEmailTemplateRecordShape(decoded any) any {
	return emailtemplates.SettingsEmailTemplateRecordShape(decoded)
}

// The email templates path templates and body keys, for the raw-verb call.
const (
	SettingsEmailTemplatesPath = emailtemplates.SettingsEmailTemplatesPath
	SettingsEmailTemplatePath  = emailtemplates.SettingsEmailTemplatePath

	SettingsEmailTemplatesKey = emailtemplates.SettingsEmailTemplatesKey
	EmailTemplateKey          = emailtemplates.EmailTemplateKey
)

// FulfillmentProfiles is the fulfillment profiles endpoints bound to a client.
type FulfillmentProfiles = fulfillmentprofiles.FulfillmentProfiles

// The fulfillment profiles views, aliased so a type assertion works under either spelling.
type (
	SettingsFulfillmentProfileList     = fulfillmentprofiles.SettingsFulfillmentProfileList
	SettingsFulfillmentProfileRecord   = fulfillmentprofiles.SettingsFulfillmentProfileRecord
	FulfillmentProfilesPagination      = fulfillmentprofiles.Pagination
	FulfillmentProfilesPaginationLinks = fulfillmentprofiles.PaginationLinks
)

// NewFulfillmentProfiles binds the resource to a requester of your own.
func NewFulfillmentProfiles(requester Requester) *FulfillmentProfiles {
	return fulfillmentprofiles.New(requester)
}

// SettingsFulfillmentProfileListShape is the DTOShape reading a body as SettingsFulfillmentProfileList, for the raw-verb call.
func SettingsFulfillmentProfileListShape(decoded any) any {
	return fulfillmentprofiles.SettingsFulfillmentProfileListShape(decoded)
}

// SettingsFulfillmentProfileRecordShape is the DTOShape reading a body as SettingsFulfillmentProfileRecord, for the raw-verb call.
func SettingsFulfillmentProfileRecordShape(decoded any) any {
	return fulfillmentprofiles.SettingsFulfillmentProfileRecordShape(decoded)
}

// The fulfillment profiles path templates and body keys, for the raw-verb call.
const (
	SettingsFulfillmentProfilesPath = fulfillmentprofiles.SettingsFulfillmentProfilesPath
	SettingsFulfillmentProfilePath  = fulfillmentprofiles.SettingsFulfillmentProfilePath

	SettingsFulfillmentProfilesKey = fulfillmentprofiles.SettingsFulfillmentProfilesKey
	FulfillmentProfileKey          = fulfillmentprofiles.FulfillmentProfileKey
)

// FulfillmentTerms is the fulfillment terms endpoints bound to a client.
type FulfillmentTerms = fulfillmentterms.FulfillmentTerms

// The fulfillment terms views, aliased so a type assertion works under either spelling.
type (
	SettingsFulfillmentTermList     = fulfillmentterms.SettingsFulfillmentTermList
	SettingsFulfillmentTermRecord   = fulfillmentterms.SettingsFulfillmentTermRecord
	FulfillmentTermsPagination      = fulfillmentterms.Pagination
	FulfillmentTermsPaginationLinks = fulfillmentterms.PaginationLinks
)

// NewFulfillmentTerms binds the resource to a requester of your own.
func NewFulfillmentTerms(requester Requester) *FulfillmentTerms {
	return fulfillmentterms.New(requester)
}

// SettingsFulfillmentTermListShape is the DTOShape reading a body as SettingsFulfillmentTermList, for the raw-verb call.
func SettingsFulfillmentTermListShape(decoded any) any {
	return fulfillmentterms.SettingsFulfillmentTermListShape(decoded)
}

// SettingsFulfillmentTermRecordShape is the DTOShape reading a body as SettingsFulfillmentTermRecord, for the raw-verb call.
func SettingsFulfillmentTermRecordShape(decoded any) any {
	return fulfillmentterms.SettingsFulfillmentTermRecordShape(decoded)
}

// The fulfillment terms path templates and body keys, for the raw-verb call.
const (
	SettingsFulfillmentTermsPath = fulfillmentterms.SettingsFulfillmentTermsPath
	SettingsFulfillmentTermPath  = fulfillmentterms.SettingsFulfillmentTermPath

	SettingsFulfillmentTermsKey = fulfillmentterms.SettingsFulfillmentTermsKey
	FulfillmentTermKey          = fulfillmentterms.FulfillmentTermKey
)

// More gift certificates views, aliased so a type assertion works under either spelling.
type (
	GiftCertificateAllocationReadRecord      = giftcertificates.GiftCertificateAllocationReadRecord
	GiftCertificateAllocationAccount         = giftcertificates.GiftCertificateAllocationAccount
	GiftCertificateTransactionReadRecord     = giftcertificates.GiftCertificateTransactionReadRecord
	GiftCertificateTransactionAccountingCode = giftcertificates.GiftCertificateTransactionAccountingCode
	GiftCertificateTransactionCurrency       = giftcertificates.GiftCertificateTransactionCurrency
)

// GiftCertificateAllocationReadRecordShape is the DTOShape reading a body as GiftCertificateAllocationReadRecord, for the raw-verb call.
func GiftCertificateAllocationReadRecordShape(decoded any) any {
	return giftcertificates.GiftCertificateAllocationReadRecordShape(decoded)
}

// GiftCertificateTransactionReadRecordShape is the DTOShape reading a body as GiftCertificateTransactionReadRecord, for the raw-verb call.
func GiftCertificateTransactionReadRecordShape(decoded any) any {
	return giftcertificates.GiftCertificateTransactionReadRecordShape(decoded)
}

// More gift certificates path templates and body keys, for the raw-verb call.
const (
	GiftCertificateAmendPath         = giftcertificates.GiftCertificateAmendPath
	GiftCertificateDeallocationsPath = giftcertificates.GiftCertificateDeallocationsPath
	GiftCertificateAllocationPath    = giftcertificates.GiftCertificateAllocationPath
	GiftCertificateTransactionPath   = giftcertificates.GiftCertificateTransactionPath

	GiftCertificateAllocationKey  = giftcertificates.GiftCertificateAllocationKey
	GiftCertificateTransactionKey = giftcertificates.GiftCertificateTransactionKey
)

// More invoices views, aliased so a type assertion works under either spelling.
type (
	InvoiceKpisReadRecord = invoices.InvoiceKpisReadRecord
	InvoiceLinesRecord    = invoices.InvoiceLinesRecord
	InvoiceLine           = invoices.InvoiceLine
	InvoiceLineItem       = invoices.InvoiceLineItem
	LineTax               = invoices.LineTax
	InvoiceLineReadRecord = invoices.InvoiceLineReadRecord
)

// InvoiceKpisReadRecordShape is the DTOShape reading a body as InvoiceKpisReadRecord, for the raw-verb call.
func InvoiceKpisReadRecordShape(decoded any) any { return invoices.InvoiceKpisReadRecordShape(decoded) }

// InvoiceLinesRecordShape is the DTOShape reading a body as InvoiceLinesRecord, for the raw-verb call.
func InvoiceLinesRecordShape(decoded any) any { return invoices.InvoiceLinesRecordShape(decoded) }

// InvoiceLineReadRecordShape is the DTOShape reading a body as InvoiceLineReadRecord, for the raw-verb call.
func InvoiceLineReadRecordShape(decoded any) any { return invoices.InvoiceLineReadRecordShape(decoded) }

// More invoices path templates and body keys, for the raw-verb call.
const (
	InvoiceKPIsPath                  = invoices.InvoiceKPIsPath
	InvoiceLinesPath                 = invoices.InvoiceLinesPath
	InvoiceLinePath                  = invoices.InvoiceLinePath
	InvoiceCreditNoteApplicationPath = invoices.InvoiceCreditNoteApplicationPath
	InvoiceCustomObjectsPath         = invoices.InvoiceCustomObjectsPath
	InvoiceCustomObjectPath          = invoices.InvoiceCustomObjectPath
	InvoiceCustomObjectInstancePath  = invoices.InvoiceCustomObjectInstancePath

	InvoiceKPIsKey = invoices.InvoiceKPIsKey
)

// InvoiceTerms is the invoice terms endpoints bound to a client.
type InvoiceTerms = invoiceterms.InvoiceTerms

// The invoice terms views, aliased so a type assertion works under either spelling.
type (
	SettingsInvoiceTermList        = invoiceterms.SettingsInvoiceTermList
	SettingsInvoiceTermRecord      = invoiceterms.SettingsInvoiceTermRecord
	SettingsInvoiceTermDuration    = invoiceterms.SettingsInvoiceTermDuration
	InvoiceTermsPagination         = invoiceterms.Pagination
	InvoiceTermsPaginationLinks    = invoiceterms.PaginationLinks
	SettingsInvoiceTermWriteRecord = invoiceterms.SettingsInvoiceTermWriteRecord
)

// NewInvoiceTerms binds the resource to a requester of your own.
func NewInvoiceTerms(requester Requester) *InvoiceTerms { return invoiceterms.New(requester) }

// SettingsInvoiceTermListShape is the DTOShape reading a body as SettingsInvoiceTermList, for the raw-verb call.
func SettingsInvoiceTermListShape(decoded any) any {
	return invoiceterms.SettingsInvoiceTermListShape(decoded)
}

// SettingsInvoiceTermRecordShape is the DTOShape reading a body as SettingsInvoiceTermRecord, for the raw-verb call.
func SettingsInvoiceTermRecordShape(decoded any) any {
	return invoiceterms.SettingsInvoiceTermRecordShape(decoded)
}

// SettingsInvoiceTermWriteRecordShape is the DTOShape reading a body as SettingsInvoiceTermWriteRecord, for the raw-verb call.
func SettingsInvoiceTermWriteRecordShape(decoded any) any {
	return invoiceterms.SettingsInvoiceTermWriteRecordShape(decoded)
}

// The invoice terms path templates and body keys, for the raw-verb call.
const (
	SettingsInvoiceTermsPath = invoiceterms.SettingsInvoiceTermsPath
	SettingsInvoiceTermPath  = invoiceterms.SettingsInvoiceTermPath

	SettingsInvoiceTermsKey = invoiceterms.SettingsInvoiceTermsKey
	InvoiceTermKey          = invoiceterms.InvoiceTermKey
)

// More item fulfillments views, aliased so a type assertion works under either spelling.
type (
	ItemFulfillmentNoteList       = itemfulfillments.ItemFulfillmentNoteList
	ItemFulfillmentNote           = itemfulfillments.ItemFulfillmentNote
	ItemFulfillmentNoteAttachment = itemfulfillments.ItemFulfillmentNoteAttachment
	ItemFulfillmentCreatedBy      = itemfulfillments.ItemFulfillmentCreatedBy
	ItemFulfillmentNotesDetail    = itemfulfillments.ItemFulfillmentNotesDetail
	ItemFulfillmentNoteWatcher    = itemfulfillments.ItemFulfillmentNoteWatcher
	ItemFulfillmentNoteReadRecord = itemfulfillments.ItemFulfillmentNoteReadRecord
)

// ItemFulfillmentNoteListShape is the DTOShape reading a body as ItemFulfillmentNoteList, for the raw-verb call.
func ItemFulfillmentNoteListShape(decoded any) any {
	return itemfulfillments.ItemFulfillmentNoteListShape(decoded)
}

// ItemFulfillmentNoteReadRecordShape is the DTOShape reading a body as ItemFulfillmentNoteReadRecord, for the raw-verb call.
func ItemFulfillmentNoteReadRecordShape(decoded any) any {
	return itemfulfillments.ItemFulfillmentNoteReadRecordShape(decoded)
}

// More item fulfillments path templates and body keys, for the raw-verb call.
const (
	ItemFulfillmentNotePath = itemfulfillments.ItemFulfillmentNotePath
)

// More item receipts path templates and body keys, for the raw-verb call.
const (
	ItemReceiptNotePath = itemreceipts.ItemReceiptNotePath
)

// More items path templates and body keys, for the raw-verb call.
const (
	ItemInventoriesAddPath = items.ItemInventoriesAddPath
	ItemInventoryPath      = items.ItemInventoryPath
	ItemKPIsPath           = items.ItemKPIsPath
)

// More labours views, aliased so a type assertion works under either spelling.
type (
	LabourProfileSessionList                  = labours.LabourProfileSessionList
	LabourProfileViewSession                  = labours.LabourProfileViewSession
	LabourProfileViewSessionAvailability      = labours.LabourProfileViewSessionAvailability
	LabourProfileViewSessionPrice             = labours.LabourProfileViewSessionPrice
	LabourProfileViewSessionPriceHour         = labours.LabourProfileViewSessionPriceHour
	LabourProfileViewSessionPricePricingLevel = labours.LabourProfileViewSessionPricePricingLevel
	LabourProfileViewSessionWorkingHour       = labours.LabourProfileViewSessionWorkingHour
	LabourProfileSessionView                  = labours.LabourProfileSessionView
	LabourCalendarRecord                      = labours.LabourCalendarRecord
	LabourAvailabilityLabourProfile           = labours.LabourAvailabilityLabourProfile
	LabourAvailabilityScheduleList            = labours.LabourAvailabilityScheduleList
	LabourAvailabilitySessionList             = labours.LabourAvailabilitySessionList
	LabourAvailabilitySessionListScheduleList = labours.LabourAvailabilitySessionListScheduleList
)

// LabourProfileSessionListShape is the DTOShape reading a body as LabourProfileSessionList, for the raw-verb call.
func LabourProfileSessionListShape(decoded any) any {
	return labours.LabourProfileSessionListShape(decoded)
}

// LabourProfileSessionViewShape is the DTOShape reading a body as LabourProfileSessionView, for the raw-verb call.
func LabourProfileSessionViewShape(decoded any) any {
	return labours.LabourProfileSessionViewShape(decoded)
}

// LabourCalendarRecordShape is the DTOShape reading a body as LabourCalendarRecord, for the raw-verb call.
func LabourCalendarRecordShape(decoded any) any { return labours.LabourCalendarRecordShape(decoded) }

// More labours path templates and body keys, for the raw-verb call.
const (
	LabourProfileSessionsPath = labours.LabourProfileSessionsPath
	LabourProfileSessionPath  = labours.LabourProfileSessionPath
	LabourCalendarPath        = labours.LabourCalendarPath

	LabourProfileSessionKey = labours.LabourProfileSessionKey
	LabourCalendarKey       = labours.LabourCalendarKey
)

// More milestones path templates and body keys, for the raw-verb call.
const (
	MilestonesEmployeeByCodePath = milestones.MilestonesEmployeeByCodePath
	MilestonesMilestonePath      = milestones.MilestonePath
)

// NoteAttachments is the note attachments endpoints bound to a client.
type NoteAttachments = noteattachments.NoteAttachments

// The note attachments views, aliased so a type assertion works under either spelling.
type (
	NoteAttachmentRecord = noteattachments.NoteAttachmentRecord
)

// NewNoteAttachments binds the resource to a requester of your own.
func NewNoteAttachments(requester Requester) *NoteAttachments { return noteattachments.New(requester) }

// NoteAttachmentRecordShape is the DTOShape reading a body as NoteAttachmentRecord, for the raw-verb call.
func NoteAttachmentRecordShape(decoded any) any {
	return noteattachments.NoteAttachmentRecordShape(decoded)
}

// The note attachments path templates and body keys, for the raw-verb call.
const (
	NoteAttachmentPath = noteattachments.NoteAttachmentPath

	NoteAttachmentKey = noteattachments.NoteAttachmentKey
)

// More note groups path templates and body keys, for the raw-verb call.
const (
	NoteGroupNotesByUserPath = notegroups.NoteGroupNotesByUserPath
)

// More orders views, aliased so a type assertion works under either spelling.
type (
	OrderCustomObjectRecord             = orders.OrderCustomObjectRecord
	OrderItemFulfillmentRecord          = orders.OrderItemFulfillmentRecord
	OrderItemFulfillmentAccount         = orders.OrderItemFulfillmentAccount
	OrderItemFulfillmentCreatedBy       = orders.OrderItemFulfillmentCreatedBy
	OrderItemFulfillmentFulfillment     = orders.OrderItemFulfillmentFulfillment
	OrderItemFulfillmentFulfillmentItem = orders.OrderItemFulfillmentFulfillmentItem
	OrderItemFulfillmentInvoice         = orders.OrderItemFulfillmentInvoice
	OrderItemFulfillmentLastUpdatedBy   = orders.OrderItemFulfillmentLastUpdatedBy
	OrderKpisRecord                     = orders.OrderKpisRecord
	OrderPaymentRecord                  = orders.OrderPaymentRecord
	OrderPaymentAccount                 = orders.OrderPaymentAccount
	OrderPaymentCreatedBy               = orders.OrderPaymentCreatedBy
	OrderPaymentInvoice                 = orders.OrderPaymentInvoice
	OrderPaymentPaymentApplied          = orders.OrderPaymentPaymentApplied
	OrderQuotePartialInvoicesRecord     = orders.OrderQuotePartialInvoicesRecord
)

// OrderCustomObjectRecordShape is the DTOShape reading a body as OrderCustomObjectRecord, for the raw-verb call.
func OrderCustomObjectRecordShape(decoded any) any {
	return orders.OrderCustomObjectRecordShape(decoded)
}

// OrderItemFulfillmentRecordShape is the DTOShape reading a body as OrderItemFulfillmentRecord, for the raw-verb call.
func OrderItemFulfillmentRecordShape(decoded any) any {
	return orders.OrderItemFulfillmentRecordShape(decoded)
}

// OrderKpisRecordShape is the DTOShape reading a body as OrderKpisRecord, for the raw-verb call.
func OrderKpisRecordShape(decoded any) any { return orders.OrderKpisRecordShape(decoded) }

// OrderPaymentRecordShape is the DTOShape reading a body as OrderPaymentRecord, for the raw-verb call.
func OrderPaymentRecordShape(decoded any) any { return orders.OrderPaymentRecordShape(decoded) }

// OrderQuotePartialInvoicesRecordShape is the DTOShape reading a body as OrderQuotePartialInvoicesRecord, for the raw-verb call.
func OrderQuotePartialInvoicesRecordShape(decoded any) any {
	return orders.OrderQuotePartialInvoicesRecordShape(decoded)
}

// More orders path templates and body keys, for the raw-verb call.
const (
	OrdersNotesPath                        = orders.OrdersNotesPath
	OrdersPreviewPath                      = orders.OrdersPreviewPath
	OrdersExpressPath                      = orders.OrdersExpressPath
	OrderAcceptPath                        = orders.OrderAcceptPath
	OrderActivateInvoiceFromDormantPath    = orders.OrderActivateInvoiceFromDormantPath
	OrderChangePreviewPath                 = orders.OrderChangePreviewPath
	OrderChangeStatusPath                  = orders.OrderChangeStatusPath
	OrderContractAdjustmentPath            = orders.OrderContractAdjustmentPath
	OrderContractAdjustmentPreviewPath     = orders.OrderContractAdjustmentPreviewPath
	OrderDowngradePath                     = orders.OrderDowngradePath
	OrderDowngradePreviewPath              = orders.OrderDowngradePreviewPath
	OrderUpgradePath                       = orders.OrderUpgradePath
	OrderUpgradePreviewPath                = orders.OrderUpgradePreviewPath
	OrderTerminateTrialPath                = orders.OrderTerminateTrialPath
	OrderCustomObjectPath                  = orders.OrderCustomObjectPath
	OrderCustomObjectInstancePath          = orders.OrderCustomObjectInstancePath
	OrderCustomObjectsInstancePath         = orders.OrderCustomObjectsInstancePath
	OrderCustomObjectsInstanceInstancePath = orders.OrderCustomObjectsInstanceInstancePath
	OrderItemFulfillmentPath               = orders.OrderItemFulfillmentPath
	OrderKPIsPath                          = orders.OrderKPIsPath
	OrderPaymentPath                       = orders.OrderPaymentPath
	OrderQuotePartialInvoicesPath          = orders.OrderQuotePartialInvoicesPath
	OrderShippingPDFPath                   = orders.OrderShippingPDFPath

	OrdersCustomObjectKey    = orders.CustomObjectKey
	OrdersItemFulfillmentKey = orders.ItemFulfillmentKey
	OrdersOrderKPIsKey       = orders.OrderKPIsKey
	OrdersPaymentKey         = orders.PaymentKey
)

// Organisation is the organisation endpoints bound to a client.
type Organisation = organisation.Organisation

// The organisation views, aliased so a type assertion works under either spelling.
type (
	SettingsOrganisationRecord             = organisation.SettingsOrganisationRecord
	SettingsOrganisationParentOrganization = organisation.SettingsOrganisationParentOrganization
)

// NewOrganisation binds the resource to a requester of your own.
func NewOrganisation(requester Requester) *Organisation { return organisation.New(requester) }

// SettingsOrganisationRecordShape is the DTOShape reading a body as SettingsOrganisationRecord, for the raw-verb call.
func SettingsOrganisationRecordShape(decoded any) any {
	return organisation.SettingsOrganisationRecordShape(decoded)
}

// The organisation path templates and body keys, for the raw-verb call.
const (
	OrganisationSettingsOrganisationPath = organisation.SettingsOrganisationPath

	OrganisationOrganizationKey = organisation.OrganizationKey
)

// Organisations is the organisations endpoints bound to a client.
type Organisations = organisations.Organisations

// NewOrganisations binds the resource to a requester of your own.
func NewOrganisations(requester Requester) *Organisations { return organisations.New(requester) }

// The organisations path templates and body keys, for the raw-verb call.
const (
	SettingsOrganisationsPath             = organisations.SettingsOrganisationsPath
	OrganisationsSettingsOrganisationPath = organisations.SettingsOrganisationPath

	SettingsOrganisationsKey     = organisations.SettingsOrganisationsKey
	OrganisationsOrganizationKey = organisations.OrganizationKey
)

// More payments views, aliased so a type assertion works under either spelling.
type (
	PaymentsNoteFileList   = payments.NoteFileList
	PaymentsNoteAttachment = payments.NoteAttachment
)

// PaymentsNoteFileListShape is the DTOShape reading a body as PaymentsNoteFileList, for the raw-verb call.
func PaymentsNoteFileListShape(decoded any) any { return payments.NoteFileListShape(decoded) }

// More payments path templates and body keys, for the raw-verb call.
const (
	InvoicePaymentPath   = payments.InvoicePaymentPath
	PaymentNoteFilesPath = payments.PaymentNoteFilesPath
)

// More portals path templates and body keys, for the raw-verb call.
const (
	PortalChangePath = portals.PortalChangePath
)

// PricingPeriods is the pricing periods endpoints bound to a client.
type PricingPeriods = pricingperiods.PricingPeriods

// The pricing periods views, aliased so a type assertion works under either spelling.
type (
	SettingsPricingPeriodList       = pricingperiods.SettingsPricingPeriodList
	SettingsPricingPeriodRecord     = pricingperiods.SettingsPricingPeriodRecord
	SettingsPricingPeriodTimePeriod = pricingperiods.SettingsPricingPeriodTimePeriod
	PricingPeriodsPagination        = pricingperiods.Pagination
	PricingPeriodsPaginationLinks   = pricingperiods.PaginationLinks
)

// NewPricingPeriods binds the resource to a requester of your own.
func NewPricingPeriods(requester Requester) *PricingPeriods { return pricingperiods.New(requester) }

// SettingsPricingPeriodListShape is the DTOShape reading a body as SettingsPricingPeriodList, for the raw-verb call.
func SettingsPricingPeriodListShape(decoded any) any {
	return pricingperiods.SettingsPricingPeriodListShape(decoded)
}

// SettingsPricingPeriodRecordShape is the DTOShape reading a body as SettingsPricingPeriodRecord, for the raw-verb call.
func SettingsPricingPeriodRecordShape(decoded any) any {
	return pricingperiods.SettingsPricingPeriodRecordShape(decoded)
}

// The pricing periods path templates and body keys, for the raw-verb call.
const (
	SettingsPricingPeriodsPath = pricingperiods.SettingsPricingPeriodsPath
	SettingsPricingPeriodPath  = pricingperiods.SettingsPricingPeriodPath

	SettingsPricingPeriodsKey = pricingperiods.SettingsPricingPeriodsKey
	PricingPeriodKey          = pricingperiods.PricingPeriodKey
)

// ProformaInvoices is the proforma invoices endpoints bound to a client.
type ProformaInvoices = proformainvoices.ProformaInvoices

// NewProformaInvoices binds the resource to a requester of your own.
func NewProformaInvoices(requester Requester) *ProformaInvoices {
	return proformainvoices.New(requester)
}

// The proforma invoices path templates and body keys, for the raw-verb call.
const (
	ProformaInvoicesPath        = proformainvoices.ProformaInvoicesPath
	ProformaInvoicePath         = proformainvoices.ProformaInvoicePath
	ProformaInvoicePaymentsPath = proformainvoices.ProformaInvoicePaymentsPath

	ProformaInvoicesKey        = proformainvoices.ProformaInvoicesKey
	ProformaInvoiceKey         = proformainvoices.ProformaInvoiceKey
	ProformaInvoicesPaymentKey = proformainvoices.PaymentKey
)

// More purchase credit notes views, aliased so a type assertion works under either spelling.
type (
	PurchaseCreditNoteCreatedRecord = purchasecreditnotes.PurchaseCreditNoteCreatedRecord
	PurchaseCreditNoteCreatedStatus = purchasecreditnotes.PurchaseCreditNoteCreatedStatus
)

// PurchaseCreditNoteCreatedRecordShape is the DTOShape reading a body as PurchaseCreditNoteCreatedRecord, for the raw-verb call.
func PurchaseCreditNoteCreatedRecordShape(decoded any) any {
	return purchasecreditnotes.PurchaseCreditNoteCreatedRecordShape(decoded)
}

// More purchase credit notes path templates and body keys, for the raw-verb call.
const (
	PurchaseCreditNoteApplyPath = purchasecreditnotes.PurchaseCreditNoteApplyPath
	PurchaseCreditNoteNotesPath = purchasecreditnotes.PurchaseCreditNoteNotesPath
	PurchaseCreditNoteNotePath  = purchasecreditnotes.PurchaseCreditNoteNotePath
)

// PurchaseExpress is the purchase express endpoints bound to a client.
type PurchaseExpress = purchaseexpress.PurchaseExpress

// The purchase express views, aliased so a type assertion works under either spelling.
type (
	PurchaseExpressRecord     = purchaseexpress.PurchaseExpressRecord
	PurchaseExpressAccount    = purchaseexpress.PurchaseExpressAccount
	PurchaseExpressCreatedBy  = purchaseexpress.PurchaseExpressCreatedBy
	PurchaseExpressCurrency   = purchaseexpress.PurchaseExpressCurrency
	PurchaseExpressCustomForm = purchaseexpress.PurchaseExpressCustomForm
	PurchaseExpressKpis       = purchaseexpress.PurchaseExpressKpis
)

// NewPurchaseExpress binds the resource to a requester of your own.
func NewPurchaseExpress(requester Requester) *PurchaseExpress { return purchaseexpress.New(requester) }

// PurchaseExpressRecordShape is the DTOShape reading a body as PurchaseExpressRecord, for the raw-verb call.
func PurchaseExpressRecordShape(decoded any) any {
	return purchaseexpress.PurchaseExpressRecordShape(decoded)
}

// The purchase express path templates and body keys, for the raw-verb call.
const (
	PurchaseExpressPath = purchaseexpress.PurchaseExpressPath

	PurchaseExpressKey = purchaseexpress.PurchaseExpressKey
)

// More purchase invoices path templates and body keys, for the raw-verb call.
const (
	PurchaseInvoicePurchaseCreditNoteApplicationPath = purchaseinvoices.PurchaseInvoicePurchaseCreditNoteApplicationPath
)

// More purchase orders views, aliased so a type assertion works under either spelling.
type (
	PurchaseOrderKpisRecord = purchaseorders.PurchaseOrderKpisRecord
)

// PurchaseOrderKpisRecordShape is the DTOShape reading a body as PurchaseOrderKpisRecord, for the raw-verb call.
func PurchaseOrderKpisRecordShape(decoded any) any {
	return purchaseorders.PurchaseOrderKpisRecordShape(decoded)
}

// More purchase orders path templates and body keys, for the raw-verb call.
const (
	PurchaseOrderKPIsPath = purchaseorders.PurchaseOrderKPIsPath

	PurchaseOrderKPIsKey = purchaseorders.PurchaseOrderKPIsKey
)

// More purchase payments views, aliased so a type assertion works under either spelling.
type (
	PurchasePaymentNoteUpdatedRecord          = purchasepayments.PurchasePaymentNoteUpdatedRecord
	PurchasePaymentNoteUpdatedAttachment      = purchasepayments.PurchasePaymentNoteUpdatedAttachment
	PurchasePaymentNoteUpdatedUser            = purchasepayments.PurchasePaymentNoteUpdatedUser
	PurchasePaymentNoteUpdatedPurchasePayment = purchasepayments.PurchasePaymentNoteUpdatedPurchasePayment
)

// PurchasePaymentNoteUpdatedRecordShape is the DTOShape reading a body as PurchasePaymentNoteUpdatedRecord, for the raw-verb call.
func PurchasePaymentNoteUpdatedRecordShape(decoded any) any {
	return purchasepayments.PurchasePaymentNoteUpdatedRecordShape(decoded)
}

// More purchase payments path templates and body keys, for the raw-verb call.
const (
	PurchasePaymentPurchaseCreditNotesPath = purchasepayments.PurchasePaymentPurchaseCreditNotesPath
)

// More purchase refunds path templates and body keys, for the raw-verb call.
const (
	PurchaseCreditNotePurchaseRefundsPath = purchaserefunds.PurchaseCreditNotePurchaseRefundsPath
)

// QuoteOrders is the quote orders endpoints bound to a client.
type QuoteOrders = quoteorders.QuoteOrders

// The quote orders views, aliased so a type assertion works under either spelling.
type (
	QuoteOrderInvoiceList       = quoteorders.QuoteOrderInvoiceList
	QuoteOrderInvoice           = quoteorders.QuoteOrderInvoice
	QuoteOrderInvoiceAccount    = quoteorders.QuoteOrderInvoiceAccount
	QuoteOrderInvoiceCreatedBy  = quoteorders.QuoteOrderInvoiceCreatedBy
	QuoteOrderInvoiceCurrency   = quoteorders.QuoteOrderInvoiceCurrency
	QuoteOrderInvoiceCustomForm = quoteorders.QuoteOrderInvoiceCustomForm
	QuoteOrderInvoiceOrder      = quoteorders.QuoteOrderInvoiceOrder
	QuoteOrdersPagination       = quoteorders.Pagination
	QuoteOrdersPaginationLinks  = quoteorders.PaginationLinks
)

// NewQuoteOrders binds the resource to a requester of your own.
func NewQuoteOrders(requester Requester) *QuoteOrders { return quoteorders.New(requester) }

// QuoteOrderInvoiceListShape is the DTOShape reading a body as QuoteOrderInvoiceList, for the raw-verb call.
func QuoteOrderInvoiceListShape(decoded any) any {
	return quoteorders.QuoteOrderInvoiceListShape(decoded)
}

// The quote orders path templates and body keys, for the raw-verb call.
const (
	QuoteOrderInvoicesPath = quoteorders.QuoteOrderInvoicesPath

	QuoteOrderInvoicesKey = quoteorders.QuoteOrderInvoicesKey
)

// ReceiveTerms is the receive terms endpoints bound to a client.
type ReceiveTerms = receiveterms.ReceiveTerms

// The receive terms views, aliased so a type assertion works under either spelling.
type (
	SettingsReceiveTermList     = receiveterms.SettingsReceiveTermList
	SettingsReceiveTermRecord   = receiveterms.SettingsReceiveTermRecord
	ReceiveTermsPagination      = receiveterms.Pagination
	ReceiveTermsPaginationLinks = receiveterms.PaginationLinks
)

// NewReceiveTerms binds the resource to a requester of your own.
func NewReceiveTerms(requester Requester) *ReceiveTerms { return receiveterms.New(requester) }

// SettingsReceiveTermListShape is the DTOShape reading a body as SettingsReceiveTermList, for the raw-verb call.
func SettingsReceiveTermListShape(decoded any) any {
	return receiveterms.SettingsReceiveTermListShape(decoded)
}

// SettingsReceiveTermRecordShape is the DTOShape reading a body as SettingsReceiveTermRecord, for the raw-verb call.
func SettingsReceiveTermRecordShape(decoded any) any {
	return receiveterms.SettingsReceiveTermRecordShape(decoded)
}

// The receive terms path templates and body keys, for the raw-verb call.
const (
	SettingsReceiveTermsPath = receiveterms.SettingsReceiveTermsPath
	SettingsReceiveTermPath  = receiveterms.SettingsReceiveTermPath

	SettingsReceiveTermsKey = receiveterms.SettingsReceiveTermsKey
	ReceiveTermKey          = receiveterms.ReceiveTermKey
)

// Reconciliations is the reconciliations endpoints bound to a client.
type Reconciliations = reconciliations.Reconciliations

// NewReconciliations binds the resource to a requester of your own.
func NewReconciliations(requester Requester) *Reconciliations { return reconciliations.New(requester) }

// The reconciliations path templates and body keys, for the raw-verb call.
const (
	ReconciliationSalePaymentsPath = reconciliations.ReconciliationSalePaymentsPath
	ReconciliationSaleRefundsPath  = reconciliations.ReconciliationSaleRefundsPath

	ReconciliationSalePaymentsKey = reconciliations.ReconciliationSalePaymentsKey
)

// More refunds views, aliased so a type assertion works under either spelling.
type (
	RefundNoteList        = refunds.RefundNoteList
	RefundNoteRecord      = refunds.RefundNoteRecord
	RefundNoteAttachment  = refunds.RefundNoteAttachment
	RefundNoteCreatedBy   = refunds.RefundNoteCreatedBy
	RefundNoteDetail      = refunds.RefundNoteDetail
	RefundNoteWatcher     = refunds.RefundNoteWatcher
	RefundsNoteFileList   = refunds.NoteFileList
	RefundsNoteAttachment = refunds.NoteAttachment
	RefundNoteFileRecord  = refunds.RefundNoteFileRecord
)

// RefundNoteListShape is the DTOShape reading a body as RefundNoteList, for the raw-verb call.
func RefundNoteListShape(decoded any) any { return refunds.RefundNoteListShape(decoded) }

// RefundNoteRecordShape is the DTOShape reading a body as RefundNoteRecord, for the raw-verb call.
func RefundNoteRecordShape(decoded any) any { return refunds.RefundNoteRecordShape(decoded) }

// RefundsNoteFileListShape is the DTOShape reading a body as RefundsNoteFileList, for the raw-verb call.
func RefundsNoteFileListShape(decoded any) any { return refunds.NoteFileListShape(decoded) }

// RefundNoteFileRecordShape is the DTOShape reading a body as RefundNoteFileRecord, for the raw-verb call.
func RefundNoteFileRecordShape(decoded any) any { return refunds.RefundNoteFileRecordShape(decoded) }

// More refunds path templates and body keys, for the raw-verb call.
const (
	RefundNotesPath     = refunds.RefundNotesPath
	RefundNotePath      = refunds.RefundNotePath
	RefundNoteFilesPath = refunds.RefundNoteFilesPath
	RefundNoteFilePath  = refunds.RefundNoteFilePath

	RefundsNoteKey = refunds.NoteKey
	RefundsFileKey = refunds.FileKey
)

// More return merchandise authorisations path templates and body keys, for the raw-verb call.
const (
	ReturnMerchandiseAuthorisationNotesPath           = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationNotesPath
	ReturnMerchandiseAuthorisationReceivePath         = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationReceivePath
	ReturnMerchandiseAuthorisationApprovePath         = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationApprovePath
	ReturnMerchandiseAuthorisationNotePath            = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationNotePath
	ReturnMerchandiseAuthorisationNoteFilePath        = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationNoteFilePath
	ReturnMerchandiseAuthorisationRejectPath          = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationRejectPath
	ReturnMerchandiseAuthorisationNotesByCodePath     = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationNotesByCodePath
	ReturnMerchandiseAuthorisationNoteByCodePath      = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationNoteByCodePath
	ReturnMerchandiseAuthorisationNoteFilesByCodePath = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationNoteFilesByCodePath
	ReturnMerchandiseAuthorisationNoteFileByCodePath  = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationNoteFileByCodePath
	ReturnMerchandiseAuthorisationReceiveByCodePath   = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationReceiveByCodePath
	ReturnMerchandiseAuthorisationReceiveByUUIDPath   = returnmerchandiseauthorisations.ReturnMerchandiseAuthorisationReceiveByUUIDPath
)

// ReturnVendorAuthorizations is the return vendor authorizations endpoints bound to a client.
type ReturnVendorAuthorizations = returnvendorauthorizations.ReturnVendorAuthorizations

// The return vendor authorizations views, aliased so a type assertion works under either spelling.
type (
	ReturnVendorAuthorizationView             = returnvendorauthorizations.ReturnVendorAuthorizationView
	ReturnVendorAuthorizationAccount          = returnvendorauthorizations.ReturnVendorAuthorizationAccount
	ReturnVendorAuthorizationCreatedBy        = returnvendorauthorizations.ReturnVendorAuthorizationCreatedBy
	ReturnVendorAuthorizationPurchaseOrder    = returnvendorauthorizations.ReturnVendorAuthorizationPurchaseOrder
	ReturnVendorAuthorizationList             = returnvendorauthorizations.ReturnVendorAuthorizationList
	ReturnVendorAuthorization                 = returnvendorauthorizations.ReturnVendorAuthorization
	ReturnVendorAuthorizationsPagination      = returnvendorauthorizations.Pagination
	ReturnVendorAuthorizationsPaginationLinks = returnvendorauthorizations.PaginationLinks
)

// NewReturnVendorAuthorizations binds the resource to a requester of your own.
func NewReturnVendorAuthorizations(requester Requester) *ReturnVendorAuthorizations {
	return returnvendorauthorizations.New(requester)
}

// ReturnVendorAuthorizationViewShape is the DTOShape reading a body as ReturnVendorAuthorizationView, for the raw-verb call.
func ReturnVendorAuthorizationViewShape(decoded any) any {
	return returnvendorauthorizations.ReturnVendorAuthorizationViewShape(decoded)
}

// ReturnVendorAuthorizationListShape is the DTOShape reading a body as ReturnVendorAuthorizationList, for the raw-verb call.
func ReturnVendorAuthorizationListShape(decoded any) any {
	return returnvendorauthorizations.ReturnVendorAuthorizationListShape(decoded)
}

// The return vendor authorizations path templates and body keys, for the raw-verb call.
const (
	ReturnVendorAuthorizationsPath        = returnvendorauthorizations.ReturnVendorAuthorizationsPath
	ReturnVendorAuthorizationPath         = returnvendorauthorizations.ReturnVendorAuthorizationPath
	ReturnVendorAuthorizationNotePath     = returnvendorauthorizations.ReturnVendorAuthorizationNotePath
	ReturnVendorAuthorizationNoteFilePath = returnvendorauthorizations.ReturnVendorAuthorizationNoteFilePath

	ReturnVendorAuthorizationsKey = returnvendorauthorizations.ReturnVendorAuthorizationsKey
	ReturnVendorAuthorizationKey  = returnvendorauthorizations.ReturnVendorAuthorizationKey
)

// Roles is the roles endpoints bound to a client.
type Roles = roles.Roles

// The roles views, aliased so a type assertion works under either spelling.
type (
	SettingsRoleList     = roles.SettingsRoleList
	SettingsRoleRecord   = roles.SettingsRoleRecord
	RolesPagination      = roles.Pagination
	RolesPaginationLinks = roles.PaginationLinks
)

// NewRoles binds the resource to a requester of your own.
func NewRoles(requester Requester) *Roles { return roles.New(requester) }

// SettingsRoleListShape is the DTOShape reading a body as SettingsRoleList, for the raw-verb call.
func SettingsRoleListShape(decoded any) any { return roles.SettingsRoleListShape(decoded) }

// SettingsRoleRecordShape is the DTOShape reading a body as SettingsRoleRecord, for the raw-verb call.
func SettingsRoleRecordShape(decoded any) any { return roles.SettingsRoleRecordShape(decoded) }

// The roles path templates and body keys, for the raw-verb call.
const (
	SettingsRolesPath = roles.SettingsRolesPath
	SettingsRolePath  = roles.SettingsRolePath

	SettingsRolesKey = roles.SettingsRolesKey
	RoleKey          = roles.RoleKey
)

// SalesReps is the sales reps endpoints bound to a client.
type SalesReps = salesreps.SalesReps

// The sales reps views, aliased so a type assertion works under either spelling.
type (
	SettingsSalesRepList     = salesreps.SettingsSalesRepList
	SettingsSalesRepRecord   = salesreps.SettingsSalesRepRecord
	SalesRepsPagination      = salesreps.Pagination
	SalesRepsPaginationLinks = salesreps.PaginationLinks
)

// NewSalesReps binds the resource to a requester of your own.
func NewSalesReps(requester Requester) *SalesReps { return salesreps.New(requester) }

// SettingsSalesRepListShape is the DTOShape reading a body as SettingsSalesRepList, for the raw-verb call.
func SettingsSalesRepListShape(decoded any) any { return salesreps.SettingsSalesRepListShape(decoded) }

// SettingsSalesRepRecordShape is the DTOShape reading a body as SettingsSalesRepRecord, for the raw-verb call.
func SettingsSalesRepRecordShape(decoded any) any {
	return salesreps.SettingsSalesRepRecordShape(decoded)
}

// The sales reps path templates and body keys, for the raw-verb call.
const (
	SettingsSalesRepsPath = salesreps.SettingsSalesRepsPath
	SettingsSalesRepPath  = salesreps.SettingsSalesRepPath

	SettingsSalesRepsKey = salesreps.SettingsSalesRepsKey
	SalesRepKey          = salesreps.SalesRepKey
)

// Salutations is the salutations endpoints bound to a client.
type Salutations = salutations.Salutations

// The salutations views, aliased so a type assertion works under either spelling.
type (
	SettingsSalutationList     = salutations.SettingsSalutationList
	SettingsSalutationRecord   = salutations.SettingsSalutationRecord
	SalutationsPagination      = salutations.Pagination
	SalutationsPaginationLinks = salutations.PaginationLinks
)

// NewSalutations binds the resource to a requester of your own.
func NewSalutations(requester Requester) *Salutations { return salutations.New(requester) }

// SettingsSalutationListShape is the DTOShape reading a body as SettingsSalutationList, for the raw-verb call.
func SettingsSalutationListShape(decoded any) any {
	return salutations.SettingsSalutationListShape(decoded)
}

// SettingsSalutationRecordShape is the DTOShape reading a body as SettingsSalutationRecord, for the raw-verb call.
func SettingsSalutationRecordShape(decoded any) any {
	return salutations.SettingsSalutationRecordShape(decoded)
}

// The salutations path templates and body keys, for the raw-verb call.
const (
	SettingsSalutationsPath = salutations.SettingsSalutationsPath
	SettingsSalutationPath  = salutations.SettingsSalutationPath

	SettingsSalutationsKey = salutations.SettingsSalutationsKey
	SalutationKey          = salutations.SalutationKey
)

// More settings views, aliased so a type assertion works under either spelling.
type (
	SettingsCustomObjectDefinitionList                                                      = settings.SettingsCustomObjectDefinitionList
	SettingsCustomObjectDefinition                                                          = settings.SettingsCustomObjectDefinition
	SettingsCustomObjectDefinitionReadRecord                                                = settings.SettingsCustomObjectDefinitionReadRecord
	SettingsCustomLogicScriptSavedRecord                                                    = settings.SettingsCustomLogicScriptSavedRecord
	SettingsDiscountProfileReadRecord                                                       = settings.SettingsDiscountProfileReadRecord
	SettingsDiscountProfileAccountingCode                                                   = settings.SettingsDiscountProfileAccountingCode
	SettingsDiscountProfileTaxCode                                                          = settings.SettingsDiscountProfileTaxCode
	SettingsFieldSettingList                                                                = settings.SettingsFieldSettingList
	SettingsFieldSettingsItemData                                                           = settings.SettingsFieldSettingsItemData
	SettingsFieldSettingReadRecord                                                          = settings.SettingsFieldSettingReadRecord
	SettingsGeneralSettingsRecord                                                           = settings.SettingsGeneralSettingsRecord
	SettingsFieldSettingsItemGeneralSettingsAccount                                         = settings.SettingsFieldSettingsItemGeneralSettingsAccount
	SettingsFieldSettingsItemGeneralSettingsAccountAccountGrouping                          = settings.SettingsFieldSettingsItemGeneralSettingsAccountAccountGrouping
	SettingsFieldSettingsItemGeneralSettingsAccountAccountManagement                        = settings.SettingsFieldSettingsItemGeneralSettingsAccountAccountManagement
	SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact1                       = settings.SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact1
	SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact2                       = settings.SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact2
	SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact3                       = settings.SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact3
	SettingsFieldSettingsItemGeneralSettingsAccountBillingContact                           = settings.SettingsFieldSettingsItemGeneralSettingsAccountBillingContact
	SettingsFieldSettingsItemGeneralSettingsAccountCreationOfSubsidiaryAccount              = settings.SettingsFieldSettingsItemGeneralSettingsAccountCreationOfSubsidiaryAccount
	SettingsFieldSettingsItemGeneralSettingsAccountReferralTrackingAtAccountLevel           = settings.SettingsFieldSettingsItemGeneralSettingsAccountReferralTrackingAtAccountLevel
	SettingsFieldSettingsItemGeneralSettingsAccountSalesRepTrackingAtAccountLevel           = settings.SettingsFieldSettingsItemGeneralSettingsAccountSalesRepTrackingAtAccountLevel
	SettingsFieldSettingsItemGeneralSettingsAccountShippingContact                          = settings.SettingsFieldSettingsItemGeneralSettingsAccountShippingContact
	SettingsGeneralSettingsDashboard                                                        = settings.SettingsGeneralSettingsDashboard
	SettingsGeneralSettingsDateAndTime                                                      = settings.SettingsGeneralSettingsDateAndTime
	SettingsGeneralSettingsItemUpdatedDetail                                                = settings.SettingsGeneralSettingsItemUpdatedDetail
	SettingsGeneralSettingsItemUpdatedDetailAvailableChargeTypes                            = settings.SettingsGeneralSettingsItemUpdatedDetailAvailableChargeTypes
	SettingsGeneralSettingsItemUpdatedDetailAvailablePricingMethods                         = settings.SettingsGeneralSettingsItemUpdatedDetailAvailablePricingMethods
	SettingsGeneralSettingsItemUpdatedDetailInventoryAdjustmentRules                        = settings.SettingsGeneralSettingsItemUpdatedDetailInventoryAdjustmentRules
	SettingsGeneralSettingsItemUpdatedDetailItemGrouping                                    = settings.SettingsGeneralSettingsItemUpdatedDetailItemGrouping
	SettingsGeneralSettingsItemUpdatedDetailItemManagement                                  = settings.SettingsGeneralSettingsItemUpdatedDetailItemManagement
	SettingsGeneralSettingsItemUpdatedDetailPreOrder                                        = settings.SettingsGeneralSettingsItemUpdatedDetailPreOrder
	SettingsGeneralSettingsItemUpdatedDetailPricingPeriodCanVaryFromChargingPeriod          = settings.SettingsGeneralSettingsItemUpdatedDetailPricingPeriodCanVaryFromChargingPeriod
	SettingsGeneralSettingsItemUpdatedDetailPurchasePriceEnteredIsInclusiveOfTax            = settings.SettingsGeneralSettingsItemUpdatedDetailPurchasePriceEnteredIsInclusiveOfTax
	SettingsGeneralSettingsItemUpdatedDetailPurchasePriceEnteredIsInclusiveOfTaxDetail      = settings.SettingsGeneralSettingsItemUpdatedDetailPurchasePriceEnteredIsInclusiveOfTaxDetail
	SettingsGeneralSettingsItemUpdatedDetailSalePriceEnteredIsInclusiveOfTax                = settings.SettingsGeneralSettingsItemUpdatedDetailSalePriceEnteredIsInclusiveOfTax
	SettingsGeneralSettingsItemUpdatedDetailUseVolumePricing                                = settings.SettingsGeneralSettingsItemUpdatedDetailUseVolumePricing
	SettingsGeneralSettingsLabour                                                           = settings.SettingsGeneralSettingsLabour
	SettingsGeneralSettingsOrder                                                            = settings.SettingsGeneralSettingsOrder
	SettingsGeneralSettingsOrderOrderManagement                                             = settings.SettingsGeneralSettingsOrderOrderManagement
	SettingsGeneralSettingsOrderReferralTrackingAtOrderLevel                                = settings.SettingsGeneralSettingsOrderReferralTrackingAtOrderLevel
	SettingsGeneralSettingsOrderSalesRepTrackingAtOrderLevel                                = settings.SettingsGeneralSettingsOrderSalesRepTrackingAtOrderLevel
	SettingsGeneralSettingsPortal                                                           = settings.SettingsGeneralSettingsPortal
	SettingsGeneralSettingsItemUpdatedProduct                                               = settings.SettingsGeneralSettingsItemUpdatedProduct
	SettingsGeneralSettingsPurchase                                                         = settings.SettingsGeneralSettingsPurchase
	SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseCreditNoteFields     = settings.SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseCreditNoteFields
	SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseInvoiceFields        = settings.SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseInvoiceFields
	SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchasePaymentFields        = settings.SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchasePaymentFields
	SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseRefundFields         = settings.SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseRefundFields
	SettingsGeneralSettingsSecurity                                                         = settings.SettingsGeneralSettingsSecurity
	SettingsGeneralSettingsTransaction                                                      = settings.SettingsGeneralSettingsTransaction
	SettingsGeneralSettingsTransactionAllowModificationOfFollowingCreditNoteFields          = settings.SettingsGeneralSettingsTransactionAllowModificationOfFollowingCreditNoteFields
	SettingsGeneralSettingsTransactionAllowModificationOfFollowingInvoiceFields             = settings.SettingsGeneralSettingsTransactionAllowModificationOfFollowingInvoiceFields
	SettingsGeneralSettingsTransactionAllowModificationOfFollowingPaymentFields             = settings.SettingsGeneralSettingsTransactionAllowModificationOfFollowingPaymentFields
	SettingsGeneralSettingsTransactionAllowModificationOfFollowingRefundFields              = settings.SettingsGeneralSettingsTransactionAllowModificationOfFollowingRefundFields
	SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreference            = settings.SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreference
	SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceEmail       = settings.SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceEmail
	SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferencePostalMail  = settings.SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferencePostalMail
	SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceTextMessage = settings.SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceTextMessage
	SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceVoiceMail   = settings.SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceVoiceMail
	SettingsGeneralSettingsTransactionDefaultAccountingCodes                                = settings.SettingsGeneralSettingsTransactionDefaultAccountingCodes
	SettingsPaymentProcessorSavedRecord                                                     = settings.SettingsPaymentProcessorSavedRecord
	SettingsPaymentProcessorSavedCurrency                                                   = settings.SettingsPaymentProcessorSavedCurrency
	SettingsPaymentProcessorReadRecord                                                      = settings.SettingsPaymentProcessorReadRecord
	SettingsPaymentProcessorAccountingCode                                                  = settings.SettingsPaymentProcessorAccountingCode
	SettingsPaymentProcessorCurrency                                                        = settings.SettingsPaymentProcessorCurrency
	SettingsPaymentTermCreatedRecord                                                        = settings.SettingsPaymentTermCreatedRecord
	SettingsPaymentTermReadRecord                                                           = settings.SettingsPaymentTermReadRecord
	SettingsPaymentTermDuration                                                             = settings.SettingsPaymentTermDuration
	SettingsPaymentTermUpdatedRecord                                                        = settings.SettingsPaymentTermUpdatedRecord
	SettingsPricingLevelCreatedRecord                                                       = settings.SettingsPricingLevelCreatedRecord
	SettingsPricingLevelReadRecord                                                          = settings.SettingsPricingLevelReadRecord
	SettingsPricingLevelTaxCode                                                             = settings.SettingsPricingLevelTaxCode
	SettingsShippingProfileReadRecord                                                       = settings.SettingsShippingProfileReadRecord
	SettingsShippingProfileAccountingCode                                                   = settings.SettingsShippingProfileAccountingCode
	SettingsShippingProfileTaxRule                                                          = settings.SettingsShippingProfileTaxRule
	SettingsWarehouseReadRecord                                                             = settings.SettingsWarehouseReadRecord
	SettingsWarehouseCountry                                                                = settings.SettingsWarehouseCountry
	SettingsTaxReadRecord                                                                   = settings.SettingsTaxReadRecord
	SettingsTaxCountry                                                                      = settings.SettingsTaxCountry
	SettingsTaxCountryTaxByState                                                            = settings.SettingsTaxCountryTaxByState
)

// SettingsCustomObjectDefinitionListShape is the DTOShape reading a body as SettingsCustomObjectDefinitionList, for the raw-verb call.
func SettingsCustomObjectDefinitionListShape(decoded any) any {
	return settings.SettingsCustomObjectDefinitionListShape(decoded)
}

// SettingsCustomObjectDefinitionReadRecordShape is the DTOShape reading a body as SettingsCustomObjectDefinitionReadRecord, for the raw-verb call.
func SettingsCustomObjectDefinitionReadRecordShape(decoded any) any {
	return settings.SettingsCustomObjectDefinitionReadRecordShape(decoded)
}

// SettingsCustomLogicScriptSavedRecordShape is the DTOShape reading a body as SettingsCustomLogicScriptSavedRecord, for the raw-verb call.
func SettingsCustomLogicScriptSavedRecordShape(decoded any) any {
	return settings.SettingsCustomLogicScriptSavedRecordShape(decoded)
}

// SettingsDiscountProfileReadRecordShape is the DTOShape reading a body as SettingsDiscountProfileReadRecord, for the raw-verb call.
func SettingsDiscountProfileReadRecordShape(decoded any) any {
	return settings.SettingsDiscountProfileReadRecordShape(decoded)
}

// SettingsFieldSettingListShape is the DTOShape reading a body as SettingsFieldSettingList, for the raw-verb call.
func SettingsFieldSettingListShape(decoded any) any {
	return settings.SettingsFieldSettingListShape(decoded)
}

// SettingsFieldSettingReadRecordShape is the DTOShape reading a body as SettingsFieldSettingReadRecord, for the raw-verb call.
func SettingsFieldSettingReadRecordShape(decoded any) any {
	return settings.SettingsFieldSettingReadRecordShape(decoded)
}

// SettingsGeneralSettingsRecordShape is the DTOShape reading a body as SettingsGeneralSettingsRecord, for the raw-verb call.
func SettingsGeneralSettingsRecordShape(decoded any) any {
	return settings.SettingsGeneralSettingsRecordShape(decoded)
}

// SettingsPaymentProcessorSavedRecordShape is the DTOShape reading a body as SettingsPaymentProcessorSavedRecord, for the raw-verb call.
func SettingsPaymentProcessorSavedRecordShape(decoded any) any {
	return settings.SettingsPaymentProcessorSavedRecordShape(decoded)
}

// SettingsPaymentProcessorReadRecordShape is the DTOShape reading a body as SettingsPaymentProcessorReadRecord, for the raw-verb call.
func SettingsPaymentProcessorReadRecordShape(decoded any) any {
	return settings.SettingsPaymentProcessorReadRecordShape(decoded)
}

// SettingsPaymentTermCreatedRecordShape is the DTOShape reading a body as SettingsPaymentTermCreatedRecord, for the raw-verb call.
func SettingsPaymentTermCreatedRecordShape(decoded any) any {
	return settings.SettingsPaymentTermCreatedRecordShape(decoded)
}

// SettingsPaymentTermReadRecordShape is the DTOShape reading a body as SettingsPaymentTermReadRecord, for the raw-verb call.
func SettingsPaymentTermReadRecordShape(decoded any) any {
	return settings.SettingsPaymentTermReadRecordShape(decoded)
}

// SettingsPaymentTermUpdatedRecordShape is the DTOShape reading a body as SettingsPaymentTermUpdatedRecord, for the raw-verb call.
func SettingsPaymentTermUpdatedRecordShape(decoded any) any {
	return settings.SettingsPaymentTermUpdatedRecordShape(decoded)
}

// SettingsPricingLevelCreatedRecordShape is the DTOShape reading a body as SettingsPricingLevelCreatedRecord, for the raw-verb call.
func SettingsPricingLevelCreatedRecordShape(decoded any) any {
	return settings.SettingsPricingLevelCreatedRecordShape(decoded)
}

// SettingsPricingLevelReadRecordShape is the DTOShape reading a body as SettingsPricingLevelReadRecord, for the raw-verb call.
func SettingsPricingLevelReadRecordShape(decoded any) any {
	return settings.SettingsPricingLevelReadRecordShape(decoded)
}

// SettingsShippingProfileReadRecordShape is the DTOShape reading a body as SettingsShippingProfileReadRecord, for the raw-verb call.
func SettingsShippingProfileReadRecordShape(decoded any) any {
	return settings.SettingsShippingProfileReadRecordShape(decoded)
}

// SettingsWarehouseReadRecordShape is the DTOShape reading a body as SettingsWarehouseReadRecord, for the raw-verb call.
func SettingsWarehouseReadRecordShape(decoded any) any {
	return settings.SettingsWarehouseReadRecordShape(decoded)
}

// SettingsTaxReadRecordShape is the DTOShape reading a body as SettingsTaxReadRecord, for the raw-verb call.
func SettingsTaxReadRecordShape(decoded any) any { return settings.SettingsTaxReadRecordShape(decoded) }

// More settings path templates and body keys, for the raw-verb call.
const (
	SettingsDiscountProfilePath            = settings.SettingsDiscountProfilePath
	SettingsFieldSettingsPath              = settings.SettingsFieldSettingsPath
	SettingsFieldSettingPath               = settings.SettingsFieldSettingPath
	SettingsGeneralSettingsGroupPath       = settings.SettingsGeneralSettingsGroupPath
	SettingsGeneralSettingPath             = settings.SettingsGeneralSettingPath
	SettingsGeneralSettingsDashboardPath   = settings.SettingsGeneralSettingsDashboardPath
	SettingsGeneralSettingsDateAndTimePath = settings.SettingsGeneralSettingsDateAndTimePath
	SettingsGeneralSettingsDisplayPath     = settings.SettingsGeneralSettingsDisplayPath
	SettingsGeneralSettingsGeneralPath     = settings.SettingsGeneralSettingsGeneralPath
	SettingsGeneralSettingsLabourPath      = settings.SettingsGeneralSettingsLabourPath
	SettingsGeneralSettingsOrderPath       = settings.SettingsGeneralSettingsOrderPath
	SettingsGeneralSettingsPortalPath      = settings.SettingsGeneralSettingsPortalPath
	SettingsGeneralSettingsPurchasePath    = settings.SettingsGeneralSettingsPurchasePath
	SettingsGeneralSettingsSecurityPath    = settings.SettingsGeneralSettingsSecurityPath
	SettingsGeneralSettingsTransactionPath = settings.SettingsGeneralSettingsTransactionPath
	SettingsPaymentProcessorPath           = settings.SettingsPaymentProcessorPath
	SettingsPaymentTermPath                = settings.SettingsPaymentTermPath
	SettingsPricingLevelPath               = settings.SettingsPricingLevelPath
	SettingsShippingProfilePath            = settings.SettingsShippingProfilePath
	SettingsWarehousePath                  = settings.SettingsWarehousePath
	SettingsTaxPath                        = settings.SettingsTaxPath

	DiscountProfileKey  = settings.DiscountProfileKey
	FieldSettingKey     = settings.FieldSettingKey
	PaymentProcessorKey = settings.PaymentProcessorKey
	PaymentTermKey      = settings.PaymentTermKey
	PricingLevelKey     = settings.PricingLevelKey
	ShippingProfileKey  = settings.ShippingProfileKey
	WarehouseKey        = settings.WarehouseKey
	TaxKey              = settings.TaxKey
)

// SettingsWorkflows is the settings workflows endpoints bound to a client.
type SettingsWorkflows = settingsworkflows.SettingsWorkflows

// The settings workflows views, aliased so a type assertion works under either spelling.
type (
	SettingsWorkflowList             = settingsworkflows.SettingsWorkflowList
	SettingsWorkflowRecord           = settingsworkflows.SettingsWorkflowRecord
	SettingsWorkflowsPagination      = settingsworkflows.Pagination
	SettingsWorkflowsPaginationLinks = settingsworkflows.PaginationLinks
)

// NewSettingsWorkflows binds the resource to a requester of your own.
func NewSettingsWorkflows(requester Requester) *SettingsWorkflows {
	return settingsworkflows.New(requester)
}

// SettingsWorkflowListShape is the DTOShape reading a body as SettingsWorkflowList, for the raw-verb call.
func SettingsWorkflowListShape(decoded any) any {
	return settingsworkflows.SettingsWorkflowListShape(decoded)
}

// SettingsWorkflowRecordShape is the DTOShape reading a body as SettingsWorkflowRecord, for the raw-verb call.
func SettingsWorkflowRecordShape(decoded any) any {
	return settingsworkflows.SettingsWorkflowRecordShape(decoded)
}

// The settings workflows path templates and body keys, for the raw-verb call.
const (
	SettingsWorkflowsPath = settingsworkflows.SettingsWorkflowsPath
	SettingsWorkflowPath  = settingsworkflows.SettingsWorkflowPath

	SettingsWorkflowsKey = settingsworkflows.SettingsWorkflowsKey
	WorkflowKey          = settingsworkflows.WorkflowKey
)

// ShippingPackages is the shipping packages endpoints bound to a client.
type ShippingPackages = shippingpackages.ShippingPackages

// The shipping packages views, aliased so a type assertion works under either spelling.
type (
	SettingsShippingPackageList     = shippingpackages.SettingsShippingPackageList
	SettingsShippingPackageRecord   = shippingpackages.SettingsShippingPackageRecord
	ShippingPackagesPagination      = shippingpackages.Pagination
	ShippingPackagesPaginationLinks = shippingpackages.PaginationLinks
)

// NewShippingPackages binds the resource to a requester of your own.
func NewShippingPackages(requester Requester) *ShippingPackages {
	return shippingpackages.New(requester)
}

// SettingsShippingPackageListShape is the DTOShape reading a body as SettingsShippingPackageList, for the raw-verb call.
func SettingsShippingPackageListShape(decoded any) any {
	return shippingpackages.SettingsShippingPackageListShape(decoded)
}

// SettingsShippingPackageRecordShape is the DTOShape reading a body as SettingsShippingPackageRecord, for the raw-verb call.
func SettingsShippingPackageRecordShape(decoded any) any {
	return shippingpackages.SettingsShippingPackageRecordShape(decoded)
}

// The shipping packages path templates and body keys, for the raw-verb call.
const (
	SettingsShippingPackagesPath = shippingpackages.SettingsShippingPackagesPath
	SettingsShippingPackagePath  = shippingpackages.SettingsShippingPackagePath

	SettingsShippingPackagesKey = shippingpackages.SettingsShippingPackagesKey
	ShippingPackageKey          = shippingpackages.ShippingPackageKey
)

// ShippingRules is the shipping rules endpoints bound to a client.
type ShippingRules = shippingrules.ShippingRules

// The shipping rules views, aliased so a type assertion works under either spelling.
type (
	SettingsShippingRuleList          = shippingrules.SettingsShippingRuleList
	SettingsShippingRuleRecord        = shippingrules.SettingsShippingRuleRecord
	SettingsShippingRuleConfiguration = shippingrules.SettingsShippingRuleConfiguration
	ShippingRulesPagination           = shippingrules.Pagination
	ShippingRulesPaginationLinks      = shippingrules.PaginationLinks
)

// NewShippingRules binds the resource to a requester of your own.
func NewShippingRules(requester Requester) *ShippingRules { return shippingrules.New(requester) }

// SettingsShippingRuleListShape is the DTOShape reading a body as SettingsShippingRuleList, for the raw-verb call.
func SettingsShippingRuleListShape(decoded any) any {
	return shippingrules.SettingsShippingRuleListShape(decoded)
}

// SettingsShippingRuleRecordShape is the DTOShape reading a body as SettingsShippingRuleRecord, for the raw-verb call.
func SettingsShippingRuleRecordShape(decoded any) any {
	return shippingrules.SettingsShippingRuleRecordShape(decoded)
}

// The shipping rules path templates and body keys, for the raw-verb call.
const (
	SettingsShippingRulesPath             = shippingrules.SettingsShippingRulesPath
	SettingsShippingRulePath              = shippingrules.SettingsShippingRulePath
	SettingsShippingRuleConfigurationPath = shippingrules.SettingsShippingRuleConfigurationPath

	SettingsShippingRulesKey = shippingrules.SettingsShippingRulesKey
	ShippingRuleKey          = shippingrules.ShippingRuleKey
	ConfigurationKey         = shippingrules.ConfigurationKey
)

// SystemEvents is the system events endpoints bound to a client.
type SystemEvents = systemevents.SystemEvents

// The system events views, aliased so a type assertion works under either spelling.
type (
	SettingsSystemEventList     = systemevents.SettingsSystemEventList
	SettingsSystemEventRecord   = systemevents.SettingsSystemEventRecord
	SystemEventsPagination      = systemevents.Pagination
	SystemEventsPaginationLinks = systemevents.PaginationLinks
)

// NewSystemEvents binds the resource to a requester of your own.
func NewSystemEvents(requester Requester) *SystemEvents { return systemevents.New(requester) }

// SettingsSystemEventListShape is the DTOShape reading a body as SettingsSystemEventList, for the raw-verb call.
func SettingsSystemEventListShape(decoded any) any {
	return systemevents.SettingsSystemEventListShape(decoded)
}

// SettingsSystemEventRecordShape is the DTOShape reading a body as SettingsSystemEventRecord, for the raw-verb call.
func SettingsSystemEventRecordShape(decoded any) any {
	return systemevents.SettingsSystemEventRecordShape(decoded)
}

// The system events path templates and body keys, for the raw-verb call.
const (
	SettingsSystemEventsPath = systemevents.SettingsSystemEventsPath
	SettingsSystemEventPath  = systemevents.SettingsSystemEventPath

	SettingsSystemEventsKey = systemevents.SettingsSystemEventsKey
	SystemEventKey          = systemevents.SystemEventKey
)

// More tasks path templates and body keys, for the raw-verb call.
const (
	TaskNoteFilePath  = tasks.TaskNoteFilePath
	TaskNoteFilesPath = tasks.TaskNoteFilesPath
)

// Teams is the teams endpoints bound to a client.
type Teams = teams.Teams

// The teams views, aliased so a type assertion works under either spelling.
type (
	SettingsTeamList     = teams.SettingsTeamList
	SettingsTeamRecord   = teams.SettingsTeamRecord
	TeamsPagination      = teams.Pagination
	TeamsPaginationLinks = teams.PaginationLinks
)

// NewTeams binds the resource to a requester of your own.
func NewTeams(requester Requester) *Teams { return teams.New(requester) }

// SettingsTeamListShape is the DTOShape reading a body as SettingsTeamList, for the raw-verb call.
func SettingsTeamListShape(decoded any) any { return teams.SettingsTeamListShape(decoded) }

// SettingsTeamRecordShape is the DTOShape reading a body as SettingsTeamRecord, for the raw-verb call.
func SettingsTeamRecordShape(decoded any) any { return teams.SettingsTeamRecordShape(decoded) }

// The teams path templates and body keys, for the raw-verb call.
const (
	SettingsTeamsPath = teams.SettingsTeamsPath
	SettingsTeamPath  = teams.SettingsTeamPath

	SettingsTeamsKey = teams.SettingsTeamsKey
	TeamKey          = teams.TeamKey
)

// TimePeriods is the time periods endpoints bound to a client.
type TimePeriods = timeperiods.TimePeriods

// The time periods views, aliased so a type assertion works under either spelling.
type (
	SettingsTimePeriodList     = timeperiods.SettingsTimePeriodList
	SettingsTimePeriodRecord   = timeperiods.SettingsTimePeriodRecord
	TimePeriodsPagination      = timeperiods.Pagination
	TimePeriodsPaginationLinks = timeperiods.PaginationLinks
)

// NewTimePeriods binds the resource to a requester of your own.
func NewTimePeriods(requester Requester) *TimePeriods { return timeperiods.New(requester) }

// SettingsTimePeriodListShape is the DTOShape reading a body as SettingsTimePeriodList, for the raw-verb call.
func SettingsTimePeriodListShape(decoded any) any {
	return timeperiods.SettingsTimePeriodListShape(decoded)
}

// SettingsTimePeriodRecordShape is the DTOShape reading a body as SettingsTimePeriodRecord, for the raw-verb call.
func SettingsTimePeriodRecordShape(decoded any) any {
	return timeperiods.SettingsTimePeriodRecordShape(decoded)
}

// The time periods path templates and body keys, for the raw-verb call.
const (
	SettingsTimePeriodsPath = timeperiods.SettingsTimePeriodsPath
	SettingsTimePeriodPath  = timeperiods.SettingsTimePeriodPath

	SettingsTimePeriodsKey = timeperiods.SettingsTimePeriodsKey
	TimePeriodKey          = timeperiods.TimePeriodKey
)

// Uom is the uom endpoints bound to a client.
type Uom = uom.Uom

// The uom views, aliased so a type assertion works under either spelling.
type (
	SettingsUomList    = uom.SettingsUomList
	SettingsUomRecord  = uom.SettingsUomRecord
	UomPagination      = uom.Pagination
	UomPaginationLinks = uom.PaginationLinks
)

// NewUom binds the resource to a requester of your own.
func NewUom(requester Requester) *Uom { return uom.New(requester) }

// SettingsUomListShape is the DTOShape reading a body as SettingsUomList, for the raw-verb call.
func SettingsUomListShape(decoded any) any { return uom.SettingsUomListShape(decoded) }

// SettingsUomRecordShape is the DTOShape reading a body as SettingsUomRecord, for the raw-verb call.
func SettingsUomRecordShape(decoded any) any { return uom.SettingsUomRecordShape(decoded) }

// The uom path templates and body keys, for the raw-verb call.
const (
	SettingsUomsPath = uom.SettingsUomsPath
	SettingsUomPath  = uom.SettingsUomPath

	SettingsUomsKey = uom.SettingsUomsKey
	UomKey          = uom.UomKey
)

// Webhooks is the webhooks endpoints bound to a client.
type Webhooks = webhooks.Webhooks

// NewWebhooks binds the resource to a requester of your own.
func NewWebhooks(requester Requester) *Webhooks { return webhooks.New(requester) }

// The webhooks path templates and body keys, for the raw-verb call.
const (
	WebhookPath = webhooks.WebhookPath

	WebhookKey = webhooks.WebhookKey
)

// More workflows views, aliased so a type assertion works under either spelling.
type (
	WorkflowList                        = workflows.WorkflowList
	WorkflowView                        = workflows.WorkflowView
	WorkflowCurrentVersion              = workflows.WorkflowCurrentVersion
	WorkflowCreatedBy                   = workflows.WorkflowCreatedBy
	WorkflowLastUpdatedBy               = workflows.WorkflowLastUpdatedBy
	WorkflowMilestoneView               = workflows.WorkflowMilestoneView
	WorkflowMilestoneActivityDefinition = workflows.WorkflowMilestoneActivityDefinition
	WorkflowMilestoneDetail             = workflows.WorkflowMilestoneDetail
	WorkflowMilestoneWorkflowVersion    = workflows.WorkflowMilestoneWorkflowVersion
)

// WorkflowListShape is the DTOShape reading a body as WorkflowList, for the raw-verb call.
func WorkflowListShape(decoded any) any { return workflows.WorkflowListShape(decoded) }

// WorkflowViewShape is the DTOShape reading a body as WorkflowView, for the raw-verb call.
func WorkflowViewShape(decoded any) any { return workflows.WorkflowViewShape(decoded) }

// WorkflowMilestoneViewShape is the DTOShape reading a body as WorkflowMilestoneView, for the raw-verb call.
func WorkflowMilestoneViewShape(decoded any) any {
	return workflows.WorkflowMilestoneViewShape(decoded)
}

// More workflows path templates and body keys, for the raw-verb call.
const (
	WorkflowsPath                         = workflows.WorkflowsPath
	WorkflowPath                          = workflows.WorkflowPath
	WorkflowStagesPath                    = workflows.WorkflowStagesPath
	WorkflowStagePath                     = workflows.WorkflowStagePath
	WorkflowMilestonePath                 = workflows.WorkflowMilestonePath
	WorkflowCustomComponentHistoryPath    = workflows.WorkflowCustomComponentHistoryPath
	WorkflowCustomComponentMilestonesPath = workflows.WorkflowCustomComponentMilestonesPath
	WorkflowCustomComponentMilestonePath  = workflows.WorkflowCustomComponentMilestonePath

	WorkflowMilestoneKey = workflows.WorkflowMilestoneKey
)

// Zones is the zones endpoints bound to a client.
type Zones = zones.Zones

// The zones views, aliased so a type assertion works under either spelling.
type (
	SettingsZoneList     = zones.SettingsZoneList
	SettingsZoneRecord   = zones.SettingsZoneRecord
	ZonesPagination      = zones.Pagination
	ZonesPaginationLinks = zones.PaginationLinks
)

// NewZones binds the resource to a requester of your own.
func NewZones(requester Requester) *Zones { return zones.New(requester) }

// SettingsZoneListShape is the DTOShape reading a body as SettingsZoneList, for the raw-verb call.
func SettingsZoneListShape(decoded any) any { return zones.SettingsZoneListShape(decoded) }

// SettingsZoneRecordShape is the DTOShape reading a body as SettingsZoneRecord, for the raw-verb call.
func SettingsZoneRecordShape(decoded any) any { return zones.SettingsZoneRecordShape(decoded) }

// The zones path templates and body keys, for the raw-verb call.
const (
	SettingsZonesPath = zones.SettingsZonesPath
	SettingsZonePath  = zones.SettingsZonePath

	SettingsZonesKey = zones.SettingsZonesKey
	ZoneKey          = zones.ZoneKey
)
