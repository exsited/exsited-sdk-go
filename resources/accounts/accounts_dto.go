package accounts

import (
	"github.com/exsited/exsited-sdk-go/v4/dto"
	"github.com/exsited/exsited-sdk-go/v4/resources/customcomponents"
)

// Three read/write asymmetries are the API's: an address reads postal_code and is
// written post_code; a contact's name reads nested under name and is written flat;
// the billing alignment is written charging_and_billing_alignment as a boolean and
// reads back under that name as a string or as charges_aligned_to_billing, a boolean.

// Reference is v4's standard embedded pointer to another record.
type Reference struct {
	Object *string
	ID     *string
	Name   *string
	Code   *string
	Href   *string
	Rate   *float64
}

// User is whoever created or last updated a record.
type User struct {
	Object *string
	ID     *string
	Name   *string
	Email  *string
	Href   *string
}

type Currency struct {
	Object *string
	ID     *string
	Name   *string
	Code   *string
	Href   *string
}

type TimeZone struct {
	Object *string
	ID     *string
	Name   *string
	Href   *string
}

type CustomForm struct {
	Object *string
	ID     *string
	Name   *string
	Href   *string
}

// SocialLinks is present on every account, never null; each value can be null.
type SocialLinks struct {
	Website  *string
	LinkedIn *string `wire:"linkedin"`
	Twitter  *string
	Facebook *string
}

// Address is an address as it reads back.
type Address struct {
	Object  *string
	ID      *string
	Line1   *string `wire:"address_line_1"`
	Line2   *string `wire:"address_line_2"`
	Line3   *string `wire:"address_line_3"`
	Line4   *string `wire:"address_line_4"`
	Line5   *string `wire:"address_line_5"`
	City    *string
	State   *string
	Country *string
	// PostalCode is the read spelling. Written as post_code; see AddressWrite.
	PostalCode      *string
	DefaultBilling  *bool
	DefaultShipping *bool
	Href            *string
	CountryCode     any
	Line1Field      *string `wire:"line_1"`
	Line2Field      *string `wire:"line_2"`
}

// ContactName is a contact's name as it reads back - nested, and split four ways.
type ContactName struct {
	Salutation  *string
	First       *string
	Middle      *string
	Last        *string
	Designation *string
}

type ContactEmail struct {
	Address       *string
	DoNotEmail    *bool
	EmailOptedOut *bool
}

// ContactPhone is a contact's phone, mobile or fax. All three share this shape.
type ContactPhone struct {
	CountryCode  *string
	AreaCode     *string
	Number       *string
	DoNotCall    *bool
	CallOptedOut *bool
	DialingCode  *string
	Formatted    *string
}

// Contact is a contact as it reads back.
type Contact struct {
	Object *string
	ID     *string
	// Type is "billing", "shipping" or "primary". This is what a v4 path takes.
	Type                       *string
	BillingContact             *bool
	ShippingContact            *bool
	ReceivesBillingInformation *bool
	Name                       *ContactName
	Email                      *ContactEmail
	// Phone can be null.
	Phone *ContactPhone
	// Mobile can be null.
	Mobile *ContactPhone
	// Fax can be null.
	Fax *ContactPhone
	// Address is untyped: an object whose shape is not established as Address.
	Address          *Address
	CustomAttributes []CustomAttribute
	CreatedOn        *string
	// LastUpdatedOn is set only on a contact that has been edited.
	LastUpdatedOn *string
	CreatedBy     *User
	LastUpdatedBy *User
	Version       *int64
}

type PaymentProcessor struct {
	Object *string
	ID     *string
	Name   *string
	Href   *string
}

// PaymentMethod is a payment method as it reads back.
type PaymentMethod struct {
	Object        *string
	ID            *string
	ProcessorType *string
	Status        *string
	// IsDefault carries the wire key "default", which Go cannot use as an exported
	// field name.
	IsDefault          *bool `wire:"default"`
	RestrictedToOrders *bool
	// SpecifiedOrderIDs is untyped: the API sends null, and its element type is unknown.
	SpecifiedOrderIDs any `wire:"specified_order_ids"`
	Processor         *PaymentProcessor
	Reference         *string
	PaymentCount      *int64
	// LastUsedAt is null on a method never charged.
	LastUsedAt    *string
	CreatedOn     *string
	LastUpdatedOn *string
	CreatedBy     *User
	LastUpdatedBy *User
	Version       *int64
}

// BillingPreferences are the preferences as they read back.
//
// Field presence varies a great deal between accounts, so most of these are absent
// rather than null on a given one, and exhttp.ApiResponse.Map is the only way to tell.
type BillingPreferences struct {
	CommunicationProfile   *string
	InvoiceMode            *string
	InvoiceTerm            *string
	BillingPeriod          *string
	BillingStartDate       *string
	BillingStartDayOfMonth *int64
	// ChargesAlignedToBilling is one of the two names the alignment setting reads
	// back under, a boolean.
	ChargesAlignedToBilling *bool
	// ChargingAndBillingAlignment is the other, a string; the write side sends this
	// same name as a boolean.
	ChargingAndBillingAlignment *string
	PaymentProcessor            *string
	PaymentMode                 *string
	PaymentTerm                 *string
	PaymentTermAlignment        *string
	ConsolidateInvoice          *bool
	FulfillmentMode             *string
	FulfillmentTerm             *string
	// ProformaProperties is untyped: an object whose shape is not established.
	ProformaProperties any
	CommunicationGroup *string
	BillingMode        *string
	ChargingPeriod     *string
	ChargingStartDate  *string
	EventProfile       *string
}

// NoteAttachment is identified by uuid, a bare UUID, not prefixed the way a
// note's id is.
type NoteAttachment struct {
	UUID        *string
	Filename    *string
	ContentType *string
	Size        *int64
	Href        *string
}

type Note struct {
	Object *string
	ID     *string
	// Content is HTML, stored and returned as written.
	Content       *string
	Account       *Reference
	Group         any
	User          any
	Attachments   []NoteAttachment
	Watchers      []Watcher
	CreatedOn     *string
	LastUpdatedOn *string
	CreatedBy     *User
	LastUpdatedBy *User
	Version       *int64
	Href          *string
}

// Pagination is the paging block, either shape.
//
// One type covers both, because a single response only ever fills one and the
// other reads nil. HasMore is present only on the cursor shape.
type Pagination struct {
	// Limit is the one field both shapes carry.
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	// PrevCursor is untyped: the API sends null.
	PrevCursor any
	// Links is untyped: an object carrying a relative next URL.
	Links *PaginationLinks

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	// Compare against the string; a truthiness test follows "NULL" as though it
	// were a link.
	NextPage *string
	// PreviousPage is "" or null, never a URL. There is no backward paging on this
	// shape.
	PreviousPage *string
}

// AccountRecord is one account.
type AccountRecord struct {
	Object *string
	// ID is the short code, "W13ZGY", and what every path takes.
	ID          *string
	Name        *string
	Status      *string
	Type        *string
	DisplayName *string
	// EmailAddress can be null.
	EmailAddress *string
	// Description can be null.
	Description *string
	Currency    *Currency
	TimeZone    *TimeZone
	// Tax is null on an account with no tax profile.
	Tax        *Reference
	CustomForm *CustomForm
	// ParentAccountID can be null.
	ParentAccountID *string
	// ImageURL is untyped: the API sends null. Use the image endpoint.
	ImageURL any
	// WorkflowStatus can be null.
	WorkflowStatus           *string
	SocialLinks              *SocialLinks
	CommunicationPreferences []CommunicationPreference
	InvoiceParentAccount     *bool
	PortalAccess             *bool
	// Origin can be null.
	Origin  *string
	Group   *Reference
	Manager *Reference
	// SalesRep is rarely set.
	SalesRep *Reference
	// UserTeam is rarely set.
	UserTeam *Reference
	// AccountingCode is untyped: an object with no "object" key of its own, so it
	// is not established as a Reference.
	AccountingCode   *AccountingCode
	PricingLevel     *Reference
	ReferralTracking *string
	CreatedOn        *string
	// LastUpdatedOn is set only once edited.
	LastUpdatedOn *string
	CreatedBy     *User
	LastUpdatedBy *User
	Version       *int64
	Href          *string

	// Sub-resources: each arrives only on the endpoint that asks for it, so all
	// are absent, not null, on a plain GetAccount.
	Addresses          []Address
	Contacts           []Contact
	PaymentMethods     []PaymentMethod
	BillingPreferences *BillingPreferences
	CustomAttributes   []CustomAttribute
	CustomObjects      []CustomObject
	Inventories        []any
	// PartialInvoices is the key the quote-partial-invoices endpoint answers
	// under, which does not match its path segment.
	PartialInvoices []any
	ActivityLogs    []any
	TaxID           any
	TaxIDNumber     any
}

// EnvelopeKey unwraps {"account": {...}}, so one type reads both a detail body and
// an entry inside a listing.
func (AccountRecord) EnvelopeKey() string { return AccountKey }

// AccountList is a listing of accounts. The records sit under "data".
type AccountList struct {
	Object     *string
	URL        *string
	Data       []AccountRecord
	Pagination *Pagination
}

// AccountImageView is a URI, not image bytes, and null when there is no image.
type AccountImageView struct {
	ImageURI  *string
	AccountID *string
}

func (AccountImageView) EnvelopeKey() string { return AccountKey }

type NoteView struct {
	Object        *string
	ID            *string
	Content       *string
	Account       *Reference
	Attachments   []NoteAttachment
	Watchers      []Watcher
	CreatedOn     *string
	LastUpdatedOn *string
	CreatedBy     *User
	LastUpdatedBy *User
	Version       *int64
	Group         any
	Href          *string
	User          any
}

func (NoteView) EnvelopeKey() string { return NoteKey }

type NoteList struct {
	Object     *string
	URL        *string
	Data       []Note
	Pagination *Pagination
}

// NoteFileList is separate from NoteList because Data holds NoteAttachment, not
// Note.
type NoteFileList struct {
	Object     *string
	URL        *string
	Data       []NoteAttachment
	Pagination *Pagination
}

// AccountCancelBody cancels an account.
//
// EffectiveDate is YYYY-MM-DD. Any other format is a 400 with code "invalid";
// the endpoint does not parse a datetime.
type AccountCancelBody struct {
	EffectiveDate *string
}

func (AccountCancelBody) EnvelopeKey() string { return AccountKey }

// AccountInformationUpdate is flatter than the read: Currency and TimeZone go in
// as plain strings while AccountRecord reads them back as nested objects, so
// posting a read body back unchanged is a 400.
//
// addresses and communication_preference are deliberately absent: the endpoint
// answers 400 not_understood for both.
type AccountInformationUpdate struct {
	Name             *string
	DisplayName      *string
	Description      *string
	EmailAddress     *string
	Origin           *string
	UserTeam         *string
	Group            *string
	Manager          *string
	ReferralTracking *string
	ParentAccount    *string
	// Currency is the code, not the nested object the read returns.
	Currency *string
	// TimeZone is the name, not the nested object the read returns.
	TimeZone            *string
	GrantPortalAccess   *bool
	GrantPortalPassword *string
	Tax                 any
}

func (AccountInformationUpdate) EnvelopeKey() string { return AccountKey }

// AccountSetPasswordBody has no envelope: these two sit at the top level, unlike
// AccountCancelBody's field.
type AccountSetPasswordBody struct {
	Password        *string
	ConfirmPassword *string
}

// AccountImageUpload is base64 in a JSON body, not a multipart upload.
type AccountImageUpload struct {
	ImageBase64 *string `wire:"image_base64"`
}

// AddressWrite is one address as it is written: PostCode where the read has
// postal_code.
type AddressWrite struct {
	Line1 *string `wire:"address_line_1"`
	Line2 *string `wire:"address_line_2"`
	Line3 *string `wire:"address_line_3"`
	Line4 *string `wire:"address_line_4"`
	Line5 *string `wire:"address_line_5"`
	City  *string
	// State is the name, not an id.
	State *string
	// Country is the name, not an id.
	Country *string
	// PostCode is the write spelling. Reads back as postal_code.
	PostCode        *string
	DefaultBilling  *bool
	DefaultShipping *bool
}

// AddressesCreate is a list under "addresses", at the top level with no envelope.
type AddressesCreate struct {
	Addresses []AddressWrite
}

// AddressUpdate is one address under a singular key, not the plural the create
// takes.
type AddressUpdate struct {
	Line1           *string `wire:"address_line_1"`
	Line2           *string `wire:"address_line_2"`
	City            *string
	State           *string
	Country         *string
	PostCode        *string
	DefaultBilling  *bool
	DefaultShipping *bool
}

func (AddressUpdate) EnvelopeKey() string { return AddressKey }

// ContactWrite is flat where the read is nested: FirstName against ContactName's
// First.
type ContactWrite struct {
	Salutation  *string
	Designation *string
	FirstName   *string
	MiddleName  *string
	LastName    *string
	Email       *ContactEmail
	Phone       *ContactPhone
	Mobile      *ContactPhone
	Fax         *ContactPhone
	Line1       *string `wire:"address_line_1"`
	Line2       *string `wire:"address_line_2"`
	City        *string
	State       *string
	Country     *string
	PostCode    *string
}

func (ContactWrite) EnvelopeKey() string { return ContactKey }

type BillingPreferencesWrite struct {
	InvoiceMode   *string
	InvoiceTerm   *string
	BillingPeriod *string
	// BillingStartDate is an enum, not a date: "DAY_OF_MONTH" and the like. A
	// calendar date is a 400 with code matches_invalid.
	BillingStartDate       *string
	BillingStartDayOfMonth *string
	// ChargingAndBillingAlignment is written as a boolean; it reads back as a
	// string under this same name. See BillingPreferences.
	ChargingAndBillingAlignment *bool
	PaymentProcessor            *string
	PaymentTerm                 *string
}

func (BillingPreferencesWrite) EnvelopeKey() string { return BillingPreferencesKey }

type PaymentMethodUpdate struct {
	IsDefault *bool `wire:"default"`
	Reference *string
	Status    *string
}

func (PaymentMethodUpdate) EnvelopeKey() string { return PaymentMethodKey }

// Str, Bool and Int are pointer helpers, since every DTO field is a pointer and
// Go has no literal for the address of one.
func Str(value string) *string { return &value }
func Bool(value bool) *bool    { return &value }
func Int(value int64) *int64   { return &value }

// AccountAddressView is the body GetAccountAddress answers.
type AccountAddressView struct {
	Addresses *Address
	ID        *string
}

// MeasuredFrom lists the endpoints AccountAddressView was measured from.
func (AccountAddressView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/addresses/{address_id}",
	}
}

// EnvelopeKey unwraps {"account": {...}}.
func (AccountAddressView) EnvelopeKey() string { return AccountKey }

// AccountContactView is the body GetAccountContact answers.
type AccountContactView struct {
	Contacts *Contact
	ID       *string
}

// MeasuredFrom lists the endpoints AccountContactView was measured from.
func (AccountContactView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/contacts/{contact_type}",
	}
}

// EnvelopeKey unwraps {"account": {...}}.
func (AccountContactView) EnvelopeKey() string { return AccountKey }

// AccountCreditNoteList is the page GetAccountCreditNotes answers.
type AccountCreditNoteList struct {
	Data       []AccountCreditNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AccountCreditNoteList was measured from.
func (AccountCreditNoteList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/credit_notes",
	}
}

// AccountCustomObjectList is the page GetAccountCustomObjects answers.
type AccountCustomObjectList struct {
	Data       []AccountCustomObjectRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AccountCustomObjectList was measured from.
func (AccountCustomObjectList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/custom_objects",
	}
}

// AccountInventoriesView is the body GetAccountInventories answers.
type AccountInventoriesView struct {
	Inventories []any
}

// MeasuredFrom lists the endpoints AccountInventoriesView was measured from.
func (AccountInventoriesView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/inventories",
	}
}

// EnvelopeKey unwraps {"account": {...}}.
func (AccountInventoriesView) EnvelopeKey() string { return AccountKey }

// AccountInvoiceList is the page GetAccountInvoices answers.
type AccountInvoiceList struct {
	Data       []AccountInvoiceRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AccountInvoiceList was measured from.
func (AccountInvoiceList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/invoices",
	}
}

// AccountNoteFileView is the body GetAccountNoteFile answers.
type AccountNoteFileView struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints AccountNoteFileView was measured from.
func (AccountNoteFileView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/notes/{note_id}/files/{file_id}",
	}
}

// EnvelopeKey unwraps {"file": {...}}.
func (AccountNoteFileView) EnvelopeKey() string { return FileKey }

// AccountPaymentMethodView is the body GetAccountPaymentMethod answers.
type AccountPaymentMethodView struct {
	ID            *string
	PaymentMethod *PaymentMethod
}

// MeasuredFrom lists the endpoints AccountPaymentMethodView was measured from.
func (AccountPaymentMethodView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/payment_methods/{payment_method_id}",
	}
}

// EnvelopeKey unwraps {"account": {...}}.
func (AccountPaymentMethodView) EnvelopeKey() string { return AccountKey }

// AccountQuotePartialInvoicesView is the body GetAccountQuotePartialInvoices answers.
type AccountQuotePartialInvoicesView struct {
	PartialInvoices []any
}

// MeasuredFrom lists the endpoints AccountQuotePartialInvoicesView was measured from.
func (AccountQuotePartialInvoicesView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/quote_partial_invoices",
	}
}

// EnvelopeKey unwraps {"account": {...}}.
func (AccountQuotePartialInvoicesView) EnvelopeKey() string { return AccountKey }

// AccountCreditNoteRecord is one row of AccountCreditNoteList.
type AccountCreditNoteRecord struct {
	Account          *Account
	Code             *string
	CreatedBy        *CreatedBy
	CreatedOn        *string
	Currency         *AccountCreditNoteCurrency
	CustomAttributes []CustomAttribute
	Href             *string
	ID               *string
	InitialAmount    *string
	Invoices         []Invoice
	IssueDate        *string
	LastUpdatedBy    *LastUpdatedBy
	LastUpdatedOn    *string
	Object           *string
	Payment          *Payment
	RefundAmount     *string
	Refundable       *bool
	RemainingAmount  *string
	Status           *string
	Version          *int64
}

// MeasuredFrom lists the endpoints AccountCreditNoteRecord was measured from.
func (AccountCreditNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/credit_notes",
	}
}

// AccountCustomObjectRecord is one row of AccountCustomObjectList.
type AccountCustomObjectRecord struct {
	CreatedBy       any
	CreatedOn       *string
	CustomComponent any
	Description     *string
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

// MeasuredFrom lists the endpoints AccountCustomObjectRecord was measured from.
func (AccountCustomObjectRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/custom_objects",
	}
}

// AccountInvoiceRecord is one row of AccountInvoiceList.
type AccountInvoiceRecord struct {
	Account                   *Account
	AlternateBillingEndDate   *string
	AlternateBillingStartDate *string
	AlternateDueDate          *string
	AlternateIssueDate        *string
	BillingEndDate            *string
	BillingStartDate          *string
	CreatedBy                 *CreatedBy
	CreatedOn                 *string
	Currency                  *AccountInvoiceCurrency
	CustomForm                *AccountInvoiceCustomForm
	CustomerPurchaseOrderID   *string
	DiscountAmount            *float64
	Due                       *string
	DueDate                   *string
	ID                        *string
	InvoiceNote               *string
	IssueDate                 *string
	LastPayment               *LastPayment
	LastUpdatedBy             *LastUpdatedBy
	LastUpdatedOn             *string
	Object                    *string
	Order                     *Order
	Origin                    *string
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

// MeasuredFrom lists the endpoints AccountInvoiceRecord was measured from.
func (AccountInvoiceRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/invoices",
	}
}

// Account is the account object.
type Account struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// AccountCreditNoteCurrency is the currency object.
type AccountCreditNoteCurrency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// AccountInvoiceCurrency is the currency object.
type AccountInvoiceCurrency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// AccountInvoiceCustomForm is the custom_form object.
type AccountInvoiceCustomForm struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// CreatedBy is the created_by object.
type CreatedBy struct {
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

// Invoice is one entry of the invoices list.
type Invoice struct {
	Href   *string
	ID     *string
	Object *string
}

// LastPayment is the last_payment object.
type LastPayment struct {
	Amount    *string
	CreatedOn *string
	ID        *string
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// Order is the order object.
type Order struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Payment is the payment object.
type Payment struct {
	Href   *string
	ID     *string
	Object *string
}

// AccountingCode is the accounting_code object.
type AccountingCode struct {
	AccountReceivable *string
}

// CommunicationPreference is one entry of the communication_preferences list.
type CommunicationPreference struct {
	Channel *string
	Enabled *bool
}

// CustomObject is one entry of the custom_objects list.
type CustomObject struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// Watcher is one entry of the watchers list.
type Watcher struct {
	Email *string
	Name  *string
	Type  *string
}

// AccountListShape reads a body as AccountList.
func AccountListShape(decoded any) any { return dto.Read[AccountList](decoded) }

// AccountRecordShape reads a body as AccountRecord.
func AccountRecordShape(decoded any) any { return dto.Read[AccountRecord](decoded) }

// AccountImageViewShape reads a body as AccountImageView.
func AccountImageViewShape(decoded any) any { return dto.Read[AccountImageView](decoded) }

// NoteViewShape reads a body as NoteView.
func NoteViewShape(decoded any) any { return dto.Read[NoteView](decoded) }

// NoteFileListShape reads a body as NoteFileList.
func NoteFileListShape(decoded any) any { return dto.Read[NoteFileList](decoded) }

// NoteListShape reads a body as NoteList.
func NoteListShape(decoded any) any { return dto.Read[NoteList](decoded) }

// AccountAddressViewShape reads a body as AccountAddressView.
func AccountAddressViewShape(decoded any) any { return dto.Read[AccountAddressView](decoded) }

// AccountContactViewShape reads a body as AccountContactView.
func AccountContactViewShape(decoded any) any { return dto.Read[AccountContactView](decoded) }

// AccountCreditNoteListShape reads a body as AccountCreditNoteList.
func AccountCreditNoteListShape(decoded any) any { return dto.Read[AccountCreditNoteList](decoded) }

// AccountCustomObjectListShape reads a body as AccountCustomObjectList.
func AccountCustomObjectListShape(decoded any) any { return dto.Read[AccountCustomObjectList](decoded) }

// AccountInventoriesViewShape reads a body as AccountInventoriesView.
func AccountInventoriesViewShape(decoded any) any { return dto.Read[AccountInventoriesView](decoded) }

// AccountInvoiceListShape reads a body as AccountInvoiceList.
func AccountInvoiceListShape(decoded any) any { return dto.Read[AccountInvoiceList](decoded) }

// AccountNoteFileViewShape reads a body as AccountNoteFileView.
func AccountNoteFileViewShape(decoded any) any { return dto.Read[AccountNoteFileView](decoded) }

// AccountPaymentMethodViewShape reads a body as AccountPaymentMethodView.
func AccountPaymentMethodViewShape(decoded any) any {
	return dto.Read[AccountPaymentMethodView](decoded)
}

// AccountQuotePartialInvoicesViewShape reads a body as AccountQuotePartialInvoicesView.
func AccountQuotePartialInvoicesViewShape(decoded any) any {
	return dto.Read[AccountQuotePartialInvoicesView](decoded)
}

// AccountKpisRecord is what GetAccountKPIs answers, under account_kpis.
type AccountKpisRecord struct {
	Object                  *string
	TotalRevenue            *string
	MonthlyRecurringRevenue *string
	TotalCollected          *string
	TotalRefunded           *string
	TotalOutstanding        *string
	TotalOverdue            *string
	TotalBalance            *string
	TotalCreditBalance      *string
	TotalOrders             *int64
	ActiveOrders            *int64
	LastReactivatedAt       any
	LastCancelledAt         any
	DeletedAt               any
	ComputedAt              *string
}

// MeasuredFrom lists the endpoints AccountKpisRecord was measured from.
func (AccountKpisRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/kpis",
	}
}

// EnvelopeKey unwraps {"account_kpis": {...}}.
func (AccountKpisRecord) EnvelopeKey() string { return AccountKPIsKey }

// AccountFinancialSummaryRecord is what GetAccountFinancialSummary answers, under account: its name
// and the financial_summary itself.
type AccountFinancialSummaryRecord struct {
	Name             *string
	FinancialSummary *AccountFinancialSummaryDetail
}

// MeasuredFrom lists the endpoints AccountFinancialSummaryRecord was measured from.
func (AccountFinancialSummaryRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/financial_summary",
	}
}

// EnvelopeKey unwraps {"account": {...}}.
func (AccountFinancialSummaryRecord) EnvelopeKey() string { return AccountKey }

// AccountFinancialSummaryDetail is a AccountFinancialSummaryRecord's financial_summary.
type AccountFinancialSummaryDetail struct {
	TotalDue                 *string
	TotalGrand               *string
	TotalPaid                *string
	InvoiceCount             *int64
	TotalDueInvoiceCount     *int64
	TotalOverdueAmount       *string
	TotalOverdueInvoiceCount *int64
	CollectedThisMonth       *string
}

// AccountCustomObjectView is what GetAccountCustomObject answers, under custom_object: the schema,
// with the account's instances of it and their own pagination.
type AccountCustomObjectView struct {
	Instances  []any
	Pagination *Pagination
}

// MeasuredFrom lists the endpoints AccountCustomObjectView was measured from.
func (AccountCustomObjectView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/custom_objects/{custom_object_id}",
	}
}

// EnvelopeKey unwraps {"custom_object": {...}}.
func (AccountCustomObjectView) EnvelopeKey() string { return CustomObjectKey }

// AccountComponentObjectRecord is what GetAccountComponentObject answers, under
// custom_component_object: the fields of one entry of the listing GetAccountComponentObjects
// answers.
type AccountComponentObjectRecord struct {
	Accounts           []any
	AttributeGroups    []any
	Attributes         []any
	Code               *string
	Component          *customcomponents.ComponentObjectsDataDetail
	CreatedBy          any
	CreatedOn          *string
	CustomForm         *customcomponents.ComponentObjectsDataCustomForm
	CustomObjects      []any
	Href               *string
	ID                 *string
	Items              []any
	LastUpdatedBy      any
	LastUpdatedOn      any
	Object             *string
	Orders             []any
	Parents            []customcomponents.ComponentObjectsDataParent
	Quotes             []any
	Restrictions       []any
	Status             *string
	Version            *int64
	WorkflowData       []any
	WorkflowStage      any
	WorkflowStatus     any
	WorkflowStatusUUID any
}

// MeasuredFrom lists the endpoints AccountComponentObjectRecord was measured from.
func (AccountComponentObjectRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/component/{component_id}/{component_object_id}",
	}
}

// EnvelopeKey unwraps {"custom_component_object": {...}}.
func (AccountComponentObjectRecord) EnvelopeKey() string { return CustomComponentObjectKey }

// AccountKpisRecordShape reads a body as AccountKpisRecord.
func AccountKpisRecordShape(decoded any) any { return dto.Read[AccountKpisRecord](decoded) }

// AccountFinancialSummaryRecordShape reads a body as AccountFinancialSummaryRecord.
func AccountFinancialSummaryRecordShape(decoded any) any {
	return dto.Read[AccountFinancialSummaryRecord](decoded)
}

// AccountCustomObjectViewShape reads a body as AccountCustomObjectView.
func AccountCustomObjectViewShape(decoded any) any { return dto.Read[AccountCustomObjectView](decoded) }

// AccountComponentObjectRecordShape reads a body as AccountComponentObjectRecord.
func AccountComponentObjectRecordShape(decoded any) any {
	return dto.Read[AccountComponentObjectRecord](decoded)
}
