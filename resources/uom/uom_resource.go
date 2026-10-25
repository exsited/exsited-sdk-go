package uom

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Uom method of the same name,
// which is the primary surface.

// GetSettingsUoms is the module form of Uom.GetSettingsUoms.
func GetSettingsUoms(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsUoms(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsUom is the module form of Uom.CreateSettingsUom.
func CreateSettingsUom(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsUom(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsUom is the module form of Uom.GetSettingsUom.
func GetSettingsUom(requester exhttp.Requester, ctx context.Context, uomID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsUom(ctx, uomID, raiseOnError, responseFormat...)
}

// DeleteSettingsUom is the module form of Uom.DeleteSettingsUom.
func DeleteSettingsUom(requester exhttp.Requester, ctx context.Context, uomID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsUom(ctx, uomID, raiseOnError, responseFormat...)
}

// UpdateSettingsUom is the module form of Uom.UpdateSettingsUom.
func UpdateSettingsUom(requester exhttp.Requester, ctx context.Context, uomID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsUom(ctx, uomID, payload, raiseOnError, responseFormat...)
}
