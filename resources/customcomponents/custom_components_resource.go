package customcomponents

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the CustomComponents method of the same name,
// which is the primary surface.

// GetComponent is the module form of CustomComponents.GetComponent.
func GetComponent(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponent(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetComponentByUUID is the module form of CustomComponents.GetComponentByUUID.
func GetComponentByUUID(requester exhttp.Requester, ctx context.Context, componentUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentByUUID(ctx, componentUUID, queryParams, raiseOnError, responseFormat...)
}

// CreateComponentByUUID is the module form of CustomComponents.CreateComponentByUUID.
func CreateComponentByUUID(requester exhttp.Requester, ctx context.Context, componentUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateComponentByUUID(ctx, componentUUID, payload, raiseOnError, responseFormat...)
}

// DeleteComponentByID is the module form of CustomComponents.DeleteComponentByID.
func DeleteComponentByID(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteComponentByID(ctx, componentUUID, componentObjectID, raiseOnError, responseFormat...)
}

// UpdateComponentByID is the module form of CustomComponents.UpdateComponentByID.
func UpdateComponentByID(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateComponentByID(ctx, componentUUID, componentObjectID, payload, raiseOnError, responseFormat...)
}

// GetComponentObjects is the module form of CustomComponents.GetComponentObjects.
func GetComponentObjects(requester exhttp.Requester, ctx context.Context, componentUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjects(ctx, componentUUID, queryParams, raiseOnError, responseFormat...)
}

// GetComponentObject is the module form of CustomComponents.GetComponentObject.
func GetComponentObject(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObject(ctx, componentUUID, componentObjectID, queryParams, raiseOnError, responseFormat...)
}

// GetComponentObjectActivity is the module form of CustomComponents.GetComponentObjectActivity.
func GetComponentObjectActivity(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectActivity(ctx, componentUUID, componentObjectID, queryParams, raiseOnError, responseFormat...)
}

// GetComponentObjectCustomObjects is the module form of CustomComponents.GetComponentObjectCustomObjects.
func GetComponentObjectCustomObjects(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectCustomObjects(ctx, componentUUID, componentObjectID, queryParams, raiseOnError, responseFormat...)
}

// GetComponentObjectPDF is the module form of CustomComponents.GetComponentObjectPDF.
func GetComponentObjectPDF(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectPDF(ctx, componentUUID, componentObjectID, queryParams, raiseOnError, responseFormat...)
}

// CreateComponent is the module form of CustomComponents.CreateComponent.
func CreateComponent(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateComponent(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateComponent is the module form of CustomComponents.UpdateComponent.
func UpdateComponent(requester exhttp.Requester, ctx context.Context, componentUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateComponent(ctx, componentUUID, payload, raiseOnError, responseFormat...)
}

// CreateComponentObject is the module form of CustomComponents.CreateComponentObject.
func CreateComponentObject(requester exhttp.Requester, ctx context.Context, componentUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateComponentObject(ctx, componentUUID, payload, raiseOnError, responseFormat...)
}

// UpdateComponentObject is the module form of CustomComponents.UpdateComponentObject.
func UpdateComponentObject(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateComponentObject(ctx, componentUUID, componentObjectID, payload, raiseOnError, responseFormat...)
}

// DeleteComponentObject is the module form of CustomComponents.DeleteComponentObject.
func DeleteComponentObject(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteComponentObject(ctx, componentUUID, componentObjectID, raiseOnError, responseFormat...)
}

// CreateComponentObjectActivity is the module form of CustomComponents.CreateComponentObjectActivity.
func CreateComponentObjectActivity(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateComponentObjectActivity(ctx, componentUUID, componentObjectID, payload, raiseOnError, responseFormat...)
}

// GetComponentObjectActivityByID is the module form of CustomComponents.GetComponentObjectActivityByID.
func GetComponentObjectActivityByID(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, activityID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectActivityByID(ctx, componentUUID, componentObjectID, activityID, raiseOnError, responseFormat...)
}

// GetComponentObjectMilestones is the module form of CustomComponents.GetComponentObjectMilestones.
func GetComponentObjectMilestones(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectMilestones(ctx, componentUUID, componentObjectID, queryParams, raiseOnError, responseFormat...)
}

// GetComponentObjectMilestone is the module form of CustomComponents.GetComponentObjectMilestone.
func GetComponentObjectMilestone(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, milestoneID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectMilestone(ctx, componentUUID, componentObjectID, milestoneID, raiseOnError, responseFormat...)
}

// GetComponentObjectHistory is the module form of CustomComponents.GetComponentObjectHistory.
func GetComponentObjectHistory(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectHistory(ctx, componentUUID, componentObjectID, queryParams, raiseOnError, responseFormat...)
}

// GetComponentObjectHistoryByID is the module form of CustomComponents.GetComponentObjectHistoryByID.
func GetComponentObjectHistoryByID(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, historyID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectHistoryByID(ctx, componentUUID, componentObjectID, historyID, raiseOnError, responseFormat...)
}

// GetComponentObjectNotes is the module form of CustomComponents.GetComponentObjectNotes.
func GetComponentObjectNotes(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectNotes(ctx, componentUUID, componentObjectID, queryParams, raiseOnError, responseFormat...)
}

// GetComponentObjectNote is the module form of CustomComponents.GetComponentObjectNote.
func GetComponentObjectNote(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectNote(ctx, componentUUID, componentObjectID, noteID, raiseOnError, responseFormat...)
}

// GetComponentObjectNoteFiles is the module form of CustomComponents.GetComponentObjectNoteFiles.
func GetComponentObjectNoteFiles(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectNoteFiles(ctx, componentUUID, componentObjectID, noteID, raiseOnError, responseFormat...)
}

// GetComponentObjectNoteFile is the module form of CustomComponents.GetComponentObjectNoteFile.
func GetComponentObjectNoteFile(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectNoteFile(ctx, componentUUID, componentObjectID, noteID, fileID, raiseOnError, responseFormat...)
}

// CreateComponentObjectNote is the module form of CustomComponents.CreateComponentObjectNote.
func CreateComponentObjectNote(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateComponentObjectNote(ctx, componentUUID, componentObjectID, note, files, raiseOnError, responseFormat...)
}

// UpdateComponentObjectNote is the module form of CustomComponents.UpdateComponentObjectNote.
func UpdateComponentObjectNote(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, noteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateComponentObjectNote(ctx, componentUUID, componentObjectID, noteID, note, files, raiseOnError, responseFormat...)
}

// DeleteComponentObjectNote is the module form of CustomComponents.DeleteComponentObjectNote.
func DeleteComponentObjectNote(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteComponentObjectNote(ctx, componentUUID, componentObjectID, noteID, raiseOnError, responseFormat...)
}

// DeleteComponentObjectNoteFile is the module form of CustomComponents.DeleteComponentObjectNoteFile.
func DeleteComponentObjectNoteFile(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteComponentObjectNoteFile(ctx, componentUUID, componentObjectID, noteID, fileID, raiseOnError, responseFormat...)
}

// GetComponentObjectCustomObject is the module form of CustomComponents.GetComponentObjectCustomObject.
func GetComponentObjectCustomObject(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectCustomObject(ctx, componentUUID, componentObjectID, customObjectID, raiseOnError, responseFormat...)
}

// CreateComponentObjectCustomObjectInstance is the module form of CustomComponents.CreateComponentObjectCustomObjectInstance.
func CreateComponentObjectCustomObjectInstance(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, customObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateComponentObjectCustomObjectInstance(ctx, componentUUID, componentObjectID, customObjectID, payload, raiseOnError, responseFormat...)
}

// GetComponentObjectCustomObjectInstance is the module form of CustomComponents.GetComponentObjectCustomObjectInstance.
func GetComponentObjectCustomObjectInstance(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetComponentObjectCustomObjectInstance(ctx, componentUUID, componentObjectID, customObjectID, instanceID, raiseOnError, responseFormat...)
}

// UpdateComponentObjectCustomObjectInstance is the module form of CustomComponents.UpdateComponentObjectCustomObjectInstance.
func UpdateComponentObjectCustomObjectInstance(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, customObjectID any, instanceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateComponentObjectCustomObjectInstance(ctx, componentUUID, componentObjectID, customObjectID, instanceID, payload, raiseOnError, responseFormat...)
}

// DeleteComponentObjectCustomObjectInstance is the module form of CustomComponents.DeleteComponentObjectCustomObjectInstance.
func DeleteComponentObjectCustomObjectInstance(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteComponentObjectCustomObjectInstance(ctx, componentUUID, componentObjectID, customObjectID, instanceID, raiseOnError, responseFormat...)
}

// LinkComponentObjectDomains is the module form of CustomComponents.LinkComponentObjectDomains.
func LinkComponentObjectDomains(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).LinkComponentObjectDomains(ctx, componentUUID, componentObjectID, payload, raiseOnError, responseFormat...)
}

// UpdateComponentObjectDomains is the module form of CustomComponents.UpdateComponentObjectDomains.
func UpdateComponentObjectDomains(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateComponentObjectDomains(ctx, componentUUID, componentObjectID, payload, raiseOnError, responseFormat...)
}

// UnlinkComponentObjectDomains is the module form of CustomComponents.UnlinkComponentObjectDomains.
func UnlinkComponentObjectDomains(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UnlinkComponentObjectDomains(ctx, componentUUID, componentObjectID, payload, raiseOnError, responseFormat...)
}

// ChangeComponentObjectStatus is the module form of CustomComponents.ChangeComponentObjectStatus.
func ChangeComponentObjectStatus(requester exhttp.Requester, ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).ChangeComponentObjectStatus(ctx, componentUUID, componentObjectID, payload, raiseOnError, responseFormat...)
}

// DeleteComponent is the module form of CustomComponents.DeleteComponent.
func DeleteComponent(requester exhttp.Requester, ctx context.Context, componentUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteComponent(ctx, componentUUID, raiseOnError, responseFormat...)
}
