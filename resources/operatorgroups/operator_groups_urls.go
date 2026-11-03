// Package operatorgroups is the v4 operator group endpoints.
package operatorgroups

// The path templates for the operator group endpoints; the client fills {api_version}.
const (
	SettingsOperatorGroupsPath = "/api/{api_version}/settings/operator_groups"
	SettingsOperatorGroupPath  = "/api/{api_version}/settings/operator_groups/{operator_group_id}"
)

// SettingsOperatorGroupsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsOperatorGroupsKey = "data"

// OperatorGroupKey is the envelope key one record arrives under.
const OperatorGroupKey = "operator_group"
