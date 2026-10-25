package payments

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

// PaymentRecord is the payment body.
type PaymentRecord struct {
	Account                *Account
	CreatedBy              *CreatedBy
	CreatedOn              *string
	CreditApplied          []CreditApplied
	GiftCertificateApplied []GiftCertificateApplied
	ID                     *string
	Invoices               []Invoice
	LastUpdatedBy          *LastUpdatedBy
	LastUpdatedOn          *string
	Object                 *string
	Origin                 any
	PaymentApplied         []PaymentApplied
	PaymentNote            *string
	ReconcileStatus        *string
	Status                 *string
	Version                *int64
}

// MeasuredFrom lists the endpoints PaymentRecord was measured from.
func (PaymentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/payments",
		"GET /api/{api_version}/orders/{order_id}/payments",
		"GET /api/{api_version}/payments",
	}
}

// PaymentList is a page of PaymentRecord records.
type PaymentList struct {
	Data       []PaymentRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PaymentList was measured from.
func (PaymentList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/payments",
		"GET /api/{api_version}/orders/{order_id}/payments",
		"GET /api/{api_version}/payments",
	}
}

// PaymentNoteRecord is the payment note body.
type PaymentNoteRecord struct {
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
	Payment       *Payment
	User          any
	Version       *int64
	Watchers      []Watcher
}

// MeasuredFrom lists the endpoints PaymentNoteRecord was measured from.
func (PaymentNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/payments/{payment_id}/notes",
	}
}

// PaymentNoteList is a page of PaymentNoteRecord records.
type PaymentNoteList struct {
	Data       []PaymentNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PaymentNoteList was measured from.
func (PaymentNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/payments/{payment_id}/notes",
	}
}

// PaymentsNoteRecord is the payments note body.
type PaymentsNoteRecord struct {
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
	User          any
	Version       *int64
	Watchers      []Watcher
}

// MeasuredFrom lists the endpoints PaymentsNoteRecord was measured from.
func (PaymentsNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/payments/notes",
	}
}

// PaymentsNoteList is a page of PaymentsNoteRecord records.
type PaymentsNoteList struct {
	Data       []PaymentsNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints PaymentsNoteList was measured from.
func (PaymentsNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/payments/notes",
	}
}

// PaymentReadRecord is the payment read body.
type PaymentReadRecord struct {
	Account                *Account
	ActivityLogs           []any
	CreatedBy              *CreatedBy
	CreatedOn              *string
	CreditApplied          []CreditApplied
	CustomAttributes       []CustomAttribute
	GiftCertificateApplied []GiftCertificateApplied
	ID                     *string
	Invoices               []Invoice
	LastUpdatedBy          *LastUpdatedBy
	LastUpdatedOn          *string
	Object                 *string
	Origin                 any
	PaymentApplied         []PaymentApplied
	PaymentNote            *string
	ReconcileStatus        *string
	Status                 *string
	Version                *int64
}

// MeasuredFrom lists the endpoints PaymentReadRecord was measured from.
func (PaymentReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/payments/{payment_id}",
		"PATCH /api/{api_version}/payments/{payment_id}",
	}
}

// EnvelopeKey unwraps {"payment": {...}}.
func (PaymentReadRecord) EnvelopeKey() string { return PaymentKey }

// AccountPaymentsRecord is the account payments body.
type AccountPaymentsRecord struct {
	Data       []PaymentRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AccountPaymentsRecord was measured from.
func (AccountPaymentsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/payments",
	}
}

// PaymentNoteReadRecord is the payment note read body.
type PaymentNoteReadRecord struct {
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
	Payment       *Payment
	User          any
	Version       *int64
	Watchers      []Watcher
}

// MeasuredFrom lists the endpoints PaymentNoteReadRecord was measured from.
func (PaymentNoteReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/payments/{payment_id}/notes/{payment_note_id}",
		"POST /api/{api_version}/payments/{payment_id}/notes",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (PaymentNoteReadRecord) EnvelopeKey() string { return NoteKey }

// PaymentNoteFileReadRecord is the body GetPaymentNoteFile answers.
type PaymentNoteFileReadRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints PaymentNoteFileReadRecord was measured from.
func (PaymentNoteFileReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/payments/{payment_id}/notes/{payment_note_id}/files/{payment_file_id}",
	}
}

// EnvelopeKey unwraps {"file": {...}}.
func (PaymentNoteFileReadRecord) EnvelopeKey() string { return FileKey }

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

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// Payment is the payment object.
type Payment struct {
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

// Account is the account object.
type Account struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// CreditApplied is one entry of the credit_applied list.
type CreditApplied struct {
	Amount *string
	Code   *string
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

// GiftCertificateApplied is one entry of the gift_certificate_applied list.
type GiftCertificateApplied struct {
	Amount *string
	Code   *string
}

// Invoice is one entry of the invoices list.
type Invoice struct {
	Applied     *string
	DueDate     *string
	ID          *string
	IssueDate   *string
	Outstanding *string
	Total       *string
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// PaymentApplied is one entry of the payment_applied list.
type PaymentApplied struct {
	Amount          *string
	GatewayResponse any
	Method          *string
	Processor       *string
	Reference       *string
}

// PaymentListShape reads a body as PaymentList.
func PaymentListShape(decoded any) any { return dto.Read[PaymentList](decoded) }

// PaymentsNoteListShape reads a body as PaymentsNoteList.
func PaymentsNoteListShape(decoded any) any { return dto.Read[PaymentsNoteList](decoded) }

// PaymentReadRecordShape reads a body as PaymentReadRecord.
func PaymentReadRecordShape(decoded any) any { return dto.Read[PaymentReadRecord](decoded) }

// AccountPaymentsRecordShape reads a body as AccountPaymentsRecord.
func AccountPaymentsRecordShape(decoded any) any { return dto.Read[AccountPaymentsRecord](decoded) }

// PaymentNoteListShape reads a body as PaymentNoteList.
func PaymentNoteListShape(decoded any) any { return dto.Read[PaymentNoteList](decoded) }

// PaymentNoteReadRecordShape reads a body as PaymentNoteReadRecord.
func PaymentNoteReadRecordShape(decoded any) any { return dto.Read[PaymentNoteReadRecord](decoded) }

// PaymentNoteFileReadRecordShape reads a body as PaymentNoteFileReadRecord.
func PaymentNoteFileReadRecordShape(decoded any) any {
	return dto.Read[PaymentNoteFileReadRecord](decoded)
}

// NoteFileList is A listing of one note's attachments - the same envelope, different records.
type NoteFileList struct {
	Object     *string
	URL        *string
	Data       []NoteAttachment
	Pagination *Pagination
}

// NoteAttachment is A file attached to a note.
type NoteAttachment struct {
	UUID        *string
	Filename    *string
	ContentType *string
	Size        *int64
	Href        *string
}

// NoteFileListShape reads a body as NoteFileList.
func NoteFileListShape(decoded any) any { return dto.Read[NoteFileList](decoded) }
