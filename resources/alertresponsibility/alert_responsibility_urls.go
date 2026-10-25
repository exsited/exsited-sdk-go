// Package alertresponsibility is the v4 alert responsibility endpoints.
package alertresponsibility

// The path templates for the alert responsibility endpoints; the client fills {api_version}.
const (
	AlertResponsibilitiesPath = "/api/{api_version}/alert_responsibility"
	AlertResponsibilityPath   = "/api/{api_version}/alert_responsibility/{alert_responsibility_id}"
)

// AlertResponsibilitiesKey is the key a listing's records arrive under, the v4 list envelope's data.
const AlertResponsibilitiesKey = "data"

// AlertResponsibilityKey is the envelope key one record arrives under.
const AlertResponsibilityKey = "alert_responsibility"
