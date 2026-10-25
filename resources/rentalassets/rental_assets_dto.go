package rentalassets

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

// RentalAssetRecord is the rental asset body.
type RentalAssetRecord struct {
	ActivityLogs                  []any
	Code                          *string
	CreatedBy                     *CreatedBy
	CreatedOn                     *string
	CustomAttributes              []any
	CustomForm                    *CustomForm
	Description                   *string
	DisplayName                   *string
	Href                          *string
	ID                            *string
	ImageName                     any
	Item                          any
	LastUpdatedBy                 *LastUpdatedBy
	LastUpdatedOn                 any
	Name                          *string
	Object                        *string
	RentalAssetAvailability       []any
	RentalAssetCustomAvailability []RentalAssetCustomAvailability
	RentalAssetProfile            *RentalAssetProfile
	Status                        *string
	UseCustomAvailability         *bool
	Version                       *int64
}

// MeasuredFrom lists the endpoints RentalAssetRecord was measured from.
func (RentalAssetRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/rental_assets",
		"GET /api/{api_version}/rental_assets/{rental_asset_uuid}",
		"POST /api/{api_version}/rental_assets",
		"PATCH /api/{api_version}/rental_assets/{rental_asset_uuid}",
		"POST /api/{api_version}/rental_assets/{rental_asset_id}/cancel",
		"POST /api/{api_version}/rental_assets/{rental_asset_id}/reactivate",
	}
}

// EnvelopeKey unwraps {"rental_asset": {...}}.
func (RentalAssetRecord) EnvelopeKey() string { return RentalAssetKey }

// RentalAssetList is a page of RentalAssetRecord records.
type RentalAssetList struct {
	Data       []RentalAssetRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints RentalAssetList was measured from.
func (RentalAssetList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/rental_assets",
	}
}

// RentalAssetProfileRecord is the rental asset profile body.
type RentalAssetProfileRecord struct {
	CreatedBy                  *CreatedBy
	CreatedOn                  *string
	Currency                   *Currency
	DisplayName                *string
	Href                       *string
	ID                         *string
	IsTaxInclusive             *bool
	LastUpdatedBy              *LastUpdatedBy
	LastUpdatedOn              any
	MaxBookingLimit            any
	Name                       *string
	Object                     *string
	RentalAssetProfileSessions []RentalAssetProfileSession
	RentalAssets               []RentalAsset
	Status                     *string
	Tax                        any
	Version                    *int64
}

// MeasuredFrom lists the endpoints RentalAssetProfileRecord was measured from.
func (RentalAssetProfileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/rental_asset_profiles",
	}
}

// RentalAssetProfileList is a page of RentalAssetProfileRecord records.
type RentalAssetProfileList struct {
	Data       []RentalAssetProfileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints RentalAssetProfileList was measured from.
func (RentalAssetProfileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/rental_asset_profiles",
	}
}

// RentalAssetAvailabilityRecord is the rental asset availability body.
type RentalAssetAvailabilityRecord struct {
	AvailableList         []any
	Code                  *string
	DisplayName           *string
	ID                    *string
	Name                  *string
	Object                *string
	ProcessedScheduleList map[string]any
	RentalAssetProfile    *RentalAssetProfile
	ScheduleList          []ScheduleList
	SessionList           []SessionList
	Status                *string
	UnavailableList       []any
	UseCustomAvailability *bool
	Version               *int64
}

// MeasuredFrom lists the endpoints RentalAssetAvailabilityRecord was measured from.
func (RentalAssetAvailabilityRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/rental_asset_availability",
	}
}

// RentalAssetAvailabilityList is a page of RentalAssetAvailabilityRecord records.
type RentalAssetAvailabilityList struct {
	Data       []RentalAssetAvailabilityRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints RentalAssetAvailabilityList was measured from.
func (RentalAssetAvailabilityList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/rental_asset_availability",
	}
}

// RentalAssetProfileReadRecord is the body GetRentalAssetProfile answers.
type RentalAssetProfileReadRecord struct {
	ActivityLogs               []any
	CreatedBy                  *CreatedBy
	CreatedOn                  *string
	Currency                   *Currency
	DisplayName                *string
	Href                       *string
	ID                         *string
	IsTaxInclusive             *bool
	LastUpdatedBy              *LastUpdatedBy
	LastUpdatedOn              *string
	MaxBookingLimit            any
	Name                       *string
	Object                     *string
	RentalAssetProfileSessions []RentalAssetProfileSession
	RentalAssets               []RentalAsset
	Status                     *string
	Tax                        any
	Version                    *int64
}

// MeasuredFrom lists the endpoints RentalAssetProfileReadRecord was measured from.
func (RentalAssetProfileReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/rental_asset_profiles/{rental_asset_profile_id}",
		"POST /api/{api_version}/rental_asset_profiles",
	}
}

// EnvelopeKey unwraps {"rental_asset_profile": {...}}.
func (RentalAssetProfileReadRecord) EnvelopeKey() string { return RentalAssetProfileKey }

// CreatedBy is the created_by object.
type CreatedBy struct {
	Name   *string
	Object *string
}

// Currency is the currency object.
type Currency struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// RentalAsset is one entry of the rental_assets list.
type RentalAsset struct {
	DisplayName           *string
	Href                  *string
	ID                    *string
	ImageName             any
	Name                  *string
	Object                *string
	Status                *string
	UseCustomAvailability *bool
}

// RentalAssetProfileSession is one entry of the rental_asset_profile_sessions list.
type RentalAssetProfileSession struct {
	EndDate                               *string
	ID                                    *string
	Name                                  *string
	Object                                *string
	RentalAssetProfileSessionAvailability []RentalAssetProfileSessionAvailability
	RentalAssetProfileSessionPrices       []RentalAssetProfileSessionPrice
	RentalAssetProfileSessionWorkingHours []RentalAssetProfileSessionWorkingHour
	StartDate                             *string
	Status                                *string
}

// RentalAssetProfileSessionAvailability is one entry of the rental_asset_profile_session_availability list.
type RentalAssetProfileSessionAvailability struct {
	Available    *bool
	EndTime      *string
	ID           *string
	Object       *string
	PreferredDay *string
	StartTime    *string
	Status       *string
}

// RentalAssetProfileSessionPrice is one entry of the rental_asset_profile_session_prices list.
type RentalAssetProfileSessionPrice struct {
	ID                                       *string
	IsEnabledProRata                         *bool
	Object                                   *string
	Price                                    *float64
	PriceType                                *string
	PricingLevel                             any
	RentalAssetProfileSessionPriceHours      []any
	RentalAssetProfileSessionPricePeakHourly []any
	Status                                   *string
}

// RentalAssetProfileSessionWorkingHour is one entry of the rental_asset_profile_session_working_hours list.
type RentalAssetProfileSessionWorkingHour struct {
	EndTime   *string
	ID        *string
	Object    *string
	StartTime *string
	Status    *string
}

// ScheduleList is one entry of the schedule_list list.
type ScheduleList struct {
	Available    *string
	EndTime      *string
	PreferredDay *string
	StartTime    *string
	UUID         *string
}

// CustomForm is the custom_form object.
type CustomForm struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// RentalAssetCustomAvailability is one entry of the rental_asset_custom_availability list.
type RentalAssetCustomAvailability struct {
	Available    *bool
	EndTime      *string
	ID           *string
	Object       *string
	PreferredDay *string
	StartTime    *string
	Status       *string
}

// RentalAssetProfile is the rental_asset_profile object.
type RentalAssetProfile struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// SessionList is one entry of the session_list list.
type SessionList struct {
	EndDate               *string
	Name                  *string
	ProcessedScheduleList map[string]any
	ScheduleList          []ScheduleList
	StartDate             *string
	UUID                  *string
}

// RentalAssetAvailabilityListShape reads a body as RentalAssetAvailabilityList.
func RentalAssetAvailabilityListShape(decoded any) any {
	return dto.Read[RentalAssetAvailabilityList](decoded)
}

// RentalAssetProfileListShape reads a body as RentalAssetProfileList.
func RentalAssetProfileListShape(decoded any) any { return dto.Read[RentalAssetProfileList](decoded) }

// RentalAssetListShape reads a body as RentalAssetList.
func RentalAssetListShape(decoded any) any { return dto.Read[RentalAssetList](decoded) }

// RentalAssetRecordShape reads a body as RentalAssetRecord.
func RentalAssetRecordShape(decoded any) any { return dto.Read[RentalAssetRecord](decoded) }

// RentalAssetProfileReadRecordShape reads a body as RentalAssetProfileReadRecord.
func RentalAssetProfileReadRecordShape(decoded any) any {
	return dto.Read[RentalAssetProfileReadRecord](decoded)
}
