// Package quoteorders is the v4 quote orders endpoints.
package quoteorders

// The path templates for the quote orders endpoints; the client fills {api_version}.
const (
	QuoteOrderInvoicesPath = "/api/{api_version}/quote_orders/{quote_order_id}/invoices"
)

// QuoteOrderInvoicesKey is the key a listing's records arrive under, the v4 list envelope's data.
const QuoteOrderInvoicesKey = "data"
