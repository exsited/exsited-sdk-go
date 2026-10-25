package salesreps

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsSalesRepList is one page of SettingsSalesRepRecord records, under data.
type SettingsSalesRepList struct {
	Data       []SettingsSalesRepRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsSalesRepList was measured from.
func (SettingsSalesRepList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/sales_reps",
	}
}

// SettingsSalesRepRecord is the settings sales rep body, under sales_rep.
type SettingsSalesRepRecord struct {
	CreatedBy     any
	CreatedOn     *string
	EmailAddress  *string
	FirstName     *string
	Href          *string
	ID            *string
	LastName      *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Status        *string
	Version       *int64
}

// MeasuredFrom lists the endpoints SettingsSalesRepRecord was measured from.
func (SettingsSalesRepRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/sales_reps",
		"GET /api/{api_version}/settings/sales_reps/{sales_rep_id}",
	}
}

// EnvelopeKey unwraps {"sales_rep": {...}}.
func (SettingsSalesRepRecord) EnvelopeKey() string { return SalesRepKey }

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

// SettingsSalesRepListShape reads a body as SettingsSalesRepList.
func SettingsSalesRepListShape(decoded any) any { return dto.Read[SettingsSalesRepList](decoded) }

// SettingsSalesRepRecordShape reads a body as SettingsSalesRepRecord.
func SettingsSalesRepRecordShape(decoded any) any { return dto.Read[SettingsSalesRepRecord](decoded) }
