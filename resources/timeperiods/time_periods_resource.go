package timeperiods

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the TimePeriods method of the same name,
// which is the primary surface.

// GetSettingsTimePeriods is the module form of TimePeriods.GetSettingsTimePeriods.
func GetSettingsTimePeriods(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsTimePeriods(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsTimePeriod is the module form of TimePeriods.CreateSettingsTimePeriod.
func CreateSettingsTimePeriod(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsTimePeriod(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsTimePeriod is the module form of TimePeriods.GetSettingsTimePeriod.
func GetSettingsTimePeriod(requester exhttp.Requester, ctx context.Context, timePeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsTimePeriod(ctx, timePeriodID, raiseOnError, responseFormat...)
}

// DeleteSettingsTimePeriod is the module form of TimePeriods.DeleteSettingsTimePeriod.
func DeleteSettingsTimePeriod(requester exhttp.Requester, ctx context.Context, timePeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsTimePeriod(ctx, timePeriodID, raiseOnError, responseFormat...)
}

// UpdateSettingsTimePeriod is the module form of TimePeriods.UpdateSettingsTimePeriod.
func UpdateSettingsTimePeriod(requester exhttp.Requester, ctx context.Context, timePeriodID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsTimePeriod(ctx, timePeriodID, payload, raiseOnError, responseFormat...)
}
