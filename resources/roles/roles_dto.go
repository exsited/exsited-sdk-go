package roles

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsRoleList is one page of SettingsRoleRecord records, under data.
type SettingsRoleList struct {
	Data       []SettingsRoleRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsRoleList was measured from.
func (SettingsRoleList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/roles",
	}
}

// SettingsRoleRecord is the settings role body, under role.
type SettingsRoleRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayName   *string
	Href          *string
	ID            *string
	IsSystemRole  *bool
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Status        *string
}

// MeasuredFrom lists the endpoints SettingsRoleRecord was measured from.
func (SettingsRoleRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/roles",
		"GET /api/{api_version}/settings/roles/{role_id}",
	}
}

// EnvelopeKey unwraps {"role": {...}}.
func (SettingsRoleRecord) EnvelopeKey() string { return RoleKey }

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

// SettingsRoleListShape reads a body as SettingsRoleList.
func SettingsRoleListShape(decoded any) any { return dto.Read[SettingsRoleList](decoded) }

// SettingsRoleRecordShape reads a body as SettingsRoleRecord.
func SettingsRoleRecordShape(decoded any) any { return dto.Read[SettingsRoleRecord](decoded) }
