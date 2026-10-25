package items

import "github.com/exsited/exsited-sdk-go/v4/dto"

// The sale, purchase and inventories reads answer a stub, {id, name, sale} with an
// int64 id, not the item, so ItemRecord keeps ID *string and the stubs have types of
// their own. quantity_on_hand is an int64 in inventories.warehouses[] and a string on
// an InventoryAdjustment; neither is coerced.

// Reference is v4's standard embedded pointer. On an item: custom_form, manager,
// group.
type Reference struct {
	Object *string
	ID     *string
	Name   *string
	Href   *string
}

// User is whoever created or last updated a record.
//
// On an item it carries {object, name} and nothing else. The rest are declared so a
// richer body still reads, but do not expect them.
type User struct {
	Object *string
	Name   *string
	ID     *string
	Email  *string
	Href   *string
}

// CustomForm is the form an item was created under. Present on every item.
type CustomForm struct {
	Object *string
	ID     *string
	Name   *string
	Href   *string
}

// ItemCode is one entry of codes: Name always a string, Value a string or null.
// An empty slot is normal, not an error.
type ItemCode struct {
	Name  *string
	Value *string
}

// ItemCurrency is one entry of currencies; all four flags are always present and
// always booleans.
type ItemCurrency struct {
	Name               *string
	UsedForSale        *bool
	DefaultForSale     *bool
	UsedForPurchase    *bool
	DefaultForPurchase *bool
	Suppliers          []any
}

// BundleDetail is one component of a bundle. Quantity is a float, not an int: 1.0
// rather than 1. UOM is rarely set.
type BundleDetail struct {
	ID       *string
	Name     *string
	Quantity *float64
	UOM      *string
}

// FamilyDetail is one member of an item family: the id and name only, with no
// object and no href.
type FamilyDetail struct {
	ItemID   *string
	ItemName *string
}

// VariationOption is one variation axis, in configure_variations or
// use_variations. The element shape of Values is not established, so it is left
// untyped.
type VariationOption struct {
	Object *string
	ID     *string
	Name   *string
	Href   *string
	Values []Value
}

// PackageDetail is one package on an item. LastUpdatedBy and LastUpdatedOn are
// always null, so both are untyped.
type PackageDetail struct {
	Object           *string
	ID               *string
	UUID             *string
	Name             *string
	DisplayName      *string
	Status           *string
	Rank             *int64
	Items            []Item
	BillingPeriods   []BillingPeriod
	CustomAttributes []any
	CreatedOn        *string
	CreatedBy        *User
	LastUpdatedOn    any
	LastUpdatedBy    any
}

// TaxConfiguration says whether a price is entered inclusive of tax.
//
// The key differs by side: the sale block carries
// sale_price_entered_inclusive_of_tax and the purchase block
// purchase_price_entered_inclusive_of_tax. Both are declared so one struct serves
// both; whichever the body did not send reads nil.
type TaxConfiguration struct {
	SalePriceEnteredInclusiveOfTax     *bool
	PurchasePriceEnteredInclusiveOfTax *bool
	SalePriceBasedOn                   *string
	TaxCode                            *TaxCode
}

// Pricing is the pricing block on a sale or purchase.
// LatestUsedPricingVersion is the field the nested filter expression addresses.
type Pricing struct {
	Type                     *string
	Version                  *int64
	LatestUsedPricingVersion *int64
	PricingModule            []PricingModule
}

// SaleCharge is the charge block on a sale.
type SaleCharge struct {
	Type        *string
	PricePeriod *string
	Properties  []Property
}

// SaleBlock is an item's sale configuration.
//
// Named ...Block rather than ItemSale because Go gives a package one namespace and
// ItemSale is already the path constant in urls.go. The fields always sent as null
// stay untyped. DefaultSalePrice is a float or null.
type SaleBlock struct {
	Enabled           *bool
	EnableBackOrder   *bool
	InvoiceNote       *string
	DefaultSalePrice  *float64
	PricingMethod     *string
	SalePriceVariant  *string
	UseOnSalePrice    *bool
	TaxExemptWhenSold *bool
	TaxConfiguration  *TaxConfiguration
	Charge            *SaleCharge
	Pricing           *Pricing
	PricingLevels     []PricingLevel
	PricingSchedules  []any
	PaymentProperties []PaymentProperty
	AccountingCode    any
	DiscountProfile   any
	PricingProfile    any
	ShippingProfile   any
	SalePrice         any
	StartDate         any
	EndDate           any
	Height            *Height
	Length            *Length
	Weight            *Weight
	Width             *Width
}

// PurchaseBlock is an item's purchase configuration. AccountingCode and
// PricingProfile are always null.
type PurchaseBlock struct {
	Enabled                   *bool
	SupplierManagementEnabled *bool
	PurchaseOrderNote         *string
	DefaultPurchasePrice      *float64
	TaxExemptWhenPurchase     *bool
	TaxConfiguration          *TaxConfiguration
	Pricing                   *Pricing
	PurchaseProperties        []PurchaseProperty
	AccountingCode            any
	PricingProfile            any
	Suppliers                 []Supplier
}

// InventoryProperties is the inventory settings block. QuantityAvailableForSale is
// a string.
type InventoryProperties struct {
	Enabled                                  *bool
	EnableBatchTracking                      *bool
	EnableSerialization                      *bool
	EnableInventoryCosting                   *bool
	EnableLowStockNotification               *bool
	EnablePreorder                           *bool
	EnableReordering                         *bool
	LowStockThresholdBasedOn                 *string
	ReorderingThresholdBasedOn               *string
	QuantityAvailableForSale                 *string
	QuantityAvailableForSaleDetermination    []any
	UseATPValueAsQuantityAvailableForWebhook *bool
	DefaultCostPerUnit                       any
	PreorderConfiguration                    []PreorderConfiguration
	PreorderPeriod                           *string
}

// WarehouseInventory is one warehouse's stock line.
//
// The quantities are numbers here, whole except quantity_available and
// quantity_promised, and each has a _value twin that is a string:
// quantity_on_hand: 79999 beside quantity_on_hand_value: "0.00". The same names on
// an InventoryAdjustment are strings instead. Nothing is coerced either way.
// Warehouse and UOM are always null.
type WarehouseInventory struct {
	QuantityOnHand           *int64
	QuantityOnHandValue      *string
	QuantityAvailable        *float64
	QuantityAvailableValue   *string
	QuantityOnOrder          *int64
	QuantityPromised         *float64
	QuantityOnReturn         *int64
	QuantityOnPurchaseReturn *int64
	Warehouse                *Warehouse
	UOM                      any
}

// InventoriesBlock is an item's inventory block.
type InventoriesBlock struct {
	Enabled                   *bool
	EnableWarehouseManagement *bool
	DefaultWarehouse          *string
	InventoryProperties       *InventoryProperties
	Warehouses                []WarehouseInventory
}

// NoteAttachment is one file on a note.
type NoteAttachment struct {
	Object      *string
	ID          *string
	Name        *string
	Href        *string
	ContentType *string
	Filename    *string
	Size        *int64
	UUID        *string
}

// NoteRecord is one note on an item. Group, UserField, LastUpdatedBy and
// LastUpdatedOn are always null.
type NoteRecord struct {
	Object        *string
	ID            *string
	Content       *string
	Version       *int64
	Href          *string
	CreatedOn     *string
	CreatedBy     *User
	Item          *Reference
	Attachments   []NoteAttachment
	Watchers      []Watcher
	Group         any
	UserField     any `wire:"user"`
	LastUpdatedOn any
	LastUpdatedBy *LastUpdatedBy
}

// EnvelopeKey unwraps {"note": {...}}.
func (NoteRecord) EnvelopeKey() string { return NoteKey }

// ItemRecord is one item, as GetItem and the listing return it.
//
// ID is the operator-assigned code, "ITEM-0001", and is what every path takes. The
// sale, purchase and inventories stubs answer an integer id instead, so they have
// types of their own.
//
// BaseUOM and ImageURL are always null and are untyped. UOMs and ActivityLogs are
// always empty, so their element types are unknown.
//
// The sub-resource fields arrive only on the endpoint that asks for them. On a
// plain listing they are absent, not null.
type ItemRecord struct {
	Object      *string
	ID          *string
	Name        *string
	DisplayName *string
	Description *string
	Type        *string
	Status      *string
	Origin      *string
	InvoiceNote *string
	Version     *int64
	Href        *string

	SaleEnabled       *bool
	PurchaseEnabled   *bool
	InventoryEnabled  *bool
	VariationsEnabled *bool

	// The read spelling. Written as is_parent / is_child.
	Parent       *bool
	Child        *bool
	ParentItemID *string

	CustomForm *CustomForm
	Manager    *Reference
	// Group is nearly always null.
	Group *Reference

	CreatedOn *string
	CreatedBy *User
	// Both null on an item never edited.
	LastUpdatedOn *string
	LastUpdatedBy *User

	Codes      []ItemCode
	Currencies []ItemCurrency

	Sale                *SaleBlock
	Purchase            *PurchaseBlock
	Inventories         *InventoriesBlock
	BundleDetails       []BundleDetail
	FamilyDetails       []FamilyDetail
	PackageDetails      []PackageDetail
	ConfigureVariations []VariationOption
	UseVariations       []VariationOption
	CustomObjects       []any

	// Always empty.
	UOMs             []any `wire:"uoms"`
	CustomAttributes []CustomAttribute
	ActivityLogs     []any

	// Always null.
	BaseUOM  any
	ImageURL any
}

// EnvelopeKey unwraps {"item": {...}}, so one type reads both a detail body and a
// single entry out of a listing's data.
func (ItemRecord) EnvelopeKey() string { return ItemKey }

// ItemList is a listing of items. The records sit under "data".
//
// Items page by cursor: every pagination block carries has_more, limit, links,
// next_cursor and prev_cursor. There is no offset shape.
type ItemList struct {
	Object     *string
	URL        *string
	Data       []ItemRecord
	Pagination *Pagination
}

// Pagination is the paging block, cursor shape only.
//
// Unlike accounts.Pagination, which declares both shapes, this declares no offset
// fields: an item listing answers the cursor shape only. If it ever answers an
// offset block, exhttp.ApiResponse.Map still carries it.
type Pagination struct {
	Limit      *int64
	HasMore    *bool
	NextCursor *string
	// PrevCursor is untyped: the API sends null.
	PrevCursor any
	// Links is untyped: an object carrying a relative next URL.
	Links *PaginationLinks
}

// ItemSaleView is what GetItemSale answers: a stub under item, not the item.
//
// Three keys, {id, name, sale}, and ID is an int64, not the string code the path
// took. Use ItemRecord for anything that returns a real item.
type ItemSaleView struct {
	ID   *int64
	Name *string
	Sale *SaleBlock
}

// EnvelopeKey unwraps {"item": {...}}.
func (ItemSaleView) EnvelopeKey() string { return ItemKey }

// ItemPurchaseView is {item: {id, name, purchase}}, ID an int64.
type ItemPurchaseView struct {
	ID       *int64
	Name     *string
	Purchase *PurchaseBlock
}

// EnvelopeKey unwraps {"item": {...}}.
func (ItemPurchaseView) EnvelopeKey() string { return ItemKey }

// ItemInventoriesView is {item: {id, name, inventories}}, ID the string code the path took.
type ItemInventoriesView struct {
	ID          *string
	Name        *string
	Inventories *InventoriesBlock
}

// EnvelopeKey unwraps {"item": {...}}.
func (ItemInventoriesView) EnvelopeKey() string { return ItemKey }

// InventoryAdjustment is one inventory movement.
//
// Every quantity is a string here, "79999", not 79999: the opposite of
// WarehouseInventory, where the same names are integers. CreatedBy and
// LastUpdatedBy are always null.
type InventoryAdjustment struct {
	Object                   *string
	ID                       *string
	ItemID                   *string
	ItemName                 *string
	WarehouseName            *string
	AdjustmentAmount         *string
	AdjustmentType           *string
	Reference                *string
	Status                   *string
	QuantityOnHand           *string
	QuantityAvailable        *string
	QuantityOnOrder          *string
	QuantityPromised         *string
	QuantityOnReturn         *string
	QuantityOnPurchaseReturn *string
	QuantityOnReserve        *string
	CreatedOn                *string
	LastUpdatedOn            *string
	Href                     *string
	CreatedBy                any
	LastUpdatedBy            any
}

// EnvelopeKey unwraps {"inventory_adjustment": {...}}.
func (InventoryAdjustment) EnvelopeKey() string { return InventoryAdjustmentKey }

// InventoryAdjustmentList is a listing of inventory movements.
type InventoryAdjustmentList struct {
	Object     *string
	URL        *string
	Data       []InventoryAdjustment
	Pagination *Pagination
}

// NoteList is a listing of an item's notes.
type NoteList struct {
	Object     *string
	URL        *string
	Data       []NoteRecord
	Pagination *Pagination
}

// NoteFileList is a listing of a note's attachments.
type NoteFileList struct {
	Object     *string
	URL        *string
	Data       []NoteAttachment
	Pagination *Pagination
}

// Every write body is wrapped in an "item" envelope, the sale, purchase and
// inventories writes included: {"item": {"sale": {...}}}, not {"sale": {...}}.
// The one exception is the image upload, whose single field sits at the top level.

// ItemCreateBody is the body CreateItem takes.
//
// Type and Name are the minimum that works. Type takes "STANDARD", "BUNDLE",
// "FAMILY" or "VARIATION_BASED".
//
// The family flags are written is_parent/is_child and read back as parent/child.
//
// Group rejects an unknown value with 400 internal_error "Invalid group".
//
// The nested blocks are `any` rather than the read types: the write shapes differ
// from the read shapes and are not established.
type ItemCreateBody struct {
	Name          *string
	Type          *string
	DisplayName   *string
	Description   *string
	Origin        *string
	InvoiceNote   *string
	Manager       *string
	CustomForm    *string
	Group         *string
	IsParent      *bool
	IsChild       *bool
	ParentItemID  *string
	Currencies    []ItemCurrency
	UOMs          []any `wire:"uoms"`
	Sale          any
	Purchase      any
	Inventories   any
	BundleDetails []any
}

// EnvelopeKey wraps the body in {"item": {...}}.
func (ItemCreateBody) EnvelopeKey() string { return ItemKey }

// ItemInformationUpdate is the body UpdateItemInformation takes. Partial: send
// only what changes.
type ItemInformationUpdate struct {
	Name        *string
	DisplayName *string
	Description *string
	Origin      *string
	InvoiceNote *string
	Manager     *string
	CustomForm  *string
}

// EnvelopeKey wraps the body in {"item": {...}}.
func (ItemInformationUpdate) EnvelopeKey() string { return ItemKey }

// ItemSaleWrite is the body the sale create and update take:
// {"item": {"sale": {...}, "currencies": [...]}}. The sale block sits inside the
// item envelope, and currencies sits beside it rather than within it.
//
// Sale is `any`: {"default_sale_price": 75} and {"charge": {"type": "RECURRING"}}
// are the known forms, not enough to declare a shape.
type ItemSaleWrite struct {
	Sale       any
	Currencies []ItemCurrency
}

// EnvelopeKey wraps the body in {"item": {...}}.
func (ItemSaleWrite) EnvelopeKey() string { return ItemKey }

// ItemPurchaseWrite is {"item": {"purchase": {...}, "currencies": [...]}}.
type ItemPurchaseWrite struct {
	Purchase   any
	Currencies []ItemCurrency
}

// EnvelopeKey wraps the body in {"item": {...}}.
func (ItemPurchaseWrite) EnvelopeKey() string { return ItemKey }

// ItemInventoriesWrite is {"item": {"inventories": {...}}}.
//
// Turning inventory off is refused with 400 "Unable to disable Inventory
// Management" once stock exists. The write spells warehouse management
// warehouse_enabled where the read returns enable_warehouse_management.
type ItemInventoriesWrite struct {
	Inventories any
}

// EnvelopeKey wraps the body in {"item": {...}}.
func (ItemInventoriesWrite) EnvelopeKey() string { return ItemKey }

// InventoryAddBody is
// {"item": {"inventories": {"warehouses": [{name, quantity, accounting_code}]}}},
// with name where the read side calls the same thing warehouse.
type InventoryAddBody struct {
	Inventories any
}

// EnvelopeKey wraps the body in {"item": {...}}.
func (InventoryAddBody) EnvelopeKey() string { return ItemKey }

// ItemVariationsWrite is the body UpdateItemVariations takes, in one of two
// forms: a parent declares ConfigureVariations; a child sets ParentItemID and
// UseVariations.
type ItemVariationsWrite struct {
	ConfigureVariations []any
	UseVariations       []any
	ParentItemID        *string
}

// EnvelopeKey wraps the body in {"item": {...}}.
func (ItemVariationsWrite) EnvelopeKey() string { return ItemKey }

// ItemBundleDetailsWrite is {"item": {"bundle_details": [{id, quantity}]}}: a
// bare list, the opposite of how family_details is written.
type ItemBundleDetailsWrite struct {
	BundleDetails []any
}

// EnvelopeKey wraps the body in {"item": {...}}.
func (ItemBundleDetailsWrite) EnvelopeKey() string { return ItemKey }

// ItemFamilyDetailsWrite is {"item": {"family_details": {"add": ["ITEM-0002"]}}}:
// an object with an add list, not a bare list.
type ItemFamilyDetailsWrite struct {
	FamilyDetails any
}

// EnvelopeKey wraps the body in {"item": {...}}.
func (ItemFamilyDetailsWrite) EnvelopeKey() string { return ItemKey }

// ItemPackageDetailsWrite is {"item": {"package_details": {...}}}.
type ItemPackageDetailsWrite struct {
	PackageDetails any
}

// EnvelopeKey wraps the body in {"item": {...}}.
func (ItemPackageDetailsWrite) EnvelopeKey() string { return ItemKey }

// ItemImageUpload is base64 in a JSON body, not a multipart upload, and the one
// write in this family with no item envelope.
type ItemImageUpload struct {
	ImageBase64 *string `wire:"image_base64"`
}

// ItemNoteFileView is the body GetItemNoteFile answers.
type ItemNoteFileView struct {
	ContentType *string
	Filename    *string
	Href        *string
	Size        *int64
	UUID        *string
}

// MeasuredFrom lists the endpoints ItemNoteFileView was measured from.
func (ItemNoteFileView) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/items/{item_id}/notes/{note_id}/files/{file_id}",
	}
}

// EnvelopeKey unwraps {"file": {...}}.
func (ItemNoteFileView) EnvelopeKey() string { return FileKey }

// Height is the height object.
type Height struct {
	Uom   *string
	Value *float64
}

// Length is the length object.
type Length struct {
	Uom   *string
	Value *float64
}

// PreorderConfiguration is one entry of the preorder_configuration list.
type PreorderConfiguration struct {
	EndDate   *string
	Name      *string
	StartDate *string
}

// Supplier is one entry of the suppliers list.
type Supplier struct {
	AccountingCode        any
	DefaultPurchasePrice  *float64
	ID                    any
	Name                  *string
	Pricing               *SupplierPricing
	PurchaseOrderNote     any
	TaxConfiguration      *SupplierTaxConfiguration
	TaxExemptWhenPurchase *bool
}

// SupplierPricing is the pricing object.
type SupplierPricing struct {
	LatestUsedPricingVersion *int64
	PricingModule            []any
	Type                     *string
	Version                  *int64
}

// SupplierTaxConfiguration is the tax_configuration object.
type SupplierTaxConfiguration struct {
	PurchasePriceEnteredInclusiveOfTax *bool
	TaxCode                            *TaxCode
}

// TaxCode is the tax_code object.
type TaxCode struct {
	Href   *string
	ID     *string
	Object *string
}

// Weight is the weight object.
type Weight struct {
	Uom   *string
	Value *float64
}

// Width is the width object.
type Width struct {
	Uom   *string
	Value *float64
}

// BillingPeriod is one entry of the billing_periods list.
type BillingPeriod struct {
	Name *string
	Rank *int64
}

// CustomAttribute is one entry of the custom_attributes list.
type CustomAttribute struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
	Type   *string
	Value  *string
}

// Item is one entry of the items list.
type Item struct {
	ItemID   *string
	ItemName *string
	Quantity *float64
	Uom      any
}

// LastUpdatedBy is the last_updated_by object.
type LastUpdatedBy struct {
	Name   *string
	Object *string
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next *string
	Prev *string
}

// PaymentProperty is one entry of the payment_properties list.
type PaymentProperty struct {
	Name  *string
	Value *string
}

// PricingLevel is one entry of the pricing_levels list.
type PricingLevel struct {
	Name *string
	UUID *string
}

// PricingModule is one entry of the pricing_module list.
type PricingModule struct {
	AveragePurchasePrice any
	Currency             *string
	LastPurchasePrice    any
	Price                *float64
	PriceType            *string
	PricingLevel         *string
	Unit                 *string
}

// Property is one entry of the properties list.
type Property struct {
	Name  *string
	Value *string
}

// PurchaseProperty is one entry of the purchase_properties list.
type PurchaseProperty struct {
	Name  *string
	Value *string
}

// Value is one entry of the values list.
type Value struct {
	Name *string
	UUID *string
}

// Warehouse is the warehouse object.
type Warehouse struct {
	Href   *string
	ID     *string
	Name   *string
	Object *string
}

// Watcher is one entry of the watchers list.
type Watcher struct {
	Email *string
	Name  *string
	Type  *string
}

// ItemListShape reads a body as ItemList.
func ItemListShape(decoded any) any { return dto.Read[ItemList](decoded) }

// ItemRecordShape reads a body as ItemRecord.
func ItemRecordShape(decoded any) any { return dto.Read[ItemRecord](decoded) }

// ItemSaleViewShape reads a body as ItemSaleView.
func ItemSaleViewShape(decoded any) any { return dto.Read[ItemSaleView](decoded) }

// ItemPurchaseViewShape reads a body as ItemPurchaseView.
func ItemPurchaseViewShape(decoded any) any { return dto.Read[ItemPurchaseView](decoded) }

// ItemInventoriesViewShape reads a body as ItemInventoriesView.
func ItemInventoriesViewShape(decoded any) any { return dto.Read[ItemInventoriesView](decoded) }

// InventoryAdjustmentListShape reads a body as InventoryAdjustmentList.
func InventoryAdjustmentListShape(decoded any) any { return dto.Read[InventoryAdjustmentList](decoded) }

// NoteListShape reads a body as NoteList.
func NoteListShape(decoded any) any { return dto.Read[NoteList](decoded) }

// NoteRecordShape reads a body as NoteRecord.
func NoteRecordShape(decoded any) any { return dto.Read[NoteRecord](decoded) }

// NoteFileListShape reads a body as NoteFileList.
func NoteFileListShape(decoded any) any { return dto.Read[NoteFileList](decoded) }

// ItemNoteFileViewShape reads a body as ItemNoteFileView.
func ItemNoteFileViewShape(decoded any) any { return dto.Read[ItemNoteFileView](decoded) }
