package alerts

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Alerts method of the same name,
// which is the primary surface.

// GetAlerts is the module form of Alerts.GetAlerts.
func GetAlerts(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAlerts(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetAlert is the module form of Alerts.GetAlert.
func GetAlert(requester exhttp.Requester, ctx context.Context, alertID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAlert(ctx, alertID, queryParams, raiseOnError, responseFormat...)
}

// GetAlertsAccountByID is the module form of Alerts.GetAlertsAccountByID.
func GetAlertsAccountByID(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAlertsAccountByID(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAlertsEmployeeByID is the module form of Alerts.GetAlertsEmployeeByID.
func GetAlertsEmployeeByID(requester exhttp.Requester, ctx context.Context, employeeID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAlertsEmployeeByID(ctx, employeeID, queryParams, raiseOnError, responseFormat...)
}

// GetAlertsLabourByID is the module form of Alerts.GetAlertsLabourByID.
func GetAlertsLabourByID(requester exhttp.Requester, ctx context.Context, labourID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAlertsLabourByID(ctx, labourID, queryParams, raiseOnError, responseFormat...)
}
