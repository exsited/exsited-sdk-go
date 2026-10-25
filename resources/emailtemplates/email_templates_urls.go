// Package emailtemplates is the v4 email templates endpoints.
package emailtemplates

// The path templates for the email templates endpoints; the client fills {api_version}.
const (
	SettingsEmailTemplatesPath = "/api/{api_version}/settings/email_templates"
	SettingsEmailTemplatePath  = "/api/{api_version}/settings/email_templates/{email_template_id}"
)

// SettingsEmailTemplatesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsEmailTemplatesKey = "data"

// EmailTemplateKey is the envelope key one record arrives under.
const EmailTemplateKey = "email_template"
