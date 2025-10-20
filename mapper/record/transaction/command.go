package transactionrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// transactionComandRecordMapper is a struct that implements the TransactionCommandRecordMapper interface.
type transactionComandRecordMapper struct {
}

// NewTransactionCommandRecordMapper returns a new instance of transactionCommandRecordMapper,
// which provides methods to map Transaction database rows to TransactionRecord domain models
// for command operations.
func NewTransactionCommandRecordMapper() TransactionCommandRecordMapper {
	return &transactionComandRecordMapper{}
}

// ToTransactionRecord maps a Transaction database row to a TransactionRecord domain model.
//
// Parameters:
//   - transaction: A pointer to a Transaction representing the database row.
//
// Returns:
//   - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionComandRecordMapper) ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord {
	var deletedAt *string

	if transaction.DeletedAt.Valid {
		formatedDeletedAt := transaction.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.TransactionRecord{
		ID:              int(transaction.TransactionID),
		TransactionNo:   transaction.TransactionNo.String(),
		CardNumber:      transaction.CardNumber,
		Amount:          int(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		MerchantID:      int(transaction.MerchantID),
		TransactionTime: transaction.TransactionTime.Format("2006-01-02 15:04:05"),
		CreatedAt:       transaction.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:       transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:       deletedAt,
	}
}
