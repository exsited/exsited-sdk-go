package attributeattachments

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// AttributeAttachments is the v4 attribute attachments endpoints, bound to one client.
type AttributeAttachments struct {
	requester exhttp.Requester
}

// GetAttributeAttachment sends GET /api/{api_version}/attribute_attachment/{file_uuid}; it has no typed view, so read ApiResponse.Map.
func (a *AttributeAttachments) GetAttributeAttachment(ctx context.Context, fileUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	if err := requireValue(fileUUID, "file_uuid"); err != nil {
		return nil, err
	}
	format, err := oneResponseFormat(responseFormat)
	if err != nil {
		return nil, err
	}
	return a.requester.SendGetRequest(ctx, AttributeAttachmentPath, &exhttp.RequestOptions{
		PathParams:   map[string]any{"file_uuid": fileUUID},
		RaiseOnError: raiseOnError, ResponseFormat: format})
}

// ResourceName is this resource's registry key, what ListResourceNames returns, snake_case.
const ResourceName = "attribute_attachments"

// oneResponseFormat reads the optional trailing responseFormat, refusing more than one.
func oneResponseFormat(responseFormat []string) (string, error) {
	switch len(responseFormat) {
	case 0:
		return "", nil
	case 1:
		return responseFormat[0], nil
	default:
		return "", exerrors.NewUsageError(fmt.Sprintf(
			"responseFormat: pass one format or none, got %d; the trailing argument is optional, "+
				"and omitting it defers to the client", len(responseFormat)))
	}
}

// isNilValue reports whether a value is a typed nil inside a non-nil any; a zero-valued struct is not.
func isNilValue(value any) bool {
	reflected := reflect.ValueOf(value)
	switch reflected.Kind() {
	case reflect.Map, reflect.Slice, reflect.Pointer, reflect.Interface, reflect.Func:
		return reflected.IsNil()
	}
	return false
}

// requireValue refuses a path value that is nil, empty or only whitespace, before any I/O.
func requireValue(value any, what string) error {
	text, isText := value.(string)
	if value == nil || isNilValue(value) || isText && strings.TrimSpace(text) == "" {
		return exerrors.UsageErrorf("%s is required: the request cannot be built without it", what)
	}
	return nil
}

// New binds the attribute attachments endpoints to any requester (the ExsitedClient).
func New(requester exhttp.Requester) *AttributeAttachments {
	return &AttributeAttachments{requester: requester}
}

// Requester returns the requester this resource was bound to.
func (a *AttributeAttachments) Requester() exhttp.Requester { return a.requester }
