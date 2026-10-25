// Package systemevents is the v4 system events endpoints.
package systemevents

// The path templates for the system events endpoints; the client fills {api_version}.
const (
	SettingsSystemEventsPath = "/api/{api_version}/settings/system_events"
	SettingsSystemEventPath  = "/api/{api_version}/settings/system_events/{system_event_id}"
)

// SettingsSystemEventsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsSystemEventsKey = "data"

// SystemEventKey is the envelope key one record arrives under.
const SystemEventKey = "system_event"
