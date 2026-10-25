package giftcertificates

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

// GiftCertificateRecord is the gift certificate body.
type GiftCertificateRecord struct {
	AccountingCode   *AccountingCode
	ActivityLogs     []ActivityLog
	Allocations      []Allocation
	Amount           *string
	Code             *string
	CreatedBy        any
	CreatedOn        *string
	Currency         *Currency
	ExpiryDate       *string
	Href             *string
	ID               *string
	IssueDate        *string
	LastUpdatedBy    any
	LastUpdatedOn    any
	LinkedItem       any
	Object           *string
	RemainingBalance *string
	Status           *string
	Transactions     []Transaction
	UsedAmount       *string
	Version          *int64
}

// MeasuredFrom lists the endpoints GiftCertificateRecord was measured from.
func (GiftCertificateRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/gift_certificates",
		"GET /api/{api_version}/gift_certificates/{gc_id}",
		"POST /api/{api_version}/gift_certificates",
		"PATCH /api/{api_version}/gift_certificates/{gc_id}",
	}
}

// EnvelopeKey unwraps {"gift_certificate": {...}}.
func (GiftCertificateRecord) EnvelopeKey() string { return GiftCertificateKey }

// GiftCertificateList is a page of GiftCertificateRecord records.
type GiftCertificateList struct {
	Data       []GiftCertificateRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints GiftCertificateList was measured from.
func (GiftCertificateList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/gift_certificates",
	}
}

// GiftCertificateTransactionRecord is the gift certificate transaction body.
type GiftCertificateTransactionRecord struct {
	Account        *Account
	AccountingCode *AccountingCode
	Amount         *string
	CreatedBy      any
	CreatedOn      *string
	Currency       *Currency
	Date           *string
	Href           *string
	ID             *string
	LastUpdatedBy  any
	LastUpdatedOn  *string
	Object         *string
	Reference      any
	Status         *string
	Type           *string
	Version        *int64
}

// MeasuredFrom lists the endpoints GiftCertificateTransactionRecord was measured from.
func (GiftCertificateTransactionRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/gift_certificates/{gc_id}/transactions",
	}
}

// GiftCertificateTransactionList is a page of GiftCertificateTransactionRecord records.
type GiftCertificateTransactionList struct {
	Data       []GiftCertificateTransactionRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints GiftCertificateTransactionList was measured from.
func (GiftCertificateTransactionList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/gift_certificates/{gc_id}/transactions",
	}
}

// GiftCertificateAllocationsRecord is the gift certificate allocations body.
type GiftCertificateAllocationsRecord struct {
	Data       []GiftCertificateTransactionRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints GiftCertificateAllocationsRecord was measured from.
func (GiftCertificateAllocationsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/gift_certificates/{gc_id}/allocations",
	}
}

// ActivityLog is one entry of the activity_logs list.
type ActivityLog struct {
	Activity *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// Account is the account object.
type Account struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// AccountingCode is the accounting_code object.
type AccountingCode struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Allocation is one entry of the allocations list.
type Allocation struct {
	Account       *Account
	CreatedBy     any
	CreatedOn     *string
	Date          *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Object        *string
	Status        *string
	Type          *string
	Version       *int64
}

// Currency is the currency object.
type Currency struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Transaction is one entry of the transactions list.
type Transaction struct {
	AccountingCode *AccountingCode
	Amount         *string
	CreatedBy      any
	CreatedOn      *string
	Currency       *Currency
	Date           *string
	Href           *string
	ID             *string
	LastUpdatedBy  any
	LastUpdatedOn  *string
	Object         *string
	Reference      any
	Status         *string
	Type           *string
	Version        *int64
}

// GiftCertificateListShape reads a body as GiftCertificateList.
func GiftCertificateListShape(decoded any) any { return dto.Read[GiftCertificateList](decoded) }

// GiftCertificateRecordShape reads a body as GiftCertificateRecord.
func GiftCertificateRecordShape(decoded any) any { return dto.Read[GiftCertificateRecord](decoded) }

// GiftCertificateAllocationsRecordShape reads a body as GiftCertificateAllocationsRecord.
func GiftCertificateAllocationsRecordShape(decoded any) any {
	return dto.Read[GiftCertificateAllocationsRecord](decoded)
}

// GiftCertificateTransactionListShape reads a body as GiftCertificateTransactionList.
func GiftCertificateTransactionListShape(decoded any) any {
	return dto.Read[GiftCertificateTransactionList](decoded)
}

// GiftCertificateAllocationReadRecord is one allocation, under gift_certificate_allocation: the
// fields an allocation carries in allocations.
type GiftCertificateAllocationReadRecord struct {
	Account       *GiftCertificateAllocationAccount
	CreatedBy     any
	CreatedOn     *string
	Date          *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Object        *string
	Status        *string
	Type          *string
	Version       *int64
}

// MeasuredFrom lists the endpoints GiftCertificateAllocationReadRecord was measured from.
func (GiftCertificateAllocationReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/gift_certificates/{gc_id}/allocations/{allocation_id}",
	}
}

// EnvelopeKey unwraps {"gift_certificate_allocation": {...}}.
func (GiftCertificateAllocationReadRecord) EnvelopeKey() string { return GiftCertificateAllocationKey }

// GiftCertificateAllocationAccount is a GiftCertificateAllocation's account.
type GiftCertificateAllocationAccount struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// GiftCertificateTransactionReadRecord is one transaction, under gift_certificate_transaction: the
// fields a transaction carries in transactions.
type GiftCertificateTransactionReadRecord struct {
	AccountingCode *GiftCertificateTransactionAccountingCode
	Amount         *string
	CreatedBy      any
	CreatedOn      *string
	Currency       *GiftCertificateTransactionCurrency
	Date           *string
	Href           *string
	ID             *string
	LastUpdatedBy  any
	LastUpdatedOn  *string
	Object         *string
	Reference      any
	Status         *string
	Type           *string
	Version        *int64
}

// MeasuredFrom lists the endpoints GiftCertificateTransactionReadRecord was measured from.
func (GiftCertificateTransactionReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/gift_certificates/{gc_id}/transactions/{transaction_id}",
	}
}

// EnvelopeKey unwraps {"gift_certificate_transaction": {...}}.
func (GiftCertificateTransactionReadRecord) EnvelopeKey() string {
	return GiftCertificateTransactionKey
}

// GiftCertificateTransactionAccountingCode is a GiftCertificateTransaction's accounting_code.
type GiftCertificateTransactionAccountingCode struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// GiftCertificateTransactionCurrency is a GiftCertificateTransaction's currency.
type GiftCertificateTransactionCurrency struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// GiftCertificateAllocationReadRecordShape reads a body as GiftCertificateAllocationReadRecord.
func GiftCertificateAllocationReadRecordShape(decoded any) any {
	return dto.Read[GiftCertificateAllocationReadRecord](decoded)
}

// GiftCertificateTransactionReadRecordShape reads a body as GiftCertificateTransactionReadRecord.
func GiftCertificateTransactionReadRecordShape(decoded any) any {
	return dto.Read[GiftCertificateTransactionReadRecord](decoded)
}
