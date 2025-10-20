package cardstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticTopupRecordMapper provides methods to map database rows to CardStatisticTopupRecord domain models.
type cardStatisticTopupRecordMapper struct{}

// NewCardStatisticTopupRecordMapper returns a new instance of CardStatisticTopupRecordMapper which provides methods
// to map database rows to CardStatisticTopupRecord domain models.
func NewCardStatisticTopupRecordMapper() CardStatisticTopupRecordMapper {
	return &cardStatisticTopupRecordMapper{}
}

// ToMonthlyTopupAmount maps a GetMonthlyTopupAmountRow database row to a CardMonthAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetMonthlyTopupAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTopupRecordMapper) ToMonthlyTopupAmount(card *db.GetMonthlyTopupAmountRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTopupAmount),
	}
}

// ToMonthlyTopupAmounts maps a slice of GetMonthlyTopupAmountRow database rows to a slice of CardMonthAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetMonthlyTopupAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTopupRecordMapper) ToMonthlyTopupAmounts(cards []*db.GetMonthlyTopupAmountRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTopupAmount(card))
	}
	return records
}

// ToYearlyTopupAmount maps a GetYearlyTopupAmountRow database row to a CardYearAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetYearlyTopupAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTopupRecordMapper) ToYearlyTopupAmount(card *db.GetYearlyTopupAmountRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTopupAmount,
	}
}

// ToYearlyTopupAmounts maps a slice of GetYearlyTopupAmountRow database rows to a slice of CardYearAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetYearlyTopupAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTopupRecordMapper) ToYearlyTopupAmounts(cards []*db.GetYearlyTopupAmountRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTopupAmount(card))
	}
	return records
}
