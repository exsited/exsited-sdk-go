package employees

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Employees method of the same name,
// which is the primary surface.

// GetEmployees is the module form of Employees.GetEmployees.
func GetEmployees(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetEmployees(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateEmployee is the module form of Employees.CreateEmployee.
func CreateEmployee(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateEmployee(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteEmployee is the module form of Employees.DeleteEmployee.
func DeleteEmployee(requester exhttp.Requester, ctx context.Context, employeeID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteEmployee(ctx, employeeID, raiseOnError, responseFormat...)
}

// GetEmployee is the module form of Employees.GetEmployee.
func GetEmployee(requester exhttp.Requester, ctx context.Context, employeeID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetEmployee(ctx, employeeID, queryParams, raiseOnError, responseFormat...)
}

// UpdateEmployee is the module form of Employees.UpdateEmployee.
func UpdateEmployee(requester exhttp.Requester, ctx context.Context, employeeID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateEmployee(ctx, employeeID, payload, raiseOnError, responseFormat...)
}
