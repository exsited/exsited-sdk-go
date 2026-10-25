package refunds

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Refunds is the v4 refunds endpoints, bound to one client.
type Refunds struct {
	requester exhttp.Requester
}

// GetRefunds sends GET /api/{api_version}/refunds.
func (r *Refunds) GetRefunds(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, RefundsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: RefundListShape})
}

// DeleteRefund sends DELETE /api/{api_version}/refunds/{refund_id}; it has no typed view, so read ApiResponse.Map.
func (r *Refunds) DeleteRefund(ctx context.Context, refundID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(refundID, "refund_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendDeleteRequest(ctx, RefundPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"refund_id": refundID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetRefund sends GET /api/{api_version}/refunds/{refund_id}.
func (r *Refunds) GetRefund(ctx context.Context, refundID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(refundID, "refund_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, RefundPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"refund_id": refundID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: RefundRecordShape})
}

// UpdateRefund sends PATCH /api/{api_version}/refunds/{refund_id}.
func (r *Refunds) UpdateRefund(ctx context.Context, refundID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(refundID, "refund_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPatchRequest(ctx, RefundPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"refund_id": refundID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: RefundUpdatedRecordShape})
}

// CreateCreditNoteRefund sends POST /api/{api_version}/credit_notes/{credit_note_id}/refunds.
func (r *Refunds) CreateCreditNoteRefund(ctx context.Context, creditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
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
	return r.requester.SendPostRequest(ctx, CreditNoteRefundsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"credit_note_id": creditNoteID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: RefundUpdatedRecordShape})
}

// GetRefundNotes sends GET /api/{api_version}/refunds/{refund_id}/notes.
func (r *Refunds) GetRefundNotes(ctx context.Context, refundID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(refundID, "refund_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, RefundNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"refund_id": refundID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: RefundNoteListShape})
}

// GetRefundNote sends GET /api/{api_version}/refunds/{refund_id}/notes/{note_id}.
func (r *Refunds) GetRefundNote(ctx context.Context, refundID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(refundID, "refund_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, RefundNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"refund_id": refundID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: RefundNoteRecordShape})
}

// CreateRefundNote sends POST /api/{api_version}/refunds/{refund_id}/notes.
func (r *Refunds) CreateRefundNote(ctx context.Context, refundID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(refundID, "refund_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, RefundNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"refund_id": refundID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: RefundNoteRecordShape})
}

// DeleteRefundNote sends DELETE /api/{api_version}/refunds/{refund_id}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (r *Refunds) DeleteRefundNote(ctx context.Context, refundID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(refundID, "refund_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendDeleteRequest(ctx, RefundNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"refund_id": refundID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetRefundNoteFiles sends GET /api/{api_version}/refunds/{refund_id}/notes/{note_id}/files.
func (r *Refunds) GetRefundNoteFiles(ctx context.Context, refundID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(refundID, "refund_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, RefundNoteFilesPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"refund_id": refundID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: NoteFileListShape})
}

// GetRefundNoteFile sends GET /api/{api_version}/refunds/{refund_id}/notes/{note_id}/files/{file_id}.
func (r *Refunds) GetRefundNoteFile(ctx context.Context, refundID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(refundID, "refund_id"); err != nil {
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
	return r.requester.SendGetRequest(ctx, RefundNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"refund_id": refundID, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: RefundNoteFileRecordShape})
}

// DeleteRefundNoteFile sends DELETE /api/{api_version}/refunds/{refund_id}/notes/{note_id}/files/{file_id}; it has no typed view, so read ApiResponse.Map.
func (r *Refunds) DeleteRefundNoteFile(ctx context.Context, refundID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(refundID, "refund_id"); err != nil {
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
	return r.requester.SendDeleteRequest(ctx, RefundNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"refund_id": refundID, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "refunds"

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

// New binds the refunds endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Refunds { return &Refunds{requester: requester} }

// Requester returns the requester this resource was bound to.
func (r *Refunds) Requester() exhttp.Requester { return r.requester }
