package notegroups

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

// NoteGroupRecord is the note group body.
type NoteGroupRecord struct {
	ActivityLogs  []any
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
	Version       *int64
}

// MeasuredFrom lists the endpoints NoteGroupRecord was measured from.
func (NoteGroupRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/note_groups",
		"GET /api/{api_version}/note_groups/{note_group_id}",
		"POST /api/{api_version}/note_groups",
		"PATCH /api/{api_version}/note_groups/{note_group_id}",
	}
}

// EnvelopeKey unwraps {"note_group": {...}}.
func (NoteGroupRecord) EnvelopeKey() string { return NoteGroupKey }

// NoteGroupList is a page of NoteGroupRecord records.
type NoteGroupList struct {
	Data       []NoteGroupRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints NoteGroupList was measured from.
func (NoteGroupList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/note_groups",
	}
}

// NoteGroupNotesRecord is the note group notes body.
type NoteGroupNotesRecord struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints NoteGroupNotesRecord was measured from.
func (NoteGroupNotesRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/note_groups/{note_group_id}/notes",
	}
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// NoteGroupListShape reads a body as NoteGroupList.
func NoteGroupListShape(decoded any) any { return dto.Read[NoteGroupList](decoded) }

// NoteGroupRecordShape reads a body as NoteGroupRecord.
func NoteGroupRecordShape(decoded any) any { return dto.Read[NoteGroupRecord](decoded) }

// NoteGroupNotesRecordShape reads a body as NoteGroupNotesRecord.
func NoteGroupNotesRecordShape(decoded any) any { return dto.Read[NoteGroupNotesRecord](decoded) }
