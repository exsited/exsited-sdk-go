package reports

import "github.com/exsited/exsited-sdk-go/v4/dto"

// Pagination is the paging block, either shape.
type Pagination struct {
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	PrevCursor any

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	NextPage     *string
	PreviousPage *string
}

// ReportsAccountsAccountOverviewRecord is the reports accounts account overview body.
type ReportsAccountsAccountOverviewRecord struct {
	Accounts *Accounts
}

// MeasuredFrom lists the endpoints ReportsAccountsAccountOverviewRecord was measured from.
func (ReportsAccountsAccountOverviewRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/reports/accounts/account_overview",
	}
}

// EnvelopeKey unwraps {"report": {...}}.
func (ReportsAccountsAccountOverviewRecord) EnvelopeKey() string { return ReportKey }

// ReportsInvoicesInvoiceOverviewReadRecord is the body GetReportsInvoicesInvoiceOverview answers.
type ReportsInvoicesInvoiceOverviewReadRecord struct {
	Invoices *Invoices
}

// MeasuredFrom lists the endpoints ReportsInvoicesInvoiceOverviewReadRecord was measured from.
func (ReportsInvoicesInvoiceOverviewReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/reports/invoices/invoice_overview",
	}
}

// EnvelopeKey unwraps {"report": {...}}.
func (ReportsInvoicesInvoiceOverviewReadRecord) EnvelopeKey() string { return ReportKey }

// ReportsOrdersOrderOverviewReadRecord is the body GetReportsOrdersOrderOverview answers.
type ReportsOrdersOrderOverviewReadRecord struct {
	Orders *Orders
}

// MeasuredFrom lists the endpoints ReportsOrdersOrderOverviewReadRecord was measured from.
func (ReportsOrdersOrderOverviewReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/reports/orders/order_overview",
	}
}

// EnvelopeKey unwraps {"report": {...}}.
func (ReportsOrdersOrderOverviewReadRecord) EnvelopeKey() string { return ReportKey }

// ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecord is the body GetReportsPurchaseInvoicesPurchaseInvoiceOverview answers.
type ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecord struct {
	PurchaseInvoices *PurchaseInvoices
}

// MeasuredFrom lists the endpoints ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecord was measured from.
func (ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/reports/purchase_invoices/purchase_invoice_overview",
	}
}

// EnvelopeKey unwraps {"report": {...}}.
func (ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecord) EnvelopeKey() string {
	return ReportKey
}

// ReportsPurchaseOrdersPurchaseOrderOverviewReadRecord is the body GetReportsPurchaseOrdersPurchaseOrderOverview answers.
type ReportsPurchaseOrdersPurchaseOrderOverviewReadRecord struct {
	PurchaseOrders *PurchaseOrders
}

// MeasuredFrom lists the endpoints ReportsPurchaseOrdersPurchaseOrderOverviewReadRecord was measured from.
func (ReportsPurchaseOrdersPurchaseOrderOverviewReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/reports/purchase_orders/purchase_order_overview",
	}
}

// EnvelopeKey unwraps {"report": {...}}.
func (ReportsPurchaseOrdersPurchaseOrderOverviewReadRecord) EnvelopeKey() string { return ReportKey }

// ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecord is the body GetReportsPurchaseTransactionsPurchasePaymentAndRefundOverview answers.
type ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecord struct {
	PurchasePaymentAndRefundOverview *PurchasePaymentAndRefundOverview
}

// MeasuredFrom lists the endpoints ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecord was measured from.
func (ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/reports/purchase_transactions/purchase_payment_and_refund_overview",
	}
}

// EnvelopeKey unwraps {"report": {...}}.
func (ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecord) EnvelopeKey() string {
	return ReportKey
}

// ReportsTransactionsPaymentAndRefundOverviewReadRecord is the body GetReportsTransactionsPaymentAndRefundOverview answers.
type ReportsTransactionsPaymentAndRefundOverviewReadRecord struct {
	PaymentAndRefundOverview *PaymentAndRefundOverview
}

// MeasuredFrom lists the endpoints ReportsTransactionsPaymentAndRefundOverviewReadRecord was measured from.
func (ReportsTransactionsPaymentAndRefundOverviewReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/reports/transactions/payment_and_refund_overview",
	}
}

// EnvelopeKey unwraps {"report": {...}}.
func (ReportsTransactionsPaymentAndRefundOverviewReadRecord) EnvelopeKey() string { return ReportKey }

// Amended is the amended object.
type Amended struct {
	Amount *string
	Count  *string
}

// Cancelled is the cancelled object.
type Cancelled struct {
	Amount *string
	Count  *string
}

// Changed is the changed object.
type Changed struct {
	Amount *string
	Count  *string
}

// Deleted is the deleted object.
type Deleted struct {
	Amount *string
	Count  *string
}

// InvoiceOverview is one entry of the invoice_overview list.
type InvoiceOverview struct {
	Amended      *Amended
	Cancelled    *Cancelled
	Date         *string
	InvoiceTotal *string
	NetGain      *NetGain
	New          *InvoiceOverviewNew
	Reactivated  *Reactivated
	Voided       *Voided
}

// InvoiceOverviewNew is the new object.
type InvoiceOverviewNew struct {
	Amount *string
	Count  *string
}

// Invoices is the invoices object.
type Invoices struct {
	InvoiceOverview []InvoiceOverview
}

// NetGain is the net_gain object.
type NetGain struct {
	Amount *string
	Count  *string
}

// NetPurchases is the net_purchases object.
type NetPurchases struct {
	Amount *string
	Count  *string
}

// OrderOverview is one entry of the order_overview list.
type OrderOverview struct {
	Cancelled   *Cancelled
	Changed     *Changed
	Date        *string
	Deleted     *Deleted
	NetGain     *NetGain
	New         *OrderOverviewNew
	OrderTotal  *string
	Reactivated *Reactivated
}

// OrderOverviewNew is the new object.
type OrderOverviewNew struct {
	Amount *string
	Count  *string
}

// Orders is the orders object.
type Orders struct {
	OrderOverview []OrderOverview
}

// Payment is one entry of the payments list.
type Payment struct {
	Date           *string
	Deleted        *Deleted
	New            *PaymentNew
	PaymentApplied *string
}

// PaymentAndRefundOverview is the payment_and_refund_overview object.
type PaymentAndRefundOverview struct {
	Payments []Payment
	Refunds  []Refund
}

// PaymentNew is the new object.
type PaymentNew struct {
	Amount *string
	Count  *string
}

// PurchaseInvoiceOverview is one entry of the purchase_invoice_overview list.
type PurchaseInvoiceOverview struct {
	Amended              *Amended
	Cancelled            *Cancelled
	Date                 *string
	NetGain              *NetGain
	New                  *PurchaseInvoiceOverviewNew
	PurchaseInvoiceTotal *string
	Reactivated          *Reactivated
	Voided               *Voided
}

// PurchaseInvoiceOverviewNew is the new object.
type PurchaseInvoiceOverviewNew struct {
	Amount *string
	Count  *string
}

// PurchaseInvoices is the purchase_invoices object.
type PurchaseInvoices struct {
	PurchaseInvoiceOverview []PurchaseInvoiceOverview
}

// PurchaseOrderOverview is one entry of the purchase_order_overview list.
type PurchaseOrderOverview struct {
	Cancelled          *Cancelled
	Changed            *Changed
	Date               *string
	Deleted            *Deleted
	NetPurchases       *NetPurchases
	New                *PurchaseOrderOverviewNew
	PurchaseOrderTotal *string
	Reactivated        *Reactivated
}

// PurchaseOrderOverviewNew is the new object.
type PurchaseOrderOverviewNew struct {
	Amount *string
	Count  *string
}

// PurchaseOrders is the purchase_orders object.
type PurchaseOrders struct {
	PurchaseOrderOverview []PurchaseOrderOverview
}

// PurchasePayment is one entry of the purchase_payments list.
type PurchasePayment struct {
	Date                   *string
	Deleted                *Deleted
	New                    *PurchasePaymentNew
	PurchasePaymentApplied *string
}

// PurchasePaymentAndRefundOverview is the purchase_payment_and_refund_overview object.
type PurchasePaymentAndRefundOverview struct {
	PurchasePayments []PurchasePayment
	PurchaseRefunds  []PurchaseRefund
}

// PurchasePaymentNew is the new object.
type PurchasePaymentNew struct {
	Amount *string
	Count  *string
}

// PurchaseRefund is one entry of the purchase_refunds list.
type PurchaseRefund struct {
	Date    *string
	Deleted *Deleted
	New     *PurchaseRefundNew
}

// PurchaseRefundNew is the new object.
type PurchaseRefundNew struct {
	Amount *string
	Count  *string
}

// Reactivated is the reactivated object.
type Reactivated struct {
	Amount *string
	Count  *string
}

// Refund is one entry of the refunds list.
type Refund struct {
	Date    *string
	Deleted *Deleted
	New     *RefundNew
}

// RefundNew is the new object.
type RefundNew struct {
	Amount *string
	Count  *string
}

// Voided is the voided object.
type Voided struct {
	Amount *string
	Count  *string
}

// AccountOverview is one entry of the account_overview list.
type AccountOverview struct {
	Cancelled   *string
	Date        *string
	Deleted     *string
	NetGain     *string
	New         *string
	Reactivated *string
}

// Accounts is the accounts object.
type Accounts struct {
	AccountOverview []AccountOverview
}

// ReportsAccountsAccountOverviewRecordShape reads a body as ReportsAccountsAccountOverviewRecord.
func ReportsAccountsAccountOverviewRecordShape(decoded any) any {
	return dto.Read[ReportsAccountsAccountOverviewRecord](decoded)
}

// ReportsInvoicesInvoiceOverviewReadRecordShape reads a body as ReportsInvoicesInvoiceOverviewReadRecord.
func ReportsInvoicesInvoiceOverviewReadRecordShape(decoded any) any {
	return dto.Read[ReportsInvoicesInvoiceOverviewReadRecord](decoded)
}

// ReportsOrdersOrderOverviewReadRecordShape reads a body as ReportsOrdersOrderOverviewReadRecord.
func ReportsOrdersOrderOverviewReadRecordShape(decoded any) any {
	return dto.Read[ReportsOrdersOrderOverviewReadRecord](decoded)
}

// ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecordShape reads a body as ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecord.
func ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecordShape(decoded any) any {
	return dto.Read[ReportsPurchaseInvoicesPurchaseInvoiceOverviewReadRecord](decoded)
}

// ReportsPurchaseOrdersPurchaseOrderOverviewReadRecordShape reads a body as ReportsPurchaseOrdersPurchaseOrderOverviewReadRecord.
func ReportsPurchaseOrdersPurchaseOrderOverviewReadRecordShape(decoded any) any {
	return dto.Read[ReportsPurchaseOrdersPurchaseOrderOverviewReadRecord](decoded)
}

// ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecordShape reads a body as ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecord.
func ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecordShape(decoded any) any {
	return dto.Read[ReportsPurchaseTransactionsPurchasePaymentAndRefundOverviewReadRecord](decoded)
}

// ReportsTransactionsPaymentAndRefundOverviewReadRecordShape reads a body as ReportsTransactionsPaymentAndRefundOverviewReadRecord.
func ReportsTransactionsPaymentAndRefundOverviewReadRecordShape(decoded any) any {
	return dto.Read[ReportsTransactionsPaymentAndRefundOverviewReadRecord](decoded)
}
