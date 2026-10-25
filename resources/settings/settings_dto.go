package settings

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

// SettingsDiscountProfileRecord is the settings discount profile body.
type SettingsDiscountProfileRecord struct {
	AccountingCode             *AccountingCode
	CreatedBy                  any
	CreatedOn                  *string
	Description                *string
	DiscountType               *string
	DisplayName                *string
	EndTime                    *string
	FixedAmount                *float64
	Href                       *string
	ID                         *string
	InvoiceNote                *string
	LastUpdatedBy              any
	LastUpdatedOn              *string
	MaxUsePerAccount           *int64
	MaxUsePerOrder             *int64
	MaxUseTotal                *int64
	Name                       *string
	Object                     *string
	OwnerType                  *string
	PercentageAmount           *float64
	PriceIsInclusiveOfDiscount *bool
	Priority                   *int64
	RedemptionRequired         *bool
	Stackable                  *bool
	StartTime                  *string
	Status                     *string
	TaxCode                    *TaxCode
	Version                    *int64
}

// MeasuredFrom lists the endpoints SettingsDiscountProfileRecord was measured from.
func (SettingsDiscountProfileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/discount_profiles",
	}
}

// SettingsDiscountProfileList is a page of SettingsDiscountProfileRecord records.
type SettingsDiscountProfileList struct {
	Data       []SettingsDiscountProfileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsDiscountProfileList was measured from.
func (SettingsDiscountProfileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/discount_profiles",
	}
}

// SettingsPaymentProcessorRecord is the settings payment processor body.
type SettingsPaymentProcessorRecord struct {
	AccountingCode    *AccountingCode
	ConversionRate    *float64
	CreatedBy         any
	CreatedOn         *string
	Currency          *Currency
	Description       any
	Disconnected      *bool
	DisplayName       *string
	GatewayType       *string
	Href              *string
	ID                *string
	IntegrationState  *string
	IsDefault         *bool
	IsForeignCurrency *bool
	LastUpdatedBy     any
	LastUpdatedOn     *string
	Name              *string
	Object            *string
	Provider          *string
	Status            *string
	Version           *int64
}

// MeasuredFrom lists the endpoints SettingsPaymentProcessorRecord was measured from.
func (SettingsPaymentProcessorRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/payment_processors",
	}
}

// SettingsPaymentProcessorList is a page of SettingsPaymentProcessorRecord records.
type SettingsPaymentProcessorList struct {
	Data       []SettingsPaymentProcessorRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsPaymentProcessorList was measured from.
func (SettingsPaymentProcessorList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/payment_processors",
	}
}

// SettingsShippingProfileRecord is the settings shipping profile body.
type SettingsShippingProfileRecord struct {
	AccountingCode        *AccountingCode
	CreatedBy             any
	CreatedOn             *string
	Description           *string
	DisplayName           *string
	FixedAmount           *float64
	Href                  *string
	ID                    *string
	IntegrationConnection any
	InvoiceNote           *string
	LastUpdatedBy         any
	LastUpdatedOn         *string
	Name                  *string
	Object                *string
	PriceTaxInclusive     *bool
	SalesTaxRate          *float64
	ShippingType          *string
	Status                *string
	TaxExempt             *bool
	TaxRule               *TaxRule
	Version               *int64
}

// MeasuredFrom lists the endpoints SettingsShippingProfileRecord was measured from.
func (SettingsShippingProfileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/shipping_profiles",
	}
}

// SettingsShippingProfileList is a page of SettingsShippingProfileRecord records.
type SettingsShippingProfileList struct {
	Data       []SettingsShippingProfileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsShippingProfileList was measured from.
func (SettingsShippingProfileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/shipping_profiles",
	}
}

// SettingsPricingLevelRecord is the settings pricing level body.
type SettingsPricingLevelRecord struct {
	BuiltInItem         any
	CreatedBy           any
	CreatedOn           *string
	Description         any
	DisplayName         *string
	Href                *string
	ID                  *string
	IsDefault           any
	IsTaxInclusive      any
	LastUpdatedBy       any
	LastUpdatedOn       *string
	Name                *string
	Object              *string
	Ordering            *int64
	Status              *string
	TaxCode             *TaxCode
	TaxExempt           any
	TaxInclusiveBasedOn any
}

// MeasuredFrom lists the endpoints SettingsPricingLevelRecord was measured from.
func (SettingsPricingLevelRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/pricing_levels",
	}
}

// SettingsPricingLevelList is a page of SettingsPricingLevelRecord records.
type SettingsPricingLevelList struct {
	Data       []SettingsPricingLevelRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsPricingLevelList was measured from.
func (SettingsPricingLevelList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/pricing_levels",
	}
}

// SettingsTimeZoneRecord is the settings time zone body.
type SettingsTimeZoneRecord struct {
	Code            *string
	CreatedBy       any
	CreatedOn       *string
	Description     any
	DstEnabled      *bool
	Hour            *int64
	Href            *string
	ID              *string
	LastUpdatedBy   any
	LastUpdatedOn   *string
	Minute          *int64
	Name            *string
	Object          *string
	Offset          *int64
	Positive        *bool
	Status          *string
	TimeZonePattern *string
}

// MeasuredFrom lists the endpoints SettingsTimeZoneRecord was measured from.
func (SettingsTimeZoneRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/time_zones",
	}
}

// SettingsTimeZoneList is a page of SettingsTimeZoneRecord records.
type SettingsTimeZoneList struct {
	Data       []SettingsTimeZoneRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsTimeZoneList was measured from.
func (SettingsTimeZoneList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/time_zones",
	}
}

// SettingsWarehouseRecord is the settings warehouse body.
type SettingsWarehouseRecord struct {
	City          any
	Country       *Country
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayName   *string
	Href          *string
	ID            *string
	IsDefault     *bool
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	PostCode      any
	State         any
	Status        *string
	Version       *int64
}

// MeasuredFrom lists the endpoints SettingsWarehouseRecord was measured from.
func (SettingsWarehouseRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/warehouses",
	}
}

// SettingsWarehouseList is a page of SettingsWarehouseRecord records.
type SettingsWarehouseList struct {
	Data       []SettingsWarehouseRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsWarehouseList was measured from.
func (SettingsWarehouseList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/warehouses",
	}
}

// SettingsCurrencyRecord is the settings currency body.
type SettingsCurrencyRecord struct {
	CashRoundingInterval *float64
	Code                 *string
	ConversionRate       *float64
	Country              *Country
	CreatedBy            any
	CreatedOn            *string
	DecimalPlaces        *int64
	Href                 *string
	ID                   *string
	LastUpdatedBy        any
	LastUpdatedOn        *string
	Name                 *string
	Object               *string
	Rounding             *string
	Status               *string
	Symbol               *string
}

// MeasuredFrom lists the endpoints SettingsCurrencyRecord was measured from.
func (SettingsCurrencyRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/currencies",
	}
}

// SettingsCurrencyList is a page of SettingsCurrencyRecord records.
type SettingsCurrencyList struct {
	Data       []SettingsCurrencyRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsCurrencyList was measured from.
func (SettingsCurrencyList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/currencies",
	}
}

// SettingsItemGroupRecord is the settings item group body.
type SettingsItemGroupRecord struct {
	AccountingCode *AccountingCode
	Code           *string
	CreatedBy      any
	CreatedOn      *string
	Description    *string
	DisplayName    *string
	Href           *string
	ID             *string
	LastUpdatedBy  any
	LastUpdatedOn  *string
	Name           *string
	Object         *string
	Status         *string
	Version        *int64
}

// MeasuredFrom lists the endpoints SettingsItemGroupRecord was measured from.
func (SettingsItemGroupRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/item_groups",
	}
}

// SettingsItemGroupList is a page of SettingsItemGroupRecord records.
type SettingsItemGroupList struct {
	Data       []SettingsItemGroupRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsItemGroupList was measured from.
func (SettingsItemGroupList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/item_groups",
	}
}

// CustomObjectRecord is the custom object body.
type CustomObjectRecord struct {
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

// MeasuredFrom lists the endpoints CustomObjectRecord was measured from.
func (CustomObjectRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/custom_objects",
	}
}

// CustomObjectList is a page of CustomObjectRecord records.
type CustomObjectList struct {
	Data       []CustomObjectRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints CustomObjectList was measured from.
func (CustomObjectList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/custom_objects",
	}
}

// SettingsTaxRecord is the settings tax body.
type SettingsTaxRecord struct {
	Code               *string
	Country            *Country
	CountryTaxByStates []CountryTaxByState
	CreatedBy          any
	CreatedOn          *string
	Href               *string
	ID                 *string
	LastUpdatedBy      any
	LastUpdatedOn      *string
	Name               *string
	Object             *string
	Rate               *float64
	Status             *string
}

// MeasuredFrom lists the endpoints SettingsTaxRecord was measured from.
func (SettingsTaxRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/taxes",
	}
}

// SettingsTaxList is a page of SettingsTaxRecord records.
type SettingsTaxList struct {
	Data       []SettingsTaxRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsTaxList was measured from.
func (SettingsTaxList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/taxes",
	}
}

// SettingsVariationRecord is the settings variation body.
type SettingsVariationRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Status        *string
}

// MeasuredFrom lists the endpoints SettingsVariationRecord was measured from.
func (SettingsVariationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/variations",
	}
}

// SettingsVariationList is a page of SettingsVariationRecord records.
type SettingsVariationList struct {
	Data       []SettingsVariationRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsVariationList was measured from.
func (SettingsVariationList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/variations",
	}
}

// SettingsPaymentTermRecord is the settings payment term body.
type SettingsPaymentTermRecord struct {
	Description any
	DisplayName *string
	Duration    *Duration
	Href        *string
	ID          *string
	Name        *string
	Object      *string
	Status      *string
}

// MeasuredFrom lists the endpoints SettingsPaymentTermRecord was measured from.
func (SettingsPaymentTermRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/payment_terms",
	}
}

// SettingsPaymentTermList is a page of SettingsPaymentTermRecord records.
type SettingsPaymentTermList struct {
	Data       []SettingsPaymentTermRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsPaymentTermList was measured from.
func (SettingsPaymentTermList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/payment_terms",
	}
}

// SettingsItemGroupsRecord is the settings item groups body.
type SettingsItemGroupsRecord struct {
	AccountingCode     *string
	BillingPreferences map[string]any
	Code               *string
	CreatedBy          *string
	CreatedOn          *string
	CustomAttributes   []any
	Description        *string
	DisplayName        *string
	Href               *string
	ID                 *string
	ImageURI           *string
	LastUpdatedBy      *string
	LastUpdatedOn      *string
	Manager            *string
	Name               *string
	Object             *string
	Status             *string
	UUID               *string
	Version            *int64
}

// MeasuredFrom lists the endpoints SettingsItemGroupsRecord was measured from.
func (SettingsItemGroupsRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/settings/item_groups",
		"GET /api/{api_version}/settings/item_groups/{item_group_id}",
		"PATCH /api/{api_version}/settings/item_groups/{item_group_id}",
		"POST /api/{api_version}/settings/item_groups/{item_group_id}/billing_preferences",
	}
}

// EnvelopeKey unwraps {"item_group": {...}}.
func (SettingsItemGroupsRecord) EnvelopeKey() string { return ItemGroupKey }

// SettingsCustomLogicScriptRecord is the settings custom logic script body.
type SettingsCustomLogicScriptRecord struct {
	ActionName    *string
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	ObjectName    *string
	Status        *string
	Type          *string
}

// MeasuredFrom lists the endpoints SettingsCustomLogicScriptRecord was measured from.
func (SettingsCustomLogicScriptRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/custom_logic_scripts/{custom_logic_script_id}",
	}
}

// EnvelopeKey unwraps {"custom_logic_script": {...}}.
func (SettingsCustomLogicScriptRecord) EnvelopeKey() string { return CustomLogicScriptKey }

// SettingsCurrencyReadRecord is the settings currency read body.
type SettingsCurrencyReadRecord struct {
	CashRoundingInterval *float64
	Code                 *string
	ConversionRate       *float64
	Country              *Country
	CreatedBy            any
	CreatedOn            *string
	DecimalPlaces        *int64
	Href                 *string
	ID                   *string
	LastUpdatedBy        any
	LastUpdatedOn        *string
	Name                 *string
	Object               *string
	Rounding             *string
	Status               *string
	Symbol               *string
}

// MeasuredFrom lists the endpoints SettingsCurrencyReadRecord was measured from.
func (SettingsCurrencyReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/currencies/{currency_id}",
	}
}

// EnvelopeKey unwraps {"currency": {...}}.
func (SettingsCurrencyReadRecord) EnvelopeKey() string { return CurrencyKey }

// SettingsCustomObjectCustomizeRecord is the settings custom object customize body.
type SettingsCustomObjectCustomizeRecord struct {
	Configuration *Configuration
	Created       *string
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Relation      *string
	Status        *string
	Updated       *string
	UseIn         *string
	UUID          *string
}

// MeasuredFrom lists the endpoints SettingsCustomObjectCustomizeRecord was measured from.
func (SettingsCustomObjectCustomizeRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/settings/custom_objects/{custom_object_id}/customize",
	}
}

// EnvelopeKey unwraps {"custom_object": {...}}.
func (SettingsCustomObjectCustomizeRecord) EnvelopeKey() string { return CustomObjectKey }

// SettingsCustomObjectsRecord is the settings custom objects body.
type SettingsCustomObjectsRecord struct {
	Created         *string
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
	Status          *string
	Updated         *string
	UseIn           *string
	UUID            *string
}

// MeasuredFrom lists the endpoints SettingsCustomObjectsRecord was measured from.
func (SettingsCustomObjectsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/custom_objects/{custom_object_id}",
		"POST /api/{api_version}/settings/custom_objects",
		"PATCH /api/{api_version}/settings/custom_objects/{custom_object_id}",
	}
}

// EnvelopeKey unwraps {"custom_object": {...}}.
func (SettingsCustomObjectsRecord) EnvelopeKey() string { return CustomObjectKey }

// SettingsVariationsRecord is the settings variations body.
type SettingsVariationsRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Options       []any
	Status        *string
	UUID          *string
}

// MeasuredFrom lists the endpoints SettingsVariationsRecord was measured from.
func (SettingsVariationsRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/settings/variations",
		"GET /api/{api_version}/settings/variations/{variation_id}",
		"PATCH /api/{api_version}/settings/variations/{variation_id}",
	}
}

// EnvelopeKey unwraps {"variation": {...}}.
func (SettingsVariationsRecord) EnvelopeKey() string { return VariationKey }

// SettingsTimeZoneReadRecord is the settings time zone read body.
type SettingsTimeZoneReadRecord struct {
	Code            *string
	CreatedBy       any
	CreatedOn       *string
	Description     any
	DstEnabled      *bool
	Hour            *int64
	Href            *string
	ID              *string
	LastUpdatedBy   any
	LastUpdatedOn   *string
	Minute          *int64
	Name            *string
	Object          *string
	Offset          *int64
	Positive        *bool
	Status          *string
	TimeZonePattern *string
}

// MeasuredFrom lists the endpoints SettingsTimeZoneReadRecord was measured from.
func (SettingsTimeZoneReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/time_zones/{time_zone_id}",
	}
}

// EnvelopeKey unwraps {"time_zone": {...}}.
func (SettingsTimeZoneReadRecord) EnvelopeKey() string { return TimeZoneKey }

// SettingsCustomLogicScriptsRecord is the settings custom logic scripts body.
type SettingsCustomLogicScriptsRecord struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsCustomLogicScriptsRecord was measured from.
func (SettingsCustomLogicScriptsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/custom_logic_scripts",
	}
}

// SettingsRecord is the settings body.
type SettingsRecord struct {
	Object    *string
	Resources []Resource
}

// MeasuredFrom lists the endpoints SettingsRecord was measured from.
func (SettingsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings",
	}
}

// EnvelopeKey unwraps {"settings": {...}}.
func (SettingsRecord) EnvelopeKey() string { return SettingsEnvelopeKey }

// SettingsFieldSettingsItemRecord is the settings field settings item body.
type SettingsFieldSettingsItemRecord struct {
	Data       []Data
	Message    *string
	Object     *string
	Pagination any
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsFieldSettingsItemRecord was measured from.
func (SettingsFieldSettingsItemRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/settings/field_settings/account",
		"PATCH /api/{api_version}/settings/field_settings/item",
		"PATCH /api/{api_version}/settings/general_settings/account",
	}
}

// SettingsGeneralSettingsAccountUpdateRecord is the body UpdateSettingsGeneralSettingsAccount answers.
type SettingsGeneralSettingsAccountUpdateRecord struct {
	Account *Account
}

// MeasuredFrom lists the endpoints SettingsGeneralSettingsAccountUpdateRecord was measured from.
func (SettingsGeneralSettingsAccountUpdateRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/settings/general_settings/account",
	}
}

// EnvelopeKey unwraps {"general_settings": {...}}.
func (SettingsGeneralSettingsAccountUpdateRecord) EnvelopeKey() string { return GeneralSettingsKey }

// SettingsGeneralSettingsItemUpdateRecord is the body UpdateSettingsGeneralSettingsItem answers.
type SettingsGeneralSettingsItemUpdateRecord struct {
	Item    *Item
	Product *Product
}

// MeasuredFrom lists the endpoints SettingsGeneralSettingsItemUpdateRecord was measured from.
func (SettingsGeneralSettingsItemUpdateRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/settings/general_settings/item",
	}
}

// EnvelopeKey unwraps {"general_settings": {...}}.
func (SettingsGeneralSettingsItemUpdateRecord) EnvelopeKey() string { return GeneralSettingsKey }

// SettingsItemGroupReadRecord is the body GetSettingsItemGroup answers.
type SettingsItemGroupReadRecord struct {
	AccountingCode *AccountingCode
	Code           *string
	CreatedBy      any
	CreatedOn      *string
	Description    *string
	DisplayName    *string
	Href           *string
	ID             *string
	LastUpdatedBy  any
	LastUpdatedOn  *string
	Name           *string
	Object         *string
	Status         *string
	Version        *int64
}

// MeasuredFrom lists the endpoints SettingsItemGroupReadRecord was measured from.
func (SettingsItemGroupReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/item_groups/{item_group_id}",
	}
}

// EnvelopeKey unwraps {"item_group": {...}}.
func (SettingsItemGroupReadRecord) EnvelopeKey() string { return ItemGroupKey }

// Account is the account object.
type Account struct {
	AccountGrouping                  *AccountGrouping
	AccountManagement                *AccountManagement
	AdditionalContact1               *AdditionalContact1
	AdditionalContact2               *AdditionalContact2
	AdditionalContact3               *AdditionalContact3
	BillingContact                   *BillingContact
	CreationOfSubsidiaryAccount      *CreationOfSubsidiaryAccount
	EnableWorkflowApprovalForAccount *string
	ReferralTrackingAtAccountLevel   *ReferralTrackingAtAccountLevel
	RequireEmailAddressForAccount    *string
	RequireNameToBeUnique            *string
	SalesRepTrackingAtAccountLevel   *SalesRepTrackingAtAccountLevel
	SendCommunicationAsGroups        *string
	ShippingContact                  *ShippingContact
	UniqueEmailAddressForAccount     *string
	UniqueEmailAddressForContacts    *string
}

// AccountGrouping is the account_grouping object.
type AccountGrouping struct {
	DefaultGroup any
	Enabled      *string
}

// AccountManagement is the account_management object.
type AccountManagement struct {
	DefaultManager any
	Enabled        *string
}

// AccountingCode is the accounting_code object.
type AccountingCode struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// AdditionalContact1 is the additional_contact_1 object.
type AdditionalContact1 struct {
	DisplayName                              *string
	Enabled                                  *string
	RequireEmailAddressForAdditionalContact1 *string
	Type                                     *string
}

// AdditionalContact2 is the additional_contact_2 object.
type AdditionalContact2 struct {
	DisplayName                              *string
	Enabled                                  *string
	RequireEmailAddressForAdditionalContact2 *string
	Type                                     *string
}

// AdditionalContact3 is the additional_contact_3 object.
type AdditionalContact3 struct {
	DisplayName                              *string
	Enabled                                  *string
	RequireEmailAddressForAdditionalContact3 *string
	Type                                     *string
}

// AvailableChargeTypes is the available_charge_types object.
type AvailableChargeTypes struct {
	Metered   *string
	OneOff    *string
	Recurring *string
}

// AvailablePricingMethods is the available_pricing_methods object.
type AvailablePricingMethods struct {
	DiscountBased *string
	MarkupBased   *string
	Standard      *string
}

// BillingContact is the billing_contact object.
type BillingContact struct {
	DisplayName                          *string
	Enabled                              *string
	RequireEmailAddressForBillingContact *string
	Type                                 *string
}

// CreationOfSubsidiaryAccount is the creation_of_subsidiary_account object.
type CreationOfSubsidiaryAccount struct {
	DefaultGroup any
	Enabled      *string
}

// Data is one entry of the data list.
type Data struct {
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayCode   *string
	Editable      any
	Enabled       *bool
	Entity        *string
	FieldName     *string
	FieldType     *string
	GlobalSearch  *bool
	GroupKey      *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Mandatory     *bool
	Note          any
	Object        *string
	SortOrder     *int64
}

// InventoryAdjustmentRules is the inventory_adjustment_rules object.
type InventoryAdjustmentRules struct {
	QuantityAvailableForSale              *string
	QuantityAvailableForSaleDetermination *string
}

// Item is the item object.
type Item struct {
	AvailableChargeTypes                         *AvailableChargeTypes
	AvailablePricingMethods                      *AvailablePricingMethods
	ChargingPeriodCanVaryFromBillingPeriod       *string
	ChargingStartDateCanVaryFromBillingStartDate *string
	EnableBackOrder                              *string
	EnableBatchTracking                          *string
	EnableBundleItem                             *string
	EnableDirectCostItem                         *string
	EnableFamilyCostItem                         *string
	EnableFixedAssetItem                         *string
	EnableGiftCertificateBasedItem               *string
	EnableInventoryCosting                       *string
	EnableInventoryManagement                    *string
	EnablePricingProfile                         *string
	EnablePurchase                               *string
	EnableSales                                  *string
	EnableSerialization                          *string
	EnableSupplierManagement                     *string
	EnableVariation                              *string
	EnableVariationBasedItem                     *string
	EnableWarehouseManagement                    *string
	EnableWorkflowApprovalForItem                *string
	InventoryAdjustmentRules                     *InventoryAdjustmentRules
	ItemGrouping                                 *ItemGrouping
	ItemManagement                               *ItemManagement
	PreOrder                                     *PreOrder
	PricingPeriodCanVaryFromChargingPeriod       *PricingPeriodCanVaryFromChargingPeriod
	PurchasePriceEnteredIsInclusiveOfTax         *PurchasePriceEnteredIsInclusiveOfTax
	RequireNameToBeUnique                        *string
	SalePriceEnteredIsInclusiveOfTax             *SalePriceEnteredIsInclusiveOfTax
	UseOnSalePrice                               *string
	UsePricingLevel                              *string
	UseSchedulingOfFuturePricing                 *string
	UseShippingProfile                           *string
	UseVolumePricing                             *UseVolumePricing
}

// ItemGrouping is the item_grouping object.
type ItemGrouping struct {
	DefaultGroup any
	Enabled      *string
}

// ItemManagement is the item_management object.
type ItemManagement struct {
	DefaultManager any
	Enabled        *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// PreOrder is the pre_order object.
type PreOrder struct {
	Enabled        *string
	NumberOfPeriod *string
	Period         *string
}

// PricingPeriodCanVaryFromChargingPeriod is the pricing_period_can_vary_from_charging_period object.
type PricingPeriodCanVaryFromChargingPeriod struct {
	DefaultPricingPeriod any
	Enabled              *string
}

// Product is the product object.
type Product struct {
	AllowItemSellMultiQuantity    *string
	AllowPlanSellMultiQuantity    *string
	AllowProductSellMultiQuantity *string
	AllowRateSellMultiQuantity    *string
	ItemQuantityUse               *string
	PlanQuantityUse               *string
	ProductQuantityUse            *string
	RateQuantityUse               *string
}

// PurchasePriceEnteredIsInclusiveOfTax is the purchase_price_entered_is_inclusive_of_tax object.
type PurchasePriceEnteredIsInclusiveOfTax struct {
	Enabled *string
	Tax     *Tax
}

// ReferralTrackingAtAccountLevel is the referral_tracking_at_account_level object.
type ReferralTrackingAtAccountLevel struct {
	DefaultReferralTracking any
	Enabled                 *string
}

// SalePriceEnteredIsInclusiveOfTax is the sale_price_entered_is_inclusive_of_tax object.
type SalePriceEnteredIsInclusiveOfTax struct {
	Enabled                          *string
	Tax                              any
	TaxIncludedInSalesPriceIsBasedOn *string
}

// SalesRepTrackingAtAccountLevel is the sales_rep_tracking_at_account_level object.
type SalesRepTrackingAtAccountLevel struct {
	DefaultSalesRepTracking any
	Enabled                 *string
}

// ShippingContact is the shipping_contact object.
type ShippingContact struct {
	DisplayName                           *string
	Enabled                               *string
	RequireEmailAddressForShippingContact *string
	Type                                  *string
}

// Tax is the tax object.
type Tax struct {
	UUID *string
}

// UseVolumePricing is the use_volume_pricing object.
type UseVolumePricing struct {
	AllUnitsFixedPricing      *string
	AllUnitsPerUnitPricing    *string
	Enabled                   *string
	IncrementalFixedPricing   *string
	IncrementalPerUnitPricing *string
}

// Attribute is one entry of the attributes list.
type Attribute struct {
	Name  *string
	Order *int64
}

// Configuration is the configuration object.
type Configuration struct {
	Attributes []Attribute
}

// Country is the country object.
type Country struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// CountryTaxByState is one entry of the country_tax_by_states list.
type CountryTaxByState struct {
	CreatedOn     *string
	ID            *string
	LastUpdatedOn *string
	Object        *string
	Rate          *float64
	State         any
	Status        *string
}

// Currency is the currency object.
type Currency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Duration is the duration object.
type Duration struct {
	Count  *int64
	Period *string
}

// Resource is one entry of the resources list.
type Resource struct {
	Href *string
	Name *string
}

// TaxCode is the tax_code object.
type TaxCode struct {
	Href   *string
	ID     *string
	Object *string
}

// TaxRule is the tax_rule object.
type TaxRule struct {
	Href   *string
	ID     *string
	Object *string
}

// CustomObjectListShape reads a body as CustomObjectList.
func CustomObjectListShape(decoded any) any { return dto.Read[CustomObjectList](decoded) }

// SettingsRecordShape reads a body as SettingsRecord.
func SettingsRecordShape(decoded any) any { return dto.Read[SettingsRecord](decoded) }

// SettingsCustomObjectsRecordShape reads a body as SettingsCustomObjectsRecord.
func SettingsCustomObjectsRecordShape(decoded any) any {
	return dto.Read[SettingsCustomObjectsRecord](decoded)
}

// SettingsCurrencyListShape reads a body as SettingsCurrencyList.
func SettingsCurrencyListShape(decoded any) any { return dto.Read[SettingsCurrencyList](decoded) }

// SettingsCustomLogicScriptsRecordShape reads a body as SettingsCustomLogicScriptsRecord.
func SettingsCustomLogicScriptsRecordShape(decoded any) any {
	return dto.Read[SettingsCustomLogicScriptsRecord](decoded)
}

// SettingsDiscountProfileListShape reads a body as SettingsDiscountProfileList.
func SettingsDiscountProfileListShape(decoded any) any {
	return dto.Read[SettingsDiscountProfileList](decoded)
}

// SettingsItemGroupListShape reads a body as SettingsItemGroupList.
func SettingsItemGroupListShape(decoded any) any { return dto.Read[SettingsItemGroupList](decoded) }

// SettingsItemGroupsRecordShape reads a body as SettingsItemGroupsRecord.
func SettingsItemGroupsRecordShape(decoded any) any {
	return dto.Read[SettingsItemGroupsRecord](decoded)
}

// SettingsPaymentProcessorListShape reads a body as SettingsPaymentProcessorList.
func SettingsPaymentProcessorListShape(decoded any) any {
	return dto.Read[SettingsPaymentProcessorList](decoded)
}

// SettingsPaymentTermListShape reads a body as SettingsPaymentTermList.
func SettingsPaymentTermListShape(decoded any) any { return dto.Read[SettingsPaymentTermList](decoded) }

// SettingsPricingLevelListShape reads a body as SettingsPricingLevelList.
func SettingsPricingLevelListShape(decoded any) any {
	return dto.Read[SettingsPricingLevelList](decoded)
}

// SettingsShippingProfileListShape reads a body as SettingsShippingProfileList.
func SettingsShippingProfileListShape(decoded any) any {
	return dto.Read[SettingsShippingProfileList](decoded)
}

// SettingsTaxListShape reads a body as SettingsTaxList.
func SettingsTaxListShape(decoded any) any { return dto.Read[SettingsTaxList](decoded) }

// SettingsTimeZoneListShape reads a body as SettingsTimeZoneList.
func SettingsTimeZoneListShape(decoded any) any { return dto.Read[SettingsTimeZoneList](decoded) }

// SettingsVariationListShape reads a body as SettingsVariationList.
func SettingsVariationListShape(decoded any) any { return dto.Read[SettingsVariationList](decoded) }

// SettingsVariationsRecordShape reads a body as SettingsVariationsRecord.
func SettingsVariationsRecordShape(decoded any) any {
	return dto.Read[SettingsVariationsRecord](decoded)
}

// SettingsWarehouseListShape reads a body as SettingsWarehouseList.
func SettingsWarehouseListShape(decoded any) any { return dto.Read[SettingsWarehouseList](decoded) }

// SettingsCurrencyReadRecordShape reads a body as SettingsCurrencyReadRecord.
func SettingsCurrencyReadRecordShape(decoded any) any {
	return dto.Read[SettingsCurrencyReadRecord](decoded)
}

// SettingsCustomLogicScriptRecordShape reads a body as SettingsCustomLogicScriptRecord.
func SettingsCustomLogicScriptRecordShape(decoded any) any {
	return dto.Read[SettingsCustomLogicScriptRecord](decoded)
}

// SettingsFieldSettingsItemRecordShape reads a body as SettingsFieldSettingsItemRecord.
func SettingsFieldSettingsItemRecordShape(decoded any) any {
	return dto.Read[SettingsFieldSettingsItemRecord](decoded)
}

// SettingsTimeZoneReadRecordShape reads a body as SettingsTimeZoneReadRecord.
func SettingsTimeZoneReadRecordShape(decoded any) any {
	return dto.Read[SettingsTimeZoneReadRecord](decoded)
}

// SettingsCustomObjectCustomizeRecordShape reads a body as SettingsCustomObjectCustomizeRecord.
func SettingsCustomObjectCustomizeRecordShape(decoded any) any {
	return dto.Read[SettingsCustomObjectCustomizeRecord](decoded)
}

// SettingsGeneralSettingsAccountUpdateRecordShape reads a body as SettingsGeneralSettingsAccountUpdateRecord.
func SettingsGeneralSettingsAccountUpdateRecordShape(decoded any) any {
	return dto.Read[SettingsGeneralSettingsAccountUpdateRecord](decoded)
}

// SettingsGeneralSettingsItemUpdateRecordShape reads a body as SettingsGeneralSettingsItemUpdateRecord.
func SettingsGeneralSettingsItemUpdateRecordShape(decoded any) any {
	return dto.Read[SettingsGeneralSettingsItemUpdateRecord](decoded)
}

// SettingsItemGroupReadRecordShape reads a body as SettingsItemGroupReadRecord.
func SettingsItemGroupReadRecordShape(decoded any) any {
	return dto.Read[SettingsItemGroupReadRecord](decoded)
}

// SettingsCustomObjectDefinitionList is one page of SettingsCustomObjectDefinition records, under
// data; paged by cursor.
type SettingsCustomObjectDefinitionList struct {
	Data       []SettingsCustomObjectDefinition
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsCustomObjectDefinitionList was measured from.
func (SettingsCustomObjectDefinitionList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/custom_objects",
	}
}

// SettingsCustomObjectDefinition is one custom object definition, as /settings/custom_objects
// serves it.
type SettingsCustomObjectDefinition struct {
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Relation      *string
	Status        *string
	UseIn         *string
}

// MeasuredFrom lists the endpoints SettingsCustomObjectDefinition was measured from.
func (SettingsCustomObjectDefinition) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/custom_objects",
		"GET /api/{api_version}/settings/custom_objects/{custom_object_id}",
	}
}

// SettingsCustomObjectDefinitionReadRecord is one custom object definition, under custom_object:
// the record the listing serves.
type SettingsCustomObjectDefinitionReadRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Relation      *string
	Status        *string
	UseIn         *string
}

// MeasuredFrom lists the endpoints SettingsCustomObjectDefinitionReadRecord was measured from.
func (SettingsCustomObjectDefinitionReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/custom_objects/{custom_object_id}",
	}
}

// EnvelopeKey unwraps {"custom_object": {...}}.
func (SettingsCustomObjectDefinitionReadRecord) EnvelopeKey() string { return CustomObjectKey }

// SettingsCustomLogicScriptSavedRecord is what the custom logic script writes answer, under
// custom_logic_script: the read's fields and code, which the server does not store.
type SettingsCustomLogicScriptSavedRecord struct {
	ActionName    *string
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayName   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	ObjectName    *string
	Status        *string
	Type          *string
	Code          any
}

// MeasuredFrom lists the endpoints SettingsCustomLogicScriptSavedRecord was measured from.
func (SettingsCustomLogicScriptSavedRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/settings/custom_logic_scripts",
		"PATCH /api/{api_version}/settings/custom_logic_scripts/{custom_logic_script_id}",
	}
}

// EnvelopeKey unwraps {"custom_logic_script": {...}}.
func (SettingsCustomLogicScriptSavedRecord) EnvelopeKey() string { return CustomLogicScriptKey }

// SettingsDiscountProfileReadRecord is one discount profile, under discount_profile: the record the
// listing serves.
type SettingsDiscountProfileReadRecord struct {
	AccountingCode             *SettingsDiscountProfileAccountingCode
	CreatedBy                  any
	CreatedOn                  *string
	Description                *string
	DiscountType               *string
	DisplayName                *string
	EndTime                    *string
	FixedAmount                *float64
	Href                       *string
	ID                         *string
	InvoiceNote                *string
	LastUpdatedBy              any
	LastUpdatedOn              *string
	MaxUsePerAccount           *int64
	MaxUsePerOrder             *int64
	MaxUseTotal                *int64
	Name                       *string
	Object                     *string
	OwnerType                  *string
	PercentageAmount           *float64
	PriceIsInclusiveOfDiscount *bool
	Priority                   *int64
	RedemptionRequired         *bool
	Stackable                  *bool
	StartTime                  *string
	Status                     *string
	TaxCode                    *SettingsDiscountProfileTaxCode
	Version                    *int64
}

// MeasuredFrom lists the endpoints SettingsDiscountProfileReadRecord was measured from.
func (SettingsDiscountProfileReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/discount_profiles/{discount_profile_id}",
		"POST /api/{api_version}/settings/discount_profiles",
		"PATCH /api/{api_version}/settings/discount_profiles/{discount_profile_id}",
	}
}

// EnvelopeKey unwraps {"discount_profile": {...}}.
func (SettingsDiscountProfileReadRecord) EnvelopeKey() string { return DiscountProfileKey }

// SettingsDiscountProfileAccountingCode is a SettingsDiscountProfile's accounting_code.
type SettingsDiscountProfileAccountingCode struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// SettingsDiscountProfileTaxCode is a SettingsDiscountProfile's tax_code.
type SettingsDiscountProfileTaxCode struct {
	Href   *string
	ID     *string
	Object *string
}

// SettingsFieldSettingList is one page of SettingsFieldSettingsItemData records, under data; paged
// by cursor.
type SettingsFieldSettingList struct {
	Data       []SettingsFieldSettingsItemData
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsFieldSettingList was measured from.
func (SettingsFieldSettingList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/field_settings",
	}
}

// SettingsFieldSettingsItemData is one entry of a SettingsFieldSettingsItem's data.
type SettingsFieldSettingsItemData struct {
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayCode   *string
	Editable      any
	Enabled       *bool
	Entity        *string
	FieldName     *string
	FieldType     *string
	GlobalSearch  *bool
	GroupKey      *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Mandatory     *bool
	Note          any
	Object        *string
	SortOrder     *int64
}

// SettingsFieldSettingReadRecord is one field setting, under field_setting: the record the listing
// serves.
type SettingsFieldSettingReadRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayCode   *string
	Editable      any
	Enabled       *bool
	Entity        *string
	FieldName     *string
	FieldType     *string
	GlobalSearch  *bool
	GroupKey      *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Mandatory     *bool
	Note          any
	Object        *string
	SortOrder     *int64
}

// MeasuredFrom lists the endpoints SettingsFieldSettingReadRecord was measured from.
func (SettingsFieldSettingReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/field_settings/{field_setting_id}",
	}
}

// EnvelopeKey unwraps {"field_setting": {...}}.
func (SettingsFieldSettingReadRecord) EnvelopeKey() string { return FieldSettingKey }

// SettingsGeneralSettingsRecord is one group of general settings, under general_settings, keyed by
// the group; item and order also carry product.
type SettingsGeneralSettingsRecord struct {
	Account     *SettingsFieldSettingsItemGeneralSettingsAccount
	Dashboard   *SettingsGeneralSettingsDashboard
	DateAndTime *SettingsGeneralSettingsDateAndTime
	Display     map[string]any
	General     map[string]any
	Item        *SettingsGeneralSettingsItemUpdatedDetail
	Labour      *SettingsGeneralSettingsLabour
	Order       *SettingsGeneralSettingsOrder
	Portal      *SettingsGeneralSettingsPortal
	Product     *SettingsGeneralSettingsItemUpdatedProduct
	Purchase    *SettingsGeneralSettingsPurchase
	Security    *SettingsGeneralSettingsSecurity
	Transaction *SettingsGeneralSettingsTransaction
}

// MeasuredFrom lists the endpoints SettingsGeneralSettingsRecord was measured from.
func (SettingsGeneralSettingsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/general_settings/{group}",
	}
}

// EnvelopeKey unwraps {"general_settings": {...}}.
func (SettingsGeneralSettingsRecord) EnvelopeKey() string { return GeneralSettingsKey }

// SettingsFieldSettingsItemGeneralSettingsAccount is a SettingsFieldSettingsItemGeneralSettings's
// account.
type SettingsFieldSettingsItemGeneralSettingsAccount struct {
	AccountGrouping                  *SettingsFieldSettingsItemGeneralSettingsAccountAccountGrouping
	AccountManagement                *SettingsFieldSettingsItemGeneralSettingsAccountAccountManagement
	AdditionalContact1               *SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact1
	AdditionalContact2               *SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact2
	AdditionalContact3               *SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact3
	BillingContact                   *SettingsFieldSettingsItemGeneralSettingsAccountBillingContact
	CreationOfSubsidiaryAccount      *SettingsFieldSettingsItemGeneralSettingsAccountCreationOfSubsidiaryAccount
	EnableWorkflowApprovalForAccount *string
	ReferralTrackingAtAccountLevel   *SettingsFieldSettingsItemGeneralSettingsAccountReferralTrackingAtAccountLevel
	RequireEmailAddressForAccount    *string
	RequireNameToBeUnique            *string
	SalesRepTrackingAtAccountLevel   *SettingsFieldSettingsItemGeneralSettingsAccountSalesRepTrackingAtAccountLevel
	SendCommunicationAsGroups        *string
	ShippingContact                  *SettingsFieldSettingsItemGeneralSettingsAccountShippingContact
	UniqueEmailAddressForAccount     *string
	UniqueEmailAddressForContacts    *string
}

// SettingsFieldSettingsItemGeneralSettingsAccountAccountGrouping is A
// SettingsFieldSettingsItemGeneralSettingsAccount's account_grouping.
type SettingsFieldSettingsItemGeneralSettingsAccountAccountGrouping struct {
	DefaultGroup any
	Enabled      *string
}

// SettingsFieldSettingsItemGeneralSettingsAccountAccountManagement is A
// SettingsFieldSettingsItemGeneralSettingsAccount's account_management.
type SettingsFieldSettingsItemGeneralSettingsAccountAccountManagement struct {
	DefaultManager any
	Enabled        *string
}

// SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact1 is A
// SettingsFieldSettingsItemGeneralSettingsAccount's additional_contact_1.
type SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact1 struct {
	DisplayName                              *string
	Enabled                                  *string
	RequireEmailAddressForAdditionalContact1 *string
	Type                                     *string
}

// SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact2 is A
// SettingsFieldSettingsItemGeneralSettingsAccount's additional_contact_2.
type SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact2 struct {
	DisplayName                              *string
	Enabled                                  *string
	RequireEmailAddressForAdditionalContact2 *string
	Type                                     *string
}

// SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact3 is A
// SettingsFieldSettingsItemGeneralSettingsAccount's additional_contact_3.
type SettingsFieldSettingsItemGeneralSettingsAccountAdditionalContact3 struct {
	DisplayName                              *string
	Enabled                                  *string
	RequireEmailAddressForAdditionalContact3 *string
	Type                                     *string
}

// SettingsFieldSettingsItemGeneralSettingsAccountBillingContact is A
// SettingsFieldSettingsItemGeneralSettingsAccount's billing_contact.
type SettingsFieldSettingsItemGeneralSettingsAccountBillingContact struct {
	DisplayName                          *string
	Enabled                              *string
	RequireEmailAddressForBillingContact *string
	Type                                 *string
}

// SettingsFieldSettingsItemGeneralSettingsAccountCreationOfSubsidiaryAccount is A
// SettingsFieldSettingsItemGeneralSettingsAccount's creation_of_subsidiary_account.
type SettingsFieldSettingsItemGeneralSettingsAccountCreationOfSubsidiaryAccount struct {
	DefaultGroup any
	Enabled      *string
}

// SettingsFieldSettingsItemGeneralSettingsAccountReferralTrackingAtAccountLevel is A
// SettingsFieldSettingsItemGeneralSettingsAccount's referral_tracking_at_account_level.
type SettingsFieldSettingsItemGeneralSettingsAccountReferralTrackingAtAccountLevel struct {
	DefaultReferralTracking any
	Enabled                 *string
}

// SettingsFieldSettingsItemGeneralSettingsAccountSalesRepTrackingAtAccountLevel is A
// SettingsFieldSettingsItemGeneralSettingsAccount's sales_rep_tracking_at_account_level.
type SettingsFieldSettingsItemGeneralSettingsAccountSalesRepTrackingAtAccountLevel struct {
	DefaultSalesRepTracking any
	Enabled                 *string
}

// SettingsFieldSettingsItemGeneralSettingsAccountShippingContact is A
// SettingsFieldSettingsItemGeneralSettingsAccount's shipping_contact.
type SettingsFieldSettingsItemGeneralSettingsAccountShippingContact struct {
	DisplayName                           *string
	Enabled                               *string
	RequireEmailAddressForShippingContact *string
	Type                                  *string
}

// SettingsGeneralSettingsDashboard is a SettingsGeneralSettingsRecord's dashboard.
type SettingsGeneralSettingsDashboard struct {
	Accounting       *string
	DefaultDashboard any `wire:"defaultDashboard"`
	Executive        *string
	Finance          *string
	Production       *string
	Purchase         *string
	Sales            *string
}

// SettingsGeneralSettingsDateAndTime is a SettingsGeneralSettingsRecord's date_and_time.
type SettingsGeneralSettingsDateAndTime struct {
	DateFormat     *string
	FirstDayOfWeek any
	TimeFormat     *string
	TimeZone       *string
}

// SettingsGeneralSettingsItemUpdatedDetail is a SettingsGeneralSettingsItemUpdated's item.
type SettingsGeneralSettingsItemUpdatedDetail struct {
	AvailableChargeTypes                         *SettingsGeneralSettingsItemUpdatedDetailAvailableChargeTypes
	AvailablePricingMethods                      *SettingsGeneralSettingsItemUpdatedDetailAvailablePricingMethods
	ChargingPeriodCanVaryFromBillingPeriod       *string
	ChargingStartDateCanVaryFromBillingStartDate *string
	EnableBackOrder                              *string
	EnableBatchTracking                          *string
	EnableBundleItem                             *string
	EnableDirectCostItem                         *string
	EnableFamilyCostItem                         *string
	EnableFixedAssetItem                         *string
	EnableGiftCertificateBasedItem               *string
	EnableInventoryCosting                       *string
	EnableInventoryManagement                    *string
	EnablePricingProfile                         *string
	EnablePurchase                               *string
	EnableSales                                  *string
	EnableSerialization                          *string
	EnableSupplierManagement                     *string
	EnableVariation                              *string
	EnableVariationBasedItem                     *string
	EnableWarehouseManagement                    *string
	EnableWorkflowApprovalForItem                *string
	InventoryAdjustmentRules                     *SettingsGeneralSettingsItemUpdatedDetailInventoryAdjustmentRules
	ItemGrouping                                 *SettingsGeneralSettingsItemUpdatedDetailItemGrouping
	ItemManagement                               *SettingsGeneralSettingsItemUpdatedDetailItemManagement
	PreOrder                                     *SettingsGeneralSettingsItemUpdatedDetailPreOrder
	PricingPeriodCanVaryFromChargingPeriod       *SettingsGeneralSettingsItemUpdatedDetailPricingPeriodCanVaryFromChargingPeriod
	PurchasePriceEnteredIsInclusiveOfTax         *SettingsGeneralSettingsItemUpdatedDetailPurchasePriceEnteredIsInclusiveOfTax
	RequireNameToBeUnique                        *string
	SalePriceEnteredIsInclusiveOfTax             *SettingsGeneralSettingsItemUpdatedDetailSalePriceEnteredIsInclusiveOfTax
	UseOnSalePrice                               *string
	UsePricingLevel                              *string
	UseSchedulingOfFuturePricing                 *string
	UseShippingProfile                           *string
	UseVolumePricing                             *SettingsGeneralSettingsItemUpdatedDetailUseVolumePricing
}

// SettingsGeneralSettingsItemUpdatedDetailAvailableChargeTypes is A
// SettingsGeneralSettingsItemUpdatedDetail's available_charge_types.
type SettingsGeneralSettingsItemUpdatedDetailAvailableChargeTypes struct {
	Metered   *string
	OneOff    *string
	Recurring *string
}

// SettingsGeneralSettingsItemUpdatedDetailAvailablePricingMethods is A
// SettingsGeneralSettingsItemUpdatedDetail's available_pricing_methods.
type SettingsGeneralSettingsItemUpdatedDetailAvailablePricingMethods struct {
	DiscountBased *string
	MarkupBased   *string
	Standard      *string
}

// SettingsGeneralSettingsItemUpdatedDetailInventoryAdjustmentRules is A
// SettingsGeneralSettingsItemUpdatedDetail's inventory_adjustment_rules.
type SettingsGeneralSettingsItemUpdatedDetailInventoryAdjustmentRules struct {
	QuantityAvailableForSale              *string
	QuantityAvailableForSaleDetermination *string
}

// SettingsGeneralSettingsItemUpdatedDetailItemGrouping is A
// SettingsGeneralSettingsItemUpdatedDetail's item_grouping.
type SettingsGeneralSettingsItemUpdatedDetailItemGrouping struct {
	DefaultGroup any
	Enabled      *string
}

// SettingsGeneralSettingsItemUpdatedDetailItemManagement is A
// SettingsGeneralSettingsItemUpdatedDetail's item_management.
type SettingsGeneralSettingsItemUpdatedDetailItemManagement struct {
	DefaultManager any
	Enabled        *string
}

// SettingsGeneralSettingsItemUpdatedDetailPreOrder is a SettingsGeneralSettingsItemUpdatedDetail's
// pre_order.
type SettingsGeneralSettingsItemUpdatedDetailPreOrder struct {
	Enabled        *string
	NumberOfPeriod *string
	Period         *string
}

// SettingsGeneralSettingsItemUpdatedDetailPricingPeriodCanVaryFromChargingPeriod is A
// SettingsGeneralSettingsItemUpdatedDetail's pricing_period_can_vary_from_charging_period.
type SettingsGeneralSettingsItemUpdatedDetailPricingPeriodCanVaryFromChargingPeriod struct {
	DefaultPricingPeriod any
	Enabled              *string
}

// SettingsGeneralSettingsItemUpdatedDetailPurchasePriceEnteredIsInclusiveOfTax is A
// SettingsGeneralSettingsItemUpdatedDetail's purchase_price_entered_is_inclusive_of_tax.
type SettingsGeneralSettingsItemUpdatedDetailPurchasePriceEnteredIsInclusiveOfTax struct {
	Enabled *string
	Tax     *SettingsGeneralSettingsItemUpdatedDetailPurchasePriceEnteredIsInclusiveOfTaxDetail
}

// SettingsGeneralSettingsItemUpdatedDetailPurchasePriceEnteredIsInclusiveOfTaxDetail is A
// SettingsGeneralSettingsItemUpdatedDetailPurchasePriceEnteredIsInclusiveOfTax's tax.
type SettingsGeneralSettingsItemUpdatedDetailPurchasePriceEnteredIsInclusiveOfTaxDetail struct {
	UUID *string
}

// SettingsGeneralSettingsItemUpdatedDetailSalePriceEnteredIsInclusiveOfTax is A
// SettingsGeneralSettingsItemUpdatedDetail's sale_price_entered_is_inclusive_of_tax.
type SettingsGeneralSettingsItemUpdatedDetailSalePriceEnteredIsInclusiveOfTax struct {
	Enabled                          *string
	Tax                              any
	TaxIncludedInSalesPriceIsBasedOn *string
}

// SettingsGeneralSettingsItemUpdatedDetailUseVolumePricing is A
// SettingsGeneralSettingsItemUpdatedDetail's use_volume_pricing.
type SettingsGeneralSettingsItemUpdatedDetailUseVolumePricing struct {
	AllUnitsFixedPricing      *string
	AllUnitsPerUnitPricing    *string
	Enabled                   *string
	IncrementalFixedPricing   *string
	IncrementalPerUnitPricing *string
}

// SettingsGeneralSettingsLabour is a SettingsGeneralSettingsRecord's labour.
type SettingsGeneralSettingsLabour struct {
	RequireEmailAddressForLabour *string
	UniqueEmailAddress           *string
	UniqueLabourName             *string
}

// SettingsGeneralSettingsOrder is a SettingsGeneralSettingsRecord's order.
type SettingsGeneralSettingsOrder struct {
	AllowNameToBeChangedInOrder           *string
	AllowUnitPriceToBeChangedInOrder      *string
	DefaultFulfillmentStatus              *string
	DefaultShippingProfile                any
	DisplayedPriceIsInclusiveOfTax        *string
	EnableBooking                         *string
	EnableLastSalePriceLookUpForAnAccount *string
	EnableOrderBackdating                 *string
	EnableQuote                           *string
	OrderManagement                       *SettingsGeneralSettingsOrderOrderManagement
	ReferralTrackingAtOrderLevel          *SettingsGeneralSettingsOrderReferralTrackingAtOrderLevel
	SalesRepTrackingAtOrderLevel          *SettingsGeneralSettingsOrderSalesRepTrackingAtOrderLevel
	UseGlobalWarehouse                    *string
}

// SettingsGeneralSettingsOrderOrderManagement is a SettingsGeneralSettingsOrder's order_management.
type SettingsGeneralSettingsOrderOrderManagement struct {
	DefaultManager any
	Enabled        *string
}

// SettingsGeneralSettingsOrderReferralTrackingAtOrderLevel is a SettingsGeneralSettingsOrder's
// referral_tracking_at_order_level.
type SettingsGeneralSettingsOrderReferralTrackingAtOrderLevel struct {
	DefaultReferralTracking any
	Enabled                 *string
}

// SettingsGeneralSettingsOrderSalesRepTrackingAtOrderLevel is a SettingsGeneralSettingsOrder's
// sales_rep_tracking_at_order_level.
type SettingsGeneralSettingsOrderSalesRepTrackingAtOrderLevel struct {
	DefaultSalesRepTracking any
	Enabled                 *string
}

// SettingsGeneralSettingsPortal is a SettingsGeneralSettingsRecord's portal.
type SettingsGeneralSettingsPortal struct {
	Contacts                              map[string]any
	EnableCustomerPortal                  *string `wire:"enableCustomerPortal"`
	EnableInviteToActivePortalForContact  *string `wire:"enableInviteToActivePortalForContact"`
	EnableSpecifyPortalPasswordForContact *string `wire:"enableSpecifyPortalPasswordForContact"`
	Settings                              map[string]any
}

// SettingsGeneralSettingsItemUpdatedProduct is a SettingsGeneralSettingsItemUpdated's product.
type SettingsGeneralSettingsItemUpdatedProduct struct {
	AllowItemSellMultiQuantity    *string
	AllowPlanSellMultiQuantity    *string
	AllowProductSellMultiQuantity *string
	AllowRateSellMultiQuantity    *string
	ItemQuantityUse               *string
	PlanQuantityUse               *string
	ProductQuantityUse            *string
	RateQuantityUse               *string
}

// SettingsGeneralSettingsPurchase is a SettingsGeneralSettingsRecord's purchase.
type SettingsGeneralSettingsPurchase struct {
	AllowModificationOfFollowingPurchaseCreditNoteFields *SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseCreditNoteFields
	AllowModificationOfFollowingPurchaseInvoiceFields    *SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseInvoiceFields
	AllowModificationOfFollowingPurchasePaymentFields    *SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchasePaymentFields
	AllowModificationOfFollowingPurchaseRefundFields     *SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseRefundFields
	EnableLastPurchasePriceLookUp                        *string
	EnablePurchaseQuote                                  *string
	EnableWorkflowApprovalForPurchaseCreditNote          *string
	EnableWorkflowApprovalForPurchaseInvoice             *string
	EnableWorkflowApprovalForPurchaseOrder               *string
	EnableWorkflowApprovalForPurchasePayment             *string
	EnableWorkflowApprovalForPurchaseRefund              *string
	InventoryIsAdjustedWhen                              *string
	SupplierInvoiceIDShouldBeUnique                      *string
	UseGlobalWarehouse                                   *string
}

// SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseCreditNoteFields is A
// SettingsGeneralSettingsPurchase's allow_modification_of_following_purchase_credit_note_fields.
type SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseCreditNoteFields struct {
	Refundable *string
}

// SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseInvoiceFields is A
// SettingsGeneralSettingsPurchase's allow_modification_of_following_purchase_invoice_fields.
type SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseInvoiceFields struct {
	DueDate   *string
	IssueDate *string
}

// SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchasePaymentFields is A
// SettingsGeneralSettingsPurchase's allow_modification_of_following_purchase_payment_fields.
type SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchasePaymentFields struct {
	PaymentDate      *string
	PaymentProcessor *string
	PaymentReference *string
}

// SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseRefundFields is A
// SettingsGeneralSettingsPurchase's allow_modification_of_following_purchase_refund_fields.
type SettingsGeneralSettingsPurchaseAllowModificationOfFollowingPurchaseRefundFields struct {
	RefundDate      *string
	RefundReference *string
}

// SettingsGeneralSettingsSecurity is a SettingsGeneralSettingsRecord's security.
type SettingsGeneralSettingsSecurity struct {
	EnableForceSessionExpire      *string
	EnableTwoFactorAuthentication *string
	ImplementerLogin              *string `wire:"implementerLogin"`
	SessionExpireLimitHour        *string
}

// SettingsGeneralSettingsTransaction is a SettingsGeneralSettingsRecord's transaction.
type SettingsGeneralSettingsTransaction struct {
	AccountingMethod                                                          any
	AllowModificationOfFollowingCreditNoteFields                              *SettingsGeneralSettingsTransactionAllowModificationOfFollowingCreditNoteFields
	AllowModificationOfFollowingInvoiceFields                                 *SettingsGeneralSettingsTransactionAllowModificationOfFollowingInvoiceFields
	AllowModificationOfFollowingPaymentFields                                 *SettingsGeneralSettingsTransactionAllowModificationOfFollowingPaymentFields
	AllowModificationOfFollowingRefundFields                                  *SettingsGeneralSettingsTransactionAllowModificationOfFollowingRefundFields
	AvailableMediumsForCommunicationPreference                                *SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreference
	CreateNegativeChargeToAccommodateAdjustmentRequiredForOrderCancelOrChange *string
	DefaultAccountingCodes                                                    *SettingsGeneralSettingsTransactionDefaultAccountingCodes
	DefaultShippingProfile                                                    any
	DisableNotificationForInvoiceWithChargesOfZeroAmount                      *string
	DisplayedPriceIsInclusiveOfTax                                            *string
	EnableCreationOfAutomaticPaymentSurcharge                                 *string
	EnableDiscountProfile                                                     *string
	UpdateCurrentInvoiceToIncludeChargeAdjustments                            *string
	UpdateInvoiceToIncludeCurrentPaymentSurcharge                             *string
}

// SettingsGeneralSettingsTransactionAllowModificationOfFollowingCreditNoteFields is A
// SettingsGeneralSettingsTransaction's allow_modification_of_following_credit_note_fields.
type SettingsGeneralSettingsTransactionAllowModificationOfFollowingCreditNoteFields struct {
	CreditNoteDate *string
	Refundable     *string
}

// SettingsGeneralSettingsTransactionAllowModificationOfFollowingInvoiceFields is A
// SettingsGeneralSettingsTransaction's allow_modification_of_following_invoice_fields.
type SettingsGeneralSettingsTransactionAllowModificationOfFollowingInvoiceFields struct {
	BillingPeriod *string
	ChargePeriod  *string
	DueDate       *string
	IssueDate     *string
}

// SettingsGeneralSettingsTransactionAllowModificationOfFollowingPaymentFields is A
// SettingsGeneralSettingsTransaction's allow_modification_of_following_payment_fields.
type SettingsGeneralSettingsTransactionAllowModificationOfFollowingPaymentFields struct {
	PaymentDate      *string
	PaymentProcessor *string
	PaymentReference *string
}

// SettingsGeneralSettingsTransactionAllowModificationOfFollowingRefundFields is A
// SettingsGeneralSettingsTransaction's allow_modification_of_following_refund_fields.
type SettingsGeneralSettingsTransactionAllowModificationOfFollowingRefundFields struct {
	RefundDate      *string
	RefundReference *string
}

// SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreference is A
// SettingsGeneralSettingsTransaction's available_mediums_for_communication_preference.
type SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreference struct {
	Email       *SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceEmail
	PostalMail  *SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferencePostalMail
	TextMessage *SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceTextMessage
	VoiceMail   *SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceVoiceMail
}

// SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceEmail is A
// SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreference's email.
type SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceEmail struct {
	Default *string
	Enabled *string
}

// SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferencePostalMail is A
// SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreference's postal_mail.
type SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferencePostalMail struct {
	Default *string
	Enabled *string
}

// SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceTextMessage is A
// SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreference's text_message.
type SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceTextMessage struct {
	Default *string
	Enabled *string
}

// SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceVoiceMail is A
// SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreference's voice_mail.
type SettingsGeneralSettingsTransactionAvailableMediumsForCommunicationPreferenceVoiceMail struct {
	Default *string
	Enabled *string
}

// SettingsGeneralSettingsTransactionDefaultAccountingCodes is A
// SettingsGeneralSettingsTransaction's default_accounting_codes.
type SettingsGeneralSettingsTransactionDefaultAccountingCodes struct {
	AccountPayable        *string
	AccountReceivable     *string
	CashAndCashEquivalent *string
	CostOfGoodsSold       *string
	CreditNotePayable     *string
	Deduction             *string
	EventCharge           *string
	Inventory             *string
	PaymentSurcharge      *string
	Rounding              *string
	SalesRevenue          *string
	Tax                   *string
}

// SettingsPaymentProcessorSavedRecord is what CreateSettingsPaymentProcessor answers, under
// payment_processor: gateway_name and credentials where the read has provider.
type SettingsPaymentProcessorSavedRecord struct {
	AccountingCode    any
	ConversionRate    *float64
	CreatedBy         any
	CreatedOn         *string
	Credentials       any
	Currency          *SettingsPaymentProcessorSavedCurrency
	Description       *string
	Disconnected      *bool
	DisplayName       *string
	GatewayName       *string
	GatewayType       *string
	Href              *string
	ID                *string
	IntegrationState  *string
	IsDefault         *bool
	IsForeignCurrency *bool
	LastUpdatedBy     any
	LastUpdatedOn     *string
	Name              *string
	Object            *string
	Provider          *string
	Status            *string
	Version           *int64
}

// MeasuredFrom lists the endpoints SettingsPaymentProcessorSavedRecord was measured from.
func (SettingsPaymentProcessorSavedRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/settings/payment_processors",
		"PATCH /api/{api_version}/settings/payment_processors/{payment_processor_id}",
	}
}

// EnvelopeKey unwraps {"payment_processor": {...}}.
func (SettingsPaymentProcessorSavedRecord) EnvelopeKey() string { return PaymentProcessorKey }

// SettingsPaymentProcessorSavedCurrency is a SettingsPaymentProcessorSavedRecord's currency.
type SettingsPaymentProcessorSavedCurrency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// SettingsPaymentProcessorReadRecord is one payment processor, under payment_processor: the record
// the listing serves.
type SettingsPaymentProcessorReadRecord struct {
	AccountingCode    *SettingsPaymentProcessorAccountingCode
	ConversionRate    *float64
	CreatedBy         any
	CreatedOn         *string
	Currency          *SettingsPaymentProcessorCurrency
	Description       any
	Disconnected      *bool
	DisplayName       *string
	GatewayType       *string
	Href              *string
	ID                *string
	IntegrationState  *string
	IsDefault         *bool
	IsForeignCurrency *bool
	LastUpdatedBy     any
	LastUpdatedOn     *string
	Name              *string
	Object            *string
	Provider          *string
	Status            *string
	Version           *int64
}

// MeasuredFrom lists the endpoints SettingsPaymentProcessorReadRecord was measured from.
func (SettingsPaymentProcessorReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/payment_processors/{payment_processor_id}",
	}
}

// EnvelopeKey unwraps {"payment_processor": {...}}.
func (SettingsPaymentProcessorReadRecord) EnvelopeKey() string { return PaymentProcessorKey }

// SettingsPaymentProcessorAccountingCode is a SettingsPaymentProcessor's accounting_code.
type SettingsPaymentProcessorAccountingCode struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// SettingsPaymentProcessorCurrency is a SettingsPaymentProcessor's currency.
type SettingsPaymentProcessorCurrency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// SettingsPaymentTermCreatedRecord is what CreateSettingsPaymentTerm answers, under payment_term: a
// flat period and duration where the read nests them.
type SettingsPaymentTermCreatedRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   any
	DisplayName   *string
	Duration      *int64
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Ordering      any
	Period        *string
	Status        *string
}

// MeasuredFrom lists the endpoints SettingsPaymentTermCreatedRecord was measured from.
func (SettingsPaymentTermCreatedRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/settings/payment_terms",
	}
}

// EnvelopeKey unwraps {"payment_term": {...}}.
func (SettingsPaymentTermCreatedRecord) EnvelopeKey() string { return PaymentTermKey }

// SettingsPaymentTermReadRecord is one payment term, under payment_term: the record the listing
// serves.
type SettingsPaymentTermReadRecord struct {
	Description any
	DisplayName *string
	Duration    *SettingsPaymentTermDuration
	Href        *string
	ID          *string
	Name        *string
	Object      *string
	Status      *string
}

// MeasuredFrom lists the endpoints SettingsPaymentTermReadRecord was measured from.
func (SettingsPaymentTermReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/payment_terms/{payment_term_id}",
	}
}

// EnvelopeKey unwraps {"payment_term": {...}}.
func (SettingsPaymentTermReadRecord) EnvelopeKey() string { return PaymentTermKey }

// SettingsPaymentTermDuration is a SettingsPaymentTerm's duration.
type SettingsPaymentTermDuration struct {
	Count  *int64
	Period *string
}

// SettingsPaymentTermUpdatedRecord is what UpdateSettingsPaymentTerm answers, under payment_term:
// the read's fields and display_order.
type SettingsPaymentTermUpdatedRecord struct {
	Description  any
	DisplayName  *string
	Duration     *SettingsPaymentTermDuration
	Href         *string
	ID           *string
	Name         *string
	Object       *string
	Status       *string
	DisplayOrder *int64
}

// MeasuredFrom lists the endpoints SettingsPaymentTermUpdatedRecord was measured from.
func (SettingsPaymentTermUpdatedRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/settings/payment_terms/{payment_term_id}",
	}
}

// EnvelopeKey unwraps {"payment_term": {...}}.
func (SettingsPaymentTermUpdatedRecord) EnvelopeKey() string { return PaymentTermKey }

// SettingsPricingLevelCreatedRecord is what CreateSettingsPricingLevel answers, under
// pricing_level: price_tax_inclusive where the read has three tax fields.
type SettingsPricingLevelCreatedRecord struct {
	BuiltInItem       any
	CreatedBy         any
	CreatedOn         *string
	Description       any
	DisplayName       *string
	Href              *string
	ID                *string
	IsDefault         any
	LastUpdatedBy     any
	LastUpdatedOn     *string
	Name              *string
	Object            *string
	Ordering          *int64
	PriceTaxInclusive any
	Status            *string
	TaxExempt         any
}

// MeasuredFrom lists the endpoints SettingsPricingLevelCreatedRecord was measured from.
func (SettingsPricingLevelCreatedRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/settings/pricing_levels",
	}
}

// EnvelopeKey unwraps {"pricing_level": {...}}.
func (SettingsPricingLevelCreatedRecord) EnvelopeKey() string { return PricingLevelKey }

// SettingsPricingLevelReadRecord is one pricing level, under pricing_level: the record the listing
// serves.
type SettingsPricingLevelReadRecord struct {
	BuiltInItem         any
	CreatedBy           any
	CreatedOn           *string
	Description         *string
	DisplayName         *string
	Href                *string
	ID                  *string
	IsDefault           any
	IsTaxInclusive      *bool
	LastUpdatedBy       any
	LastUpdatedOn       *string
	Name                *string
	Object              *string
	Ordering            *int64
	Status              *string
	TaxCode             *SettingsPricingLevelTaxCode
	TaxExempt           *bool
	TaxInclusiveBasedOn *string
}

// MeasuredFrom lists the endpoints SettingsPricingLevelReadRecord was measured from.
func (SettingsPricingLevelReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/pricing_levels/{pricing_level_id}",
		"PATCH /api/{api_version}/settings/pricing_levels/{pricing_level_id}",
	}
}

// EnvelopeKey unwraps {"pricing_level": {...}}.
func (SettingsPricingLevelReadRecord) EnvelopeKey() string { return PricingLevelKey }

// SettingsPricingLevelTaxCode is a SettingsPricingLevel's tax_code.
type SettingsPricingLevelTaxCode struct {
	Href   *string
	ID     *string
	Object *string
}

// SettingsShippingProfileReadRecord is one shipping profile, under shipping_profile: the record the
// listing serves.
type SettingsShippingProfileReadRecord struct {
	AccountingCode        *SettingsShippingProfileAccountingCode
	CreatedBy             any
	CreatedOn             *string
	Description           *string
	DisplayName           *string
	FixedAmount           *float64
	Href                  *string
	ID                    *string
	IntegrationConnection any
	InvoiceNote           *string
	LastUpdatedBy         any
	LastUpdatedOn         *string
	Name                  *string
	Object                *string
	PriceTaxInclusive     *bool
	SalesTaxRate          *float64
	ShippingType          *string
	Status                *string
	TaxExempt             *bool
	TaxRule               *SettingsShippingProfileTaxRule
	Version               *int64
}

// MeasuredFrom lists the endpoints SettingsShippingProfileReadRecord was measured from.
func (SettingsShippingProfileReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/shipping_profiles/{shipping_profile_id}",
		"POST /api/{api_version}/settings/shipping_profiles",
	}
}

// EnvelopeKey unwraps {"shipping_profile": {...}}.
func (SettingsShippingProfileReadRecord) EnvelopeKey() string { return ShippingProfileKey }

// SettingsShippingProfileAccountingCode is a SettingsShippingProfile's accounting_code.
type SettingsShippingProfileAccountingCode struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// SettingsShippingProfileTaxRule is a SettingsShippingProfile's tax_rule.
type SettingsShippingProfileTaxRule struct {
	Href   *string
	ID     *string
	Object *string
}

// SettingsWarehouseReadRecord is one warehouse, under warehouse: the record the listing serves.
type SettingsWarehouseReadRecord struct {
	City          *string
	Country       *SettingsWarehouseCountry
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Href          *string
	ID            *string
	IsDefault     *bool
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	PostCode      *string
	State         *string
	Status        *string
	Version       *int64
}

// MeasuredFrom lists the endpoints SettingsWarehouseReadRecord was measured from.
func (SettingsWarehouseReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/warehouses/{warehouse_id}",
		"POST /api/{api_version}/settings/warehouses",
		"PATCH /api/{api_version}/settings/warehouses/{warehouse_id}",
	}
}

// EnvelopeKey unwraps {"warehouse": {...}}.
func (SettingsWarehouseReadRecord) EnvelopeKey() string { return WarehouseKey }

// SettingsWarehouseCountry is a SettingsWarehouse's country.
type SettingsWarehouseCountry struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// SettingsTaxReadRecord is one tax, under tax: the record the listing serves.
type SettingsTaxReadRecord struct {
	Code               *string
	Country            *SettingsTaxCountry
	CountryTaxByStates []SettingsTaxCountryTaxByState
	CreatedBy          any
	CreatedOn          *string
	Href               *string
	ID                 *string
	LastUpdatedBy      any
	LastUpdatedOn      *string
	Name               *string
	Object             *string
	Rate               *float64
	Status             *string
}

// MeasuredFrom lists the endpoints SettingsTaxReadRecord was measured from.
func (SettingsTaxReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/taxes/{tax_id}",
	}
}

// EnvelopeKey unwraps {"tax": {...}}.
func (SettingsTaxReadRecord) EnvelopeKey() string { return TaxKey }

// SettingsTaxCountry is a SettingsTax's country.
type SettingsTaxCountry struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// SettingsTaxCountryTaxByState is one entry of a SettingsTax's country_tax_by_states.
type SettingsTaxCountryTaxByState struct {
	CreatedOn     *string
	ID            *string
	LastUpdatedOn *string
	Object        *string
	Rate          *float64
	State         any
	Status        *string
}

// SettingsCustomObjectDefinitionListShape reads a body as SettingsCustomObjectDefinitionList.
func SettingsCustomObjectDefinitionListShape(decoded any) any {
	return dto.Read[SettingsCustomObjectDefinitionList](decoded)
}

// SettingsCustomObjectDefinitionReadRecordShape reads a body as SettingsCustomObjectDefinitionReadRecord.
func SettingsCustomObjectDefinitionReadRecordShape(decoded any) any {
	return dto.Read[SettingsCustomObjectDefinitionReadRecord](decoded)
}

// SettingsCustomLogicScriptSavedRecordShape reads a body as SettingsCustomLogicScriptSavedRecord.
func SettingsCustomLogicScriptSavedRecordShape(decoded any) any {
	return dto.Read[SettingsCustomLogicScriptSavedRecord](decoded)
}

// SettingsDiscountProfileReadRecordShape reads a body as SettingsDiscountProfileReadRecord.
func SettingsDiscountProfileReadRecordShape(decoded any) any {
	return dto.Read[SettingsDiscountProfileReadRecord](decoded)
}

// SettingsFieldSettingListShape reads a body as SettingsFieldSettingList.
func SettingsFieldSettingListShape(decoded any) any {
	return dto.Read[SettingsFieldSettingList](decoded)
}

// SettingsFieldSettingReadRecordShape reads a body as SettingsFieldSettingReadRecord.
func SettingsFieldSettingReadRecordShape(decoded any) any {
	return dto.Read[SettingsFieldSettingReadRecord](decoded)
}

// SettingsGeneralSettingsRecordShape reads a body as SettingsGeneralSettingsRecord.
func SettingsGeneralSettingsRecordShape(decoded any) any {
	return dto.Read[SettingsGeneralSettingsRecord](decoded)
}

// SettingsPaymentProcessorSavedRecordShape reads a body as SettingsPaymentProcessorSavedRecord.
func SettingsPaymentProcessorSavedRecordShape(decoded any) any {
	return dto.Read[SettingsPaymentProcessorSavedRecord](decoded)
}

// SettingsPaymentProcessorReadRecordShape reads a body as SettingsPaymentProcessorReadRecord.
func SettingsPaymentProcessorReadRecordShape(decoded any) any {
	return dto.Read[SettingsPaymentProcessorReadRecord](decoded)
}

// SettingsPaymentTermCreatedRecordShape reads a body as SettingsPaymentTermCreatedRecord.
func SettingsPaymentTermCreatedRecordShape(decoded any) any {
	return dto.Read[SettingsPaymentTermCreatedRecord](decoded)
}

// SettingsPaymentTermReadRecordShape reads a body as SettingsPaymentTermReadRecord.
func SettingsPaymentTermReadRecordShape(decoded any) any {
	return dto.Read[SettingsPaymentTermReadRecord](decoded)
}

// SettingsPaymentTermUpdatedRecordShape reads a body as SettingsPaymentTermUpdatedRecord.
func SettingsPaymentTermUpdatedRecordShape(decoded any) any {
	return dto.Read[SettingsPaymentTermUpdatedRecord](decoded)
}

// SettingsPricingLevelCreatedRecordShape reads a body as SettingsPricingLevelCreatedRecord.
func SettingsPricingLevelCreatedRecordShape(decoded any) any {
	return dto.Read[SettingsPricingLevelCreatedRecord](decoded)
}

// SettingsPricingLevelReadRecordShape reads a body as SettingsPricingLevelReadRecord.
func SettingsPricingLevelReadRecordShape(decoded any) any {
	return dto.Read[SettingsPricingLevelReadRecord](decoded)
}

// SettingsShippingProfileReadRecordShape reads a body as SettingsShippingProfileReadRecord.
func SettingsShippingProfileReadRecordShape(decoded any) any {
	return dto.Read[SettingsShippingProfileReadRecord](decoded)
}

// SettingsWarehouseReadRecordShape reads a body as SettingsWarehouseReadRecord.
func SettingsWarehouseReadRecordShape(decoded any) any {
	return dto.Read[SettingsWarehouseReadRecord](decoded)
}

// SettingsTaxReadRecordShape reads a body as SettingsTaxReadRecord.
func SettingsTaxReadRecordShape(decoded any) any { return dto.Read[SettingsTaxReadRecord](decoded) }
