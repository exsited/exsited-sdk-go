// Package alertresponsibilities is the v4 alert responsibilities endpoints.
package alertresponsibilities

// The path templates for the alert responsibilities endpoints; the client fills {api_version}.
const (
	SettingsAlertResponsibilitiesPath = "/api/{api_version}/settings/alert_responsibilities"
	SettingsAlertResponsibilityPath   = "/api/{api_version}/settings/alert_responsibilities/{alert_responsibility_id}"
)

// SettingsAlertResponsibilitiesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsAlertResponsibilitiesKey = "data"

// AlertResponsibilityKey is the envelope key one record arrives under.
const AlertResponsibilityKey = "alert_responsibility"
