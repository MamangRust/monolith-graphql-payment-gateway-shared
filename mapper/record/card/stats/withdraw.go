package cardstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticWithdrawRecordMapper provides methods to map database rows related to card withdrawals to domain models.
type cardStatisticWithdrawRecordMapper struct {
}

// NewCardStatisticWithdrawRecordMapper creates and returns a new instance of CardStatisticWithdrawRecordMapper,
// which provides methods to map database rows related to card withdrawals to domain models.
func NewCardStatisticWithdrawRecordMapper() CardStatisticWithdrawRecordMapper {
	return &cardStatisticWithdrawRecordMapper{}
}

// ToMonthlyWithdrawAmount maps a GetMonthlyWithdrawAmountRow database row to a CardMonthAmount domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyWithdrawAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticWithdrawRecordMapper) ToMonthlyWithdrawAmount(card *db.GetMonthlyWithdrawAmountRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalWithdrawAmount),
	}
}

// ToMonthlyWithdrawAmounts maps a slice of GetMonthlyWithdrawAmountRow database rows to a slice of CardMonthAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyWithdrawAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardStatisticWithdrawRecordMapper) ToMonthlyWithdrawAmounts(cards []*db.GetMonthlyWithdrawAmountRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyWithdrawAmount(card))
	}
	return records
}

// ToYearlyWithdrawAmount maps a GetYearlyWithdrawAmountRow database row to a CardYearAmount domain model.
//
// Args:
//   - card: A pointer to a GetYearlyWithdrawAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticWithdrawRecordMapper) ToYearlyWithdrawAmount(card *db.GetYearlyWithdrawAmountRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalWithdrawAmount,
	}
}

// ToYearlyWithdrawAmounts maps a slice of GetYearlyWithdrawAmountRow database rows to a slice of CardYearAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyWithdrawAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardStatisticWithdrawRecordMapper) ToYearlyWithdrawAmounts(cards []*db.GetYearlyWithdrawAmountRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyWithdrawAmount(card))
	}
	return records
}
