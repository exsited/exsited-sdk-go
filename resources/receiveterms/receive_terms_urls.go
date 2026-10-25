// Package receiveterms is the v4 receive terms endpoints.
package receiveterms

// The path templates for the receive terms endpoints; the client fills {api_version}.
const (
	SettingsReceiveTermsPath = "/api/{api_version}/settings/receive_terms"
	SettingsReceiveTermPath  = "/api/{api_version}/settings/receive_terms/{receive_term_id}"
)

// SettingsReceiveTermsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsReceiveTermsKey = "data"

// ReceiveTermKey is the envelope key one record arrives under.
const ReceiveTermKey = "receive_term"
