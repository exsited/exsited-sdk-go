// Package roles is the v4 roles endpoints.
package roles

// The path templates for the roles endpoints; the client fills {api_version}.
const (
	SettingsRolesPath = "/api/{api_version}/settings/roles"
	SettingsRolePath  = "/api/{api_version}/settings/roles/{role_id}"
)

// SettingsRolesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsRolesKey = "data"

// RoleKey is the envelope key one record arrives under.
const RoleKey = "role"
