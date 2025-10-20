package withdrawstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// WithdrawStatisticStatusRecordMapper is an interface that provides methods to map
// database rows related to withdraw status (success/failed) into domain models.
type withdrawStatisticStatusRecordMapper struct{}

// NewWithdrawStatisticStatusRecordMapper creates and returns a new instance of
// WithdrawStatisticStatusRecordMapper, which provides methods to map database
// rows related to withdraw status (success/failed) into domain models.
func NewWithdrawStatisticStatusRecordMapper() WithdrawStatisticStatusRecordMapper {
	return &withdrawStatisticStatusRecordMapper{}
}

// ToWithdrawRecordMonthStatusSuccess maps a database row representing monthly
// successful withdraw statistics to a WithdrawRecordMonthStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to a GetMonthWithdrawStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordMonthStatusSuccess containing the mapped data,
//     including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawStatisticStatusRecordMapper) ToWithdrawRecordMonthStatusSuccess(s *db.GetMonthWithdrawStatusSuccessRow) *record.WithdrawRecordMonthStatusSuccess {
	return &record.WithdrawRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToWithdrawRecordsMonthStatusSuccess converts a slice of GetMonthWithdrawStatusSuccessRow
// database rows to a slice of WithdrawRecordMonthStatusSuccess domain models.
//
// Parameters:
//   - Withdraws: A slice of pointers to GetMonthWithdrawStatusSuccessRow representing
//     the database rows containing monthly successful withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawRecordMonthStatusSuccess, each containing
//     the mapped data including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawStatisticStatusRecordMapper) ToWithdrawRecordsMonthStatusSuccess(Withdraws []*db.GetMonthWithdrawStatusSuccessRow) []*record.WithdrawRecordMonthStatusSuccess {
	var WithdrawRecords []*record.WithdrawRecordMonthStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordMonthStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordYearStatusSuccess maps a database row representing yearly
// successful withdraw statistics to a WithdrawRecordYearStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to a GetYearlyWithdrawStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordYearStatusSuccess containing the mapped data,
//     including Year, TotalSuccess, and TotalAmount.
func (t *withdrawStatisticStatusRecordMapper) ToWithdrawRecordYearStatusSuccess(s *db.GetYearlyWithdrawStatusSuccessRow) *record.WithdrawRecordYearStatusSuccess {
	return &record.WithdrawRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToWithdrawRecordsYearStatusSuccess maps a slice of GetYearlyWithdrawStatusSuccessRow
// database rows to a slice of WithdrawRecordYearStatusSuccess domain models.
//
// Parameters:
//   - Withdraws: A slice of pointers to GetYearlyWithdrawStatusSuccessRow representing
//     the database rows containing yearly successful withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawRecordYearStatusSuccess, each containing
//     the mapped data including Year, TotalSuccess, and TotalAmount.
func (t *withdrawStatisticStatusRecordMapper) ToWithdrawRecordsYearStatusSuccess(Withdraws []*db.GetYearlyWithdrawStatusSuccessRow) []*record.WithdrawRecordYearStatusSuccess {
	var WithdrawRecords []*record.WithdrawRecordYearStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordYearStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordMonthStatusFailed maps a database row representing monthly
// failed withdraw statistics to a WithdrawRecordMonthStatusFailed domain model.
//
// Parameters:
//   - s: A pointer to a GetMonthWithdrawStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordMonthStatusFailed containing the mapped data,
//     including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawStatisticStatusRecordMapper) ToWithdrawRecordMonthStatusFailed(s *db.GetMonthWithdrawStatusFailedRow) *record.WithdrawRecordMonthStatusFailed {
	return &record.WithdrawRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToWithdrawRecordsMonthStatusFailed maps a slice of GetMonthWithdrawStatusFailedRow
// database rows to a slice of WithdrawRecordMonthStatusFailed domain models.
//
// Parameters:
//   - Withdraws: A slice of pointers to GetMonthWithdrawStatusFailedRow representing
//     the database rows containing monthly failed withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawRecordMonthStatusFailed, each containing
//     the mapped data including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawStatisticStatusRecordMapper) ToWithdrawRecordsMonthStatusFailed(Withdraws []*db.GetMonthWithdrawStatusFailedRow) []*record.WithdrawRecordMonthStatusFailed {
	var WithdrawRecords []*record.WithdrawRecordMonthStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordMonthStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordYearStatusFailed maps a database row representing yearly
// failed withdraw statistics to a WithdrawRecordYearStatusFailed domain model.
//
// Parameters:
//   - s: A pointer to a GetYearlyWithdrawStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordYearStatusFailed containing the mapped data,
//     including Year, TotalFailed, and TotalAmount.
func (t *withdrawStatisticStatusRecordMapper) ToWithdrawRecordYearStatusFailed(s *db.GetYearlyWithdrawStatusFailedRow) *record.WithdrawRecordYearStatusFailed {
	return &record.WithdrawRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToWithdrawRecordsYearStatusFailed maps a slice of GetYearlyWithdrawStatusFailedRow
// database rows to a slice of WithdrawRecordYearStatusFailed domain models.
//
// Parameters:
//   - Withdraws: A slice of pointers to GetYearlyWithdrawStatusFailedRow representing
//     the database rows containing yearly failed withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawRecordYearStatusFailed, each containing
//     the mapped data including Year, TotalFailed, and TotalAmount.
func (t *withdrawStatisticStatusRecordMapper) ToWithdrawRecordsYearStatusFailed(Withdraws []*db.GetYearlyWithdrawStatusFailedRow) []*record.WithdrawRecordYearStatusFailed {
	var WithdrawRecords []*record.WithdrawRecordYearStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordYearStatusFailed(Withdraw))
	}

	return WithdrawRecords
}
