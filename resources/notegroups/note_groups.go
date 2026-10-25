package notegroups

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// NoteGroups is the v4 note groups endpoints, bound to one client.
type NoteGroups struct {
	requester exhttp.Requester
}

// GetNoteGroups sends GET /api/{api_version}/note_groups.
func (n *NoteGroups) GetNoteGroups(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return n.requester.SendGetRequest(ctx, NoteGroupsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: NoteGroupListShape})
}

// CreateNoteGroup sends POST /api/{api_version}/note_groups.
func (n *NoteGroups) CreateNoteGroup(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return n.requester.SendPostRequest(ctx, NoteGroupsPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: NoteGroupRecordShape})
}

// DeleteNoteGroup sends DELETE /api/{api_version}/note_groups/{note_group_id}; it has no typed view, so read ApiResponse.Map.
func (n *NoteGroups) DeleteNoteGroup(ctx context.Context, noteGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(noteGroupID, "note_group_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return n.requester.SendDeleteRequest(ctx, NoteGroupPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"note_group_id": noteGroupID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetNoteGroup sends GET /api/{api_version}/note_groups/{note_group_id}.
func (n *NoteGroups) GetNoteGroup(ctx context.Context, noteGroupID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(noteGroupID, "note_group_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return n.requester.SendGetRequest(ctx, NoteGroupPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"note_group_id": noteGroupID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: NoteGroupRecordShape})
}

// UpdateNoteGroup sends PATCH /api/{api_version}/note_groups/{note_group_id}.
func (n *NoteGroups) UpdateNoteGroup(ctx context.Context, noteGroupID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(noteGroupID, "note_group_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return n.requester.SendPatchRequest(ctx, NoteGroupPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"note_group_id": noteGroupID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: NoteGroupRecordShape})
}

// GetNoteGroupNotes sends GET /api/{api_version}/note_groups/{note_group_id}/notes.
func (n *NoteGroups) GetNoteGroupNotes(ctx context.Context, noteGroupID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(noteGroupID, "note_group_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return n.requester.SendGetRequest(ctx, NoteGroupNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"note_group_id": noteGroupID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: NoteGroupNotesRecordShape})
}

// GetNoteGroupNotesByUser sends GET /api/{api_version}/note_groups/{note_group_id}/notes/user/{note_user}; it has no typed view, so read ApiResponse.Map.
func (n *NoteGroups) GetNoteGroupNotesByUser(ctx context.Context, noteGroupID any, noteUser any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(noteGroupID, "note_group_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteUser, "note_user"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return n.requester.SendGetRequest(ctx, NoteGroupNotesByUserPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"note_group_id": noteGroupID, "note_user": noteUser}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "note_groups"

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

// New binds the note groups endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *NoteGroups { return &NoteGroups{requester: requester} }

// Requester returns the requester this resource was bound to.
func (n *NoteGroups) Requester() exhttp.Requester { return n.requester }
