package httperrors

import (
	"encoding/json"
	"errors"
	"runtime/debug"
	"time"
)

// Details represents individual error details
type Details struct {
	Field string `json:"field,omitempty"`
	Error string `json:"error"`
	Hint  string `json:"hint,omitempty"`
	Stack string `json:"-"`
}

// Error represents the structured error format
type Error struct {
	Code      string    `json:"code"`
	Message   string    `json:"message"`
	Details   []Details `json:"details,omitempty"`
	Timestamp string    `json:"timestamp"`
	Err       error     `json:"-"`
	Stack     string    `json:"-"`
}

// Error implements the error interface
func (e *Error) Error() string {
	if len(e.Details) > 0 {
		return e.Details[0].Error
	}

	return e.Message
}

// New creates a basic error without details
func New(code, message string) *Error {
	return &Error{
		Code:      code,
		Message:   message,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Stack:     string(debug.Stack()),
	}
}

// NewErrorWithDetails creates an error with multiple details
func NewErrorWithDetails(code, message string, details []Details) *Error {
	return &Error{
		Code:      code,
		Message:   message,
		Details:   details,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Stack:     string(debug.Stack()),
	}
}

// AddDetails appends multiple details to an existing error
func (e *Error) AddDetails(details []Details) {
	for i := range details {
		details[i].Stack = string(debug.Stack())
	}

	e.Details = append(e.Details, details...)
}

// AddDetail appends a single detail to an existing error
func (e *Error) AddDetail(detail Details) {
	detail.Stack = string(debug.Stack())

	e.Details = append(e.Details, detail)
}

// ToJSON returns the error as a JSON-formatted byte slice
func (e *Error) ToJSON() ([]byte, error) {
	return json.Marshal(ErrorResponse{Errors: *e})
}

func AppendDetails(err error, details ...Details) error {
	var httpErr *Error
	if errors.As(err, &httpErr) {
		httpErr.AddDetails(details)
		return httpErr
	}

	// Convert generic error into structured httperrors.Error
	return &Error{
		Code:      "UNKNOWN_ERROR",
		Message:   err.Error(),
		Details:   details,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Err:       err,
	}
}

func ExtractError(err error) *Error {
	var httpErr *Error
	if errors.As(err, &httpErr) {
		return httpErr
	}

	// Convert a normal error into an httperrors.Error for uniformity
	return &Error{
		Code:      "INTERNAL_ERROR",
		Message:   err.Error(),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Stack:     string(debug.Stack()),
		Err:       err,
	}
}

// GetStackTrace returns the stack trace for internal logging
func (e *Error) GetStackTrace() string {
	return e.Stack
}

// GetDetailsStackTrace returns the stack trace for internal logging
func (e *Error) GetDetailsStackTrace(field string) string {
	for i := range e.Details {
		if e.Details[i].Field == field {
			return e.Stack
		}
	}

	return ""
}
