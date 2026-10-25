// Package applications is the v4 applications endpoints.
package applications

// The path templates for the applications endpoints; the client fills {api_version}.
const (
	SettingsApplicationsPath = "/api/{api_version}/settings/applications"
	SettingsApplicationPath  = "/api/{api_version}/settings/applications/{application_id}"
)

// SettingsApplicationsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsApplicationsKey = "data"

// ApplicationKey is the envelope key one record arrives under.
const ApplicationKey = "application"
