package accountingcodes

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

// SettingsAccountingCodeRecord is the settings accounting code body.
type SettingsAccountingCodeRecord struct {
	AccountCodeType *string
	CreatedBy       any
	CreatedOn       *string
	DefaultMatchKey *string
	Description     any
	DisplayName     *string
	GlAccountName   *string
	GlAccountNumber *string
	Href            *string
	ID              *string
	LastUpdatedBy   any
	LastUpdatedOn   *string
	Name            *string
	Object          *string
	Status          *string
	Version         *int64
}

// MeasuredFrom lists the endpoints SettingsAccountingCodeRecord was measured from.
func (SettingsAccountingCodeRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/accounting_codes",
		"GET /api/{api_version}/settings/accounting_codes/{accounting_code_id}",
		"POST /api/{api_version}/settings/accounting_codes",
		"PATCH /api/{api_version}/settings/accounting_codes/{accounting_code_id}",
	}
}

// EnvelopeKey unwraps {"accounting_code": {...}}.
func (SettingsAccountingCodeRecord) EnvelopeKey() string { return AccountingCodeKey }

// SettingsAccountingCodeList is a page of SettingsAccountingCodeRecord records.
type SettingsAccountingCodeList struct {
	Data       []SettingsAccountingCodeRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsAccountingCodeList was measured from.
func (SettingsAccountingCodeList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/accounting_codes",
	}
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// SettingsAccountingCodeListShape reads a body as SettingsAccountingCodeList.
func SettingsAccountingCodeListShape(decoded any) any {
	return dto.Read[SettingsAccountingCodeList](decoded)
}

// SettingsAccountingCodeRecordShape reads a body as SettingsAccountingCodeRecord.
func SettingsAccountingCodeRecordShape(decoded any) any {
	return dto.Read[SettingsAccountingCodeRecord](decoded)
}
