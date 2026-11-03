// Package usersettings is the v4 user settings endpoints.
package usersettings

// The path templates for the user settings endpoints; the client fills {api_version}.
//
// {user_id} is the user's usr_ id, which carries the hex of the user's op_ operator id;
const (
	UserSettingsPath             = "/api/{api_version}/user_settings/{user_id}"
	UserSettingsProfileImagePath = "/api/{api_version}/user_settings/{user_id}/profile_image"
)

// UserSettingsKey is the envelope key one record arrives under.
const UserSettingsKey = "user_settings"
