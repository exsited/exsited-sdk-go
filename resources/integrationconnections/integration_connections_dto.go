package integrationconnections

import "github.com/exsited/exsited-sdk-go/v4/dto"

// Pagination is the paging block, either shape.
type Pagination struct {
	Limit *int64

	// Cursor shape.
	HasMore    *bool
	NextCursor *string
	PrevCursor any
	Links      *PaginationLinks

	// Offset shape.
	Records *int64
	Offset  *int64
	// NextPage is the literal string "NULL" at the end of a listing, not null.
	NextPage     *string
	PreviousPage *string
}

// IntegrationConnectionsRecord is the integration connections body.
type IntegrationConnectionsRecord struct {
	Data       []any
	Object     *string
	Pagination *Pagination
	URL        *string
}

// MeasuredFrom lists the endpoints IntegrationConnectionsRecord was measured from.
func (IntegrationConnectionsRecord) MeasuredFrom() []string {
	return []string{
		"GET /api/{api_version}/integration_connections",
	}
}

// PaginationLinks is the links block of a cursor Pagination.
type PaginationLinks struct {
	Next any
	Prev any
}

// IntegrationConnectionsRecordShape reads a body as IntegrationConnectionsRecord.
func IntegrationConnectionsRecordShape(decoded any) any {
	return dto.Read[IntegrationConnectionsRecord](decoded)
}
