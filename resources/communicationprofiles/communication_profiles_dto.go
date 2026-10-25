package communicationprofiles

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsCommunicationProfileList is one page of SettingsCommunicationProfileRecord records, under
// data.
type SettingsCommunicationProfileList struct {
	Data       []SettingsCommunicationProfileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsCommunicationProfileList was measured from.
func (SettingsCommunicationProfileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/communication_profiles",
	}
}

// SettingsCommunicationProfileRecord is the settings communication profile body, under
// communication_profile.
type SettingsCommunicationProfileRecord struct {
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

// MeasuredFrom lists the endpoints SettingsCommunicationProfileRecord was measured from.
func (SettingsCommunicationProfileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/communication_profiles",
		"GET /api/{api_version}/settings/communication_profiles/{communication_profile_id}",
	}
}

// EnvelopeKey unwraps {"communication_profile": {...}}.
func (SettingsCommunicationProfileRecord) EnvelopeKey() string { return CommunicationProfileKey }

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

// SettingsCommunicationProfileListShape reads a body as SettingsCommunicationProfileList.
func SettingsCommunicationProfileListShape(decoded any) any {
	return dto.Read[SettingsCommunicationProfileList](decoded)
}

// SettingsCommunicationProfileRecordShape reads a body as SettingsCommunicationProfileRecord.
func SettingsCommunicationProfileRecordShape(decoded any) any {
	return dto.Read[SettingsCommunicationProfileRecord](decoded)
}
