package usergrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcValidateCreateUser = response.NewGrpcError("error", "validation failed: invalid create User request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateUser = response.NewGrpcError("error", "validation failed: invalid update User request", int(codes.InvalidArgument))
)
