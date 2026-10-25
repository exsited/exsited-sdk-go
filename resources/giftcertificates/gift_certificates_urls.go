// Package giftcertificates is the v4 gift certificates endpoints.
package giftcertificates

// The path templates for the gift certificates endpoints; the client fills {api_version}.
const (
	GiftCertificatesPath             = "/api/{api_version}/gift_certificates"
	GiftCertificatePath              = "/api/{api_version}/gift_certificates/{gc_id}"
	GiftCertificateAllocationsPath   = "/api/{api_version}/gift_certificates/{gc_id}/allocations"
	GiftCertificateTransactionsPath  = "/api/{api_version}/gift_certificates/{gc_id}/transactions"
	GiftCertificateAmendPath         = "/api/{api_version}/gift_certificates/{gc_id}/amend"
	GiftCertificateDeallocationsPath = "/api/{api_version}/gift_certificates/{gc_id}/deallocations"
	GiftCertificateAllocationPath    = "/api/{api_version}/gift_certificates/{gc_id}/allocations/{allocation_id}"
	GiftCertificateTransactionPath   = "/api/{api_version}/gift_certificates/{gc_id}/transactions/{transaction_id}"
)

// GiftCertificatesKey is the key a listing's records arrive under, the v4 list envelope's data.
const GiftCertificatesKey = "data"

// GiftCertificateKey is the envelope key one record arrives under.
const GiftCertificateKey = "gift_certificate"

// GiftCertificateAllocationKey is the envelope key one record arrives under.
const GiftCertificateAllocationKey = "gift_certificate_allocation"

// GiftCertificateTransactionKey is the envelope key one record arrives under.
const GiftCertificateTransactionKey = "gift_certificate_transaction"
