// Package customforms is the v4 custom forms endpoints.
package customforms

// The path templates for the custom forms endpoints; the client fills {api_version}.
const (
	SettingsCustomFormsPath = "/api/{api_version}/settings/custom_forms"
	SettingsCustomFormPath  = "/api/{api_version}/settings/custom_forms/{custom_form_id}"
)

// SettingsCustomFormsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsCustomFormsKey = "data"

// CustomFormKey is the envelope key one record arrives under.
const CustomFormKey = "custom_form"
