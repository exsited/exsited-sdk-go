package timesheets

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

// TimesheetRecord is the timesheet body.
type TimesheetRecord struct {
	CreatedBy     *CreatedBy
	CreatedOn     *string
	EndTime       *string
	Href          *string
	ID            *string
	Labour        *Labour
	LastUpdatedBy *LastUpdatedBy
	LastUpdatedOn any
	Object        *string
	SaleOrder     *SaleOrder
	SaleOrderItem *SaleOrderItem
	StartTime     *string
	Status        *string
	Version       *int64
}

// MeasuredFrom lists the endpoints TimesheetRecord was measured from.
func (TimesheetRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/timesheets",
	}
}

// TimesheetList is a page of TimesheetRecord records.
type TimesheetList struct {
	Data       []TimesheetRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints TimesheetList was measured from.
func (TimesheetList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/timesheets",
	}
}

// TimesheetCreateRecord is the body CreateTimesheet answers.
type TimesheetCreateRecord struct {
	ActivityLogs  []ActivityLog
	CreatedBy     *CreatedBy
	CreatedOn     *string
	EndTime       *string
	Href          *string
	ID            *string
	Labour        *Labour
	LastUpdatedBy *LastUpdatedBy
	LastUpdatedOn *string
	Object        *string
	SaleOrder     *SaleOrder
	SaleOrderItem *SaleOrderItem
	StartTime     *string
	Status        *string
	Version       *int64
}

// MeasuredFrom lists the endpoints TimesheetCreateRecord was measured from.
func (TimesheetCreateRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/timesheets",
		"PATCH /api/{api_version}/timesheets/{timesheet_id}",
	}
}

// EnvelopeKey unwraps {"timesheet": {...}}.
func (TimesheetCreateRecord) EnvelopeKey() string { return TimesheetKey }

// ActivityLog is one entry of the activity_logs list.
type ActivityLog struct {
	Activity *string
}

// CreatedBy is the created_by object.
type CreatedBy struct {
	Name   *string
	Object *string
}

// Labour is the labour object.
type Labour struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// SaleOrder is the sale_order object.
type SaleOrder struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// SaleOrderItem is the sale_order_item object.
type SaleOrderItem struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// TimesheetListShape reads a body as TimesheetList.
func TimesheetListShape(decoded any) any { return dto.Read[TimesheetList](decoded) }

// TimesheetCreateRecordShape reads a body as TimesheetCreateRecord.
func TimesheetCreateRecordShape(decoded any) any { return dto.Read[TimesheetCreateRecord](decoded) }
