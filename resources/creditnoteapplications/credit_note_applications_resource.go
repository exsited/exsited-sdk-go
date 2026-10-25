package creditnoteapplications

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the CreditNoteApplications method of the same name,
// which is the primary surface.

// GetCreditNoteApplications is the module form of CreditNoteApplications.GetCreditNoteApplications.
func GetCreditNoteApplications(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCreditNoteApplications(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetCreditNoteApplication is the module form of CreditNoteApplications.GetCreditNoteApplication.
func GetCreditNoteApplication(requester exhttp.Requester, ctx context.Context, creditNoteApplicationID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCreditNoteApplication(ctx, creditNoteApplicationID, queryParams, raiseOnError, responseFormat...)
}

// GetInvoiceCreditNoteApplications is the module form of CreditNoteApplications.GetInvoiceCreditNoteApplications.
func GetInvoiceCreditNoteApplications(requester exhttp.Requester, ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInvoiceCreditNoteApplications(ctx, invoiceID, queryParams, raiseOnError, responseFormat...)
}
