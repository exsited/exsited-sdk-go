// Package returnvendorauthorizations is the v4 return vendor authorizations endpoints.
package returnvendorauthorizations

// The path templates for the return vendor authorizations endpoints; the client fills {api_version}.
const (
	ReturnVendorAuthorizationsPath        = "/api/{api_version}/return_vendor_authorizations"
	ReturnVendorAuthorizationPath         = "/api/{api_version}/return_vendor_authorizations/{rva_id}"
	ReturnVendorAuthorizationNotePath     = "/api/{api_version}/return_vendor_authorizations/{rva_id}/notes/{note_id}"
	ReturnVendorAuthorizationNoteFilePath = "/api/{api_version}/return_vendor_authorizations/{rva_id}/notes/{note_id}/files/{file_id}"
)

// ReturnVendorAuthorizationsKey is the key a listing's records arrive under, the v4 list envelope's data.
const ReturnVendorAuthorizationsKey = "data"

// ReturnVendorAuthorizationKey is the envelope key one record arrives under.
const ReturnVendorAuthorizationKey = "return_vendor_authorization"
