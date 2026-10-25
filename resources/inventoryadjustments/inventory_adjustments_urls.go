// Package inventoryadjustments is the v4 inventory adjustments endpoints.
package inventoryadjustments

// The path templates for the inventory adjustments endpoints; the client fills {api_version}.
const (
	InventoryAdjustmentsPath = "/api/{api_version}/inventory_adjustments"
	InventoryAdjustmentPath  = "/api/{api_version}/inventory_adjustments/{inventory_adjustment_id}"
)

// InventoryAdjustmentsKey is the key a listing's records arrive under, the v4 list envelope's data.
const InventoryAdjustmentsKey = "data"

// InventoryAdjustmentKey is the envelope key one record arrives under.
const InventoryAdjustmentKey = "inventory_adjustment"
