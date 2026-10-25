package creditnotes

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// CreditNotes is the v4 credit notes endpoints, bound to one client.
type CreditNotes struct {
	requester exhttp.Requester
}

// GetCreditNotes sends GET /api/{api_version}/credit_notes.
func (c *CreditNotes) GetCreditNotes(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CreditNotesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: CreditNoteListShape})
}

// GetCreditNote sends GET /api/{api_version}/credit_notes/{credit_note_id}.
func (c *CreditNotes) GetCreditNote(ctx context.Context, creditNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CreditNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"credit_note_id": creditNoteID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: CreditNoteReadRecordShape})
}

// UpdateCreditNote sends PATCH /api/{api_version}/credit_notes/{credit_note_id}.
func (c *CreditNotes) UpdateCreditNote(ctx context.Context, creditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPatchRequest(ctx, CreditNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"credit_note_id": creditNoteID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: CreditNoteUpdateRecordShape})
}

// CreateCreditNote sends POST /api/{api_version}/credit_notes; it has no typed view, so read ApiResponse.Map.
func (c *CreditNotes) CreateCreditNote(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, CreditNotesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteCreditNote sends DELETE /api/{api_version}/credit_notes/{credit_note_id}; it has no typed view, so read ApiResponse.Map.
func (c *CreditNotes) DeleteCreditNote(ctx context.Context, creditNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendDeleteRequest(ctx, CreditNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"credit_note_id": creditNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ApplyCreditNote sends POST /api/{api_version}/credit_notes/{credit_note_id}/apply; it has no typed view, so read ApiResponse.Map.
func (c *CreditNotes) ApplyCreditNote(ctx context.Context, creditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, CreditNoteApplyPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"credit_note_id": creditNoteID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetCreditNoteItems sends GET /api/{api_version}/credit_notes/{credit_note_id}/items.
func (c *CreditNotes) GetCreditNoteItems(ctx context.Context, creditNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CreditNoteItemsPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"credit_note_id": creditNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: CreditNoteItemsRecordShape})
}

// GetCreditNoteCreditNoteApplications sends GET /api/{api_version}/credit_notes/{credit_note_id}/credit_note_applications.
func (c *CreditNotes) GetCreditNoteCreditNoteApplications(ctx context.Context, creditNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CreditNoteCreditNoteApplicationsPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"credit_note_id": creditNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: CreditNoteApplicationsRecordShape})
}

// GetCreditNoteNotes sends GET /api/{api_version}/credit_notes/{credit_note_id}/notes.
func (c *CreditNotes) GetCreditNoteNotes(ctx context.Context, creditNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CreditNoteNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"credit_note_id": creditNoteID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: CreditNoteNoteListShape})
}

// GetCreditNoteNote sends GET /api/{api_version}/credit_notes/{credit_note_id}/notes/{note_id}.
func (c *CreditNotes) GetCreditNoteNote(ctx context.Context, creditNoteID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CreditNoteNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"credit_note_id": creditNoteID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: CreditNoteNoteRecordShape})
}

// CreateCreditNoteNote sends POST /api/{api_version}/credit_notes/{credit_note_id}/notes; it has no typed view, so read ApiResponse.Map.
func (c *CreditNotes) CreateCreditNoteNote(ctx context.Context, creditNoteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, CreditNoteNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"credit_note_id": creditNoteID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteCreditNoteNote sends DELETE /api/{api_version}/credit_notes/{credit_note_id}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (c *CreditNotes) DeleteCreditNoteNote(ctx context.Context, creditNoteID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendDeleteRequest(ctx, CreditNoteNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"credit_note_id": creditNoteID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetCreditNoteNoteFiles sends GET /api/{api_version}/credit_notes/{credit_note_id}/notes/{note_id}/files; it has no typed view, so read ApiResponse.Map.
func (c *CreditNotes) GetCreditNoteNoteFiles(ctx context.Context, creditNoteID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CreditNoteNoteFilesPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"credit_note_id": creditNoteID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetCreditNoteNoteFile sends GET /api/{api_version}/credit_notes/{credit_note_id}/notes/{note_id}/files/{file_id}; it has no typed view, so read ApiResponse.Map.
func (c *CreditNotes) GetCreditNoteNoteFile(ctx context.Context, creditNoteID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteID, "credit_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(fileID, "file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CreditNoteNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"credit_note_id": creditNoteID, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "credit_notes"

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

// requirePayload refuses a nil body before any I/O; the client's validator would let one through bodyless.
func requirePayload(payload any) error {
	if payload == nil || isNilValue(payload) {
		return exerrors.NewUsageError("payload is required: this endpoint takes a body, and nil would " +
			"send the request with no body at all")
	}
	return nil
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

// requireNote refuses an empty note before any I/O; the text is the whole body of a note write.
func requireNote(note string) error {
	if note == "" {
		return exerrors.NewUsageError("note is required: the note text is the body of this request")
	}
	return nil
}

// New binds the credit notes endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *CreditNotes { return &CreditNotes{requester: requester} }

// Requester returns the requester this resource was bound to.
func (c *CreditNotes) Requester() exhttp.Requester { return c.requester }
