package proformainvoices

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// ProformaInvoices is the v4 proforma invoices endpoints, bound to one client.
type ProformaInvoices struct {
	requester exhttp.Requester
}

// GetProformaInvoices sends GET /api/{api_version}/proforma_invoices; it has no typed view, so read ApiResponse.Map.
func (p *ProformaInvoices) GetProformaInvoices(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, ProformaInvoicesPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetProformaInvoice sends GET /api/{api_version}/proforma_invoices/{proforma_id}; it has no typed view, so read ApiResponse.Map.
func (p *ProformaInvoices) GetProformaInvoice(ctx context.Context, proformaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(proformaID, "proforma_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, ProformaInvoicePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"proforma_id": proformaID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteProformaInvoice sends DELETE /api/{api_version}/proforma_invoices/{proforma_id}; it has no typed view, so read ApiResponse.Map.
func (p *ProformaInvoices) DeleteProformaInvoice(ctx context.Context, proformaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(proformaID, "proforma_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, ProformaInvoicePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"proforma_id": proformaID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateProformaInvoicePayment sends POST /api/{api_version}/proforma_invoices/{proforma_id}/payments; it has no typed view, so read ApiResponse.Map.
func (p *ProformaInvoices) CreateProformaInvoicePayment(ctx context.Context, proformaID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(proformaID, "proforma_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, ProformaInvoicePaymentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"proforma_id": proformaID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "proforma_invoices"

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

// New binds the proforma invoices endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *ProformaInvoices {
	return &ProformaInvoices{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (p *ProformaInvoices) Requester() exhttp.Requester { return p.requester }
