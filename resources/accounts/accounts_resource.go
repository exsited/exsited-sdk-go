package accounts

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Accounts method of the same name,
// which is the primary surface.

// GetAccounts is the module form of Accounts.GetAccounts.
func GetAccounts(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccounts(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetAccount is the module form of Accounts.GetAccount.
func GetAccount(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccount(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountInformation is the module form of Accounts.GetAccountInformation.
func GetAccountInformation(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountInformation(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountImage is the module form of Accounts.GetAccountImage.
func GetAccountImage(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountImage(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountContacts is the module form of Accounts.GetAccountContacts.
func GetAccountContacts(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountContacts(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountAddresses is the module form of Accounts.GetAccountAddresses.
func GetAccountAddresses(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountAddresses(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountBillingPreferences is the module form of Accounts.GetAccountBillingPreferences.
func GetAccountBillingPreferences(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountBillingPreferences(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountPaymentMethods is the module form of Accounts.GetAccountPaymentMethods.
func GetAccountPaymentMethods(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountPaymentMethods(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountInventories is the module form of Accounts.GetAccountInventories.
func GetAccountInventories(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountInventories(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountQuotePartialInvoices is the module form of Accounts.GetAccountQuotePartialInvoices.
func GetAccountQuotePartialInvoices(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountQuotePartialInvoices(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountContact is the module form of Accounts.GetAccountContact.
func GetAccountContact(requester exhttp.Requester, ctx context.Context, accountID any, contactType any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountContact(ctx, accountID, contactType, raiseOnError, responseFormat...)
}

// GetAccountAddress is the module form of Accounts.GetAccountAddress.
func GetAccountAddress(requester exhttp.Requester, ctx context.Context, accountID any, addressID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountAddress(ctx, accountID, addressID, raiseOnError, responseFormat...)
}

// GetAccountPaymentMethod is the module form of Accounts.GetAccountPaymentMethod.
func GetAccountPaymentMethod(requester exhttp.Requester, ctx context.Context, accountID any, paymentMethodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountPaymentMethod(ctx, accountID, paymentMethodID, raiseOnError, responseFormat...)
}

// GetAccountNote is the module form of Accounts.GetAccountNote.
func GetAccountNote(requester exhttp.Requester, ctx context.Context, accountID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountNote(ctx, accountID, noteID, raiseOnError, responseFormat...)
}

// GetAccountNoteFiles is the module form of Accounts.GetAccountNoteFiles.
func GetAccountNoteFiles(requester exhttp.Requester, ctx context.Context, accountID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountNoteFiles(ctx, accountID, noteID, raiseOnError, responseFormat...)
}

// GetAccountNoteFile is the module form of Accounts.GetAccountNoteFile.
func GetAccountNoteFile(requester exhttp.Requester, ctx context.Context, accountID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountNoteFile(ctx, accountID, noteID, fileID, raiseOnError, responseFormat...)
}

// GetAccountEmail is the module form of Accounts.GetAccountEmail.
func GetAccountEmail(requester exhttp.Requester, ctx context.Context, accountID any, emailID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountEmail(ctx, accountID, emailID, raiseOnError, responseFormat...)
}

// GetAccountNotes is the module form of Accounts.GetAccountNotes.
func GetAccountNotes(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountNotes(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountEmails is the module form of Accounts.GetAccountEmails.
func GetAccountEmails(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountEmails(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountOrders is the module form of Accounts.GetAccountOrders.
func GetAccountOrders(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountOrders(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountInvoices is the module form of Accounts.GetAccountInvoices.
func GetAccountInvoices(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountInvoices(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountPayments is the module form of Accounts.GetAccountPayments.
func GetAccountPayments(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountPayments(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountRefunds is the module form of Accounts.GetAccountRefunds.
func GetAccountRefunds(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountRefunds(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountCreditNotes is the module form of Accounts.GetAccountCreditNotes.
func GetAccountCreditNotes(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountCreditNotes(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountPurchaseOrders is the module form of Accounts.GetAccountPurchaseOrders.
func GetAccountPurchaseOrders(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountPurchaseOrders(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountPurchaseInvoices is the module form of Accounts.GetAccountPurchaseInvoices.
func GetAccountPurchaseInvoices(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountPurchaseInvoices(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountPurchasePayments is the module form of Accounts.GetAccountPurchasePayments.
func GetAccountPurchasePayments(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountPurchasePayments(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountPurchaseRefunds is the module form of Accounts.GetAccountPurchaseRefunds.
func GetAccountPurchaseRefunds(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountPurchaseRefunds(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountCustomObjects is the module form of Accounts.GetAccountCustomObjects.
func GetAccountCustomObjects(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountCustomObjects(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// CancelAccount is the module form of Accounts.CancelAccount.
func CancelAccount(requester exhttp.Requester, ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CancelAccount(ctx, accountID, payload, raiseOnError, responseFormat...)
}

// SetAccountPassword is the module form of Accounts.SetAccountPassword.
func SetAccountPassword(requester exhttp.Requester, ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).SetAccountPassword(ctx, accountID, payload, raiseOnError, responseFormat...)
}

// UpdateAccountInformation is the module form of Accounts.UpdateAccountInformation.
func UpdateAccountInformation(requester exhttp.Requester, ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateAccountInformation(ctx, accountID, payload, raiseOnError, responseFormat...)
}

// CreateAccountAddresses is the module form of Accounts.CreateAccountAddresses.
func CreateAccountAddresses(requester exhttp.Requester, ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateAccountAddresses(ctx, accountID, payload, raiseOnError, responseFormat...)
}

// UpdateAccountAddress is the module form of Accounts.UpdateAccountAddress.
func UpdateAccountAddress(requester exhttp.Requester, ctx context.Context, accountID any, addressID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateAccountAddress(ctx, accountID, addressID, payload, raiseOnError, responseFormat...)
}

// UpdateAccountContact is the module form of Accounts.UpdateAccountContact.
func UpdateAccountContact(requester exhttp.Requester, ctx context.Context, accountID any, contactType any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateAccountContact(ctx, accountID, contactType, payload, raiseOnError, responseFormat...)
}

// ReplaceAccountContact is the module form of Accounts.ReplaceAccountContact.
func ReplaceAccountContact(requester exhttp.Requester, ctx context.Context, accountID any, contactType any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReplaceAccountContact(ctx, accountID, contactType, payload, raiseOnError, responseFormat...)
}

// ReplaceAccountBillingPreferences is the module form of Accounts.ReplaceAccountBillingPreferences.
func ReplaceAccountBillingPreferences(requester exhttp.Requester, ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReplaceAccountBillingPreferences(ctx, accountID, payload, raiseOnError, responseFormat...)
}

// UpdateAccountPaymentMethod is the module form of Accounts.UpdateAccountPaymentMethod.
func UpdateAccountPaymentMethod(requester exhttp.Requester, ctx context.Context, accountID any, paymentMethodID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateAccountPaymentMethod(ctx, accountID, paymentMethodID, payload, raiseOnError, responseFormat...)
}

// CreateAccountNote is the module form of Accounts.CreateAccountNote.
func CreateAccountNote(requester exhttp.Requester, ctx context.Context, accountID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateAccountNote(ctx, accountID, note, files, raiseOnError, responseFormat...)
}

// UpdateAccountNote is the module form of Accounts.UpdateAccountNote.
func UpdateAccountNote(requester exhttp.Requester, ctx context.Context, accountID any, noteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateAccountNote(ctx, accountID, noteID, note, files, raiseOnError, responseFormat...)
}

// UploadAccountImage is the module form of Accounts.UploadAccountImage.
func UploadAccountImage(requester exhttp.Requester, ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UploadAccountImage(ctx, accountID, payload, raiseOnError, responseFormat...)
}

// DeleteAccount is the module form of Accounts.DeleteAccount.
func DeleteAccount(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteAccount(ctx, accountID, raiseOnError, responseFormat...)
}

// DeleteAccountContact is the module form of Accounts.DeleteAccountContact.
func DeleteAccountContact(requester exhttp.Requester, ctx context.Context, accountID any, contactType any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteAccountContact(ctx, accountID, contactType, raiseOnError, responseFormat...)
}

// DeleteAccountAddress is the module form of Accounts.DeleteAccountAddress.
func DeleteAccountAddress(requester exhttp.Requester, ctx context.Context, accountID any, addressID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteAccountAddress(ctx, accountID, addressID, raiseOnError, responseFormat...)
}

// DeleteAccountPaymentMethod is the module form of Accounts.DeleteAccountPaymentMethod.
func DeleteAccountPaymentMethod(requester exhttp.Requester, ctx context.Context, accountID any, paymentMethodID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteAccountPaymentMethod(ctx, accountID, paymentMethodID, raiseOnError, responseFormat...)
}

// DeleteAccountNote is the module form of Accounts.DeleteAccountNote.
func DeleteAccountNote(requester exhttp.Requester, ctx context.Context, accountID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteAccountNote(ctx, accountID, noteID, raiseOnError, responseFormat...)
}

// DeleteAccountImage is the module form of Accounts.DeleteAccountImage.
func DeleteAccountImage(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteAccountImage(ctx, accountID, raiseOnError, responseFormat...)
}

// CreateAccount is the module form of Accounts.CreateAccount.
func CreateAccount(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateAccount(ctx, payload, raiseOnError, responseFormat...)
}

// ReplaceAccount is the module form of Accounts.ReplaceAccount.
func ReplaceAccount(requester exhttp.Requester, ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReplaceAccount(ctx, accountID, payload, raiseOnError, responseFormat...)
}

// ReactivateAccount is the module form of Accounts.ReactivateAccount.
func ReactivateAccount(requester exhttp.Requester, ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReactivateAccount(ctx, accountID, payload, raiseOnError, responseFormat...)
}

// CreateAccountContacts is the module form of Accounts.CreateAccountContacts.
func CreateAccountContacts(requester exhttp.Requester, ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateAccountContacts(ctx, accountID, payload, raiseOnError, responseFormat...)
}

// ReplaceAccountAddress is the module form of Accounts.ReplaceAccountAddress.
func ReplaceAccountAddress(requester exhttp.Requester, ctx context.Context, accountID any, addressID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReplaceAccountAddress(ctx, accountID, addressID, payload, raiseOnError, responseFormat...)
}

// CreateAccountPaymentMethod is the module form of Accounts.CreateAccountPaymentMethod.
func CreateAccountPaymentMethod(requester exhttp.Requester, ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateAccountPaymentMethod(ctx, accountID, payload, raiseOnError, responseFormat...)
}

// ReplaceAccountPaymentMethod is the module form of Accounts.ReplaceAccountPaymentMethod.
func ReplaceAccountPaymentMethod(requester exhttp.Requester, ctx context.Context, accountID any, paymentMethodID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReplaceAccountPaymentMethod(ctx, accountID, paymentMethodID, payload, raiseOnError, responseFormat...)
}

// CreateAccountCustomObjectInstance is the module form of Accounts.CreateAccountCustomObjectInstance.
func CreateAccountCustomObjectInstance(requester exhttp.Requester, ctx context.Context, accountID any, customObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateAccountCustomObjectInstance(ctx, accountID, customObjectID, payload, raiseOnError, responseFormat...)
}

// UpdateAccountCustomObjectInstance is the module form of Accounts.UpdateAccountCustomObjectInstance.
func UpdateAccountCustomObjectInstance(requester exhttp.Requester, ctx context.Context, accountID any, customObjectInstanceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateAccountCustomObjectInstance(ctx, accountID, customObjectInstanceID, payload, raiseOnError, responseFormat...)
}

// SendAccountEmail is the module form of Accounts.SendAccountEmail.
func SendAccountEmail(requester exhttp.Requester, ctx context.Context, accountID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).SendAccountEmail(ctx, accountID, payload, raiseOnError, responseFormat...)
}

// PurgeAccount is the module form of Accounts.PurgeAccount.
func PurgeAccount(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).PurgeAccount(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountKPIs is the module form of Accounts.GetAccountKPIs.
func GetAccountKPIs(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountKPIs(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountFinancialSummary is the module form of Accounts.GetAccountFinancialSummary.
func GetAccountFinancialSummary(requester exhttp.Requester, ctx context.Context, accountID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountFinancialSummary(ctx, accountID, raiseOnError, responseFormat...)
}

// GetAccountCustomObject is the module form of Accounts.GetAccountCustomObject.
func GetAccountCustomObject(requester exhttp.Requester, ctx context.Context, accountID any, customObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountCustomObject(ctx, accountID, customObjectID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountCustomObjectInstance is the module form of Accounts.GetAccountCustomObjectInstance.
func GetAccountCustomObjectInstance(requester exhttp.Requester, ctx context.Context, accountID any, customObjectID any, customObjectInstanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountCustomObjectInstance(ctx, accountID, customObjectID, customObjectInstanceID, raiseOnError, responseFormat...)
}

// GetAccountComponentObjects is the module form of Accounts.GetAccountComponentObjects.
func GetAccountComponentObjects(requester exhttp.Requester, ctx context.Context, accountID any, componentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountComponentObjects(ctx, accountID, componentID, queryParams, raiseOnError, responseFormat...)
}

// GetAccountComponentObject is the module form of Accounts.GetAccountComponentObject.
func GetAccountComponentObject(requester exhttp.Requester, ctx context.Context, accountID any, componentID any, componentObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAccountComponentObject(ctx, accountID, componentID, componentObjectID, raiseOnError, responseFormat...)
}

// DeleteAccountNoteFile is the module form of Accounts.DeleteAccountNoteFile.
func DeleteAccountNoteFile(requester exhttp.Requester, ctx context.Context, accountID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteAccountNoteFile(ctx, accountID, noteID, fileID, raiseOnError, responseFormat...)
}
