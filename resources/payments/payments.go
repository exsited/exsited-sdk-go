package payments

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Payments is the v4 payments endpoints, bound to one client.
type Payments struct {
	requester exhttp.Requester
}

// GetPayments sends GET /api/{api_version}/payments.
func (p *Payments) GetPayments(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PaymentsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PaymentListShape})
}

// GetPaymentsNotes sends GET /api/{api_version}/payments/notes.
func (p *Payments) GetPaymentsNotes(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PaymentsNotesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PaymentsNoteListShape})
}

// DeletePayment sends DELETE /api/{api_version}/payments/{payment_id}; it has no typed view, so read ApiResponse.Map.
func (p *Payments) DeletePayment(ctx context.Context, paymentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PaymentPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"payment_id": paymentID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// UpdatePayment sends PATCH /api/{api_version}/payments/{payment_id}.
func (p *Payments) UpdatePayment(ctx context.Context, paymentID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPatchRequest(ctx, PaymentPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"payment_id": paymentID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PaymentReadRecordShape})
}

// GetPayment sends GET /api/{api_version}/payments/{payment_id}.
func (p *Payments) GetPayment(ctx context.Context, paymentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PaymentPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"payment_id": paymentID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PaymentReadRecordShape})
}

// GetInvoicePayments sends GET /api/{api_version}/invoices/{invoice_id}/payments.
func (p *Payments) GetInvoicePayments(ctx context.Context, invoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, InvoicePaymentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"invoice_id": invoiceID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PaymentListShape})
}

// GetOrderPayments sends GET /api/{api_version}/orders/{order_id}/payments.
func (p *Payments) GetOrderPayments(ctx context.Context, orderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(orderID, "order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, OrderPaymentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"order_id": orderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PaymentListShape})
}

// GetPaymentNotes sends GET /api/{api_version}/payments/{payment_id}/notes.
func (p *Payments) GetPaymentNotes(ctx context.Context, paymentID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PaymentNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"payment_id": paymentID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PaymentNoteListShape})
}

// CreatePaymentNote sends POST /api/{api_version}/payments/{payment_id}/notes.
// The note travels as the form field note, and any files go beside it as multipart parts.
func (p *Payments) CreatePaymentNote(ctx context.Context, paymentID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PaymentNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"payment_id": paymentID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PaymentNoteReadRecordShape})
}

// GetPaymentNote sends GET /api/{api_version}/payments/{payment_id}/notes/{payment_note_id}.
func (p *Payments) GetPaymentNote(ctx context.Context, paymentID any, paymentNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(paymentNoteID, "payment_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PaymentNotePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"payment_id": paymentID, "payment_note_id": paymentNoteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PaymentNoteReadRecordShape})
}

// DeletePaymentNote sends DELETE /api/{api_version}/payments/{payment_id}/notes/{payment_note_id}; it has no typed view, so read ApiResponse.Map.
func (p *Payments) DeletePaymentNote(ctx context.Context, paymentID any, paymentNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(paymentNoteID, "payment_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PaymentNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"payment_id": paymentID, "payment_note_id": paymentNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetPaymentNoteFile sends GET /api/{api_version}/payments/{payment_id}/notes/{payment_note_id}/files/{payment_file_id}.
func (p *Payments) GetPaymentNoteFile(ctx context.Context, paymentID any, paymentNoteID any, paymentFileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(paymentNoteID, "payment_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(paymentFileID, "payment_file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PaymentNoteFilePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"payment_id": paymentID, "payment_note_id": paymentNoteID, "payment_file_id": paymentFileID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PaymentNoteFileReadRecordShape})
}

// CreatePayment sends POST /api/{api_version}/payments.
func (p *Payments) CreatePayment(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PaymentsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PaymentReadRecordShape})
}

// GetInvoicePayment sends GET /api/{api_version}/invoices/{invoice_id}/payments/{payment_id}.
func (p *Payments) GetInvoicePayment(ctx context.Context, invoiceID any, paymentID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(invoiceID, "invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, InvoicePaymentPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"invoice_id": invoiceID, "payment_id": paymentID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PaymentReadRecordShape})
}

// GetPaymentNoteFiles sends GET /api/{api_version}/payments/{payment_id}/notes/{payment_note_id}/files.
func (p *Payments) GetPaymentNoteFiles(ctx context.Context, paymentID any, paymentNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(paymentID, "payment_id"); err != nil {
		return nil, err
	}
	if err := requireValue(paymentNoteID, "payment_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PaymentNoteFilesPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"payment_id": paymentID, "payment_note_id": paymentNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: NoteFileListShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "payments"

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

// New binds the payments endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Payments { return &Payments{requester: requester} }

// Requester returns the requester this resource was bound to.
func (p *Payments) Requester() exhttp.Requester { return p.requester }
