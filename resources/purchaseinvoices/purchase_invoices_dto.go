package purchaseinvoices

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

// PurchaseInvoiceRecord is the purchase invoice body.
type PurchaseInvoiceRecord struct {
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
	Origin             *string
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

// MeasuredFrom lists the endpoints PurchaseInvoiceRecord was measured from.
func (PurchaseInvoiceRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices",
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/information",
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_invoices",
	}
}

// EnvelopeKey unwraps {"purchase_invoice": {...}}.
func (PurchaseInvoiceRecord) EnvelopeKey() string { return PurchaseInvoiceKey }

// PurchaseInvoiceList is a page of PurchaseInvoiceRecord records.
type PurchaseInvoiceList struct {
	Data       []PurchaseInvoiceRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseInvoiceList was measured from.
func (PurchaseInvoiceList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices",
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_invoices",
	}
}

// PurchaseInvoicePurchasePaymentRecord is the purchase invoice purchase payment body.
type PurchaseInvoicePurchasePaymentRecord struct {
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
	LastUpdatedOn               any
	Object                      *string
	Origin                      *string
	PaidTotal                   *string
	PaymentDate                 *string
	PaymentMethodApplications   []PaymentMethodApplication
	PaymentNote                 *string
	PaymentReference            any
	PurchaseInvoiceApplications []PurchaseInvoiceApplication
	Status                      *string
	Version                     *int64
}

// MeasuredFrom lists the endpoints PurchaseInvoicePurchasePaymentRecord was measured from.
func (PurchaseInvoicePurchasePaymentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_payments",
	}
}

// PurchaseInvoicePurchasePaymentList is a page of PurchaseInvoicePurchasePaymentRecord
// records.
type PurchaseInvoicePurchasePaymentList struct {
	Data       []PurchaseInvoicePurchasePaymentRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseInvoicePurchasePaymentList was measured from.
func (PurchaseInvoicePurchasePaymentList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_payments",
	}
}

// PurchaseInvoiceNoteRecord is the purchase invoice note body.
type PurchaseInvoiceNoteRecord struct {
	Attachments     []Attachment
	Content         *string
	CreatedBy       *CreatedBy
	CreatedOn       *string
	Group           any
	Href            *string
	ID              *string
	LastUpdatedBy   *LastUpdatedBy
	LastUpdatedOn   any
	Object          *string
	PurchaseInvoice *PurchaseInvoice
	User            any
	Version         *int64
	Watchers        []Watcher
}

// MeasuredFrom lists the endpoints PurchaseInvoiceNoteRecord was measured from.
func (PurchaseInvoiceNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes",
	}
}

// PurchaseInvoiceNoteList is a page of PurchaseInvoiceNoteRecord records.
type PurchaseInvoiceNoteList struct {
	Data       []PurchaseInvoiceNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseInvoiceNoteList was measured from.
func (PurchaseInvoiceNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes",
	}
}

// PurchaseInvoicesNoteRecord is the purchase invoices note body.
type PurchaseInvoicesNoteRecord struct {
	Attachments   []any
	Content       *string
	CreatedBy     *CreatedBy
	CreatedOn     *string
	Group         any
	Href          *string
	ID            *string
	LastUpdatedBy *LastUpdatedBy
	LastUpdatedOn any
	Object        *string
	User          any
	Version       *int64
	Watchers      []Watcher
}

// MeasuredFrom lists the endpoints PurchaseInvoicesNoteRecord was measured from.
func (PurchaseInvoicesNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/notes",
	}
}

// PurchaseInvoicesNoteList is a page of PurchaseInvoicesNoteRecord records.
type PurchaseInvoicesNoteList struct {
	Data       []PurchaseInvoicesNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseInvoicesNoteList was measured from.
func (PurchaseInvoicesNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/notes",
	}
}

// PurchaseInvoiceReadRecord is the purchase invoice read body.
type PurchaseInvoiceReadRecord struct {
	Account            *Account
	ActivityLogs       []any
	AlternateDueDate   *string
	AlternateIssueDate *string
	CreatedBy          *CreatedBy
	CreatedOn          *string
	Currency           *Currency
	CustomAttributes   []CustomAttribute
	CustomForm         any
	DueDate            *string
	ID                 *string
	IssueDate          *string
	Kpis               *Kpis
	LastUpdatedBy      *LastUpdatedBy
	LastUpdatedOn      *string
	Lines              []Line
	Object             *string
	Origin             *string
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

// MeasuredFrom lists the endpoints PurchaseInvoiceReadRecord was measured from.
func (PurchaseInvoiceReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}",
		"POST /api/{api_version}/purchase_invoices",
		"POST /api/{api_version}/purchase_invoices/{purchase_invoice_id}/amend",
		"POST /api/{api_version}/purchase_invoices/{purchase_invoice_id}/reactivate",
		"POST /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_invoices",
	}
}

// EnvelopeKey unwraps {"purchase_invoice": {...}}.
func (PurchaseInvoiceReadRecord) EnvelopeKey() string { return PurchaseInvoiceKey }

// AccountPurchaseInvoicesRecord is the account purchase invoices body.
type AccountPurchaseInvoicesRecord struct {
	Data       []AccountPurchaseInvoice
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AccountPurchaseInvoicesRecord was measured from.
func (AccountPurchaseInvoicesRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/purchase_invoices",
	}
}

// PurchaseInvoiceInformationRecord is the purchase invoice information body.
type PurchaseInvoiceInformationRecord struct {
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
	InvoiceNote        *string
	IssueDate          *string
	Kpis               *Kpis
	LastUpdatedBy      *LastUpdatedBy
	LastUpdatedOn      *string
	Lines              []Line
	Object             *string
	Origin             *string
	PaymentStatus      *string
	PriceTaxInclusive  *bool
	PurchaseOrder      any
	Status             *string
	Subtotal           *string
	Tax                *string
	Total              *string
	UUID               *string
	Version            *int64
}

// MeasuredFrom lists the endpoints PurchaseInvoiceInformationRecord was measured from.
func (PurchaseInvoiceInformationRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/purchase_invoices/{purchase_invoice_id}/information",
	}
}

// EnvelopeKey unwraps {"purchase_invoice": {...}}.
func (PurchaseInvoiceInformationRecord) EnvelopeKey() string { return PurchaseInvoiceKey }

// PurchaseInvoiceCancelRecord is the purchase invoice cancel body.
type PurchaseInvoiceCancelRecord struct {
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
	EventUUID          *string
	ID                 *string
	IssueDate          *string
	Kpis               *Kpis
	LastUpdatedBy      *LastUpdatedBy
	LastUpdatedOn      *string
	Lines              []Line
	Object             *string
	Origin             *string
	PaymentStatus      *string
	PriceTaxInclusive  *bool
	PurchaseOrder      any
	Status             *string
	Subtotal           *string
	Tax                *string
	Total              *string
	UUID               *string
	Version            *int64
}

// MeasuredFrom lists the endpoints PurchaseInvoiceCancelRecord was measured from.
func (PurchaseInvoiceCancelRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/purchase_invoices/{purchase_invoice_id}/cancel",
	}
}

// EnvelopeKey unwraps {"purchase_invoice": {...}}.
func (PurchaseInvoiceCancelRecord) EnvelopeKey() string { return PurchaseInvoiceKey }

// PurchaseInvoiceNotesRecord is the purchase invoice notes body.
type PurchaseInvoiceNotesRecord struct {
	Attachments      []Attachment
	Content          *string
	CreatedBy        *CreatedBy
	CreatedOn        *string
	ExternalWatchers []any
	Group            any
	Href             *string
	ID               *string
	InternalWatchers []any
	LastUpdatedBy    any
	LastUpdatedOn    any
	Object           *string
	PurchaseInvoice  *PurchaseInvoice
	User             any
	Version          *int64
	Watchers         []Watcher
}

// MeasuredFrom lists the endpoints PurchaseInvoiceNotesRecord was measured from.
func (PurchaseInvoiceNotesRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (PurchaseInvoiceNotesRecord) EnvelopeKey() string { return NoteKey }

// PurchaseInvoiceLinesRecord is the purchase invoice lines body.
type PurchaseInvoiceLinesRecord struct {
	Lines  []Line
	Object *string
}

// MeasuredFrom lists the endpoints PurchaseInvoiceLinesRecord was measured from.
func (PurchaseInvoiceLinesRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/lines",
	}
}

// EnvelopeKey unwraps {"purchase_invoice": {...}}.
func (PurchaseInvoiceLinesRecord) EnvelopeKey() string { return PurchaseInvoiceKey }

// PurchaseInvoiceLineRecord is the purchase invoice line body.
type PurchaseInvoiceLineRecord struct {
	Line   *Line
	Object *string
}

// MeasuredFrom lists the endpoints PurchaseInvoiceLineRecord was measured from.
func (PurchaseInvoiceLineRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/lines/{purchase_invoice_line_uuid}",
	}
}

// EnvelopeKey unwraps {"purchase_invoice": {...}}.
func (PurchaseInvoiceLineRecord) EnvelopeKey() string { return PurchaseInvoiceKey }

// PurchaseInvoiceNoteReadRecord is the purchase invoice note read body.
type PurchaseInvoiceNoteReadRecord struct {
	Attachments     []Attachment
	Content         *string
	CreatedBy       *CreatedBy
	CreatedOn       *string
	Group           any
	Href            *string
	ID              *string
	LastUpdatedBy   *LastUpdatedBy
	LastUpdatedOn   *string
	Note            map[string]any
	Object          *string
	PurchaseInvoice *PurchaseInvoice
	User            any
	Version         *int64
	Watchers        []Watcher
}

// MeasuredFrom lists the endpoints PurchaseInvoiceNoteReadRecord was measured from.
func (PurchaseInvoiceNoteReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}",
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}/files",
		"PATCH /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (PurchaseInvoiceNoteReadRecord) EnvelopeKey() string { return NoteKey }

// PurchaseInvoiceNoteFileList is the page GetPurchaseInvoiceNoteFiles answers.
type PurchaseInvoiceNoteFileList struct {
	Data       []PurchaseInvoiceNoteFileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseInvoiceNoteFileList was measured from.
func (PurchaseInvoiceNoteFileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}/files",
	}
}

// PurchaseInvoiceNoteFileReadRecord is the body GetPurchaseInvoiceNoteFile answers.
type PurchaseInvoiceNoteFileReadRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints PurchaseInvoiceNoteFileReadRecord was measured from.
func (PurchaseInvoiceNoteFileReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}/files/{file_id}",
	}
}

// EnvelopeKey unwraps {"file": {...}}.
func (PurchaseInvoiceNoteFileReadRecord) EnvelopeKey() string { return FileKey }

// PurchaseInvoiceNoteFileRecord is one row of PurchaseInvoiceNoteFileList.
type PurchaseInvoiceNoteFileRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints PurchaseInvoiceNoteFileRecord was measured from.
func (PurchaseInvoiceNoteFileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}/files",
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

// Item is the item object.
type Item struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
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

// PurchaseInvoice is the purchase_invoice object.
type PurchaseInvoice struct {
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

// PurchaseInvoicePurchaseCreditNoteApplicationList is the page GetPurchaseInvoicePurchaseCreditNoteApplications answers.
type PurchaseInvoicePurchaseCreditNoteApplicationList struct {
	Data       []PurchaseInvoicePurchaseCreditNoteApplicationRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchaseInvoicePurchaseCreditNoteApplicationList was measured from.
func (PurchaseInvoicePurchaseCreditNoteApplicationList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_credit_note_applications",
	}
}

// PurchaseInvoicePurchaseCreditNoteApplicationRecord is one row of PurchaseInvoicePurchaseCreditNoteApplicationList.
type PurchaseInvoicePurchaseCreditNoteApplicationRecord struct {
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

// MeasuredFrom lists the endpoints PurchaseInvoicePurchaseCreditNoteApplicationRecord was measured from.
func (PurchaseInvoicePurchaseCreditNoteApplicationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_credit_note_applications",
	}
}

// AccountPurchaseInvoice is one row of AccountPurchaseInvoicesRecord.
type AccountPurchaseInvoice struct {
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
	Origin             *string
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
	Name   *string
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

// PurchaseOrder is the purchase_order object.
type PurchaseOrder struct {
	Href   *string
	ID     *string
	Object *string
}

// PurchasePayment is the purchase_payment object.
type PurchasePayment struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PurchaseInvoiceListShape reads a body as PurchaseInvoiceList.
func PurchaseInvoiceListShape(decoded any) any { return dto.Read[PurchaseInvoiceList](decoded) }

// PurchaseInvoicesNoteListShape reads a body as PurchaseInvoicesNoteList.
func PurchaseInvoicesNoteListShape(decoded any) any {
	return dto.Read[PurchaseInvoicesNoteList](decoded)
}

// PurchaseInvoiceReadRecordShape reads a body as PurchaseInvoiceReadRecord.
func PurchaseInvoiceReadRecordShape(decoded any) any {
	return dto.Read[PurchaseInvoiceReadRecord](decoded)
}

// AccountPurchaseInvoicesRecordShape reads a body as AccountPurchaseInvoicesRecord.
func AccountPurchaseInvoicesRecordShape(decoded any) any {
	return dto.Read[AccountPurchaseInvoicesRecord](decoded)
}

// PurchaseInvoiceCancelRecordShape reads a body as PurchaseInvoiceCancelRecord.
func PurchaseInvoiceCancelRecordShape(decoded any) any {
	return dto.Read[PurchaseInvoiceCancelRecord](decoded)
}

// PurchaseInvoiceRecordShape reads a body as PurchaseInvoiceRecord.
func PurchaseInvoiceRecordShape(decoded any) any { return dto.Read[PurchaseInvoiceRecord](decoded) }

// PurchaseInvoiceInformationRecordShape reads a body as PurchaseInvoiceInformationRecord.
func PurchaseInvoiceInformationRecordShape(decoded any) any {
	return dto.Read[PurchaseInvoiceInformationRecord](decoded)
}

// PurchaseInvoiceLinesRecordShape reads a body as PurchaseInvoiceLinesRecord.
func PurchaseInvoiceLinesRecordShape(decoded any) any {
	return dto.Read[PurchaseInvoiceLinesRecord](decoded)
}

// PurchaseInvoiceNoteListShape reads a body as PurchaseInvoiceNoteList.
func PurchaseInvoiceNoteListShape(decoded any) any { return dto.Read[PurchaseInvoiceNoteList](decoded) }

// PurchaseInvoiceNotesRecordShape reads a body as PurchaseInvoiceNotesRecord.
func PurchaseInvoiceNotesRecordShape(decoded any) any {
	return dto.Read[PurchaseInvoiceNotesRecord](decoded)
}

// PurchaseInvoicePurchasePaymentListShape reads a body as PurchaseInvoicePurchasePaymentList.
func PurchaseInvoicePurchasePaymentListShape(decoded any) any {
	return dto.Read[PurchaseInvoicePurchasePaymentList](decoded)
}

// PurchaseInvoiceLineRecordShape reads a body as PurchaseInvoiceLineRecord.
func PurchaseInvoiceLineRecordShape(decoded any) any {
	return dto.Read[PurchaseInvoiceLineRecord](decoded)
}

// PurchaseInvoiceNoteReadRecordShape reads a body as PurchaseInvoiceNoteReadRecord.
func PurchaseInvoiceNoteReadRecordShape(decoded any) any {
	return dto.Read[PurchaseInvoiceNoteReadRecord](decoded)
}

// PurchaseInvoiceNoteFileListShape reads a body as PurchaseInvoiceNoteFileList.
func PurchaseInvoiceNoteFileListShape(decoded any) any {
	return dto.Read[PurchaseInvoiceNoteFileList](decoded)
}

// PurchaseInvoiceNoteFileReadRecordShape reads a body as PurchaseInvoiceNoteFileReadRecord.
func PurchaseInvoiceNoteFileReadRecordShape(decoded any) any {
	return dto.Read[PurchaseInvoiceNoteFileReadRecord](decoded)
}

// PurchaseInvoicePurchaseCreditNoteApplicationListShape reads a body as PurchaseInvoicePurchaseCreditNoteApplicationList.
func PurchaseInvoicePurchaseCreditNoteApplicationListShape(decoded any) any {
	return dto.Read[PurchaseInvoicePurchaseCreditNoteApplicationList](decoded)
}
