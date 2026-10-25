package shippingpackages

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ShippingPackages method of the same name,
// which is the primary surface.

// GetSettingsShippingPackages is the module form of ShippingPackages.GetSettingsShippingPackages.
func GetSettingsShippingPackages(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsShippingPackages(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsShippingPackage is the module form of ShippingPackages.CreateSettingsShippingPackage.
func CreateSettingsShippingPackage(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsShippingPackage(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsShippingPackage is the module form of ShippingPackages.GetSettingsShippingPackage.
func GetSettingsShippingPackage(requester exhttp.Requester, ctx context.Context, shippingPackageID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsShippingPackage(ctx, shippingPackageID, raiseOnError, responseFormat...)
}

// DeleteSettingsShippingPackage is the module form of ShippingPackages.DeleteSettingsShippingPackage.
func DeleteSettingsShippingPackage(requester exhttp.Requester, ctx context.Context, shippingPackageID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsShippingPackage(ctx, shippingPackageID, raiseOnError, responseFormat...)
}

// UpdateSettingsShippingPackage is the module form of ShippingPackages.UpdateSettingsShippingPackage.
func UpdateSettingsShippingPackage(requester exhttp.Requester, ctx context.Context, shippingPackageID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsShippingPackage(ctx, shippingPackageID, payload, raiseOnError, responseFormat...)
}
