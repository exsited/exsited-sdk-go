// Package timesheets is the v4 timesheets endpoints.
package timesheets

// The path templates for the timesheets endpoints; the client fills {api_version}.
const (
	TimesheetsPath = "/api/{api_version}/timesheets"
	TimesheetPath  = "/api/{api_version}/timesheets/{timesheet_id}"
)

// TimesheetsKey is the key a listing's records arrive under, the v4 list envelope's data.
const TimesheetsKey = "data"

// TimesheetKey is the envelope key one record arrives under.
const TimesheetKey = "timesheet"
