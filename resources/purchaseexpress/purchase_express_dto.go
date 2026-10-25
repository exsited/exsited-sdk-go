package purchaseexpress

import "github.com/exsited/exsited-sdk-go/v4/dto"

// PurchaseExpressRecord is what CreatePurchaseExpress answers, under purchase_express.
type PurchaseExpressRecord struct {
	Account                *PurchaseExpressAccount
	ActivityLogs           []any
	BillTotal              *string
	Code                   *string
	CreatedBy              *PurchaseExpressCreatedBy
	CreatedOn              *string
	Currency               *PurchaseExpressCurrency
	CustomAttributes       []any
	CustomForm             *PurchaseExpressCustomForm
	DueDate                *string
	DueTotal               *string
	ExpectedCompletionDate any
	Href                   *string
	ID                     *string
	IsPriceTaxInclusive    any
	IssueDate              *string
	Kpis                   *PurchaseExpressKpis
	LastUpdatedBy          any
	LastUpdatedOn          *string
	Object                 *string
	OperationalStatus      *string
	Origin                 any
	PaidAmount             *string
	PurchaseOrderNote      any
	SaleOrder              any
	Status                 *string
	Subtotal               *string
	SupplierInvoiceID      any
	TaxTotal               *string
	Total                  *string
	Version                *int64
	Warehouse              any
}

// MeasuredFrom lists the endpoints PurchaseExpressRecord was measured from.
func (PurchaseExpressRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/purchase_express",
	}
}

// EnvelopeKey unwraps {"purchase_express": {...}}.
func (PurchaseExpressRecord) EnvelopeKey() string { return PurchaseExpressKey }

// PurchaseExpressAccount is a PurchaseExpressRecord's account.
type PurchaseExpressAccount struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PurchaseExpressCreatedBy is a PurchaseExpressRecord's created_by.
type PurchaseExpressCreatedBy struct {
	Name   *string
	Object *string
}

// PurchaseExpressCurrency is a PurchaseExpressRecord's currency.
type PurchaseExpressCurrency struct {
	Code   *string
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PurchaseExpressCustomForm is a PurchaseExpressRecord's custom_form.
type PurchaseExpressCustomForm struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// PurchaseExpressKpis is a PurchaseExpressRecord's kpis.
type PurchaseExpressKpis struct {
	EstimatedTotal       *string
	IssueDate            any
	LastCancelledOn      any
	LastChangedOn        any
	LastDeletedOn        any
	LastInvoiceIssueDate any
	LastInvoiceTotal     *string
	LastReactivatedOn    any
	Object               *string
	TotalExpense         *string
	TotalOutstanding     *string
	TotalOverdue         *string
	TotalPurchaseInvoice *int64
}

// PurchaseExpressRecordShape reads a body as PurchaseExpressRecord.
func PurchaseExpressRecordShape(decoded any) any { return dto.Read[PurchaseExpressRecord](decoded) }
