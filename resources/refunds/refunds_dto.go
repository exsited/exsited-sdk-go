package refunds

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

// RefundRecord is the refund body.
type RefundRecord struct {
	Account          *Account
	Amount           *string
	CreatedBy        *string
	CreatedOn        *string
	CreditNote       *CreditNote
	CustomAttributes []any
	GatewayResponse  any
	ID               *string
	IssueDate        *string
	LastUpdatedBy    *string
	LastUpdatedOn    *string
	Note             any
	Object           *string
	PaymentMethod    *string
	PaymentProcessor *string
	ReconcileStatus  *string
	Reference        *string
	Status           *string
	UUID             *string
	Version          *int64
}

// MeasuredFrom lists the endpoints RefundRecord was measured from.
func (RefundRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/refunds",
		"GET /api/{api_version}/refunds/{refund_id}",
	}
}

// EnvelopeKey unwraps {"refund": {...}}.
func (RefundRecord) EnvelopeKey() string { return RefundKey }

// RefundList is a page of RefundRecord records.
type RefundList struct {
	Data       []RefundRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints RefundList was measured from.
func (RefundList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/refunds",
	}
}

// RefundUpdatedRecord is the refund updated body.
type RefundUpdatedRecord struct {
	Account          *Account
	Amount           *string
	CreatedBy        *string
	CreatedOn        *string
	CreditNote       *CreditNote
	CreditNoteID     *string
	CustomAttributes []any
	GatewayResponse  *string
	ID               *string
	IssueDate        *string
	LastUpdatedBy    *string
	LastUpdatedOn    *string
	Note             *string
	Object           *string
	PaymentMethod    *string
	PaymentProcessor *string
	ReconcileStatus  *string
	Reference        *string
	Status           *string
	UUID             *string
	Version          *int64
}

// MeasuredFrom lists the endpoints RefundUpdatedRecord was measured from.
func (RefundUpdatedRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/credit_notes/{credit_note_id}/refunds",
		"PATCH /api/{api_version}/refunds/{refund_id}",
	}
}

// EnvelopeKey unwraps {"refund": {...}}.
func (RefundUpdatedRecord) EnvelopeKey() string { return RefundKey }

// AccountRefundsRecord is the account refunds body.
type AccountRefundsRecord struct {
	Data       []AccountRefund
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AccountRefundsRecord was measured from.
func (AccountRefundsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/refunds",
	}
}

// Account is the account object.
type Account struct {
	Href   *string
	ID     *string
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

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// AccountRefund is one row of AccountRefundsRecord.
type AccountRefund struct {
	Account          *Account
	Amount           *string
	CreatedBy        *string
	CreatedOn        *string
	CreditNote       *CreditNote
	CustomAttributes []any
	GatewayResponse  any
	ID               *string
	IssueDate        *string
	LastUpdatedBy    any
	LastUpdatedOn    any
	Note             any
	Object           *string
	PaymentMethod    *string
	PaymentProcessor *string
	ReconcileStatus  *string
	Reference        *string
	Status           *string
	UUID             *string
	Version          *int64
}

// RefundListShape reads a body as RefundList.
func RefundListShape(decoded any) any { return dto.Read[RefundList](decoded) }

// RefundRecordShape reads a body as RefundRecord.
func RefundRecordShape(decoded any) any { return dto.Read[RefundRecord](decoded) }

// RefundUpdatedRecordShape reads a body as RefundUpdatedRecord.
func RefundUpdatedRecordShape(decoded any) any { return dto.Read[RefundUpdatedRecord](decoded) }

// AccountRefundsRecordShape reads a body as AccountRefundsRecord.
func AccountRefundsRecordShape(decoded any) any { return dto.Read[AccountRefundsRecord](decoded) }

// RefundNoteList is one page of RefundNoteRecord records, under data.
type RefundNoteList struct {
	Data       []RefundNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints RefundNoteList was measured from.
func (RefundNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/refunds/{refund_id}/notes",
	}
}

// RefundNoteRecord is one note on a refund, under note.
type RefundNoteRecord struct {
	Attachments   []RefundNoteAttachment
	Content       *string
	CreatedBy     *RefundNoteCreatedBy
	CreatedOn     *string
	Group         any
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn any
	Object        *string
	Refund        *RefundNoteDetail
	User          any
	Version       *int64
	Watchers      []RefundNoteWatcher
}

// MeasuredFrom lists the endpoints RefundNoteRecord was measured from.
func (RefundNoteRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/refunds/{refund_id}/notes",
		"GET /api/{api_version}/refunds/{refund_id}/notes",
		"GET /api/{api_version}/refunds/{refund_id}/notes/{note_id}",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (RefundNoteRecord) EnvelopeKey() string { return NoteKey }

// RefundNoteAttachment is one entry of a RefundNoteRecord's attachments.
type RefundNoteAttachment struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// RefundNoteCreatedBy is a RefundNoteRecord's created_by.
type RefundNoteCreatedBy struct {
	Name   *string
	Object *string
}

// RefundNoteDetail is a RefundNoteRecord's refund.
type RefundNoteDetail struct {
	Href   *string
	ID     *string
	Object *string
}

// RefundNoteWatcher is one entry of a RefundNoteRecord's watchers.
type RefundNoteWatcher struct {
	Email any
	Name  *string
	Type  *string
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

// RefundNoteFileRecord is one attachment's descriptor, under file.
type RefundNoteFileRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints RefundNoteFileRecord was measured from.
func (RefundNoteFileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/refunds/{refund_id}/notes/{note_id}/files/{file_id}",
	}
}

// EnvelopeKey unwraps {"file": {...}}.
func (RefundNoteFileRecord) EnvelopeKey() string { return FileKey }

// RefundNoteListShape reads a body as RefundNoteList.
func RefundNoteListShape(decoded any) any { return dto.Read[RefundNoteList](decoded) }

// RefundNoteRecordShape reads a body as RefundNoteRecord.
func RefundNoteRecordShape(decoded any) any { return dto.Read[RefundNoteRecord](decoded) }

// NoteFileListShape reads a body as NoteFileList.
func NoteFileListShape(decoded any) any { return dto.Read[NoteFileList](decoded) }

// RefundNoteFileRecordShape reads a body as RefundNoteFileRecord.
func RefundNoteFileRecordShape(decoded any) any { return dto.Read[RefundNoteFileRecord](decoded) }
