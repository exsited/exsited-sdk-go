package creditnotes

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the CreditNotes method of the same name,
// which is the primary surface.

// GetCreditNotes is the module form of CreditNotes.GetCreditNotes.
func GetCreditNotes(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCreditNotes(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetCreditNote is the module form of CreditNotes.GetCreditNote.
func GetCreditNote(requester exhttp.Requester, ctx context.Context, creditNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCreditNote(ctx, creditNoteID, queryParams, raiseOnError, responseFormat...)
}

// UpdateCreditNote is the module form of CreditNotes.UpdateCreditNote.
func UpdateCreditNote(requester exhttp.Requester, ctx context.Context, creditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateCreditNote(ctx, creditNoteID, payload, raiseOnError, responseFormat...)
}

// CreateCreditNote is the module form of CreditNotes.CreateCreditNote.
func CreateCreditNote(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateCreditNote(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteCreditNote is the module form of CreditNotes.DeleteCreditNote.
func DeleteCreditNote(requester exhttp.Requester, ctx context.Context, creditNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteCreditNote(ctx, creditNoteID, raiseOnError, responseFormat...)
}

// ApplyCreditNote is the module form of CreditNotes.ApplyCreditNote.
func ApplyCreditNote(requester exhttp.Requester, ctx context.Context, creditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ApplyCreditNote(ctx, creditNoteID, payload, raiseOnError, responseFormat...)
}

// GetCreditNoteItems is the module form of CreditNotes.GetCreditNoteItems.
func GetCreditNoteItems(requester exhttp.Requester, ctx context.Context, creditNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCreditNoteItems(ctx, creditNoteID, raiseOnError, responseFormat...)
}

// GetCreditNoteCreditNoteApplications is the module form of CreditNotes.GetCreditNoteCreditNoteApplications.
func GetCreditNoteCreditNoteApplications(requester exhttp.Requester, ctx context.Context, creditNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCreditNoteCreditNoteApplications(ctx, creditNoteID, raiseOnError, responseFormat...)
}

// GetCreditNoteNotes is the module form of CreditNotes.GetCreditNoteNotes.
func GetCreditNoteNotes(requester exhttp.Requester, ctx context.Context, creditNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCreditNoteNotes(ctx, creditNoteID, queryParams, raiseOnError, responseFormat...)
}

// GetCreditNoteNote is the module form of CreditNotes.GetCreditNoteNote.
func GetCreditNoteNote(requester exhttp.Requester, ctx context.Context, creditNoteID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCreditNoteNote(ctx, creditNoteID, noteID, raiseOnError, responseFormat...)
}

// CreateCreditNoteNote is the module form of CreditNotes.CreateCreditNoteNote.
func CreateCreditNoteNote(requester exhttp.Requester, ctx context.Context, creditNoteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateCreditNoteNote(ctx, creditNoteID, note, files, raiseOnError, responseFormat...)
}

// DeleteCreditNoteNote is the module form of CreditNotes.DeleteCreditNoteNote.
func DeleteCreditNoteNote(requester exhttp.Requester, ctx context.Context, creditNoteID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteCreditNoteNote(ctx, creditNoteID, noteID, raiseOnError, responseFormat...)
}

// GetCreditNoteNoteFiles is the module form of CreditNotes.GetCreditNoteNoteFiles.
func GetCreditNoteNoteFiles(requester exhttp.Requester, ctx context.Context, creditNoteID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCreditNoteNoteFiles(ctx, creditNoteID, noteID, raiseOnError, responseFormat...)
}

// GetCreditNoteNoteFile is the module form of CreditNotes.GetCreditNoteNoteFile.
func GetCreditNoteNoteFile(requester exhttp.Requester, ctx context.Context, creditNoteID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCreditNoteNoteFile(ctx, creditNoteID, noteID, fileID, raiseOnError, responseFormat...)
}
