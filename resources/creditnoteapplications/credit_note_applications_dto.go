package creditnoteapplications

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

// CreditNoteApplicationRecord is the credit note application body.
type CreditNoteApplicationRecord struct {
	// Amount is untyped: the API sends it as a string or a number.
	Amount any
	// CreatedBy is untyped: the API sends it as a string or an object.
	CreatedBy     *CreatedBy
	CreatedOn     *string
	CreditNote    *CreditNote
	CreditNoteID  *string
	Date          *string
	Href          *string
	ID            *string
	IssueDate     *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Object        *string
	Payment       *Payment
	PaymentID     *string
	Primaryid     *int64 `wire:"primaryId"`
	Refund        *Refund
	RefundID      *string
	// RemainingBalance is untyped: the API sends it as a string or a number.
	RemainingBalance any
	Status           *string
	UUID             *string
	// Version is untyped: the API sends it as a string or an integer.
	Version any
}

// MeasuredFrom lists the endpoints CreditNoteApplicationRecord was measured from.
func (CreditNoteApplicationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/credit_note_applications",
	}
}

// CreditNoteApplicationList is a page of CreditNoteApplicationRecord records.
type CreditNoteApplicationList struct {
	Data       []CreditNoteApplicationRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints CreditNoteApplicationList was measured from.
func (CreditNoteApplicationList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/credit_note_applications",
	}
}

// CreditNoteApplicationReadRecord is the credit note application read body.
type CreditNoteApplicationReadRecord struct {
	Amount           *float64
	CreatedBy        *CreatedBy
	CreatedOn        *string
	CreditNote       *CreditNote
	Href             *string
	ID               *string
	IssueDate        *string
	LastUpdatedBy    any
	LastUpdatedOn    *string
	Object           *string
	Payment          *Payment
	Primaryid        *int64 `wire:"primaryId"`
	Refund           *Refund
	RemainingBalance *float64
	Status           *string
	Version          *int64
}

// MeasuredFrom lists the endpoints CreditNoteApplicationReadRecord was measured from.
func (CreditNoteApplicationReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/credit_note_applications/{credit_note_application_id}",
	}
}

// EnvelopeKey unwraps {"credit_note_application": {...}}.
func (CreditNoteApplicationReadRecord) EnvelopeKey() string { return CreditNoteApplicationKey }

// InvoiceCreditNoteApplicationsRecord is the invoice credit note applications body.
type InvoiceCreditNoteApplicationsRecord struct {
	Data       []CreditNoteApplicationRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints InvoiceCreditNoteApplicationsRecord was measured from.
func (InvoiceCreditNoteApplicationsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/credit_note_applications",
	}
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// CreatedBy is the created_by object.
type CreatedBy struct {
	ID     *string
	Name   *string
	Object *string
}

// CreditNote is the credit_note object.
type CreditNote struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Payment is the payment object.
type Payment struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Refund is the refund object.
type Refund struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// CreditNoteApplicationListShape reads a body as CreditNoteApplicationList.
func CreditNoteApplicationListShape(decoded any) any {
	return dto.Read[CreditNoteApplicationList](decoded)
}

// CreditNoteApplicationReadRecordShape reads a body as CreditNoteApplicationReadRecord.
func CreditNoteApplicationReadRecordShape(decoded any) any {
	return dto.Read[CreditNoteApplicationReadRecord](decoded)
}

// InvoiceCreditNoteApplicationsRecordShape reads a body as InvoiceCreditNoteApplicationsRecord.
func InvoiceCreditNoteApplicationsRecordShape(decoded any) any {
	return dto.Read[InvoiceCreditNoteApplicationsRecord](decoded)
}
