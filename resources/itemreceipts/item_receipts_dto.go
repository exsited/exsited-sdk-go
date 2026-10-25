package itemreceipts

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

// ItemReceiptRecord is the item receipt body.
type ItemReceiptRecord struct {
	Code             *string
	CreatedBy        *CreatedBy
	CreatedOn        *string
	Href             *string
	ID               *string
	LastUpdatedBy    *LastUpdatedBy
	LastUpdatedOn    *string
	Lines            []Line
	Object           *string
	PurchaseOrder    *PurchaseOrder
	ReceiptDate      *string
	ReceiptNote      *string
	ReceiptStatus    *string
	ReceiveCompleted *bool
	Status           *string
	TrackingNumber   *string
	Version          *int64
}

// MeasuredFrom lists the endpoints ItemReceiptRecord was measured from.
func (ItemReceiptRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/item_receipts",
		"GET /api/{api_version}/item_receipts/{item_receipt_id}",
	}
}

// EnvelopeKey unwraps {"item_receipt": {...}}.
func (ItemReceiptRecord) EnvelopeKey() string { return ItemReceiptKey }

// ItemReceiptList is a page of ItemReceiptRecord records.
type ItemReceiptList struct {
	Data       []ItemReceiptRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ItemReceiptList was measured from.
func (ItemReceiptList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/item_receipts",
	}
}

// ItemReceiptsRecord is the item receipts body.
type ItemReceiptsRecord struct {
	ActivityLogs    []any
	CreatedBy       *string
	CreatedOn       *string
	Date            *string
	ID              *string
	LastUpdatedBy   *string
	LastUpdatedOn   *string
	Note            *string
	PurchaseOrderID *string
	ReceiptStatus   *string
	Receipts        []any
	TrackingNumber  *string
	UUID            *string
	Version         *string
}

// MeasuredFrom lists the endpoints ItemReceiptsRecord was measured from.
func (ItemReceiptsRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/item_receipts",
	}
}

// EnvelopeKey unwraps {"item-receipt": {...}}.
func (ItemReceiptsRecord) EnvelopeKey() string { return ItemReceiptKey }

// ItemReceiptUpdatedRecord is the item receipt updated body.
type ItemReceiptUpdatedRecord struct {
	Code             *string
	CreatedBy        *CreatedBy
	CreatedOn        *string
	Date             *string
	Href             *string
	ID               *string
	LastUpdatedBy    *LastUpdatedBy
	LastUpdatedOn    *string
	Lines            []Line
	Note             *string
	Object           *string
	PurchaseOrder    *PurchaseOrder
	PurchaseOrderID  *string
	ReceiptDate      *string
	ReceiptNote      any
	ReceiptStatus    *string
	Receipts         []any
	ReceiveCompleted *bool
	Status           *string
	TrackingNumber   *string
	UUID             *string
	Version          *int64
}

// MeasuredFrom lists the endpoints ItemReceiptUpdatedRecord was measured from.
func (ItemReceiptUpdatedRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/item_receipts/{item_receipt_id}",
	}
}

// EnvelopeKey unwraps {"item_receipt": {...}}.
func (ItemReceiptUpdatedRecord) EnvelopeKey() string { return ItemReceiptKey }

// PurchaseOrderItemReceiptsRecord is the purchase order item receipts body.
type PurchaseOrderItemReceiptsRecord struct {
	ItemReceipts []any
	Pagination   map[string]any
}

// MeasuredFrom lists the endpoints PurchaseOrderItemReceiptsRecord was measured from.
func (PurchaseOrderItemReceiptsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/purchase_orders/{purchase_order_id}/item_receipts",
	}
}

// EnvelopeKey unwraps {"purchase_order": {...}}.
func (PurchaseOrderItemReceiptsRecord) EnvelopeKey() string { return PurchaseOrderKey }

// CreatedBy is the created_by object.
type CreatedBy struct {
	Name   *string
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

// Line is one entry of the lines list.
type Line struct {
	Item                  *Item
	PendingReturnQuantity *float64
	PurchaseOrderLine     *PurchaseOrderLine
	ReceiptQuantity       *float64
	ReceivedQuantity      *float64
	ReturnedQuantity      *float64
	Uom                   *Uom
	Warehouse             *Warehouse
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// PurchaseOrder is the purchase_order object.
type PurchaseOrder struct {
	Code   *string
	Href   *string
	ID     *string
	Object *string
}

// PurchaseOrderLine is the purchase_order_line object.
type PurchaseOrderLine struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
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

// ItemReceiptListShape reads a body as ItemReceiptList.
func ItemReceiptListShape(decoded any) any { return dto.Read[ItemReceiptList](decoded) }

// ItemReceiptsRecordShape reads a body as ItemReceiptsRecord.
func ItemReceiptsRecordShape(decoded any) any { return dto.Read[ItemReceiptsRecord](decoded) }

// ItemReceiptUpdatedRecordShape reads a body as ItemReceiptUpdatedRecord.
func ItemReceiptUpdatedRecordShape(decoded any) any {
	return dto.Read[ItemReceiptUpdatedRecord](decoded)
}

// ItemReceiptRecordShape reads a body as ItemReceiptRecord.
func ItemReceiptRecordShape(decoded any) any { return dto.Read[ItemReceiptRecord](decoded) }

// PurchaseOrderItemReceiptsRecordShape reads a body as PurchaseOrderItemReceiptsRecord.
func PurchaseOrderItemReceiptsRecordShape(decoded any) any {
	return dto.Read[PurchaseOrderItemReceiptsRecord](decoded)
}
