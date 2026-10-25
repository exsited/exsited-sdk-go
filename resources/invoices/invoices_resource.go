package invoices

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Invoices method of the same name,
// which is the primary surface.

// GetInvoices is the module form of Invoices.GetInvoices.
func GetInvoices(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoices(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateInvoice is the module form of Invoices.CreateInvoice.
func CreateInvoice(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateInvoice(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteInvoice is the module form of Invoices.DeleteInvoice.
func DeleteInvoice(requester exhttp.Requester, ctx context.Context, invoiceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteInvoice(ctx, invoiceID, raiseOnError, responseFormat...)
}

// GetInvoice is the module form of Invoices.GetInvoice.
func GetInvoice(requester exhttp.Requester, ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoice(ctx, invoiceID, queryParams, raiseOnError, responseFormat...)
}

// AmendInvoice is the module form of Invoices.AmendInvoice.
func AmendInvoice(requester exhttp.Requester, ctx context.Context, invoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).AmendInvoice(ctx, invoiceID, payload, raiseOnError, responseFormat...)
}

// CancelInvoice is the module form of Invoices.CancelInvoice.
func CancelInvoice(requester exhttp.Requester, ctx context.Context, invoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CancelInvoice(ctx, invoiceID, payload, raiseOnError, responseFormat...)
}

// UpdateInvoiceInformation is the module form of Invoices.UpdateInvoiceInformation.
func UpdateInvoiceInformation(requester exhttp.Requester, ctx context.Context, invoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateInvoiceInformation(ctx, invoiceID, payload, raiseOnError, responseFormat...)
}

// GetInvoiceInformation is the module form of Invoices.GetInvoiceInformation.
func GetInvoiceInformation(requester exhttp.Requester, ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceInformation(ctx, invoiceID, queryParams, raiseOnError, responseFormat...)
}

// GetInvoiceItemFulfillments is the module form of Invoices.GetInvoiceItemFulfillments.
func GetInvoiceItemFulfillments(requester exhttp.Requester, ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceItemFulfillments(ctx, invoiceID, queryParams, raiseOnError, responseFormat...)
}

// GetInvoiceNotes is the module form of Invoices.GetInvoiceNotes.
func GetInvoiceNotes(requester exhttp.Requester, ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceNotes(ctx, invoiceID, queryParams, raiseOnError, responseFormat...)
}

// CreateInvoiceNote is the module form of Invoices.CreateInvoiceNote.
func CreateInvoiceNote(requester exhttp.Requester, ctx context.Context, invoiceID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateInvoiceNote(ctx, invoiceID, note, files, raiseOnError, responseFormat...)
}

// GetInvoicePDF is the module form of Invoices.GetInvoicePDF.
func GetInvoicePDF(requester exhttp.Requester, ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoicePDF(ctx, invoiceID, queryParams, raiseOnError, responseFormat...)
}

// ReactivateInvoice is the module form of Invoices.ReactivateInvoice.
func ReactivateInvoice(requester exhttp.Requester, ctx context.Context, invoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReactivateInvoice(ctx, invoiceID, payload, raiseOnError, responseFormat...)
}

// GetInvoiceReturnMerchandiseAuthorisations is the module form of Invoices.GetInvoiceReturnMerchandiseAuthorisations.
func GetInvoiceReturnMerchandiseAuthorisations(requester exhttp.Requester, ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceReturnMerchandiseAuthorisations(ctx, invoiceID, queryParams, raiseOnError, responseFormat...)
}

// GetOrderInvoices is the module form of Invoices.GetOrderInvoices.
func GetOrderInvoices(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderInvoices(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// CreateOrderInvoice is the module form of Invoices.CreateOrderInvoice.
func CreateOrderInvoice(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateOrderInvoice(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// GetInvoiceItemFulfillment is the module form of Invoices.GetInvoiceItemFulfillment.
func GetInvoiceItemFulfillment(requester exhttp.Requester, ctx context.Context, invoiceID any, fulfillmentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceItemFulfillment(ctx, invoiceID, fulfillmentID, queryParams, raiseOnError, responseFormat...)
}

// GetInvoiceNote is the module form of Invoices.GetInvoiceNote.
func GetInvoiceNote(requester exhttp.Requester, ctx context.Context, invoiceID any, invoiceNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceNote(ctx, invoiceID, invoiceNoteID, queryParams, raiseOnError, responseFormat...)
}

// DeleteInvoiceNote is the module form of Invoices.DeleteInvoiceNote.
func DeleteInvoiceNote(requester exhttp.Requester, ctx context.Context, invoiceID any, invoiceNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteInvoiceNote(ctx, invoiceID, invoiceNoteID, raiseOnError, responseFormat...)
}

// GetInvoiceReturnMerchandiseAuthorisation is the module form of Invoices.GetInvoiceReturnMerchandiseAuthorisation.
func GetInvoiceReturnMerchandiseAuthorisation(requester exhttp.Requester, ctx context.Context, invoiceID any, rmaID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceReturnMerchandiseAuthorisation(ctx, invoiceID, rmaID, queryParams, raiseOnError, responseFormat...)
}

// GetInvoiceNoteFiles is the module form of Invoices.GetInvoiceNoteFiles.
func GetInvoiceNoteFiles(requester exhttp.Requester, ctx context.Context, invoiceID any, invoiceNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceNoteFiles(ctx, invoiceID, invoiceNoteID, queryParams, raiseOnError, responseFormat...)
}

// GetInvoiceNoteFile is the module form of Invoices.GetInvoiceNoteFile.
func GetInvoiceNoteFile(requester exhttp.Requester, ctx context.Context, invoiceID any, invoiceNoteID any, invoiceFileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceNoteFile(ctx, invoiceID, invoiceNoteID, invoiceFileID, queryParams, raiseOnError, responseFormat...)
}

// GetInvoiceKPIs is the module form of Invoices.GetInvoiceKPIs.
func GetInvoiceKPIs(requester exhttp.Requester, ctx context.Context, invoiceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceKPIs(ctx, invoiceID, raiseOnError, responseFormat...)
}

// GetInvoiceLines is the module form of Invoices.GetInvoiceLines.
func GetInvoiceLines(requester exhttp.Requester, ctx context.Context, invoiceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceLines(ctx, invoiceID, raiseOnError, responseFormat...)
}

// GetInvoiceLine is the module form of Invoices.GetInvoiceLine.
func GetInvoiceLine(requester exhttp.Requester, ctx context.Context, invoiceID any, lineID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceLine(ctx, invoiceID, lineID, raiseOnError, responseFormat...)
}

// GetInvoiceCreditNoteApplication is the module form of Invoices.GetInvoiceCreditNoteApplication.
func GetInvoiceCreditNoteApplication(requester exhttp.Requester, ctx context.Context, invoiceID any, creditNoteApplicationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceCreditNoteApplication(ctx, invoiceID, creditNoteApplicationID, raiseOnError, responseFormat...)
}

// GetInvoiceCustomObjects is the module form of Invoices.GetInvoiceCustomObjects.
func GetInvoiceCustomObjects(requester exhttp.Requester, ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceCustomObjects(ctx, invoiceID, queryParams, raiseOnError, responseFormat...)
}

// GetInvoiceCustomObject is the module form of Invoices.GetInvoiceCustomObject.
func GetInvoiceCustomObject(requester exhttp.Requester, ctx context.Context, invoiceID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceCustomObject(ctx, invoiceID, customObjectID, raiseOnError, responseFormat...)
}

// GetInvoiceCustomObjectInstance is the module form of Invoices.GetInvoiceCustomObjectInstance.
func GetInvoiceCustomObjectInstance(requester exhttp.Requester, ctx context.Context, invoiceID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceCustomObjectInstance(ctx, invoiceID, customObjectID, instanceID, raiseOnError, responseFormat...)
}

// DeleteInvoiceNoteFile is the module form of Invoices.DeleteInvoiceNoteFile.
func DeleteInvoiceNoteFile(requester exhttp.Requester, ctx context.Context, invoiceID any, invoiceNoteID any, invoiceFileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteInvoiceNoteFile(ctx, invoiceID, invoiceNoteID, invoiceFileID, raiseOnError, responseFormat...)
}
