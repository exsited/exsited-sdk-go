// Package customattributes is the v4 custom attributes endpoints.
package customattributes

// The path templates for the custom attributes endpoints; the client fills {api_version}.
const (
	CustomAttributesPath = "/api/{api_version}/custom_attributes"
	CustomAttributePath  = "/api/{api_version}/custom_attributes/{custom_attribute_id}"
)

// CustomAttributesKey is the key a listing's records arrive under, the v4 list envelope's data.
const CustomAttributesKey = "data"

// CustomAttributeKey is the envelope key one record arrives under.
const CustomAttributeKey = "custom_attribute"
