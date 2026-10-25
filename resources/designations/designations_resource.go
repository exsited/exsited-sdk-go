package designations

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Designations method of the same name,
// which is the primary surface.

// GetSettingsDesignations is the module form of Designations.GetSettingsDesignations.
func GetSettingsDesignations(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsDesignations(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsDesignation is the module form of Designations.CreateSettingsDesignation.
func CreateSettingsDesignation(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsDesignation(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsDesignation is the module form of Designations.GetSettingsDesignation.
func GetSettingsDesignation(requester exhttp.Requester, ctx context.Context, designationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsDesignation(ctx, designationID, raiseOnError, responseFormat...)
}

// DeleteSettingsDesignation is the module form of Designations.DeleteSettingsDesignation.
func DeleteSettingsDesignation(requester exhttp.Requester, ctx context.Context, designationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsDesignation(ctx, designationID, raiseOnError, responseFormat...)
}

// UpdateSettingsDesignation is the module form of Designations.UpdateSettingsDesignation.
func UpdateSettingsDesignation(requester exhttp.Requester, ctx context.Context, designationID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsDesignation(ctx, designationID, payload, raiseOnError, responseFormat...)
}
