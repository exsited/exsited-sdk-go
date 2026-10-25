package reconciliations

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Reconciliations method of the same name,
// which is the primary surface.

// GetReconciliationSalePayments is the module form of Reconciliations.GetReconciliationSalePayments.
func GetReconciliationSalePayments(requester exhttp.Requester, ctx context.Context, reconciliationID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReconciliationSalePayments(ctx, reconciliationID, queryParams, raiseOnError, responseFormat...)
}

// GetReconciliationSaleRefunds is the module form of Reconciliations.GetReconciliationSaleRefunds.
func GetReconciliationSaleRefunds(requester exhttp.Requester, ctx context.Context, reconciliationID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReconciliationSaleRefunds(ctx, reconciliationID, queryParams, raiseOnError, responseFormat...)
}
