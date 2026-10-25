package systemevents

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the SystemEvents method of the same name,
// which is the primary surface.

// GetSettingsSystemEvents is the module form of SystemEvents.GetSettingsSystemEvents.
func GetSettingsSystemEvents(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsSystemEvents(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsSystemEvent is the module form of SystemEvents.GetSettingsSystemEvent.
func GetSettingsSystemEvent(requester exhttp.Requester, ctx context.Context, systemEventID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsSystemEvent(ctx, systemEventID, raiseOnError, responseFormat...)
}
