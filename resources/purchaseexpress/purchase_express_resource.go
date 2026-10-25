package purchaseexpress

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the PurchaseExpress method of the same name,
// which is the primary surface.

// CreatePurchaseExpress is the module form of PurchaseExpress.CreatePurchaseExpress.
func CreatePurchaseExpress(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreatePurchaseExpress(ctx, payload, raiseOnError, responseFormat...)
}
