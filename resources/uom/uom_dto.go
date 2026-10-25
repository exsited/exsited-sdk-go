package uom

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsUomList is one page of SettingsUomRecord records, under data.
type SettingsUomList struct {
	Data       []SettingsUomRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsUomList was measured from.
func (SettingsUomList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/uom",
	}
}

// SettingsUomRecord is the settings unit of measurement body, under uom.
type SettingsUomRecord struct {
	CreatedBy       any
	CreatedOn       *string
	CustomType      any
	DecimalPlaces   *int64
	Description     *string
	DisplayName     *string
	Href            *string
	ID              *string
	LastUpdatedBy   any
	LastUpdatedOn   *string
	Name            *string
	Object          *string
	Rounding        *string
	Status          *string
	Type            *string
	TypeDisplayText *string
}

// MeasuredFrom lists the endpoints SettingsUomRecord was measured from.
func (SettingsUomRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/uom",
		"GET /api/{api_version}/settings/uom/{uom_id}",
	}
}

// EnvelopeKey unwraps {"uom": {...}}.
func (SettingsUomRecord) EnvelopeKey() string { return UomKey }

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

// SettingsUomListShape reads a body as SettingsUomList.
func SettingsUomListShape(decoded any) any { return dto.Read[SettingsUomList](decoded) }

// SettingsUomRecordShape reads a body as SettingsUomRecord.
func SettingsUomRecordShape(decoded any) any { return dto.Read[SettingsUomRecord](decoded) }
