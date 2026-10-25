package abafiles

// This family has no typed views: no successful response body is known to read field types from.
//
// Every endpoint here answers ApiResponse.Map and leaves ApiResponse.DTO() nil.
// A view declared from the endpoint names alone would be a guess, and a wrong type in Go
// does not panic, it leaves the field nil and loses the value in silence.
