// Package webhooks is the v4 webhooks endpoints.
package webhooks

// The path templates for the webhooks endpoints; the client fills {api_version}.
const (
	WebhookPath = "/api/{api_version}/webhooks/{webhook_id}"
)

// WebhookKey is the envelope key one record arrives under.
const WebhookKey = "webhook"
