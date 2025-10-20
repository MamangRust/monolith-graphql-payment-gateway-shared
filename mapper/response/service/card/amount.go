package cardservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// CardStatisticAmountResponseMapper provides methods to map CardMonthAmount and CardYearAmount domain models to
// CardResponseMonthAmount and CardResponseYearAmount domain models.
type cardStatsAmountResponseMapper struct {
}

// NewCardStatsAmountResponseMapper returns a new instance of cardStatsAmountResponseMapper, which provides methods to map
// CardMonthAmount and CardYearAmount domain models to CardResponseMonthAmount and CardResponseYearAmount
// domain models.
func NewCardStatsAmountResponseMapper() CardStatisticAmountResponseMapper {
	return &cardStatsAmountResponseMapper{}
}

// ToGetMonthlyAmount maps a CardMonthAmount domain model to a CardResponseMonthAmount
// domain model.
//
// Parameters:
//   - card: A pointer to a CardMonthAmount representing the domain model.
//
// Returns:
//   - A pointer to a CardResponseMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatsAmountResponseMapper) ToGetMonthlyAmount(card *record.CardMonthAmount) *response.CardResponseMonthAmount {
	return &response.CardResponseMonthAmount{
		Month:       card.Month,
		TotalAmount: card.TotalAmount,
	}
}

// ToGetMonthlyAmounts converts a slice of CardMonthAmount domain models into a slice
// of CardResponseMonthAmount response objects.
//
// Parameters:
//   - cards: A slice of pointers to CardMonthAmount domain models.
//
// Returns:
//   - A slice of pointers to CardResponseMonthAmount containing the mapped data for each month.
func (s *cardStatsAmountResponseMapper) ToGetMonthlyAmounts(cards []*record.CardMonthAmount) []*response.CardResponseMonthAmount {
	var records []*response.CardResponseMonthAmount

	for _, card := range cards {
		records = append(records, s.ToGetMonthlyAmount(card))
	}

	return records
}

// ToGetYearlyAmount maps a CardYearAmount domain model to a CardResponseYearAmount
// domain model.
//
// Parameters:
//   - card: A pointer to a CardYearAmount representing the domain model.
//
// Returns:
//   - A pointer to a CardResponseYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatsAmountResponseMapper) ToGetYearlyAmount(card *record.CardYearAmount) *response.CardResponseYearAmount {
	return &response.CardResponseYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalAmount,
	}
}

// ToGetYearlyAmounts converts a slice of CardYearAmount domain models into a slice
// of CardResponseYearAmount response objects.
//
// Parameters:
//   - cards: A slice of pointers to CardYearAmount domain models.
//
// Returns:
//   - A slice of pointers to CardResponseYearAmount containing the mapped data for each year.
func (s *cardStatsAmountResponseMapper) ToGetYearlyAmounts(cards []*record.CardYearAmount) []*response.CardResponseYearAmount {
	var records []*response.CardResponseYearAmount

	for _, card := range cards {
		records = append(records, s.ToGetYearlyAmount(card))
	}

	return records
}
