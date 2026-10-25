// Package customcomponents is the v4 custom components endpoints.
package customcomponents

// The path templates for the custom components endpoints; the client fills {api_version}.
const (
	ComponentPath                           = "/api/{api_version}/component"
	ComponentByUUIDPath                     = "/api/{api_version}/component/{component_uuid}"
	ComponentByIDPath                       = "/api/{api_version}/component/{component_uuid}/{component_object_id}"
	ComponentObjectsPath                    = "/api/{api_version}/components/{component_uuid}/objects"
	ComponentObjectPath                     = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}"
	ComponentObjectActivityPath             = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/activity"
	ComponentObjectCustomObjectsPath        = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/custom_objects"
	ComponentObjectPDFPath                  = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/pdf"
	ComponentObjectActivityEntryPath        = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/activity/{activity_id}"
	ComponentObjectMilestonesPath           = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/milestones"
	ComponentObjectMilestonePath            = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/milestones/{milestone_id}"
	ComponentObjectHistoryPath              = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/history"
	ComponentObjectHistoryEntryPath         = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/history/{history_id}"
	ComponentObjectNotesPath                = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes"
	ComponentObjectNotePath                 = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}"
	ComponentObjectNoteFilesPath            = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}/files"
	ComponentObjectNoteFilePath             = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}/files/{file_id}"
	ComponentObjectCustomObjectPath         = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/custom_objects/{custom_object_id}"
	ComponentObjectCustomObjectInstancePath = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/custom_objects/{custom_object_id}/instances/{instance_id}"
	ComponentObjectLinkDomainsPath          = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/link_domains"
	ComponentObjectUnlinkDomainsPath        = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/unlink_domains"
	ComponentObjectChangeStatusPath         = "/api/{api_version}/components/{component_uuid}/objects/{component_object_id}/change_status"
)

// CustomComponentsKey is the key a listing's records arrive under, the v4 list envelope's data.
const CustomComponentsKey = "data"

// AccCusCompKey is the key one component instance arrives under: the component's name,
// which differs per component, so no view unwraps it.
const AccCusCompKey = "Acc Cus Comp"

// ComponentKey is the envelope key one record arrives under.
const ComponentKey = "component"

// CustomComponentObjectKey is the envelope key one record arrives under.
const CustomComponentObjectKey = "custom_component_object"

// CustomComponentObjectActivityKey is the envelope key one record arrives under.
const CustomComponentObjectActivityKey = "custom_component_object_activity"

// NoteKey is the envelope key one record arrives under.
const NoteKey = "note"

// FileKey is the envelope key one record arrives under.
const FileKey = "file"
