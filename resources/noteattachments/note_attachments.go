package noteattachments

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// NoteAttachments is the v4 note attachments endpoints, bound to one client.
type NoteAttachments struct {
	requester exhttp.Requester
}

// GetNoteAttachment sends GET /api/{api_version}/note_attachment/{attachment_uuid}.
func (n *NoteAttachments) GetNoteAttachment(ctx context.Context, attachmentUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(attachmentUUID, "attachment_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return n.requester.SendGetRequest(ctx, NoteAttachmentPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"attachment_uuid": attachmentUUID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: NoteAttachmentRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "note_attachments"

// oneResponseFormat reads the optional trailing responseFormat, refusing more than one.
func oneResponseFormat(responseFormat []string) (string, error) {
	switch len(responseFormat) {
	case 0:
		return "", nil
	case 1:
		return responseFormat[0], nil
	default:
		return "", exerrors.NewUsageError(fmt.Sprintf(
			"responseFormat: pass one format or none, got %d; the trailing argument is optional, "+
				"and omitting it defers to the client", len(responseFormat)))
	}
}

// isNilValue reports whether a value is a typed nil inside a non-nil any; a zero-valued struct is not.
func isNilValue(value any) bool {
	reflected := reflect.ValueOf(value)
	switch reflected.Kind() {
	case reflect.Map, reflect.Slice, reflect.Pointer, reflect.Interface, reflect.Func:
		return reflected.IsNil()
	}
	return false
}

// requireValue refuses a path value that is nil, empty or only whitespace, before any I/O.
func requireValue(value any, what string) error {
	text, isText := value.(string)
	if value == nil || isNilValue(value) || isText && strings.TrimSpace(text) == "" {
		return exerrors.UsageErrorf("%s is required: the request cannot be built without it", what)
	}
	return nil
}

// New binds the note attachments endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *NoteAttachments { return &NoteAttachments{requester: requester} }

// Requester returns the requester this resource was bound to.
func (n *NoteAttachments) Requester() exhttp.Requester { return n.requester }
