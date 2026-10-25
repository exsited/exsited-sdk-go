package shippingrules

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// ShippingRules is the v4 shipping rules endpoints, bound to one client.
type ShippingRules struct {
	requester exhttp.Requester
}

// GetSettingsShippingRules sends GET /api/{api_version}/settings/shipping_rules.
func (s *ShippingRules) GetSettingsShippingRules(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsShippingRulesPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsShippingRuleListShape})
}

// CreateSettingsShippingRule sends POST /api/{api_version}/settings/shipping_rules.
func (s *ShippingRules) CreateSettingsShippingRule(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsShippingRulesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsShippingRuleRecordShape})
}

// GetSettingsShippingRule sends GET /api/{api_version}/settings/shipping_rules/{shipping_rule_id}.
func (s *ShippingRules) GetSettingsShippingRule(ctx context.Context, shippingRuleID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(shippingRuleID, "shipping_rule_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsShippingRulePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"shipping_rule_id": shippingRuleID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsShippingRuleRecordShape})
}

// DeleteSettingsShippingRule sends DELETE /api/{api_version}/settings/shipping_rules/{shipping_rule_id}; it has no typed view, so read ApiResponse.Map.
func (s *ShippingRules) DeleteSettingsShippingRule(ctx context.Context, shippingRuleID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(shippingRuleID, "shipping_rule_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsShippingRulePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"shipping_rule_id": shippingRuleID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsShippingRule sends PATCH /api/{api_version}/settings/shipping_rules/{shipping_rule_id}.
func (s *ShippingRules) UpdateSettingsShippingRule(ctx context.Context, shippingRuleID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(shippingRuleID, "shipping_rule_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsShippingRulePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"shipping_rule_id": shippingRuleID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsShippingRuleRecordShape})
}

// CreateSettingsShippingRuleConfiguration sends POST /api/{api_version}/settings/shipping_rules/{shipping_rule_id}/configuration.
func (s *ShippingRules) CreateSettingsShippingRuleConfiguration(ctx context.Context, shippingRuleID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(shippingRuleID, "shipping_rule_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsShippingRuleConfigurationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"shipping_rule_id": shippingRuleID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsShippingRuleRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "shipping_rules"

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

// requirePayload refuses a nil payload, a typed nil included, before any I/O.
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

// New binds the shipping rules endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *ShippingRules { return &ShippingRules{requester: requester} }

// Requester returns the requester this resource was bound to.
func (s *ShippingRules) Requester() exhttp.Requester { return s.requester }
