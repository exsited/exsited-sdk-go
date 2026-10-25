// Package timeperiods is the v4 time periods endpoints.
package timeperiods

// The path templates for the time periods endpoints; the client fills {api_version}.
const (
	SettingsTimePeriodsPath = "/api/{api_version}/settings/time_periods"
	SettingsTimePeriodPath  = "/api/{api_version}/settings/time_periods/{time_period_id}"
)

// SettingsTimePeriodsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsTimePeriodsKey = "data"

// TimePeriodKey is the envelope key one record arrives under.
const TimePeriodKey = "time_period"
