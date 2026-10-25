package documenttemplates

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsDocumentTemplateList is one page of SettingsDocumentTemplateRecord records, under data.
type SettingsDocumentTemplateList struct {
	Data       []SettingsDocumentTemplateRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsDocumentTemplateList was measured from.
func (SettingsDocumentTemplateList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/document_templates",
	}
}

// SettingsDocumentTemplateRecord is the settings document template body, under document_template.
type SettingsDocumentTemplateRecord struct {
	CreatedBy        any
	CreatedOn        *string
	Description      *string
	DisplayName      *string
	Href             *string
	ID               *string
	IsCustomTemplate *bool
	IsDefault        *bool
	IsEdited         *bool
	IsSystemTemplate *bool
	LastUpdatedBy    any
	LastUpdatedOn    *string
	Name             *string
	Object           *string
	Status           *string
	Type             *string
	Version          *int64
}

// MeasuredFrom lists the endpoints SettingsDocumentTemplateRecord was measured from.
func (SettingsDocumentTemplateRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/document_templates",
		"GET /api/{api_version}/settings/document_templates/{document_template_id}",
		"POST /api/{api_version}/settings/document_templates",
		"PATCH /api/{api_version}/settings/document_templates/{document_template_id}",
	}
}

// EnvelopeKey unwraps {"document_template": {...}}.
func (SettingsDocumentTemplateRecord) EnvelopeKey() string { return DocumentTemplateKey }

// Pagination is a list envelope's paging block. One response fills only one shape: cursor (HasMore
// present), offset (Records present) or the external database listings' (HasNext present), and the
// others read nil.
type Pagination struct {
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	PrevCursor *string
	Links      *PaginationLinks

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	NextPage     *string
	PreviousPage *string

	// The external database listings' shape.
	HasNext *bool
	HasPrev *bool
	PerPage *int64
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// SettingsDocumentTemplateListShape reads a body as SettingsDocumentTemplateList.
func SettingsDocumentTemplateListShape(decoded any) any {
	return dto.Read[SettingsDocumentTemplateList](decoded)
}

// SettingsDocumentTemplateRecordShape reads a body as SettingsDocumentTemplateRecord.
func SettingsDocumentTemplateRecordShape(decoded any) any {
	return dto.Read[SettingsDocumentTemplateRecord](decoded)
}
