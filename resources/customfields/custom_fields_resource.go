package customfields

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the CustomFields method of the same name,
// which is the primary surface.

// GetSettingsCustomFields is the module form of CustomFields.GetSettingsCustomFields.
func GetSettingsCustomFields(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCustomFields(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsCustomField is the module form of CustomFields.GetSettingsCustomField.
func GetSettingsCustomField(requester exhttp.Requester, ctx context.Context, customFieldID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCustomField(ctx, customFieldID, raiseOnError, responseFormat...)
}
