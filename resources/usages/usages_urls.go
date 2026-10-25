// Package usages is the v4 usages endpoints.
package usages

// The path templates for the usages endpoints; the client fills {api_version}.
const (
	UsagesPath = "/api/{api_version}/usages"
	UsagePath  = "/api/{api_version}/usages/{usage_id}"
)

// UsagesKey is the key a listing's records arrive under, the v4 list envelope's data.
const UsagesKey = "data"

// UsageKey is the envelope key one record arrives under.
const UsageKey = "usage"
