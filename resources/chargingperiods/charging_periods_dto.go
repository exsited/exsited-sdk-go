package chargingperiods

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsChargingPeriodList is one page of SettingsChargingPeriodRecord records, under data.
type SettingsChargingPeriodList struct {
	Data       []SettingsChargingPeriodRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsChargingPeriodList was measured from.
func (SettingsChargingPeriodList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/charging_periods",
	}
}

// SettingsChargingPeriodRecord is the settings charging period body, under charging_period.
type SettingsChargingPeriodRecord struct {
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
	TimePeriod    *SettingsChargingPeriodTimePeriod
}

// MeasuredFrom lists the endpoints SettingsChargingPeriodRecord was measured from.
func (SettingsChargingPeriodRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/charging_periods",
		"GET /api/{api_version}/settings/charging_periods/{charging_period_id}",
		"POST /api/{api_version}/settings/charging_periods",
		"PATCH /api/{api_version}/settings/charging_periods/{charging_period_id}",
	}
}

// EnvelopeKey unwraps {"charging_period": {...}}.
func (SettingsChargingPeriodRecord) EnvelopeKey() string { return ChargingPeriodKey }

// SettingsChargingPeriodTimePeriod is a SettingsChargingPeriodRecord's time_period.
type SettingsChargingPeriodTimePeriod struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

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

// SettingsChargingPeriodListShape reads a body as SettingsChargingPeriodList.
func SettingsChargingPeriodListShape(decoded any) any {
	return dto.Read[SettingsChargingPeriodList](decoded)
}

// SettingsChargingPeriodRecordShape reads a body as SettingsChargingPeriodRecord.
func SettingsChargingPeriodRecordShape(decoded any) any {
	return dto.Read[SettingsChargingPeriodRecord](decoded)
}
