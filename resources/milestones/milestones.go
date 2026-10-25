package milestones

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Milestones is the v4 milestones endpoints, bound to one client.
type Milestones struct {
	requester exhttp.Requester
}

// GetMilestones sends GET /api/{api_version}/milestones.
func (m *Milestones) GetMilestones(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return m.requester.SendGetRequest(ctx, MilestonesPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: MilestonesRecordShape})
}

// GetMilestonesAccountByID sends GET /api/{api_version}/milestones/account/{account_id}.
func (m *Milestones) GetMilestonesAccountByID(ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(accountID, "account_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return m.requester.SendGetRequest(ctx, MilestonesAccountByIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"account_id": accountID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: MilestonesRecordShape})
}

// GetMilestonesLabourByID sends GET /api/{api_version}/milestones/labour/{labour_id}.
func (m *Milestones) GetMilestonesLabourByID(ctx context.Context, labourID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(labourID, "labour_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return m.requester.SendGetRequest(ctx, MilestonesLabourByIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"labour_id": labourID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: MilestonesLabourByIDListShape})
}

// GetMilestonesEmployeeByCode sends GET /api/{api_version}/milestones/employee/{employee_code}.
func (m *Milestones) GetMilestonesEmployeeByCode(ctx context.Context, employeeCode any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(employeeCode, "employee_code"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return m.requester.SendGetRequest(ctx, MilestonesEmployeeByCodePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"employee_code": employeeCode}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: MilestonesRecordShape})
}

// GetMilestone sends GET /api/{api_version}/milestones/{milestone_uuid}; it has no typed view, so read ApiResponse.Map.
func (m *Milestones) GetMilestone(ctx context.Context, milestoneUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(milestoneUUID, "milestone_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return m.requester.SendGetRequest(ctx, MilestonePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"milestone_uuid": milestoneUUID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "milestones"

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

// New binds the milestones endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Milestones { return &Milestones{requester: requester} }

// Requester returns the requester this resource was bound to.
func (m *Milestones) Requester() exhttp.Requester { return m.requester }
