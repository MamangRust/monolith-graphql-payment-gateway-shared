package cardgrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcValidateCreateCardRequest = response.NewGrpcError("card", "Invalid input for create card", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateCardRequest = response.NewGrpcError("card", "Invalid input for update card", int(codes.InvalidArgument))
)
