package purchaseinvoices

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchasecreditnoteapplications"
)

// PurchaseInvoices is the v4 purchase invoices endpoints, bound to one client.
type PurchaseInvoices struct {
	requester exhttp.Requester
}

// GetPurchaseInvoices sends GET /api/{api_version}/purchase_invoices.
func (p *PurchaseInvoices) GetPurchaseInvoices(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoicesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceListShape})
}

// CreatePurchaseInvoice sends POST /api/{api_version}/purchase_invoices.
func (p *PurchaseInvoices) CreatePurchaseInvoice(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseInvoicesPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceReadRecordShape})
}

// GetPurchaseInvoicesNotes sends GET /api/{api_version}/purchase_invoices/notes.
func (p *PurchaseInvoices) GetPurchaseInvoicesNotes(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoicesNotesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoicesNoteListShape})
}

// DeletePurchaseInvoice sends DELETE /api/{api_version}/purchase_invoices/{purchase_invoice_id}; it has no typed view, so read ApiResponse.Map.
func (p *PurchaseInvoices) DeletePurchaseInvoice(ctx context.Context, purchaseInvoiceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PurchaseInvoicePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_invoice_id": purchaseInvoiceID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetPurchaseInvoice sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}.
func (p *PurchaseInvoices) GetPurchaseInvoice(ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoicePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_invoice_id": purchaseInvoiceID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceReadRecordShape})
}

// AmendPurchaseInvoice sends POST /api/{api_version}/purchase_invoices/{purchase_invoice_id}/amend.
func (p *PurchaseInvoices) AmendPurchaseInvoice(ctx context.Context, purchaseInvoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
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
	return p.requester.SendPostRequest(ctx, PurchaseInvoiceAmendPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_invoice_id": purchaseInvoiceID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceReadRecordShape})
}

// CancelPurchaseInvoice sends POST /api/{api_version}/purchase_invoices/{purchase_invoice_id}/cancel.
func (p *PurchaseInvoices) CancelPurchaseInvoice(ctx context.Context, purchaseInvoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
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
	return p.requester.SendPostRequest(ctx, PurchaseInvoiceCancelPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_invoice_id": purchaseInvoiceID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceCancelRecordShape})
}

// GetPurchaseInvoiceInformation sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/information.
func (p *PurchaseInvoices) GetPurchaseInvoiceInformation(ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoiceInformationPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_invoice_id": purchaseInvoiceID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceRecordShape})
}

// UpdatePurchaseInvoiceInformation sends PATCH /api/{api_version}/purchase_invoices/{purchase_invoice_id}/information.
func (p *PurchaseInvoices) UpdatePurchaseInvoiceInformation(ctx context.Context, purchaseInvoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
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
	return p.requester.SendPatchRequest(ctx, PurchaseInvoiceInformationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_invoice_id": purchaseInvoiceID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceInformationRecordShape})
}

// GetPurchaseInvoiceLines sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/lines.
func (p *PurchaseInvoices) GetPurchaseInvoiceLines(ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoiceLinesPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_invoice_id": purchaseInvoiceID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceLinesRecordShape})
}

// GetPurchaseInvoiceNotes sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes.
func (p *PurchaseInvoices) GetPurchaseInvoiceNotes(ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoiceNotesPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_invoice_id": purchaseInvoiceID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceNoteListShape})
}

// CreatePurchaseInvoiceNote sends POST /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes.
// The note travels as the form field note, and any files go beside it as multipart parts.
func (p *PurchaseInvoices) CreatePurchaseInvoiceNote(ctx context.Context, purchaseInvoiceID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseInvoiceNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_invoice_id": purchaseInvoiceID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceNotesRecordShape})
}

// GetPurchaseInvoicePDF sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/pdf; it has no typed view, so read ApiResponse.Map.
func (p *PurchaseInvoices) GetPurchaseInvoicePDF(ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoicePDFPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_invoice_id": purchaseInvoiceID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetPurchaseInvoicePurchaseCreditNoteApplications sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_credit_note_applications.
func (p *PurchaseInvoices) GetPurchaseInvoicePurchaseCreditNoteApplications(ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoicePurchaseCreditNoteApplicationsPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_invoice_id": purchaseInvoiceID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoicePurchaseCreditNoteApplicationListShape})
}

// GetPurchaseInvoicePurchasePayments sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_payments.
func (p *PurchaseInvoices) GetPurchaseInvoicePurchasePayments(ctx context.Context, purchaseInvoiceID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoicePurchasePaymentsPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_invoice_id": purchaseInvoiceID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoicePurchasePaymentListShape})
}

// ReactivatePurchaseInvoice sends POST /api/{api_version}/purchase_invoices/{purchase_invoice_id}/reactivate.
func (p *PurchaseInvoices) ReactivatePurchaseInvoice(ctx context.Context, purchaseInvoiceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
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
	return p.requester.SendPostRequest(ctx, PurchaseInvoiceReactivatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_invoice_id": purchaseInvoiceID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceReadRecordShape})
}

// GetPurchaseInvoiceLine sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/lines/{purchase_invoice_line_uuid}.
func (p *PurchaseInvoices) GetPurchaseInvoiceLine(ctx context.Context, purchaseInvoiceID any, purchaseInvoiceLineUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(purchaseInvoiceLineUUID, "purchase_invoice_line_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoiceLinePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_invoice_id": purchaseInvoiceID, "purchase_invoice_line_uuid": purchaseInvoiceLineUUID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceLineRecordShape})
}

// UpdatePurchaseInvoiceNote sends PATCH /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}.
// The note travels as the form field note, and any files go beside it as multipart parts.
func (p *PurchaseInvoices) UpdatePurchaseInvoiceNote(ctx context.Context, purchaseInvoiceID any, noteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPatchRequest(ctx, PurchaseInvoiceNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_invoice_id": purchaseInvoiceID, "note_id": noteID},
		FormData:   map[string]any{"note": note}, Files: files, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceNoteReadRecordShape})
}

// GetPurchaseInvoiceNote sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}.
func (p *PurchaseInvoices) GetPurchaseInvoiceNote(ctx context.Context, purchaseInvoiceID any, noteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoiceNotePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_invoice_id": purchaseInvoiceID, "note_id": noteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceNoteReadRecordShape})
}

// DeletePurchaseInvoiceNote sends DELETE /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (p *PurchaseInvoices) DeletePurchaseInvoiceNote(ctx context.Context, purchaseInvoiceID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PurchaseInvoiceNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_invoice_id": purchaseInvoiceID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetPurchaseInvoiceNoteFiles sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}/files.
func (p *PurchaseInvoices) GetPurchaseInvoiceNoteFiles(ctx context.Context, purchaseInvoiceID any, noteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoiceNoteFilesPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_invoice_id": purchaseInvoiceID, "note_id": noteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceNoteFileListShape})
}

// GetPurchaseInvoiceNoteFile sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/notes/{note_id}/files/{file_id}.
func (p *PurchaseInvoices) GetPurchaseInvoiceNoteFile(ctx context.Context, purchaseInvoiceID any, noteID any, fileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
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
	return p.requester.SendGetRequest(ctx, PurchaseInvoiceNoteFilePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_invoice_id": purchaseInvoiceID, "note_id": noteID, "file_id": fileID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseInvoiceNoteFileReadRecordShape})
}

// GetPurchaseInvoicePurchaseCreditNoteApplication sends GET /api/{api_version}/purchase_invoices/{purchase_invoice_id}/purchase_credit_note_applications/{application_id}.
func (p *PurchaseInvoices) GetPurchaseInvoicePurchaseCreditNoteApplication(ctx context.Context, purchaseInvoiceID any, applicationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseInvoiceID, "purchase_invoice_id"); err != nil {
		return nil, err
	}
	if err := requireValue(applicationID, "application_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseInvoicePurchaseCreditNoteApplicationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_invoice_id": purchaseInvoiceID, "application_id": applicationID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: purchasecreditnoteapplications.PurchaseCreditNoteApplicationReadRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "purchase_invoices"

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

// New binds the purchase invoices endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *PurchaseInvoices {
	return &PurchaseInvoices{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (p *PurchaseInvoices) Requester() exhttp.Requester { return p.requester }
