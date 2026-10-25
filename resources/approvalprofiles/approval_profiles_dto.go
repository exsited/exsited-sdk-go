package approvalprofiles

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsApprovalProfileList is one page of SettingsApprovalProfileRecord records, under data.
type SettingsApprovalProfileList struct {
	Data       []SettingsApprovalProfileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsApprovalProfileList was measured from.
func (SettingsApprovalProfileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/approval_profiles",
	}
}

// SettingsApprovalProfileRecord is the settings approval profile body, under approval_profile.
type SettingsApprovalProfileRecord struct {
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

// MeasuredFrom lists the endpoints SettingsApprovalProfileRecord was measured from.
func (SettingsApprovalProfileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/approval_profiles",
		"GET /api/{api_version}/settings/approval_profiles/{approval_profile_id}",
	}
}

// EnvelopeKey unwraps {"approval_profile": {...}}.
func (SettingsApprovalProfileRecord) EnvelopeKey() string { return ApprovalProfileKey }

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

// SettingsApprovalProfileListShape reads a body as SettingsApprovalProfileList.
func SettingsApprovalProfileListShape(decoded any) any {
	return dto.Read[SettingsApprovalProfileList](decoded)
}

// SettingsApprovalProfileRecordShape reads a body as SettingsApprovalProfileRecord.
func SettingsApprovalProfileRecordShape(decoded any) any {
	return dto.Read[SettingsApprovalProfileRecord](decoded)
}
