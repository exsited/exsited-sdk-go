package refunds

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Refunds method of the same name,
// which is the primary surface.

// GetRefunds is the module form of Refunds.GetRefunds.
func GetRefunds(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetRefunds(ctx, queryParams, raiseOnError, responseFormat...)
}

// DeleteRefund is the module form of Refunds.DeleteRefund.
func DeleteRefund(requester exhttp.Requester, ctx context.Context, refundID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteRefund(ctx, refundID, raiseOnError, responseFormat...)
}

// GetRefund is the module form of Refunds.GetRefund.
func GetRefund(requester exhttp.Requester, ctx context.Context, refundID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetRefund(ctx, refundID, queryParams, raiseOnError, responseFormat...)
}

// UpdateRefund is the module form of Refunds.UpdateRefund.
func UpdateRefund(requester exhttp.Requester, ctx context.Context, refundID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateRefund(ctx, refundID, payload, raiseOnError, responseFormat...)
}

// CreateCreditNoteRefund is the module form of Refunds.CreateCreditNoteRefund.
func CreateCreditNoteRefund(requester exhttp.Requester, ctx context.Context, creditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateCreditNoteRefund(ctx, creditNoteID, payload, raiseOnError, responseFormat...)
}

// GetRefundNotes is the module form of Refunds.GetRefundNotes.
func GetRefundNotes(requester exhttp.Requester, ctx context.Context, refundID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetRefundNotes(ctx, refundID, queryParams, raiseOnError, responseFormat...)
}

// GetRefundNote is the module form of Refunds.GetRefundNote.
func GetRefundNote(requester exhttp.Requester, ctx context.Context, refundID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetRefundNote(ctx, refundID, noteID, raiseOnError, responseFormat...)
}

// CreateRefundNote is the module form of Refunds.CreateRefundNote.
func CreateRefundNote(requester exhttp.Requester, ctx context.Context, refundID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateRefundNote(ctx, refundID, note, files, raiseOnError, responseFormat...)
}

// DeleteRefundNote is the module form of Refunds.DeleteRefundNote.
func DeleteRefundNote(requester exhttp.Requester, ctx context.Context, refundID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteRefundNote(ctx, refundID, noteID, raiseOnError, responseFormat...)
}

// GetRefundNoteFiles is the module form of Refunds.GetRefundNoteFiles.
func GetRefundNoteFiles(requester exhttp.Requester, ctx context.Context, refundID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetRefundNoteFiles(ctx, refundID, noteID, raiseOnError, responseFormat...)
}

// GetRefundNoteFile is the module form of Refunds.GetRefundNoteFile.
func GetRefundNoteFile(requester exhttp.Requester, ctx context.Context, refundID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetRefundNoteFile(ctx, refundID, noteID, fileID, raiseOnError, responseFormat...)
}

// DeleteRefundNoteFile is the module form of Refunds.DeleteRefundNoteFile.
func DeleteRefundNoteFile(requester exhttp.Requester, ctx context.Context, refundID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteRefundNoteFile(ctx, refundID, noteID, fileID, raiseOnError, responseFormat...)
}
