package express

import "github.com/exsited/exsited-sdk-go/v4/dto"

// ExpressCreateRecord is the body CreateExpress answers.
type ExpressCreateRecord struct {
	Account                  *Account
	ActivityLogs             []ActivityLog
	AllowContract            *bool
	BillingAddress           any
	BillingPreferences       *BillingPreferences
	BillingStartDate         *string
	CommunicationPreferences []any
	CreatedBy                any
	CreatedOn                *string
	Currency                 *Currency
	CustomAttributes         []any
	CustomForm               *CustomForm
	CustomerPurchaseOrderID  any
	Description              any
	DiscountProfile          any
	DisplayName              *string
	ID                       *string
	InvoiceNote              any
	Kpis                     *Kpis
	LastUpdatedBy            any
	LastUpdatedOn            any
	Lines                    []Line
	Manager                  any
	Name                     *string
	NextBillingFromDate      *string
	Object                   *string
	OrderStartDate           *string
	OrderType                *string
	Origin                   any
	OriginalTotal            *string
	PreOrder                 *bool
	PriceTaxInclusive        *bool
	QuoteConverted           *bool
	QuoteOrder               *bool
	ReferralAccount          any
	SalesRep                 any
	ShippingAddress          any
	ShippingCost             *float64
	ShippingLineItems        []any
	ShippingProfile          any
	Status                   *string
	Subtotal                 *string
	Tax                      *string
	TimeZone                 *TimeZone
	Total                    *string
	Version                  *int64
}

// MeasuredFrom lists the endpoints ExpressCreateRecord was measured from.
func (ExpressCreateRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/express",
	}
}

// EnvelopeKey unwraps {"express": {...}}.
func (ExpressCreateRecord) EnvelopeKey() string { return ExpressKey }

// Account is the account object.
type Account struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// ActivityLog is one entry of the activity_logs list.
type ActivityLog struct {
	Activity *string
}

// BillingPreferences is the billing_preferences object.
type BillingPreferences struct {
	BillingPeriod        *string
	CommunicationProfile *string
	ConsolidateInvoice   *bool
	FulfillmentMode      *string
	FulfillmentTerm      *string
	InvoiceMode          *string
	InvoiceTerm          *string
	PaymentMode          *string
	PaymentProcessor     *string
	PaymentTerm          *string
	PaymentTermAlignment *string
	ProformaProperties   *ProformaProperties
}

// Currency is the currency object.
type Currency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// CustomForm is the custom_form object.
type CustomForm struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Item is the item object.
type Item struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
	UUID   *string
}

// ItemAccountingCode is the item_accounting_code object.
type ItemAccountingCode struct {
	SalesRevenue *string
}

// ItemPriceSnapshot is the item_price_snapshot object.
type ItemPriceSnapshot struct {
	PricingRule *PricingRule
}

// ItemPriceTax is the item_price_tax object.
type ItemPriceTax struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
	Rate   *float64
}

// ItemSaleTaxConfiguration is the item_sale_tax_configuration object.
type ItemSaleTaxConfiguration struct {
	SalePriceIsBasedOn *string
	TaxCode            *TaxCode
}

// Kpis is the kpis object.
type Kpis struct {
	DeletedAt               any
	EstimatedTotal          *string
	LastCancelledAt         any
	LastChangedAt           any
	LastInvoiceIssueDate    any
	LastInvoiceTotal        *string
	LastReactivatedAt       any
	MonthlyRecurringRevenue *string
	NextInvoiceIssueDate    any
	Object                  *string
	StartDate               any
	TotalCollected          *string
	TotalDue                *string
	TotalInvoice            *int64
	TotalOutstanding        *string
	TotalRevenue            *string
}

// Line is one entry of the lines list.
type Line struct {
	ChargeItemUUID           *string
	ExpectedDeliveryDate     *string
	ID                       *string
	Item                     *Item
	ItemAccountingCode       *ItemAccountingCode
	ItemChargeType           *string
	ItemCustomAttributes     []any
	ItemDescription          *string
	ItemInvoiceNote          *string
	ItemOrderQuantity        *int64
	ItemPriceSnapshot        *ItemPriceSnapshot
	ItemPriceTax             *ItemPriceTax
	ItemSaleTaxConfiguration *ItemSaleTaxConfiguration
	ItemType                 *string
	Object                   *string
	PreOrderStockDetails     *PreOrderStockDetails
	ShippingCost             *float64
	Subtotal                 *string
	Tax                      *string
	TaxExemptWhenSold        *bool
	Total                    *string
	Version                  *int64
}

// PreOrderStockDetails is the pre_order_stock_details object.
type PreOrderStockDetails struct {
	PendingReserve *string
	Reserved       *string
	Sold           *string
}

// PricingRule is the pricing_rule object.
type PricingRule struct {
	Price           *float64
	PricePeriod     *string
	PriceType       *string
	PricingLevel    *string
	PricingMethod   *string
	PricingSchedule *string
	Uom             *string
	UUID            *string
	Version         *string
	Warehouse       *string
}

// ProformaProperties is the proforma_properties object.
type ProformaProperties struct {
	AcceptanceDate     any
	AcceptanceRequired *bool
	AcceptanceTerm     any
	PaymentRequired    *bool
	TermsAndConditions any
}

// TaxCode is the tax_code object.
type TaxCode struct {
	Code *string
	Link *string
	Rate *string
	UUID *string
}

// TimeZone is the time_zone object.
type TimeZone struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// ExpressCreateRecordShape reads a body as ExpressCreateRecord.
func ExpressCreateRecordShape(decoded any) any { return dto.Read[ExpressCreateRecord](decoded) }
