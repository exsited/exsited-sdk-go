package pricingperiods

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the PricingPeriods method of the same name,
// which is the primary surface.

// GetSettingsPricingPeriods is the module form of PricingPeriods.GetSettingsPricingPeriods.
func GetSettingsPricingPeriods(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsPricingPeriods(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsPricingPeriod is the module form of PricingPeriods.CreateSettingsPricingPeriod.
func CreateSettingsPricingPeriod(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsPricingPeriod(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsPricingPeriod is the module form of PricingPeriods.GetSettingsPricingPeriod.
func GetSettingsPricingPeriod(requester exhttp.Requester, ctx context.Context, pricingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsPricingPeriod(ctx, pricingPeriodID, raiseOnError, responseFormat...)
}

// DeleteSettingsPricingPeriod is the module form of PricingPeriods.DeleteSettingsPricingPeriod.
func DeleteSettingsPricingPeriod(requester exhttp.Requester, ctx context.Context, pricingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsPricingPeriod(ctx, pricingPeriodID, raiseOnError, responseFormat...)
}

// UpdateSettingsPricingPeriod is the module form of PricingPeriods.UpdateSettingsPricingPeriod.
func UpdateSettingsPricingPeriod(requester exhttp.Requester, ctx context.Context, pricingPeriodID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsPricingPeriod(ctx, pricingPeriodID, payload, raiseOnError, responseFormat...)
}
