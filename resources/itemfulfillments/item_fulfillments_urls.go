// Package itemfulfillments is the v4 item fulfillments endpoints.
package itemfulfillments

// The path templates for the item fulfillments endpoints; the client fills {api_version}.
const (
	ItemFulfillmentsPath     = "/api/{api_version}/item_fulfillments"
	ItemFulfillmentPath      = "/api/{api_version}/item_fulfillments/{fulfillment_id}"
	ItemFulfillmentNotesPath = "/api/{api_version}/item_fulfillments/{fulfillment_id}/notes"
	ItemFulfillmentNotePath  = "/api/{api_version}/item_fulfillments/{fulfillment_id}/notes/{note_id}"
)

// ItemFulfillmentsKey is the key a listing's records arrive under, the v4 list envelope's data.
const ItemFulfillmentsKey = "data"

// InvoiceKey is the envelope key one record arrives under.
const InvoiceKey = "invoice"

// ItemFulfillmentKey is the envelope key one record arrives under.
const ItemFulfillmentKey = "item_fulfillment"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"
