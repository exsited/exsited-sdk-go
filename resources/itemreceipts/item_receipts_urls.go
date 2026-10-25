// Package itemreceipts is the v4 item receipts endpoints.
package itemreceipts

// The path templates for the item receipts endpoints; the client fills {api_version}.
const (
	ItemReceiptsPath              = "/api/{api_version}/item_receipts"
	ItemReceiptPath               = "/api/{api_version}/item_receipts/{item_receipt_id}"
	PurchaseOrderItemReceiptsPath = "/api/{api_version}/purchase_orders/{purchase_order_id}/item_receipts"
	ItemReceiptNotePath           = "/api/{api_version}/item_receipts/{item_receipt_id}/notes/{note_id}"
)

// ItemReceiptsKey is the key a listing's records arrive under, the v4 list envelope's data.
const ItemReceiptsKey = "data"

// ItemReceiptKey is the envelope key one record arrives under.
const ItemReceiptKey = "item_receipt"

// PurchaseOrderKey is the envelope key one record arrives under.
const PurchaseOrderKey = "purchase_order"
