package accountingcodes

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the AccountingCodes method of the same name,
// which is the primary surface.

// GetSettingsAccountingCodes is the module form of AccountingCodes.GetSettingsAccountingCodes.
func GetSettingsAccountingCodes(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsAccountingCodes(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsAccountingCode is the module form of AccountingCodes.CreateSettingsAccountingCode.
func CreateSettingsAccountingCode(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsAccountingCode(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsAccountingCode is the module form of AccountingCodes.GetSettingsAccountingCode.
func GetSettingsAccountingCode(requester exhttp.Requester, ctx context.Context, accountingCodeID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsAccountingCode(ctx, accountingCodeID, queryParams, raiseOnError, responseFormat...)
}

// DeleteSettingsAccountingCode is the module form of AccountingCodes.DeleteSettingsAccountingCode.
func DeleteSettingsAccountingCode(requester exhttp.Requester, ctx context.Context, accountingCodeID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsAccountingCode(ctx, accountingCodeID, raiseOnError, responseFormat...)
}

// UpdateSettingsAccountingCode is the module form of AccountingCodes.UpdateSettingsAccountingCode.
func UpdateSettingsAccountingCode(requester exhttp.Requester, ctx context.Context, accountingCodeID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsAccountingCode(ctx, accountingCodeID, payload, raiseOnError, responseFormat...)
}
