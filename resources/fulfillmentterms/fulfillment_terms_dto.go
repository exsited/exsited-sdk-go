package fulfillmentterms

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsFulfillmentTermList is one page of SettingsFulfillmentTermRecord records, under data.
type SettingsFulfillmentTermList struct {
	Data       []SettingsFulfillmentTermRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsFulfillmentTermList was measured from.
func (SettingsFulfillmentTermList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/fulfillment_terms",
	}
}

// SettingsFulfillmentTermRecord is the settings fulfillment term body, under fulfillment_term.
type SettingsFulfillmentTermRecord struct {
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

// MeasuredFrom lists the endpoints SettingsFulfillmentTermRecord was measured from.
func (SettingsFulfillmentTermRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/fulfillment_terms",
		"GET /api/{api_version}/settings/fulfillment_terms/{fulfillment_term_id}",
		"POST /api/{api_version}/settings/fulfillment_terms",
		"PATCH /api/{api_version}/settings/fulfillment_terms/{fulfillment_term_id}",
	}
}

// EnvelopeKey unwraps {"fulfillment_term": {...}}.
func (SettingsFulfillmentTermRecord) EnvelopeKey() string { return FulfillmentTermKey }

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

// SettingsFulfillmentTermListShape reads a body as SettingsFulfillmentTermList.
func SettingsFulfillmentTermListShape(decoded any) any {
	return dto.Read[SettingsFulfillmentTermList](decoded)
}

// SettingsFulfillmentTermRecordShape reads a body as SettingsFulfillmentTermRecord.
func SettingsFulfillmentTermRecordShape(decoded any) any {
	return dto.Read[SettingsFulfillmentTermRecord](decoded)
}
