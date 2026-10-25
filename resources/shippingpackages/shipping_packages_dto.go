package shippingpackages

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsShippingPackageList is one page of SettingsShippingPackageRecord records, under data.
type SettingsShippingPackageList struct {
	Data       []SettingsShippingPackageRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints SettingsShippingPackageList was measured from.
func (SettingsShippingPackageList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/shipping_packages",
	}
}

// SettingsShippingPackageRecord is the settings shipping package body, under shipping_package.
type SettingsShippingPackageRecord struct {
	CreatedBy     any
	CreatedOn     *string
	Cubic         any
	CubicUom      any
	Description   *string
	DisplayName   *string
	Height        *float64
	HeightUom     any
	Href          *string
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn *string
	Length        *float64
	LengthUom     any
	Name          *string
	Object        *string
	Status        *string
	Weight        *float64
	WeightUom     any
	Width         *float64
	WidthUom      any
}

// MeasuredFrom lists the endpoints SettingsShippingPackageRecord was measured from.
func (SettingsShippingPackageRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/shipping_packages",
		"GET /api/{api_version}/settings/shipping_packages/{shipping_package_id}",
		"POST /api/{api_version}/settings/shipping_packages",
		"PATCH /api/{api_version}/settings/shipping_packages/{shipping_package_id}",
	}
}

// EnvelopeKey unwraps {"shipping_package": {...}}.
func (SettingsShippingPackageRecord) EnvelopeKey() string { return ShippingPackageKey }

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

// SettingsShippingPackageListShape reads a body as SettingsShippingPackageList.
func SettingsShippingPackageListShape(decoded any) any {
	return dto.Read[SettingsShippingPackageList](decoded)
}

// SettingsShippingPackageRecordShape reads a body as SettingsShippingPackageRecord.
func SettingsShippingPackageRecordShape(decoded any) any {
	return dto.Read[SettingsShippingPackageRecord](decoded)
}
