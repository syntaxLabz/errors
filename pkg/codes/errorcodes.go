package codes

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const (
	OK                      = "OK"
	Created                 = "CREATED"
	Accepted                = "ACCEPTED"
	NoContent               = "NO_CONTENT"
	MovedPermanently        = "MOVED_PERMANENTLY"
	Found                   = "FOUND"
	BadRequest              = "BAD_REQUEST"
	Unauthorized            = "UNAUTHORIZED"
	Forbidden               = "FORBIDDEN"
	NotFound                = "NOT_FOUND"
	MethodNotAllowed        = "METHOD_NOT_ALLOWED"
	NotAcceptable           = "NOT_ACCEPTABLE"
	RequestTimeout          = "REQUEST_TIMEOUT"
	Conflict                  = "CONFLICT"
	Gone                    = "GONE"
	LengthRequired          = "LENGTH_REQUIRED"
	PreconditionFailed      = "PRECONDITION_FAILED"
	PayloadTooLarge         = "PAYLOAD_TOO_LARGE"
	URITooLong              = "URI_TOO_LONG"
	UnsupportedMediaType    = "UNSUPPORTED_MEDIA_TYPE"
	RangeNotSatisfiable     = "RANGE_NOT_SATISFIABLE"
	ExpectationFailed       = "EXPECTATION_FAILED"
	ImATeapot               = "IM_A_TEAPOT"
	UpgradeRequired         = "UPGRADE_REQUIRED"
	PreconditionRequired    = "PRECONDITION_REQUIRED"
	TooManyRequests         = "TOO_MANY_REQUESTS"
	InternalServerError     = "INTERNAL_SERVER_ERROR"
	NotImplemented          = "NOT_IMPLEMENTED"
	BadGateway              = "BAD_GATEWAY"
	ServiceUnavailable      = "SERVICE_UNAVAILABLE"
	GatewayTimeout          = "GATEWAY_TIMEOUT"
	HTTPVersionNotSupported = "HTTP_VERSION_NOT_SUPPORTED"
)

// RESTCodeMap maps error codes to HTTP status codes.
var restCodeMap = map[string]int{
	OK:                      http.StatusOK,
	Created:                 http.StatusCreated,
	Accepted:                http.StatusAccepted,
	NoContent:               http.StatusNoContent,
	MovedPermanently:        http.StatusMovedPermanently,
	Found:                   http.StatusFound,
	BadRequest:              http.StatusBadRequest,
	Unauthorized:            http.StatusUnauthorized,
	Forbidden:               http.StatusForbidden,
	NotFound:                http.StatusNotFound,
	MethodNotAllowed:        http.StatusMethodNotAllowed,
	NotAcceptable:           http.StatusNotAcceptable,
	RequestTimeout:          http.StatusRequestTimeout,
	Conflict:                http.StatusConflict,
	Gone:                    http.StatusGone,
	LengthRequired:          http.StatusLengthRequired,
	PreconditionFailed:      http.StatusPreconditionFailed,
	PayloadTooLarge:         http.StatusRequestEntityTooLarge,
	URITooLong:              http.StatusRequestURITooLong,
	UnsupportedMediaType:    http.StatusUnsupportedMediaType,
	RangeNotSatisfiable:     http.StatusRequestedRangeNotSatisfiable,
	ExpectationFailed:       http.StatusExpectationFailed,
	ImATeapot:               http.StatusTeapot,
	UpgradeRequired:         http.StatusUpgradeRequired,
	PreconditionRequired:    http.StatusPreconditionRequired,
	TooManyRequests:         http.StatusTooManyRequests,
	InternalServerError:     http.StatusInternalServerError,
	NotImplemented:          http.StatusNotImplemented,
	BadGateway:              http.StatusBadGateway,
	ServiceUnavailable:      http.StatusServiceUnavailable,
	GatewayTimeout:          http.StatusGatewayTimeout,
	HTTPVersionNotSupported: http.StatusHTTPVersionNotSupported,
}

// GRPCCodeMap maps error codes to gRPC codes.
var gRPCCodeMap = map[string]codes.Code{
	OK:                      codes.OK,
	Created:                 codes.OK,            // gRPC has no equivalent for Created
	Accepted:                codes.OK,            // Treated as success
	NoContent:               codes.OK,            // No Content treated as success
	MovedPermanently:        codes.Unimplemented, // gRPC doesn't support redirects
	Found:                   codes.Unimplemented, // Same as above
	BadRequest:              codes.InvalidArgument,
	Unauthorized:            codes.Unauthenticated,
	Forbidden:               codes.PermissionDenied,
	NotFound:                codes.NotFound,
	MethodNotAllowed:        codes.Unimplemented,
	NotAcceptable:           codes.InvalidArgument,
	RequestTimeout:          codes.DeadlineExceeded,
	Conflict:                codes.Aborted,
	Gone:                    codes.NotFound,
	LengthRequired:          codes.FailedPrecondition,
	PreconditionFailed:      codes.FailedPrecondition,
	PayloadTooLarge:         codes.ResourceExhausted,
	URITooLong:              codes.InvalidArgument,
	UnsupportedMediaType:    codes.Unimplemented,
	RangeNotSatisfiable:     codes.OutOfRange,
	ExpectationFailed:       codes.FailedPrecondition,
	ImATeapot:               codes.Internal, // Fun HTTP status, maps to Internal
	UpgradeRequired:         codes.Unimplemented,
	PreconditionRequired:    codes.FailedPrecondition,
	TooManyRequests:         codes.ResourceExhausted,
	InternalServerError:     codes.Internal,
	NotImplemented:          codes.Unimplemented,
	BadGateway:              codes.Unavailable,
	ServiceUnavailable:      codes.Unavailable,
	GatewayTimeout:          codes.DeadlineExceeded,
	HTTPVersionNotSupported: codes.Unimplemented,
}
