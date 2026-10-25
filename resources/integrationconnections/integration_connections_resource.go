package integrationconnections

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the IntegrationConnections method of the same name,
// which is the primary surface.

// GetIntegrationConnections is the module form of IntegrationConnections.GetIntegrationConnections.
func GetIntegrationConnections(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetIntegrationConnections(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetIntegrationConnection is the module form of IntegrationConnections.GetIntegrationConnection.
func GetIntegrationConnection(requester exhttp.Requester, ctx context.Context, integrationConnectionID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetIntegrationConnection(ctx, integrationConnectionID, queryParams, raiseOnError, responseFormat...)
}
