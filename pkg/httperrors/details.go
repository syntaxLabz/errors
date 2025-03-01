package httperrors

import (
	"fmt"
	"runtime/debug"
)

var (
	// Validation Errors
	MissingParameter = func(field string) Details {
		return Details{
			Field: field,
			Error: fmt.Sprintf("Parameter %s is required.", field),
			Hint:  fmt.Sprintf("Ensure %s is included in the request.", field),
		}
	}

	InvalidParameter = func(field string) Details {
		return Details{
			Field: field,
			Error: fmt.Sprintf("Parameter %s is invalid.", field),
			Hint:  fmt.Sprintf("Check the value of %s for correctness.", field),
		}
	}

	InvalidFormat = func(field string) Details {
		return Details{
			Field: field,
			Error: fmt.Sprintf("Parameter %s has an invalid format.", field),
			Hint:  fmt.Sprintf("Ensure %s follows the expected format.", field),
		}
	}

	LengthExceeded = func(field string, max int) Details {
		return Details{
			Field: field,
			Error: fmt.Sprintf("Parameter %s exceeds maximum length of %d.", field, max),
			Hint:  fmt.Sprintf("Ensure %s does not exceed %d characters.", field, max),
		}
	}

	OutOfRange = func(field string, min, max int) Details {
		return Details{
			Field: field,
			Error: fmt.Sprintf("Parameter %s must be between %d and %d.", field, min, max),
			Hint:  fmt.Sprintf("Provide a value between %d and %d for %s.", min, max, field),
		}
	}

	InvalidEnumValue = func(field string, allowedValues []string) Details {
		return Details{
			Field: field,
			Error: fmt.Sprintf("Parameter %s must be one of [%s].", field, formatAllowedValues(allowedValues)),
			Hint:  fmt.Sprintf("Allowed values: %s.", formatAllowedValues(allowedValues)),
			Stack:     string(debug.Stack()),
		}
	}

	// Header Errors
	MissingHeader = func(header string) Details {
		return Details{
			Field: header,
			Error: fmt.Sprintf("Header %s is required.", header),
			Hint:  fmt.Sprintf("Ensure %s is included in the request headers.", header),
			Stack:     string(debug.Stack()),
		}
	}

	InvalidHeader = func(header string) Details {
		return Details{
			Field: header,
			Error: fmt.Sprintf("Header %s is invalid.", header),
			Hint:  fmt.Sprintf("Verify the value of header %s.", header),
			Stack:     string(debug.Stack()),
		}
	}

	MissingCorrelationID = Details{
		Field: "correlation_id",
		Error: "Missing Correlation ID in request headers.",
		Hint:  "Include 'correlation_id' in the request headers with a valid UUID.",
		Stack:     string(debug.Stack()),
	}

	// Authentication & Authorization Errors
	Unauthorized = Details{
		Field: "auth",
		Error: "Authentication failed. Invalid credentials.",
		Hint:  "Ensure correct credentials are provided.",
		Stack:     string(debug.Stack()),
	}

	Forbidden = Details{
		Field: "auth",
		Error: "Access denied. You do not have permission to perform this action.",
		Hint:  "Check user roles and permissions.",
		Stack:     string(debug.Stack()),
	}

    TokenExpired = Details{
		Field: "auth",
		Error: "Authentication token has expired.",
		Hint:  "Request a new token and retry.",
		Stack:     string(debug.Stack()),
	}

    InvalidToken = Details{
		Field: "auth",
		Error: "Invalid authentication token.",
		Hint:  "Ensure the token is valid and not tampered with.",
		Stack:     string(debug.Stack()),
	}

	// Resource Errors
	NotFound = func(resource string) Details {
		return Details{
			Field: resource,
			Error: fmt.Sprintf("%s not found.", resource),
			Hint:  fmt.Sprintf("Check if %s exists before making the request.", resource),
			Stack:     string(debug.Stack()),
		}
	}

	AlreadyExists = func(resource string) Details {
		return Details{
			Field: resource,
			Error: fmt.Sprintf("%s already exists.", resource),
			Hint:  fmt.Sprintf("Ensure %s does not already exist before attempting to create it.", resource),
			Stack:     string(debug.Stack()),
		}
	}

	Conflict = func(resource string) Details {
		return Details{
			Field: resource,
			Error: fmt.Sprintf("A conflict occurred with %s.", resource),
			Hint:  fmt.Sprintf("Resolve conflicts with %s before retrying.", resource),
			Stack:     string(debug.Stack()),
		}
	}

	// Server Errors
	Internal = Details{
		Field: "server",
		Error: "An unexpected internal error occurred.",
		Stack:     string(debug.Stack()),
	}

    Database = Details{
		Field: "database",
		Error: "A database error occurred.",
		Stack:     string(debug.Stack()),
	}

	ServiceDown = Details{
		Field: "service",
		Error: "Service is temporarily unavailable.",
		Stack:     string(debug.Stack()),
	}

	RateLimitExceeded = Details{
		Field: "rate_limit",
		Error: "Rate limit exceeded. Please try again later.",
		Stack:     string(debug.Stack()),
	}

	// Request & Payload Errors
	InvalidJSON = Details{
		Field: "request",
		Error: "Invalid JSON payload.",
		Hint:  "Ensure the request body is a valid JSON object.",
		Stack:     string(debug.Stack()),
	}

	RequestTimeout = Details{
		Field: "request",
		Error: "Request timed out. Please try again.",
		Hint:  "Ensure the server is reachable and retry the request.",
		Stack:     string(debug.Stack()),
	}

	UnsupportedMediaType = Details{
		Field: "content_type",
		Error: "Unsupported media type. Please check the request format.",
		Stack:     string(debug.Stack()),
		Hint:  "Use 'application/json' as the Content-Type.",
	}

	MissingQueryParam = func(param string) Details {
		return Details{
			Field: param,
			Error: fmt.Sprintf("Query parameter %s is required.", param),
			Hint:  fmt.Sprintf("Include %s in the request URL.", param),
			Stack:     string(debug.Stack()),
		}
	}

	InvalidQueryParam = func(param string) Details {
		return Details{
			Field: param,
			Error: fmt.Sprintf("Query parameter %s is invalid.", param),
			Hint:  fmt.Sprintf("Provide a valid value for %s.", param),
			Stack:     string(debug.Stack()),
		}
	}

	PaginationLimitExceeded = Details{
		Field: "pagination",
		Error: "Pagination limit exceeded. Reduce the page size.",
		Hint:  "Use a smaller page size in the request.",
		Stack:     string(debug.Stack()),
	}

	RequiredTogether = func(fields ...string) Details {
		return Details{
			Field: fmt.Sprintf("%v", fields),
			Error: fmt.Sprintf("Fields %v are required together.", fields),
			Hint:  fmt.Sprintf("Ensure all of %v are included in the request.", fields),
			Stack:     string(debug.Stack()),
		}
	}
)

// Helper function to format allowed enum values
func formatAllowedValues(values []string) string {
	return fmt.Sprintf("%s", values)
}
