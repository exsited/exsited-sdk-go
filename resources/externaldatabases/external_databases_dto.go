package externaldatabases

import "github.com/exsited/exsited-sdk-go/v4/dto"

// Pagination is the paging block, either shape.
type Pagination struct {
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	PrevCursor any
	PerPage    any
	Links      *PaginationLinks
	HasPrev    *bool
	HasNext    *bool

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	NextPage     *string
	PreviousPage *string
}

// ExternalDatabasesRecord is the external databases body.
type ExternalDatabasesRecord struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ExternalDatabasesRecord was measured from.
func (ExternalDatabasesRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/external_databases",
		"GET /api/{api_version}/external_databases/{db_name}",
	}
}

// ExternalDatabaseCreateRecord is the body CreateExternalDatabase answers.
type ExternalDatabaseCreateRecord struct {
	ActivityLogs []any
	Message      *string
	Name         *string
}

// MeasuredFrom lists the endpoints ExternalDatabaseCreateRecord was measured from.
func (ExternalDatabaseCreateRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/external_databases",
	}
}

// EnvelopeKey unwraps {"external_database": {...}}.
func (ExternalDatabaseCreateRecord) EnvelopeKey() string { return ExternalDatabaseKey }

// ExternalDatabaseInsertRecordCreateRecord is the body CreateExternalDatabaseInsertRecord answers.
type ExternalDatabaseInsertRecordCreateRecord struct {
	InsertedData []map[string]any
	Message      *string
	Success      *bool
}

// MeasuredFrom lists the endpoints ExternalDatabaseInsertRecordCreateRecord was measured from.
func (ExternalDatabaseInsertRecordCreateRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/external_databases/{db_name}/insert_records",
	}
}

// EnvelopeKey unwraps {"external_database": {...}}.
func (ExternalDatabaseInsertRecordCreateRecord) EnvelopeKey() string { return ExternalDatabaseKey }

// ExternalDatabaseSearchCreateRecord is the body CreateExternalDatabaseSearch answers.
type ExternalDatabaseSearchCreateRecord struct {
	Count   *int64
	Limit   *int64
	List    map[string]any
	Offset  *int64
	Success *bool
}

// MeasuredFrom lists the endpoints ExternalDatabaseSearchCreateRecord was measured from.
func (ExternalDatabaseSearchCreateRecord) MeasuredFrom() []string {
	return []string{
		"POST /api/{api_version}/external_databases/{db_name}/search",
	}
}

// EnvelopeKey unwraps {"external_database": {...}}.
func (ExternalDatabaseSearchCreateRecord) EnvelopeKey() string { return ExternalDatabaseKey }

// ExternalDatabaseTableList is the page GetExternalDatabaseTable answers.
type ExternalDatabaseTableList struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints ExternalDatabaseTableList was measured from.
func (ExternalDatabaseTableList) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/external_databases/{db_name}/tables/{table_name}",
	}
}

// ExternalDatabaseUpdateRecordsUpdateRecord is the body UpdateExternalDatabaseUpdateRecords answers.
type ExternalDatabaseUpdateRecordsUpdateRecord struct {
	Message *string
}

// MeasuredFrom lists the endpoints ExternalDatabaseUpdateRecordsUpdateRecord was measured from.
func (ExternalDatabaseUpdateRecordsUpdateRecord) MeasuredFrom() []string {
	return []string{
		"PATCH /api/{api_version}/external_databases/{db_name}/update_records",
	}
}

// EnvelopeKey unwraps {"external_database": {...}}.
func (ExternalDatabaseUpdateRecordsUpdateRecord) EnvelopeKey() string { return ExternalDatabaseKey }

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// ExternalDatabasesRecordShape reads a body as ExternalDatabasesRecord.
func ExternalDatabasesRecordShape(decoded any) any { return dto.Read[ExternalDatabasesRecord](decoded) }

// ExternalDatabaseCreateRecordShape reads a body as ExternalDatabaseCreateRecord.
func ExternalDatabaseCreateRecordShape(decoded any) any {
	return dto.Read[ExternalDatabaseCreateRecord](decoded)
}

// ExternalDatabaseInsertRecordCreateRecordShape reads a body as ExternalDatabaseInsertRecordCreateRecord.
func ExternalDatabaseInsertRecordCreateRecordShape(decoded any) any {
	return dto.Read[ExternalDatabaseInsertRecordCreateRecord](decoded)
}

// ExternalDatabaseSearchCreateRecordShape reads a body as ExternalDatabaseSearchCreateRecord.
func ExternalDatabaseSearchCreateRecordShape(decoded any) any {
	return dto.Read[ExternalDatabaseSearchCreateRecord](decoded)
}

// ExternalDatabaseTableListShape reads a body as ExternalDatabaseTableList.
func ExternalDatabaseTableListShape(decoded any) any {
	return dto.Read[ExternalDatabaseTableList](decoded)
}

// ExternalDatabaseUpdateRecordsUpdateRecordShape reads a body as ExternalDatabaseUpdateRecordsUpdateRecord.
func ExternalDatabaseUpdateRecordsUpdateRecordShape(decoded any) any {
	return dto.Read[ExternalDatabaseUpdateRecordsUpdateRecord](decoded)
}
