package topupgrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcTopupInvalidID    = response.NewGrpcError("topup", "Invalid Topup ID", int(codes.InvalidArgument))
	ErrGrpcTopupInvalidMonth = response.NewGrpcError("month", "Invalid Topup Month", int(codes.InvalidArgument))
	ErrGrpcInvalidCardNumber = response.NewGrpcError("card_id", "Invalid card number", int(codes.InvalidArgument))
	ErrGrpcTopupInvalidYear  = response.NewGrpcError("year", "Invalid Topup Year", int(codes.InvalidArgument))
)
