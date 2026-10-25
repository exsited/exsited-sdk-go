package items

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Items is the v4 items endpoints, bound to one client.
type Items struct {
	requester exhttp.Requester
}

// GetItems lists items, with optional filtering, sorting and paging: GET /api/{api_version}/items.
//
// Four parameters control the listing: limit, cursor, and the order_by +
// direction pair; build the last two with SortParams. Everything else is a
// filter: top-level fields take filter_<field>, built with ItemFilterParam;
// nested fields take the single filter expression, e.g.
// {"filter": "items/sale(pricing/latest_used_pricing_version) eq '1'"}.
func (i *Items) GetItems(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ItemListShape})
}

// GetItem retrieves one item: GET /api/{api_version}/items/{item_id}.
//
// itemID is the item's id as a listing returns it, "ITEM-0001". Not the uuid,
// and not the integer id the sale, purchase and inventories reads hand back.
func (i *Items) GetItem(ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: ItemRecordShape})
}

// GetItemInformation retrieves an item's core information. Pairs with
// UpdateItemInformation.
//
// GET /api/{api_version}/items/{item_id}/information.
func (i *Items) GetItemInformation(ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemInformationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: ItemRecordShape})
}

// CreateItem creates an item: POST /api/{api_version}/items.
//
// {"item": {"type": "STANDARD", "name": "Widget"}} is the minimum that works. An
// unknown group is a 400 with code internal_error and the message "Invalid group",
// a validation failure wearing a server fault's name.
func (i *Items) CreateItem(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemsPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ItemRecordShape})
}

// UpdateItemInformation updates an item's core information. Partial: send only
// what changes.
//
// PATCH /api/{api_version}/items/{item_id}/information.
func (i *Items) UpdateItemInformation(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, ItemInformationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: ItemRecordShape})
}

// DeactivateItem deactivates an item. Takes no body; answers the whole item.
//
// POST /api/{api_version}/items/{item_id}/deactivate.
func (i *Items) DeactivateItem(ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemDeactivatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// ReactivateItem reactivates a deactivated item. Takes no body; answers the whole
// item.
//
// POST /api/{api_version}/items/{item_id}/reactivate.
func (i *Items) ReactivateItem(ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemReactivatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// DeleteItem deletes an item: DELETE /api/{api_version}/items/{item_id}.
//
// A 409 with code internal_error means the item is referenced: "All Orders with
// this Item need to be deleted". A business rule, not a server fault despite the
// code, and not retryable.
func (i *Items) DeleteItem(ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendDeleteRequest(ctx, ItemPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// UploadItemImage sets the item's image. {"image_base64": "..."} at the top
// level, the one write in this family with no item envelope. Answers the whole
// item.
//
// POST /api/{api_version}/items/{item_id}/image.
func (i *Items) UploadItemImage(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemImagePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// GetItemSale retrieves an item's sale configuration: GET /api/{api_version}/items/{item_id}/sale.
//
// A stub, not the item: {"item": {"id": 2, "name": "...", "sale": {...}}}, and id
// is an integer, not the string code the path took.
func (i *Items) GetItemSale(ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemSalePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: ItemSaleViewShape})
}

// CreateItemSale creates an item's sale configuration.
// {"item": {"sale": {...}, "currencies": [...]}}. Answers 409 already_exists when
// the item already has one.
//
// POST /api/{api_version}/items/{item_id}/sale.
func (i *Items) CreateItemSale(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemSalePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// UpdateItemSale updates an item's sale configuration. {"item": {"sale": {...}}},
// partial.
//
// PATCH /api/{api_version}/items/{item_id}/sale.
func (i *Items) UpdateItemSale(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, ItemSalePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// GetItemPurchase retrieves an item's purchase configuration. A stub, as the sale
// read is.
//
// GET /api/{api_version}/items/{item_id}/purchase.
func (i *Items) GetItemPurchase(ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemPurchasePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: ItemPurchaseViewShape})
}

// CreateItemPurchase creates an item's purchase configuration. Answers
// 409 already_exists when one exists.
//
// POST /api/{api_version}/items/{item_id}/purchase.
func (i *Items) CreateItemPurchase(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemPurchasePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// UpdateItemPurchase updates an item's purchase configuration: PATCH /api/{api_version}/items/{item_id}/purchase.
func (i *Items) UpdateItemPurchase(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, ItemPurchasePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// GetItemInventories retrieves an item's inventory configuration and stock: GET /api/{api_version}/items/{item_id}/inventories.
//
// A stub under item whose id is the string code the path took. The per-warehouse
// quantities inside are numbers, each beside a _value twin that is a string.
func (i *Items) GetItemInventories(ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemInventoriesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: ItemInventoriesViewShape})
}

// UpdateItemInventories updates an item's inventory configuration.
// {"item": {"inventories": {...}}}.
//
// The write spells warehouse management warehouse_enabled where the read returns
// enable_warehouse_management. Turning inventory off once stock exists is refused
// with 400 "Unable to disable Inventory Management".
//
// PATCH /api/{api_version}/items/{item_id}/inventories.
func (i *Items) UpdateItemInventories(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, ItemInventoriesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// AddItemInventory adds stock to a warehouse. The path segment is inventory,
// singular.
//
// The warehouse is keyed name in the body and warehouse on the read side.
//
// POST /api/{api_version}/items/{item_id}/inventory/add.
func (i *Items) AddItemInventory(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemInventoryAddPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// GetItemInventoryAdjustments lists an item's inventory movements: GET /api/{api_version}/items/{item_id}/inventory_adjustments.
//
// Every quantity on these records is a string, "79999", not 79999: the opposite
// of what GetItemInventories returns for the same names.
func (i *Items) GetItemInventoryAdjustments(ctx context.Context, itemID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemInventoryAdjustmentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InventoryAdjustmentListShape})
}

// GetItemWarehouseInventories retrieves one warehouse's stock for an item: GET /api/{api_version}/items/{item_id}/warehouses/{warehouse_name}/inventories.
//
// warehouseName is the warehouse name, not an id: the one slot in this family
// that takes a name. A 404 names it, e.g. "warehouse 'BareHouse' not found".
func (i *Items) GetItemWarehouseInventories(ctx context.Context, itemID any, warehouseName any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemWarehouseInventoriesPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"item_id": itemID, "warehouse_name": warehouseName},
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ItemInventoriesViewShape})
}

// UpdateItemBundleDetails sets the components of a bundle item.
// {"item": {"bundle_details": [{id, quantity}]}}: a bare list, unlike
// family_details, which takes an object with an add key.
//
// PATCH /api/{api_version}/items/{item_id}/bundle_details.
func (i *Items) UpdateItemBundleDetails(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, ItemBundleDetailsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// UpdateItemVariations configures an item's variations. A parent declares
// configure_variations; a child sets parent_item_id and use_variations.
//
// PATCH /api/{api_version}/items/{item_id}/variations.
func (i *Items) UpdateItemVariations(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, ItemVariationsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// GetItemFamilyDetails lists the members of an item family. Each carries item_id
// and item_name only.
//
// GET /api/{api_version}/items/{item_id}/family_details.
func (i *Items) GetItemFamilyDetails(ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemFamilyDetailsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: ItemRecordShape})
}

// CreateItemFamilyDetails adds members to an item family.
// {"item": {"family_details": {"add": ["ITEM-0002"]}}}: an object with an add
// list.
//
// POST /api/{api_version}/items/{item_id}/family_details.
func (i *Items) CreateItemFamilyDetails(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemFamilyDetailsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// GetItemPackageDetails lists an item's packages: GET /api/{api_version}/items/{item_id}/package_details.
func (i *Items) GetItemPackageDetails(ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemPackageDetailsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: ItemRecordShape})
}

// GetItemPackageDetail retrieves one package. packageDetailID is the package's
// uuid, as package_details[].uuid carries it.
//
// GET /api/{api_version}/items/{item_id}/package_details/{package_detail_id}.
func (i *Items) GetItemPackageDetail(ctx context.Context, itemID any, packageDetailID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemPackageDetailPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"item_id": itemID, "package_detail_id": packageDetailID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ItemRecordShape})
}

// CreateItemPackageDetails adds a package to an item.
//
// POST /api/{api_version}/items/{item_id}/package_details.
func (i *Items) CreateItemPackageDetails(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemPackageDetailsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       ItemRecordShape})
}

// DeleteItemPackageDetail deletes one package: DELETE /api/{api_version}/items/{item_id}/package_details/{package_detail_id}.
func (i *Items) DeleteItemPackageDetail(ctx context.Context, itemID any, packageDetailID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendDeleteRequest(ctx, ItemPackageDetailPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"item_id": itemID, "package_detail_id": packageDetailID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetItemNotes lists an item's notes. The list envelope, as on an account: GET /api/{api_version}/items/{item_id}/notes.
func (i *Items) GetItemNotes(ctx context.Context, itemID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: NoteListShape})
}

// GetItemNote retrieves one note. Answers {"note": {...}}: GET /api/{api_version}/items/{item_id}/notes/{note_id}.
func (i *Items) GetItemNote(ctx context.Context, itemID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID, "note_id": noteID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: NoteRecordShape})
}

// CreateItemNote attaches a note to an item: POST /api/{api_version}/items/{item_id}/notes.
//
// A form, not JSON: the text travels as a form field named note, so this takes the
// text directly rather than a payload object. Any files go beside it as multipart
// parts. HTML is accepted and comes back as written.
func (i *Items) CreateItemNote(ctx context.Context, itemID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: NoteRecordShape})
}

// UpdateItemNote updates a note's text. A form, with any files as multipart parts,
// as the create is. Answers 400 "Invalid UUID" when noteID is not one; the message
// names the format, not the field.
//
// PATCH /api/{api_version}/items/{item_id}/notes/{note_id}.
func (i *Items) UpdateItemNote(ctx context.Context, itemID any, noteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, ItemNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID, "note_id": noteID},
		FormData:   map[string]any{"note": note}, Files: files, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: NoteRecordShape})
}

// DeleteItemNote deletes one note. Answers 400 "Invalid note_id" when the id is
// not one.
//
// DELETE /api/{api_version}/items/{item_id}/notes/{note_id}.
func (i *Items) DeleteItemNote(ctx context.Context, itemID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendDeleteRequest(ctx, ItemNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID, "note_id": noteID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetItemNoteFiles lists a note's attachments: a page of file descriptors, not the
// files.
//
// GET /api/{api_version}/items/{item_id}/notes/{note_id}/files.
func (i *Items) GetItemNoteFiles(ctx context.Context, itemID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemNoteFilesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID, "note_id": noteID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: NoteFileListShape})
}

// GetItemNoteFile retrieves one attachment. exhttp.ApiResponse.Map may be bytes for a binary
// attachment, in which case JSON and dto.Dto are both nil; otherwise it answers the
// file's descriptor.
//
// GET /api/{api_version}/items/{item_id}/notes/{note_id}/files/{file_id}.
func (i *Items) GetItemNoteFile(ctx context.Context, itemID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"item_id": itemID, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ItemNoteFileViewShape})
}

// ItemFilterParam builds a filter_* parameter name, refusing one the API rejects.
//
// The same construction as exhttp.FilterParam (filter_<field> for equality and
// filter_<field>_<operator> otherwise), with one addition: a field in
// DeprecatedFilters returns an error here rather than costing a request to be told
// the syntax is deprecated.
func ItemFilterParam(field string, operator ...string) (string, error) {
	if DeprecatedFilters[field] {
		return "", exerrors.NewUsageError(
			"filter field " + field + ": the API answers 400 'Filter syntax " +
				"deprecated bracket' for this one, because it addresses a field " +
				"inside a nested block rather than a top-level item field. Use the " +
				"'filter' expression parameter instead, e.g. " +
				`{"filter": "items/sale(pricing/latest_used_pricing_version) eq '1'"}` +
				" - see GetItems.")
	}
	return exhttp.FilterParam(field, operator...)
}

// MustItemFilterParam is ItemFilterParam for a literal known good at author time.
func MustItemFilterParam(field string, operator ...string) string {
	name, err := ItemFilterParam(field, operator...)
	if err != nil {
		panic(err)
	}
	return name
}

// SortParams builds this family's sort parameters.
//
// Items sort with two parameters, not one. Where an account listing takes
// sort="-created_on", an item listing takes order_by="created_on" and
// direction="desc". exhttp.SortParam builds the account spelling and does not
// apply here; passing it to an item listing is accepted and silently ignored.
//
// The construction is shared with the four other families that take the pair,
// so this is exhttp.SortParams under the name of the family that needs it.
func SortParams(field string, descending bool) (map[string]any, error) {
	return exhttp.SortParams(field, descending)
}

// MustSortParams is SortParams for a literal known good at author time.
func MustSortParams(field string, descending bool) map[string]any {
	params, err := SortParams(field, descending)
	if err != nil {
		panic(err)
	}
	return params
}

// CreateItemInventories sends POST /api/{api_version}/items/{item_id}/inventories.
func (i *Items) CreateItemInventories(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemID, "item_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemInventoriesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ItemRecordShape})
}

// AddItemInventories sends POST /api/{api_version}/items/{item_id}/inventories/add.
func (i *Items) AddItemInventories(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemID, "item_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemInventoriesAddPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ItemRecordShape})
}

// CreateItemInventory sends POST /api/{api_version}/items/{item_id}/inventory.
func (i *Items) CreateItemInventory(ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemID, "item_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemInventoryPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ItemRecordShape})
}

// DeleteItemNoteFile sends DELETE /api/{api_version}/items/{item_id}/notes/{note_id}/files/{file_id}; it has no typed view, so read ApiResponse.Map.
func (i *Items) DeleteItemNoteFile(ctx context.Context, itemID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemID, "item_id"); err != nil {
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
	return i.requester.SendDeleteRequest(ctx, ItemNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"item_id": itemID, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateItemPackageDetail sends PATCH /api/{api_version}/items/{item_id}/package_details/{package_detail_id}.
func (i *Items) UpdateItemPackageDetail(ctx context.Context, itemID any, packageDetailID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemID, "item_id"); err != nil {
		return nil, err
	}
	if err := requireValue(packageDetailID, "package_detail_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, ItemPackageDetailPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_id": itemID, "package_detail_id": packageDetailID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ItemRecordShape})
}

// GetItemKPIs sends GET /api/{api_version}/items/{item_id}/kpis; it has no typed view, so read ApiResponse.Map.
func (i *Items) GetItemKPIs(ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemID, "item_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemKPIsPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"item_id": itemID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "items"

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

// New binds the items endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Items { return &Items{requester: requester} }

// Requester returns the requester this resource was bound to.
func (i *Items) Requester() exhttp.Requester { return i.requester }
