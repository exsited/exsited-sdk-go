package portals

import (
	"context"
	"fmt"
	"reflect"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Portals is the v4 portals endpoints, bound to one client.
type Portals struct {
	requester exhttp.Requester
}

// ActivatePortal sends POST /api/{api_version}/portal/activate; it has no typed view, so read ApiResponse.Map.
func (p *Portals) ActivatePortal(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PortalActivatePath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreatePortalChangePassword sends POST /api/{api_version}/portal/change_password.
func (p *Portals) CreatePortalChangePassword(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PortalChangePasswordPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PortalChangePasswordCreateRecordShape})
}

// CreatePortalLogin sends POST /api/{api_version}/portal/login.
func (p *Portals) CreatePortalLogin(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PortalLoginPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PortalLoginRecordShape})
}

// CreatePortalReset sends POST /api/{api_version}/portal/reset; it has no typed view, so read ApiResponse.Map.
func (p *Portals) CreatePortalReset(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PortalResetPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreatePortalResetLink sends POST /api/{api_version}/portal/reset_link; it has no typed view, so read ApiResponse.Map.
func (p *Portals) CreatePortalResetLink(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PortalResetLinkPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreatePortalValidate sends POST /api/{api_version}/portal/validate; it has no typed view, so read ApiResponse.Map.
func (p *Portals) CreatePortalValidate(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PortalValidatePath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format})
}

// SetPortalPassword sends POST /api/{api_version}/portal/change; it has no typed view, so read ApiResponse.Map.
func (p *Portals) SetPortalPassword(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PortalChangePath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "portals"

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

// requirePayload refuses a nil body before any I/O; the client's validator would let one through bodyless.
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

// New binds the portals endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Portals { return &Portals{requester: requester} }

// Requester returns the requester this resource was bound to.
func (p *Portals) Requester() exhttp.Requester { return p.requester }
