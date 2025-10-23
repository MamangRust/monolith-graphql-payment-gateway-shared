package errors

import (
	"encoding/json"

	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
)

func GrpcErrorToJson(err *pb.ErrorResponse) string {
	jsonData, _ := json.Marshal(err)
	return string(jsonData)
}
