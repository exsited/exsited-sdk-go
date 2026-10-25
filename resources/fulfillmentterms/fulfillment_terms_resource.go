package fulfillmentterms

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the FulfillmentTerms method of the same name,
// which is the primary surface.

// GetSettingsFulfillmentTerms is the module form of FulfillmentTerms.GetSettingsFulfillmentTerms.
func GetSettingsFulfillmentTerms(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsFulfillmentTerms(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsFulfillmentTerm is the module form of FulfillmentTerms.CreateSettingsFulfillmentTerm.
func CreateSettingsFulfillmentTerm(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsFulfillmentTerm(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsFulfillmentTerm is the module form of FulfillmentTerms.GetSettingsFulfillmentTerm.
func GetSettingsFulfillmentTerm(requester exhttp.Requester, ctx context.Context, fulfillmentTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsFulfillmentTerm(ctx, fulfillmentTermID, raiseOnError, responseFormat...)
}

// DeleteSettingsFulfillmentTerm is the module form of FulfillmentTerms.DeleteSettingsFulfillmentTerm.
func DeleteSettingsFulfillmentTerm(requester exhttp.Requester, ctx context.Context, fulfillmentTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsFulfillmentTerm(ctx, fulfillmentTermID, raiseOnError, responseFormat...)
}

// UpdateSettingsFulfillmentTerm is the module form of FulfillmentTerms.UpdateSettingsFulfillmentTerm.
func UpdateSettingsFulfillmentTerm(requester exhttp.Requester, ctx context.Context, fulfillmentTermID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsFulfillmentTerm(ctx, fulfillmentTermID, payload, raiseOnError, responseFormat...)
}
