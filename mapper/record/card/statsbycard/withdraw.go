package cardstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticWithdrawByCardRecordMapper provides methods to map database rows to CardStatisticWithdrawByCardRecord domain models.
type cardStatisticWithdrawRecordMapper struct {
}

// NewCardStatisticWithdrawRecordMapper returns a new instance of CardStatisticWithdrawByCardRecordMapper
// which provides methods to map GetMonthlyWithdrawAmountByCardNumberRow and GetYearlyWithdrawAmountByCardNumberRow
// database rows to CardMonthAmount and CardYearAmount domain models.
func NewCardStatisticWithdrawRecordMapper() CardStatisticWithdrawByCardRecordMapper {
	return &cardStatisticWithdrawRecordMapper{}
}

// ToMonthlyWithdrawAmountByCardNumber maps a GetMonthlyWithdrawAmountByCardNumberRow database row to a CardMonthAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetMonthlyWithdrawAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticWithdrawRecordMapper) ToMonthlyWithdrawAmountByCardNumber(card *db.GetMonthlyWithdrawAmountByCardNumberRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalWithdrawAmount),
	}
}

// ToMonthlyWithdrawAmountsByCardNumber maps a slice of GetMonthlyWithdrawAmountByCardNumberRow database rows
// to a slice of CardMonthAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetMonthlyWithdrawAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticWithdrawRecordMapper) ToMonthlyWithdrawAmountsByCardNumber(cards []*db.GetMonthlyWithdrawAmountByCardNumberRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyWithdrawAmountByCardNumber(card))
	}
	return records
}

// ToYearlyWithdrawAmountByCardNumber maps a GetYearlyWithdrawAmountByCardNumberRow database row to a CardYearAmount domain model.
//
// Parameters:
//   - card: A pointer to a GetYearlyWithdrawAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticWithdrawRecordMapper) ToYearlyWithdrawAmountByCardNumber(card *db.GetYearlyWithdrawAmountByCardNumberRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalWithdrawAmount,
	}
}

// ToYearlyWithdrawAmountsByCardNumber maps a slice of GetYearlyWithdrawAmountByCardNumberRow database rows
// to a slice of CardYearAmount domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetYearlyWithdrawAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticWithdrawRecordMapper) ToYearlyWithdrawAmountsByCardNumber(cards []*db.GetYearlyWithdrawAmountByCardNumberRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyWithdrawAmountByCardNumber(card))
	}
	return records
}
