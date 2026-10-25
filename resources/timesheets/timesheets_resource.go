package timesheets

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Timesheets method of the same name,
// which is the primary surface.

// GetTimesheets is the module form of Timesheets.GetTimesheets.
func GetTimesheets(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetTimesheets(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateTimesheet is the module form of Timesheets.CreateTimesheet.
func CreateTimesheet(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateTimesheet(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteTimesheet is the module form of Timesheets.DeleteTimesheet.
func DeleteTimesheet(requester exhttp.Requester, ctx context.Context, createdTimesheetID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteTimesheet(ctx, createdTimesheetID, raiseOnError, responseFormat...)
}

// UpdateTimesheet is the module form of Timesheets.UpdateTimesheet.
func UpdateTimesheet(requester exhttp.Requester, ctx context.Context, createdTimesheetID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateTimesheet(ctx, createdTimesheetID, payload, raiseOnError, responseFormat...)
}

// GetTimesheet is the module form of Timesheets.GetTimesheet.
func GetTimesheet(requester exhttp.Requester, ctx context.Context, createdTimesheetID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetTimesheet(ctx, createdTimesheetID, raiseOnError, responseFormat...)
}
