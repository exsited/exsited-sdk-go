// Package refunds is the v4 refunds endpoints.
package refunds

// The path templates for the refunds endpoints; the client fills {api_version}.
const (
	RefundsPath           = "/api/{api_version}/refunds"
	RefundPath            = "/api/{api_version}/refunds/{refund_id}"
	CreditNoteRefundsPath = "/api/{api_version}/credit_notes/{credit_note_id}/refunds"
	RefundNotesPath       = "/api/{api_version}/refunds/{refund_id}/notes"
	RefundNotePath        = "/api/{api_version}/refunds/{refund_id}/notes/{note_id}"
	RefundNoteFilesPath   = "/api/{api_version}/refunds/{refund_id}/notes/{note_id}/files"
	RefundNoteFilePath    = "/api/{api_version}/refunds/{refund_id}/notes/{note_id}/files/{file_id}"
)

// RefundsKey is the key a listing's records arrive under, the v4 list envelope's data.
const RefundsKey = "data"

// RefundKey is the envelope key one record arrives under.
const RefundKey = "refund"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"

// FileKey is the envelope key one record arrives under.
const FileKey = "file"
