package salesreps

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the SalesReps method of the same name,
// which is the primary surface.

// GetSettingsSalesReps is the module form of SalesReps.GetSettingsSalesReps.
func GetSettingsSalesReps(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsSalesReps(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsSalesRep is the module form of SalesReps.CreateSettingsSalesRep.
func CreateSettingsSalesRep(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsSalesRep(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsSalesRep is the module form of SalesReps.GetSettingsSalesRep.
func GetSettingsSalesRep(requester exhttp.Requester, ctx context.Context, salesRepID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsSalesRep(ctx, salesRepID, raiseOnError, responseFormat...)
}

// DeleteSettingsSalesRep is the module form of SalesReps.DeleteSettingsSalesRep.
func DeleteSettingsSalesRep(requester exhttp.Requester, ctx context.Context, salesRepID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsSalesRep(ctx, salesRepID, raiseOnError, responseFormat...)
}

// UpdateSettingsSalesRep is the module form of SalesReps.UpdateSettingsSalesRep.
func UpdateSettingsSalesRep(requester exhttp.Requester, ctx context.Context, salesRepID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsSalesRep(ctx, salesRepID, payload, raiseOnError, responseFormat...)
}
