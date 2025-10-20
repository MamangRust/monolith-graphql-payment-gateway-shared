package saldoservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// SaldoBalanceResponseMapper provides methods to map SaldoMonthSaldoBalance and SaldoYearSaldoBalance records to SaldoMonthBalanceResponse and SaldoYearBalanceResponse API-compatible responses.
type saldoStatsBalanceResponseMapper struct {
}

// NewSaldoBalanceResponseMapper creates and returns a new instance of SaldoBalanceResponseMapper.
func NewSaldoStatsBalanceResponseMapper() SaldoStatisticBalanceResponseMapper {
	return &saldoStatsBalanceResponseMapper{}
}

// ToSaldoMonthBalanceResponse maps a SaldoMonthSaldoBalance record to a SaldoMonthBalanceResponse.
//
// Args:
//   - ss: A pointer to a SaldoMonthSaldoBalance containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoMonthBalanceResponse containing the mapped data, including Month, and TotalBalance.
func (s *saldoStatsBalanceResponseMapper) ToSaldoMonthBalanceResponse(ss *record.SaldoMonthSaldoBalance) *response.SaldoMonthBalanceResponse {
	return &response.SaldoMonthBalanceResponse{
		Month:        ss.Month,
		TotalBalance: ss.TotalBalance,
	}
}

// ToSaldoMonthBalanceResponses maps a list of SaldoMonthSaldoBalance records to a list of SaldoMonthBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to SaldoMonthSaldoBalance containing the data to be mapped.
//
// Returns:
//
//	A slice of pointers to SaldoMonthBalanceResponse containing the mapped data, including Month, and TotalBalance.
func (s *saldoStatsBalanceResponseMapper) ToSaldoMonthBalanceResponses(ss []*record.SaldoMonthSaldoBalance) []*response.SaldoMonthBalanceResponse {
	var saldoResponses []*response.SaldoMonthBalanceResponse
	for _, saldo := range ss {
		saldoResponses = append(saldoResponses, s.ToSaldoMonthBalanceResponse(saldo))
	}
	return saldoResponses
}

// ToSaldoYearBalanceResponse maps a SaldoYearSaldoBalance record to a SaldoYearBalanceResponse.
//
// Args:
//   - ss: A pointer to a SaldoYearSaldoBalance containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoYearBalanceResponse containing the mapped data, including Year, and TotalBalance.
func (s *saldoStatsBalanceResponseMapper) ToSaldoYearBalanceResponse(ss *record.SaldoYearSaldoBalance) *response.SaldoYearBalanceResponse {
	return &response.SaldoYearBalanceResponse{
		Year:         ss.Year,
		TotalBalance: ss.TotalBalance,
	}
}

// ToSaldoYearBalanceResponses maps a list of SaldoYearSaldoBalance records to a list of SaldoYearBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to SaldoYearSaldoBalance containing the data to be mapped.
//
// Returns:
//
//	A slice of pointers to SaldoYearBalanceResponse containing the mapped data, including Year, and TotalBalance.
func (s *saldoStatsBalanceResponseMapper) ToSaldoYearBalanceResponses(ss []*record.SaldoYearSaldoBalance) []*response.SaldoYearBalanceResponse {
	var saldoResponses []*response.SaldoYearBalanceResponse
	for _, saldo := range ss {
		saldoResponses = append(saldoResponses, s.ToSaldoYearBalanceResponse(saldo))
	}
	return saldoResponses
}
