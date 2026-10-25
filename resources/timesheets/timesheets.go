package timesheets

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Timesheets is the v4 timesheets endpoints, bound to one client.
type Timesheets struct {
	requester exhttp.Requester
}

// GetTimesheets sends GET /api/{api_version}/timesheets.
func (t *Timesheets) GetTimesheets(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendGetRequest(ctx, TimesheetsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: TimesheetListShape})
}

// CreateTimesheet sends POST /api/{api_version}/timesheets.
func (t *Timesheets) CreateTimesheet(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendPostRequest(ctx, TimesheetsPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: TimesheetCreateRecordShape})
}

// DeleteTimesheet sends DELETE /api/{api_version}/timesheets/{timesheet_id}; it has no typed view, so read ApiResponse.Map.
func (t *Timesheets) DeleteTimesheet(ctx context.Context, createdTimesheetID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(createdTimesheetID, "timesheet_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendDeleteRequest(ctx, TimesheetPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"timesheet_id": createdTimesheetID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateTimesheet sends PATCH /api/{api_version}/timesheets/{timesheet_id}.
func (t *Timesheets) UpdateTimesheet(ctx context.Context, createdTimesheetID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(createdTimesheetID, "timesheet_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendPatchRequest(ctx, TimesheetPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"timesheet_id": createdTimesheetID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: TimesheetCreateRecordShape})
}

// GetTimesheet sends GET /api/{api_version}/timesheets/{timesheet_id}.
func (t *Timesheets) GetTimesheet(ctx context.Context, createdTimesheetID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(createdTimesheetID, "timesheet_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendGetRequest(ctx, TimesheetPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"timesheet_id": createdTimesheetID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: TimesheetCreateRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "timesheets"

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

// requirePayload refuses a nil body before any I/O; the client's validator would let one through bodyless.
func requirePayload(payload any) error {
	if payload == nil || isNilValue(payload) {
		return exerrors.NewUsageError("payload is required: this endpoint takes a body, and nil would " +
			"send the request with no body at all")
	}
	return nil
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

// New binds the timesheets endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Timesheets { return &Timesheets{requester: requester} }

// Requester returns the requester this resource was bound to.
func (t *Timesheets) Requester() exhttp.Requester { return t.requester }
