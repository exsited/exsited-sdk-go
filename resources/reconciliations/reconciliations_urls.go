// Package reconciliations is the v4 reconciliations endpoints.
package reconciliations

// The path templates for the reconciliations endpoints; the client fills {api_version}.
const (
	ReconciliationSalePaymentsPath = "/api/{api_version}/reconciliations/{reconciliation_id}/sale_payments"
	ReconciliationSaleRefundsPath  = "/api/{api_version}/reconciliations/{reconciliation_id}/sale_refunds"
)

// ReconciliationSalePaymentsKey is the key a listing's records arrive under, the v4 list envelope's data.
const ReconciliationSalePaymentsKey = "data"
