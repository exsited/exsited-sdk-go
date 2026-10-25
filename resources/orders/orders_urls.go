// Package orders is the v4 orders endpoints.
package orders

// The path templates for the orders endpoints; the client fills {api_version}.
const (
	OrdersPath                             = "/api/{api_version}/orders"
	OrderPath                              = "/api/{api_version}/orders/{order_id}"
	OrderActivatePath                      = "/api/{api_version}/orders/{order_id}/activate"
	OrderActivateDraftPath                 = "/api/{api_version}/orders/{order_id}/activate_draft"
	OrderBillingPreferencesPath            = "/api/{api_version}/orders/{order_id}/billing_preferences"
	OrderCancelPath                        = "/api/{api_version}/orders/{order_id}/cancel"
	OrderChangePath                        = "/api/{api_version}/orders/{order_id}/change"
	OrderConvertToOrderPath                = "/api/{api_version}/orders/{order_id}/convert_to_order"
	OrderCustomObjectsPath                 = "/api/{api_version}/orders/{order_id}/custom_objects"
	OrderCustomObjectsInstancesPath        = "/api/{api_version}/orders/{order_id}/custom_objects_instances"
	OrderDiscardDraftPath                  = "/api/{api_version}/orders/{order_id}/discard_draft"
	OrderExtendTrialPath                   = "/api/{api_version}/orders/{order_id}/extend_trial"
	OrderInformationPath                   = "/api/{api_version}/orders/{order_id}/information"
	OrderItemFulfillmentsPath              = "/api/{api_version}/orders/{order_id}/item_fulfillments"
	OrderLinesPath                         = "/api/{api_version}/orders/{order_id}/lines"
	OrderNotesPath                         = "/api/{api_version}/orders/{order_id}/notes"
	OrderPaymentMethodsPath                = "/api/{api_version}/orders/{order_id}/payment_methods"
	OrderPDFPath                           = "/api/{api_version}/orders/{order_id}/pdf"
	OrderReactivatePath                    = "/api/{api_version}/orders/{order_id}/reactivate"
	OrderRelinquishPath                    = "/api/{api_version}/orders/{order_id}/relinquish"
	OrderLinePath                          = "/api/{api_version}/orders/{order_id}/lines/{line_uuid}"
	OrderNotePath                          = "/api/{api_version}/orders/{order_id}/notes/{order_note_id}"
	OrderPaymentMethodPath                 = "/api/{api_version}/orders/{order_id}/payment_methods/{payment_method_ref}"
	OrderLineBillingPreferencesPath        = "/api/{api_version}/orders/{order_id}/lines/{line_uuid}/billing_preferences"
	OrderNoteFilesPath                     = "/api/{api_version}/orders/{order_id}/notes/{order_note_id}/files"
	OrderNoteFilePath                      = "/api/{api_version}/orders/{order_id}/notes/{order_note_id}/files/{order_file_id}"
	OrdersNotesPath                        = "/api/{api_version}/orders/notes"
	OrdersPreviewPath                      = "/api/{api_version}/orders/preview"
	OrdersExpressPath                      = "/api/{api_version}/orders/express"
	OrderAcceptPath                        = "/api/{api_version}/orders/{order_id}/accept"
	OrderActivateInvoiceFromDormantPath    = "/api/{api_version}/orders/{order_id}/activate_invoice_from_dormant/{invoice_uuid}"
	OrderChangePreviewPath                 = "/api/{api_version}/orders/{order_id}/change/preview"
	OrderChangeStatusPath                  = "/api/{api_version}/orders/{order_id}/change_status"
	OrderContractAdjustmentPath            = "/api/{api_version}/orders/{order_id}/contract_adjustment"
	OrderContractAdjustmentPreviewPath     = "/api/{api_version}/orders/{order_id}/contract_adjustment/preview"
	OrderDowngradePath                     = "/api/{api_version}/orders/{order_id}/downgrade"
	OrderDowngradePreviewPath              = "/api/{api_version}/orders/{order_id}/downgrade/preview"
	OrderUpgradePath                       = "/api/{api_version}/orders/{order_id}/upgrade"
	OrderUpgradePreviewPath                = "/api/{api_version}/orders/{order_id}/upgrade/preview"
	OrderTerminateTrialPath                = "/api/{api_version}/orders/{order_id}/terminate_trial"
	OrderCustomObjectPath                  = "/api/{api_version}/orders/{order_id}/custom_objects/{custom_object_id}"
	OrderCustomObjectInstancePath          = "/api/{api_version}/orders/{order_id}/custom_objects/{custom_object_id}/instances/{instance_id}"
	OrderCustomObjectsInstancePath         = "/api/{api_version}/orders/{order_id}/custom_objects_instances/{custom_object_id}"
	OrderCustomObjectsInstanceInstancePath = "/api/{api_version}/orders/{order_id}/custom_objects_instances/{custom_object_id}/instances/{instance_id}"
	OrderItemFulfillmentPath               = "/api/{api_version}/orders/{order_id}/item_fulfillments/{fulfillment_id}"
	OrderKPIsPath                          = "/api/{api_version}/orders/{order_id}/kpis"
	OrderPaymentPath                       = "/api/{api_version}/orders/{order_id}/payments/{payment_id}"
	OrderQuotePartialInvoicesPath          = "/api/{api_version}/orders/{order_id}/quote_partial_invoices"
	OrderShippingPDFPath                   = "/api/{api_version}/orders/{order_id}/shipping_pdf/{shipping_line_id}"
)

// OrdersKey is the key a listing's records arrive under, the v4 list envelope's data.
const OrdersKey = "data"

// OrderKey is the envelope key one record arrives under.
const OrderKey = "order"

// OrderLineKey is the envelope key one record arrives under.
const OrderLineKey = "order_line"

// FileKey is the envelope key one record arrives under.
const FileKey = "file"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"

// CustomObjectKey is the envelope key one record arrives under.
const CustomObjectKey = "custom_object"

// ItemFulfillmentKey is the envelope key one record arrives under.
const ItemFulfillmentKey = "item_fulfillment"

// OrderKPIsKey is the envelope key one record arrives under.
const OrderKPIsKey = "order_kpis"

// PaymentKey is the envelope key one record arrives under.
const PaymentKey = "payment"
