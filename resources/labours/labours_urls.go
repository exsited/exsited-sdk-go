// Package labours is the v4 labours endpoints.
package labours

// The path templates for the labours endpoints; the client fills {api_version}.
const (
	LabourAvailabilityPath         = "/api/{api_version}/labour_availability"
	LabourProfilesPath             = "/api/{api_version}/labour_profiles"
	LaboursPath                    = "/api/{api_version}/labours"
	LabourProfilePath              = "/api/{api_version}/labour_profiles/{labour_profile_id}"
	LaboursAvailableForBookingPath = "/api/{api_version}/labours/available_for_booking"
	LabourPath                     = "/api/{api_version}/labours/{labour_uuid}"
	LabourCancelPath               = "/api/{api_version}/labours/{labour_id}/cancel"
	LabourReactivatePath           = "/api/{api_version}/labours/{labour_id}/reactivate"
	LabourOrdersPath               = "/api/{api_version}/labours/{labour_uuid}/orders"
	LabourProfileSessionsPath      = "/api/{api_version}/labour_profiles/{labour_profile_id}/sessions"
	LabourProfileSessionPath       = "/api/{api_version}/labour_profiles/{labour_profile_id}/sessions/{session_id}"
	LabourCalendarPath             = "/api/{api_version}/labours/{labour_id}/availability"
)

// LaboursKey is the key a listing's records arrive under, the v4 list envelope's data.
const LaboursKey = "data"

// LabourKey is the envelope key one record arrives under.
const LabourKey = "labour"

// LabourProfileKey is the envelope key one record arrives under.
const LabourProfileKey = "labour_profile"

// LabourProfileSessionKey is the envelope key one record arrives under.
const LabourProfileSessionKey = "labour_profile_session"

// LabourCalendarKey is the envelope key one record arrives under.
const LabourCalendarKey = "labour_calendar"
