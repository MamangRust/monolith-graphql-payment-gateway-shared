package transactiongrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcValidateCreateTransactionRequest = response.NewGrpcError("transaction", "Invalid input for create card", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateTransactionRequest = response.NewGrpcError("transaction", "Invalid input for update card", int(codes.InvalidArgument))
)
