package giftcertificates

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the GiftCertificates method of the same name,
// which is the primary surface.

// GetGiftCertificates is the module form of GiftCertificates.GetGiftCertificates.
func GetGiftCertificates(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetGiftCertificates(ctx, queryParams, raiseOnError, responseFormat...)
}

// CreateGiftCertificate is the module form of GiftCertificates.CreateGiftCertificate.
func CreateGiftCertificate(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateGiftCertificate(ctx, payload, raiseOnError, responseFormat...)
}

// DeleteGiftCertificate is the module form of GiftCertificates.DeleteGiftCertificate.
func DeleteGiftCertificate(requester exhttp.Requester, ctx context.Context, gcID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteGiftCertificate(ctx, gcID, raiseOnError, responseFormat...)
}

// GetGiftCertificate is the module form of GiftCertificates.GetGiftCertificate.
func GetGiftCertificate(requester exhttp.Requester, ctx context.Context, gcID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetGiftCertificate(ctx, gcID, queryParams, raiseOnError, responseFormat...)
}

// UpdateGiftCertificate is the module form of GiftCertificates.UpdateGiftCertificate.
func UpdateGiftCertificate(requester exhttp.Requester, ctx context.Context, gcID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateGiftCertificate(ctx, gcID, payload, raiseOnError, responseFormat...)
}

// GetGiftCertificateAllocations is the module form of GiftCertificates.GetGiftCertificateAllocations.
func GetGiftCertificateAllocations(requester exhttp.Requester, ctx context.Context, gcID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetGiftCertificateAllocations(ctx, gcID, queryParams, raiseOnError, responseFormat...)
}

// GetGiftCertificateTransactions is the module form of GiftCertificates.GetGiftCertificateTransactions.
func GetGiftCertificateTransactions(requester exhttp.Requester, ctx context.Context, gcID any, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetGiftCertificateTransactions(ctx, gcID, queryParams, raiseOnError, responseFormat...)
}

// AmendGiftCertificate is the module form of GiftCertificates.AmendGiftCertificate.
func AmendGiftCertificate(requester exhttp.Requester, ctx context.Context, gcID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).AmendGiftCertificate(ctx, gcID, payload, raiseOnError, responseFormat...)
}

// AllocateGiftCertificate is the module form of GiftCertificates.AllocateGiftCertificate.
func AllocateGiftCertificate(requester exhttp.Requester, ctx context.Context, gcID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).AllocateGiftCertificate(ctx, gcID, payload, raiseOnError, responseFormat...)
}

// DeallocateGiftCertificate is the module form of GiftCertificates.DeallocateGiftCertificate.
func DeallocateGiftCertificate(requester exhttp.Requester, ctx context.Context, gcID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeallocateGiftCertificate(ctx, gcID, payload, raiseOnError, responseFormat...)
}

// GetGiftCertificateAllocation is the module form of GiftCertificates.GetGiftCertificateAllocation.
func GetGiftCertificateAllocation(requester exhttp.Requester, ctx context.Context, gcID any, allocationID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetGiftCertificateAllocation(ctx, gcID, allocationID, raiseOnError, responseFormat...)
}

// GetGiftCertificateTransaction is the module form of GiftCertificates.GetGiftCertificateTransaction.
func GetGiftCertificateTransaction(requester exhttp.Requester, ctx context.Context, gcID any, transactionID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetGiftCertificateTransaction(ctx, gcID, transactionID, raiseOnError, responseFormat...)
}
