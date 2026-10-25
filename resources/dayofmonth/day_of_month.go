package dayofmonth

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// DayOfMonth is the v4 day of month endpoints, bound to one client.
type DayOfMonth struct {
	requester exhttp.Requester
}

// GetSettingsDaysOfMonth sends GET /api/{api_version}/settings/day_of_month.
func (d *DayOfMonth) GetSettingsDaysOfMonth(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return d.requester.SendGetRequest(ctx, SettingsDaysOfMonthPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsDayOfMonthListShape})
}

// GetSettingsDayOfMonth sends GET /api/{api_version}/settings/day_of_month/{day_of_month_id}.
func (d *DayOfMonth) GetSettingsDayOfMonth(ctx context.Context, dayOfMonthID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(dayOfMonthID, "day_of_month_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return d.requester.SendGetRequest(ctx, SettingsDayOfMonthPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"day_of_month_id": dayOfMonthID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsDayOfMonthRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "day_of_month"

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

// New binds the day of month endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *DayOfMonth { return &DayOfMonth{requester: requester} }

// Requester returns the requester this resource was bound to.
func (d *DayOfMonth) Requester() exhttp.Requester { return d.requester }
