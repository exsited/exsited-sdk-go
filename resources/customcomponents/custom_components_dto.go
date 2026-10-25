package customcomponents

import "github.com/exsited/exsited-sdk-go/v4/dto"

// Pagination is the paging block, either shape.
type Pagination struct {
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	PrevCursor any
	Links      *PaginationLinks

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	NextPage     *string
	PreviousPage *string
}

// ComponentRecord is the component body.
type ComponentRecord struct {
	ChildCustomComponents []any
	CreatedBy             any
	CreatedOn             *string
	CustomComponent       any
	CustomForms           []CustomForm
	CustomObjects         []any
	Description           any
	DisplayName           *string
	Href                  *string
	ID                    *string
	LastUpdatedBy         any
	LastUpdatedOn         *string
	Name                  *string
	Object                *string
	Relation              *string
	Status                *string
	UseIn                 *string
	Version               *int64
}

// MeasuredFrom lists the endpoints ComponentRecord was measured from.
func (ComponentRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/component",
		"GET /api/{api_version}/component/{component_uuid}",
	}
}

// EnvelopeKey unwraps {"component": {...}}.
func (ComponentRecord) EnvelopeKey() string { return ComponentKey }

// ComponentList is a page of ComponentRecord records.
type ComponentList struct {
	Data       []ComponentRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ComponentList was measured from.
func (ComponentList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/component",
	}
}

// ComponentByUuidRecord is the component instance the create and update answer, under a
// key that is the component's own name, so ComponentByUuidRecordShape unwraps whatever
// single key the body has.
type ComponentByUuidRecord struct {
	ActivityLogs    []ActivityLog
	AttributeGroups []any
	Attributes      []Attribute
	CreatedBy       *string
	CreatedOn       *string
	CustomForm      *CustomForm
	CustomObjects   []any
	ID              *string
	LastUpdatedBy   *string
	LastUpdatedOn   *string
	Parents         []Parent
	Restrictions    []any
	Status          *string
	UUID            *string
	Version         *string
}

// MeasuredFrom lists the endpoints ComponentByUuidRecord was measured from.
func (ComponentByUuidRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/component/{component_uuid}",
		"PATCH /api/{api_version}/component/{component_uuid}/{component_object_id}",
	}
}

// ComponentObjectRecord is the component object body.
type ComponentObjectRecord struct {
	Accounts           []any
	AttributeGroups    []any
	Attributes         []any
	Code               *string
	Component          *Component
	CreatedBy          *string
	CreatedOn          *string
	CustomForm         *CustomForm
	CustomObjects      []any
	Href               *string
	ID                 *string
	Items              []any
	LastUpdatedBy      *string
	LastUpdatedOn      *string
	Object             *string
	Orders             []any
	Parents            []Parent
	Quotes             []any
	Restrictions       []any
	Status             *string
	UUID               *string
	Version            *int64
	WorkflowData       []any
	WorkflowStage      any
	WorkflowStatus     any
	WorkflowStatusUUID any
}

// MeasuredFrom lists the endpoints ComponentObjectRecord was measured from.
func (ComponentObjectRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}",
	}
}

// EnvelopeKey unwraps {"custom_component_object": {...}}.
func (ComponentObjectRecord) EnvelopeKey() string { return CustomComponentObjectKey }

// ComponentObjectsRecord is the component objects body.
type ComponentObjectsRecord struct {
	Data       []ComponentObject
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ComponentObjectsRecord was measured from.
func (ComponentObjectsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/components/{component_uuid}/objects",
	}
}

// ComponentObjectActivityRecord is the component object activity body.
type ComponentObjectActivityRecord struct {
	Activites  []any
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ComponentObjectActivityRecord was measured from.
func (ComponentObjectActivityRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/activity",
	}
}

// ComponentObjectCustomObjectsRecord is the component object custom objects body.
type ComponentObjectCustomObjectsRecord struct {
	CustomObjects []any
	Data          []any
	Object        *string
	Pagination    *Pagination
	URL           *string
}

// MeasuredFrom lists the endpoints ComponentObjectCustomObjectsRecord was measured from.
func (ComponentObjectCustomObjectsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/custom_objects",
	}
}

// Component is the component object.
type Component struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// ActivityLog is one entry of the activity_logs list.
type ActivityLog struct {
	Activity *string
}

// Attribute is one entry of the attributes list.
type Attribute struct {
	Name  *string
	Value *string
}

// ComponentObject is one row of ComponentObjectsRecord.
type ComponentObject struct {
	Accounts           []any
	AttributeGroups    []any
	Attributes         []any
	Code               *string
	Component          *Component
	CreatedBy          any
	CreatedOn          *string
	CustomForm         *CustomForm
	CustomObjects      []any
	Href               *string
	ID                 *string
	Items              []any
	LastUpdatedBy      any
	LastUpdatedOn      any
	Object             *string
	Orders             []any
	Parents            []Parent
	Quotes             []any
	Restrictions       []any
	Status             *string
	Version            *int64
	WorkflowData       []any
	WorkflowStage      any
	WorkflowStatus     any
	WorkflowStatusUUID any
}

// CustomForm is the custom_form or custom_forms object, alone or as a list entry.
type CustomForm struct {
	ColumnCount  *int64
	Data         *CustomFormRow
	ID           *string
	Name         *string
	Object       *string
	TemplateName *string
	UUID         *string
}

// CustomFormRow is the data object of CustomForm.
type CustomFormRow struct {
	FirstColumnAttribute  []FirstColumnAttribute  `wire:"firstColumnAttribute"`
	SecondColumnAttribute []SecondColumnAttribute `wire:"secondColumnAttribute"`
	ThirdColumnAttribute  []any                   `wire:"thirdColumnAttribute"`
}

// FirstColumnAttribute is one entry of the firstColumnAttribute list.
type FirstColumnAttribute struct {
	CanEncrypt                *bool `wire:"canEncrypt"`
	Description               *string
	DisplayName               *string `wire:"displayName"`
	Enabled                   *bool
	EncryptData               *bool `wire:"encryptData"`
	Entity                    any
	FieldRequired             *bool `wire:"fieldRequired"`
	FormMaxLengthInputed      *bool `wire:"formMaxLengthInputed"`
	FormMinLengthInputed      *bool `wire:"formMinLengthInputed"`
	ID                        *int64
	Ids                       *int64
	IsRequired                *bool   `wire:"isRequired"`
	IsUnique                  *bool   `wire:"isUnique"`
	MaxLength                 any     `wire:"maxLength"`
	MetaData                  *string `wire:"metaData"`
	MinLength                 any     `wire:"minLength"`
	Name                      *string
	Options                   []Option
	Position                  *string
	ShowInListView            *bool `wire:"showInListView"`
	TimeFormat                any   `wire:"timeFormat"`
	Type                      *string
	TypeValue                 *string `wire:"typeValue"`
	ValidationMessage         *string `wire:"validationMessage"`
	ValidationMessageRequired *bool   `wire:"validationMessageRequired"`
	Value                     *string
}

// Option is one entry of the options list.
type Option struct {
	ID   *string
	Name *string
}

// Parent is one entry of the parents list.
type Parent struct {
	Email  *string
	ID     *string
	Link   *string
	Name   *string
	Object *string
	Type   *string
}

// SecondColumnAttribute is one entry of the secondColumnAttribute list.
type SecondColumnAttribute struct {
	DisplayName               *string `wire:"displayName"`
	FieldEnabled              *bool   `wire:"fieldEnabled"`
	FieldRequired             *bool   `wire:"fieldRequired"`
	Ids                       *int64
	Name                      *string
	ShowInListView            *bool `wire:"showInListView"`
	Type                      *string
	TypeActual                *TypeActual `wire:"typeActual"`
	ValidationMessageRequired *bool       `wire:"validationMessageRequired"`
}

// TypeActual is the typeActual object.
type TypeActual struct {
	Key   *string
	Value *string
}

// ComponentListShape reads a body as ComponentList.
func ComponentListShape(decoded any) any { return dto.Read[ComponentList](decoded) }

// ComponentRecordShape reads a body as ComponentRecord.
func ComponentRecordShape(decoded any) any { return dto.Read[ComponentRecord](decoded) }

// ComponentByUuidRecordShape reads a body as ComponentByUuidRecord, unwrapping its one key.
func ComponentByUuidRecordShape(decoded any) any {
	if body, isObject := decoded.(map[string]any); isObject && len(body) == 1 {
		for _, inner := range body {
			if record, isRecord := inner.(map[string]any); isRecord {
				return dto.Read[ComponentByUuidRecord](record)
			}
		}
	}
	return dto.Read[ComponentByUuidRecord](decoded)
}

// ComponentObjectsRecordShape reads a body as ComponentObjectsRecord.
func ComponentObjectsRecordShape(decoded any) any { return dto.Read[ComponentObjectsRecord](decoded) }

// ComponentObjectRecordShape reads a body as ComponentObjectRecord.
func ComponentObjectRecordShape(decoded any) any { return dto.Read[ComponentObjectRecord](decoded) }

// ComponentObjectActivityRecordShape reads a body as ComponentObjectActivityRecord.
func ComponentObjectActivityRecordShape(decoded any) any {
	return dto.Read[ComponentObjectActivityRecord](decoded)
}

// ComponentObjectCustomObjectsRecordShape reads a body as ComponentObjectCustomObjectsRecord.
func ComponentObjectCustomObjectsRecordShape(decoded any) any {
	return dto.Read[ComponentObjectCustomObjectsRecord](decoded)
}

// ComponentObjectsDataDetail is a ComponentObjectsData's component.
type ComponentObjectsDataDetail struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// ComponentObjectsDataCustomForm is a ComponentObjectsData's custom_form.
type ComponentObjectsDataCustomForm struct {
	ID     *string
	Name   *string
	Object *string
}

// ComponentObjectsDataParent is one entry of a ComponentObjectsData's parents.
type ComponentObjectsDataParent struct {
	Email  *string
	ID     *string
	Name   *string
	Object *string
	Type   *string
}

// ComponentObjectRecordView is the component object body a v4 write answers, under
// custom_component_object: the fields a read nests there, with activity_logs beside them.
type ComponentObjectRecordView struct {
	Accounts           []any
	AttributeGroups    []any
	Attributes         []any
	Code               *string
	Component          *ComponentObjectCustomComponentObjectDetail
	CreatedBy          any
	CreatedOn          *string
	CustomForm         *ComponentObjectCustomComponentObjectCustomForm
	CustomObjects      []any
	Href               *string
	ID                 *string
	Items              []any
	LastUpdatedBy      any
	Object             *string
	Orders             []any
	Parents            []ComponentObjectCustomComponentObjectParent
	Quotes             []any
	Restrictions       []any
	Status             *string
	Version            *int64
	WorkflowData       []any
	WorkflowStage      any
	WorkflowStatus     any
	WorkflowStatusUUID any
	ActivityLogs       []ComponentObjectRecordActivityLog
	LastUpdatedOn      *string
}

// MeasuredFrom lists the endpoints ComponentObjectRecordView was measured from.
func (ComponentObjectRecordView) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/components/{component_uuid}/objects",
		"PATCH /api/{api_version}/components/{component_uuid}/objects/{component_object_id}",
	}
}

// EnvelopeKey unwraps {"custom_component_object": {...}}.
func (ComponentObjectRecordView) EnvelopeKey() string { return CustomComponentObjectKey }

// ComponentObjectCustomComponentObjectDetail is a ComponentObjectCustomComponentObject's component.
type ComponentObjectCustomComponentObjectDetail struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// ComponentObjectCustomComponentObjectCustomForm is a ComponentObjectCustomComponentObject's
// custom_form.
type ComponentObjectCustomComponentObjectCustomForm struct {
	ID     *string
	Name   *string
	Object *string
}

// ComponentObjectCustomComponentObjectParent is one entry of a
// ComponentObjectCustomComponentObject's parents.
type ComponentObjectCustomComponentObjectParent struct {
	Email  *string
	ID     *string
	Name   *string
	Object *string
	Type   *string
}

// ComponentObjectRecordActivityLog is one entry of a ComponentObjectRecordView's activity_logs.
type ComponentObjectRecordActivityLog struct {
	Activity *string
}

// ComponentObjectActivityEntryRecord is one component object activity, under
// custom_component_object_activity.
type ComponentObjectActivityEntryRecord struct {
	CreatedBy any
	CreatedOn *string
	Href      *string
	ID        *string
	Message   *string
	Object    *string
	Status    *string
	Type      *string
	UserName  *string
	Version   *int64
}

// MeasuredFrom lists the endpoints ComponentObjectActivityEntryRecord was measured from.
func (ComponentObjectActivityEntryRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/activity",
		"GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/activity/{activity_id}",
	}
}

// EnvelopeKey unwraps {"custom_component_object_activity": {...}}.
func (ComponentObjectActivityEntryRecord) EnvelopeKey() string {
	return CustomComponentObjectActivityKey
}

// ComponentObjectNotesRecord is one page of ComponentObjectNoteRecord records, under data.
type ComponentObjectNotesRecord struct {
	Data       []ComponentObjectNoteRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ComponentObjectNotesRecord was measured from.
func (ComponentObjectNotesRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes",
	}
}

// ComponentObjectNoteRecord is A component object's note, under note.
type ComponentObjectNoteRecord struct {
	Attachments           []ComponentObjectNoteAttachment
	Content               *string
	CreatedBy             *ComponentObjectNoteCreatedBy
	CreatedOn             *string
	CustomComponentObject *ComponentObjectNoteCustomComponentObject
	Group                 any
	Href                  *string
	ID                    *string
	LastUpdatedBy         *ComponentObjectNoteLastUpdatedBy
	LastUpdatedOn         *string
	Object                *string
	User                  any
	Version               *int64
	Watchers              []any
}

// MeasuredFrom lists the endpoints ComponentObjectNoteRecord was measured from.
func (ComponentObjectNoteRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes",
		"GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}",
		"POST /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes",
		"PATCH /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}",
	}
}

// EnvelopeKey unwraps {"note": {...}}.
func (ComponentObjectNoteRecord) EnvelopeKey() string { return NoteKey }

// ComponentObjectNoteAttachment is one entry of a ComponentObjectNoteRecord's attachments; uuid is
// the file id.
type ComponentObjectNoteAttachment struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// ComponentObjectNoteCreatedBy is a ComponentObjectNoteRecord's created_by.
type ComponentObjectNoteCreatedBy struct {
	Name   *string
	Object *string
}

// ComponentObjectNoteCustomComponentObject is a ComponentObjectNoteRecord's
// custom_component_object, the note's owner.
type ComponentObjectNoteCustomComponentObject struct {
	Href   *string
	ID     *string
	Object *string
}

// ComponentObjectNoteLastUpdatedBy is a ComponentObjectNoteRecord's last_updated_by, set once the
// note has been updated.
type ComponentObjectNoteLastUpdatedBy struct {
	Name   *string
	Object *string
}

// ComponentObjectNoteFilesRecord is a note's files, under data; its pagination carries has_more
// alone.
type ComponentObjectNoteFilesRecord struct {
	Data       []ComponentObjectNoteAttachment
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ComponentObjectNoteFilesRecord was measured from.
func (ComponentObjectNoteFilesRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}/files",
	}
}

// ComponentObjectNoteFileRecord is one file of a note, under file.
type ComponentObjectNoteFileRecord struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints ComponentObjectNoteFileRecord was measured from.
func (ComponentObjectNoteFileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/components/{component_uuid}/objects/{component_object_id}/notes/{note_id}/files/{file_id}",
	}
}

// EnvelopeKey unwraps {"file": {...}}.
func (ComponentObjectNoteFileRecord) EnvelopeKey() string { return FileKey }

// ComponentObjectRecordViewShape reads a body as ComponentObjectRecordView.
func ComponentObjectRecordViewShape(decoded any) any {
	return dto.Read[ComponentObjectRecordView](decoded)
}

// ComponentObjectActivityEntryRecordShape reads a body as ComponentObjectActivityEntryRecord.
func ComponentObjectActivityEntryRecordShape(decoded any) any {
	return dto.Read[ComponentObjectActivityEntryRecord](decoded)
}

// ComponentObjectNotesRecordShape reads a body as ComponentObjectNotesRecord.
func ComponentObjectNotesRecordShape(decoded any) any {
	return dto.Read[ComponentObjectNotesRecord](decoded)
}

// ComponentObjectNoteRecordShape reads a body as ComponentObjectNoteRecord.
func ComponentObjectNoteRecordShape(decoded any) any {
	return dto.Read[ComponentObjectNoteRecord](decoded)
}

// ComponentObjectNoteFilesRecordShape reads a body as ComponentObjectNoteFilesRecord.
func ComponentObjectNoteFilesRecordShape(decoded any) any {
	return dto.Read[ComponentObjectNoteFilesRecord](decoded)
}

// ComponentObjectNoteFileRecordShape reads a body as ComponentObjectNoteFileRecord.
func ComponentObjectNoteFileRecordShape(decoded any) any {
	return dto.Read[ComponentObjectNoteFileRecord](decoded)
}
