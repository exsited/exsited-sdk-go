// Package rentalassets is the v4 rental assets endpoints.
package rentalassets

// The path templates for the rental assets endpoints; the client fills {api_version}.
const (
	RentalAssetAvailabilityPath = "/api/{api_version}/rental_asset_availability"
	RentalAssetProfilesPath     = "/api/{api_version}/rental_asset_profiles"
	RentalAssetsPath            = "/api/{api_version}/rental_assets"
	RentalAssetProfilePath      = "/api/{api_version}/rental_asset_profiles/{rental_asset_profile_id}"
	RentalAssetPath             = "/api/{api_version}/rental_assets/{rental_asset_uuid}"
	RentalAssetCancelPath       = "/api/{api_version}/rental_assets/{rental_asset_id}/cancel"
	RentalAssetReactivatePath   = "/api/{api_version}/rental_assets/{rental_asset_id}/reactivate"
)

// RentalAssetsKey is the key a listing's records arrive under, the v4 list envelope's data.
const RentalAssetsKey = "data"

// RentalAssetKey is the envelope key one record arrives under.
const RentalAssetKey = "rental_asset"

// RentalAssetProfileKey is the envelope key one record arrives under.
const RentalAssetProfileKey = "rental_asset_profile"
