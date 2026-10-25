// Package settings is the v4 settings endpoints.
package settings

// The path templates for the settings endpoints; the client fills {api_version}.
const (
	CustomObjectsPath                       = "/api/{api_version}/custom_objects"
	SettingsPath                            = "/api/{api_version}/settings"
	CustomObjectPath                        = "/api/{api_version}/custom_objects/{custom_object_id}"
	SettingsCurrenciesPath                  = "/api/{api_version}/settings/currencies"
	SettingsCustomLogicScriptsPath          = "/api/{api_version}/settings/custom_logic_scripts"
	SettingsCustomObjectsPath               = "/api/{api_version}/settings/custom_objects"
	SettingsDiscountProfilesPath            = "/api/{api_version}/settings/discount_profiles"
	SettingsItemGroupsPath                  = "/api/{api_version}/settings/item_groups"
	SettingsPaymentProcessorsPath           = "/api/{api_version}/settings/payment_processors"
	SettingsPaymentTermsPath                = "/api/{api_version}/settings/payment_terms"
	SettingsPricingLevelsPath               = "/api/{api_version}/settings/pricing_levels"
	SettingsShippingProfilesPath            = "/api/{api_version}/settings/shipping_profiles"
	SettingsTaxesPath                       = "/api/{api_version}/settings/taxes"
	SettingsTimeZonesPath                   = "/api/{api_version}/settings/time_zones"
	SettingsVariationsPath                  = "/api/{api_version}/settings/variations"
	SettingsWarehousesPath                  = "/api/{api_version}/settings/warehouses"
	SettingsCurrencyPath                    = "/api/{api_version}/settings/currencies/{currency_id}"
	SettingsCustomLogicScriptPath           = "/api/{api_version}/settings/custom_logic_scripts/{custom_logic_script_id}"
	SettingsCustomObjectPath                = "/api/{api_version}/settings/custom_objects/{custom_object_id}"
	SettingsFieldSettingsAccountPath        = "/api/{api_version}/settings/field_settings/account"
	SettingsFieldSettingsItemPath           = "/api/{api_version}/settings/field_settings/item"
	SettingsGeneralSettingsAccountPath      = "/api/{api_version}/settings/general_settings/account"
	SettingsGeneralSettingsItemPath         = "/api/{api_version}/settings/general_settings/item"
	SettingsItemGroupPath                   = "/api/{api_version}/settings/item_groups/{item_group_id}"
	SettingsTimeZonePath                    = "/api/{api_version}/settings/time_zones/{time_zone_id}"
	SettingsVariationPath                   = "/api/{api_version}/settings/variations/{variation_id}"
	SettingsCustomObjectCustomizePath       = "/api/{api_version}/settings/custom_objects/{custom_object_id}/customize"
	SettingsItemGroupBillingPreferencesPath = "/api/{api_version}/settings/item_groups/{item_group_id}/billing_preferences"
	SettingsDiscountProfilePath             = "/api/{api_version}/settings/discount_profiles/{discount_profile_id}"
	SettingsFieldSettingsPath               = "/api/{api_version}/settings/field_settings"
	SettingsFieldSettingPath                = "/api/{api_version}/settings/field_settings/{field_setting_id}"
	SettingsGeneralSettingsGroupPath        = "/api/{api_version}/settings/general_settings/{group}"
	SettingsGeneralSettingPath              = "/api/{api_version}/settings/general_settings/{group}/{general_setting_id}"
	SettingsGeneralSettingsDashboardPath    = "/api/{api_version}/settings/general_settings/dashboard"
	SettingsGeneralSettingsDateAndTimePath  = "/api/{api_version}/settings/general_settings/date_and_time"
	SettingsGeneralSettingsDisplayPath      = "/api/{api_version}/settings/general_settings/display"
	SettingsGeneralSettingsGeneralPath      = "/api/{api_version}/settings/general_settings/general"
	SettingsGeneralSettingsLabourPath       = "/api/{api_version}/settings/general_settings/labour"
	SettingsGeneralSettingsOrderPath        = "/api/{api_version}/settings/general_settings/order"
	SettingsGeneralSettingsPortalPath       = "/api/{api_version}/settings/general_settings/portal"
	SettingsGeneralSettingsPurchasePath     = "/api/{api_version}/settings/general_settings/purchase"
	SettingsGeneralSettingsSecurityPath     = "/api/{api_version}/settings/general_settings/security"
	SettingsGeneralSettingsTransactionPath  = "/api/{api_version}/settings/general_settings/transaction"
	SettingsPaymentProcessorPath            = "/api/{api_version}/settings/payment_processors/{payment_processor_id}"
	SettingsPaymentTermPath                 = "/api/{api_version}/settings/payment_terms/{payment_term_id}"
	SettingsPricingLevelPath                = "/api/{api_version}/settings/pricing_levels/{pricing_level_id}"
	SettingsShippingProfilePath             = "/api/{api_version}/settings/shipping_profiles/{shipping_profile_id}"
	SettingsWarehousePath                   = "/api/{api_version}/settings/warehouses/{warehouse_id}"
	SettingsTaxPath                         = "/api/{api_version}/settings/taxes/{tax_id}"
)

// SettingsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsKey = "data"

// CurrencyKey is the envelope key one record arrives under.
const CurrencyKey = "currency"

// CustomLogicScriptKey is the envelope key one record arrives under.
const CustomLogicScriptKey = "custom_logic_script"

// CustomObjectKey is the envelope key one record arrives under.
const CustomObjectKey = "custom_object"

// ItemGroupKey is the envelope key one record arrives under.
const ItemGroupKey = "item_group"

// TimeZoneKey is the envelope key one record arrives under.
const TimeZoneKey = "time_zone"

// VariationKey is the envelope key one record arrives under.
const VariationKey = "variation"

// SettingsEnvelopeKey is the envelope key one record arrives under.
const SettingsEnvelopeKey = "settings"

// GeneralSettingsKey is the envelope key one record arrives under.
const GeneralSettingsKey = "general_settings"

// DiscountProfileKey is the envelope key one record arrives under.
const DiscountProfileKey = "discount_profile"

// FieldSettingKey is the envelope key one record arrives under.
const FieldSettingKey = "field_setting"

// PaymentProcessorKey is the envelope key one record arrives under.
const PaymentProcessorKey = "payment_processor"

// PaymentTermKey is the envelope key one record arrives under.
const PaymentTermKey = "payment_term"

// PricingLevelKey is the envelope key one record arrives under.
const PricingLevelKey = "pricing_level"

// ShippingProfileKey is the envelope key one record arrives under.
const ShippingProfileKey = "shipping_profile"

// WarehouseKey is the envelope key one record arrives under.
const WarehouseKey = "warehouse"

// TaxKey is the envelope key one record arrives under.
const TaxKey = "tax"
