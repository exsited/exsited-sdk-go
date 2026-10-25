// Package customfields is the v4 custom fields endpoints.
package customfields

// The path templates for the custom fields endpoints; the client fills {api_version}.
const (
	SettingsCustomFieldsPath = "/api/{api_version}/settings/custom_fields"
	SettingsCustomFieldPath  = "/api/{api_version}/settings/custom_fields/{custom_field_id}"
)

// SettingsCustomFieldsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsCustomFieldsKey = "data"

// CustomFieldKey is the envelope key one record arrives under.
const CustomFieldKey = "custom_field"
