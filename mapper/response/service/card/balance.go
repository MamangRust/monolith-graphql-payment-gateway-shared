package cardservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// CardStatisticBalanceResponseMapper provides methods to map CardMonthBalance and CardYearlyBalance domain models to
// CardResponseMonthBalance and CardResponseYearlyBalance domain models.
type cardStatsBalanceResponseMapper struct {
}

// NewCardStatsBalanceResponseMapper creates a new instance of CardStatisticBalanceResponseMapper.
//
// It returns a pointer to a cardStatsBalanceResponseMapper which implements the
// CardStatisticBalanceResponseMapper interface for mapping database rows to domain
// models.
func NewCardStatsBalanceResponseMapper() CardStatisticBalanceResponseMapper {
	return &cardStatsBalanceResponseMapper{}
}

// ToGetMonthlyBalance maps a CardMonthBalance domain model to a CardResponseMonthBalance
// response object.
//
// Parameters:
//   - card: A pointer to a CardMonthBalance domain model.
//
// Returns:
//   - A pointer to a CardResponseMonthBalance containing the mapped data, including
//     Month and TotalBalance.
func (s *cardStatsBalanceResponseMapper) ToGetMonthlyBalance(card *record.CardMonthBalance) *response.CardResponseMonthBalance {
	return &response.CardResponseMonthBalance{
		Month:        card.Month,
		TotalBalance: card.TotalBalance,
	}
}

// ToGetMonthlyBalances converts a slice of CardMonthBalance domain models into a slice
// of CardResponseMonthBalance response objects.
//
// Parameters:
//   - cards: A slice of pointers to CardMonthBalance domain models.
//
// Returns:
//   - A slice of pointers to CardResponseMonthBalance containing the mapped data for each month.
func (s *cardStatsBalanceResponseMapper) ToGetMonthlyBalances(cards []*record.CardMonthBalance) []*response.CardResponseMonthBalance {
	var records []*response.CardResponseMonthBalance

	for _, card := range cards {
		records = append(records, s.ToGetMonthlyBalance(card))
	}

	return records
}

// ToGetYearlyBalance maps a CardYearlyBalance domain model to a CardResponseYearlyBalance
// response object.
//
// Parameters:
//   - card: A pointer to a CardYearlyBalance domain model.
//
// Returns:
//   - A pointer to a CardResponseYearlyBalance containing the mapped data, including
//     Year and TotalBalance.
func (s *cardStatsBalanceResponseMapper) ToGetYearlyBalance(card *record.CardYearlyBalance) *response.CardResponseYearlyBalance {
	return &response.CardResponseYearlyBalance{
		Year:         card.Year,
		TotalBalance: card.TotalBalance,
	}
}

// ToGetYearlyBalances converts a slice of CardYearlyBalance domain models into a slice
// of CardResponseYearlyBalance response objects.
//
// Parameters:
//   - cards: A slice of pointers to CardYearlyBalance domain models.
//
// Returns:
//   - A slice of pointers to CardResponseYearlyBalance containing the mapped data for each year.
func (s *cardStatsBalanceResponseMapper) ToGetYearlyBalances(cards []*record.CardYearlyBalance) []*response.CardResponseYearlyBalance {
	var records []*response.CardResponseYearlyBalance

	for _, card := range cards {
		records = append(records, s.ToGetYearlyBalance(card))
	}

	return records
}
