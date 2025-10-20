package merchantgrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcValidateCreateMerchant       = response.NewGrpcError("merchant", "Invalid input for create merchant", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateMerchant       = response.NewGrpcError("merchant", "Invalid input for update merchant", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateMerchantStatus = response.NewGrpcError("merchant", "Invalid input for update merchant status", int(codes.InvalidArgument))
)
