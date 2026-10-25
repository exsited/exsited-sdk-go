// Package notegroups is the v4 note groups endpoints.
package notegroups

// The path templates for the note groups endpoints; the client fills {api_version}.
const (
	NoteGroupsPath           = "/api/{api_version}/note_groups"
	NoteGroupPath            = "/api/{api_version}/note_groups/{note_group_id}"
	NoteGroupNotesPath       = "/api/{api_version}/note_groups/{note_group_id}/notes"
	NoteGroupNotesByUserPath = "/api/{api_version}/note_groups/{note_group_id}/notes/user/{note_user}"
)

// NoteGroupsKey is the key a listing's records arrive under, the v4 list envelope's data.
const NoteGroupsKey = "data"

// NoteGroupKey is the envelope key one record arrives under.
const NoteGroupKey = "note_group"
