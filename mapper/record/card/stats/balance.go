package cardstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticBalanceRecordMapper is an interface that provides methods to map Card database rows to CardRecord domain models for statistic operations.
type cardStatisticBalanceRecordMapper struct {
}

// NewCardStatisticBalanceRecordMapper creates a new instance of CardStatisticBalanceRecordMapper.
// It initializes and returns a pointer to cardStatisticBalanceRecordMapper which implements the
// CardStatisticBalanceRecordMapper interface for mapping database rows to domain models.
func NewCardStatisticBalanceRecordMapper() CardStatisticBalanceRecordMapper {
	return &cardStatisticBalanceRecordMapper{}
}

// ToMonthlyBalance maps a GetMonthlyBalancesRow database row to a CardMonthBalance domain model.
//
// Parameters:
//   - card: A pointer to a GetMonthlyBalancesRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthBalance containing the mapped data, including Month and TotalBalance.
func (s *cardStatisticBalanceRecordMapper) ToMonthlyBalance(card *db.GetMonthlyBalancesRow) *record.CardMonthBalance {
	return &record.CardMonthBalance{
		Month:        card.Month,
		TotalBalance: int64(card.TotalBalance),
	}
}

// ToMonthlyBalances maps a slice of GetMonthlyBalancesRow database rows to a slice of CardMonthBalance domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetMonthlyBalancesRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthBalance containing the mapped data, including Month and TotalBalance.
func (s *cardStatisticBalanceRecordMapper) ToMonthlyBalances(cards []*db.GetMonthlyBalancesRow) []*record.CardMonthBalance {
	var records []*record.CardMonthBalance
	for _, card := range cards {
		records = append(records, s.ToMonthlyBalance(card))
	}
	return records
}

// ToYearlyBalance maps a GetYearlyBalancesRow database row to a CardYearlyBalance domain model.
//
// Parameters:
//   - card: A pointer to a GetYearlyBalancesRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearlyBalance containing the mapped data, including Year and TotalBalance.
func (s *cardStatisticBalanceRecordMapper) ToYearlyBalance(card *db.GetYearlyBalancesRow) *record.CardYearlyBalance {
	return &record.CardYearlyBalance{
		Year:         card.Year,
		TotalBalance: card.TotalBalance,
	}
}

// ToYearlyBalances maps a slice of GetYearlyBalancesRow database rows to a slice of CardYearlyBalance domain models.
//
// Parameters:
//   - cards: A slice of pointers to GetYearlyBalancesRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearlyBalance containing the mapped data, including Year and TotalBalance.
func (s *cardStatisticBalanceRecordMapper) ToYearlyBalances(cards []*db.GetYearlyBalancesRow) []*record.CardYearlyBalance {
	var records []*record.CardYearlyBalance
	for _, card := range cards {
		records = append(records, s.ToYearlyBalance(card))
	}
	return records
}
