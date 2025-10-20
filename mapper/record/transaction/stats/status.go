package transactionstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// transactionStatisticStatusRecordMapper is a struct that implements the TransactionStatisticStatusRecordMapper interface.
type transactionStatisticStatusRecordMapper struct {
}

// NewTransactionStatisticStatusRecordMapper creates and returns a new instance of the
// transactionStatisticStatusRecordMapper, which implements the TransactionStatisticStatusRecordMapper
// interface. This mapper is used to convert database rows into domain models for transaction statistics.
func NewTransactionStatisticStatusRecordMapper() TransactionStatisticStatusRecordMapper {
	return &transactionStatisticStatusRecordMapper{}
}

// ToTransactionRecordMonthStatusSuccess maps a single GetMonthTransactionStatusSuccessRow database row
// to a TransactionRecordMonthStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to GetMonthTransactionStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionStatisticStatusRecordMapper) ToTransactionRecordMonthStatusSuccess(s *db.GetMonthTransactionStatusSuccessRow) *record.TransactionRecordMonthStatusSuccess {
	return &record.TransactionRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransactionRecordsMonthStatusSuccess maps a slice of GetMonthTransactionStatusSuccessRow database rows to a slice of
// TransactionRecordMonthStatusSuccess domain models.
//
// Parameters:
//   - Transactions: A slice of pointers to GetMonthTransactionStatusSuccessRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionStatisticStatusRecordMapper) ToTransactionRecordsMonthStatusSuccess(Transactions []*db.GetMonthTransactionStatusSuccessRow) []*record.TransactionRecordMonthStatusSuccess {
	var TransactionRecords []*record.TransactionRecordMonthStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordMonthStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordYearStatusSuccess maps a GetYearlyTransactionStatusSuccessRow database row
// to a TransactionRecordYearStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to GetYearlyTransactionStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionStatisticStatusRecordMapper) ToTransactionRecordYearStatusSuccess(s *db.GetYearlyTransactionStatusSuccessRow) *record.TransactionRecordYearStatusSuccess {
	return &record.TransactionRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransactionRecordsYearStatusSuccess maps a slice of GetYearlyTransactionStatusSuccessRow database rows
// to a slice of TransactionRecordYearStatusSuccess domain models.
//
// Parameters:
//   - Transactions: A slice of pointers to GetYearlyTransactionStatusSuccessRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionStatisticStatusRecordMapper) ToTransactionRecordsYearStatusSuccess(Transactions []*db.GetYearlyTransactionStatusSuccessRow) []*record.TransactionRecordYearStatusSuccess {
	var TransactionRecords []*record.TransactionRecordYearStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordYearStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordMonthStatusFailed maps a GetMonthTransactionStatusFailedRow database row
// to a TransactionRecordMonthStatusFailed domain model.
//
// Parameters:
//   - s: A pointer to GetMonthTransactionStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transactionStatisticStatusRecordMapper) ToTransactionRecordMonthStatusFailed(s *db.GetMonthTransactionStatusFailedRow) *record.TransactionRecordMonthStatusFailed {
	return &record.TransactionRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionRecordsMonthStatusFailed maps a slice of GetMonthTransactionStatusFailedRow database rows
// to a slice of TransactionRecordMonthStatusFailed domain models.
//
// Parameters:
//   - Transactions: A slice of pointers to GetMonthTransactionStatusFailedRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transactionStatisticStatusRecordMapper) ToTransactionRecordsMonthStatusFailed(Transactions []*db.GetMonthTransactionStatusFailedRow) []*record.TransactionRecordMonthStatusFailed {
	var TransactionRecords []*record.TransactionRecordMonthStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordMonthStatusFailed(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordYearStatusFailed maps a GetYearlyTransactionStatusFailedRow database row
// to a TransactionRecordYearStatusFailed domain model.
//
// Parameters:
//   - s: A pointer to GetYearlyTransactionStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transactionStatisticStatusRecordMapper) ToTransactionRecordYearStatusFailed(s *db.GetYearlyTransactionStatusFailedRow) *record.TransactionRecordYearStatusFailed {
	return &record.TransactionRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionRecordsYearStatusFailed maps a slice of GetYearlyTransactionStatusFailedRow database rows
// to a slice of TransactionRecordYearStatusFailed domain models.
//
// Parameters:
//   - Transactions: A slice of pointers to GetYearlyTransactionStatusFailedRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transactionStatisticStatusRecordMapper) ToTransactionRecordsYearStatusFailed(Transactions []*db.GetYearlyTransactionStatusFailedRow) []*record.TransactionRecordYearStatusFailed {
	var TransactionRecords []*record.TransactionRecordYearStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordYearStatusFailed(Transaction))
	}

	return TransactionRecords
}
