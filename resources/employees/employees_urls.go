// Package employees is the v4 employees endpoints.
package employees

// The path templates for the employees endpoints; the client fills {api_version}.
const (
	EmployeesPath = "/api/{api_version}/employees"
	EmployeePath  = "/api/{api_version}/employees/{employee_id}"
)

// EmployeesKey is the key a listing's records arrive under, the v4 list envelope's data.
const EmployeesKey = "data"

// EmployeeKey is the envelope key one record arrives under.
const EmployeeKey = "employee"
