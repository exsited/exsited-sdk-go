package orders

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

// OrderRecord is the order body.
type OrderRecord struct {
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

// MeasuredFrom lists the endpoints OrderRecord was measured from.
func (OrderRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders",
	}
}

// OrderList is a page of OrderRecord records.
type OrderList struct {
	Data       []OrderRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints OrderList was measured from.
func (OrderList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders",
	}
}

// OrderNoteRecord is the order note body.
type OrderNoteRecord struct {
	Attachments   []Attachment
	Content       *string
	CreatedBy     *CreatedBy
	CreatedOn     *string
	Group         any
	Href          *string
	ID            *string
	LastUpdatedBy *LastUpdatedBy
	LastUpdatedOn any
	Object        *string
	Order         *Order
	User          any
	Version       *int64
	Watchers      []Watcher
}

// MeasuredFrom lists the endpoints OrderNoteRecord was measured from.
func (OrderNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/notes",
	}
}

// OrderNoteList is a page of OrderNoteRecord records.
type OrderNoteList struct {
	Data       []OrderNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints OrderNoteList was measured from.
func (OrderNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/notes",
	}
}

// OrderReadRecord is the order read body.
type OrderReadRecord struct {
	Account                  *Account
	ActivityLogs             []ActivityLog
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
	CustomAttributes         []CustomAttribute
	CustomForm               *CustomForm
	CustomerPurchaseOrderID  any
	Description              *string
	DiscountProfile          *DiscountProfile
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
	ShippingAddress          *ShippingAddress
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

// MeasuredFrom lists the endpoints OrderReadRecord was measured from.
func (OrderReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}",
		"POST /api/{api_version}/orders",
		"POST /api/{api_version}/orders/{order_id}/activate_draft",
		"PATCH /api/{api_version}/orders/{order_id}/billing_preferences",
		"PUT /api/{api_version}/orders/{order_id}/billing_preferences",
		"POST /api/{api_version}/orders/{order_id}/cancel",
		"POST /api/{api_version}/orders/{order_id}/change",
		"POST /api/{api_version}/orders/{order_id}/convert_to_order",
		"PATCH /api/{api_version}/orders/{order_id}/information",
		"POST /api/{api_version}/orders/{order_id}/reactivate",
		"POST /api/{api_version}/orders/{order_id}/relinquish",
		"PATCH /api/{api_version}/orders/{order_id}/lines/{line_uuid}",
		"PATCH /api/{api_version}/orders/{order_id}/lines/{line_uuid}/billing_preferences",
		"PUT /api/{api_version}/orders/{order_id}/lines/{line_uuid}/billing_preferences",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (OrderReadRecord) EnvelopeKey() string { return OrderKey }

// OrderInformationRecord is the order information body.
type OrderInformationRecord struct {
	Account                  *Account
	AllowContract            *bool
	BillingAddress           *BillingAddress
	BillingStartDate         *string
	BillingStartDayOfMonth   *string
	CommunicationPreferences []CommunicationPreference
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
	ReferralAccount          any
	SalesRep                 any
	ShippingAddress          *ShippingAddress
	ShippingCost             *float64
	ShippingProfile          any
	Status                   *string
	TimeZone                 *TimeZone
	Version                  *int64
}

// MeasuredFrom lists the endpoints OrderInformationRecord was measured from.
func (OrderInformationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/information",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (OrderInformationRecord) EnvelopeKey() string { return OrderKey }

// AccountOrdersRecord is the account orders body.
type AccountOrdersRecord struct {
	Data       []OrderRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AccountOrdersRecord was measured from.
func (AccountOrdersRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/orders",
		"GET /api/{api_version}/orders/{order_id}/custom_objects",
		"GET /api/{api_version}/orders/{order_id}/custom_objects_instances",
		"GET /api/{api_version}/orders/{order_id}/item_fulfillments",
	}
}

// OrderBillingPreferencesRecord is the order billing preferences body.
type OrderBillingPreferencesRecord struct {
	BillingPreferences *BillingPreferences
	Object             *string
}

// MeasuredFrom lists the endpoints OrderBillingPreferencesRecord was measured from.
func (OrderBillingPreferencesRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/billing_preferences",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (OrderBillingPreferencesRecord) EnvelopeKey() string { return OrderKey }

// OrderDiscardDraftCreateRecord is the body CreateOrderDiscardDraft answers.
type OrderDiscardDraftCreateRecord struct {
	Message *string
}

// MeasuredFrom lists the endpoints OrderDiscardDraftCreateRecord was measured from.
func (OrderDiscardDraftCreateRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/orders/{order_id}/discard_draft",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (OrderDiscardDraftCreateRecord) EnvelopeKey() string { return OrderKey }

// OrderExtendTrialCreateRecord is the body CreateOrderExtendTrial answers.
type OrderExtendTrialCreateRecord struct {
	Message *string
}

// MeasuredFrom lists the endpoints OrderExtendTrialCreateRecord was measured from.
func (OrderExtendTrialCreateRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/orders/{order_id}/extend_trial",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (OrderExtendTrialCreateRecord) EnvelopeKey() string { return OrderKey }

// OrderLineReadRecord is the body GetOrderLine answers.
type OrderLineReadRecord struct {
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

// MeasuredFrom lists the endpoints OrderLineReadRecord was measured from.
func (OrderLineReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/lines/{line_uuid}",
	}
}

// EnvelopeKey unwraps {"order_line": {...}}.
func (OrderLineReadRecord) EnvelopeKey() string { return OrderLineKey }

// OrderLinesReadRecord is the body GetOrderLines answers.
type OrderLinesReadRecord struct {
	Lines             []Line
	Object            *string
	ShippingLineItems []any
	Subtotal          *string
	Tax               *string
	Total             *string
}

// MeasuredFrom lists the endpoints OrderLinesReadRecord was measured from.
func (OrderLinesReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/lines",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (OrderLinesReadRecord) EnvelopeKey() string { return OrderKey }

// OrderNoteFileList is the page GetOrderNoteFiles answers.
type OrderNoteFileList struct {
	Data       []OrderNoteFileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints OrderNoteFileList was measured from.
func (OrderNoteFileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/notes/{order_note_id}/files",
	}
}

// OrderNoteFileReadRecord is the body GetOrderNoteFile answers.
type OrderNoteFileReadRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints OrderNoteFileReadRecord was measured from.
func (OrderNoteFileReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/notes/{order_note_id}/files/{order_file_id}",
	}
}

// EnvelopeKey unwraps {"file": {...}}.
func (OrderNoteFileReadRecord) EnvelopeKey() string { return FileKey }

// OrderNoteReadRecord is the body GetOrderNote answers.
type OrderNoteReadRecord struct {
	Attachments   []Attachment
	Content       *string
	CreatedBy     *CreatedBy
	CreatedOn     *string
	Group         any
	Href          *string
	ID            *string
	LastUpdatedBy *LastUpdatedBy
	LastUpdatedOn *string
	Object        *string
	Order         *Order
	User          any
	Version       *int64
	Watchers      []Watcher
}

// MeasuredFrom lists the endpoints OrderNoteReadRecord was measured from.
func (OrderNoteReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/notes/{order_note_id}",
		"POST /api/{api_version}/orders/{order_id}/notes",
		"PATCH /api/{api_version}/orders/{order_id}/notes/{order_note_id}",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (OrderNoteReadRecord) EnvelopeKey() string { return NoteKey }

// OrderPDFReadRecord is the body GetOrderPDF answers.
type OrderPDFReadRecord struct {
	PDF *PDF
}

// MeasuredFrom lists the endpoints OrderPDFReadRecord was measured from.
func (OrderPDFReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/pdf",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (OrderPDFReadRecord) EnvelopeKey() string { return OrderKey }

// OrderPaymentMethodReadRecord is the body GetOrderPaymentMethod answers.
type OrderPaymentMethodReadRecord struct {
	PaymentMethod *PaymentMethod
}

// MeasuredFrom lists the endpoints OrderPaymentMethodReadRecord was measured from.
func (OrderPaymentMethodReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/payment_methods/{payment_method_ref}",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (OrderPaymentMethodReadRecord) EnvelopeKey() string { return OrderKey }

// OrderPaymentMethodsReadRecord is the body GetOrderPaymentMethods answers.
type OrderPaymentMethodsReadRecord struct {
	PaymentMethods []PaymentMethod
}

// MeasuredFrom lists the endpoints OrderPaymentMethodsReadRecord was measured from.
func (OrderPaymentMethodsReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/payment_methods",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (OrderPaymentMethodsReadRecord) EnvelopeKey() string { return OrderKey }

// OrderNoteFileRecord is one row of OrderNoteFileList.
type OrderNoteFileRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints OrderNoteFileRecord was measured from.
func (OrderNoteFileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/notes/{order_note_id}/files",
	}
}

// ActivityLog is one entry of the activity_logs list.
type ActivityLog struct {
	Activity *string
}

// Attachment is one entry of the attachments list.
type Attachment struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// ContractProperties is the contract_properties object.
type ContractProperties struct {
	AllowDowngrade                       *bool
	AllowDowngradeCharge                 *bool
	AllowEarlyTermination                *bool
	AllowPostponement                    *bool
	AllowTrial                           *bool
	AllowUpgrade                         *bool
	ApplyEarlyTerminationCharge          *bool
	AutoRenewalEndsOn                    *string
	AutoRenewalRequireCustomerAcceptance *bool
	AutoRenewalTerm                      *string
	DowngradeChargeFixed                 *string
	DowngradeChargeType                  *string
	EarlyTerminationChargeType           *string
	EarlyTerminationMinimumPeriod        *string
	InitialContractTerm                  *string
	PeriodBeforeDowngrade                *string
	RenewAutomatically                   *string
	RequireCustomerAcceptance            *bool
	RequiresPaymentMethod                *bool
	StartContractAfterTrialEnds          *bool
	TerminationAccountingCode            *string
	TerminationFixedCharge               *string
	TerminationNoticePeriod              *string
	TrialEndDate                         *string
	TrialPeriod                          *string
	TrialRequiresPaymentMethod           *bool
}

// CreatedBy is the created_by object.
type CreatedBy struct {
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
	SalesRevenue any
}

// ItemCustomAttribute is one entry of the item_custom_attributes list.
type ItemCustomAttribute struct {
	Name *string
	// Value is a string for a STRING attribute and a list of strings for a multi-value one.
	Value any
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
	BillingMode                  *string
	ChargingAndBillingAlignment  *string
	ChargingPeriod               *string
	ChargingStartDate            *string
	ProRataPartialChargingPeriod *string
	ProRataPartialPricingPeriod  *string
}

// ItemSaleTaxConfiguration is the item_sale_tax_configuration object.
type ItemSaleTaxConfiguration struct {
	SalePriceIsBasedOn *string
	TaxCode            *TaxCode
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// Line is one entry of the lines list.
type Line struct {
	Booking                  *Booking
	ChargeItemUUID           *string
	ExpectedDeliveryDate     *string
	ID                       *string
	Item                     *Item
	ItemAccountingCode       *ItemAccountingCode
	ItemChargeType           *string
	ItemCustomAttributes     []ItemCustomAttribute
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

// Order is the order object.
type Order struct {
	Href   *string
	ID     *string
	Object *string
}

// PDF is the pdf object.
type PDF struct {
	DownloadURL *string
	FileSize    *int64
	ID          *string
	Object      *string
	PreviewURL  *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// PaymentMethod is the payment_method or payment_methods object, alone or as a list entry.
type PaymentMethod struct {
	CreatedBy          *string
	CreatedOn          *string
	Default            *bool
	ID                 *string
	LastUpdatedBy      *string
	LastUpdatedOn      *string
	LastUsedAt         any
	Object             *string
	PaymentCount       *int64
	Processor          *Processor
	ProcessorType      *string
	Reference          *string
	RestrictedToOrders *bool
	SpecifiedOrderIds  *string
	Status             *string
	Version            *int64
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

// Processor is the processor object.
type Processor struct {
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

// Watcher is one entry of the watchers list.
type Watcher struct {
	Email *string
	Name  *string
	Type  *string
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

// Booking is the booking object.
type Booking struct {
	BookingEndTime   *string
	BookingStartTime *string
	LabourUUID       *string
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
	// Value is a string for a STRING attribute and a list of strings for a multi-value one.
	Value any
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

// Kpis is the kpis object.
type Kpis struct {
	CreatedOn               any
	DeletedAt               any
	EstimatedTotal          *string
	LastCancelledAt         any
	LastChangedAt           any
	LastInvoiceIssueDate    any
	LastInvoiceTotal        *string
	LastReactivatedAt       any
	LastUpdatedOn           any
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

// Manager is the manager object.
type Manager struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// ProformaProperties is the proforma_properties object.
type ProformaProperties struct {
	AcceptanceDate     any
	AcceptanceRequired *bool
	AcceptanceTerm     *string
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

// TimeZone is the time_zone object.
type TimeZone struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// OrderListShape reads a body as OrderList.
func OrderListShape(decoded any) any { return dto.Read[OrderList](decoded) }

// OrderReadRecordShape reads a body as OrderReadRecord.
func OrderReadRecordShape(decoded any) any { return dto.Read[OrderReadRecord](decoded) }

// AccountOrdersRecordShape reads a body as AccountOrdersRecord.
func AccountOrdersRecordShape(decoded any) any { return dto.Read[AccountOrdersRecord](decoded) }

// OrderBillingPreferencesRecordShape reads a body as OrderBillingPreferencesRecord.
func OrderBillingPreferencesRecordShape(decoded any) any {
	return dto.Read[OrderBillingPreferencesRecord](decoded)
}

// OrderInformationRecordShape reads a body as OrderInformationRecord.
func OrderInformationRecordShape(decoded any) any { return dto.Read[OrderInformationRecord](decoded) }

// OrderNoteListShape reads a body as OrderNoteList.
func OrderNoteListShape(decoded any) any { return dto.Read[OrderNoteList](decoded) }

// OrderDiscardDraftCreateRecordShape reads a body as OrderDiscardDraftCreateRecord.
func OrderDiscardDraftCreateRecordShape(decoded any) any {
	return dto.Read[OrderDiscardDraftCreateRecord](decoded)
}

// OrderExtendTrialCreateRecordShape reads a body as OrderExtendTrialCreateRecord.
func OrderExtendTrialCreateRecordShape(decoded any) any {
	return dto.Read[OrderExtendTrialCreateRecord](decoded)
}

// OrderLineReadRecordShape reads a body as OrderLineReadRecord.
func OrderLineReadRecordShape(decoded any) any { return dto.Read[OrderLineReadRecord](decoded) }

// OrderLinesReadRecordShape reads a body as OrderLinesReadRecord.
func OrderLinesReadRecordShape(decoded any) any { return dto.Read[OrderLinesReadRecord](decoded) }

// OrderNoteFileListShape reads a body as OrderNoteFileList.
func OrderNoteFileListShape(decoded any) any { return dto.Read[OrderNoteFileList](decoded) }

// OrderNoteFileReadRecordShape reads a body as OrderNoteFileReadRecord.
func OrderNoteFileReadRecordShape(decoded any) any { return dto.Read[OrderNoteFileReadRecord](decoded) }

// OrderNoteReadRecordShape reads a body as OrderNoteReadRecord.
func OrderNoteReadRecordShape(decoded any) any { return dto.Read[OrderNoteReadRecord](decoded) }

// OrderPDFReadRecordShape reads a body as OrderPDFReadRecord.
func OrderPDFReadRecordShape(decoded any) any { return dto.Read[OrderPDFReadRecord](decoded) }

// OrderPaymentMethodReadRecordShape reads a body as OrderPaymentMethodReadRecord.
func OrderPaymentMethodReadRecordShape(decoded any) any {
	return dto.Read[OrderPaymentMethodReadRecord](decoded)
}

// OrderPaymentMethodsReadRecordShape reads a body as OrderPaymentMethodsReadRecord.
func OrderPaymentMethodsReadRecordShape(decoded any) any {
	return dto.Read[OrderPaymentMethodsReadRecord](decoded)
}

// OrderCustomObjectRecord is one custom object with the order's instances of it, under
// custom_object.
type OrderCustomObjectRecord struct {
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

// MeasuredFrom lists the endpoints OrderCustomObjectRecord was measured from.
func (OrderCustomObjectRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/custom_objects/{custom_object_id}",
		"GET /api/{api_version}/orders/{order_id}/custom_objects_instances/{custom_object_id}",
	}
}

// EnvelopeKey unwraps {"custom_object": {...}}.
func (OrderCustomObjectRecord) EnvelopeKey() string { return CustomObjectKey }

// OrderItemFulfillmentRecord is what GetOrderItemFulfillment answers, under item_fulfillment.
type OrderItemFulfillmentRecord struct {
	Account              *OrderItemFulfillmentAccount
	Code                 *string
	CreatedBy            *OrderItemFulfillmentCreatedBy
	CreatedOn            *string
	CustomAttributes     map[string]any
	DeliveryStatus       any
	FulfillmentCompleted *bool
	FulfillmentDate      *string
	FulfillmentNote      any
	FulfillmentStatus    *string
	Fulfillments         []OrderItemFulfillmentFulfillment
	Href                 *string
	ID                   *string
	Invoice              *OrderItemFulfillmentInvoice
	LastUpdatedBy        *OrderItemFulfillmentLastUpdatedBy
	LastUpdatedOn        *string
	Object               *string
	Order                any
	Status               *string
	TrackingNumber       *string
	Version              *int64
}

// MeasuredFrom lists the endpoints OrderItemFulfillmentRecord was measured from.
func (OrderItemFulfillmentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/item_fulfillments/{fulfillment_id}",
	}
}

// EnvelopeKey unwraps {"item_fulfillment": {...}}.
func (OrderItemFulfillmentRecord) EnvelopeKey() string { return ItemFulfillmentKey }

// OrderItemFulfillmentAccount is a OrderItemFulfillmentRecord's account.
type OrderItemFulfillmentAccount struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// OrderItemFulfillmentCreatedBy is a OrderItemFulfillmentRecord's created_by.
type OrderItemFulfillmentCreatedBy struct {
	Name   *string
	Object *string
}

// OrderItemFulfillmentFulfillment is one entry of a OrderItemFulfillmentRecord's fulfillments.
type OrderItemFulfillmentFulfillment struct {
	FulfillmentIndex   *int64
	FulfillmentQty     *float64
	FulfillmentSentQty *float64
	ID                 *string
	InventoryTracking  any
	Item               *OrderItemFulfillmentFulfillmentItem
	Object             *string
	PendingReturn      *float64
	Returned           any
	ServingStatus      *string
	Uom                any
	UomPrecision       any
}

// OrderItemFulfillmentFulfillmentItem is a OrderItemFulfillmentFulfillment's item.
type OrderItemFulfillmentFulfillmentItem struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// OrderItemFulfillmentInvoice is a OrderItemFulfillmentRecord's invoice.
type OrderItemFulfillmentInvoice struct {
	Href   *string
	ID     *string
	Object *string
}

// OrderItemFulfillmentLastUpdatedBy is a OrderItemFulfillmentRecord's last_updated_by.
type OrderItemFulfillmentLastUpdatedBy struct {
	Name   *string
	Object *string
}

// OrderKpisRecord is what GetOrderKPIs answers, under order_kpis.
type OrderKpisRecord struct {
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

// MeasuredFrom lists the endpoints OrderKpisRecord was measured from.
func (OrderKpisRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/kpis",
	}
}

// EnvelopeKey unwraps {"order_kpis": {...}}.
func (OrderKpisRecord) EnvelopeKey() string { return OrderKPIsKey }

// OrderPaymentRecord is what GetOrderPayment answers, under payment.
type OrderPaymentRecord struct {
	Account                *OrderPaymentAccount
	CreatedBy              *OrderPaymentCreatedBy
	CreatedOn              *string
	CreditApplied          []any
	CustomAttributes       []any
	GiftCertificateApplied []any
	ID                     *string
	Invoices               []OrderPaymentInvoice
	LastUpdatedBy          any
	LastUpdatedOn          any
	Object                 *string
	Origin                 any
	PaymentApplied         []OrderPaymentPaymentApplied
	PaymentNote            *string
	ReconcileStatus        *string
	Status                 *string
	Version                *int64
}

// MeasuredFrom lists the endpoints OrderPaymentRecord was measured from.
func (OrderPaymentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/payments/{payment_id}",
	}
}

// EnvelopeKey unwraps {"payment": {...}}.
func (OrderPaymentRecord) EnvelopeKey() string { return PaymentKey }

// OrderPaymentAccount is a OrderPaymentRecord's account.
type OrderPaymentAccount struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// OrderPaymentCreatedBy is a OrderPaymentRecord's created_by.
type OrderPaymentCreatedBy struct {
	Name   *string
	Object *string
}

// OrderPaymentInvoice is one entry of a OrderPaymentRecord's invoices.
type OrderPaymentInvoice struct {
	Applied     *string
	DueDate     *string
	ID          *string
	IssueDate   *string
	Outstanding *string
	Total       *string
}

// OrderPaymentPaymentApplied is one entry of a OrderPaymentRecord's payment_applied.
type OrderPaymentPaymentApplied struct {
	Amount          *string
	GatewayResponse any
	Method          *string
	Processor       *string
	Reference       *string
}

// OrderQuotePartialInvoicesRecord is what GetOrderQuotePartialInvoices answers, under order.
type OrderQuotePartialInvoicesRecord struct {
	PartialInvoices []any
}

// MeasuredFrom lists the endpoints OrderQuotePartialInvoicesRecord was measured from.
func (OrderQuotePartialInvoicesRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/orders/{order_id}/quote_partial_invoices",
	}
}

// EnvelopeKey unwraps {"order": {...}}.
func (OrderQuotePartialInvoicesRecord) EnvelopeKey() string { return OrderKey }

// OrderCustomObjectRecordShape reads a body as OrderCustomObjectRecord.
func OrderCustomObjectRecordShape(decoded any) any { return dto.Read[OrderCustomObjectRecord](decoded) }

// OrderItemFulfillmentRecordShape reads a body as OrderItemFulfillmentRecord.
func OrderItemFulfillmentRecordShape(decoded any) any {
	return dto.Read[OrderItemFulfillmentRecord](decoded)
}

// OrderKpisRecordShape reads a body as OrderKpisRecord.
func OrderKpisRecordShape(decoded any) any { return dto.Read[OrderKpisRecord](decoded) }

// OrderPaymentRecordShape reads a body as OrderPaymentRecord.
func OrderPaymentRecordShape(decoded any) any { return dto.Read[OrderPaymentRecord](decoded) }

// OrderQuotePartialInvoicesRecordShape reads a body as OrderQuotePartialInvoicesRecord.
func OrderQuotePartialInvoicesRecordShape(decoded any) any {
	return dto.Read[OrderQuotePartialInvoicesRecord](decoded)
}
