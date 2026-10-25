package purchaserefunds

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the PurchaseRefunds method of the same name,
// which is the primary surface.

// GetPurchaseRefunds is the module form of PurchaseRefunds.GetPurchaseRefunds.
func GetPurchaseRefunds(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseRefunds(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseRefund is the module form of PurchaseRefunds.GetPurchaseRefund.
func GetPurchaseRefund(requester exhttp.Requester, ctx context.Context, purchaseRefundID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseRefund(ctx, purchaseRefundID, queryParams, raiseOnError, responseFormat...)
}

// CreatePurchaseCreditNotePurchaseRefund is the module form of PurchaseRefunds.CreatePurchaseCreditNotePurchaseRefund.
func CreatePurchaseCreditNotePurchaseRefund(requester exhttp.Requester, ctx context.Context, purchaseCreditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchaseCreditNotePurchaseRefund(ctx, purchaseCreditNoteID, payload, raiseOnError, responseFormat...)
}

// UpdatePurchaseRefund is the module form of PurchaseRefunds.UpdatePurchaseRefund.
func UpdatePurchaseRefund(requester exhttp.Requester, ctx context.Context, purchaseRefundID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdatePurchaseRefund(ctx, purchaseRefundID, payload, raiseOnError, responseFormat...)
}

// DeletePurchaseRefund is the module form of PurchaseRefunds.DeletePurchaseRefund.
func DeletePurchaseRefund(requester exhttp.Requester, ctx context.Context, purchaseRefundID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeletePurchaseRefund(ctx, purchaseRefundID, raiseOnError, responseFormat...)
}
