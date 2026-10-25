package inventoryadjustments

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// InventoryAdjustments is the v4 inventory adjustments endpoints, bound to one client.
type InventoryAdjustments struct {
	requester exhttp.Requester
}

// GetInventoryAdjustments sends GET /api/{api_version}/inventory_adjustments.
func (i *InventoryAdjustments) GetInventoryAdjustments(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InventoryAdjustmentsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InventoryAdjustmentListShape})
}

// GetInventoryAdjustment sends GET /api/{api_version}/inventory_adjustments/{inventory_adjustment_id}.
func (i *InventoryAdjustments) GetInventoryAdjustment(ctx context.Context, inventoryAdjustmentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(inventoryAdjustmentID, "inventory_adjustment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InventoryAdjustmentPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"inventory_adjustment_id": inventoryAdjustmentID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InventoryAdjustmentRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "inventory_adjustments"

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

// New binds the inventory adjustments endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *InventoryAdjustments {
	return &InventoryAdjustments{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (i *InventoryAdjustments) Requester() exhttp.Requester { return i.requester }
