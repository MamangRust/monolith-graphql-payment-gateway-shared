package topupservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// TopupMethodResponseMapper is an interface that provides methods to map
type topupMethodResponseMapper struct{}

// NewTopupMethodResponseMapper returns a pointer to a new topupMethodResponseMapper instance.
// This struct provides methods to map TopupMonthMethod and TopupYearlyMethod domain models to
// TopupMonthMethodResponse and TopupYearlyMethodResponse API-compatible response types.
func NewTopupMethodResponseMapper() TopupStatsMethodResponseMapper {
	return &topupMethodResponseMapper{}
}

// ToTopupMonthlyMethodResponse converts a TopupMonthMethod domain model
// into a TopupMonthMethodResponse API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupMonthMethod containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupMonthMethodResponse containing the mapped data, including
//     Month, TopupMethod, TotalTopups, and TotalAmount.
func (t *topupMethodResponseMapper) ToTopupMonthlyMethodResponse(s *record.TopupMonthMethod) *response.TopupMonthMethodResponse {
	return &response.TopupMonthMethodResponse{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyMethodResponses maps a slice of TopupMonthMethod domain models
// to a slice of TopupMonthMethodResponse API-compatible response types.
//
// Args:
//   - s: A slice of TopupMonthMethod containing the data to be mapped.
//
// Returns:
//   - A slice of TopupMonthMethodResponse containing the mapped data, including
//     Month, TopupMethod, TotalTopups, and TotalAmount.
func (t *topupMethodResponseMapper) ToTopupMonthlyMethodResponses(s []*record.TopupMonthMethod) []*response.TopupMonthMethodResponse {
	var topupResponses []*response.TopupMonthMethodResponse
	for _, topup := range s {
		topupResponses = append(topupResponses, t.ToTopupMonthlyMethodResponse(topup))
	}
	return topupResponses
}

// ToTopupYearlyMethodResponse converts a TopupYearlyMethod domain model
// into a TopupYearlyMethodResponse API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupYearlyMethod containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupYearlyMethodResponse containing the mapped data, including
//     Year, TopupMethod, TotalTopups, and TotalAmount.
func (t *topupMethodResponseMapper) ToTopupYearlyMethodResponse(s *record.TopupYearlyMethod) *response.TopupYearlyMethodResponse {
	return &response.TopupYearlyMethodResponse{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyMethodResponses maps a slice of TopupYearlyMethod domain models
// to a slice of TopupYearlyMethodResponse API-compatible response types.
//
// Args:
//   - s: A slice of TopupYearlyMethod containing the data to be mapped.
//
// Returns:
//   - A slice of TopupYearlyMethodResponse containing the mapped data, including
//     Year, TopupMethod, TotalTopups, and TotalAmount.
func (t *topupMethodResponseMapper) ToTopupYearlyMethodResponses(s []*record.TopupYearlyMethod) []*response.TopupYearlyMethodResponse {
	var topupResponses []*response.TopupYearlyMethodResponse
	for _, topup := range s {
		topupResponses = append(topupResponses, t.ToTopupYearlyMethodResponse(topup))
	}
	return topupResponses
}
