// Package salesreps is the v4 sales reps endpoints.
package salesreps

// The path templates for the sales reps endpoints; the client fills {api_version}.
const (
	SettingsSalesRepsPath = "/api/{api_version}/settings/sales_reps"
	SettingsSalesRepPath  = "/api/{api_version}/settings/sales_reps/{sales_rep_id}"
)

// SettingsSalesRepsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsSalesRepsKey = "data"

// SalesRepKey is the envelope key one record arrives under.
const SalesRepKey = "sales_rep"
