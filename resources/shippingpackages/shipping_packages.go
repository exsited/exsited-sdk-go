package shippingpackages

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// ShippingPackages is the v4 shipping packages endpoints, bound to one client.
type ShippingPackages struct {
	requester exhttp.Requester
}

// GetSettingsShippingPackages sends GET /api/{api_version}/settings/shipping_packages.
func (s *ShippingPackages) GetSettingsShippingPackages(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsShippingPackagesPath, &exhttp.RequestOptions{
		QueryParams:  queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsShippingPackageListShape})
}

// CreateSettingsShippingPackage sends POST /api/{api_version}/settings/shipping_packages.
func (s *ShippingPackages) CreateSettingsShippingPackage(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPostRequest(ctx, SettingsShippingPackagesPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsShippingPackageRecordShape})
}

// GetSettingsShippingPackage sends GET /api/{api_version}/settings/shipping_packages/{shipping_package_id}.
func (s *ShippingPackages) GetSettingsShippingPackage(ctx context.Context, shippingPackageID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(shippingPackageID, "shipping_package_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendGetRequest(ctx, SettingsShippingPackagePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"shipping_package_id": shippingPackageID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsShippingPackageRecordShape})
}

// DeleteSettingsShippingPackage sends DELETE /api/{api_version}/settings/shipping_packages/{shipping_package_id}; it has no typed view, so read ApiResponse.Map.
func (s *ShippingPackages) DeleteSettingsShippingPackage(ctx context.Context, shippingPackageID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(shippingPackageID, "shipping_package_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendDeleteRequest(ctx, SettingsShippingPackagePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"shipping_package_id": shippingPackageID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateSettingsShippingPackage sends PATCH /api/{api_version}/settings/shipping_packages/{shipping_package_id}.
func (s *ShippingPackages) UpdateSettingsShippingPackage(ctx context.Context, shippingPackageID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(shippingPackageID, "shipping_package_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return s.requester.SendPatchRequest(ctx, SettingsShippingPackagePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"shipping_package_id": shippingPackageID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: SettingsShippingPackageRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "shipping_packages"

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

// New binds the shipping packages endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *ShippingPackages {
	return &ShippingPackages{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (s *ShippingPackages) Requester() exhttp.Requester { return s.requester }
