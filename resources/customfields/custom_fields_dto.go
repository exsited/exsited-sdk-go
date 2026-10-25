package customfields

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsCustomFieldList is one page of SettingsCustomFieldRecord records, under data.
type SettingsCustomFieldList struct {
	Data       []SettingsCustomFieldRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsCustomFieldList was measured from.
func (SettingsCustomFieldList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/custom_fields",
	}
}

// SettingsCustomFieldRecord is the settings custom field body, under custom_field.
type SettingsCustomFieldRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	IsMandatory   *bool
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Status        *string
	Type          *string
	Version       *int64
}

// MeasuredFrom lists the endpoints SettingsCustomFieldRecord was measured from.
func (SettingsCustomFieldRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/custom_fields",
		"GET /api/{api_version}/settings/custom_fields/{custom_field_id}",
	}
}

// EnvelopeKey unwraps {"custom_field": {...}}.
func (SettingsCustomFieldRecord) EnvelopeKey() string { return CustomFieldKey }

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

// SettingsCustomFieldListShape reads a body as SettingsCustomFieldList.
func SettingsCustomFieldListShape(decoded any) any { return dto.Read[SettingsCustomFieldList](decoded) }

// SettingsCustomFieldRecordShape reads a body as SettingsCustomFieldRecord.
func SettingsCustomFieldRecordShape(decoded any) any {
	return dto.Read[SettingsCustomFieldRecord](decoded)
}
