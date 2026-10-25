// Package portals is the v4 portals endpoints.
package portals

// The path templates for the portals endpoints; the client fills {api_version}.
const (
	PortalActivatePath       = "/api/{api_version}/portal/activate"
	PortalChangePasswordPath = "/api/{api_version}/portal/change_password"
	PortalLoginPath          = "/api/{api_version}/portal/login"
	PortalResetPath          = "/api/{api_version}/portal/reset"
	PortalResetLinkPath      = "/api/{api_version}/portal/reset_link"
	PortalValidatePath       = "/api/{api_version}/portal/validate"
	PortalChangePath         = "/api/{api_version}/portal/change"
)

// PortalKey is the envelope key one record arrives under.
const PortalKey = "portal"
