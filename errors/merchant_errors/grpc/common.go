package merchantgrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcMerchantInvalidID     = response.NewGrpcError("merchant", "Invalid Merchant ID", int(codes.InvalidArgument))
	ErrGrpcMerchantInvalidUserID = response.NewGrpcError("merchant", "Invalid Merchant User ID", int(codes.InvalidArgument))
	ErrGrpcMerchantInvalidApiKey = response.NewGrpcError("merchant", "Invalid Merchant Api Key", int(codes.InvalidArgument))
	ErrGrpcMerchantInvalidMonth  = response.NewGrpcError("month", "Invalid Merchant Month", int(codes.InvalidArgument))
	ErrGrpcMerchantInvalidYear   = response.NewGrpcError("year", "Invalid Merchant Year", int(codes.InvalidArgument))
)
