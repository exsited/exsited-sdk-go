package customfeatures

import "github.com/exsited/exsited-sdk-go/v4/dto"

// CustomFeatureRecord is the custom feature body.
type CustomFeatureRecord struct {
	ActivityLogs          []any
	ChildCustomComponents []any
	CreatedBy             any
	CreatedOn             *string
	CustomComponent       any
	CustomForms           []CustomForm
	CustomObjects         []any
	Description           *string
	DisplayName           *string
	Href                  *string
	ID                    *string
	LastUpdatedBy         any
	LastUpdatedOn         *string
	Message               *string
	Name                  *string
	Object                *string
	Relation              *string
	Status                *string
	UseIn                 *string
	Version               *int64
}

// MeasuredFrom lists the endpoints CustomFeatureRecord was measured from.
func (CustomFeatureRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/custom_feature",
		"PATCH /api/{api_version}/custom_feature/{custom_feature_id}",
		"POST /api/{api_version}/custom_feature/{custom_feature_id}",
	}
}

// EnvelopeKey unwraps {"custom_feature": {...}}.
func (CustomFeatureRecord) EnvelopeKey() string { return CustomFeatureKey }

// Attribute is one entry of the attributes list.
type Attribute struct {
	DataType       *string
	DisplayName    *string
	Name           *string
	ShowInListView *bool
	Type           *string
}

// Column1 is the column_1 object.
type Column1 struct {
	Attributes []Attribute
}

// Column2 is the column_2 object.
type Column2 struct {
	Attributes []Attribute
}

// Configuration is the configuration object.
type Configuration struct {
	Column1  *Column1
	Column2  *Column2
	Settings *Settings
}

// CustomForm is one entry of the custom_forms list.
type CustomForm struct {
	Configuration *Configuration
}

// Settings is the settings object.
type Settings struct {
	CustomForm *string
	FormLayout *string
}

// CustomFeatureRecordShape reads a body as CustomFeatureRecord.
func CustomFeatureRecordShape(decoded any) any { return dto.Read[CustomFeatureRecord](decoded) }

// CustomFeatureList is one page of what GetCustomFeatures lists, under data.
type CustomFeatureList struct {
	Data       []CustomFeatureDetail
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints CustomFeatureList was measured from.
func (CustomFeatureList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/custom_feature",
	}
}

// CustomFeatureDetail is a CustomFeature's custom_feature.
type CustomFeatureDetail struct {
	ActivityLogs          []any
	ChildCustomComponents []any
	CreatedBy             any
	CreatedOn             *string
	CustomComponent       any
	CustomForms           []CustomFeatureDetailCustomForm
	CustomObjects         []any
	Description           *string
	DisplayName           *string
	Href                  *string
	ID                    *string
	LastUpdatedBy         any
	LastUpdatedOn         *string
	Name                  *string
	Object                *string
	Relation              *string
	Status                *string
	UseIn                 *string
	Version               *int64
}

// CustomFeatureDetailCustomForm is one entry of a CustomFeatureDetail's custom_forms.
type CustomFeatureDetailCustomForm struct {
	Configuration *CustomFeatureDetailCustomFormConfiguration
}

// CustomFeatureDetailCustomFormConfiguration is a CustomFeatureDetailCustomForm's configuration.
type CustomFeatureDetailCustomFormConfiguration struct {
	Column1  *CustomFeatureDetailCustomFormConfigurationColumn1
	Column2  *CustomFeatureDetailCustomFormConfigurationColumn2
	Settings *CustomFeatureDetailCustomFormConfigurationSettings
}

// CustomFeatureDetailCustomFormConfigurationColumn1 is A
// CustomFeatureDetailCustomFormConfiguration's column_1.
type CustomFeatureDetailCustomFormConfigurationColumn1 struct {
	Attributes []CustomFeatureDetailCustomFormConfigurationColumn1Attribute
}

// CustomFeatureDetailCustomFormConfigurationColumn1Attribute is one entry of a
// CustomFeatureDetailCustomFormConfigurationColumn1's attributes.
type CustomFeatureDetailCustomFormConfigurationColumn1Attribute struct {
	DataType       *string
	DisplayName    *string
	Name           *string
	ShowInListView *bool
	Type           *string
}

// CustomFeatureDetailCustomFormConfigurationColumn2 is A
// CustomFeatureDetailCustomFormConfiguration's column_2.
type CustomFeatureDetailCustomFormConfigurationColumn2 struct {
	Attributes []CustomFeatureDetailCustomFormConfigurationColumn2Attribute
}

// CustomFeatureDetailCustomFormConfigurationColumn2Attribute is one entry of a
// CustomFeatureDetailCustomFormConfigurationColumn2's attributes.
type CustomFeatureDetailCustomFormConfigurationColumn2Attribute struct {
	Name *string
	Type *string
}

// CustomFeatureDetailCustomFormConfigurationSettings is A
// CustomFeatureDetailCustomFormConfiguration's settings.
type CustomFeatureDetailCustomFormConfigurationSettings struct {
	CustomForm *string
	FormLayout *string
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

// CustomFeatureListShape reads a body as CustomFeatureList.
func CustomFeatureListShape(decoded any) any { return dto.Read[CustomFeatureList](decoded) }
