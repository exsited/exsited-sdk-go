package timeperiods

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsTimePeriodList is one page of SettingsTimePeriodRecord records, under data.
type SettingsTimePeriodList struct {
	Data       []SettingsTimePeriodRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsTimePeriodList was measured from.
func (SettingsTimePeriodList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/time_periods",
	}
}

// SettingsTimePeriodRecord is the settings time period body, under time_period.
type SettingsTimePeriodRecord struct {
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

// MeasuredFrom lists the endpoints SettingsTimePeriodRecord was measured from.
func (SettingsTimePeriodRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/time_periods",
		"GET /api/{api_version}/settings/time_periods/{time_period_id}",
		"POST /api/{api_version}/settings/time_periods",
		"PATCH /api/{api_version}/settings/time_periods/{time_period_id}",
	}
}

// EnvelopeKey unwraps {"time_period": {...}}.
func (SettingsTimePeriodRecord) EnvelopeKey() string { return TimePeriodKey }

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

// SettingsTimePeriodListShape reads a body as SettingsTimePeriodList.
func SettingsTimePeriodListShape(decoded any) any { return dto.Read[SettingsTimePeriodList](decoded) }

// SettingsTimePeriodRecordShape reads a body as SettingsTimePeriodRecord.
func SettingsTimePeriodRecordShape(decoded any) any {
	return dto.Read[SettingsTimePeriodRecord](decoded)
}
