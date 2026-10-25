package purchasepayments

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchasecreditnotes"
)

// PurchasePayments is the v4 purchase payments endpoints, bound to one client.
type PurchasePayments struct {
	requester exhttp.Requester
}

// GetPurchasePayments sends GET /api/{api_version}/purchase_payments.
func (p *PurchasePayments) GetPurchasePayments(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchasePaymentsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchasePaymentListShape})
}

// CreatePurchasePayment sends POST /api/{api_version}/purchase_payments.
func (p *PurchasePayments) CreatePurchasePayment(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchasePaymentsPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchasePaymentsRecordShape})
}

// DeletePurchasePayment sends DELETE /api/{api_version}/purchase_payments/{purchase_payment_id}; it has no typed view, so read ApiResponse.Map.
func (p *PurchasePayments) DeletePurchasePayment(ctx context.Context, purchasePaymentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PurchasePaymentPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_payment_id": purchasePaymentID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetPurchasePayment sends GET /api/{api_version}/purchase_payments/{purchase_payment_id}.
func (p *PurchasePayments) GetPurchasePayment(ctx context.Context, purchasePaymentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchasePaymentPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_payment_id": purchasePaymentID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchasePaymentReadRecordShape})
}

// UpdatePurchasePayment sends PATCH /api/{api_version}/purchase_payments/{purchase_payment_id}.
func (p *PurchasePayments) UpdatePurchasePayment(ctx context.Context, purchasePaymentID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPatchRequest(ctx, PurchasePaymentPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_payment_id": purchasePaymentID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchasePaymentUpdatedRecordShape})
}

// CreatePurchaseInvoicePurchasePayment sends POST /api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_payments.
func (p *PurchasePayments) CreatePurchaseInvoicePurchasePayment(ctx context.Context, purchaseInvoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseInvoicePurchasePaymentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_invoice_id": purchaseInvoiceID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PurchasePaymentsRecordShape})
}

// GetPurchasePaymentNotes sends GET /api/{api_version}/purchase_payments/{purchase_payment_id}/notes.
func (p *PurchasePayments) GetPurchasePaymentNotes(ctx context.Context, purchasePaymentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchasePaymentNotesPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_payment_id": purchasePaymentID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchasePaymentNoteListShape})
}

// CreatePurchasePaymentNote sends POST /api/{api_version}/purchase_payments/{purchase_payment_id}/notes.
// The note travels as the form field note, and any files go beside it as multipart parts.
func (p *PurchasePayments) CreatePurchasePaymentNote(ctx context.Context, purchasePaymentID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchasePaymentNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_payment_id": purchasePaymentID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchasePaymentNotesRecordShape})
}

// GetPurchasePaymentNote sends GET /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}.
func (p *PurchasePayments) GetPurchasePaymentNote(ctx context.Context, purchasePaymentID any, ppNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(ppNoteID, "pp_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchasePaymentNotePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_payment_id": purchasePaymentID, "pp_note_id": ppNoteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchasePaymentNoteReadRecordShape})
}

// DeletePurchasePaymentNote sends DELETE /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}; it has no typed view, so read ApiResponse.Map.
func (p *PurchasePayments) DeletePurchasePaymentNote(ctx context.Context, purchasePaymentID any, ppNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(ppNoteID, "pp_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PurchasePaymentNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_payment_id": purchasePaymentID, "pp_note_id": ppNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetPurchasePaymentNoteFiles sends GET /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}/files.
func (p *PurchasePayments) GetPurchasePaymentNoteFiles(ctx context.Context, purchasePaymentID any, ppNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(ppNoteID, "pp_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchasePaymentNoteFilesPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_payment_id": purchasePaymentID, "pp_note_id": ppNoteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchasePaymentNoteFileListShape})
}

// GetPurchasePaymentNoteFile sends GET /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}/files/{file_id}.
func (p *PurchasePayments) GetPurchasePaymentNoteFile(ctx context.Context, purchasePaymentID any, ppNoteID any, fileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(ppNoteID, "pp_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(fileID, "file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchasePaymentNoteFilePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_payment_id": purchasePaymentID, "pp_note_id": ppNoteID, "file_id": fileID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchasePaymentNoteFileReadRecordShape})
}

// GetPurchasePaymentPurchaseCreditNotes sends GET /api/{api_version}/purchase_payments/{purchase_payment_id}/purchase_credit_notes.
func (p *PurchasePayments) GetPurchasePaymentPurchaseCreditNotes(ctx context.Context, purchasePaymentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchasePaymentPurchaseCreditNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_payment_id": purchasePaymentID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: purchasecreditnotes.PurchaseCreditNoteListShape})
}

// UpdatePurchasePaymentNote sends PATCH /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}.
func (p *PurchasePayments) UpdatePurchasePaymentNote(ctx context.Context, purchasePaymentID any, ppNoteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(ppNoteID, "pp_note_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPatchRequest(ctx, PurchasePaymentNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_payment_id": purchasePaymentID, "pp_note_id": ppNoteID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PurchasePaymentNoteUpdatedRecordShape})
}

// DeletePurchasePaymentNoteFile sends DELETE /api/{api_version}/purchase_payments/{purchase_payment_id}/notes/{pp_note_id}/files/{file_id}; it has no typed view, so read ApiResponse.Map.
func (p *PurchasePayments) DeletePurchasePaymentNoteFile(ctx context.Context, purchasePaymentID any, ppNoteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchasePaymentID, "purchase_payment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(ppNoteID, "pp_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(fileID, "file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PurchasePaymentNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_payment_id": purchasePaymentID, "pp_note_id": ppNoteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "purchase_payments"

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

// New binds the purchase payments endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *PurchasePayments {
	return &PurchasePayments{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (p *PurchasePayments) Requester() exhttp.Requester { return p.requester }
