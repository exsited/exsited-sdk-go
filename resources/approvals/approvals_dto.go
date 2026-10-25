package approvals

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

// ApprovalsAccountRecord is the approvals account body.
type ApprovalsAccountRecord struct {
	CreatedOn     *string
	Description   any
	DisplayName   *string
	Entity        *string
	Event         *string
	Href          *string
	ID            *string
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Status        *string
}

// MeasuredFrom lists the endpoints ApprovalsAccountRecord was measured from.
func (ApprovalsAccountRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/approvals/account",
	}
}

// ApprovalsAccountList is a page of ApprovalsAccountRecord records.
type ApprovalsAccountList struct {
	Data       []ApprovalsAccountRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ApprovalsAccountList was measured from.
func (ApprovalsAccountList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/approvals/account",
	}
}

// ApprovalsSaleorderRecord is the approvals saleorder body.
type ApprovalsSaleorderRecord struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ApprovalsSaleorderRecord was measured from.
func (ApprovalsSaleorderRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/approvals/saleOrder",
	}
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// ApprovalsAccountListShape reads a body as ApprovalsAccountList.
func ApprovalsAccountListShape(decoded any) any { return dto.Read[ApprovalsAccountList](decoded) }

// ApprovalsSaleorderRecordShape reads a body as ApprovalsSaleorderRecord.
func ApprovalsSaleorderRecordShape(decoded any) any {
	return dto.Read[ApprovalsSaleorderRecord](decoded)
}
