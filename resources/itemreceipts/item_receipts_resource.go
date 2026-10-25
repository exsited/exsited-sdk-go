package itemreceipts

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ItemReceipts method of the same name,
// which is the primary surface.

// GetItemReceipts is the module form of ItemReceipts.GetItemReceipts.
func GetItemReceipts(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemReceipts(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateItemReceipt is the module form of ItemReceipts.CreateItemReceipt.
func CreateItemReceipt(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateItemReceipt(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateItemReceipt is the module form of ItemReceipts.UpdateItemReceipt.
func UpdateItemReceipt(requester exhttp.Requester, ctx context.Context, itemReceiptID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateItemReceipt(ctx, itemReceiptID, payload, raiseOnError, responseFormat...)
}

// GetItemReceipt is the module form of ItemReceipts.GetItemReceipt.
func GetItemReceipt(requester exhttp.Requester, ctx context.Context, itemReceiptID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemReceipt(ctx, itemReceiptID, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseOrderItemReceipts is the module form of ItemReceipts.GetPurchaseOrderItemReceipts.
func GetPurchaseOrderItemReceipts(requester exhttp.Requester, ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseOrderItemReceipts(ctx, purchaseOrderID, queryParams, raiseOnError, responseFormat...)
}

// DeleteItemReceipt is the module form of ItemReceipts.DeleteItemReceipt.
func DeleteItemReceipt(requester exhttp.Requester, ctx context.Context, itemReceiptID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteItemReceipt(ctx, itemReceiptID, raiseOnError, responseFormat...)
}

// DeleteItemReceiptNote is the module form of ItemReceipts.DeleteItemReceiptNote.
func DeleteItemReceiptNote(requester exhttp.Requester, ctx context.Context, itemReceiptID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteItemReceiptNote(ctx, itemReceiptID, noteID, raiseOnError, responseFormat...)
}
