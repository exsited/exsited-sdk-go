package usersettings

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// UserSettings is the v4 user settings endpoints, bound to one client. The family has no
// listing.
type UserSettings struct {
	requester exhttp.Requester
}

// GetUserSettings sends GET /api/{api_version}/user_settings/{user_id}, under user_settings.
//
// userID is the usr_ id, which carries the same hex as the user's op_ id.
func (u *UserSettings) GetUserSettings(ctx context.Context, userID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(userID, "user_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return u.requester.SendGetRequest(ctx, UserSettingsPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"user_id": userID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: UserSettingsRecordShape})
}

// GetUserSettingsProfileImage sends GET /api/{api_version}/user_settings/{user_id}/profile_image;
// base64 is the image without a data URI prefix, and empty when the read has none to serve.
func (u *UserSettings) GetUserSettingsProfileImage(ctx context.Context, userID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(userID, "user_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return u.requester.SendGetRequest(ctx, UserSettingsProfileImagePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"user_id": userID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: UserSettingsProfileImageRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "user_settings"

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

// isNilValue reports whether a value is a typed nil inside a non-nil any; a zero-valued struct is not.
func isNilValue(value any) bool {
	reflected := reflect.ValueOf(value)
	switch reflected.Kind() {
	case reflect.Map, reflect.Slice, reflect.Pointer, reflect.Interface, reflect.Func:
		return reflected.IsNil()
	}
	return false
}

// requireValue refuses a path value that is nil, empty or only whitespace, before any I/O.
func requireValue(value any, what string) error {
	text, isText := value.(string)
	if value == nil || isNilValue(value) || isText && strings.TrimSpace(text) == "" {
		return exerrors.UsageErrorf("%s is required: the request cannot be built without it", what)
	}
	return nil
}

// New binds the user settings endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *UserSettings { return &UserSettings{requester: requester} }

// Requester returns the requester this resource was bound to.
func (u *UserSettings) Requester() exhttp.Requester { return u.requester }
