package accountgroups

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsAccountGroupList is one page of SettingsAccountGroupRecord records, under data.
type SettingsAccountGroupList struct {
	Data       []SettingsAccountGroupRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsAccountGroupList was measured from.
func (SettingsAccountGroupList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/account_groups",
	}
}

// SettingsAccountGroupRecord is the settings account group body, under account_group.
type SettingsAccountGroupRecord struct {
	AccountingCode      any
	CustomAttributes    []any
	Description         any
	DisplayName         *string
	Href                *string
	ID                  *string
	InvoiceAccountMedia []string
	ManagedBy           any
	Name                *string
	Object              *string
	PricingLevel        any
	StateRule           *SettingsAccountGroupStateRule
	Status              *string
}

// MeasuredFrom lists the endpoints SettingsAccountGroupRecord was measured from.
func (SettingsAccountGroupRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/account_groups",
		"GET /api/{api_version}/settings/account_groups/{account_group_id}",
		"POST /api/{api_version}/settings/account_groups",
	}
}

// EnvelopeKey unwraps {"account_group": {...}}.
func (SettingsAccountGroupRecord) EnvelopeKey() string { return AccountGroupKey }

// SettingsAccountGroupStateRule is a SettingsAccountGroupRecord's state_rule.
type SettingsAccountGroupStateRule struct {
	Href   *string
	ID     *string
	Object *string
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

// SettingsAccountGroupListShape reads a body as SettingsAccountGroupList.
func SettingsAccountGroupListShape(decoded any) any {
	return dto.Read[SettingsAccountGroupList](decoded)
}

// SettingsAccountGroupRecordShape reads a body as SettingsAccountGroupRecord.
func SettingsAccountGroupRecordShape(decoded any) any {
	return dto.Read[SettingsAccountGroupRecord](decoded)
}
