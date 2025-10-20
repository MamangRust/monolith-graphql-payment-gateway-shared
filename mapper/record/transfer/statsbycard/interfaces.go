package transferstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TransferStatisticStatusByCardRecordMapper maps transfer status (success/failed) filtered by card_number.
type TransferStatisticStatusByCardRecordMapper interface {
	// ToTransferRecordMonthStatusSuccessCardNumber maps a GetMonthTransferStatusSuccessCardNumberRow database row
	// to a TransferRecordMonthStatusSuccess domain model.
	//
	// Parameters:
	//   - s: A pointer to GetMonthTransferStatusSuccessCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
	//     Month, TotalSuccess, and TotalAmount.
	ToTransferRecordMonthStatusSuccessCardNumber(s *db.GetMonthTransferStatusSuccessCardNumberRow) *record.TransferRecordMonthStatusSuccess
	// ToTransferRecordsMonthStatusSuccessCardNumber maps a slice of GetMonthTransferStatusSuccessCardNumberRow database rows
	// to a slice of TransferRecordMonthStatusSuccess domain models.
	//
	// Parameters:
	//   - Transfers: A slice of pointers to GetMonthTransferStatusSuccessCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
	//     Month, TotalSuccess, and TotalAmount.
	ToTransferRecordsMonthStatusSuccessCardNumber(s []*db.GetMonthTransferStatusSuccessCardNumberRow) []*record.TransferRecordMonthStatusSuccess
	// ToTransferRecordYearStatusSuccessCardNumber maps a GetYearlyTransferStatusSuccessCardNumberRow database row
	// to a TransferRecordYearStatusSuccess domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetYearlyTransferStatusSuccessCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferRecordYearStatusSuccess containing the mapped data, including Year,
	//     TotalSuccess, and TotalAmount.
	ToTransferRecordYearStatusSuccessCardNumber(s *db.GetYearlyTransferStatusSuccessCardNumberRow) *record.TransferRecordYearStatusSuccess
	// ToTransferRecordsYearStatusSuccessCardNumber maps a slice of GetYearlyTransferStatusSuccessCardNumberRow database rows
	// to a slice of TransferRecordYearStatusSuccess domain models.
	//
	// Parameters:
	//   - Transfers: A slice of pointers to GetYearlyTransferStatusSuccessCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferRecordYearStatusSuccess containing the mapped data, including Year,
	//     TotalSuccess, and TotalAmount.
	ToTransferRecordsYearStatusSuccessCardNumber(s []*db.GetYearlyTransferStatusSuccessCardNumberRow) []*record.TransferRecordYearStatusSuccess

	// ToTransferRecordMonthStatusFailedCardNumber maps a GetMonthTransferStatusFailedCardNumberRow database row
	// to a TransferRecordMonthStatusFailed domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetMonthTransferStatusFailedCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferRecordMonthStatusFailed containing the mapped data, including Year,
	//     Month, TotalFailed, and TotalAmount.
	ToTransferRecordMonthStatusFailedCardNumber(s *db.GetMonthTransferStatusFailedCardNumberRow) *record.TransferRecordMonthStatusFailed
	// ToTransferRecordsMonthStatusFailedCardNumber maps a slice of GetMonthTransferStatusFailedCardNumberRow database rows
	// to a slice of TransferRecordMonthStatusFailed domain models.
	//
	// Parameters:
	//   - Transfers: A slice of pointers to GetMonthTransferStatusFailedCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferRecordMonthStatusFailed containing the mapped data, including Year,
	//     Month, TotalFailed, and TotalAmount.
	ToTransferRecordsMonthStatusFailedCardNumber(s []*db.GetMonthTransferStatusFailedCardNumberRow) []*record.TransferRecordMonthStatusFailed
	// ToTransferRecordYearStatusFailedCardNumber maps a GetYearlyTransferStatusFailedCardNumberRow database row
	// to a TransferRecordYearStatusFailed domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetYearlyTransferStatusFailedCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferRecordYearStatusFailed containing the mapped data, including Year,
	//     TotalFailed, and TotalAmount.
	ToTransferRecordYearStatusFailedCardNumber(s *db.GetYearlyTransferStatusFailedCardNumberRow) *record.TransferRecordYearStatusFailed
	// ToTransferRecordsYearStatusFailedCardNumber maps a slice of GetYearlyTransferStatusFailedCardNumberRow database rows
	// to a slice of TransferRecordYearStatusFailed domain models.
	//
	// Parameters:
	//   - Transfers: A slice of pointers to GetYearlyTransferStatusFailedCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferRecordYearStatusFailed containing the mapped data, including Year,
	//     TotalFailed, and TotalAmount.
	ToTransferRecordsYearStatusFailedCardNumber(s []*db.GetYearlyTransferStatusFailedCardNumberRow) []*record.TransferRecordYearStatusFailed
}

// TransferStatisticSenderAmountByCardRecordMapper maps transfer amounts filtered by sender card number.
type TransferStatisticSenderAmountByCardRecordMapper interface {
	// ToTransferMonthAmountSender maps a GetMonthlyTransferAmountsBySenderCardNumberRow database row
	// to a TransferMonthAmount domain model.
	//
	// Parameters:
	//   - ss: A pointer to a GetMonthlyTransferAmountsBySenderCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferMonthAmount containing the mapped data, including Month
	//     and TotalAmount.
	ToTransferMonthAmountSender(s *db.GetMonthlyTransferAmountsBySenderCardNumberRow) *record.TransferMonthAmount
	// ToTransferMonthAmountsSender maps a slice of GetMonthlyTransferAmountsBySenderCardNumberRow
	// database rows to a slice of TransferMonthAmount domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetMonthlyTransferAmountsBySenderCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferMonthAmount containing the mapped data, including Month and TotalAmount.
	ToTransferMonthAmountsSender(s []*db.GetMonthlyTransferAmountsBySenderCardNumberRow) []*record.TransferMonthAmount
	// ToTransferYearAmountSender maps a GetYearlyTransferAmountsBySenderCardNumberRow database row
	// to a TransferYearAmount domain model.
	//
	// Parameters:
	//   - ss: A pointer to a GetYearlyTransferAmountsBySenderCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.
	ToTransferYearAmountSender(s *db.GetYearlyTransferAmountsBySenderCardNumberRow) *record.TransferYearAmount
	// ToTransferYearAmountsSender maps a slice of GetYearlyTransferAmountsBySenderCardNumberRow database rows
	// to a slice of TransferYearAmount domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetYearlyTransferAmountsBySenderCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.
	ToTransferYearAmountsSender(s []*db.GetYearlyTransferAmountsBySenderCardNumberRow) []*record.TransferYearAmount
}

// TransferStatisticReceiverAmountByCardRecordMapper maps transfer amounts filtered by receiver card number.
type TransferStatisticReceiverAmountByCardRecordMapper interface {
	// ToTransferMonthAmountReceiver maps a GetMonthlyTransferAmountsByReceiverCardNumberRow database row
	// to a TransferMonthAmount domain model.
	//
	// Parameters:
	//   - ss: A pointer to a GetMonthlyTransferAmountsByReceiverCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferMonthAmount containing the mapped data, including Month and TotalAmount.
	ToTransferMonthAmountReceiver(s *db.GetMonthlyTransferAmountsByReceiverCardNumberRow) *record.TransferMonthAmount
	// ToTransferMonthAmountsReceiver maps a slice of GetMonthlyTransferAmountsByReceiverCardNumberRow database rows
	// to a slice of TransferMonthAmount domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetMonthlyTransferAmountsByReceiverCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferMonthAmount containing the mapped data, including Month and TotalAmount.
	ToTransferMonthAmountsReceiver(s []*db.GetMonthlyTransferAmountsByReceiverCardNumberRow) []*record.TransferMonthAmount
	// ToTransferYearAmountReceiver maps a GetYearlyTransferAmountsByReceiverCardNumberRow database row
	// to a TransferYearAmount domain model.
	//
	// Parameters:
	//   - ss: A pointer to a GetYearlyTransferAmountsByReceiverCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.
	ToTransferYearAmountReceiver(s *db.GetYearlyTransferAmountsByReceiverCardNumberRow) *record.TransferYearAmount
	// ToTransferYearAmountsReceiver maps a slice of GetYearlyTransferAmountsByReceiverCardNumberRow database rows
	// to a slice of TransferYearAmount domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetYearlyTransferAmountsByReceiverCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.
	ToTransferYearAmountsReceiver(s []*db.GetYearlyTransferAmountsByReceiverCardNumberRow) []*record.TransferYearAmount
}
