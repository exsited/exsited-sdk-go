package customforms

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsCustomFormList is one page of SettingsCustomFormRecord records, under data.
type SettingsCustomFormList struct {
	Data       []SettingsCustomFormRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsCustomFormList was measured from.
func (SettingsCustomFormList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/custom_forms",
	}
}

// SettingsCustomFormRecord is the settings custom form body, under custom_form.
type SettingsCustomFormRecord struct {
	CreatedBy          any
	CreatedOn          *string
	Description        *string
	DisplayName        *string
	ForDefaultTemplate *bool
	Href               *string
	ID                 *string
	IsSystemTemplate   *bool
	LastUpdatedBy      any
	LastUpdatedOn      *string
	Name               *string
	Object             *string
	Status             *string
}

// MeasuredFrom lists the endpoints SettingsCustomFormRecord was measured from.
func (SettingsCustomFormRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/custom_forms",
		"GET /api/{api_version}/settings/custom_forms/{custom_form_id}",
		"POST /api/{api_version}/settings/custom_forms",
		"PATCH /api/{api_version}/settings/custom_forms/{custom_form_id}",
	}
}

// EnvelopeKey unwraps {"custom_form": {...}}.
func (SettingsCustomFormRecord) EnvelopeKey() string { return CustomFormKey }

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

// SettingsCustomFormListShape reads a body as SettingsCustomFormList.
func SettingsCustomFormListShape(decoded any) any { return dto.Read[SettingsCustomFormList](decoded) }

// SettingsCustomFormRecordShape reads a body as SettingsCustomFormRecord.
func SettingsCustomFormRecordShape(decoded any) any {
	return dto.Read[SettingsCustomFormRecord](decoded)
}
