// Package pricingperiods is the v4 pricing periods endpoints.
package pricingperiods

// The path templates for the pricing periods endpoints; the client fills {api_version}.
const (
	SettingsPricingPeriodsPath = "/api/{api_version}/settings/pricing_periods"
	SettingsPricingPeriodPath  = "/api/{api_version}/settings/pricing_periods/{pricing_period_id}"
)

// SettingsPricingPeriodsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsPricingPeriodsKey = "data"

// PricingPeriodKey is the envelope key one record arrives under.
const PricingPeriodKey = "pricing_period"
