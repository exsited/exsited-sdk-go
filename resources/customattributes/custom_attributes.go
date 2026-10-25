package customattributes

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// CustomAttributes is the v4 custom attributes endpoints, bound to one client.
type CustomAttributes struct {
	requester exhttp.Requester
}

// GetCustomAttributes sends GET /api/{api_version}/custom_attributes.
func (c *CustomAttributes) GetCustomAttributes(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CustomAttributesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: CustomAttributeListShape})
}

// CreateCustomAttribute sends POST /api/{api_version}/custom_attributes.
func (c *CustomAttributes) CreateCustomAttribute(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, CustomAttributesPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: CustomAttributesRecordShape})
}

// DeleteCustomAttribute sends DELETE /api/{api_version}/custom_attributes/{custom_attribute_id}; it has no typed view, so read ApiResponse.Map.
func (c *CustomAttributes) DeleteCustomAttribute(ctx context.Context, customAttributeID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customAttributeID, "custom_attribute_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendDeleteRequest(ctx, CustomAttributePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"custom_attribute_id": customAttributeID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetCustomAttribute sends GET /api/{api_version}/custom_attributes/{custom_attribute_id}.
func (c *CustomAttributes) GetCustomAttribute(ctx context.Context, customAttributeID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customAttributeID, "custom_attribute_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CustomAttributePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"custom_attribute_id": customAttributeID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: CustomAttributeRecordShape})
}

// UpdateCustomAttribute sends PATCH /api/{api_version}/custom_attributes/{custom_attribute_id}.
func (c *CustomAttributes) UpdateCustomAttribute(ctx context.Context, customAttributeID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customAttributeID, "custom_attribute_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPatchRequest(ctx, CustomAttributePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"custom_attribute_id": customAttributeID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: CustomAttributeRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "custom_attributes"

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

// New binds the custom attributes endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *CustomAttributes {
	return &CustomAttributes{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (c *CustomAttributes) Requester() exhttp.Requester { return c.requester }
