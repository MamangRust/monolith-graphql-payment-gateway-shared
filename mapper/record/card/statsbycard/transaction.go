package cardstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticTransactionRecordMapper provides methods to map database rows to CardStatisticTransactionRecord domain models.
type cardStatisticTransactionRecordMapper struct {
}

// NewCardStatisticTransactionRecordMapper returns a new instance of
// CardStatisticTransactionRecordMapper, which provides methods to map
// database rows to CardStatisticTransactionRecord domain models.
func NewCardStatisticTransactionRecordMapper() CardStatisticTransactionByCardRecordMapper {
	return &cardStatisticTransactionRecordMapper{}
}

// ToMonthlyTransactionAmountByCardNumber maps a GetMonthlyTransactionAmountByCardNumberRow database row
// to a CardMonthAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetMonthlyTransactionAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransactionRecordMapper) ToMonthlyTransactionAmountByCardNumber(card *db.GetMonthlyTransactionAmountByCardNumberRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTransactionAmount),
	}
}

// ToMonthlyTransactionAmountsByCardNumber maps a slice of GetMonthlyTransactionAmountByCardNumberRow database rows
// to a slice of CardMonthAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetMonthlyTransactionAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTransactionRecordMapper) ToMonthlyTransactionAmountsByCardNumber(cards []*db.GetMonthlyTransactionAmountByCardNumberRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransactionAmountByCardNumber(card))
	}
	return records
}

// ToYearlyTransactionAmountByCardNumber maps a GetYearlyTransactionAmountByCardNumberRow database row
// to a CardYearAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetYearlyTransactionAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransactionRecordMapper) ToYearlyTransactionAmountByCardNumber(card *db.GetYearlyTransactionAmountByCardNumberRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTransactionAmount,
	}
}

// ToYearlyTransactionAmountsByCardNumber maps a slice of GetYearlyTransactionAmountByCardNumberRow database rows
// to a slice of CardYearAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetYearlyTransactionAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTransactionRecordMapper) ToYearlyTransactionAmountsByCardNumber(cards []*db.GetYearlyTransactionAmountByCardNumberRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransactionAmountByCardNumber(card))
	}
	return records
}
