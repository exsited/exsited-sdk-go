// Package tasks is the v4 tasks endpoints.
package tasks

// The path templates for the tasks endpoints; the client fills {api_version}.
const (
	TasksPath               = "/api/{api_version}/tasks"
	TaskPath                = "/api/{api_version}/tasks/{task_id}"
	TaskChecklistPath       = "/api/{api_version}/tasks/{task_id}/checklist"
	TaskNotesPath           = "/api/{api_version}/tasks/{task_code}/notes"
	TaskChecklistByUUIDPath = "/api/{api_version}/tasks/{task_id}/checklist/{task_checklist_uuid}"
	TaskNotePath            = "/api/{api_version}/tasks/{task_code}/notes/{note_id}"
	TaskNoteFilePath        = "/api/{api_version}/tasks/{task_code}/notes/{note_id}/files/{file_id}"
	TaskNoteFilesPath       = "/api/{api_version}/tasks/{task_code}/notes/{note_id}/files"
)

// TasksKey is the key a listing's records arrive under, the v4 list envelope's data.
const TasksKey = "data"

// WorkTaskKey is the envelope key one record arrives under.
const WorkTaskKey = "work_task"
