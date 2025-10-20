package cardstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticBalanceByCardRecordMapper provides methods to map database rows to CardStatisticBalanceByCardRecord domain models.
type cardStatisticBalanceByCardRecordMapper struct {
}

// NewCardStatisticBalanceByCardRecordMapper creates a new CardStatisticBalanceByCardRecordMapper.
func NewCardStatisticBalanceByCardRecordMapper() CardStatisticBalanceByCardRecordMapper {
	return &cardStatisticBalanceByCardRecordMapper{}
}

// ToMonthlyBalanceCardNumber maps a GetMonthlyBalancesByCardNumberRow database row to a CardMonthBalance domain model.
//
// Parameters:
//   - card: A pointer to a GetMonthlyBalancesByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthBalance containing the mapped data, including Month and TotalBalance.
func (s *cardStatisticBalanceByCardRecordMapper) ToMonthlyBalanceCardNumber(card *db.GetMonthlyBalancesByCardNumberRow) *record.CardMonthBalance {
	return &record.CardMonthBalance{
		Month:        card.Month,
		TotalBalance: int64(card.TotalBalance),
	}
}

// ToMonthlyBalancesCardNumber maps a slice of GetMonthlyBalancesByCardNumberRow database rows to a slice of CardMonthBalance domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetMonthlyBalancesByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthBalance containing the mapped data, including Month and TotalBalance.
func (s *cardStatisticBalanceByCardRecordMapper) ToMonthlyBalancesCardNumber(cards []*db.GetMonthlyBalancesByCardNumberRow) []*record.CardMonthBalance {
	var records []*record.CardMonthBalance
	for _, card := range cards {
		records = append(records, s.ToMonthlyBalanceCardNumber(card))
	}
	return records
}

// ToYearlyBalanceCardNumber maps a GetYearlyBalancesByCardNumberRow database row to a CardYearlyBalance domain model.
//
// Parameters:
//   - card: A pointer to a GetYearlyBalancesByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearlyBalance containing the mapped data, including Year and TotalBalance.
func (s *cardStatisticBalanceByCardRecordMapper) ToYearlyBalanceCardNumber(card *db.GetYearlyBalancesByCardNumberRow) *record.CardYearlyBalance {
	return &record.CardYearlyBalance{
		Year:         card.Year,
		TotalBalance: card.TotalBalance,
	}
}

// ToYearlyBalancesCardNumber maps a slice of GetYearlyBalancesByCardNumberRow database rows to a slice of CardYearlyBalance domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetYearlyBalancesByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearlyBalance containing the mapped data, including Year and TotalBalance.
func (s *cardStatisticBalanceByCardRecordMapper) ToYearlyBalancesCardNumber(cards []*db.GetYearlyBalancesByCardNumberRow) []*record.CardYearlyBalance {
	var records []*record.CardYearlyBalance
	for _, card := range cards {
		records = append(records, s.ToYearlyBalanceCardNumber(card))
	}
	return records
}
