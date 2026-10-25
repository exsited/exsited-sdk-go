package purchaseorders

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the PurchaseOrders method of the same name,
// which is the primary surface.

// GetPurchaseOrders is the module form of PurchaseOrders.GetPurchaseOrders.
func GetPurchaseOrders(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrders(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreatePurchaseOrder is the module form of PurchaseOrders.CreatePurchaseOrder.
func CreatePurchaseOrder(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchaseOrder(ctx, payload, raiseOnError, responseFormat...)
}

// GetPurchaseOrdersNotes is the module form of PurchaseOrders.GetPurchaseOrdersNotes.
func GetPurchaseOrdersNotes(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrdersNotes(ctx, queryParams, raiseOnError, responseFormat...)
}

// DeletePurchaseOrder is the module form of PurchaseOrders.DeletePurchaseOrder.
func DeletePurchaseOrder(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePurchaseOrder(ctx, purchaseOrderID, raiseOnError, responseFormat...)
}

// UpdatePurchaseOrder is the module form of PurchaseOrders.UpdatePurchaseOrder.
func UpdatePurchaseOrder(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdatePurchaseOrder(ctx, purchaseOrderID, payload, raiseOnError, responseFormat...)
}

// GetPurchaseOrder is the module form of PurchaseOrders.GetPurchaseOrder.
func GetPurchaseOrder(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrder(ctx, purchaseOrderID, queryParams, raiseOnError, responseFormat...)
}

// CancelPurchaseOrder is the module form of PurchaseOrders.CancelPurchaseOrder.
func CancelPurchaseOrder(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CancelPurchaseOrder(ctx, purchaseOrderID, payload, raiseOnError, responseFormat...)
}

// ChangePurchaseOrder is the module form of PurchaseOrders.ChangePurchaseOrder.
func ChangePurchaseOrder(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ChangePurchaseOrder(ctx, purchaseOrderID, payload, raiseOnError, responseFormat...)
}

// GetPurchaseOrderInformation is the module form of PurchaseOrders.GetPurchaseOrderInformation.
func GetPurchaseOrderInformation(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrderInformation(ctx, purchaseOrderID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseOrderLines is the module form of PurchaseOrders.GetPurchaseOrderLines.
func GetPurchaseOrderLines(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrderLines(ctx, purchaseOrderID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseOrderNotes is the module form of PurchaseOrders.GetPurchaseOrderNotes.
func GetPurchaseOrderNotes(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrderNotes(ctx, purchaseOrderID, queryParams, raiseOnError, responseFormat...)
}

// CreatePurchaseOrderNote is the module form of PurchaseOrders.CreatePurchaseOrderNote.
func CreatePurchaseOrderNote(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchaseOrderNote(ctx, purchaseOrderID, note, files, raiseOnError, responseFormat...)
}

// CreatePurchaseOrderPurchaseInvoice is the module form of PurchaseOrders.CreatePurchaseOrderPurchaseInvoice.
func CreatePurchaseOrderPurchaseInvoice(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchaseOrderPurchaseInvoice(ctx, purchaseOrderID, payload, raiseOnError, responseFormat...)
}

// GetPurchaseOrderPurchaseInvoices is the module form of PurchaseOrders.GetPurchaseOrderPurchaseInvoices.
func GetPurchaseOrderPurchaseInvoices(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrderPurchaseInvoices(ctx, purchaseOrderID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseOrderPurchasePayments is the module form of PurchaseOrders.GetPurchaseOrderPurchasePayments.
func GetPurchaseOrderPurchasePayments(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrderPurchasePayments(ctx, purchaseOrderID, queryParams, raiseOnError, responseFormat...)
}

// ReactivatePurchaseOrder is the module form of PurchaseOrders.ReactivatePurchaseOrder.
func ReactivatePurchaseOrder(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReactivatePurchaseOrder(ctx, purchaseOrderID, payload, raiseOnError, responseFormat...)
}

// GetPurchaseOrderLine is the module form of PurchaseOrders.GetPurchaseOrderLine.
func GetPurchaseOrderLine(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, purchaseOrderLineUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrderLine(ctx, purchaseOrderID, purchaseOrderLineUUID, queryParams, raiseOnError, responseFormat...)
}

// DeletePurchaseOrderNote is the module form of PurchaseOrders.DeletePurchaseOrderNote.
func DeletePurchaseOrderNote(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, purchaseOrderNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePurchaseOrderNote(ctx, purchaseOrderID, purchaseOrderNoteID, raiseOnError, responseFormat...)
}

// GetPurchaseOrderNote is the module form of PurchaseOrders.GetPurchaseOrderNote.
func GetPurchaseOrderNote(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, purchaseOrderNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrderNote(ctx, purchaseOrderID, purchaseOrderNoteID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseOrderNoteFiles is the module form of PurchaseOrders.GetPurchaseOrderNoteFiles.
func GetPurchaseOrderNoteFiles(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, purchaseOrderNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrderNoteFiles(ctx, purchaseOrderID, purchaseOrderNoteID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseOrderNoteFile is the module form of PurchaseOrders.GetPurchaseOrderNoteFile.
func GetPurchaseOrderNoteFile(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, purchaseOrderNoteID any, purchaseOrderFileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrderNoteFile(ctx, purchaseOrderID, purchaseOrderNoteID, purchaseOrderFileID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseOrderKPIs is the module form of PurchaseOrders.GetPurchaseOrderKPIs.
func GetPurchaseOrderKPIs(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrderKPIs(ctx, purchaseOrderID, raiseOnError, responseFormat...)
}

// DeletePurchaseOrderNoteFile is the module form of PurchaseOrders.DeletePurchaseOrderNoteFile.
func DeletePurchaseOrderNoteFile(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, purchaseOrderNoteID any, purchaseOrderFileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePurchaseOrderNoteFile(ctx, purchaseOrderID, purchaseOrderNoteID, purchaseOrderFileID, raiseOnError, responseFormat...)
}
