package accountingcodes

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// AccountingCodes is the v4 accounting codes endpoints, bound to one client.
type AccountingCodes struct {
	requester exhttp.Requester
}

// GetSettingsAccountingCodes sends GET /api/{api_version}/settings/accounting_codes.
func (a *AccountingCodes) GetSettingsAccountingCodes(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, SettingsAccountingCodesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsAccountingCodeListShape})
}

// CreateSettingsAccountingCode sends POST /api/{api_version}/settings/accounting_codes.
func (a *AccountingCodes) CreateSettingsAccountingCode(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPostRequest(ctx, SettingsAccountingCodesPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsAccountingCodeRecordShape})
}

// GetSettingsAccountingCode sends GET /api/{api_version}/settings/accounting_codes/{accounting_code_id}.
func (a *AccountingCodes) GetSettingsAccountingCode(ctx context.Context, accountingCodeID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountingCodeID, "accounting_code_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, SettingsAccountingCodePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"accounting_code_id": accountingCodeID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsAccountingCodeRecordShape})
}

// DeleteSettingsAccountingCode sends DELETE /api/{api_version}/settings/accounting_codes/{accounting_code_id}; it has no typed view, so read ApiResponse.Map.
func (a *AccountingCodes) DeleteSettingsAccountingCode(ctx context.Context, accountingCodeID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountingCodeID, "accounting_code_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendDeleteRequest(ctx, SettingsAccountingCodePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"accounting_code_id": accountingCodeID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsAccountingCode sends PATCH /api/{api_version}/settings/accounting_codes/{accounting_code_id}.
func (a *AccountingCodes) UpdateSettingsAccountingCode(ctx context.Context, accountingCodeID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountingCodeID, "accounting_code_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendPatchRequest(ctx, SettingsAccountingCodePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"accounting_code_id": accountingCodeID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: SettingsAccountingCodeRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "accounting_codes"

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

// New binds the accounting codes endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *AccountingCodes { return &AccountingCodes{requester: requester} }

// Requester returns the requester this resource was bound to.
func (a *AccountingCodes) Requester() exhttp.Requester { return a.requester }
