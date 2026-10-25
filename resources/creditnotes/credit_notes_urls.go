// Package creditnotes is the v4 credit notes endpoints.
package creditnotes

// The path templates for the credit notes endpoints; the client fills {api_version}.
const (
	CreditNotesPath                      = "/api/{api_version}/credit_notes"
	CreditNotePath                       = "/api/{api_version}/credit_notes/{credit_note_id}"
	CreditNoteApplyPath                  = "/api/{api_version}/credit_notes/{credit_note_id}/apply"
	CreditNoteItemsPath                  = "/api/{api_version}/credit_notes/{credit_note_id}/items"
	CreditNoteCreditNoteApplicationsPath = "/api/{api_version}/credit_notes/{credit_note_id}/credit_note_applications"
	CreditNoteNotesPath                  = "/api/{api_version}/credit_notes/{credit_note_id}/notes"
	CreditNoteNotePath                   = "/api/{api_version}/credit_notes/{credit_note_id}/notes/{note_id}"
	CreditNoteNoteFilesPath              = "/api/{api_version}/credit_notes/{credit_note_id}/notes/{note_id}/files"
	CreditNoteNoteFilePath               = "/api/{api_version}/credit_notes/{credit_note_id}/notes/{note_id}/files/{file_id}"
)

// CreditNotesKey is the key a listing's records arrive under, the v4 list envelope's data.
const CreditNotesKey = "data"

// CreditNoteKey is the envelope key one record arrives under.
const CreditNoteKey = "credit_note"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"
