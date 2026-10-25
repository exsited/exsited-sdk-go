package client

import (
	"context"
	"fmt"
	"iter"
	"sort"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The query parameters the pager owns.
const (
	CursorParam = "cursor"
	OffsetParam = "offset"
	LimitParam  = "limit"
)

// DefaultRecordKey is the envelope key holding a page's records on v4; offset listings key them by family name.
const DefaultRecordKey = "data"

// NoNextPage is the string the offset shape uses to mean "no next page"; never truth-test next_page.
const NoNextPage = "NULL"

type PaginateOptions struct {
	// QueryParams travel with every page: a cursor or offset is a position in one filtered, sorted listing.
	QueryParams map[string]any
	// MaxPages caps the walk; 0 means unlimited.
	MaxPages int
	// RaiseOnError: nil or true stops with an error on a failed page; false yields that page last.
	RaiseOnError *bool
	// PathParams fill the path template on every page.
	PathParams map[string]any
	// RecordKey is the envelope key an offset page's records sit under, which sizes the next offset.
	RecordKey string
}

// PageOptions is an alias of PaginateOptions.
type PageOptions = PaginateOptions

// IterPages walks a paginated listing, one page per yield, in whichever of v4's two shapes it is served.
func IterPages(ctx context.Context, client exhttp.Requester, path string, options PaginateOptions) iter.Seq2[*exhttp.ApiResponse, error] {
	return func(yield func(*exhttp.ApiResponse, error) bool) {
		for _, owned := range []string{CursorParam, OffsetParam} {
			if _, present := options.QueryParams[owned]; present {
				yield(nil, exerrors.NewUsageError(fmt.Sprintf(
					"queryParams: '%s' is owned by the pager, which sets it from each page; "+
						"to resume from a saved position call the endpoint directly with it", owned)))
				return
			}
		}
		if options.MaxPages < 0 {
			yield(nil, exerrors.NewUsageError(fmt.Sprintf("maxPages: expected a positive integer or 0 for unlimited, got %d", options.MaxPages)))
			return
		}
		raise := options.RaiseOnError == nil || *options.RaiseOnError
		recordKey := options.RecordKey
		if recordKey == "" {
			recordKey = DefaultRecordKey
		}

		seenCursors := map[string]bool{}
		yielded := 0
		next := withPosition(options.QueryParams, "", nil)

		for {
			response, err := client.SendGetRequest(ctx, path, &exhttp.RequestOptions{
				PathParams: options.PathParams, QueryParams: next, RaiseOnError: &raise})
			if err != nil {
				yield(nil, err)
				return
			}
			if !yield(response, nil) {
				return
			}
			yielded++

			if !response.Succeeded() {
				return // only reachable with RaiseOnError=false
			}
			if options.MaxPages > 0 && yielded >= options.MaxPages {
				return
			}
			// Map, not Data: the paging block is a fact about the body, whatever Data presents.
			block := paginationBlock(response.Map)
			if block == nil {
				return
			}
			if cursor := ReadNextCursor(block); cursor != "" {
				if seenCursors[cursor] {
					yield(nil, &exerrors.SdkError{Message: fmt.Sprintf(
						"pagination loop: after %d page(s) of a %s listing the API handed back a next_cursor "+
							"it had already served, so following it again would never terminate - the pages "+
							"already yielded are good", yielded, exerrors.EndpointLabel(path, ""))})
					return
				}
				seenCursors[cursor] = true
				next = withPosition(options.QueryParams, CursorParam, cursor)
				continue
			}
			onPage := 0
			if records, isList := recordsUnder(response, recordKey); isList {
				onPage = len(records)
			}
			offset, ok := ReadNextOffset(block, onPage)
			if !ok {
				return
			}
			// The offset is computed, not the link followed: a next_page link can point at another path.
			next = withPosition(options.QueryParams, OffsetParam, offset)
		}
	}
}

// withPosition copies the caller's query and adds the pager's own position parameter, when there is one.
func withPosition(query map[string]any, name string, value any) map[string]any {
	out := make(map[string]any, len(query)+1)
	for key, entry := range query {
		out[key] = entry
	}
	if name != "" {
		out[name] = value
	}
	return out
}

// IterRecords walks every record of a listing across every page; a failed page always errors and stops.
func IterRecords(ctx context.Context, client exhttp.Requester, path, recordKey string, options PaginateOptions) iter.Seq2[any, error] {
	return func(yield func(any, error) bool) {
		if strings.TrimSpace(recordKey) == "" {
			yield(nil, exerrors.NewUsageError(fmt.Sprintf(
				"recordKey: expected the envelope key holding the records, like 'data' or 'accounts', got '%s'", recordKey)))
			return
		}
		alwaysRaise := true
		options.RaiseOnError = &alwaysRaise
		options.RecordKey = recordKey
		pages := 0
		for page, err := range IterPages(ctx, client, path, options) {
			if err != nil {
				yield(nil, err)
				return
			}
			pages++
			records, isList := recordsUnder(page, recordKey)
			if !isList {
				// On v4 a later page with no list under the key ends the records.
				if pages > 1 {
					return
				}
				yield(nil, missingRecordKeyError(page, recordKey))
				return
			}
			for _, record := range records {
				if !yield(record, nil) {
					return
				}
			}
		}
	}
}

// IterPages is the package IterPages over this client, calling visit per page; visit returns false to stop.
func (c *ExsitedClient) IterPages(
	ctx context.Context, path string, options PageOptions,
	visit func(*exhttp.ApiResponse) bool,
) error {
	lenient := false
	options.RaiseOnError = &lenient
	for page, err := range IterPages(ctx, c, path, options) {
		if err != nil {
			return err
		}
		if !visit(page) {
			return nil
		}
		if !page.Succeeded() {
			return page.RaiseForStatus()
		}
	}
	return nil
}

// IterRecords collects every record of a listing across every page, under RecordKey or DefaultRecordKey.
func (c *ExsitedClient) IterRecords(
	ctx context.Context, path string, options PageOptions,
) ([]any, error) {
	recordKey := options.RecordKey
	if recordKey == "" {
		recordKey = DefaultRecordKey
	}
	var out []any
	for record, err := range IterRecords(ctx, c, path, recordKey, options) {
		if err != nil {
			return nil, err
		}
		out = append(out, record)
	}
	return out, nil
}

// ReadNextCursor is the cursor for the next page, or "" when this is not the cursor shape or the listing ended.
func ReadNextCursor(block map[string]any) string {
	// has_more is authoritative: the last page of some listings carries a stale next_cursor beside has_more:false.
	if hasMore, present := block["has_more"]; present {
		if flag, ok := hasMore.(bool); !ok || !flag {
			return ""
		}
	}
	cursor, ok := block["next_cursor"].(string)
	if !ok || strings.TrimSpace(cursor) == "" {
		return ""
	}
	return cursor
}

// ReadNextOffset is the offset for the next page, and false at the end; next_page is compared to NoNextPage.
func ReadNextOffset(block map[string]any, recordsOnPage int) (int64, bool) {
	raw, present := block["next_page"]
	if !present || raw == nil {
		return 0, false
	}
	if text, ok := raw.(string); ok {
		trimmed := strings.TrimSpace(text)
		if trimmed == "" || strings.EqualFold(trimmed, NoNextPage) {
			return 0, false
		}
	}
	// NumberAsInt64, not a float64 assertion: the transport decodes with UseNumber.
	current := int64(0)
	if offset, ok := exhttp.NumberAsInt64(block[OffsetParam]); ok {
		current = offset
	}
	step := int64(recordsOnPage)
	if limit, ok := exhttp.NumberAsInt64(block[LimitParam]); ok && limit > 0 {
		step = limit
	}
	if step <= 0 {
		return 0, false
	}
	return current + step, true
}

func paginationBlock(body any) map[string]any {
	record, ok := body.(map[string]any)
	if !ok {
		return nil
	}
	block, ok := record["pagination"].(map[string]any)
	if !ok {
		return nil
	}
	return block
}

func recordsUnder(page *exhttp.ApiResponse, recordKey string) ([]any, bool) {
	if data, ok := page.Map.(map[string]any); ok {
		if list, ok := data[recordKey].([]any); ok {
			return list, true
		}
	}
	return nil, false
}

// missingRecordKeyError names the keys the page does carry, so a wrong key diagnoses itself.
func missingRecordKeyError(page *exhttp.ApiResponse, recordKey string) error {
	// The envelope's own type, so a list, a string and a number each diagnose differently.
	available := fmt.Sprintf("a %T body", page.Map)
	if data, ok := page.Map.(map[string]any); ok {
		keys := make([]string, 0, len(data))
		for k := range data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		available = strings.Join(keys, ", ")
	}
	return exerrors.NewUsageError(fmt.Sprintf(
		"recordKey '%s': this page has no list under that key; it carries %s", recordKey, available))
}
