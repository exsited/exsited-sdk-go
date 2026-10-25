package externaldatabases

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// ExternalDatabases is the v4 external databases endpoints, bound to one client.
type ExternalDatabases struct {
	requester exhttp.Requester
}

// GetExternalDatabases sends GET /api/{api_version}/external_databases.
func (e *ExternalDatabases) GetExternalDatabases(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return e.requester.SendGetRequest(ctx, ExternalDatabasesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ExternalDatabasesRecordShape})
}

// CreateExternalDatabase sends POST /api/{api_version}/external_databases.
func (e *ExternalDatabases) CreateExternalDatabase(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return e.requester.SendPostRequest(ctx, ExternalDatabasesPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ExternalDatabaseCreateRecordShape})
}

// GetExternalDatabase sends GET /api/{api_version}/external_databases/{db_name}.
func (e *ExternalDatabases) GetExternalDatabase(ctx context.Context, dbName any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(dbName, "db_name"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return e.requester.SendGetRequest(ctx, ExternalDatabasePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"db_name": dbName}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ExternalDatabasesRecordShape})
}

// CreateExternalDatabaseInsertRecord sends POST /api/{api_version}/external_databases/{db_name}/insert_records.
func (e *ExternalDatabases) CreateExternalDatabaseInsertRecord(ctx context.Context, dbName any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(dbName, "db_name"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return e.requester.SendPostRequest(ctx, ExternalDatabaseInsertRecordsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"db_name": dbName}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ExternalDatabaseInsertRecordCreateRecordShape})
}

// CreateExternalDatabaseSearch sends POST /api/{api_version}/external_databases/{db_name}/search.
func (e *ExternalDatabases) CreateExternalDatabaseSearch(ctx context.Context, dbName any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(dbName, "db_name"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return e.requester.SendPostRequest(ctx, ExternalDatabaseSearchPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"db_name": dbName}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ExternalDatabaseSearchCreateRecordShape})
}

// UpdateExternalDatabaseUpdateRecords sends PATCH /api/{api_version}/external_databases/{db_name}/update_records.
func (e *ExternalDatabases) UpdateExternalDatabaseUpdateRecords(ctx context.Context, dbName any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(dbName, "db_name"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return e.requester.SendPatchRequest(ctx, ExternalDatabaseUpdateRecordsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"db_name": dbName}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ExternalDatabaseUpdateRecordsUpdateRecordShape})
}

// GetExternalDatabaseTable sends GET /api/{api_version}/external_databases/{db_name}/tables/{table_name}.
func (e *ExternalDatabases) GetExternalDatabaseTable(ctx context.Context, dbName any, tableName any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(dbName, "db_name"); err != nil {
		return nil, err
	}
	if err := requireValue(tableName, "table_name"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return e.requester.SendGetRequest(ctx, ExternalDatabaseTablePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"db_name": dbName, "table_name": tableName},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ExternalDatabaseTableListShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "external_databases"

// oneResponseFormat reads the optional trailing responseFormat, refusing more than one.
func oneResponseFormat(responseFormat []string) (string, error) {
	switch len(responseFormat) {
	case 0:
		return "", nil
	case 1:
		return responseFormat[0], nil
	default:
		return "", exerrors.NewUsageError(fmt.Sprintf(
			"responseFormat: pass one format or none, got %d; the trailing argument is optional, "+
				"and omitting it defers to the client", len(responseFormat)))
	}
}

// requirePayload refuses a nil body before any I/O; the client's validator would let one through bodyless.
func requirePayload(payload any) error {
	if payload == nil || isNilValue(payload) {
		return exerrors.NewUsageError("payload is required: this endpoint takes a body, and nil would " +
			"send the request with no body at all")
	}
	return nil
}

// isNilValue reports whether a value is a typed nil inside a non-nil any; a zero-valued struct is not.
func isNilValue(value any) bool {
	reflected := reflect.ValueOf(value)
	switch reflected.Kind() {
	case reflect.Map, reflect.Slice, reflect.Pointer, reflect.Interface, reflect.Func:
		return reflected.IsNil()
	}
	return false
}

// requireValue refuses a path value that is nil, empty or only whitespace, before any I/O.
func requireValue(value any, what string) error {
	text, isText := value.(string)
	if value == nil || isNilValue(value) || isText && strings.TrimSpace(text) == "" {
		return exerrors.UsageErrorf("%s is required: the request cannot be built without it", what)
	}
	return nil
}

// New binds the external databases endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *ExternalDatabases {
	return &ExternalDatabases{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (e *ExternalDatabases) Requester() exhttp.Requester { return e.requester }
