package topupservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// TopupStatsAmountResponseMapper is an interface that provides methods to map TopupMonthAmount and TopupYearlyAmount
type topupStatsAmountResponseMapper struct{}

// NewTopupStatsAmountResponseMapper returns a pointer to a new topupStatsAmountResponseMapper instance.
// This struct provides methods to map TopupMonthAmount and TopupYearlyAmount domain models to
// TopupMonthAmountResponse and TopupYearlyAmountResponse API-compatible response types,
// respectively.
func NewTopupStatsAmountResponseMapper() TopupStatsAmountResponseMapper {
	return &topupStatsAmountResponseMapper{}
}

// ToTopupMonthlyAmountResponse converts a TopupMonthAmount domain model
// into a TopupMonthAmountResponse API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupMonthAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupMonthAmountResponse containing the mapped data, including
//     Month, TotalAmount.
func (t *topupStatsAmountResponseMapper) ToTopupMonthlyAmountResponse(s *record.TopupMonthAmount) *response.TopupMonthAmountResponse {
	return &response.TopupMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyAmountResponses converts a slice of TopupMonthAmount domain models
// into a slice of TopupMonthAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of TopupMonthAmount containing the data to be mapped.
//
// Returns:
//   - A slice of TopupMonthAmountResponse containing the mapped data, including
//     Month and TotalAmount.
func (t *topupStatsAmountResponseMapper) ToTopupMonthlyAmountResponses(s []*record.TopupMonthAmount) []*response.TopupMonthAmountResponse {
	var topupResponses []*response.TopupMonthAmountResponse
	for _, topup := range s {
		topupResponses = append(topupResponses, t.ToTopupMonthlyAmountResponse(topup))
	}
	return topupResponses
}

// ToTopupYearlyAmountResponse converts a TopupYearlyAmount domain model
// into a TopupYearlyAmountResponse API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupYearlyAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupYearlyAmountResponse containing the mapped data, including
//     Year, TotalAmount.
func (t *topupStatsAmountResponseMapper) ToTopupYearlyAmountResponse(s *record.TopupYearlyAmount) *response.TopupYearlyAmountResponse {
	return &response.TopupYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyAmountResponses maps a slice of TopupYearlyAmount domain models
// to a slice of TopupYearlyAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of TopupYearlyAmount containing the data to be mapped.
//
// Returns:
//   - A slice of TopupYearlyAmountResponse containing the mapped data, including
//     Year, TotalAmount.
func (t *topupStatsAmountResponseMapper) ToTopupYearlyAmountResponses(s []*record.TopupYearlyAmount) []*response.TopupYearlyAmountResponse {
	var topupResponses []*response.TopupYearlyAmountResponse
	for _, topup := range s {
		topupResponses = append(topupResponses, t.ToTopupYearlyAmountResponse(topup))
	}
	return topupResponses
}
