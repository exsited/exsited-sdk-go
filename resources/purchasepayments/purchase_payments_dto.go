package purchasepayments

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

// PurchasePaymentRecord is the purchase payment body.
type PurchasePaymentRecord struct {
	Account *Account
	Code    *string
	// CreatedBy is untyped: the API sends it as a string or an object.
	CreatedBy              *CreatedBy
	CreatedOn              *string
	CreditNoteApplications []CreditNoteApplication
	CreditNoteTotal        *string
	Currency               *Currency
	CustomAttributes       []any
	CustomObjects          []any
	Href                   *string
	ID                     *string
	// LastUpdatedBy is untyped: the API sends it as a string or an object.
	LastUpdatedBy               *LastUpdatedBy
	LastUpdatedOn               *string
	Object                      *string
	Origin                      *string
	PaidTotal                   *string
	PaymentDate                 *string
	PaymentMethodApplications   []PaymentMethodApplication
	PaymentNote                 *string
	PaymentReference            any
	PurchaseCreditApplied       []any
	PurchaseInvoiceApplications []PurchaseInvoiceApplication
	PurchaseInvoices            []any
	PurchaseOrderID             *string
	PurchasePaymentApplied      []any
	PurchasePaymentDate         *string
	PurchasePaymentNote         *string
	Status                      *string
	TotalApplied                *string
	UUID                        *string
	// Version is untyped: the API sends it as a string or an integer.
	Version any
}

// MeasuredFrom lists the endpoints PurchasePaymentRecord was measured from.
func (PurchasePaymentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_payments",
	}
}

// PurchasePaymentList is a page of PurchasePaymentRecord records.
type PurchasePaymentList struct {
	Data       []PurchasePaymentRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchasePaymentList was measured from.
func (PurchasePaymentList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_payments",
	}
}

// PurchasePaymentReadRecord is the purchase payment read body.
type PurchasePaymentReadRecord struct {
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

// MeasuredFrom lists the endpoints PurchasePaymentReadRecord was measured from.
func (PurchasePaymentReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_payments",
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_payments",
		"GET /api/{api_version}/purchase_payments/{purchase_payment_id}",
	}
}

// EnvelopeKey unwraps {"purchase_payment": {...}}.
func (PurchasePaymentReadRecord) EnvelopeKey() string { return PurchasePaymentKey }

// PurchasePaymentNoteRecord is the purchase payment note body.
type PurchasePaymentNoteRecord struct {
	Attachments     []Attachment
	Content         *string
	CreatedBy       *CreatedBy
	CreatedOn       *string
	Group           any
	Href            *string
	ID              *string
	LastUpdatedBy   any
	LastUpdatedOn   any
	Object          *string
	PurchasePayment *PurchasePayment
	User            any
	Version         *int64
	Watchers        []Watcher
}

// MeasuredFrom lists the endpoints PurchasePaymentNoteRecord was measured from.
func (PurchasePaymentNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_payments/{purchase_payment_id}/notes",
	}
}

// PurchasePaymentNoteList is a page of PurchasePaymentNoteRecord records.
type PurchasePaymentNoteList struct {
	Data       []PurchasePaymentNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchasePaymentNoteList was measured from.
func (PurchasePaymentNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_payments/{purchase_payment_id}/notes",
	}
}

// PurchasePaymentUpdatedRecord is the purchase payment updated body.
type PurchasePaymentUpdatedRecord struct {
	Account                     *Account
	ActivityLogs                []any
	Code                        *string
	CreatedBy                   *CreatedBy
	CreatedOn                   *string
	CreditNoteApplications      []any
	CreditNoteTotal             *string
	Currency                    *Currency
	CustomAttributes            []any
	CustomObjects               []any
	Href                        *string
	ID                          *string
	LastUpdatedBy               *LastUpdatedBy
	LastUpdatedOn               *string
	Object                      *string
	Origin                      *string
	PaidTotal                   *string
	PaymentDate                 *string
	PaymentMethodApplications   []PaymentMethodApplication
	PaymentNote                 *string
	PaymentReference            any
	PurchaseCreditApplied       []any
	PurchaseInvoiceApplications []PurchaseInvoiceApplication
	PurchaseInvoices            []any
	PurchaseOrderID             *string
	PurchasePaymentApplied      []any
	PurchasePaymentDate         *string
	PurchasePaymentNote         *string
	Status                      *string
	TotalApplied                *string
	UUID                        *string
	Version                     *int64
}

// MeasuredFrom lists the endpoints PurchasePaymentUpdatedRecord was measured from.
func (PurchasePaymentUpdatedRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/purchase_payments/{purchase_payment_id}",
	}
}

// EnvelopeKey unwraps {"purchase_payment": {...}}.
func (PurchasePaymentUpdatedRecord) EnvelopeKey() string { return PurchasePaymentKey }

// PurchasePaymentsRecord is the purchase payments body.
type PurchasePaymentsRecord struct {
	Account                     *Account
	ActivityLogs                []ActivityLog
	Code                        *string
	CreatedBy                   *CreatedBy
	CreatedOn                   *string
	CreditNoteApplications      []CreditNoteApplication
	CreditNoteTotal             *string
	Currency                    *Currency
	CustomAttributes            []any
	CustomObjects               []any
	Href                        *string
	ID                          *string
	LastUpdatedBy               *string
	LastUpdatedOn               *string
	Object                      *string
	Origin                      *string
	PaidTotal                   *string
	PaymentDate                 *string
	PaymentMethodApplications   []PaymentMethodApplication
	PaymentNote                 *string
	PaymentReference            any
	PurchaseCreditApplied       []any
	PurchaseInvoiceApplications []PurchaseInvoiceApplication
	PurchaseInvoices            []any
	PurchaseOrderID             *string
	PurchasePaymentApplied      []any
	PurchasePaymentDate         *string
	PurchasePaymentNote         *string
	Status                      *string
	TotalApplied                *string
	UUID                        *string
	Version                     *int64
}

// MeasuredFrom lists the endpoints PurchasePaymentsRecord was measured from.
func (PurchasePaymentsRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_payments",
		"POST /api/{api_version}/purchase_payments",
	}
}

// EnvelopeKey unwraps {"purchase_payment": {...}}.
func (PurchasePaymentsRecord) EnvelopeKey() string { return PurchasePaymentKey }

// AccountPurchasePaymentsRecord is the account purchase payments body.
type AccountPurchasePaymentsRecord struct {
	Data       []PurchasePaymentRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AccountPurchasePaymentsRecord was measured from.
func (AccountPurchasePaymentsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/purchase_payments",
	}
}

// PurchasePaymentNotesRecord is the purchase payment notes body.
type PurchasePaymentNotesRecord struct {
	Attachments     []Attachment
	Content         *string
	CreatedBy       *CreatedBy
	CreatedOn       *string
	Group           any
	Href            *string
	ID              *string
	LastUpdatedBy   any
	LastUpdatedOn   any
	Link            *string
	Name            *string
	Object          *string
	PurchasePayment *PurchasePayment
	User            any
	UUID            *string
	Version         *int64
	Watchers        []Watcher
}

// MeasuredFrom lists the endpoints PurchasePaymentNotesRecord was measured from.
func (PurchasePaymentNotesRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/purchase_payments/{purchase_payment_id}/notes",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (PurchasePaymentNotesRecord) EnvelopeKey() string { return NoteKey }

// PurchasePaymentNoteReadRecord is the purchase payment note read body.
type PurchasePaymentNoteReadRecord struct {
	Attachments     []Attachment
	Content         *string
	CreatedBy       *CreatedBy
	CreatedOn       *string
	Group           any
	Href            *string
	ID              *string
	LastUpdatedBy   any
	LastUpdatedOn   any
	Note            map[string]any
	Object          *string
	PurchasePayment *PurchasePayment
	User            any
	Version         *int64
	Watchers        []Watcher
}

// MeasuredFrom lists the endpoints PurchasePaymentNoteReadRecord was measured from.
func (PurchasePaymentNoteReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}",
		"GET /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}/files",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (PurchasePaymentNoteReadRecord) EnvelopeKey() string { return NoteKey }

// PurchasePaymentNoteFileList is the page GetPurchasePaymentNoteFiles answers.
type PurchasePaymentNoteFileList struct {
	Data       []PurchasePaymentNoteFileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PurchasePaymentNoteFileList was measured from.
func (PurchasePaymentNoteFileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}/files",
	}
}

// PurchasePaymentNoteFileReadRecord is the body GetPurchasePaymentNoteFile answers.
type PurchasePaymentNoteFileReadRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints PurchasePaymentNoteFileReadRecord was measured from.
func (PurchasePaymentNoteFileReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}/files/{file_id}",
	}
}

// EnvelopeKey unwraps {"file": {...}}.
func (PurchasePaymentNoteFileReadRecord) EnvelopeKey() string { return FileKey }

// PurchasePaymentNoteFileRecord is one row of PurchasePaymentNoteFileList.
type PurchasePaymentNoteFileRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints PurchasePaymentNoteFileRecord was measured from.
func (PurchasePaymentNoteFileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}/files",
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

// ActivityLog is one entry of the activity_logs list.
type ActivityLog struct {
	Activity *string
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

// CreditNoteApplication is one entry of the credit_note_applications list.
type CreditNoteApplication struct {
	Amount             *string
	PurchaseCreditNote *PurchaseCreditNote
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

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev any
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

// PurchasePayment is the purchase_payment object.
type PurchasePayment struct {
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

// PurchasePaymentListShape reads a body as PurchasePaymentList.
func PurchasePaymentListShape(decoded any) any { return dto.Read[PurchasePaymentList](decoded) }

// PurchasePaymentsRecordShape reads a body as PurchasePaymentsRecord.
func PurchasePaymentsRecordShape(decoded any) any { return dto.Read[PurchasePaymentsRecord](decoded) }

// PurchasePaymentReadRecordShape reads a body as PurchasePaymentReadRecord.
func PurchasePaymentReadRecordShape(decoded any) any {
	return dto.Read[PurchasePaymentReadRecord](decoded)
}

// PurchasePaymentUpdatedRecordShape reads a body as PurchasePaymentUpdatedRecord.
func PurchasePaymentUpdatedRecordShape(decoded any) any {
	return dto.Read[PurchasePaymentUpdatedRecord](decoded)
}

// AccountPurchasePaymentsRecordShape reads a body as AccountPurchasePaymentsRecord.
func AccountPurchasePaymentsRecordShape(decoded any) any {
	return dto.Read[AccountPurchasePaymentsRecord](decoded)
}

// PurchasePaymentNoteListShape reads a body as PurchasePaymentNoteList.
func PurchasePaymentNoteListShape(decoded any) any { return dto.Read[PurchasePaymentNoteList](decoded) }

// PurchasePaymentNotesRecordShape reads a body as PurchasePaymentNotesRecord.
func PurchasePaymentNotesRecordShape(decoded any) any {
	return dto.Read[PurchasePaymentNotesRecord](decoded)
}

// PurchasePaymentNoteReadRecordShape reads a body as PurchasePaymentNoteReadRecord.
func PurchasePaymentNoteReadRecordShape(decoded any) any {
	return dto.Read[PurchasePaymentNoteReadRecord](decoded)
}

// PurchasePaymentNoteFileListShape reads a body as PurchasePaymentNoteFileList.
func PurchasePaymentNoteFileListShape(decoded any) any {
	return dto.Read[PurchasePaymentNoteFileList](decoded)
}

// PurchasePaymentNoteFileReadRecordShape reads a body as PurchasePaymentNoteFileReadRecord.
func PurchasePaymentNoteFileReadRecordShape(decoded any) any {
	return dto.Read[PurchasePaymentNoteFileReadRecord](decoded)
}

// PurchasePaymentNoteUpdatedRecord is what UpdatePurchasePaymentNote answers, under note.
type PurchasePaymentNoteUpdatedRecord struct {
	Attachments     []PurchasePaymentNoteUpdatedAttachment
	Content         *string
	CreatedBy       *PurchasePaymentNoteUpdatedUser
	CreatedOn       *string
	Group           any
	Href            *string
	ID              *string
	LastUpdatedBy   *PurchasePaymentNoteUpdatedUser
	LastUpdatedOn   *string
	Object          *string
	PurchasePayment *PurchasePaymentNoteUpdatedPurchasePayment
	User            any
	Version         *int64
	Watchers        []any
}

// MeasuredFrom lists the endpoints PurchasePaymentNoteUpdatedRecord was measured from.
func (PurchasePaymentNoteUpdatedRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (PurchasePaymentNoteUpdatedRecord) EnvelopeKey() string { return NoteKey }

// PurchasePaymentNoteUpdatedAttachment is one entry of a PurchasePaymentNoteUpdatedRecord's
// attachments.
type PurchasePaymentNoteUpdatedAttachment struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// PurchasePaymentNoteUpdatedUser is a PurchasePaymentNoteUpdatedRecord's created_by and
// last_updated_by.
type PurchasePaymentNoteUpdatedUser struct {
	Name   *string
	Object *string
}

// PurchasePaymentNoteUpdatedPurchasePayment is a PurchasePaymentNoteUpdatedRecord's
// purchase_payment.
type PurchasePaymentNoteUpdatedPurchasePayment struct {
	Href   *string
	ID     *string
	Object *string
}

// PurchasePaymentNoteUpdatedRecordShape reads a body as PurchasePaymentNoteUpdatedRecord.
func PurchasePaymentNoteUpdatedRecordShape(decoded any) any {
	return dto.Read[PurchasePaymentNoteUpdatedRecord](decoded)
}
