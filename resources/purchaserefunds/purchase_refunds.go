package purchaserefunds

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// PurchaseRefunds is the v4 purchase refunds endpoints, bound to one client.
type PurchaseRefunds struct {
	requester exhttp.Requester
}

// GetPurchaseRefunds sends GET /api/{api_version}/purchase_refunds.
func (p *PurchaseRefunds) GetPurchaseRefunds(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseRefundsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: PurchaseRefundsRecordShape})
}

// GetPurchaseRefund sends GET /api/{api_version}/purchase_refunds/{purchase_refund_id}.
func (p *PurchaseRefunds) GetPurchaseRefund(ctx context.Context, purchaseRefundID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseRefundID, "purchase_refund_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendGetRequest(ctx, PurchaseRefundPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_refund_id": purchaseRefundID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PurchaseRefundRecordShape})
}

// CreatePurchaseCreditNotePurchaseRefund sends POST /api/{api_version}/purchase_credit_notes/{purchase_credit_note_id}/purchase_refunds.
func (p *PurchaseRefunds) CreatePurchaseCreditNotePurchaseRefund(ctx context.Context, purchaseCreditNoteID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseCreditNoteID, "purchase_credit_note_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPostRequest(ctx, PurchaseCreditNotePurchaseRefundsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_credit_note_id": purchaseCreditNoteID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PurchaseRefundRecordShape})
}

// UpdatePurchaseRefund sends PATCH /api/{api_version}/purchase_refunds/{purchase_refund_id}.
func (p *PurchaseRefunds) UpdatePurchaseRefund(ctx context.Context, purchaseRefundID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseRefundID, "purchase_refund_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendPatchRequest(ctx, PurchaseRefundPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"purchase_refund_id": purchaseRefundID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: PurchaseRefundRecordShape})
}

// DeletePurchaseRefund sends DELETE /api/{api_version}/purchase_refunds/{purchase_refund_id}; it has no typed view, so read ApiResponse.Map.
func (p *PurchaseRefunds) DeletePurchaseRefund(ctx context.Context, purchaseRefundID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(purchaseRefundID, "purchase_refund_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return p.requester.SendDeleteRequest(ctx, PurchaseRefundPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"purchase_refund_id": purchaseRefundID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "purchase_refunds"

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

// requirePayload refuses a nil payload, a typed nil included, before any I/O.
func requirePayload(payload any) error {
	if payload == nil || isNilValue(payload) {
		return exerrors.NewUsageError("payload is required: this endpoint takes a body, and nil would " +
			"send the request with no body at all")
	}
	return nil
}

// New binds the purchase refunds endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *PurchaseRefunds { return &PurchaseRefunds{requester: requester} }

// Requester returns the requester this resource was bound to.
func (p *PurchaseRefunds) Requester() exhttp.Requester { return p.requester }
