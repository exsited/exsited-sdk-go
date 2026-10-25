package events

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

// EventRecord is the event body.
type EventRecord struct {
	CreatedOn            *string
	EventName            *string
	ExecutedOn           *string
	Href                 *string
	ID                   *string
	Object               *string
	ResourceType         *string
	ResourceUUID         *string
	ScheduledToExecuteOn *string
	Status               any
}

// MeasuredFrom lists the endpoints EventRecord was measured from.
func (EventRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/events",
		"GET /api/{api_version}/events/{event_uuid}",
		"POST /api/{api_version}/events/{event_uuid}/remove",
	}
}

// EnvelopeKey unwraps {"event": {...}}.
func (EventRecord) EnvelopeKey() string { return EventKey }

// EventList is a page of EventRecord records.
type EventList struct {
	Data       []EventRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints EventList was measured from.
func (EventList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/events",
	}
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// EventListShape reads a body as EventList.
func EventListShape(decoded any) any { return dto.Read[EventList](decoded) }

// EventRecordShape reads a body as EventRecord.
func EventRecordShape(decoded any) any { return dto.Read[EventRecord](decoded) }
