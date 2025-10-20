package transactionservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type transactionStatsStatusResponseMapper struct{}

func NewTransactionStatsStatusResponseMapper() TransactionStatsStatusResponseMapper {
	return &transactionStatsStatusResponseMapper{}
}

// ToTransactionResponseMonthStatusSuccess maps a record of monthly transaction success status
// into a TransactionResponseMonthStatusSuccess.
//
// Args:
//   - s: A pointer to TransactionRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionResponseMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionStatsStatusResponseMapper) ToTransactionResponseMonthStatusSuccess(s *record.TransactionRecordMonthStatusSuccess) *response.TransactionResponseMonthStatusSuccess {
	return &response.TransactionResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTransactionResponsesMonthStatusSuccess converts multiple records of monthly transaction success status into a slice of
// TransactionResponseMonthStatusSuccess.
//
// Args:
//   - Transactions: A slice of pointers to TransactionRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponseMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionStatsStatusResponseMapper) ToTransactionResponsesMonthStatusSuccess(Transactions []*record.TransactionRecordMonthStatusSuccess) []*response.TransactionResponseMonthStatusSuccess {
	var TransactionRecords []*response.TransactionResponseMonthStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionResponseMonthStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// ToTransactionResponseYearStatusSuccess maps a record of yearly transaction success status
// into a TransactionResponseYearStatusSuccess.
//
// Args:
//   - s: A pointer to TransactionRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionResponseYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionStatsStatusResponseMapper) ToTransactionResponseYearStatusSuccess(s *record.TransactionRecordYearStatusSuccess) *response.TransactionResponseYearStatusSuccess {
	return &response.TransactionResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTransactionResponsesYearStatusSuccess converts multiple records of yearly transaction success status into a slice of
// TransactionResponseYearStatusSuccess.
//
// Args:
//   - Transactions: A slice of pointers to TransactionRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponseYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionStatsStatusResponseMapper) ToTransactionResponsesYearStatusSuccess(Transactions []*record.TransactionRecordYearStatusSuccess) []*response.TransactionResponseYearStatusSuccess {
	var TransactionRecords []*response.TransactionResponseYearStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionResponseYearStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// ToTransactionResponseMonthStatusFailed maps a record of monthly transaction failed status into a TransactionResponseMonthStatusFailed.
//
// Args:
//   - s: A pointer to TransactionRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionResponseMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transactionStatsStatusResponseMapper) ToTransactionResponseMonthStatusFailed(s *record.TransactionRecordMonthStatusFailed) *response.TransactionResponseMonthStatusFailed {
	return &response.TransactionResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTransactionResponsesMonthStatusFailed converts multiple records of monthly transaction failed status into a slice of
// TransactionResponseMonthStatusFailed.
//
// Args:
//   - Transactions: A slice of pointers to TransactionRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponseMonthStatusFailed containing the mapped data, including Year, Month, TotalFailed, and TotalAmount.
func (t *transactionStatsStatusResponseMapper) ToTransactionResponsesMonthStatusFailed(Transactions []*record.TransactionRecordMonthStatusFailed) []*response.TransactionResponseMonthStatusFailed {
	var TransactionRecords []*response.TransactionResponseMonthStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionResponseMonthStatusFailed(Transaction))
	}

	return TransactionRecords
}

// ToTransactionResponseYearStatusFailed maps a yearly transaction failed status record
// into a TransactionResponseYearStatusFailed.
//
// Args:
//   - s: A pointer to TransactionRecordYearStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *transactionStatsStatusResponseMapper) ToTransactionResponseYearStatusFailed(s *record.TransactionRecordYearStatusFailed) *response.TransactionResponseYearStatusFailed {
	return &response.TransactionResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTransactionResponsesYearStatusFailed converts multiple yearly transaction failed status records
// into a slice of TransactionResponseYearStatusFailed.
//
// Args:
//   - Transactions: A slice of pointers to TransactionRecordYearStatusFailed containing the records to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *transactionStatsStatusResponseMapper) ToTransactionResponsesYearStatusFailed(Transactions []*record.TransactionRecordYearStatusFailed) []*response.TransactionResponseYearStatusFailed {
	var TransactionRecords []*response.TransactionResponseYearStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionResponseYearStatusFailed(Transaction))
	}

	return TransactionRecords
}
