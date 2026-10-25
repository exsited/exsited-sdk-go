package tasks

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the Tasks method of the same name,
// which is the primary surface.

// GetTasks is the module form of Tasks.GetTasks.
func GetTasks(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetTasks(ctx, queryParams, raiseOnError, responseFormat...)
}

// DeleteTask is the module form of Tasks.DeleteTask.
func DeleteTask(requester exhttp.Requester, ctx context.Context, taskID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteTask(ctx, taskID, raiseOnError, responseFormat...)
}

// GetTask is the module form of Tasks.GetTask.
func GetTask(requester exhttp.Requester, ctx context.Context, taskID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetTask(ctx, taskID, queryParams, raiseOnError, responseFormat...)
}

// UpdateTask is the module form of Tasks.UpdateTask.
func UpdateTask(requester exhttp.Requester, ctx context.Context, taskID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateTask(ctx, taskID, payload, raiseOnError, responseFormat...)
}

// GetTaskChecklist is the module form of Tasks.GetTaskChecklist.
func GetTaskChecklist(requester exhttp.Requester, ctx context.Context, taskID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetTaskChecklist(ctx, taskID, queryParams, raiseOnError, responseFormat...)
}

// CreateTaskChecklist is the module form of Tasks.CreateTaskChecklist.
func CreateTaskChecklist(requester exhttp.Requester, ctx context.Context, taskID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateTaskChecklist(ctx, taskID, payload, raiseOnError, responseFormat...)
}

// GetTaskNotes is the module form of Tasks.GetTaskNotes.
func GetTaskNotes(requester exhttp.Requester, ctx context.Context, taskID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetTaskNotes(ctx, taskID, queryParams, raiseOnError, responseFormat...)
}

// CreateTaskNote is the module form of Tasks.CreateTaskNote.
func CreateTaskNote(requester exhttp.Requester, ctx context.Context, taskID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateTaskNote(ctx, taskID, note, files, raiseOnError, responseFormat...)
}

// GetTaskChecklistByUUID is the module form of Tasks.GetTaskChecklistByUUID.
func GetTaskChecklistByUUID(requester exhttp.Requester, ctx context.Context, taskID any, taskChecklistUUID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetTaskChecklistByUUID(ctx, taskID, taskChecklistUUID, queryParams, raiseOnError, responseFormat...)
}

// DeleteTaskNote is the module form of Tasks.DeleteTaskNote.
func DeleteTaskNote(requester exhttp.Requester, ctx context.Context, taskID any, taskNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteTaskNote(ctx, taskID, taskNoteID, raiseOnError, responseFormat...)
}

// CreateTask is the module form of Tasks.CreateTask.
func CreateTask(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateTask(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteTaskChecklistByUUID is the module form of Tasks.DeleteTaskChecklistByUUID.
func DeleteTaskChecklistByUUID(requester exhttp.Requester, ctx context.Context, taskID any, taskChecklistUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteTaskChecklistByUUID(ctx, taskID, taskChecklistUUID, raiseOnError, responseFormat...)
}

// UpdateTaskChecklistByUUID is the module form of Tasks.UpdateTaskChecklistByUUID.
func UpdateTaskChecklistByUUID(requester exhttp.Requester, ctx context.Context, taskID any, taskChecklistUUID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateTaskChecklistByUUID(ctx, taskID, taskChecklistUUID, payload, raiseOnError, responseFormat...)
}

// DeleteTaskNoteFile is the module form of Tasks.DeleteTaskNoteFile.
func DeleteTaskNoteFile(requester exhttp.Requester, ctx context.Context, taskCode any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteTaskNoteFile(ctx, taskCode, noteID, fileID, raiseOnError, responseFormat...)
}

// GetTaskNoteFiles is the module form of Tasks.GetTaskNoteFiles.
func GetTaskNoteFiles(requester exhttp.Requester, ctx context.Context, taskCode any, noteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetTaskNoteFiles(ctx, taskCode, noteID, raiseOnError, responseFormat...)
}

// GetTaskNoteFile is the module form of Tasks.GetTaskNoteFile.
func GetTaskNoteFile(requester exhttp.Requester, ctx context.Context, taskCode any, noteID any, fileID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetTaskNoteFile(ctx, taskCode, noteID, fileID, raiseOnError, responseFormat...)
}

// GetTaskNote is the module form of Tasks.GetTaskNote.
func GetTaskNote(requester exhttp.Requester, ctx context.Context, taskID any, taskNoteID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetTaskNote(ctx, taskID, taskNoteID, raiseOnError, responseFormat...)
}

// UpdateTaskNote is the module form of Tasks.UpdateTaskNote.
func UpdateTaskNote(requester exhttp.Requester, ctx context.Context, taskID any, taskNoteID any, note string, files map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateTaskNote(ctx, taskID, taskNoteID, note, files, raiseOnError, responseFormat...)
}
