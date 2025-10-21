package errors

import (
	"encoding/json"

	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
)

// GrpcErrorToJson takes a pointer to a gRPC ErrorResponse and returns a string containing
// the JSON representation of the error.
//
// The function does not return an error and instead panics if the marshaling fails.
func GrpcErrorToJson(err *pb.ErrorResponse) string {
	jsonData, _ := json.Marshal(err)
	return string(jsonData)
}
