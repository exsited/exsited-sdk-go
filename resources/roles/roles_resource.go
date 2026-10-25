package roles

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Roles method of the same name,
// which is the primary surface.

// GetSettingsRoles is the module form of Roles.GetSettingsRoles.
func GetSettingsRoles(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsRoles(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsRole is the module form of Roles.CreateSettingsRole.
func CreateSettingsRole(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsRole(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsRole is the module form of Roles.GetSettingsRole.
func GetSettingsRole(requester exhttp.Requester, ctx context.Context, roleID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsRole(ctx, roleID, raiseOnError, responseFormat...)
}

// DeleteSettingsRole is the module form of Roles.DeleteSettingsRole.
func DeleteSettingsRole(requester exhttp.Requester, ctx context.Context, roleID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsRole(ctx, roleID, raiseOnError, responseFormat...)
}

// UpdateSettingsRole is the module form of Roles.UpdateSettingsRole.
func UpdateSettingsRole(requester exhttp.Requester, ctx context.Context, roleID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsRole(ctx, roleID, payload, raiseOnError, responseFormat...)
}
