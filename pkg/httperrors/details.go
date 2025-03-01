package httperrors

import "fmt"

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
		}
	}

	// Header Errors
	MissingHeader = func(header string) Details {
		return Details{
			Field: header,
			Error: fmt.Sprintf("Header %s is required.", header),
			Hint:  fmt.Sprintf("Ensure %s is included in the request headers.", header),
		}
	}

	InvalidHeader = func(header string) Details {
		return Details{
			Field: header,
			Error: fmt.Sprintf("Header %s is invalid.", header),
			Hint:  fmt.Sprintf("Verify the value of header %s.", header),
		}
	}

	MissingCorrelationID = Details{
		Field: "correlation_id",
		Error: "Missing Correlation ID in request headers.",
		Hint:  "Include 'correlation_id' in the request headers with a valid UUID.",
	}

	// Authentication & Authorization Errors
	Unauthorized = Details{
		Field: "auth",
		Error: "Authentication failed. Invalid credentials.",
		Hint:  "Ensure correct credentials are provided.",
	}

	Forbidden = Details{
		Field: "auth",
		Error: "Access denied. You do not have permission to perform this action.",
		Hint:  "Check user roles and permissions.",
	}

    TokenExpired = Details{
		Field: "auth",
		Error: "Authentication token has expired.",
		Hint:  "Request a new token and retry.",
	}

    InvalidToken = Details{
		Field: "auth",
		Error: "Invalid authentication token.",
		Hint:  "Ensure the token is valid and not tampered with.",
	}

	// Resource Errors
	NotFound = func(resource string) Details {
		return Details{
			Field: resource,
			Error: fmt.Sprintf("%s not found.", resource),
			Hint:  fmt.Sprintf("Check if %s exists before making the request.", resource),
		}
	}

	AlreadyExists = func(resource string) Details {
		return Details{
			Field: resource,
			Error: fmt.Sprintf("%s already exists.", resource),
			Hint:  fmt.Sprintf("Ensure %s does not already exist before attempting to create it.", resource),
		}
	}

	Conflict = func(resource string) Details {
		return Details{
			Field: resource,
			Error: fmt.Sprintf("A conflict occurred with %s.", resource),
			Hint:  fmt.Sprintf("Resolve conflicts with %s before retrying.", resource),
		}
	}

	// Server Errors
	Internal = Details{
		Field: "server",
		Error: "An unexpected internal error occurred.",
	}

    Database = Details{
		Field: "database",
		Error: "A database error occurred.",
	}

	ServiceDown = Details{
		Field: "service",
		Error: "Service is temporarily unavailable.",
	}

	RateLimitExceeded = Details{
		Field: "rate_limit",
		Error: "Rate limit exceeded. Please try again later.",
	}

	// Request & Payload Errors
	InvalidJSON = Details{
		Field: "request",
		Error: "Invalid JSON payload.",
		Hint:  "Ensure the request body is a valid JSON object.",
	}

	RequestTimeout = Details{
		Field: "request",
		Error: "Request timed out. Please try again.",
		Hint:  "Ensure the server is reachable and retry the request.",
	}

	UnsupportedMediaType = Details{
		Field: "content_type",
		Error: "Unsupported media type. Please check the request format.",
		Hint:  "Use 'application/json' as the Content-Type.",
	}

	MissingQueryParam = func(param string) Details {
		return Details{
			Field: param,
			Error: fmt.Sprintf("Query parameter %s is required.", param),
			Hint:  fmt.Sprintf("Include %s in the request URL.", param),
		}
	}

	InvalidQueryParam = func(param string) Details {
		return Details{
			Field: param,
			Error: fmt.Sprintf("Query parameter %s is invalid.", param),
			Hint:  fmt.Sprintf("Provide a valid value for %s.", param),
		}
	}

	PaginationLimitExceeded = Details{
		Field: "pagination",
		Error: "Pagination limit exceeded. Reduce the page size.",
		Hint:  "Use a smaller page size in the request.",
	}

	RequiredTogether = func(fields ...string) Details {
		return Details{
			Field: fmt.Sprintf("%v", fields),
			Error: fmt.Sprintf("Fields %v are required together.", fields),
			Hint:  fmt.Sprintf("Ensure all of %v are included in the request.", fields),
		}
	}
)

// Helper function to format allowed enum values
func formatAllowedValues(values []string) string {
	return fmt.Sprintf("%s", values)
}
