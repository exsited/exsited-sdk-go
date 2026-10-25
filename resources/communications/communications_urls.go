// Package communications is the v4 communications endpoints.
package communications

// The path templates for the communications endpoints; the client fills {api_version}.
const (
	CommunicationsSendEmailPath = "/api/{api_version}/communications/send_email"
)

// EmailKey is the envelope key one record arrives under.
const EmailKey = "email"
