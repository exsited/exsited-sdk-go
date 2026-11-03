package operators

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsOperatorList is one page of SettingsOperatorRecord records, under data.
type SettingsOperatorList struct {
	Data       []SettingsOperatorRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsOperatorList is returned by.
func (SettingsOperatorList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/operators",
	}
}

// SettingsOperatorRecord is the operator body, under operator. Description, CreatedBy and
// LastUpdatedBy are untyped.
type SettingsOperatorRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayName   *string
	Email         *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
}

// MeasuredFrom lists the endpoints SettingsOperatorRecord is returned by.
func (SettingsOperatorRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/operators",
		"GET /api/{api_version}/settings/operators/{operator_id}",
		"POST /api/{api_version}/settings/operators",
		"PATCH /api/{api_version}/settings/operators/{operator_id}",
	}
}

// EnvelopeKey unwraps {"operator": {...}}.
func (SettingsOperatorRecord) EnvelopeKey() string { return OperatorKey }

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

// SettingsOperatorListShape reads a body as SettingsOperatorList.
func SettingsOperatorListShape(decoded any) any { return dto.Read[SettingsOperatorList](decoded) }

// SettingsOperatorRecordShape reads a body as SettingsOperatorRecord.
func SettingsOperatorRecordShape(decoded any) any { return dto.Read[SettingsOperatorRecord](decoded) }
