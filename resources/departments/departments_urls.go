// Package departments is the v4 departments endpoints.
package departments

// The path templates for the departments endpoints; the client fills {api_version}.
const (
	SettingsDepartmentsPath = "/api/{api_version}/settings/departments"
	SettingsDepartmentPath  = "/api/{api_version}/settings/departments/{department_id}"
)

// SettingsDepartmentsKey is the key a listing's records arrive under, the v4 list envelope's data.
const SettingsDepartmentsKey = "data"

// DepartmentKey is the envelope key one record arrives under.
const DepartmentKey = "department"
