package invoiceterms

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// InvoiceTerms is the v4 invoice terms endpoints, bound to one client.
type InvoiceTerms struct {
	requester exhttp.Requester
}

// GetSettingsInvoiceTerms sends GET /api/{api_version}/settings/invoice_terms.
func (i *InvoiceTerms) GetSettingsInvoiceTerms(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, SettingsInvoiceTermsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsInvoiceTermListShape})
}

// CreateSettingsInvoiceTerm sends POST /api/{api_version}/settings/invoice_terms.
func (i *InvoiceTerms) CreateSettingsInvoiceTerm(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, SettingsInvoiceTermsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsInvoiceTermWriteRecordShape})
}

// GetSettingsInvoiceTerm sends GET /api/{api_version}/settings/invoice_terms/{invoice_term_id}.
func (i *InvoiceTerms) GetSettingsInvoiceTerm(ctx context.Context, invoiceTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceTermID, "invoice_term_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, SettingsInvoiceTermPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"invoice_term_id": invoiceTermID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsInvoiceTermRecordShape})
}

// DeleteSettingsInvoiceTerm sends DELETE /api/{api_version}/settings/invoice_terms/{invoice_term_id}; it has no typed view, so read ApiResponse.Map.
func (i *InvoiceTerms) DeleteSettingsInvoiceTerm(ctx context.Context, invoiceTermID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceTermID, "invoice_term_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendDeleteRequest(ctx, SettingsInvoiceTermPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"invoice_term_id": invoiceTermID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsInvoiceTerm sends PATCH /api/{api_version}/settings/invoice_terms/{invoice_term_id}.
func (i *InvoiceTerms) UpdateSettingsInvoiceTerm(ctx context.Context, invoiceTermID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceTermID, "invoice_term_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, SettingsInvoiceTermPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_term_id": invoiceTermID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsInvoiceTermWriteRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "invoice_terms"

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

// New binds the invoice terms endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *InvoiceTerms { return &InvoiceTerms{requester: requester} }

// Requester returns the requester this resource was bound to.
func (i *InvoiceTerms) Requester() exhttp.Requester { return i.requester }
