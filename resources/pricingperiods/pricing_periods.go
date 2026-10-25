package pricingperiods

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// PricingPeriods is the v4 pricing periods endpoints, bound to one client.
type PricingPeriods struct {
	requester exhttp.Requester
}

// GetSettingsPricingPeriods sends GET /api/{api_version}/settings/pricing_periods.
func (p *PricingPeriods) GetSettingsPricingPeriods(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, SettingsPricingPeriodsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPricingPeriodListShape})
}

// CreateSettingsPricingPeriod sends POST /api/{api_version}/settings/pricing_periods.
func (p *PricingPeriods) CreateSettingsPricingPeriod(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, SettingsPricingPeriodsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPricingPeriodRecordShape})
}

// GetSettingsPricingPeriod sends GET /api/{api_version}/settings/pricing_periods/{pricing_period_id}.
func (p *PricingPeriods) GetSettingsPricingPeriod(ctx context.Context, pricingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(pricingPeriodID, "pricing_period_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, SettingsPricingPeriodPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"pricing_period_id": pricingPeriodID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPricingPeriodRecordShape})
}

// DeleteSettingsPricingPeriod sends DELETE /api/{api_version}/settings/pricing_periods/{pricing_period_id}; it has no typed view, so read ApiResponse.Map.
func (p *PricingPeriods) DeleteSettingsPricingPeriod(ctx context.Context, pricingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(pricingPeriodID, "pricing_period_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, SettingsPricingPeriodPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"pricing_period_id": pricingPeriodID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsPricingPeriod sends PATCH /api/{api_version}/settings/pricing_periods/{pricing_period_id}.
func (p *PricingPeriods) UpdateSettingsPricingPeriod(ctx context.Context, pricingPeriodID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(pricingPeriodID, "pricing_period_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPatchRequest(ctx, SettingsPricingPeriodPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"pricing_period_id": pricingPeriodID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsPricingPeriodRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "pricing_periods"

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

// New binds the pricing periods endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *PricingPeriods { return &PricingPeriods{requester: requester} }

// Requester returns the requester this resource was bound to.
func (p *PricingPeriods) Requester() exhttp.Requester { return p.requester }
