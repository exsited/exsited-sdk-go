package purchaseinvoices

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the PurchaseInvoices method of the same name,
// which is the primary surface.

// GetPurchaseInvoices is the module form of PurchaseInvoices.GetPurchaseInvoices.
func GetPurchaseInvoices(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoices(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreatePurchaseInvoice is the module form of PurchaseInvoices.CreatePurchaseInvoice.
func CreatePurchaseInvoice(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchaseInvoice(ctx, payload, raiseOnError, responseFormat...)
}

// GetPurchaseInvoicesNotes is the module form of PurchaseInvoices.GetPurchaseInvoicesNotes.
func GetPurchaseInvoicesNotes(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoicesNotes(ctx, queryParams, raiseOnError, responseFormat...)
}

// DeletePurchaseInvoice is the module form of PurchaseInvoices.DeletePurchaseInvoice.
func DeletePurchaseInvoice(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePurchaseInvoice(ctx, purchaseInvoiceID, raiseOnError, responseFormat...)
}

// GetPurchaseInvoice is the module form of PurchaseInvoices.GetPurchaseInvoice.
func GetPurchaseInvoice(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoice(ctx, purchaseInvoiceID, queryParams, raiseOnError, responseFormat...)
}

// AmendPurchaseInvoice is the module form of PurchaseInvoices.AmendPurchaseInvoice.
func AmendPurchaseInvoice(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).AmendPurchaseInvoice(ctx, purchaseInvoiceID, payload, raiseOnError, responseFormat...)
}

// CancelPurchaseInvoice is the module form of PurchaseInvoices.CancelPurchaseInvoice.
func CancelPurchaseInvoice(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CancelPurchaseInvoice(ctx, purchaseInvoiceID, payload, raiseOnError, responseFormat...)
}

// GetPurchaseInvoiceInformation is the module form of PurchaseInvoices.GetPurchaseInvoiceInformation.
func GetPurchaseInvoiceInformation(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoiceInformation(ctx, purchaseInvoiceID, queryParams, raiseOnError, responseFormat...)
}

// UpdatePurchaseInvoiceInformation is the module form of PurchaseInvoices.UpdatePurchaseInvoiceInformation.
func UpdatePurchaseInvoiceInformation(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdatePurchaseInvoiceInformation(ctx, purchaseInvoiceID, payload, raiseOnError, responseFormat...)
}

// GetPurchaseInvoiceLines is the module form of PurchaseInvoices.GetPurchaseInvoiceLines.
func GetPurchaseInvoiceLines(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoiceLines(ctx, purchaseInvoiceID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseInvoiceNotes is the module form of PurchaseInvoices.GetPurchaseInvoiceNotes.
func GetPurchaseInvoiceNotes(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoiceNotes(ctx, purchaseInvoiceID, queryParams, raiseOnError, responseFormat...)
}

// CreatePurchaseInvoiceNote is the module form of PurchaseInvoices.CreatePurchaseInvoiceNote.
func CreatePurchaseInvoiceNote(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchaseInvoiceNote(ctx, purchaseInvoiceID, note, files, raiseOnError, responseFormat...)
}

// GetPurchaseInvoicePDF is the module form of PurchaseInvoices.GetPurchaseInvoicePDF.
func GetPurchaseInvoicePDF(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoicePDF(ctx, purchaseInvoiceID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseInvoicePurchaseCreditNoteApplications is the module form of PurchaseInvoices.GetPurchaseInvoicePurchaseCreditNoteApplications.
func GetPurchaseInvoicePurchaseCreditNoteApplications(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoicePurchaseCreditNoteApplications(ctx, purchaseInvoiceID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseInvoicePurchasePayments is the module form of PurchaseInvoices.GetPurchaseInvoicePurchasePayments.
func GetPurchaseInvoicePurchasePayments(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoicePurchasePayments(ctx, purchaseInvoiceID, queryParams, raiseOnError, responseFormat...)
}

// ReactivatePurchaseInvoice is the module form of PurchaseInvoices.ReactivatePurchaseInvoice.
func ReactivatePurchaseInvoice(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReactivatePurchaseInvoice(ctx, purchaseInvoiceID, payload, raiseOnError, responseFormat...)
}

// GetPurchaseInvoiceLine is the module form of PurchaseInvoices.GetPurchaseInvoiceLine.
func GetPurchaseInvoiceLine(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, purchaseInvoiceLineUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoiceLine(ctx, purchaseInvoiceID, purchaseInvoiceLineUUID, queryParams, raiseOnError, responseFormat...)
}

// UpdatePurchaseInvoiceNote is the module form of PurchaseInvoices.UpdatePurchaseInvoiceNote.
func UpdatePurchaseInvoiceNote(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, noteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdatePurchaseInvoiceNote(ctx, purchaseInvoiceID, noteID, note, files, raiseOnError, responseFormat...)
}

// GetPurchaseInvoiceNote is the module form of PurchaseInvoices.GetPurchaseInvoiceNote.
func GetPurchaseInvoiceNote(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, noteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoiceNote(ctx, purchaseInvoiceID, noteID, queryParams, raiseOnError, responseFormat...)
}

// DeletePurchaseInvoiceNote is the module form of PurchaseInvoices.DeletePurchaseInvoiceNote.
func DeletePurchaseInvoiceNote(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePurchaseInvoiceNote(ctx, purchaseInvoiceID, noteID, raiseOnError, responseFormat...)
}

// GetPurchaseInvoiceNoteFiles is the module form of PurchaseInvoices.GetPurchaseInvoiceNoteFiles.
func GetPurchaseInvoiceNoteFiles(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, noteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoiceNoteFiles(ctx, purchaseInvoiceID, noteID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseInvoiceNoteFile is the module form of PurchaseInvoices.GetPurchaseInvoiceNoteFile.
func GetPurchaseInvoiceNoteFile(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, noteID any, fileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoiceNoteFile(ctx, purchaseInvoiceID, noteID, fileID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseInvoicePurchaseCreditNoteApplication is the module form of PurchaseInvoices.GetPurchaseInvoicePurchaseCreditNoteApplication.
func GetPurchaseInvoicePurchaseCreditNoteApplication(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, applicationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseInvoicePurchaseCreditNoteApplication(ctx, purchaseInvoiceID, applicationID, raiseOnError, responseFormat...)
}
