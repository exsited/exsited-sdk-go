package teams

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Teams method of the same name,
// which is the primary surface.

// GetSettingsTeams is the module form of Teams.GetSettingsTeams.
func GetSettingsTeams(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsTeams(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsTeam is the module form of Teams.CreateSettingsTeam.
func CreateSettingsTeam(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsTeam(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsTeam is the module form of Teams.GetSettingsTeam.
func GetSettingsTeam(requester exhttp.Requester, ctx context.Context, teamID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsTeam(ctx, teamID, raiseOnError, responseFormat...)
}

// DeleteSettingsTeam is the module form of Teams.DeleteSettingsTeam.
func DeleteSettingsTeam(requester exhttp.Requester, ctx context.Context, teamID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsTeam(ctx, teamID, raiseOnError, responseFormat...)
}

// UpdateSettingsTeam is the module form of Teams.UpdateSettingsTeam.
func UpdateSettingsTeam(requester exhttp.Requester, ctx context.Context, teamID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsTeam(ctx, teamID, payload, raiseOnError, responseFormat...)
}
