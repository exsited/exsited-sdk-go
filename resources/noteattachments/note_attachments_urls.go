// Package noteattachments is the v4 note attachments endpoints.
package noteattachments

// The path templates for the note attachments endpoints; the client fills {api_version}.
const (
	NoteAttachmentPath = "/api/{api_version}/note_attachment/{attachment_uuid}"
)

// NoteAttachmentKey is the envelope key one record arrives under.
const NoteAttachmentKey = "note_attachment"
