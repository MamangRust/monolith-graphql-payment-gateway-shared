package cardstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticByCardRecordMapper provides methods to map database rows to CardStatisticByCardRecord domain models.
type CardStatisticBalanceByCardRecordMapper interface {
	// ToMonthlyBalanceCardNumber maps a GetMonthlyBalancesByCardNumberRow database row to a CardMonthBalance domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetMonthlyBalancesByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardMonthBalance containing the mapped data, including Month and TotalBalance.
	ToMonthlyBalanceCardNumber(card *db.GetMonthlyBalancesByCardNumberRow) *record.CardMonthBalance
	// ToMonthlyBalancesCardNumber maps a slice of GetMonthlyBalancesByCardNumberRow database rows to a slice of CardMonthBalance domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetMonthlyBalancesByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardMonthBalance containing the mapped data, including Month and TotalBalance.
	ToMonthlyBalancesCardNumber(cards []*db.GetMonthlyBalancesByCardNumberRow) []*record.CardMonthBalance

	// ToYearlyBalanceCardNumber maps a GetYearlyBalancesByCardNumberRow database row to a CardYearlyBalance domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetYearlyBalancesByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardYearlyBalance containing the mapped data, including Year and TotalBalance.
	ToYearlyBalanceCardNumber(card *db.GetYearlyBalancesByCardNumberRow) *record.CardYearlyBalance
	// ToYearlyBalancesCardNumber maps a slice of GetYearlyBalancesByCardNumberRow database rows to a slice of CardYearlyBalance domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetYearlyBalancesByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardYearlyBalance containing the mapped data, including Year and TotalBalance.
	ToYearlyBalancesCardNumber(cards []*db.GetYearlyBalancesByCardNumberRow) []*record.CardYearlyBalance
}

// CardStatisticTopupByCardRecordMapper provides methods to map database rows to CardStatisticTopupByCardRecord domain models.
type CardStatisticTopupByCardRecordMapper interface {
	// ToMonthlyTopupAmountByCardNumber maps a GetMonthlyTopupAmountByCardNumberRow database row to a CardMonthAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetMonthlyTopupAmountByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTopupAmountByCardNumber(card *db.GetMonthlyTopupAmountByCardNumberRow) *record.CardMonthAmount
	// ToMonthlyTopupAmountsByCardNumber maps a slice of GetMonthlyTopupAmountByCardNumberRow database rows to a slice of CardMonthAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetMonthlyTopupAmountByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTopupAmountsByCardNumber(cards []*db.GetMonthlyTopupAmountByCardNumberRow) []*record.CardMonthAmount

	// ToYearlyTopupAmountByCardNumber maps a GetYearlyTopupAmountByCardNumberRow database row to a CardYearAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetYearlyTopupAmountByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTopupAmountByCardNumber(card *db.GetYearlyTopupAmountByCardNumberRow) *record.CardYearAmount
	// ToYearlyTopupAmountsByCardNumber maps a slice of GetYearlyTopupAmountByCardNumberRow database rows
	// to a slice of CardYearAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetYearlyTopupAmountByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTopupAmountsByCardNumber(cards []*db.GetYearlyTopupAmountByCardNumberRow) []*record.CardYearAmount
}

// CardStatisticWithdrawByCardRecordMapper provides methods to map database rows to CardStatisticWithdrawByCardRecord domain models.
type CardStatisticWithdrawByCardRecordMapper interface {
	// ToMonthlyWithdrawAmountByCardNumber maps a GetMonthlyWithdrawAmountByCardNumberRow database row to a CardMonthAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetMonthlyWithdrawAmountByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyWithdrawAmountByCardNumber(card *db.GetMonthlyWithdrawAmountByCardNumberRow) *record.CardMonthAmount
	// ToMonthlyWithdrawAmountsByCardNumber maps a slice of GetMonthlyWithdrawAmountByCardNumberRow database rows
	// to a slice of CardMonthAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetMonthlyWithdrawAmountByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyWithdrawAmountsByCardNumber(cards []*db.GetMonthlyWithdrawAmountByCardNumberRow) []*record.CardMonthAmount

	// ToYearlyWithdrawAmountByCardNumber maps a GetYearlyWithdrawAmountByCardNumberRow database row to a CardYearAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetYearlyWithdrawAmountByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyWithdrawAmountByCardNumber(card *db.GetYearlyWithdrawAmountByCardNumberRow) *record.CardYearAmount
	// ToYearlyWithdrawAmountsByCardNumber maps a slice of GetYearlyWithdrawAmountByCardNumberRow database rows
	// to a slice of CardYearAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetYearlyWithdrawAmountByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyWithdrawAmountsByCardNumber(cards []*db.GetYearlyWithdrawAmountByCardNumberRow) []*record.CardYearAmount
}

// CardStatisticTransactionByCardRecordMapper provides methods to map database rows to CardStatisticTransactionByCardRecord domain models.
type CardStatisticTransactionByCardRecordMapper interface {
	// ToMonthlyTransactionAmountByCardNumber maps a GetMonthlyTransactionAmountByCardNumberRow database row
	// to a CardMonthAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetMonthlyTransactionAmountByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransactionAmountByCardNumber(card *db.GetMonthlyTransactionAmountByCardNumberRow) *record.CardMonthAmount
	// ToMonthlyTransactionAmountsByCardNumber maps a slice of GetMonthlyTransactionAmountByCardNumberRow database rows
	// to a slice of CardMonthAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetMonthlyTransactionAmountByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransactionAmountsByCardNumber(cards []*db.GetMonthlyTransactionAmountByCardNumberRow) []*record.CardMonthAmount

	// ToYearlyTransactionAmountByCardNumber maps a GetYearlyTransactionAmountByCardNumberRow database row
	// to a CardYearAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetYearlyTransactionAmountByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTransactionAmountByCardNumber(card *db.GetYearlyTransactionAmountByCardNumberRow) *record.CardYearAmount
	// ToYearlyTransactionAmountsByCardNumber maps a slice of GetYearlyTransactionAmountByCardNumberRow database rows
	// to a slice of CardYearAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetYearlyTransactionAmountByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.

	ToYearlyTransactionAmountsByCardNumber(cards []*db.GetYearlyTransactionAmountByCardNumberRow) []*record.CardYearAmount
}

// CardStatisticTransferByCardRecordMapper provides methods to map database rows to CardStatisticTransferByCardRecord domain models.
type CardStatisticTransferByCardRecordMapper interface {
	// ToMonthlyTransferSenderAmountByCardNumber maps a GetMonthlyTransferAmountBySenderRow database row to a CardMonthAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetMonthlyTransferAmountBySenderRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransferSenderAmountByCardNumber(card *db.GetMonthlyTransferAmountBySenderRow) *record.CardMonthAmount
	// ToMonthlyTransferSenderAmountsByCardNumber maps a slice of GetMonthlyTransferAmountBySenderRow database rows
	// to a slice of CardMonthAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetMonthlyTransferAmountBySenderRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransferSenderAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountBySenderRow) []*record.CardMonthAmount

	// ToYearlyTransferSenderAmountByCardNumber maps a GetYearlyTransferAmountBySenderRow database row
	// to a CardYearAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetYearlyTransferAmountBySenderRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTransferSenderAmountByCardNumber(card *db.GetYearlyTransferAmountBySenderRow) *record.CardYearAmount
	// ToYearlyTransferSenderAmountsByCardNumber maps a slice of GetYearlyTransferAmountBySenderRow database rows
	// to a slice of CardYearAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetYearlyTransferAmountBySenderRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTransferSenderAmountsByCardNumber(cards []*db.GetYearlyTransferAmountBySenderRow) []*record.CardYearAmount

	// ToMonthlyTransferReceiverAmountByCardNumber maps a GetMonthlyTransferAmountByReceiverRow database row
	// to a CardMonthAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetMonthlyTransferAmountByReceiverRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransferReceiverAmountByCardNumber(card *db.GetMonthlyTransferAmountByReceiverRow) *record.CardMonthAmount
	// ToMonthlyTransferReceiverAmountsByCardNumber maps a slice of GetMonthlyTransferAmountByReceiverRow database rows
	// to a slice of CardMonthAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetMonthlyTransferAmountByReceiverRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransferReceiverAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountByReceiverRow) []*record.CardMonthAmount

	// ToYearlyTransferReceiverAmountByCardNumber maps a GetYearlyTransferAmountByReceiverRow database row
	// to a CardYearAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetYearlyTransferAmountByReceiverRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTransferReceiverAmountByCardNumber(card *db.GetYearlyTransferAmountByReceiverRow) *record.CardYearAmount
	// ToYearlyTransferReceiverAmountsByCardNumber maps a slice of GetYearlyTransferAmountByReceiverRow database rows
	// to a slice of CardYearAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetYearlyTransferAmountByReceiverRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTransferReceiverAmountsByCardNumber(cards []*db.GetYearlyTransferAmountByReceiverRow) []*record.CardYearAmount
}
