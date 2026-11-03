package usersettings

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the UserSettings method of the same name,
// which is the primary surface.

// GetUserSettings is the module form of UserSettings.GetUserSettings.
func GetUserSettings(requester exhttp.Requester, ctx context.Context, userID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetUserSettings(ctx, userID, raiseOnError, responseFormat...)
}

// GetUserSettingsProfileImage is the module form of UserSettings.GetUserSettingsProfileImage.
func GetUserSettingsProfileImage(requester exhttp.Requester, ctx context.Context, userID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetUserSettingsProfileImage(ctx, userID, raiseOnError, responseFormat...)
}
