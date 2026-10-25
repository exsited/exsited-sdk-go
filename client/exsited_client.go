package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	nethttp "net/http"
	neturl "net/url"
	"os"
	"reflect"
	"slices"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/exsited/exsited-sdk-go/v4/auth"
	"github.com/exsited/exsited-sdk-go/v4/config"
	"github.com/exsited/exsited-sdk-go/v4/dto"
	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
	"github.com/exsited/exsited-sdk-go/v4/redact"
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

// ExsitedClient sends requests; it satisfies http.Requester, which is what a resource module takes.
type ExsitedClient struct {
	config    *config.SdkConfig
	transport *exhttp.HTTPTransport
	tokens    *auth.TokenProvider
	log       func(string)

	// ownsTransport is true when no WithTransport was applied, so Close releases the transport.
	ownsTransport bool

	// closeMutex guards logFile, which Close read-modify-writes.
	closeMutex sync.Mutex
	logFile    *os.File // the LogFile sink, released by Close; nil when unset

	// closed is the lifetime latch Close sets and refuseIfClosed reads.
	closed atomic.Bool

	// lifetime is the window between a send being accepted and that send being finished.
	lifetime sync.RWMutex

	abaFilesOnce sync.Once
	abaFiles     *abafiles.AbaFiles

	accountEmailsOnce sync.Once
	accountEmails     *accountemails.AccountEmails

	accountGroupsOnce sync.Once
	accountGroups     *accountgroups.AccountGroups

	accountingCodesOnce sync.Once
	accountingCodes     *accountingcodes.AccountingCodes

	accountsOnce sync.Once
	accounts     *accounts.Accounts

	actionScriptsOnce sync.Once
	actionScripts     *actionscripts.ActionScripts

	additionalCodesOnce sync.Once
	additionalCodes     *additionalcodes.AdditionalCodes

	alertResponsibilitiesOnce sync.Once
	alertResponsibilities     *alertresponsibilities.AlertResponsibilities

	alertResponsibilityOnce sync.Once
	alertResponsibility     *alertresponsibility.AlertResponsibility

	alertsOnce sync.Once
	alerts     *alerts.Alerts

	applicationsOnce sync.Once
	applications     *applications.Applications

	approvalProfilesOnce sync.Once
	approvalProfiles     *approvalprofiles.ApprovalProfiles

	approvalsOnce sync.Once
	approvals     *approvals.Approvals

	approvalUserGroupsOnce sync.Once
	approvalUserGroups     *approvalusergroups.ApprovalUserGroups

	attributeAttachmentsOnce sync.Once
	attributeAttachments     *attributeattachments.AttributeAttachments

	billingPeriodsOnce sync.Once
	billingPeriods     *billingperiods.BillingPeriods

	bookingsOnce sync.Once
	bookings     *bookings.Bookings

	chargingPeriodsOnce sync.Once
	chargingPeriods     *chargingperiods.ChargingPeriods

	communicationProfilesOnce sync.Once
	communicationProfiles     *communicationprofiles.CommunicationProfiles

	communicationsOnce sync.Once
	communications     *communications.Communications

	countriesOnce sync.Once
	countries     *countries.Countries

	creditNoteApplicationsOnce sync.Once
	creditNoteApplications     *creditnoteapplications.CreditNoteApplications

	creditNotesOnce sync.Once
	creditNotes     *creditnotes.CreditNotes

	customAttributesOnce sync.Once
	customAttributes     *customattributes.CustomAttributes

	customComponentsOnce sync.Once
	customComponents     *customcomponents.CustomComponents

	customEventsOnce sync.Once
	customEvents     *customevents.CustomEvents

	customFeaturesOnce sync.Once
	customFeatures     *customfeatures.CustomFeatures

	customFieldsOnce sync.Once
	customFields     *customfields.CustomFields

	customFormsOnce sync.Once
	customForms     *customforms.CustomForms

	dayOfMonthOnce sync.Once
	dayOfMonth     *dayofmonth.DayOfMonth

	departmentsOnce sync.Once
	departments     *departments.Departments

	designationsOnce sync.Once
	designations     *designations.Designations

	documentTemplatesOnce sync.Once
	documentTemplates     *documenttemplates.DocumentTemplates

	emailTemplatesOnce sync.Once
	emailTemplates     *emailtemplates.EmailTemplates

	employeesOnce sync.Once
	employees     *employees.Employees

	eventsOnce sync.Once
	events     *events.Events

	expressOnce sync.Once
	express     *express.Express

	externalDatabasesOnce sync.Once
	externalDatabases     *externaldatabases.ExternalDatabases

	fulfillmentProfilesOnce sync.Once
	fulfillmentProfiles     *fulfillmentprofiles.FulfillmentProfiles

	fulfillmentTermsOnce sync.Once
	fulfillmentTerms     *fulfillmentterms.FulfillmentTerms

	giftCertificatesOnce sync.Once
	giftCertificates     *giftcertificates.GiftCertificates

	integrationConnectionsOnce sync.Once
	integrationConnections     *integrationconnections.IntegrationConnections

	inventoryAdjustmentsOnce sync.Once
	inventoryAdjustments     *inventoryadjustments.InventoryAdjustments

	invoicesOnce sync.Once
	invoices     *invoices.Invoices

	invoiceTermsOnce sync.Once
	invoiceTerms     *invoiceterms.InvoiceTerms

	itemFulfillmentsOnce sync.Once
	itemFulfillments     *itemfulfillments.ItemFulfillments

	itemReceiptsOnce sync.Once
	itemReceipts     *itemreceipts.ItemReceipts

	itemsOnce sync.Once
	items     *items.Items

	laboursOnce sync.Once
	labours     *labours.Labours

	milestonesOnce sync.Once
	milestones     *milestones.Milestones

	noteAttachmentsOnce sync.Once
	noteAttachments     *noteattachments.NoteAttachments

	noteGroupsOnce sync.Once
	noteGroups     *notegroups.NoteGroups

	ordersOnce sync.Once
	orders     *orders.Orders

	organisationOnce sync.Once
	organisation     *organisation.Organisation

	organisationsOnce sync.Once
	organisations     *organisations.Organisations

	paymentsOnce sync.Once
	payments     *payments.Payments

	portalsOnce sync.Once
	portals     *portals.Portals

	pricingPeriodsOnce sync.Once
	pricingPeriods     *pricingperiods.PricingPeriods

	proformaInvoicesOnce sync.Once
	proformaInvoices     *proformainvoices.ProformaInvoices

	purchaseCreditNoteApplicationsOnce sync.Once
	purchaseCreditNoteApplications     *purchasecreditnoteapplications.PurchaseCreditNoteApplications

	purchaseCreditNotesOnce sync.Once
	purchaseCreditNotes     *purchasecreditnotes.PurchaseCreditNotes

	purchaseExpressOnce sync.Once
	purchaseExpress     *purchaseexpress.PurchaseExpress

	purchaseInvoicesOnce sync.Once
	purchaseInvoices     *purchaseinvoices.PurchaseInvoices

	purchaseOrdersOnce sync.Once
	purchaseOrders     *purchaseorders.PurchaseOrders

	purchasePaymentsOnce sync.Once
	purchasePayments     *purchasepayments.PurchasePayments

	purchaseRefundsOnce sync.Once
	purchaseRefunds     *purchaserefunds.PurchaseRefunds

	quoteOrdersOnce sync.Once
	quoteOrders     *quoteorders.QuoteOrders

	receiveTermsOnce sync.Once
	receiveTerms     *receiveterms.ReceiveTerms

	reconciliationsOnce sync.Once
	reconciliations     *reconciliations.Reconciliations

	refundsOnce sync.Once
	refunds     *refunds.Refunds

	rentalAssetsOnce sync.Once
	rentalAssets     *rentalassets.RentalAssets

	reportsOnce sync.Once
	reports     *reports.Reports

	returnMerchandiseAuthorisationsOnce sync.Once
	returnMerchandiseAuthorisations     *returnmerchandiseauthorisations.ReturnMerchandiseAuthorisations

	returnVendorAuthorizationsOnce sync.Once
	returnVendorAuthorizations     *returnvendorauthorizations.ReturnVendorAuthorizations

	rolesOnce sync.Once
	roles     *roles.Roles

	salesRepsOnce sync.Once
	salesReps     *salesreps.SalesReps

	salutationsOnce sync.Once
	salutations     *salutations.Salutations

	settingsOnce sync.Once
	settings     *settings.Settings

	settingsWorkflowsOnce sync.Once
	settingsWorkflows     *settingsworkflows.SettingsWorkflows

	shippingPackagesOnce sync.Once
	shippingPackages     *shippingpackages.ShippingPackages

	shippingRulesOnce sync.Once
	shippingRules     *shippingrules.ShippingRules

	systemEventsOnce sync.Once
	systemEvents     *systemevents.SystemEvents

	tasksOnce sync.Once
	tasks     *tasks.Tasks

	teamsOnce sync.Once
	teams     *teams.Teams

	timePeriodsOnce sync.Once
	timePeriods     *timeperiods.TimePeriods

	timesheetsOnce sync.Once
	timesheets     *timesheets.Timesheets

	uomOnce sync.Once
	uom     *uom.Uom

	usagesOnce sync.Once
	usages     *usages.Usages

	webhooksOnce sync.Once
	webhooks     *webhooks.Webhooks

	workflowsOnce sync.Once
	workflows     *workflows.Workflows

	zonesOnce sync.Once
	zones     *zones.Zones
}

const (
	// LoggerName prefixes each log-file line, followed by the escaped client name.
	LoggerName = "exsited"

	apiVersionParam = "api_version"

	// authorizationHeader is the canonical spelling the bearer token is always sent under.
	authorizationHeader = "Authorization"

	// noResponseStatus stands where a status would go; text, so it cannot look like a real code.
	noResponseStatus = "no response"

	// The record levels. LogLevel gates which of them the log hook receives.
	levelDebug   = "DEBUG"
	levelInfo    = "INFO"
	levelWarning = "WARNING"
	levelError   = "ERROR"
)

// logLevelRanks orders the five levels so a configured LogLevel drops every record below it.
var logLevelRanks = map[string]int{
	levelDebug: 10, levelInfo: 20, levelWarning: 30, levelError: 40, "CRITICAL": 50,
}

var logRecordPrefixes = map[string]string{
	levelDebug: "DEBUG ", levelInfo: "INFO ", levelWarning: "WARN ", levelError: "ERROR ",
}

func logAllows(sdkConfig config.SdkConfig, level string) bool {
	return logLevelRanks[level] >= logLevelRanks[sdkConfig.LogLevel]
}

type clientOptions struct {
	transport     *exhttp.HTTPTransport
	tokenStore    auth.TokenStore
	tokenProvider *auth.TokenProvider
	retryPolicy   *exhttp.RetryPolicy
	log           func(string)
}

// ClientOption configures a client at construction: the injection points, mainly for testing.
type ClientOption func(*clientOptions)

// WithTransport injects a transport in place of the one the config would build.
func WithTransport(t *exhttp.HTTPTransport) ClientOption {
	return func(applied *clientOptions) { applied.transport = t }
}

func WithTokenStore(s auth.TokenStore) ClientOption {
	return func(applied *clientOptions) { applied.tokenStore = s }
}

// WithTokenProvider injects a built token provider, so several clients can share one token cache.
func WithTokenProvider(p *auth.TokenProvider) ClientOption {
	return func(applied *clientOptions) { applied.tokenProvider = p }
}

// WithRetryPolicy injects a retry policy (built with NewRetryPolicy) into the default transport.
func WithRetryPolicy(p *exhttp.RetryPolicy) ClientOption {
	return func(applied *clientOptions) { applied.retryPolicy = p }
}

// WithLogger installs a plain log hook: a func(string) called with one line per exchange.
func WithLogger(f func(string)) ClientOption { return func(applied *clientOptions) { applied.log = f } }

func NewClient(given config.SdkConfig, options ...ClientOption) (*ExsitedClient, error) {
	// Validated raw, then normalised once, the response format resolved: the only place EXSITED_RESPONSE_FORMAT is read.
	resolved, err := config.NewSdkConfig(given)
	if err != nil {
		return nil, err
	}
	applied := &clientOptions{}
	for _, option := range options {
		option(applied)
	}
	if applied.transport != nil && applied.retryPolicy != nil {
		return nil, exerrors.NewUsageError("pass WithTransport or WithRetryPolicy, not both - an injected " +
			"transport already carries its own policy, so this one would be discarded unused; build it " +
			"into the transport instead, as NewHTTPTransport(config) with maxRetries and the " +
			"retry fields set on that config")
	}

	logFn, logFile, err := attachLogFile(*resolved, applied.log)
	if err != nil {
		return nil, err
	}
	closeOnError := func() {
		if logFile != nil {
			_ = logFile.Close()
		}
	}

	warnAboutInsecureTransport(*resolved, logFn)

	transport := applied.transport

	ownsTransport := transport == nil
	if transport == nil {
		transport, err = exhttp.NewHTTPTransport(*resolved,
			exhttp.WithProgress(progressFunc(*resolved, logFn)),
			exhttp.WithRetryPolicy(applied.retryPolicy))
		if err != nil {
			closeOnError()
			return nil, err
		}
	}
	tokenProvider := applied.tokenProvider
	if tokenProvider == nil {
		store := applied.tokenStore
		if store == nil {
			store = buildDefaultTokenStore(*resolved)
		}
		tokenProvider = auth.NewTokenProvider(resolved, store, transport)
	}
	return &ExsitedClient{config: resolved, transport: transport, ownsTransport: ownsTransport,
		tokens: tokenProvider, log: logFn, logFile: logFile}, nil
}

// NewExsitedClient is an alias of NewClient.
func NewExsitedClient(given config.SdkConfig, options ...ClientOption) (*ExsitedClient, error) {
	return NewClient(given, options...)
}

// buildDefaultTokenStore is the store AuthConfig asks for: this process only, or the shared cache file.
func buildDefaultTokenStore(sdkConfig config.SdkConfig) auth.TokenStore {
	if sdkConfig.Auth.InMemoryToken {
		return auth.NewMemoryTokenStore()
	}
	path := sdkConfig.Auth.TokenCachePath
	if path == "" {
		path = auth.DefaultTokenCachePath()
	}
	return auth.NewFileTokenStore(path)
}

// loggerSafe flattens a client name into one log-name segment, so no instance nests under another.
func loggerSafe(clientName string) string {
	var escaped strings.Builder
	for index := 0; index < len(clientName); index++ {
		c := clientName[index]
		switch {
		case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9',
			c == '-', c == '_', c == '~':
			escaped.WriteByte(c)
		default:
			fmt.Fprintf(&escaped, "%%%02X", c)
		}
	}
	return escaped.String()
}

// attachLogFile opens config.LogFile for append and returns a hook that also writes timestamped lines there.
func attachLogFile(sdkConfig config.SdkConfig, base func(string)) (func(string), *os.File, error) {
	if sdkConfig.LogFile == "" {
		return base, nil, nil
	}
	// 0600, like the token cache: the exchange lines logged here carry record ids in their URLs.
	f, err := os.OpenFile(sdkConfig.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		// The path and the OS error go to CauseText, not the message callers are asked to retain.
		refusal := exerrors.NewConfigError(fmt.Sprintf(
			"logFile: %s could not be opened for appending", redact.DescribePath(sdkConfig.LogFile)))
		refusal.Cause = err
		refusal.CauseText = fmt.Sprintf("%s: %v", sdkConfig.LogFile, err)
		return nil, nil, refusal
	}
	name := loggerSafe(sdkConfig.Auth.ClientName)
	var writing sync.Mutex
	sink := func(line string) {
		// One Fprintf per record, under a lock, so two concurrent sends cannot interleave a row.
		writing.Lock()
		fmt.Fprintf(f, "%s %s.%s %s\n", time.Now().Format("2006-01-02 15:04:05"), LoggerName, name, line)
		writing.Unlock()
		if base != nil {
			base(line)
		}
	}
	return sink, f, nil
}

// emittedTransportWarnings dedupes the stderr copy of each warning per process; the log hook gets all.
var (
	emittedTransportWarningsMu sync.Mutex
	emittedTransportWarnings   = map[string]bool{}
)

// warnAboutInsecureTransport warns when TLS verification is off or the base URL is http:// to a remote host.
func warnAboutInsecureTransport(sdkConfig config.SdkConfig, log func(string)) {
	emit := func(message string) {
		w := exerrors.NewInsecureTransportWarning(message)
		if log != nil && logAllows(sdkConfig, levelWarning) {
			log("WARN " + w.Error())
		}
		emittedTransportWarningsMu.Lock()
		seen := emittedTransportWarnings[message]
		emittedTransportWarnings[message] = true
		emittedTransportWarningsMu.Unlock()
		if !seen {
			fmt.Fprintln(warningWriter(sdkConfig), "InsecureTransportWarning: "+w.Error())
		}
	}
	if !sdkConfig.VerifySslEnabled() {
		emit("TLS verification is disabled (verifySsl=false); responses cannot be trusted")
	}
	if warning := config.PlaintextTransportWarning(sdkConfig.Auth.BaseURL); warning != "" {
		emit(warning)
	}
}

// printWriter is where the printed exchange goes: DiagnosticWriter when set, else stdout.
func printWriter(sdkConfig config.SdkConfig) io.Writer {
	if sdkConfig.DiagnosticWriter != nil {
		return sdkConfig.DiagnosticWriter
	}
	return os.Stdout
}

// warningWriter is where a transport warning goes: DiagnosticWriter when set, else stderr.
func warningWriter(sdkConfig config.SdkConfig) io.Writer {
	if sdkConfig.DiagnosticWriter != nil {
		return sdkConfig.DiagnosticWriter
	}
	return os.Stderr
}

// Config returns a deep copy of the configuration this client was built with.
func (c *ExsitedClient) Config() config.SdkConfig { return c.config.DeepCopy() }

func (c *ExsitedClient) Name() string { return c.config.Auth.ClientName }

// AbaFiles is built and cached on first use.
func (c *ExsitedClient) AbaFiles() *abafiles.AbaFiles {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.abaFilesOnce.Do(func() {
		bound, _ := NewBoundResource(c, abafiles.ResourceName)
		c.abaFiles, _ = bound.Value.(*abafiles.AbaFiles)
	})
	return c.abaFiles
}

// AccountEmails is built and cached on first use.
func (c *ExsitedClient) AccountEmails() *accountemails.AccountEmails {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.accountEmailsOnce.Do(func() {
		bound, _ := NewBoundResource(c, accountemails.ResourceName)
		c.accountEmails, _ = bound.Value.(*accountemails.AccountEmails)
	})
	return c.accountEmails
}

// AccountGroups is built and cached on first use.
func (c *ExsitedClient) AccountGroups() *accountgroups.AccountGroups {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.accountGroupsOnce.Do(func() {
		bound, _ := NewBoundResource(c, accountgroups.ResourceName)
		c.accountGroups, _ = bound.Value.(*accountgroups.AccountGroups)
	})
	return c.accountGroups
}

// AccountingCodes is built and cached on first use.
func (c *ExsitedClient) AccountingCodes() *accountingcodes.AccountingCodes {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.accountingCodesOnce.Do(func() {
		bound, _ := NewBoundResource(c, accountingcodes.ResourceName)
		c.accountingCodes, _ = bound.Value.(*accountingcodes.AccountingCodes)
	})
	return c.accountingCodes
}

// Accounts is built and cached on first use.
func (c *ExsitedClient) Accounts() *accounts.Accounts {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.accountsOnce.Do(func() {
		bound, _ := NewBoundResource(c, accounts.ResourceName)
		c.accounts, _ = bound.Value.(*accounts.Accounts)
	})
	return c.accounts
}

// ActionScripts is built and cached on first use.
func (c *ExsitedClient) ActionScripts() *actionscripts.ActionScripts {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.actionScriptsOnce.Do(func() {
		bound, _ := NewBoundResource(c, actionscripts.ResourceName)
		c.actionScripts, _ = bound.Value.(*actionscripts.ActionScripts)
	})
	return c.actionScripts
}

// AdditionalCodes is built and cached on first use.
func (c *ExsitedClient) AdditionalCodes() *additionalcodes.AdditionalCodes {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.additionalCodesOnce.Do(func() {
		bound, _ := NewBoundResource(c, additionalcodes.ResourceName)
		c.additionalCodes, _ = bound.Value.(*additionalcodes.AdditionalCodes)
	})
	return c.additionalCodes
}

// AlertResponsibilities is built and cached on first use.
func (c *ExsitedClient) AlertResponsibilities() *alertresponsibilities.AlertResponsibilities {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.alertResponsibilitiesOnce.Do(func() {
		bound, _ := NewBoundResource(c, alertresponsibilities.ResourceName)
		c.alertResponsibilities, _ = bound.Value.(*alertresponsibilities.AlertResponsibilities)
	})
	return c.alertResponsibilities
}

// AlertResponsibility is built and cached on first use.
func (c *ExsitedClient) AlertResponsibility() *alertresponsibility.AlertResponsibility {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.alertResponsibilityOnce.Do(func() {
		bound, _ := NewBoundResource(c, alertresponsibility.ResourceName)
		c.alertResponsibility, _ = bound.Value.(*alertresponsibility.AlertResponsibility)
	})
	return c.alertResponsibility
}

// Alerts is built and cached on first use.
func (c *ExsitedClient) Alerts() *alerts.Alerts {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.alertsOnce.Do(func() {
		bound, _ := NewBoundResource(c, alerts.ResourceName)
		c.alerts, _ = bound.Value.(*alerts.Alerts)
	})
	return c.alerts
}

// Applications is built and cached on first use.
func (c *ExsitedClient) Applications() *applications.Applications {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.applicationsOnce.Do(func() {
		bound, _ := NewBoundResource(c, applications.ResourceName)
		c.applications, _ = bound.Value.(*applications.Applications)
	})
	return c.applications
}

// ApprovalProfiles is built and cached on first use.
func (c *ExsitedClient) ApprovalProfiles() *approvalprofiles.ApprovalProfiles {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.approvalProfilesOnce.Do(func() {
		bound, _ := NewBoundResource(c, approvalprofiles.ResourceName)
		c.approvalProfiles, _ = bound.Value.(*approvalprofiles.ApprovalProfiles)
	})
	return c.approvalProfiles
}

// Approvals is built and cached on first use.
func (c *ExsitedClient) Approvals() *approvals.Approvals {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.approvalsOnce.Do(func() {
		bound, _ := NewBoundResource(c, approvals.ResourceName)
		c.approvals, _ = bound.Value.(*approvals.Approvals)
	})
	return c.approvals
}

// ApprovalUserGroups is built and cached on first use.
func (c *ExsitedClient) ApprovalUserGroups() *approvalusergroups.ApprovalUserGroups {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.approvalUserGroupsOnce.Do(func() {
		bound, _ := NewBoundResource(c, approvalusergroups.ResourceName)
		c.approvalUserGroups, _ = bound.Value.(*approvalusergroups.ApprovalUserGroups)
	})
	return c.approvalUserGroups
}

// AttributeAttachments is built and cached on first use.
func (c *ExsitedClient) AttributeAttachments() *attributeattachments.AttributeAttachments {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.attributeAttachmentsOnce.Do(func() {
		bound, _ := NewBoundResource(c, attributeattachments.ResourceName)
		c.attributeAttachments, _ = bound.Value.(*attributeattachments.AttributeAttachments)
	})
	return c.attributeAttachments
}

// BillingPeriods is built and cached on first use.
func (c *ExsitedClient) BillingPeriods() *billingperiods.BillingPeriods {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.billingPeriodsOnce.Do(func() {
		bound, _ := NewBoundResource(c, billingperiods.ResourceName)
		c.billingPeriods, _ = bound.Value.(*billingperiods.BillingPeriods)
	})
	return c.billingPeriods
}

// Bookings is built and cached on first use.
func (c *ExsitedClient) Bookings() *bookings.Bookings {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.bookingsOnce.Do(func() {
		bound, _ := NewBoundResource(c, bookings.ResourceName)
		c.bookings, _ = bound.Value.(*bookings.Bookings)
	})
	return c.bookings
}

// ChargingPeriods is built and cached on first use.
func (c *ExsitedClient) ChargingPeriods() *chargingperiods.ChargingPeriods {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.chargingPeriodsOnce.Do(func() {
		bound, _ := NewBoundResource(c, chargingperiods.ResourceName)
		c.chargingPeriods, _ = bound.Value.(*chargingperiods.ChargingPeriods)
	})
	return c.chargingPeriods
}

// CommunicationProfiles is built and cached on first use.
func (c *ExsitedClient) CommunicationProfiles() *communicationprofiles.CommunicationProfiles {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.communicationProfilesOnce.Do(func() {
		bound, _ := NewBoundResource(c, communicationprofiles.ResourceName)
		c.communicationProfiles, _ = bound.Value.(*communicationprofiles.CommunicationProfiles)
	})
	return c.communicationProfiles
}

// Communications is built and cached on first use.
func (c *ExsitedClient) Communications() *communications.Communications {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.communicationsOnce.Do(func() {
		bound, _ := NewBoundResource(c, communications.ResourceName)
		c.communications, _ = bound.Value.(*communications.Communications)
	})
	return c.communications
}

// Countries is built and cached on first use.
func (c *ExsitedClient) Countries() *countries.Countries {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.countriesOnce.Do(func() {
		bound, _ := NewBoundResource(c, countries.ResourceName)
		c.countries, _ = bound.Value.(*countries.Countries)
	})
	return c.countries
}

// CreditNoteApplications is built and cached on first use.
func (c *ExsitedClient) CreditNoteApplications() *creditnoteapplications.CreditNoteApplications {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.creditNoteApplicationsOnce.Do(func() {
		bound, _ := NewBoundResource(c, creditnoteapplications.ResourceName)
		c.creditNoteApplications, _ = bound.Value.(*creditnoteapplications.CreditNoteApplications)
	})
	return c.creditNoteApplications
}

// CreditNotes is built and cached on first use.
func (c *ExsitedClient) CreditNotes() *creditnotes.CreditNotes {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.creditNotesOnce.Do(func() {
		bound, _ := NewBoundResource(c, creditnotes.ResourceName)
		c.creditNotes, _ = bound.Value.(*creditnotes.CreditNotes)
	})
	return c.creditNotes
}

// CustomAttributes is built and cached on first use.
func (c *ExsitedClient) CustomAttributes() *customattributes.CustomAttributes {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.customAttributesOnce.Do(func() {
		bound, _ := NewBoundResource(c, customattributes.ResourceName)
		c.customAttributes, _ = bound.Value.(*customattributes.CustomAttributes)
	})
	return c.customAttributes
}

// CustomComponents is built and cached on first use.
func (c *ExsitedClient) CustomComponents() *customcomponents.CustomComponents {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.customComponentsOnce.Do(func() {
		bound, _ := NewBoundResource(c, customcomponents.ResourceName)
		c.customComponents, _ = bound.Value.(*customcomponents.CustomComponents)
	})
	return c.customComponents
}

// CustomEvents is built and cached on first use.
func (c *ExsitedClient) CustomEvents() *customevents.CustomEvents {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.customEventsOnce.Do(func() {
		bound, _ := NewBoundResource(c, customevents.ResourceName)
		c.customEvents, _ = bound.Value.(*customevents.CustomEvents)
	})
	return c.customEvents
}

// CustomFeatures is built and cached on first use.
func (c *ExsitedClient) CustomFeatures() *customfeatures.CustomFeatures {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.customFeaturesOnce.Do(func() {
		bound, _ := NewBoundResource(c, customfeatures.ResourceName)
		c.customFeatures, _ = bound.Value.(*customfeatures.CustomFeatures)
	})
	return c.customFeatures
}

// CustomFields is built and cached on first use.
func (c *ExsitedClient) CustomFields() *customfields.CustomFields {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.customFieldsOnce.Do(func() {
		bound, _ := NewBoundResource(c, customfields.ResourceName)
		c.customFields, _ = bound.Value.(*customfields.CustomFields)
	})
	return c.customFields
}

// CustomForms is built and cached on first use.
func (c *ExsitedClient) CustomForms() *customforms.CustomForms {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.customFormsOnce.Do(func() {
		bound, _ := NewBoundResource(c, customforms.ResourceName)
		c.customForms, _ = bound.Value.(*customforms.CustomForms)
	})
	return c.customForms
}

// DayOfMonth is built and cached on first use.
func (c *ExsitedClient) DayOfMonth() *dayofmonth.DayOfMonth {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.dayOfMonthOnce.Do(func() {
		bound, _ := NewBoundResource(c, dayofmonth.ResourceName)
		c.dayOfMonth, _ = bound.Value.(*dayofmonth.DayOfMonth)
	})
	return c.dayOfMonth
}

// Departments is built and cached on first use.
func (c *ExsitedClient) Departments() *departments.Departments {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.departmentsOnce.Do(func() {
		bound, _ := NewBoundResource(c, departments.ResourceName)
		c.departments, _ = bound.Value.(*departments.Departments)
	})
	return c.departments
}

// Designations is built and cached on first use.
func (c *ExsitedClient) Designations() *designations.Designations {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.designationsOnce.Do(func() {
		bound, _ := NewBoundResource(c, designations.ResourceName)
		c.designations, _ = bound.Value.(*designations.Designations)
	})
	return c.designations
}

// DocumentTemplates is built and cached on first use.
func (c *ExsitedClient) DocumentTemplates() *documenttemplates.DocumentTemplates {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.documentTemplatesOnce.Do(func() {
		bound, _ := NewBoundResource(c, documenttemplates.ResourceName)
		c.documentTemplates, _ = bound.Value.(*documenttemplates.DocumentTemplates)
	})
	return c.documentTemplates
}

// EmailTemplates is built and cached on first use.
func (c *ExsitedClient) EmailTemplates() *emailtemplates.EmailTemplates {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.emailTemplatesOnce.Do(func() {
		bound, _ := NewBoundResource(c, emailtemplates.ResourceName)
		c.emailTemplates, _ = bound.Value.(*emailtemplates.EmailTemplates)
	})
	return c.emailTemplates
}

// Employees is built and cached on first use.
func (c *ExsitedClient) Employees() *employees.Employees {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.employeesOnce.Do(func() {
		bound, _ := NewBoundResource(c, employees.ResourceName)
		c.employees, _ = bound.Value.(*employees.Employees)
	})
	return c.employees
}

// Events is built and cached on first use.
func (c *ExsitedClient) Events() *events.Events {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.eventsOnce.Do(func() {
		bound, _ := NewBoundResource(c, events.ResourceName)
		c.events, _ = bound.Value.(*events.Events)
	})
	return c.events
}

// Express is built and cached on first use.
func (c *ExsitedClient) Express() *express.Express {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.expressOnce.Do(func() {
		bound, _ := NewBoundResource(c, express.ResourceName)
		c.express, _ = bound.Value.(*express.Express)
	})
	return c.express
}

// ExternalDatabases is built and cached on first use.
func (c *ExsitedClient) ExternalDatabases() *externaldatabases.ExternalDatabases {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.externalDatabasesOnce.Do(func() {
		bound, _ := NewBoundResource(c, externaldatabases.ResourceName)
		c.externalDatabases, _ = bound.Value.(*externaldatabases.ExternalDatabases)
	})
	return c.externalDatabases
}

// FulfillmentProfiles is built and cached on first use.
func (c *ExsitedClient) FulfillmentProfiles() *fulfillmentprofiles.FulfillmentProfiles {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.fulfillmentProfilesOnce.Do(func() {
		bound, _ := NewBoundResource(c, fulfillmentprofiles.ResourceName)
		c.fulfillmentProfiles, _ = bound.Value.(*fulfillmentprofiles.FulfillmentProfiles)
	})
	return c.fulfillmentProfiles
}

// FulfillmentTerms is built and cached on first use.
func (c *ExsitedClient) FulfillmentTerms() *fulfillmentterms.FulfillmentTerms {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.fulfillmentTermsOnce.Do(func() {
		bound, _ := NewBoundResource(c, fulfillmentterms.ResourceName)
		c.fulfillmentTerms, _ = bound.Value.(*fulfillmentterms.FulfillmentTerms)
	})
	return c.fulfillmentTerms
}

// GiftCertificates is built and cached on first use.
func (c *ExsitedClient) GiftCertificates() *giftcertificates.GiftCertificates {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.giftCertificatesOnce.Do(func() {
		bound, _ := NewBoundResource(c, giftcertificates.ResourceName)
		c.giftCertificates, _ = bound.Value.(*giftcertificates.GiftCertificates)
	})
	return c.giftCertificates
}

// IntegrationConnections is built and cached on first use.
func (c *ExsitedClient) IntegrationConnections() *integrationconnections.IntegrationConnections {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.integrationConnectionsOnce.Do(func() {
		bound, _ := NewBoundResource(c, integrationconnections.ResourceName)
		c.integrationConnections, _ = bound.Value.(*integrationconnections.IntegrationConnections)
	})
	return c.integrationConnections
}

// InventoryAdjustments is built and cached on first use.
func (c *ExsitedClient) InventoryAdjustments() *inventoryadjustments.InventoryAdjustments {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.inventoryAdjustmentsOnce.Do(func() {
		bound, _ := NewBoundResource(c, inventoryadjustments.ResourceName)
		c.inventoryAdjustments, _ = bound.Value.(*inventoryadjustments.InventoryAdjustments)
	})
	return c.inventoryAdjustments
}

// Invoices is built and cached on first use.
func (c *ExsitedClient) Invoices() *invoices.Invoices {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.invoicesOnce.Do(func() {
		bound, _ := NewBoundResource(c, invoices.ResourceName)
		c.invoices, _ = bound.Value.(*invoices.Invoices)
	})
	return c.invoices
}

// InvoiceTerms is built and cached on first use.
func (c *ExsitedClient) InvoiceTerms() *invoiceterms.InvoiceTerms {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.invoiceTermsOnce.Do(func() {
		bound, _ := NewBoundResource(c, invoiceterms.ResourceName)
		c.invoiceTerms, _ = bound.Value.(*invoiceterms.InvoiceTerms)
	})
	return c.invoiceTerms
}

// ItemFulfillments is built and cached on first use.
func (c *ExsitedClient) ItemFulfillments() *itemfulfillments.ItemFulfillments {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.itemFulfillmentsOnce.Do(func() {
		bound, _ := NewBoundResource(c, itemfulfillments.ResourceName)
		c.itemFulfillments, _ = bound.Value.(*itemfulfillments.ItemFulfillments)
	})
	return c.itemFulfillments
}

// ItemReceipts is built and cached on first use.
func (c *ExsitedClient) ItemReceipts() *itemreceipts.ItemReceipts {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.itemReceiptsOnce.Do(func() {
		bound, _ := NewBoundResource(c, itemreceipts.ResourceName)
		c.itemReceipts, _ = bound.Value.(*itemreceipts.ItemReceipts)
	})
	return c.itemReceipts
}

// Items is built and cached on first use.
func (c *ExsitedClient) Items() *items.Items {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.itemsOnce.Do(func() {
		bound, _ := NewBoundResource(c, items.ResourceName)
		c.items, _ = bound.Value.(*items.Items)
	})
	return c.items
}

// Labours is built and cached on first use.
func (c *ExsitedClient) Labours() *labours.Labours {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.laboursOnce.Do(func() {
		bound, _ := NewBoundResource(c, labours.ResourceName)
		c.labours, _ = bound.Value.(*labours.Labours)
	})
	return c.labours
}

// Milestones is built and cached on first use.
func (c *ExsitedClient) Milestones() *milestones.Milestones {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.milestonesOnce.Do(func() {
		bound, _ := NewBoundResource(c, milestones.ResourceName)
		c.milestones, _ = bound.Value.(*milestones.Milestones)
	})
	return c.milestones
}

// NoteAttachments is built and cached on first use.
func (c *ExsitedClient) NoteAttachments() *noteattachments.NoteAttachments {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.noteAttachmentsOnce.Do(func() {
		bound, _ := NewBoundResource(c, noteattachments.ResourceName)
		c.noteAttachments, _ = bound.Value.(*noteattachments.NoteAttachments)
	})
	return c.noteAttachments
}

// NoteGroups is built and cached on first use.
func (c *ExsitedClient) NoteGroups() *notegroups.NoteGroups {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.noteGroupsOnce.Do(func() {
		bound, _ := NewBoundResource(c, notegroups.ResourceName)
		c.noteGroups, _ = bound.Value.(*notegroups.NoteGroups)
	})
	return c.noteGroups
}

// Orders is built and cached on first use.
func (c *ExsitedClient) Orders() *orders.Orders {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.ordersOnce.Do(func() {
		bound, _ := NewBoundResource(c, orders.ResourceName)
		c.orders, _ = bound.Value.(*orders.Orders)
	})
	return c.orders
}

// Organisation is built and cached on first use.
func (c *ExsitedClient) Organisation() *organisation.Organisation {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.organisationOnce.Do(func() {
		bound, _ := NewBoundResource(c, organisation.ResourceName)
		c.organisation, _ = bound.Value.(*organisation.Organisation)
	})
	return c.organisation
}

// Organisations is built and cached on first use.
func (c *ExsitedClient) Organisations() *organisations.Organisations {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.organisationsOnce.Do(func() {
		bound, _ := NewBoundResource(c, organisations.ResourceName)
		c.organisations, _ = bound.Value.(*organisations.Organisations)
	})
	return c.organisations
}

// Payments is built and cached on first use.
func (c *ExsitedClient) Payments() *payments.Payments {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.paymentsOnce.Do(func() {
		bound, _ := NewBoundResource(c, payments.ResourceName)
		c.payments, _ = bound.Value.(*payments.Payments)
	})
	return c.payments
}

// Portals is built and cached on first use.
func (c *ExsitedClient) Portals() *portals.Portals {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.portalsOnce.Do(func() {
		bound, _ := NewBoundResource(c, portals.ResourceName)
		c.portals, _ = bound.Value.(*portals.Portals)
	})
	return c.portals
}

// PricingPeriods is built and cached on first use.
func (c *ExsitedClient) PricingPeriods() *pricingperiods.PricingPeriods {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.pricingPeriodsOnce.Do(func() {
		bound, _ := NewBoundResource(c, pricingperiods.ResourceName)
		c.pricingPeriods, _ = bound.Value.(*pricingperiods.PricingPeriods)
	})
	return c.pricingPeriods
}

// ProformaInvoices is built and cached on first use.
func (c *ExsitedClient) ProformaInvoices() *proformainvoices.ProformaInvoices {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.proformaInvoicesOnce.Do(func() {
		bound, _ := NewBoundResource(c, proformainvoices.ResourceName)
		c.proformaInvoices, _ = bound.Value.(*proformainvoices.ProformaInvoices)
	})
	return c.proformaInvoices
}

// PurchaseCreditNoteApplications is built and cached on first use.
func (c *ExsitedClient) PurchaseCreditNoteApplications() *purchasecreditnoteapplications.PurchaseCreditNoteApplications {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.purchaseCreditNoteApplicationsOnce.Do(func() {
		bound, _ := NewBoundResource(c, purchasecreditnoteapplications.ResourceName)
		c.purchaseCreditNoteApplications, _ = bound.Value.(*purchasecreditnoteapplications.PurchaseCreditNoteApplications)
	})
	return c.purchaseCreditNoteApplications
}

// PurchaseCreditNotes is built and cached on first use.
func (c *ExsitedClient) PurchaseCreditNotes() *purchasecreditnotes.PurchaseCreditNotes {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.purchaseCreditNotesOnce.Do(func() {
		bound, _ := NewBoundResource(c, purchasecreditnotes.ResourceName)
		c.purchaseCreditNotes, _ = bound.Value.(*purchasecreditnotes.PurchaseCreditNotes)
	})
	return c.purchaseCreditNotes
}

// PurchaseExpress is built and cached on first use.
func (c *ExsitedClient) PurchaseExpress() *purchaseexpress.PurchaseExpress {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.purchaseExpressOnce.Do(func() {
		bound, _ := NewBoundResource(c, purchaseexpress.ResourceName)
		c.purchaseExpress, _ = bound.Value.(*purchaseexpress.PurchaseExpress)
	})
	return c.purchaseExpress
}

// PurchaseInvoices is built and cached on first use.
func (c *ExsitedClient) PurchaseInvoices() *purchaseinvoices.PurchaseInvoices {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.purchaseInvoicesOnce.Do(func() {
		bound, _ := NewBoundResource(c, purchaseinvoices.ResourceName)
		c.purchaseInvoices, _ = bound.Value.(*purchaseinvoices.PurchaseInvoices)
	})
	return c.purchaseInvoices
}

// PurchaseOrders is built and cached on first use.
func (c *ExsitedClient) PurchaseOrders() *purchaseorders.PurchaseOrders {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.purchaseOrdersOnce.Do(func() {
		bound, _ := NewBoundResource(c, purchaseorders.ResourceName)
		c.purchaseOrders, _ = bound.Value.(*purchaseorders.PurchaseOrders)
	})
	return c.purchaseOrders
}

// PurchasePayments is built and cached on first use.
func (c *ExsitedClient) PurchasePayments() *purchasepayments.PurchasePayments {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.purchasePaymentsOnce.Do(func() {
		bound, _ := NewBoundResource(c, purchasepayments.ResourceName)
		c.purchasePayments, _ = bound.Value.(*purchasepayments.PurchasePayments)
	})
	return c.purchasePayments
}

// PurchaseRefunds is built and cached on first use.
func (c *ExsitedClient) PurchaseRefunds() *purchaserefunds.PurchaseRefunds {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.purchaseRefundsOnce.Do(func() {
		bound, _ := NewBoundResource(c, purchaserefunds.ResourceName)
		c.purchaseRefunds, _ = bound.Value.(*purchaserefunds.PurchaseRefunds)
	})
	return c.purchaseRefunds
}

// QuoteOrders is built and cached on first use.
func (c *ExsitedClient) QuoteOrders() *quoteorders.QuoteOrders {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.quoteOrdersOnce.Do(func() {
		bound, _ := NewBoundResource(c, quoteorders.ResourceName)
		c.quoteOrders, _ = bound.Value.(*quoteorders.QuoteOrders)
	})
	return c.quoteOrders
}

// ReceiveTerms is built and cached on first use.
func (c *ExsitedClient) ReceiveTerms() *receiveterms.ReceiveTerms {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.receiveTermsOnce.Do(func() {
		bound, _ := NewBoundResource(c, receiveterms.ResourceName)
		c.receiveTerms, _ = bound.Value.(*receiveterms.ReceiveTerms)
	})
	return c.receiveTerms
}

// Reconciliations is built and cached on first use.
func (c *ExsitedClient) Reconciliations() *reconciliations.Reconciliations {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.reconciliationsOnce.Do(func() {
		bound, _ := NewBoundResource(c, reconciliations.ResourceName)
		c.reconciliations, _ = bound.Value.(*reconciliations.Reconciliations)
	})
	return c.reconciliations
}

// Refunds is built and cached on first use.
func (c *ExsitedClient) Refunds() *refunds.Refunds {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.refundsOnce.Do(func() {
		bound, _ := NewBoundResource(c, refunds.ResourceName)
		c.refunds, _ = bound.Value.(*refunds.Refunds)
	})
	return c.refunds
}

// RentalAssets is built and cached on first use.
func (c *ExsitedClient) RentalAssets() *rentalassets.RentalAssets {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.rentalAssetsOnce.Do(func() {
		bound, _ := NewBoundResource(c, rentalassets.ResourceName)
		c.rentalAssets, _ = bound.Value.(*rentalassets.RentalAssets)
	})
	return c.rentalAssets
}

// Reports is built and cached on first use.
func (c *ExsitedClient) Reports() *reports.Reports {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.reportsOnce.Do(func() {
		bound, _ := NewBoundResource(c, reports.ResourceName)
		c.reports, _ = bound.Value.(*reports.Reports)
	})
	return c.reports
}

// ReturnMerchandiseAuthorisations is built and cached on first use.
func (c *ExsitedClient) ReturnMerchandiseAuthorisations() *returnmerchandiseauthorisations.ReturnMerchandiseAuthorisations {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.returnMerchandiseAuthorisationsOnce.Do(func() {
		bound, _ := NewBoundResource(c, returnmerchandiseauthorisations.ResourceName)
		c.returnMerchandiseAuthorisations, _ = bound.Value.(*returnmerchandiseauthorisations.ReturnMerchandiseAuthorisations)
	})
	return c.returnMerchandiseAuthorisations
}

// ReturnVendorAuthorizations is built and cached on first use.
func (c *ExsitedClient) ReturnVendorAuthorizations() *returnvendorauthorizations.ReturnVendorAuthorizations {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.returnVendorAuthorizationsOnce.Do(func() {
		bound, _ := NewBoundResource(c, returnvendorauthorizations.ResourceName)
		c.returnVendorAuthorizations, _ = bound.Value.(*returnvendorauthorizations.ReturnVendorAuthorizations)
	})
	return c.returnVendorAuthorizations
}

// Roles is built and cached on first use.
func (c *ExsitedClient) Roles() *roles.Roles {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.rolesOnce.Do(func() {
		bound, _ := NewBoundResource(c, roles.ResourceName)
		c.roles, _ = bound.Value.(*roles.Roles)
	})
	return c.roles
}

// SalesReps is built and cached on first use.
func (c *ExsitedClient) SalesReps() *salesreps.SalesReps {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.salesRepsOnce.Do(func() {
		bound, _ := NewBoundResource(c, salesreps.ResourceName)
		c.salesReps, _ = bound.Value.(*salesreps.SalesReps)
	})
	return c.salesReps
}

// Salutations is built and cached on first use.
func (c *ExsitedClient) Salutations() *salutations.Salutations {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.salutationsOnce.Do(func() {
		bound, _ := NewBoundResource(c, salutations.ResourceName)
		c.salutations, _ = bound.Value.(*salutations.Salutations)
	})
	return c.salutations
}

// Settings is built and cached on first use.
func (c *ExsitedClient) Settings() *settings.Settings {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.settingsOnce.Do(func() {
		bound, _ := NewBoundResource(c, settings.ResourceName)
		c.settings, _ = bound.Value.(*settings.Settings)
	})
	return c.settings
}

// SettingsWorkflows is built and cached on first use.
func (c *ExsitedClient) SettingsWorkflows() *settingsworkflows.SettingsWorkflows {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.settingsWorkflowsOnce.Do(func() {
		bound, _ := NewBoundResource(c, settingsworkflows.ResourceName)
		c.settingsWorkflows, _ = bound.Value.(*settingsworkflows.SettingsWorkflows)
	})
	return c.settingsWorkflows
}

// ShippingPackages is built and cached on first use.
func (c *ExsitedClient) ShippingPackages() *shippingpackages.ShippingPackages {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.shippingPackagesOnce.Do(func() {
		bound, _ := NewBoundResource(c, shippingpackages.ResourceName)
		c.shippingPackages, _ = bound.Value.(*shippingpackages.ShippingPackages)
	})
	return c.shippingPackages
}

// ShippingRules is built and cached on first use.
func (c *ExsitedClient) ShippingRules() *shippingrules.ShippingRules {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.shippingRulesOnce.Do(func() {
		bound, _ := NewBoundResource(c, shippingrules.ResourceName)
		c.shippingRules, _ = bound.Value.(*shippingrules.ShippingRules)
	})
	return c.shippingRules
}

// SystemEvents is built and cached on first use.
func (c *ExsitedClient) SystemEvents() *systemevents.SystemEvents {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.systemEventsOnce.Do(func() {
		bound, _ := NewBoundResource(c, systemevents.ResourceName)
		c.systemEvents, _ = bound.Value.(*systemevents.SystemEvents)
	})
	return c.systemEvents
}

// Tasks is built and cached on first use.
func (c *ExsitedClient) Tasks() *tasks.Tasks {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.tasksOnce.Do(func() {
		bound, _ := NewBoundResource(c, tasks.ResourceName)
		c.tasks, _ = bound.Value.(*tasks.Tasks)
	})
	return c.tasks
}

// Teams is built and cached on first use.
func (c *ExsitedClient) Teams() *teams.Teams {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.teamsOnce.Do(func() {
		bound, _ := NewBoundResource(c, teams.ResourceName)
		c.teams, _ = bound.Value.(*teams.Teams)
	})
	return c.teams
}

// TimePeriods is built and cached on first use.
func (c *ExsitedClient) TimePeriods() *timeperiods.TimePeriods {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.timePeriodsOnce.Do(func() {
		bound, _ := NewBoundResource(c, timeperiods.ResourceName)
		c.timePeriods, _ = bound.Value.(*timeperiods.TimePeriods)
	})
	return c.timePeriods
}

// Timesheets is built and cached on first use.
func (c *ExsitedClient) Timesheets() *timesheets.Timesheets {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.timesheetsOnce.Do(func() {
		bound, _ := NewBoundResource(c, timesheets.ResourceName)
		c.timesheets, _ = bound.Value.(*timesheets.Timesheets)
	})
	return c.timesheets
}

// Uom is built and cached on first use.
func (c *ExsitedClient) Uom() *uom.Uom {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.uomOnce.Do(func() {
		bound, _ := NewBoundResource(c, uom.ResourceName)
		c.uom, _ = bound.Value.(*uom.Uom)
	})
	return c.uom
}

// Usages is built and cached on first use.
func (c *ExsitedClient) Usages() *usages.Usages {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.usagesOnce.Do(func() {
		bound, _ := NewBoundResource(c, usages.ResourceName)
		c.usages, _ = bound.Value.(*usages.Usages)
	})
	return c.usages
}

// Webhooks is built and cached on first use.
func (c *ExsitedClient) Webhooks() *webhooks.Webhooks {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.webhooksOnce.Do(func() {
		bound, _ := NewBoundResource(c, webhooks.ResourceName)
		c.webhooks, _ = bound.Value.(*webhooks.Webhooks)
	})
	return c.webhooks
}

// Workflows is built and cached on first use.
func (c *ExsitedClient) Workflows() *workflows.Workflows {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.workflowsOnce.Do(func() {
		bound, _ := NewBoundResource(c, workflows.ResourceName)
		c.workflows, _ = bound.Value.(*workflows.Workflows)
	})
	return c.workflows
}

// Zones is built and cached on first use.
func (c *ExsitedClient) Zones() *zones.Zones {
	// Through the registry, so this accessor and ListResourceNames cannot disagree.
	c.zonesOnce.Do(func() {
		bound, _ := NewBoundResource(c, zones.ResourceName)
		c.zones, _ = bound.Value.(*zones.Zones)
	})
	return c.zones
}

// ListResourceNames returns every resource reachable on this client, e.g. ["accounts"].
func (c *ExsitedClient) ListResourceNames() []string {
	return RegisteredResourceNames()
}

// SendApiRequest sends one request and returns the response: every argument on the options.
func (c *ExsitedClient) SendApiRequest(
	ctx context.Context, httpMethod, path string, options *exhttp.RequestOptions,
) (*exhttp.ApiResponse, error) {
	// Ahead of every other check and of the token grant; see refuseIfClosed and acceptSend.
	finished, err := c.acceptSend()
	if err != nil {
		return nil, err
	}
	defer finished()
	if ctx == nil {
		ctx = context.Background()
	}
	if err := refuseIfBudgetSpent(ctx, "before the request was prepared"); err != nil {
		return nil, err
	}
	// Frozen before anything reads it, so later caller writes cannot reach this call.
	var requestOptions exhttp.RequestOptions
	if options != nil {
		requestOptions = options.Frozen()
	}
	method, err := exhttp.NormaliseMethod(httpMethod)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(path) == "" {
		return nil, exerrors.NewUsageError(fmt.Sprintf(
			"path must name an endpoint, got '%s' - an empty path requests the instance root", path))
	}
	responseFormat, err := c.resolveResponseFormat(requestOptions.ResponseFormat)
	if err != nil {
		return nil, err
	}

	// One key for the whole call, so the transport's retry and the 401 replay collapse into the first attempt.
	headers := exhttp.WithIdempotencyKey(method, c.mergeRequestHeaders(requestOptions.Headers))
	if err := exhttp.CheckRequestHeaders(headers); err != nil {
		return nil, err
	}
	resolvedPath, err := exhttp.SubstitutePathParams(path, c.seedApiVersion(path, requestOptions.PathParams))
	if err != nil {
		return nil, err
	}
	url, err := exhttp.BuildURL(c.config.Auth.BaseURL, resolvedPath)
	if err != nil {
		return nil, err
	}
	queryParams, err := exhttp.BuildQueryParams(requestOptions.QueryParams)
	if err != nil {
		return nil, err
	}
	if err := rejectDuplicatedQueryParams(url, queryParams); err != nil {
		return nil, err
	}

	// Judged and encoded only after the target is settled: the order is load-bearing.
	wireJSONBody, err := validatePayload(&requestOptions)
	if err != nil {
		return nil, err
	}
	reportedBody := redact.RedactValue(requestOptions.ReportedBody())
	formEncoded := requestOptions.JSONBody == nil && requestOptions.FormData != nil

	shape := requestOptions.DTOShape
	if shape == nil && requestOptions.DtoTarget != nil {
		shape = shapeOfTarget(requestOptions.DtoTarget)
	}
	materialised, err := exhttp.TransportRequest{
		HTTPMethod:         method,
		URL:                url,
		EndpointTemplate:   endpointTemplateOf(path),
		QueryParams:        queryParams,
		JSONBody:           wireJSONBody,
		FormData:           requestOptions.FormData,
		Files:              requestOptions.Files,
		Headers:            headers,
		RequestInfoHeaders: redact.RedactMapping(headers),
		RequestInfoBody:    redact.RedactValue(summarisePayload(reportedBody, requestOptions.Files)),
		ResponseFormat:     responseFormat,
		DTOShape:           shape,
	}.Materialised()
	if err != nil {
		return nil, err
	}
	// From the materialised copy, not the caller's map, so the report describes what was sent.
	reportedFiles := materialised.Files

	// Read again: a deadline that expired during preparation must not now buy a token and send.
	if err := refuseIfBudgetSpent(ctx, "while the request was being prepared"); err != nil {
		return nil, err
	}

	accessToken := ""
	if !requestOptions.Unauthenticated {
		accessToken, err = c.tokens.AccessToken(ctx)
		if err != nil {
			return nil, err
		}
	}

	response, err := c.transport.SendHTTPRequest(ctx, withBearerToken(materialised, accessToken))
	if err != nil {
		c.reportFailedSend(err, reportedBody, reportedFiles, formEncoded)
		return nil, err
	}

	// Map, not Data: whether a 401 says the token is dead is read off the body, whatever form Data takes.
	if !requestOptions.Unauthenticated && response.StatusCode == 401 &&
		exhttp.MeansExpiredToken(response.Map, method) {
		// Reported before the replay, so a 401 -> refresh -> 200 sequence shows the 401.
		c.reportExchange(response, reportedBody, reportedFiles, formEncoded, levelWarning)
		fresh, refreshErr := c.tokens.Refresh(ctx, accessToken)
		if refreshErr != nil {
			// Compare-and-clear, so the refused token leaves the shared cache and a token minted meanwhile stays;
			// the refresh failure is the error the caller is owed, so this one's own failure is dropped.
			_, _ = c.tokens.InvalidateAccessToken(accessToken)
			return nil, refreshErr
		}
		// The same request under the same headers, and so under the same idempotency key.
		response, err = c.transport.SendHTTPRequest(ctx, withBearerToken(materialised, fresh))
		if err != nil {
			c.reportFailedSend(err, reportedBody, reportedFiles, formEncoded)
			return nil, err
		}
	}
	c.reportExchange(response, reportedBody, reportedFiles, formEncoded, levelInfo)

	raise := c.config.RaiseOnError
	if requestOptions.RaiseOnError != nil {
		raise = *requestOptions.RaiseOnError
	}
	if raise {
		if statusError := response.RaiseForStatus(); statusError != nil {
			return response, statusError
		}
	}
	return response, nil
}

// shapeOfTarget reads a body into the pointer an endpoint named, the typed view DTO() hands back.
func shapeOfTarget(target any) exhttp.DTOShape {
	return func(decoded any) any {
		if err := dto.DecodeDto(decoded, target); err != nil {
			return nil
		}
		return target
	}
}

// SendGetRequest and its siblings are SendApiRequest with the method filled in.
func (c *ExsitedClient) SendGetRequest(
	ctx context.Context, path string, o *exhttp.RequestOptions,
) (*exhttp.ApiResponse, error) {
	return c.SendApiRequest(ctx, nethttp.MethodGet, path, o)
}

func (c *ExsitedClient) SendPostRequest(
	ctx context.Context, path string, o *exhttp.RequestOptions,
) (*exhttp.ApiResponse, error) {
	return c.SendApiRequest(ctx, nethttp.MethodPost, path, o)
}

func (c *ExsitedClient) SendPutRequest(
	ctx context.Context, path string, o *exhttp.RequestOptions,
) (*exhttp.ApiResponse, error) {
	return c.SendApiRequest(ctx, nethttp.MethodPut, path, o)
}

func (c *ExsitedClient) SendPatchRequest(
	ctx context.Context, path string, o *exhttp.RequestOptions,
) (*exhttp.ApiResponse, error) {
	return c.SendApiRequest(ctx, nethttp.MethodPatch, path, o)
}

func (c *ExsitedClient) SendDeleteRequest(
	ctx context.Context, path string, o *exhttp.RequestOptions,
) (*exhttp.ApiResponse, error) {
	return c.SendApiRequest(ctx, nethttp.MethodDelete, path, o)
}

// Close releases the LogFile handle and the connections of a transport this client built.
func (c *ExsitedClient) Close() (err error) {
	// Latched first, so a send racing this Close is refused rather than handed a closing transport.
	c.closed.Store(true)
	// Then the lifetime write lock, granted once every accepted send has returned.
	c.lifetime.Lock()
	defer c.lifetime.Unlock()
	defer func() {
		c.closeMutex.Lock()
		handle := c.logFile
		c.logFile = nil // a second Close is a no-op, not a spurious error
		c.closeMutex.Unlock()
		if handle == nil {
			return
		}
		if closeErr := handle.Close(); closeErr != nil && err == nil {
			refusal := exerrors.NewConfigError(fmt.Sprintf(
				"logFile: %s failed to flush its last records to disk",
				redact.DescribePath(c.config.LogFile)))
			refusal.Cause = closeErr
			refusal.CauseText = fmt.Sprintf("%s: %v", c.config.LogFile, closeErr)
			err = refusal
		}
	}()
	if c.ownsTransport {
		c.transport.Close()
	}
	return nil
}

// CloseClient is an alias for Close.
func (c *ExsitedClient) CloseClient() error {
	return c.Close()
}

// refuseIfBudgetSpent refuses a request whose context is already done, before more work goes into it.
func refuseIfBudgetSpent(ctx context.Context, when string) error {
	cause := ctx.Err()
	if cause == nil {
		return nil
	}
	verb, errorType := "cancelled", exhttp.ErrorTypeUnknown
	if errors.Is(cause, context.DeadlineExceeded) {
		verb, errorType = "timed out", exhttp.ErrorTypeTimeout
	}
	return &exerrors.SdkError{
		Message:   fmt.Sprintf("%s %s, so nothing was sent", verb, when),
		ErrorType: errorType,
		Cause:     cause,
	}
}

// acceptSend refuses a closed client, else takes the lifetime read lock the send holds until it returns.
func (c *ExsitedClient) acceptSend() (func(), error) {
	if err := c.refuseIfClosed(); err != nil {
		return nil, err
	}
	c.lifetime.RLock()
	if err := c.refuseIfClosed(); err != nil {
		c.lifetime.RUnlock()
		return nil, err
	}
	return c.lifetime.RUnlock, nil
}

func (c *ExsitedClient) refuseIfClosed() error {
	if !c.closed.Load() {
		return nil
	}
	return exerrors.NewUsageError(fmt.Sprintf(
		"client '%s' is closed: Close() has already run, so this request was not sent - "+
			"build a new client rather than reusing a closed one",
		c.config.Auth.ClientName))
}

// withBearerToken applies one attempt's bearer token: the only change a stale-token replay may make.
func withBearerToken(materialised exhttp.TransportRequest, accessToken string) exhttp.TransportRequest {
	if accessToken == "" {
		return materialised
	}
	headers := make(map[string]string, len(materialised.Headers)+1)
	for name, value := range materialised.Headers {
		// Any caller spelling of Authorization is replaced, not sent beside the SDK's own.
		if !strings.EqualFold(name, authorizationHeader) {
			headers[name] = value
		}
	}
	headers[authorizationHeader] = "Bearer " + accessToken
	out := materialised
	out.Headers = headers
	out.RequestInfoHeaders = redact.RedactMapping(headers)
	return out
}

// resolveResponseFormat picks the per-call format if given, otherwise the client's own.
func (c *ExsitedClient) resolveResponseFormat(perCall string) (string, error) {
	if perCall == "" {
		return c.config.ResponseFormat, nil
	}
	format := strings.ToLower(strings.TrimSpace(perCall))
	if !config.IsResponseFormat(format) {
		return "", exerrors.NewUsageError(fmt.Sprintf(
			"responseFormat: expected one of %s, got '%s'",
			strings.Join(config.ResponseFormatNames(), ", "), perCall))
	}
	return format, nil
}

// mergeRequestHeaders merges SDK, config and caller headers case-insensitively; the caller's values win.
func (c *ExsitedClient) mergeRequestHeaders(callerHeaders map[string]string) map[string]string {
	canonical := map[string]string{}   // lower -> canonical spelling
	byLowerName := map[string]string{} // lower -> value

	putHeader(canonical, byLowerName, "Accept", "application/json")
	putHeader(canonical, byLowerName, "User-Agent", c.config.UserAgent)
	for _, name := range sortedStringKeys(c.config.DefaultHeaders) {
		putHeader(canonical, byLowerName, name, c.config.DefaultHeaders[name])
	}
	for _, name := range sortedStringKeys(callerHeaders) {
		putHeader(canonical, byLowerName, name, callerHeaders[name])
	}
	merged := make(map[string]string, len(byLowerName))
	for lower, value := range byLowerName {
		merged[canonical[lower]] = value
	}
	return merged
}

func putHeader(canonical, byLowerName map[string]string, name, value string) {
	lower := strings.ToLower(name)
	// The last spelling seen wins; only the three SDK-owned names below are restored to canonical form.
	canonical[lower] = name
	switch lower {
	case "authorization":
		canonical[lower] = authorizationHeader
	case "accept":
		canonical[lower] = "Accept"
	case "user-agent":
		canonical[lower] = "User-Agent"
	}
	byLowerName[lower] = value
}

func (c *ExsitedClient) seedApiVersion(path string, pathParams map[string]any) map[string]any {
	resolved := make(map[string]any, len(pathParams)+1)
	maps.Copy(resolved, pathParams)
	if slices.Contains(exhttp.FindPathParamNames(path), apiVersionParam) {
		if _, ok := resolved[apiVersionParam]; !ok {
			resolved[apiVersionParam] = string(c.config.APIVersion)
		}
	}
	return resolved
}

// validatePayload judges the payload before the token grant, so a bad body costs no OAuth2 round trip.
func validatePayload(requestOptions *exhttp.RequestOptions) (any, error) {
	// A typed payload is a view over the map form, never a second wire format.
	if payload, isDTO := requestOptions.JSONBody.(exhttp.DTOPayload); isDTO {
		rendered := payload.ToMap()
		if rendered == nil {
			return nil, exerrors.NewUsageError(fmt.Sprintf(
				"jsonBody: %T rendered no body at all - a DTO payload must return the map to send, "+
					"and nil would send the request with no body", requestOptions.JSONBody))
		}

		requestOptions.JSONBody = exhttp.SnapshotJSONValue(rendered)
	} else if isStructPayload(requestOptions.JSONBody) {
		// Typed views have no ToMap of their own; the dto base renders them.
		rendered, err := dto.EncodeDto(requestOptions.JSONBody)
		if err != nil {
			return nil, err
		}
		requestOptions.JSONBody = exhttp.SnapshotJSONValue(rendered)
	}
	if requestOptions.JSONBody != nil && requestOptions.FormData != nil {
		return nil, exerrors.NewUsageError("pass jsonBody or formData, not both - they are different content types " +
			"and only one can be sent")
	}
	if requestOptions.JSONBody != nil && requestOptions.Files != nil {
		return nil, exerrors.NewUsageError("pass jsonBody or files, not both - a multipart body replaces the JSON " +
			"one, so the JSON would be discarded unsent")
	}
	var wireJSONBody any
	// A string body is JSON the caller wrote, verified here and then sent byte for byte.
	if raw, isRawJSON := requestOptions.JSONBody.(string); isRawJSON {
		if err := validateRawJSONBody(raw); err != nil {
			return nil, err
		}
		wireJSONBody = raw
	} else if requestOptions.JSONBody != nil {
		switch reflect.ValueOf(requestOptions.JSONBody).Kind() {
		case reflect.Map, reflect.Slice, reflect.Array:
		default:
			return nil, exerrors.NewUsageError(fmt.Sprintf(
				"jsonBody must be a mapping or list, got %T", requestOptions.JSONBody))
		}

		encoded, err := exhttp.EncodeJSONBody(requestOptions.JSONBody)
		if err != nil {
			return nil, exhttp.RefuseUnencodableJSONBody(err)
		}
		wireJSONBody = string(encoded)
	}
	// The same pre-flight for a multipart body, so the grant is not spent on one that cannot be built.
	if err := exhttp.CheckFiles(requestOptions.Files); err != nil {
		return nil, err
	}
	if len(requestOptions.Files) > 0 {
		// Only a multipart body puts formData names into part headers, where a line break could forge one.
		if err := exhttp.CheckFormDataNames(requestOptions.FormData); err != nil {
			return nil, err
		}
	}
	return wireJSONBody, nil
}

// isStructPayload reports a struct, or a pointer to one, handed over as the body.
func isStructPayload(body any) bool {
	value := reflect.ValueOf(body)
	for value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return true
		}
		value = value.Elem()
	}
	return value.Kind() == reflect.Struct
}

func validateRawJSONBody(jsonBody string) error {
	var decoded any
	if err := json.Unmarshal([]byte(jsonBody), &decoded); err != nil {
		// The byte offset, not the decoder's words, which quote the caller's own document.
		refusal := exerrors.NewUsageError(fmt.Sprintf(
			"jsonBody: the text is not valid JSON%s", jsonSyntaxPosition(err)))
		refusal.Cause = err
		refusal.CauseText = err.Error()
		return refusal
	}
	switch decoded.(type) {
	case map[string]any, []any:
		return nil
	}
	return exerrors.NewUsageError(fmt.Sprintf(
		"jsonBody: a JSON string must hold an object or an array, got %s", jsonKindName(decoded)))
}

// jsonSyntaxPosition names the byte offset where a JSON document stopped parsing, when the decoder says.
func jsonSyntaxPosition(err error) string {
	var syntax *json.SyntaxError
	if errors.As(err, &syntax) {
		return fmt.Sprintf(" - it stops parsing at byte %d", syntax.Offset)
	}
	var unmarshalType *json.UnmarshalTypeError
	if errors.As(err, &unmarshalType) {
		return fmt.Sprintf(" - a %s where a value was expected, at byte %d",
			unmarshalType.Value, unmarshalType.Offset)
	}
	return ""
}

// jsonKindName names what a JSON document decoded to in JSON's vocabulary: "a number", not "float64".
func jsonKindName(decoded any) string {
	switch decoded.(type) {
	case nil:
		return "null"
	case bool:
		return "a boolean"
	case float64:
		return "a number"
	case string:
		return "a string"
	}
	return fmt.Sprintf("%T", decoded)
}

func rejectDuplicatedQueryParams(url string, queryParams map[string]any) error {
	if len(queryParams) == 0 {
		return nil
	}
	// Parsed with net/url, so a '?' in the fragment cannot fake a query and encoded names still clash.
	parsed, err := neturl.Parse(url)
	if err != nil {
		return nil // the URL builder has already vetted the origin
	}
	// ParseQuery decodes what it can; the decoded names are still the honest comparison set.
	existing, _ := neturl.ParseQuery(parsed.RawQuery)
	if len(existing) == 0 {
		return nil
	}
	var clash []string
	for k := range queryParams {
		if _, ok := existing[k]; ok {
			clash = append(clash, k)
		}
	}
	sort.Strings(clash)
	if len(clash) > 0 {
		return exerrors.NewUsageError(fmt.Sprintf("queryParams %s already appear in the path's own query string, "+
			"so each would be sent twice with two values and the API chooses - pass them once, in one place",
			strings.Join(clash, ", ")))
	}
	return nil
}

func progressFunc(sdkConfig config.SdkConfig, log func(string)) func(string) {
	return func(message string) {
		if sdkConfig.PrintsApiInfo() && !sdkConfig.PrintOnlyOnError {
			fmt.Fprintln(printWriter(sdkConfig), neutralise(message))
		}
		// An info record, so LogLevel "INFO" gets the timeline without touching stdout.
		if log != nil && logAllows(sdkConfig, levelInfo) {
			log(neutralise(message))
		}
	}
}

func (c *ExsitedClient) reportFailedSend(err error, reportedBody any,
	files map[string]any, formEncoded bool) {
	var sdkError *exerrors.SdkError
	if !errors.As(err, &sdkError) {
		return
	}
	info, ok := sdkError.RequestInfo.(*exhttp.RequestInfo)
	if !ok || info == nil {
		return
	}
	failed := &exhttp.ApiResponse{
		StatusCode:  0, // rendered as "no response"
		Data:        nil,
		Headers:     map[string]string{},
		RequestInfo: info,
	}
	c.reportExchange(failed, reportedBody, files, formEncoded, levelError)
}

func (c *ExsitedClient) reportExchange(response *exhttp.ApiResponse, reportedBody any,
	files map[string]any, formEncoded bool, level string) {
	wantsPrint := c.config.PrintsApiInfo() || c.config.PrintRequestPayload || c.config.PrintResponseBody ||
		c.config.PrintHeaders || c.config.PrintCurl
	// Reported even without a request record: Redacted absorbs a nil one.
	if wantsPrint && !(c.config.PrintOnlyOnError && response.Succeeded()) {
		fmt.Fprintln(printWriter(*c.config), renderExchange(response, reportedBody, files, formEncoded, *c.config))
	}

	if c.log != nil {
		if logAllows(*c.config, level) {
			// Redacted(), not the raw record: a host-built &RequestInfo{} never met the redacting constructor.
			info := response.RequestInfo.Redacted()
			// Named as SdkError.Error() names a call: by endpoint template, else by the resolved path's shape.
			c.log(neutralise(fmt.Sprintf("%s%s %s -> %s in %ss (attempt %d/%d)", logRecordPrefixes[level],
				info.HTTPMethod, exerrors.EndpointLabel(info.EndpointTemplate, info.URL),
				statusLabel(response.StatusCode),
				exerrors.FormatSeconds(info.ElapsedSeconds), info.AttemptCount, info.MaxAttempts)))
		}
		// Outside the record guard: the bodies do not come from it.
		if c.config.LogBodies && logAllows(*c.config, levelDebug) {
			// RenderJSONValue, not %v, which would write raw newlines and terminal escapes into the log.
			c.log(fmt.Sprintf("DEBUG payload=%s response=%s",
				RenderJSONValue(redact.RedactValue(reportedBody), true),
				RenderJSONValue(redact.RedactValue(response.Map), true)))
		}
	}
}

// endpointTemplateOf is the path as the caller wrote it, placeholders unresolved, for error rendering.
func endpointTemplateOf(path string) string {
	lowered := strings.ToLower(strings.TrimSpace(path))
	if strings.HasPrefix(lowered, "http://") || strings.HasPrefix(lowered, "https://") {
		return ""
	}
	return path
}
