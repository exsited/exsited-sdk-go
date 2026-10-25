// Package additionalcodes is the v4 additional codes endpoints.
package additionalcodes

// The path templates for the additional codes endpoints; the client fills {api_version}.
const (
	SettingsAdditionalCodesPath = "/api/{api_version}/settings/additional_codes"
	SettingsAdditionalCodePath  = "/api/{api_version}/settings/additional_codes/{additional_code_id}"
)

// SettingsAdditionalCodesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsAdditionalCodesKey = "data"

// AdditionalCodeKey is the envelope key one record arrives under.
const AdditionalCodeKey = "additional_code"
