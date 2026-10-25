// Package zones is the v4 zones endpoints.
package zones

// The path templates for the zones endpoints; the client fills {api_version}.
const (
	SettingsZonesPath = "/api/{api_version}/settings/zones"
	SettingsZonePath  = "/api/{api_version}/settings/zones/{zone_id}"
)

// SettingsZonesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsZonesKey = "data"

// ZoneKey is the envelope key one record arrives under.
const ZoneKey = "zone"
