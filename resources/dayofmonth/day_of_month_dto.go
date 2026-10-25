package dayofmonth

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsDayOfMonthList is one page of SettingsDayOfMonthRecord records, under data.
type SettingsDayOfMonthList struct {
	Data       []SettingsDayOfMonthRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsDayOfMonthList was measured from.
func (SettingsDayOfMonthList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/day_of_month",
	}
}

// SettingsDayOfMonthRecord is the settings day of month body, under day_of_month.
type SettingsDayOfMonthRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Day           *int64
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

// MeasuredFrom lists the endpoints SettingsDayOfMonthRecord was measured from.
func (SettingsDayOfMonthRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/day_of_month",
		"GET /api/{api_version}/settings/day_of_month/{day_of_month_id}",
	}
}

// EnvelopeKey unwraps {"day_of_month": {...}}.
func (SettingsDayOfMonthRecord) EnvelopeKey() string { return DayOfMonthKey }

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

// SettingsDayOfMonthListShape reads a body as SettingsDayOfMonthList.
func SettingsDayOfMonthListShape(decoded any) any { return dto.Read[SettingsDayOfMonthList](decoded) }

// SettingsDayOfMonthRecordShape reads a body as SettingsDayOfMonthRecord.
func SettingsDayOfMonthRecordShape(decoded any) any {
	return dto.Read[SettingsDayOfMonthRecord](decoded)
}
