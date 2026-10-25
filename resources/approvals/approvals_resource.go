package approvals

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Approvals method of the same name,
// which is the primary surface.

// GetApprovalsAccount is the module form of Approvals.GetApprovalsAccount.
func GetApprovalsAccount(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetApprovalsAccount(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetApprovalsSaleorder is the module form of Approvals.GetApprovalsSaleorder.
func GetApprovalsSaleorder(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetApprovalsSaleorder(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetApprovalsAccountByUUID is the module form of Approvals.GetApprovalsAccountByUUID.
func GetApprovalsAccountByUUID(requester exhttp.Requester, ctx context.Context, accountUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetApprovalsAccountByUUID(ctx, accountUUID, queryParams, raiseOnError, responseFormat...)
}

// ApproveApproval is the module form of Approvals.ApproveApproval.
func ApproveApproval(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ApproveApproval(ctx, payload, raiseOnError, responseFormat...)
}

// RejectApproval is the module form of Approvals.RejectApproval.
func RejectApproval(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).RejectApproval(ctx, payload, raiseOnError, responseFormat...)
}
