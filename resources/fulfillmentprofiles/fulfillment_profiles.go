package fulfillmentprofiles

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// FulfillmentProfiles is the v4 fulfillment profiles endpoints, bound to one client.
type FulfillmentProfiles struct {
	requester exhttp.Requester
}

// GetSettingsFulfillmentProfiles sends GET /api/{api_version}/settings/fulfillment_profiles.
func (f *FulfillmentProfiles) GetSettingsFulfillmentProfiles(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return f.requester.SendGetRequest(ctx, SettingsFulfillmentProfilesPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsFulfillmentProfileListShape})
}

// CreateSettingsFulfillmentProfile sends POST /api/{api_version}/settings/fulfillment_profiles.
func (f *FulfillmentProfiles) CreateSettingsFulfillmentProfile(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return f.requester.SendPostRequest(ctx, SettingsFulfillmentProfilesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsFulfillmentProfileRecordShape})
}

// GetSettingsFulfillmentProfile sends GET /api/{api_version}/settings/fulfillment_profiles/{fulfillment_profile_id}.
func (f *FulfillmentProfiles) GetSettingsFulfillmentProfile(ctx context.Context, fulfillmentProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fulfillmentProfileID, "fulfillment_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return f.requester.SendGetRequest(ctx, SettingsFulfillmentProfilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"fulfillment_profile_id": fulfillmentProfileID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsFulfillmentProfileRecordShape})
}

// DeleteSettingsFulfillmentProfile sends DELETE /api/{api_version}/settings/fulfillment_profiles/{fulfillment_profile_id}; it has no typed view, so read ApiResponse.Map.
func (f *FulfillmentProfiles) DeleteSettingsFulfillmentProfile(ctx context.Context, fulfillmentProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fulfillmentProfileID, "fulfillment_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return f.requester.SendDeleteRequest(ctx, SettingsFulfillmentProfilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"fulfillment_profile_id": fulfillmentProfileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsFulfillmentProfile sends PATCH /api/{api_version}/settings/fulfillment_profiles/{fulfillment_profile_id}.
func (f *FulfillmentProfiles) UpdateSettingsFulfillmentProfile(ctx context.Context, fulfillmentProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fulfillmentProfileID, "fulfillment_profile_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return f.requester.SendPatchRequest(ctx, SettingsFulfillmentProfilePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"fulfillment_profile_id": fulfillmentProfileID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsFulfillmentProfileRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "fulfillment_profiles"

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

// New binds the fulfillment profiles endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *FulfillmentProfiles {
	return &FulfillmentProfiles{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (f *FulfillmentProfiles) Requester() exhttp.Requester { return f.requester }
