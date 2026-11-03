package operatorgroups

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsOperatorGroupList is one page of SettingsOperatorGroupRecord records, under data.
type SettingsOperatorGroupList struct {
	Data       []SettingsOperatorGroupRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsOperatorGroupList is returned by.
func (SettingsOperatorGroupList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/operator_groups",
	}
}

// SettingsOperatorGroupRecord is the operator group body, under operator_group. The create
// answer carries only object, id, name, display_name, description and href. CreatedBy and
// LastUpdatedBy are untyped.
type SettingsOperatorGroupRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Departments   []SettingsOperatorGroupDepartment
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Role          *SettingsOperatorGroupRole
	Status        *string
}

// MeasuredFrom lists the endpoints SettingsOperatorGroupRecord is returned by.
func (SettingsOperatorGroupRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/operator_groups",
		"GET /api/{api_version}/settings/operator_groups/{operator_group_id}",
		"POST /api/{api_version}/settings/operator_groups",
		"PATCH /api/{api_version}/settings/operator_groups/{operator_group_id}",
	}
}

// EnvelopeKey unwraps {"operator_group": {...}}.
func (SettingsOperatorGroupRecord) EnvelopeKey() string { return OperatorGroupKey }

// SettingsOperatorGroupRole is a SettingsOperatorGroupRecord's role. A read names it as a role
// with its rol_ id and href; the update answer names it as a reference with the dashed uuid and
// no href.
type SettingsOperatorGroupRole struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// SettingsOperatorGroupDepartment is one of a SettingsOperatorGroupRecord's departments, a
// reference; ID can be null.
type SettingsOperatorGroupDepartment struct {
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

// SettingsOperatorGroupListShape reads a body as SettingsOperatorGroupList.
func SettingsOperatorGroupListShape(decoded any) any {
	return dto.Read[SettingsOperatorGroupList](decoded)
}

// SettingsOperatorGroupRecordShape reads a body as SettingsOperatorGroupRecord.
func SettingsOperatorGroupRecordShape(decoded any) any {
	return dto.Read[SettingsOperatorGroupRecord](decoded)
}
