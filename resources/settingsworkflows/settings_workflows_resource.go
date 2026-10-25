package settingsworkflows

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the SettingsWorkflows method of the same name,
// which is the primary surface.

// GetSettingsWorkflows is the module form of SettingsWorkflows.GetSettingsWorkflows.
func GetSettingsWorkflows(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsWorkflows(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetSettingsWorkflow is the module form of SettingsWorkflows.GetSettingsWorkflow.
func GetSettingsWorkflow(requester exhttp.Requester, ctx context.Context, workflowID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsWorkflow(ctx, workflowID, raiseOnError, responseFormat...)
}
