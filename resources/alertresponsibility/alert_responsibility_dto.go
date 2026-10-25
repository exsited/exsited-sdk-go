package alertresponsibility

import "github.com/exsited/exsited-sdk-go/v4/dto"

// AlertResponsibilityList is one page of AlertResponsibilityRecord records, under data.
type AlertResponsibilityList struct {
	Data       []AlertResponsibilityRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AlertResponsibilityList was measured from.
func (AlertResponsibilityList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/alert_responsibility",
	}
}

// AlertResponsibilityRecord is the alert responsibility body, under alert_responsibility;
// activity_logs comes on the create answer only.
type AlertResponsibilityRecord struct {
	ActivityLogs  []any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Status        *string
}

// MeasuredFrom lists the endpoints AlertResponsibilityRecord was measured from.
func (AlertResponsibilityRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/alert_responsibility",
		"GET /api/{api_version}/alert_responsibility/{alert_responsibility_id}",
		"POST /api/{api_version}/alert_responsibility",
		"PATCH /api/{api_version}/alert_responsibility/{alert_responsibility_id}",
	}
}

// EnvelopeKey unwraps {"alert_responsibility": {...}}.
func (AlertResponsibilityRecord) EnvelopeKey() string { return AlertResponsibilityKey }

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

// AlertResponsibilityListShape reads a body as AlertResponsibilityList.
func AlertResponsibilityListShape(decoded any) any { return dto.Read[AlertResponsibilityList](decoded) }

// AlertResponsibilityRecordShape reads a body as AlertResponsibilityRecord.
func AlertResponsibilityRecordShape(decoded any) any {
	return dto.Read[AlertResponsibilityRecord](decoded)
}
