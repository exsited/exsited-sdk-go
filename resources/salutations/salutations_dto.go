package salutations

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsSalutationList is one page of SettingsSalutationRecord records, under data.
type SettingsSalutationList struct {
	Data       []SettingsSalutationRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsSalutationList was measured from.
func (SettingsSalutationList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/salutations",
	}
}

// SettingsSalutationRecord is the settings salutation body, under salutation.
type SettingsSalutationRecord struct {
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
}

// MeasuredFrom lists the endpoints SettingsSalutationRecord was measured from.
func (SettingsSalutationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/salutations",
		"GET /api/{api_version}/settings/salutations/{salutation_id}",
		"POST /api/{api_version}/settings/salutations",
		"PATCH /api/{api_version}/settings/salutations/{salutation_id}",
	}
}

// EnvelopeKey unwraps {"salutation": {...}}.
func (SettingsSalutationRecord) EnvelopeKey() string { return SalutationKey }

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

// SettingsSalutationListShape reads a body as SettingsSalutationList.
func SettingsSalutationListShape(decoded any) any { return dto.Read[SettingsSalutationList](decoded) }

// SettingsSalutationRecordShape reads a body as SettingsSalutationRecord.
func SettingsSalutationRecordShape(decoded any) any {
	return dto.Read[SettingsSalutationRecord](decoded)
}
