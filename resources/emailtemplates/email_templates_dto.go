package emailtemplates

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsEmailTemplateList is one page of SettingsEmailTemplateRecord records, under data.
type SettingsEmailTemplateList struct {
	Data       []SettingsEmailTemplateRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsEmailTemplateList was measured from.
func (SettingsEmailTemplateList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/email_templates",
	}
}

// SettingsEmailTemplateRecord is the settings email template body, under email_template.
type SettingsEmailTemplateRecord struct {
	CreatedBy        any
	CreatedOn        *string
	Description      *string
	DisplayName      *string
	Entity           any
	Href             *string
	ID               *string
	IsEdited         *bool
	IsSystemTemplate *bool
	LastUpdatedBy    any
	LastUpdatedOn    *string
	Name             *string
	Object           *string
	ReplyAddress     any
	RichEditor       *bool
	SenderAddress    any
	SenderName       any
	Status           *string
	Subject          *string
	TrackEmailOpen   *bool
	TrackLinkClick   *bool
}

// MeasuredFrom lists the endpoints SettingsEmailTemplateRecord was measured from.
func (SettingsEmailTemplateRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/email_templates",
		"GET /api/{api_version}/settings/email_templates/{email_template_id}",
		"POST /api/{api_version}/settings/email_templates",
		"PATCH /api/{api_version}/settings/email_templates/{email_template_id}",
	}
}

// EnvelopeKey unwraps {"email_template": {...}}.
func (SettingsEmailTemplateRecord) EnvelopeKey() string { return EmailTemplateKey }

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

// SettingsEmailTemplateListShape reads a body as SettingsEmailTemplateList.
func SettingsEmailTemplateListShape(decoded any) any {
	return dto.Read[SettingsEmailTemplateList](decoded)
}

// SettingsEmailTemplateRecordShape reads a body as SettingsEmailTemplateRecord.
func SettingsEmailTemplateRecordShape(decoded any) any {
	return dto.Read[SettingsEmailTemplateRecord](decoded)
}
