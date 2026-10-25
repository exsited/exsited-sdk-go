// Package proformainvoices is the v4 proforma invoices endpoints.
package proformainvoices

// The path templates for the proforma invoices endpoints; the client fills {api_version}.
const (
	ProformaInvoicesPath        = "/api/{api_version}/proforma_invoices"
	ProformaInvoicePath         = "/api/{api_version}/proforma_invoices/{proforma_id}"
	ProformaInvoicePaymentsPath = "/api/{api_version}/proforma_invoices/{proforma_id}/payments"
)

// ProformaInvoicesKey is the key a listing's records arrive under, the v4 list envelope's data.
const ProformaInvoicesKey = "data"

// ProformaInvoiceKey is the envelope key one record arrives under.
const ProformaInvoiceKey = "proforma_invoice"

// PaymentKey is the envelope key one record arrives under.
const PaymentKey = "payment"
