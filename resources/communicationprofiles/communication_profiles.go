package communicationprofiles

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// CommunicationProfiles is the v4 communication profiles endpoints, bound to one client.
type CommunicationProfiles struct {
	requester exhttp.Requester
}

// GetSettingsCommunicationProfiles sends GET /api/{api_version}/settings/communication_profiles.
func (c *CommunicationProfiles) GetSettingsCommunicationProfiles(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, SettingsCommunicationProfilesPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsCommunicationProfileListShape})
}

// CreateSettingsCommunicationProfile sends POST /api/{api_version}/settings/communication_profiles; it has no typed view, so read ApiResponse.Map.
func (c *CommunicationProfiles) CreateSettingsCommunicationProfile(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, SettingsCommunicationProfilesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetSettingsCommunicationProfile sends GET /api/{api_version}/settings/communication_profiles/{communication_profile_id}.
func (c *CommunicationProfiles) GetSettingsCommunicationProfile(ctx context.Context, communicationProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(communicationProfileID, "communication_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, SettingsCommunicationProfilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"communication_profile_id": communicationProfileID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsCommunicationProfileRecordShape})
}

// DeleteSettingsCommunicationProfile sends DELETE /api/{api_version}/settings/communication_profiles/{communication_profile_id}; it has no typed view, so read ApiResponse.Map.
func (c *CommunicationProfiles) DeleteSettingsCommunicationProfile(ctx context.Context, communicationProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(communicationProfileID, "communication_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendDeleteRequest(ctx, SettingsCommunicationProfilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"communication_profile_id": communicationProfileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsCommunicationProfile sends PATCH /api/{api_version}/settings/communication_profiles/{communication_profile_id}; it has no typed view, so read ApiResponse.Map.
func (c *CommunicationProfiles) UpdateSettingsCommunicationProfile(ctx context.Context, communicationProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(communicationProfileID, "communication_profile_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPatchRequest(ctx, SettingsCommunicationProfilePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"communication_profile_id": communicationProfileID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "communication_profiles"

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

// New binds the communication profiles endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *CommunicationProfiles {
	return &CommunicationProfiles{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (c *CommunicationProfiles) Requester() exhttp.Requester { return c.requester }
