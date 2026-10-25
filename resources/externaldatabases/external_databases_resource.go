package externaldatabases

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ExternalDatabases method of the same name,
// which is the primary surface.

// GetExternalDatabases is the module form of ExternalDatabases.GetExternalDatabases.
func GetExternalDatabases(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetExternalDatabases(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateExternalDatabase is the module form of ExternalDatabases.CreateExternalDatabase.
func CreateExternalDatabase(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateExternalDatabase(ctx, payload, raiseOnError, responseFormat...)
}

// GetExternalDatabase is the module form of ExternalDatabases.GetExternalDatabase.
func GetExternalDatabase(requester exhttp.Requester, ctx context.Context, dbName any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetExternalDatabase(ctx, dbName, queryParams, raiseOnError, responseFormat...)
}

// CreateExternalDatabaseInsertRecord is the module form of ExternalDatabases.CreateExternalDatabaseInsertRecord.
func CreateExternalDatabaseInsertRecord(requester exhttp.Requester, ctx context.Context, dbName any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateExternalDatabaseInsertRecord(ctx, dbName, payload, raiseOnError, responseFormat...)
}

// CreateExternalDatabaseSearch is the module form of ExternalDatabases.CreateExternalDatabaseSearch.
func CreateExternalDatabaseSearch(requester exhttp.Requester, ctx context.Context, dbName any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateExternalDatabaseSearch(ctx, dbName, payload, raiseOnError, responseFormat...)
}

// UpdateExternalDatabaseUpdateRecords is the module form of ExternalDatabases.UpdateExternalDatabaseUpdateRecords.
func UpdateExternalDatabaseUpdateRecords(requester exhttp.Requester, ctx context.Context, dbName any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateExternalDatabaseUpdateRecords(ctx, dbName, payload, raiseOnError, responseFormat...)
}

// GetExternalDatabaseTable is the module form of ExternalDatabases.GetExternalDatabaseTable.
func GetExternalDatabaseTable(requester exhttp.Requester, ctx context.Context, dbName any, tableName any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetExternalDatabaseTable(ctx, dbName, tableName, queryParams, raiseOnError, responseFormat...)
}
