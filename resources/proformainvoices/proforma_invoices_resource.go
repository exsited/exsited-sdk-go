package proformainvoices

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ProformaInvoices method of the same name,
// which is the primary surface.

// GetProformaInvoices is the module form of ProformaInvoices.GetProformaInvoices.
func GetProformaInvoices(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetProformaInvoices(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetProformaInvoice is the module form of ProformaInvoices.GetProformaInvoice.
func GetProformaInvoice(requester exhttp.Requester, ctx context.Context, proformaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetProformaInvoice(ctx, proformaID, raiseOnError, responseFormat...)
}

// DeleteProformaInvoice is the module form of ProformaInvoices.DeleteProformaInvoice.
func DeleteProformaInvoice(requester exhttp.Requester, ctx context.Context, proformaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteProformaInvoice(ctx, proformaID, raiseOnError, responseFormat...)
}

// CreateProformaInvoicePayment is the module form of ProformaInvoices.CreateProformaInvoicePayment.
func CreateProformaInvoicePayment(requester exhttp.Requester, ctx context.Context, proformaID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateProformaInvoicePayment(ctx, proformaID, payload, raiseOnError, responseFormat...)
}
