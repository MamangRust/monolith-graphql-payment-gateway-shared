package saldoservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// SaldoTotalBalanceResponseMapper provides methods to map SaldoMonthTotalBalance and SaldoYearTotalBalance records to SaldoMonthTotalBalanceResponse and SaldoYearTotalBalanceResponse API-compatible responses.
type saldoTotalBalanceResponseMapper struct {
}

// NewSaldoTotalBalanceResponseMapper creates a new SaldoTotalBalanceResponseMapper.
//
// Returns:
//   - A pointer to a SaldoTotalBalanceResponseMapper.
func NewSaldoTotalBalanceResponseMapper() SaldoStatisticTotalBalanceResponseMapper {
	return &saldoTotalBalanceResponseMapper{}
}

// ToSaldoMonthTotalBalanceResponse maps a SaldoMonthTotalBalance record to a SaldoMonthTotalBalanceResponse.
//
// Args:
//   - ss: A pointer to a SaldoMonthTotalBalance containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoMonthTotalBalanceResponse containing the mapped data, including Month, Year, and TotalBalance.
func (s *saldoTotalBalanceResponseMapper) ToSaldoMonthTotalBalanceResponse(ss *record.SaldoMonthTotalBalance) *response.SaldoMonthTotalBalanceResponse {
	totalBalance := 0

	if ss.TotalBalance != 0 {
		totalBalance = ss.TotalBalance
	}

	return &response.SaldoMonthTotalBalanceResponse{
		Month:        ss.Month,
		Year:         ss.Year,
		TotalBalance: totalBalance,
	}
}

// ToSaldoMonthTotalBalanceResponses maps a list of SaldoMonthTotalBalance records to a list of SaldoMonthTotalBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to SaldoMonthTotalBalance containing the data to be mapped.
//
// Returns:
//
//	A slice of pointers to SaldoMonthTotalBalanceResponse containing the mapped data, including Month, Year, and TotalBalance.
func (s *saldoTotalBalanceResponseMapper) ToSaldoMonthTotalBalanceResponses(ss []*record.SaldoMonthTotalBalance) []*response.SaldoMonthTotalBalanceResponse {
	var saldoResponses []*response.SaldoMonthTotalBalanceResponse
	for _, saldo := range ss {
		saldoResponses = append(saldoResponses, s.ToSaldoMonthTotalBalanceResponse(saldo))
	}
	return saldoResponses
}

// ToSaldoYearTotalBalanceResponse maps a SaldoYearTotalBalance record to a SaldoYearTotalBalanceResponse.
//
// Args:
//   - ss: A pointer to a SaldoYearTotalBalance containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoYearTotalBalanceResponse containing the mapped data, including Year, and TotalBalance.
func (s *saldoTotalBalanceResponseMapper) ToSaldoYearTotalBalanceResponse(ss *record.SaldoYearTotalBalance) *response.SaldoYearTotalBalanceResponse {
	return &response.SaldoYearTotalBalanceResponse{
		Year:         ss.Year,
		TotalBalance: ss.TotalBalance,
	}
}

// ToSaldoYearTotalBalanceResponses maps a list of SaldoYearTotalBalance records to a list of SaldoYearTotalBalanceResponse.
//
// Args:
//   - ss: A slice of pointers to SaldoYearTotalBalance containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to SaldoYearTotalBalanceResponse containing the mapped data, including Year, and TotalBalance.
func (s *saldoTotalBalanceResponseMapper) ToSaldoYearTotalBalanceResponses(ss []*record.SaldoYearTotalBalance) []*response.SaldoYearTotalBalanceResponse {
	var saldoResponses []*response.SaldoYearTotalBalanceResponse
	for _, saldo := range ss {
		saldoResponses = append(saldoResponses, s.ToSaldoYearTotalBalanceResponse(saldo))
	}
	return saldoResponses
}
