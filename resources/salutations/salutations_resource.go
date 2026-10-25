package salutations

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Salutations method of the same name,
// which is the primary surface.

// GetSettingsSalutations is the module form of Salutations.GetSettingsSalutations.
func GetSettingsSalutations(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsSalutations(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsSalutation is the module form of Salutations.CreateSettingsSalutation.
func CreateSettingsSalutation(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsSalutation(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsSalutation is the module form of Salutations.GetSettingsSalutation.
func GetSettingsSalutation(requester exhttp.Requester, ctx context.Context, salutationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsSalutation(ctx, salutationID, raiseOnError, responseFormat...)
}

// DeleteSettingsSalutation is the module form of Salutations.DeleteSettingsSalutation.
func DeleteSettingsSalutation(requester exhttp.Requester, ctx context.Context, salutationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsSalutation(ctx, salutationID, raiseOnError, responseFormat...)
}

// UpdateSettingsSalutation is the module form of Salutations.UpdateSettingsSalutation.
func UpdateSettingsSalutation(requester exhttp.Requester, ctx context.Context, salutationID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsSalutation(ctx, salutationID, payload, raiseOnError, responseFormat...)
}
