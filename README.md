# errors: Advanced Error pkg for your Go Applications 

## Overview

`errors` package provides a structured way to handle HTTP errors in a Go application. It standardizes error responses, maps HTTP status codes, and offers utilities for creating detailed error messages and provide stack for each error for easy debugging.

## Features

- Predefined error codes mapped to HTTP responses
- Structured error messages with details
- Utility functions for creating validation, authentication, and server errors
- JSON serialization support for API responses
- Stack support for both main error and details.

## Installation:

To install the module, use the following command:

```
go get github.com/syntaxLabz/errors
```

## Error Response:

``` {
  "errors": {
    "code": "BAD_REQUEST",
    "message": "Header validation failed.",
    "details": [
      {
        "field": "userID",
        "error": "Parameter userID is required.",
        "hint": "Ensure userID is included in the request."
      }
    ]
  }
}

```
## Example Usage

```
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

			log.Print(err.Stack)

			statusCode, errResp := err.ErrorResponse()
			w.WriteHeader(statusCode)
			w.Write(errResp.ToJSON())
		}

		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", nil)
}

```

# `httperrors` Package Documentation

\
---

## Folder Structure

```plaintext

codes/
│-- errorcodes.go
├─ getters.go       
httperrors/
│-- validations.go
├─ details.go        # Predefined error templates
├─ errorResponse.go  # Error response struct with JSON serialization
└─ new.go           # Core logic to create and manage errors
```

---

## Error Types

### 1. Validation Errors

| Error Name         | Description                |
| ------------------ | -------------------------- |
| `MissingParameter` | Missing required parameter |
| `InvalidParameter` | Invalid parameter value    |
| `LengthExceeded`   | Exceeded parameter length  |
| `OutOfRange`       | Value out of allowed range |

#### Example

```go
err := httperrors.MissingParameter("email")
fmt.Println(err.Error) // Output: Parameter email is required.
```

---

### 2. Authentication & Authorization Errors

| Error Name     | Description              |
| -------------- | ------------------------ |
| `Unauthorized` | Invalid authentication   |
| `Forbidden`    | Insufficient permissions |
| `TokenExpired` | Token expired            |
| `InvalidToken` | Malformed token          |

#### Example

```go
err := httperrors.Unauthorized
fmt.Println(err.Error) // Output: Unauthorized request.
```

---

### 3. Resource Errors

| Error Name      | Description                 |
| --------------- | --------------------------- |
| `NotFound`      | Resource not found          |
| `AlreadyExists` | Resource already exists     |
| `Conflict`      | Conflict with current state |

#### Example

```go
err := httperrors.NotFound("user")
fmt.Println(err.Error) // Output: User not found.
```

---

## Error Response Format

All errors are returned in the following JSON format:

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

---

## Layered Architecture Example

### Folder Structure

```plaintext
project/
│
├─ handlers/
│   └─ user.go       # Request Validation
├─ services/
│   └─ user.go       # Business Logic
└─ store/
    └─ user.go       # Database Queries
```

### Handler Layer

```go
package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/syntaxLabz/httperrors"
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

	ctx.JSON(http.StatusOK, gin.H{"message": "User created"})
}
```

---

### Service Layer

```go
package services

import "github.com/syntaxLabz/httperrors"

func CreateUser(email string) error {
	if email == "" {
		return httperrors.MissingParameter("email")
	}
	return nil
}
```

---

### Store Layer

```go
package store

import "github.com/syntaxLabz/httperrors"

func GetUserByEmail(email string) error {
	if email == "exists@example.com" {
		return httperrors.AlreadyExists("user")
	}
	return nil
}
```

---

## Normal Architecture Example

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/syntaxLabz/httperrors"
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

---

## Best Practices

- Use **httperrors.New()** for custom error codes
- Include **hint** fields for better debugging
- Use **stack traces** only for internal logs
- Wrap external errors inside `httperrors`

---



## Contributions
Feel free to fork the repository, make changes, and create pull requests! We welcome contributions that improve functionality or fix bugs.





