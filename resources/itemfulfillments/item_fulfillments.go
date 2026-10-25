package itemfulfillments

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// ItemFulfillments is the v4 item fulfillments endpoints, bound to one client.
type ItemFulfillments struct {
	requester exhttp.Requester
}

// GetItemFulfillments sends GET /api/{api_version}/item_fulfillments.
func (i *ItemFulfillments) GetItemFulfillments(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemFulfillmentsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ItemFulfillmentListShape})
}

// CreateItemFulfillment sends POST /api/{api_version}/item_fulfillments.
func (i *ItemFulfillments) CreateItemFulfillment(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemFulfillmentsPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ItemFulfillmentsRecordShape})
}

// GetItemFulfillment sends GET /api/{api_version}/item_fulfillments/{fulfillment_id}.
func (i *ItemFulfillments) GetItemFulfillment(ctx context.Context, fulfillmentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fulfillmentID, "fulfillment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemFulfillmentPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"fulfillment_id": fulfillmentID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ItemFulfillmentReadRecordShape})
}

// UpdateItemFulfillment sends PATCH /api/{api_version}/item_fulfillments/{fulfillment_id}.
func (i *ItemFulfillments) UpdateItemFulfillment(ctx context.Context, fulfillmentID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fulfillmentID, "fulfillment_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, ItemFulfillmentPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"fulfillment_id": fulfillmentID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ItemFulfillmentUpdatedRecordShape})
}

// CreateItemFulfillmentNote sends POST /api/{api_version}/item_fulfillments/{fulfillment_id}/notes.
// The note travels as the form field note, and any files go beside it as multipart parts.
func (i *ItemFulfillments) CreateItemFulfillmentNote(ctx context.Context, fulfillmentID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fulfillmentID, "fulfillment_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemFulfillmentNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"fulfillment_id": fulfillmentID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ItemFulfillmentNoteCreateRecordShape})
}

// DeleteItemFulfillment sends DELETE /api/{api_version}/item_fulfillments/{fulfillment_id}; it has no typed view, so read ApiResponse.Map.
func (i *ItemFulfillments) DeleteItemFulfillment(ctx context.Context, fulfillmentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fulfillmentID, "fulfillment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendDeleteRequest(ctx, ItemFulfillmentPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"fulfillment_id": fulfillmentID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteItemFulfillmentNote sends DELETE /api/{api_version}/item_fulfillments/{fulfillment_id}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (i *ItemFulfillments) DeleteItemFulfillmentNote(ctx context.Context, fulfillmentID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fulfillmentID, "fulfillment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendDeleteRequest(ctx, ItemFulfillmentNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"fulfillment_id": fulfillmentID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetItemFulfillmentNotes sends GET /api/{api_version}/item_fulfillments/{fulfillment_id}/notes.
func (i *ItemFulfillments) GetItemFulfillmentNotes(ctx context.Context, fulfillmentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fulfillmentID, "fulfillment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemFulfillmentNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"fulfillment_id": fulfillmentID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ItemFulfillmentNoteListShape})
}

// GetItemFulfillmentNote sends GET /api/{api_version}/item_fulfillments/{fulfillment_id}/notes/{note_id}.
func (i *ItemFulfillments) GetItemFulfillmentNote(ctx context.Context, fulfillmentID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fulfillmentID, "fulfillment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemFulfillmentNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"fulfillment_id": fulfillmentID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ItemFulfillmentNoteReadRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "item_fulfillments"

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

// requireNote refuses an empty note before any I/O; the text is the whole body of a note write.
func requireNote(note string) error {
	if note == "" {
		return exerrors.NewUsageError("note is required: the note text is the body of this request")
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

// New binds the item fulfillments endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *ItemFulfillments {
	return &ItemFulfillments{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (i *ItemFulfillments) Requester() exhttp.Requester { return i.requester }
