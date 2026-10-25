// Package accounts is the v4 accounts endpoints.
package accounts

// The path templates for the accounts endpoints; the client fills {api_version}.
const (
	AccountsPath                      = "/api/{api_version}/accounts"
	AccountPath                       = "/api/{api_version}/accounts/{account_id}"
	AccountInformationPath            = "/api/{api_version}/accounts/{account_id}/information"
	AccountCancelPath                 = "/api/{api_version}/accounts/{account_id}/cancel"
	AccountSetPasswordPath            = "/api/{api_version}/accounts/{account_id}/set-password"
	AccountContactsPath               = "/api/{api_version}/accounts/{account_id}/contacts"
	AccountContactPath                = "/api/{api_version}/accounts/{account_id}/contacts/{contact_type}"
	AccountAddressesPath              = "/api/{api_version}/accounts/{account_id}/addresses"
	AccountAddressPath                = "/api/{api_version}/accounts/{account_id}/addresses/{address_id}"
	AccountBillingPreferencesPath     = "/api/{api_version}/accounts/{account_id}/billing_preferences"
	AccountPaymentMethodsPath         = "/api/{api_version}/accounts/{account_id}/payment_methods"
	AccountPaymentMethodPath          = "/api/{api_version}/accounts/{account_id}/payment_methods/{payment_method_id}"
	AccountNotesPath                  = "/api/{api_version}/accounts/{account_id}/notes"
	AccountNotePath                   = "/api/{api_version}/accounts/{account_id}/notes/{note_id}"
	AccountNoteFilesPath              = "/api/{api_version}/accounts/{account_id}/notes/{note_id}/files"
	AccountNoteFilePath               = "/api/{api_version}/accounts/{account_id}/notes/{note_id}/files/{file_id}"
	AccountImagePath                  = "/api/{api_version}/accounts/{account_id}/image"
	AccountEmailsPath                 = "/api/{api_version}/accounts/{account_id}/communications/emails"
	AccountEmailPath                  = "/api/{api_version}/accounts/{account_id}/communications/emails/{email_id}"
	AccountOrdersPath                 = "/api/{api_version}/accounts/{account_id}/orders"
	AccountInvoicesPath               = "/api/{api_version}/accounts/{account_id}/invoices"
	AccountPaymentsPath               = "/api/{api_version}/accounts/{account_id}/payments"
	AccountRefundsPath                = "/api/{api_version}/accounts/{account_id}/refunds"
	AccountCreditNotesPath            = "/api/{api_version}/accounts/{account_id}/credit_notes"
	AccountPurchaseOrdersPath         = "/api/{api_version}/accounts/{account_id}/purchase_orders"
	AccountPurchaseInvoicesPath       = "/api/{api_version}/accounts/{account_id}/purchase_invoices"
	AccountPurchasePaymentsPath       = "/api/{api_version}/accounts/{account_id}/purchase_payments"
	AccountPurchaseRefundsPath        = "/api/{api_version}/accounts/{account_id}/purchase_refunds"
	AccountCustomObjectsPath          = "/api/{api_version}/accounts/{account_id}/custom_objects"
	AccountInventoriesPath            = "/api/{api_version}/accounts/{account_id}/inventories"
	AccountQuotePartialInvoicesPath   = "/api/{api_version}/accounts/{account_id}/quote_partial_invoices"
	AccountReactivatePath             = "/api/{api_version}/accounts/{account_id}/reactivate"
	AccountCustomObjectPath           = "/api/{api_version}/accounts/{account_id}/custom_objects/{custom_object_id}"
	AccountCustomObjectInstancePath   = "/api/{api_version}/accounts/{account_id}/custom_objects_instances/{custom_object_instance_id}"
	AccountPurgePath                  = "/api/{api_version}/accounts/{account_id}/purge"
	AccountKPIsPath                   = "/api/{api_version}/accounts/{account_id}/kpis"
	AccountFinancialSummaryPath       = "/api/{api_version}/accounts/{account_id}/financial_summary"
	AccountCustomObjectInstanceOfPath = "/api/{api_version}/accounts/{account_id}/custom_objects/{custom_object_id}/instances/{custom_object_instance_id}"
	AccountComponentObjectsPath       = "/api/{api_version}/accounts/{account_id}/component/{component_id}"
	AccountComponentObjectPath        = "/api/{api_version}/accounts/{account_id}/component/{component_id}/{component_object_id}"
)

// AccountsKey is the key a listing's records arrive under, the v4 list envelope's data.
const AccountsKey = "data"

// AccountKey is the envelope key one record arrives under.
const AccountKey = "account"

// AddressKey is the envelope key one record arrives under.
const AddressKey = "address"

// BillingPreferencesKey is the envelope key one record arrives under.
const BillingPreferencesKey = "billing_preferences"

// ContactKey is the envelope key one record arrives under.
const ContactKey = "contact"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"

// PaymentMethodKey is the envelope key one record arrives under.
const PaymentMethodKey = "payment_method"

// accountsQueryParams are the names the listing takes besides its filter_<field> filters.
var accountsQueryParams = [3]string{"limit", "cursor", "sort"}

// AccountsQueryParams returns a fresh copy of the names the listing takes besides its filters.
func AccountsQueryParams() []string {
	return append(make([]string, 0, len(accountsQueryParams)), accountsQueryParams[:]...)
}

// FileKey is the envelope key one record arrives under.
const FileKey = "file"

// AccountKPIsKey is the envelope key one record arrives under.
const AccountKPIsKey = "account_kpis"

// CustomObjectKey is the envelope key one record arrives under.
const CustomObjectKey = "custom_object"

// CustomComponentObjectKey is the envelope key one record arrives under.
const CustomComponentObjectKey = "custom_component_object"
