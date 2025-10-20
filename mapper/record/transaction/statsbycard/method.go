package transactionstatbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// transactionStatisticMethodByCardMapper is a struct that implements the TransactionStatisticMethodByCardMapper interface.
type transactionStatisticMethodByCardMapper struct {
}

// NewTransactionStatisticMethodByCardMapper creates a new instance of
// transactionStatisticMethodByCardMapper, which provides methods to map
// database rows related to transaction methods by card number to domain models.
// It returns a pointer to transactionStatisticMethodByCardMapper.
func NewTransactionStatisticMethodByCardMapper() *transactionStatisticMethodByCardMapper {
	return &transactionStatisticMethodByCardMapper{}
}

// ToTransactionMonthlyMethodByCardNumber maps a GetMonthlyPaymentMethodsByCardNumberRow database row to a TransactionMonthMethod
// domain model.
//
// Parameters:
//   - ss: A pointer to GetMonthlyPaymentMethodsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionMonthMethod containing the mapped data, including Month,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionStatisticMethodByCardMapper) ToTransactionMonthlyMethodByCardNumber(ss *db.GetMonthlyPaymentMethodsByCardNumberRow) *record.TransactionMonthMethod {
	return &record.TransactionMonthMethod{
		Month:             ss.Month,
		PaymentMethod:     ss.PaymentMethod,
		TotalTransactions: int(ss.TotalTransactions),
		TotalAmount:       int(ss.TotalAmount),
	}
}

// ToTransactionMonthlyMethodsByCardNumber maps a slice of GetMonthlyPaymentMethodsByCardNumberRow database rows
// to a slice of TransactionMonthMethod domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetMonthlyPaymentMethodsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionMonthMethod containing the mapped data, including Month,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionStatisticMethodByCardMapper) ToTransactionMonthlyMethodsByCardNumber(ss []*db.GetMonthlyPaymentMethodsByCardNumberRow) []*record.TransactionMonthMethod {
	var transactionRecords []*record.TransactionMonthMethod
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionMonthlyMethodByCardNumber(transaction))
	}
	return transactionRecords
}

// ToTransactionYearlyMethodByCardNumber maps a GetYearlyPaymentMethodsByCardNumberRow database row to a TransactionYearMethod
// domain model.
//
// Parameters:
//   - ss: A pointer to GetYearlyPaymentMethodsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionYearMethod containing the mapped data, including Year,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionStatisticMethodByCardMapper) ToTransactionYearlyMethodByCardNumber(ss *db.GetYearlyPaymentMethodsByCardNumberRow) *record.TransactionYearMethod {
	return &record.TransactionYearMethod{
		Year:              ss.Year,
		PaymentMethod:     ss.PaymentMethod,
		TotalTransactions: int(ss.TotalTransactions),
		TotalAmount:       int(ss.TotalAmount),
	}
}

// ToTransactionYearlyMethodsByCardNumber maps a slice of GetYearlyPaymentMethodsByCardNumberRow database rows
// to a slice of TransactionYearMethod domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetYearlyPaymentMethodsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionYearMethod containing the mapped data, including Year,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionStatisticMethodByCardMapper) ToTransactionYearlyMethodsByCardNumber(ss []*db.GetYearlyPaymentMethodsByCardNumberRow) []*record.TransactionYearMethod {
	var transactionRecords []*record.TransactionYearMethod
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionYearlyMethodByCardNumber(transaction))
	}
	return transactionRecords
}
