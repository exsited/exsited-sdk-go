package systemevents

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsSystemEventList is one page of SettingsSystemEventRecord records, under data.
type SettingsSystemEventList struct {
	Data       []SettingsSystemEventRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsSystemEventList was measured from.
func (SettingsSystemEventList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/system_events",
	}
}

// SettingsSystemEventRecord is the settings system event body, under system_event.
type SettingsSystemEventRecord struct {
	BaseEvent     *bool
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Status        *string
}

// MeasuredFrom lists the endpoints SettingsSystemEventRecord was measured from.
func (SettingsSystemEventRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/system_events",
		"GET /api/{api_version}/settings/system_events/{system_event_id}",
	}
}

// EnvelopeKey unwraps {"system_event": {...}}.
func (SettingsSystemEventRecord) EnvelopeKey() string { return SystemEventKey }

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

// SettingsSystemEventListShape reads a body as SettingsSystemEventList.
func SettingsSystemEventListShape(decoded any) any { return dto.Read[SettingsSystemEventList](decoded) }

// SettingsSystemEventRecordShape reads a body as SettingsSystemEventRecord.
func SettingsSystemEventRecordShape(decoded any) any {
	return dto.Read[SettingsSystemEventRecord](decoded)
}
