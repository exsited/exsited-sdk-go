// Package approvalprofiles is the v4 approval profiles endpoints.
package approvalprofiles

// The path templates for the approval profiles endpoints; the client fills {api_version}.
const (
	SettingsApprovalProfilesPath = "/api/{api_version}/settings/approval_profiles"
	SettingsApprovalProfilePath  = "/api/{api_version}/settings/approval_profiles/{approval_profile_id}"
)

// SettingsApprovalProfilesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsApprovalProfilesKey = "data"

// ApprovalProfileKey is the envelope key one record arrives under.
const ApprovalProfileKey = "approval_profile"
