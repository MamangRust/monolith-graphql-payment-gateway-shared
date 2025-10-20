package transactionstatbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TransactionStatisticByCardStatusMapper maps transaction status by card number to domain models.
type TransactionStatisticByCardStatusMapper interface {
	// ToTransactionRecordMonthStatusSuccessCardNumber maps a GetMonthTransactionStatusSuccessCardNumberRow database row
	// to a TransactionRecordMonthStatusSuccess domain model.
	//
	// Parameters:
	//   - s: A pointer to GetMonthTransactionStatusSuccessCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
	//     Month, TotalSuccess, and TotalAmount.
	ToTransactionRecordMonthStatusSuccessCardNumber(s *db.GetMonthTransactionStatusSuccessCardNumberRow) *record.TransactionRecordMonthStatusSuccess
	// ToTransactionRecordsMonthStatusSuccessCardNumber maps a slice of GetMonthTransactionStatusSuccessCardNumberRow database rows
	// to a slice of TransactionRecordMonthStatusSuccess domain models.
	//
	// Parameters:
	//   - Transactions: A slice of pointers to GetMonthTransactionStatusSuccessCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
	//     Month, TotalSuccess, and TotalAmount.
	ToTransactionRecordsMonthStatusSuccessCardNumber(s []*db.GetMonthTransactionStatusSuccessCardNumberRow) []*record.TransactionRecordMonthStatusSuccess
	// ToTransactionRecordYearStatusSuccessCardNumber maps a GetYearlyTransactionStatusSuccessCardNumberRow database row
	// to a TransactionRecordYearStatusSuccess domain model.
	//
	// Parameters:
	//   - s: A pointer to GetYearlyTransactionStatusSuccessCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
	//     TotalSuccess, and TotalAmount.
	ToTransactionRecordYearStatusSuccessCardNumber(s *db.GetYearlyTransactionStatusSuccessCardNumberRow) *record.TransactionRecordYearStatusSuccess
	// ToTransactionRecordsYearStatusSuccessCardNumber maps a slice of GetYearlyTransactionStatusSuccessCardNumberRow database rows
	// to a slice of TransactionRecordYearStatusSuccess domain models.
	//
	// Parameters:
	//   - Transactions: A slice of pointers to GetYearlyTransactionStatusSuccessCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
	//     TotalSuccess, and TotalAmount.
	ToTransactionRecordsYearStatusSuccessCardNumber(s []*db.GetYearlyTransactionStatusSuccessCardNumberRow) []*record.TransactionRecordYearStatusSuccess

	// ToTransactionRecordMonthStatusFailedCardNumber maps a GetMonthTransactionStatusFailedCardNumberRow
	// database row to a TransactionRecordMonthStatusFailed domain model.
	//
	// Parameters:
	//   - s: A pointer to GetMonthTransactionStatusFailedCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
	//     Month, TotalFailed, and TotalAmount.
	ToTransactionRecordMonthStatusFailedCardNumber(s *db.GetMonthTransactionStatusFailedCardNumberRow) *record.TransactionRecordMonthStatusFailed
	// ToTransactionRecordsMonthStatusFailedCardNumber maps a slice of GetMonthTransactionStatusFailedCardNumberRow
	// database rows to a slice of TransactionRecordMonthStatusFailed domain models.
	//
	// Parameters:
	//   - Transactions: A slice of pointers to GetMonthTransactionStatusFailedCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
	//     Month, TotalFailed, and TotalAmount.
	ToTransactionRecordsMonthStatusFailedCardNumber(s []*db.GetMonthTransactionStatusFailedCardNumberRow) []*record.TransactionRecordMonthStatusFailed

	// ToTransactionRecordYearStatusFailedCardNumber maps a GetYearlyTransactionStatusFailedCardNumberRow
	// database row to a TransactionRecordYearStatusFailed domain model.
	//
	// Parameters:
	//   - s: A pointer to GetYearlyTransactionStatusFailedCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionRecordYearStatusFailed containing the mapped data, including Year,
	//     TotalFailed, and TotalAmount.
	ToTransactionRecordYearStatusFailedCardNumber(s *db.GetYearlyTransactionStatusFailedCardNumberRow) *record.TransactionRecordYearStatusFailed
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
	ToTransactionRecordsYearStatusFailedCardNumber(s []*db.GetYearlyTransactionStatusFailedCardNumberRow) []*record.TransactionRecordYearStatusFailed
}

// TransactionStatisticByCardMethodMapper maps transaction method by card number to domain models.
type TransactionStatisticByCardMethodMapper interface {
	// ToTransactionMonthlyMethodByCardNumber maps a GetMonthlyPaymentMethodsByCardNumberRow database row to a TransactionMonthMethod
	// domain model.
	//
	// Parameters:
	//   - ss: A pointer to GetMonthlyPaymentMethodsByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionMonthMethod containing the mapped data, including Month,
	//     PaymentMethod, TotalTransactions, and TotalAmount.
	ToTransactionMonthlyMethodByCardNumber(s *db.GetMonthlyPaymentMethodsByCardNumberRow) *record.TransactionMonthMethod
	// ToTransactionMonthlyMethodsByCardNumber maps a slice of GetMonthlyPaymentMethodsByCardNumberRow database rows
	// to a slice of TransactionMonthMethod domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetMonthlyPaymentMethodsByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionMonthMethod containing the mapped data, including Month,
	//     PaymentMethod, TotalTransactions, and TotalAmount.
	ToTransactionMonthlyMethodsByCardNumber(s []*db.GetMonthlyPaymentMethodsByCardNumberRow) []*record.TransactionMonthMethod

	// ToTransactionYearlyMethodByCardNumber maps a GetYearlyPaymentMethodsByCardNumberRow database row to a TransactionYearMethod
	// domain model.
	//
	// Parameters:
	//   - ss: A pointer to GetYearlyPaymentMethodsByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionYearMethod containing the mapped data, including Year,
	//     PaymentMethod, TotalTransactions, and TotalAmount.
	ToTransactionYearlyMethodByCardNumber(s *db.GetYearlyPaymentMethodsByCardNumberRow) *record.TransactionYearMethod
	// ToTransactionYearlyMethodsByCardNumber maps a slice of GetYearlyPaymentMethodsByCardNumberRow database rows
	// to a slice of TransactionYearMethod domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetYearlyPaymentMethodsByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionYearMethod containing the mapped data, including Year,
	//     PaymentMethod, TotalTransactions, and TotalAmount.
	ToTransactionYearlyMethodsByCardNumber(s []*db.GetYearlyPaymentMethodsByCardNumberRow) []*record.TransactionYearMethod
}

// TransactionStatisticByCardAmountMapper maps transaction amount by card number to domain models.
type TransactionStatisticByCardAmountMapper interface {
	// ToTransactionMonthlyAmountByCardNumber maps a GetMonthlyAmountsByCardNumberRow database row
	// to a TransactionMonthAmount domain model.
	//
	// Parameters:
	//   - ss: A pointer to GetMonthlyAmountsByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
	ToTransactionMonthlyAmountByCardNumber(s *db.GetMonthlyAmountsByCardNumberRow) *record.TransactionMonthAmount
	// ToTransactionMonthlyAmountsByCardNumber maps a slice of GetMonthlyAmountsByCardNumberRow database rows
	// to a slice of TransactionMonthAmount domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetMonthlyAmountsByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
	ToTransactionMonthlyAmountsByCardNumber(s []*db.GetMonthlyAmountsByCardNumberRow) []*record.TransactionMonthAmount

	// ToTransactionYearlyAmountByCardNumber maps a GetYearlyAmountsByCardNumberRow database row
	// to a TransactionYearlyAmount domain model.
	//
	// Parameters:
	//   - ss: A pointer to GetYearlyAmountsByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
	ToTransactionYearlyAmountByCardNumber(s *db.GetYearlyAmountsByCardNumberRow) *record.TransactionYearlyAmount
	// ToTransactionYearlyAmountsByCardNumber maps a slice of GetYearlyAmountsByCardNumberRow database rows
	// to a slice of TransactionYearlyAmount domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetYearlyAmountsByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
	ToTransactionYearlyAmountsByCardNumber(s []*db.GetYearlyAmountsByCardNumberRow) []*record.TransactionYearlyAmount
}
