// Package invoiceterms is the v4 invoice terms endpoints.
package invoiceterms

// The path templates for the invoice terms endpoints; the client fills {api_version}.
const (
	SettingsInvoiceTermsPath = "/api/{api_version}/settings/invoice_terms"
	SettingsInvoiceTermPath  = "/api/{api_version}/settings/invoice_terms/{invoice_term_id}"
)

// SettingsInvoiceTermsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsInvoiceTermsKey = "data"

// InvoiceTermKey is the envelope key one record arrives under.
const InvoiceTermKey = "invoice_term"
