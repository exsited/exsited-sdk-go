package applications

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Applications method of the same name,
// which is the primary surface.

// GetSettingsApplications is the module form of Applications.GetSettingsApplications.
func GetSettingsApplications(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsApplications(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsApplication is the module form of Applications.CreateSettingsApplication.
func CreateSettingsApplication(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsApplication(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsApplication is the module form of Applications.GetSettingsApplication.
func GetSettingsApplication(requester exhttp.Requester, ctx context.Context, applicationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsApplication(ctx, applicationID, raiseOnError, responseFormat...)
}

// DeleteSettingsApplication is the module form of Applications.DeleteSettingsApplication.
func DeleteSettingsApplication(requester exhttp.Requester, ctx context.Context, applicationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsApplication(ctx, applicationID, raiseOnError, responseFormat...)
}

// UpdateSettingsApplication is the module form of Applications.UpdateSettingsApplication.
func UpdateSettingsApplication(requester exhttp.Requester, ctx context.Context, applicationID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsApplication(ctx, applicationID, payload, raiseOnError, responseFormat...)
}
