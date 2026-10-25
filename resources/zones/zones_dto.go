package zones

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsZoneList is one page of SettingsZoneRecord records, under data.
type SettingsZoneList struct {
	Data       []SettingsZoneRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsZoneList was measured from.
func (SettingsZoneList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/zones",
	}
}

// SettingsZoneRecord is the settings zone body, under zone.
type SettingsZoneRecord struct {
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

// MeasuredFrom lists the endpoints SettingsZoneRecord was measured from.
func (SettingsZoneRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/zones",
		"GET /api/{api_version}/settings/zones/{zone_id}",
		"POST /api/{api_version}/settings/zones",
		"PATCH /api/{api_version}/settings/zones/{zone_id}",
	}
}

// EnvelopeKey unwraps {"zone": {...}}.
func (SettingsZoneRecord) EnvelopeKey() string { return ZoneKey }

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

// SettingsZoneListShape reads a body as SettingsZoneList.
func SettingsZoneListShape(decoded any) any { return dto.Read[SettingsZoneList](decoded) }

// SettingsZoneRecordShape reads a body as SettingsZoneRecord.
func SettingsZoneRecordShape(decoded any) any { return dto.Read[SettingsZoneRecord](decoded) }
