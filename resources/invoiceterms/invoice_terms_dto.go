package invoiceterms

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsInvoiceTermList is one page of SettingsInvoiceTermRecord records, under data.
type SettingsInvoiceTermList struct {
	Data       []SettingsInvoiceTermRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsInvoiceTermList was measured from.
func (SettingsInvoiceTermList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/invoice_terms",
	}
}

// SettingsInvoiceTermRecord is the settings invoice term body, under invoice_term.
type SettingsInvoiceTermRecord struct {
	Description *string
	DisplayName *string
	Duration    *SettingsInvoiceTermDuration
	Href        *string
	ID          *string
	Name        *string
	Object      *string
	Status      *string
}

// MeasuredFrom lists the endpoints SettingsInvoiceTermRecord was measured from.
func (SettingsInvoiceTermRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/invoice_terms",
		"GET /api/{api_version}/settings/invoice_terms/{invoice_term_id}",
	}
}

// EnvelopeKey unwraps {"invoice_term": {...}}.
func (SettingsInvoiceTermRecord) EnvelopeKey() string { return InvoiceTermKey }

// SettingsInvoiceTermDuration is a SettingsInvoiceTermRecord's duration.
type SettingsInvoiceTermDuration struct {
	Count  *int64
	Period *string
}

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

// SettingsInvoiceTermWriteRecord is the settings invoice term body a create or an update answers,
// under invoice_term.
type SettingsInvoiceTermWriteRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	DisplayName   *string
	Duration      *int64
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	Ordering      *int64
	Period        *string
	Status        *string
}

// MeasuredFrom lists the endpoints SettingsInvoiceTermWriteRecord was measured from.
func (SettingsInvoiceTermWriteRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/settings/invoice_terms",
		"PATCH /api/{api_version}/settings/invoice_terms/{invoice_term_id}",
	}
}

// EnvelopeKey unwraps {"invoice_term": {...}}.
func (SettingsInvoiceTermWriteRecord) EnvelopeKey() string { return InvoiceTermKey }

// SettingsInvoiceTermListShape reads a body as SettingsInvoiceTermList.
func SettingsInvoiceTermListShape(decoded any) any { return dto.Read[SettingsInvoiceTermList](decoded) }

// SettingsInvoiceTermRecordShape reads a body as SettingsInvoiceTermRecord.
func SettingsInvoiceTermRecordShape(decoded any) any {
	return dto.Read[SettingsInvoiceTermRecord](decoded)
}

// SettingsInvoiceTermWriteRecordShape reads a body as SettingsInvoiceTermWriteRecord.
func SettingsInvoiceTermWriteRecordShape(decoded any) any {
	return dto.Read[SettingsInvoiceTermWriteRecord](decoded)
}
