package chargingperiods

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// ChargingPeriods is the v4 charging periods endpoints, bound to one client.
type ChargingPeriods struct {
	requester exhttp.Requester
}

// GetSettingsChargingPeriods sends GET /api/{api_version}/settings/charging_periods.
func (c *ChargingPeriods) GetSettingsChargingPeriods(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, SettingsChargingPeriodsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsChargingPeriodListShape})
}

// CreateSettingsChargingPeriod sends POST /api/{api_version}/settings/charging_periods.
func (c *ChargingPeriods) CreateSettingsChargingPeriod(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, SettingsChargingPeriodsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsChargingPeriodRecordShape})
}

// GetSettingsChargingPeriod sends GET /api/{api_version}/settings/charging_periods/{charging_period_id}.
func (c *ChargingPeriods) GetSettingsChargingPeriod(ctx context.Context, chargingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(chargingPeriodID, "charging_period_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, SettingsChargingPeriodPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"charging_period_id": chargingPeriodID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsChargingPeriodRecordShape})
}

// DeleteSettingsChargingPeriod sends DELETE /api/{api_version}/settings/charging_periods/{charging_period_id}; it has no typed view, so read ApiResponse.Map.
func (c *ChargingPeriods) DeleteSettingsChargingPeriod(ctx context.Context, chargingPeriodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(chargingPeriodID, "charging_period_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendDeleteRequest(ctx, SettingsChargingPeriodPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"charging_period_id": chargingPeriodID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsChargingPeriod sends PATCH /api/{api_version}/settings/charging_periods/{charging_period_id}.
func (c *ChargingPeriods) UpdateSettingsChargingPeriod(ctx context.Context, chargingPeriodID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(chargingPeriodID, "charging_period_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPatchRequest(ctx, SettingsChargingPeriodPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"charging_period_id": chargingPeriodID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsChargingPeriodRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "charging_periods"

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

// New binds the charging periods endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *ChargingPeriods { return &ChargingPeriods{requester: requester} }

// Requester returns the requester this resource was bound to.
func (c *ChargingPeriods) Requester() exhttp.Requester { return c.requester }
