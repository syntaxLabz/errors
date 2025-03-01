package codes

import "google.golang.org/grpc/codes"

func GRPCCode(code string) codes.Code {
	return gRPCCodeMap[code]
}

func RESTCode(code string) int {
	return restCodeMap[code]
}
