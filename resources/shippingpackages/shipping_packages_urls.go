// Package shippingpackages is the v4 shipping packages endpoints.
package shippingpackages

// The path templates for the shipping packages endpoints; the client fills {api_version}.
const (
	SettingsShippingPackagesPath = "/api/{api_version}/settings/shipping_packages"
	SettingsShippingPackagePath  = "/api/{api_version}/settings/shipping_packages/{shipping_package_id}"
)

// SettingsShippingPackagesKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsShippingPackagesKey = "data"

// ShippingPackageKey is the envelope key one record arrives under.
const ShippingPackageKey = "shipping_package"
