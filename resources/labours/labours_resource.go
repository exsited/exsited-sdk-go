package labours

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Labours method of the same name,
// which is the primary surface.

// GetLabourAvailability is the module form of Labours.GetLabourAvailability.
func GetLabourAvailability(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetLabourAvailability(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetLabourProfiles is the module form of Labours.GetLabourProfiles.
func GetLabourProfiles(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetLabourProfiles(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateLabourProfile is the module form of Labours.CreateLabourProfile.
func CreateLabourProfile(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateLabourProfile(ctx, payload, raiseOnError, responseFormat...)
}

// GetLabours is the module form of Labours.GetLabours.
func GetLabours(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetLabours(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateLabour is the module form of Labours.CreateLabour.
func CreateLabour(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateLabour(ctx, payload, raiseOnError, responseFormat...)
}

// GetLabourProfile is the module form of Labours.GetLabourProfile.
func GetLabourProfile(requester exhttp.Requester, ctx context.Context, labourProfileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetLabourProfile(ctx, labourProfileID, queryParams, raiseOnError, responseFormat...)
}

// UpdateLabourProfile is the module form of Labours.UpdateLabourProfile.
func UpdateLabourProfile(requester exhttp.Requester, ctx context.Context, labourProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateLabourProfile(ctx, labourProfileID, payload, raiseOnError, responseFormat...)
}

// DeleteLabourProfile is the module form of Labours.DeleteLabourProfile.
func DeleteLabourProfile(requester exhttp.Requester, ctx context.Context, labourProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteLabourProfile(ctx, labourProfileID, raiseOnError, responseFormat...)
}

// GetLaboursAvailableForBooking is the module form of Labours.GetLaboursAvailableForBooking.
func GetLaboursAvailableForBooking(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetLaboursAvailableForBooking(ctx, queryParams, raiseOnError, responseFormat...)
}

// UpdateLabour is the module form of Labours.UpdateLabour.
func UpdateLabour(requester exhttp.Requester, ctx context.Context, labourUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateLabour(ctx, labourUUID, payload, raiseOnError, responseFormat...)
}

// DeleteLabour is the module form of Labours.DeleteLabour.
func DeleteLabour(requester exhttp.Requester, ctx context.Context, labourUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteLabour(ctx, labourUUID, raiseOnError, responseFormat...)
}

// GetLabour is the module form of Labours.GetLabour.
func GetLabour(requester exhttp.Requester, ctx context.Context, labourUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetLabour(ctx, labourUUID, queryParams, raiseOnError, responseFormat...)
}

// CancelLabour is the module form of Labours.CancelLabour.
func CancelLabour(requester exhttp.Requester, ctx context.Context, labourID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CancelLabour(ctx, labourID, payload, raiseOnError, responseFormat...)
}

// ReactivateLabour is the module form of Labours.ReactivateLabour.
func ReactivateLabour(requester exhttp.Requester, ctx context.Context, labourID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ReactivateLabour(ctx, labourID, payload, raiseOnError, responseFormat...)
}

// GetLabourOrders is the module form of Labours.GetLabourOrders.
func GetLabourOrders(requester exhttp.Requester, ctx context.Context, labourUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetLabourOrders(ctx, labourUUID, queryParams, raiseOnError, responseFormat...)
}

// UpdateLabourAvailability is the module form of Labours.UpdateLabourAvailability.
func UpdateLabourAvailability(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateLabourAvailability(ctx, payload, raiseOnError, responseFormat...)
}

// GetLabourProfileSessions is the module form of Labours.GetLabourProfileSessions.
func GetLabourProfileSessions(requester exhttp.Requester, ctx context.Context, labourProfileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetLabourProfileSessions(ctx, labourProfileID, queryParams, raiseOnError, responseFormat...)
}

// GetLabourProfileSession is the module form of Labours.GetLabourProfileSession.
func GetLabourProfileSession(requester exhttp.Requester, ctx context.Context, labourProfileID any, sessionID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetLabourProfileSession(ctx, labourProfileID, sessionID, raiseOnError, responseFormat...)
}

// GetLabourCalendar is the module form of Labours.GetLabourCalendar.
func GetLabourCalendar(requester exhttp.Requester, ctx context.Context, labourID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetLabourCalendar(ctx, labourID, queryParams, raiseOnError, responseFormat...)
}
