package invoices

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

// InvoiceRecord is one invoice.
type InvoiceRecord struct {
	Account                   *Account
	AccountID                 *string
	AccountName               *string
	ActivityLogs              []ActivityLog
	AlternateBillingEndDate   *string
	AlternateBillingStartDate *string
	AlternateDueDate          *string
	AlternateIssueDate        *string
	BillingEndDate            *string
	BillingStartDate          *string
	Code                      *int64
	// CreatedBy is untyped: no single type reads every record.
	CreatedBy               *CreatedBy
	CreatedOn               *string
	CreditApplied           *string
	CreditIssued            *string
	Currency                *Currency
	CustomAttributes        []any
	CustomForm              *CustomForm
	CustomObjects           []any
	CustomerPurchaseOrderID *string
	// DiscountAmount is untyped, and the one money field that is not text.
	DiscountAmount   any
	Due              *string
	DueDate          *string
	ID               *string
	InvoiceNote      *string
	Invoices         []any
	IssueDate        *string
	ItemFulfillment  map[string]any
	ItemFulfillments []any
	Kpis             map[string]any
	LastPayment      *LastPayment
	// LastUpdatedBy is untyped, the same two shapes as CreatedBy.
	LastUpdatedBy  *LastUpdatedBy
	LastUpdatedOn  *string
	Lines          []Line
	Message        *string
	Note           map[string]any
	Notes          []any
	Object         *string
	Order          *Order
	OrderID        *string
	Origin         *string
	Outstanding    *string
	Pagination     map[string]any
	Paid           *string
	PaymentApplied *string
	PaymentStatus  *string
	// PriceTaxInclusive is untyped.
	PriceTaxInclusive               any
	RecurringRateCount              *string
	ReturnMerchandiseAuthorisation  map[string]any
	ReturnMerchandiseAuthorisations []any
	ShippingCost                    *string
	Status                          *string
	Subtotal                        *string
	Tax                             *string
	TaxTotal                        *string
	Total                           *string
	Type                            *string
	UsageRateCount                  *string
	UUID                            *string
	// Version is untyped.
	Version any
}

// EnvelopeKey unwraps {"invoice": {...}}.
func (InvoiceRecord) EnvelopeKey() string { return InvoiceKey }

// InvoiceList is a page of invoices.
type InvoiceList struct {
	Data       []InvoiceRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// InvoiceNoteList is a page of an invoice's notes.
type InvoiceNoteList struct {
	Data       []InvoiceNote
	Object     *string
	Pagination *Pagination
	URL        *string
}

// InvoiceItemFulfillmentList is the page GetInvoiceItemFulfillments answers.
type InvoiceItemFulfillmentList struct {
	Data       []InvoiceItemFulfillmentRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints InvoiceItemFulfillmentList was measured from.
func (InvoiceItemFulfillmentList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/item_fulfillments",
	}
}

// InvoiceItemFulfillmentView is the body GetInvoiceItemFulfillment answers.
type InvoiceItemFulfillmentView struct {
	Account              *Account
	Code                 *string
	CreatedBy            *CreatedBy
	CreatedOn            *string
	CustomAttributes     map[string]any
	DeliveryStatus       any
	FulfillmentCompleted *bool
	FulfillmentDate      *string
	FulfillmentNote      *string
	FulfillmentStatus    *string
	Fulfillments         []Fulfillment
	Href                 *string
	ID                   *string
	Invoice              *Invoice
	LastUpdatedBy        *LastUpdatedBy
	LastUpdatedOn        *string
	Object               *string
	Order                any
	Status               *string
	TrackingNumber       *string
	Version              *int64
}

// MeasuredFrom lists the endpoints InvoiceItemFulfillmentView was measured from.
func (InvoiceItemFulfillmentView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/item_fulfillments/{fulfillment_id}",
	}
}

// EnvelopeKey unwraps {"item_fulfillment": {...}}.
func (InvoiceItemFulfillmentView) EnvelopeKey() string { return ItemFulfillmentKey }

// InvoiceNoteFileList is the page GetInvoiceNoteFiles answers.
type InvoiceNoteFileList struct {
	Data       []InvoiceNoteFileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints InvoiceNoteFileList was measured from.
func (InvoiceNoteFileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}/files",
	}
}

// InvoiceNoteFileView is the body GetInvoiceNoteFile answers.
type InvoiceNoteFileView struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints InvoiceNoteFileView was measured from.
func (InvoiceNoteFileView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}/files/{invoice_file_id}",
	}
}

// EnvelopeKey unwraps {"file": {...}}.
func (InvoiceNoteFileView) EnvelopeKey() string { return FileKey }

// InvoiceNoteView is the body GetInvoiceNote answers.
type InvoiceNoteView struct {
	Attachments   []Attachment
	Content       *string
	CreatedBy     *CreatedBy
	CreatedOn     *string
	Group         any
	Href          *string
	ID            *string
	Invoice       *Invoice
	LastUpdatedBy any
	LastUpdatedOn any
	Object        *string
	User          any
	Version       *int64
	Watchers      []Watcher
}

// MeasuredFrom lists the endpoints InvoiceNoteView was measured from.
func (InvoiceNoteView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}",
		"POST /api/{api_version}/invoices/{invoice_id}/notes",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (InvoiceNoteView) EnvelopeKey() string { return NoteKey }

// InvoicePDFView is the body GetInvoicePDF answers.
type InvoicePDFView struct {
	PDF *PDF
}

// MeasuredFrom lists the endpoints InvoicePDFView was measured from.
func (InvoicePDFView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/pdf",
	}
}

// EnvelopeKey unwraps {"invoice": {...}}.
func (InvoicePDFView) EnvelopeKey() string { return InvoiceKey }

// InvoiceReturnMerchandiseAuthorisationList is the page GetInvoiceReturnMerchandiseAuthorisations answers.
type InvoiceReturnMerchandiseAuthorisationList struct {
	Data       []InvoiceReturnMerchandiseAuthorisationRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints InvoiceReturnMerchandiseAuthorisationList was measured from.
func (InvoiceReturnMerchandiseAuthorisationList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations",
	}
}

// InvoiceReturnMerchandiseAuthorisationView is the body GetInvoiceReturnMerchandiseAuthorisation answers.
type InvoiceReturnMerchandiseAuthorisationView struct {
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
	RmaDate       *string
	SaleInvoice   *SaleInvoice
	SaleOrder     any
	Status        *string
	Version       *int64
	Warehouse     any
}

// MeasuredFrom lists the endpoints InvoiceReturnMerchandiseAuthorisationView was measured from.
func (InvoiceReturnMerchandiseAuthorisationView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations/{rma_id}",
	}
}

// EnvelopeKey unwraps {"return_merchandise_authorisation": {...}}.
func (InvoiceReturnMerchandiseAuthorisationView) EnvelopeKey() string {
	return ReturnMerchandiseAuthorisationKey
}

// InvoiceItemFulfillmentRecord is one row of InvoiceItemFulfillmentList.
type InvoiceItemFulfillmentRecord struct {
	Account              *Account
	Code                 *string
	CreatedBy            *CreatedBy
	CreatedOn            *string
	CustomAttributes     map[string]any
	DeliveryStatus       any
	FulfillmentCompleted *bool
	FulfillmentDate      *string
	FulfillmentNote      *string
	FulfillmentStatus    *string
	Fulfillments         []Fulfillment
	Href                 *string
	ID                   *string
	Invoice              *Invoice
	LastUpdatedBy        *LastUpdatedBy
	LastUpdatedOn        *string
	Object               *string
	Order                any
	Status               *string
	TrackingNumber       *string
	Version              *int64
}

// MeasuredFrom lists the endpoints InvoiceItemFulfillmentRecord was measured from.
func (InvoiceItemFulfillmentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/item_fulfillments",
	}
}

// InvoiceNoteFileRecord is one row of InvoiceNoteFileList.
type InvoiceNoteFileRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints InvoiceNoteFileRecord was measured from.
func (InvoiceNoteFileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}/files",
	}
}

// InvoiceReturnMerchandiseAuthorisationRecord is one row of InvoiceReturnMerchandiseAuthorisationList.
type InvoiceReturnMerchandiseAuthorisationRecord struct {
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
	RmaDate       *string
	SaleInvoice   *SaleInvoice
	SaleOrder     any
	Status        *string
	Version       *int64
	Warehouse     any
}

// MeasuredFrom lists the endpoints InvoiceReturnMerchandiseAuthorisationRecord was measured from.
func (InvoiceReturnMerchandiseAuthorisationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations",
	}
}

// Account is the account object.
type Account struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Attachment is one entry of the attachments list.
type Attachment struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// CreatedBy is the created_by object.
type CreatedBy struct {
	Name   *string
	Object *string
}

// Fulfillment is one entry of the fulfillments list.
type Fulfillment struct {
	FulfillmentIndex   *int64
	FulfillmentQty     *float64
	FulfillmentSentQty *float64
	ID                 *string
	InventoryTracking  any
	Item               *Item
	Object             *string
	PendingReturn      *float64
	Returned           *float64
	ServingStatus      *string
	Uom                *string
	UomPrecision       *int64
}

// Invoice is the invoice object.
type Invoice struct {
	Href   *string
	ID     *string
	Object *string
}

// Item is the item object.
type Item struct {
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

// Order is the order object.
type Order struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PDF is the pdf object.
type PDF struct {
	DownloadURL *string
	FileSize    *int64
	ID          *string
	Object      *string
	PreviewURL  *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
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

// Watcher is one entry of the watchers list.
type Watcher struct {
	Email *string
	Name  *string
	Type  *string
}

// AccountingCode is the accounting_code object.
type AccountingCode struct {
	SalesRevenue any
}

// ActivityLog is one entry of the activity_logs list.
type ActivityLog struct {
	Activity *string
}

// Currency is the currency object.
type Currency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// CustomForm is the custom_form object.
type CustomForm struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// InvoiceNote is one row of InvoiceNoteList.
type InvoiceNote struct {
	Attachments   []Attachment
	Content       *string
	CreatedBy     *CreatedBy
	CreatedOn     *string
	Group         any
	Href          *string
	ID            *string
	Invoice       *Invoice
	LastUpdatedBy any
	LastUpdatedOn any
	Object        *string
	User          any
	Version       *int64
	Watchers      []Watcher
}

// LastPayment is the last_payment object.
type LastPayment struct {
	Amount    *string
	CreatedOn *string
	ID        *string
}

// Line is one entry of the lines list.
type Line struct {
	AccountingCode             *AccountingCode
	AlternateChargingEndDate   *string
	AlternateChargingStartDate *string
	ChargingEndDate            *string
	ChargingStartDate          *string
	ID                         *string
	Item                       *Item
	ItemChargeType             *string
	ItemCustomAttributes       []any
	ItemID                     *string
	ItemName                   *string
	ItemOrderQuantity          *int64
	ItemUom                    any
	ItemUUID                   *string
	ItemWarehouse              *string
	Object                     *string
	PricingSnapshotPrice       *float64
	PricingSnapshotUUID        *string
	Subtotal                   *string
	Tax                        *Tax
	Total                      *string
	Version                    *int64
}

// Tax is the tax object.
type Tax struct {
	Amount *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
	Rate   *float64
}

// InvoiceListShape reads a body as InvoiceList.
func InvoiceListShape(decoded any) any { return dto.Read[InvoiceList](decoded) }

// InvoiceRecordShape reads a body as InvoiceRecord.
func InvoiceRecordShape(decoded any) any { return dto.Read[InvoiceRecord](decoded) }

// InvoiceNoteListShape reads a body as InvoiceNoteList.
func InvoiceNoteListShape(decoded any) any { return dto.Read[InvoiceNoteList](decoded) }

// InvoiceItemFulfillmentListShape reads a body as InvoiceItemFulfillmentList.
func InvoiceItemFulfillmentListShape(decoded any) any {
	return dto.Read[InvoiceItemFulfillmentList](decoded)
}

// InvoiceItemFulfillmentViewShape reads a body as InvoiceItemFulfillmentView.
func InvoiceItemFulfillmentViewShape(decoded any) any {
	return dto.Read[InvoiceItemFulfillmentView](decoded)
}

// InvoiceNoteFileListShape reads a body as InvoiceNoteFileList.
func InvoiceNoteFileListShape(decoded any) any { return dto.Read[InvoiceNoteFileList](decoded) }

// InvoiceNoteFileViewShape reads a body as InvoiceNoteFileView.
func InvoiceNoteFileViewShape(decoded any) any { return dto.Read[InvoiceNoteFileView](decoded) }

// InvoiceNoteViewShape reads a body as InvoiceNoteView.
func InvoiceNoteViewShape(decoded any) any { return dto.Read[InvoiceNoteView](decoded) }

// InvoicePDFViewShape reads a body as InvoicePDFView.
func InvoicePDFViewShape(decoded any) any { return dto.Read[InvoicePDFView](decoded) }

// InvoiceReturnMerchandiseAuthorisationListShape reads a body as InvoiceReturnMerchandiseAuthorisationList.
func InvoiceReturnMerchandiseAuthorisationListShape(decoded any) any {
	return dto.Read[InvoiceReturnMerchandiseAuthorisationList](decoded)
}

// InvoiceReturnMerchandiseAuthorisationViewShape reads a body as InvoiceReturnMerchandiseAuthorisationView.
func InvoiceReturnMerchandiseAuthorisationViewShape(decoded any) any {
	return dto.Read[InvoiceReturnMerchandiseAuthorisationView](decoded)
}

// InvoiceKpisReadRecord is what GetInvoiceKPIs answers, under invoice_kpis: the block an invoice
// carries as kpis, with an object beside it.
type InvoiceKpisReadRecord struct {
	CreditApplied     *string
	CreditIssued      *string
	PaymentApplied    *string
	LastPaymentDate   *string
	LastAmendedOn     *string
	LastCancelledOn   *string
	LastReactivatedOn *string
	DeletedOn         *string
	VoidedOn          *string
	Outstanding       any
	Overdue           any
	Object            *string
}

// EnvelopeKey unwraps {"invoice_kpis": {...}}.
func (InvoiceKpisReadRecord) EnvelopeKey() string { return InvoiceKPIsKey }

// InvoiceLinesRecord is what GetInvoiceLines answers, under invoice: its object and lines, and
// nothing else of the invoice.
type InvoiceLinesRecord struct {
	Object *string
	Lines  []InvoiceLine
}

// EnvelopeKey unwraps {"invoice": {...}}.
func (InvoiceLinesRecord) EnvelopeKey() string { return InvoiceKey }

// InvoiceLine is one line of an invoice.
type InvoiceLine struct {
	Object                     *string
	ID                         *string
	Version                    *int64
	Item                       *InvoiceLineItem
	ItemID                     *string
	ItemUUID                   *string
	ItemName                   *string
	ItemUom                    *string
	ItemWarehouse              *string
	ItemSerialOrBatchNumber    *string
	ItemOrderQuantity          *int64
	ItemChargeType             *string
	ItemCustomAttributes       []any
	PricingSnapshotPrice       *float64
	PricingSnapshotUUID        *string
	ChargingStartDate          *string
	ChargingEndDate            *string
	AlternateChargingStartDate *string
	AlternateChargingEndDate   *string
	Subtotal                   *string
	Total                      *string
	Tax                        *LineTax
	AccountingCode             *AccountingCode
}

// InvoiceLineItem is a InvoiceLine's item.
type InvoiceLineItem struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// LineTax is the tax block on a line. Present on every line.
type LineTax struct {
	Object *string
	ID     *string
	Name   *string
	Rate   *float64
	Amount *string
	Href   *string
}

// InvoiceLineReadRecord is what GetInvoiceLine answers, under invoice: its object and one line.
type InvoiceLineReadRecord struct {
	Object *string
	Line   *InvoiceLine
}

// EnvelopeKey unwraps {"invoice": {...}}.
func (InvoiceLineReadRecord) EnvelopeKey() string { return InvoiceKey }

// InvoiceKpisReadRecordShape reads a body as InvoiceKpisReadRecord.
func InvoiceKpisReadRecordShape(decoded any) any { return dto.Read[InvoiceKpisReadRecord](decoded) }

// InvoiceLinesRecordShape reads a body as InvoiceLinesRecord.
func InvoiceLinesRecordShape(decoded any) any { return dto.Read[InvoiceLinesRecord](decoded) }

// InvoiceLineReadRecordShape reads a body as InvoiceLineReadRecord.
func InvoiceLineReadRecordShape(decoded any) any { return dto.Read[InvoiceLineReadRecord](decoded) }
