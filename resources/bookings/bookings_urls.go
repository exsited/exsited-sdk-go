// Package bookings is the v4 bookings endpoints.
package bookings

// The path templates for the bookings endpoints; the client fills {api_version}.
const (
	BookingPath                              = "/api/{api_version}/booking"
	BookingByIDPath                          = "/api/{api_version}/booking/{booking_id}"
	BookingPreviewPath                       = "/api/{api_version}/booking/preview"
	BookingNotesPath                         = "/api/{api_version}/booking/notes"
	BookingBillingPreferencesPath            = "/api/{api_version}/booking/{booking_id}/billing_preferences"
	BookingChangePath                        = "/api/{api_version}/booking/{booking_id}/change"
	BookingChangePreviewPath                 = "/api/{api_version}/booking/{booking_id}/change/preview"
	BookingInformationPath                   = "/api/{api_version}/booking/{booking_id}/information"
	BookingLinesPath                         = "/api/{api_version}/booking/{booking_id}/lines"
	BookingLinePath                          = "/api/{api_version}/booking/{booking_id}/lines/{line_uuid}"
	BookingLineBillingPreferencesPath        = "/api/{api_version}/booking/{booking_id}/lines/{line_uuid}/billing_preferences"
	BookingCustomObjectsPath                 = "/api/{api_version}/booking/{booking_id}/custom_objects"
	BookingCustomObjectPath                  = "/api/{api_version}/booking/{booking_id}/custom_objects/{custom_object_id}"
	BookingCustomObjectInstancePath          = "/api/{api_version}/booking/{booking_id}/custom_objects/{custom_object_id}/instances/{instance_id}"
	BookingCustomObjectsInstancesPath        = "/api/{api_version}/booking/{booking_id}/custom_objects_instances"
	BookingCustomObjectsInstancePath         = "/api/{api_version}/booking/{booking_id}/custom_objects_instances/{custom_object_id}"
	BookingCustomObjectsInstanceInstancePath = "/api/{api_version}/booking/{booking_id}/custom_objects_instances/{custom_object_id}/instances/{instance_id}"
	BookingInvoicesPath                      = "/api/{api_version}/booking/{booking_id}/invoices"
	BookingItemFulfillmentsPath              = "/api/{api_version}/booking/{booking_id}/item_fulfillments"
	BookingItemFulfillmentPath               = "/api/{api_version}/booking/{booking_id}/item_fulfillments/{fulfillment_id}"
	BookingKPIsPath                          = "/api/{api_version}/booking/{booking_id}/kpis"
	BookingPaymentMethodsPath                = "/api/{api_version}/booking/{booking_id}/payment_methods"
	BookingPaymentMethodPath                 = "/api/{api_version}/booking/{booking_id}/payment_methods/{payment_method_ref}"
	BookingPaymentsPath                      = "/api/{api_version}/booking/{booking_id}/payments"
	BookingPaymentPath                       = "/api/{api_version}/booking/{booking_id}/payments/{payment_id}"
	BookingPDFPath                           = "/api/{api_version}/booking/{booking_id}/pdf"
	BookingQuotePartialInvoicesPath          = "/api/{api_version}/booking/{booking_id}/quote_partial_invoices"
	BookingShippingPDFPath                   = "/api/{api_version}/booking/{booking_id}/shipping_pdf/{shipping_line_id}"
)

// BookingsKey is the key a listing's records arrive under, the v4 list envelope's data.
const BookingsKey = "data"

// OrderKey is the envelope key one record arrives under.
const OrderKey = "order"

// OrderLineKey is the envelope key one record arrives under.
const OrderLineKey = "order_line"

// CustomObjectKey is the envelope key one record arrives under.
const CustomObjectKey = "custom_object"

// ItemFulfillmentKey is the envelope key one record arrives under.
const ItemFulfillmentKey = "item_fulfillment"

// OrderKPIsKey is the envelope key one record arrives under.
const OrderKPIsKey = "order_kpis"

// PaymentKey is the envelope key one record arrives under.
const PaymentKey = "payment"
