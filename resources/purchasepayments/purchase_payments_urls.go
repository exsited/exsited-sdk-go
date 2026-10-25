// Package purchasepayments is the v4 purchase payments endpoints.
package purchasepayments

// The path templates for the purchase payments endpoints; the client fills {api_version}.
const (
	PurchasePaymentsPath                   = "/api/{api_version}/purchase_payments"
	PurchasePaymentPath                    = "/api/{api_version}/purchase_payments/{purchase_payment_id}"
	PurchaseInvoicePurchasePaymentsPath    = "/api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_payments"
	PurchasePaymentNotesPath               = "/api/{api_version}/purchase_payments/{purchase_payment_id}/notes"
	PurchasePaymentNotePath                = "/api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}"
	PurchasePaymentNoteFilesPath           = "/api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}/files"
	PurchasePaymentNoteFilePath            = "/api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}/files/{file_id}"
	PurchasePaymentPurchaseCreditNotesPath = "/api/{api_version}/purchase_payments/{purchase_payment_id}/purchase_credit_notes"
)

// PurchasePaymentsKey is the key a listing's records arrive under, the v4 list envelope's data.
const PurchasePaymentsKey = "data"

// FilesKey is the envelope key one record arrives under.
const FilesKey = "files"

// PurchasePaymentKey is the envelope key one record arrives under.
const PurchasePaymentKey = "purchase_payment"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"

// FileKey is the envelope key one record arrives under.
const FileKey = "file"
