package purchasepayments

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the PurchasePayments method of the same name,
// which is the primary surface.

// GetPurchasePayments is the module form of PurchasePayments.GetPurchasePayments.
func GetPurchasePayments(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchasePayments(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreatePurchasePayment is the module form of PurchasePayments.CreatePurchasePayment.
func CreatePurchasePayment(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchasePayment(ctx, payload, raiseOnError, responseFormat...)
}

// DeletePurchasePayment is the module form of PurchasePayments.DeletePurchasePayment.
func DeletePurchasePayment(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePurchasePayment(ctx, purchasePaymentID, raiseOnError, responseFormat...)
}

// GetPurchasePayment is the module form of PurchasePayments.GetPurchasePayment.
func GetPurchasePayment(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchasePayment(ctx, purchasePaymentID, queryParams, raiseOnError, responseFormat...)
}

// UpdatePurchasePayment is the module form of PurchasePayments.UpdatePurchasePayment.
func UpdatePurchasePayment(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdatePurchasePayment(ctx, purchasePaymentID, payload, raiseOnError, responseFormat...)
}

// CreatePurchaseInvoicePurchasePayment is the module form of PurchasePayments.CreatePurchaseInvoicePurchasePayment.
func CreatePurchaseInvoicePurchasePayment(requester exhttp.Requester, ctx context.Context, purchaseInvoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchaseInvoicePurchasePayment(ctx, purchaseInvoiceID, payload, raiseOnError, responseFormat...)
}

// GetPurchasePaymentNotes is the module form of PurchasePayments.GetPurchasePaymentNotes.
func GetPurchasePaymentNotes(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchasePaymentNotes(ctx, purchasePaymentID, queryParams, raiseOnError, responseFormat...)
}

// CreatePurchasePaymentNote is the module form of PurchasePayments.CreatePurchasePaymentNote.
func CreatePurchasePaymentNote(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchasePaymentNote(ctx, purchasePaymentID, note, files, raiseOnError, responseFormat...)
}

// GetPurchasePaymentNote is the module form of PurchasePayments.GetPurchasePaymentNote.
func GetPurchasePaymentNote(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, ppNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchasePaymentNote(ctx, purchasePaymentID, ppNoteID, queryParams, raiseOnError, responseFormat...)
}

// DeletePurchasePaymentNote is the module form of PurchasePayments.DeletePurchasePaymentNote.
func DeletePurchasePaymentNote(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, ppNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePurchasePaymentNote(ctx, purchasePaymentID, ppNoteID, raiseOnError, responseFormat...)
}

// GetPurchasePaymentNoteFiles is the module form of PurchasePayments.GetPurchasePaymentNoteFiles.
func GetPurchasePaymentNoteFiles(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, ppNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchasePaymentNoteFiles(ctx, purchasePaymentID, ppNoteID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchasePaymentNoteFile is the module form of PurchasePayments.GetPurchasePaymentNoteFile.
func GetPurchasePaymentNoteFile(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, ppNoteID any, fileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchasePaymentNoteFile(ctx, purchasePaymentID, ppNoteID, fileID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchasePaymentPurchaseCreditNotes is the module form of PurchasePayments.GetPurchasePaymentPurchaseCreditNotes.
func GetPurchasePaymentPurchaseCreditNotes(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchasePaymentPurchaseCreditNotes(ctx, purchasePaymentID, queryParams, raiseOnError, responseFormat...)
}

// UpdatePurchasePaymentNote is the module form of PurchasePayments.UpdatePurchasePaymentNote.
func UpdatePurchasePaymentNote(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, ppNoteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdatePurchasePaymentNote(ctx, purchasePaymentID, ppNoteID, note, files, raiseOnError, responseFormat...)
}

// DeletePurchasePaymentNoteFile is the module form of PurchasePayments.DeletePurchasePaymentNoteFile.
func DeletePurchasePaymentNoteFile(requester exhttp.Requester, ctx context.Context, purchasePaymentID any, ppNoteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePurchasePaymentNoteFile(ctx, purchasePaymentID, ppNoteID, fileID, raiseOnError, responseFormat...)
}
