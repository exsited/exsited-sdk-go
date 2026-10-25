package bookings

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Bookings is the v4 bookings endpoints, bound to one client.
type Bookings struct {
	requester exhttp.Requester
}

// GetBooking sends GET /api/{api_version}/booking.
func (b *Bookings) GetBooking(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: BookingListShape})
}

// GetBookingByID sends GET /api/{api_version}/booking/{booking_id}.
func (b *Bookings) GetBookingByID(ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingByIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingByIdRecordShape})
}

// CreateBooking sends POST /api/{api_version}/booking; it has no typed view, so read ApiResponse.Map.
func (b *Bookings) CreateBooking(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendPostRequest(ctx, BookingPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// PreviewBooking sends POST /api/{api_version}/booking/preview; it has no typed view, so read ApiResponse.Map.
func (b *Bookings) PreviewBooking(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendPostRequest(ctx, BookingPreviewPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteBooking sends DELETE /api/{api_version}/booking/{booking_id}; it has no typed view, so read ApiResponse.Map.
func (b *Bookings) DeleteBooking(ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendDeleteRequest(ctx, BookingByIDPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetBookingNotes sends GET /api/{api_version}/booking/notes.
func (b *Bookings) GetBookingNotes(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingNotesPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingNoteListShape})
}

// GetBookingBillingPreferences sends GET /api/{api_version}/booking/{booking_id}/billing_preferences.
func (b *Bookings) GetBookingBillingPreferences(ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingBillingPreferencesPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingBillingPreferencesViewShape})
}

// UpdateBookingBillingPreferences sends PATCH /api/{api_version}/booking/{booking_id}/billing_preferences.
func (b *Bookings) UpdateBookingBillingPreferences(ctx context.Context, bookingID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendPatchRequest(ctx, BookingBillingPreferencesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingByIdRecordShape})
}

// ChangeBooking sends POST /api/{api_version}/booking/{booking_id}/change; it has no typed view, so read ApiResponse.Map.
func (b *Bookings) ChangeBooking(ctx context.Context, bookingID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendPostRequest(ctx, BookingChangePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// PreviewBookingChange sends POST /api/{api_version}/booking/{booking_id}/change/preview; it has no typed view, so read ApiResponse.Map.
func (b *Bookings) PreviewBookingChange(ctx context.Context, bookingID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendPostRequest(ctx, BookingChangePreviewPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetBookingInformation sends GET /api/{api_version}/booking/{booking_id}/information.
func (b *Bookings) GetBookingInformation(ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingInformationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingByIdRecordShape})
}

// UpdateBookingInformation sends PATCH /api/{api_version}/booking/{booking_id}/information.
func (b *Bookings) UpdateBookingInformation(ctx context.Context, bookingID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendPatchRequest(ctx, BookingInformationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingByIdRecordShape})
}

// GetBookingLines sends GET /api/{api_version}/booking/{booking_id}/lines.
func (b *Bookings) GetBookingLines(ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingLinesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingByIdRecordShape})
}

// GetBookingLine sends GET /api/{api_version}/booking/{booking_id}/lines/{line_uuid}.
func (b *Bookings) GetBookingLine(ctx context.Context, bookingID any, lineUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	if err := requireValue(lineUUID, "line_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingLinePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID, "line_uuid": lineUUID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingLineRecordShape})
}

// UpdateBookingLine sends PATCH /api/{api_version}/booking/{booking_id}/lines/{line_uuid}.
func (b *Bookings) UpdateBookingLine(ctx context.Context, bookingID any, lineUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
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
	return b.requester.SendPatchRequest(ctx, BookingLinePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID, "line_uuid": lineUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingByIdRecordShape})
}

// UpdateBookingLineBillingPreferences sends PATCH /api/{api_version}/booking/{booking_id}/lines/{line_uuid}/billing_preferences.
func (b *Bookings) UpdateBookingLineBillingPreferences(ctx context.Context, bookingID any, lineUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
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
	return b.requester.SendPatchRequest(ctx, BookingLineBillingPreferencesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID, "line_uuid": lineUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingByIdRecordShape})
}

// GetBookingCustomObjects sends GET /api/{api_version}/booking/{booking_id}/custom_objects.
func (b *Bookings) GetBookingCustomObjects(ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingCustomObjectsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingCustomObjectListShape})
}

// GetBookingCustomObject sends GET /api/{api_version}/booking/{booking_id}/custom_objects/{custom_object_id}.
func (b *Bookings) GetBookingCustomObject(ctx context.Context, bookingID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingCustomObjectPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID, "custom_object_id": customObjectID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingCustomObjectRecordShape})
}

// GetBookingCustomObjectInstance sends GET /api/{api_version}/booking/{booking_id}/custom_objects/{custom_object_id}/instances/{instance_id}; it has no typed view, so read ApiResponse.Map.
func (b *Bookings) GetBookingCustomObjectInstance(ctx context.Context, bookingID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
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
	return b.requester.SendGetRequest(ctx, BookingCustomObjectInstancePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID, "custom_object_id": customObjectID, "instance_id": instanceID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetBookingCustomObjectsInstances sends GET /api/{api_version}/booking/{booking_id}/custom_objects_instances.
func (b *Bookings) GetBookingCustomObjectsInstances(ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingCustomObjectsInstancesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingCustomObjectListShape})
}

// GetBookingCustomObjectsInstance sends GET /api/{api_version}/booking/{booking_id}/custom_objects_instances/{custom_object_id}.
func (b *Bookings) GetBookingCustomObjectsInstance(ctx context.Context, bookingID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingCustomObjectsInstancePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID, "custom_object_id": customObjectID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingCustomObjectRecordShape})
}

// GetBookingCustomObjectsInstanceInstance sends GET /api/{api_version}/booking/{booking_id}/custom_objects_instances/{custom_object_id}/instances/{instance_id}; it has no typed view, so read ApiResponse.Map.
func (b *Bookings) GetBookingCustomObjectsInstanceInstance(ctx context.Context, bookingID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
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
	return b.requester.SendGetRequest(ctx, BookingCustomObjectsInstanceInstancePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID, "custom_object_id": customObjectID, "instance_id": instanceID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetBookingInvoices sends GET /api/{api_version}/booking/{booking_id}/invoices.
func (b *Bookings) GetBookingInvoices(ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingInvoicesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingInvoiceListShape})
}

// GetBookingItemFulfillments sends GET /api/{api_version}/booking/{booking_id}/item_fulfillments; it has no typed view, so read ApiResponse.Map.
func (b *Bookings) GetBookingItemFulfillments(ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingItemFulfillmentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetBookingItemFulfillment sends GET /api/{api_version}/booking/{booking_id}/item_fulfillments/{fulfillment_id}.
func (b *Bookings) GetBookingItemFulfillment(ctx context.Context, bookingID any, fulfillmentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	if err := requireValue(fulfillmentID, "fulfillment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingItemFulfillmentPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID, "fulfillment_id": fulfillmentID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingItemFulfillmentRecordShape})
}

// GetBookingKPIs sends GET /api/{api_version}/booking/{booking_id}/kpis.
func (b *Bookings) GetBookingKPIs(ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingKPIsPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingKpisRecordShape})
}

// GetBookingPaymentMethods sends GET /api/{api_version}/booking/{booking_id}/payment_methods.
func (b *Bookings) GetBookingPaymentMethods(ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingPaymentMethodsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingPaymentMethodsRecordShape})
}

// GetBookingPaymentMethod sends GET /api/{api_version}/booking/{booking_id}/payment_methods/{payment_method_ref}.
func (b *Bookings) GetBookingPaymentMethod(ctx context.Context, bookingID any, paymentMethodRef any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	if err := requireValue(paymentMethodRef, "payment_method_ref"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingPaymentMethodPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID, "payment_method_ref": paymentMethodRef},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingPaymentMethodRecordShape})
}

// GetBookingPayments sends GET /api/{api_version}/booking/{booking_id}/payments.
func (b *Bookings) GetBookingPayments(ctx context.Context, bookingID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingPaymentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"booking_id": bookingID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingPaymentListShape})
}

// GetBookingPayment sends GET /api/{api_version}/booking/{booking_id}/payments/{payment_id}.
func (b *Bookings) GetBookingPayment(ctx context.Context, bookingID any, paymentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingPaymentPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID, "payment_id": paymentID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingPaymentViewShape})
}

// GetBookingPDF sends GET /api/{api_version}/booking/{booking_id}/pdf.
func (b *Bookings) GetBookingPDF(ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingPDFPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingPdfRecordShape})
}

// GetBookingQuotePartialInvoices sends GET /api/{api_version}/booking/{booking_id}/quote_partial_invoices.
func (b *Bookings) GetBookingQuotePartialInvoices(ctx context.Context, bookingID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingQuotePartialInvoicesPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: BookingQuotePartialInvoicesRecordShape})
}

// GetBookingShippingPDF sends GET /api/{api_version}/booking/{booking_id}/shipping_pdf/{shipping_line_id}; it has no typed view, so read ApiResponse.Map.
func (b *Bookings) GetBookingShippingPDF(ctx context.Context, bookingID any, shippingLineID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(bookingID, "booking_id"); err != nil {
		return nil, err
	}
	if err := requireValue(shippingLineID, "shipping_line_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return b.requester.SendGetRequest(ctx, BookingShippingPDFPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"booking_id": bookingID, "shipping_line_id": shippingLineID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "bookings"

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

// requirePayload refuses a nil payload, a typed nil included, before any I/O.
func requirePayload(payload any) error {
	if payload == nil || isNilValue(payload) {
		return exerrors.NewUsageError("payload is required: this endpoint takes a body, and nil would " +
			"send the request with no body at all")
	}
	return nil
}

// New binds the bookings endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Bookings { return &Bookings{requester: requester} }

// Requester returns the requester this resource was bound to.
func (b *Bookings) Requester() exhttp.Requester { return b.requester }
