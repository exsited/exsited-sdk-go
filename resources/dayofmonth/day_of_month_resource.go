package dayofmonth

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the DayOfMonth method of the same name,
// which is the primary surface.

// GetSettingsDaysOfMonth is the module form of DayOfMonth.GetSettingsDaysOfMonth.
func GetSettingsDaysOfMonth(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsDaysOfMonth(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsDayOfMonth is the module form of DayOfMonth.GetSettingsDayOfMonth.
func GetSettingsDayOfMonth(requester exhttp.Requester, ctx context.Context, dayOfMonthID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsDayOfMonth(ctx, dayOfMonthID, raiseOnError, responseFormat...)
}
