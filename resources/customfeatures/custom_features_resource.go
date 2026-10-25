package customfeatures

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the CustomFeatures method of the same name,
// which is the primary surface.

// CreateCustomFeature is the module form of CustomFeatures.CreateCustomFeature.
func CreateCustomFeature(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateCustomFeature(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteComponentByID is the module form of CustomFeatures.DeleteComponentByID.
func DeleteComponentByID(requester exhttp.Requester, ctx context.Context, customFeatureID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteComponentByID(ctx, customFeatureID, raiseOnError, responseFormat...)
}

// UpdateCustomFeatureByID is the module form of CustomFeatures.UpdateCustomFeatureByID.
func UpdateCustomFeatureByID(requester exhttp.Requester, ctx context.Context, customFeatureID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateCustomFeatureByID(ctx, customFeatureID, payload, raiseOnError, responseFormat...)
}

// CreateCustomFeatureByID is the module form of CustomFeatures.CreateCustomFeatureByID.
func CreateCustomFeatureByID(requester exhttp.Requester, ctx context.Context, customFeatureID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateCustomFeatureByID(ctx, customFeatureID, payload, raiseOnError, responseFormat...)
}

// GetCustomFeatures is the module form of CustomFeatures.GetCustomFeatures.
func GetCustomFeatures(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCustomFeatures(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetCustomFeatureByID is the module form of CustomFeatures.GetCustomFeatureByID.
func GetCustomFeatureByID(requester exhttp.Requester, ctx context.Context, customFeatureID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetCustomFeatureByID(ctx, customFeatureID, raiseOnError, responseFormat...)
}

// DeleteCustomFeatureByID is the module form of CustomFeatures.DeleteCustomFeatureByID.
func DeleteCustomFeatureByID(requester exhttp.Requester, ctx context.Context, customFeatureID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteCustomFeatureByID(ctx, customFeatureID, raiseOnError, responseFormat...)
}
