package transfergrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcValidateCreateTransferRequest = response.NewGrpcError("transfer", "Invalid input for create transfer", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateTransferRequest = response.NewGrpcError("transfer", "Invalid input for update transfer", int(codes.InvalidArgument))
)
