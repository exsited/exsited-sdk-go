// Package purchaserefunds is the v4 purchase refunds endpoints.
package purchaserefunds

// The path templates for the purchase refunds endpoints; the client fills {api_version}.
const (
	PurchaseRefundsPath                   = "/api/{api_version}/purchase_refunds"
	PurchaseRefundPath                    = "/api/{api_version}/purchase_refunds/{purchase_refund_id}"
	PurchaseCreditNotePurchaseRefundsPath = "/api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}/purchase_refunds"
)

// PurchaseRefundsKey is the key a listing's records arrive under, the v4 list envelope's data.
const PurchaseRefundsKey = "data"

// PurchaseRefundKey is the envelope key one record arrives under.
const PurchaseRefundKey = "purchase_refund"
