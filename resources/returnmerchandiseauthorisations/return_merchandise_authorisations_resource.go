package returnmerchandiseauthorisations

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ReturnMerchandiseAuthorisations method of the same name,
// which is the primary surface.

// GetReturnMerchandiseAuthorisations is the module form of ReturnMerchandiseAuthorisations.GetReturnMerchandiseAuthorisations.
func GetReturnMerchandiseAuthorisations(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReturnMerchandiseAuthorisations(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateReturnMerchandiseAuthorisation is the module form of ReturnMerchandiseAuthorisations.CreateReturnMerchandiseAuthorisation.
func CreateReturnMerchandiseAuthorisation(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateReturnMerchandiseAuthorisation(ctx, payload, raiseOnError, responseFormat...)
}

// GetReturnMerchandiseAuthorisation is the module form of ReturnMerchandiseAuthorisations.GetReturnMerchandiseAuthorisation.
func GetReturnMerchandiseAuthorisation(requester exhttp.Requester, ctx context.Context, rmaID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReturnMerchandiseAuthorisation(ctx, rmaID, queryParams, raiseOnError, responseFormat...)
}

// GetInvoiceReturnMerchandiseAuthorisationReceive is the module form of ReturnMerchandiseAuthorisations.GetInvoiceReturnMerchandiseAuthorisationReceive.
func GetInvoiceReturnMerchandiseAuthorisationReceive(requester exhttp.Requester, ctx context.Context, invoiceID any, rmaID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceReturnMerchandiseAuthorisationReceive(ctx, invoiceID, rmaID, queryParams, raiseOnError, responseFormat...)
}

// GetInvoiceReturnMerchandiseAuthorisationReceiveByUUID is the module form of ReturnMerchandiseAuthorisations.GetInvoiceReturnMerchandiseAuthorisationReceiveByUUID.
func GetInvoiceReturnMerchandiseAuthorisationReceiveByUUID(requester exhttp.Requester, ctx context.Context, invoiceID any, rmaID any, receiveUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceReturnMerchandiseAuthorisationReceiveByUUID(ctx, invoiceID, rmaID, receiveUUID, queryParams, raiseOnError, responseFormat...)
}

// CreateReturnMerchandiseAuthorisationNote is the module form of ReturnMerchandiseAuthorisations.CreateReturnMerchandiseAuthorisationNote.
func CreateReturnMerchandiseAuthorisationNote(requester exhttp.Requester, ctx context.Context, rmaID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateReturnMerchandiseAuthorisationNote(ctx, rmaID, payload, raiseOnError, responseFormat...)
}

// ReceiveReturnMerchandiseAuthorisation is the module form of ReturnMerchandiseAuthorisations.ReceiveReturnMerchandiseAuthorisation.
func ReceiveReturnMerchandiseAuthorisation(requester exhttp.Requester, ctx context.Context, rmaID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReceiveReturnMerchandiseAuthorisation(ctx, rmaID, payload, raiseOnError, responseFormat...)
}

// DeleteReturnMerchandiseAuthorisation is the module form of ReturnMerchandiseAuthorisations.DeleteReturnMerchandiseAuthorisation.
func DeleteReturnMerchandiseAuthorisation(requester exhttp.Requester, ctx context.Context, rmaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteReturnMerchandiseAuthorisation(ctx, rmaID, raiseOnError, responseFormat...)
}

// UpdateReturnMerchandiseAuthorisation is the module form of ReturnMerchandiseAuthorisations.UpdateReturnMerchandiseAuthorisation.
func UpdateReturnMerchandiseAuthorisation(requester exhttp.Requester, ctx context.Context, rmaID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateReturnMerchandiseAuthorisation(ctx, rmaID, payload, raiseOnError, responseFormat...)
}

// ApproveReturnMerchandiseAuthorisation is the module form of ReturnMerchandiseAuthorisations.ApproveReturnMerchandiseAuthorisation.
func ApproveReturnMerchandiseAuthorisation(requester exhttp.Requester, ctx context.Context, rmaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ApproveReturnMerchandiseAuthorisation(ctx, rmaID, raiseOnError, responseFormat...)
}

// DeleteReturnMerchandiseAuthorisationNote is the module form of ReturnMerchandiseAuthorisations.DeleteReturnMerchandiseAuthorisationNote.
func DeleteReturnMerchandiseAuthorisationNote(requester exhttp.Requester, ctx context.Context, rmaID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteReturnMerchandiseAuthorisationNote(ctx, rmaID, noteID, raiseOnError, responseFormat...)
}

// DeleteReturnMerchandiseAuthorisationNoteFile is the module form of ReturnMerchandiseAuthorisations.DeleteReturnMerchandiseAuthorisationNoteFile.
func DeleteReturnMerchandiseAuthorisationNoteFile(requester exhttp.Requester, ctx context.Context, rmaID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteReturnMerchandiseAuthorisationNoteFile(ctx, rmaID, noteID, fileID, raiseOnError, responseFormat...)
}

// RejectReturnMerchandiseAuthorisation is the module form of ReturnMerchandiseAuthorisations.RejectReturnMerchandiseAuthorisation.
func RejectReturnMerchandiseAuthorisation(requester exhttp.Requester, ctx context.Context, rmaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).RejectReturnMerchandiseAuthorisation(ctx, rmaID, raiseOnError, responseFormat...)
}

// GetReturnMerchandiseAuthorisationNotes is the module form of ReturnMerchandiseAuthorisations.GetReturnMerchandiseAuthorisationNotes.
func GetReturnMerchandiseAuthorisationNotes(requester exhttp.Requester, ctx context.Context, rmaCode any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReturnMerchandiseAuthorisationNotes(ctx, rmaCode, queryParams, raiseOnError, responseFormat...)
}

// GetReturnMerchandiseAuthorisationNote is the module form of ReturnMerchandiseAuthorisations.GetReturnMerchandiseAuthorisationNote.
func GetReturnMerchandiseAuthorisationNote(requester exhttp.Requester, ctx context.Context, rmaCode any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReturnMerchandiseAuthorisationNote(ctx, rmaCode, noteID, raiseOnError, responseFormat...)
}

// GetReturnMerchandiseAuthorisationNoteFiles is the module form of ReturnMerchandiseAuthorisations.GetReturnMerchandiseAuthorisationNoteFiles.
func GetReturnMerchandiseAuthorisationNoteFiles(requester exhttp.Requester, ctx context.Context, rmaCode any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReturnMerchandiseAuthorisationNoteFiles(ctx, rmaCode, noteID, raiseOnError, responseFormat...)
}

// GetReturnMerchandiseAuthorisationNoteFile is the module form of ReturnMerchandiseAuthorisations.GetReturnMerchandiseAuthorisationNoteFile.
func GetReturnMerchandiseAuthorisationNoteFile(requester exhttp.Requester, ctx context.Context, rmaCode any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReturnMerchandiseAuthorisationNoteFile(ctx, rmaCode, noteID, fileID, raiseOnError, responseFormat...)
}

// GetReturnMerchandiseAuthorisationReceive is the module form of ReturnMerchandiseAuthorisations.GetReturnMerchandiseAuthorisationReceive.
func GetReturnMerchandiseAuthorisationReceive(requester exhttp.Requester, ctx context.Context, rmaCode any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReturnMerchandiseAuthorisationReceive(ctx, rmaCode, queryParams, raiseOnError, responseFormat...)
}

// GetReturnMerchandiseAuthorisationReceiveByUUID is the module form of ReturnMerchandiseAuthorisations.GetReturnMerchandiseAuthorisationReceiveByUUID.
func GetReturnMerchandiseAuthorisationReceiveByUUID(requester exhttp.Requester, ctx context.Context, rmaCode any, receiveUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReturnMerchandiseAuthorisationReceiveByUUID(ctx, rmaCode, receiveUUID, raiseOnError, responseFormat...)
}
