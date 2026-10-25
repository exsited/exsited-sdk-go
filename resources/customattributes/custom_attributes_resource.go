package customattributes

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the CustomAttributes method of the same name,
// which is the primary surface.

// GetCustomAttributes is the module form of CustomAttributes.GetCustomAttributes.
func GetCustomAttributes(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCustomAttributes(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateCustomAttribute is the module form of CustomAttributes.CreateCustomAttribute.
func CreateCustomAttribute(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateCustomAttribute(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteCustomAttribute is the module form of CustomAttributes.DeleteCustomAttribute.
func DeleteCustomAttribute(requester exhttp.Requester, ctx context.Context, customAttributeID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteCustomAttribute(ctx, customAttributeID, raiseOnError, responseFormat...)
}

// GetCustomAttribute is the module form of CustomAttributes.GetCustomAttribute.
func GetCustomAttribute(requester exhttp.Requester, ctx context.Context, customAttributeID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCustomAttribute(ctx, customAttributeID, queryParams, raiseOnError, responseFormat...)
}

// UpdateCustomAttribute is the module form of CustomAttributes.UpdateCustomAttribute.
func UpdateCustomAttribute(requester exhttp.Requester, ctx context.Context, customAttributeID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateCustomAttribute(ctx, customAttributeID, payload, raiseOnError, responseFormat...)
}
