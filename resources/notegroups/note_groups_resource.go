package notegroups

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the NoteGroups method of the same name,
// which is the primary surface.

// GetNoteGroups is the module form of NoteGroups.GetNoteGroups.
func GetNoteGroups(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetNoteGroups(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateNoteGroup is the module form of NoteGroups.CreateNoteGroup.
func CreateNoteGroup(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateNoteGroup(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteNoteGroup is the module form of NoteGroups.DeleteNoteGroup.
func DeleteNoteGroup(requester exhttp.Requester, ctx context.Context, noteGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteNoteGroup(ctx, noteGroupID, raiseOnError, responseFormat...)
}

// GetNoteGroup is the module form of NoteGroups.GetNoteGroup.
func GetNoteGroup(requester exhttp.Requester, ctx context.Context, noteGroupID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetNoteGroup(ctx, noteGroupID, queryParams, raiseOnError, responseFormat...)
}

// UpdateNoteGroup is the module form of NoteGroups.UpdateNoteGroup.
func UpdateNoteGroup(requester exhttp.Requester, ctx context.Context, noteGroupID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateNoteGroup(ctx, noteGroupID, payload, raiseOnError, responseFormat...)
}

// GetNoteGroupNotes is the module form of NoteGroups.GetNoteGroupNotes.
func GetNoteGroupNotes(requester exhttp.Requester, ctx context.Context, noteGroupID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetNoteGroupNotes(ctx, noteGroupID, queryParams, raiseOnError, responseFormat...)
}

// GetNoteGroupNotesByUser is the module form of NoteGroups.GetNoteGroupNotesByUser.
func GetNoteGroupNotesByUser(requester exhttp.Requester, ctx context.Context, noteGroupID any, noteUser any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetNoteGroupNotesByUser(ctx, noteGroupID, noteUser, queryParams, raiseOnError, responseFormat...)
}
