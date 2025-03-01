package httperrors

import "github.com/syntaxLabz/errors/pkg/codes"

// ErrorResponse is the standard response format for errors
type ErrorResponse struct {
	Errors Error `json:"errors"`
}

func (e *ErrorResponse) Error() string {
	return e.Errors.Error()
}

func (e *ErrorResponse) ToJSON() []byte {
	json, _ := e.Errors.ToJSON()
	return json
}

func (e *Error) ErrorResponse() (statusCode int, err *ErrorResponse) {
	return codes.RESTCode(e.Code), &ErrorResponse{Errors: *e}
}
