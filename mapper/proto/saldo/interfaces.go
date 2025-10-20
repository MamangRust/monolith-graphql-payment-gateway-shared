package saldoprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pbhelpers "github.com/MamangRust/monolith-payment-gateway-pb"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
)

type SaldoBaseProtoMapper interface {
	// ToProtoResponseSaldo maps a single saldo record to a protobuf response.
	ToProtoResponseSaldo(status string, message string, pbResponse *response.SaldoResponse) *pb.ApiResponseSaldo
}

type SaldoQueryProtoMapper interface {
	SaldoBaseProtoMapper

	// ToProtoResponsePaginationSaldo maps paginated saldo records to a protobuf response.
	ToProtoResponsePaginationSaldo(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponse) *pb.ApiResponsePaginationSaldo

	// ToProtoResponsePaginationSaldoDeleteAt maps paginated soft-deleted saldo records to a protobuf response.
	ToProtoResponsePaginationSaldoDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponseDeleteAt) *pb.ApiResponsePaginationSaldoDeleteAt
}

type SaldoCommandProtoMapper interface {
	SaldoBaseProtoMapper

	ToProtoResponseSaldoDeleteAt(status string, message string, pbResponse *response.SaldoResponseDeleteAt) *pb.ApiResponseSaldoDeleteAt

	// ToProtoResponseSaldoDelete returns a response indicating a saldo record has been deleted.
	ToProtoResponseSaldoDelete(status string, message string) *pb.ApiResponseSaldoDelete

	// ToProtoResponseSaldoAll returns all saldo records without pagination.
	ToProtoResponseSaldoAll(status string, message string) *pb.ApiResponseSaldoAll
}

type SaldoStatsTotalSaldoProtoMapper interface {
	// ToProtoResponseMonthTotalSaldo maps monthly total saldo to a protobuf response.
	ToProtoResponseMonthTotalSaldo(status string, message string, pbResponse []*response.SaldoMonthTotalBalanceResponse) *pb.ApiResponseMonthTotalSaldo

	// ToProtoResponseYearTotalSaldo maps yearly total saldo to a protobuf response.
	ToProtoResponseYearTotalSaldo(status string, message string, pbResponse []*response.SaldoYearTotalBalanceResponse) *pb.ApiResponseYearTotalSaldo
}

type SaldoStatsBalanceProtoMapper interface {
	// ToProtoResponseMonthSaldoBalances maps monthly saldo balances to a protobuf response.
	ToProtoResponseMonthSaldoBalances(status string, message string, pbResponse []*response.SaldoMonthBalanceResponse) *pb.ApiResponseMonthSaldoBalances

	// ToProtoResponseYearSaldoBalances maps yearly saldo balances to a protobuf response.
	ToProtoResponseYearSaldoBalances(status string, message string, pbResponse []*response.SaldoYearBalanceResponse) *pb.ApiResponseYearSaldoBalances
}
