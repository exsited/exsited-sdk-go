package fulfillmentprofiles

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the FulfillmentProfiles method of the same name,
// which is the primary surface.

// GetSettingsFulfillmentProfiles is the module form of FulfillmentProfiles.GetSettingsFulfillmentProfiles.
func GetSettingsFulfillmentProfiles(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsFulfillmentProfiles(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsFulfillmentProfile is the module form of FulfillmentProfiles.CreateSettingsFulfillmentProfile.
func CreateSettingsFulfillmentProfile(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsFulfillmentProfile(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsFulfillmentProfile is the module form of FulfillmentProfiles.GetSettingsFulfillmentProfile.
func GetSettingsFulfillmentProfile(requester exhttp.Requester, ctx context.Context, fulfillmentProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsFulfillmentProfile(ctx, fulfillmentProfileID, raiseOnError, responseFormat...)
}

// DeleteSettingsFulfillmentProfile is the module form of FulfillmentProfiles.DeleteSettingsFulfillmentProfile.
func DeleteSettingsFulfillmentProfile(requester exhttp.Requester, ctx context.Context, fulfillmentProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsFulfillmentProfile(ctx, fulfillmentProfileID, raiseOnError, responseFormat...)
}

// UpdateSettingsFulfillmentProfile is the module form of FulfillmentProfiles.UpdateSettingsFulfillmentProfile.
func UpdateSettingsFulfillmentProfile(requester exhttp.Requester, ctx context.Context, fulfillmentProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsFulfillmentProfile(ctx, fulfillmentProfileID, payload, raiseOnError, responseFormat...)
}
