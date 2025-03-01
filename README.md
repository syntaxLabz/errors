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

## Contributions
Feel free to fork the repository, make changes, and create pull requests! We welcome contributions that improve functionality or fix bugs.





