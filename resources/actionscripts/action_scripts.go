package actionscripts

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// ActionScripts is the v4 action scripts endpoints, bound to one client.
type ActionScripts struct {
	requester exhttp.Requester
}

// GetActionScripts sends GET /api/{api_version}/action_scripts.
func (a *ActionScripts) GetActionScripts(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, ActionScriptsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ActionScriptListShape})
}

// GetActionScript sends GET /api/{api_version}/action_scripts/{action_script_id}.
func (a *ActionScripts) GetActionScript(ctx context.Context, actionScriptID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(actionScriptID, "action_script_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, ActionScriptPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"action_script_id": actionScriptID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ActionScriptRecordShape})
}

// CreateActionScript sends POST /api/{api_version}/action_scripts.
func (a *ActionScripts) CreateActionScript(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, ActionScriptsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ActionScriptRecordShape})
}

// UpdateActionScript sends PATCH /api/{api_version}/action_scripts/{action_script_id}.
func (a *ActionScripts) UpdateActionScript(ctx context.Context, actionScriptID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(actionScriptID, "action_script_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPatchRequest(ctx, ActionScriptPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"action_script_id": actionScriptID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ActionScriptRecordShape})
}

// DeleteActionScript sends DELETE /api/{api_version}/action_scripts/{action_script_id}; it has no typed view, so read ApiResponse.Map.
func (a *ActionScripts) DeleteActionScript(ctx context.Context, actionScriptID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(actionScriptID, "action_script_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendDeleteRequest(ctx, ActionScriptPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"action_script_id": actionScriptID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "action_scripts"

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

// New binds the action scripts endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *ActionScripts { return &ActionScripts{requester: requester} }

// Requester returns the requester this resource was bound to.
func (a *ActionScripts) Requester() exhttp.Requester { return a.requester }
