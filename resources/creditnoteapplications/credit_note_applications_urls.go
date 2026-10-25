// Package creditnoteapplications is the v4 credit note applications endpoints.
package creditnoteapplications

// The path templates for the credit note applications endpoints; the client fills {api_version}.
const (
	CreditNoteApplicationsPath        = "/api/{api_version}/credit_note_applications"
	CreditNoteApplicationPath         = "/api/{api_version}/credit_note_applications/{credit_note_application_id}"
	InvoiceCreditNoteApplicationsPath = "/api/{api_version}/invoices/{invoice_id}/credit_note_applications"
)

// CreditNoteApplicationsKey is the key a listing's records arrive under, the v4 list envelope's data.
const CreditNoteApplicationsKey = "data"

// CreditNoteApplicationKey is the envelope key one record arrives under.
const CreditNoteApplicationKey = "credit_note_application"
