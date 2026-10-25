// Package approvals is the v4 approvals endpoints.
package approvals

// The path templates for the approvals endpoints; the client fills {api_version}.
const (
	ApprovalsAccountPath       = "/api/{api_version}/approvals/account"
	ApprovalsSaleorderPath     = "/api/{api_version}/approvals/saleOrder"
	ApprovalsAccountByUUIDPath = "/api/{api_version}/approvals/account/{account_uuid}"
	ApprovalsApprovePath       = "/api/{api_version}/approvals/approve"
	ApprovalsRejectPath        = "/api/{api_version}/approvals/reject"
)

// ApprovalsKey is the key a listing's records arrive under, the v4 list envelope's data.
const ApprovalsKey = "data"
