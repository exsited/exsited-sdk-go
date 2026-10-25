package items

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Items method of the same name,
// which is the primary surface.

// GetItems is the module form of Items.GetItems.
func GetItems(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItems(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetItem is the module form of Items.GetItem.
func GetItem(requester exhttp.Requester, ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItem(ctx, itemID, raiseOnError, responseFormat...)
}

// GetItemInformation is the module form of Items.GetItemInformation.
func GetItemInformation(requester exhttp.Requester, ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemInformation(ctx, itemID, raiseOnError, responseFormat...)
}

// CreateItem is the module form of Items.CreateItem.
func CreateItem(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateItem(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateItemInformation is the module form of Items.UpdateItemInformation.
func UpdateItemInformation(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateItemInformation(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// DeactivateItem is the module form of Items.DeactivateItem.
func DeactivateItem(requester exhttp.Requester, ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeactivateItem(ctx, itemID, raiseOnError, responseFormat...)
}

// ReactivateItem is the module form of Items.ReactivateItem.
func ReactivateItem(requester exhttp.Requester, ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReactivateItem(ctx, itemID, raiseOnError, responseFormat...)
}

// DeleteItem is the module form of Items.DeleteItem.
func DeleteItem(requester exhttp.Requester, ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteItem(ctx, itemID, raiseOnError, responseFormat...)
}

// UploadItemImage is the module form of Items.UploadItemImage.
func UploadItemImage(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UploadItemImage(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// GetItemSale is the module form of Items.GetItemSale.
func GetItemSale(requester exhttp.Requester, ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemSale(ctx, itemID, raiseOnError, responseFormat...)
}

// CreateItemSale is the module form of Items.CreateItemSale.
func CreateItemSale(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateItemSale(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// UpdateItemSale is the module form of Items.UpdateItemSale.
func UpdateItemSale(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateItemSale(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// GetItemPurchase is the module form of Items.GetItemPurchase.
func GetItemPurchase(requester exhttp.Requester, ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemPurchase(ctx, itemID, raiseOnError, responseFormat...)
}

// CreateItemPurchase is the module form of Items.CreateItemPurchase.
func CreateItemPurchase(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateItemPurchase(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// UpdateItemPurchase is the module form of Items.UpdateItemPurchase.
func UpdateItemPurchase(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateItemPurchase(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// GetItemInventories is the module form of Items.GetItemInventories.
func GetItemInventories(requester exhttp.Requester, ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemInventories(ctx, itemID, raiseOnError, responseFormat...)
}

// UpdateItemInventories is the module form of Items.UpdateItemInventories.
func UpdateItemInventories(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateItemInventories(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// AddItemInventory is the module form of Items.AddItemInventory.
func AddItemInventory(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).AddItemInventory(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// GetItemInventoryAdjustments is the module form of Items.GetItemInventoryAdjustments.
func GetItemInventoryAdjustments(requester exhttp.Requester, ctx context.Context, itemID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemInventoryAdjustments(ctx, itemID, queryParams, raiseOnError, responseFormat...)
}

// GetItemWarehouseInventories is the module form of Items.GetItemWarehouseInventories.
func GetItemWarehouseInventories(requester exhttp.Requester, ctx context.Context, itemID any, warehouseName any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemWarehouseInventories(ctx, itemID, warehouseName, raiseOnError, responseFormat...)
}

// UpdateItemBundleDetails is the module form of Items.UpdateItemBundleDetails.
func UpdateItemBundleDetails(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateItemBundleDetails(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// UpdateItemVariations is the module form of Items.UpdateItemVariations.
func UpdateItemVariations(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateItemVariations(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// GetItemFamilyDetails is the module form of Items.GetItemFamilyDetails.
func GetItemFamilyDetails(requester exhttp.Requester, ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemFamilyDetails(ctx, itemID, raiseOnError, responseFormat...)
}

// CreateItemFamilyDetails is the module form of Items.CreateItemFamilyDetails.
func CreateItemFamilyDetails(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateItemFamilyDetails(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// GetItemPackageDetails is the module form of Items.GetItemPackageDetails.
func GetItemPackageDetails(requester exhttp.Requester, ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemPackageDetails(ctx, itemID, raiseOnError, responseFormat...)
}

// GetItemPackageDetail is the module form of Items.GetItemPackageDetail.
func GetItemPackageDetail(requester exhttp.Requester, ctx context.Context, itemID any, packageDetailID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemPackageDetail(ctx, itemID, packageDetailID, raiseOnError, responseFormat...)
}

// CreateItemPackageDetails is the module form of Items.CreateItemPackageDetails.
func CreateItemPackageDetails(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateItemPackageDetails(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// DeleteItemPackageDetail is the module form of Items.DeleteItemPackageDetail.
func DeleteItemPackageDetail(requester exhttp.Requester, ctx context.Context, itemID any, packageDetailID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteItemPackageDetail(ctx, itemID, packageDetailID, raiseOnError, responseFormat...)
}

// GetItemNotes is the module form of Items.GetItemNotes.
func GetItemNotes(requester exhttp.Requester, ctx context.Context, itemID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemNotes(ctx, itemID, queryParams, raiseOnError, responseFormat...)
}

// GetItemNote is the module form of Items.GetItemNote.
func GetItemNote(requester exhttp.Requester, ctx context.Context, itemID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemNote(ctx, itemID, noteID, raiseOnError, responseFormat...)
}

// CreateItemNote is the module form of Items.CreateItemNote.
func CreateItemNote(requester exhttp.Requester, ctx context.Context, itemID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateItemNote(ctx, itemID, note, files, raiseOnError, responseFormat...)
}

// UpdateItemNote is the module form of Items.UpdateItemNote.
func UpdateItemNote(requester exhttp.Requester, ctx context.Context, itemID any, noteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateItemNote(ctx, itemID, noteID, note, files, raiseOnError, responseFormat...)
}

// DeleteItemNote is the module form of Items.DeleteItemNote.
func DeleteItemNote(requester exhttp.Requester, ctx context.Context, itemID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteItemNote(ctx, itemID, noteID, raiseOnError, responseFormat...)
}

// GetItemNoteFiles is the module form of Items.GetItemNoteFiles.
func GetItemNoteFiles(requester exhttp.Requester, ctx context.Context, itemID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemNoteFiles(ctx, itemID, noteID, raiseOnError, responseFormat...)
}

// GetItemNoteFile is the module form of Items.GetItemNoteFile.
func GetItemNoteFile(requester exhttp.Requester, ctx context.Context, itemID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemNoteFile(ctx, itemID, noteID, fileID, raiseOnError, responseFormat...)
}

// CreateItemInventories is the module form of Items.CreateItemInventories.
func CreateItemInventories(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateItemInventories(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// AddItemInventories is the module form of Items.AddItemInventories.
func AddItemInventories(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).AddItemInventories(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// CreateItemInventory is the module form of Items.CreateItemInventory.
func CreateItemInventory(requester exhttp.Requester, ctx context.Context, itemID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateItemInventory(ctx, itemID, payload, raiseOnError, responseFormat...)
}

// DeleteItemNoteFile is the module form of Items.DeleteItemNoteFile.
func DeleteItemNoteFile(requester exhttp.Requester, ctx context.Context, itemID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteItemNoteFile(ctx, itemID, noteID, fileID, raiseOnError, responseFormat...)
}

// UpdateItemPackageDetail is the module form of Items.UpdateItemPackageDetail.
func UpdateItemPackageDetail(requester exhttp.Requester, ctx context.Context, itemID any, packageDetailID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateItemPackageDetail(ctx, itemID, packageDetailID, payload, raiseOnError, responseFormat...)
}

// GetItemKPIs is the module form of Items.GetItemKPIs.
func GetItemKPIs(requester exhttp.Requester, ctx context.Context, itemID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemKPIs(ctx, itemID, raiseOnError, responseFormat...)
}
