package actionscripts

import "github.com/exsited/exsited-sdk-go/v4/dto"

// ActionScriptList is one page of ActionScriptRecord records, under data.
type ActionScriptList struct {
	Data       []ActionScriptRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ActionScriptList was measured from.
func (ActionScriptList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/action_scripts",
	}
}

// ActionScriptRecord is the action script body, under action_script; activity_logs comes on the
// create answer only.
type ActionScriptRecord struct {
	ActionName    any
	ActivityLogs  []any
	CreatedBy     *ActionScriptCreatedBy
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy *ActionScriptLastUpdatedBy
	LastUpdatedOn any
	Name          *string
	Object        *string
	ObjectName    any
	Status        *string
}

// MeasuredFrom lists the endpoints ActionScriptRecord was measured from.
func (ActionScriptRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/action_scripts",
		"GET /api/{api_version}/action_scripts/{action_script_id}",
		"POST /api/{api_version}/action_scripts",
		"PATCH /api/{api_version}/action_scripts/{action_script_id}",
	}
}

// EnvelopeKey unwraps {"action_script": {...}}.
func (ActionScriptRecord) EnvelopeKey() string { return ActionScriptKey }

// ActionScriptCreatedBy is an ActionScriptRecord's created_by.
type ActionScriptCreatedBy struct {
	Name   *string
	Object *string
}

// ActionScriptLastUpdatedBy is an ActionScriptRecord's last_updated_by, set once the script has
// been updated.
type ActionScriptLastUpdatedBy struct {
	Name   *string
	Object *string
}

// Pagination is a list envelope's paging block. One response fills only one shape: cursor (HasMore
// present), offset (Records present) or the external database listings' (HasNext present), and the
// others read nil.
type Pagination struct {
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	PrevCursor *string
	Links      *PaginationLinks

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	NextPage     *string
	PreviousPage *string

	// The external database listings' shape.
	HasNext *bool
	HasPrev *bool
	PerPage *int64
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// ActionScriptListShape reads a body as ActionScriptList.
func ActionScriptListShape(decoded any) any { return dto.Read[ActionScriptList](decoded) }

// ActionScriptRecordShape reads a body as ActionScriptRecord.
func ActionScriptRecordShape(decoded any) any { return dto.Read[ActionScriptRecord](decoded) }
