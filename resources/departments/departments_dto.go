package departments

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsDepartmentList is one page of SettingsDepartmentRecord records, under data.
type SettingsDepartmentList struct {
	Data       []SettingsDepartmentRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsDepartmentList was measured from.
func (SettingsDepartmentList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/departments",
	}
}

// SettingsDepartmentRecord is the settings department body, under department.
type SettingsDepartmentRecord struct {
	Code          *string
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Division      *SettingsDepartmentDivision
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Status        *string
	Version       *int64
}

// MeasuredFrom lists the endpoints SettingsDepartmentRecord was measured from.
func (SettingsDepartmentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/departments",
		"GET /api/{api_version}/settings/departments/{department_id}",
	}
}

// EnvelopeKey unwraps {"department": {...}}.
func (SettingsDepartmentRecord) EnvelopeKey() string { return DepartmentKey }

// SettingsDepartmentDivision is a SettingsDepartmentRecord's division.
type SettingsDepartmentDivision struct {
	ID     *string
	Name   *string
	Object *string
}

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

// SettingsDepartmentListShape reads a body as SettingsDepartmentList.
func SettingsDepartmentListShape(decoded any) any { return dto.Read[SettingsDepartmentList](decoded) }

// SettingsDepartmentRecordShape reads a body as SettingsDepartmentRecord.
func SettingsDepartmentRecordShape(decoded any) any {
	return dto.Read[SettingsDepartmentRecord](decoded)
}
