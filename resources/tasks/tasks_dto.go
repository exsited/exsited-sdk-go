package tasks

import "github.com/exsited/exsited-sdk-go/v4/dto"

// Pagination is the paging block, either shape.
type Pagination struct {
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	PrevCursor any
	Links      *PaginationLinks

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	NextPage     *string
	PreviousPage *string
}

// TaskRecord is the task body.
type TaskRecord struct {
	ActualHours      any
	Code             *string
	CompletedAt      any
	CompletedBy      *CompletedBy
	CreatedBy        *CreatedBy
	CreatedOn        *string
	CustomAttributes []any
	Description      *string
	DisplayName      *string
	DueDate          any
	EstimatedHours   any
	Href             *string
	ID               *string
	LastUpdatedBy    *LastUpdatedBy
	LastUpdatedOn    any
	Milestone        any
	Name             *string
	Object           *string
	ParentTask       any
	Priority         *string
	Resource         any
	SortOrder        *int64
	StartAfter       any
	StartDate        any
	State            any
	Status           *string
	TaskStatus       *string
	Version          *int64
	Workflow         any
}

// MeasuredFrom lists the endpoints TaskRecord was measured from.
func (TaskRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/tasks",
	}
}

// TaskList is a page of TaskRecord records.
type TaskList struct {
	Data       []TaskRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints TaskList was measured from.
func (TaskList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/tasks",
	}
}

// TaskChecklistReadRecord is the body GetTaskChecklist answers.
type TaskChecklistReadRecord struct {
	Checklist []Checklist
}

// MeasuredFrom lists the endpoints TaskChecklistReadRecord was measured from.
func (TaskChecklistReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/tasks/{task_id}/checklist",
	}
}

// EnvelopeKey unwraps {"work_task": {...}}.
func (TaskChecklistReadRecord) EnvelopeKey() string { return WorkTaskKey }

// TaskNoteList is the page GetTaskNotes answers.
type TaskNoteList struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints TaskNoteList was measured from.
func (TaskNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/tasks/{task_code}/notes",
	}
}

// TaskReadRecord is the body GetTask answers.
type TaskReadRecord struct {
	ActivityLogs     []ActivityLog
	ActualHours      any
	Alerts           []any
	Assignees        *Assignees
	Checklist        []Checklist
	Code             *string
	CompletedAt      *string
	CompletedBy      *CompletedBy
	CreatedBy        *CreatedBy
	CreatedOn        *string
	CustomAttributes []any
	Description      *string
	DisplayName      *string
	DueDate          *string
	EstimatedHours   *int64
	Href             *string
	ID               *string
	LastUpdatedBy    *LastUpdatedBy
	LastUpdatedOn    *string
	Milestone        any
	Name             *string
	Object           *string
	ParentTask       any
	Priority         *string
	Resource         any
	SortOrder        *int64
	StartAfter       any
	StartDate        *string
	State            any
	Status           *string
	TaskStatus       *string
	Version          *int64
	Workflow         any
}

// MeasuredFrom lists the endpoints TaskReadRecord was measured from.
func (TaskReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/tasks/{task_id}",
		"PATCH /api/{api_version}/tasks/{task_id}",
		"POST /api/{api_version}/tasks/{task_id}/checklist",
	}
}

// EnvelopeKey unwraps {"work_task": {...}}.
func (TaskReadRecord) EnvelopeKey() string { return WorkTaskKey }

// ActivityLog is one entry of the activity_logs list.
type ActivityLog struct {
	Activity *string
}

// Assignees is the assignees object.
type Assignees struct {
	ManuallyAddedAccountContacts []any
	SelectedAccountContacts      []any
	SelectedEmployees            []SelectedEmployee
	SelectedLabours              []SelectedLabour
	SelectedOperators            []any
}

// Checklist is one entry of the checklist list.
type Checklist struct {
	CompletedAt   any
	CompletedBy   any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	ID            *string
	LastUpdatedOn *string
	Name          *string
	Object        *string
	SortOrder     *int64
	Status        *string
	Version       *int64
}

// CompletedBy is the completed_by object.
type CompletedBy struct {
	Name   *string
	Object *string
}

// CreatedBy is the created_by object.
type CreatedBy struct {
	Name   *string
	Object *string
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev any
}

// SelectedEmployee is one entry of the selected_employees list.
type SelectedEmployee struct {
	DisplayOrder   *int64
	Email          *string
	EmployeeCode   *string
	Name           *string
	Responsibility any
	UUID           *string
}

// SelectedLabour is one entry of the selected_labours list.
type SelectedLabour struct {
	DisplayOrder   *int64
	Email          *string
	LabourCode     *string
	Name           *string
	Responsibility any
	UUID           *string
}

// TaskChecklistByUUIDReadRecord is the body GetTaskChecklistByUUID answers.
type TaskChecklistByUUIDReadRecord struct {
	Checklist *Checklist
}

// MeasuredFrom lists the endpoints TaskChecklistByUUIDReadRecord was measured from.
func (TaskChecklistByUUIDReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/tasks/{task_id}/checklist/{task_checklist_uuid}",
	}
}

// EnvelopeKey unwraps {"work_task": {...}}.
func (TaskChecklistByUUIDReadRecord) EnvelopeKey() string { return WorkTaskKey }

// Email is the email object.
type Email struct {
	Address       *string
	EmailOptedOut *bool
}

// Name is the name object.
type Name struct {
	Designation any
	First       *string
	Last        *string
	Middle      any
	Salutation  any
}

// TaskListShape reads a body as TaskList.
func TaskListShape(decoded any) any { return dto.Read[TaskList](decoded) }

// TaskChecklistReadRecordShape reads a body as TaskChecklistReadRecord.
func TaskChecklistReadRecordShape(decoded any) any { return dto.Read[TaskChecklistReadRecord](decoded) }

// TaskNoteListShape reads a body as TaskNoteList.
func TaskNoteListShape(decoded any) any { return dto.Read[TaskNoteList](decoded) }

// TaskReadRecordShape reads a body as TaskReadRecord.
func TaskReadRecordShape(decoded any) any { return dto.Read[TaskReadRecord](decoded) }

// TaskChecklistByUUIDReadRecordShape reads a body as TaskChecklistByUUIDReadRecord.
func TaskChecklistByUUIDReadRecordShape(decoded any) any {
	return dto.Read[TaskChecklistByUUIDReadRecord](decoded)
}
