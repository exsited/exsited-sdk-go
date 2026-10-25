// Package shippingrules is the v4 shipping rules endpoints.
package shippingrules

// The path templates for the shipping rules endpoints; the client fills {api_version}.
const (
	SettingsShippingRulesPath             = "/api/{api_version}/settings/shipping_rules"
	SettingsShippingRulePath              = "/api/{api_version}/settings/shipping_rules/{shipping_rule_id}"
	SettingsShippingRuleConfigurationPath = "/api/{api_version}/settings/shipping_rules/{shipping_rule_id}/configuration"
)

// SettingsShippingRulesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsShippingRulesKey = "data"

// ShippingRuleKey is the envelope key one record arrives under.
const ShippingRuleKey = "shipping_rule"

// ConfigurationKey is the envelope key one record arrives under.
const ConfigurationKey = "configuration"
