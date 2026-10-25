package purchasecreditnoteapplications

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the PurchaseCreditNoteApplications method of the same name,
// which is the primary surface.

// GetPurchaseCreditNoteApplications is the module form of PurchaseCreditNoteApplications.GetPurchaseCreditNoteApplications.
func GetPurchaseCreditNoteApplications(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseCreditNoteApplications(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetPurchaseCreditNoteApplication is the module form of PurchaseCreditNoteApplications.GetPurchaseCreditNoteApplication.
func GetPurchaseCreditNoteApplication(requester exhttp.Requester, ctx context.Context, purchaseCreditNoteApplicationID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetPurchaseCreditNoteApplication(ctx, purchaseCreditNoteApplicationID, queryParams, raiseOnError, responseFormat...)
}
