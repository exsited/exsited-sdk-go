package usages

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Usages is the v4 usages endpoints, bound to one client.
type Usages struct {
	requester exhttp.Requester
}

// GetUsages sends GET /api/{api_version}/usages.
func (u *Usages) GetUsages(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return u.requester.SendGetRequest(ctx, UsagesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: UsageListShape})
}

// CreateUsage sends POST /api/{api_version}/usages.
func (u *Usages) CreateUsage(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return u.requester.SendPostRequest(ctx, UsagesPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: UsagesRecordShape})
}

// GetUsage sends GET /api/{api_version}/usages/{usage_id}.
func (u *Usages) GetUsage(ctx context.Context, usageID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(usageID, "usage_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return u.requester.SendGetRequest(ctx, UsagePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"usage_id": usageID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: UsageRecordShape})
}

// DeleteUsage sends DELETE /api/{api_version}/usages/{usage_id}; it has no typed view, so read ApiResponse.Map.
func (u *Usages) DeleteUsage(ctx context.Context, usageID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(usageID, "usage_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return u.requester.SendDeleteRequest(ctx, UsagePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"usage_id": usageID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// UpdateUsage sends PATCH /api/{api_version}/usages/{usage_id}.
func (u *Usages) UpdateUsage(ctx context.Context, usageID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(usageID, "usage_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return u.requester.SendPatchRequest(ctx, UsagePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"usage_id": usageID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: UsageUpdatedRecordShape})
}

// ReplaceUsage sends PUT /api/{api_version}/usages/{usage_id}.
func (u *Usages) ReplaceUsage(ctx context.Context, usageID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(usageID, "usage_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return u.requester.SendPutRequest(ctx, UsagePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"usage_id": usageID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: UsageUpdatedRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "usages"

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

// New binds the usages endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Usages { return &Usages{requester: requester} }

// Requester returns the requester this resource was bound to.
func (u *Usages) Requester() exhttp.Requester { return u.requester }
