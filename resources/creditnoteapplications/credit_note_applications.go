package creditnoteapplications

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// CreditNoteApplications is the v4 credit note applications endpoints, bound to one client.
type CreditNoteApplications struct {
	requester exhttp.Requester
}

// GetCreditNoteApplications sends GET /api/{api_version}/credit_note_applications.
func (c *CreditNoteApplications) GetCreditNoteApplications(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CreditNoteApplicationsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: CreditNoteApplicationListShape})
}

// GetCreditNoteApplication sends GET /api/{api_version}/credit_note_applications/{credit_note_application_id}.
func (c *CreditNoteApplications) GetCreditNoteApplication(ctx context.Context, creditNoteApplicationID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(creditNoteApplicationID, "credit_note_application_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, CreditNoteApplicationPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"credit_note_application_id": creditNoteApplicationID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: CreditNoteApplicationReadRecordShape})
}

// GetInvoiceCreditNoteApplications sends GET /api/{api_version}/invoices/{invoice_id}/credit_note_applications.
func (c *CreditNoteApplications) GetInvoiceCreditNoteApplications(ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, InvoiceCreditNoteApplicationsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceCreditNoteApplicationsRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "credit_note_applications"

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

// New binds the credit note applications endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *CreditNoteApplications {
	return &CreditNoteApplications{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (c *CreditNoteApplications) Requester() exhttp.Requester { return c.requester }
