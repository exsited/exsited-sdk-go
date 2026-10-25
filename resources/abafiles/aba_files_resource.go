package abafiles

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the AbaFiles method of the same name,
// which is the primary surface.

// GenerateABAFile is the module form of AbaFiles.GenerateABAFile.
func GenerateABAFile(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GenerateABAFile(ctx, payload, raiseOnError, responseFormat...)
}
