// Package purchasecreditnotes is the v4 purchase credit notes endpoints.
package purchasecreditnotes

// The path templates for the purchase credit notes endpoints; the client fills {api_version}.
const (
	PurchaseCreditNotesPath     = "/api/{api_version}/purchase_credit_notes"
	PurchaseCreditNotePath      = "/api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}"
	PurchaseCreditNoteApplyPath = "/api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}/apply"
	PurchaseCreditNoteNotesPath = "/api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}/notes"
	PurchaseCreditNoteNotePath  = "/api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}/notes/{note_id}"
)

// PurchaseCreditNotesKey is the key a listing's records arrive under, the v4 list envelope's data.
const PurchaseCreditNotesKey = "data"

// PurchaseCreditNoteKey is the envelope key one record arrives under.
const PurchaseCreditNoteKey = "purchase_credit_note"
