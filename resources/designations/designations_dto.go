package designations

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsDesignationList is one page of SettingsDesignationRecord records, under data.
type SettingsDesignationList struct {
	Data       []SettingsDesignationRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsDesignationList was measured from.
func (SettingsDesignationList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/designations",
	}
}

// SettingsDesignationRecord is the settings designation body, under designation.
type SettingsDesignationRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Status        *string
}

// MeasuredFrom lists the endpoints SettingsDesignationRecord was measured from.
func (SettingsDesignationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/designations",
		"GET /api/{api_version}/settings/designations/{designation_id}",
	}
}

// EnvelopeKey unwraps {"designation": {...}}.
func (SettingsDesignationRecord) EnvelopeKey() string { return DesignationKey }

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

// SettingsDesignationListShape reads a body as SettingsDesignationList.
func SettingsDesignationListShape(decoded any) any { return dto.Read[SettingsDesignationList](decoded) }

// SettingsDesignationRecordShape reads a body as SettingsDesignationRecord.
func SettingsDesignationRecordShape(decoded any) any {
	return dto.Read[SettingsDesignationRecord](decoded)
}
