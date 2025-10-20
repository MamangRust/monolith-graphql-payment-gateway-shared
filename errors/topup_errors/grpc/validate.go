package topupgrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcValidateCreateTopup = response.NewGrpcError("topup", "Invalid input for create topup", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateTopup = response.NewGrpcError("topup", "Invalid input for update topup", int(codes.InvalidArgument))
)
