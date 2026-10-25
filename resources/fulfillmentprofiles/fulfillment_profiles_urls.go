// Package fulfillmentprofiles is the v4 fulfillment profiles endpoints.
package fulfillmentprofiles

// The path templates for the fulfillment profiles endpoints; the client fills {api_version}.
const (
	SettingsFulfillmentProfilesPath = "/api/{api_version}/settings/fulfillment_profiles"
	SettingsFulfillmentProfilePath  = "/api/{api_version}/settings/fulfillment_profiles/{fulfillment_profile_id}"
)

// SettingsFulfillmentProfilesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsFulfillmentProfilesKey = "data"

// FulfillmentProfileKey is the envelope key one record arrives under.
const FulfillmentProfileKey = "fulfillment_profile"
