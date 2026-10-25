package returnvendorauthorizations

import "github.com/exsited/exsited-sdk-go/v4/dto"

// ReturnVendorAuthorizationView is what GetReturnVendorAuthorization answers, under
// return_vendor_authorization: the listing's record with its lines and receive_rvas.
type ReturnVendorAuthorizationView struct {
	Account       *ReturnVendorAuthorizationAccount
	Code          *string
	CreatedBy     *ReturnVendorAuthorizationCreatedBy
	CreatedOn     *string
	Href          *string
	ID            *string
	ItemCount     *int64
	LastUpdatedBy any
	LastUpdatedOn any
	Note          *string
	Object        *string
	PurchaseOrder *ReturnVendorAuthorizationPurchaseOrder
	ReturnStatus  *string
	RvaDate       *string
	Status        *string
	Version       *int64
	Lines         []any
	ReceiveRvas   []any
}

// MeasuredFrom lists the endpoints ReturnVendorAuthorizationView was measured from.
func (ReturnVendorAuthorizationView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/return_vendor_authorizations/{rva_id}",
	}
}

// EnvelopeKey unwraps {"return_vendor_authorization": {...}}.
func (ReturnVendorAuthorizationView) EnvelopeKey() string { return ReturnVendorAuthorizationKey }

// ReturnVendorAuthorizationAccount is a ReturnVendorAuthorization's account.
type ReturnVendorAuthorizationAccount struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// ReturnVendorAuthorizationCreatedBy is a ReturnVendorAuthorization's created_by.
type ReturnVendorAuthorizationCreatedBy struct {
	Name   *string
	Object *string
}

// ReturnVendorAuthorizationPurchaseOrder is a ReturnVendorAuthorization's purchase_order.
type ReturnVendorAuthorizationPurchaseOrder struct {
	Code   *string
	Href   *string
	ID     *string
	Object *string
}

// ReturnVendorAuthorizationList is one page of ReturnVendorAuthorization records, under data; paged
// by cursor.
type ReturnVendorAuthorizationList struct {
	Data       []ReturnVendorAuthorization
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ReturnVendorAuthorizationList was measured from.
func (ReturnVendorAuthorizationList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/return_vendor_authorizations",
	}
}

// ReturnVendorAuthorization is one return vendor authorization, as the listing returns it.
type ReturnVendorAuthorization struct {
	Account       *ReturnVendorAuthorizationAccount
	Code          *string
	CreatedBy     *ReturnVendorAuthorizationCreatedBy
	CreatedOn     *string
	Href          *string
	ID            *string
	ItemCount     *int64
	LastUpdatedBy any
	LastUpdatedOn any
	Note          *string
	Object        *string
	PurchaseOrder *ReturnVendorAuthorizationPurchaseOrder
	ReturnStatus  *string
	RvaDate       *string
	Status        *string
	Version       *int64
}

// MeasuredFrom lists the endpoints ReturnVendorAuthorization was measured from.
func (ReturnVendorAuthorization) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/return_vendor_authorizations",
	}
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

// ReturnVendorAuthorizationViewShape reads a body as ReturnVendorAuthorizationView.
func ReturnVendorAuthorizationViewShape(decoded any) any {
	return dto.Read[ReturnVendorAuthorizationView](decoded)
}

// ReturnVendorAuthorizationListShape reads a body as ReturnVendorAuthorizationList.
func ReturnVendorAuthorizationListShape(decoded any) any {
	return dto.Read[ReturnVendorAuthorizationList](decoded)
}
