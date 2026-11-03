package operators

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Operators method of the same name,
// which is the primary surface.

// GetSettingsOperators is the module form of Operators.GetSettingsOperators.
func GetSettingsOperators(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsOperators(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsOperator is the module form of Operators.CreateSettingsOperator.
func CreateSettingsOperator(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsOperator(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsOperator is the module form of Operators.GetSettingsOperator.
func GetSettingsOperator(requester exhttp.Requester, ctx context.Context, operatorID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsOperator(ctx, operatorID, raiseOnError, responseFormat...)
}

// UpdateSettingsOperator is the module form of Operators.UpdateSettingsOperator.
func UpdateSettingsOperator(requester exhttp.Requester, ctx context.Context, operatorID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsOperator(ctx, operatorID, payload, raiseOnError, responseFormat...)
}
