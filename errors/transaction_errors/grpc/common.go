package transactiongrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcTransactionInvalidID         = response.NewGrpcError("transaction", "Invalid Transaction ID", int(codes.InvalidArgument))
	ErrGrpcTransactionInvalidMerchantID = response.NewGrpcError("transaction", "Invalid Transaction Merchant ID", int(codes.InvalidArgument))
	ErrGrpcInvalidCardNumber            = response.NewGrpcError("card_id", "Invalid card number", int(codes.InvalidArgument))
	ErrGrpcInvalidMonth                 = response.NewGrpcError("month", "Invalid month", int(codes.InvalidArgument))
	ErrGrpcInvalidYear                  = response.NewGrpcError("year", "Invalid year", int(codes.InvalidArgument))
)
