// Package accountgroups is the v4 account groups endpoints.
package accountgroups

// The path templates for the account groups endpoints; the client fills {api_version}.
const (
	SettingsAccountGroupsPath = "/api/{api_version}/settings/account_groups"
	SettingsAccountGroupPath  = "/api/{api_version}/settings/account_groups/{account_group_id}"
)

// SettingsAccountGroupsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsAccountGroupsKey = "data"

// AccountGroupKey is the envelope key one record arrives under.
const AccountGroupKey = "account_group"
