package organisation

import (
	"context"
	"fmt"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Organisation is the v4 organisation endpoints, bound to one client.
type Organisation struct {
	requester exhttp.Requester
}

// GetSettingsOrganisation sends GET /api/{api_version}/settings/organisation.
func (o *Organisation) GetSettingsOrganisation(ctx context.Context, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, SettingsOrganisationPath, &exhttp.RequestOptions{
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsOrganisationRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "organisation"

// oneResponseFormat reads the optional trailing responseFormat, refusing more than one.
func oneResponseFormat(responseFormat []string) (string, error) {
	switch len(responseFormat) {
	case 0:
		return "", nil
	case 1:
		return responseFormat[0], nil
	default:
		return "", exerrors.NewUsageError(fmt.Sprintf(
			"responseFormat: pass one format or none, got %d; the trailing argument is optional, "+
				"and omitting it defers to the client", len(responseFormat)))
	}
}

// New binds the organisation endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Organisation { return &Organisation{requester: requester} }

// Requester returns the requester this resource was bound to.
func (o *Organisation) Requester() exhttp.Requester { return o.requester }
