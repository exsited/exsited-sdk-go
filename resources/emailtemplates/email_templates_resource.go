package emailtemplates

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the EmailTemplates method of the same name,
// which is the primary surface.

// GetSettingsEmailTemplates is the module form of EmailTemplates.GetSettingsEmailTemplates.
func GetSettingsEmailTemplates(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsEmailTemplates(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsEmailTemplate is the module form of EmailTemplates.CreateSettingsEmailTemplate.
func CreateSettingsEmailTemplate(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsEmailTemplate(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsEmailTemplate is the module form of EmailTemplates.GetSettingsEmailTemplate.
func GetSettingsEmailTemplate(requester exhttp.Requester, ctx context.Context, emailTemplateID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsEmailTemplate(ctx, emailTemplateID, raiseOnError, responseFormat...)
}

// DeleteSettingsEmailTemplate is the module form of EmailTemplates.DeleteSettingsEmailTemplate.
func DeleteSettingsEmailTemplate(requester exhttp.Requester, ctx context.Context, emailTemplateID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsEmailTemplate(ctx, emailTemplateID, raiseOnError, responseFormat...)
}

// UpdateSettingsEmailTemplate is the module form of EmailTemplates.UpdateSettingsEmailTemplate.
func UpdateSettingsEmailTemplate(requester exhttp.Requester, ctx context.Context, emailTemplateID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsEmailTemplate(ctx, emailTemplateID, payload, raiseOnError, responseFormat...)
}
