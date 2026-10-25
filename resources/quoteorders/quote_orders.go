package quoteorders

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// QuoteOrders is the v4 quote orders endpoints, bound to one client.
type QuoteOrders struct {
	requester exhttp.Requester
}

// GetQuoteOrderInvoices sends GET /api/{api_version}/quote_orders/{quote_order_id}/invoices.
func (q *QuoteOrders) GetQuoteOrderInvoices(ctx context.Context, quoteOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(quoteOrderID, "quote_order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return q.requester.SendGetRequest(ctx, QuoteOrderInvoicesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"quote_order_id": quoteOrderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: QuoteOrderInvoiceListShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "quote_orders"

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

// isNilValue reports whether a value is a typed nil inside a non-nil any; a zero-valued struct is not.
func isNilValue(value any) bool {
	reflected := reflect.ValueOf(value)
	switch reflected.Kind() {
	case reflect.Map, reflect.Slice, reflect.Pointer, reflect.Interface, reflect.Func:
		return reflected.IsNil()
	}
	return false
}

// requireValue refuses a path value that is nil, empty or only whitespace, before any I/O.
func requireValue(value any, what string) error {
	text, isText := value.(string)
	if value == nil || isNilValue(value) || isText && strings.TrimSpace(text) == "" {
		return exerrors.UsageErrorf("%s is required: the request cannot be built without it", what)
	}
	return nil
}

// New binds the quote orders endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *QuoteOrders { return &QuoteOrders{requester: requester} }

// Requester returns the requester this resource was bound to.
func (q *QuoteOrders) Requester() exhttp.Requester { return q.requester }
