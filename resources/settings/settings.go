package settings

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Settings is the v4 settings endpoints, bound to one client.
type Settings struct {
	requester exhttp.Requester
}

// GetCustomObjects sends GET /api/{api_version}/custom_objects.
func (s *Settings) GetCustomObjects(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, CustomObjectsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: CustomObjectListShape})
}

// GetSettings sends GET /api/{api_version}/settings.
func (s *Settings) GetSettings(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsRecordShape})
}

// GetCustomObject sends GET /api/{api_version}/custom_objects/{custom_object_id}.
func (s *Settings) GetCustomObject(ctx context.Context, customObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, CustomObjectPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"custom_object_id": customObjectID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsCustomObjectsRecordShape})
}

// GetSettingsCurrencies sends GET /api/{api_version}/settings/currencies.
func (s *Settings) GetSettingsCurrencies(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsCurrenciesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsCurrencyListShape})
}

// GetSettingsCustomLogicScripts sends GET /api/{api_version}/settings/custom_logic_scripts.
func (s *Settings) GetSettingsCustomLogicScripts(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsCustomLogicScriptsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsCustomLogicScriptsRecordShape})
}

// CreateSettingsCustomObject sends POST /api/{api_version}/settings/custom_objects.
func (s *Settings) CreateSettingsCustomObject(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsCustomObjectsPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsCustomObjectsRecordShape})
}

// GetSettingsDiscountProfiles sends GET /api/{api_version}/settings/discount_profiles.
func (s *Settings) GetSettingsDiscountProfiles(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsDiscountProfilesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsDiscountProfileListShape})
}

// GetSettingsItemGroups sends GET /api/{api_version}/settings/item_groups.
func (s *Settings) GetSettingsItemGroups(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsItemGroupsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsItemGroupListShape})
}

// CreateSettingsItemGroup sends POST /api/{api_version}/settings/item_groups.
func (s *Settings) CreateSettingsItemGroup(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsItemGroupsPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsItemGroupsRecordShape})
}

// GetSettingsPaymentProcessors sends GET /api/{api_version}/settings/payment_processors.
func (s *Settings) GetSettingsPaymentProcessors(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsPaymentProcessorsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsPaymentProcessorListShape})
}

// GetSettingsPaymentTerms sends GET /api/{api_version}/settings/payment_terms.
func (s *Settings) GetSettingsPaymentTerms(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsPaymentTermsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsPaymentTermListShape})
}

// GetSettingsPricingLevels sends GET /api/{api_version}/settings/pricing_levels.
func (s *Settings) GetSettingsPricingLevels(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsPricingLevelsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsPricingLevelListShape})
}

// GetSettingsShippingProfiles sends GET /api/{api_version}/settings/shipping_profiles.
func (s *Settings) GetSettingsShippingProfiles(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsShippingProfilesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsShippingProfileListShape})
}

// GetSettingsTaxes sends GET /api/{api_version}/settings/taxes.
func (s *Settings) GetSettingsTaxes(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsTaxesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsTaxListShape})
}

// GetSettingsTimeZones sends GET /api/{api_version}/settings/time_zones.
func (s *Settings) GetSettingsTimeZones(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsTimeZonesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsTimeZoneListShape})
}

// GetSettingsVariations sends GET /api/{api_version}/settings/variations.
func (s *Settings) GetSettingsVariations(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsVariationsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsVariationListShape})
}

// CreateSettingsVariation sends POST /api/{api_version}/settings/variations.
func (s *Settings) CreateSettingsVariation(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsVariationsPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsVariationsRecordShape})
}

// GetSettingsWarehouses sends GET /api/{api_version}/settings/warehouses.
func (s *Settings) GetSettingsWarehouses(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsWarehousesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsWarehouseListShape})
}

// GetSettingsCurrency sends GET /api/{api_version}/settings/currencies/{currency_id}.
func (s *Settings) GetSettingsCurrency(ctx context.Context, currencyID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(currencyID, "currency_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsCurrencyPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"currency_id": currencyID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsCurrencyReadRecordShape})
}

// GetSettingsCustomLogicScript sends GET /api/{api_version}/settings/custom_logic_scripts/{custom_logic_script_id}.
func (s *Settings) GetSettingsCustomLogicScript(ctx context.Context, customLogicScriptID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customLogicScriptID, "custom_logic_script_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsCustomLogicScriptPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"custom_logic_script_id": customLogicScriptID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsCustomLogicScriptRecordShape})
}

// DeleteSettingsCustomObject sends DELETE /api/{api_version}/settings/custom_objects/{custom_object_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) DeleteSettingsCustomObject(ctx context.Context, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsCustomObjectPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"custom_object_id": customObjectID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// UpdateSettingsCustomObject sends PATCH /api/{api_version}/settings/custom_objects/{custom_object_id}.
func (s *Settings) UpdateSettingsCustomObject(ctx context.Context, customObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsCustomObjectPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"custom_object_id": customObjectID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsCustomObjectsRecordShape})
}

// UpdateSettingsFieldSettingsAccount sends PATCH /api/{api_version}/settings/field_settings/account.
func (s *Settings) UpdateSettingsFieldSettingsAccount(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsFieldSettingsAccountPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsFieldSettingsItemRecordShape})
}

// UpdateSettingsFieldSettingsItem sends PATCH /api/{api_version}/settings/field_settings/item.
func (s *Settings) UpdateSettingsFieldSettingsItem(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsFieldSettingsItemPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsFieldSettingsItemRecordShape})
}

// UpdateSettingsGeneralSettingsAccount sends PATCH /api/{api_version}/settings/general_settings/account.
func (s *Settings) UpdateSettingsGeneralSettingsAccount(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsAccountPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsGeneralSettingsAccountUpdateRecordShape})
}

// UpdateSettingsGeneralSettingsItem sends PATCH /api/{api_version}/settings/general_settings/item.
func (s *Settings) UpdateSettingsGeneralSettingsItem(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsItemPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsGeneralSettingsItemUpdateRecordShape})
}

// DeleteSettingsItemGroup sends DELETE /api/{api_version}/settings/item_groups/{item_group_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) DeleteSettingsItemGroup(ctx context.Context, itemGroupID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemGroupID, "item_group_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsItemGroupPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_group_id": itemGroupID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetSettingsItemGroup sends GET /api/{api_version}/settings/item_groups/{item_group_id}.
func (s *Settings) GetSettingsItemGroup(ctx context.Context, itemGroupID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemGroupID, "item_group_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsItemGroupPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_group_id": itemGroupID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsItemGroupReadRecordShape})
}

// UpdateSettingsItemGroup sends PATCH /api/{api_version}/settings/item_groups/{item_group_id}.
func (s *Settings) UpdateSettingsItemGroup(ctx context.Context, itemGroupID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemGroupID, "item_group_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsItemGroupPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_group_id": itemGroupID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsItemGroupsRecordShape})
}

// GetSettingsTimeZone sends GET /api/{api_version}/settings/time_zones/{time_zone_id}.
func (s *Settings) GetSettingsTimeZone(ctx context.Context, timeZoneID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(timeZoneID, "time_zone_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsTimeZonePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"time_zone_id": timeZoneID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsTimeZoneReadRecordShape})
}

// GetSettingsVariation sends GET /api/{api_version}/settings/variations/{variation_id}.
func (s *Settings) GetSettingsVariation(ctx context.Context, variationID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(variationID, "variation_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsVariationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"variation_id": variationID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsVariationsRecordShape})
}

// UpdateSettingsVariation sends PATCH /api/{api_version}/settings/variations/{variation_id}.
func (s *Settings) UpdateSettingsVariation(ctx context.Context, variationID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(variationID, "variation_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsVariationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"variation_id": variationID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsVariationsRecordShape})
}

// CreateSettingsCustomObjectCustomize sends POST /api/{api_version}/settings/custom_objects/{custom_object_id}/customize.
func (s *Settings) CreateSettingsCustomObjectCustomize(ctx context.Context, customObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsCustomObjectCustomizePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"custom_object_id": customObjectID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsCustomObjectCustomizeRecordShape})
}

// CreateSettingsItemGroupBillingPreference sends POST /api/{api_version}/settings/item_groups/{item_group_id}/billing_preferences.
func (s *Settings) CreateSettingsItemGroupBillingPreference(ctx context.Context, itemGroupID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemGroupID, "item_group_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsItemGroupBillingPreferencesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_group_id": itemGroupID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsItemGroupsRecordShape})
}

// GetSettingsCustomObjects sends GET /api/{api_version}/settings/custom_objects.
func (s *Settings) GetSettingsCustomObjects(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsCustomObjectsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsCustomObjectDefinitionListShape})
}

// GetSettingsCustomObject sends GET /api/{api_version}/settings/custom_objects/{custom_object_id}.
func (s *Settings) GetSettingsCustomObject(ctx context.Context, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsCustomObjectPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"custom_object_id": customObjectID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsCustomObjectDefinitionReadRecordShape})
}

// CreateCustomObject sends POST /api/{api_version}/custom_objects.
func (s *Settings) CreateCustomObject(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, CustomObjectsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsCustomObjectsRecordShape})
}

// UpdateCustomObject sends PATCH /api/{api_version}/custom_objects/{custom_object_id}.
func (s *Settings) UpdateCustomObject(ctx context.Context, customObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, CustomObjectPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"custom_object_id": customObjectID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsCustomObjectsRecordShape})
}

// DeleteCustomObject sends DELETE /api/{api_version}/custom_objects/{custom_object_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) DeleteCustomObject(ctx context.Context, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, CustomObjectPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"custom_object_id": customObjectID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateSettingsCurrency sends POST /api/{api_version}/settings/currencies; it has no typed view, so read ApiResponse.Map.
func (s *Settings) CreateSettingsCurrency(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsCurrenciesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsCurrency sends PATCH /api/{api_version}/settings/currencies/{currency_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) UpdateSettingsCurrency(ctx context.Context, currencyID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(currencyID, "currency_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsCurrencyPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"currency_id": currencyID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateSettingsCustomLogicScript sends POST /api/{api_version}/settings/custom_logic_scripts.
func (s *Settings) CreateSettingsCustomLogicScript(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsCustomLogicScriptsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsCustomLogicScriptSavedRecordShape})
}

// UpdateSettingsCustomLogicScript sends PATCH /api/{api_version}/settings/custom_logic_scripts/{custom_logic_script_id}.
func (s *Settings) UpdateSettingsCustomLogicScript(ctx context.Context, customLogicScriptID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customLogicScriptID, "custom_logic_script_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsCustomLogicScriptPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"custom_logic_script_id": customLogicScriptID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsCustomLogicScriptSavedRecordShape})
}

// DeleteSettingsCustomLogicScript sends DELETE /api/{api_version}/settings/custom_logic_scripts/{custom_logic_script_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) DeleteSettingsCustomLogicScript(ctx context.Context, customLogicScriptID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(customLogicScriptID, "custom_logic_script_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsCustomLogicScriptPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"custom_logic_script_id": customLogicScriptID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateSettingsDiscountProfile sends POST /api/{api_version}/settings/discount_profiles.
func (s *Settings) CreateSettingsDiscountProfile(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsDiscountProfilesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsDiscountProfileReadRecordShape})
}

// GetSettingsDiscountProfile sends GET /api/{api_version}/settings/discount_profiles/{discount_profile_id}.
func (s *Settings) GetSettingsDiscountProfile(ctx context.Context, discountProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(discountProfileID, "discount_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsDiscountProfilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"discount_profile_id": discountProfileID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsDiscountProfileReadRecordShape})
}

// UpdateSettingsDiscountProfile sends PATCH /api/{api_version}/settings/discount_profiles/{discount_profile_id}.
func (s *Settings) UpdateSettingsDiscountProfile(ctx context.Context, discountProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(discountProfileID, "discount_profile_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsDiscountProfilePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"discount_profile_id": discountProfileID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsDiscountProfileReadRecordShape})
}

// DeleteSettingsDiscountProfile sends DELETE /api/{api_version}/settings/discount_profiles/{discount_profile_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) DeleteSettingsDiscountProfile(ctx context.Context, discountProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(discountProfileID, "discount_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsDiscountProfilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"discount_profile_id": discountProfileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetSettingsFieldSettings sends GET /api/{api_version}/settings/field_settings.
func (s *Settings) GetSettingsFieldSettings(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsFieldSettingsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsFieldSettingListShape})
}

// GetSettingsFieldSetting sends GET /api/{api_version}/settings/field_settings/{field_setting_id}.
func (s *Settings) GetSettingsFieldSetting(ctx context.Context, fieldSettingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fieldSettingID, "field_setting_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsFieldSettingPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"field_setting_id": fieldSettingID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsFieldSettingReadRecordShape})
}

// GetSettingsGeneralSettings sends GET /api/{api_version}/settings/general_settings/{group}.
func (s *Settings) GetSettingsGeneralSettings(ctx context.Context, group any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(group, "group"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsGeneralSettingsGroupPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"group": group},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsGeneralSettingsRecordShape})
}

// GetSettingsGeneralSetting sends GET /api/{api_version}/settings/general_settings/{group}/{general_setting_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) GetSettingsGeneralSetting(ctx context.Context, group any, generalSettingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(group, "group"); err != nil {
		return nil, err
	}
	if err := requireValue(generalSettingID, "general_setting_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsGeneralSettingPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"group": group, "general_setting_id": generalSettingID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsGeneralSettingsDashboard sends PATCH /api/{api_version}/settings/general_settings/dashboard.
func (s *Settings) UpdateSettingsGeneralSettingsDashboard(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsDashboardPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsGeneralSettingsRecordShape})
}

// UpdateSettingsGeneralSettingsDateAndTime sends PATCH /api/{api_version}/settings/general_settings/date_and_time.
func (s *Settings) UpdateSettingsGeneralSettingsDateAndTime(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsDateAndTimePath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsGeneralSettingsRecordShape})
}

// UpdateSettingsGeneralSettingsDisplay sends PATCH /api/{api_version}/settings/general_settings/display.
func (s *Settings) UpdateSettingsGeneralSettingsDisplay(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsDisplayPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsGeneralSettingsRecordShape})
}

// UpdateSettingsGeneralSettingsGeneral sends PATCH /api/{api_version}/settings/general_settings/general.
func (s *Settings) UpdateSettingsGeneralSettingsGeneral(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsGeneralPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsGeneralSettingsRecordShape})
}

// UpdateSettingsGeneralSettingsLabour sends PATCH /api/{api_version}/settings/general_settings/labour.
func (s *Settings) UpdateSettingsGeneralSettingsLabour(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsLabourPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsGeneralSettingsRecordShape})
}

// UpdateSettingsGeneralSettingsOrder sends PATCH /api/{api_version}/settings/general_settings/order.
func (s *Settings) UpdateSettingsGeneralSettingsOrder(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsOrderPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsGeneralSettingsRecordShape})
}

// UpdateSettingsGeneralSettingsPortal sends PATCH /api/{api_version}/settings/general_settings/portal.
func (s *Settings) UpdateSettingsGeneralSettingsPortal(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsPortalPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsGeneralSettingsRecordShape})
}

// UpdateSettingsGeneralSettingsPurchase sends PATCH /api/{api_version}/settings/general_settings/purchase.
func (s *Settings) UpdateSettingsGeneralSettingsPurchase(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsPurchasePath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsGeneralSettingsRecordShape})
}

// UpdateSettingsGeneralSettingsSecurity sends PATCH /api/{api_version}/settings/general_settings/security.
func (s *Settings) UpdateSettingsGeneralSettingsSecurity(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsSecurityPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsGeneralSettingsRecordShape})
}

// UpdateSettingsGeneralSettingsTransaction sends PATCH /api/{api_version}/settings/general_settings/transaction.
func (s *Settings) UpdateSettingsGeneralSettingsTransaction(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsGeneralSettingsTransactionPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsGeneralSettingsRecordShape})
}

// CreateSettingsPaymentProcessor sends POST /api/{api_version}/settings/payment_processors.
func (s *Settings) CreateSettingsPaymentProcessor(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsPaymentProcessorsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPaymentProcessorSavedRecordShape})
}

// GetSettingsPaymentProcessor sends GET /api/{api_version}/settings/payment_processors/{payment_processor_id}.
func (s *Settings) GetSettingsPaymentProcessor(ctx context.Context, paymentProcessorID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentProcessorID, "payment_processor_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsPaymentProcessorPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"payment_processor_id": paymentProcessorID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPaymentProcessorReadRecordShape})
}

// UpdateSettingsPaymentProcessor sends PATCH /api/{api_version}/settings/payment_processors/{payment_processor_id}.
func (s *Settings) UpdateSettingsPaymentProcessor(ctx context.Context, paymentProcessorID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentProcessorID, "payment_processor_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsPaymentProcessorPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"payment_processor_id": paymentProcessorID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPaymentProcessorSavedRecordShape})
}

// DeleteSettingsPaymentProcessor sends DELETE /api/{api_version}/settings/payment_processors/{payment_processor_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) DeleteSettingsPaymentProcessor(ctx context.Context, paymentProcessorID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentProcessorID, "payment_processor_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsPaymentProcessorPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"payment_processor_id": paymentProcessorID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateSettingsPaymentTerm sends POST /api/{api_version}/settings/payment_terms.
func (s *Settings) CreateSettingsPaymentTerm(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsPaymentTermsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPaymentTermCreatedRecordShape})
}

// GetSettingsPaymentTerm sends GET /api/{api_version}/settings/payment_terms/{payment_term_id}.
func (s *Settings) GetSettingsPaymentTerm(ctx context.Context, paymentTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentTermID, "payment_term_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsPaymentTermPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"payment_term_id": paymentTermID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPaymentTermReadRecordShape})
}

// UpdateSettingsPaymentTerm sends PATCH /api/{api_version}/settings/payment_terms/{payment_term_id}.
func (s *Settings) UpdateSettingsPaymentTerm(ctx context.Context, paymentTermID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentTermID, "payment_term_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsPaymentTermPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"payment_term_id": paymentTermID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPaymentTermUpdatedRecordShape})
}

// DeleteSettingsPaymentTerm sends DELETE /api/{api_version}/settings/payment_terms/{payment_term_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) DeleteSettingsPaymentTerm(ctx context.Context, paymentTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentTermID, "payment_term_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsPaymentTermPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"payment_term_id": paymentTermID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateSettingsPricingLevel sends POST /api/{api_version}/settings/pricing_levels.
func (s *Settings) CreateSettingsPricingLevel(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsPricingLevelsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPricingLevelCreatedRecordShape})
}

// GetSettingsPricingLevel sends GET /api/{api_version}/settings/pricing_levels/{pricing_level_id}.
func (s *Settings) GetSettingsPricingLevel(ctx context.Context, pricingLevelID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(pricingLevelID, "pricing_level_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsPricingLevelPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"pricing_level_id": pricingLevelID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPricingLevelReadRecordShape})
}

// UpdateSettingsPricingLevel sends PATCH /api/{api_version}/settings/pricing_levels/{pricing_level_id}.
func (s *Settings) UpdateSettingsPricingLevel(ctx context.Context, pricingLevelID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(pricingLevelID, "pricing_level_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsPricingLevelPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"pricing_level_id": pricingLevelID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPricingLevelReadRecordShape})
}

// DeleteSettingsPricingLevel sends DELETE /api/{api_version}/settings/pricing_levels/{pricing_level_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) DeleteSettingsPricingLevel(ctx context.Context, pricingLevelID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(pricingLevelID, "pricing_level_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsPricingLevelPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"pricing_level_id": pricingLevelID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateSettingsShippingProfile sends POST /api/{api_version}/settings/shipping_profiles.
func (s *Settings) CreateSettingsShippingProfile(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsShippingProfilesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsShippingProfileReadRecordShape})
}

// GetSettingsShippingProfile sends GET /api/{api_version}/settings/shipping_profiles/{shipping_profile_id}.
func (s *Settings) GetSettingsShippingProfile(ctx context.Context, shippingProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(shippingProfileID, "shipping_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsShippingProfilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"shipping_profile_id": shippingProfileID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsShippingProfileReadRecordShape})
}

// UpdateSettingsShippingProfile sends PATCH /api/{api_version}/settings/shipping_profiles/{shipping_profile_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) UpdateSettingsShippingProfile(ctx context.Context, shippingProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(shippingProfileID, "shipping_profile_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsShippingProfilePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"shipping_profile_id": shippingProfileID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteSettingsShippingProfile sends DELETE /api/{api_version}/settings/shipping_profiles/{shipping_profile_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) DeleteSettingsShippingProfile(ctx context.Context, shippingProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(shippingProfileID, "shipping_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsShippingProfilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"shipping_profile_id": shippingProfileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateSettingsWarehouse sends POST /api/{api_version}/settings/warehouses.
func (s *Settings) CreateSettingsWarehouse(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsWarehousesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsWarehouseReadRecordShape})
}

// GetSettingsWarehouse sends GET /api/{api_version}/settings/warehouses/{warehouse_id}.
func (s *Settings) GetSettingsWarehouse(ctx context.Context, warehouseID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(warehouseID, "warehouse_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsWarehousePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"warehouse_id": warehouseID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsWarehouseReadRecordShape})
}

// UpdateSettingsWarehouse sends PATCH /api/{api_version}/settings/warehouses/{warehouse_id}.
func (s *Settings) UpdateSettingsWarehouse(ctx context.Context, warehouseID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(warehouseID, "warehouse_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsWarehousePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"warehouse_id": warehouseID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsWarehouseReadRecordShape})
}

// DeleteSettingsWarehouse sends DELETE /api/{api_version}/settings/warehouses/{warehouse_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) DeleteSettingsWarehouse(ctx context.Context, warehouseID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(warehouseID, "warehouse_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsWarehousePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"warehouse_id": warehouseID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateSettingsTax sends POST /api/{api_version}/settings/taxes; it has no typed view, so read ApiResponse.Map.
func (s *Settings) CreateSettingsTax(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsTaxesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetSettingsTax sends GET /api/{api_version}/settings/taxes/{tax_id}.
func (s *Settings) GetSettingsTax(ctx context.Context, taxID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taxID, "tax_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsTaxPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"tax_id": taxID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsTaxReadRecordShape})
}

// UpdateSettingsTax sends PATCH /api/{api_version}/settings/taxes/{tax_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) UpdateSettingsTax(ctx context.Context, taxID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taxID, "tax_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsTaxPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"tax_id": taxID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteSettingsVariation sends DELETE /api/{api_version}/settings/variations/{variation_id}; it has no typed view, so read ApiResponse.Map.
func (s *Settings) DeleteSettingsVariation(ctx context.Context, variationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(variationID, "variation_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsVariationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"variation_id": variationID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "settings"

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

// requirePayload refuses a nil body before any I/O; the client's validator would let one through bodyless.
func requirePayload(payload any) error {
	if payload == nil || isNilValue(payload) {
		return exerrors.NewUsageError("payload is required: this endpoint takes a body, and nil would " +
			"send the request with no body at all")
	}
	return nil
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

// New binds the settings endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Settings { return &Settings{requester: requester} }

// Requester returns the requester this resource was bound to.
func (s *Settings) Requester() exhttp.Requester { return s.requester }
