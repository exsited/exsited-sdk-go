package http

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	"github.com/exsited/exsited-sdk-go/v4/redact"
)

// SortDirection is a query-parameter sort direction; it serialises to its wire value.
//
// Which v4 parameter carries it is a fact about the family; see SortParam and SortParams.
type SortDirection string

// The two sort directions.
const (
	SortAsc  SortDirection = "asc"
	SortDesc SortDirection = "desc"
)

// BuildQueryParams normalises every parameter to its query-string text: bools, ISO-8601 instants, repeated slices.
func BuildQueryParams(params map[string]any) (map[string]any, error) {
	out := make(map[string]any, len(params))
	names := make([]string, 0, len(params))
	for key := range params {
		names = append(names, key)
	}
	sort.Strings(names)
	for _, key := range names {
		value := params[key]
		if value == nil {
			continue
		}
		normalised, err := normaliseQueryValue(key, value, 0, map[uintptr]bool{})
		if err != nil {
			return nil, err
		}
		out[key] = normalised
	}
	return out, nil
}

// normaliseQueryValue normalises one value; seen names a cyclic slice rather than tripping the depth limit.
func normaliseQueryValue(key string, value any, depth int, seen map[uintptr]bool) (any, error) {
	if depth >= redact.MaxRedactionDepth {
		return nil, exerrors.NewUsageError(fmt.Sprintf(
			"query param '%s': nested more than %d levels deep, which is too deep to "+
				"encode in a query string", key, redact.MaxRedactionDepth))
	}
	switch v := value.(type) {
	case bool:
		if v {
			return "true", nil
		}
		return "false", nil
	case string:
		return v, nil
	case json.Number:
		return v, nil // kept as a number; stringified at the wire
	case SortDirection:
		return string(v), nil
	case time.Time:
		return isoTime(v), nil
	case []byte:
		return string(v), nil
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return v, nil
	}

	reflectValue := reflect.ValueOf(value)
	switch reflectValue.Kind() {
	case reflect.Map:
		return nil, exerrors.NewUsageError(fmt.Sprintf(
			"query param '%s': nested objects cannot be encoded in a query string; flatten it "+
				"or send it in the body", key))
	case reflect.Slice, reflect.Array:
		// Self-reference before the depth rule: a cyclic []any is its own diagnosis.
		identity := uintptr(0)
		if reflectValue.Kind() == reflect.Slice {
			identity = reflectValue.Pointer()
			if seen[identity] {
				return nil, exerrors.NewUsageError(fmt.Sprintf(
					"query param '%s': the value refers to itself, so it has no finite "+
						"query-string encoding", key))
			}
		}
		if depth >= 1 {
			return nil, exerrors.NewUsageError(fmt.Sprintf(
				"query param '%s': a sequence inside a sequence has no query-string encoding; "+
					"flatten it to one list of scalars", key))
		}
		if identity != 0 {
			seen[identity] = true
		}
		list := make([]any, 0, reflectValue.Len())
		for i := 0; i < reflectValue.Len(); i++ {
			item := reflectValue.Index(i).Interface()
			if item == nil {
				continue
			}
			normalised, err := normaliseQueryValue(key, item, depth+1, seen)
			if err != nil {
				return nil, err
			}
			list = append(list, normalised)
		}
		return list, nil
	}
	return fmt.Sprint(value), nil
}

// isoTime always carries the UTC offset, with a fractional part only when the instant has one.
func isoTime(t time.Time) string {
	fraction := ""
	if nanoseconds := t.Nanosecond(); nanoseconds != 0 {
		if nanoseconds%1000 == 0 {
			fraction = fmt.Sprintf(".%06d", nanoseconds/1000)
		} else {
			fraction = fmt.Sprintf(".%09d", nanoseconds)
		}
	}
	return t.Format("2006-01-02T15:04:05") + fraction + t.Format("-07:00")
}

// FilterOperators are the suffixes a filter name can carry, the union of v4's four per-type lists;
// "eq" is the absence of a suffix, never a literal _eq.
//
// starts_with is a name the Text list advertises, but the API answers it with 200 and no rows.
var FilterOperators = []string{
	"eq", "neq", "in", "nin", "contains", "not_contains", "starts_with",
	"gt", "gte", "lt", "lte", "has_value", "does_not_have_value",
}

// PresenceOperators take no right-hand side: the value is not read, so pass true, never nil, which is dropped.
var PresenceOperators = []string{"has_value", "does_not_have_value"}

// ListOperators take a comma-separated list in one parameter; build the value with FilterValues.
var ListOperators = []string{"in", "nin"}

// Ascending and Descending are aliases of SortAsc and SortDesc.
const (
	Ascending  = SortAsc
	Descending = SortDesc
)

// FilterParam builds one filter's parameter name, filter_<field> or filter_<field>_<operator>.
//
// Only the operator's existence is checked: the API answers a wrong type for it with a better 400.
func FilterParam(field string, operator ...string) (string, error) {
	if strings.TrimSpace(field) == "" {
		return "", exerrors.UsageErrorf("filter field: expected a name like 'status', got %q", field)
	}
	resolved := "eq"
	if len(operator) > 0 && operator[0] != "" {
		resolved = strings.ToLower(operator[0])
	}
	known := false
	for _, candidate := range FilterOperators {
		if candidate == resolved {
			known = true
			break
		}
	}
	if !known {
		return "", exerrors.UsageErrorf("filter operator %q: expected one of %s",
			resolved, strings.Join(FilterOperators, ", "))
	}
	stem := "filter_" + strings.TrimSpace(field)
	if resolved == "eq" {
		return stem, nil
	}
	return stem + "_" + resolved, nil
}

// MustFilterParam is FilterParam for a literal known good at author time.
func MustFilterParam(field string, operator ...string) string {
	name, err := FilterParam(field, operator...)
	if err != nil {
		panic(err)
	}
	return name
}

// FilterValues renders an "in" or "nin" value: comma separated in one key, as v4 splits it.
//
// A slice would repeat the key, which v4 answers 200 with no rows; a member holding a comma or an empty list is refused.
func FilterValues(values ...any) (string, error) {
	rendered := make([]string, 0, len(values))
	for _, value := range values {
		if value == nil {
			return "", exerrors.UsageErrorf(
				"filter values: a member is nil, which has no value to match; drop it, " +
					"or use the 'does_not_have_value' operator if what was meant is " +
					"'this field is unset'")
		}
		normalised, err := normaliseQueryValue("filter values", value, 1, map[uintptr]bool{})
		if err != nil {
			return "", err
		}
		text := stringifyValue(normalised)
		if strings.Contains(text, ",") {
			return "", exerrors.UsageErrorf(
				"filter values: the member %q contains a comma, which the API reads as "+
					"the separator between two values, so it would match something "+
					"other than what was asked for", text)
		}
		rendered = append(rendered, text)
	}
	if len(rendered) == 0 {
		return "", exerrors.UsageErrorf(
			"filter values: expected at least one value; an empty 'in' filter is sent " +
				"as 'filter_<field>_in=' and answered 200 with no rows, which reads as " +
				"'nothing matched' rather than as the mistake it is")
	}
	return strings.Join(rendered, ","), nil
}

// MustFilterValues is FilterValues for literals known good at author time.
func MustFilterValues(values ...any) string {
	value, err := FilterValues(values...)
	if err != nil {
		panic(err)
	}
	return value
}

// SortParam builds a `sort` value, "-" prefixed to descend, for Account, Accounting Codes, Alert,
// Integration Connections and Inventory Adjustments; the wrong spelling is ignored with a 200.
func SortParam(field string, direction ...SortDirection) (string, error) {
	if strings.TrimSpace(field) == "" {
		return "", exerrors.UsageErrorf("sort field: expected a name like 'created_on', got %q", field)
	}
	resolved := SortAsc
	if len(direction) > 0 && direction[0] != "" {
		resolved = SortDirection(strings.ToLower(string(direction[0])))
	}
	switch resolved {
	case SortAsc:
		return strings.TrimSpace(field), nil
	case SortDesc:
		return "-" + strings.TrimSpace(field), nil
	default:
		return "", exerrors.UsageErrorf("sort direction %q: expected %q or %q",
			direction[0], SortAsc, SortDesc)
	}
}

// MustSortParam is SortParam for a literal known good at author time.
func MustSortParam(field string, direction ...SortDirection) string {
	value, err := SortParam(field, direction...)
	if err != nil {
		panic(err)
	}
	return value
}

// SortParams builds the `order_by` and `direction` pair Item, Order, Usage, Purchase Order and
// Return Merchandise Authorisation take instead of `sort`.
func SortParams(field string, descending bool) (map[string]any, error) {
	if strings.TrimSpace(field) == "" {
		return nil, exerrors.UsageErrorf("sort field: expected a name like 'created_on', got %q", field)
	}
	direction := SortAsc
	if descending {
		direction = SortDesc
	}
	return map[string]any{
		"order_by":  strings.TrimSpace(field),
		"direction": string(direction),
	}, nil
}

// MustSortParams is SortParams for a literal known good at author time.
func MustSortParams(field string, descending bool) map[string]any {
	params, err := SortParams(field, descending)
	if err != nil {
		panic(err)
	}
	return params
}

// QueryPair is one normalised name and value, in the order it is sent.
type QueryPair struct {
	Name  string
	Value string
}

// QueryPairs is BuildQueryParams expanded the way the wire carries it: sorted by name, a list repeating its key.
func QueryPairs(params map[string]any) ([]QueryPair, error) {
	normalised, err := BuildQueryParams(params)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(normalised))
	for name := range normalised {
		names = append(names, name)
	}
	sort.Strings(names)
	var out []QueryPair
	for _, name := range names {
		for _, single := range ExpandValue(normalised[name]) {
			out = append(out, QueryPair{Name: name, Value: single})
		}
	}
	return out, nil
}

// EncodeQueryParams renders normalised pairs, percent-encoded.
func EncodeQueryParams(pairs []QueryPair) string {
	if len(pairs) == 0 {
		return ""
	}
	parts := make([]string, 0, len(pairs))
	for _, pair := range pairs {
		parts = append(parts, url.QueryEscape(pair.Name)+"="+url.QueryEscape(pair.Value))
	}
	return strings.Join(parts, "&")
}
