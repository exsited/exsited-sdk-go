package purchaseorders

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// PurchaseOrders is the v4 purchase orders endpoints, bound to one client.
type PurchaseOrders struct {
	requester exhttp.Requester
}

// GetPurchaseOrders sends GET /api/{api_version}/purchase_orders.
func (p *PurchaseOrders) GetPurchaseOrders(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrdersPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderListShape})
}

// CreatePurchaseOrder sends POST /api/{api_version}/purchase_orders.
func (p *PurchaseOrders) CreatePurchaseOrder(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseOrdersPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrdersRecordShape})
}

// GetPurchaseOrdersNotes sends GET /api/{api_version}/purchase_orders/notes.
func (p *PurchaseOrders) GetPurchaseOrdersNotes(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrdersNotesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrdersNoteListShape})
}

// DeletePurchaseOrder sends DELETE /api/{api_version}/purchase_orders/{purchase_order_id}; it has no typed view, so read ApiResponse.Map.
func (p *PurchaseOrders) DeletePurchaseOrder(ctx context.Context, purchaseOrderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PurchaseOrderPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// UpdatePurchaseOrder sends PATCH /api/{api_version}/purchase_orders/{purchase_order_id}.
func (p *PurchaseOrders) UpdatePurchaseOrder(ctx context.Context, purchaseOrderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPatchRequest(ctx, PurchaseOrderPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderUpdatedRecordShape})
}

// GetPurchaseOrder sends GET /api/{api_version}/purchase_orders/{purchase_order_id}.
func (p *PurchaseOrders) GetPurchaseOrder(ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrderPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PurchaseOrderReadRecordShape})
}

// CancelPurchaseOrder sends POST /api/{api_version}/purchase_orders/{purchase_order_id}/cancel.
func (p *PurchaseOrders) CancelPurchaseOrder(ctx context.Context, purchaseOrderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseOrderCancelPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderReadRecordShape})
}

// ChangePurchaseOrder sends POST /api/{api_version}/purchase_orders/{purchase_order_id}/change.
func (p *PurchaseOrders) ChangePurchaseOrder(ctx context.Context, purchaseOrderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseOrderChangePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderReadRecordShape})
}

// GetPurchaseOrderInformation sends GET /api/{api_version}/purchase_orders/{purchase_order_id}/information.
func (p *PurchaseOrders) GetPurchaseOrderInformation(ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrderInformationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderInformationRecordShape})
}

// GetPurchaseOrderLines sends GET /api/{api_version}/purchase_orders/{purchase_order_id}/lines.
func (p *PurchaseOrders) GetPurchaseOrderLines(ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrderLinesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PurchaseOrderLineListShape})
}

// GetPurchaseOrderNotes sends GET /api/{api_version}/purchase_orders/{purchase_order_id}/notes.
func (p *PurchaseOrders) GetPurchaseOrderNotes(ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrderNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PurchaseOrderNotesRecordShape})
}

// CreatePurchaseOrderNote sends POST /api/{api_version}/purchase_orders/{purchase_order_id}/notes.
// The note travels as the form field note, and any files go beside it as multipart parts.
func (p *PurchaseOrders) CreatePurchaseOrderNote(ctx context.Context, purchaseOrderID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseOrderNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderNoteRecordShape})
}

// CreatePurchaseOrderPurchaseInvoice sends POST /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_invoices.
func (p *PurchaseOrders) CreatePurchaseOrderPurchaseInvoice(ctx context.Context, purchaseOrderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseOrderPurchaseInvoicesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderPurchaseInvoiceCreateRecordShape})
}

// GetPurchaseOrderPurchaseInvoices sends GET /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_invoices.
func (p *PurchaseOrders) GetPurchaseOrderPurchaseInvoices(ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrderPurchaseInvoicesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderPurchaseInvoiceListShape})
}

// GetPurchaseOrderPurchasePayments sends GET /api/{api_version}/purchase_orders/{purchase_order_id}/purchase_payments.
func (p *PurchaseOrders) GetPurchaseOrderPurchasePayments(ctx context.Context, purchaseOrderID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrderPurchasePaymentsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderPurchasePaymentListShape})
}

// ReactivatePurchaseOrder sends POST /api/{api_version}/purchase_orders/{purchase_order_id}/reactivate.
func (p *PurchaseOrders) ReactivatePurchaseOrder(ctx context.Context, purchaseOrderID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseOrderReactivatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_order_id": purchaseOrderID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderReadRecordShape})
}

// GetPurchaseOrderLine sends GET /api/{api_version}/purchase_orders/{purchase_order_id}/lines/{purchase_order_line_uuid}.
func (p *PurchaseOrders) GetPurchaseOrderLine(ctx context.Context, purchaseOrderID any, purchaseOrderLineUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(purchaseOrderLineUUID, "purchase_order_line_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrderLinePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_order_id": purchaseOrderID, "purchase_order_line_uuid": purchaseOrderLineUUID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderLineReadRecordShape})
}

// DeletePurchaseOrderNote sends DELETE /api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}; it has no typed view, so read ApiResponse.Map.
func (p *PurchaseOrders) DeletePurchaseOrderNote(ctx context.Context, purchaseOrderID any, purchaseOrderNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(purchaseOrderNoteID, "purchase_order_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PurchaseOrderNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_order_id": purchaseOrderID, "purchase_order_note_id": purchaseOrderNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetPurchaseOrderNote sends GET /api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}.
func (p *PurchaseOrders) GetPurchaseOrderNote(ctx context.Context, purchaseOrderID any, purchaseOrderNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(purchaseOrderNoteID, "purchase_order_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrderNotePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_order_id": purchaseOrderID, "purchase_order_note_id": purchaseOrderNoteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderNoteRecordShape})
}

// GetPurchaseOrderNoteFiles sends GET /api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}/files.
func (p *PurchaseOrders) GetPurchaseOrderNoteFiles(ctx context.Context, purchaseOrderID any, purchaseOrderNoteID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(purchaseOrderNoteID, "purchase_order_note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrderNoteFilesPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_order_id": purchaseOrderID, "purchase_order_note_id": purchaseOrderNoteID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderNoteFileListShape})
}

// GetPurchaseOrderNoteFile sends GET /api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}/files/{purchase_order_file_id}.
func (p *PurchaseOrders) GetPurchaseOrderNoteFile(ctx context.Context, purchaseOrderID any, purchaseOrderNoteID any, purchaseOrderFileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(purchaseOrderNoteID, "purchase_order_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(purchaseOrderFileID, "purchase_order_file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrderNoteFilePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"purchase_order_id": purchaseOrderID, "purchase_order_note_id": purchaseOrderNoteID, "purchase_order_file_id": purchaseOrderFileID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseOrderNoteFileReadRecordShape})
}

// GetPurchaseOrderKPIs sends GET /api/{api_version}/purchase_orders/{purchase_order_id}/kpis.
func (p *PurchaseOrders) GetPurchaseOrderKPIs(ctx context.Context, purchaseOrderID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseOrderKPIsPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_order_id": purchaseOrderID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PurchaseOrderKpisRecordShape})
}

// DeletePurchaseOrderNoteFile sends DELETE /api/{api_version}/purchase_orders/{purchase_order_id}/notes/{purchase_order_note_id}/files/{purchase_order_file_id}; it has no typed view, so read ApiResponse.Map.
func (p *PurchaseOrders) DeletePurchaseOrderNoteFile(ctx context.Context, purchaseOrderID any, purchaseOrderNoteID any, purchaseOrderFileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseOrderID, "purchase_order_id"); err != nil {
		return nil, err
	}
	if err := requireValue(purchaseOrderNoteID, "purchase_order_note_id"); err != nil {
		return nil, err
	}
	if err := requireValue(purchaseOrderFileID, "purchase_order_file_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PurchaseOrderNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_order_id": purchaseOrderID, "purchase_order_note_id": purchaseOrderNoteID, "purchase_order_file_id": purchaseOrderFileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "purchase_orders"

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

// New binds the purchase orders endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *PurchaseOrders { return &PurchaseOrders{requester: requester} }

// Requester returns the requester this resource was bound to.
func (p *PurchaseOrders) Requester() exhttp.Requester { return p.requester }
