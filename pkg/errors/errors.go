package errors

import "fmt"

type ResourceNotFoundError struct {
	Err             string
	Message         string
	Code            int
	ResourceType    string
	IdentifierType  string
	IdentifierValue string
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("Resource not found for %s with %s = %s", e.ResourceType, e.IdentifierType, e.IdentifierValue)
}
