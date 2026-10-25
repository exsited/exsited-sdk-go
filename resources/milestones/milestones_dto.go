package milestones

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

// MilestonesRecord is the milestones body.
type MilestonesRecord struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints MilestonesRecord was measured from.
func (MilestonesRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/milestones",
		"GET /api/{api_version}/milestones/account/{account_id}",
	}
}

// MilestonesLabourByIDList is the page GetMilestonesLabourByID answers.
type MilestonesLabourByIDList struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints MilestonesLabourByIDList was measured from.
func (MilestonesLabourByIDList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/milestones/labour/{labour_id}",
	}
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// MilestonesRecordShape reads a body as MilestonesRecord.
func MilestonesRecordShape(decoded any) any { return dto.Read[MilestonesRecord](decoded) }

// MilestonesLabourByIDListShape reads a body as MilestonesLabourByIDList.
func MilestonesLabourByIDListShape(decoded any) any {
	return dto.Read[MilestonesLabourByIDList](decoded)
}
