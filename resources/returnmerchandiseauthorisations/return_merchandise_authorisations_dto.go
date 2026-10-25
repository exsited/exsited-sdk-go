package returnmerchandiseauthorisations

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

// ReturnMerchandiseAuthorisationRecord is the return merchandise authorisation body.
type ReturnMerchandiseAuthorisationRecord struct {
	Account       *Account
	Code          *string
	CreatedBy     *CreatedBy
	CreatedOn     *string
	Href          *string
	ID            *string
	ItemCount     *int64
	LastUpdatedBy any
	LastUpdatedOn any
	Note          *string
	Object        *string
	ReturnStatus  *string
	RMADate       *string
	SaleInvoice   *SaleInvoice
	SaleOrder     any
	Status        *string
	Version       *int64
	Warehouse     any
}

// MeasuredFrom lists the endpoints ReturnMerchandiseAuthorisationRecord was measured from.
func (ReturnMerchandiseAuthorisationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/return_merchandise_authorisations",
	}
}

// ReturnMerchandiseAuthorisationList is a page of ReturnMerchandiseAuthorisationRecord
// records.
type ReturnMerchandiseAuthorisationList struct {
	Data       []ReturnMerchandiseAuthorisationRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ReturnMerchandiseAuthorisationList was measured from.
func (ReturnMerchandiseAuthorisationList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/return_merchandise_authorisations",
	}
}

// InvoiceReturnMerchandiseAuthorisationReceiveRecord is the invoice return merchandise
// authorisation receive body.
type InvoiceReturnMerchandiseAuthorisationReceiveRecord struct {
	Code             *string
	CreatedBy        *CreatedBy
	CreatedOn        *string
	Href             *string
	ID               *string
	LastUpdatedBy    any
	LastUpdatedOn    any
	Object           *string
	QuantityReceived *float64
	ReceiveReturns   []any
	Status           *string
	UUID             *string
}

// MeasuredFrom lists the endpoints InvoiceReturnMerchandiseAuthorisationReceiveRecord was measured from.
func (InvoiceReturnMerchandiseAuthorisationReceiveRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations/{rma_code}/receive",
		"GET /api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations/{rma_code}/receive/{receive_uuid}",
	}
}

// EnvelopeKey unwraps {"receive_rma": {...}}.
func (InvoiceReturnMerchandiseAuthorisationReceiveRecord) EnvelopeKey() string { return ReceiveRmaKey }

// InvoiceReturnMerchandiseAuthorisationReceiveList is a page of
// InvoiceReturnMerchandiseAuthorisationReceiveRecord records.
type InvoiceReturnMerchandiseAuthorisationReceiveList struct {
	ReceiveReturnMerchandiseAuthorisations []InvoiceReturnMerchandiseAuthorisationReceiveRecord
	Pagination                             *Pagination
	Data                                   []InvoiceReturnMerchandiseAuthorisationReceiveRecord
	Object                                 *string
	URL                                    *string
}

// MeasuredFrom lists the endpoints InvoiceReturnMerchandiseAuthorisationReceiveList was measured from.
func (InvoiceReturnMerchandiseAuthorisationReceiveList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations/{rma_code}/receive",
	}
}

// ReturnMerchandiseAuthorisationReadRecord is the return merchandise authorisation read
// body.
type ReturnMerchandiseAuthorisationReadRecord struct {
	Account       *Account
	Code          *string
	CreatedBy     *CreatedBy
	CreatedOn     *string
	Href          *string
	ID            *string
	ItemCount     *int64
	LastUpdatedBy any
	LastUpdatedOn any
	Lines         []any
	Note          *string
	Object        *string
	ReceiveRmas   []ReceiveRma
	ReturnStatus  *string
	RMADate       *string
	SaleInvoice   *SaleInvoice
	SaleOrder     any
	Status        *string
	Version       *int64
	Warehouse     any
}

// MeasuredFrom lists the endpoints ReturnMerchandiseAuthorisationReadRecord was measured from.
func (ReturnMerchandiseAuthorisationReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations/{rma_id}",
		"GET /api/{api_version}/return_merchandise_authorisations/{rma_id}",
	}
}

// EnvelopeKey unwraps {"return_merchandise_authorisation": {...}}.
func (ReturnMerchandiseAuthorisationReadRecord) EnvelopeKey() string {
	return "return_merchandise_authorisation"
}

// ReturnMerchandiseAuthorisationsRecord is the return merchandise authorisations body.
type ReturnMerchandiseAuthorisationsRecord struct {
	ActivityLogs                   []any
	Code                           *string
	Href                           *string
	ID                             *string
	Object                         *string
	ReturnMerchandiseAuthorisation map[string]any
	Status                         *string
	Version                        *int64
}

// MeasuredFrom lists the endpoints ReturnMerchandiseAuthorisationsRecord was measured from.
func (ReturnMerchandiseAuthorisationsRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/return_merchandise_authorisations",
	}
}

// EnvelopeKey unwraps {"return_merchandise_authorisation": {...}}.
func (ReturnMerchandiseAuthorisationsRecord) EnvelopeKey() string {
	return ReturnMerchandiseAuthorisationKey
}

// CreatedBy is the created_by object.
type CreatedBy struct {
	Name   *string
	Object *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
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

// ReceiveRma is one entry of the receive_rmas list.
type ReceiveRma struct {
	Code             *string
	CreatedBy        *CreatedBy
	CreatedOn        *string
	Href             *string
	ID               *string
	LastUpdatedBy    any
	LastUpdatedOn    any
	Object           *string
	QuantityReceived *float64
	Status           *string
}

// SaleInvoice is the sale_invoice object.
type SaleInvoice struct {
	Code   *string
	Href   *string
	ID     *string
	Object *string
}

// ReturnMerchandiseAuthorisationListShape reads a body as ReturnMerchandiseAuthorisationList.
func ReturnMerchandiseAuthorisationListShape(decoded any) any {
	return dto.Read[ReturnMerchandiseAuthorisationList](decoded)
}

// ReturnMerchandiseAuthorisationsRecordShape reads a body as ReturnMerchandiseAuthorisationsRecord.
func ReturnMerchandiseAuthorisationsRecordShape(decoded any) any {
	return dto.Read[ReturnMerchandiseAuthorisationsRecord](decoded)
}

// ReturnMerchandiseAuthorisationReadRecordShape reads a body as ReturnMerchandiseAuthorisationReadRecord.
func ReturnMerchandiseAuthorisationReadRecordShape(decoded any) any {
	return dto.Read[ReturnMerchandiseAuthorisationReadRecord](decoded)
}

// InvoiceReturnMerchandiseAuthorisationReceiveListShape reads a body as InvoiceReturnMerchandiseAuthorisationReceiveList.
func InvoiceReturnMerchandiseAuthorisationReceiveListShape(decoded any) any {
	return dto.Read[InvoiceReturnMerchandiseAuthorisationReceiveList](decoded)
}

// InvoiceReturnMerchandiseAuthorisationReceiveRecordShape reads a body as InvoiceReturnMerchandiseAuthorisationReceiveRecord.
func InvoiceReturnMerchandiseAuthorisationReceiveRecordShape(decoded any) any {
	return dto.Read[InvoiceReturnMerchandiseAuthorisationReceiveRecord](decoded)
}
