package communicationprofiles

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the CommunicationProfiles method of the same name,
// which is the primary surface.

// GetSettingsCommunicationProfiles is the module form of CommunicationProfiles.GetSettingsCommunicationProfiles.
func GetSettingsCommunicationProfiles(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCommunicationProfiles(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsCommunicationProfile is the module form of CommunicationProfiles.CreateSettingsCommunicationProfile.
func CreateSettingsCommunicationProfile(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsCommunicationProfile(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsCommunicationProfile is the module form of CommunicationProfiles.GetSettingsCommunicationProfile.
func GetSettingsCommunicationProfile(requester exhttp.Requester, ctx context.Context, communicationProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCommunicationProfile(ctx, communicationProfileID, raiseOnError, responseFormat...)
}

// DeleteSettingsCommunicationProfile is the module form of CommunicationProfiles.DeleteSettingsCommunicationProfile.
func DeleteSettingsCommunicationProfile(requester exhttp.Requester, ctx context.Context, communicationProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsCommunicationProfile(ctx, communicationProfileID, raiseOnError, responseFormat...)
}

// UpdateSettingsCommunicationProfile is the module form of CommunicationProfiles.UpdateSettingsCommunicationProfile.
func UpdateSettingsCommunicationProfile(requester exhttp.Requester, ctx context.Context, communicationProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsCommunicationProfile(ctx, communicationProfileID, payload, raiseOnError, responseFormat...)
}
