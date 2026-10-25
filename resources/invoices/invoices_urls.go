// Package invoices is the v4 invoices endpoints.
package invoices

// The path templates for the invoices endpoints; the client fills {api_version}.
const (
	InvoicesPath                               = "/api/{api_version}/invoices"
	InvoicePath                                = "/api/{api_version}/invoices/{invoice_id}"
	InvoiceAmendPath                           = "/api/{api_version}/invoices/{invoice_id}/amend"
	InvoiceCancelPath                          = "/api/{api_version}/invoices/{invoice_id}/cancel"
	InvoiceInformationPath                     = "/api/{api_version}/invoices/{invoice_id}/information"
	InvoiceItemFulfillmentsPath                = "/api/{api_version}/invoices/{invoice_id}/item_fulfillments"
	InvoiceNotesPath                           = "/api/{api_version}/invoices/{invoice_id}/notes"
	InvoicePDFPath                             = "/api/{api_version}/invoices/{invoice_id}/pdf"
	InvoiceReactivatePath                      = "/api/{api_version}/invoices/{invoice_id}/reactivate"
	InvoiceReturnMerchandiseAuthorisationsPath = "/api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations"
	OrderInvoicesPath                          = "/api/{api_version}/orders/{order_id}/invoices"
	InvoiceItemFulfillmentPath                 = "/api/{api_version}/invoices/{invoice_id}/item_fulfillments/{fulfillment_id}"
	InvoiceNotePath                            = "/api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}"
	InvoiceReturnMerchandiseAuthorisationPath  = "/api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations/{rma_id}"
	InvoiceNoteFilesPath                       = "/api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}/files"
	InvoiceNoteFilePath                        = "/api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}/files/{invoice_file_id}"
	InvoiceKPIsPath                            = "/api/{api_version}/invoices/{invoice_id}/kpis"
	InvoiceLinesPath                           = "/api/{api_version}/invoices/{invoice_id}/lines"
	InvoiceLinePath                            = "/api/{api_version}/invoices/{invoice_id}/lines/{line_id}"
	InvoiceCreditNoteApplicationPath           = "/api/{api_version}/invoices/{invoice_id}/credit_note_applications/{credit_note_application_id}"
	InvoiceCustomObjectsPath                   = "/api/{api_version}/invoices/{invoice_id}/custom_objects"
	InvoiceCustomObjectPath                    = "/api/{api_version}/invoices/{invoice_id}/custom_objects/{custom_object_id}"
	InvoiceCustomObjectInstancePath            = "/api/{api_version}/invoices/{invoice_id}/custom_objects/{custom_object_id}/instances/{instance_id}"
)

// InvoicesKey is the key a listing's records arrive under, the v4 list envelope's data.
const InvoicesKey = "data"

// InvoiceKey is the envelope key one record arrives under.
const InvoiceKey = "invoice"

// ItemFulfillmentKey is the envelope key one record arrives under.
const ItemFulfillmentKey = "item_fulfillment"

// FileKey is the envelope key one record arrives under.
const FileKey = "file"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"

// ReturnMerchandiseAuthorisationKey is the envelope key one record arrives under.
const ReturnMerchandiseAuthorisationKey = "return_merchandise_authorisation"

// InvoiceKPIsKey is the envelope key one record arrives under.
const InvoiceKPIsKey = "invoice_kpis"
