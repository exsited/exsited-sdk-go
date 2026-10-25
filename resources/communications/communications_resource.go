package communications

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Communications method of the same name,
// which is the primary surface.

// SendEmail is the module form of Communications.SendEmail.
func SendEmail(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).SendEmail(ctx, payload, raiseOnError, responseFormat...)
}
