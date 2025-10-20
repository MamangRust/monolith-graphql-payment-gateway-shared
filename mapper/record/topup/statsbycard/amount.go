package topupstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TopupStatisticAmountByCardNumberMapper maps GetMonthlyTopupAmountsByCardNumberRow database rows to TopupMonthAmount domain models.
type topupStatisticAmountByCardNumberMapper struct {
}

// NewTopupStatisticAmountByCardNumberMapper returns a new instance of TopupStatisticAmountByCardNumberMapper, which provides methods to map GetMonthlyTopupAmountsByCardNumberRow and GetYearlyTopupAmountsByCardNumberRow database rows to TopupMonthAmount and TopupYearlyAmount domain models, respectively.
func NewTopupStatisticAmountByCardNumberMapper() TopupStatisticAmountByCardNumberMapper {
	return &topupStatisticAmountByCardNumberMapper{}
}

// ToTopupMonthlyAmountByCardNumber maps a GetMonthlyTopupAmountsByCardNumberRow to a TopupMonthAmount domain model.
//
// Parameters:
//   - s: A pointer to a GetMonthlyTopupAmountsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TopupMonthAmount containing the mapped data, including Month and TotalAmount.
func (t *topupStatisticAmountByCardNumberMapper) ToTopupMonthlyAmountByCardNumber(s *db.GetMonthlyTopupAmountsByCardNumberRow) *record.TopupMonthAmount {
	return &record.TopupMonthAmount{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyAmountsByCardNumber maps a slice of GetMonthlyTopupAmountsByCardNumberRow database rows to a slice of TopupMonthAmount domain models.
// It iterates over the provided slice, converting each element using the ToTopupMonthlyAmountByCardNumber method,
// and appends the result to a new slice. The function returns a slice of pointers to TopupMonthAmount.
func (t *topupStatisticAmountByCardNumberMapper) ToTopupMonthlyAmountsByCardNumber(s []*db.GetMonthlyTopupAmountsByCardNumberRow) []*record.TopupMonthAmount {
	var topupRecords []*record.TopupMonthAmount

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupMonthlyAmountByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupYearlyAmountByCardNumber maps a GetYearlyTopupAmountsByCardNumberRow to a TopupYearlyAmount domain model.
//
// Parameters:
//   - s: A pointer to a GetYearlyTopupAmountsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TopupYearlyAmount containing the mapped data, including Year and TotalAmount.
func (t *topupStatisticAmountByCardNumberMapper) ToTopupYearlyAmountByCardNumber(s *db.GetYearlyTopupAmountsByCardNumberRow) *record.TopupYearlyAmount {
	return &record.TopupYearlyAmount{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyAmountsByCardNumber maps a slice of GetYearlyTopupAmountsByCardNumberRow to a
// slice of TopupYearlyAmount domain models.
//
// Parameters:
//   - s: A slice of pointers to GetYearlyTopupAmountsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TopupYearlyAmount containing the mapped data, including Year and TotalAmount.
func (t *topupStatisticAmountByCardNumberMapper) ToTopupYearlyAmountsByCardNumber(s []*db.GetYearlyTopupAmountsByCardNumberRow) []*record.TopupYearlyAmount {
	var topupRecords []*record.TopupYearlyAmount

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupYearlyAmountByCardNumber(topup))
	}

	return topupRecords
}
