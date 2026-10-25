// Package accountemails is the v4 account emails endpoints.
package accountemails

// The path templates for the account emails endpoints; the client fills {api_version}.
const (
	AccountCommunicationsEmailPath = "/api/{api_version}/accounts/{account_id}/communications/emails/{email_uuid}"
)

// AccountEmailsKey is the key a listing's records arrive under, the v4 list envelope's data.
const AccountEmailsKey = "data"

// EmailKey is the envelope key one record arrives under.
const EmailKey = "email"
