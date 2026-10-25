package giftcertificates

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// GiftCertificates is the v4 gift certificates endpoints, bound to one client.
type GiftCertificates struct {
	requester exhttp.Requester
}

// GetGiftCertificates sends GET /api/{api_version}/gift_certificates.
func (g *GiftCertificates) GetGiftCertificates(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendGetRequest(ctx, GiftCertificatesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: GiftCertificateListShape})
}

// CreateGiftCertificate sends POST /api/{api_version}/gift_certificates.
func (g *GiftCertificates) CreateGiftCertificate(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendPostRequest(ctx, GiftCertificatesPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: GiftCertificateRecordShape})
}

// DeleteGiftCertificate sends DELETE /api/{api_version}/gift_certificates/{gc_id}; it has no typed view, so read ApiResponse.Map.
func (g *GiftCertificates) DeleteGiftCertificate(ctx context.Context, gcID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(gcID, "gc_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendDeleteRequest(ctx, GiftCertificatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"gc_id": gcID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetGiftCertificate sends GET /api/{api_version}/gift_certificates/{gc_id}.
func (g *GiftCertificates) GetGiftCertificate(ctx context.Context, gcID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(gcID, "gc_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendGetRequest(ctx, GiftCertificatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"gc_id": gcID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: GiftCertificateRecordShape})
}

// UpdateGiftCertificate sends PATCH /api/{api_version}/gift_certificates/{gc_id}.
func (g *GiftCertificates) UpdateGiftCertificate(ctx context.Context, gcID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(gcID, "gc_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendPatchRequest(ctx, GiftCertificatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"gc_id": gcID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       GiftCertificateRecordShape})
}

// GetGiftCertificateAllocations sends GET /api/{api_version}/gift_certificates/{gc_id}/allocations.
func (g *GiftCertificates) GetGiftCertificateAllocations(ctx context.Context, gcID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(gcID, "gc_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendGetRequest(ctx, GiftCertificateAllocationsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"gc_id": gcID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: GiftCertificateAllocationsRecordShape})
}

// GetGiftCertificateTransactions sends GET /api/{api_version}/gift_certificates/{gc_id}/transactions.
func (g *GiftCertificates) GetGiftCertificateTransactions(ctx context.Context, gcID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(gcID, "gc_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendGetRequest(ctx, GiftCertificateTransactionsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"gc_id": gcID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: GiftCertificateTransactionListShape})
}

// AmendGiftCertificate sends POST /api/{api_version}/gift_certificates/{gc_id}/amend.
func (g *GiftCertificates) AmendGiftCertificate(ctx context.Context, gcID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(gcID, "gc_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendPostRequest(ctx, GiftCertificateAmendPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"gc_id": gcID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: GiftCertificateRecordShape})
}

// AllocateGiftCertificate sends POST /api/{api_version}/gift_certificates/{gc_id}/allocations.
func (g *GiftCertificates) AllocateGiftCertificate(ctx context.Context, gcID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(gcID, "gc_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendPostRequest(ctx, GiftCertificateAllocationsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"gc_id": gcID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: GiftCertificateRecordShape})
}

// DeallocateGiftCertificate sends POST /api/{api_version}/gift_certificates/{gc_id}/deallocations.
func (g *GiftCertificates) DeallocateGiftCertificate(ctx context.Context, gcID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(gcID, "gc_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendPostRequest(ctx, GiftCertificateDeallocationsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"gc_id": gcID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: GiftCertificateRecordShape})
}

// GetGiftCertificateAllocation sends GET /api/{api_version}/gift_certificates/{gc_id}/allocations/{allocation_id}.
func (g *GiftCertificates) GetGiftCertificateAllocation(ctx context.Context, gcID any, allocationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(gcID, "gc_id"); err != nil {
		return nil, err
	}
	if err := requireValue(allocationID, "allocation_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendGetRequest(ctx, GiftCertificateAllocationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"gc_id": gcID, "allocation_id": allocationID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: GiftCertificateAllocationReadRecordShape})
}

// GetGiftCertificateTransaction sends GET /api/{api_version}/gift_certificates/{gc_id}/transactions/{transaction_id}.
func (g *GiftCertificates) GetGiftCertificateTransaction(ctx context.Context, gcID any, transactionID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(gcID, "gc_id"); err != nil {
		return nil, err
	}
	if err := requireValue(transactionID, "transaction_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return g.requester.SendGetRequest(ctx, GiftCertificateTransactionPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"gc_id": gcID, "transaction_id": transactionID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: GiftCertificateTransactionReadRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "gift_certificates"

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

// New binds the gift certificates endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *GiftCertificates {
	return &GiftCertificates{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (g *GiftCertificates) Requester() exhttp.Requester { return g.requester }
