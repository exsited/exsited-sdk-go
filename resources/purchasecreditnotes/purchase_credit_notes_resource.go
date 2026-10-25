package purchasecreditnotes

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the PurchaseCreditNotes method of the same name,
// which is the primary surface.

// GetPurchaseCreditNotes is the module form of PurchaseCreditNotes.GetPurchaseCreditNotes.
func GetPurchaseCreditNotes(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseCreditNotes(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseCreditNote is the module form of PurchaseCreditNotes.GetPurchaseCreditNote.
func GetPurchaseCreditNote(requester exhttp.Requester, ctx context.Context, purchaseCreditNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseCreditNote(ctx, purchaseCreditNoteID, queryParams, raiseOnError, responseFormat...)
}

// UpdatePurchaseCreditNote is the module form of PurchaseCreditNotes.UpdatePurchaseCreditNote.
func UpdatePurchaseCreditNote(requester exhttp.Requester, ctx context.Context, purchaseCreditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdatePurchaseCreditNote(ctx, purchaseCreditNoteID, payload, raiseOnError, responseFormat...)
}

// CreatePurchaseCreditNote is the module form of PurchaseCreditNotes.CreatePurchaseCreditNote.
func CreatePurchaseCreditNote(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchaseCreditNote(ctx, payload, raiseOnError, responseFormat...)
}

// DeletePurchaseCreditNote is the module form of PurchaseCreditNotes.DeletePurchaseCreditNote.
func DeletePurchaseCreditNote(requester exhttp.Requester, ctx context.Context, purchaseCreditNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePurchaseCreditNote(ctx, purchaseCreditNoteID, raiseOnError, responseFormat...)
}

// ApplyPurchaseCreditNote is the module form of PurchaseCreditNotes.ApplyPurchaseCreditNote.
func ApplyPurchaseCreditNote(requester exhttp.Requester, ctx context.Context, purchaseCreditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ApplyPurchaseCreditNote(ctx, purchaseCreditNoteID, payload, raiseOnError, responseFormat...)
}

// CreatePurchaseCreditNoteNote is the module form of PurchaseCreditNotes.CreatePurchaseCreditNoteNote.
func CreatePurchaseCreditNoteNote(requester exhttp.Requester, ctx context.Context, purchaseCreditNoteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchaseCreditNoteNote(ctx, purchaseCreditNoteID, note, files, raiseOnError, responseFormat...)
}

// DeletePurchaseCreditNoteNote is the module form of PurchaseCreditNotes.DeletePurchaseCreditNoteNote.
func DeletePurchaseCreditNoteNote(requester exhttp.Requester, ctx context.Context, purchaseCreditNoteID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePurchaseCreditNoteNote(ctx, purchaseCreditNoteID, noteID, raiseOnError, responseFormat...)
}
