package httperrors

import (
	"fmt"

	"github.com/syntaxLabz/errors/pkg/codes"
)

// BodyValidationError creates an API error for validation issues
func BodyValidationError(details ...Details) *Error {
	return NewErrorWithDetails(codes.BadRequest, "Body validation failed.", details)
}

// HeaderValidationError creates an API error for validation issues
func HeaderValidationError(details ...Details) *Error {
	return NewErrorWithDetails(codes.BadRequest, "Header validation failed.", details)
}

// RequestValidationError creates an API error for validation issues
func RequestValidationError(details ...Details) *Error {
	return NewErrorWithDetails(codes.BadRequest, "Request validation failed.", details)
}

// NewAuthError creates an API error for authentication issues
func NewAuthError() *Error {
	return New("codes.CodeUnauthorized", "Authentication failed.")
}

// NewNotFoundError creates an API error for missing resources
func NewNotFoundError(resource, value string) *Error {
	return New(codes.NotFound, fmt.Sprintf("no %s component found for %s.", resource, value))
}

// NewConflictError creates an API error for conflicting resources
func NewConflictError(details ...Details) *Error {
	return NewErrorWithDetails(codes.Conflict, "Entity already exists.", details)
}

// NewServerError creates an API error for internal issues
func NewServerError() *Error {
	return New(codes.InternalServerError, "An unexpected internal error occurred.")
}

// NewDBError creates an API error for internal issues
func NewDBError() *Error {
	return New(codes.InternalServerError, "An unexpected database error occurred.")
}
