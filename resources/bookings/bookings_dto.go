package bookings

import "github.com/exsited/exsited-sdk-go/v4/dto"

// Pagination is the paging block, either shape.
type Pagination struct {
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	PrevCursor any
	Links      *PaginationLinks

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	NextPage     *string
	PreviousPage *string
}

// BookingRecord is the booking body.
type BookingRecord struct {
	Account                  *Account
	AllowContract            *bool
	BillingAddress           *BillingAddress
	BillingPreferences       *BillingPreferences
	BillingStartDate         *string
	BillingStartDayOfMonth   *string
	CommunicationPreferences []CommunicationPreference
	ContractProperties       *ContractProperties
	CreatedBy                any
	CreatedOn                *string
	Currency                 *Currency
	CustomForm               *CustomForm
	CustomerPurchaseOrderID  any
	Description              *string
	DiscountProfile          *DiscountProfile
	DisplayName              *string
	ID                       *string
	InvoiceNote              *string
	LastUpdatedBy            any
	LastUpdatedOn            *string
	Manager                  *Manager
	Name                     *string
	NextBillingFromDate      *string
	NextBillingFromDateUtc   *string
	Object                   *string
	OrderStartDate           *string
	OrderType                *string
	Origin                   *string
	OriginalTotal            any
	PreOrder                 *bool
	PriceTaxInclusive        *bool
	QuoteConverted           *bool
	QuoteOrder               *bool
	ReferralAccount          *ReferralAccount
	SalesRep                 *SalesRep
	ShippingAddress          *ShippingAddress
	ShippingCost             *float64
	ShippingProfile          *ShippingProfile
	Status                   *string
	TimeZone                 *TimeZone
	Version                  *int64
}

// MeasuredFrom lists the endpoints BookingRecord was measured from.
func (BookingRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking",
	}
}

// BookingList is a page of BookingRecord records.
type BookingList struct {
	Data       []BookingRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints BookingList was measured from.
func (BookingList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking",
	}
}

// BookingByIdRecord is the booking by id body.
type BookingByIdRecord struct {
	Account                  *Account
	AllowContract            *bool
	BillingAddress           *BillingAddress
	BillingPreferences       *BillingPreferences
	BillingStartDate         *string
	CommunicationPreferences []CommunicationPreference
	CreatedBy                any
	CreatedOn                *string
	Currency                 *Currency
	CustomAttributes         []CustomAttribute
	CustomForm               *CustomForm
	CustomerPurchaseOrderID  any
	Description              *string
	DiscountProfile          any
	DisplayName              *string
	ID                       *string
	InvoiceNote              *string
	Kpis                     *Kpis
	LastUpdatedBy            any
	LastUpdatedOn            *string
	Lines                    []Line
	Manager                  *Manager
	Name                     *string
	NextBillingFromDate      *string
	NextBillingFromDateUtc   *string
	Object                   *string
	OrderStartDate           *string
	OrderType                *string
	Origin                   *string
	OriginalTotal            any
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

// MeasuredFrom lists the endpoints BookingByIdRecord was measured from.
func (BookingByIdRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (BookingByIdRecord) EnvelopeKey() string { return OrderKey }

// ContractProperties is the contract_properties object.
type ContractProperties struct {
	AllowDowngrade                       *bool
	AllowDowngradeCharge                 *bool
	AllowEarlyTermination                *bool
	AllowPostponement                    *bool
	AllowTrial                           *bool
	AllowUpgrade                         *bool
	AutoRenewalEndsOn                    *string
	AutoRenewalRequireCustomerAcceptance *bool
	AutoRenewalTerm                      *string
	DowngradeChargeFixed                 *string
	DowngradeChargeType                  *string
	InitialContractTerm                  *string
	PeriodBeforeDowngrade                *string
	RenewAutomatically                   *string
	RequireCustomerAcceptance            *bool
	RequiresPaymentMethod                *bool
	TerminationAccountingCode            *string
	TerminationNoticePeriod              *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// Account is the account object.
type Account struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BillingAddress is the billing_address object.
type BillingAddress struct {
	AddressLine1 *string
	AddressLine2 *string
	AddressLine3 *string
	AddressLine4 *string
	AddressLine5 *string
	City         *string
	Country      *string
	Name         *string
	PostalCode   *string
	State        *string
}

// BillingPreferences is the billing_preferences object.
type BillingPreferences struct {
	BillingPeriod        *string
	CommunicationProfile *string
	ConsolidateInvoice   *bool
	ConsolidateKey       *string
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

// CommunicationPreference is one entry of the communication_preferences list.
type CommunicationPreference struct {
	Channel *string
	Enabled *bool
}

// Currency is the currency object.
type Currency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// CustomAttribute is one entry of the custom_attributes list.
type CustomAttribute struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
	Type   *string
	Value  any
}

// CustomForm is the custom_form object.
type CustomForm struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// DiscountProfile is the discount_profile object.
type DiscountProfile struct {
	Amount         *string
	Href           *string
	ID             *string
	Name           *string
	Object         *string
	RedemptionCode *string
	Type           *string
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
	SalesRevenue any
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

// ItemProperties is the item_properties object.
type ItemProperties struct {
	BillingMode                 *string
	ChargingAndBillingAlignment *string
	ChargingStartDate           *string
	ProRataPartialPricingPeriod *string
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
	ItemProperties           *ItemProperties
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

// Manager is the manager object.
type Manager struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
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

// ReferralAccount is the referral_account object.
type ReferralAccount struct {
	Href   *string
	ID     *string
	Object *string
}

// SalesRep is the sales_rep object.
type SalesRep struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// ShippingAddress is the shipping_address object.
type ShippingAddress struct {
	AddressLine1 *string
	AddressLine2 *string
	AddressLine3 *string
	AddressLine4 *string
	AddressLine5 *string
	City         *string
	Country      *string
	Name         *string
	PostalCode   *string
	State        *string
}

// ShippingProfile is the shipping_profile object.
type ShippingProfile struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// TaxCode is the tax_code object.
type TaxCode struct {
	Code *string
	Link *string
	Rate any
	UUID *string
}

// TimeZone is the time_zone object.
type TimeZone struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BookingListShape reads a body as BookingList.
func BookingListShape(decoded any) any { return dto.Read[BookingList](decoded) }

// BookingByIdRecordShape reads a body as BookingByIdRecord.
func BookingByIdRecordShape(decoded any) any { return dto.Read[BookingByIdRecord](decoded) }

// BookingNoteList is one page of every order's notes, under data.
type BookingNoteList struct {
	Data       []BookingNote
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints BookingNoteList was measured from.
func (BookingNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/notes",
	}
}

// BookingNote is one entry of a BookingNoteList's data.
type BookingNote struct {
	Attachments   []BookingNoteAttachment
	Content       *string
	CreatedBy     *BookingNoteCreatedBy
	CreatedOn     *string
	Group         any
	Href          *string
	ID            *string
	LastUpdatedBy *BookingNoteLastUpdatedBy
	LastUpdatedOn *string
	Object        *string
	User          any
	Version       *int64
	Watchers      []BookingNoteWatcher
}

// BookingNoteAttachment is one entry of a BookingNote's attachments.
type BookingNoteAttachment struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// BookingNoteCreatedBy is a BookingNote's created_by.
type BookingNoteCreatedBy struct {
	Name   *string
	Object *string
}

// BookingNoteLastUpdatedBy is a BookingNote's last_updated_by.
type BookingNoteLastUpdatedBy struct {
	Name   *string
	Object *string
}

// BookingNoteWatcher is one entry of a BookingNote's watchers.
type BookingNoteWatcher struct {
	Email *string
	Name  *string
	Type  *string
}

// BookingBillingPreferencesView is what GetBookingBillingPreferences answers, under order.
type BookingBillingPreferencesView struct {
	BillingPreferences *BookingBillingPreferencesDetail
	Object             *string
}

// MeasuredFrom lists the endpoints BookingBillingPreferencesView was measured from.
func (BookingBillingPreferencesView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/billing_preferences",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (BookingBillingPreferencesView) EnvelopeKey() string { return OrderKey }

// BookingBillingPreferencesDetail is a BookingBillingPreferencesView's billing_preferences.
type BookingBillingPreferencesDetail struct {
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
	ProformaProperties   *BookingBillingPreferencesDetailProformaProperties
}

// BookingBillingPreferencesDetailProformaProperties is a BookingBillingPreferencesDetail's
// proforma_properties.
type BookingBillingPreferencesDetailProformaProperties struct {
	AcceptanceDate     any
	AcceptanceRequired *bool
	AcceptanceTerm     any
	PaymentRequired    *bool
	TermsAndConditions any
}

// BookingLineRecord is what GetBookingLine answers, under order_line.
type BookingLineRecord struct {
	ChargeItemUUID           *string
	ExpectedDeliveryDate     *string
	ID                       *string
	Item                     *BookingLineItem
	ItemAccountingCode       *BookingLineItemAccountingCode
	ItemChargeType           *string
	ItemCustomAttributes     []any
	ItemDescription          *string
	ItemInvoiceNote          *string
	ItemOrderQuantity        *int64
	ItemPriceSnapshot        *BookingLineItemPriceSnapshot
	ItemPriceTax             *BookingLineItemPriceTax
	ItemProperties           *BookingLineItemProperties
	ItemSaleTaxConfiguration *BookingLineItemSaleTaxConfiguration
	ItemType                 *string
	Object                   *string
	PreOrderStockDetails     *BookingLinePreOrderStockDetails
	ShippingCost             *float64
	Subtotal                 *string
	Tax                      *string
	TaxExemptWhenSold        *bool
	Total                    *string
	Version                  *int64
}

// MeasuredFrom lists the endpoints BookingLineRecord was measured from.
func (BookingLineRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/lines/{line_uuid}",
	}
}

// EnvelopeKey unwraps {"order_line": {...}}.
func (BookingLineRecord) EnvelopeKey() string { return OrderLineKey }

// BookingLineItem is a BookingLineRecord's item.
type BookingLineItem struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
	UUID   *string
}

// BookingLineItemAccountingCode is a BookingLineRecord's item_accounting_code.
type BookingLineItemAccountingCode struct {
	SalesRevenue *string
}

// BookingLineItemPriceSnapshot is a BookingLineRecord's item_price_snapshot.
type BookingLineItemPriceSnapshot struct {
	PricingRule *BookingLineItemPriceSnapshotPricingRule
}

// BookingLineItemPriceSnapshotPricingRule is a BookingLineItemPriceSnapshot's pricing_rule.
type BookingLineItemPriceSnapshotPricingRule struct {
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

// BookingLineItemPriceTax is a BookingLineRecord's item_price_tax.
type BookingLineItemPriceTax struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
	Rate   *float64
}

// BookingLineItemProperties is a BookingLineRecord's item_properties.
type BookingLineItemProperties struct {
	BillingMode                 *string
	ChargingAndBillingAlignment *string
	ChargingStartDate           *string
	ProRataPartialPricingPeriod *string
}

// BookingLineItemSaleTaxConfiguration is a BookingLineRecord's item_sale_tax_configuration.
type BookingLineItemSaleTaxConfiguration struct {
	SalePriceIsBasedOn *string
	TaxCode            *BookingLineItemSaleTaxConfigurationTaxCode
}

// BookingLineItemSaleTaxConfigurationTaxCode is a BookingLineItemSaleTaxConfiguration's tax_code.
type BookingLineItemSaleTaxConfigurationTaxCode struct {
	Code *string
	Link *string
	Rate *string
	UUID *string
}

// BookingLinePreOrderStockDetails is a BookingLineRecord's pre_order_stock_details.
type BookingLinePreOrderStockDetails struct {
	PendingReserve *string
	Reserved       *string
	Sold           *string
}

// BookingCustomObjectList is one page of the custom objects orders use, under data.
type BookingCustomObjectList struct {
	Data       []BookingCustomObjectEntry
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints BookingCustomObjectList was measured from.
func (BookingCustomObjectList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/custom_objects",
		"GET /api/{api_version}/booking/{booking_id}/custom_objects_instances",
	}
}

// BookingCustomObjectEntry is one entry of a BookingCustomObjectList's data.
type BookingCustomObjectEntry struct {
	CreatedBy       any
	CreatedOn       *string
	CustomComponent any
	Description     any
	DisplayName     *string
	Href            *string
	ID              *string
	LastUpdatedBy   any
	LastUpdatedOn   *string
	Name            *string
	Object          *string
	Relation        *string
	UseIn           *string
}

// BookingCustomObjectRecord is one custom object with the booking's instances of it, under
// custom_object.
type BookingCustomObjectRecord struct {
	CreatedBy       any
	CreatedOn       *string
	CustomComponent any
	Description     any
	DisplayName     *string
	Href            *string
	ID              *string
	Instances       []any
	LastUpdatedBy   any
	LastUpdatedOn   *string
	Name            *string
	Object          *string
	Pagination      *Pagination
	Relation        *string
	UseIn           *string
}

// MeasuredFrom lists the endpoints BookingCustomObjectRecord was measured from.
func (BookingCustomObjectRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/custom_objects/{custom_object_id}",
		"GET /api/{api_version}/booking/{booking_id}/custom_objects_instances/{custom_object_id}",
	}
}

// EnvelopeKey unwraps {"custom_object": {...}}.
func (BookingCustomObjectRecord) EnvelopeKey() string { return CustomObjectKey }

// BookingInvoiceList is one page of a booking's invoices, under data.
type BookingInvoiceList struct {
	Data       []BookingInvoice
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints BookingInvoiceList was measured from.
func (BookingInvoiceList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/invoices",
	}
}

// BookingInvoice is one entry of a BookingInvoiceList's data.
type BookingInvoice struct {
	Account                   *BookingInvoiceAccount
	AlternateBillingEndDate   *string
	AlternateBillingStartDate *string
	AlternateDueDate          *string
	AlternateIssueDate        *string
	BillingEndDate            *string
	BillingStartDate          *string
	CreatedBy                 *BookingInvoiceCreatedBy
	CreatedOn                 *string
	Currency                  *BookingInvoiceCurrency
	CustomForm                *BookingInvoiceCustomForm
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
	Order                     *BookingInvoiceOrder
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

// BookingInvoiceAccount is a BookingInvoice's account.
type BookingInvoiceAccount struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BookingInvoiceCreatedBy is a BookingInvoice's created_by.
type BookingInvoiceCreatedBy struct {
	Name   *string
	Object *string
}

// BookingInvoiceCurrency is a BookingInvoice's currency.
type BookingInvoiceCurrency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BookingInvoiceCustomForm is a BookingInvoice's custom_form.
type BookingInvoiceCustomForm struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BookingInvoiceOrder is a BookingInvoice's order.
type BookingInvoiceOrder struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BookingItemFulfillmentRecord is what GetBookingItemFulfillment answers, under item_fulfillment.
type BookingItemFulfillmentRecord struct {
	Account              *BookingItemFulfillmentAccount
	Code                 *string
	CreatedBy            *BookingItemFulfillmentCreatedBy
	CreatedOn            *string
	CustomAttributes     map[string]any
	DeliveryStatus       any
	FulfillmentCompleted *bool
	FulfillmentDate      *string
	FulfillmentNote      any
	FulfillmentStatus    *string
	Fulfillments         []BookingItemFulfillmentFulfillment
	Href                 *string
	ID                   *string
	Invoice              *BookingItemFulfillmentInvoice
	LastUpdatedBy        *BookingItemFulfillmentLastUpdatedBy
	LastUpdatedOn        *string
	Object               *string
	Order                any
	Status               *string
	TrackingNumber       *string
	Version              *int64
}

// MeasuredFrom lists the endpoints BookingItemFulfillmentRecord was measured from.
func (BookingItemFulfillmentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/item_fulfillments/{fulfillment_id}",
	}
}

// EnvelopeKey unwraps {"item_fulfillment": {...}}.
func (BookingItemFulfillmentRecord) EnvelopeKey() string { return ItemFulfillmentKey }

// BookingItemFulfillmentAccount is a BookingItemFulfillmentRecord's account.
type BookingItemFulfillmentAccount struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BookingItemFulfillmentCreatedBy is a BookingItemFulfillmentRecord's created_by.
type BookingItemFulfillmentCreatedBy struct {
	Name   *string
	Object *string
}

// BookingItemFulfillmentFulfillment is one entry of a BookingItemFulfillmentRecord's fulfillments.
type BookingItemFulfillmentFulfillment struct {
	FulfillmentIndex   *int64
	FulfillmentQty     *float64
	FulfillmentSentQty *float64
	ID                 *string
	InventoryTracking  any
	Item               *BookingItemFulfillmentFulfillmentItem
	Object             *string
	PendingReturn      *float64
	Returned           any
	ServingStatus      *string
	Uom                any
	UomPrecision       any
}

// BookingItemFulfillmentFulfillmentItem is a BookingItemFulfillmentFulfillment's item.
type BookingItemFulfillmentFulfillmentItem struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BookingItemFulfillmentInvoice is a BookingItemFulfillmentRecord's invoice.
type BookingItemFulfillmentInvoice struct {
	Href   *string
	ID     *string
	Object *string
}

// BookingItemFulfillmentLastUpdatedBy is a BookingItemFulfillmentRecord's last_updated_by.
type BookingItemFulfillmentLastUpdatedBy struct {
	Name   *string
	Object *string
}

// BookingKpisRecord is what GetBookingKPIs answers, under order_kpis.
type BookingKpisRecord struct {
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

// MeasuredFrom lists the endpoints BookingKpisRecord was measured from.
func (BookingKpisRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/kpis",
	}
}

// EnvelopeKey unwraps {"order_kpis": {...}}.
func (BookingKpisRecord) EnvelopeKey() string { return OrderKPIsKey }

// BookingPaymentMethodsRecord is what GetBookingPaymentMethods answers, under order.
type BookingPaymentMethodsRecord struct {
	PaymentMethods []BookingPaymentMethodsPaymentMethod
}

// MeasuredFrom lists the endpoints BookingPaymentMethodsRecord was measured from.
func (BookingPaymentMethodsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/payment_methods",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (BookingPaymentMethodsRecord) EnvelopeKey() string { return OrderKey }

// BookingPaymentMethodsPaymentMethod is one entry of a BookingPaymentMethodsRecord's
// payment_methods.
type BookingPaymentMethodsPaymentMethod struct {
	CreatedBy          *string
	CreatedOn          *string
	Default            *bool
	ID                 *string
	LastUpdatedBy      *string
	LastUpdatedOn      *string
	LastUsedAt         *string
	Object             *string
	PaymentCount       *int64
	Processor          *BookingPaymentMethodsPaymentMethodProcessor
	ProcessorType      *string
	Reference          *string
	RestrictedToOrders *bool
	SpecifiedOrderIds  *string
	Status             *string
	Version            *int64
}

// BookingPaymentMethodsPaymentMethodProcessor is a BookingPaymentMethodsPaymentMethod's processor.
type BookingPaymentMethodsPaymentMethodProcessor struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BookingPaymentMethodRecord is what GetBookingPaymentMethod answers, under order.
type BookingPaymentMethodRecord struct {
	PaymentMethod *BookingPaymentMethodPaymentMethod
}

// MeasuredFrom lists the endpoints BookingPaymentMethodRecord was measured from.
func (BookingPaymentMethodRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/payment_methods/{payment_method_ref}",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (BookingPaymentMethodRecord) EnvelopeKey() string { return OrderKey }

// BookingPaymentMethodPaymentMethod is a BookingPaymentMethodRecord's payment_method.
type BookingPaymentMethodPaymentMethod struct {
	CreatedBy          *string
	CreatedOn          *string
	Default            *bool
	ID                 *string
	LastUpdatedBy      *string
	LastUpdatedOn      *string
	LastUsedAt         *string
	Object             *string
	PaymentCount       *int64
	Processor          *BookingPaymentMethodPaymentMethodProcessor
	ProcessorType      *string
	Reference          *string
	RestrictedToOrders *bool
	SpecifiedOrderIds  *string
	Status             *string
	Version            *int64
}

// BookingPaymentMethodPaymentMethodProcessor is a BookingPaymentMethodPaymentMethod's processor.
type BookingPaymentMethodPaymentMethodProcessor struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BookingPaymentList is one page of a booking's payments, under data.
type BookingPaymentList struct {
	Data       []BookingPayment
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints BookingPaymentList was measured from.
func (BookingPaymentList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/payments",
	}
}

// BookingPayment is one entry of a BookingPaymentList's data.
type BookingPayment struct {
	Account                *BookingPaymentAccount
	CreatedBy              *BookingPaymentCreatedBy
	CreatedOn              *string
	CreditApplied          []any
	GiftCertificateApplied []any
	ID                     *string
	Invoices               []BookingPaymentInvoice
	LastUpdatedBy          any
	LastUpdatedOn          any
	Object                 *string
	Origin                 any
	PaymentApplied         []BookingPaymentPaymentApplied
	PaymentNote            *string
	ReconcileStatus        *string
	Status                 *string
	Version                *int64
}

// BookingPaymentAccount is a BookingPayment's account.
type BookingPaymentAccount struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BookingPaymentCreatedBy is a BookingPayment's created_by.
type BookingPaymentCreatedBy struct {
	Name   *string
	Object *string
}

// BookingPaymentInvoice is one entry of a BookingPayment's invoices.
type BookingPaymentInvoice struct {
	Applied     *string
	DueDate     *string
	ID          *string
	IssueDate   *string
	Outstanding *string
	Total       *string
}

// BookingPaymentPaymentApplied is one entry of a BookingPayment's payment_applied.
type BookingPaymentPaymentApplied struct {
	Amount          *string
	GatewayResponse any
	Method          *string
	Processor       *string
	Reference       *string
}

// BookingPaymentView is what GetBookingPayment answers, under payment.
type BookingPaymentView struct {
	Account                *BookingPaymentViewAccount
	CreatedBy              *BookingPaymentViewCreatedBy
	CreatedOn              *string
	CreditApplied          []any
	CustomAttributes       []any
	GiftCertificateApplied []any
	ID                     *string
	Invoices               []BookingPaymentViewInvoice
	LastUpdatedBy          any
	LastUpdatedOn          any
	Object                 *string
	Origin                 any
	PaymentApplied         []BookingPaymentViewPaymentApplied
	PaymentNote            *string
	ReconcileStatus        *string
	Status                 *string
	Version                *int64
}

// MeasuredFrom lists the endpoints BookingPaymentView was measured from.
func (BookingPaymentView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/payments/{payment_id}",
	}
}

// EnvelopeKey unwraps {"payment": {...}}.
func (BookingPaymentView) EnvelopeKey() string { return PaymentKey }

// BookingPaymentViewAccount is a BookingPaymentView's account.
type BookingPaymentViewAccount struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// BookingPaymentViewCreatedBy is a BookingPaymentView's created_by.
type BookingPaymentViewCreatedBy struct {
	Name   *string
	Object *string
}

// BookingPaymentViewInvoice is one entry of a BookingPaymentView's invoices.
type BookingPaymentViewInvoice struct {
	Applied     *string
	DueDate     *string
	ID          *string
	IssueDate   *string
	Outstanding *string
	Total       *string
}

// BookingPaymentViewPaymentApplied is one entry of a BookingPaymentView's payment_applied.
type BookingPaymentViewPaymentApplied struct {
	Amount          *string
	GatewayResponse any
	Method          *string
	Processor       *string
	Reference       *string
}

// BookingPdfRecord is what GetBookingPDF answers, under order.
type BookingPdfRecord struct {
	PDF *BookingPdfPdf
}

// MeasuredFrom lists the endpoints BookingPdfRecord was measured from.
func (BookingPdfRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/pdf",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (BookingPdfRecord) EnvelopeKey() string { return OrderKey }

// BookingPdfPdf is a BookingPdfRecord's pdf.
type BookingPdfPdf struct {
	DownloadURL *string
	FileSize    *int64
	ID          *string
	Object      *string
	PreviewURL  *string
}

// BookingQuotePartialInvoicesRecord is what GetBookingQuotePartialInvoices answers, under order.
type BookingQuotePartialInvoicesRecord struct {
	PartialInvoices []any
}

// MeasuredFrom lists the endpoints BookingQuotePartialInvoicesRecord was measured from.
func (BookingQuotePartialInvoicesRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/booking/{booking_id}/quote_partial_invoices",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (BookingQuotePartialInvoicesRecord) EnvelopeKey() string { return OrderKey }

// BookingNoteListShape reads a body as BookingNoteList.
func BookingNoteListShape(decoded any) any { return dto.Read[BookingNoteList](decoded) }

// BookingBillingPreferencesViewShape reads a body as BookingBillingPreferencesView.
func BookingBillingPreferencesViewShape(decoded any) any {
	return dto.Read[BookingBillingPreferencesView](decoded)
}

// BookingLineRecordShape reads a body as BookingLineRecord.
func BookingLineRecordShape(decoded any) any { return dto.Read[BookingLineRecord](decoded) }

// BookingCustomObjectListShape reads a body as BookingCustomObjectList.
func BookingCustomObjectListShape(decoded any) any { return dto.Read[BookingCustomObjectList](decoded) }

// BookingCustomObjectRecordShape reads a body as BookingCustomObjectRecord.
func BookingCustomObjectRecordShape(decoded any) any {
	return dto.Read[BookingCustomObjectRecord](decoded)
}

// BookingInvoiceListShape reads a body as BookingInvoiceList.
func BookingInvoiceListShape(decoded any) any { return dto.Read[BookingInvoiceList](decoded) }

// BookingItemFulfillmentRecordShape reads a body as BookingItemFulfillmentRecord.
func BookingItemFulfillmentRecordShape(decoded any) any {
	return dto.Read[BookingItemFulfillmentRecord](decoded)
}

// BookingKpisRecordShape reads a body as BookingKpisRecord.
func BookingKpisRecordShape(decoded any) any { return dto.Read[BookingKpisRecord](decoded) }

// BookingPaymentMethodsRecordShape reads a body as BookingPaymentMethodsRecord.
func BookingPaymentMethodsRecordShape(decoded any) any {
	return dto.Read[BookingPaymentMethodsRecord](decoded)
}

// BookingPaymentMethodRecordShape reads a body as BookingPaymentMethodRecord.
func BookingPaymentMethodRecordShape(decoded any) any {
	return dto.Read[BookingPaymentMethodRecord](decoded)
}

// BookingPaymentListShape reads a body as BookingPaymentList.
func BookingPaymentListShape(decoded any) any { return dto.Read[BookingPaymentList](decoded) }

// BookingPaymentViewShape reads a body as BookingPaymentView.
func BookingPaymentViewShape(decoded any) any { return dto.Read[BookingPaymentView](decoded) }

// BookingPdfRecordShape reads a body as BookingPdfRecord.
func BookingPdfRecordShape(decoded any) any { return dto.Read[BookingPdfRecord](decoded) }

// BookingQuotePartialInvoicesRecordShape reads a body as BookingQuotePartialInvoicesRecord.
func BookingQuotePartialInvoicesRecordShape(decoded any) any {
	return dto.Read[BookingQuotePartialInvoicesRecord](decoded)
}
