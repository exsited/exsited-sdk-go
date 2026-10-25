package attributeattachments

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the AttributeAttachments method of the same name,
// which is the primary surface.

// GetAttributeAttachment is the module form of AttributeAttachments.GetAttributeAttachment.
func GetAttributeAttachment(requester exhttp.Requester, ctx context.Context, fileUUID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetAttributeAttachment(ctx, fileUUID, raiseOnError, responseFormat...)
}
