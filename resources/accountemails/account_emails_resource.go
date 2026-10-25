package accountemails

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the AccountEmails method of the same name,
// which is the primary surface.

// DeleteAccountCommunicationsEmail is the module form of AccountEmails.DeleteAccountCommunicationsEmail.
func DeleteAccountCommunicationsEmail(requester exhttp.Requester, ctx context.Context, accountID any, emailUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteAccountCommunicationsEmail(ctx, accountID, emailUUID, raiseOnError, responseFormat...)
}
