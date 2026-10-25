package client

import (
	"sort"
	"strings"
	"sync"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
	"github.com/exsited/exsited-sdk-go/v4/resources/abafiles"
	"github.com/exsited/exsited-sdk-go/v4/resources/accountemails"
	"github.com/exsited/exsited-sdk-go/v4/resources/accountgroups"
	"github.com/exsited/exsited-sdk-go/v4/resources/accountingcodes"
	"github.com/exsited/exsited-sdk-go/v4/resources/accounts"
	"github.com/exsited/exsited-sdk-go/v4/resources/actionscripts"
	"github.com/exsited/exsited-sdk-go/v4/resources/additionalcodes"
	"github.com/exsited/exsited-sdk-go/v4/resources/alertresponsibilities"
	"github.com/exsited/exsited-sdk-go/v4/resources/alertresponsibility"
	"github.com/exsited/exsited-sdk-go/v4/resources/alerts"
	"github.com/exsited/exsited-sdk-go/v4/resources/applications"
	"github.com/exsited/exsited-sdk-go/v4/resources/approvalprofiles"
	"github.com/exsited/exsited-sdk-go/v4/resources/approvals"
	"github.com/exsited/exsited-sdk-go/v4/resources/approvalusergroups"
	"github.com/exsited/exsited-sdk-go/v4/resources/attributeattachments"
	"github.com/exsited/exsited-sdk-go/v4/resources/billingperiods"
	"github.com/exsited/exsited-sdk-go/v4/resources/bookings"
	"github.com/exsited/exsited-sdk-go/v4/resources/chargingperiods"
	"github.com/exsited/exsited-sdk-go/v4/resources/communicationprofiles"
	"github.com/exsited/exsited-sdk-go/v4/resources/communications"
	"github.com/exsited/exsited-sdk-go/v4/resources/countries"
	"github.com/exsited/exsited-sdk-go/v4/resources/creditnoteapplications"
	"github.com/exsited/exsited-sdk-go/v4/resources/creditnotes"
	"github.com/exsited/exsited-sdk-go/v4/resources/customattributes"
	"github.com/exsited/exsited-sdk-go/v4/resources/customcomponents"
	"github.com/exsited/exsited-sdk-go/v4/resources/customevents"
	"github.com/exsited/exsited-sdk-go/v4/resources/customfeatures"
	"github.com/exsited/exsited-sdk-go/v4/resources/customfields"
	"github.com/exsited/exsited-sdk-go/v4/resources/customforms"
	"github.com/exsited/exsited-sdk-go/v4/resources/dayofmonth"
	"github.com/exsited/exsited-sdk-go/v4/resources/departments"
	"github.com/exsited/exsited-sdk-go/v4/resources/designations"
	"github.com/exsited/exsited-sdk-go/v4/resources/documenttemplates"
	"github.com/exsited/exsited-sdk-go/v4/resources/emailtemplates"
	"github.com/exsited/exsited-sdk-go/v4/resources/employees"
	"github.com/exsited/exsited-sdk-go/v4/resources/events"
	"github.com/exsited/exsited-sdk-go/v4/resources/express"
	"github.com/exsited/exsited-sdk-go/v4/resources/externaldatabases"
	"github.com/exsited/exsited-sdk-go/v4/resources/fulfillmentprofiles"
	"github.com/exsited/exsited-sdk-go/v4/resources/fulfillmentterms"
	"github.com/exsited/exsited-sdk-go/v4/resources/giftcertificates"
	"github.com/exsited/exsited-sdk-go/v4/resources/integrationconnections"
	"github.com/exsited/exsited-sdk-go/v4/resources/inventoryadjustments"
	"github.com/exsited/exsited-sdk-go/v4/resources/invoices"
	"github.com/exsited/exsited-sdk-go/v4/resources/invoiceterms"
	"github.com/exsited/exsited-sdk-go/v4/resources/itemfulfillments"
	"github.com/exsited/exsited-sdk-go/v4/resources/itemreceipts"
	"github.com/exsited/exsited-sdk-go/v4/resources/items"
	"github.com/exsited/exsited-sdk-go/v4/resources/labours"
	"github.com/exsited/exsited-sdk-go/v4/resources/milestones"
	"github.com/exsited/exsited-sdk-go/v4/resources/noteattachments"
	"github.com/exsited/exsited-sdk-go/v4/resources/notegroups"
	"github.com/exsited/exsited-sdk-go/v4/resources/orders"
	"github.com/exsited/exsited-sdk-go/v4/resources/organisation"
	"github.com/exsited/exsited-sdk-go/v4/resources/organisations"
	"github.com/exsited/exsited-sdk-go/v4/resources/payments"
	"github.com/exsited/exsited-sdk-go/v4/resources/portals"
	"github.com/exsited/exsited-sdk-go/v4/resources/pricingperiods"
	"github.com/exsited/exsited-sdk-go/v4/resources/proformainvoices"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchasecreditnoteapplications"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchasecreditnotes"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchaseexpress"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchaseinvoices"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchaseorders"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchasepayments"
	"github.com/exsited/exsited-sdk-go/v4/resources/purchaserefunds"
	"github.com/exsited/exsited-sdk-go/v4/resources/quoteorders"
	"github.com/exsited/exsited-sdk-go/v4/resources/receiveterms"
	"github.com/exsited/exsited-sdk-go/v4/resources/reconciliations"
	"github.com/exsited/exsited-sdk-go/v4/resources/refunds"
	"github.com/exsited/exsited-sdk-go/v4/resources/rentalassets"
	"github.com/exsited/exsited-sdk-go/v4/resources/reports"
	"github.com/exsited/exsited-sdk-go/v4/resources/returnmerchandiseauthorisations"
	"github.com/exsited/exsited-sdk-go/v4/resources/returnvendorauthorizations"
	"github.com/exsited/exsited-sdk-go/v4/resources/roles"
	"github.com/exsited/exsited-sdk-go/v4/resources/salesreps"
	"github.com/exsited/exsited-sdk-go/v4/resources/salutations"
	"github.com/exsited/exsited-sdk-go/v4/resources/settings"
	"github.com/exsited/exsited-sdk-go/v4/resources/settingsworkflows"
	"github.com/exsited/exsited-sdk-go/v4/resources/shippingpackages"
	"github.com/exsited/exsited-sdk-go/v4/resources/shippingrules"
	"github.com/exsited/exsited-sdk-go/v4/resources/systemevents"
	"github.com/exsited/exsited-sdk-go/v4/resources/tasks"
	"github.com/exsited/exsited-sdk-go/v4/resources/teams"
	"github.com/exsited/exsited-sdk-go/v4/resources/timeperiods"
	"github.com/exsited/exsited-sdk-go/v4/resources/timesheets"
	"github.com/exsited/exsited-sdk-go/v4/resources/uom"
	"github.com/exsited/exsited-sdk-go/v4/resources/usages"
	"github.com/exsited/exsited-sdk-go/v4/resources/webhooks"
	"github.com/exsited/exsited-sdk-go/v4/resources/workflows"
	"github.com/exsited/exsited-sdk-go/v4/resources/zones"
)

// An explicit registry of endpoint families: one is added without touching the client, and a typo is refused.

// ResourceFactory builds one resource family over a client.
type ResourceFactory func(requester exhttp.Requester) any

// resourceModules is the registry, guarded because registration is process-wide and may race.
var (
	resourceModulesMutex sync.RWMutex
	resourceModules      = map[string]ResourceFactory{
		abafiles.ResourceName:                        func(requester exhttp.Requester) any { return abafiles.New(requester) },
		accountemails.ResourceName:                   func(requester exhttp.Requester) any { return accountemails.New(requester) },
		accountgroups.ResourceName:                   func(requester exhttp.Requester) any { return accountgroups.New(requester) },
		accountingcodes.ResourceName:                 func(requester exhttp.Requester) any { return accountingcodes.New(requester) },
		accounts.ResourceName:                        func(requester exhttp.Requester) any { return accounts.New(requester) },
		actionscripts.ResourceName:                   func(requester exhttp.Requester) any { return actionscripts.New(requester) },
		additionalcodes.ResourceName:                 func(requester exhttp.Requester) any { return additionalcodes.New(requester) },
		alertresponsibilities.ResourceName:           func(requester exhttp.Requester) any { return alertresponsibilities.New(requester) },
		alertresponsibility.ResourceName:             func(requester exhttp.Requester) any { return alertresponsibility.New(requester) },
		alerts.ResourceName:                          func(requester exhttp.Requester) any { return alerts.New(requester) },
		applications.ResourceName:                    func(requester exhttp.Requester) any { return applications.New(requester) },
		approvalprofiles.ResourceName:                func(requester exhttp.Requester) any { return approvalprofiles.New(requester) },
		approvals.ResourceName:                       func(requester exhttp.Requester) any { return approvals.New(requester) },
		approvalusergroups.ResourceName:              func(requester exhttp.Requester) any { return approvalusergroups.New(requester) },
		attributeattachments.ResourceName:            func(requester exhttp.Requester) any { return attributeattachments.New(requester) },
		billingperiods.ResourceName:                  func(requester exhttp.Requester) any { return billingperiods.New(requester) },
		bookings.ResourceName:                        func(requester exhttp.Requester) any { return bookings.New(requester) },
		chargingperiods.ResourceName:                 func(requester exhttp.Requester) any { return chargingperiods.New(requester) },
		communicationprofiles.ResourceName:           func(requester exhttp.Requester) any { return communicationprofiles.New(requester) },
		communications.ResourceName:                  func(requester exhttp.Requester) any { return communications.New(requester) },
		countries.ResourceName:                       func(requester exhttp.Requester) any { return countries.New(requester) },
		creditnoteapplications.ResourceName:          func(requester exhttp.Requester) any { return creditnoteapplications.New(requester) },
		creditnotes.ResourceName:                     func(requester exhttp.Requester) any { return creditnotes.New(requester) },
		customattributes.ResourceName:                func(requester exhttp.Requester) any { return customattributes.New(requester) },
		customcomponents.ResourceName:                func(requester exhttp.Requester) any { return customcomponents.New(requester) },
		customevents.ResourceName:                    func(requester exhttp.Requester) any { return customevents.New(requester) },
		customfeatures.ResourceName:                  func(requester exhttp.Requester) any { return customfeatures.New(requester) },
		customfields.ResourceName:                    func(requester exhttp.Requester) any { return customfields.New(requester) },
		customforms.ResourceName:                     func(requester exhttp.Requester) any { return customforms.New(requester) },
		dayofmonth.ResourceName:                      func(requester exhttp.Requester) any { return dayofmonth.New(requester) },
		departments.ResourceName:                     func(requester exhttp.Requester) any { return departments.New(requester) },
		designations.ResourceName:                    func(requester exhttp.Requester) any { return designations.New(requester) },
		documenttemplates.ResourceName:               func(requester exhttp.Requester) any { return documenttemplates.New(requester) },
		emailtemplates.ResourceName:                  func(requester exhttp.Requester) any { return emailtemplates.New(requester) },
		employees.ResourceName:                       func(requester exhttp.Requester) any { return employees.New(requester) },
		events.ResourceName:                          func(requester exhttp.Requester) any { return events.New(requester) },
		express.ResourceName:                         func(requester exhttp.Requester) any { return express.New(requester) },
		externaldatabases.ResourceName:               func(requester exhttp.Requester) any { return externaldatabases.New(requester) },
		fulfillmentprofiles.ResourceName:             func(requester exhttp.Requester) any { return fulfillmentprofiles.New(requester) },
		fulfillmentterms.ResourceName:                func(requester exhttp.Requester) any { return fulfillmentterms.New(requester) },
		giftcertificates.ResourceName:                func(requester exhttp.Requester) any { return giftcertificates.New(requester) },
		integrationconnections.ResourceName:          func(requester exhttp.Requester) any { return integrationconnections.New(requester) },
		inventoryadjustments.ResourceName:            func(requester exhttp.Requester) any { return inventoryadjustments.New(requester) },
		invoices.ResourceName:                        func(requester exhttp.Requester) any { return invoices.New(requester) },
		invoiceterms.ResourceName:                    func(requester exhttp.Requester) any { return invoiceterms.New(requester) },
		itemfulfillments.ResourceName:                func(requester exhttp.Requester) any { return itemfulfillments.New(requester) },
		itemreceipts.ResourceName:                    func(requester exhttp.Requester) any { return itemreceipts.New(requester) },
		items.ResourceName:                           func(requester exhttp.Requester) any { return items.New(requester) },
		labours.ResourceName:                         func(requester exhttp.Requester) any { return labours.New(requester) },
		milestones.ResourceName:                      func(requester exhttp.Requester) any { return milestones.New(requester) },
		noteattachments.ResourceName:                 func(requester exhttp.Requester) any { return noteattachments.New(requester) },
		notegroups.ResourceName:                      func(requester exhttp.Requester) any { return notegroups.New(requester) },
		orders.ResourceName:                          func(requester exhttp.Requester) any { return orders.New(requester) },
		organisation.ResourceName:                    func(requester exhttp.Requester) any { return organisation.New(requester) },
		organisations.ResourceName:                   func(requester exhttp.Requester) any { return organisations.New(requester) },
		payments.ResourceName:                        func(requester exhttp.Requester) any { return payments.New(requester) },
		portals.ResourceName:                         func(requester exhttp.Requester) any { return portals.New(requester) },
		pricingperiods.ResourceName:                  func(requester exhttp.Requester) any { return pricingperiods.New(requester) },
		proformainvoices.ResourceName:                func(requester exhttp.Requester) any { return proformainvoices.New(requester) },
		purchasecreditnoteapplications.ResourceName:  func(requester exhttp.Requester) any { return purchasecreditnoteapplications.New(requester) },
		purchasecreditnotes.ResourceName:             func(requester exhttp.Requester) any { return purchasecreditnotes.New(requester) },
		purchaseexpress.ResourceName:                 func(requester exhttp.Requester) any { return purchaseexpress.New(requester) },
		purchaseinvoices.ResourceName:                func(requester exhttp.Requester) any { return purchaseinvoices.New(requester) },
		purchaseorders.ResourceName:                  func(requester exhttp.Requester) any { return purchaseorders.New(requester) },
		purchasepayments.ResourceName:                func(requester exhttp.Requester) any { return purchasepayments.New(requester) },
		purchaserefunds.ResourceName:                 func(requester exhttp.Requester) any { return purchaserefunds.New(requester) },
		quoteorders.ResourceName:                     func(requester exhttp.Requester) any { return quoteorders.New(requester) },
		receiveterms.ResourceName:                    func(requester exhttp.Requester) any { return receiveterms.New(requester) },
		reconciliations.ResourceName:                 func(requester exhttp.Requester) any { return reconciliations.New(requester) },
		refunds.ResourceName:                         func(requester exhttp.Requester) any { return refunds.New(requester) },
		rentalassets.ResourceName:                    func(requester exhttp.Requester) any { return rentalassets.New(requester) },
		reports.ResourceName:                         func(requester exhttp.Requester) any { return reports.New(requester) },
		returnmerchandiseauthorisations.ResourceName: func(requester exhttp.Requester) any { return returnmerchandiseauthorisations.New(requester) },
		returnvendorauthorizations.ResourceName:      func(requester exhttp.Requester) any { return returnvendorauthorizations.New(requester) },
		roles.ResourceName:                           func(requester exhttp.Requester) any { return roles.New(requester) },
		salesreps.ResourceName:                       func(requester exhttp.Requester) any { return salesreps.New(requester) },
		salutations.ResourceName:                     func(requester exhttp.Requester) any { return salutations.New(requester) },
		settings.ResourceName:                        func(requester exhttp.Requester) any { return settings.New(requester) },
		settingsworkflows.ResourceName:               func(requester exhttp.Requester) any { return settingsworkflows.New(requester) },
		shippingpackages.ResourceName:                func(requester exhttp.Requester) any { return shippingpackages.New(requester) },
		shippingrules.ResourceName:                   func(requester exhttp.Requester) any { return shippingrules.New(requester) },
		systemevents.ResourceName:                    func(requester exhttp.Requester) any { return systemevents.New(requester) },
		tasks.ResourceName:                           func(requester exhttp.Requester) any { return tasks.New(requester) },
		teams.ResourceName:                           func(requester exhttp.Requester) any { return teams.New(requester) },
		timeperiods.ResourceName:                     func(requester exhttp.Requester) any { return timeperiods.New(requester) },
		timesheets.ResourceName:                      func(requester exhttp.Requester) any { return timesheets.New(requester) },
		uom.ResourceName:                             func(requester exhttp.Requester) any { return uom.New(requester) },
		usages.ResourceName:                          func(requester exhttp.Requester) any { return usages.New(requester) },
		webhooks.ResourceName:                        func(requester exhttp.Requester) any { return webhooks.New(requester) },
		workflows.ResourceName:                       func(requester exhttp.Requester) any { return workflows.New(requester) },
		zones.ResourceName:                           func(requester exhttp.Requester) any { return zones.New(requester) },
	}
)

// clientAttributeNames are the names a client already answers by itself.
var clientAttributeNames = map[string]bool{
	"aba_files":                         false,
	"account_emails":                    false,
	"account_groups":                    false,
	"accounting_codes":                  false,
	"accounts":                          false,
	"action_scripts":                    false,
	"additional_codes":                  false,
	"alert_responsibilities":            false,
	"alert_responsibility":              false,
	"alerts":                            false,
	"applications":                      false,
	"approval_profiles":                 false,
	"approval_user_groups":              false,
	"approvals":                         false,
	"attribute_attachments":             false,
	"billing_periods":                   false,
	"bookings":                          false,
	"charging_periods":                  false,
	"communication_profiles":            false,
	"communications":                    false,
	"countries":                         false,
	"credit_note_applications":          false,
	"credit_notes":                      false,
	"custom_attributes":                 false,
	"custom_components":                 false,
	"custom_events":                     false,
	"custom_features":                   false,
	"custom_fields":                     false,
	"custom_forms":                      false,
	"day_of_month":                      false,
	"departments":                       false,
	"designations":                      false,
	"document_templates":                false,
	"email_templates":                   false,
	"employees":                         false,
	"events":                            false,
	"express":                           false,
	"external_databases":                false,
	"fulfillment_profiles":              false,
	"fulfillment_terms":                 false,
	"gift_certificates":                 false,
	"integration_connections":           false,
	"inventory_adjustments":             false,
	"invoice_terms":                     false,
	"invoices":                          false,
	"item_fulfillments":                 false,
	"item_receipts":                     false,
	"items":                             false,
	"labours":                           false,
	"milestones":                        false,
	"note_attachments":                  false,
	"note_groups":                       false,
	"orders":                            false,
	"organisation":                      false,
	"organisations":                     false,
	"payments":                          false,
	"portals":                           false,
	"pricing_periods":                   false,
	"proforma_invoices":                 false,
	"purchase_credit_note_applications": false,
	"purchase_credit_notes":             false,
	"purchase_express":                  false,
	"purchase_invoices":                 false,
	"purchase_orders":                   false,
	"purchase_payments":                 false,
	"purchase_refunds":                  false,
	"quote_orders":                      false,
	"receive_terms":                     false,
	"reconciliations":                   false,
	"refunds":                           false,
	"rental_assets":                     false,
	"reports":                           false,
	"return_merchandise_authorisations": false,
	"return_vendor_authorizations":      false,
	"roles":                             false,
	"sales_reps":                        false,
	"salutations":                       false,
	"settings":                          false,
	"settings_workflows":                false,
	"shipping_packages":                 false,
	"shipping_rules":                    false,
	"system_events":                     false,
	"tasks":                             false,
	"teams":                             false,
	"time_periods":                      false,
	"timesheets":                        false,
	"uom":                               false,
	"usages":                            false,
	"webhooks":                          false,
	"workflows":                         false,
	"zones":                             false,
	"close":                             true,
	"close_client":                      true,
	"config":                            true,
	"name":                              true,
	"resource":                          true,
	"iter_pages":                        true,
	"iter_records":                      true,
	"send_api_request":                  true,
	"send_get_request":                  true,
	"send_post_request":                 true,
	"send_put_request":                  true,
	"send_patch_request":                true,
	"send_delete_request":               true,
	"list_resource_names":               true,
}

// RegisterResourceModule makes a resource family reachable as one of a client's resources.
func RegisterResourceModule(resourceName string, factory ResourceFactory) error {
	if err := validateResourceName(resourceName); err != nil {
		return err
	}
	if factory == nil {
		return exerrors.NewUsageError("register_resource_module: " + resourceName +
			" has no factory - nothing would be reachable under that name")
	}
	resourceModulesMutex.Lock()
	defer resourceModulesMutex.Unlock()
	if _, found := resourceModules[resourceName]; found {
		// Any re-registration is refused: Go cannot compare function values, so a repeat is a clash.
		return exerrors.NewUsageError("register_resource_module: " + resourceName +
			" is already registered - one name cannot answer to two resources")
	}
	resourceModules[resourceName] = factory
	return nil
}

// validateResourceName refuses a name no lookup could ever reach.
func validateResourceName(resourceName string) error {
	if resourceName == "" {
		return exerrors.NewUsageError("register_resource_module: the resource name is blank")
	}
	if strings.HasPrefix(resourceName, "_") {
		return exerrors.NewUsageError("register_resource_module: " + resourceName +
			" is private, so it could never be reached")
	}
	for _, r := range resourceName {
		if !(r >= 'a' && r <= 'z') && !(r >= '0' && r <= '9') && r != '_' {
			return exerrors.NewUsageError("register_resource_module: " + resourceName +
				" is not lower snake_case, which is how a resource name is spelled")
		}
	}
	// Every family name starts with a letter, so a name that does not could never be one.
	if resourceName[0] < 'a' || resourceName[0] > 'z' {
		return exerrors.NewUsageError("register_resource_module: " + resourceName +
			" does not start with a lower-case letter, as every family name does")
	}
	if clientAttributeNames[resourceName] {
		return exerrors.NewUsageError("register_resource_module: " + resourceName +
			" is already a member of the client, so a resource of that name could never be reached")
	}
	return nil
}

// RegisteredResourceNames returns every registered resource name, sorted.
func RegisteredResourceNames() []string {
	resourceModulesMutex.RLock()
	defer resourceModulesMutex.RUnlock()
	return registeredNamesLocked()
}

// ResolveResourceModule returns the factory registered under a name.
func ResolveResourceModule(resourceName string) (ResourceFactory, error) {
	resourceModulesMutex.RLock()
	defer resourceModulesMutex.RUnlock()
	factory, found := resourceModules[resourceName]
	if !found {
		return nil, exerrors.NewAttributeError("no resource named " + resourceName +
			"; registered: " + strings.Join(registeredNamesLocked(), ", "))
	}
	return factory, nil
}

// registeredNamesLocked lists the names with the lock already held.
func registeredNamesLocked() []string {
	names := make([]string, 0, len(resourceModules))
	for name := range resourceModules {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// BoundResource is a resource family with a client already applied, constructed once.
type BoundResource struct {
	// Name is the registered resource name, e.g. "purchase_credit_notes".
	Name string
	// Value is the constructed family; assert it to its concrete type, or use the client's typed accessor.
	Value any
}

// NewBoundResource constructs the resource registered under a name over one requester.
func NewBoundResource(requester exhttp.Requester, resourceName string) (*BoundResource, error) {
	factory, err := ResolveResourceModule(resourceName)
	if err != nil {
		return nil, err
	}
	// Refused here, because a family bound to no requester would fail at its first call instead.
	if requester == nil {
		return nil, exerrors.UsageErrorf("cannot bind resource %q: it needs a requester to send with", resourceName)
	}
	return &BoundResource{Name: resourceName, Value: factory(requester)}, nil
}

// Resource returns one family bound to this client, by its registered name.
func (c *ExsitedClient) Resource(resourceName string) (*BoundResource, error) {
	return NewBoundResource(c, resourceName)
}
