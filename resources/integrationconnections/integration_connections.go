package integrationconnections

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// IntegrationConnections is the v4 integration connections endpoints, bound to one client.
type IntegrationConnections struct {
	requester exhttp.Requester
}

// GetIntegrationConnections sends GET /api/{api_version}/integration_connections. Filter by status with
// filter_status, which FilterParam("status") builds: ACTIVE, INACTIVE or DISCONNECTED. v4 ignores a plain
// status parameter and answers 200 whatever its value.
func (i *IntegrationConnections) GetIntegrationConnections(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, IntegrationConnectionsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: IntegrationConnectionsRecordShape})
}

// GetIntegrationConnection sends GET /api/{api_version}/integration_connections/{integration_connection_id}; it has no typed view, so read ApiResponse.Map.
func (i *IntegrationConnections) GetIntegrationConnection(ctx context.Context, integrationConnectionID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(integrationConnectionID, "integration_connection_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, IntegrationConnectionPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"integration_connection_id": integrationConnectionID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "integration_connections"

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

// New binds the integration connections endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *IntegrationConnections {
	return &IntegrationConnections{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (i *IntegrationConnections) Requester() exhttp.Requester { return i.requester }
