// Package communicationprofiles is the v4 communication profiles endpoints.
package communicationprofiles

// The path templates for the communication profiles endpoints; the client fills {api_version}.
const (
	SettingsCommunicationProfilesPath = "/api/{api_version}/settings/communication_profiles"
	SettingsCommunicationProfilePath  = "/api/{api_version}/settings/communication_profiles/{communication_profile_id}"
)

// SettingsCommunicationProfilesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsCommunicationProfilesKey = "data"

// CommunicationProfileKey is the envelope key one record arrives under.
const CommunicationProfileKey = "communication_profile"
