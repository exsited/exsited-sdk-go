package returnvendorauthorizations

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ReturnVendorAuthorizations method of the same name,
// which is the primary surface.

// CreateReturnVendorAuthorization is the module form of ReturnVendorAuthorizations.CreateReturnVendorAuthorization.
func CreateReturnVendorAuthorization(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateReturnVendorAuthorization(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteReturnVendorAuthorization is the module form of ReturnVendorAuthorizations.DeleteReturnVendorAuthorization.
func DeleteReturnVendorAuthorization(requester exhttp.Requester, ctx context.Context, rvaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteReturnVendorAuthorization(ctx, rvaID, raiseOnError, responseFormat...)
}

// UpdateReturnVendorAuthorization is the module form of ReturnVendorAuthorizations.UpdateReturnVendorAuthorization.
func UpdateReturnVendorAuthorization(requester exhttp.Requester, ctx context.Context, rvaID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateReturnVendorAuthorization(ctx, rvaID, payload, raiseOnError, responseFormat...)
}

// DeleteReturnVendorAuthorizationNote is the module form of ReturnVendorAuthorizations.DeleteReturnVendorAuthorizationNote.
func DeleteReturnVendorAuthorizationNote(requester exhttp.Requester, ctx context.Context, rvaID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteReturnVendorAuthorizationNote(ctx, rvaID, noteID, raiseOnError, responseFormat...)
}

// DeleteReturnVendorAuthorizationNoteFile is the module form of ReturnVendorAuthorizations.DeleteReturnVendorAuthorizationNoteFile.
func DeleteReturnVendorAuthorizationNoteFile(requester exhttp.Requester, ctx context.Context, rvaID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteReturnVendorAuthorizationNoteFile(ctx, rvaID, noteID, fileID, raiseOnError, responseFormat...)
}

// GetReturnVendorAuthorizations is the module form of ReturnVendorAuthorizations.GetReturnVendorAuthorizations.
func GetReturnVendorAuthorizations(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReturnVendorAuthorizations(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetReturnVendorAuthorization is the module form of ReturnVendorAuthorizations.GetReturnVendorAuthorization.
func GetReturnVendorAuthorization(requester exhttp.Requester, ctx context.Context, rvaID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetReturnVendorAuthorization(ctx, rvaID, raiseOnError, responseFormat...)
}
