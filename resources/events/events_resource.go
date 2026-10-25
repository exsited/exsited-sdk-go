package events

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Events method of the same name,
// which is the primary surface.

// GetEvents is the module form of Events.GetEvents.
func GetEvents(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetEvents(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetEvent is the module form of Events.GetEvent.
func GetEvent(requester exhttp.Requester, ctx context.Context, eventUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetEvent(ctx, eventUUID, queryParams, raiseOnError, responseFormat...)
}

// CreateEventExecute is the module form of Events.CreateEventExecute.
func CreateEventExecute(requester exhttp.Requester, ctx context.Context, eventUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateEventExecute(ctx, eventUUID, payload, raiseOnError, responseFormat...)
}

// CreateEventRemove is the module form of Events.CreateEventRemove.
func CreateEventRemove(requester exhttp.Requester, ctx context.Context, eventUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateEventRemove(ctx, eventUUID, payload, raiseOnError, responseFormat...)
}
