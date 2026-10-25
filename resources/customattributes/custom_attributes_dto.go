package customattributes

import "github.com/exsited/exsited-sdk-go/v4/dto"

// Pagination is the paging block, either shape.
type Pagination struct {
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	PrevCursor any
	Links      *PaginationLinks

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	NextPage     *string
	PreviousPage *string
}

// CustomAttributeRecord is the custom attribute body.
type CustomAttributeRecord struct {
	CreatedBy      any
	CreatedOn      *string
	Description    any
	DisplayName    *string
	EncryptData    *bool
	Href           *string
	ID             *string
	IncludeSeconds any
	IsMandatory    *bool
	LastUpdatedBy  any
	LastUpdatedOn  *string
	MaxLength      any
	MaxSize        any
	MaxValue       any
	MetaData       any
	MinValue       any
	Name           *string
	Object         *string
	Options        []Option
	Status         *string
	TimeFormat     any
	Type           *string
	Unit           any
	Uom            any
	UseIn          []UseIn
	Version        *int64
}

// MeasuredFrom lists the endpoints CustomAttributeRecord was measured from.
func (CustomAttributeRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/custom_attributes",
		"GET /api/{api_version}/custom_attributes/{custom_attribute_id}",
		"PATCH /api/{api_version}/custom_attributes/{custom_attribute_id}",
	}
}

// EnvelopeKey unwraps {"custom_attribute": {...}}.
func (CustomAttributeRecord) EnvelopeKey() string { return CustomAttributeKey }

// CustomAttributeList is a page of CustomAttributeRecord records.
type CustomAttributeList struct {
	Data       []CustomAttributeRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints CustomAttributeList was measured from.
func (CustomAttributeList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/custom_attributes",
	}
}

// CustomAttributesRecord is the custom attributes body.
type CustomAttributesRecord struct {
	ActivityLogs   []ActivityLog
	CreatedBy      any
	CreatedOn      *string
	Description    *string
	DisplayName    *string
	EncryptData    *bool
	Href           *string
	ID             *string
	IncludeSeconds any
	IsMandatory    *bool
	LastUpdatedBy  any
	LastUpdatedOn  *string
	MaxLength      *int64
	MaxSize        any
	MaxValue       any
	MetaData       *string
	MinValue       any
	Name           *string
	Object         *string
	Options        []Option
	Status         *string
	TimeFormat     any
	Type           *string
	Unit           any
	Uom            any
	UseIn          []UseIn
	UUID           *string
	Version        *int64
}

// MeasuredFrom lists the endpoints CustomAttributesRecord was measured from.
func (CustomAttributesRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/custom_attributes",
	}
}

// EnvelopeKey unwraps {"custom_attribute": {...}}.
func (CustomAttributesRecord) EnvelopeKey() string { return CustomAttributeKey }

// ActivityLog is one entry of the activity_logs list.
type ActivityLog struct {
	Activity *string
}

// Option is one entry of the options list.
type Option struct {
	DisplayOrder *int64
	Name         *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// UseIn is one entry of the use_in list.
type UseIn struct {
	AssociatedAccountGroups []any
	AssociatedItemGroups    []any
	AssociatedUserGroups    []any
	Enabled                 *bool
	Required                *bool
	Resource                *string
	Unique                  *bool
}

// CustomAttributeListShape reads a body as CustomAttributeList.
func CustomAttributeListShape(decoded any) any { return dto.Read[CustomAttributeList](decoded) }

// CustomAttributesRecordShape reads a body as CustomAttributesRecord.
func CustomAttributesRecordShape(decoded any) any { return dto.Read[CustomAttributesRecord](decoded) }

// CustomAttributeRecordShape reads a body as CustomAttributeRecord.
func CustomAttributeRecordShape(decoded any) any { return dto.Read[CustomAttributeRecord](decoded) }
