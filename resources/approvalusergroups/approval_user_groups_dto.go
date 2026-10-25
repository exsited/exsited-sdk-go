package approvalusergroups

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsApprovalUserGroupList is one page of SettingsApprovalUserGroupRecord records, under data.
type SettingsApprovalUserGroupList struct {
	Data       []SettingsApprovalUserGroupRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsApprovalUserGroupList was measured from.
func (SettingsApprovalUserGroupList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/approval_user_groups",
	}
}

// SettingsApprovalUserGroupRecord is the settings approval user group body, under
// approval_user_group.
type SettingsApprovalUserGroupRecord struct {
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

// MeasuredFrom lists the endpoints SettingsApprovalUserGroupRecord was measured from.
func (SettingsApprovalUserGroupRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/approval_user_groups",
		"GET /api/{api_version}/settings/approval_user_groups/{approval_user_group_id}",
		"POST /api/{api_version}/settings/approval_user_groups",
		"PATCH /api/{api_version}/settings/approval_user_groups/{approval_user_group_id}",
	}
}

// EnvelopeKey unwraps {"approval_user_group": {...}}.
func (SettingsApprovalUserGroupRecord) EnvelopeKey() string { return ApprovalUserGroupKey }

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

// SettingsApprovalUserGroupListShape reads a body as SettingsApprovalUserGroupList.
func SettingsApprovalUserGroupListShape(decoded any) any {
	return dto.Read[SettingsApprovalUserGroupList](decoded)
}

// SettingsApprovalUserGroupRecordShape reads a body as SettingsApprovalUserGroupRecord.
func SettingsApprovalUserGroupRecordShape(decoded any) any {
	return dto.Read[SettingsApprovalUserGroupRecord](decoded)
}
