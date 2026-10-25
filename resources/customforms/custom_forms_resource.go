package customforms

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the CustomForms method of the same name,
// which is the primary surface.

// GetSettingsCustomForms is the module form of CustomForms.GetSettingsCustomForms.
func GetSettingsCustomForms(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCustomForms(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsCustomForm is the module form of CustomForms.CreateSettingsCustomForm.
func CreateSettingsCustomForm(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsCustomForm(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsCustomForm is the module form of CustomForms.GetSettingsCustomForm.
func GetSettingsCustomForm(requester exhttp.Requester, ctx context.Context, customFormID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCustomForm(ctx, customFormID, raiseOnError, responseFormat...)
}

// DeleteSettingsCustomForm is the module form of CustomForms.DeleteSettingsCustomForm.
func DeleteSettingsCustomForm(requester exhttp.Requester, ctx context.Context, customFormID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsCustomForm(ctx, customFormID, raiseOnError, responseFormat...)
}

// UpdateSettingsCustomForm is the module form of CustomForms.UpdateSettingsCustomForm.
func UpdateSettingsCustomForm(requester exhttp.Requester, ctx context.Context, customFormID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsCustomForm(ctx, customFormID, payload, raiseOnError, responseFormat...)
}
