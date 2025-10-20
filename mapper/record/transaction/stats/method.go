package transactionstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

type transactionStatisticMethodRecordMapper struct{}

func NewTransactionStatisticMethodRecordMapper() TransactionStatisticMethodRecordMapper {
	return &transactionStatisticMethodRecordMapper{}
}

// ToTransactionMonthlyMethod maps a GetMonthlyPaymentMethodsRow database row to a TransactionMonthMethod
// domain model.
//
// Parameters:
//   - s: A pointer to GetMonthlyPaymentMethodsRow representing the database row.
//
// Returns:
//   - A pointer to TransactionMonthMethod containing the mapped data, including Month,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionStatisticMethodRecordMapper) ToTransactionMonthlyMethod(ss *db.GetMonthlyPaymentMethodsRow) *record.TransactionMonthMethod {
	return &record.TransactionMonthMethod{
		Month:             ss.Month,
		PaymentMethod:     ss.PaymentMethod,
		TotalTransactions: int(ss.TotalTransactions),
		TotalAmount:       int(ss.TotalAmount),
	}
}

// ToTransactionMonthlyMethods maps a slice of GetMonthlyPaymentMethodsRow database rows
// to a slice of TransactionMonthMethod domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetMonthlyPaymentMethodsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionMonthMethod containing the mapped data, including Month,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionStatisticMethodRecordMapper) ToTransactionMonthlyMethods(ss []*db.GetMonthlyPaymentMethodsRow) []*record.TransactionMonthMethod {
	var transactionRecords []*record.TransactionMonthMethod
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionMonthlyMethod(transaction))
	}
	return transactionRecords
}

// ToTransactionYearlyMethod maps a GetYearlyPaymentMethodsRow database row to a TransactionYearMethod
// domain model.
//
// Parameters:
//   - s: A pointer to GetYearlyPaymentMethodsRow representing the database row.
//
// Returns:
//   - A pointer to TransactionYearMethod containing the mapped data, including Year,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionStatisticMethodRecordMapper) ToTransactionYearlyMethod(ss *db.GetYearlyPaymentMethodsRow) *record.TransactionYearMethod {
	return &record.TransactionYearMethod{
		Year:              ss.Year,
		PaymentMethod:     ss.PaymentMethod,
		TotalTransactions: int(ss.TotalTransactions),
		TotalAmount:       int(ss.TotalAmount),
	}
}

// ToTransactionYearlyMethods maps a slice of GetYearlyPaymentMethodsRow database rows
// to a slice of TransactionYearMethod domain models.
//
// Parameters:
//   - ss: A slice of pointers to GetYearlyPaymentMethodsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionYearMethod containing the mapped data, including Year,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionStatisticMethodRecordMapper) ToTransactionYearlyMethods(ss []*db.GetYearlyPaymentMethodsRow) []*record.TransactionYearMethod {
	var transactionRecords []*record.TransactionYearMethod
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionYearlyMethod(transaction))
	}
	return transactionRecords
}
