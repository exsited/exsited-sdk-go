package additionalcodes

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the AdditionalCodes method of the same name,
// which is the primary surface.

// GetSettingsAdditionalCodes is the module form of AdditionalCodes.GetSettingsAdditionalCodes.
func GetSettingsAdditionalCodes(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsAdditionalCodes(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsAdditionalCode is the module form of AdditionalCodes.CreateSettingsAdditionalCode.
func CreateSettingsAdditionalCode(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsAdditionalCode(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsAdditionalCode is the module form of AdditionalCodes.GetSettingsAdditionalCode.
func GetSettingsAdditionalCode(requester exhttp.Requester, ctx context.Context, additionalCodeID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsAdditionalCode(ctx, additionalCodeID, raiseOnError, responseFormat...)
}

// DeleteSettingsAdditionalCode is the module form of AdditionalCodes.DeleteSettingsAdditionalCode.
func DeleteSettingsAdditionalCode(requester exhttp.Requester, ctx context.Context, additionalCodeID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsAdditionalCode(ctx, additionalCodeID, raiseOnError, responseFormat...)
}

// UpdateSettingsAdditionalCode is the module form of AdditionalCodes.UpdateSettingsAdditionalCode.
func UpdateSettingsAdditionalCode(requester exhttp.Requester, ctx context.Context, additionalCodeID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsAdditionalCode(ctx, additionalCodeID, payload, raiseOnError, responseFormat...)
}
