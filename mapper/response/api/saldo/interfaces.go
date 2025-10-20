package saldoapimapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
)

type SaldoBaseResponseMapper interface {
	// Converts a single saldo gRPC response to an API response.
	ToApiResponseSaldo(pbResponse *pb.ApiResponseSaldo) *response.ApiResponseSaldo
}

type SaldoQueryResponseMapper interface {
	SaldoBaseResponseMapper

	// Converts paginated saldo responses into an API response.
	ToApiResponsePaginationSaldo(pbResponse *pb.ApiResponsePaginationSaldo) *response.ApiResponsePaginationSaldo

	// Converts paginated soft-deleted saldo responses into an API response.
	ToApiResponsePaginationSaldoDeleteAt(pbResponse *pb.ApiResponsePaginationSaldoDeleteAt) *response.ApiResponsePaginationSaldoDeleteAt
}

type SaldoCommandResponseMapper interface {
	SaldoBaseResponseMapper

	ToApiResponseSaldoDeleteAt(pbResponse *pb.ApiResponseSaldoDeleteAt) *response.ApiResponseSaldoDeleteAt

	// Converts a deleted saldo response into an API response.
	ToApiResponseSaldoDelete(pbResponse *pb.ApiResponseSaldoDelete) *response.ApiResponseSaldoDelete

	// Converts all saldo records into a comprehensive API response.
	ToApiResponseSaldoAll(pbResponse *pb.ApiResponseSaldoAll) *response.ApiResponseSaldoAll
}

type SaldoStatsTotalResponseMapper interface {
	// Converts monthly total saldo values into an API response.
	ToApiResponseMonthTotalSaldo(pbResponse *pb.ApiResponseMonthTotalSaldo) *response.ApiResponseMonthTotalSaldo

	// Converts yearly total saldo values into an API response.
	ToApiResponseYearTotalSaldo(pbResponse *pb.ApiResponseYearTotalSaldo) *response.ApiResponseYearTotalSaldo
}

type SaldoStatsBalanceResponseMapper interface {
	// Converts monthly saldo balances into an API response.
	ToApiResponseMonthSaldoBalances(pbResponse *pb.ApiResponseMonthSaldoBalances) *response.ApiResponseMonthSaldoBalances

	// Converts yearly saldo balances into an API response.
	ToApiResponseYearSaldoBalances(pbResponse *pb.ApiResponseYearSaldoBalances) *response.ApiResponseYearSaldoBalances
}
