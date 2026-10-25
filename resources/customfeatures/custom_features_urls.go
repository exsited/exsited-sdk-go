// Package customfeatures is the v4 custom features endpoints.
package customfeatures

// The path templates for the custom features endpoints; the client fills {api_version}.
const (
	CustomFeaturePath     = "/api/{api_version}/custom_feature"
	ComponentByIDPath     = "/api/{api_version}/component/{custom_feature_id}"
	CustomFeatureByIDPath = "/api/{api_version}/custom_feature/{custom_feature_id}"
)

// CustomFeaturesKey is the key a listing's records arrive under, the v4 list envelope's data.
const CustomFeaturesKey = "data"

// CustomFeatureKey is the envelope key one record arrives under.
const CustomFeatureKey = "custom_feature"
