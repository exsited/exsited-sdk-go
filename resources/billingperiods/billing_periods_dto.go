package billingperiods

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsBillingPeriodList is one page of SettingsBillingPeriodRecord records, under data.
type SettingsBillingPeriodList struct {
	Data       []SettingsBillingPeriodRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsBillingPeriodList was measured from.
func (SettingsBillingPeriodList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/billing_periods",
	}
}

// SettingsBillingPeriodRecord is the settings billing period body, under billing_period.
type SettingsBillingPeriodRecord struct {
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
	Ordering      *int64
	Period        *string
	Status        *string
	TimePeriod    *string
}

// MeasuredFrom lists the endpoints SettingsBillingPeriodRecord was measured from.
func (SettingsBillingPeriodRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/billing_periods",
		"GET /api/{api_version}/settings/billing_periods/{billing_period_id}",
		"POST /api/{api_version}/settings/billing_periods",
		"PATCH /api/{api_version}/settings/billing_periods/{billing_period_id}",
	}
}

// EnvelopeKey unwraps {"billing_period": {...}}.
func (SettingsBillingPeriodRecord) EnvelopeKey() string { return BillingPeriodKey }

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

// SettingsBillingPeriodListShape reads a body as SettingsBillingPeriodList.
func SettingsBillingPeriodListShape(decoded any) any {
	return dto.Read[SettingsBillingPeriodList](decoded)
}

// SettingsBillingPeriodRecordShape reads a body as SettingsBillingPeriodRecord.
func SettingsBillingPeriodRecordShape(decoded any) any {
	return dto.Read[SettingsBillingPeriodRecord](decoded)
}
