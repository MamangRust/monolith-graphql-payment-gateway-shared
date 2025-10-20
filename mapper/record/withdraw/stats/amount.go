package withdrawstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// WithdrawStatisticAmountRecordMapper is an interface that provides methods to map
type withdrawStatisticAmountRecordMapper struct{}

// NewWithdrawStatisticAmountRecordMapper returns a new instance of WithdrawStatisticAmountRecordMapper
// which provides methods to map GetMonthlyWithdrawsRow and GetYearlyWithdrawsRow
// database rows to WithdrawMonthlyAmount and WithdrawYearlyAmount domain models.
func NewWithdrawStatisticAmountRecordMapper() WithdrawStatisticAmountRecordMapper {
	return &withdrawStatisticAmountRecordMapper{}
}

// ToWithdrawAmountMonthly maps a single database row to a WithdrawMonthlyAmount
// domain model. It is intended for use with database rows that contain monthly
// withdraw statistics.
//
// Parameters:
//   - ss: A pointer to a GetMonthlyWithdrawsRow representing the database row
//     with monthly withdraw statistics.
//
// Returns:
//   - A pointer to a WithdrawMonthlyAmount containing the mapped data,
//     including Month and TotalAmount.
func (r *withdrawStatisticAmountRecordMapper) ToWithdrawAmountMonthly(ss *db.GetMonthlyWithdrawsRow) *record.WithdrawMonthlyAmount {
	return &record.WithdrawMonthlyAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalWithdrawAmount),
	}
}

// ToWithdrawsAmountMonthly maps a slice of GetMonthlyWithdrawsRow database rows
// to a slice of WithdrawMonthlyAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetMonthlyWithdrawsRow representing the database
//     rows with monthly withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawMonthlyAmount, each containing the mapped
//     data including Month and TotalAmount.
func (s *withdrawStatisticAmountRecordMapper) ToWithdrawsAmountMonthly(ss []*db.GetMonthlyWithdrawsRow) []*record.WithdrawMonthlyAmount {
	var withdrawRecords []*record.WithdrawMonthlyAmount

	for _, withdraw := range ss {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawAmountMonthly(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawAmountYearly maps a single database row to a WithdrawYearlyAmount
// domain model. It is intended for use with database rows that contain yearly
// withdraw statistics.
//
// Parameters:
//   - ss: A pointer to a GetYearlyWithdrawsRow representing the database row
//     with yearly withdraw statistics.
//
// Returns:
//   - A pointer to a WithdrawYearlyAmount containing the mapped data,
//     including Year and TotalAmount.
func (r *withdrawStatisticAmountRecordMapper) ToWithdrawAmountYearly(ss *db.GetYearlyWithdrawsRow) *record.WithdrawYearlyAmount {
	return &record.WithdrawYearlyAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalWithdrawAmount),
	}
}

// ToWithdrawsAmountYearly maps a slice of GetYearlyWithdrawsRow database rows to a slice
// of WithdrawYearlyAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetYearlyWithdrawsRow representing the database rows
//     with yearly withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawYearlyAmount, each containing the mapped
//     data including Year and TotalAmount.
func (s *withdrawStatisticAmountRecordMapper) ToWithdrawsAmountYearly(ss []*db.GetYearlyWithdrawsRow) []*record.WithdrawYearlyAmount {
	var withdrawRecords []*record.WithdrawYearlyAmount

	for _, withdraw := range ss {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawAmountYearly(withdraw))
	}

	return withdrawRecords
}
