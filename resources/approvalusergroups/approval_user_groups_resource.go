package approvalusergroups

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ApprovalUserGroups method of the same name,
// which is the primary surface.

// GetSettingsApprovalUserGroups is the module form of ApprovalUserGroups.GetSettingsApprovalUserGroups.
func GetSettingsApprovalUserGroups(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsApprovalUserGroups(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsApprovalUserGroup is the module form of ApprovalUserGroups.CreateSettingsApprovalUserGroup.
func CreateSettingsApprovalUserGroup(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsApprovalUserGroup(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsApprovalUserGroup is the module form of ApprovalUserGroups.GetSettingsApprovalUserGroup.
func GetSettingsApprovalUserGroup(requester exhttp.Requester, ctx context.Context, approvalUserGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsApprovalUserGroup(ctx, approvalUserGroupID, raiseOnError, responseFormat...)
}

// DeleteSettingsApprovalUserGroup is the module form of ApprovalUserGroups.DeleteSettingsApprovalUserGroup.
func DeleteSettingsApprovalUserGroup(requester exhttp.Requester, ctx context.Context, approvalUserGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsApprovalUserGroup(ctx, approvalUserGroupID, raiseOnError, responseFormat...)
}

// UpdateSettingsApprovalUserGroup is the module form of ApprovalUserGroups.UpdateSettingsApprovalUserGroup.
func UpdateSettingsApprovalUserGroup(requester exhttp.Requester, ctx context.Context, approvalUserGroupID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsApprovalUserGroup(ctx, approvalUserGroupID, payload, raiseOnError, responseFormat...)
}
