package openapi

import (
	"github.com/AlexeyKluev/ogen/internal/location"
	"github.com/AlexeyKluev/ogen/jsonschema"
)

// Example is an OpenAPI Example.
type Example struct {
	Ref Ref

	Summary       string
	Description   string
	Value         jsonschema.Example
	ExternalValue string

	location.Pointer `json:"-" yaml:"-"`
}
