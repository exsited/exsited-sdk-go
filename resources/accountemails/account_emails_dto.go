package accountemails

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

// AccountCommunicationsEmailRecord is the account communications email body.
type AccountCommunicationsEmailRecord struct {
	AdditionalTemplate any
	Attachments        []Attachment
	Bcc                []any
	Body               any
	Cc                 []any
	CreatedBy          any
	CreatedOn          *string
	From               any
	ID                 *string
	LastUpdatedBy      any
	LastUpdatedOn      any
	MailType           *string
	MessageType        *string
	Object             *string
	Status             *string
	Subject            any
	To                 *string
	UUID               *string
}

// MeasuredFrom lists the endpoints AccountCommunicationsEmailRecord was measured from.
func (AccountCommunicationsEmailRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/communications/emails",
		"GET /api/{api_version}/accounts/{account_id}/communications/emails/{email_uuid}",
		"GET /api/{api_version}/accounts/{account_id}/communications/emails/{email_id}",
	}
}

// EnvelopeKey unwraps {"email": {...}}.
func (AccountCommunicationsEmailRecord) EnvelopeKey() string { return EmailKey }

// AccountCommunicationsEmailList is a page of AccountCommunicationsEmailRecord records.
type AccountCommunicationsEmailList struct {
	Data       []AccountCommunicationsEmailRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints AccountCommunicationsEmailList was measured from.
func (AccountCommunicationsEmailList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/accounts/{account_id}/communications/emails",
	}
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// Attachment is one entry of the attachments list.
type Attachment struct {
	FilePath *string
	Name     *string
	Path     *string
	Size     *int64
	Type     *string
}

// AccountCommunicationsEmailListShape reads a body as AccountCommunicationsEmailList.
func AccountCommunicationsEmailListShape(decoded any) any {
	return dto.Read[AccountCommunicationsEmailList](decoded)
}

// AccountCommunicationsEmailRecordShape reads a body as AccountCommunicationsEmailRecord.
func AccountCommunicationsEmailRecordShape(decoded any) any {
	return dto.Read[AccountCommunicationsEmailRecord](decoded)
}
