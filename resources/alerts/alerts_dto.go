package alerts

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

// AlertsRecord is the alerts body.
type AlertsRecord struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AlertsRecord was measured from.
func (AlertsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/alerts",
		"GET /api/{api_version}/alerts/account/{account_id}",
	}
}

// AlertsEmployeeByIDList is the page GetAlertsEmployeeByID answers.
type AlertsEmployeeByIDList struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AlertsEmployeeByIDList was measured from.
func (AlertsEmployeeByIDList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/alerts/employee/{employee_id}",
	}
}

// AlertsLabourByIDList is the page GetAlertsLabourByID answers.
type AlertsLabourByIDList struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AlertsLabourByIDList was measured from.
func (AlertsLabourByIDList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/alerts/labour/{labour_id}",
	}
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// AlertsRecordShape reads a body as AlertsRecord.
func AlertsRecordShape(decoded any) any { return dto.Read[AlertsRecord](decoded) }

// AlertsEmployeeByIDListShape reads a body as AlertsEmployeeByIDList.
func AlertsEmployeeByIDListShape(decoded any) any { return dto.Read[AlertsEmployeeByIDList](decoded) }

// AlertsLabourByIDListShape reads a body as AlertsLabourByIDList.
func AlertsLabourByIDListShape(decoded any) any { return dto.Read[AlertsLabourByIDList](decoded) }
