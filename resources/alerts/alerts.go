package alerts

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Alerts is the v4 alerts endpoints, bound to one client.
type Alerts struct {
	requester exhttp.Requester
}

// GetAlerts sends GET /api/{api_version}/alerts.
func (a *Alerts) GetAlerts(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AlertsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AlertsRecordShape})
}

// GetAlert sends GET /api/{api_version}/alerts/{alert_id}; it has no typed view, so read ApiResponse.Map.
func (a *Alerts) GetAlert(ctx context.Context, alertID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(alertID, "alert_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AlertPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"alert_id": alertID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetAlertsAccountByID sends GET /api/{api_version}/alerts/account/{account_id}.
func (a *Alerts) GetAlertsAccountByID(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AlertsAccountByIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AlertsRecordShape})
}

// GetAlertsEmployeeByID sends GET /api/{api_version}/alerts/employee/{employee_id}.
func (a *Alerts) GetAlertsEmployeeByID(ctx context.Context, employeeID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(employeeID, "employee_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AlertsEmployeeByIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"employee_id": employeeID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AlertsEmployeeByIDListShape})
}

// GetAlertsLabourByID sends GET /api/{api_version}/alerts/labour/{labour_id}.
func (a *Alerts) GetAlertsLabourByID(ctx context.Context, labourID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourID, "labour_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AlertsLabourByIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_id": labourID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AlertsLabourByIDListShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "alerts"

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

// New binds the alerts endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Alerts { return &Alerts{requester: requester} }

// Requester returns the requester this resource was bound to.
func (a *Alerts) Requester() exhttp.Requester { return a.requester }
