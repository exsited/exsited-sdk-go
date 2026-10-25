package itemfulfillments

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ItemFulfillments method of the same name,
// which is the primary surface.

// GetItemFulfillments is the module form of ItemFulfillments.GetItemFulfillments.
func GetItemFulfillments(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemFulfillments(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateItemFulfillment is the module form of ItemFulfillments.CreateItemFulfillment.
func CreateItemFulfillment(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateItemFulfillment(ctx, payload, raiseOnError, responseFormat...)
}

// GetItemFulfillment is the module form of ItemFulfillments.GetItemFulfillment.
func GetItemFulfillment(requester exhttp.Requester, ctx context.Context, fulfillmentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemFulfillment(ctx, fulfillmentID, queryParams, raiseOnError, responseFormat...)
}

// UpdateItemFulfillment is the module form of ItemFulfillments.UpdateItemFulfillment.
func UpdateItemFulfillment(requester exhttp.Requester, ctx context.Context, fulfillmentID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateItemFulfillment(ctx, fulfillmentID, payload, raiseOnError, responseFormat...)
}

// CreateItemFulfillmentNote is the module form of ItemFulfillments.CreateItemFulfillmentNote.
func CreateItemFulfillmentNote(requester exhttp.Requester, ctx context.Context, fulfillmentID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateItemFulfillmentNote(ctx, fulfillmentID, note, files, raiseOnError, responseFormat...)
}

// DeleteItemFulfillment is the module form of ItemFulfillments.DeleteItemFulfillment.
func DeleteItemFulfillment(requester exhttp.Requester, ctx context.Context, fulfillmentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteItemFulfillment(ctx, fulfillmentID, raiseOnError, responseFormat...)
}

// DeleteItemFulfillmentNote is the module form of ItemFulfillments.DeleteItemFulfillmentNote.
func DeleteItemFulfillmentNote(requester exhttp.Requester, ctx context.Context, fulfillmentID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteItemFulfillmentNote(ctx, fulfillmentID, noteID, raiseOnError, responseFormat...)
}

// GetItemFulfillmentNotes is the module form of ItemFulfillments.GetItemFulfillmentNotes.
func GetItemFulfillmentNotes(requester exhttp.Requester, ctx context.Context, fulfillmentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemFulfillmentNotes(ctx, fulfillmentID, queryParams, raiseOnError, responseFormat...)
}

// GetItemFulfillmentNote is the module form of ItemFulfillments.GetItemFulfillmentNote.
func GetItemFulfillmentNote(requester exhttp.Requester, ctx context.Context, fulfillmentID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetItemFulfillmentNote(ctx, fulfillmentID, noteID, raiseOnError, responseFormat...)
}
