// Package actionscripts is the v4 action scripts endpoints.
package actionscripts

// The path templates for the action scripts endpoints; the client fills {api_version}.
const (
	ActionScriptsPath = "/api/{api_version}/action_scripts"
	ActionScriptPath  = "/api/{api_version}/action_scripts/{action_script_id}"
)

// ActionScriptsKey is the key a listing's records arrive under, the v4 list envelope's data.
const ActionScriptsKey = "data"

// ActionScriptKey is the envelope key one record arrives under.
const ActionScriptKey = "action_script"
