package receiveterms

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ReceiveTerms method of the same name,
// which is the primary surface.

// GetSettingsReceiveTerms is the module form of ReceiveTerms.GetSettingsReceiveTerms.
func GetSettingsReceiveTerms(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsReceiveTerms(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsReceiveTerm is the module form of ReceiveTerms.CreateSettingsReceiveTerm.
func CreateSettingsReceiveTerm(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsReceiveTerm(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsReceiveTerm is the module form of ReceiveTerms.GetSettingsReceiveTerm.
func GetSettingsReceiveTerm(requester exhttp.Requester, ctx context.Context, receiveTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsReceiveTerm(ctx, receiveTermID, raiseOnError, responseFormat...)
}

// DeleteSettingsReceiveTerm is the module form of ReceiveTerms.DeleteSettingsReceiveTerm.
func DeleteSettingsReceiveTerm(requester exhttp.Requester, ctx context.Context, receiveTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsReceiveTerm(ctx, receiveTermID, raiseOnError, responseFormat...)
}

// UpdateSettingsReceiveTerm is the module form of ReceiveTerms.UpdateSettingsReceiveTerm.
func UpdateSettingsReceiveTerm(requester exhttp.Requester, ctx context.Context, receiveTermID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsReceiveTerm(ctx, receiveTermID, payload, raiseOnError, responseFormat...)
}
