package purchaseorders

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

// PurchaseOrderLineRecord is the purchase order line body.
type PurchaseOrderLineRecord struct {
	CreatedOn                    *string
	FlatDiscount                 *float64
	Href                         *string
	ID                           *string
	IsPriceTaxInclusive          *bool
	IsTaxExempt                  *bool
	Item                         *Item
	ItemAccountingCode           *ItemAccountingCode
	ItemDiscountAmount           *float64
	ItemExpectedCompletionDate   any
	ItemInvoiceNote              any
	ItemPriceSnapshot            *ItemPriceSnapshot
	ItemPriceTax                 *ItemPriceTax
	ItemPriceTaxExempt           *bool
	ItemPurchaseAccountingCode   *ItemPurchaseAccountingCode
	ItemPurchaseTaxConfiguration *ItemPurchaseTaxConfiguration
	LastUpdatedOn                *string
	Name                         *string
	Object                       *string
	Price                        *float64
	PurchaseLineNote             any
	Quantity                     *int64
	SortIndex                    *int64
	Status                       *string
	Subtotal                     *string
	Tax                          *Tax
	TaxTotal                     *string
	Total                        *string
	Uom                          *Uom
	Version                      *int64
	Warehouse                    *Warehouse
}

// MeasuredFrom lists the endpoints PurchaseOrderLineRecord was measured from.
func (PurchaseOrderLineRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/lines",
	}
}

// PurchaseOrderLineList is a page of PurchaseOrderLineRecord records.
type PurchaseOrderLineList struct {
	Data       []PurchaseOrderLineRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseOrderLineList was measured from.
func (PurchaseOrderLineList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/lines",
	}
}

// PurchaseOrderRecord is the purchase order body.
type PurchaseOrderRecord struct {
	Account                *Account
	BillTotal              *string
	Code                   *string
	CreatedBy              *CreatedBy
	CreatedOn              *string
	Currency               *Currency
	CustomAttributes       []any
	CustomForm             *CustomForm
	DueDate                *string
	DueTotal               *string
	ExpectedCompletionDate *string
	Href                   *string
	ID                     *string
	IsPriceTaxInclusive    *bool
	IssueDate              *string
	LastUpdatedBy          *LastUpdatedBy
	LastUpdatedOn          *string
	Object                 *string
	OperationalStatus      *string
	Origin                 *string
	PaidAmount             *string
	PurchaseOrderNote      *string
	SaleOrder              any
	Status                 *string
	Subtotal               *string
	SupplierInvoiceID      *string
	TaxTotal               *string
	Total                  *string
	Version                *int64
	Warehouse              any
}

// MeasuredFrom lists the endpoints PurchaseOrderRecord was measured from.
func (PurchaseOrderRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders",
	}
}

// PurchaseOrderList is a page of PurchaseOrderRecord records.
type PurchaseOrderList struct {
	Data       []PurchaseOrderRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseOrderList was measured from.
func (PurchaseOrderList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders",
	}
}

// PurchaseOrderPurchaseInvoiceRecord is the purchase order purchase invoice body.
type PurchaseOrderPurchaseInvoiceRecord struct {
	Account            *Account
	AlternateDueDate   *string
	AlternateIssueDate *string
	CreatedBy          *CreatedBy
	CreatedOn          *string
	Currency           *Currency
	CustomForm         any
	DueDate            *string
	ID                 *string
	IssueDate          *string
	LastUpdatedBy      *LastUpdatedBy
	LastUpdatedOn      *string
	Object             *string
	Origin             any
	PaymentStatus      *string
	PriceTaxInclusive  *bool
	PurchaseOrder      *PurchaseOrder
	Status             *string
	Subtotal           *string
	Tax                *string
	Total              *string
	UUID               *string
	Version            *int64
}

// MeasuredFrom lists the endpoints PurchaseOrderPurchaseInvoiceRecord was measured from.
func (PurchaseOrderPurchaseInvoiceRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_invoices",
	}
}

// PurchaseOrderPurchaseInvoiceList is a page of PurchaseOrderPurchaseInvoiceRecord
// records.
type PurchaseOrderPurchaseInvoiceList struct {
	Data       []PurchaseOrderPurchaseInvoiceRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseOrderPurchaseInvoiceList was measured from.
func (PurchaseOrderPurchaseInvoiceList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_invoices",
	}
}

// PurchaseOrderPurchasePaymentRecord is the purchase order purchase payment body.
type PurchaseOrderPurchasePaymentRecord struct {
	Account                     *Account
	Code                        *string
	CreatedBy                   *CreatedBy
	CreatedOn                   *string
	CreditNoteApplications      []CreditNoteApplication
	CreditNoteTotal             *string
	Currency                    *Currency
	Href                        *string
	ID                          *string
	LastUpdatedBy               *LastUpdatedBy
	LastUpdatedOn               *string
	Object                      *string
	Origin                      any
	PaidTotal                   *string
	PaymentDate                 *string
	PaymentMethodApplications   []PaymentMethodApplication
	PaymentNote                 *string
	PaymentReference            any
	PurchaseInvoiceApplications []PurchaseInvoiceApplication
	Status                      *string
	Version                     *int64
}

// MeasuredFrom lists the endpoints PurchaseOrderPurchasePaymentRecord was measured from.
func (PurchaseOrderPurchasePaymentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_payments",
	}
}

// PurchaseOrderPurchasePaymentList is a page of PurchaseOrderPurchasePaymentRecord
// records.
type PurchaseOrderPurchasePaymentList struct {
	Data       []PurchaseOrderPurchasePaymentRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseOrderPurchasePaymentList was measured from.
func (PurchaseOrderPurchasePaymentList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_payments",
	}
}

// PurchaseOrdersNoteRecord is the purchase orders note body.
type PurchaseOrdersNoteRecord struct {
	Attachments   []any
	Content       *string
	CreatedBy     *CreatedBy
	CreatedOn     *string
	Group         any
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn any
	Object        *string
	User          any
	Version       *int64
	Watchers      []Watcher
}

// MeasuredFrom lists the endpoints PurchaseOrdersNoteRecord was measured from.
func (PurchaseOrdersNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/notes",
	}
}

// PurchaseOrdersNoteList is a page of PurchaseOrdersNoteRecord records.
type PurchaseOrdersNoteList struct {
	Data       []PurchaseOrdersNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseOrdersNoteList was measured from.
func (PurchaseOrdersNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/notes",
	}
}

// PurchaseOrderReadRecord is the purchase order read body.
type PurchaseOrderReadRecord struct {
	Account                *Account
	ActivityLogs           []any
	BillTotal              *string
	Code                   *string
	CreatedBy              *CreatedBy
	CreatedOn              *string
	Currency               *Currency
	CustomAttributes       []CustomAttribute
	CustomForm             *CustomForm
	DueDate                *string
	DueTotal               *string
	ExpectedCompletionDate *string
	Href                   *string
	ID                     *string
	IsPriceTaxInclusive    *bool
	IssueDate              *string
	Kpis                   *PurchaseOrderKpis
	LastUpdatedBy          *LastUpdatedBy
	LastUpdatedOn          *string
	Object                 *string
	OperationalStatus      *string
	Origin                 *string
	PaidAmount             *string
	PurchaseOrderNote      *string
	SaleOrder              any
	Status                 *string
	Subtotal               *string
	SupplierInvoiceID      *string
	TaxTotal               *string
	Total                  *string
	Version                *int64
	Warehouse              any
}

// MeasuredFrom lists the endpoints PurchaseOrderReadRecord was measured from.
func (PurchaseOrderReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}",
		"POST /api/{api_version}/purchase_orders/{purchase_order_id}/change",
		"POST /api/{api_version}/purchase_orders/{purchase_order_id}/cancel",
		"POST /api/{api_version}/purchase_orders/{purchase_order_id}/reactivate",
	}
}

// EnvelopeKey unwraps {"purchase_order": {...}}.
func (PurchaseOrderReadRecord) EnvelopeKey() string { return PurchaseOrderKey }

// PurchaseOrderUpdatedRecord is the purchase order updated body.
type PurchaseOrderUpdatedRecord struct {
	Account                *Account
	AccountID              *string
	ActivityLogs           []any
	BillTotal              *string
	Code                   *string
	CreatedBy              *CreatedBy
	CreatedOn              *string
	Currency               *Currency
	CustomAttributes       []any
	CustomForm             *CustomForm
	CustomObjects          []any
	DueDate                *string
	DueTotal               *string
	ExpectedCompletionDate *string
	Href                   *string
	ID                     *string
	IsPriceTaxInclusive    any
	IssueDate              *string
	Kpis                   *PurchaseOrderUpdatedKpis
	LastUpdatedBy          *LastUpdatedBy
	LastUpdatedOn          *string
	Lines                  []any
	Object                 *string
	OperationalStatus      *string
	Origin                 *string
	PaidAmount             *string
	PriceTaxInclusive      *string
	PurchaseOrderNote      *string
	SaleOrder              any
	SaleOrderID            *string
	Status                 *string
	Subtotal               *string
	SupplierInvoiceID      *string
	Tax                    *string
	TaxTotal               *string
	Total                  *string
	UUID                   *string
	Version                *int64
	Warehouse              any
}

// MeasuredFrom lists the endpoints PurchaseOrderUpdatedRecord was measured from.
func (PurchaseOrderUpdatedRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/purchase_orders/{purchase_order_id}",
	}
}

// EnvelopeKey unwraps {"purchase_order": {...}}.
func (PurchaseOrderUpdatedRecord) EnvelopeKey() string { return PurchaseOrderKey }

// PurchaseOrdersRecord is the purchase orders body.
type PurchaseOrdersRecord struct {
	Account                *Account
	AccountID              *string
	ActivityLogs           []any
	BillTotal              *string
	Code                   *string
	CreatedBy              *CreatedBy
	CreatedOn              *string
	Currency               *Currency
	CustomAttributes       []any
	CustomForm             *CustomForm
	CustomObjects          []any
	DueDate                *string
	DueTotal               *string
	ExpectedCompletionDate *string
	Href                   *string
	ID                     *string
	IsPriceTaxInclusive    any
	IssueDate              *string
	Kpis                   *PurchaseOrdersKpis
	LastUpdatedBy          *string
	LastUpdatedOn          *string
	Lines                  []any
	Object                 *string
	OperationalStatus      *string
	Origin                 *string
	PaidAmount             *string
	PriceTaxInclusive      *string
	PurchaseOrderNote      *string
	SaleOrder              any
	Status                 *string
	Subtotal               *string
	SupplierInvoiceID      *string
	Tax                    *string
	TaxTotal               *string
	Total                  *string
	UUID                   *string
	Version                *int64
	Warehouse              any
}

// MeasuredFrom lists the endpoints PurchaseOrdersRecord was measured from.
func (PurchaseOrdersRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/purchase_orders",
	}
}

// EnvelopeKey unwraps {"purchase_order": {...}}.
func (PurchaseOrdersRecord) EnvelopeKey() string { return PurchaseOrderKey }

// PurchaseOrderInformationRecord is the purchase order information body.
type PurchaseOrderInformationRecord struct {
	BillTotal              *string
	Code                   *string
	CreatedBy              *CreatedBy
	CreatedOn              *string
	CustomAttributes       []CustomAttribute
	CustomForm             *CustomForm
	DueDate                *string
	DueTotal               *string
	ExpectedCompletionDate *string
	Href                   *string
	ID                     *string
	IsPriceTaxInclusive    *bool
	IssueDate              *string
	LastUpdatedBy          *LastUpdatedBy
	LastUpdatedOn          *string
	Object                 *string
	OperationalStatus      *string
	Origin                 *string
	PaidAmount             *string
	PurchaseOrderNote      *string
	Status                 *string
	Subtotal               *string
	SupplierInvoiceID      *string
	TaxTotal               *string
	Total                  *string
	Version                *int64
}

// MeasuredFrom lists the endpoints PurchaseOrderInformationRecord was measured from.
func (PurchaseOrderInformationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/information",
	}
}

// EnvelopeKey unwraps {"purchase_order": {...}}.
func (PurchaseOrderInformationRecord) EnvelopeKey() string { return PurchaseOrderKey }

// AccountPurchaseOrdersRecord is the account purchase orders body.
type AccountPurchaseOrdersRecord struct {
	Data       []PurchaseOrderRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AccountPurchaseOrdersRecord was measured from.
func (AccountPurchaseOrdersRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/purchase_orders",
	}
}

// PurchaseOrderNotesRecord is the purchase order notes body.
type PurchaseOrderNotesRecord struct {
	Data       []Data
	Notes      []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseOrderNotesRecord was measured from.
func (PurchaseOrderNotesRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/notes",
	}
}

// PurchaseOrderNoteRecord is the purchase order note body.
type PurchaseOrderNoteRecord struct {
	Attachments   []Attachment
	Content       *string
	CreatedBy     *CreatedBy
	CreatedOn     *string
	Group         any
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn any
	Note          map[string]any
	Object        *string
	PurchaseOrder *PurchaseOrder
	User          any
	Version       *int64
	Watchers      []Watcher
}

// MeasuredFrom lists the endpoints PurchaseOrderNoteRecord was measured from.
func (PurchaseOrderNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}",
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}/files",
		"POST /api/{api_version}/purchase_orders/{purchase_order_id}/notes",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (PurchaseOrderNoteRecord) EnvelopeKey() string { return NoteKey }

// PurchaseOrderLineReadRecord is the purchase order line read body.
type PurchaseOrderLineReadRecord struct {
	CreatedOn                    *string
	FlatDiscount                 *float64
	Href                         *string
	ID                           *string
	IsPriceTaxInclusive          *bool
	IsTaxExempt                  *bool
	Item                         *Item
	ItemAccountingCode           *ItemAccountingCode
	ItemDiscountAmount           *float64
	ItemExpectedCompletionDate   *string
	ItemInvoiceNote              any
	ItemPriceSnapshot            *ItemPriceSnapshot
	ItemPriceTax                 *ItemPriceTax
	ItemPriceTaxExempt           *bool
	ItemPurchaseAccountingCode   *ItemPurchaseAccountingCode
	ItemPurchaseTaxConfiguration *ItemPurchaseTaxConfiguration
	LastUpdatedOn                *string
	Line                         map[string]any
	Name                         *string
	Object                       *string
	Price                        *float64
	PurchaseLineNote             any
	Quantity                     *int64
	SortIndex                    *int64
	Status                       *string
	Subtotal                     *string
	Tax                          *Tax
	TaxTotal                     *string
	Total                        *string
	Uom                          *Uom
	Version                      *int64
	Warehouse                    *Warehouse
}

// MeasuredFrom lists the endpoints PurchaseOrderLineReadRecord was measured from.
func (PurchaseOrderLineReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/lines/{purchase_order_line_uuid}",
	}
}

// EnvelopeKey unwraps {"purchase_order_line": {...}}.
func (PurchaseOrderLineReadRecord) EnvelopeKey() string { return PurchaseOrderLineKey }

// PurchaseOrderNoteFileList is the page GetPurchaseOrderNoteFiles answers.
type PurchaseOrderNoteFileList struct {
	Data       []PurchaseOrderNoteFileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseOrderNoteFileList was measured from.
func (PurchaseOrderNoteFileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}/files",
	}
}

// PurchaseOrderNoteFileReadRecord is the body GetPurchaseOrderNoteFile answers.
type PurchaseOrderNoteFileReadRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints PurchaseOrderNoteFileReadRecord was measured from.
func (PurchaseOrderNoteFileReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}/files/{purchase_order_file_id}",
	}
}

// EnvelopeKey unwraps {"file": {...}}.
func (PurchaseOrderNoteFileReadRecord) EnvelopeKey() string { return FileKey }

// PurchaseOrderPurchaseInvoiceCreateRecord is the body CreatePurchaseOrderPurchaseInvoice answers.
type PurchaseOrderPurchaseInvoiceCreateRecord struct {
	Account            *Account
	ActivityLogs       []any
	AlternateDueDate   *string
	AlternateIssueDate *string
	CreatedBy          *CreatedBy
	CreatedOn          *string
	Currency           *Currency
	CustomAttributes   []any
	CustomForm         any
	DueDate            *string
	ID                 *string
	IssueDate          *string
	Kpis               *Kpis
	LastUpdatedBy      any
	LastUpdatedOn      any
	Lines              []Line
	Object             *string
	Origin             any
	PaymentStatus      *string
	PriceTaxInclusive  *bool
	PurchaseOrder      *PurchaseOrder
	Status             *string
	Subtotal           *string
	Tax                *string
	Total              *string
	UUID               *string
	Version            *int64
}

// MeasuredFrom lists the endpoints PurchaseOrderPurchaseInvoiceCreateRecord was measured from.
func (PurchaseOrderPurchaseInvoiceCreateRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_invoices",
	}
}

// EnvelopeKey unwraps {"purchase_invoice": {...}}.
func (PurchaseOrderPurchaseInvoiceCreateRecord) EnvelopeKey() string { return PurchaseInvoiceKey }

// PurchaseOrderNoteFileRecord is one row of PurchaseOrderNoteFileList.
type PurchaseOrderNoteFileRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints PurchaseOrderNoteFileRecord was measured from.
func (PurchaseOrderNoteFileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}/files",
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

// Currency is the currency object.
type Currency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Data is one entry of the data list.
type Data struct {
	Attachments   []Attachment
	Content       *string
	CreatedBy     *CreatedBy
	CreatedOn     *string
	Group         any
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn any
	Object        *string
	PurchaseOrder *PurchaseOrder
	User          any
	Version       *int64
	Watchers      []Watcher
}

// Item is the item object.
type Item struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// ItemAccountingCode is the item_accounting_code object.
type ItemAccountingCode struct {
	CostOfGoodsSold *string
}

// ItemPriceSnapshot is the item_price_snapshot object.
type ItemPriceSnapshot struct {
	PricingRule *PricingRule
}

// ItemPriceTax is the item_price_tax object.
type ItemPriceTax struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
	Rate   *float64
}

// ItemPurchaseAccountingCode is the item_purchase_accounting_code object.
type ItemPurchaseAccountingCode struct {
	CostOfGoodsSold *string
}

// ItemPurchaseTaxConfiguration is the item_purchase_tax_configuration object.
type ItemPurchaseTaxConfiguration struct {
	PurchasePriceIsTaxInclusive *bool
	TaxCode                     *TaxCode
	TaxExemptWhenPurchase       *bool
}

// Kpis is the kpis object.
type Kpis struct {
	CreditApplied     *string
	CreditIssued      *string
	DeletedOn         any
	LastAmendedOn     any
	LastCancelledOn   any
	LastPaymentDate   any
	LastReactivatedOn any
	Object            *string
	Outstanding       *string
	Overdue           *string
	PaymentApplied    *string
	VoidedOn          any
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// Line is one entry of the lines list.
type Line struct {
	AccountingCode *string
	ID             *string
	Item           *Item
	Object         *string
	Quantity       *int64
	Subtotal       *string
	Tax            *string
	Total          *string
	Uom            *string
	Version        *int64
	Warehouse      *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// PricingRule is the pricing_rule object.
type PricingRule struct {
	Price          *float64
	PriceType      *string
	PricingVersion *string
	Uom            *string
	UUID           *string
	Version        *string
}

// PurchaseOrder is the purchase_order object.
type PurchaseOrder struct {
	Href   *string
	ID     *string
	Object *string
}

// Tax is the tax object.
type Tax struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// TaxCode is the tax_code object.
type TaxCode struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
	Rate   *string
}

// Uom is the uom object.
type Uom struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Warehouse is the warehouse object.
type Warehouse struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Watcher is one entry of the watchers list.
type Watcher struct {
	Email *string
	Name  *string
	Type  *string
}

// CreditNoteApplication is one entry of the credit_note_applications list.
type CreditNoteApplication struct {
	Amount             *string
	PurchaseCreditNote *PurchaseCreditNote
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

// CustomForm is the custom_form object.
type CustomForm struct {
	Href   *string
	ID     *string
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

// PaymentMethodApplication is one entry of the payment_method_applications list.
type PaymentMethodApplication struct {
	Amount           *string
	PaymentMethod    *PaymentMethod
	PaymentProcessor *PaymentProcessor
	Reference        *string
}

// PaymentProcessor is the payment_processor object.
type PaymentProcessor struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PurchaseCreditNote is the purchase_credit_note object.
type PurchaseCreditNote struct {
	Code   *string
	Href   *string
	ID     *string
	Object *string
}

// PurchaseInvoice is the purchase_invoice object.
type PurchaseInvoice struct {
	Code   *string
	Href   *string
	ID     *string
	Object *string
}

// PurchaseInvoiceApplication is one entry of the purchase_invoice_applications list.
type PurchaseInvoiceApplication struct {
	Applied         *string
	DueDate         *string
	IssueDate       *string
	Outstanding     *string
	PurchaseInvoice *PurchaseInvoice
	Total           *string
}

// PurchaseOrderKpis is the kpis object.
type PurchaseOrderKpis struct {
	EstimatedTotal       *string
	IssueDate            any
	LastCancelledOn      any
	LastChangedOn        any
	LastDeletedOn        any
	LastInvoiceIssueDate any
	LastInvoiceTotal     *string
	LastReactivatedOn    any
	Object               *string
	TotalExpense         *string
	TotalOutstanding     *string
	TotalOverdue         *string
	TotalPurchaseInvoice *int64
}

// PurchaseOrderUpdatedKpis is the kpis object.
type PurchaseOrderUpdatedKpis struct {
	EstimatedTotal       *string
	IssueDate            any
	LastCancelledOn      any
	LastChangedOn        any
	LastDeletedOn        any
	LastInvoiceIssueDate any
	LastInvoiceTotal     *string
	LastReactivatedOn    any
	Object               *string
	TotalExpense         *string
	TotalOutstanding     *string
	TotalOverdue         *string
	TotalPurchaseInvoice *int64
}

// PurchaseOrdersKpis is the kpis object.
type PurchaseOrdersKpis struct {
	EstimatedTotal       *string
	IssueDate            any
	LastCancelledOn      any
	LastChangedOn        any
	LastDeletedOn        any
	LastInvoiceIssueDate any
	LastInvoiceTotal     *string
	LastReactivatedOn    any
	Object               *string
	TotalExpense         *string
	TotalOutstanding     *string
	TotalOverdue         *string
	TotalPurchaseInvoice *int64
}

// PurchaseOrderListShape reads a body as PurchaseOrderList.
func PurchaseOrderListShape(decoded any) any { return dto.Read[PurchaseOrderList](decoded) }

// PurchaseOrdersRecordShape reads a body as PurchaseOrdersRecord.
func PurchaseOrdersRecordShape(decoded any) any { return dto.Read[PurchaseOrdersRecord](decoded) }

// PurchaseOrdersNoteListShape reads a body as PurchaseOrdersNoteList.
func PurchaseOrdersNoteListShape(decoded any) any { return dto.Read[PurchaseOrdersNoteList](decoded) }

// PurchaseOrderUpdatedRecordShape reads a body as PurchaseOrderUpdatedRecord.
func PurchaseOrderUpdatedRecordShape(decoded any) any {
	return dto.Read[PurchaseOrderUpdatedRecord](decoded)
}

// PurchaseOrderReadRecordShape reads a body as PurchaseOrderReadRecord.
func PurchaseOrderReadRecordShape(decoded any) any { return dto.Read[PurchaseOrderReadRecord](decoded) }

// AccountPurchaseOrdersRecordShape reads a body as AccountPurchaseOrdersRecord.
func AccountPurchaseOrdersRecordShape(decoded any) any {
	return dto.Read[AccountPurchaseOrdersRecord](decoded)
}

// PurchaseOrderInformationRecordShape reads a body as PurchaseOrderInformationRecord.
func PurchaseOrderInformationRecordShape(decoded any) any {
	return dto.Read[PurchaseOrderInformationRecord](decoded)
}

// PurchaseOrderLineListShape reads a body as PurchaseOrderLineList.
func PurchaseOrderLineListShape(decoded any) any { return dto.Read[PurchaseOrderLineList](decoded) }

// PurchaseOrderNotesRecordShape reads a body as PurchaseOrderNotesRecord.
func PurchaseOrderNotesRecordShape(decoded any) any {
	return dto.Read[PurchaseOrderNotesRecord](decoded)
}

// PurchaseOrderPurchaseInvoiceListShape reads a body as PurchaseOrderPurchaseInvoiceList.
func PurchaseOrderPurchaseInvoiceListShape(decoded any) any {
	return dto.Read[PurchaseOrderPurchaseInvoiceList](decoded)
}

// PurchaseOrderPurchasePaymentListShape reads a body as PurchaseOrderPurchasePaymentList.
func PurchaseOrderPurchasePaymentListShape(decoded any) any {
	return dto.Read[PurchaseOrderPurchasePaymentList](decoded)
}

// PurchaseOrderLineReadRecordShape reads a body as PurchaseOrderLineReadRecord.
func PurchaseOrderLineReadRecordShape(decoded any) any {
	return dto.Read[PurchaseOrderLineReadRecord](decoded)
}

// PurchaseOrderNoteRecordShape reads a body as PurchaseOrderNoteRecord.
func PurchaseOrderNoteRecordShape(decoded any) any { return dto.Read[PurchaseOrderNoteRecord](decoded) }

// PurchaseOrderNoteFileListShape reads a body as PurchaseOrderNoteFileList.
func PurchaseOrderNoteFileListShape(decoded any) any {
	return dto.Read[PurchaseOrderNoteFileList](decoded)
}

// PurchaseOrderNoteFileReadRecordShape reads a body as PurchaseOrderNoteFileReadRecord.
func PurchaseOrderNoteFileReadRecordShape(decoded any) any {
	return dto.Read[PurchaseOrderNoteFileReadRecord](decoded)
}

// PurchaseOrderPurchaseInvoiceCreateRecordShape reads a body as PurchaseOrderPurchaseInvoiceCreateRecord.
func PurchaseOrderPurchaseInvoiceCreateRecordShape(decoded any) any {
	return dto.Read[PurchaseOrderPurchaseInvoiceCreateRecord](decoded)
}

// PurchaseOrderKpisRecord is what GetPurchaseOrderKPIs answers, under purchase_order_kpis: the
// block a purchase order carries as kpis.
type PurchaseOrderKpisRecord struct {
	EstimatedTotal       *string
	IssueDate            any
	LastCancelledOn      any
	LastChangedOn        any
	LastDeletedOn        any
	LastInvoiceIssueDate any
	LastInvoiceTotal     *string
	LastReactivatedOn    any
	Object               *string
	TotalExpense         *string
	TotalOutstanding     *string
	TotalOverdue         *string
	TotalPurchaseInvoice *int64
}

// MeasuredFrom lists the endpoints PurchaseOrderKpisRecord was measured from.
func (PurchaseOrderKpisRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/kpis",
	}
}

// EnvelopeKey unwraps {"purchase_order_kpis": {...}}.
func (PurchaseOrderKpisRecord) EnvelopeKey() string { return PurchaseOrderKPIsKey }

// PurchaseOrderKpisRecordShape reads a body as PurchaseOrderKpisRecord.
func PurchaseOrderKpisRecordShape(decoded any) any { return dto.Read[PurchaseOrderKpisRecord](decoded) }
