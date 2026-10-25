// Package payments is the v4 payments endpoints.
package payments

// The path templates for the payments endpoints; the client fills {api_version}.
const (
	PaymentsPath         = "/api/{api_version}/payments"
	PaymentsNotesPath    = "/api/{api_version}/payments/notes"
	PaymentPath          = "/api/{api_version}/payments/{payment_id}"
	InvoicePaymentsPath  = "/api/{api_version}/invoices/{invoice_id}/payments"
	OrderPaymentsPath    = "/api/{api_version}/orders/{order_id}/payments"
	PaymentNotesPath     = "/api/{api_version}/payments/{payment_id}/notes"
	PaymentNotePath      = "/api/{api_version}/payments/{payment_id}/notes/{payment_note_id}"
	PaymentNoteFilePath  = "/api/{api_version}/payments/{payment_id}/notes/{payment_note_id}/files/{payment_file_id}"
	InvoicePaymentPath   = "/api/{api_version}/invoices/{invoice_id}/payments/{payment_id}"
	PaymentNoteFilesPath = "/api/{api_version}/payments/{payment_id}/notes/{payment_note_id}/files"
)

// PaymentsKey is the key a listing's records arrive under, the v4 list envelope's data.
const PaymentsKey = "data"

// PaymentKey is the envelope key one record arrives under.
const PaymentKey = "payment"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"

// FileKey is the envelope key one record arrives under.
const FileKey = "file"
