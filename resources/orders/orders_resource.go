package orders

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Orders method of the same name,
// which is the primary surface.

// GetOrders is the module form of Orders.GetOrders.
func GetOrders(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrders(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateOrder is the module form of Orders.CreateOrder.
func CreateOrder(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateOrder(ctx, payload, raiseOnError, responseFormat...)
}

// GetOrder is the module form of Orders.GetOrder.
func GetOrder(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrder(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// DeleteOrder is the module form of Orders.DeleteOrder.
func DeleteOrder(requester exhttp.Requester, ctx context.Context, orderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteOrder(ctx, orderID, raiseOnError, responseFormat...)
}

// ActivateOrder is the module form of Orders.ActivateOrder.
func ActivateOrder(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ActivateOrder(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// CreateOrderActivateDraft is the module form of Orders.CreateOrderActivateDraft.
func CreateOrderActivateDraft(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateOrderActivateDraft(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// GetOrderBillingPreferences is the module form of Orders.GetOrderBillingPreferences.
func GetOrderBillingPreferences(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderBillingPreferences(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// UpdateOrderBillingPreferences is the module form of Orders.UpdateOrderBillingPreferences.
func UpdateOrderBillingPreferences(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateOrderBillingPreferences(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// ReplaceOrderBillingPreferences is the module form of Orders.ReplaceOrderBillingPreferences.
func ReplaceOrderBillingPreferences(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReplaceOrderBillingPreferences(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// CancelOrder is the module form of Orders.CancelOrder.
func CancelOrder(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CancelOrder(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// ChangeOrder is the module form of Orders.ChangeOrder.
func ChangeOrder(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ChangeOrder(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// CreateOrderConvertToOrder is the module form of Orders.CreateOrderConvertToOrder.
func CreateOrderConvertToOrder(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateOrderConvertToOrder(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// GetOrderCustomObjects is the module form of Orders.GetOrderCustomObjects.
func GetOrderCustomObjects(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderCustomObjects(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// GetOrderCustomObjectsInstances is the module form of Orders.GetOrderCustomObjectsInstances.
func GetOrderCustomObjectsInstances(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderCustomObjectsInstances(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// CreateOrderDiscardDraft is the module form of Orders.CreateOrderDiscardDraft.
func CreateOrderDiscardDraft(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateOrderDiscardDraft(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// CreateOrderExtendTrial is the module form of Orders.CreateOrderExtendTrial.
func CreateOrderExtendTrial(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateOrderExtendTrial(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// GetOrderInformation is the module form of Orders.GetOrderInformation.
func GetOrderInformation(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderInformation(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// UpdateOrderInformation is the module form of Orders.UpdateOrderInformation.
func UpdateOrderInformation(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateOrderInformation(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// GetOrderItemFulfillments is the module form of Orders.GetOrderItemFulfillments.
func GetOrderItemFulfillments(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderItemFulfillments(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// GetOrderLines is the module form of Orders.GetOrderLines.
func GetOrderLines(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderLines(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// GetOrderNotes is the module form of Orders.GetOrderNotes.
func GetOrderNotes(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderNotes(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// CreateOrderNote is the module form of Orders.CreateOrderNote.
func CreateOrderNote(requester exhttp.Requester, ctx context.Context, orderID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateOrderNote(ctx, orderID, note, files, raiseOnError, responseFormat...)
}

// GetOrderPaymentMethods is the module form of Orders.GetOrderPaymentMethods.
func GetOrderPaymentMethods(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderPaymentMethods(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// GetOrderPDF is the module form of Orders.GetOrderPDF.
func GetOrderPDF(requester exhttp.Requester, ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderPDF(ctx, orderID, queryParams, raiseOnError, responseFormat...)
}

// ReactivateOrder is the module form of Orders.ReactivateOrder.
func ReactivateOrder(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReactivateOrder(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// CreateOrderRelinquish is the module form of Orders.CreateOrderRelinquish.
func CreateOrderRelinquish(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateOrderRelinquish(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// UpdateOrderLine is the module form of Orders.UpdateOrderLine.
func UpdateOrderLine(requester exhttp.Requester, ctx context.Context, orderID any, lineUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateOrderLine(ctx, orderID, lineUUID, payload, raiseOnError, responseFormat...)
}

// GetOrderLine is the module form of Orders.GetOrderLine.
func GetOrderLine(requester exhttp.Requester, ctx context.Context, orderID any, lineUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderLine(ctx, orderID, lineUUID, queryParams, raiseOnError, responseFormat...)
}

// DeleteOrderNote is the module form of Orders.DeleteOrderNote.
func DeleteOrderNote(requester exhttp.Requester, ctx context.Context, orderID any, orderNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteOrderNote(ctx, orderID, orderNoteID, raiseOnError, responseFormat...)
}

// GetOrderNote is the module form of Orders.GetOrderNote.
func GetOrderNote(requester exhttp.Requester, ctx context.Context, orderID any, orderNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderNote(ctx, orderID, orderNoteID, queryParams, raiseOnError, responseFormat...)
}

// UpdateOrderNote is the module form of Orders.UpdateOrderNote.
func UpdateOrderNote(requester exhttp.Requester, ctx context.Context, orderID any, orderNoteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateOrderNote(ctx, orderID, orderNoteID, note, files, raiseOnError, responseFormat...)
}

// GetOrderPaymentMethod is the module form of Orders.GetOrderPaymentMethod.
func GetOrderPaymentMethod(requester exhttp.Requester, ctx context.Context, orderID any, paymentMethodRef any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderPaymentMethod(ctx, orderID, paymentMethodRef, queryParams, raiseOnError, responseFormat...)
}

// UpdateOrderLineBillingPreferences is the module form of Orders.UpdateOrderLineBillingPreferences.
func UpdateOrderLineBillingPreferences(requester exhttp.Requester, ctx context.Context, orderID any, lineUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateOrderLineBillingPreferences(ctx, orderID, lineUUID, payload, raiseOnError, responseFormat...)
}

// ReplaceOrderLineBillingPreferences is the module form of Orders.ReplaceOrderLineBillingPreferences.
func ReplaceOrderLineBillingPreferences(requester exhttp.Requester, ctx context.Context, orderID any, lineUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReplaceOrderLineBillingPreferences(ctx, orderID, lineUUID, payload, raiseOnError, responseFormat...)
}

// GetOrderNoteFiles is the module form of Orders.GetOrderNoteFiles.
func GetOrderNoteFiles(requester exhttp.Requester, ctx context.Context, orderID any, orderNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderNoteFiles(ctx, orderID, orderNoteID, queryParams, raiseOnError, responseFormat...)
}

// GetOrderNoteFile is the module form of Orders.GetOrderNoteFile.
func GetOrderNoteFile(requester exhttp.Requester, ctx context.Context, orderID any, orderNoteID any, orderFileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderNoteFile(ctx, orderID, orderNoteID, orderFileID, queryParams, raiseOnError, responseFormat...)
}

// GetOrdersNotes is the module form of Orders.GetOrdersNotes.
func GetOrdersNotes(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrdersNotes(ctx, queryParams, raiseOnError, responseFormat...)
}

// PreviewOrder is the module form of Orders.PreviewOrder.
func PreviewOrder(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).PreviewOrder(ctx, payload, raiseOnError, responseFormat...)
}

// CreateOrderExpress is the module form of Orders.CreateOrderExpress.
func CreateOrderExpress(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateOrderExpress(ctx, payload, raiseOnError, responseFormat...)
}

// AcceptOrder is the module form of Orders.AcceptOrder.
func AcceptOrder(requester exhttp.Requester, ctx context.Context, orderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).AcceptOrder(ctx, orderID, raiseOnError, responseFormat...)
}

// ActivateOrderInvoiceFromDormant is the module form of Orders.ActivateOrderInvoiceFromDormant.
func ActivateOrderInvoiceFromDormant(requester exhttp.Requester, ctx context.Context, orderID any, invoiceUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ActivateOrderInvoiceFromDormant(ctx, orderID, invoiceUUID, payload, raiseOnError, responseFormat...)
}

// PreviewOrderChange is the module form of Orders.PreviewOrderChange.
func PreviewOrderChange(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).PreviewOrderChange(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// ChangeOrderStatus is the module form of Orders.ChangeOrderStatus.
func ChangeOrderStatus(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ChangeOrderStatus(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// AdjustOrderContract is the module form of Orders.AdjustOrderContract.
func AdjustOrderContract(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).AdjustOrderContract(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// PreviewOrderContractAdjustment is the module form of Orders.PreviewOrderContractAdjustment.
func PreviewOrderContractAdjustment(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).PreviewOrderContractAdjustment(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// DowngradeOrder is the module form of Orders.DowngradeOrder.
func DowngradeOrder(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DowngradeOrder(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// PreviewOrderDowngrade is the module form of Orders.PreviewOrderDowngrade.
func PreviewOrderDowngrade(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).PreviewOrderDowngrade(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// UpgradeOrder is the module form of Orders.UpgradeOrder.
func UpgradeOrder(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpgradeOrder(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// PreviewOrderUpgrade is the module form of Orders.PreviewOrderUpgrade.
func PreviewOrderUpgrade(requester exhttp.Requester, ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).PreviewOrderUpgrade(ctx, orderID, payload, raiseOnError, responseFormat...)
}

// TerminateOrderTrial is the module form of Orders.TerminateOrderTrial.
func TerminateOrderTrial(requester exhttp.Requester, ctx context.Context, orderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).TerminateOrderTrial(ctx, orderID, raiseOnError, responseFormat...)
}

// DeleteOrderNoteFile is the module form of Orders.DeleteOrderNoteFile.
func DeleteOrderNoteFile(requester exhttp.Requester, ctx context.Context, orderID any, orderNoteID any, orderFileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteOrderNoteFile(ctx, orderID, orderNoteID, orderFileID, raiseOnError, responseFormat...)
}

// GetOrderCustomObject is the module form of Orders.GetOrderCustomObject.
func GetOrderCustomObject(requester exhttp.Requester, ctx context.Context, orderID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderCustomObject(ctx, orderID, customObjectID, raiseOnError, responseFormat...)
}

// GetOrderCustomObjectInstance is the module form of Orders.GetOrderCustomObjectInstance.
func GetOrderCustomObjectInstance(requester exhttp.Requester, ctx context.Context, orderID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderCustomObjectInstance(ctx, orderID, customObjectID, instanceID, raiseOnError, responseFormat...)
}

// GetOrderCustomObjectsInstance is the module form of Orders.GetOrderCustomObjectsInstance.
func GetOrderCustomObjectsInstance(requester exhttp.Requester, ctx context.Context, orderID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderCustomObjectsInstance(ctx, orderID, customObjectID, raiseOnError, responseFormat...)
}

// GetOrderCustomObjectsInstanceInstance is the module form of Orders.GetOrderCustomObjectsInstanceInstance.
func GetOrderCustomObjectsInstanceInstance(requester exhttp.Requester, ctx context.Context, orderID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderCustomObjectsInstanceInstance(ctx, orderID, customObjectID, instanceID, raiseOnError, responseFormat...)
}

// GetOrderItemFulfillment is the module form of Orders.GetOrderItemFulfillment.
func GetOrderItemFulfillment(requester exhttp.Requester, ctx context.Context, orderID any, fulfillmentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderItemFulfillment(ctx, orderID, fulfillmentID, raiseOnError, responseFormat...)
}

// GetOrderKPIs is the module form of Orders.GetOrderKPIs.
func GetOrderKPIs(requester exhttp.Requester, ctx context.Context, orderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderKPIs(ctx, orderID, raiseOnError, responseFormat...)
}

// GetOrderPayment is the module form of Orders.GetOrderPayment.
func GetOrderPayment(requester exhttp.Requester, ctx context.Context, orderID any, paymentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderPayment(ctx, orderID, paymentID, raiseOnError, responseFormat...)
}

// GetOrderQuotePartialInvoices is the module form of Orders.GetOrderQuotePartialInvoices.
func GetOrderQuotePartialInvoices(requester exhttp.Requester, ctx context.Context, orderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderQuotePartialInvoices(ctx, orderID, raiseOnError, responseFormat...)
}

// GetOrderShippingPDF is the module form of Orders.GetOrderShippingPDF.
func GetOrderShippingPDF(requester exhttp.Requester, ctx context.Context, orderID any, shippingLineID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetOrderShippingPDF(ctx, orderID, shippingLineID, raiseOnError, responseFormat...)
}
