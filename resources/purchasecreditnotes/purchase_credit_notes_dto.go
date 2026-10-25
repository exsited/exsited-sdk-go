package purchasecreditnotes

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

// PurchaseCreditNoteRecord is the purchase credit note body.
type PurchaseCreditNoteRecord struct {
	Account *Account
	Amount  *string
	Code    *string
	// CreatedBy is untyped: the API sends it as a string or an object.
	CreatedBy        *CreatedBy
	CreatedOn        *string
	Currency         *Currency
	CustomAttributes []any
	Date             *string
	Href             *string
	ID               *string
	InitialAmount    *string
	IssueDate        *string
	// LastUpdatedBy is untyped: the API sends it as a string or an object.
	LastUpdatedBy          *LastUpdatedBy
	LastUpdatedOn          *string
	Object                 *string
	PaymentID              *string
	PurchaseCreditNoteRows []PurchaseCreditNoteRow
	PurchasePayment        *PurchasePayment
	RefundAmount           *string
	// Refundable is untyped: the API sends it as a string or a boolean.
	Refundable       any
	RemainingAmount  *string
	RemainingBalance *string
	Status           *string
	UUID             *string
	// Version is untyped: the API sends it as a string or an integer.
	Version any
}

// MeasuredFrom lists the endpoints PurchaseCreditNoteRecord was measured from.
func (PurchaseCreditNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_credit_notes",
	}
}

// PurchaseCreditNoteList is a page of PurchaseCreditNoteRecord records.
type PurchaseCreditNoteList struct {
	Data       []PurchaseCreditNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseCreditNoteList was measured from.
func (PurchaseCreditNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_credit_notes",
	}
}

// PurchaseCreditNoteReadRecord is the purchase credit note read body.
type PurchaseCreditNoteReadRecord struct {
	Account                *Account
	Code                   *string
	CreatedBy              *CreatedBy
	CreatedOn              *string
	Currency               *Currency
	Href                   *string
	ID                     *string
	InitialAmount          *string
	IssueDate              *string
	LastUpdatedBy          *LastUpdatedBy
	LastUpdatedOn          *string
	Object                 *string
	PurchaseCreditNoteRows []PurchaseCreditNoteRow
	PurchasePayment        *PurchasePayment
	RefundAmount           *string
	Refundable             *bool
	RemainingAmount        *string
	Status                 *string
	Version                *int64
}

// MeasuredFrom lists the endpoints PurchaseCreditNoteReadRecord was measured from.
func (PurchaseCreditNoteReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}",
	}
}

// EnvelopeKey unwraps {"purchase_credit_note": {...}}.
func (PurchaseCreditNoteReadRecord) EnvelopeKey() string { return PurchaseCreditNoteKey }

// PurchaseCreditNoteUpdateRecord is the body UpdatePurchaseCreditNote answers.
type PurchaseCreditNoteUpdateRecord struct {
	Code    *string
	Href    *string
	ID      *string
	Object  *string
	Status  *Status
	Version *int64
}

// MeasuredFrom lists the endpoints PurchaseCreditNoteUpdateRecord was measured from.
func (PurchaseCreditNoteUpdateRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}",
	}
}

// EnvelopeKey unwraps {"purchase_credit_note": {...}}.
func (PurchaseCreditNoteUpdateRecord) EnvelopeKey() string { return PurchaseCreditNoteKey }

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// Status is the status object.
type Status struct {
	Key   *string
	Value *string
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

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// PaymentMethod is the payment_method object.
type PaymentMethod struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PaymentProcessor is the payment_processor object.
type PaymentProcessor struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PurchaseCreditNoteRow is one entry of the purchase_credit_note_rows list.
type PurchaseCreditNoteRow struct {
	CreditNoteFromBlankMethod   *string
	CreditNoteFromCreditNote    *string
	CreditNoteFromPaymentMethod *string
	CreditNoteTotal             *string
	PaymentMethod               *PaymentMethod
	PaymentProcessor            *PaymentProcessor
	RefundFromBlankMethod       *string
	RefundFromCreditNote        *string
	RefundFromPaymentMethod     *string
	RefundTotal                 *string
}

// PurchasePayment is the purchase_payment object.
type PurchasePayment struct {
	Code   *string
	Href   *string
	ID     *string
	Object *string
}

// PurchaseCreditNoteListShape reads a body as PurchaseCreditNoteList.
func PurchaseCreditNoteListShape(decoded any) any { return dto.Read[PurchaseCreditNoteList](decoded) }

// PurchaseCreditNoteReadRecordShape reads a body as PurchaseCreditNoteReadRecord.
func PurchaseCreditNoteReadRecordShape(decoded any) any {
	return dto.Read[PurchaseCreditNoteReadRecord](decoded)
}

// PurchaseCreditNoteUpdateRecordShape reads a body as PurchaseCreditNoteUpdateRecord.
func PurchaseCreditNoteUpdateRecordShape(decoded any) any {
	return dto.Read[PurchaseCreditNoteUpdateRecord](decoded)
}

// PurchaseCreditNoteCreatedRecord is what CreatePurchaseCreditNote answers, under
// purchase_credit_note: a shorter body than the reads, and status an object where they send a
// string.
type PurchaseCreditNoteCreatedRecord struct {
	ActivityLogs    []any
	Code            *string
	Href            *string
	ID              *string
	InitialAmount   *string
	IssueDate       *string
	Object          *string
	RemainingAmount *string
	Status          *PurchaseCreditNoteCreatedStatus
	Version         *int64
}

// MeasuredFrom lists the endpoints PurchaseCreditNoteCreatedRecord was measured from.
func (PurchaseCreditNoteCreatedRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/purchase_credit_notes",
	}
}

// EnvelopeKey unwraps {"purchase_credit_note": {...}}.
func (PurchaseCreditNoteCreatedRecord) EnvelopeKey() string { return PurchaseCreditNoteKey }

// PurchaseCreditNoteCreatedStatus is a PurchaseCreditNoteCreatedRecord's status.
type PurchaseCreditNoteCreatedStatus struct {
	Key   *string
	Value *string
}

// PurchaseCreditNoteCreatedRecordShape reads a body as PurchaseCreditNoteCreatedRecord.
func PurchaseCreditNoteCreatedRecordShape(decoded any) any {
	return dto.Read[PurchaseCreditNoteCreatedRecord](decoded)
}
