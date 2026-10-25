package documenttemplates

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the DocumentTemplates method of the same name,
// which is the primary surface.

// GetSettingsDocumentTemplates is the module form of DocumentTemplates.GetSettingsDocumentTemplates.
func GetSettingsDocumentTemplates(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsDocumentTemplates(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateSettingsDocumentTemplate is the module form of DocumentTemplates.CreateSettingsDocumentTemplate.
func CreateSettingsDocumentTemplate(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateSettingsDocumentTemplate(ctx, payload, raiseOnError, responseFormat...)
}

// GetSettingsDocumentTemplate is the module form of DocumentTemplates.GetSettingsDocumentTemplate.
func GetSettingsDocumentTemplate(requester exhttp.Requester, ctx context.Context, documentTemplateID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetSettingsDocumentTemplate(ctx, documentTemplateID, raiseOnError, responseFormat...)
}

// DeleteSettingsDocumentTemplate is the module form of DocumentTemplates.DeleteSettingsDocumentTemplate.
func DeleteSettingsDocumentTemplate(requester exhttp.Requester, ctx context.Context, documentTemplateID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteSettingsDocumentTemplate(ctx, documentTemplateID, raiseOnError, responseFormat...)
}

// UpdateSettingsDocumentTemplate is the module form of DocumentTemplates.UpdateSettingsDocumentTemplate.
func UpdateSettingsDocumentTemplate(requester exhttp.Requester, ctx context.Context, documentTemplateID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateSettingsDocumentTemplate(ctx, documentTemplateID, payload, raiseOnError, responseFormat...)
}
