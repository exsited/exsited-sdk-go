package portals

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Portals method of the same name,
// which is the primary surface.

// ActivatePortal is the module form of Portals.ActivatePortal.
func ActivatePortal(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ActivatePortal(ctx, payload, raiseOnError, responseFormat...)
}

// CreatePortalChangePassword is the module form of Portals.CreatePortalChangePassword.
func CreatePortalChangePassword(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePortalChangePassword(ctx, payload, raiseOnError, responseFormat...)
}

// CreatePortalLogin is the module form of Portals.CreatePortalLogin.
func CreatePortalLogin(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePortalLogin(ctx, payload, raiseOnError, responseFormat...)
}

// CreatePortalReset is the module form of Portals.CreatePortalReset.
func CreatePortalReset(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePortalReset(ctx, payload, raiseOnError, responseFormat...)
}

// CreatePortalResetLink is the module form of Portals.CreatePortalResetLink.
func CreatePortalResetLink(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePortalResetLink(ctx, payload, raiseOnError, responseFormat...)
}

// CreatePortalValidate is the module form of Portals.CreatePortalValidate.
func CreatePortalValidate(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePortalValidate(ctx, payload, raiseOnError, responseFormat...)
}

// SetPortalPassword is the module form of Portals.SetPortalPassword.
func SetPortalPassword(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).SetPortalPassword(ctx, payload, raiseOnError, responseFormat...)
}
