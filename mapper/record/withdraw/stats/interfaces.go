package withdrawstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// WithdrawStatisticStatusRecordMapper maps withdraw status success/failed (monthly & yearly).
type WithdrawStatisticStatusRecordMapper interface {
	// ToWithdrawRecordMonthStatusSuccess maps a database row representing monthly
	// successful withdraw statistics to a WithdrawRecordMonthStatusSuccess domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetMonthWithdrawStatusSuccessRow representing the database row.
	//
	// Returns:
	//   - A pointer to a WithdrawRecordMonthStatusSuccess containing the mapped data,
	//     including Year, Month, TotalSuccess, and TotalAmount.
	ToWithdrawRecordMonthStatusSuccess(s *db.GetMonthWithdrawStatusSuccessRow) *record.WithdrawRecordMonthStatusSuccess
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
	ToWithdrawRecordsMonthStatusSuccess(s []*db.GetMonthWithdrawStatusSuccessRow) []*record.WithdrawRecordMonthStatusSuccess
	// ToWithdrawRecordYearStatusSuccess maps a database row representing yearly
	// successful withdraw statistics to a WithdrawRecordYearStatusSuccess domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetYearlyWithdrawStatusSuccessRow representing the database row.
	//
	// Returns:
	//   - A pointer to a WithdrawRecordYearStatusSuccess containing the mapped data,
	//     including Year, TotalSuccess, and TotalAmount.
	ToWithdrawRecordYearStatusSuccess(s *db.GetYearlyWithdrawStatusSuccessRow) *record.WithdrawRecordYearStatusSuccess
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
	ToWithdrawRecordsYearStatusSuccess(s []*db.GetYearlyWithdrawStatusSuccessRow) []*record.WithdrawRecordYearStatusSuccess

	// ToWithdrawRecordMonthStatusFailed maps a database row representing monthly
	// failed withdraw statistics to a WithdrawRecordMonthStatusFailed domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetMonthWithdrawStatusFailedRow representing the database row.
	//
	// Returns:
	//   - A pointer to a WithdrawRecordMonthStatusFailed containing the mapped data,
	//     including Year, Month, TotalFailed, and TotalAmount.
	ToWithdrawRecordMonthStatusFailed(s *db.GetMonthWithdrawStatusFailedRow) *record.WithdrawRecordMonthStatusFailed
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
	ToWithdrawRecordsMonthStatusFailed(s []*db.GetMonthWithdrawStatusFailedRow) []*record.WithdrawRecordMonthStatusFailed
	// ToWithdrawRecordYearStatusFailed maps a database row representing yearly
	// failed withdraw statistics to a WithdrawRecordYearStatusFailed domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetYearlyWithdrawStatusFailedRow representing the database row.
	//
	// Returns:
	//   - A pointer to a WithdrawRecordYearStatusFailed containing the mapped data,
	//     including Year, TotalFailed, and TotalAmount.
	ToWithdrawRecordYearStatusFailed(s *db.GetYearlyWithdrawStatusFailedRow) *record.WithdrawRecordYearStatusFailed
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
	ToWithdrawRecordsYearStatusFailed(s []*db.GetYearlyWithdrawStatusFailedRow) []*record.WithdrawRecordYearStatusFailed
}

// WithdrawStatisticAmountRecordMapper maps withdraw amount statistics (monthly & yearly).
type WithdrawStatisticAmountRecordMapper interface {
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
	ToWithdrawAmountMonthly(s *db.GetMonthlyWithdrawsRow) *record.WithdrawMonthlyAmount
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
	ToWithdrawsAmountMonthly(s []*db.GetMonthlyWithdrawsRow) []*record.WithdrawMonthlyAmount
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
	ToWithdrawAmountYearly(s *db.GetYearlyWithdrawsRow) *record.WithdrawYearlyAmount
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
	ToWithdrawsAmountYearly(s []*db.GetYearlyWithdrawsRow) []*record.WithdrawYearlyAmount
}
