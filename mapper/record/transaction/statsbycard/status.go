package transactionstatbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// transactionStatisticStatusByCardMapper provides methods to map database rows
// related to transaction status by card number to domain models.
type transactionStatisticStatusByCardMapper struct {
}

// NewTransactionStatisticStatusByCardMapper creates a new instance of
// transactionStatisticStatusByCardMapper, which provides methods to map
// database rows related to transaction status by card number to domain models.
// It returns a pointer to transactionStatisticStatusByCardMapper.
func NewTransactionStatisticStatusByCardMapper() TransactionStatisticByCardStatusMapper {
	return &transactionStatisticStatusByCardMapper{}
}

// ToTransactionRecordMonthStatusSuccessCardNumber maps a GetMonthTransactionStatusSuccessCardNumberRow database row
// to a TransactionRecordMonthStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to GetMonthTransactionStatusSuccessCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionStatisticStatusByCardMapper) ToTransactionRecordMonthStatusSuccessCardNumber(s *db.GetMonthTransactionStatusSuccessCardNumberRow) *record.TransactionRecordMonthStatusSuccess {
	return &record.TransactionRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransactionRecordsMonthStatusSuccessCardNumber maps a slice of GetMonthTransactionStatusSuccessCardNumberRow database rows
// to a slice of TransactionRecordMonthStatusSuccess domain models.
//
// Parameters:
//   - Transactions: A slice of pointers to GetMonthTransactionStatusSuccessCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionStatisticStatusByCardMapper) ToTransactionRecordsMonthStatusSuccessCardNumber(Transactions []*db.GetMonthTransactionStatusSuccessCardNumberRow) []*record.TransactionRecordMonthStatusSuccess {
	var TransactionRecords []*record.TransactionRecordMonthStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordMonthStatusSuccessCardNumber(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordYearStatusSuccessCardNumber maps a GetYearlyTransactionStatusSuccessCardNumberRow database row
// to a TransactionRecordYearStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to GetYearlyTransactionStatusSuccessCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionStatisticStatusByCardMapper) ToTransactionRecordYearStatusSuccessCardNumber(s *db.GetYearlyTransactionStatusSuccessCardNumberRow) *record.TransactionRecordYearStatusSuccess {
	return &record.TransactionRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransactionRecordsYearStatusSuccessCardNumber maps a slice of GetYearlyTransactionStatusSuccessCardNumberRow database rows
// to a slice of TransactionRecordYearStatusSuccess domain models.
//
// Parameters:
//   - Transactions: A slice of pointers to GetYearlyTransactionStatusSuccessCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionStatisticStatusByCardMapper) ToTransactionRecordsYearStatusSuccessCardNumber(Transactions []*db.GetYearlyTransactionStatusSuccessCardNumberRow) []*record.TransactionRecordYearStatusSuccess {
	var TransactionRecords []*record.TransactionRecordYearStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordYearStatusSuccessCardNumber(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordMonthStatusFailedCardNumber maps a GetMonthTransactionStatusFailedCardNumberRow
// database row to a TransactionRecordMonthStatusFailed domain model.
//
// Parameters:
//   - s: A pointer to GetMonthTransactionStatusFailedCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transactionStatisticStatusByCardMapper) ToTransactionRecordMonthStatusFailedCardNumber(s *db.GetMonthTransactionStatusFailedCardNumberRow) *record.TransactionRecordMonthStatusFailed {
	return &record.TransactionRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionRecordsMonthStatusFailedCardNumber maps a slice of GetMonthTransactionStatusFailedCardNumberRow
// database rows to a slice of TransactionRecordMonthStatusFailed domain models.
//
// Parameters:
//   - Transactions: A slice of pointers to GetMonthTransactionStatusFailedCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transactionStatisticStatusByCardMapper) ToTransactionRecordsMonthStatusFailedCardNumber(Transactions []*db.GetMonthTransactionStatusFailedCardNumberRow) []*record.TransactionRecordMonthStatusFailed {
	var TransactionRecords []*record.TransactionRecordMonthStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordMonthStatusFailedCardNumber(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordYearStatusFailedCardNumber maps a GetYearlyTransactionStatusFailedCardNumberRow
// database row to a TransactionRecordYearStatusFailed domain model.
//
// Parameters:
//   - s: A pointer to GetYearlyTransactionStatusFailedCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transactionStatisticStatusByCardMapper) ToTransactionRecordYearStatusFailedCardNumber(s *db.GetYearlyTransactionStatusFailedCardNumberRow) *record.TransactionRecordYearStatusFailed {
	return &record.TransactionRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionRecordsYearStatusFailedCardNumber maps a slice of GetYearlyTransactionStatusFailedCardNumberRow
// database rows to a slice of TransactionRecordYearStatusFailed domain models.
//
// Parameters:
//   - Transactions: A slice of pointers to GetYearlyTransactionStatusFailedCardNumberRow representing the
//     database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transactionStatisticStatusByCardMapper) ToTransactionRecordsYearStatusFailedCardNumber(Transactions []*db.GetYearlyTransactionStatusFailedCardNumberRow) []*record.TransactionRecordYearStatusFailed {
	var TransactionRecords []*record.TransactionRecordYearStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordYearStatusFailedCardNumber(Transaction))
	}

	return TransactionRecords
}
