package reports

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Reports method of the same name,
// which is the primary surface.

// GetReportsAccountsAccountOverview is the module form of Reports.GetReportsAccountsAccountOverview.
func GetReportsAccountsAccountOverview(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReportsAccountsAccountOverview(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetReportsInvoicesInvoiceOverview is the module form of Reports.GetReportsInvoicesInvoiceOverview.
func GetReportsInvoicesInvoiceOverview(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReportsInvoicesInvoiceOverview(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetReportsOrdersOrderOverview is the module form of Reports.GetReportsOrdersOrderOverview.
func GetReportsOrdersOrderOverview(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReportsOrdersOrderOverview(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetReportsPurchaseInvoicesPurchaseInvoiceOverview is the module form of Reports.GetReportsPurchaseInvoicesPurchaseInvoiceOverview.
func GetReportsPurchaseInvoicesPurchaseInvoiceOverview(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReportsPurchaseInvoicesPurchaseInvoiceOverview(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetReportsPurchaseOrdersPurchaseOrderOverview is the module form of Reports.GetReportsPurchaseOrdersPurchaseOrderOverview.
func GetReportsPurchaseOrdersPurchaseOrderOverview(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReportsPurchaseOrdersPurchaseOrderOverview(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetReportsPurchaseTransactionsPurchasePaymentAndRefundOverview is the module form of Reports.GetReportsPurchaseTransactionsPurchasePaymentAndRefundOverview.
func GetReportsPurchaseTransactionsPurchasePaymentAndRefundOverview(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReportsPurchaseTransactionsPurchasePaymentAndRefundOverview(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetReportsTransactionsPaymentAndRefundOverview is the module form of Reports.GetReportsTransactionsPaymentAndRefundOverview.
func GetReportsTransactionsPaymentAndRefundOverview(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReportsTransactionsPaymentAndRefundOverview(ctx, queryParams, raiseOnError, responseFormat...)
}
