package saldogrpcerrors

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcValidateCreateSaldo = response.NewGrpcError("saldo", "Invalid input for create saldo", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateSaldo = response.NewGrpcError("saldo", "Invalid input for update saldo", int(codes.InvalidArgument))
)
