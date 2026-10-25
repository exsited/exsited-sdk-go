package receiveterms

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// ReceiveTerms is the v4 receive terms endpoints, bound to one client.
type ReceiveTerms struct {
	requester exhttp.Requester
}

// GetSettingsReceiveTerms sends GET /api/{api_version}/settings/receive_terms.
func (r *ReceiveTerms) GetSettingsReceiveTerms(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, SettingsReceiveTermsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsReceiveTermListShape})
}

// CreateSettingsReceiveTerm sends POST /api/{api_version}/settings/receive_terms.
func (r *ReceiveTerms) CreateSettingsReceiveTerm(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, SettingsReceiveTermsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsReceiveTermRecordShape})
}

// GetSettingsReceiveTerm sends GET /api/{api_version}/settings/receive_terms/{receive_term_id}.
func (r *ReceiveTerms) GetSettingsReceiveTerm(ctx context.Context, receiveTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(receiveTermID, "receive_term_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, SettingsReceiveTermPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"receive_term_id": receiveTermID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsReceiveTermRecordShape})
}

// DeleteSettingsReceiveTerm sends DELETE /api/{api_version}/settings/receive_terms/{receive_term_id}; it has no typed view, so read ApiResponse.Map.
func (r *ReceiveTerms) DeleteSettingsReceiveTerm(ctx context.Context, receiveTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(receiveTermID, "receive_term_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendDeleteRequest(ctx, SettingsReceiveTermPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"receive_term_id": receiveTermID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsReceiveTerm sends PATCH /api/{api_version}/settings/receive_terms/{receive_term_id}.
func (r *ReceiveTerms) UpdateSettingsReceiveTerm(ctx context.Context, receiveTermID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(receiveTermID, "receive_term_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPatchRequest(ctx, SettingsReceiveTermPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"receive_term_id": receiveTermID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsReceiveTermRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "receive_terms"

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

// New binds the receive terms endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *ReceiveTerms { return &ReceiveTerms{requester: requester} }

// Requester returns the requester this resource was bound to.
func (r *ReceiveTerms) Requester() exhttp.Requester { return r.requester }
