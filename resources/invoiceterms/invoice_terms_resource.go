package invoiceterms

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the InvoiceTerms method of the same name,
// which is the primary surface.

// GetSettingsInvoiceTerms is the module form of InvoiceTerms.GetSettingsInvoiceTerms.
func GetSettingsInvoiceTerms(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsInvoiceTerms(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsInvoiceTerm is the module form of InvoiceTerms.CreateSettingsInvoiceTerm.
func CreateSettingsInvoiceTerm(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsInvoiceTerm(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsInvoiceTerm is the module form of InvoiceTerms.GetSettingsInvoiceTerm.
func GetSettingsInvoiceTerm(requester exhttp.Requester, ctx context.Context, invoiceTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsInvoiceTerm(ctx, invoiceTermID, raiseOnError, responseFormat...)
}

// DeleteSettingsInvoiceTerm is the module form of InvoiceTerms.DeleteSettingsInvoiceTerm.
func DeleteSettingsInvoiceTerm(requester exhttp.Requester, ctx context.Context, invoiceTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsInvoiceTerm(ctx, invoiceTermID, raiseOnError, responseFormat...)
}

// UpdateSettingsInvoiceTerm is the module form of InvoiceTerms.UpdateSettingsInvoiceTerm.
func UpdateSettingsInvoiceTerm(requester exhttp.Requester, ctx context.Context, invoiceTermID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsInvoiceTerm(ctx, invoiceTermID, payload, raiseOnError, responseFormat...)
}
