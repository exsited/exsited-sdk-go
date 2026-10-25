// Package purchaseinvoices is the v4 purchase invoices endpoints.
package purchaseinvoices

// The path templates for the purchase invoices endpoints; the client fills {api_version}.
const (
	PurchaseInvoicesPath                              = "/api/{api_version}/purchase_invoices"
	PurchaseInvoicesNotesPath                         = "/api/{api_version}/purchase_invoices/notes"
	PurchaseInvoicePath                               = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}"
	PurchaseInvoiceAmendPath                          = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/amend"
	PurchaseInvoiceCancelPath                         = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/cancel"
	PurchaseInvoiceInformationPath                    = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/information"
	PurchaseInvoiceLinesPath                          = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/lines"
	PurchaseInvoiceNotesPath                          = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes"
	PurchaseInvoicePDFPath                            = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/pdf"
	PurchaseInvoicePurchaseCreditNoteApplicationsPath = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_credit_note_applications"
	PurchaseInvoicePurchasePaymentsPath               = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_payments"
	PurchaseInvoiceReactivatePath                     = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/reactivate"
	PurchaseInvoiceLinePath                           = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/lines/{purchase_invoice_line_uuid}"
	PurchaseInvoiceNotePath                           = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}"
	PurchaseInvoiceNoteFilesPath                      = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}/files"
	PurchaseInvoiceNoteFilePath                       = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}/files/{file_id}"
	PurchaseInvoicePurchaseCreditNoteApplicationPath  = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_credit_note_applications/{application_id}"
)

// PurchaseInvoicesKey is the key a listing's records arrive under, the v4 list envelope's data.
const PurchaseInvoicesKey = "data"

// PurchaseInvoiceKey is the envelope key one record arrives under.
const PurchaseInvoiceKey = "purchase_invoice"

// WatcherListKey is the envelope key one record arrives under.
const WatcherListKey = "watcherList"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"

// FileKey is the envelope key one record arrives under.
const FileKey = "file"
