package approvalprofiles

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ApprovalProfiles method of the same name,
// which is the primary surface.

// GetSettingsApprovalProfiles is the module form of ApprovalProfiles.GetSettingsApprovalProfiles.
func GetSettingsApprovalProfiles(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsApprovalProfiles(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsApprovalProfile is the module form of ApprovalProfiles.CreateSettingsApprovalProfile.
func CreateSettingsApprovalProfile(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsApprovalProfile(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsApprovalProfile is the module form of ApprovalProfiles.GetSettingsApprovalProfile.
func GetSettingsApprovalProfile(requester exhttp.Requester, ctx context.Context, approvalProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsApprovalProfile(ctx, approvalProfileID, raiseOnError, responseFormat...)
}

// DeleteSettingsApprovalProfile is the module form of ApprovalProfiles.DeleteSettingsApprovalProfile.
func DeleteSettingsApprovalProfile(requester exhttp.Requester, ctx context.Context, approvalProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsApprovalProfile(ctx, approvalProfileID, raiseOnError, responseFormat...)
}

// UpdateSettingsApprovalProfile is the module form of ApprovalProfiles.UpdateSettingsApprovalProfile.
func UpdateSettingsApprovalProfile(requester exhttp.Requester, ctx context.Context, approvalProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsApprovalProfile(ctx, approvalProfileID, payload, raiseOnError, responseFormat...)
}
