package inventoryadjustments

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

// InventoryAdjustmentRecord is the inventory adjustment body.
type InventoryAdjustmentRecord struct {
	AdjustmentAmount         *string
	AdjustmentType           *string
	CreatedBy                any
	CreatedOn                *string
	Href                     *string
	ID                       *string
	ItemID                   *string
	ItemName                 *string
	LastUpdatedBy            any
	LastUpdatedOn            *string
	Object                   *string
	QuantityAvailable        *string
	QuantityOnHand           *string
	QuantityOnOrder          *string
	QuantityOnPurchaseReturn *string
	QuantityOnReserve        *string
	QuantityOnReturn         *string
	QuantityPromised         *string
	Reference                *string
	Status                   *string
	WarehouseName            *string
}

// MeasuredFrom lists the endpoints InventoryAdjustmentRecord was measured from.
func (InventoryAdjustmentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/inventory_adjustments",
		"GET /api/{api_version}/inventory_adjustments/{inventory_adjustment_id}",
		"GET /api/{api_version}/items/{item_id}/inventory_adjustments",
	}
}

// EnvelopeKey unwraps {"inventory_adjustment": {...}}.
func (InventoryAdjustmentRecord) EnvelopeKey() string { return InventoryAdjustmentKey }

// InventoryAdjustmentList is a page of InventoryAdjustmentRecord records.
type InventoryAdjustmentList struct {
	Data       []InventoryAdjustmentRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints InventoryAdjustmentList was measured from.
func (InventoryAdjustmentList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/inventory_adjustments",
		"GET /api/{api_version}/items/{item_id}/inventory_adjustments",
	}
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// InventoryAdjustmentListShape reads a body as InventoryAdjustmentList.
func InventoryAdjustmentListShape(decoded any) any { return dto.Read[InventoryAdjustmentList](decoded) }

// InventoryAdjustmentRecordShape reads a body as InventoryAdjustmentRecord.
func InventoryAdjustmentRecordShape(decoded any) any {
	return dto.Read[InventoryAdjustmentRecord](decoded)
}
