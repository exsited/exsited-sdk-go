package additionalcodes

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsAdditionalCodeList is one page of SettingsAdditionalCodeRecord records, under data.
type SettingsAdditionalCodeList struct {
	Data       []SettingsAdditionalCodeRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsAdditionalCodeList was measured from.
func (SettingsAdditionalCodeList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/additional_codes",
	}
}

// SettingsAdditionalCodeRecord is the settings additional code body, under additional_code.
type SettingsAdditionalCodeRecord struct {
	Description *string
	DisplayName *string
	Href        *string
	ID          *string
	Name        *string
	Object      *string
	Status      *string
}

// MeasuredFrom lists the endpoints SettingsAdditionalCodeRecord was measured from.
func (SettingsAdditionalCodeRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/additional_codes",
		"GET /api/{api_version}/settings/additional_codes/{additional_code_id}",
		"POST /api/{api_version}/settings/additional_codes",
	}
}

// EnvelopeKey unwraps {"additional_code": {...}}.
func (SettingsAdditionalCodeRecord) EnvelopeKey() string { return AdditionalCodeKey }

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

// SettingsAdditionalCodeListShape reads a body as SettingsAdditionalCodeList.
func SettingsAdditionalCodeListShape(decoded any) any {
	return dto.Read[SettingsAdditionalCodeList](decoded)
}

// SettingsAdditionalCodeRecordShape reads a body as SettingsAdditionalCodeRecord.
func SettingsAdditionalCodeRecordShape(decoded any) any {
	return dto.Read[SettingsAdditionalCodeRecord](decoded)
}
