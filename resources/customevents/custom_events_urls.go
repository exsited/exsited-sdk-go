// Package customevents is the v4 custom events endpoints.
package customevents

// The path templates for the custom events endpoints; the client fills {api_version}.
const (
	SettingsCustomEventsPath = "/api/{api_version}/settings/custom_events"
	SettingsCustomEventPath  = "/api/{api_version}/settings/custom_events/{custom_event_id}"
)

// SettingsCustomEventsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsCustomEventsKey = "data"

// CustomEventKey is the envelope key one record arrives under.
const CustomEventKey = "custom_event"
