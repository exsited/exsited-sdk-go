package quoteorders

import "github.com/exsited/exsited-sdk-go/v4/dto"

// QuoteOrderInvoiceList is one page of a quote's invoices, under data.
type QuoteOrderInvoiceList struct {
	Data       []QuoteOrderInvoice
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints QuoteOrderInvoiceList was measured from.
func (QuoteOrderInvoiceList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/quote_orders/{quote_order_id}/invoices",
	}
}

// QuoteOrderInvoice is one entry of a QuoteOrderInvoiceList's data.
type QuoteOrderInvoice struct {
	Account                   *QuoteOrderInvoiceAccount
	AlternateBillingEndDate   *string
	AlternateBillingStartDate *string
	AlternateDueDate          *string
	AlternateIssueDate        *string
	BillingEndDate            *string
	BillingStartDate          *string
	CreatedBy                 *QuoteOrderInvoiceCreatedBy
	CreatedOn                 *string
	Currency                  *QuoteOrderInvoiceCurrency
	CustomForm                *QuoteOrderInvoiceCustomForm
	CustomerPurchaseOrderID   any
	DiscountAmount            *float64
	Due                       *string
	DueDate                   *string
	ID                        *string
	InvoiceNote               any
	IssueDate                 *string
	LastPayment               any
	LastUpdatedBy             any
	LastUpdatedOn             any
	Object                    *string
	Order                     *QuoteOrderInvoiceOrder
	Origin                    any
	Paid                      *string
	PaymentStatus             *string
	PriceTaxInclusive         *bool
	Status                    *string
	Subtotal                  *string
	Tax                       *string
	TaxTotal                  *string
	Total                     *string
	Type                      *string
	Version                   *int64
}

// QuoteOrderInvoiceAccount is a QuoteOrderInvoice's account.
type QuoteOrderInvoiceAccount struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// QuoteOrderInvoiceCreatedBy is a QuoteOrderInvoice's created_by.
type QuoteOrderInvoiceCreatedBy struct {
	Name   *string
	Object *string
}

// QuoteOrderInvoiceCurrency is a QuoteOrderInvoice's currency.
type QuoteOrderInvoiceCurrency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// QuoteOrderInvoiceCustomForm is a QuoteOrderInvoice's custom_form.
type QuoteOrderInvoiceCustomForm struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// QuoteOrderInvoiceOrder is a QuoteOrderInvoice's order.
type QuoteOrderInvoiceOrder struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Pagination is a list envelope's paging block. One response fills only one shape: cursor (HasMore
// present), offset (Records present) or the external database listings' (HasNext present), and the
// others read nil.
type Pagination struct {
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	PrevCursor *string
	Links      *PaginationLinks

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	NextPage     *string
	PreviousPage *string

	// The external database listings' shape.
	HasNext *bool
	HasPrev *bool
	PerPage *int64
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// QuoteOrderInvoiceListShape reads a body as QuoteOrderInvoiceList.
func QuoteOrderInvoiceListShape(decoded any) any { return dto.Read[QuoteOrderInvoiceList](decoded) }
