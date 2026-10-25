// Package purchaseorders is the v4 purchase orders endpoints.
package purchaseorders

// The path templates for the purchase orders endpoints; the client fills {api_version}.
const (
	PurchaseOrdersPath                = "/api/{api_version}/purchase_orders"
	PurchaseOrdersNotesPath           = "/api/{api_version}/purchase_orders/notes"
	PurchaseOrderPath                 = "/api/{api_version}/purchase_orders/{purchase_order_id}"
	PurchaseOrderCancelPath           = "/api/{api_version}/purchase_orders/{purchase_order_id}/cancel"
	PurchaseOrderChangePath           = "/api/{api_version}/purchase_orders/{purchase_order_id}/change"
	PurchaseOrderInformationPath      = "/api/{api_version}/purchase_orders/{purchase_order_id}/information"
	PurchaseOrderLinesPath            = "/api/{api_version}/purchase_orders/{purchase_order_id}/lines"
	PurchaseOrderNotesPath            = "/api/{api_version}/purchase_orders/{purchase_order_id}/notes"
	PurchaseOrderPurchaseInvoicesPath = "/api/{api_version}/purchase_orders/{purchase_order_id}/purchase_invoices"
	PurchaseOrderPurchasePaymentsPath = "/api/{api_version}/purchase_orders/{purchase_order_id}/purchase_payments"
	PurchaseOrderReactivatePath       = "/api/{api_version}/purchase_orders/{purchase_order_id}/reactivate"
	PurchaseOrderLinePath             = "/api/{api_version}/purchase_orders/{purchase_order_id}/lines/{purchase_order_line_uuid}"
	PurchaseOrderNotePath             = "/api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}"
	PurchaseOrderNoteFilesPath        = "/api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}/files"
	PurchaseOrderNoteFilePath         = "/api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}/files/{purchase_order_file_id}"
	PurchaseOrderKPIsPath             = "/api/{api_version}/purchase_orders/{purchase_order_id}/kpis"
)

// PurchaseOrdersKey is the key a listing's records arrive under, the v4 list envelope's data.
const PurchaseOrdersKey = "data"

// PurchaseOrderKey is the envelope key one record arrives under.
const PurchaseOrderKey = "purchase_order"

// PurchaseOrderLineKey is the envelope key one record arrives under.
const PurchaseOrderLineKey = "purchase_order_line"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"

// FileKey is the envelope key one record arrives under.
const FileKey = "file"

// PurchaseInvoiceKey is the envelope key one record arrives under.
const PurchaseInvoiceKey = "purchase_invoice"

// PurchaseOrderKPIsKey is the envelope key one record arrives under.
const PurchaseOrderKPIsKey = "purchase_order_kpis"
