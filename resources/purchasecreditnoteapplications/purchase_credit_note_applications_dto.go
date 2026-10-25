package purchasecreditnoteapplications

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

// PurchaseCreditNoteApplicationRecord is the purchase credit note application body.
type PurchaseCreditNoteApplicationRecord struct {
	Amount *string
	// CreatedBy is untyped: the API sends it as a string or an object.
	CreatedBy          *CreatedBy
	CreatedOn          *string
	CreditNoteID       *string
	Date               *string
	Href               *string
	ID                 *string
	IssueDate          *string
	LastUpdatedBy      any
	LastUpdatedOn      *string
	Object             *string
	PaymentID          *string
	PurchaseCreditNote *PurchaseCreditNote
	PurchasePayment    *PurchasePayment
	RemainingBalance   *string
	Status             *string
	UUID               *string
	// Version is untyped: the API sends it as a string or an integer.
	Version any
}

// MeasuredFrom lists the endpoints PurchaseCreditNoteApplicationRecord was measured from.
func (PurchaseCreditNoteApplicationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_credit_note_applications",
	}
}

// PurchaseCreditNoteApplicationList is a page of PurchaseCreditNoteApplicationRecord
// records.
type PurchaseCreditNoteApplicationList struct {
	Data       []PurchaseCreditNoteApplicationRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseCreditNoteApplicationList was measured from.
func (PurchaseCreditNoteApplicationList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_credit_note_applications",
	}
}

// PurchaseCreditNoteApplicationReadRecord is the purchase credit note application read
// body.
type PurchaseCreditNoteApplicationReadRecord struct {
	Amount             *string
	CreatedBy          *CreatedBy
	CreatedOn          *string
	Href               *string
	ID                 *string
	IssueDate          *string
	LastUpdatedBy      any
	LastUpdatedOn      *string
	Object             *string
	PurchaseCreditNote *PurchaseCreditNote
	PurchasePayment    *PurchasePayment
	RemainingBalance   *string
	Status             *string
	Version            *int64
}

// MeasuredFrom lists the endpoints PurchaseCreditNoteApplicationReadRecord was measured from.
func (PurchaseCreditNoteApplicationReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_credit_note_applications/{purchase_credit_note_application_id}",
	}
}

// EnvelopeKey unwraps {"purchase_credit_note_application": {...}}.
func (PurchaseCreditNoteApplicationReadRecord) EnvelopeKey() string {
	return "purchase_credit_note_application"
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// CreatedBy is the created_by object.
type CreatedBy struct {
	Name   *string
	Object *string
}

// PurchaseCreditNote is the purchase_credit_note object.
type PurchaseCreditNote struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PurchasePayment is the purchase_payment object.
type PurchasePayment struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PurchaseCreditNoteApplicationListShape reads a body as PurchaseCreditNoteApplicationList.
func PurchaseCreditNoteApplicationListShape(decoded any) any {
	return dto.Read[PurchaseCreditNoteApplicationList](decoded)
}

// PurchaseCreditNoteApplicationReadRecordShape reads a body as PurchaseCreditNoteApplicationReadRecord.
func PurchaseCreditNoteApplicationReadRecordShape(decoded any) any {
	return dto.Read[PurchaseCreditNoteApplicationReadRecord](decoded)
}
