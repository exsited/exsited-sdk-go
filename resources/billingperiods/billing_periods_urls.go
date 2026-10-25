// Package billingperiods is the v4 billing periods endpoints.
package billingperiods

// The path templates for the billing periods endpoints; the client fills {api_version}.
const (
	SettingsBillingPeriodsPath = "/api/{api_version}/settings/billing_periods"
	SettingsBillingPeriodPath  = "/api/{api_version}/settings/billing_periods/{billing_period_id}"
)

// SettingsBillingPeriodsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsBillingPeriodsKey = "data"

// BillingPeriodKey is the envelope key one record arrives under.
const BillingPeriodKey = "billing_period"
