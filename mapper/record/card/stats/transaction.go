package cardstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticTransactionRecordMapper provides methods to map database rows to CardStatisticTransactionRecord domain models.
type cardStatisticTransactionRecordMapper struct {
}

// NewCardStatisticTransactionRecordMapper returns a new instance of CardStatisticTransactionRecordMapper which provides methods to map database rows to CardStatisticTransactionRecord domain models.
func NewCardStatisticTransactionRecordMapper() CardStatisticTransactionRecordMapper {
	return &cardStatisticTransactionRecordMapper{}
}

// ToMonthlyTransactionAmount maps a GetMonthlyTransactionAmountRow database row to a CardMonthAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetMonthlyTransactionAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransactionRecordMapper) ToMonthlyTransactionAmount(card *db.GetMonthlyTransactionAmountRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTransactionAmount),
	}
}

// ToMonthlyTransactionAmounts maps a slice of GetMonthlyTransactionAmountRow database rows to a slice of CardMonthAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetMonthlyTransactionAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransactionRecordMapper) ToMonthlyTransactionAmounts(cards []*db.GetMonthlyTransactionAmountRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransactionAmount(card))
	}
	return records
}

// ToYearlyTransactionAmount maps a GetYearlyTransactionAmountRow database row to a CardYearAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetYearlyTransactionAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransactionRecordMapper) ToYearlyTransactionAmount(card *db.GetYearlyTransactionAmountRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTransactionAmount,
	}
}

// ToYearlyTransactionAmounts maps a slice of GetYearlyTransactionAmountRow database rows to a slice of CardYearAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetYearlyTransactionAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransactionRecordMapper) ToYearlyTransactionAmounts(cards []*db.GetYearlyTransactionAmountRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransactionAmount(card))
	}
	return records
}
