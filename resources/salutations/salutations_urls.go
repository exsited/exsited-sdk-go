// Package salutations is the v4 salutations endpoints.
package salutations

// The path templates for the salutations endpoints; the client fills {api_version}.
const (
	SettingsSalutationsPath = "/api/{api_version}/settings/salutations"
	SettingsSalutationPath  = "/api/{api_version}/settings/salutations/{salutation_id}"
)

// SettingsSalutationsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsSalutationsKey = "data"

// SalutationKey is the envelope key one record arrives under.
const SalutationKey = "salutation"
