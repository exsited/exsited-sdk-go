package returnmerchandiseauthorisations

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
	"github.com/exsited/exsited-sdk-go/v4/resources/invoices"
)

// ReturnMerchandiseAuthorisations is the v4 return merchandise authorisations endpoints, bound to one client.
type ReturnMerchandiseAuthorisations struct {
	requester exhttp.Requester
}

// GetReturnMerchandiseAuthorisations sends GET /api/{api_version}/return_merchandise_authorisations.
func (r *ReturnMerchandiseAuthorisations) GetReturnMerchandiseAuthorisations(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReturnMerchandiseAuthorisationsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ReturnMerchandiseAuthorisationListShape})
}

// CreateReturnMerchandiseAuthorisation sends POST /api/{api_version}/return_merchandise_authorisations.
func (r *ReturnMerchandiseAuthorisations) CreateReturnMerchandiseAuthorisation(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, ReturnMerchandiseAuthorisationsPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ReturnMerchandiseAuthorisationsRecordShape})
}

// GetReturnMerchandiseAuthorisation sends GET /api/{api_version}/return_merchandise_authorisations/{rma_id}.
func (r *ReturnMerchandiseAuthorisations) GetReturnMerchandiseAuthorisation(ctx context.Context, rmaID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaID, "rma_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReturnMerchandiseAuthorisationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rma_id": rmaID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ReturnMerchandiseAuthorisationReadRecordShape})
}

// GetInvoiceReturnMerchandiseAuthorisationReceive sends GET /api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations/{rma_code}/receive.
func (r *ReturnMerchandiseAuthorisations) GetInvoiceReturnMerchandiseAuthorisationReceive(ctx context.Context, invoiceID any, rmaID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(rmaID, "rma_code"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, InvoiceReturnMerchandiseAuthorisationReceivePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"invoice_id": invoiceID, "rma_code": rmaID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceReturnMerchandiseAuthorisationReceiveListShape})
}

// GetInvoiceReturnMerchandiseAuthorisationReceiveByUUID sends GET /api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations/{rma_code}/receive/{receive_uuid}.
func (r *ReturnMerchandiseAuthorisations) GetInvoiceReturnMerchandiseAuthorisationReceiveByUUID(ctx context.Context, invoiceID any, rmaID any, receiveUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(rmaID, "rma_code"); err != nil {
		return nil, err
	}
	if err := requireValue(receiveUUID, "receive_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, InvoiceReturnMerchandiseAuthorisationReceiveByUUIDPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"invoice_id": invoiceID, "rma_code": rmaID, "receive_uuid": receiveUUID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: InvoiceReturnMerchandiseAuthorisationReceiveRecordShape})
}

// CreateReturnMerchandiseAuthorisationNote sends POST /api/{api_version}/return_merchandise_authorisations/{rma_id}/notes; it has no typed view, so read ApiResponse.Map.
func (r *ReturnMerchandiseAuthorisations) CreateReturnMerchandiseAuthorisationNote(ctx context.Context, rmaID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaID, "rma_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, ReturnMerchandiseAuthorisationNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rma_id": rmaID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ReceiveReturnMerchandiseAuthorisation sends POST /api/{api_version}/return_merchandise_authorisations/{rma_id}/receive.
func (r *ReturnMerchandiseAuthorisations) ReceiveReturnMerchandiseAuthorisation(ctx context.Context, rmaID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaID, "rma_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, ReturnMerchandiseAuthorisationReceivePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rma_id": rmaID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: invoices.InvoiceReturnMerchandiseAuthorisationViewShape})
}

// DeleteReturnMerchandiseAuthorisation sends DELETE /api/{api_version}/return_merchandise_authorisations/{rma_id}; it has no typed view, so read ApiResponse.Map.
func (r *ReturnMerchandiseAuthorisations) DeleteReturnMerchandiseAuthorisation(ctx context.Context, rmaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaID, "rma_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendDeleteRequest(ctx, ReturnMerchandiseAuthorisationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rma_id": rmaID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateReturnMerchandiseAuthorisation sends PATCH /api/{api_version}/return_merchandise_authorisations/{rma_id}.
func (r *ReturnMerchandiseAuthorisations) UpdateReturnMerchandiseAuthorisation(ctx context.Context, rmaID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaID, "rma_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPatchRequest(ctx, ReturnMerchandiseAuthorisationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rma_id": rmaID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: invoices.InvoiceReturnMerchandiseAuthorisationViewShape})
}

// ApproveReturnMerchandiseAuthorisation sends POST /api/{api_version}/return_merchandise_authorisations/{rma_id}/approve.
func (r *ReturnMerchandiseAuthorisations) ApproveReturnMerchandiseAuthorisation(ctx context.Context, rmaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaID, "rma_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, ReturnMerchandiseAuthorisationApprovePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rma_id": rmaID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: invoices.InvoiceReturnMerchandiseAuthorisationViewShape})
}

// DeleteReturnMerchandiseAuthorisationNote sends DELETE /api/{api_version}/return_merchandise_authorisations/{rma_id}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (r *ReturnMerchandiseAuthorisations) DeleteReturnMerchandiseAuthorisationNote(ctx context.Context, rmaID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaID, "rma_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendDeleteRequest(ctx, ReturnMerchandiseAuthorisationNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rma_id": rmaID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteReturnMerchandiseAuthorisationNoteFile sends DELETE /api/{api_version}/return_merchandise_authorisations/{rma_id}/notes/{note_id}/files/{file_id}; it has no typed view, so read ApiResponse.Map.
func (r *ReturnMerchandiseAuthorisations) DeleteReturnMerchandiseAuthorisationNoteFile(ctx context.Context, rmaID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaID, "rma_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(fileID, "file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendDeleteRequest(ctx, ReturnMerchandiseAuthorisationNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rma_id": rmaID, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// RejectReturnMerchandiseAuthorisation sends POST /api/{api_version}/return_merchandise_authorisations/{rma_id}/reject.
func (r *ReturnMerchandiseAuthorisations) RejectReturnMerchandiseAuthorisation(ctx context.Context, rmaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaID, "rma_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, ReturnMerchandiseAuthorisationRejectPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rma_id": rmaID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: invoices.InvoiceReturnMerchandiseAuthorisationViewShape})
}

// GetReturnMerchandiseAuthorisationNotes sends GET /api/{api_version}/return_merchandise_authorisations/{rma_code}/notes; it has no typed view, so read ApiResponse.Map.
func (r *ReturnMerchandiseAuthorisations) GetReturnMerchandiseAuthorisationNotes(ctx context.Context, rmaCode any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaCode, "rma_code"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReturnMerchandiseAuthorisationNotesByCodePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rma_code": rmaCode}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetReturnMerchandiseAuthorisationNote sends GET /api/{api_version}/return_merchandise_authorisations/{rma_code}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (r *ReturnMerchandiseAuthorisations) GetReturnMerchandiseAuthorisationNote(ctx context.Context, rmaCode any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaCode, "rma_code"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReturnMerchandiseAuthorisationNoteByCodePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rma_code": rmaCode, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetReturnMerchandiseAuthorisationNoteFiles sends GET /api/{api_version}/return_merchandise_authorisations/{rma_code}/notes/{note_id}/files; it has no typed view, so read ApiResponse.Map.
func (r *ReturnMerchandiseAuthorisations) GetReturnMerchandiseAuthorisationNoteFiles(ctx context.Context, rmaCode any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaCode, "rma_code"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReturnMerchandiseAuthorisationNoteFilesByCodePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rma_code": rmaCode, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetReturnMerchandiseAuthorisationNoteFile sends GET /api/{api_version}/return_merchandise_authorisations/{rma_code}/notes/{note_id}/files/{file_id}; it has no typed view, so read ApiResponse.Map.
func (r *ReturnMerchandiseAuthorisations) GetReturnMerchandiseAuthorisationNoteFile(ctx context.Context, rmaCode any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaCode, "rma_code"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(fileID, "file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReturnMerchandiseAuthorisationNoteFileByCodePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rma_code": rmaCode, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetReturnMerchandiseAuthorisationReceive sends GET /api/{api_version}/return_merchandise_authorisations/{rma_code}/receive.
func (r *ReturnMerchandiseAuthorisations) GetReturnMerchandiseAuthorisationReceive(ctx context.Context, rmaCode any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaCode, "rma_code"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReturnMerchandiseAuthorisationReceiveByCodePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rma_code": rmaCode}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceReturnMerchandiseAuthorisationReceiveListShape})
}

// GetReturnMerchandiseAuthorisationReceiveByUUID sends GET /api/{api_version}/return_merchandise_authorisations/{rma_code}/receive/{receive_uuid}.
func (r *ReturnMerchandiseAuthorisations) GetReturnMerchandiseAuthorisationReceiveByUUID(ctx context.Context, rmaCode any, receiveUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rmaCode, "rma_code"); err != nil {
		return nil, err
	}
	if err := requireValue(receiveUUID, "receive_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReturnMerchandiseAuthorisationReceiveByUUIDPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rma_code": rmaCode, "receive_uuid": receiveUUID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: InvoiceReturnMerchandiseAuthorisationReceiveRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "return_merchandise_authorisations"

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

// New binds the return merchandise authorisations endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *ReturnMerchandiseAuthorisations {
	return &ReturnMerchandiseAuthorisations{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (r *ReturnMerchandiseAuthorisations) Requester() exhttp.Requester { return r.requester }
