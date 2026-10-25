package usages

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

// UsageRecord is the usage body.
type UsageRecord struct {
	ChargeItem        *ChargeItem
	ChargingEndDate   *string
	ChargingPeriod    *string
	ChargingStartDate *string
	CreatedBy         *CreatedBy
	CreatedOn         *string
	CustomAttributes  []CustomAttribute
	EndTime           *string
	Href              *string
	ID                *string
	LastUpdatedBy     *LastUpdatedBy
	LastUpdatedOn     any
	Object            *string
	Quantity          *float64
	Source            any
	StartTime         *string
	Status            *string
	Type              *string
	Uom               any
	Version           *int64
}

// MeasuredFrom lists the endpoints UsageRecord was measured from.
func (UsageRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/usages",
		"GET /api/{api_version}/usages/{usage_id}",
	}
}

// EnvelopeKey unwraps {"usage": {...}}.
func (UsageRecord) EnvelopeKey() string { return UsageKey }

// UsageList is a page of UsageRecord records.
type UsageList struct {
	Data       []UsageRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints UsageList was measured from.
func (UsageList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/usages",
	}
}

// UsagesRecord is the usages body.
type UsagesRecord struct {
	ActivityLogs     []any
	ChargeItemName   *string
	ChargeItemUUID   *string
	ChargeStatus     *string
	ChargingPeriod   *string
	CreatedBy        *string
	CreatedOn        *string
	CustomAttributes []any
	EndTime          *string
	EventID          *string
	Failed           []any
	Href             *string
	ID               *string
	LastUpdatedBy    *string
	LastUpdatedOn    *string
	Object           *string
	Quantity         *string
	Source           *string
	StartTime        *string
	Success          []any
	Type             *string
	Uom              *string
	UUID             *string
	Version          *string
}

// MeasuredFrom lists the endpoints UsagesRecord was measured from.
func (UsagesRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/usages",
	}
}

// EnvelopeKey unwraps {"usage": {...}}.
func (UsagesRecord) EnvelopeKey() string { return UsageKey }

// UsageUpdatedRecord is the usage updated body.
type UsageUpdatedRecord struct {
	Href   *string
	ID     *string
	Note   *string
	Object *string
	Status *string
}

// MeasuredFrom lists the endpoints UsageUpdatedRecord was measured from.
func (UsageUpdatedRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/usages/{usage_id}",
		"PUT /api/{api_version}/usages/{usage_id}",
	}
}

// EnvelopeKey unwraps {"usage": {...}}.
func (UsageUpdatedRecord) EnvelopeKey() string { return UsageKey }

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// ChargeItem is the charge_item object.
type ChargeItem struct {
	ID     *string
	Name   *string
	Object *string
}

// CreatedBy is the created_by object.
type CreatedBy struct {
	Name   *string
	Object *string
}

// CustomAttribute is one entry of the custom_attributes list.
type CustomAttribute struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
	Type   *string
	Value  *string
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// UsageListShape reads a body as UsageList.
func UsageListShape(decoded any) any { return dto.Read[UsageList](decoded) }

// UsagesRecordShape reads a body as UsagesRecord.
func UsagesRecordShape(decoded any) any { return dto.Read[UsagesRecord](decoded) }

// UsageRecordShape reads a body as UsageRecord.
func UsageRecordShape(decoded any) any { return dto.Read[UsageRecord](decoded) }

// UsageUpdatedRecordShape reads a body as UsageUpdatedRecord.
func UsageUpdatedRecordShape(decoded any) any { return dto.Read[UsageUpdatedRecord](decoded) }
