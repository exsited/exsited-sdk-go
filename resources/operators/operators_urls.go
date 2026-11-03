// Package operators is the v4 operator endpoints.
package operators

// The path templates for the operator endpoints; the client fills {api_version}.
const (
	SettingsOperatorsPath = "/api/{api_version}/settings/operators"
	SettingsOperatorPath  = "/api/{api_version}/settings/operators/{operator_id}"
)

// SettingsOperatorsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsOperatorsKey = "data"

// OperatorKey is the envelope key one record arrives under.
const OperatorKey = "operator"
