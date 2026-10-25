package billingperiods

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the BillingPeriods method of the same name,
// which is the primary surface.

// GetSettingsBillingPeriods is the module form of BillingPeriods.GetSettingsBillingPeriods.
func GetSettingsBillingPeriods(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsBillingPeriods(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsBillingPeriod is the module form of BillingPeriods.CreateSettingsBillingPeriod.
func CreateSettingsBillingPeriod(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsBillingPeriod(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsBillingPeriod is the module form of BillingPeriods.GetSettingsBillingPeriod.
func GetSettingsBillingPeriod(requester exhttp.Requester, ctx context.Context, billingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsBillingPeriod(ctx, billingPeriodID, raiseOnError, responseFormat...)
}

// DeleteSettingsBillingPeriod is the module form of BillingPeriods.DeleteSettingsBillingPeriod.
func DeleteSettingsBillingPeriod(requester exhttp.Requester, ctx context.Context, billingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsBillingPeriod(ctx, billingPeriodID, raiseOnError, responseFormat...)
}

// UpdateSettingsBillingPeriod is the module form of BillingPeriods.UpdateSettingsBillingPeriod.
func UpdateSettingsBillingPeriod(requester exhttp.Requester, ctx context.Context, billingPeriodID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsBillingPeriod(ctx, billingPeriodID, payload, raiseOnError, responseFormat...)
}
