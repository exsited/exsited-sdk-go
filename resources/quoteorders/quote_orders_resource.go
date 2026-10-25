package quoteorders

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the QuoteOrders method of the same name,
// which is the primary surface.

// GetQuoteOrderInvoices is the module form of QuoteOrders.GetQuoteOrderInvoices.
func GetQuoteOrderInvoices(requester exhttp.Requester, ctx context.Context, quoteOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetQuoteOrderInvoices(ctx, quoteOrderID, queryParams, raiseOnError, responseFormat...)
}
