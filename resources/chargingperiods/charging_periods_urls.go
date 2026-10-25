// Package chargingperiods is the v4 charging periods endpoints.
package chargingperiods

// The path templates for the charging periods endpoints; the client fills {api_version}.
const (
	SettingsChargingPeriodsPath = "/api/{api_version}/settings/charging_periods"
	SettingsChargingPeriodPath  = "/api/{api_version}/settings/charging_periods/{charging_period_id}"
)

// SettingsChargingPeriodsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsChargingPeriodsKey = "data"

// ChargingPeriodKey is the envelope key one record arrives under.
const ChargingPeriodKey = "charging_period"
