// Package workflows is the v4 workflows endpoints.
package workflows

// The path templates for the workflows endpoints; the client fills {api_version}.
const (
	MilestonePath                              = "/api/{api_version}/milestones/{milestone_uuid}"
	WorkflowComponentByIDPath                  = "/api/{api_version}/workflow/component/{cc_id}"
	WorkflowMilestonesPath                     = "/api/{api_version}/workflows/{workflow_uuid}/milestones"
	WorkflowComponentCcMilestonePath           = "/api/{api_version}/workflow/component/{cc_id}/milestone"
	WorkflowComponentHistoryPath               = "/api/{api_version}/workflow/component/{component_object_id}/history"
	WorkflowComponentCcInstanceCcMilestonePath = "/api/{api_version}/workflow/component/{component_uuid}/{component_object_id}/milestone"
	WorkflowsPath                              = "/api/{api_version}/workflows"
	WorkflowPath                               = "/api/{api_version}/workflows/{workflow_id}"
	WorkflowStagesPath                         = "/api/{api_version}/workflows/{workflow_id}/stages"
	WorkflowStagePath                          = "/api/{api_version}/workflows/{workflow_id}/stages/{stage_id}"
	WorkflowMilestonePath                      = "/api/{api_version}/workflows/{workflow_id}/milestones/{milestone_id}"
	WorkflowCustomComponentHistoryPath         = "/api/{api_version}/workflows/custom_components/{component_uuid}/{object_code}/history"
	WorkflowCustomComponentMilestonesPath      = "/api/{api_version}/workflows/custom_components/{component_uuid}/{object_code}/milestones"
	WorkflowCustomComponentMilestonePath       = "/api/{api_version}/workflows/custom_components/{component_uuid}/{object_code}/milestones/{milestone_id}"
)

// WorkflowMilestoneKey is the envelope key one record arrives under.
const WorkflowMilestoneKey = "workflow_milestone"
