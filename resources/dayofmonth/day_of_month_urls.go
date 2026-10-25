// Package dayofmonth is the v4 day of month endpoints.
package dayofmonth

// The path templates for the day of month endpoints; the client fills {api_version}.
const (
	SettingsDaysOfMonthPath = "/api/{api_version}/settings/day_of_month"
	SettingsDayOfMonthPath  = "/api/{api_version}/settings/day_of_month/{day_of_month_id}"
)

// SettingsDaysOfMonthKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsDaysOfMonthKey = "data"

// DayOfMonthKey is the envelope key one record arrives under.
const DayOfMonthKey = "day_of_month"
