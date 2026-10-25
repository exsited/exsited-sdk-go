// Package milestones is the v4 milestones endpoints.
package milestones

// The path templates for the milestones endpoints; the client fills {api_version}.
const (
	MilestonesPath               = "/api/{api_version}/milestones"
	MilestonesAccountByIDPath    = "/api/{api_version}/milestones/account/{account_id}"
	MilestonesLabourByIDPath     = "/api/{api_version}/milestones/labour/{labour_id}"
	MilestonesEmployeeByCodePath = "/api/{api_version}/milestones/employee/{employee_code}"
	MilestonePath                = "/api/{api_version}/milestones/{milestone_uuid}"
)

// MilestonesKey is the key a listing's records arrive under, the v4 list envelope's data.
const MilestonesKey = "data"
