package additionalcodes

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// AdditionalCodes is the v4 additional codes endpoints, bound to one client.
type AdditionalCodes struct {
	requester exhttp.Requester
}

// GetSettingsAdditionalCodes sends GET /api/{api_version}/settings/additional_codes.
func (a *AdditionalCodes) GetSettingsAdditionalCodes(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, SettingsAdditionalCodesPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsAdditionalCodeListShape})
}

// CreateSettingsAdditionalCode sends POST /api/{api_version}/settings/additional_codes.
func (a *AdditionalCodes) CreateSettingsAdditionalCode(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, SettingsAdditionalCodesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsAdditionalCodeRecordShape})
}

// GetSettingsAdditionalCode sends GET /api/{api_version}/settings/additional_codes/{additional_code_id}.
func (a *AdditionalCodes) GetSettingsAdditionalCode(ctx context.Context, additionalCodeID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(additionalCodeID, "additional_code_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, SettingsAdditionalCodePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"additional_code_id": additionalCodeID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsAdditionalCodeRecordShape})
}

// DeleteSettingsAdditionalCode sends DELETE /api/{api_version}/settings/additional_codes/{additional_code_id}; it has no typed view, so read ApiResponse.Map.
func (a *AdditionalCodes) DeleteSettingsAdditionalCode(ctx context.Context, additionalCodeID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(additionalCodeID, "additional_code_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendDeleteRequest(ctx, SettingsAdditionalCodePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"additional_code_id": additionalCodeID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsAdditionalCode sends PATCH /api/{api_version}/settings/additional_codes/{additional_code_id}; it has no typed view, so read ApiResponse.Map.
func (a *AdditionalCodes) UpdateSettingsAdditionalCode(ctx context.Context, additionalCodeID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(additionalCodeID, "additional_code_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPatchRequest(ctx, SettingsAdditionalCodePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"additional_code_id": additionalCodeID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "additional_codes"

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

// requirePayload refuses a nil payload, a typed nil included, before any I/O.
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

// New binds the additional codes endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *AdditionalCodes { return &AdditionalCodes{requester: requester} }

// Requester returns the requester this resource was bound to.
func (a *AdditionalCodes) Requester() exhttp.Requester { return a.requester }
