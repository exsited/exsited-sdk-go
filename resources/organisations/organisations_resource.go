package organisations

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Organisations method of the same name,
// which is the primary surface.

// GetSettingsOrganisations is the module form of Organisations.GetSettingsOrganisations.
func GetSettingsOrganisations(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsOrganisations(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsOrganisation is the module form of Organisations.CreateSettingsOrganisation.
func CreateSettingsOrganisation(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsOrganisation(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsOrganisation is the module form of Organisations.GetSettingsOrganisation.
func GetSettingsOrganisation(requester exhttp.Requester, ctx context.Context, organisationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsOrganisation(ctx, organisationID, raiseOnError, responseFormat...)
}

// DeleteSettingsOrganisation is the module form of Organisations.DeleteSettingsOrganisation.
func DeleteSettingsOrganisation(requester exhttp.Requester, ctx context.Context, organisationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsOrganisation(ctx, organisationID, raiseOnError, responseFormat...)
}

// UpdateSettingsOrganisation is the module form of Organisations.UpdateSettingsOrganisation.
func UpdateSettingsOrganisation(requester exhttp.Requester, ctx context.Context, organisationID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsOrganisation(ctx, organisationID, payload, raiseOnError, responseFormat...)
}
