package portals

import "github.com/exsited/exsited-sdk-go/v4/dto"

// Pagination is the paging block, either shape.
type Pagination struct {
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	PrevCursor any

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	NextPage     *string
	PreviousPage *string
}

// PortalLoginRecord is the portal login body.
type PortalLoginRecord struct {
	User *User
}

// MeasuredFrom lists the endpoints PortalLoginRecord was measured from.
func (PortalLoginRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/portal/login",
	}
}

// EnvelopeKey unwraps {"portal": {...}}.
func (PortalLoginRecord) EnvelopeKey() string { return PortalKey }

// PortalChangePasswordCreateRecord is the body CreatePortalChangePassword answers.
type PortalChangePasswordCreateRecord struct {
	Message *string
}

// MeasuredFrom lists the endpoints PortalChangePasswordCreateRecord was measured from.
func (PortalChangePasswordCreateRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/portal/change_password",
	}
}

// EnvelopeKey unwraps {"portal": {...}}.
func (PortalChangePasswordCreateRecord) EnvelopeKey() string { return PortalKey }

// Portal is one entry of the portals list.
type Portal struct {
	Account     *string
	DisplayName *string
	UUID        *string
}

// User is the user object.
type User struct {
	EmailAddress      *string
	GrantAccessPortal *bool
	ImageURI          *string
	Name              *string
	Portals           []Portal
	Type              *string
}

// PortalLoginRecordShape reads a body as PortalLoginRecord.
func PortalLoginRecordShape(decoded any) any { return dto.Read[PortalLoginRecord](decoded) }

// PortalChangePasswordCreateRecordShape reads a body as PortalChangePasswordCreateRecord.
func PortalChangePasswordCreateRecordShape(decoded any) any {
	return dto.Read[PortalChangePasswordCreateRecord](decoded)
}
