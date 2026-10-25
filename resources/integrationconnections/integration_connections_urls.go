// Package integrationconnections is the v4 integration connections endpoints.
package integrationconnections

// The path templates for the integration connections endpoints; the client fills {api_version}.
const (
	IntegrationConnectionsPath = "/api/{api_version}/integration_connections"
	IntegrationConnectionPath  = "/api/{api_version}/integration_connections/{integration_connection_id}"
)

// IntegrationConnectionsKey is the key a listing's records arrive under, the v4 list envelope's data.
const IntegrationConnectionsKey = "data"
