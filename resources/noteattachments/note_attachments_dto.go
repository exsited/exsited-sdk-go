package noteattachments

import "github.com/exsited/exsited-sdk-go/v4/dto"

// NoteAttachmentRecord is the note attachment body, under note_attachment; base64 is the file
// itself.
type NoteAttachmentRecord struct {
	Base64 *string `wire:"base64"`
	ID     *string
	Name   *string
	Object *string
	Size   *int64
}

// MeasuredFrom lists the endpoints NoteAttachmentRecord was measured from.
func (NoteAttachmentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/note_attachment/{attachment_uuid}",
	}
}

// EnvelopeKey unwraps {"note_attachment": {...}}.
func (NoteAttachmentRecord) EnvelopeKey() string { return NoteAttachmentKey }

// NoteAttachmentRecordShape reads a body as NoteAttachmentRecord.
func NoteAttachmentRecordShape(decoded any) any { return dto.Read[NoteAttachmentRecord](decoded) }
