package tasks

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// Tasks is the v4 tasks endpoints, bound to one client.
type Tasks struct {
	requester exhttp.Requester
}

// GetTasks sends GET /api/{api_version}/tasks.
func (t *Tasks) GetTasks(ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendGetRequest(ctx, TasksPath, &exhttp.RequestOptions{
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: TaskListShape})
}

// DeleteTask sends DELETE /api/{api_version}/tasks/{task_id}; it has no typed view, so read ApiResponse.Map.
func (t *Tasks) DeleteTask(ctx context.Context, taskID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendDeleteRequest(ctx, TaskPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"task_id": taskID}, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetTask sends GET /api/{api_version}/tasks/{task_id}.
func (t *Tasks) GetTask(ctx context.Context, taskID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendGetRequest(ctx, TaskPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"task_id": taskID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: TaskReadRecordShape})
}

// UpdateTask sends PATCH /api/{api_version}/tasks/{task_id}.
func (t *Tasks) UpdateTask(ctx context.Context, taskID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendPatchRequest(ctx, TaskPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"task_id": taskID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       TaskReadRecordShape})
}

// GetTaskChecklist sends GET /api/{api_version}/tasks/{task_id}/checklist.
func (t *Tasks) GetTaskChecklist(ctx context.Context, taskID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendGetRequest(ctx, TaskChecklistPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"task_id": taskID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: TaskChecklistReadRecordShape})
}

// CreateTaskChecklist sends POST /api/{api_version}/tasks/{task_id}/checklist.
func (t *Tasks) CreateTaskChecklist(ctx context.Context, taskID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_id"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendPostRequest(ctx, TaskChecklistPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"task_id": taskID}, JSONBody: payload, RaiseOnError: raiseOnError,
		ResponseFormat: format,
		DTOShape:       TaskReadRecordShape})
}

// GetTaskNotes sends GET /api/{api_version}/tasks/{task_code}/notes.
func (t *Tasks) GetTaskNotes(ctx context.Context, taskID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_code"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendGetRequest(ctx, TaskNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"task_code": taskID}, QueryParams: queryParams,
		RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: TaskNoteListShape})
}

// CreateTaskNote sends POST /api/{api_version}/tasks/{task_code}/notes; it has no typed view, so read ApiResponse.Map.
// The note travels as the form field note, and any files go beside it as multipart parts.
func (t *Tasks) CreateTaskNote(ctx context.Context, taskID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_code"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendPostRequest(ctx, TaskNotesPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"task_code": taskID}, FormData: map[string]any{"note": note}, Files: files, RaiseOnError: raiseOnError,
		ResponseFormat: format})
}

// GetTaskChecklistByUUID sends GET /api/{api_version}/tasks/{task_id}/checklist/{task_checklist_uuid}.
func (t *Tasks) GetTaskChecklistByUUID(ctx context.Context, taskID any, taskChecklistUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_id"); err != nil {
		return nil, err
	}
	if err := requireValue(taskChecklistUUID, "task_checklist_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendGetRequest(ctx, TaskChecklistByUUIDPath, &exhttp.RequestOptions{
		PathParams:  map[string]any{"task_id": taskID, "task_checklist_uuid": taskChecklistUUID},
		QueryParams: queryParams, RaiseOnError: raiseOnError, ResponseFormat: format,
		DTOShape: TaskChecklistByUUIDReadRecordShape})
}

// DeleteTaskNote sends DELETE /api/{api_version}/tasks/{task_code}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (t *Tasks) DeleteTaskNote(ctx context.Context, taskID any, taskNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_code"); err != nil {
		return nil, err
	}
	if err := requireValue(taskNoteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendDeleteRequest(ctx, TaskNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"task_code": taskID, "note_id": taskNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// CreateTask sends POST /api/{api_version}/tasks.
func (t *Tasks) CreateTask(ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendPostRequest(ctx, TasksPath, &exhttp.RequestOptions{
		JSONBody:     payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: TaskReadRecordShape})
}

// DeleteTaskChecklistByUUID sends DELETE /api/{api_version}/tasks/{task_id}/checklist/{task_checklist_uuid}; it has no typed view, so read ApiResponse.Map.
func (t *Tasks) DeleteTaskChecklistByUUID(ctx context.Context, taskID any, taskChecklistUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_id"); err != nil {
		return nil, err
	}
	if err := requireValue(taskChecklistUUID, "task_checklist_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendDeleteRequest(ctx, TaskChecklistByUUIDPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"task_id": taskID, "task_checklist_uuid": taskChecklistUUID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateTaskChecklistByUUID sends PATCH /api/{api_version}/tasks/{task_id}/checklist/{task_checklist_uuid}.
func (t *Tasks) UpdateTaskChecklistByUUID(ctx context.Context, taskID any, taskChecklistUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_id"); err != nil {
		return nil, err
	}
	if err := requireValue(taskChecklistUUID, "task_checklist_uuid"); err != nil {
		return nil, err
	}
	if err := requirePayload(payload); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendPatchRequest(ctx, TaskChecklistByUUIDPath, &exhttp.RequestOptions{
		PathParams: map[string]any{"task_id": taskID, "task_checklist_uuid": taskChecklistUUID}, JSONBody: payload,
		RaiseOnError: raiseOnError, ResponseFormat: format, DTOShape: TaskReadRecordShape})
}

// DeleteTaskNoteFile sends DELETE /api/{api_version}/tasks/{task_code}/notes/{note_id}/files/{file_id}; it has no typed view, so read ApiResponse.Map.
func (t *Tasks) DeleteTaskNoteFile(ctx context.Context, taskCode any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskCode, "task_code"); err != nil {
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
	return t.requester.SendDeleteRequest(ctx, TaskNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"task_code": taskCode, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetTaskNoteFiles sends GET /api/{api_version}/tasks/{task_code}/notes/{note_id}/files; it has no typed view, so read ApiResponse.Map.
func (t *Tasks) GetTaskNoteFiles(ctx context.Context, taskCode any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskCode, "task_code"); err != nil {
		return nil, err
	}
	if err := requireValue(noteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendGetRequest(ctx, TaskNoteFilesPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"task_code": taskCode, "note_id": noteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetTaskNoteFile sends GET /api/{api_version}/tasks/{task_code}/notes/{note_id}/files/{file_id}; it has no typed view, so read ApiResponse.Map.
func (t *Tasks) GetTaskNoteFile(ctx context.Context, taskCode any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskCode, "task_code"); err != nil {
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
	return t.requester.SendGetRequest(ctx, TaskNoteFilePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"task_code": taskCode, "note_id": noteID, "file_id": fileID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// GetTaskNote sends GET /api/{api_version}/tasks/{task_code}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (t *Tasks) GetTaskNote(ctx context.Context, taskID any, taskNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_code"); err != nil {
		return nil, err
	}
	if err := requireValue(taskNoteID, "note_id"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendGetRequest(ctx, TaskNotePath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"task_code": taskID, "note_id": taskNoteID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// UpdateTaskNote sends PATCH /api/{api_version}/tasks/{task_code}/notes/{note_id}; it has no typed view, so read ApiResponse.Map.
func (t *Tasks) UpdateTaskNote(ctx context.Context, taskID any, taskNoteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(taskID, "task_code"); err != nil {
		return nil, err
	}
	if err := requireValue(taskNoteID, "note_id"); err != nil {
		return nil, err
	}
	if err := requireNote(note); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return t.requester.SendPatchRequest(ctx, TaskNotePath, &exhttp.RequestOptions{
		PathParams: map[string]any{"task_code": taskID, "note_id": taskNoteID}, FormData: map[string]any{"note": note}, Files: files,
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "tasks"

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

// requireNote refuses an empty note before any I/O; the text is the whole body of a note write.
func requireNote(note string) error {
	if note == "" {
		return exerrors.NewUsageError("note is required: the note text is the body of this request")
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

// New binds the tasks endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *Tasks { return &Tasks{requester: requester} }

// Requester returns the requester this resource was bound to.
func (t *Tasks) Requester() exhttp.Requester { return t.requester }
