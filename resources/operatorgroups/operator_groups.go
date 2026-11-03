package operatorgroups

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// OperatorGroups is the v4 operator group endpoints, bound to one client.
type OperatorGroups struct {
	requester exhttp.Requester
}

// GetSettingsOperatorGroups sends GET /api/{api_version}/settings/operator_groups; paged by cursor.
func (o *OperatorGroups) GetSettingsOperatorGroups(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, SettingsOperatorGroupsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsOperatorGroupListShape})
}

// CreateSettingsOperatorGroup sends POST /api/{api_version}/settings/operator_groups, answered 201.
//
// role is required and takes the role's dashed uuid, not its rol_ id.
func (o *OperatorGroups) CreateSettingsOperatorGroup(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, SettingsOperatorGroupsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsOperatorGroupRecordShape})
}

// GetSettingsOperatorGroup sends GET /api/{api_version}/settings/operator_groups/{operator_group_id}.
func (o *OperatorGroups) GetSettingsOperatorGroup(ctx context.Context, operatorGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(operatorGroupID, "operator_group_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, SettingsOperatorGroupPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"operator_group_id": operatorGroupID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsOperatorGroupRecordShape})
}

// DeleteSettingsOperatorGroup sends DELETE /api/{api_version}/settings/operator_groups/{operator_group_id}, answered 204; it has no typed view, so read ApiResponse.Map.
func (o *OperatorGroups) DeleteSettingsOperatorGroup(ctx context.Context, operatorGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(operatorGroupID, "operator_group_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendDeleteRequest(ctx, SettingsOperatorGroupPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"operator_group_id": operatorGroupID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsOperatorGroup sends PATCH /api/{api_version}/settings/operator_groups/{operator_group_id};
// the answer names role as a reference carrying the dashed uuid.
func (o *OperatorGroups) UpdateSettingsOperatorGroup(ctx context.Context, operatorGroupID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(operatorGroupID, "operator_group_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPatchRequest(ctx, SettingsOperatorGroupPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"operator_group_id": operatorGroupID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsOperatorGroupRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "operator_groups"

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

// New binds the operator group endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *OperatorGroups { return &OperatorGroups{requester: requester} }

// Requester returns the requester this resource was bound to.
func (o *OperatorGroups) Requester() exhttp.Requester { return o.requester }
