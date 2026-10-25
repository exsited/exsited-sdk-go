package shippingrules

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ShippingRules method of the same name,
// which is the primary surface.

// GetSettingsShippingRules is the module form of ShippingRules.GetSettingsShippingRules.
func GetSettingsShippingRules(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsShippingRules(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsShippingRule is the module form of ShippingRules.CreateSettingsShippingRule.
func CreateSettingsShippingRule(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsShippingRule(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsShippingRule is the module form of ShippingRules.GetSettingsShippingRule.
func GetSettingsShippingRule(requester exhttp.Requester, ctx context.Context, shippingRuleID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsShippingRule(ctx, shippingRuleID, raiseOnError, responseFormat...)
}

// DeleteSettingsShippingRule is the module form of ShippingRules.DeleteSettingsShippingRule.
func DeleteSettingsShippingRule(requester exhttp.Requester, ctx context.Context, shippingRuleID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsShippingRule(ctx, shippingRuleID, raiseOnError, responseFormat...)
}

// UpdateSettingsShippingRule is the module form of ShippingRules.UpdateSettingsShippingRule.
func UpdateSettingsShippingRule(requester exhttp.Requester, ctx context.Context, shippingRuleID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsShippingRule(ctx, shippingRuleID, payload, raiseOnError, responseFormat...)
}

// CreateSettingsShippingRuleConfiguration is the module form of ShippingRules.CreateSettingsShippingRuleConfiguration.
func CreateSettingsShippingRuleConfiguration(requester exhttp.Requester, ctx context.Context, shippingRuleID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsShippingRuleConfiguration(ctx, shippingRuleID, payload, raiseOnError, responseFormat...)
}
