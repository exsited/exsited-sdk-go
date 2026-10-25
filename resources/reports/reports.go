package reports

import (
	"context"
	"fmt"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Reports is the v4 reports endpoints, bound to one client.
type Reports struct {
	requester exhttp.Requester
}

// GetReportsAccountsAccountOverview sends GET /api/{api_version}/reports/accounts/account_overview.
func (r *Reports) GetReportsAccountsAccountOverview(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReportsAccountsAccountOverviewPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ReportsAccountsAccountOverviewRecordShape})
}

// GetReportsInvoicesInvoiceOverview sends GET /api/{api_version}/reports/invoices/invoice_overview.
func (r *Reports) GetReportsInvoicesInvoiceOverview(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReportsInvoicesInvoiceOverviewPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ReportsInvoicesInvoiceOverviewReadRecordShape})
}

// GetReportsOrdersOrderOverview sends GET /api/{api_version}/reports/orders/order_overview.
func (r *Reports) GetReportsOrdersOrderOverview(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReportsOrdersOrderOverviewPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ReportsOrdersOrderOverviewReadRecordShape})
}

// GetReportsPurchaseInvoicesPurchaseInvoiceOverview sends GET /api/{api_version}/reports/purchase_invoices/purchase_invoice_overview.
func (r *Reports) GetReportsPurchaseInvoicesPurchaseInvoiceOverview(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReportsPurchaseInvoicesPurchaseInvoiceOverviewPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecordShape})
}

// GetReportsPurchaseOrdersPurchaseOrderOverview sends GET /api/{api_version}/reports/purchase_orders/purchase_order_overview.
func (r *Reports) GetReportsPurchaseOrdersPurchaseOrderOverview(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReportsPurchaseOrdersPurchaseOrderOverviewPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ReportsPurchaseOrdersPurchaseOrderOverviewReadRecordShape})
}

// GetReportsPurchaseTransactionsPurchasePaymentAndRefundOverview sends GET /api/{api_version}/reports/purchase_transactions/purchase_payment_and_refund_overview.
func (r *Reports) GetReportsPurchaseTransactionsPurchasePaymentAndRefundOverview(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecordShape})
}

// GetReportsTransactionsPaymentAndRefundOverview sends GET /api/{api_version}/reports/transactions/payment_and_refund_overview.
func (r *Reports) GetReportsTransactionsPaymentAndRefundOverview(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReportsTransactionsPaymentAndRefundOverviewPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ReportsTransactionsPaymentAndRefundOverviewReadRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "reports"

// oneResponseFormat reads the optional trailing responseFormat, refusing more than one.
func oneResponseFormat(responseFormat []string) (string, error) {
	switch len(responseFormat) {
	case 0:
		return "", nil
	case 1:
		return responseFormat[0], nil
	default:
		return "", exerrors.NewUsageError(fmt.Sprintf(
			"responseFormat: pass one format or none, got %d; the trailing argument is optional, "+
				"and omitting it defers to the client", len(responseFormat)))
	}
}

// New binds the reports endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Reports { return &Reports{requester: requester} }

// Requester returns the requester this resource was bound to.
func (r *Reports) Requester() exhttp.Requester { return r.requester }
