package transactionrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TransactionBaseRecordMapper provides methods to map database rows to TransactionRecord domain models.
type TransactionBaseRecordMapper interface {
	// ToTransactionRecord maps a Transaction database row to a TransactionRecord domain model.
	//
	// Parameters:
	//   - transaction: A pointer to a Transaction representing the database row.
	//
	// Returns:
	//   - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
	//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord
}

// TransactionQueryRecordMapper provides methods to map database rows to TransactionRecord domain models.
type TransactionQueryRecordMapper interface {
	TransactionBaseRecordMapper

	ToTransactionsRecord(transactions []*db.Transaction) []*record.TransactionRecord

	// ToTransactionByCardNumberRecord maps a GetTransactionsByCardNumberRow database row to a TransactionRecord domain model.
	//
	// Parameters:
	//   - transaction: A pointer to a GetTransactionsByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
	//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToTransactionByCardNumberRecord(transaction *db.GetTransactionsByCardNumberRow) *record.TransactionRecord
	// ToTransactionsByCardNumberRecord maps a slice of GetTransactionsByCardNumberRow database rows
	// to a slice of TransactionRecord domain models.
	//
	// Parameters:
	//   - transactions: A slice of pointers to GetTransactionsByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
	//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToTransactionsByCardNumberRecord(transactions []*db.GetTransactionsByCardNumberRow) []*record.TransactionRecord

	// ToTransactionRecordAll maps a GetTransactionsRow database row to a TransactionRecord domain model.
	//
	// Parameters:
	//   - transaction: A pointer to a GetTransactionsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
	//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToTransactionRecordAll(transaction *db.GetTransactionsRow) *record.TransactionRecord
	// ToTransactionsRecordAll maps a slice of GetTransactionsRow database rows to a slice of TransactionRecord domain models.
	//
	// Parameters:
	//   - transactions: A slice of pointers to GetTransactionsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
	//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToTransactionsRecordAll(transactions []*db.GetTransactionsRow) []*record.TransactionRecord

	// ToTransactionRecordActive maps a GetActiveTransactionsRow database row to a TransactionRecord domain model.
	// It is intended for use with database rows that contain active transaction records.
	// It returns a pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
	// CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToTransactionRecordActive(transaction *db.GetActiveTransactionsRow) *record.TransactionRecord
	// ToTransactionsRecordActive maps a slice of GetActiveTransactionsRow database rows
	// to a slice of TransactionRecord domain models.
	//
	// Parameters:
	//   - transactions: A slice of pointers to GetActiveTransactionsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
	//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToTransactionsRecordActive(transactions []*db.GetActiveTransactionsRow) []*record.TransactionRecord
	// ToTransactionRecordTrashed maps a GetTrashedTransactionsRow database row to a TransactionRecord domain model.
	// It is intended for use with database rows that contain trashed transaction records.
	// It returns a pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
	// CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToTransactionRecordTrashed(transaction *db.GetTrashedTransactionsRow) *record.TransactionRecord
	// ToTransactionsRecordTrashed maps a slice of GetTrashedTransactionsRow database rows to a slice of
	// TransactionRecord domain models.
	//
	// Parameters:
	//   - transactions: A slice of pointers to GetTrashedTransactionsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
	//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToTransactionsRecordTrashed(transactions []*db.GetTrashedTransactionsRow) []*record.TransactionRecord
}

// TransactionCommandRecordMapper provides methods to map database rows to TransactionRecord domain models.
type TransactionCommandRecordMapper interface {
	TransactionBaseRecordMapper
}
