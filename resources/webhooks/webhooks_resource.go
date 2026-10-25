package webhooks

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Webhooks method of the same name,
// which is the primary surface.

// GetWebhook is the module form of Webhooks.GetWebhook.
func GetWebhook(requester exhttp.Requester, ctx context.Context, webhookID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetWebhook(ctx, webhookID, raiseOnError, responseFormat...)
}
