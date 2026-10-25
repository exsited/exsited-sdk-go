package labours

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

// LabourRecord is the labour body.
type LabourRecord struct {
	ActivityLogs             []any
	Code                     *string
	CreatedBy                *CreatedBy
	CreatedOn                *string
	CustomForm               *CustomForm
	Description              *string
	DisplayName              *string
	EmailAddress             *string
	Href                     *string
	ID                       *string
	ImageName                any
	LabourAvailability       []any
	LabourContacts           []LabourContact
	LabourCustomAvailability []LabourCustomAvailability
	LabourProfile            *LabourProfile
	LastUpdatedBy            *LastUpdatedBy
	LastUpdatedOn            *string
	Name                     *string
	Object                   *string
	Status                   *string
	UseCustomAvailability    *bool
	Version                  *int64
}

// MeasuredFrom lists the endpoints LabourRecord was measured from.
func (LabourRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labours",
		"GET /api/{api_version}/labours/available_for_booking",
		"GET /api/{api_version}/labours/{labour_uuid}",
		"POST /api/{api_version}/labours",
		"PATCH /api/{api_version}/labours/{labour_uuid}",
		"POST /api/{api_version}/labours/{labour_id}/cancel",
		"POST /api/{api_version}/labours/{labour_id}/reactivate",
	}
}

// EnvelopeKey unwraps {"labour": {...}}.
func (LabourRecord) EnvelopeKey() string { return LabourKey }

// LabourList is a page of LabourRecord records.
type LabourList struct {
	Data       []LabourRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints LabourList was measured from.
func (LabourList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labours",
		"GET /api/{api_version}/labours/available_for_booking",
	}
}

// LabourProfileRecord is the labour profile body.
type LabourProfileRecord struct {
	CreatedBy       *CreatedBy
	CreatedOn       *string
	Currency        *Currency
	DisplayName     *string
	Href            *string
	ID              *string
	Labours         []Labour
	LastUpdatedBy   *LastUpdatedBy
	LastUpdatedOn   *string
	MaxBookingLimit any
	Name            *string
	Object          *string
	Sessions        []Session
	Status          *string
	Tax             *Tax
	TaxInclusive    *bool
	Version         *int64
}

// MeasuredFrom lists the endpoints LabourProfileRecord was measured from.
func (LabourProfileRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labour_profiles",
	}
}

// LabourProfileList is a page of LabourProfileRecord records.
type LabourProfileList struct {
	Data       []LabourProfileRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints LabourProfileList was measured from.
func (LabourProfileList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labour_profiles",
	}
}

// LabourAvailabilityRecord is the labour availability body.
type LabourAvailabilityRecord struct {
	AvailableList         []any
	Code                  *string
	DisplayName           *string
	EmailAddress          *string
	ID                    *string
	LabourProfile         *LabourProfile
	Name                  *string
	Object                *string
	ProcessedScheduleList map[string]any
	ScheduleList          []ScheduleList
	SessionList           []SessionList
	Status                *string
	UnavailableList       []any
	UseCustomAvailability *bool
	Version               *int64
}

// MeasuredFrom lists the endpoints LabourAvailabilityRecord was measured from.
func (LabourAvailabilityRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labour_availability",
	}
}

// LabourAvailabilityList is a page of LabourAvailabilityRecord records.
type LabourAvailabilityList struct {
	Data       []LabourAvailabilityRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints LabourAvailabilityList was measured from.
func (LabourAvailabilityList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labour_availability",
	}
}

// LabourOrderList is the page GetLabourOrders answers.
type LabourOrderList struct {
	Data       []LabourOrderRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints LabourOrderList was measured from.
func (LabourOrderList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labours/{labour_uuid}/orders",
	}
}

// LabourProfileReadRecord is the body GetLabourProfile answers.
type LabourProfileReadRecord struct {
	ActivityLogs    []any
	CreatedBy       *CreatedBy
	CreatedOn       *string
	Currency        *Currency
	DisplayName     *string
	Href            *string
	ID              *string
	Labours         []Labour
	LastUpdatedBy   *LastUpdatedBy
	LastUpdatedOn   *string
	MaxBookingLimit *int64
	Name            *string
	Object          *string
	Sessions        []Session
	Status          *string
	Tax             *Tax
	TaxInclusive    *bool
	Version         *int64
}

// MeasuredFrom lists the endpoints LabourProfileReadRecord was measured from.
func (LabourProfileReadRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labour_profiles/{labour_profile_id}",
		"POST /api/{api_version}/labour_profiles",
	}
}

// EnvelopeKey unwraps {"labour_profile": {...}}.
func (LabourProfileReadRecord) EnvelopeKey() string { return LabourProfileKey }

// LabourOrderRecord is one row of LabourOrderList.
type LabourOrderRecord struct {
	Account                  *Account
	AllowContract            *bool
	BillingAddress           any
	BillingPreferences       *BillingPreferences
	BillingStartDate         *string
	CommunicationPreferences []CommunicationPreference
	CreatedBy                any
	CreatedOn                *string
	Currency                 *Currency
	CustomForm               *CustomForm
	CustomerPurchaseOrderID  any
	Description              any
	DiscountProfile          any
	DisplayName              *string
	ID                       *string
	InvoiceNote              any
	LastUpdatedBy            any
	LastUpdatedOn            *string
	Manager                  any
	Name                     *string
	NextBillingFromDate      *string
	Object                   *string
	OrderStartDate           *string
	OrderType                *string
	Origin                   any
	OriginalTotal            any
	PreOrder                 *bool
	PriceTaxInclusive        *bool
	QuoteConverted           *bool
	QuoteOrder               *bool
	ReferralAccount          any
	SalesRep                 any
	ShippingAddress          any
	ShippingCost             *float64
	ShippingProfile          any
	Status                   *string
	TimeZone                 *TimeZone
	Version                  *int64
}

// MeasuredFrom lists the endpoints LabourOrderRecord was measured from.
func (LabourOrderRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labours/{labour_uuid}/orders",
	}
}

// Account is the account object.
type Account struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Availability is one entry of the availability list.
type Availability struct {
	Available    *bool
	EndTime      *string
	ID           *string
	Object       *string
	PreferredDay *string
	StartTime    *string
	Status       *string
}

// BillingPreferences is the billing_preferences object.
type BillingPreferences struct {
	BillingPeriod        *string
	CommunicationProfile *string
	ConsolidateInvoice   *bool
	FulfillmentMode      *string
	FulfillmentTerm      *string
	InvoiceMode          *string
	InvoiceTerm          *string
	PaymentMode          *string
	PaymentProcessor     *string
	PaymentTerm          *string
	PaymentTermAlignment *string
	ProformaProperties   *ProformaProperties
}

// CommunicationPreference is one entry of the communication_preferences list.
type CommunicationPreference struct {
	Channel *string
	Enabled *bool
}

// CreatedBy is the created_by object.
type CreatedBy struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Currency is the currency object.
type Currency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// CustomForm is the custom_form object.
type CustomForm struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Hour is one entry of the hours list.
type Hour struct {
	EndTime   *string
	ID        *string
	Object    *string
	Price     *float64
	StartTime *string
	Status    *string
}

// Labour is one entry of the labours list.
type Labour struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev any
}

// Price is one entry of the prices list.
type Price struct {
	EnabledProRata *bool
	Hours          []Hour
	ID             *string
	Object         *string
	PeakHourly     []any
	Price          *float64
	PriceType      *string
	PricingLevel   *PricingLevel
	Status         *string
}

// PricingLevel is the pricing_level object.
type PricingLevel struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// ProformaProperties is the proforma_properties object.
type ProformaProperties struct {
	AcceptanceDate     any
	AcceptanceRequired *bool
	AcceptanceTerm     any
	PaymentRequired    *bool
	TermsAndConditions any
}

// Session is one entry of the sessions list.
type Session struct {
	AllowOvertime *bool
	Availability  []Availability
	EndDate       *string
	Href          *string
	ID            *string
	Name          *string
	Object        *string
	Prices        []Price
	StartDate     *string
	Status        *string
	WorkingHours  []WorkingHour
}

// Tax is the tax object.
type Tax struct {
	Href   *string
	ID     *string
	Object *string
}

// TimeZone is the time_zone object.
type TimeZone struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// WorkingHour is one entry of the working_hours list.
type WorkingHour struct {
	EndTime   *string
	ID        *string
	Object    *string
	StartTime *string
	Status    *string
}

// Fax is the fax object.
type Fax struct {
	AreaCode    *string
	CountryCode *string
	DoNotCall   *bool
	Full        *string
	Number      *string
}

// LabourContact is one entry of the labour_contacts list.
type LabourContact struct {
	AddressLine1 *string
	AddressLine2 *string
	AddressLine3 *string
	AddressLine4 *string
	AddressLine5 *string
	City         *string
	ContactType  *string
	Country      *string
	Designation  *string
	EmailAddress *string
	Fax          *Fax
	FirstName    *string
	ID           *string
	LastName     *string
	MiddleName   any
	Mobile       *Mobile
	Object       *string
	Phone        *Phone
	PostCode     *string
	ReceiveEmail *bool
	Salutation   *string
	State        *string
	Status       *string
}

// LabourCustomAvailability is one entry of the labour_custom_availability list.
type LabourCustomAvailability struct {
	Available    *bool
	EndTime      *string
	ID           *string
	Object       *string
	PreferredDay *string
	StartTime    *string
}

// LabourProfile is the labour_profile object.
type LabourProfile struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Mobile is the mobile object.
type Mobile struct {
	AreaCode    *string
	CountryCode *string
	DoNotCall   *bool
	Full        *string
	Number      *string
}

// Phone is the phone object.
type Phone struct {
	AreaCode    *string
	CountryCode *string
	DoNotCall   *bool
	Full        *string
	Number      *string
}

// ScheduleList is one entry of the schedule_list list.
type ScheduleList struct {
	Available    any
	EndTime      *string
	ID           *string
	Object       *string
	PreferredDay *string
	StartTime    *string
	UUID         *string
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

// LabourAvailabilityListShape reads a body as LabourAvailabilityList.
func LabourAvailabilityListShape(decoded any) any { return dto.Read[LabourAvailabilityList](decoded) }

// LabourProfileListShape reads a body as LabourProfileList.
func LabourProfileListShape(decoded any) any { return dto.Read[LabourProfileList](decoded) }

// LabourListShape reads a body as LabourList.
func LabourListShape(decoded any) any { return dto.Read[LabourList](decoded) }

// LabourRecordShape reads a body as LabourRecord.
func LabourRecordShape(decoded any) any { return dto.Read[LabourRecord](decoded) }

// LabourOrderListShape reads a body as LabourOrderList.
func LabourOrderListShape(decoded any) any { return dto.Read[LabourOrderList](decoded) }

// LabourProfileReadRecordShape reads a body as LabourProfileReadRecord.
func LabourProfileReadRecordShape(decoded any) any { return dto.Read[LabourProfileReadRecord](decoded) }

// LabourProfileSessionList is one page of a labour profile's sessions, under data, each a
// LabourProfileViewSession.
type LabourProfileSessionList struct {
	Data       []LabourProfileViewSession
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints LabourProfileSessionList was measured from.
func (LabourProfileSessionList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labour_profiles/{labour_profile_id}/sessions",
	}
}

// LabourProfileViewSession is one entry of a LabourProfileView's sessions.
type LabourProfileViewSession struct {
	AllowOvertime *bool
	Availability  []LabourProfileViewSessionAvailability
	EndDate       *string
	Href          *string
	ID            *string
	Name          *string
	Object        *string
	Prices        []LabourProfileViewSessionPrice
	StartDate     *string
	Status        *string
	WorkingHours  []LabourProfileViewSessionWorkingHour
}

// LabourProfileViewSessionAvailability is one entry of a LabourProfileViewSession's availability.
type LabourProfileViewSessionAvailability struct {
	Available    *bool
	EndTime      *string
	ID           *string
	Object       *string
	PreferredDay *string
	StartTime    *string
	Status       *string
}

// LabourProfileViewSessionPrice is one entry of a LabourProfileViewSession's prices.
type LabourProfileViewSessionPrice struct {
	EnabledProRata *bool
	Hours          []LabourProfileViewSessionPriceHour
	ID             *string
	Object         *string
	PeakHourly     []any
	Price          *float64
	PriceType      *string
	PricingLevel   *LabourProfileViewSessionPricePricingLevel
	Status         *string
}

// LabourProfileViewSessionPriceHour is one entry of a LabourProfileViewSessionPrice's hours.
type LabourProfileViewSessionPriceHour struct {
	EndTime   *string
	ID        *string
	Object    *string
	Price     *float64
	StartTime *string
	Status    *string
}

// LabourProfileViewSessionPricePricingLevel is a LabourProfileViewSessionPrice's pricing_level.
type LabourProfileViewSessionPricePricingLevel struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// LabourProfileViewSessionWorkingHour is one entry of a LabourProfileViewSession's working_hours.
type LabourProfileViewSessionWorkingHour struct {
	EndTime   *string
	ID        *string
	Object    *string
	StartTime *string
	Status    *string
}

// LabourProfileSessionView is what GetLabourProfileSession answers, under labour_profile_session.
type LabourProfileSessionView struct {
	AllowOvertime *bool
	Availability  []LabourProfileViewSessionAvailability
	EndDate       *string
	Href          *string
	ID            *string
	Name          *string
	Object        *string
	Prices        []LabourProfileViewSessionPrice
	StartDate     *string
	Status        *string
	WorkingHours  []LabourProfileViewSessionWorkingHour
}

// MeasuredFrom lists the endpoints LabourProfileSessionView was measured from.
func (LabourProfileSessionView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labour_profiles/{labour_profile_id}/sessions/{session_id}",
	}
}

// EnvelopeKey unwraps {"labour_profile_session": {...}}.
func (LabourProfileSessionView) EnvelopeKey() string { return LabourProfileSessionKey }

// LabourCalendarRecord is what GetLabourCalendar answers, under labour_calendar: a
// LabourAvailability record for one labour.
type LabourCalendarRecord struct {
	AvailableList         []any
	Code                  *string
	DisplayName           *string
	EmailAddress          *string
	ID                    *string
	LabourProfile         *LabourAvailabilityLabourProfile
	Name                  *string
	Object                *string
	ProcessedScheduleList map[string]any
	ScheduleList          []LabourAvailabilityScheduleList
	SessionList           []LabourAvailabilitySessionList
	Status                *string
	UnavailableList       []any
	UseCustomAvailability *bool
	Version               *int64
}

// MeasuredFrom lists the endpoints LabourCalendarRecord was measured from.
func (LabourCalendarRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/labours/{labour_id}/availability",
	}
}

// EnvelopeKey unwraps {"labour_calendar": {...}}.
func (LabourCalendarRecord) EnvelopeKey() string { return LabourCalendarKey }

// LabourAvailabilityLabourProfile is a LabourAvailability's labour_profile.
type LabourAvailabilityLabourProfile struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// LabourAvailabilityScheduleList is one entry of a LabourAvailability's schedule_list.
type LabourAvailabilityScheduleList struct {
	Available    *bool
	EndTime      *string
	ID           *string
	Object       *string
	PreferredDay *string
	StartTime    *string
}

// LabourAvailabilitySessionList is one entry of a LabourAvailability's session_list.
type LabourAvailabilitySessionList struct {
	EndDate               *string
	Name                  *string
	ProcessedScheduleList map[string]any
	ScheduleList          []LabourAvailabilitySessionListScheduleList
	StartDate             *string
	UUID                  *string
}

// LabourAvailabilitySessionListScheduleList is one entry of a LabourAvailabilitySessionList's
// schedule_list.
type LabourAvailabilitySessionListScheduleList struct {
	Available    *string
	EndTime      *string
	PreferredDay *string
	StartTime    *string
	UUID         *string
}

// LabourProfileSessionListShape reads a body as LabourProfileSessionList.
func LabourProfileSessionListShape(decoded any) any {
	return dto.Read[LabourProfileSessionList](decoded)
}

// LabourProfileSessionViewShape reads a body as LabourProfileSessionView.
func LabourProfileSessionViewShape(decoded any) any {
	return dto.Read[LabourProfileSessionView](decoded)
}

// LabourCalendarRecordShape reads a body as LabourCalendarRecord.
func LabourCalendarRecordShape(decoded any) any { return dto.Read[LabourCalendarRecord](decoded) }
