package saldogrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcSaldoInvalidID         = response.NewGrpcError("saldo", "Invalid Saldo ID", int(codes.InvalidArgument))
	ErrGrpcSaldoInvalidCardNumber = response.NewGrpcError("saldo", "Invalid Saldo Card Number", int(codes.InvalidArgument))
	ErrGrpcSaldoInvalidMonth      = response.NewGrpcError("saldo", "Invalid Saldo Month", int(codes.InvalidArgument))
	ErrGrpcSaldoInvalidYear       = response.NewGrpcError("saldo", "Invalid Saldo Year", int(codes.InvalidArgument))
)
