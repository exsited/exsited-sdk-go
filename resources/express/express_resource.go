package express

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Express method of the same name,
// which is the primary surface.

// CreateExpress is the module form of Express.CreateExpress.
func CreateExpress(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateExpress(ctx, payload, raiseOnError, responseFormat...)
}
