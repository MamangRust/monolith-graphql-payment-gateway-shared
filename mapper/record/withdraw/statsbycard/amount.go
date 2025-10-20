package withdrawstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// withdrawStatisticAmountByCardRecordMapper provides methods to map database rows
// related to withdraw statistics filtered by card number to domain models.
type withdrawStatisticAmountByCardRecordMapper struct{}

// NewWithdrawStatisticAmountByCardRecordMapper creates and returns a new instance
// of WithdrawStatisticAmountByCardRecordMapper, which provides methods to map
// database rows related to withdraw statistics filtered by card number to
// domain models.
func NewWithdrawStatisticAmountByCardRecordMapper() *withdrawStatisticAmountByCardRecordMapper {
	return &withdrawStatisticAmountByCardRecordMapper{}
}

// ToWithdrawAmountMonthlyByCardNumber maps a database row representing monthly
// withdraw statistics, filtered by card number, to a WithdrawMonthlyAmount
// domain model.
//
// Parameters:
//   - ss: A pointer to a GetMonthlyWithdrawsByCardNumberRow representing the
//     database row with monthly withdraw statistics filtered by card number.
//
// Returns:
//   - A pointer to a WithdrawMonthlyAmount containing the mapped data,
//     including Month and TotalAmount.
func (r *withdrawStatisticAmountByCardRecordMapper) ToWithdrawAmountMonthlyByCardNumber(ss *db.GetMonthlyWithdrawsByCardNumberRow) *record.WithdrawMonthlyAmount {
	return &record.WithdrawMonthlyAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalWithdrawAmount),
	}
}

// ToWithdrawsAmountMonthlyByCardNumber maps a slice of
// GetMonthlyWithdrawsByCardNumberRow database rows to a slice of
// WithdrawMonthlyAmount domain models, filtered by card number.
//
// Parameters:
//   - ss: A slice of pointers to GetMonthlyWithdrawsByCardNumberRow
//     representing the database rows with monthly withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawMonthlyAmount, each containing the
//     mapped data including Month and TotalAmount.
func (s *withdrawStatisticAmountByCardRecordMapper) ToWithdrawsAmountMonthlyByCardNumber(ss []*db.GetMonthlyWithdrawsByCardNumberRow) []*record.WithdrawMonthlyAmount {
	var withdrawRecords []*record.WithdrawMonthlyAmount

	for _, withdraw := range ss {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawAmountMonthlyByCardNumber(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawAmountYearlyByCardNumber maps a database row representing yearly
// withdraw statistics, filtered by card number, to a WithdrawYearlyAmount
// domain model.
//
// Parameters:
//   - ss: A pointer to a GetYearlyWithdrawsByCardNumberRow representing the
//     database row with yearly withdraw statistics filtered by card number.
//
// Returns:
//   - A pointer to a WithdrawYearlyAmount containing the mapped data,
//     including Year and TotalAmount.
func (r *withdrawStatisticAmountByCardRecordMapper) ToWithdrawAmountYearlyByCardNumber(ss *db.GetYearlyWithdrawsByCardNumberRow) *record.WithdrawYearlyAmount {
	return &record.WithdrawYearlyAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalWithdrawAmount),
	}
}

// ToWithdrawsAmountYearlyByCardNumber maps a slice of
// GetYearlyWithdrawsByCardNumberRow database rows to a slice of
// WithdrawYearlyAmount domain models, filtered by card number.
//
// Parameters:
//   - ss: A slice of pointers to GetYearlyWithdrawsByCardNumberRow
//     representing the database rows with yearly withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawYearlyAmount, each containing the
//     mapped data including Year and TotalAmount.
func (s *withdrawStatisticAmountByCardRecordMapper) ToWithdrawsAmountYearlyByCardNumber(ss []*db.GetYearlyWithdrawsByCardNumberRow) []*record.WithdrawYearlyAmount {
	var withdrawRecords []*record.WithdrawYearlyAmount

	for _, withdraw := range ss {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawAmountYearlyByCardNumber(withdraw))
	}

	return withdrawRecords
}
