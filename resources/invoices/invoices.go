package invoices

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
	"github.com/exsited/exsited-sdk-go/v4/resources/creditnoteapplications"
)

// Invoices is the v4 invoices endpoints, bound to one client.
type Invoices struct {
	requester exhttp.Requester
}

// GetInvoices sends GET /api/{api_version}/invoices.
func (i *Invoices) GetInvoices(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoicesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceListShape})
}

// CreateInvoice sends POST /api/{api_version}/invoices.
func (i *Invoices) CreateInvoice(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, InvoicesPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceRecordShape})
}

// DeleteInvoice sends DELETE /api/{api_version}/invoices/{invoice_id}.
func (i *Invoices) DeleteInvoice(ctx context.Context, invoiceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendDeleteRequest(ctx, InvoicePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, RaiseOnError: raiseOnError,
		ResponseFormat: format, DTOShape: InvoiceRecordShape})
}

// GetInvoice sends GET /api/{api_version}/invoices/{invoice_id}.
func (i *Invoices) GetInvoice(ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoicePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceRecordShape})
}

// AmendInvoice sends POST /api/{api_version}/invoices/{invoice_id}/amend.
func (i *Invoices) AmendInvoice(ctx context.Context, invoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, InvoiceAmendPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceRecordShape})
}

// CancelInvoice sends POST /api/{api_version}/invoices/{invoice_id}/cancel.
func (i *Invoices) CancelInvoice(ctx context.Context, invoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, InvoiceCancelPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceRecordShape})
}

// UpdateInvoiceInformation sends PATCH /api/{api_version}/invoices/{invoice_id}/information.
func (i *Invoices) UpdateInvoiceInformation(ctx context.Context, invoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPatchRequest(ctx, InvoiceInformationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceRecordShape})
}

// GetInvoiceInformation sends GET /api/{api_version}/invoices/{invoice_id}/information.
func (i *Invoices) GetInvoiceInformation(ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceInformationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceRecordShape})
}

// GetInvoiceItemFulfillments sends GET /api/{api_version}/invoices/{invoice_id}/item_fulfillments.
func (i *Invoices) GetInvoiceItemFulfillments(ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceItemFulfillmentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceItemFulfillmentListShape})
}

// GetInvoiceNotes sends GET /api/{api_version}/invoices/{invoice_id}/notes.
func (i *Invoices) GetInvoiceNotes(ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceNoteListShape})
}

// CreateInvoiceNote sends POST /api/{api_version}/invoices/{invoice_id}/notes.
// The note travels as the form field note, and any files go beside it as multipart parts.
func (i *Invoices) CreateInvoiceNote(ctx context.Context, invoiceID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, InvoiceNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceNoteViewShape})
}

// GetInvoicePDF sends GET /api/{api_version}/invoices/{invoice_id}/pdf.
func (i *Invoices) GetInvoicePDF(ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoicePDFPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoicePDFViewShape})
}

// ReactivateInvoice sends POST /api/{api_version}/invoices/{invoice_id}/reactivate.
func (i *Invoices) ReactivateInvoice(ctx context.Context, invoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, InvoiceReactivatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceRecordShape})
}

// GetInvoiceReturnMerchandiseAuthorisations sends GET /api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations.
func (i *Invoices) GetInvoiceReturnMerchandiseAuthorisations(ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceReturnMerchandiseAuthorisationsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceReturnMerchandiseAuthorisationListShape})
}

// GetOrderInvoices sends GET /api/{api_version}/orders/{order_id}/invoices.
func (i *Invoices) GetOrderInvoices(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, OrderInvoicesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceListShape})
}

// CreateOrderInvoice sends POST /api/{api_version}/orders/{order_id}/invoices.
func (i *Invoices) CreateOrderInvoice(ctx context.Context, orderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendPostRequest(ctx, OrderInvoicesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceRecordShape})
}

// GetInvoiceItemFulfillment sends GET /api/{api_version}/invoices/{invoice_id}/item_fulfillments/{fulfillment_id}.
func (i *Invoices) GetInvoiceItemFulfillment(ctx context.Context, invoiceID any, fulfillmentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(fulfillmentID, "fulfillment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceItemFulfillmentPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"invoice_id": invoiceID, "fulfillment_id": fulfillmentID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceItemFulfillmentViewShape})
}

// GetInvoiceNote sends GET /api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}.
func (i *Invoices) GetInvoiceNote(ctx context.Context, invoiceID any, invoiceNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(invoiceNoteID, "invoice_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceNotePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"invoice_id": invoiceID, "invoice_note_id": invoiceNoteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceNoteViewShape})
}

// DeleteInvoiceNote sends DELETE /api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}; it has no typed view, so read ApiResponse.Map.
func (i *Invoices) DeleteInvoiceNote(ctx context.Context, invoiceID any, invoiceNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(invoiceNoteID, "invoice_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendDeleteRequest(ctx, InvoiceNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"invoice_id": invoiceID, "invoice_note_id": invoiceNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetInvoiceReturnMerchandiseAuthorisation sends GET /api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations/{rma_id}.
func (i *Invoices) GetInvoiceReturnMerchandiseAuthorisation(ctx context.Context, invoiceID any, rmaID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(rmaID, "rma_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceReturnMerchandiseAuthorisationPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"invoice_id": invoiceID, "rma_id": rmaID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceReturnMerchandiseAuthorisationViewShape})
}

// GetInvoiceNoteFiles sends GET /api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}/files.
func (i *Invoices) GetInvoiceNoteFiles(ctx context.Context, invoiceID any, invoiceNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(invoiceNoteID, "invoice_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceNoteFilesPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"invoice_id": invoiceID, "invoice_note_id": invoiceNoteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceNoteFileListShape})
}

// GetInvoiceNoteFile sends GET /api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}/files/{invoice_file_id}.
func (i *Invoices) GetInvoiceNoteFile(ctx context.Context, invoiceID any, invoiceNoteID any, invoiceFileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(invoiceNoteID, "invoice_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(invoiceFileID, "invoice_file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceNoteFilePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"invoice_id": invoiceID, "invoice_note_id": invoiceNoteID, "invoice_file_id": invoiceFileID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceNoteFileViewShape})
}

// GetInvoiceKPIs sends GET /api/{api_version}/invoices/{invoice_id}/kpis.
func (i *Invoices) GetInvoiceKPIs(ctx context.Context, invoiceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceKPIsPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"invoice_id": invoiceID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceKpisReadRecordShape})
}

// GetInvoiceLines sends GET /api/{api_version}/invoices/{invoice_id}/lines.
func (i *Invoices) GetInvoiceLines(ctx context.Context, invoiceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceLinesPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"invoice_id": invoiceID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceLinesRecordShape})
}

// GetInvoiceLine sends GET /api/{api_version}/invoices/{invoice_id}/lines/{line_id}.
func (i *Invoices) GetInvoiceLine(ctx context.Context, invoiceID any, lineID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(lineID, "line_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceLinePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"invoice_id": invoiceID, "line_id": lineID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceLineReadRecordShape})
}

// GetInvoiceCreditNoteApplication sends GET /api/{api_version}/invoices/{invoice_id}/credit_note_applications/{credit_note_application_id}.
func (i *Invoices) GetInvoiceCreditNoteApplication(ctx context.Context, invoiceID any, creditNoteApplicationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(creditNoteApplicationID, "credit_note_application_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceCreditNoteApplicationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"invoice_id": invoiceID, "credit_note_application_id": creditNoteApplicationID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: creditnoteapplications.CreditNoteApplicationReadRecordShape})
}

// GetInvoiceCustomObjects sends GET /api/{api_version}/invoices/{invoice_id}/custom_objects; it has no typed view, so read ApiResponse.Map.
func (i *Invoices) GetInvoiceCustomObjects(ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceCustomObjectsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetInvoiceCustomObject sends GET /api/{api_version}/invoices/{invoice_id}/custom_objects/{custom_object_id}; it has no typed view, so read ApiResponse.Map.
func (i *Invoices) GetInvoiceCustomObject(ctx context.Context, invoiceID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceCustomObjectPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"invoice_id": invoiceID, "custom_object_id": customObjectID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetInvoiceCustomObjectInstance sends GET /api/{api_version}/invoices/{invoice_id}/custom_objects/{custom_object_id}/instances/{instance_id}; it has no typed view, so read ApiResponse.Map.
func (i *Invoices) GetInvoiceCustomObjectInstance(ctx context.Context, invoiceID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(instanceID, "instance_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendGetRequest(ctx, InvoiceCustomObjectInstancePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"invoice_id": invoiceID, "custom_object_id": customObjectID, "instance_id": instanceID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteInvoiceNoteFile sends DELETE /api/{api_version}/invoices/{invoice_id}/notes/{invoice_note_id}/files/{invoice_file_id}; it has no typed view, so read ApiResponse.Map.
func (i *Invoices) DeleteInvoiceNoteFile(ctx context.Context, invoiceID any, invoiceNoteID any, invoiceFileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(invoiceNoteID, "invoice_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(invoiceFileID, "invoice_file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return i.requester.SendDeleteRequest(ctx, InvoiceNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"invoice_id": invoiceID, "invoice_note_id": invoiceNoteID, "invoice_file_id": invoiceFileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "invoices"

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

// requireNote refuses an empty note before any I/O; the text is the whole body of a note write.
func requireNote(note string) error {
	if note == "" {
		return exerrors.NewUsageError("note is required: the note text is the body of this request")
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

// New binds the invoices endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Invoices { return &Invoices{requester: requester} }

// Requester returns the requester this resource was bound to.
func (i *Invoices) Requester() exhttp.Requester { return i.requester }
