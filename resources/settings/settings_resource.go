package settings

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Settings method of the same name,
// which is the primary surface.

// GetCustomObjects is the module form of Settings.GetCustomObjects.
func GetCustomObjects(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCustomObjects(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettings is the module form of Settings.GetSettings.
func GetSettings(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettings(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetCustomObject is the module form of Settings.GetCustomObject.
func GetCustomObject(requester exhttp.Requester, ctx context.Context, customObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCustomObject(ctx, customObjectID, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsCurrencies is the module form of Settings.GetSettingsCurrencies.
func GetSettingsCurrencies(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCurrencies(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsCustomLogicScripts is the module form of Settings.GetSettingsCustomLogicScripts.
func GetSettingsCustomLogicScripts(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCustomLogicScripts(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsCustomObject is the module form of Settings.CreateSettingsCustomObject.
func CreateSettingsCustomObject(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsCustomObject(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsDiscountProfiles is the module form of Settings.GetSettingsDiscountProfiles.
func GetSettingsDiscountProfiles(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsDiscountProfiles(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsItemGroups is the module form of Settings.GetSettingsItemGroups.
func GetSettingsItemGroups(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsItemGroups(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsItemGroup is the module form of Settings.CreateSettingsItemGroup.
func CreateSettingsItemGroup(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsItemGroup(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsPaymentProcessors is the module form of Settings.GetSettingsPaymentProcessors.
func GetSettingsPaymentProcessors(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsPaymentProcessors(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsPaymentTerms is the module form of Settings.GetSettingsPaymentTerms.
func GetSettingsPaymentTerms(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsPaymentTerms(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsPricingLevels is the module form of Settings.GetSettingsPricingLevels.
func GetSettingsPricingLevels(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsPricingLevels(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsShippingProfiles is the module form of Settings.GetSettingsShippingProfiles.
func GetSettingsShippingProfiles(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsShippingProfiles(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsTaxes is the module form of Settings.GetSettingsTaxes.
func GetSettingsTaxes(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsTaxes(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsTimeZones is the module form of Settings.GetSettingsTimeZones.
func GetSettingsTimeZones(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsTimeZones(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsVariations is the module form of Settings.GetSettingsVariations.
func GetSettingsVariations(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsVariations(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsVariation is the module form of Settings.CreateSettingsVariation.
func CreateSettingsVariation(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsVariation(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsWarehouses is the module form of Settings.GetSettingsWarehouses.
func GetSettingsWarehouses(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsWarehouses(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsCurrency is the module form of Settings.GetSettingsCurrency.
func GetSettingsCurrency(requester exhttp.Requester, ctx context.Context, currencyID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCurrency(ctx, currencyID, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsCustomLogicScript is the module form of Settings.GetSettingsCustomLogicScript.
func GetSettingsCustomLogicScript(requester exhttp.Requester, ctx context.Context, customLogicScriptID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCustomLogicScript(ctx, customLogicScriptID, queryParams, raiseOnError, responseFormat...)
}

// DeleteSettingsCustomObject is the module form of Settings.DeleteSettingsCustomObject.
func DeleteSettingsCustomObject(requester exhttp.Requester, ctx context.Context, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsCustomObject(ctx, customObjectID, raiseOnError, responseFormat...)
}

// UpdateSettingsCustomObject is the module form of Settings.UpdateSettingsCustomObject.
func UpdateSettingsCustomObject(requester exhttp.Requester, ctx context.Context, customObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsCustomObject(ctx, customObjectID, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsFieldSettingsAccount is the module form of Settings.UpdateSettingsFieldSettingsAccount.
func UpdateSettingsFieldSettingsAccount(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsFieldSettingsAccount(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsFieldSettingsItem is the module form of Settings.UpdateSettingsFieldSettingsItem.
func UpdateSettingsFieldSettingsItem(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsFieldSettingsItem(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsAccount is the module form of Settings.UpdateSettingsGeneralSettingsAccount.
func UpdateSettingsGeneralSettingsAccount(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsAccount(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsItem is the module form of Settings.UpdateSettingsGeneralSettingsItem.
func UpdateSettingsGeneralSettingsItem(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsItem(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteSettingsItemGroup is the module form of Settings.DeleteSettingsItemGroup.
func DeleteSettingsItemGroup(requester exhttp.Requester, ctx context.Context, itemGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsItemGroup(ctx, itemGroupID, raiseOnError, responseFormat...)
}

// GetSettingsItemGroup is the module form of Settings.GetSettingsItemGroup.
func GetSettingsItemGroup(requester exhttp.Requester, ctx context.Context, itemGroupID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsItemGroup(ctx, itemGroupID, queryParams, raiseOnError, responseFormat...)
}

// UpdateSettingsItemGroup is the module form of Settings.UpdateSettingsItemGroup.
func UpdateSettingsItemGroup(requester exhttp.Requester, ctx context.Context, itemGroupID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsItemGroup(ctx, itemGroupID, payload, raiseOnError, responseFormat...)
}

// GetSettingsTimeZone is the module form of Settings.GetSettingsTimeZone.
func GetSettingsTimeZone(requester exhttp.Requester, ctx context.Context, timeZoneID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsTimeZone(ctx, timeZoneID, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsVariation is the module form of Settings.GetSettingsVariation.
func GetSettingsVariation(requester exhttp.Requester, ctx context.Context, variationID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsVariation(ctx, variationID, queryParams, raiseOnError, responseFormat...)
}

// UpdateSettingsVariation is the module form of Settings.UpdateSettingsVariation.
func UpdateSettingsVariation(requester exhttp.Requester, ctx context.Context, variationID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsVariation(ctx, variationID, payload, raiseOnError, responseFormat...)
}

// CreateSettingsCustomObjectCustomize is the module form of Settings.CreateSettingsCustomObjectCustomize.
func CreateSettingsCustomObjectCustomize(requester exhttp.Requester, ctx context.Context, customObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsCustomObjectCustomize(ctx, customObjectID, payload, raiseOnError, responseFormat...)
}

// CreateSettingsItemGroupBillingPreference is the module form of Settings.CreateSettingsItemGroupBillingPreference.
func CreateSettingsItemGroupBillingPreference(requester exhttp.Requester, ctx context.Context, itemGroupID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsItemGroupBillingPreference(ctx, itemGroupID, payload, raiseOnError, responseFormat...)
}

// GetSettingsCustomObjects is the module form of Settings.GetSettingsCustomObjects.
func GetSettingsCustomObjects(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCustomObjects(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsCustomObject is the module form of Settings.GetSettingsCustomObject.
func GetSettingsCustomObject(requester exhttp.Requester, ctx context.Context, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCustomObject(ctx, customObjectID, raiseOnError, responseFormat...)
}

// CreateCustomObject is the module form of Settings.CreateCustomObject.
func CreateCustomObject(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateCustomObject(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateCustomObject is the module form of Settings.UpdateCustomObject.
func UpdateCustomObject(requester exhttp.Requester, ctx context.Context, customObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateCustomObject(ctx, customObjectID, payload, raiseOnError, responseFormat...)
}

// DeleteCustomObject is the module form of Settings.DeleteCustomObject.
func DeleteCustomObject(requester exhttp.Requester, ctx context.Context, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteCustomObject(ctx, customObjectID, raiseOnError, responseFormat...)
}

// CreateSettingsCurrency is the module form of Settings.CreateSettingsCurrency.
func CreateSettingsCurrency(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsCurrency(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsCurrency is the module form of Settings.UpdateSettingsCurrency.
func UpdateSettingsCurrency(requester exhttp.Requester, ctx context.Context, currencyID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsCurrency(ctx, currencyID, payload, raiseOnError, responseFormat...)
}

// CreateSettingsCustomLogicScript is the module form of Settings.CreateSettingsCustomLogicScript.
func CreateSettingsCustomLogicScript(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsCustomLogicScript(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsCustomLogicScript is the module form of Settings.UpdateSettingsCustomLogicScript.
func UpdateSettingsCustomLogicScript(requester exhttp.Requester, ctx context.Context, customLogicScriptID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsCustomLogicScript(ctx, customLogicScriptID, payload, raiseOnError, responseFormat...)
}

// DeleteSettingsCustomLogicScript is the module form of Settings.DeleteSettingsCustomLogicScript.
func DeleteSettingsCustomLogicScript(requester exhttp.Requester, ctx context.Context, customLogicScriptID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsCustomLogicScript(ctx, customLogicScriptID, raiseOnError, responseFormat...)
}

// CreateSettingsDiscountProfile is the module form of Settings.CreateSettingsDiscountProfile.
func CreateSettingsDiscountProfile(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsDiscountProfile(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsDiscountProfile is the module form of Settings.GetSettingsDiscountProfile.
func GetSettingsDiscountProfile(requester exhttp.Requester, ctx context.Context, discountProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsDiscountProfile(ctx, discountProfileID, raiseOnError, responseFormat...)
}

// UpdateSettingsDiscountProfile is the module form of Settings.UpdateSettingsDiscountProfile.
func UpdateSettingsDiscountProfile(requester exhttp.Requester, ctx context.Context, discountProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsDiscountProfile(ctx, discountProfileID, payload, raiseOnError, responseFormat...)
}

// DeleteSettingsDiscountProfile is the module form of Settings.DeleteSettingsDiscountProfile.
func DeleteSettingsDiscountProfile(requester exhttp.Requester, ctx context.Context, discountProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsDiscountProfile(ctx, discountProfileID, raiseOnError, responseFormat...)
}

// GetSettingsFieldSettings is the module form of Settings.GetSettingsFieldSettings.
func GetSettingsFieldSettings(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsFieldSettings(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsFieldSetting is the module form of Settings.GetSettingsFieldSetting.
func GetSettingsFieldSetting(requester exhttp.Requester, ctx context.Context, fieldSettingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsFieldSetting(ctx, fieldSettingID, raiseOnError, responseFormat...)
}

// GetSettingsGeneralSettings is the module form of Settings.GetSettingsGeneralSettings.
func GetSettingsGeneralSettings(requester exhttp.Requester, ctx context.Context, group any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsGeneralSettings(ctx, group, raiseOnError, responseFormat...)
}

// GetSettingsGeneralSetting is the module form of Settings.GetSettingsGeneralSetting.
func GetSettingsGeneralSetting(requester exhttp.Requester, ctx context.Context, group any, generalSettingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsGeneralSetting(ctx, group, generalSettingID, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsDashboard is the module form of Settings.UpdateSettingsGeneralSettingsDashboard.
func UpdateSettingsGeneralSettingsDashboard(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsDashboard(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsDateAndTime is the module form of Settings.UpdateSettingsGeneralSettingsDateAndTime.
func UpdateSettingsGeneralSettingsDateAndTime(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsDateAndTime(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsDisplay is the module form of Settings.UpdateSettingsGeneralSettingsDisplay.
func UpdateSettingsGeneralSettingsDisplay(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsDisplay(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsGeneral is the module form of Settings.UpdateSettingsGeneralSettingsGeneral.
func UpdateSettingsGeneralSettingsGeneral(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsGeneral(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsLabour is the module form of Settings.UpdateSettingsGeneralSettingsLabour.
func UpdateSettingsGeneralSettingsLabour(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsLabour(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsOrder is the module form of Settings.UpdateSettingsGeneralSettingsOrder.
func UpdateSettingsGeneralSettingsOrder(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsOrder(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsPortal is the module form of Settings.UpdateSettingsGeneralSettingsPortal.
func UpdateSettingsGeneralSettingsPortal(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsPortal(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsPurchase is the module form of Settings.UpdateSettingsGeneralSettingsPurchase.
func UpdateSettingsGeneralSettingsPurchase(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsPurchase(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsSecurity is the module form of Settings.UpdateSettingsGeneralSettingsSecurity.
func UpdateSettingsGeneralSettingsSecurity(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsSecurity(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateSettingsGeneralSettingsTransaction is the module form of Settings.UpdateSettingsGeneralSettingsTransaction.
func UpdateSettingsGeneralSettingsTransaction(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsGeneralSettingsTransaction(ctx, payload, raiseOnError, responseFormat...)
}

// CreateSettingsPaymentProcessor is the module form of Settings.CreateSettingsPaymentProcessor.
func CreateSettingsPaymentProcessor(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsPaymentProcessor(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsPaymentProcessor is the module form of Settings.GetSettingsPaymentProcessor.
func GetSettingsPaymentProcessor(requester exhttp.Requester, ctx context.Context, paymentProcessorID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsPaymentProcessor(ctx, paymentProcessorID, raiseOnError, responseFormat...)
}

// UpdateSettingsPaymentProcessor is the module form of Settings.UpdateSettingsPaymentProcessor.
func UpdateSettingsPaymentProcessor(requester exhttp.Requester, ctx context.Context, paymentProcessorID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsPaymentProcessor(ctx, paymentProcessorID, payload, raiseOnError, responseFormat...)
}

// DeleteSettingsPaymentProcessor is the module form of Settings.DeleteSettingsPaymentProcessor.
func DeleteSettingsPaymentProcessor(requester exhttp.Requester, ctx context.Context, paymentProcessorID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsPaymentProcessor(ctx, paymentProcessorID, raiseOnError, responseFormat...)
}

// CreateSettingsPaymentTerm is the module form of Settings.CreateSettingsPaymentTerm.
func CreateSettingsPaymentTerm(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsPaymentTerm(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsPaymentTerm is the module form of Settings.GetSettingsPaymentTerm.
func GetSettingsPaymentTerm(requester exhttp.Requester, ctx context.Context, paymentTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsPaymentTerm(ctx, paymentTermID, raiseOnError, responseFormat...)
}

// UpdateSettingsPaymentTerm is the module form of Settings.UpdateSettingsPaymentTerm.
func UpdateSettingsPaymentTerm(requester exhttp.Requester, ctx context.Context, paymentTermID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsPaymentTerm(ctx, paymentTermID, payload, raiseOnError, responseFormat...)
}

// DeleteSettingsPaymentTerm is the module form of Settings.DeleteSettingsPaymentTerm.
func DeleteSettingsPaymentTerm(requester exhttp.Requester, ctx context.Context, paymentTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsPaymentTerm(ctx, paymentTermID, raiseOnError, responseFormat...)
}

// CreateSettingsPricingLevel is the module form of Settings.CreateSettingsPricingLevel.
func CreateSettingsPricingLevel(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsPricingLevel(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsPricingLevel is the module form of Settings.GetSettingsPricingLevel.
func GetSettingsPricingLevel(requester exhttp.Requester, ctx context.Context, pricingLevelID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsPricingLevel(ctx, pricingLevelID, raiseOnError, responseFormat...)
}

// UpdateSettingsPricingLevel is the module form of Settings.UpdateSettingsPricingLevel.
func UpdateSettingsPricingLevel(requester exhttp.Requester, ctx context.Context, pricingLevelID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsPricingLevel(ctx, pricingLevelID, payload, raiseOnError, responseFormat...)
}

// DeleteSettingsPricingLevel is the module form of Settings.DeleteSettingsPricingLevel.
func DeleteSettingsPricingLevel(requester exhttp.Requester, ctx context.Context, pricingLevelID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsPricingLevel(ctx, pricingLevelID, raiseOnError, responseFormat...)
}

// CreateSettingsShippingProfile is the module form of Settings.CreateSettingsShippingProfile.
func CreateSettingsShippingProfile(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsShippingProfile(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsShippingProfile is the module form of Settings.GetSettingsShippingProfile.
func GetSettingsShippingProfile(requester exhttp.Requester, ctx context.Context, shippingProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsShippingProfile(ctx, shippingProfileID, raiseOnError, responseFormat...)
}

// UpdateSettingsShippingProfile is the module form of Settings.UpdateSettingsShippingProfile.
func UpdateSettingsShippingProfile(requester exhttp.Requester, ctx context.Context, shippingProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsShippingProfile(ctx, shippingProfileID, payload, raiseOnError, responseFormat...)
}

// DeleteSettingsShippingProfile is the module form of Settings.DeleteSettingsShippingProfile.
func DeleteSettingsShippingProfile(requester exhttp.Requester, ctx context.Context, shippingProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsShippingProfile(ctx, shippingProfileID, raiseOnError, responseFormat...)
}

// CreateSettingsWarehouse is the module form of Settings.CreateSettingsWarehouse.
func CreateSettingsWarehouse(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsWarehouse(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsWarehouse is the module form of Settings.GetSettingsWarehouse.
func GetSettingsWarehouse(requester exhttp.Requester, ctx context.Context, warehouseID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsWarehouse(ctx, warehouseID, raiseOnError, responseFormat...)
}

// UpdateSettingsWarehouse is the module form of Settings.UpdateSettingsWarehouse.
func UpdateSettingsWarehouse(requester exhttp.Requester, ctx context.Context, warehouseID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsWarehouse(ctx, warehouseID, payload, raiseOnError, responseFormat...)
}

// DeleteSettingsWarehouse is the module form of Settings.DeleteSettingsWarehouse.
func DeleteSettingsWarehouse(requester exhttp.Requester, ctx context.Context, warehouseID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsWarehouse(ctx, warehouseID, raiseOnError, responseFormat...)
}

// CreateSettingsTax is the module form of Settings.CreateSettingsTax.
func CreateSettingsTax(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsTax(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsTax is the module form of Settings.GetSettingsTax.
func GetSettingsTax(requester exhttp.Requester, ctx context.Context, taxID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsTax(ctx, taxID, raiseOnError, responseFormat...)
}

// UpdateSettingsTax is the module form of Settings.UpdateSettingsTax.
func UpdateSettingsTax(requester exhttp.Requester, ctx context.Context, taxID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsTax(ctx, taxID, payload, raiseOnError, responseFormat...)
}

// DeleteSettingsVariation is the module form of Settings.DeleteSettingsVariation.
func DeleteSettingsVariation(requester exhttp.Requester, ctx context.Context, variationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsVariation(ctx, variationID, raiseOnError, responseFormat...)
}
