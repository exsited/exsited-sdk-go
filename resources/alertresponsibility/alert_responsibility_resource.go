package alertresponsibility

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the AlertResponsibility method of the same name,
// which is the primary surface.

// GetAlertResponsibilities is the module form of AlertResponsibility.GetAlertResponsibilities.
func GetAlertResponsibilities(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAlertResponsibilities(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetAlertResponsibility is the module form of AlertResponsibility.GetAlertResponsibility.
func GetAlertResponsibility(requester exhttp.Requester, ctx context.Context, alertResponsibilityID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAlertResponsibility(ctx, alertResponsibilityID, raiseOnError, responseFormat...)
}

// CreateAlertResponsibility is the module form of AlertResponsibility.CreateAlertResponsibility.
func CreateAlertResponsibility(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateAlertResponsibility(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateAlertResponsibility is the module form of AlertResponsibility.UpdateAlertResponsibility.
func UpdateAlertResponsibility(requester exhttp.Requester, ctx context.Context, alertResponsibilityID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateAlertResponsibility(ctx, alertResponsibilityID, payload, raiseOnError, responseFormat...)
}

// DeleteAlertResponsibility is the module form of AlertResponsibility.DeleteAlertResponsibility.
func DeleteAlertResponsibility(requester exhttp.Requester, ctx context.Context, alertResponsibilityID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteAlertResponsibility(ctx, alertResponsibilityID, raiseOnError, responseFormat...)
}
