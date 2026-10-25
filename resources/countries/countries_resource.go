package countries

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Countries method of the same name,
// which is the primary surface.

// GetSettingsCountries is the module form of Countries.GetSettingsCountries.
func GetSettingsCountries(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCountries(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsCountry is the module form of Countries.GetSettingsCountry.
func GetSettingsCountry(requester exhttp.Requester, ctx context.Context, countryID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsCountry(ctx, countryID, raiseOnError, responseFormat...)
}

// UpdateSettingsCountry is the module form of Countries.UpdateSettingsCountry.
func UpdateSettingsCountry(requester exhttp.Requester, ctx context.Context, countryID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsCountry(ctx, countryID, payload, raiseOnError, responseFormat...)
}
