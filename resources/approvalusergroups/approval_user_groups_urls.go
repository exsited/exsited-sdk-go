// Package approvalusergroups is the v4 approval user groups endpoints.
package approvalusergroups

// The path templates for the approval user groups endpoints; the client fills {api_version}.
const (
	SettingsApprovalUserGroupsPath = "/api/{api_version}/settings/approval_user_groups"
	SettingsApprovalUserGroupPath  = "/api/{api_version}/settings/approval_user_groups/{approval_user_group_id}"
)

// SettingsApprovalUserGroupsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsApprovalUserGroupsKey = "data"

// ApprovalUserGroupKey is the envelope key one record arrives under.
const ApprovalUserGroupKey = "approval_user_group"
