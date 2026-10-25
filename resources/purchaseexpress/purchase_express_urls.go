// Package purchaseexpress is the v4 purchase express endpoints.
package purchaseexpress

// The path templates for the purchase express endpoints; the client fills {api_version}.
const (
	PurchaseExpressPath = "/api/{api_version}/purchase_express"
)

// PurchaseExpressKey is the envelope key one record arrives under.
const PurchaseExpressKey = "purchase_express"
