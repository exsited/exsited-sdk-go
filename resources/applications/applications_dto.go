package applications

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsApplicationList is one page of SettingsApplicationRecord records, under data.
type SettingsApplicationList struct {
	Data       []SettingsApplicationRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsApplicationList was measured from.
func (SettingsApplicationList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/applications",
	}
}

// SettingsApplicationRecord is the settings application body, under application.
type SettingsApplicationRecord struct {
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

// MeasuredFrom lists the endpoints SettingsApplicationRecord was measured from.
func (SettingsApplicationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/applications",
		"GET /api/{api_version}/settings/applications/{application_id}",
		"POST /api/{api_version}/settings/applications",
		"PATCH /api/{api_version}/settings/applications/{application_id}",
	}
}

// EnvelopeKey unwraps {"application": {...}}.
func (SettingsApplicationRecord) EnvelopeKey() string { return ApplicationKey }

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

// SettingsApplicationListShape reads a body as SettingsApplicationList.
func SettingsApplicationListShape(decoded any) any { return dto.Read[SettingsApplicationList](decoded) }

// SettingsApplicationRecordShape reads a body as SettingsApplicationRecord.
func SettingsApplicationRecordShape(decoded any) any {
	return dto.Read[SettingsApplicationRecord](decoded)
}
