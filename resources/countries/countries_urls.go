// Package countries is the v4 countries endpoints.
package countries

// The path templates for the countries endpoints; the client fills {api_version}.
const (
	SettingsCountriesPath = "/api/{api_version}/settings/countries"
	SettingsCountryPath   = "/api/{api_version}/settings/countries/{country_id}"
)

// SettingsCountriesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsCountriesKey = "data"

// CountryKey is the envelope key one record arrives under.
const CountryKey = "country"
