package customcomponents

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// CustomComponents is the v4 custom components endpoints, bound to one client.
type CustomComponents struct {
	requester exhttp.Requester
}

// GetComponent sends GET /api/{api_version}/component.
func (c *CustomComponents) GetComponent(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ComponentListShape})
}

// GetComponentByUUID sends GET /api/{api_version}/component/{component_uuid}.
func (c *CustomComponents) GetComponentByUUID(ctx context.Context, componentUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentByUUIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentRecordShape})
}

// CreateComponentByUUID sends POST /api/{api_version}/component/{component_uuid}; it answers 201 under the
// component's own name.
func (c *CustomComponents) CreateComponentByUUID(ctx context.Context, componentUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, ComponentByUUIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentByUuidRecordShape})
}

// DeleteComponentByID sends DELETE /api/{api_version}/component/{component_uuid}/{component_object_id};
// it answers 204. It has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) DeleteComponentByID(ctx context.Context, componentUUID any, componentObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendDeleteRequest(ctx, ComponentByIDPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateComponentByID sends PATCH /api/{api_version}/component/{component_uuid}/{component_object_id};
// it answers 202 under the component's own name.
func (c *CustomComponents) UpdateComponentByID(ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPatchRequest(ctx, ComponentByIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID},
		JSONBody:   payload, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ComponentByUuidRecordShape})
}

// GetComponentObjects sends GET /api/{api_version}/components/{component_uuid}/objects.
func (c *CustomComponents) GetComponentObjects(ctx context.Context, componentUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectsRecordShape})
}

// GetComponentObject sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}.
func (c *CustomComponents) GetComponentObject(ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ComponentObjectRecordShape})
}

// GetComponentObjectActivity sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/activity.
func (c *CustomComponents) GetComponentObjectActivity(ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectActivityPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ComponentObjectActivityRecordShape})
}

// GetComponentObjectCustomObjects sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/custom_objects.
func (c *CustomComponents) GetComponentObjectCustomObjects(ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectCustomObjectsPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: ComponentObjectCustomObjectsRecordShape})
}

// GetComponentObjectPDF sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/pdf; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) GetComponentObjectPDF(ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectPDFPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateComponent sends POST /api/{api_version}/component.
func (c *CustomComponents) CreateComponent(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, ComponentPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentRecordShape})
}

// UpdateComponent sends PATCH /api/{api_version}/component/{component_uuid}.
func (c *CustomComponents) UpdateComponent(ctx context.Context, componentUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPatchRequest(ctx, ComponentByUUIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentRecordShape})
}

// CreateComponentObject sends POST /api/{api_version}/components/{component_uuid}/objects.
func (c *CustomComponents) CreateComponentObject(ctx context.Context, componentUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, ComponentObjectsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectRecordViewShape})
}

// UpdateComponentObject sends PATCH /api/{api_version}/components/{component_uuid}/objects/{component_object_id}.
func (c *CustomComponents) UpdateComponentObject(ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPatchRequest(ctx, ComponentObjectPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectRecordViewShape})
}

// DeleteComponentObject sends DELETE /api/{api_version}/components/{component_uuid}/objects/{component_object_id}; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) DeleteComponentObject(ctx context.Context, componentUUID any, componentObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendDeleteRequest(ctx, ComponentObjectPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateComponentObjectActivity sends POST /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/activity.
func (c *CustomComponents) CreateComponentObjectActivity(ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, ComponentObjectActivityPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectActivityEntryRecordShape})
}

// GetComponentObjectActivityByID sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/activity/{activity_id}.
func (c *CustomComponents) GetComponentObjectActivityByID(ctx context.Context, componentUUID any, componentObjectID any, activityID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(activityID, "activity_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectActivityEntryPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "activity_id": activityID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectActivityEntryRecordShape})
}

// GetComponentObjectMilestones sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/milestones; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) GetComponentObjectMilestones(ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectMilestonesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetComponentObjectMilestone sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/milestones/{milestone_id}; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) GetComponentObjectMilestone(ctx context.Context, componentUUID any, componentObjectID any, milestoneID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(milestoneID, "milestone_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectMilestonePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "milestone_id": milestoneID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetComponentObjectHistory sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/history; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) GetComponentObjectHistory(ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectHistoryPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetComponentObjectHistoryByID sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/history/{history_id}; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) GetComponentObjectHistoryByID(ctx context.Context, componentUUID any, componentObjectID any, historyID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(historyID, "history_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectHistoryEntryPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "history_id": historyID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetComponentObjectNotes sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes.
func (c *CustomComponents) GetComponentObjectNotes(ctx context.Context, componentUUID any, componentObjectID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectNotesRecordShape})
}

// GetComponentObjectNote sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}.
func (c *CustomComponents) GetComponentObjectNote(ctx context.Context, componentUUID any, componentObjectID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectNoteRecordShape})
}

// GetComponentObjectNoteFiles sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}/files.
func (c *CustomComponents) GetComponentObjectNoteFiles(ctx context.Context, componentUUID any, componentObjectID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectNoteFilesPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectNoteFilesRecordShape})
}

// GetComponentObjectNoteFile sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}/files/{file_id}.
func (c *CustomComponents) GetComponentObjectNoteFile(ctx context.Context, componentUUID any, componentObjectID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
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
	return c.requester.SendGetRequest(ctx, ComponentObjectNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectNoteFileRecordShape})
}

// CreateComponentObjectNote sends POST /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes.
func (c *CustomComponents) CreateComponentObjectNote(ctx context.Context, componentUUID any, componentObjectID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, ComponentObjectNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectNoteRecordShape})
}

// UpdateComponentObjectNote sends PATCH /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}.
func (c *CustomComponents) UpdateComponentObjectNote(ctx context.Context, componentUUID any, componentObjectID any, noteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPatchRequest(ctx, ComponentObjectNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "note_id": noteID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectNoteRecordShape})
}

// DeleteComponentObjectNote sends DELETE /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) DeleteComponentObjectNote(ctx context.Context, componentUUID any, componentObjectID any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendDeleteRequest(ctx, ComponentObjectNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteComponentObjectNoteFile sends DELETE /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}/files/{file_id}; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) DeleteComponentObjectNoteFile(ctx context.Context, componentUUID any, componentObjectID any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
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
	return c.requester.SendDeleteRequest(ctx, ComponentObjectNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetComponentObjectCustomObject sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/custom_objects/{custom_object_id}; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) GetComponentObjectCustomObject(ctx context.Context, componentUUID any, componentObjectID any, customObjectID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectCustomObjectPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "custom_object_id": customObjectID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateComponentObjectCustomObjectInstance sends POST /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/custom_objects/{custom_object_id}; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) CreateComponentObjectCustomObjectInstance(ctx context.Context, componentUUID any, componentObjectID any, customObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, ComponentObjectCustomObjectPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "custom_object_id": customObjectID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetComponentObjectCustomObjectInstance sends GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/custom_objects/{custom_object_id}/instances/{instance_id}; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) GetComponentObjectCustomObjectInstance(ctx context.Context, componentUUID any, componentObjectID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(instanceID, "instance_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendGetRequest(ctx, ComponentObjectCustomObjectInstancePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "custom_object_id": customObjectID, "instance_id": instanceID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateComponentObjectCustomObjectInstance sends PATCH /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/custom_objects/{custom_object_id}/instances/{instance_id}; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) UpdateComponentObjectCustomObjectInstance(ctx context.Context, componentUUID any, componentObjectID any, customObjectID any, instanceID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(instanceID, "instance_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPatchRequest(ctx, ComponentObjectCustomObjectInstancePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "custom_object_id": customObjectID, "instance_id": instanceID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// DeleteComponentObjectCustomObjectInstance sends DELETE /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/custom_objects/{custom_object_id}/instances/{instance_id}; it has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) DeleteComponentObjectCustomObjectInstance(ctx context.Context, componentUUID any, componentObjectID any, customObjectID any, instanceID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(customObjectID, "custom_object_id"); err != nil {
		return nil, err
	}
	if err := requireValue(instanceID, "instance_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendDeleteRequest(ctx, ComponentObjectCustomObjectInstancePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID, "custom_object_id": customObjectID, "instance_id": instanceID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// LinkComponentObjectDomains sends POST /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/link_domains.
func (c *CustomComponents) LinkComponentObjectDomains(ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, ComponentObjectLinkDomainsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectRecordViewShape})
}

// UpdateComponentObjectDomains sends PATCH /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/link_domains.
func (c *CustomComponents) UpdateComponentObjectDomains(ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPatchRequest(ctx, ComponentObjectLinkDomainsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectRecordViewShape})
}

// UnlinkComponentObjectDomains sends POST /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/unlink_domains.
func (c *CustomComponents) UnlinkComponentObjectDomains(ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, ComponentObjectUnlinkDomainsPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectRecordViewShape})
}

// ChangeComponentObjectStatus sends POST /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/change_status.
func (c *CustomComponents) ChangeComponentObjectStatus(ctx context.Context, componentUUID any, componentObjectID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	if err := requireValue(componentObjectID, "component_object_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendPostRequest(ctx, ComponentObjectChangeStatusPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"component_uuid": componentUUID, "component_object_id": componentObjectID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: ComponentObjectRecordViewShape})
}

// DeleteComponent sends DELETE /api/{api_version}/component/{component_uuid}, removing the component itself;
// it answers 204. It has no typed view, so read ApiResponse.Map.
func (c *CustomComponents) DeleteComponent(ctx context.Context, componentUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(componentUUID, "component_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return c.requester.SendDeleteRequest(ctx, ComponentByUUIDPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"component_uuid": componentUUID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "custom_components"

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

// requireNote refuses an empty note before any I/O; the text is the whole body of a note write.
func requireNote(note string) error {
	if note == "" {
		return exerrors.NewUsageError("note is required: the note text is the body of this request")
	}
	return nil
}

// New binds the custom components endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *CustomComponents {
	return &CustomComponents{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (c *CustomComponents) Requester() exhttp.Requester { return c.requester }
