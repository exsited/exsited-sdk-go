package purchasecreditnotes

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// PurchaseCreditNotes is the v4 purchase credit notes endpoints, bound to one client.
type PurchaseCreditNotes struct {
	requester exhttp.Requester
}

// GetPurchaseCreditNotes sends GET /api/{api_version}/purchase_credit_notes.
func (p *PurchaseCreditNotes) GetPurchaseCreditNotes(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseCreditNotesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseCreditNoteListShape})
}

// GetPurchaseCreditNote sends GET /api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}.
func (p *PurchaseCreditNotes) GetPurchaseCreditNote(ctx context.Context, purchaseCreditNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseCreditNoteID, "purchase_credit_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseCreditNotePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_credit_note_id": purchaseCreditNoteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseCreditNoteReadRecordShape})
}

// UpdatePurchaseCreditNote sends PATCH /api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}.
func (p *PurchaseCreditNotes) UpdatePurchaseCreditNote(ctx context.Context, purchaseCreditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseCreditNoteID, "purchase_credit_note_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPatchRequest(ctx, PurchaseCreditNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_credit_note_id": purchaseCreditNoteID},
		JSONBody:   payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseCreditNoteUpdateRecordShape})
}

// CreatePurchaseCreditNote sends POST /api/{api_version}/purchase_credit_notes.
func (p *PurchaseCreditNotes) CreatePurchaseCreditNote(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseCreditNotesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PurchaseCreditNoteCreatedRecordShape})
}

// DeletePurchaseCreditNote sends DELETE /api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}; it has no typed view, so read ApiResponse.Map.
func (p *PurchaseCreditNotes) DeletePurchaseCreditNote(ctx context.Context, purchaseCreditNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseCreditNoteID, "purchase_credit_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PurchaseCreditNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_credit_note_id": purchaseCreditNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ApplyPurchaseCreditNote sends POST /api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}/apply.
func (p *PurchaseCreditNotes) ApplyPurchaseCreditNote(ctx context.Context, purchaseCreditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseCreditNoteID, "purchase_credit_note_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseCreditNoteApplyPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_credit_note_id": purchaseCreditNoteID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PurchaseCreditNoteCreatedRecordShape})
}

// CreatePurchaseCreditNoteNote sends POST /api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}/notes; it has no typed view, so read ApiResponse.Map.
func (p *PurchaseCreditNotes) CreatePurchaseCreditNoteNote(ctx context.Context, purchaseCreditNoteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseCreditNoteID, "purchase_credit_note_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseCreditNoteNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_credit_note_id": purchaseCreditNoteID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeletePurchaseCreditNoteNote sends DELETE /api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (p *PurchaseCreditNotes) DeletePurchaseCreditNoteNote(ctx context.Context, purchaseCreditNoteID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseCreditNoteID, "purchase_credit_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PurchaseCreditNoteNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_credit_note_id": purchaseCreditNoteID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "purchase_credit_notes"

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

// New binds the purchase credit notes endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *PurchaseCreditNotes {
	return &PurchaseCreditNotes{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (p *PurchaseCreditNotes) Requester() exhttp.Requester { return p.requester }
