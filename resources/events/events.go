package events

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Events is the v4 events endpoints, bound to one client.
type Events struct {
	requester exhttp.Requester
}

// GetEvents sends GET /api/{api_version}/events.
func (e *Events) GetEvents(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return e.requester.SendGetRequest(ctx, EventsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: EventListShape})
}

// GetEvent sends GET /api/{api_version}/events/{event_uuid}.
func (e *Events) GetEvent(ctx context.Context, eventUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(eventUUID, "event_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return e.requester.SendGetRequest(ctx, EventPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"event_uuid": eventUUID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: EventRecordShape})
}

// CreateEventExecute sends POST /api/{api_version}/events/{event_uuid}/execute; it has no typed view, so read ApiResponse.Map.
func (e *Events) CreateEventExecute(ctx context.Context, eventUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(eventUUID, "event_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return e.requester.SendPostRequest(ctx, EventExecutePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"event_uuid": eventUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateEventRemove sends POST /api/{api_version}/events/{event_uuid}/remove.
func (e *Events) CreateEventRemove(ctx context.Context, eventUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(eventUUID, "event_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return e.requester.SendPostRequest(ctx, EventRemovePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"event_uuid": eventUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: EventRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "events"

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

// New binds the events endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Events { return &Events{requester: requester} }

// Requester returns the requester this resource was bound to.
func (e *Events) Requester() exhttp.Requester { return e.requester }
