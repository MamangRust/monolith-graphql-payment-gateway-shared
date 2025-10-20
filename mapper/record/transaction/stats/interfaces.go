package transactionstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TransactionStatisticStatusRecordMapper maps transaction status rows to domain models.
type TransactionStatisticStatusRecordMapper interface {
	// ToTransactionRecordMonthStatusSuccess maps a single GetMonthTransactionStatusSuccessRow database row
	// to a TransactionRecordMonthStatusSuccess domain model.
	//
	// Parameters:
	//   - s: A pointer to GetMonthTransactionStatusSuccessRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
	//     Month, TotalSuccess, and TotalAmount.
	ToTransactionRecordMonthStatusSuccess(s *db.GetMonthTransactionStatusSuccessRow) *record.TransactionRecordMonthStatusSuccess
	// ToTransactionRecordsMonthStatusSuccess maps a slice of GetMonthTransactionStatusSuccessRow database rows to a slice of
	// TransactionRecordMonthStatusSuccess domain models.
	//
	// Parameters:
	//   - Transactions: A slice of pointers to GetMonthTransactionStatusSuccessRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
	//     Month, TotalSuccess, and TotalAmount.
	ToTransactionRecordsMonthStatusSuccess(transactions []*db.GetMonthTransactionStatusSuccessRow) []*record.TransactionRecordMonthStatusSuccess

	// ToTransactionRecordYearStatusSuccess maps a GetYearlyTransactionStatusSuccessRow database row
	// to a TransactionRecordYearStatusSuccess domain model.
	//
	// Parameters:
	//   - s: A pointer to GetYearlyTransactionStatusSuccessRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
	//     TotalSuccess, and TotalAmount.
	ToTransactionRecordYearStatusSuccess(s *db.GetYearlyTransactionStatusSuccessRow) *record.TransactionRecordYearStatusSuccess
	// ToTransactionRecordsYearStatusSuccess maps a slice of GetYearlyTransactionStatusSuccessRow database rows
	// to a slice of TransactionRecordYearStatusSuccess domain models.
	//
	// Parameters:
	//   - Transactions: A slice of pointers to GetYearlyTransactionStatusSuccessRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
	//     TotalSuccess, and TotalAmount.
	ToTransactionRecordsYearStatusSuccess(transactions []*db.GetYearlyTransactionStatusSuccessRow) []*record.TransactionRecordYearStatusSuccess

	// ToTransactionRecordMonthStatusFailed maps a GetMonthTransactionStatusFailedRow database row
	// to a TransactionRecordMonthStatusFailed domain model.
	//
	// Parameters:
	//   - s: A pointer to GetMonthTransactionStatusFailedRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
	//     Month, TotalFailed, and TotalAmount.
	ToTransactionRecordMonthStatusFailed(s *db.GetMonthTransactionStatusFailedRow) *record.TransactionRecordMonthStatusFailed
	// ToTransactionRecordsMonthStatusFailed maps a slice of GetMonthTransactionStatusFailedRow database rows
	// to a slice of TransactionRecordMonthStatusFailed domain models.
	//
	// Parameters:
	//   - Transactions: A slice of pointers to GetMonthTransactionStatusFailedRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
	//     Month, TotalFailed, and TotalAmount.
	ToTransactionRecordsMonthStatusFailed(transactions []*db.GetMonthTransactionStatusFailedRow) []*record.TransactionRecordMonthStatusFailed

	// ToTransactionRecordYearStatusFailed maps a GetYearlyTransactionStatusFailedRow database row
	// to a TransactionRecordYearStatusFailed domain model.
	//
	// Parameters:
	//   - s: A pointer to GetYearlyTransactionStatusFailedRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionRecordYearStatusFailed containing the mapped data, including Year,
	//     TotalFailed, and TotalAmount.
	ToTransactionRecordYearStatusFailed(s *db.GetYearlyTransactionStatusFailedRow) *record.TransactionRecordYearStatusFailed
	// ToTransactionRecordsYearStatusFailed maps a slice of GetYearlyTransactionStatusFailedRow database rows
	// to a slice of TransactionRecordYearStatusFailed domain models.
	//
	// Parameters:
	//   - Transactions: A slice of pointers to GetYearlyTransactionStatusFailedRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionRecordYearStatusFailed containing the mapped data, including Year,
	//     TotalFailed, and TotalAmount.
	ToTransactionRecordsYearStatusFailed(transactions []*db.GetYearlyTransactionStatusFailedRow) []*record.TransactionRecordYearStatusFailed
}

// TransactionStatisticMethodRecordMapper maps payment method statistics rows to domain models.
type TransactionStatisticMethodRecordMapper interface {
	// ToTransactionMonthlyMethod maps a GetMonthlyPaymentMethodsRow database row to a TransactionMonthMethod
	// domain model.
	//
	// Parameters:
	//   - s: A pointer to GetMonthlyPaymentMethodsRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionMonthMethod containing the mapped data, including Month,
	//     PaymentMethod, TotalTransactions, and TotalAmount.
	ToTransactionMonthlyMethod(ss *db.GetMonthlyPaymentMethodsRow) *record.TransactionMonthMethod
	// ToTransactionMonthlyMethods maps a slice of GetMonthlyPaymentMethodsRow database rows
	// to a slice of TransactionMonthMethod domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetMonthlyPaymentMethodsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionMonthMethod containing the mapped data, including Month,
	//     PaymentMethod, TotalTransactions, and TotalAmount.
	ToTransactionMonthlyMethods(ss []*db.GetMonthlyPaymentMethodsRow) []*record.TransactionMonthMethod

	// ToTransactionYearlyMethod maps a GetYearlyPaymentMethodsRow database row to a TransactionYearMethod
	// domain model.
	//
	// Parameters:
	//   - s: A pointer to GetYearlyPaymentMethodsRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionYearMethod containing the mapped data, including Year,
	//     PaymentMethod, TotalTransactions, and TotalAmount.
	ToTransactionYearlyMethod(ss *db.GetYearlyPaymentMethodsRow) *record.TransactionYearMethod
	// ToTransactionYearlyMethods maps a slice of GetYearlyPaymentMethodsRow database rows
	// to a slice of TransactionYearMethod domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetYearlyPaymentMethodsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionYearMethod containing the mapped data, including Year,
	//     PaymentMethod, TotalTransactions, and TotalAmount.
	ToTransactionYearlyMethods(ss []*db.GetYearlyPaymentMethodsRow) []*record.TransactionYearMethod
}

// TransactionStatisticAmountRecordMapper maps amount statistics rows to domain models.
type TransactionStatisticAmountRecordMapper interface {
	// ToTransactionMonthlyAmount maps a GetMonthlyAmountsRow database row to a TransactionMonthAmount
	// domain model.
	//
	// Parameters:
	//   - ss: A pointer to GetMonthlyAmountsRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
	ToTransactionMonthlyAmount(ss *db.GetMonthlyAmountsRow) *record.TransactionMonthAmount
	// ToTransactionMonthlyAmounts maps a slice of GetMonthlyAmountsRow database rows
	// to a slice of TransactionMonthAmount domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetMonthlyAmountsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
	ToTransactionMonthlyAmounts(ss []*db.GetMonthlyAmountsRow) []*record.TransactionMonthAmount

	// ToTransactionYearlyAmount maps a GetYearlyAmountsRow database row to a TransactionYearlyAmount
	// domain model.
	//
	// Parameters:
	//   - ss: A pointer to GetYearlyAmountsRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
	ToTransactionYearlyAmount(ss *db.GetYearlyAmountsRow) *record.TransactionYearlyAmount
	// ToTransactionYearlyAmounts maps a slice of GetYearlyAmountsRow database rows
	// to a slice of TransactionYearlyAmount domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetYearlyAmountsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
	ToTransactionYearlyAmounts(ss []*db.GetYearlyAmountsRow) []*record.TransactionYearlyAmount
}
