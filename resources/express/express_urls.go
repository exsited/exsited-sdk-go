// Package express is the v4 express endpoints.
package express

// The path templates for the express endpoints; the client fills {api_version}.
const (
	ExpressPath = "/api/{api_version}/express"
)

// ExpressKey is the envelope key one record arrives under.
const ExpressKey = "express"
