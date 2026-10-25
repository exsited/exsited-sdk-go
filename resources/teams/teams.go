package teams

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Teams is the v4 teams endpoints, bound to one client.
type Teams struct {
	requester exhttp.Requester
}

// GetSettingsTeams sends GET /api/{api_version}/settings/teams.
func (t *Teams) GetSettingsTeams(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendGetRequest(ctx, SettingsTeamsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsTeamListShape})
}

// CreateSettingsTeam sends POST /api/{api_version}/settings/teams.
func (t *Teams) CreateSettingsTeam(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendPostRequest(ctx, SettingsTeamsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsTeamRecordShape})
}

// GetSettingsTeam sends GET /api/{api_version}/settings/teams/{team_id}.
func (t *Teams) GetSettingsTeam(ctx context.Context, teamID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(teamID, "team_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendGetRequest(ctx, SettingsTeamPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"team_id": teamID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsTeamRecordShape})
}

// DeleteSettingsTeam sends DELETE /api/{api_version}/settings/teams/{team_id}; it has no typed view, so read ApiResponse.Map.
func (t *Teams) DeleteSettingsTeam(ctx context.Context, teamID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(teamID, "team_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendDeleteRequest(ctx, SettingsTeamPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"team_id": teamID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsTeam sends PATCH /api/{api_version}/settings/teams/{team_id}.
func (t *Teams) UpdateSettingsTeam(ctx context.Context, teamID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(teamID, "team_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendPatchRequest(ctx, SettingsTeamPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"team_id": teamID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsTeamRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "teams"

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

// New binds the teams endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Teams { return &Teams{requester: requester} }

// Requester returns the requester this resource was bound to.
func (t *Teams) Requester() exhttp.Requester { return t.requester }
