package cardstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// CardStatisticBalanceRecordMapper provides methods to map database rows to CardStatisticBalanceRecord domain models.
type CardStatisticBalanceRecordMapper interface {
	// ToMonthlyBalance maps a GetMonthlyBalancesRow database row to a CardMonthBalance domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetMonthlyBalancesRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardMonthBalance containing the mapped data, including Month and TotalBalance.
	ToMonthlyBalance(card *db.GetMonthlyBalancesRow) *record.CardMonthBalance
	// ToMonthlyBalances maps a slice of GetMonthlyBalancesRow database rows to a slice of CardMonthBalance domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetMonthlyBalancesRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardMonthBalance containing the mapped data, including Month and TotalBalance.
	ToMonthlyBalances(cards []*db.GetMonthlyBalancesRow) []*record.CardMonthBalance

	// ToYearlyBalance maps a GetYearlyBalancesRow database row to a CardYearlyBalance domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetYearlyBalancesRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardYearlyBalance containing the mapped data, including Year and TotalBalance.
	ToYearlyBalance(card *db.GetYearlyBalancesRow) *record.CardYearlyBalance
	// ToYearlyBalances maps a slice of GetYearlyBalancesRow database rows to a slice of CardYearlyBalance domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetYearlyBalancesRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardYearlyBalance containing the mapped data, including Year and TotalBalance.
	ToYearlyBalances(cards []*db.GetYearlyBalancesRow) []*record.CardYearlyBalance
}

// CardStatisticTopupRecordMapper provides methods to map database rows to CardStatisticTopupRecord domain models.
type CardStatisticTopupRecordMapper interface {
	// ToMonthlyTopupAmount maps a GetMonthlyTopupAmountRow database row to a CardMonthAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetMonthlyTopupAmountRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTopupAmount(card *db.GetMonthlyTopupAmountRow) *record.CardMonthAmount
	// ToMonthlyTopupAmounts maps a slice of GetMonthlyTopupAmountRow database rows to a slice of CardMonthAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetMonthlyTopupAmountRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTopupAmounts(cards []*db.GetMonthlyTopupAmountRow) []*record.CardMonthAmount
	// ToYearlyTopupAmount maps a GetYearlyTopupAmountRow database row to a CardYearAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetYearlyTopupAmountRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTopupAmount(card *db.GetYearlyTopupAmountRow) *record.CardYearAmount
	// ToYearlyTopupAmounts maps a slice of GetYearlyTopupAmountRow database rows to a slice of CardYearAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetYearlyTopupAmountRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTopupAmounts(cards []*db.GetYearlyTopupAmountRow) []*record.CardYearAmount
}

// CardStatisticWithdrawRecordMapper provides methods to map database rows to CardStatisticWithdrawRecord domain models.
type CardStatisticWithdrawRecordMapper interface {
	ToMonthlyWithdrawAmount(card *db.GetMonthlyWithdrawAmountRow) *record.CardMonthAmount
	ToMonthlyWithdrawAmounts(cards []*db.GetMonthlyWithdrawAmountRow) []*record.CardMonthAmount

	ToYearlyWithdrawAmount(card *db.GetYearlyWithdrawAmountRow) *record.CardYearAmount
	ToYearlyWithdrawAmounts(cards []*db.GetYearlyWithdrawAmountRow) []*record.CardYearAmount
}

// CardStatisticTransactionRecordMapper provides methods to map database rows to CardStatisticTransactionRecord domain models.
type CardStatisticTransactionRecordMapper interface {
	// ToMonthlyTransactionAmount maps a GetMonthlyTransactionAmountRow database row to a CardMonthAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetMonthlyTransactionAmountRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransactionAmount(card *db.GetMonthlyTransactionAmountRow) *record.CardMonthAmount
	// ToMonthlyTransactionAmounts maps a slice of GetMonthlyTransactionAmountRow database rows to a slice of CardMonthAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetMonthlyTransactionAmountRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransactionAmounts(cards []*db.GetMonthlyTransactionAmountRow) []*record.CardMonthAmount

	// ToYearlyTransactionAmount maps a GetYearlyTransactionAmountRow database row to a CardYearAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetYearlyTransactionAmountRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTransactionAmount(card *db.GetYearlyTransactionAmountRow) *record.CardYearAmount
	// ToYearlyTransactionAmounts maps a slice of GetYearlyTransactionAmountRow database rows to a slice of CardYearAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetYearlyTransactionAmountRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTransactionAmounts(cards []*db.GetYearlyTransactionAmountRow) []*record.CardYearAmount
}

// CardStatisticTransferRecordMapper provides methods to map database rows to CardStatisticTransferRecord domain models.
type CardStatisticTransferRecordMapper interface {
	// ToMonthlyTransferSenderAmount maps a GetMonthlyTransferAmountSenderRow database row to a CardMonthAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetMonthlyTransferAmountSenderRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransferSenderAmount(card *db.GetMonthlyTransferAmountSenderRow) *record.CardMonthAmount
	// ToMonthlyTransferSenderAmounts maps a slice of GetMonthlyTransferAmountSenderRow database rows to a slice of CardMonthAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetMonthlyTransferAmountSenderRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransferSenderAmounts(cards []*db.GetMonthlyTransferAmountSenderRow) []*record.CardMonthAmount

	// ToYearlyTransferSenderAmount maps a GetYearlyTransferAmountSenderRow database row to a CardYearAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetYearlyTransferAmountSenderRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTransferSenderAmount(card *db.GetYearlyTransferAmountSenderRow) *record.CardYearAmount
	// ToYearlyTransferSenderAmounts maps a slice of GetYearlyTransferAmountSenderRow database rows to a slice of CardYearAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetYearlyTransferAmountSenderRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTransferSenderAmounts(cards []*db.GetYearlyTransferAmountSenderRow) []*record.CardYearAmount

	// ToMonthlyTransferReceiverAmount maps a GetMonthlyTransferAmountReceiverRow database row to a CardMonthAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetMonthlyTransferAmountReceiverRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransferReceiverAmount(card *db.GetMonthlyTransferAmountReceiverRow) *record.CardMonthAmount
	// ToMonthlyTransferReceiverAmounts maps a slice of GetMonthlyTransferAmountReceiverRow database rows to a slice of CardMonthAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetMonthlyTransferAmountReceiverRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
	ToMonthlyTransferReceiverAmounts(cards []*db.GetMonthlyTransferAmountReceiverRow) []*record.CardMonthAmount

	// ToYearlyTransferReceiverAmount maps a GetYearlyTransferAmountReceiverRow database row to a CardYearAmount domain model.
	//
	// Parameters:
	//   - card: A pointer to a GetYearlyTransferAmountReceiverRow representing the database row.
	//
	// Returns:
	//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTransferReceiverAmount(card *db.GetYearlyTransferAmountReceiverRow) *record.CardYearAmount
	// ToYearlyTransferReceiverAmounts maps a slice of GetYearlyTransferAmountReceiverRow database rows to a slice of CardYearAmount domain models.
	//
	// Parameters:
	//   - cards: A slice of pointers to GetYearlyTransferAmountReceiverRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
	ToYearlyTransferReceiverAmounts(cards []*db.GetYearlyTransferAmountReceiverRow) []*record.CardYearAmount
}
