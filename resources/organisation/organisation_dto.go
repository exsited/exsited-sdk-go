package organisation

import "github.com/exsited/exsited-sdk-go/v4/dto"

// SettingsOrganisationRecord is the root organisation body, under organization.
type SettingsOrganisationRecord struct {
	AddressLine1       any `wire:"address_line1"`
	AddressLine2       any `wire:"address_line2"`
	AddressLine3       any `wire:"address_line3"`
	AddressLine4       any `wire:"address_line4"`
	AddressLine5       any `wire:"address_line5"`
	City               any
	CreatedBy          any
	CreatedOn          *string
	Description        any
	Details            any
	DisplayName        *string
	Email              any
	Fax                any
	FaxAreaCode        any
	FaxCountryCode     any
	Href               *string
	ID                 *string
	LastUpdatedBy      any
	LastUpdatedOn      *string
	Mobile             any
	MobileAreaCode     any
	MobileCountryCode  any
	Name               *string
	Nodes              []any
	Object             *string
	ParentOrganization *SettingsOrganisationParentOrganization
	Phone              any
	PhoneAreaCode      any
	PhoneCountryCode   any
	PostCode           any
	State              any
	Status             *string
	Type               *string
	Website            any
}

// MeasuredFrom lists the endpoints SettingsOrganisationRecord was measured from.
func (SettingsOrganisationRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/settings/organisation",
	}
}

// EnvelopeKey unwraps {"organization": {...}}.
func (SettingsOrganisationRecord) EnvelopeKey() string { return OrganizationKey }

// SettingsOrganisationParentOrganization is a SettingsOrganisationRecord's parent_organization.
type SettingsOrganisationParentOrganization struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// SettingsOrganisationRecordShape reads a body as SettingsOrganisationRecord.
func SettingsOrganisationRecordShape(decoded any) any {
	return dto.Read[SettingsOrganisationRecord](decoded)
}
