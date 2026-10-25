package employees

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

// EmployeeRecord is the employee body.
type EmployeeRecord struct {
	ActivityLogs     []any
	Code             *string
	Contacts         []Contact
	CreatedBy        *CreatedBy
	CreatedOn        *string
	CustomAttributes []any
	CustomForm       *CustomForm
	Description      *string
	DisplayName      *string
	EmailAddress     *string
	Href             *string
	ID               *string
	ImageUUID        any
	LastUpdatedBy    *LastUpdatedBy
	LastUpdatedOn    any
	Name             *string
	Object           *string
	Status           *string
	Version          *int64
}

// MeasuredFrom lists the endpoints EmployeeRecord was measured from.
func (EmployeeRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/employees",
		"GET /api/{api_version}/employees/{employee_id}",
		"POST /api/{api_version}/employees",
		"PATCH /api/{api_version}/employees/{employee_id}",
	}
}

// EnvelopeKey unwraps {"employee": {...}}.
func (EmployeeRecord) EnvelopeKey() string { return EmployeeKey }

// EmployeeList is a page of EmployeeRecord records.
type EmployeeList struct {
	Data       []EmployeeRecord
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints EmployeeList was measured from.
func (EmployeeList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/employees",
	}
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// Contact is one entry of the contacts list.
type Contact struct {
	Address       any
	CreatedBy     *CreatedBy
	CreatedOn     *string
	Email         *Email
	Fax           any
	ID            *string
	LastUpdatedBy any
	LastUpdatedOn any
	Mobile        any
	Name          *Name
	Object        *string
	Phone         any
	Status        *string
	Type          *string
	Version       *int64
}

// CreatedBy is the created_by object.
type CreatedBy struct {
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

// Email is the email object.
type Email struct {
	Address       *string
	EmailOptedOut *bool
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// Name is the name object.
type Name struct {
	Designation any
	First       *string
	Last        *string
	Middle      any
	Salutation  any
}

// EmployeeListShape reads a body as EmployeeList.
func EmployeeListShape(decoded any) any { return dto.Read[EmployeeList](decoded) }

// EmployeeRecordShape reads a body as EmployeeRecord.
func EmployeeRecordShape(decoded any) any { return dto.Read[EmployeeRecord](decoded) }
