package bookings

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Bookings method of the same name,
// which is the primary surface.

// GetBooking is the module form of Bookings.GetBooking.
func GetBooking(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBooking(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetBookingByID is the module form of Bookings.GetBookingByID.
func GetBookingByID(requester exhttp.Requester, ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingByID(ctx, bookingID, queryParams, raiseOnError, responseFormat...)
}

// CreateBooking is the module form of Bookings.CreateBooking.
func CreateBooking(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateBooking(ctx, payload, raiseOnError, responseFormat...)
}

// PreviewBooking is the module form of Bookings.PreviewBooking.
func PreviewBooking(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).PreviewBooking(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteBooking is the module form of Bookings.DeleteBooking.
func DeleteBooking(requester exhttp.Requester, ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteBooking(ctx, bookingID, raiseOnError, responseFormat...)
}

// GetBookingNotes is the module form of Bookings.GetBookingNotes.
func GetBookingNotes(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingNotes(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetBookingBillingPreferences is the module form of Bookings.GetBookingBillingPreferences.
func GetBookingBillingPreferences(requester exhttp.Requester, ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingBillingPreferences(ctx, bookingID, raiseOnError, responseFormat...)
}

// UpdateBookingBillingPreferences is the module form of Bookings.UpdateBookingBillingPreferences.
func UpdateBookingBillingPreferences(requester exhttp.Requester, ctx context.Context, bookingID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateBookingBillingPreferences(ctx, bookingID, payload, raiseOnError, responseFormat...)
}

// ChangeBooking is the module form of Bookings.ChangeBooking.
func ChangeBooking(requester exhttp.Requester, ctx context.Context, bookingID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ChangeBooking(ctx, bookingID, payload, raiseOnError, responseFormat...)
}

// PreviewBookingChange is the module form of Bookings.PreviewBookingChange.
func PreviewBookingChange(requester exhttp.Requester, ctx context.Context, bookingID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).PreviewBookingChange(ctx, bookingID, payload, raiseOnError, responseFormat...)
}

// GetBookingInformation is the module form of Bookings.GetBookingInformation.
func GetBookingInformation(requester exhttp.Requester, ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingInformation(ctx, bookingID, raiseOnError, responseFormat...)
}

// UpdateBookingInformation is the module form of Bookings.UpdateBookingInformation.
func UpdateBookingInformation(requester exhttp.Requester, ctx context.Context, bookingID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateBookingInformation(ctx, bookingID, payload, raiseOnError, responseFormat...)
}

// GetBookingLines is the module form of Bookings.GetBookingLines.
func GetBookingLines(requester exhttp.Requester, ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingLines(ctx, bookingID, queryParams, raiseOnError, responseFormat...)
}

// GetBookingLine is the module form of Bookings.GetBookingLine.
func GetBookingLine(requester exhttp.Requester, ctx context.Context, bookingID any, lineUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingLine(ctx, bookingID, lineUUID, raiseOnError, responseFormat...)
}

// UpdateBookingLine is the module form of Bookings.UpdateBookingLine.
func UpdateBookingLine(requester exhttp.Requester, ctx context.Context, bookingID any, lineUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateBookingLine(ctx, bookingID, lineUUID, payload, raiseOnError, responseFormat...)
}

// UpdateBookingLineBillingPreferences is the module form of Bookings.UpdateBookingLineBillingPreferences.
func UpdateBookingLineBillingPreferences(requester exhttp.Requester, ctx context.Context, bookingID any, lineUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateBookingLineBillingPreferences(ctx, bookingID, lineUUID, payload, raiseOnError, responseFormat...)
}

// GetBookingCustomObjects is the module form of Bookings.GetBookingCustomObjects.
func GetBookingCustomObjects(requester exhttp.Requester, ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingCustomObjects(ctx, bookingID, queryParams, raiseOnError, responseFormat...)
}

// GetBookingCustomObject is the module form of Bookings.GetBookingCustomObject.
func GetBookingCustomObject(requester exhttp.Requester, ctx context.Context, bookingID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingCustomObject(ctx, bookingID, customObjectID, raiseOnError, responseFormat...)
}

// GetBookingCustomObjectInstance is the module form of Bookings.GetBookingCustomObjectInstance.
func GetBookingCustomObjectInstance(requester exhttp.Requester, ctx context.Context, bookingID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingCustomObjectInstance(ctx, bookingID, customObjectID, instanceID, raiseOnError, responseFormat...)
}

// GetBookingCustomObjectsInstances is the module form of Bookings.GetBookingCustomObjectsInstances.
func GetBookingCustomObjectsInstances(requester exhttp.Requester, ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingCustomObjectsInstances(ctx, bookingID, queryParams, raiseOnError, responseFormat...)
}

// GetBookingCustomObjectsInstance is the module form of Bookings.GetBookingCustomObjectsInstance.
func GetBookingCustomObjectsInstance(requester exhttp.Requester, ctx context.Context, bookingID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingCustomObjectsInstance(ctx, bookingID, customObjectID, raiseOnError, responseFormat...)
}

// GetBookingCustomObjectsInstanceInstance is the module form of Bookings.GetBookingCustomObjectsInstanceInstance.
func GetBookingCustomObjectsInstanceInstance(requester exhttp.Requester, ctx context.Context, bookingID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingCustomObjectsInstanceInstance(ctx, bookingID, customObjectID, instanceID, raiseOnError, responseFormat...)
}

// GetBookingInvoices is the module form of Bookings.GetBookingInvoices.
func GetBookingInvoices(requester exhttp.Requester, ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingInvoices(ctx, bookingID, queryParams, raiseOnError, responseFormat...)
}

// GetBookingItemFulfillments is the module form of Bookings.GetBookingItemFulfillments.
func GetBookingItemFulfillments(requester exhttp.Requester, ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingItemFulfillments(ctx, bookingID, queryParams, raiseOnError, responseFormat...)
}

// GetBookingItemFulfillment is the module form of Bookings.GetBookingItemFulfillment.
func GetBookingItemFulfillment(requester exhttp.Requester, ctx context.Context, bookingID any, fulfillmentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingItemFulfillment(ctx, bookingID, fulfillmentID, raiseOnError, responseFormat...)
}

// GetBookingKPIs is the module form of Bookings.GetBookingKPIs.
func GetBookingKPIs(requester exhttp.Requester, ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingKPIs(ctx, bookingID, raiseOnError, responseFormat...)
}

// GetBookingPaymentMethods is the module form of Bookings.GetBookingPaymentMethods.
func GetBookingPaymentMethods(requester exhttp.Requester, ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingPaymentMethods(ctx, bookingID, queryParams, raiseOnError, responseFormat...)
}

// GetBookingPaymentMethod is the module form of Bookings.GetBookingPaymentMethod.
func GetBookingPaymentMethod(requester exhttp.Requester, ctx context.Context, bookingID any, paymentMethodRef any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingPaymentMethod(ctx, bookingID, paymentMethodRef, raiseOnError, responseFormat...)
}

// GetBookingPayments is the module form of Bookings.GetBookingPayments.
func GetBookingPayments(requester exhttp.Requester, ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingPayments(ctx, bookingID, queryParams, raiseOnError, responseFormat...)
}

// GetBookingPayment is the module form of Bookings.GetBookingPayment.
func GetBookingPayment(requester exhttp.Requester, ctx context.Context, bookingID any, paymentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingPayment(ctx, bookingID, paymentID, raiseOnError, responseFormat...)
}

// GetBookingPDF is the module form of Bookings.GetBookingPDF.
func GetBookingPDF(requester exhttp.Requester, ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingPDF(ctx, bookingID, raiseOnError, responseFormat...)
}

// GetBookingQuotePartialInvoices is the module form of Bookings.GetBookingQuotePartialInvoices.
func GetBookingQuotePartialInvoices(requester exhttp.Requester, ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingQuotePartialInvoices(ctx, bookingID, raiseOnError, responseFormat...)
}

// GetBookingShippingPDF is the module form of Bookings.GetBookingShippingPDF.
func GetBookingShippingPDF(requester exhttp.Requester, ctx context.Context, bookingID any, shippingLineID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetBookingShippingPDF(ctx, bookingID, shippingLineID, raiseOnError, responseFormat...)
}
