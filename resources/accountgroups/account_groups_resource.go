package accountgroups

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the AccountGroups method of the same name,
// which is the primary surface.

// GetSettingsAccountGroups is the module form of AccountGroups.GetSettingsAccountGroups.
func GetSettingsAccountGroups(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsAccountGroups(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsAccountGroup is the module form of AccountGroups.CreateSettingsAccountGroup.
func CreateSettingsAccountGroup(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsAccountGroup(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsAccountGroup is the module form of AccountGroups.GetSettingsAccountGroup.
func GetSettingsAccountGroup(requester exhttp.Requester, ctx context.Context, accountGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsAccountGroup(ctx, accountGroupID, raiseOnError, responseFormat...)
}

// DeleteSettingsAccountGroup is the module form of AccountGroups.DeleteSettingsAccountGroup.
func DeleteSettingsAccountGroup(requester exhttp.Requester, ctx context.Context, accountGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsAccountGroup(ctx, accountGroupID, raiseOnError, responseFormat...)
}

// UpdateSettingsAccountGroup is the module form of AccountGroups.UpdateSettingsAccountGroup.
func UpdateSettingsAccountGroup(requester exhttp.Requester, ctx context.Context, accountGroupID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsAccountGroup(ctx, accountGroupID, payload, raiseOnError, responseFormat...)
}
