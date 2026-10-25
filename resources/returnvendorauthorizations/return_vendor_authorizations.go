package returnvendorauthorizations

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// ReturnVendorAuthorizations is the v4 return vendor authorizations endpoints, bound to one client.
type ReturnVendorAuthorizations struct {
	requester exhttp.Requester
}

// CreateReturnVendorAuthorization sends POST /api/{api_version}/return_vendor_authorizations.
func (r *ReturnVendorAuthorizations) CreateReturnVendorAuthorization(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, ReturnVendorAuthorizationsPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ReturnVendorAuthorizationViewShape})
}

// DeleteReturnVendorAuthorization sends DELETE /api/{api_version}/return_vendor_authorizations/{rva_id}; it has no typed view, so read ApiResponse.Map.
func (r *ReturnVendorAuthorizations) DeleteReturnVendorAuthorization(ctx context.Context, rvaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rvaID, "rva_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendDeleteRequest(ctx, ReturnVendorAuthorizationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rva_id": rvaID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateReturnVendorAuthorization sends PATCH /api/{api_version}/return_vendor_authorizations/{rva_id}.
func (r *ReturnVendorAuthorizations) UpdateReturnVendorAuthorization(ctx context.Context, rvaID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rvaID, "rva_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPatchRequest(ctx, ReturnVendorAuthorizationPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rva_id": rvaID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ReturnVendorAuthorizationViewShape})
}

// DeleteReturnVendorAuthorizationNote sends DELETE /api/{api_version}/return_vendor_authorizations/{rva_id}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (r *ReturnVendorAuthorizations) DeleteReturnVendorAuthorizationNote(ctx context.Context, rvaID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rvaID, "rva_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendDeleteRequest(ctx, ReturnVendorAuthorizationNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rva_id": rvaID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteReturnVendorAuthorizationNoteFile sends DELETE /api/{api_version}/return_vendor_authorizations/{rva_id}/notes/{note_id}/files/{file_id}; it has no typed view, so read ApiResponse.Map.
func (r *ReturnVendorAuthorizations) DeleteReturnVendorAuthorizationNoteFile(ctx context.Context, rvaID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rvaID, "rva_id"); err != nil {
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
	return r.requester.SendDeleteRequest(ctx, ReturnVendorAuthorizationNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rva_id": rvaID, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetReturnVendorAuthorizations sends GET /api/{api_version}/return_vendor_authorizations.
func (r *ReturnVendorAuthorizations) GetReturnVendorAuthorizations(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReturnVendorAuthorizationsPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ReturnVendorAuthorizationListShape})
}

// GetReturnVendorAuthorization sends GET /api/{api_version}/return_vendor_authorizations/{rva_id}.
func (r *ReturnVendorAuthorizations) GetReturnVendorAuthorization(ctx context.Context, rvaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rvaID, "rva_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, ReturnVendorAuthorizationPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rva_id": rvaID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ReturnVendorAuthorizationViewShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "return_vendor_authorizations"

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

// requirePayload refuses a nil payload, a typed nil included, before any I/O.
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

// New binds the return vendor authorizations endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *ReturnVendorAuthorizations {
	return &ReturnVendorAuthorizations{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (r *ReturnVendorAuthorizations) Requester() exhttp.Requester { return r.requester }
