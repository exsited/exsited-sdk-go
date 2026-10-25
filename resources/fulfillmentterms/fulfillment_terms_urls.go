// Package fulfillmentterms is the v4 fulfillment terms endpoints.
package fulfillmentterms

// The path templates for the fulfillment terms endpoints; the client fills {api_version}.
const (
	SettingsFulfillmentTermsPath = "/api/{api_version}/settings/fulfillment_terms"
	SettingsFulfillmentTermPath  = "/api/{api_version}/settings/fulfillment_terms/{fulfillment_term_id}"
)

// SettingsFulfillmentTermsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsFulfillmentTermsKey = "data"

// FulfillmentTermKey is the envelope key one record arrives under.
const FulfillmentTermKey = "fulfillment_term"
