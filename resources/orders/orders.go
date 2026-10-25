package orders

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Orders is the v4 orders endpoints, bound to one client.
type Orders struct {
	requester exhttp.Requester
}

// GetOrders sends GET /api/{api_version}/orders.
func (o *Orders) GetOrders(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrdersPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderListShape})
}

// CreateOrder sends POST /api/{api_version}/orders; it answers 201. A contract needs a recurring or
// metered item on a line.
func (o *Orders) CreateOrder(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrdersPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// GetOrder sends GET /api/{api_version}/orders/{order_id}.
func (o *Orders) GetOrder(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: OrderReadRecordShape})
}

// DeleteOrder sends DELETE /api/{api_version}/orders/{order_id}; it has no typed view, so read ApiResponse.Map.
func (o *Orders) DeleteOrder(ctx context.Context, orderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendDeleteRequest(ctx, OrderPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// ActivateOrder sends POST /api/{api_version}/orders/{order_id}/activate; it has no typed view, so read ApiResponse.Map.
func (o *Orders) ActivateOrder(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderActivatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateOrderActivateDraft sends POST /api/{api_version}/orders/{order_id}/activate_draft.
func (o *Orders) CreateOrderActivateDraft(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderActivateDraftPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// GetOrderBillingPreferences sends GET /api/{api_version}/orders/{order_id}/billing_preferences.
func (o *Orders) GetOrderBillingPreferences(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderBillingPreferencesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderBillingPreferencesRecordShape})
}

// UpdateOrderBillingPreferences sends PATCH /api/{api_version}/orders/{order_id}/billing_preferences.
func (o *Orders) UpdateOrderBillingPreferences(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPatchRequest(ctx, OrderBillingPreferencesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// ReplaceOrderBillingPreferences sends PUT /api/{api_version}/orders/{order_id}/billing_preferences.
func (o *Orders) ReplaceOrderBillingPreferences(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPutRequest(ctx, OrderBillingPreferencesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// CancelOrder sends POST /api/{api_version}/orders/{order_id}/cancel.
func (o *Orders) CancelOrder(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderCancelPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// ChangeOrder sends POST /api/{api_version}/orders/{order_id}/change.
func (o *Orders) ChangeOrder(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderChangePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// CreateOrderConvertToOrder sends POST /api/{api_version}/orders/{order_id}/convert_to_order.
func (o *Orders) CreateOrderConvertToOrder(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderConvertToOrderPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// GetOrderCustomObjects sends GET /api/{api_version}/orders/{order_id}/custom_objects.
func (o *Orders) GetOrderCustomObjects(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderCustomObjectsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountOrdersRecordShape})
}

// GetOrderCustomObjectsInstances sends GET /api/{api_version}/orders/{order_id}/custom_objects_instances.
func (o *Orders) GetOrderCustomObjectsInstances(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderCustomObjectsInstancesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountOrdersRecordShape})
}

// CreateOrderDiscardDraft sends POST /api/{api_version}/orders/{order_id}/discard_draft.
func (o *Orders) CreateOrderDiscardDraft(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderDiscardDraftPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderDiscardDraftCreateRecordShape})
}

// CreateOrderExtendTrial sends POST /api/{api_version}/orders/{order_id}/extend_trial; the new
// window starts after the current trial ends.
func (o *Orders) CreateOrderExtendTrial(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderExtendTrialPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderExtendTrialCreateRecordShape})
}

// GetOrderInformation sends GET /api/{api_version}/orders/{order_id}/information.
func (o *Orders) GetOrderInformation(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderInformationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: OrderInformationRecordShape})
}

// UpdateOrderInformation sends PATCH /api/{api_version}/orders/{order_id}/information.
func (o *Orders) UpdateOrderInformation(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPatchRequest(ctx, OrderInformationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// GetOrderItemFulfillments sends GET /api/{api_version}/orders/{order_id}/item_fulfillments.
func (o *Orders) GetOrderItemFulfillments(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderItemFulfillmentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: AccountOrdersRecordShape})
}

// GetOrderLines sends GET /api/{api_version}/orders/{order_id}/lines.
func (o *Orders) GetOrderLines(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderLinesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderLinesReadRecordShape})
}

// GetOrderNotes sends GET /api/{api_version}/orders/{order_id}/notes.
func (o *Orders) GetOrderNotes(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: OrderNoteListShape})
}

// CreateOrderNote sends POST /api/{api_version}/orders/{order_id}/notes.
// The note travels as the form field note, and any files go beside it as multipart parts.
func (o *Orders) CreateOrderNote(ctx context.Context, orderID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderNoteReadRecordShape})
}

// GetOrderPaymentMethods sends GET /api/{api_version}/orders/{order_id}/payment_methods.
func (o *Orders) GetOrderPaymentMethods(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderPaymentMethodsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderPaymentMethodsReadRecordShape})
}

// GetOrderPDF sends GET /api/{api_version}/orders/{order_id}/pdf.
func (o *Orders) GetOrderPDF(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderPDFPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderPDFReadRecordShape})
}

// ReactivateOrder sends POST /api/{api_version}/orders/{order_id}/reactivate.
func (o *Orders) ReactivateOrder(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderReactivatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// CreateOrderRelinquish sends POST /api/{api_version}/orders/{order_id}/relinquish.
func (o *Orders) CreateOrderRelinquish(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderRelinquishPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// UpdateOrderLine sends PATCH /api/{api_version}/orders/{order_id}/lines/{line_uuid}.
func (o *Orders) UpdateOrderLine(ctx context.Context, orderID any, lineUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(lineUUID, "line_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPatchRequest(ctx, OrderLinePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID, "line_uuid": lineUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// GetOrderLine sends GET /api/{api_version}/orders/{order_id}/lines/{line_uuid}.
func (o *Orders) GetOrderLine(ctx context.Context, orderID any, lineUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(lineUUID, "line_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderLinePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"order_id": orderID, "line_uuid": lineUUID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderLineReadRecordShape})
}

// DeleteOrderNote sends DELETE /api/{api_version}/orders/{order_id}/notes/{order_note_id}; it has no typed view, so read ApiResponse.Map.
func (o *Orders) DeleteOrderNote(ctx context.Context, orderID any, orderNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(orderNoteID, "order_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendDeleteRequest(ctx, OrderNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID, "order_note_id": orderNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetOrderNote sends GET /api/{api_version}/orders/{order_id}/notes/{order_note_id}.
func (o *Orders) GetOrderNote(ctx context.Context, orderID any, orderNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(orderNoteID, "order_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderNotePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"order_id": orderID, "order_note_id": orderNoteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderNoteReadRecordShape})
}

// UpdateOrderNote sends PATCH /api/{api_version}/orders/{order_id}/notes/{order_note_id}.
// The note travels as the form field note, and any files go beside it as multipart parts.
func (o *Orders) UpdateOrderNote(ctx context.Context, orderID any, orderNoteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(orderNoteID, "order_note_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPatchRequest(ctx, OrderNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID, "order_note_id": orderNoteID},
		FormData:   map[string]any{"note": note}, Files: files, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderNoteReadRecordShape})
}

// GetOrderPaymentMethod sends GET /api/{api_version}/orders/{order_id}/payment_methods/{payment_method_ref}.
func (o *Orders) GetOrderPaymentMethod(ctx context.Context, orderID any, paymentMethodRef any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(paymentMethodRef, "payment_method_ref"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderPaymentMethodPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"order_id": orderID, "payment_method_ref": paymentMethodRef},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderPaymentMethodReadRecordShape})
}

// UpdateOrderLineBillingPreferences sends PATCH /api/{api_version}/orders/{order_id}/lines/{line_uuid}/billing_preferences.
func (o *Orders) UpdateOrderLineBillingPreferences(ctx context.Context, orderID any, lineUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(lineUUID, "line_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPatchRequest(ctx, OrderLineBillingPreferencesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID, "line_uuid": lineUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// ReplaceOrderLineBillingPreferences sends PUT /api/{api_version}/orders/{order_id}/lines/{line_uuid}/billing_preferences.
func (o *Orders) ReplaceOrderLineBillingPreferences(ctx context.Context, orderID any, lineUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(lineUUID, "line_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPutRequest(ctx, OrderLineBillingPreferencesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID, "line_uuid": lineUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderReadRecordShape})
}

// GetOrderNoteFiles sends GET /api/{api_version}/orders/{order_id}/notes/{order_note_id}/files.
func (o *Orders) GetOrderNoteFiles(ctx context.Context, orderID any, orderNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(orderNoteID, "order_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderNoteFilesPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"order_id": orderID, "order_note_id": orderNoteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderNoteFileListShape})
}

// GetOrderNoteFile sends GET /api/{api_version}/orders/{order_id}/notes/{order_note_id}/files/{order_file_id}.
func (o *Orders) GetOrderNoteFile(ctx context.Context, orderID any, orderNoteID any, orderFileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(orderNoteID, "order_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(orderFileID, "order_file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderNoteFilePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"order_id": orderID, "order_note_id": orderNoteID, "order_file_id": orderFileID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: OrderNoteFileReadRecordShape})
}

// GetOrdersNotes sends GET /api/{api_version}/orders/notes.
func (o *Orders) GetOrdersNotes(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrdersNotesPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: OrderNoteListShape})
}

// PreviewOrder sends POST /api/{api_version}/orders/preview; it has no typed view, so read ApiResponse.Map.
func (o *Orders) PreviewOrder(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrdersPreviewPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateOrderExpress sends POST /api/{api_version}/orders/express; it has no typed view, so read ApiResponse.Map.
func (o *Orders) CreateOrderExpress(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrdersExpressPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// AcceptOrder sends POST /api/{api_version}/orders/{order_id}/accept; it has no typed view, so read ApiResponse.Map.
func (o *Orders) AcceptOrder(ctx context.Context, orderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderAcceptPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ActivateOrderInvoiceFromDormant sends POST /api/{api_version}/orders/{order_id}/activate_invoice_from_dormant/{invoice_uuid}; it has no typed view, so read ApiResponse.Map.
func (o *Orders) ActivateOrderInvoiceFromDormant(ctx context.Context, orderID any, invoiceUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(invoiceUUID, "invoice_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderActivateInvoiceFromDormantPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID, "invoice_uuid": invoiceUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// PreviewOrderChange sends POST /api/{api_version}/orders/{order_id}/change/preview.
func (o *Orders) PreviewOrderChange(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderChangePreviewPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: OrderReadRecordShape})
}

// ChangeOrderStatus sends POST /api/{api_version}/orders/{order_id}/change_status; it has no typed view, so read ApiResponse.Map.
func (o *Orders) ChangeOrderStatus(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderChangeStatusPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// AdjustOrderContract sends POST /api/{api_version}/orders/{order_id}/contract_adjustment; it has no typed view, so read ApiResponse.Map.
func (o *Orders) AdjustOrderContract(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderContractAdjustmentPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// PreviewOrderContractAdjustment sends POST /api/{api_version}/orders/{order_id}/contract_adjustment/preview; it has no typed view, so read ApiResponse.Map.
func (o *Orders) PreviewOrderContractAdjustment(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderContractAdjustmentPreviewPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DowngradeOrder sends POST /api/{api_version}/orders/{order_id}/downgrade; it has no typed view, so read ApiResponse.Map.
func (o *Orders) DowngradeOrder(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderDowngradePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// PreviewOrderDowngrade sends POST /api/{api_version}/orders/{order_id}/downgrade/preview; it has no typed view, so read ApiResponse.Map.
func (o *Orders) PreviewOrderDowngrade(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderDowngradePreviewPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpgradeOrder sends POST /api/{api_version}/orders/{order_id}/upgrade; it has no typed view, so read ApiResponse.Map.
func (o *Orders) UpgradeOrder(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderUpgradePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// PreviewOrderUpgrade sends POST /api/{api_version}/orders/{order_id}/upgrade/preview; it has no typed view, so read ApiResponse.Map.
func (o *Orders) PreviewOrderUpgrade(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderUpgradePreviewPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// TerminateOrderTrial sends POST /api/{api_version}/orders/{order_id}/terminate_trial; it has no typed view, so read ApiResponse.Map.
func (o *Orders) TerminateOrderTrial(ctx context.Context, orderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendPostRequest(ctx, OrderTerminateTrialPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteOrderNoteFile sends DELETE /api/{api_version}/orders/{order_id}/notes/{order_note_id}/files/{order_file_id}; it has no typed view, so read ApiResponse.Map.
func (o *Orders) DeleteOrderNoteFile(ctx context.Context, orderID any, orderNoteID any, orderFileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(orderNoteID, "order_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(orderFileID, "order_file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendDeleteRequest(ctx, OrderNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID, "order_note_id": orderNoteID, "order_file_id": orderFileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetOrderCustomObject sends GET /api/{api_version}/orders/{order_id}/custom_objects/{custom_object_id}.
func (o *Orders) GetOrderCustomObject(ctx context.Context, orderID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderCustomObjectPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID, "custom_object_id": customObjectID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: OrderCustomObjectRecordShape})
}

// GetOrderCustomObjectInstance sends GET /api/{api_version}/orders/{order_id}/custom_objects/{custom_object_id}/instances/{instance_id}; it has no typed view, so read ApiResponse.Map.
func (o *Orders) GetOrderCustomObjectInstance(ctx context.Context, orderID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(instanceID, "instance_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderCustomObjectInstancePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID, "custom_object_id": customObjectID, "instance_id": instanceID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetOrderCustomObjectsInstance sends GET /api/{api_version}/orders/{order_id}/custom_objects_instances/{custom_object_id}.
func (o *Orders) GetOrderCustomObjectsInstance(ctx context.Context, orderID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderCustomObjectsInstancePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID, "custom_object_id": customObjectID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: OrderCustomObjectRecordShape})
}

// GetOrderCustomObjectsInstanceInstance sends GET /api/{api_version}/orders/{order_id}/custom_objects_instances/{custom_object_id}/instances/{instance_id}; it has no typed view, so read ApiResponse.Map.
func (o *Orders) GetOrderCustomObjectsInstanceInstance(ctx context.Context, orderID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(instanceID, "instance_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderCustomObjectsInstanceInstancePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID, "custom_object_id": customObjectID, "instance_id": instanceID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetOrderItemFulfillment sends GET /api/{api_version}/orders/{order_id}/item_fulfillments/{fulfillment_id}.
func (o *Orders) GetOrderItemFulfillment(ctx context.Context, orderID any, fulfillmentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(fulfillmentID, "fulfillment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderItemFulfillmentPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID, "fulfillment_id": fulfillmentID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: OrderItemFulfillmentRecordShape})
}

// GetOrderKPIs sends GET /api/{api_version}/orders/{order_id}/kpis.
func (o *Orders) GetOrderKPIs(ctx context.Context, orderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderKPIsPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: OrderKpisRecordShape})
}

// GetOrderPayment sends GET /api/{api_version}/orders/{order_id}/payments/{payment_id}.
func (o *Orders) GetOrderPayment(ctx context.Context, orderID any, paymentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderPaymentPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID, "payment_id": paymentID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: OrderPaymentRecordShape})
}

// GetOrderQuotePartialInvoices sends GET /api/{api_version}/orders/{order_id}/quote_partial_invoices.
func (o *Orders) GetOrderQuotePartialInvoices(ctx context.Context, orderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderQuotePartialInvoicesPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: OrderQuotePartialInvoicesRecordShape})
}

// GetOrderShippingPDF sends GET /api/{api_version}/orders/{order_id}/shipping_pdf/{shipping_line_id}; it has no typed view, so read ApiResponse.Map.
func (o *Orders) GetOrderShippingPDF(ctx context.Context, orderID any, shippingLineID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(shippingLineID, "shipping_line_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return o.requester.SendGetRequest(ctx, OrderShippingPDFPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"order_id": orderID, "shipping_line_id": shippingLineID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "orders"

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

// requireNote refuses an empty note before any I/O; the text is the whole body of a note write.
func requireNote(note string) error {
	if note == "" {
		return exerrors.NewUsageError("note is required: the note text is the body of this request")
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

// New binds the orders endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Orders { return &Orders{requester: requester} }

// Requester returns the requester this resource was bound to.
func (o *Orders) Requester() exhttp.Requester { return o.requester }
