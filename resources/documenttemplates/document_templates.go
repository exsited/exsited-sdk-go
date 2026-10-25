package documenttemplates

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// DocumentTemplates is the v4 document templates endpoints, bound to one client.
type DocumentTemplates struct {
	requester exhttp.Requester
}

// GetSettingsDocumentTemplates sends GET /api/{api_version}/settings/document_templates.
func (d *DocumentTemplates) GetSettingsDocumentTemplates(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return d.requester.SendGetRequest(ctx, SettingsDocumentTemplatesPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsDocumentTemplateListShape})
}

// CreateSettingsDocumentTemplate sends POST /api/{api_version}/settings/document_templates.
func (d *DocumentTemplates) CreateSettingsDocumentTemplate(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return d.requester.SendPostRequest(ctx, SettingsDocumentTemplatesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsDocumentTemplateRecordShape})
}

// GetSettingsDocumentTemplate sends GET /api/{api_version}/settings/document_templates/{document_template_id}.
func (d *DocumentTemplates) GetSettingsDocumentTemplate(ctx context.Context, documentTemplateID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(documentTemplateID, "document_template_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return d.requester.SendGetRequest(ctx, SettingsDocumentTemplatePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"document_template_id": documentTemplateID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsDocumentTemplateRecordShape})
}

// DeleteSettingsDocumentTemplate sends DELETE /api/{api_version}/settings/document_templates/{document_template_id}; it has no typed view, so read ApiResponse.Map.
func (d *DocumentTemplates) DeleteSettingsDocumentTemplate(ctx context.Context, documentTemplateID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(documentTemplateID, "document_template_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return d.requester.SendDeleteRequest(ctx, SettingsDocumentTemplatePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"document_template_id": documentTemplateID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsDocumentTemplate sends PATCH /api/{api_version}/settings/document_templates/{document_template_id}.
func (d *DocumentTemplates) UpdateSettingsDocumentTemplate(ctx context.Context, documentTemplateID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(documentTemplateID, "document_template_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return d.requester.SendPatchRequest(ctx, SettingsDocumentTemplatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"document_template_id": documentTemplateID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsDocumentTemplateRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "document_templates"

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

// requirePayload refuses a nil payload, a typed nil included, before any I/O.
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

// New binds the document templates endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *DocumentTemplates {
	return &DocumentTemplates{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (d *DocumentTemplates) Requester() exhttp.Requester { return d.requester }
