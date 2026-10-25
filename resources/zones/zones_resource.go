package zones

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Zones method of the same name,
// which is the primary surface.

// GetSettingsZones is the module form of Zones.GetSettingsZones.
func GetSettingsZones(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsZones(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsZone is the module form of Zones.CreateSettingsZone.
func CreateSettingsZone(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsZone(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsZone is the module form of Zones.GetSettingsZone.
func GetSettingsZone(requester exhttp.Requester, ctx context.Context, zoneID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsZone(ctx, zoneID, raiseOnError, responseFormat...)
}

// DeleteSettingsZone is the module form of Zones.DeleteSettingsZone.
func DeleteSettingsZone(requester exhttp.Requester, ctx context.Context, zoneID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsZone(ctx, zoneID, raiseOnError, responseFormat...)
}

// UpdateSettingsZone is the module form of Zones.UpdateSettingsZone.
func UpdateSettingsZone(requester exhttp.Requester, ctx context.Context, zoneID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsZone(ctx, zoneID, payload, raiseOnError, responseFormat...)
}
