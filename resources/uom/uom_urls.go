// Package uom is the v4 uom endpoints.
package uom

// The path templates for the uom endpoints; the client fills {api_version}.
const (
	SettingsUomsPath = "/api/{api_version}/settings/uom"
	SettingsUomPath  = "/api/{api_version}/settings/uom/{uom_id}"
)

// SettingsUomsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsUomsKey = "data"

// UomKey is the envelope key one record arrives under.
const UomKey = "uom"
