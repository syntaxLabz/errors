# syntaxLabz/errors Package Documentation

## Overview

The `syntaxLabz/errors` package provides a structured and standardized way to handle HTTP errors in Go applications. It offers predefined error codes, structured error messages, stack trace support, and utilities for creating detailed error responses that can be easily serialized to JSON for API responses.

## Table of Contents

1. [Installation](#installation)
2. [Package Structure](#package-structure)
3. [Core Features](#core-features)
4. [Error Types](#error-types)
5. [Usage Examples](#usage-examples)
6. [Best Practices](#best-practices)
7. [API Reference](#api-reference)
8. [Architecture Integration](#architecture-integration)
9. [Common Patterns](#common-patterns)
10. [When to Use](#when-to-use)

## Installation

```bash
go get github.com/syntaxLabz/errors
```

## Package Structure

```
syntaxLabz/errors/
├── pkg/
│   ├── codes/
│   │   ├── errorcodes.go      # Error code definitions
│   │   └── getters.go         # Error code getters
│   └── httperrors/
│       ├── validations.go     # Validation error functions
│       ├── details.go         # Predefined error templates
│       ├── errorResponse.go   # Error response struct with JSON serialization
│       └── new.go             # Core logic to create and manage errors
```

## Core Features

### 1. **Predefined Error Codes**
- HTTP status codes are automatically mapped to appropriate error responses
- Standardized error codes across your application
- Consistent error messaging

### 2. **Structured Error Messages**
- Detailed error information with field-specific details
- Helpful hints for debugging and fixing errors
- Timestamp inclusion for logging and tracking

### 3. **Stack Trace Support**
- Full stack traces for both main errors and detailed errors
- Useful for debugging and development environments
- Can be logged separately from user-facing responses

### 4. **JSON Serialization**
- Automatic JSON serialization for API responses
- Consistent error response format
- Easy integration with web frameworks

## Error Types

### Validation Errors

| Error Function | Description | HTTP Status |
|----------------|-------------|-------------|
| `MissingParameter(field string)` | Missing required parameter | 400 |
| `InvalidParameter(field string)` | Invalid parameter value | 400 |
| `LengthExceeded(field string)` | Parameter length exceeded | 400 |
| `OutOfRange(field string)` | Value out of allowed range | 400 |
| `HeaderValidationError(details...)` | Header validation failed | 400 |
| `InvalidJSON` | Invalid JSON format | 400 |

### Authentication Errors

| Error Function | Description | HTTP Status |
|----------------|-------------|-------------|
| `Unauthorized` | Invalid authentication | 401 |
| `Forbidden` | Insufficient permissions | 403 |
| `TokenExpired` | Token expired | 401 |
| `InvalidToken` | Malformed token | 401 |

### Resource Errors

| Error Function | Description | HTTP Status |
|----------------|-------------|-------------|
| `NotFound(resource string)` | Resource not found | 404 |
| `AlreadyExists(resource string)` | Resource already exists | 409 |
| `Conflict(resource string)` | Conflict with current state | 409 |

### Server Errors

| Error Function | Description | HTTP Status |
|----------------|-------------|-------------|
| `InternalServerError` | Internal server error | 500 |
| `ServiceUnavailable` | Service temporarily unavailable | 503 |

## Usage Examples

### Basic Error Handling

```go
package main

import (
    "log"
    "net/http"
    "github.com/syntaxLabz/errors/pkg/httperrors"
)

func main() {
    http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
        userID := r.Header.Get("userID")
        if userID == "" {
            err := httperrors.HeaderValidationError(httperrors.MissingHeader("userID"))
            
            // Log the stack trace for debugging
            log.Print(err.Stack)
            
            // Return structured error response
            statusCode, errResp := err.ErrorResponse()
            w.WriteHeader(statusCode)
            w.Write(errResp.ToJSON())
            return
        }
        
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Success"))
    })
    
    http.ListenAndServe(":8080", nil)
}
```

### Gin Framework Integration

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/syntaxLabz/errors/pkg/httperrors"
)

func main() {
    r := gin.Default()
    
    r.GET("/user", func(ctx *gin.Context) {
        email := ctx.Query("email")
        if email == "" {
            status, err := httperrors.MissingParameter("email").ErrorResponse()
            ctx.JSON(status, err)
            return
        }
        
        ctx.JSON(http.StatusOK, gin.H{"email": email})
    })
    
    r.Run()
}
```

### Layered Architecture Example

#### Handler Layer (Request Validation)
```go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/syntaxLabz/errors/pkg/httperrors"
)

func CreateUser(ctx *gin.Context) {
    var req map[string]interface{}
    if err := ctx.BindJSON(&req); err != nil {
        status, httpErr := httperrors.InvalidJSON.ErrorResponse()
        ctx.JSON(status, httpErr)
        return
    }
    
    if _, ok := req["email"]; !ok {
        status, httpErr := httperrors.MissingParameter("email").ErrorResponse()
        ctx.JSON(status, httpErr)
        return
    }
    
    // Call service layer
    if err := services.CreateUser(req["email"].(string)); err != nil {
        status, httpErr := err.ErrorResponse()
        ctx.JSON(status, httpErr)
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{"message": "User created"})
}
```

#### Service Layer (Business Logic)
```go
package services

import "github.com/syntaxLabz/errors/pkg/httperrors"

func CreateUser(email string) error {
    if email == "" {
        return httperrors.MissingParameter("email")
    }
    
    // Check if user already exists
    if err := store.GetUserByEmail(email); err != nil {
        return err
    }
    
    // Create user logic here
    return nil
}
```

#### Data Layer (Database Operations)
```go
package store

import "github.com/syntaxLabz/errors/pkg/httperrors"

func GetUserByEmail(email string) error {
    // Simulate database check
    if email == "exists@example.com" {
        return httperrors.AlreadyExists("user")
    }
    
    return nil
}
```

## Best Practices

### 1. **Error Propagation**
- Always propagate errors up the call stack
- Use appropriate error types for each layer
- Don't swallow errors silently

```go
// Good
func processRequest(data string) error {
    if data == "" {
        return httperrors.MissingParameter("data")
    }
    
    if err := validateData(data); err != nil {
        return err // Propagate the error
    }
    
    return nil
}

// Bad
func processRequest(data string) error {
    if data == "" {
        log.Println("Missing data") // Silent failure
        return nil
    }
    
    return nil
}
```

### 2. **Stack Trace Usage**
- Log stack traces for debugging but don't expose them to users
- Use stack traces in development and testing environments
- Consider log levels for production

```go
err := httperrors.InternalServerError
log.Printf("Internal error occurred: %s", err.Stack) // For debugging
status, errResp := err.ErrorResponse()
ctx.JSON(status, errResp) // Clean response to user
```

### 3. **Custom Error Details**
- Provide helpful hints in error messages
- Include relevant context information
- Make error messages actionable

```go
// Good - Provides context and hints
err := httperrors.MissingParameter("email")
err.AddHint("Ensure 'email' is included in the request body")

// Even better - Custom error with context
err := httperrors.New(
    httperrors.BAD_REQUEST,
    "User registration failed",
    httperrors.Detail{
        Field: "email",
        Error: "Email is required for user registration",
        Hint:  "Please provide a valid email address",
    },
)
```

### 4. **Validation Patterns**
- Validate early and return immediately
- Use specific validation functions
- Group related validations

```go
func ValidateUserRequest(req UserRequest) error {
    if req.Email == "" {
        return httperrors.MissingParameter("email")
    }
    
    if !isValidEmail(req.Email) {
        return httperrors.InvalidParameter("email")
    }
    
    if len(req.Password) < 8 {
        return httperrors.LengthExceeded("password")
    }
    
    return nil
}
```

## API Reference

### Error Response Format

All errors are returned in the following standardized JSON format:

```json
{
    "errors": {
        "code": "INVALID_PARAMETER",
        "message": "Invalid request parameter.",
        "details": [
            {
                "field": "email",
                "error": "Email is required",
                "hint": "Ensure 'email' is included."
            }
        ],
        "timestamp": "2025-03-02T10:00:00Z"
    }
}
```

### Core Methods

#### `ErrorResponse() (int, *ErrorResponse)`
Returns HTTP status code and error response object.

```go
err := httperrors.NotFound("user")
status, response := err.ErrorResponse()
// status: 404
// response: Contains structured error data
```

#### `ToJSON() []byte`
Converts error response to JSON bytes.

```go
err := httperrors.InvalidParameter("email")
_, response := err.ErrorResponse()
jsonData := response.ToJSON()
```

#### `Stack`
Provides stack trace information for debugging.

```go
err := httperrors.InternalServerError
log.Printf("Stack trace: %s", err.Stack)
```

## Architecture Integration

### Middleware Integration

```go
func ErrorMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        
        // Check if there were any errors
        if len(c.Errors) > 0 {
            err := c.Errors.Last()
            
            // Check if it's our custom error type
            if httpErr, ok := err.Err.(*httperrors.HTTPError); ok {
                status, response := httpErr.ErrorResponse()
                c.JSON(status, response)
                return
            }
            
            // Handle other error types
            status, response := httperrors.InternalServerError.ErrorResponse()
            c.JSON(status, response)
        }
    }
}
```

### Database Integration

```go
func handleDBError(err error) error {
    if err == sql.ErrNoRows {
        return httperrors.NotFound("resource")
    }
    
    if isDuplicateKeyError(err) {
        return httperrors.AlreadyExists("resource")
    }
    
    // For other database errors
    return httperrors.InternalServerError
}
```

## Common Patterns

### 1. **Input Validation Pattern**

```go
func validateAndProcess(req Request) error {
    // Validate required fields
    if req.Email == "" {
        return httperrors.MissingParameter("email")
    }
    
    if req.Password == "" {
        return httperrors.MissingParameter("password")
    }
    
    // Validate format
    if !isValidEmail(req.Email) {
        return httperrors.InvalidParameter("email")
    }
    
    // Process the request
    return processRequest(req)
}
```

### 2. **Resource Management Pattern**

```go
func getUserByID(id string) (*User, error) {
    if id == "" {
        return nil, httperrors.MissingParameter("id")
    }
    
    user, err := db.FindUserByID(id)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, httperrors.NotFound("user")
        }
        return nil, httperrors.InternalServerError
    }
    
    return user, nil
}
```

### 3. **Authentication Pattern**

```go
func authenticateUser(token string) error {
    if token == "" {
        return httperrors.MissingParameter("token")
    }
    
    if !isValidTokenFormat(token) {
        return httperrors.InvalidToken
    }
    
    if isTokenExpired(token) {
        return httperrors.TokenExpired
    }
    
    if !hasValidPermissions(token) {
        return httperrors.Forbidden
    }
    
    return nil
}
```

## When to Use

### ✅ **Use syntaxLabz/errors When:**

1. **Building HTTP APIs**
   - RESTful APIs that need consistent error responses
   - Web services with standard HTTP status codes
   - APIs that require detailed error information

2. **Structured Error Handling**
   - Applications requiring consistent error formats
   - Systems with multiple error types and categories
   - Applications needing detailed error context

3. **Development and Debugging**
   - Applications where stack traces are valuable for debugging
   - Systems requiring detailed error logging
   - Development environments needing rich error information

4. **Multi-layered Applications**
   - Applications with handler, service, and data layers
   - Systems requiring error propagation across layers
   - Applications with complex validation requirements

5. **Team Development**
   - Projects with multiple developers needing consistent error handling
   - Applications requiring standardized error responses
   - Systems with strict error handling requirements

### ❌ **Don't Use When:**

1. **Simple Applications**
   - Basic command-line tools
   - Simple scripts with minimal error handling
   - Applications not requiring structured errors

2. **Non-HTTP Applications**
   - Desktop applications
   - System services not serving HTTP
   - Applications not requiring HTTP status codes

3. **Performance-Critical Applications**
   - Applications where error handling overhead is a concern
   - Systems requiring minimal memory footprint
   - High-performance computing applications

4. **Legacy Systems**
   - Applications with existing error handling patterns
   - Systems that can't be easily refactored
   - Applications with conflicting error handling libraries

## Conclusion

The `syntaxLabz/errors` package provides a comprehensive solution for HTTP error handling in Go applications. It standardizes error responses, provides rich debugging information, and integrates well with popular web frameworks. By following the patterns and best practices outlined in this documentation, you can build robust, maintainable applications with consistent error handling.

The package is particularly valuable for API development, where consistent error responses are crucial for client integration and debugging. The structured approach to error handling makes it easier to maintain and extend your application's error handling capabilities over time.
