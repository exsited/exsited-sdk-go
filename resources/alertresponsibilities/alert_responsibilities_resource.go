package alertresponsibilities

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the AlertResponsibilities method of the same name,
// which is the primary surface.

// GetSettingsAlertResponsibilities is the module form of AlertResponsibilities.GetSettingsAlertResponsibilities.
func GetSettingsAlertResponsibilities(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsAlertResponsibilities(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsAlertResponsibility is the module form of AlertResponsibilities.CreateSettingsAlertResponsibility.
func CreateSettingsAlertResponsibility(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsAlertResponsibility(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsAlertResponsibility is the module form of AlertResponsibilities.GetSettingsAlertResponsibility.
func GetSettingsAlertResponsibility(requester exhttp.Requester, ctx context.Context, alertResponsibilityID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsAlertResponsibility(ctx, alertResponsibilityID, raiseOnError, responseFormat...)
}

// DeleteSettingsAlertResponsibility is the module form of AlertResponsibilities.DeleteSettingsAlertResponsibility.
func DeleteSettingsAlertResponsibility(requester exhttp.Requester, ctx context.Context, alertResponsibilityID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsAlertResponsibility(ctx, alertResponsibilityID, raiseOnError, responseFormat...)
}

// UpdateSettingsAlertResponsibility is the module form of AlertResponsibilities.UpdateSettingsAlertResponsibility.
func UpdateSettingsAlertResponsibility(requester exhttp.Requester, ctx context.Context, alertResponsibilityID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsAlertResponsibility(ctx, alertResponsibilityID, payload, raiseOnError, responseFormat...)
}
