package receiveterms

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsReceiveTermList is one page of SettingsReceiveTermRecord records, under data.
type SettingsReceiveTermList struct {
	Data       []SettingsReceiveTermRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsReceiveTermList was measured from.
func (SettingsReceiveTermList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/receive_terms",
	}
}

// SettingsReceiveTermRecord is the settings receive term body, under receive_term.
type SettingsReceiveTermRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Duration      *int64
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Period        *string
	Status        *string
}

// MeasuredFrom lists the endpoints SettingsReceiveTermRecord was measured from.
func (SettingsReceiveTermRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/receive_terms",
		"GET /api/{api_version}/settings/receive_terms/{receive_term_id}",
		"POST /api/{api_version}/settings/receive_terms",
		"PATCH /api/{api_version}/settings/receive_terms/{receive_term_id}",
	}
}

// EnvelopeKey unwraps {"receive_term": {...}}.
func (SettingsReceiveTermRecord) EnvelopeKey() string { return ReceiveTermKey }

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

// SettingsReceiveTermListShape reads a body as SettingsReceiveTermList.
func SettingsReceiveTermListShape(decoded any) any { return dto.Read[SettingsReceiveTermList](decoded) }

// SettingsReceiveTermRecordShape reads a body as SettingsReceiveTermRecord.
func SettingsReceiveTermRecordShape(decoded any) any {
	return dto.Read[SettingsReceiveTermRecord](decoded)
}
