package accounts

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
	"github.com/exsited/exsited-sdk-go/v4/resources/accountemails"
	"github.com/exsited/exsited-sdk-go/v4/resources/customcomponents"
	"github.com/exsited/exsited-sdk-go/v4/resources/orders"
	"github.com/exsited/exsited-sdk-go/v4/resources/payments"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchaseinvoices"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchaseorders"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchasepayments"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchaserefunds"
	"github.com/exsited/exsited-sdk-go/v4/resources/refunds"
)

// Accounts is the v4 account endpoints, bound to one client.
type Accounts struct {
	requester exhttp.Requester
}

// GetAccounts lists accounts, with optional filtering, sorting and paging: GET /api/{api_version}/accounts.
//
// Three query parameters control the listing (limit, cursor and sort), and
// everything else is a filter, spelled filter_<field> for equality and
// filter_<field>_<operator> otherwise. Build the names with exhttp.FilterParam
// and exhttp.SortParam rather than concatenating, since equality is the
// irregular case and takes no suffix.
func (a *Accounts) GetAccounts(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AccountListShape})
}

// GetAccount retrieves one account: GET /api/{api_version}/accounts/{account_id}.
//
// accountID is the account's id - the short code a response carries, "W13ZGY".
// Not the uuid: a response carries both, only id addresses the record in a path,
// and the uuid answers 404. Use the uuid with filter_uuid on GetAccounts instead.
func (a *Accounts) GetAccount(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: AccountRecordShape})
}

// GetAccountInformation retrieves an account's core information: GET /api/{api_version}/accounts/{account_id}/information.
//
// The read that pairs with UpdateAccountInformation. It answers the same body as
// GetAccount.
func (a *Accounts) GetAccountInformation(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountInformationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: AccountRecordShape})
}

// GetAccountImage retrieves the profile image reference: GET /api/{api_version}/accounts/{account_id}/image.
//
// A URI, not image bytes, and null when the account has no image.
func (a *Accounts) GetAccountImage(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountImagePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: AccountImageViewShape})
}

// GetAccountContacts lists an account's contacts: GET /api/{api_version}/accounts/{account_id}/contacts.
//
// The account envelope with the contacts on it, not the list envelope and not a
// bare list.
func (a *Accounts) GetAccountContacts(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountContactsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: AccountRecordShape})
}

// GetAccountAddresses lists an account's addresses: GET /api/{api_version}/accounts/{account_id}/addresses.
//
// Each address carries an "addr_…" id - what UpdateAccountAddress and
// DeleteAccountAddress take.
func (a *Accounts) GetAccountAddresses(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountAddressesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: AccountRecordShape})
}

// GetAccountBillingPreferences retrieves an account's billing preferences: GET /api/{api_version}/accounts/{account_id}/billing_preferences.
//
// The alignment setting reads back under two names; see BillingPreferences.
func (a *Accounts) GetAccountBillingPreferences(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountBillingPreferencesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: AccountRecordShape})
}

// GetAccountPaymentMethods lists an account's payment methods: GET /api/{api_version}/accounts/{account_id}/payment_methods.
//
// Each carries a "pm_…" id - what the update and delete take.
func (a *Accounts) GetAccountPaymentMethods(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountPaymentMethodsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: AccountRecordShape})
}

// GetAccountInventories lists an account's inventories: GET /api/{api_version}/accounts/{account_id}/inventories.
//
// Not the list envelope despite being a listing: unpaged, and the records are one
// level in under "account".
func (a *Accounts) GetAccountInventories(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountInventoriesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: AccountInventoriesViewShape})
}

// GetAccountQuotePartialInvoices lists partial invoices from quotes: GET /api/{api_version}/accounts/{account_id}/quote_partial_invoices.
//
// Unpaged, and the key inside does not match the path: the records are under
// "partial_invoices".
func (a *Accounts) GetAccountQuotePartialInvoices(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountQuotePartialInvoicesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: AccountQuotePartialInvoicesViewShape})
}

// GetAccountContact retrieves one contact, addressed by its type: GET /api/{api_version}/accounts/{account_id}/contacts/{contact_type}.
//
// contactType is "billing", "shipping" or "primary" - not its id. An account holds
// at most one contact of each type, so the type addresses it; passing the "con_…"
// id answers 404.
func (a *Accounts) GetAccountContact(ctx context.Context, accountID any, contactType any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountContactPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "contact_type": contactType},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountContactViewShape})
}

// GetAccountAddress retrieves one address by its "addr_…" id: GET /api/{api_version}/accounts/{account_id}/addresses/{address_id}.
func (a *Accounts) GetAccountAddress(ctx context.Context, accountID any, addressID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountAddressPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "address_id": addressID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountAddressViewShape})
}

// GetAccountPaymentMethod retrieves one payment method by its "pm_…" id: GET /api/{api_version}/accounts/{account_id}/payment_methods/{payment_method_id}.
//
// For a method id that does not belong to the account this answers 200 with no
// payment_methods key at all, not 404. Check for the key before indexing.
func (a *Accounts) GetAccountPaymentMethod(ctx context.Context, accountID any, paymentMethodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountPaymentMethodPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "payment_method_id": paymentMethodID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountPaymentMethodViewShape})
}

// GetAccountNote retrieves one note by its "note_…" id: GET /api/{api_version}/accounts/{account_id}/notes/{note_id}.
func (a *Accounts) GetAccountNote(ctx context.Context, accountID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: NoteViewShape})
}

// GetAccountNoteFiles lists the files attached to a note: GET /api/{api_version}/accounts/{account_id}/notes/{note_id}/files.
//
// Each entry is identified by uuid, not id.
func (a *Accounts) GetAccountNoteFiles(ctx context.Context, accountID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountNoteFilesPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: NoteFileListShape})
}

// GetAccountNoteFile retrieves one file attached to a note: GET /api/{api_version}/accounts/{account_id}/notes/{note_id}/files/{file_id}.
//
// exhttp.ApiResponse.Map may be []byte for a binary attachment, in which case JSON is empty
// and dto.Dto is nil.
func (a *Accounts) GetAccountNoteFile(ctx context.Context, accountID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AccountNoteFileViewShape})
}

// GetAccountEmail retrieves one email by its uuid: GET /api/{api_version}/accounts/{account_id}/communications/emails/{email_id}.
func (a *Accounts) GetAccountEmail(ctx context.Context, accountID any, emailID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountEmailPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "email_id": emailID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: accountemails.AccountCommunicationsEmailRecordShape})
}

// GetAccountNotes lists an account's notes: GET /api/{api_version}/accounts/{account_id}/notes.
func (a *Accounts) GetAccountNotes(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: NoteListShape})
}

// GetAccountEmails lists an account's emails: GET /api/{api_version}/accounts/{account_id}/communications/emails.
func (a *Accounts) GetAccountEmails(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountEmailsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: accountemails.AccountCommunicationsEmailListShape})
}

// GetAccountOrders lists an account's orders: GET /api/{api_version}/accounts/{account_id}/orders.
func (a *Accounts) GetAccountOrders(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountOrdersPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: orders.AccountOrdersRecordShape})
}

// GetAccountInvoices lists an account's invoices: GET /api/{api_version}/accounts/{account_id}/invoices.
func (a *Accounts) GetAccountInvoices(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountInvoicesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AccountInvoiceListShape})
}

// GetAccountPayments lists an account's payments: GET /api/{api_version}/accounts/{account_id}/payments.
func (a *Accounts) GetAccountPayments(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountPaymentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: payments.AccountPaymentsRecordShape})
}

// GetAccountRefunds lists an account's refunds: GET /api/{api_version}/accounts/{account_id}/refunds.
func (a *Accounts) GetAccountRefunds(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountRefundsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: refunds.AccountRefundsRecordShape})
}

// GetAccountCreditNotes lists an account's credit notes: GET /api/{api_version}/accounts/{account_id}/credit_notes.
func (a *Accounts) GetAccountCreditNotes(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountCreditNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AccountCreditNoteListShape})
}

// GetAccountPurchaseOrders lists an account's purchase orders: GET /api/{api_version}/accounts/{account_id}/purchase_orders.
func (a *Accounts) GetAccountPurchaseOrders(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountPurchaseOrdersPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: purchaseorders.AccountPurchaseOrdersRecordShape})
}

// GetAccountPurchaseInvoices lists an account's purchase invoices: GET /api/{api_version}/accounts/{account_id}/purchase_invoices.
func (a *Accounts) GetAccountPurchaseInvoices(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountPurchaseInvoicesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: purchaseinvoices.AccountPurchaseInvoicesRecordShape})
}

// GetAccountPurchasePayments lists an account's purchase payments: GET /api/{api_version}/accounts/{account_id}/purchase_payments.
func (a *Accounts) GetAccountPurchasePayments(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountPurchasePaymentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: purchasepayments.AccountPurchasePaymentsRecordShape})
}

// GetAccountPurchaseRefunds lists an account's purchase refunds: GET /api/{api_version}/accounts/{account_id}/purchase_refunds.
func (a *Accounts) GetAccountPurchaseRefunds(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountPurchaseRefundsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: purchaserefunds.AccountPurchaseRefundsRecordShape})
}

// GetAccountCustomObjects lists an account's custom objects: GET /api/{api_version}/accounts/{account_id}/custom_objects.
func (a *Accounts) GetAccountCustomObjects(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountCustomObjectsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AccountCustomObjectListShape})
}

// CancelAccount cancels an account, effective on a date: POST /api/{api_version}/accounts/{account_id}/cancel.
//
// payload is a map, a JSON string sent byte-for-byte, or an *AccountCancelBody.
// All three put the same bytes on the wire. effective_date is YYYY-MM-DD; any
// other format is a 400 with code "invalid".
func (a *Accounts) CancelAccount(ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountCancelPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AccountRecordShape})
}

// SetAccountPassword sets the account's portal password: POST /api/{api_version}/accounts/{account_id}/set-password.
//
// payload is {password, confirm_password} at the top level, not under an
// "account" envelope the way cancel's body is.
func (a *Accounts) SetAccountPassword(ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountSetPasswordPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// UpdateAccountInformation updates an account's core information: PATCH /api/{api_version}/accounts/{account_id}/information.
//
// A partial update: send only the fields to change. Writes are flatter than reads
// - currency and time_zone go in as plain strings while the read returns nested
// objects, so posting a read body back unchanged is a 400. addresses and
// communication_preference are refused with code not_understood.
func (a *Accounts) UpdateAccountInformation(ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPatchRequest(ctx, AccountInformationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// CreateAccountAddresses adds one or more addresses: POST /api/{api_version}/accounts/{account_id}/addresses.
//
// payload is a list under "addresses", at the top level with no "account"
// envelope. state and country are names, not ids, and the write spelling is
// post_code where the read returns postal_code.
func (a *Accounts) CreateAccountAddresses(ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountAddressesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// UpdateAccountAddress updates one address: PATCH /api/{api_version}/accounts/{account_id}/addresses/{address_id}.
//
// payload is one address under a singular "address" key, not the plural
// "addresses" that the create takes.
func (a *Accounts) UpdateAccountAddress(ctx context.Context, accountID any, addressID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPatchRequest(ctx, AccountAddressPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID, "address_id": addressID},
		JSONBody:   payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AccountRecordShape})
}

// UpdateAccountContact updates one contact, leaving unsent fields alone: PATCH /api/{api_version}/accounts/{account_id}/contacts/{contact_type}.
//
// This takes flat first_name, middle_name and last_name, while the read returns
// them nested under "name".
func (a *Accounts) UpdateAccountContact(ctx context.Context, accountID any, contactType any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPatchRequest(ctx, AccountContactPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID, "contact_type": contactType},
		JSONBody:   payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AccountRecordShape})
}

// ReplaceAccountContact replaces one contact wholesale: PUT /api/{api_version}/accounts/{account_id}/contacts/{contact_type}.
//
// A field the payload omits is cleared, not left alone.
func (a *Accounts) ReplaceAccountContact(ctx context.Context, accountID any, contactType any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPutRequest(ctx, AccountContactPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID, "contact_type": contactType},
		JSONBody:   payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AccountRecordShape})
}

// ReplaceAccountBillingPreferences replaces an account's billing preferences: PUT /api/{api_version}/accounts/{account_id}/billing_preferences.
//
// billing_start_date is an enum, not a date - "DAY_OF_MONTH" and the like. A
// calendar date is a 400 with code matches_invalid, which reads as a formatting
// complaint and is really a type one.
func (a *Accounts) ReplaceAccountBillingPreferences(ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPutRequest(ctx, AccountBillingPreferencesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AccountRecordShape})
}

// UpdateAccountPaymentMethod updates one payment method: PATCH /api/{api_version}/accounts/{account_id}/payment_methods/{payment_method_id}.
func (a *Accounts) UpdateAccountPaymentMethod(ctx context.Context, accountID any, paymentMethodID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPatchRequest(ctx, AccountPaymentMethodPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID, "payment_method_id": paymentMethodID},
		JSONBody:   payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: AccountRecordShape})
}

// CreateAccountNote attaches a note to an account: POST /api/{api_version}/accounts/{account_id}/notes.
//
// This endpoint takes a form, not JSON: the note text travels as the form field
// "note". Any files go beside it as multipart parts; with none the form is
// URL-encoded. HTML is accepted and comes back as written.
func (a *Accounts) CreateAccountNote(ctx context.Context, accountID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: NoteViewShape})
}

// UpdateAccountNote updates a note's content, a form like CreateAccountNote's: PATCH /api/{api_version}/accounts/{account_id}/notes/{note_id}.
func (a *Accounts) UpdateAccountNote(ctx context.Context, accountID any, noteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPatchRequest(ctx, AccountNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID, "note_id": noteID},
		FormData:   map[string]any{"note": note}, Files: files, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: NoteViewShape})
}

// UploadAccountImage sets the account's profile image: POST /api/{api_version}/accounts/{account_id}/image.
//
// Base64 in a JSON body, not a multipart upload, unlike the note endpoints.
// Encode the bytes yourself.
func (a *Accounts) UploadAccountImage(ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountImagePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// DeleteAccount deletes an account: DELETE /api/{api_version}/accounts/{account_id}.
//
// A 409 means the account is locked: one with orders, invoices or payments
// against it cannot be deleted, and the API answers "Account has been locked, this
// action cannot be performed" with code internal_error. It is a business rule, not
// a server fault, so do not retry it.
func (a *Accounts) DeleteAccount(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendDeleteRequest(ctx, AccountPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// DeleteAccountContact deletes one contact, addressed by its type: DELETE /api/{api_version}/accounts/{account_id}/contacts/{contact_type}.
func (a *Accounts) DeleteAccountContact(ctx context.Context, accountID any, contactType any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendDeleteRequest(ctx, AccountContactPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "contact_type": contactType},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteAccountAddress deletes one address: DELETE /api/{api_version}/accounts/{account_id}/addresses/{address_id}.
func (a *Accounts) DeleteAccountAddress(ctx context.Context, accountID any, addressID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendDeleteRequest(ctx, AccountAddressPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "address_id": addressID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteAccountPaymentMethod deletes one payment method: DELETE /api/{api_version}/accounts/{account_id}/payment_methods/{payment_method_id}.
func (a *Accounts) DeleteAccountPaymentMethod(ctx context.Context, accountID any, paymentMethodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendDeleteRequest(ctx, AccountPaymentMethodPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "payment_method_id": paymentMethodID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteAccountNote deletes one note: DELETE /api/{api_version}/accounts/{account_id}/notes/{note_id}.
func (a *Accounts) DeleteAccountNote(ctx context.Context, accountID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendDeleteRequest(ctx, AccountNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteAccountImage removes the account's profile image. Answers 204, so Map is nil: DELETE /api/{api_version}/accounts/{account_id}/image.
func (a *Accounts) DeleteAccountImage(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendDeleteRequest(ctx, AccountImagePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// CreateAccount sends POST /api/{api_version}/accounts.
func (a *Accounts) CreateAccount(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// ReplaceAccount sends PUT /api/{api_version}/accounts/{account_id}.
func (a *Accounts) ReplaceAccount(ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPutRequest(ctx, AccountPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// ReactivateAccount sends POST /api/{api_version}/accounts/{account_id}/reactivate.
func (a *Accounts) ReactivateAccount(ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountReactivatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// CreateAccountContacts sends POST /api/{api_version}/accounts/{account_id}/contacts.
func (a *Accounts) CreateAccountContacts(ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountContactsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// ReplaceAccountAddress sends PUT /api/{api_version}/accounts/{account_id}/addresses/{address_id}.
func (a *Accounts) ReplaceAccountAddress(ctx context.Context, accountID any, addressID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requireValue(addressID, "address_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPutRequest(ctx, AccountAddressPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID, "address_id": addressID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// CreateAccountPaymentMethod sends POST /api/{api_version}/accounts/{account_id}/payment_methods.
func (a *Accounts) CreateAccountPaymentMethod(ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountPaymentMethodsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// ReplaceAccountPaymentMethod sends PUT /api/{api_version}/accounts/{account_id}/payment_methods/{payment_method_id}.
func (a *Accounts) ReplaceAccountPaymentMethod(ctx context.Context, accountID any, paymentMethodID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requireValue(paymentMethodID, "payment_method_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPutRequest(ctx, AccountPaymentMethodPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID, "payment_method_id": paymentMethodID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// CreateAccountCustomObjectInstance sends POST /api/{api_version}/accounts/{account_id}/custom_objects/{custom_object_id}; it has no typed view, so read ApiResponse.Map.
func (a *Accounts) CreateAccountCustomObjectInstance(ctx context.Context, accountID any, customObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountCustomObjectPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID, "custom_object_id": customObjectID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateAccountCustomObjectInstance sends PATCH /api/{api_version}/accounts/{account_id}/custom_objects_instances/{custom_object_instance_id}.
func (a *Accounts) UpdateAccountCustomObjectInstance(ctx context.Context, accountID any, customObjectInstanceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectInstanceID, "custom_object_instance_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPatchRequest(ctx, AccountCustomObjectInstancePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID, "custom_object_instance_id": customObjectInstanceID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountRecordShape})
}

// SendAccountEmail sends POST /api/{api_version}/accounts/{account_id}/communications/emails; it has no typed view, so read ApiResponse.Map.
func (a *Accounts) SendAccountEmail(ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountEmailsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// PurgeAccount sends POST /api/{api_version}/accounts/{account_id}/purge; it has no typed view, so read ApiResponse.Map.
func (a *Accounts) PurgeAccount(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, AccountPurgePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetAccountKPIs sends GET /api/{api_version}/accounts/{account_id}/kpis.
func (a *Accounts) GetAccountKPIs(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountKPIsPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountKpisRecordShape})
}

// GetAccountFinancialSummary sends GET /api/{api_version}/accounts/{account_id}/financial_summary.
func (a *Accounts) GetAccountFinancialSummary(ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountFinancialSummaryPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountFinancialSummaryRecordShape})
}

// GetAccountCustomObject sends GET /api/{api_version}/accounts/{account_id}/custom_objects/{custom_object_id}.
func (a *Accounts) GetAccountCustomObject(ctx context.Context, accountID any, customObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountCustomObjectPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID, "custom_object_id": customObjectID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountCustomObjectViewShape})
}

// GetAccountCustomObjectInstance sends GET /api/{api_version}/accounts/{account_id}/custom_objects/{custom_object_id}/instances/{custom_object_instance_id}; it has no typed view, so read ApiResponse.Map.
func (a *Accounts) GetAccountCustomObjectInstance(ctx context.Context, accountID any, customObjectID any, customObjectInstanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectInstanceID, "custom_object_instance_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountCustomObjectInstanceOfPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "custom_object_id": customObjectID, "custom_object_instance_id": customObjectInstanceID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetAccountComponentObjects sends GET /api/{api_version}/accounts/{account_id}/component/{component_id}.
func (a *Accounts) GetAccountComponentObjects(ctx context.Context, accountID any, componentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requireValue(componentID, "component_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountComponentObjectsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID, "component_id": componentID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: customcomponents.ComponentObjectsRecordShape})
}

// GetAccountComponentObject sends GET /api/{api_version}/accounts/{account_id}/component/{component_id}/{component_object_id}.
func (a *Accounts) GetAccountComponentObject(ctx context.Context, accountID any, componentID any, componentObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requireValue(componentID, "component_id"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AccountComponentObjectPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "component_id": componentID, "component_object_id": componentObjectID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountComponentObjectRecordShape})
}

// DeleteAccountNoteFile sends DELETE /api/{api_version}/accounts/{account_id}/notes/{note_id}/files/{file_id}; it has no typed view, so read ApiResponse.Map.
func (a *Accounts) DeleteAccountNoteFile(ctx context.Context, accountID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(fileID, "file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendDeleteRequest(ctx, AccountNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"account_id": accountID, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "accounts"

// oneResponseFormat reads the optional trailing responseFormat, refusing more than one.
func oneResponseFormat(responseFormat []string) (string, error) {
	switch len(responseFormat) {
	case 0:
		return "", nil
	case 1:
		return responseFormat[0], nil
	default:
		return "", exerrors.NewUsageError(fmt.Sprintf(
			"responseFormat: pass one format or none, got %d; the trailing argument is optional, "+
				"and omitting it defers to the client", len(responseFormat)))
	}
}

// requirePayload refuses a nil body before any I/O; the client's validator would let one through bodyless.
func requirePayload(payload any) error {
	if payload == nil || isNilValue(payload) {
		return exerrors.NewUsageError("payload is required: this endpoint takes a body, and nil would " +
			"send the request with no body at all")
	}
	return nil
}

// isNilValue reports whether a value is a typed nil inside a non-nil any; a zero-valued struct is not.
func isNilValue(value any) bool {
	reflected := reflect.ValueOf(value)
	switch reflected.Kind() {
	case reflect.Map, reflect.Slice, reflect.Pointer, reflect.Interface, reflect.Func:
		return reflected.IsNil()
	}
	return false
}

// requireNote refuses an empty note, which is the whole body of the request.
func requireNote(note string) error {
	if note == "" {
		return exerrors.NewUsageError("note is required: the note text is the body of this request")
	}
	return nil
}

// requireValue refuses a path value that is nil, empty or only whitespace, before any I/O.
func requireValue(value any, what string) error {
	text, isText := value.(string)
	if value == nil || isNilValue(value) || isText && strings.TrimSpace(text) == "" {
		return exerrors.UsageErrorf("%s is required: the request cannot be built without it", what)
	}
	return nil
}

// New binds the accounts endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Accounts { return &Accounts{requester: requester} }

// Requester returns the requester this resource was bound to.
func (a *Accounts) Requester() exhttp.Requester { return a.requester }
