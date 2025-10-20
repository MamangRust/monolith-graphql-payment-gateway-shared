package withdrawstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// withdrawStatisticStatusByCardRecordMapper provides methods to map database rows
// related to withdraw status (success/failed) filtered by card number to domain models.
type withdrawStatisticStatusByCardRecordMapper struct {
}

// NewWithdrawStatisticStatusByCardRecordMapper creates and returns a new instance of
// WithdrawStatisticStatusByCardRecordMapper, which provides methods to map database
// rows related to withdraw status (success/failed) filtered by card number to domain
// models.
func NewWithdrawStatisticStatusByCardRecordMapper() *withdrawStatisticStatusByCardRecordMapper {
	return &withdrawStatisticStatusByCardRecordMapper{}
}

// ToWithdrawRecordMonthStatusSuccessCardNumber maps a database row representing
// monthly successful withdraw statistics, filtered by card number, to a
// WithdrawRecordMonthStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to a GetMonthWithdrawStatusSuccessCardNumberRow representing
//     the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordMonthStatusSuccess containing the mapped data,
//     including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawStatisticStatusByCardRecordMapper) ToWithdrawRecordMonthStatusSuccessCardNumber(s *db.GetMonthWithdrawStatusSuccessCardNumberRow) *record.WithdrawRecordMonthStatusSuccess {
	return &record.WithdrawRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToWithdrawRecordsMonthStatusSuccessCardNumber maps a slice of
// GetMonthWithdrawStatusSuccessCardNumberRow database rows to a slice of
// WithdrawRecordMonthStatusSuccess domain models, filtered by card number.
//
// Parameters:
//   - Withdraws: A slice of pointers to GetMonthWithdrawStatusSuccessCardNumberRow
//     representing the database rows with monthly successful withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawRecordMonthStatusSuccess, each containing
//     the mapped data including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawStatisticStatusByCardRecordMapper) ToWithdrawRecordsMonthStatusSuccessCardNumber(Withdraws []*db.GetMonthWithdrawStatusSuccessCardNumberRow) []*record.WithdrawRecordMonthStatusSuccess {
	var WithdrawRecords []*record.WithdrawRecordMonthStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordMonthStatusSuccessCardNumber(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordYearStatusSuccessCardNumber maps a database row representing yearly
// successful withdraw statistics, filtered by card number, to a
// WithdrawRecordYearStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to a GetYearlyWithdrawStatusSuccessCardNumberRow representing
//     the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordYearStatusSuccess containing the mapped data,
//     including Year, TotalSuccess, and TotalAmount.
func (t *withdrawStatisticStatusByCardRecordMapper) ToWithdrawRecordYearStatusSuccessCardNumber(s *db.GetYearlyWithdrawStatusSuccessCardNumberRow) *record.WithdrawRecordYearStatusSuccess {
	return &record.WithdrawRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToWithdrawRecordsYearStatusSuccessCardNumber maps a slice of
// GetYearlyWithdrawStatusSuccessCardNumberRow database rows to a slice of
// WithdrawRecordYearStatusSuccess domain models, filtered by card number.
//
// Parameters:
//   - Withdraws: A slice of pointers to GetYearlyWithdrawStatusSuccessCardNumberRow
//     representing the database rows with yearly successful withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawRecordYearStatusSuccess, each containing
//     the mapped data including Year, TotalSuccess, and TotalAmount.
func (t *withdrawStatisticStatusByCardRecordMapper) ToWithdrawRecordsYearStatusSuccessCardNumber(Withdraws []*db.GetYearlyWithdrawStatusSuccessCardNumberRow) []*record.WithdrawRecordYearStatusSuccess {
	var WithdrawRecords []*record.WithdrawRecordYearStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordYearStatusSuccessCardNumber(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordMonthStatusFailedCardNumber maps a GetMonthWithdrawStatusFailedCardNumberRow database row
// to a WithdrawRecordMonthStatusFailed domain model, filtered by card number.
//
// Parameters:
//   - s: A pointer to a GetMonthWithdrawStatusFailedCardNumberRow representing the
//     database row.
//
// Returns:
//   - A pointer to a WithdrawRecordMonthStatusFailed containing the mapped data,
//     including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawStatisticStatusByCardRecordMapper) ToWithdrawRecordMonthStatusFailedCardNumber(s *db.GetMonthWithdrawStatusFailedCardNumberRow) *record.WithdrawRecordMonthStatusFailed {
	return &record.WithdrawRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToWithdrawRecordsMonthStatusFailedCardNumber maps a slice of
// GetMonthWithdrawStatusFailedCardNumberRow database rows to a slice of
// WithdrawRecordMonthStatusFailed domain models, filtered by card number.
//
// Parameters:
//   - Withdraws: A slice of pointers to GetMonthWithdrawStatusFailedCardNumberRow
//     representing the database rows with monthly failed withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawRecordMonthStatusFailed, each containing
//     the mapped data including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawStatisticStatusByCardRecordMapper) ToWithdrawRecordsMonthStatusFailedCardNumber(Withdraws []*db.GetMonthWithdrawStatusFailedCardNumberRow) []*record.WithdrawRecordMonthStatusFailed {
	var WithdrawRecords []*record.WithdrawRecordMonthStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordMonthStatusFailedCardNumber(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordYearStatusFailedCardNumber maps a single
// GetYearlyWithdrawStatusFailedCardNumberRow database row to a
// WithdrawRecordYearStatusFailed domain model, filtered by card number.
//
// Parameters:
//   - s: A pointer to a GetYearlyWithdrawStatusFailedCardNumberRow
//     representing the database row with yearly failed withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A pointer to a WithdrawRecordYearStatusFailed containing the mapped
//     data, including Year, TotalFailed, and TotalAmount.
func (t *withdrawStatisticStatusByCardRecordMapper) ToWithdrawRecordYearStatusFailedCardNumber(s *db.GetYearlyWithdrawStatusFailedCardNumberRow) *record.WithdrawRecordYearStatusFailed {
	return &record.WithdrawRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToWithdrawRecordsYearStatusFailedCardNumber maps a slice of
// GetYearlyWithdrawStatusFailedCardNumberRow database rows to a slice of
// WithdrawRecordYearStatusFailed domain models, filtered by card number.
//
// Parameters:
//   - Withdraws: A slice of pointers to GetYearlyWithdrawStatusFailedCardNumberRow
//     representing the database rows with yearly failed withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawRecordYearStatusFailed, each containing
//     the mapped data including Year, TotalFailed, and TotalAmount.
func (t *withdrawStatisticStatusByCardRecordMapper) ToWithdrawRecordsYearStatusFailedCardNumber(Withdraws []*db.GetYearlyWithdrawStatusFailedCardNumberRow) []*record.WithdrawRecordYearStatusFailed {
	var WithdrawRecords []*record.WithdrawRecordYearStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordYearStatusFailedCardNumber(Withdraw))
	}

	return WithdrawRecords
}
