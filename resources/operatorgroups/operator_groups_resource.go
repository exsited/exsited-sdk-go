package operatorgroups

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the OperatorGroups method of the same name,
// which is the primary surface.

// GetSettingsOperatorGroups is the module form of OperatorGroups.GetSettingsOperatorGroups.
func GetSettingsOperatorGroups(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsOperatorGroups(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsOperatorGroup is the module form of OperatorGroups.CreateSettingsOperatorGroup.
func CreateSettingsOperatorGroup(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsOperatorGroup(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsOperatorGroup is the module form of OperatorGroups.GetSettingsOperatorGroup.
func GetSettingsOperatorGroup(requester exhttp.Requester, ctx context.Context, operatorGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsOperatorGroup(ctx, operatorGroupID, raiseOnError, responseFormat...)
}

// DeleteSettingsOperatorGroup is the module form of OperatorGroups.DeleteSettingsOperatorGroup.
func DeleteSettingsOperatorGroup(requester exhttp.Requester, ctx context.Context, operatorGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsOperatorGroup(ctx, operatorGroupID, raiseOnError, responseFormat...)
}

// UpdateSettingsOperatorGroup is the module form of OperatorGroups.UpdateSettingsOperatorGroup.
func UpdateSettingsOperatorGroup(requester exhttp.Requester, ctx context.Context, operatorGroupID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsOperatorGroup(ctx, operatorGroupID, payload, raiseOnError, responseFormat...)
}
