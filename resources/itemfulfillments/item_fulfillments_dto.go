package itemfulfillments

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

// ItemFulfillmentRecord is the item fulfillment body.
type ItemFulfillmentRecord struct {
	Account *Account
	Code    *string
	// CreatedBy is untyped: the API sends it as a string or an object.
	CreatedBy            *CreatedBy
	CreatedOn            *string
	CustomAttributes     map[string]any
	Date                 *string
	DeliveryStatus       any
	FulfillmentCompleted *bool
	FulfillmentDate      *string
	FulfillmentNote      any
	FulfillmentStatus    *string
	Fulfillments         []Fulfillment
	Href                 *string
	ID                   *string
	Invoice              *Invoice
	InvoiceID            *string
	// LastUpdatedBy is untyped: the API sends it as a string or an object.
	LastUpdatedBy  *LastUpdatedBy
	LastUpdatedOn  *string
	Note           *string
	Object         *string
	Order          any
	OrderID        *string
	Status         *string
	TrackingNumber *string
	UUID           *string
	// Version is untyped: the API sends it as a string or an integer.
	Version any
}

// MeasuredFrom lists the endpoints ItemFulfillmentRecord was measured from.
func (ItemFulfillmentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/item_fulfillments",
	}
}

// ItemFulfillmentList is a page of ItemFulfillmentRecord records.
type ItemFulfillmentList struct {
	Data       []ItemFulfillmentRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ItemFulfillmentList was measured from.
func (ItemFulfillmentList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/item_fulfillments",
	}
}

// ItemFulfillmentReadRecord is the item fulfillment read body.
type ItemFulfillmentReadRecord struct {
	Account              *Account
	Code                 *string
	CreatedBy            *CreatedBy
	CreatedOn            *string
	CustomAttributes     map[string]any
	DeliveryStatus       any
	FulfillmentCompleted *bool
	FulfillmentDate      *string
	FulfillmentNote      any
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

// MeasuredFrom lists the endpoints ItemFulfillmentReadRecord was measured from.
func (ItemFulfillmentReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/invoices/{invoice_id}/item_fulfillments",
		"GET /api/{api_version}/invoices/{invoice_id}/item_fulfillments/{fulfillment_id}",
		"GET /api/{api_version}/item_fulfillments/{fulfillment_id}",
	}
}

// EnvelopeKey unwraps {"item_fulfillment": {...}}.
func (ItemFulfillmentReadRecord) EnvelopeKey() string { return ItemFulfillmentKey }

// ItemFulfillmentsRecord is the item fulfillments body.
type ItemFulfillmentsRecord struct {
	ActivityLogs      []any
	Code              *string
	FulfillmentDate   *string
	FulfillmentNote   any
	FulfillmentStatus *string
	Href              *string
	ID                *string
	ItemFulfillment   map[string]any
	Object            *string
	Status            *string
	Version           *int64
}

// MeasuredFrom lists the endpoints ItemFulfillmentsRecord was measured from.
func (ItemFulfillmentsRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/item_fulfillments",
	}
}

// EnvelopeKey unwraps {"item_fulfillment": {...}}.
func (ItemFulfillmentsRecord) EnvelopeKey() string { return ItemFulfillmentKey }

// ItemFulfillmentUpdatedRecord is the item fulfillment updated body.
type ItemFulfillmentUpdatedRecord struct {
	Code    *string
	Href    *string
	ID      *string
	Note    *string
	Object  *string
	Status  *string
	Version *int64
}

// MeasuredFrom lists the endpoints ItemFulfillmentUpdatedRecord was measured from.
func (ItemFulfillmentUpdatedRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/item_fulfillments/{fulfillment_id}",
	}
}

// EnvelopeKey unwraps {"item_fulfillment": {...}}.
func (ItemFulfillmentUpdatedRecord) EnvelopeKey() string { return ItemFulfillmentKey }

// ItemFulfillmentNoteCreateRecord is the body CreateItemFulfillmentNote answers.
type ItemFulfillmentNoteCreateRecord struct {
	Href            *string
	ID              *string
	ItemFulfillment *ItemFulfillment
	Object          *string
}

// MeasuredFrom lists the endpoints ItemFulfillmentNoteCreateRecord was measured from.
func (ItemFulfillmentNoteCreateRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/item_fulfillments/{fulfillment_id}/notes",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (ItemFulfillmentNoteCreateRecord) EnvelopeKey() string { return NoteKey }

// ItemFulfillment is the item_fulfillment object.
type ItemFulfillment struct {
	Href   *string
	ID     *string
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

// ItemFulfillmentListShape reads a body as ItemFulfillmentList.
func ItemFulfillmentListShape(decoded any) any { return dto.Read[ItemFulfillmentList](decoded) }

// ItemFulfillmentsRecordShape reads a body as ItemFulfillmentsRecord.
func ItemFulfillmentsRecordShape(decoded any) any { return dto.Read[ItemFulfillmentsRecord](decoded) }

// ItemFulfillmentReadRecordShape reads a body as ItemFulfillmentReadRecord.
func ItemFulfillmentReadRecordShape(decoded any) any {
	return dto.Read[ItemFulfillmentReadRecord](decoded)
}

// ItemFulfillmentUpdatedRecordShape reads a body as ItemFulfillmentUpdatedRecord.
func ItemFulfillmentUpdatedRecordShape(decoded any) any {
	return dto.Read[ItemFulfillmentUpdatedRecord](decoded)
}

// ItemFulfillmentNoteCreateRecordShape reads a body as ItemFulfillmentNoteCreateRecord.
func ItemFulfillmentNoteCreateRecordShape(decoded any) any {
	return dto.Read[ItemFulfillmentNoteCreateRecord](decoded)
}

// ItemFulfillmentNoteList is one page of ItemFulfillmentNote records, under data.
type ItemFulfillmentNoteList struct {
	Data       []ItemFulfillmentNote
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ItemFulfillmentNoteList was measured from.
func (ItemFulfillmentNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/item_fulfillments/{fulfillment_id}/notes",
	}
}

// ItemFulfillmentNote is one fulfillment note as the notes listing and read answer it; the note is
// stored on the fulfillment's sale invoice.
type ItemFulfillmentNote struct {
	Attachments     []ItemFulfillmentNoteAttachment
	Content         *string
	CreatedBy       *ItemFulfillmentCreatedBy
	CreatedOn       *string
	Group           any
	Href            *string
	ID              *string
	ItemFulfillment *ItemFulfillmentNotesDetail
	LastUpdatedBy   any
	LastUpdatedOn   any
	Object          *string
	User            any
	Version         *int64
	Watchers        []ItemFulfillmentNoteWatcher
}

// ItemFulfillmentNoteAttachment is one entry of a ItemFulfillmentNote's attachments.
type ItemFulfillmentNoteAttachment struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// ItemFulfillmentCreatedBy is a ItemFulfillment's created_by.
type ItemFulfillmentCreatedBy struct {
	Name   *string
	Object *string
}

// ItemFulfillmentNotesDetail is a ItemFulfillmentNotes's item_fulfillment.
type ItemFulfillmentNotesDetail struct {
	Href   *string
	ID     *string
	Object *string
}

// ItemFulfillmentNoteWatcher is one entry of a ItemFulfillmentNote's watchers.
type ItemFulfillmentNoteWatcher struct {
	Email *string
	Name  *string
	Type  *string
}

// ItemFulfillmentNoteReadRecord is what GetItemFulfillmentNote answers: the note under
// item_fulfillment.note.
type ItemFulfillmentNoteReadRecord struct {
	Note *ItemFulfillmentNote
}

// MeasuredFrom lists the endpoints ItemFulfillmentNoteReadRecord was measured from.
func (ItemFulfillmentNoteReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/item_fulfillments/{fulfillment_id}/notes/{note_id}",
	}
}

// EnvelopeKey unwraps {"item_fulfillment": {...}}.
func (ItemFulfillmentNoteReadRecord) EnvelopeKey() string { return ItemFulfillmentKey }

// ItemFulfillmentNoteListShape reads a body as ItemFulfillmentNoteList.
func ItemFulfillmentNoteListShape(decoded any) any { return dto.Read[ItemFulfillmentNoteList](decoded) }

// ItemFulfillmentNoteReadRecordShape reads a body as ItemFulfillmentNoteReadRecord.
func ItemFulfillmentNoteReadRecordShape(decoded any) any {
	return dto.Read[ItemFulfillmentNoteReadRecord](decoded)
}
