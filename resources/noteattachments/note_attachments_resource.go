package noteattachments

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the NoteAttachments method of the same name,
// which is the primary surface.

// GetNoteAttachment is the module form of NoteAttachments.GetNoteAttachment.
func GetNoteAttachment(requester exhttp.Requester, ctx context.Context, attachmentUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetNoteAttachment(ctx, attachmentUUID, raiseOnError, responseFormat...)
}
