package inventoryadjustments

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the InventoryAdjustments method of the same name,
// which is the primary surface.

// GetInventoryAdjustments is the module form of InventoryAdjustments.GetInventoryAdjustments.
func GetInventoryAdjustments(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInventoryAdjustments(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetInventoryAdjustment is the module form of InventoryAdjustments.GetInventoryAdjustment.
func GetInventoryAdjustment(requester exhttp.Requester, ctx context.Context, inventoryAdjustmentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetInventoryAdjustment(ctx, inventoryAdjustmentID, queryParams, raiseOnError, responseFormat...)
}
