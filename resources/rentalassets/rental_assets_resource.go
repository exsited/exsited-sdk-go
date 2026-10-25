package rentalassets

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the RentalAssets method of the same name,
// which is the primary surface.

// GetRentalAssetAvailability is the module form of RentalAssets.GetRentalAssetAvailability.
func GetRentalAssetAvailability(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetRentalAssetAvailability(ctx, queryParams, raiseOnError, responseFormat...)
}

// UpdateRentalAssetAvailability is the module form of RentalAssets.UpdateRentalAssetAvailability.
func UpdateRentalAssetAvailability(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateRentalAssetAvailability(ctx, payload, raiseOnError, responseFormat...)
}

// GetRentalAssetProfiles is the module form of RentalAssets.GetRentalAssetProfiles.
func GetRentalAssetProfiles(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetRentalAssetProfiles(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateRentalAssetProfile is the module form of RentalAssets.CreateRentalAssetProfile.
func CreateRentalAssetProfile(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateRentalAssetProfile(ctx, payload, raiseOnError, responseFormat...)
}

// GetRentalAssets is the module form of RentalAssets.GetRentalAssets.
func GetRentalAssets(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetRentalAssets(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateRentalAsset is the module form of RentalAssets.CreateRentalAsset.
func CreateRentalAsset(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateRentalAsset(ctx, payload, raiseOnError, responseFormat...)
}

// GetRentalAssetProfile is the module form of RentalAssets.GetRentalAssetProfile.
func GetRentalAssetProfile(requester exhttp.Requester, ctx context.Context, rentalAssetProfileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetRentalAssetProfile(ctx, rentalAssetProfileID, queryParams, raiseOnError, responseFormat...)
}

// UpdateRentalAssetProfile is the module form of RentalAssets.UpdateRentalAssetProfile.
func UpdateRentalAssetProfile(requester exhttp.Requester, ctx context.Context, rentalAssetProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateRentalAssetProfile(ctx, rentalAssetProfileID, payload, raiseOnError, responseFormat...)
}

// DeleteRentalAssetProfile is the module form of RentalAssets.DeleteRentalAssetProfile.
func DeleteRentalAssetProfile(requester exhttp.Requester, ctx context.Context, rentalAssetProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteRentalAssetProfile(ctx, rentalAssetProfileID, raiseOnError, responseFormat...)
}

// UpdateRentalAsset is the module form of RentalAssets.UpdateRentalAsset.
func UpdateRentalAsset(requester exhttp.Requester, ctx context.Context, rentalAssetUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateRentalAsset(ctx, rentalAssetUUID, payload, raiseOnError, responseFormat...)
}

// DeleteRentalAsset is the module form of RentalAssets.DeleteRentalAsset.
func DeleteRentalAsset(requester exhttp.Requester, ctx context.Context, rentalAssetUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteRentalAsset(ctx, rentalAssetUUID, raiseOnError, responseFormat...)
}

// GetRentalAsset is the module form of RentalAssets.GetRentalAsset.
func GetRentalAsset(requester exhttp.Requester, ctx context.Context, rentalAssetUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetRentalAsset(ctx, rentalAssetUUID, queryParams, raiseOnError, responseFormat...)
}

// CancelRentalAsset is the module form of RentalAssets.CancelRentalAsset.
func CancelRentalAsset(requester exhttp.Requester, ctx context.Context, rentalAssetID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CancelRentalAsset(ctx, rentalAssetID, payload, raiseOnError, responseFormat...)
}

// ReactivateRentalAsset is the module form of RentalAssets.ReactivateRentalAsset.
func ReactivateRentalAsset(requester exhttp.Requester, ctx context.Context, rentalAssetID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReactivateRentalAsset(ctx, rentalAssetID, payload, raiseOnError, responseFormat...)
}
