package operators

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Operators is the v4 operator endpoints, bound to one client. An operator is a user of the
// instance, and this family has no delete.
type Operators struct {
	requester exhttp.Requester
}

// GetSettingsOperators sends GET /api/{api_version}/settings/operators; paged by cursor.
func (o *Operators) GetSettingsOperators(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, SettingsOperatorsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsOperatorListShape})
}

// CreateSettingsOperator sends POST /api/{api_version}/settings/operators, answered 201. role and
// operator_group take dashed uuids.
func (o *Operators) CreateSettingsOperator(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, SettingsOperatorsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsOperatorRecordShape})
}

// GetSettingsOperator sends GET /api/{api_version}/settings/operators/{operator_id}.
func (o *Operators) GetSettingsOperator(ctx context.Context, operatorID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(operatorID, "operator_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, SettingsOperatorPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"operator_id": operatorID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsOperatorRecordShape})
}

// UpdateSettingsOperator sends PATCH /api/{api_version}/settings/operators/{operator_id};
// description is refused as not understood.
func (o *Operators) UpdateSettingsOperator(ctx context.Context, operatorID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(operatorID, "operator_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPatchRequest(ctx, SettingsOperatorPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"operator_id": operatorID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsOperatorRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "operators"

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

// New binds the operator endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Operators { return &Operators{requester: requester} }

// Requester returns the requester this resource was bound to.
func (o *Operators) Requester() exhttp.Requester { return o.requester }
