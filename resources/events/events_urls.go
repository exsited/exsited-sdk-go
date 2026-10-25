// Package events is the v4 events endpoints.
package events

// The path templates for the events endpoints; the client fills {api_version}.
const (
	EventsPath       = "/api/{api_version}/events"
	EventPath        = "/api/{api_version}/events/{event_uuid}"
	EventExecutePath = "/api/{api_version}/events/{event_uuid}/execute"
	EventRemovePath  = "/api/{api_version}/events/{event_uuid}/remove"
)

// EventsKey is the key a listing's records arrive under, the v4 list envelope's data.
const EventsKey = "data"

// EventKey is the envelope key one record arrives under.
const EventKey = "event"
