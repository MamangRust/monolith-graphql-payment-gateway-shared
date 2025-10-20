package withdrawresponseservice

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type withdrawStatsAmountResponseMapper struct{}

func NewWithdrawStatsAmountResponseMapper() WithdrawStatsAmountResponseMapper {
	return &withdrawStatsAmountResponseMapper{}
}

// ToWithdrawAmountMonthlyResponse converts a single WithdrawMonthlyAmount record
// into a WithdrawMonthlyAmountResponse.
//
// Args:
//   - s: The WithdrawMonthlyAmount that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawMonthlyAmountResponse containing the mapped data,
//     including Month and TotalAmount.
func (w *withdrawStatsAmountResponseMapper) ToWithdrawAmountMonthlyResponse(s *record.WithdrawMonthlyAmount) *response.WithdrawMonthlyAmountResponse {
	return &response.WithdrawMonthlyAmountResponse{
		Month:       s.Month,
		TotalAmount: s.TotalAmount,
	}
}

// ToWithdrawsAmountMonthlyResponses converts a slice of WithdrawMonthlyAmount
// records into a slice of WithdrawMonthlyAmountResponse objects.
//
// Args:
//   - s: A slice of pointers to WithdrawMonthlyAmount records that need to be converted.
//
// Returns:
//   - A slice of pointers to WithdrawMonthlyAmountResponse, each containing the mapped data,
//     including Month and TotalAmount.
func (w *withdrawStatsAmountResponseMapper) ToWithdrawsAmountMonthlyResponses(s []*record.WithdrawMonthlyAmount) []*response.WithdrawMonthlyAmountResponse {
	var withdrawResponses []*response.WithdrawMonthlyAmountResponse
	for _, withdraw := range s {
		withdrawResponses = append(withdrawResponses, w.ToWithdrawAmountMonthlyResponse(withdraw))
	}
	return withdrawResponses
}

// ToWithdrawAmountYearlyResponse converts a WithdrawYearlyAmount record
// into a WithdrawYearlyAmountResponse.
//
// Args:
//   - s: The WithdrawYearlyAmount that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawYearlyAmountResponse containing the mapped data, including
//     Year and TotalAmount.
func (w *withdrawStatsAmountResponseMapper) ToWithdrawAmountYearlyResponse(s *record.WithdrawYearlyAmount) *response.WithdrawYearlyAmountResponse {
	return &response.WithdrawYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: s.TotalAmount,
	}
}

// ToWithdrawsAmountYearlyResponses converts a slice of WithdrawYearlyAmount records
// into a slice of WithdrawYearlyAmountResponse.
//
// Args:
//   - s: A slice of pointers to WithdrawYearlyAmount records that need to be converted.
//
// Returns:
//   - A slice of pointers to WithdrawYearlyAmountResponse, each containing the mapped data,
//     including Year and TotalAmount.
func (w *withdrawStatsAmountResponseMapper) ToWithdrawsAmountYearlyResponses(s []*record.WithdrawYearlyAmount) []*response.WithdrawYearlyAmountResponse {
	var withdrawResponses []*response.WithdrawYearlyAmountResponse
	for _, withdraw := range s {
		withdrawResponses = append(withdrawResponses, w.ToWithdrawAmountYearlyResponse(withdraw))
	}
	return withdrawResponses
}
