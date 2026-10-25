package organisations

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Organisations is the v4 organisations endpoints, bound to one client.
type Organisations struct {
	requester exhttp.Requester
}

// GetSettingsOrganisations sends GET /api/{api_version}/settings/organisations; it has no typed view, so read ApiResponse.Map.
func (o *Organisations) GetSettingsOrganisations(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, SettingsOrganisationsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateSettingsOrganisation sends POST /api/{api_version}/settings/organisations; it has no typed view, so read ApiResponse.Map.
func (o *Organisations) CreateSettingsOrganisation(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, SettingsOrganisationsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetSettingsOrganisation sends GET /api/{api_version}/settings/organisations/{organisation_id}; it has no typed view, so read ApiResponse.Map.
func (o *Organisations) GetSettingsOrganisation(ctx context.Context, organisationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(organisationID, "organisation_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, SettingsOrganisationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"organisation_id": organisationID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteSettingsOrganisation sends DELETE /api/{api_version}/settings/organisations/{organisation_id}; it has no typed view, so read ApiResponse.Map.
func (o *Organisations) DeleteSettingsOrganisation(ctx context.Context, organisationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(organisationID, "organisation_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendDeleteRequest(ctx, SettingsOrganisationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"organisation_id": organisationID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsOrganisation sends PATCH /api/{api_version}/settings/organisations/{organisation_id}; it has no typed view, so read ApiResponse.Map.
func (o *Organisations) UpdateSettingsOrganisation(ctx context.Context, organisationID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(organisationID, "organisation_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPatchRequest(ctx, SettingsOrganisationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"organisation_id": organisationID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "organisations"

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

// New binds the organisations endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Organisations { return &Organisations{requester: requester} }

// Requester returns the requester this resource was bound to.
func (o *Organisations) Requester() exhttp.Requester { return o.requester }
