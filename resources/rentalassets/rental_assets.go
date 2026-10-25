package rentalassets

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// RentalAssets is the v4 rental assets endpoints, bound to one client.
type RentalAssets struct {
	requester exhttp.Requester
}

// GetRentalAssetAvailability sends GET /api/{api_version}/rental_asset_availability.
func (r *RentalAssets) GetRentalAssetAvailability(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, RentalAssetAvailabilityPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: RentalAssetAvailabilityListShape})
}

// UpdateRentalAssetAvailability sends PATCH /api/{api_version}/rental_asset_availability; it has no typed view, so read ApiResponse.Map.
func (r *RentalAssets) UpdateRentalAssetAvailability(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPatchRequest(ctx, RentalAssetAvailabilityPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetRentalAssetProfiles sends GET /api/{api_version}/rental_asset_profiles.
func (r *RentalAssets) GetRentalAssetProfiles(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, RentalAssetProfilesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: RentalAssetProfileListShape})
}

// CreateRentalAssetProfile sends POST /api/{api_version}/rental_asset_profiles.
func (r *RentalAssets) CreateRentalAssetProfile(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, RentalAssetProfilesPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: RentalAssetProfileReadRecordShape})
}

// GetRentalAssets sends GET /api/{api_version}/rental_assets.
func (r *RentalAssets) GetRentalAssets(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, RentalAssetsPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: RentalAssetListShape})
}

// CreateRentalAsset sends POST /api/{api_version}/rental_assets.
func (r *RentalAssets) CreateRentalAsset(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, RentalAssetsPath, &exhttp.RequestOptions{
		JSONBody: payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: RentalAssetRecordShape})
}

// GetRentalAssetProfile sends GET /api/{api_version}/rental_asset_profiles/{rental_asset_profile_id}.
func (r *RentalAssets) GetRentalAssetProfile(ctx context.Context, rentalAssetProfileID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rentalAssetProfileID, "rental_asset_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, RentalAssetProfilePath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"rental_asset_profile_id": rentalAssetProfileID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: RentalAssetProfileReadRecordShape})
}

// UpdateRentalAssetProfile sends PATCH /api/{api_version}/rental_asset_profiles/{rental_asset_profile_id}; it has no typed view, so read ApiResponse.Map.
func (r *RentalAssets) UpdateRentalAssetProfile(ctx context.Context, rentalAssetProfileID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rentalAssetProfileID, "rental_asset_profile_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPatchRequest(ctx, RentalAssetProfilePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rental_asset_profile_id": rentalAssetProfileID},
		JSONBody:   payload, RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteRentalAssetProfile sends DELETE /api/{api_version}/rental_asset_profiles/{rental_asset_profile_id}; it has no typed view, so read ApiResponse.Map.
func (r *RentalAssets) DeleteRentalAssetProfile(ctx context.Context, rentalAssetProfileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rentalAssetProfileID, "rental_asset_profile_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendDeleteRequest(ctx, RentalAssetProfilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"rental_asset_profile_id": rentalAssetProfileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateRentalAsset sends PATCH /api/{api_version}/rental_assets/{rental_asset_uuid}.
func (r *RentalAssets) UpdateRentalAsset(ctx context.Context, rentalAssetUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rentalAssetUUID, "rental_asset_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPatchRequest(ctx, RentalAssetPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rental_asset_uuid": rentalAssetUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: RentalAssetRecordShape})
}

// DeleteRentalAsset sends DELETE /api/{api_version}/rental_assets/{rental_asset_uuid}; it has no typed view, so read ApiResponse.Map.
func (r *RentalAssets) DeleteRentalAsset(ctx context.Context, rentalAssetUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rentalAssetUUID, "rental_asset_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendDeleteRequest(ctx, RentalAssetPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rental_asset_uuid": rentalAssetUUID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetRentalAsset sends GET /api/{api_version}/rental_assets/{rental_asset_uuid}.
func (r *RentalAssets) GetRentalAsset(ctx context.Context, rentalAssetUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rentalAssetUUID, "rental_asset_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendGetRequest(ctx, RentalAssetPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rental_asset_uuid": rentalAssetUUID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: RentalAssetRecordShape})
}

// CancelRentalAsset sends POST /api/{api_version}/rental_assets/{rental_asset_id}/cancel.
func (r *RentalAssets) CancelRentalAsset(ctx context.Context, rentalAssetID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rentalAssetID, "rental_asset_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, RentalAssetCancelPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rental_asset_id": rentalAssetID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: RentalAssetRecordShape})
}

// ReactivateRentalAsset sends POST /api/{api_version}/rental_assets/{rental_asset_id}/reactivate.
func (r *RentalAssets) ReactivateRentalAsset(ctx context.Context, rentalAssetID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(rentalAssetID, "rental_asset_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return r.requester.SendPostRequest(ctx, RentalAssetReactivatePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"rental_asset_id": rentalAssetID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: RentalAssetRecordShape})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "rental_assets"

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

// New binds the rental assets endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *RentalAssets { return &RentalAssets{requester: requester} }

// Requester returns the requester this resource was bound to.
func (r *RentalAssets) Requester() exhttp.Requester { return r.requester }
