// Package alerts is the v4 alerts endpoints.
package alerts

// The path templates for the alerts endpoints; the client fills {api_version}.
const (
	AlertsPath             = "/api/{api_version}/alerts"
	AlertPath              = "/api/{api_version}/alerts/{alert_id}"
	AlertsAccountByIDPath  = "/api/{api_version}/alerts/account/{account_id}"
	AlertsEmployeeByIDPath = "/api/{api_version}/alerts/employee/{employee_id}"
	AlertsLabourByIDPath   = "/api/{api_version}/alerts/labour/{labour_id}"
)

// AlertsKey is the key a listing's records arrive under, the v4 list envelope's data.
const AlertsKey = "data"
