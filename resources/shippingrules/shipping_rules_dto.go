package shippingrules

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsShippingRuleList is one page of SettingsShippingRuleRecord records, under data.
type SettingsShippingRuleList struct {
	Data       []SettingsShippingRuleRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsShippingRuleList was measured from.
func (SettingsShippingRuleList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/shipping_rules",
	}
}

// SettingsShippingRuleRecord is the settings shipping rule body, under shipping_rule.
type SettingsShippingRuleRecord struct {
	Configuration *SettingsShippingRuleConfiguration
	CreatedBy     any
	CreatedOn     *string
	Description   *string
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Name          *string
	Object        *string
	ShippingClass any
	Status        *string
	Version       *int64
}

// MeasuredFrom lists the endpoints SettingsShippingRuleRecord was measured from.
func (SettingsShippingRuleRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/shipping_rules",
		"GET /api/{api_version}/settings/shipping_rules/{shipping_rule_id}",
		"POST /api/{api_version}/settings/shipping_rules",
		"PATCH /api/{api_version}/settings/shipping_rules/{shipping_rule_id}",
	}
}

// EnvelopeKey unwraps {"shipping_rule": {...}}.
func (SettingsShippingRuleRecord) EnvelopeKey() string { return ShippingRuleKey }

// SettingsShippingRuleConfiguration is a SettingsShippingRuleRecord's configuration.
type SettingsShippingRuleConfiguration struct {
	Rate  any
	Zones []any
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

// SettingsShippingRuleListShape reads a body as SettingsShippingRuleList.
func SettingsShippingRuleListShape(decoded any) any {
	return dto.Read[SettingsShippingRuleList](decoded)
}

// SettingsShippingRuleRecordShape reads a body as SettingsShippingRuleRecord.
func SettingsShippingRuleRecordShape(decoded any) any {
	return dto.Read[SettingsShippingRuleRecord](decoded)
}
