package customevents

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsCustomEventList is one page of SettingsCustomEventRecord records, under data.
type SettingsCustomEventList struct {
	Data       []SettingsCustomEventRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsCustomEventList was measured from.
func (SettingsCustomEventList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/custom_events",
	}
}

// SettingsCustomEventRecord is the settings custom event body, under custom_event.
type SettingsCustomEventRecord struct {
	BaseEvent     *bool
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Status        *string
}

// MeasuredFrom lists the endpoints SettingsCustomEventRecord was measured from.
func (SettingsCustomEventRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/custom_events",
		"GET /api/{api_version}/settings/custom_events/{custom_event_id}",
	}
}

// EnvelopeKey unwraps {"custom_event": {...}}.
func (SettingsCustomEventRecord) EnvelopeKey() string { return CustomEventKey }

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

// SettingsCustomEventListShape reads a body as SettingsCustomEventList.
func SettingsCustomEventListShape(decoded any) any { return dto.Read[SettingsCustomEventList](decoded) }

// SettingsCustomEventRecordShape reads a body as SettingsCustomEventRecord.
func SettingsCustomEventRecordShape(decoded any) any {
	return dto.Read[SettingsCustomEventRecord](decoded)
}
