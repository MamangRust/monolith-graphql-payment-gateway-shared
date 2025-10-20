package withdrawstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// WithdrawStatisticStatusByCardRecordMapper maps withdraw status success/failed filtered by card_number.
type WithdrawStatisticStatusByCardRecordMapper interface {
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
	ToWithdrawRecordMonthStatusSuccessCardNumber(s *db.GetMonthWithdrawStatusSuccessCardNumberRow) *record.WithdrawRecordMonthStatusSuccess
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
	ToWithdrawRecordsMonthStatusSuccessCardNumber(s []*db.GetMonthWithdrawStatusSuccessCardNumberRow) []*record.WithdrawRecordMonthStatusSuccess
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
	ToWithdrawRecordYearStatusSuccessCardNumber(s *db.GetYearlyWithdrawStatusSuccessCardNumberRow) *record.WithdrawRecordYearStatusSuccess
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
	ToWithdrawRecordsYearStatusSuccessCardNumber(s []*db.GetYearlyWithdrawStatusSuccessCardNumberRow) []*record.WithdrawRecordYearStatusSuccess

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
	ToWithdrawRecordMonthStatusFailedCardNumber(s *db.GetMonthWithdrawStatusFailedCardNumberRow) *record.WithdrawRecordMonthStatusFailed
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
	ToWithdrawRecordsMonthStatusFailedCardNumber(s []*db.GetMonthWithdrawStatusFailedCardNumberRow) []*record.WithdrawRecordMonthStatusFailed
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
	ToWithdrawRecordYearStatusFailedCardNumber(s *db.GetYearlyWithdrawStatusFailedCardNumberRow) *record.WithdrawRecordYearStatusFailed
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
	ToWithdrawRecordsYearStatusFailedCardNumber(s []*db.GetYearlyWithdrawStatusFailedCardNumberRow) []*record.WithdrawRecordYearStatusFailed
}

// WithdrawStatisticAmountByCardRecordMapper maps withdraw amount statistics filtered by card_number.
type WithdrawStatisticAmountByCardRecordMapper interface {
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
	ToWithdrawAmountMonthlyByCardNumber(s *db.GetMonthlyWithdrawsByCardNumberRow) *record.WithdrawMonthlyAmount
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
	ToWithdrawsAmountMonthlyByCardNumber(s []*db.GetMonthlyWithdrawsByCardNumberRow) []*record.WithdrawMonthlyAmount
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
	ToWithdrawAmountYearlyByCardNumber(s *db.GetYearlyWithdrawsByCardNumberRow) *record.WithdrawYearlyAmount
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
	ToWithdrawsAmountYearlyByCardNumber(s []*db.GetYearlyWithdrawsByCardNumberRow) []*record.WithdrawYearlyAmount
}
