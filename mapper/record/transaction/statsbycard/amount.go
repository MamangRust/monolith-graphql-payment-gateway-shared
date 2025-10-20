package transactionstatbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// transactionStatisticAmountByCardMapper is a struct that implements the TransactionStatisticAmountByCardMapper interface.
type transactionStatisticAmountByCardMapper struct {
}

// NewTransactionStatisticAmountByCardMapper creates a new instance of
// transactionStatisticAmountByCardMapper, which provides methods to map
// database rows related to transaction amounts by card number to domain models.
// It returns a pointer to transactionStatisticAmountByCardMapper.
func NewTransactionStatisticAmountByCardMapper() *transactionStatisticAmountByCardMapper {
	return &transactionStatisticAmountByCardMapper{}
}

// ToTransactionMonthlyAmountByCardNumber maps a GetMonthlyAmountsByCardNumberRow database row
// to a TransactionMonthAmount domain model.
//
// Parameters:
//   - ss: A pointer to GetMonthlyAmountsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transactionStatisticAmountByCardMapper) ToTransactionMonthlyAmountByCardNumber(ss *db.GetMonthlyAmountsByCardNumberRow) *record.TransactionMonthAmount {
	return &record.TransactionMonthAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalAmount),
	}
}

// ToTransactionMonthlyAmountsByCardNumber maps a slice of GetMonthlyAmountsByCardNumberRow database rows
// to a slice of TransactionMonthAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetMonthlyAmountsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transactionStatisticAmountByCardMapper) ToTransactionMonthlyAmountsByCardNumber(ss []*db.GetMonthlyAmountsByCardNumberRow) []*record.TransactionMonthAmount {
	var transactionRecords []*record.TransactionMonthAmount
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionMonthlyAmountByCardNumber(transaction))
	}
	return transactionRecords

}

// ToTransactionYearlyAmountByCardNumber maps a GetYearlyAmountsByCardNumberRow database row
// to a TransactionYearlyAmount domain model.
//
// Parameters:
//   - ss: A pointer to GetYearlyAmountsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
func (s *transactionStatisticAmountByCardMapper) ToTransactionYearlyAmountByCardNumber(ss *db.GetYearlyAmountsByCardNumberRow) *record.TransactionYearlyAmount {
	return &record.TransactionYearlyAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalAmount),
	}
}

// ToTransactionYearlyAmountsByCardNumber maps a slice of GetYearlyAmountsByCardNumberRow database rows
// to a slice of TransactionYearlyAmount domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetYearlyAmountsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
func (s *transactionStatisticAmountByCardMapper) ToTransactionYearlyAmountsByCardNumber(ss []*db.GetYearlyAmountsByCardNumberRow) []*record.TransactionYearlyAmount {
	var transactionRecords []*record.TransactionYearlyAmount
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionYearlyAmountByCardNumber(transaction))
	}
	return transactionRecords
}
