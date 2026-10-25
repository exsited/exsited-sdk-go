package actionscripts

import (
	"context"

	exhttp "github.com/exsited/exsited-sdk-go/v4/http"
)

// The module form: each function binds the requester and calls the ActionScripts method of the same name,
// which is the primary surface.

// GetActionScripts is the module form of ActionScripts.GetActionScripts.
func GetActionScripts(requester exhttp.Requester, ctx context.Context, queryParams map[string]any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetActionScripts(ctx, queryParams, raiseOnError, responseFormat...)
}

// GetActionScript is the module form of ActionScripts.GetActionScript.
func GetActionScript(requester exhttp.Requester, ctx context.Context, actionScriptID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).GetActionScript(ctx, actionScriptID, raiseOnError, responseFormat...)
}

// CreateActionScript is the module form of ActionScripts.CreateActionScript.
func CreateActionScript(requester exhttp.Requester, ctx context.Context, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).CreateActionScript(ctx, payload, raiseOnError, responseFormat...)
}

// UpdateActionScript is the module form of ActionScripts.UpdateActionScript.
func UpdateActionScript(requester exhttp.Requester, ctx context.Context, actionScriptID any, payload any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).UpdateActionScript(ctx, actionScriptID, payload, raiseOnError, responseFormat...)
}

// DeleteActionScript is the module form of ActionScripts.DeleteActionScript.
func DeleteActionScript(requester exhttp.Requester, ctx context.Context, actionScriptID any, raiseOnError *bool, responseFormat ...string) (*exhttp.ApiResponse, error) {
	return New(requester).DeleteActionScript(ctx, actionScriptID, raiseOnError, responseFormat...)
}
