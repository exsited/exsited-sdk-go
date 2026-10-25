package payments

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Payments method of the same name,
// which is the primary surface.

// GetPayments is the module form of Payments.GetPayments.
func GetPayments(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPayments(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetPaymentsNotes is the module form of Payments.GetPaymentsNotes.
func GetPaymentsNotes(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPaymentsNotes(ctx, queryParams, raiseOnError, responseFormat...)
}

// DeletePayment is the module form of Payments.DeletePayment.
func DeletePayment(requester exhttp.Requester, ctx context.Context, paymentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePayment(ctx, paymentID, raiseOnError, responseFormat...)
}

// UpdatePayment is the module form of Payments.UpdatePayment.
func UpdatePayment(requester exhttp.Requester, ctx context.Context, paymentID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdatePayment(ctx, paymentID, payload, raiseOnError, responseFormat...)
}

// GetPayment is the module form of Payments.GetPayment.
func GetPayment(requester exhttp.Requester, ctx context.Context, paymentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPayment(ctx, paymentID, queryParams, raiseOnError, responseFormat...)
}

// GetInvoicePayments is the module form of Payments.GetInvoicePayments.
func GetInvoicePayments(requester exhttp.Requester, ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoicePayments(ctx, invoiceID, queryParams, raiseOnError, responseFormat...)
}

// GetOrderPayments is the module form of Payments.GetOrderPayments.
func GetOrderPayments(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderPayments(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// GetPaymentNotes is the module form of Payments.GetPaymentNotes.
func GetPaymentNotes(requester exhttp.Requester, ctx context.Context, paymentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPaymentNotes(ctx, paymentID, queryParams, raiseOnError, responseFormat...)
}

// CreatePaymentNote is the module form of Payments.CreatePaymentNote.
func CreatePaymentNote(requester exhttp.Requester, ctx context.Context, paymentID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePaymentNote(ctx, paymentID, note, files, raiseOnError, responseFormat...)
}

// GetPaymentNote is the module form of Payments.GetPaymentNote.
func GetPaymentNote(requester exhttp.Requester, ctx context.Context, paymentID any, paymentNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPaymentNote(ctx, paymentID, paymentNoteID, queryParams, raiseOnError, responseFormat...)
}

// DeletePaymentNote is the module form of Payments.DeletePaymentNote.
func DeletePaymentNote(requester exhttp.Requester, ctx context.Context, paymentID any, paymentNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePaymentNote(ctx, paymentID, paymentNoteID, raiseOnError, responseFormat...)
}

// GetPaymentNoteFile is the module form of Payments.GetPaymentNoteFile.
func GetPaymentNoteFile(requester exhttp.Requester, ctx context.Context, paymentID any, paymentNoteID any, paymentFileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPaymentNoteFile(ctx, paymentID, paymentNoteID, paymentFileID, queryParams, raiseOnError, responseFormat...)
}

// CreatePayment is the module form of Payments.CreatePayment.
func CreatePayment(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePayment(ctx, payload, raiseOnError, responseFormat...)
}

// GetInvoicePayment is the module form of Payments.GetInvoicePayment.
func GetInvoicePayment(requester exhttp.Requester, ctx context.Context, invoiceID any, paymentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoicePayment(ctx, invoiceID, paymentID, raiseOnError, responseFormat...)
}

// GetPaymentNoteFiles is the module form of Payments.GetPaymentNoteFiles.
func GetPaymentNoteFiles(requester exhttp.Requester, ctx context.Context, paymentID any, paymentNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPaymentNoteFiles(ctx, paymentID, paymentNoteID, raiseOnError, responseFormat...)
}
