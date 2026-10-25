package usages

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Usages method of the same name,
// which is the primary surface.

// GetUsages is the module form of Usages.GetUsages.
func GetUsages(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetUsages(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateUsage is the module form of Usages.CreateUsage.
func CreateUsage(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateUsage(ctx, payload, raiseOnError, responseFormat...)
}

// GetUsage is the module form of Usages.GetUsage.
func GetUsage(requester exhttp.Requester, ctx context.Context, usageID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetUsage(ctx, usageID, queryParams, raiseOnError, responseFormat...)
}

// DeleteUsage is the module form of Usages.DeleteUsage.
func DeleteUsage(requester exhttp.Requester, ctx context.Context, usageID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteUsage(ctx, usageID, raiseOnError, responseFormat...)
}

// UpdateUsage is the module form of Usages.UpdateUsage.
func UpdateUsage(requester exhttp.Requester, ctx context.Context, usageID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateUsage(ctx, usageID, payload, raiseOnError, responseFormat...)
}

// ReplaceUsage is the module form of Usages.ReplaceUsage.
func ReplaceUsage(requester exhttp.Requester, ctx context.Context, usageID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReplaceUsage(ctx, usageID, payload, raiseOnError, responseFormat...)
}
