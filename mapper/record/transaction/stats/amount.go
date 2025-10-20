package transactionstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

type transactionStatisticAmountRecordMapper struct {
}

func NewTransactionStatisticAmountRecordMapper() TransactionStatisticAmountRecordMapper {
	return &transactionStatisticAmountRecordMapper{}
}

// ToTransactionMonthlyAmount maps a GetMonthlyAmountsRow database row to a TransactionMonthAmount
// domain model.
//
// Parameters:
//   - ss: A pointer to GetMonthlyAmountsRow representing the database row.
//
// Returns:
//   - A pointer to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transactionStatisticAmountRecordMapper) ToTransactionMonthlyAmount(ss *db.GetMonthlyAmountsRow) *record.TransactionMonthAmount {
	return &record.TransactionMonthAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalAmount),
	}
}

// ToTransactionMonthlyAmounts maps a slice of GetMonthlyAmountsRow database rows
// to a slice of TransactionMonthAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetMonthlyAmountsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transactionStatisticAmountRecordMapper) ToTransactionMonthlyAmounts(ss []*db.GetMonthlyAmountsRow) []*record.TransactionMonthAmount {
	var transactionRecords []*record.TransactionMonthAmount
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionMonthlyAmount(transaction))
	}
	return transactionRecords
}

// ToTransactionYearlyAmount maps a GetYearlyAmountsRow database row to a TransactionYearlyAmount
// domain model.
//
// Parameters:
//   - ss: A pointer to GetYearlyAmountsRow representing the database row.
//
// Returns:
//   - A pointer to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
func (s *transactionStatisticAmountRecordMapper) ToTransactionYearlyAmount(ss *db.GetYearlyAmountsRow) *record.TransactionYearlyAmount {
	return &record.TransactionYearlyAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalAmount),
	}
}

// ToTransactionYearlyAmounts maps a slice of GetYearlyAmountsRow database rows
// to a slice of TransactionYearlyAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetYearlyAmountsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
func (s *transactionStatisticAmountRecordMapper) ToTransactionYearlyAmounts(ss []*db.GetYearlyAmountsRow) []*record.TransactionYearlyAmount {
	var transactionRecords []*record.TransactionYearlyAmount
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionYearlyAmount(transaction))
	}
	return transactionRecords
}
