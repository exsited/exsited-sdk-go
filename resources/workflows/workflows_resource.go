package workflows

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Workflows method of the same name,
// which is the primary surface.

// DeleteMilestone is the module form of Workflows.DeleteMilestone.
func DeleteMilestone(requester exhttp.Requester, ctx context.Context, milestoneUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteMilestone(ctx, milestoneUUID, raiseOnError, responseFormat...)
}

// UpdateWorkflowComponentByID is the module form of Workflows.UpdateWorkflowComponentByID.
func UpdateWorkflowComponentByID(requester exhttp.Requester, ctx context.Context, ccID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateWorkflowComponentByID(ctx, ccID, payload, raiseOnError, responseFormat...)
}

// GetWorkflowMilestones is the module form of Workflows.GetWorkflowMilestones.
func GetWorkflowMilestones(requester exhttp.Requester, ctx context.Context, workflowUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWorkflowMilestones(ctx, workflowUUID, queryParams, raiseOnError, responseFormat...)
}

// CreateWorkflowComponentCcMilestone is the module form of Workflows.CreateWorkflowComponentCcMilestone.
func CreateWorkflowComponentCcMilestone(requester exhttp.Requester, ctx context.Context, ccID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateWorkflowComponentCcMilestone(ctx, ccID, payload, raiseOnError, responseFormat...)
}

// GetWorkflowComponentHistory is the module form of Workflows.GetWorkflowComponentHistory.
func GetWorkflowComponentHistory(requester exhttp.Requester, ctx context.Context, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWorkflowComponentHistory(ctx, componentObjectID, queryParams, raiseOnError, responseFormat...)
}

// GetWorkflowComponentCcInstanceCcMilestone is the module form of Workflows.GetWorkflowComponentCcInstanceCcMilestone.
func GetWorkflowComponentCcInstanceCcMilestone(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWorkflowComponentCcInstanceCcMilestone(ctx, componentUUID, componentObjectID, queryParams, raiseOnError, responseFormat...)
}

// GetWorkflows is the module form of Workflows.GetWorkflows.
func GetWorkflows(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWorkflows(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetWorkflow is the module form of Workflows.GetWorkflow.
func GetWorkflow(requester exhttp.Requester, ctx context.Context, workflowID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWorkflow(ctx, workflowID, raiseOnError, responseFormat...)
}

// CreateWorkflow is the module form of Workflows.CreateWorkflow.
func CreateWorkflow(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateWorkflow(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateWorkflow is the module form of Workflows.UpdateWorkflow.
func UpdateWorkflow(requester exhttp.Requester, ctx context.Context, workflowID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateWorkflow(ctx, workflowID, payload, raiseOnError, responseFormat...)
}

// DeleteWorkflow is the module form of Workflows.DeleteWorkflow.
func DeleteWorkflow(requester exhttp.Requester, ctx context.Context, workflowID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteWorkflow(ctx, workflowID, raiseOnError, responseFormat...)
}

// GetWorkflowStages is the module form of Workflows.GetWorkflowStages.
func GetWorkflowStages(requester exhttp.Requester, ctx context.Context, workflowID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWorkflowStages(ctx, workflowID, queryParams, raiseOnError, responseFormat...)
}

// GetWorkflowStage is the module form of Workflows.GetWorkflowStage.
func GetWorkflowStage(requester exhttp.Requester, ctx context.Context, workflowID any, stageID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWorkflowStage(ctx, workflowID, stageID, raiseOnError, responseFormat...)
}

// GetWorkflowMilestone is the module form of Workflows.GetWorkflowMilestone.
func GetWorkflowMilestone(requester exhttp.Requester, ctx context.Context, workflowID any, milestoneID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWorkflowMilestone(ctx, workflowID, milestoneID, raiseOnError, responseFormat...)
}

// GetWorkflowCustomComponentHistory is the module form of Workflows.GetWorkflowCustomComponentHistory.
func GetWorkflowCustomComponentHistory(requester exhttp.Requester, ctx context.Context, componentUUID any, objectCode any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWorkflowCustomComponentHistory(ctx, componentUUID, objectCode, queryParams, raiseOnError, responseFormat...)
}

// GetWorkflowCustomComponentMilestones is the module form of Workflows.GetWorkflowCustomComponentMilestones.
func GetWorkflowCustomComponentMilestones(requester exhttp.Requester, ctx context.Context, componentUUID any, objectCode any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWorkflowCustomComponentMilestones(ctx, componentUUID, objectCode, queryParams, raiseOnError, responseFormat...)
}

// GetWorkflowCustomComponentMilestone is the module form of Workflows.GetWorkflowCustomComponentMilestone.
func GetWorkflowCustomComponentMilestone(requester exhttp.Requester, ctx context.Context, componentUUID any, objectCode any, milestoneID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWorkflowCustomComponentMilestone(ctx, componentUUID, objectCode, milestoneID, raiseOnError, responseFormat...)
}
