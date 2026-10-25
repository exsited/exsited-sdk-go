package workflows

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Workflows is the v4 workflows endpoints, bound to one client.
type Workflows struct {
	requester exhttp.Requester
}

// DeleteMilestone sends DELETE /api/{api_version}/milestones/{milestone_uuid}.
func (w *Workflows) DeleteMilestone(ctx context.Context, milestoneUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(milestoneUUID, "milestone_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendDeleteRequest(ctx, MilestonePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"milestone_uuid": milestoneUUID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// UpdateWorkflowComponentByID sends PATCH /api/{api_version}/workflow/component/{cc_id}. v4 answers this
// path 405; the documented route takes two slots where this path has one, and no body is verified.
func (w *Workflows) UpdateWorkflowComponentByID(ctx context.Context, ccID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(ccID, "cc_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendPatchRequest(ctx, WorkflowComponentByIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"cc_id": ccID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetWorkflowMilestones sends GET /api/{api_version}/workflows/{workflow_uuid}/milestones.
func (w *Workflows) GetWorkflowMilestones(ctx context.Context, workflowUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(workflowUUID, "workflow_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendGetRequest(ctx, WorkflowMilestonesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"workflow_uuid": workflowUUID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: WorkflowMilestoneListShape})
}

// CreateWorkflowComponentCcMilestone sends POST /api/{api_version}/workflow/component/{cc_id}/milestone. v4
// answers this path 405; the documented route takes two slots where this path has one, and no body is
// verified.
func (w *Workflows) CreateWorkflowComponentCcMilestone(ctx context.Context, ccID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(ccID, "cc_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendPostRequest(ctx, WorkflowComponentCcMilestonePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"cc_id": ccID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetWorkflowComponentHistory reads the workflow history of one component instance, by its code:
// GET /api/{api_version}/workflow/component/{component_object_id}/history, answered under histories.
func (w *Workflows) GetWorkflowComponentHistory(ctx context.Context, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendGetRequest(ctx, WorkflowComponentHistoryPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"component_object_id": componentObjectID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetWorkflowComponentCcInstanceCcMilestone reads the milestones of one component instance, by the
// component's uuid and the instance's code:
// GET /api/{api_version}/workflow/component/{component_uuid}/{component_object_id}/milestone, answered under
// milestones.
func (w *Workflows) GetWorkflowComponentCcInstanceCcMilestone(ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendGetRequest(ctx, WorkflowComponentCcInstanceCcMilestonePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetWorkflows sends GET /api/{api_version}/workflows.
func (w *Workflows) GetWorkflows(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendGetRequest(ctx, WorkflowsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: WorkflowListShape})
}

// GetWorkflow sends GET /api/{api_version}/workflows/{workflow_id}.
func (w *Workflows) GetWorkflow(ctx context.Context, workflowID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(workflowID, "workflow_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendGetRequest(ctx, WorkflowPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"workflow_id": workflowID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: WorkflowViewShape})
}

// CreateWorkflow sends POST /api/{api_version}/workflows.
func (w *Workflows) CreateWorkflow(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendPostRequest(ctx, WorkflowsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: WorkflowViewShape})
}

// UpdateWorkflow sends PATCH /api/{api_version}/workflows/{workflow_id}.
func (w *Workflows) UpdateWorkflow(ctx context.Context, workflowID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(workflowID, "workflow_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendPatchRequest(ctx, WorkflowPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"workflow_id": workflowID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: WorkflowViewShape})
}

// DeleteWorkflow sends DELETE /api/{api_version}/workflows/{workflow_id}; it has no typed view, so read ApiResponse.Map.
func (w *Workflows) DeleteWorkflow(ctx context.Context, workflowID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(workflowID, "workflow_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendDeleteRequest(ctx, WorkflowPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"workflow_id": workflowID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetWorkflowStages sends GET /api/{api_version}/workflows/{workflow_id}/stages; it has no typed view, so read ApiResponse.Map.
func (w *Workflows) GetWorkflowStages(ctx context.Context, workflowID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(workflowID, "workflow_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendGetRequest(ctx, WorkflowStagesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"workflow_id": workflowID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetWorkflowStage sends GET /api/{api_version}/workflows/{workflow_id}/stages/{stage_id}; it has no typed view, so read ApiResponse.Map.
func (w *Workflows) GetWorkflowStage(ctx context.Context, workflowID any, stageID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(workflowID, "workflow_id"); err != nil {
		return nil, err
	}
	if err := requireValue(stageID, "stage_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendGetRequest(ctx, WorkflowStagePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"workflow_id": workflowID, "stage_id": stageID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetWorkflowMilestone sends GET /api/{api_version}/workflows/{workflow_id}/milestones/{milestone_id}.
func (w *Workflows) GetWorkflowMilestone(ctx context.Context, workflowID any, milestoneID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(workflowID, "workflow_id"); err != nil {
		return nil, err
	}
	if err := requireValue(milestoneID, "milestone_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendGetRequest(ctx, WorkflowMilestonePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"workflow_id": workflowID, "milestone_id": milestoneID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: WorkflowMilestoneViewShape})
}

// GetWorkflowCustomComponentHistory sends GET /api/{api_version}/workflows/custom_components/{component_uuid}/{object_code}/history; it has no typed view, so read ApiResponse.Map.
func (w *Workflows) GetWorkflowCustomComponentHistory(ctx context.Context, componentUUID any, objectCode any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(objectCode, "object_code"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendGetRequest(ctx, WorkflowCustomComponentHistoryPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "object_code": objectCode}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetWorkflowCustomComponentMilestones sends GET /api/{api_version}/workflows/custom_components/{component_uuid}/{object_code}/milestones; it has no typed view, so read ApiResponse.Map.
func (w *Workflows) GetWorkflowCustomComponentMilestones(ctx context.Context, componentUUID any, objectCode any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(objectCode, "object_code"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendGetRequest(ctx, WorkflowCustomComponentMilestonesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "object_code": objectCode}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetWorkflowCustomComponentMilestone sends GET /api/{api_version}/workflows/custom_components/{component_uuid}/{object_code}/milestones/{milestone_id}; it has no typed view, so read ApiResponse.Map.
func (w *Workflows) GetWorkflowCustomComponentMilestone(ctx context.Context, componentUUID any, objectCode any, milestoneID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(objectCode, "object_code"); err != nil {
		return nil, err
	}
	if err := requireValue(milestoneID, "milestone_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return w.requester.SendGetRequest(ctx, WorkflowCustomComponentMilestonePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "object_code": objectCode, "milestone_id": milestoneID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "workflows"

// oneResponseFormat reads the optional trailing responseFormat, refusing more than one.
func oneResponseFormat(responseFormat []string) (string, error) {
	switch len(responseFormat) {
	case 0:
		return "", nil
	case 1:
		return responseFormat[0], nil
	default:
		return "", exerrors.NewUsageError(fmt.Sprintf(
			"responseFormat: pass one format or none, got %d; the trailing argument is optional, "+
				"and omitting it defers to the client", len(responseFormat)))
	}
}

// requirePayload refuses a nil body before any I/O; the client's validator would let one through bodyless.
func requirePayload(payload any) error {
	if payload == nil || isNilValue(payload) {
		return exerrors.NewUsageError("payload is required: this endpoint takes a body, and nil would " +
			"send the request with no body at all")
	}
	return nil
}

// isNilValue reports whether a value is a typed nil inside a non-nil any; a zero-valued struct is not.
func isNilValue(value any) bool {
	reflected := reflect.ValueOf(value)
	switch reflected.Kind() {
	case reflect.Map, reflect.Slice, reflect.Pointer, reflect.Interface, reflect.Func:
		return reflected.IsNil()
	}
	return false
}

// requireValue refuses a path value that is nil, empty or only whitespace, before any I/O.
func requireValue(value any, what string) error {
	text, isText := value.(string)
	if value == nil || isNilValue(value) || isText && strings.TrimSpace(text) == "" {
		return exerrors.UsageErrorf("%s is required: the request cannot be built without it", what)
	}
	return nil
}

// New binds the workflows endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Workflows { return &Workflows{requester: requester} }

// Requester returns the requester this resource was bound to.
func (w *Workflows) Requester() exhttp.Requester { return w.requester }
