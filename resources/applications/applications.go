package applications

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Applications is the v4 applications endpoints, bound to one client.
type Applications struct {
	requester exhttp.Requester
}

// GetSettingsApplications sends GET /api/{api_version}/settings/applications.
func (a *Applications) GetSettingsApplications(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, SettingsApplicationsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsApplicationListShape})
}

// CreateSettingsApplication sends POST /api/{api_version}/settings/applications.
func (a *Applications) CreateSettingsApplication(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, SettingsApplicationsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsApplicationRecordShape})
}

// GetSettingsApplication sends GET /api/{api_version}/settings/applications/{application_id}.
func (a *Applications) GetSettingsApplication(ctx context.Context, applicationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(applicationID, "application_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, SettingsApplicationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"application_id": applicationID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsApplicationRecordShape})
}

// DeleteSettingsApplication sends DELETE /api/{api_version}/settings/applications/{application_id}; it has no typed view, so read ApiResponse.Map.
func (a *Applications) DeleteSettingsApplication(ctx context.Context, applicationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(applicationID, "application_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendDeleteRequest(ctx, SettingsApplicationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"application_id": applicationID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsApplication sends PATCH /api/{api_version}/settings/applications/{application_id}.
func (a *Applications) UpdateSettingsApplication(ctx context.Context, applicationID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(applicationID, "application_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPatchRequest(ctx, SettingsApplicationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"application_id": applicationID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsApplicationRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "applications"

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

// New binds the applications endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Applications { return &Applications{requester: requester} }

// Requester returns the requester this resource was bound to.
func (a *Applications) Requester() exhttp.Requester { return a.requester }
