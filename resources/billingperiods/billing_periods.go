package billingperiods

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// BillingPeriods is the v4 billing periods endpoints, bound to one client.
type BillingPeriods struct {
	requester exhttp.Requester
}

// GetSettingsBillingPeriods sends GET /api/{api_version}/settings/billing_periods.
func (b *BillingPeriods) GetSettingsBillingPeriods(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, SettingsBillingPeriodsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsBillingPeriodListShape})
}

// CreateSettingsBillingPeriod sends POST /api/{api_version}/settings/billing_periods.
func (b *BillingPeriods) CreateSettingsBillingPeriod(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendPostRequest(ctx, SettingsBillingPeriodsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsBillingPeriodRecordShape})
}

// GetSettingsBillingPeriod sends GET /api/{api_version}/settings/billing_periods/{billing_period_id}.
func (b *BillingPeriods) GetSettingsBillingPeriod(ctx context.Context, billingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(billingPeriodID, "billing_period_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, SettingsBillingPeriodPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"billing_period_id": billingPeriodID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsBillingPeriodRecordShape})
}

// DeleteSettingsBillingPeriod sends DELETE /api/{api_version}/settings/billing_periods/{billing_period_id}; it has no typed view, so read ApiResponse.Map.
func (b *BillingPeriods) DeleteSettingsBillingPeriod(ctx context.Context, billingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(billingPeriodID, "billing_period_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendDeleteRequest(ctx, SettingsBillingPeriodPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"billing_period_id": billingPeriodID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsBillingPeriod sends PATCH /api/{api_version}/settings/billing_periods/{billing_period_id}.
func (b *BillingPeriods) UpdateSettingsBillingPeriod(ctx context.Context, billingPeriodID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(billingPeriodID, "billing_period_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendPatchRequest(ctx, SettingsBillingPeriodPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"billing_period_id": billingPeriodID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsBillingPeriodRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "billing_periods"

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

// New binds the billing periods endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *BillingPeriods { return &BillingPeriods{requester: requester} }

// Requester returns the requester this resource was bound to.
func (b *BillingPeriods) Requester() exhttp.Requester { return b.requester }
