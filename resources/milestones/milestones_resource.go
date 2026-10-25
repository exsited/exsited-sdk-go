package milestones

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Milestones method of the same name,
// which is the primary surface.

// GetMilestones is the module form of Milestones.GetMilestones.
func GetMilestones(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetMilestones(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetMilestonesAccountByID is the module form of Milestones.GetMilestonesAccountByID.
func GetMilestonesAccountByID(requester exhttp.Requester, ctx context.Context, accountID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetMilestonesAccountByID(ctx, accountID, queryParams, raiseOnError, responseFormat...)
}

// GetMilestonesLabourByID is the module form of Milestones.GetMilestonesLabourByID.
func GetMilestonesLabourByID(requester exhttp.Requester, ctx context.Context, labourID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetMilestonesLabourByID(ctx, labourID, queryParams, raiseOnError, responseFormat...)
}

// GetMilestonesEmployeeByCode is the module form of Milestones.GetMilestonesEmployeeByCode.
func GetMilestonesEmployeeByCode(requester exhttp.Requester, ctx context.Context, employeeCode any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetMilestonesEmployeeByCode(ctx, employeeCode, queryParams, raiseOnError, responseFormat...)
}

// GetMilestone is the module form of Milestones.GetMilestone.
func GetMilestone(requester exhttp.Requester, ctx context.Context, milestoneUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetMilestone(ctx, milestoneUUID, raiseOnError, responseFormat...)
}
