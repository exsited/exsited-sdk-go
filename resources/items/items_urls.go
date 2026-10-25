// Package items is the v4 items endpoints.
package items

// The path templates for the items endpoints; the client fills {api_version}.
const (
	ItemsPath                    = "/api/{api_version}/items"
	ItemPath                     = "/api/{api_version}/items/{item_id}"
	ItemInformationPath          = "/api/{api_version}/items/{item_id}/information"
	ItemDeactivatePath           = "/api/{api_version}/items/{item_id}/deactivate"
	ItemReactivatePath           = "/api/{api_version}/items/{item_id}/reactivate"
	ItemImagePath                = "/api/{api_version}/items/{item_id}/image"
	ItemSalePath                 = "/api/{api_version}/items/{item_id}/sale"
	ItemPurchasePath             = "/api/{api_version}/items/{item_id}/purchase"
	ItemInventoriesPath          = "/api/{api_version}/items/{item_id}/inventories"
	ItemInventoryAddPath         = "/api/{api_version}/items/{item_id}/inventory/add"
	ItemInventoryAdjustmentsPath = "/api/{api_version}/items/{item_id}/inventory_adjustments"
	ItemWarehouseInventoriesPath = "/api/{api_version}/items/{item_id}/warehouses/{warehouse_name}/inventories"
	ItemBundleDetailsPath        = "/api/{api_version}/items/{item_id}/bundle_details"
	ItemVariationsPath           = "/api/{api_version}/items/{item_id}/variations"
	ItemFamilyDetailsPath        = "/api/{api_version}/items/{item_id}/family_details"
	ItemPackageDetailsPath       = "/api/{api_version}/items/{item_id}/package_details"
	ItemPackageDetailPath        = "/api/{api_version}/items/{item_id}/package_details/{package_detail_id}"
	ItemNotesPath                = "/api/{api_version}/items/{item_id}/notes"
	ItemNotePath                 = "/api/{api_version}/items/{item_id}/notes/{note_id}"
	ItemNoteFilesPath            = "/api/{api_version}/items/{item_id}/notes/{note_id}/files"
	ItemNoteFilePath             = "/api/{api_version}/items/{item_id}/notes/{note_id}/files/{file_id}"
	ItemInventoriesAddPath       = "/api/{api_version}/items/{item_id}/inventories/add"
	ItemInventoryPath            = "/api/{api_version}/items/{item_id}/inventory"
	ItemKPIsPath                 = "/api/{api_version}/items/{item_id}/kpis"
)

// ItemsKey is the key a listing's records arrive under, the v4 list envelope's data.
const ItemsKey = "data"

// InventoryAdjustmentKey is the envelope key one record arrives under.
const InventoryAdjustmentKey = "inventory_adjustment"

// ItemKey is the envelope key one record arrives under.
const ItemKey = "item"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"

// itemsQueryParams are the names the listing takes besides its filter_<field> filters.
var itemsQueryParams = [5]string{"limit", "cursor", "order_by", "direction", "filter"}

// ItemsQueryParams returns a fresh copy of the names the listing takes besides its filters.
func ItemsQueryParams() []string {
	return append(make([]string, 0, len(itemsQueryParams)), itemsQueryParams[:]...)
}

// FilterableFields are the top-level item fields that filter_<field> accepts.
// Every one also takes _neq; the text ones additionally take _contains and
// _not_contains.
var FilterableFields = map[string]bool{
	"id": true, "uuid": true, "name": true, "display_name": true,
	"description": true, "status": true, "type": true, "origin": true,
	"invoice_note": true, "image_uri": true, "base_uom": true, "group": true,
	"manager": true, "custom_form": true, "created_by": true,
	"last_updated_by": true, "version": true,
}

// DeprecatedFilters are the fields documented as filter_<field> that the API
// rejects with 400 "Filter syntax deprecated bracket". Each lives inside a nested
// block and is reachable through the filter expression instead.
var DeprecatedFilters = map[string]bool{
	"item_number": true, "upc_code": true,
	"sale_enabled": true, "sales_accounting_code_revenue": true,
	"sales_billing_mode": true, "sales_charge_type": true,
	"sales_currency": true, "sales_fulfillment_mode": true,
	"sales_payment_mode": true, "sales_pricing_method": true,
	"sales_tax_inclusive_based_on": true,
	"purchase_enabled":             true, "purchase_currency": true,
	"purchase_order_note": true, "purchase_profile": true,
	"purchase_receive_mode": true,
	"inventory_enabled":     true, "inventory_quantity_available": true,
	"inventory_quantity_on_hand":              true,
	"inventory_qty_avl_on_sale_determination": true,
}

// FileKey is the envelope key one record arrives under.
const FileKey = "file"
