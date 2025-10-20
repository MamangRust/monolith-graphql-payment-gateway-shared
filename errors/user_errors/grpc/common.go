package usergrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var ErrGrpcUserInvalidId = response.NewGrpcError("error", "Invalid User ID", int(codes.NotFound))
