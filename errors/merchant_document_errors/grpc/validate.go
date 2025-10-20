package merchantdocumentgrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcValidateCreateMerchantDocument = response.NewGrpcError("merchant_document", "Invalid input for create merchant document", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateMerchantDocument = response.NewGrpcError("merchant_document", "Invalid input for update merchant document", int(codes.InvalidArgument))
)
