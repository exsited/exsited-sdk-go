package teams

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsTeamList is one page of SettingsTeamRecord records, under data.
type SettingsTeamList struct {
	Data       []SettingsTeamRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsTeamList was measured from.
func (SettingsTeamList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/teams",
	}
}

// SettingsTeamRecord is the settings team body, under team.
type SettingsTeamRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	ManagedBy     any
	Name          *string
	Object        *string
	Status        *string
	Version       any
}

// MeasuredFrom lists the endpoints SettingsTeamRecord was measured from.
func (SettingsTeamRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/teams",
		"GET /api/{api_version}/settings/teams/{team_id}",
		"POST /api/{api_version}/settings/teams",
		"PATCH /api/{api_version}/settings/teams/{team_id}",
	}
}

// EnvelopeKey unwraps {"team": {...}}.
func (SettingsTeamRecord) EnvelopeKey() string { return TeamKey }

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

// SettingsTeamListShape reads a body as SettingsTeamList.
func SettingsTeamListShape(decoded any) any { return dto.Read[SettingsTeamList](decoded) }

// SettingsTeamRecordShape reads a body as SettingsTeamRecord.
func SettingsTeamRecordShape(decoded any) any { return dto.Read[SettingsTeamRecord](decoded) }
