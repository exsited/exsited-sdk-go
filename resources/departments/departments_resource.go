package departments

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Departments method of the same name,
// which is the primary surface.

// GetSettingsDepartments is the module form of Departments.GetSettingsDepartments.
func GetSettingsDepartments(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsDepartments(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsDepartment is the module form of Departments.GetSettingsDepartment.
func GetSettingsDepartment(requester exhttp.Requester, ctx context.Context, departmentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsDepartment(ctx, departmentID, raiseOnError, responseFormat...)
}
