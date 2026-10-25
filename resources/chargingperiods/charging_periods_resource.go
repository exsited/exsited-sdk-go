package chargingperiods

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ChargingPeriods method of the same name,
// which is the primary surface.

// GetSettingsChargingPeriods is the module form of ChargingPeriods.GetSettingsChargingPeriods.
func GetSettingsChargingPeriods(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsChargingPeriods(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsChargingPeriod is the module form of ChargingPeriods.CreateSettingsChargingPeriod.
func CreateSettingsChargingPeriod(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsChargingPeriod(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsChargingPeriod is the module form of ChargingPeriods.GetSettingsChargingPeriod.
func GetSettingsChargingPeriod(requester exhttp.Requester, ctx context.Context, chargingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsChargingPeriod(ctx, chargingPeriodID, raiseOnError, responseFormat...)
}

// DeleteSettingsChargingPeriod is the module form of ChargingPeriods.DeleteSettingsChargingPeriod.
func DeleteSettingsChargingPeriod(requester exhttp.Requester, ctx context.Context, chargingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsChargingPeriod(ctx, chargingPeriodID, raiseOnError, responseFormat...)
}

// UpdateSettingsChargingPeriod is the module form of ChargingPeriods.UpdateSettingsChargingPeriod.
func UpdateSettingsChargingPeriod(requester exhttp.Requester, ctx context.Context, chargingPeriodID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsChargingPeriod(ctx, chargingPeriodID, payload, raiseOnError, responseFormat...)
}
