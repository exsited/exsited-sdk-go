// Package teams is the v4 teams endpoints.
package teams

// The path templates for the teams endpoints; the client fills {api_version}.
const (
	SettingsTeamsPath = "/api/{api_version}/settings/teams"
	SettingsTeamPath  = "/api/{api_version}/settings/teams/{team_id}"
)

// SettingsTeamsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsTeamsKey = "data"

// TeamKey is the envelope key one record arrives under.
const TeamKey = "team"
