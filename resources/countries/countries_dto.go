package countries

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsCountryList is one page of SettingsCountryRecord records, under data.
type SettingsCountryList struct {
	Data       []SettingsCountryRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsCountryList was measured from.
func (SettingsCountryList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/countries",
	}
}

// SettingsCountryRecord is the settings country body, under country.
type SettingsCountryRecord struct {
	Code            *string
	CreatedBy       any
	CreatedOn       *string
	DefaultAreaCode any
	DefaultCity     *string
	DefaultPostCode *string
	DefaultState    *string
	Description     *string
	DisplayName     *string
	FlagExists      *bool
	Href            *string
	ID              *string
	LastUpdatedBy   any
	LastUpdatedOn   *string
	Name            *string
	Object          *string
	Status          *string
}

// MeasuredFrom lists the endpoints SettingsCountryRecord was measured from.
func (SettingsCountryRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/countries",
		"GET /api/{api_version}/settings/countries/{country_id}",
		"PATCH /api/{api_version}/settings/countries/{country_id}",
	}
}

// EnvelopeKey unwraps {"country": {...}}.
func (SettingsCountryRecord) EnvelopeKey() string { return CountryKey }

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

// SettingsCountryListShape reads a body as SettingsCountryList.
func SettingsCountryListShape(decoded any) any { return dto.Read[SettingsCountryList](decoded) }

// SettingsCountryRecordShape reads a body as SettingsCountryRecord.
func SettingsCountryRecordShape(decoded any) any { return dto.Read[SettingsCountryRecord](decoded) }
