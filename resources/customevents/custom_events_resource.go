package customevents

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the CustomEvents method of the same name,
// which is the primary surface.

// GetSettingsCustomEvents is the module form of CustomEvents.GetSettingsCustomEvents.
func GetSettingsCustomEvents(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCustomEvents(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsCustomEvent is the module form of CustomEvents.CreateSettingsCustomEvent.
func CreateSettingsCustomEvent(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsCustomEvent(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsCustomEvent is the module form of CustomEvents.GetSettingsCustomEvent.
func GetSettingsCustomEvent(requester exhttp.Requester, ctx context.Context, customEventID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCustomEvent(ctx, customEventID, raiseOnError, responseFormat...)
}

// DeleteSettingsCustomEvent is the module form of CustomEvents.DeleteSettingsCustomEvent.
func DeleteSettingsCustomEvent(requester exhttp.Requester, ctx context.Context, customEventID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsCustomEvent(ctx, customEventID, raiseOnError, responseFormat...)
}

// UpdateSettingsCustomEvent is the module form of CustomEvents.UpdateSettingsCustomEvent.
func UpdateSettingsCustomEvent(requester exhttp.Requester, ctx context.Context, customEventID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsCustomEvent(ctx, customEventID, payload, raiseOnError, responseFormat...)
}
