package settingsworkflows

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsWorkflowList is one page of SettingsWorkflowRecord records, under data.
type SettingsWorkflowList struct {
	Data       []SettingsWorkflowRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsWorkflowList was measured from.
func (SettingsWorkflowList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/workflows",
	}
}

// SettingsWorkflowRecord is the settings workflow body, under workflow.
type SettingsWorkflowRecord struct {
	Code          any
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayName   *string
	EnableStage   *bool
	Entity        *string
	Event         *string
	Href          *string
	ID            *string
	IsPreBuilt    *bool
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Status        *string
}

// MeasuredFrom lists the endpoints SettingsWorkflowRecord was measured from.
func (SettingsWorkflowRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/workflows",
		"GET /api/{api_version}/settings/workflows/{workflow_id}",
	}
}

// EnvelopeKey unwraps {"workflow": {...}}.
func (SettingsWorkflowRecord) EnvelopeKey() string { return WorkflowKey }

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

// SettingsWorkflowListShape reads a body as SettingsWorkflowList.
func SettingsWorkflowListShape(decoded any) any { return dto.Read[SettingsWorkflowList](decoded) }

// SettingsWorkflowRecordShape reads a body as SettingsWorkflowRecord.
func SettingsWorkflowRecordShape(decoded any) any { return dto.Read[SettingsWorkflowRecord](decoded) }
