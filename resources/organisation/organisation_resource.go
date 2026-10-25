package organisation

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Organisation method of the same name,
// which is the primary surface.

// GetSettingsOrganisation is the module form of Organisation.GetSettingsOrganisation.
func GetSettingsOrganisation(requester exhttp.Requester, ctx context.Context, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsOrganisation(ctx, raiseOnError, responseFormat...)
}
