// Package returnmerchandiseauthorisations is the v4 return merchandise authorisations endpoints.
package returnmerchandiseauthorisations

// The path templates for the return merchandise authorisations endpoints; the client fills {api_version}.
const (
	ReturnMerchandiseAuthorisationsPath                    = "/api/{api_version}/return_merchandise_authorisations"
	ReturnMerchandiseAuthorisationPath                     = "/api/{api_version}/return_merchandise_authorisations/{rma_id}"
	InvoiceReturnMerchandiseAuthorisationReceivePath       = "/api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations/{rma_code}/receive"
	InvoiceReturnMerchandiseAuthorisationReceiveByUUIDPath = "/api/{api_version}/invoices/{invoice_id}/return_merchandise_authorisations/{rma_code}/receive/{receive_uuid}"
	ReturnMerchandiseAuthorisationNotesPath                = "/api/{api_version}/return_merchandise_authorisations/{rma_id}/notes"
	ReturnMerchandiseAuthorisationReceivePath              = "/api/{api_version}/return_merchandise_authorisations/{rma_id}/receive"
	ReturnMerchandiseAuthorisationApprovePath              = "/api/{api_version}/return_merchandise_authorisations/{rma_id}/approve"
	ReturnMerchandiseAuthorisationNotePath                 = "/api/{api_version}/return_merchandise_authorisations/{rma_id}/notes/{note_id}"
	ReturnMerchandiseAuthorisationNoteFilePath             = "/api/{api_version}/return_merchandise_authorisations/{rma_id}/notes/{note_id}/files/{file_id}"
	ReturnMerchandiseAuthorisationRejectPath               = "/api/{api_version}/return_merchandise_authorisations/{rma_id}/reject"
	ReturnMerchandiseAuthorisationNotesByCodePath          = "/api/{api_version}/return_merchandise_authorisations/{rma_code}/notes"
	ReturnMerchandiseAuthorisationNoteByCodePath           = "/api/{api_version}/return_merchandise_authorisations/{rma_code}/notes/{note_id}"
	ReturnMerchandiseAuthorisationNoteFilesByCodePath      = "/api/{api_version}/return_merchandise_authorisations/{rma_code}/notes/{note_id}/files"
	ReturnMerchandiseAuthorisationNoteFileByCodePath       = "/api/{api_version}/return_merchandise_authorisations/{rma_code}/notes/{note_id}/files/{file_id}"
	ReturnMerchandiseAuthorisationReceiveByCodePath        = "/api/{api_version}/return_merchandise_authorisations/{rma_code}/receive"
	ReturnMerchandiseAuthorisationReceiveByUUIDPath        = "/api/{api_version}/return_merchandise_authorisations/{rma_code}/receive/{receive_uuid}"
)

// ReturnMerchandiseAuthorisationsKey is the key a listing's records arrive under, the v4 list envelope's data.
const ReturnMerchandiseAuthorisationsKey = "data"

// InvoiceKey is the envelope key one record arrives under.
const InvoiceKey = "invoice"

// ReceiveRmaKey is the envelope key one record arrives under.
const ReceiveRmaKey = "receive_rma"

// ReturnMerchandiseAuthorisationKey is the envelope key one record arrives under.
const ReturnMerchandiseAuthorisationKey = "return_merchandise_authorisation"
