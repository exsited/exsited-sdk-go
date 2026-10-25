// Package designations is the v4 designations endpoints.
package designations

// The path templates for the designations endpoints; the client fills {api_version}.
const (
	SettingsDesignationsPath = "/api/{api_version}/settings/designations"
	SettingsDesignationPath  = "/api/{api_version}/settings/designations/{designation_id}"
)

// SettingsDesignationsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsDesignationsKey = "data"

// DesignationKey is the envelope key one record arrives under.
const DesignationKey = "designation"
