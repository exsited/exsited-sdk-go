package itemreceipts

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// ItemReceipts is the v4 item receipts endpoints, bound to one client.
type ItemReceipts struct {
	requester exhttp.Requester
}

// GetItemReceipts sends GET /api/{api_version}/item_receipts.
func (i *ItemReceipts) GetItemReceipts(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemReceiptsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ItemReceiptListShape})
}

// CreateItemReceipt sends POST /api/{api_version}/item_receipts.
func (i *ItemReceipts) CreateItemReceipt(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, ItemReceiptsPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ItemReceiptsRecordShape})
}

// UpdateItemReceipt sends PATCH /api/{api_version}/item_receipts/{item_receipt_id}.
func (i *ItemReceipts) UpdateItemReceipt(ctx context.Context, itemReceiptID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemReceiptID, "item_receipt_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, ItemReceiptPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_receipt_id": itemReceiptID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ItemReceiptUpdatedRecordShape})
}

// GetItemReceipt sends GET /api/{api_version}/item_receipts/{item_receipt_id}.
func (i *ItemReceipts) GetItemReceipt(ctx context.Context, itemReceiptID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemReceiptID, "item_receipt_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, ItemReceiptPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"item_receipt_id": itemReceiptID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ItemReceiptRecordShape})
}

// GetPurchaseOrderItemReceipts sends GET /api/{api_version}/purchase_orders/{purchase_order_id}/item_receipts.
func (i *ItemReceipts) GetPurchaseOrderItemReceipts(ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, PurchaseOrderItemReceiptsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderItemReceiptsRecordShape})
}

// DeleteItemReceipt sends DELETE /api/{api_version}/item_receipts/{item_receipt_id}; it has no typed view, so read ApiResponse.Map.
func (i *ItemReceipts) DeleteItemReceipt(ctx context.Context, itemReceiptID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemReceiptID, "item_receipt_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendDeleteRequest(ctx, ItemReceiptPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"item_receipt_id": itemReceiptID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteItemReceiptNote sends DELETE /api/{api_version}/item_receipts/{item_receipt_id}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (i *ItemReceipts) DeleteItemReceiptNote(ctx context.Context, itemReceiptID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(itemReceiptID, "item_receipt_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendDeleteRequest(ctx, ItemReceiptNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"item_receipt_id": itemReceiptID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "item_receipts"

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

// requireValue refuses a path value that is nil, empty or only whitespace, before any I/O.
func requireValue(value any, what string) error {
	text, isText := value.(string)
	if value == nil || isNilValue(value) || isText && strings.TrimSpace(text) == "" {
		return exerrors.UsageErrorf("%s is required: the request cannot be built without it", what)
	}
	return nil
}

// New binds the item receipts endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *ItemReceipts { return &ItemReceipts{requester: requester} }

// Requester returns the requester this resource was bound to.
func (i *ItemReceipts) Requester() exhttp.Requester { return i.requester }
