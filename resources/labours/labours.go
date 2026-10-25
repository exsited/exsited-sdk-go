package labours

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Labours is the v4 labours endpoints, bound to one client.
type Labours struct {
	requester exhttp.Requester
}

// GetLabourAvailability sends GET /api/{api_version}/labour_availability.
func (l *Labours) GetLabourAvailability(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendGetRequest(ctx, LabourAvailabilityPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: LabourAvailabilityListShape})
}

// GetLabourProfiles sends GET /api/{api_version}/labour_profiles.
func (l *Labours) GetLabourProfiles(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendGetRequest(ctx, LabourProfilesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: LabourProfileListShape})
}

// CreateLabourProfile sends POST /api/{api_version}/labour_profiles.
func (l *Labours) CreateLabourProfile(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendPostRequest(ctx, LabourProfilesPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: LabourProfileReadRecordShape})
}

// GetLabours sends GET /api/{api_version}/labours.
func (l *Labours) GetLabours(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendGetRequest(ctx, LaboursPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: LabourListShape})
}

// CreateLabour sends POST /api/{api_version}/labours.
func (l *Labours) CreateLabour(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendPostRequest(ctx, LaboursPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: LabourRecordShape})
}

// GetLabourProfile sends GET /api/{api_version}/labour_profiles/{labour_profile_id}.
func (l *Labours) GetLabourProfile(ctx context.Context, labourProfileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourProfileID, "labour_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendGetRequest(ctx, LabourProfilePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_profile_id": labourProfileID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: LabourProfileReadRecordShape})
}

// UpdateLabourProfile sends PATCH /api/{api_version}/labour_profiles/{labour_profile_id}; it has no typed view, so read ApiResponse.Map.
func (l *Labours) UpdateLabourProfile(ctx context.Context, labourProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourProfileID, "labour_profile_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendPatchRequest(ctx, LabourProfilePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_profile_id": labourProfileID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteLabourProfile sends DELETE /api/{api_version}/labour_profiles/{labour_profile_id}; it has no typed view, so read ApiResponse.Map.
func (l *Labours) DeleteLabourProfile(ctx context.Context, labourProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourProfileID, "labour_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendDeleteRequest(ctx, LabourProfilePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_profile_id": labourProfileID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetLaboursAvailableForBooking sends GET /api/{api_version}/labours/available_for_booking.
func (l *Labours) GetLaboursAvailableForBooking(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendGetRequest(ctx, LaboursAvailableForBookingPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: LabourListShape})
}

// UpdateLabour sends PATCH /api/{api_version}/labours/{labour_uuid}.
func (l *Labours) UpdateLabour(ctx context.Context, labourUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourUUID, "labour_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendPatchRequest(ctx, LabourPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_uuid": labourUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: LabourRecordShape})
}

// DeleteLabour sends DELETE /api/{api_version}/labours/{labour_uuid}; it has no typed view, so read ApiResponse.Map.
func (l *Labours) DeleteLabour(ctx context.Context, labourUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourUUID, "labour_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendDeleteRequest(ctx, LabourPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_uuid": labourUUID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetLabour sends GET /api/{api_version}/labours/{labour_uuid}.
func (l *Labours) GetLabour(ctx context.Context, labourUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourUUID, "labour_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendGetRequest(ctx, LabourPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_uuid": labourUUID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: LabourRecordShape})
}

// CancelLabour sends POST /api/{api_version}/labours/{labour_id}/cancel.
func (l *Labours) CancelLabour(ctx context.Context, labourID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourID, "labour_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendPostRequest(ctx, LabourCancelPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_id": labourID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: LabourRecordShape})
}

// ReactivateLabour sends POST /api/{api_version}/labours/{labour_id}/reactivate.
func (l *Labours) ReactivateLabour(ctx context.Context, labourID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourID, "labour_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendPostRequest(ctx, LabourReactivatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_id": labourID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: LabourRecordShape})
}

// GetLabourOrders sends GET /api/{api_version}/labours/{labour_uuid}/orders.
func (l *Labours) GetLabourOrders(ctx context.Context, labourUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourUUID, "labour_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendGetRequest(ctx, LabourOrdersPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_uuid": labourUUID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: LabourOrderListShape})
}

// UpdateLabourAvailability sends PATCH /api/{api_version}/labour_availability; it has no typed view, so read ApiResponse.Map.
func (l *Labours) UpdateLabourAvailability(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendPatchRequest(ctx, LabourAvailabilityPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetLabourProfileSessions sends GET /api/{api_version}/labour_profiles/{labour_profile_id}/sessions.
func (l *Labours) GetLabourProfileSessions(ctx context.Context, labourProfileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourProfileID, "labour_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendGetRequest(ctx, LabourProfileSessionsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_profile_id": labourProfileID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: LabourProfileSessionListShape})
}

// GetLabourProfileSession sends GET /api/{api_version}/labour_profiles/{labour_profile_id}/sessions/{session_id}.
func (l *Labours) GetLabourProfileSession(ctx context.Context, labourProfileID any, sessionID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourProfileID, "labour_profile_id"); err != nil {
		return nil, err
	}
	if err := requireValue(sessionID, "session_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendGetRequest(ctx, LabourProfileSessionPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"labour_profile_id": labourProfileID, "session_id": sessionID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: LabourProfileSessionViewShape})
}

// GetLabourCalendar sends GET /api/{api_version}/labours/{labour_id}/availability.
func (l *Labours) GetLabourCalendar(ctx context.Context, labourID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourID, "labour_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return l.requester.SendGetRequest(ctx, LabourCalendarPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_id": labourID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: LabourCalendarRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "labours"

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

// New binds the labours endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Labours { return &Labours{requester: requester} }

// Requester returns the requester this resource was bound to.
func (l *Labours) Requester() exhttp.Requester { return l.requester }
