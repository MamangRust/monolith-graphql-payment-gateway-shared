package saldoservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// SaldoBaseResponseMapper provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types.
type SaldoBaseResponseMapper interface {
	// Converts a single SaldoRecord into SaldoResponse.
	ToSaldoResponse(saldo *record.SaldoRecord) *response.SaldoResponse
}

// SaldoQueryResponseMapper provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types.
type SaldoQueryResponseMapper interface {
	SaldoBaseResponseMapper

	// Converts a list of SaldoRecord into SaldoResponse list.
	ToSaldoResponses(saldos []*record.SaldoRecord) []*response.SaldoResponse

	// Converts multiple soft-deleted SaldoRecords into SaldoResponseDeleteAt list.
	ToSaldoResponsesDeleteAt(saldos []*record.SaldoRecord) []*response.SaldoResponseDeleteAt
}

// SaldoCommandResponseMapper provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types.
type SaldoCommandResponseMapper interface {
	SaldoBaseResponseMapper

	// Converts a soft-deleted SaldoRecord into SaldoResponseDeleteAt.
	ToSaldoResponseDeleteAt(saldo *record.SaldoRecord) *response.SaldoResponseDeleteAt
}

type SaldoStatisticTotalBalanceResponseMapper interface {
	// Converts a monthly total balance record into API response.
	ToSaldoMonthTotalBalanceResponse(ss *record.SaldoMonthTotalBalance) *response.SaldoMonthTotalBalanceResponse

	// Converts multiple monthly total balance records into API response list.
	ToSaldoMonthTotalBalanceResponses(ss []*record.SaldoMonthTotalBalance) []*response.SaldoMonthTotalBalanceResponse

	// Converts a yearly total balance record into API response.
	ToSaldoYearTotalBalanceResponse(ss *record.SaldoYearTotalBalance) *response.SaldoYearTotalBalanceResponse

	// Converts multiple yearly total balance records into API response list.
	ToSaldoYearTotalBalanceResponses(ss []*record.SaldoYearTotalBalance) []*response.SaldoYearTotalBalanceResponse
}

// SaldoStatisticsResponseMapper provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types.
type SaldoStatisticBalanceResponseMapper interface {

	// Converts a monthly balance record into API response.
	ToSaldoMonthBalanceResponse(ss *record.SaldoMonthSaldoBalance) *response.SaldoMonthBalanceResponse

	// Converts multiple monthly balance records into API response list.
	ToSaldoMonthBalanceResponses(ss []*record.SaldoMonthSaldoBalance) []*response.SaldoMonthBalanceResponse

	// Converts a yearly balance record into API response.
	ToSaldoYearBalanceResponse(ss *record.SaldoYearSaldoBalance) *response.SaldoYearBalanceResponse

	// Converts multiple yearly balance records into API response list.
	ToSaldoYearBalanceResponses(ss []*record.SaldoYearSaldoBalance) []*response.SaldoYearBalanceResponse
}
