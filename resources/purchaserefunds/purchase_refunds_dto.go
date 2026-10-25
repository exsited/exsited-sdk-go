package purchaserefunds

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

// PurchaseRefundsRecord is the purchase refunds body.
type PurchaseRefundsRecord struct {
	AlternateDate    *string
	Amount           *string
	CreatedBy        *string
	CreatedOn        *string
	CreditNoteID     *string
	CustomAttributes []any
	Data             []PurchaseRefund
	Date             *string
	ID               *string
	LastUpdatedBy    *string
	LastUpdatedOn    *string
	Note             *string
	Object           *string
	Pagination       *Pagination
	PaymentMethod    *string
	PaymentProcessor *string
	Reference        *string
	Status           *string
	URL              *string
	UUID             *string
	Version          *string
}

// MeasuredFrom lists the endpoints PurchaseRefundsRecord was measured from.
func (PurchaseRefundsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_refunds",
	}
}

// PurchaseRefundRecord is the purchase refund body.
type PurchaseRefundRecord struct {
	Account          *Account
	AlternateDate    *string
	Amount           *string
	CreatedBy        *CreatedBy
	CreatedOn        *string
	CreditNote       *CreditNote
	CreditNoteID     *string
	CustomAttributes []CustomAttribute
	Date             *string
	ID               *string
	IssueDate        *string
	LastUpdatedBy    *string
	LastUpdatedOn    *string
	Note             *string
	Object           *string
	PaymentMethod    *string
	PaymentProcessor *string
	PurchaseRefund   any
	Reference        *string
	Status           *string
	UUID             *string
	Version          *int64
}

// MeasuredFrom lists the endpoints PurchaseRefundRecord was measured from.
func (PurchaseRefundRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_refunds/{purchase_refund_id}",
	}
}

// EnvelopeKey unwraps {"purchase_refund": {...}}.
func (PurchaseRefundRecord) EnvelopeKey() string { return PurchaseRefundKey }

// AccountPurchaseRefundsRecord is the account purchase refunds body.
type AccountPurchaseRefundsRecord struct {
	Data       []AccountPurchaseRefund
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AccountPurchaseRefundsRecord was measured from.
func (AccountPurchaseRefundsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/purchase_refunds",
	}
}

// Account is the account object.
type Account struct {
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

// CreditNote is the credit_note object.
type CreditNote struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// AccountPurchaseRefund is one row of AccountPurchaseRefundsRecord.
type AccountPurchaseRefund struct {
	Account          *Account
	Amount           *string
	CreatedBy        *CreatedBy
	CreatedOn        *string
	CreditNote       *CreditNote
	CustomAttributes []CustomAttribute
	ID               *string
	IssueDate        *string
	LastUpdatedBy    any
	LastUpdatedOn    any
	Note             *string
	Object           *string
	PaymentMethod    *string
	PaymentProcessor *string
	Reference        *string
	Status           *string
	UUID             *string
	Version          *int64
}

// CustomAttribute is one entry of the custom_attributes list.
type CustomAttribute struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
	Type   *string
	Value  *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// PurchaseRefund is one row of PurchaseRefundsRecord.
type PurchaseRefund struct {
	Account          *Account
	Amount           *string
	CreatedBy        *CreatedBy
	CreatedOn        *string
	CreditNote       *CreditNote
	CustomAttributes []CustomAttribute
	ID               *string
	IssueDate        *string
	LastUpdatedBy    any
	LastUpdatedOn    any
	Note             *string
	Object           *string
	PaymentMethod    *string
	PaymentProcessor *string
	Reference        *string
	Status           *string
	UUID             *string
	Version          *int64
}

// PurchaseRefundsRecordShape reads a body as PurchaseRefundsRecord.
func PurchaseRefundsRecordShape(decoded any) any { return dto.Read[PurchaseRefundsRecord](decoded) }

// PurchaseRefundRecordShape reads a body as PurchaseRefundRecord.
func PurchaseRefundRecordShape(decoded any) any { return dto.Read[PurchaseRefundRecord](decoded) }

// AccountPurchaseRefundsRecordShape reads a body as AccountPurchaseRefundsRecord.
func AccountPurchaseRefundsRecordShape(decoded any) any {
	return dto.Read[AccountPurchaseRefundsRecord](decoded)
}
