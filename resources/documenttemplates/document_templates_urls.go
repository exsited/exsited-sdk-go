// Package documenttemplates is the v4 document templates endpoints.
package documenttemplates

// The path templates for the document templates endpoints; the client fills {api_version}.
const (
	SettingsDocumentTemplatesPath = "/api/{api_version}/settings/document_templates"
	SettingsDocumentTemplatePath  = "/api/{api_version}/settings/document_templates/{document_template_id}"
)

// SettingsDocumentTemplatesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsDocumentTemplatesKey = "data"

// DocumentTemplateKey is the envelope key one record arrives under.
const DocumentTemplateKey = "document_template"
