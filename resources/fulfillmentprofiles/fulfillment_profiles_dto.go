package fulfillmentprofiles

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsFulfillmentProfileList is one page of SettingsFulfillmentProfileRecord records, under
// data.
type SettingsFulfillmentProfileList struct {
	Data       []SettingsFulfillmentProfileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsFulfillmentProfileList was measured from.
func (SettingsFulfillmentProfileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/fulfillment_profiles",
	}
}

// SettingsFulfillmentProfileRecord is the settings fulfillment profile body, under
// fulfillment_profile.
type SettingsFulfillmentProfileRecord struct {
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
	Version       *int64
}

// MeasuredFrom lists the endpoints SettingsFulfillmentProfileRecord was measured from.
func (SettingsFulfillmentProfileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/fulfillment_profiles",
		"GET /api/{api_version}/settings/fulfillment_profiles/{fulfillment_profile_id}",
		"POST /api/{api_version}/settings/fulfillment_profiles",
		"PATCH /api/{api_version}/settings/fulfillment_profiles/{fulfillment_profile_id}",
	}
}

// EnvelopeKey unwraps {"fulfillment_profile": {...}}.
func (SettingsFulfillmentProfileRecord) EnvelopeKey() string { return FulfillmentProfileKey }

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

// SettingsFulfillmentProfileListShape reads a body as SettingsFulfillmentProfileList.
func SettingsFulfillmentProfileListShape(decoded any) any {
	return dto.Read[SettingsFulfillmentProfileList](decoded)
}

// SettingsFulfillmentProfileRecordShape reads a body as SettingsFulfillmentProfileRecord.
func SettingsFulfillmentProfileRecordShape(decoded any) any {
	return dto.Read[SettingsFulfillmentProfileRecord](decoded)
}
