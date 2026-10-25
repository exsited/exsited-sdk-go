// Package purchasecreditnoteapplications is the v4 purchase credit note applications endpoints.
package purchasecreditnoteapplications

// The path templates for the purchase credit note applications endpoints; the client fills {api_version}.
const (
	PurchaseCreditNoteApplicationsPath = "/api/{api_version}/purchase_credit_note_applications"
	PurchaseCreditNoteApplicationPath  = "/api/{api_version}/purchase_credit_note_applications/{purchase_credit_note_application_id}"
)

// PurchaseCreditNoteApplicationsKey is the key a listing's records arrive under, the v4 list envelope's data.
const PurchaseCreditNoteApplicationsKey = "data"
