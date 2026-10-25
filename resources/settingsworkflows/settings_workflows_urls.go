// Package settingsworkflows is the v4 settings workflows endpoints.
package settingsworkflows

// The path templates for the settings workflows endpoints; the client fills {api_version}.
const (
	SettingsWorkflowsPath = "/api/{api_version}/settings/workflows"
	SettingsWorkflowPath  = "/api/{api_version}/settings/workflows/{workflow_id}"
)

// SettingsWorkflowsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsWorkflowsKey = "data"

// WorkflowKey is the envelope key one record arrives under.
const WorkflowKey = "workflow"
