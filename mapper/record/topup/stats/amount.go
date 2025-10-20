package topupstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// topupStatisticAmountRecordMapper provides methods to map GetMonthlyTopupAmountsRow database rows to TopupMonthAmount domain models.
type topupStatisticAmountRecordMapper struct {
}

// NewTopupStatisticAmountRecordMapper creates a new instance of topupStatisticAmountRecordMapper.
func NewTopupStatisticAmountRecordMapper() TopupStatisticAmountMapper {
	return &topupStatisticAmountRecordMapper{}
}

// ToTopupMonthlyAmount maps a GetMonthlyTopupAmountsRow database row to a TopupMonthAmount domain model.
// It takes a pointer to a GetMonthlyTopupAmountsRow as a parameter and returns a pointer to a TopupMonthAmount.
// The function maps the fields of GetMonthlyTopupAmountsRow to the corresponding fields of TopupMonthAmount.
func (t *topupStatisticAmountRecordMapper) ToTopupMonthlyAmount(s *db.GetMonthlyTopupAmountsRow) *record.TopupMonthAmount {
	return &record.TopupMonthAmount{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyAmounts maps a slice of GetMonthlyTopupAmountsRow database rows to a slice of TopupMonthAmount domain models.
// It iterates over the provided slice, converting each element using the ToTopupMonthlyAmount method,
// and appends the result to a new slice. The function returns a slice of pointers to TopupMonthAmount.
func (t *topupStatisticAmountRecordMapper) ToTopupMonthlyAmounts(s []*db.GetMonthlyTopupAmountsRow) []*record.TopupMonthAmount {
	var topupRecords []*record.TopupMonthAmount

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupMonthlyAmount(topup))
	}

	return topupRecords
}

// ToTopupYearlyAmount maps a GetYearlyTopupAmountsRow database row to a TopupYearlyAmount domain model.
// It takes a pointer to a GetYearlyTopupAmountsRow as a parameter and returns a pointer to a TopupYearlyAmount.
// The function maps the fields of GetYearlyTopupAmountsRow to the corresponding fields of TopupYearlyAmount.
func (t *topupStatisticAmountRecordMapper) ToTopupYearlyAmount(s *db.GetYearlyTopupAmountsRow) *record.TopupYearlyAmount {
	return &record.TopupYearlyAmount{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyAmounts maps a slice of GetYearlyTopupAmountsRow database rows to a slice of TopupYearlyAmount domain models.
// It iterates over the provided slice, converting each element using the ToTopupYearlyAmount method,
// and appends the result to a new slice. The function returns a slice of pointers to TopupYearlyAmount.
func (t *topupStatisticAmountRecordMapper) ToTopupYearlyAmounts(s []*db.GetYearlyTopupAmountsRow) []*record.TopupYearlyAmount {
	var topupRecords []*record.TopupYearlyAmount

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupYearlyAmount(topup))
	}

	return topupRecords
}
