package merchantdocumentgrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcFailedCreateMerchantDocument = response.NewGrpcError("merchant_document", "Failed to create merchant document", int(codes.Internal))
	ErrGrpcFailedUpdateMerchantDocument = response.NewGrpcError("merchant_document", "Failed to update merchant document", int(codes.Internal))
)
