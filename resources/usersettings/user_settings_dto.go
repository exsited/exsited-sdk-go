package usersettings

import "github.com/exsited/exsited-sdk-go/v4/dto"

// UserSettingsRecord is one user's settings, under user_settings. CreatedBy and LastUpdatedBy
// are untyped.
type UserSettingsRecord struct {
	CreatedBy       any
	CreatedOn       *string
	Dashboard       *UserSettingsDashboard
	DisplaySettings *UserSettingsDisplaySettings
	Href            *string
	ID              *string
	LastUpdatedBy   any
	LastUpdatedOn   *string
	Object          *string
	Profile         *UserSettingsProfile
}

// MeasuredFrom lists the endpoints UserSettingsRecord is returned by.
func (UserSettingsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/user_settings/{user_id}",
	}
}

// EnvelopeKey unwraps {"user_settings": {...}}.
func (UserSettingsRecord) EnvelopeKey() string { return UserSettingsKey }

// UserSettingsProfile is a UserSettingsRecord's profile. TimeZone is untyped.
type UserSettingsProfile struct {
	DisplayName *string
	Email       *string
	ImageLink   *string
	Name        *string
	TimeZone    any
}

// UserSettingsDashboard is a UserSettingsRecord's dashboard. EnabledDashboards and
// DefaultDashboard are untyped.
type UserSettingsDashboard struct {
	DefaultDashboard  any
	EnabledDashboards map[string]any
	UseDefault        *bool
}

// UserSettingsDisplaySettings is a UserSettingsRecord's display_settings. Modules has no known
// element shape.
type UserSettingsDisplaySettings struct {
	Modules    []any
	UseDefault *bool
}

// UserSettingsProfileImageRecord is a user's profile image, under user_settings. Base64 is the
// image without a data URI prefix, and empty when the read has none to serve.
type UserSettingsProfileImageRecord struct {
	Base64 *string `wire:"base64"`
}

// MeasuredFrom lists the endpoints UserSettingsProfileImageRecord is returned by.
func (UserSettingsProfileImageRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/user_settings/{user_id}/profile_image",
	}
}

// EnvelopeKey unwraps {"user_settings": {...}}.
func (UserSettingsProfileImageRecord) EnvelopeKey() string { return UserSettingsKey }

// UserSettingsRecordShape reads a body as UserSettingsRecord.
func UserSettingsRecordShape(decoded any) any { return dto.Read[UserSettingsRecord](decoded) }

// UserSettingsProfileImageRecordShape reads a body as UserSettingsProfileImageRecord.
func UserSettingsProfileImageRecordShape(decoded any) any {
	return dto.Read[UserSettingsProfileImageRecord](decoded)
}
