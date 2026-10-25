// Package reports is the v4 reports endpoints.
package reports

// The path templates for the reports endpoints; the client fills {api_version}.
const (
	ReportsAccountsAccountOverviewPath                              = "/api/{api_version}/reports/accounts/account_overview"
	ReportsInvoicesInvoiceOverviewPath                              = "/api/{api_version}/reports/invoices/invoice_overview"
	ReportsOrdersOrderOverviewPath                                  = "/api/{api_version}/reports/orders/order_overview"
	ReportsPurchaseInvoicesPurchaseInvoiceOverviewPath              = "/api/{api_version}/reports/purchase_invoices/purchase_invoice_overview"
	ReportsPurchaseOrdersPurchaseOrderOverviewPath                  = "/api/{api_version}/reports/purchase_orders/purchase_order_overview"
	ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewPath = "/api/{api_version}/reports/purchase_transactions/purchase_payment_and_refund_overview"
	ReportsTransactionsPaymentAndRefundOverviewPath                 = "/api/{api_version}/reports/transactions/payment_and_refund_overview"
)

// ReportKey is the envelope key one record arrives under.
const ReportKey = "report"
