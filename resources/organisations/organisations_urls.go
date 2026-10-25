// Package organisations is the v4 organisations endpoints.
package organisations

// The path templates for the organisations endpoints; the client fills {api_version}.
const (
	SettingsOrganisationsPath = "/api/{api_version}/settings/organisations"
	SettingsOrganisationPath  = "/api/{api_version}/settings/organisations/{organisation_id}"
)

// SettingsOrganisationsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsOrganisationsKey = "data"

// OrganizationKey is the envelope key one record arrives under.
const OrganizationKey = "organization"
