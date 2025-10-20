package merchantdocumentgrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var ErrGrpcMerchantInvalidID = response.NewGrpcError("merchant_document", "Invalid merchant id", int(codes.InvalidArgument))
