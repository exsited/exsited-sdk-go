package workflows

import "github.com/exsited/exsited-sdk-go/v4/dto"

// WorkflowMilestoneList is the page GetWorkflowMilestones answers.
type WorkflowMilestoneList struct {
	Data       []WorkflowMilestoneRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints WorkflowMilestoneList was measured from.
func (WorkflowMilestoneList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/workflows/{workflow_uuid}/milestones",
	}
}

// WorkflowMilestoneRecord is one row of WorkflowMilestoneList.
type WorkflowMilestoneRecord struct {
	ActivityDefinition *ActivityDefinition
	AutoCreate         any
	CreatedOn          *string
	Description        any
	DisplayName        *string
	EstimatedHours     any
	Href               *string
	Icon               any
	ID                 *string
	LastUpdatedOn      *string
	Name               *string
	Object             *string
	Priority           *string
	SortOrder          any
	Status             *string
	Workflow           *Workflow
	WorkflowVersion    *WorkflowVersion
}

// MeasuredFrom lists the endpoints WorkflowMilestoneRecord was measured from.
func (WorkflowMilestoneRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/workflows/{workflow_uuid}/milestones",
	}
}

// ActivityDefinition is the activity_definition object.
type ActivityDefinition struct {
	ID     *string
	Name   *string
	Object *string
}

// Pagination is the cursor paging block.
type Pagination struct {
	HasMore    *bool
	Limit      *int64
	Links      *PaginationLinks
	NextCursor any
	PrevCursor any
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// Workflow is the workflow object.
type Workflow struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// WorkflowVersion is the workflow_version object.
type WorkflowVersion struct {
	ID     *string
	Name   *string
	Object *string
}

// WorkflowMilestoneListShape reads a body as WorkflowMilestoneList.
func WorkflowMilestoneListShape(decoded any) any { return dto.Read[WorkflowMilestoneList](decoded) }

// WorkflowList is one page of what GetWorkflows lists, under data.
type WorkflowList struct {
	Data       []WorkflowView
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints WorkflowList was measured from.
func (WorkflowList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/workflows",
	}
}

// WorkflowView is one entry of a WorkflowList's data.
type WorkflowView struct {
	Object         *string
	ID             *string
	Status         *string
	Code           any
	Name           *string
	DisplayName    *string
	Description    any
	Entity         *string
	EntityID       *int64
	Event          *string
	IsPreBuilt     *bool
	EnableStage    *bool
	CurrentVersion *WorkflowCurrentVersion
	CustomForm     any
	CreatedOn      *string
	LastUpdatedOn  *string
	CreatedBy      *WorkflowCreatedBy
	LastUpdatedBy  *WorkflowLastUpdatedBy
	Href           *string
}

// WorkflowCurrentVersion is a WorkflowView's current_version.
type WorkflowCurrentVersion struct {
	Object *string
	ID     *string
	Name   *string
}

// WorkflowCreatedBy is a WorkflowView's created_by.
type WorkflowCreatedBy struct {
	Object *string
	Name   *string
}

// WorkflowLastUpdatedBy is a WorkflowView's last_updated_by.
type WorkflowLastUpdatedBy struct {
	Object *string
	Name   *string
}

// WorkflowMilestoneView is one milestone definition, under workflow_milestone: what
// GetWorkflowMilestone answers, the same fields as a listed one.
type WorkflowMilestoneView struct {
	ActivityDefinition *WorkflowMilestoneActivityDefinition
	AutoCreate         any
	CreatedOn          *string
	Description        any
	DisplayName        *string
	EstimatedHours     any
	Href               *string
	Icon               any
	ID                 *string
	LastUpdatedOn      *string
	Name               *string
	Object             *string
	Priority           *string
	SortOrder          any
	Status             *string
	Workflow           *WorkflowMilestoneDetail
	WorkflowVersion    *WorkflowMilestoneWorkflowVersion
}

// MeasuredFrom lists the endpoints WorkflowMilestoneView was measured from.
func (WorkflowMilestoneView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/workflows/{workflow_id}/milestones/{milestone_id}",
	}
}

// EnvelopeKey unwraps {"workflow_milestone": {...}}.
func (WorkflowMilestoneView) EnvelopeKey() string { return WorkflowMilestoneKey }

// WorkflowMilestoneActivityDefinition is a WorkflowMilestone's activity_definition.
type WorkflowMilestoneActivityDefinition struct {
	ID     *string
	Name   *string
	Object *string
}

// WorkflowMilestoneDetail is a WorkflowMilestone's workflow.
type WorkflowMilestoneDetail struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// WorkflowMilestoneWorkflowVersion is a WorkflowMilestone's workflow_version.
type WorkflowMilestoneWorkflowVersion struct {
	ID     *string
	Name   *string
	Object *string
}

// WorkflowListShape reads a body as WorkflowList.
func WorkflowListShape(decoded any) any { return dto.Read[WorkflowList](decoded) }

// WorkflowViewShape reads a body as WorkflowView.
func WorkflowViewShape(decoded any) any { return dto.Read[WorkflowView](decoded) }

// WorkflowMilestoneViewShape reads a body as WorkflowMilestoneView.
func WorkflowMilestoneViewShape(decoded any) any { return dto.Read[WorkflowMilestoneView](decoded) }
