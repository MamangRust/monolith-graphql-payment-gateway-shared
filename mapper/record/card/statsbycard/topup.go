package cardstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticTopupByCardRecordMapper provides methods to map database rows to CardStatisticTopupByCardRecord domain models.
type cardStatisticTopupRecordMapper struct {
}

// NewCardStatisticTopupRecordMapper returns a new instance of CardStatisticTopupByCardRecordMapper
// which provides methods to map GetMonthlyTopupAmountByCardNumberRow and GetYearlyTopupAmountByCardNumberRow
// database rows to CardMonthAmount and CardYearAmount domain models.
func NewCardStatisticTopupRecordMapper() CardStatisticTopupByCardRecordMapper {
	return &cardStatisticTopupRecordMapper{}
}

// ToMonthlyTopupAmountByCardNumber maps a GetMonthlyTopupAmountByCardNumberRow database row to a CardMonthAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetMonthlyTopupAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTopupRecordMapper) ToMonthlyTopupAmountByCardNumber(card *db.GetMonthlyTopupAmountByCardNumberRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTopupAmount),
	}
}

// ToMonthlyTopupAmountsByCardNumber maps a slice of GetMonthlyTopupAmountByCardNumberRow database rows to a slice of CardMonthAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetMonthlyTopupAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticTopupRecordMapper) ToMonthlyTopupAmountsByCardNumber(cards []*db.GetMonthlyTopupAmountByCardNumberRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTopupAmountByCardNumber(card))
	}
	return records
}

// ToYearlyTopupAmountByCardNumber maps a GetYearlyTopupAmountByCardNumberRow database row to a CardYearAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetYearlyTopupAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTopupRecordMapper) ToYearlyTopupAmountByCardNumber(card *db.GetYearlyTopupAmountByCardNumberRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTopupAmount,
	}
}

// ToYearlyTopupAmountsByCardNumber maps a slice of GetYearlyTopupAmountByCardNumberRow database rows
// to a slice of CardYearAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetYearlyTopupAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticTopupRecordMapper) ToYearlyTopupAmountsByCardNumber(cards []*db.GetYearlyTopupAmountByCardNumberRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTopupAmountByCardNumber(card))
	}
	return records
}
