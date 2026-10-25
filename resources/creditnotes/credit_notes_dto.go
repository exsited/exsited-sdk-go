package creditnotes

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

// CreditNoteRecord is the credit note body.
type CreditNoteRecord struct {
	Account   *Account
	AccountID *string
	Amount    *string
	Code      *string
	// CreatedBy is untyped: the API sends it as a string or an object.
	CreatedBy        *CreatedBy
	CreatedOn        *string
	Currency         *Currency
	CustomAttributes []CustomAttribute
	CustomObjects    []any
	Date             *string
	Href             *string
	ID               *string
	InitialAmount    *string
	InvoiceID        *string
	Invoices         []Invoice
	IssueDate        *string
	LastUpdatedBy    *LastUpdatedBy
	LastUpdatedOn    *string
	Object           *string
	Payment          *Payment
	PaymentID        *string
	RefundAmount     *string
	// Refundable is untyped: the API sends it as a string or a boolean.
	Refundable       any
	RemainingAmount  *string
	RemainingBalance *string
	Status           *string
	UpdatedBy        *string
	UpdatedOn        *string
	UUID             *string
	// Version is untyped: the API sends it as a string or an integer.
	Version any
}

// MeasuredFrom lists the endpoints CreditNoteRecord was measured from.
func (CreditNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/credit_notes",
	}
}

// CreditNoteList is a page of CreditNoteRecord records.
type CreditNoteList struct {
	Data       []CreditNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints CreditNoteList was measured from.
func (CreditNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/credit_notes",
	}
}

// CreditNoteReadRecord is the credit note read body.
type CreditNoteReadRecord struct {
	Account          *Account
	Code             *string
	CreatedBy        *CreatedBy
	CreatedOn        *string
	Currency         *Currency
	CustomAttributes []CustomAttribute
	Href             *string
	ID               *string
	InitialAmount    *string
	Invoices         []Invoice
	IssueDate        *string
	LastUpdatedBy    *LastUpdatedBy
	LastUpdatedOn    *string
	Object           *string
	Payment          *Payment
	RefundAmount     *string
	Refundable       *bool
	RemainingAmount  *string
	Status           *string
	Version          *int64
}

// MeasuredFrom lists the endpoints CreditNoteReadRecord was measured from.
func (CreditNoteReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/credit_notes/{credit_note_id}",
	}
}

// EnvelopeKey unwraps {"credit_note": {...}}.
func (CreditNoteReadRecord) EnvelopeKey() string { return CreditNoteKey }

// CreditNoteUpdateRecord is the body UpdateCreditNote answers.
type CreditNoteUpdateRecord struct {
	Code       *string
	Href       *string
	ID         *string
	IssueDate  *string
	Object     *string
	Refundable *bool
	Status     *string
	Version    *int64
}

// MeasuredFrom lists the endpoints CreditNoteUpdateRecord was measured from.
func (CreditNoteUpdateRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/credit_notes/{credit_note_id}",
	}
}

// EnvelopeKey unwraps {"credit_note": {...}}.
func (CreditNoteUpdateRecord) EnvelopeKey() string { return CreditNoteKey }

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// Account is the account object.
type Account struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// CreatedBy is the created_by object.
type CreatedBy struct {
	Name   *string
	Object *string
}

// Currency is the currency object.
type Currency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// CustomAttribute is one entry of the custom_attributes list.
type CustomAttribute struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
	Type   *string
	Value  any
}

// Invoice is one entry of the invoices list.
type Invoice struct {
	Href   *string
	ID     *string
	Object *string
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// Payment is the payment object.
type Payment struct {
	Href   *string
	ID     *string
	Object *string
}

// CreditNoteListShape reads a body as CreditNoteList.
func CreditNoteListShape(decoded any) any { return dto.Read[CreditNoteList](decoded) }

// CreditNoteReadRecordShape reads a body as CreditNoteReadRecord.
func CreditNoteReadRecordShape(decoded any) any { return dto.Read[CreditNoteReadRecord](decoded) }

// CreditNoteUpdateRecordShape reads a body as CreditNoteUpdateRecord.
func CreditNoteUpdateRecordShape(decoded any) any { return dto.Read[CreditNoteUpdateRecord](decoded) }

// CreditNoteItemsRecord is A credit note's balance rows, under credit_note.
type CreditNoteItemsRecord struct {
	Href   *string
	Items  []CreditNoteItem
	Object *string
}

// MeasuredFrom lists the endpoints CreditNoteItemsRecord was measured from.
func (CreditNoteItemsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/credit_notes/{credit_note_id}/items",
	}
}

// EnvelopeKey unwraps {"credit_note": {...}}.
func (CreditNoteItemsRecord) EnvelopeKey() string { return CreditNoteKey }

// CreditNoteItem is one entry of a CreditNoteItemsRecord's items: what was credited and refunded,
// by source, each amount a string.
type CreditNoteItem struct {
	CreditNoteFromBlankMethod     *string
	CreditNoteFromCreditNote      *string
	CreditNoteFromGiftCertificate *string
	CreditNoteFromInvoiceCredit   *string
	CreditNoteFromPaymentMethod   *string
	CreditNoteTotal               *string
	Invoice                       any
	Object                        *string
	RefundFromBlankMethod         *string
	RefundFromCreditNote          *string
	RefundFromGiftCertificate     *string
	RefundFromInvoiceCredit       *string
	RefundFromPaymentMethod       *string
	RefundTotal                   *string
}

// CreditNoteApplicationsRecord is A credit note's applications, under credit_note.
type CreditNoteApplicationsRecord struct {
	Applications []any
	Href         *string
	Object       *string
}

// MeasuredFrom lists the endpoints CreditNoteApplicationsRecord was measured from.
func (CreditNoteApplicationsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/credit_notes/{credit_note_id}/credit_note_applications",
	}
}

// EnvelopeKey unwraps {"credit_note": {...}}.
func (CreditNoteApplicationsRecord) EnvelopeKey() string { return CreditNoteKey }

// CreditNoteNoteList is one page of CreditNoteNoteRecord records, under data.
type CreditNoteNoteList struct {
	Data       []CreditNoteNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints CreditNoteNoteList was measured from.
func (CreditNoteNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/credit_notes/{credit_note_id}/notes",
	}
}

// CreditNoteNoteRecord is one note on a credit note, under note.
type CreditNoteNoteRecord struct {
	Attachments   []any
	Content       *string
	CreatedBy     *CreditNoteNoteCreatedBy
	CreatedOn     *string
	CreditNote    *CreditNoteNoteDetail
	Group         any
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn any
	Object        *string
	User          any
	Version       *int64
	Watchers      []any
}

// MeasuredFrom lists the endpoints CreditNoteNoteRecord was measured from.
func (CreditNoteNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/credit_notes/{credit_note_id}/notes",
		"GET /api/{api_version}/credit_notes/{credit_note_id}/notes/{note_id}",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (CreditNoteNoteRecord) EnvelopeKey() string { return NoteKey }

// CreditNoteNoteCreatedBy is a CreditNoteNoteRecord's created_by.
type CreditNoteNoteCreatedBy struct {
	ID     *string
	Name   *string
	Object *string
}

// CreditNoteNoteDetail is a CreditNoteNoteRecord's credit_note.
type CreditNoteNoteDetail struct {
	Href   *string
	ID     *string
	Object *string
}

// CreditNoteItemsRecordShape reads a body as CreditNoteItemsRecord.
func CreditNoteItemsRecordShape(decoded any) any { return dto.Read[CreditNoteItemsRecord](decoded) }

// CreditNoteApplicationsRecordShape reads a body as CreditNoteApplicationsRecord.
func CreditNoteApplicationsRecordShape(decoded any) any {
	return dto.Read[CreditNoteApplicationsRecord](decoded)
}

// CreditNoteNoteListShape reads a body as CreditNoteNoteList.
func CreditNoteNoteListShape(decoded any) any { return dto.Read[CreditNoteNoteList](decoded) }

// CreditNoteNoteRecordShape reads a body as CreditNoteNoteRecord.
func CreditNoteNoteRecordShape(decoded any) any { return dto.Read[CreditNoteNoteRecord](decoded) }
