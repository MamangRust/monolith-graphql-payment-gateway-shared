package withdrawgrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcValidateCreateWithdrawRequest = response.NewGrpcError("withdraw", "Invalid input for create withdraw", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateWithdrawRequest = response.NewGrpcError("withdraw", "Invalid input for update withdraw", int(codes.InvalidArgument))
)
