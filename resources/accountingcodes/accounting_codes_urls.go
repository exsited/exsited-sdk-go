// Package accountingcodes is the v4 accounting codes endpoints.
package accountingcodes

// The path templates for the accounting codes endpoints; the client fills {api_version}.
const (
	SettingsAccountingCodesPath = "/api/{api_version}/settings/accounting_codes"
	SettingsAccountingCodePath  = "/api/{api_version}/settings/accounting_codes/{accounting_code_id}"
)

// AccountingCodesKey is the key a listing's records arrive under, the v4 list envelope's data.
const AccountingCodesKey = "data"

// AccountingCodeKey is the envelope key one record arrives under.
const AccountingCodeKey = "accounting_code"
