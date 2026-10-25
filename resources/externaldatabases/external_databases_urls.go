// Package externaldatabases is the v4 external databases endpoints.
package externaldatabases

// The path templates for the external databases endpoints; the client fills {api_version}.
const (
	ExternalDatabasesPath             = "/api/{api_version}/external_databases"
	ExternalDatabasePath              = "/api/{api_version}/external_databases/{db_name}"
	ExternalDatabaseInsertRecordsPath = "/api/{api_version}/external_databases/{db_name}/insert_records"
	ExternalDatabaseSearchPath        = "/api/{api_version}/external_databases/{db_name}/search"
	ExternalDatabaseUpdateRecordsPath = "/api/{api_version}/external_databases/{db_name}/update_records"
	ExternalDatabaseTablePath         = "/api/{api_version}/external_databases/{db_name}/tables/{table_name}"
)

// ExternalDatabasesKey is the key a listing's records arrive under, the v4 list envelope's data.
const ExternalDatabasesKey = "data"

// ExternalDatabaseKey is the envelope key one record arrives under.
const ExternalDatabaseKey = "external_database"
