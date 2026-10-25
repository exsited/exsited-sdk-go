package pricingperiods

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsPricingPeriodList is one page of SettingsPricingPeriodRecord records, under data.
type SettingsPricingPeriodList struct {
	Data       []SettingsPricingPeriodRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsPricingPeriodList was measured from.
func (SettingsPricingPeriodList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/pricing_periods",
	}
}

// SettingsPricingPeriodRecord is the settings pricing period body, under pricing_period.
type SettingsPricingPeriodRecord struct {
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
	TimePeriod    *SettingsPricingPeriodTimePeriod
}

// MeasuredFrom lists the endpoints SettingsPricingPeriodRecord was measured from.
func (SettingsPricingPeriodRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/pricing_periods",
		"GET /api/{api_version}/settings/pricing_periods/{pricing_period_id}",
		"POST /api/{api_version}/settings/pricing_periods",
		"PATCH /api/{api_version}/settings/pricing_periods/{pricing_period_id}",
	}
}

// EnvelopeKey unwraps {"pricing_period": {...}}.
func (SettingsPricingPeriodRecord) EnvelopeKey() string { return PricingPeriodKey }

// SettingsPricingPeriodTimePeriod is a SettingsPricingPeriodRecord's time_period.
type SettingsPricingPeriodTimePeriod struct {
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

// SettingsPricingPeriodListShape reads a body as SettingsPricingPeriodList.
func SettingsPricingPeriodListShape(decoded any) any {
	return dto.Read[SettingsPricingPeriodList](decoded)
}

// SettingsPricingPeriodRecordShape reads a body as SettingsPricingPeriodRecord.
func SettingsPricingPeriodRecordShape(decoded any) any {
	return dto.Read[SettingsPricingPeriodRecord](decoded)
}
