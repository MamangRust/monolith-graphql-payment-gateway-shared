package transactionrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// transactionQueryRecordMapper is a struct that provides methods to map database rows to TransactionRecord domain models.
type transactionQueryRecordMapper struct {
}

// NewTransactionQueryRecordMapper returns a new instance of
// transactionQueryRecordMapper, which provides methods to map
// database rows to TransactionRecord domain models.
func NewTransactionQueryRecordMapper() TransactionQueryRecordMapper {
	return &transactionQueryRecordMapper{}
}

// ToTransactionRecord maps a Transaction database row to a TransactionRecord domain model.
//
// Parameters:
//   - transaction: A pointer to a Transaction representing the database row.
//
// Returns:
//   - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionQueryRecordMapper) ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord {
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

// ToTransactionsRecord maps a slice of Transaction database rows to a slice of TransactionRecord domain models.
//
// Parameters:
//   - transactions: A slice of pointers to Transaction representing the database rows.
//
// Returns:
//   - A slice of TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionQueryRecordMapper) ToTransactionsRecord(transactions []*db.Transaction) []*record.TransactionRecord {
	var records []*record.TransactionRecord
	for _, transaction := range transactions {
		records = append(records, s.ToTransactionRecord(transaction))
	}
	return records
}

// ToTransactionByCardNumberRecord maps a GetTransactionsByCardNumberRow database row to a TransactionRecord domain model.
//
// Parameters:
//   - transaction: A pointer to a GetTransactionsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionQueryRecordMapper) ToTransactionByCardNumberRecord(transaction *db.GetTransactionsByCardNumberRow) *record.TransactionRecord {
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

// ToTransactionsByCardNumberRecord maps a slice of GetTransactionsByCardNumberRow database rows
// to a slice of TransactionRecord domain models.
//
// Parameters:
//   - transactions: A slice of pointers to GetTransactionsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionQueryRecordMapper) ToTransactionsByCardNumberRecord(transactions []*db.GetTransactionsByCardNumberRow) []*record.TransactionRecord {
	var transactionRecords []*record.TransactionRecord
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, s.ToTransactionByCardNumberRecord(transaction))
	}
	return transactionRecords
}

// ToTransactionRecordAll maps a GetTransactionsRow database row to a TransactionRecord domain model.
//
// Parameters:
//   - transaction: A pointer to a GetTransactionsRow representing the database row.
//
// Returns:
//   - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionQueryRecordMapper) ToTransactionRecordAll(transaction *db.GetTransactionsRow) *record.TransactionRecord {
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

// ToTransactionsRecordAll maps a slice of GetTransactionsRow database rows to a slice of TransactionRecord domain models.
//
// Parameters:
//   - transactions: A slice of pointers to GetTransactionsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionQueryRecordMapper) ToTransactionsRecordAll(transactions []*db.GetTransactionsRow) []*record.TransactionRecord {
	var transactionRecords []*record.TransactionRecord
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, s.ToTransactionRecordAll(transaction))
	}
	return transactionRecords
}

// ToTransactionRecordActive maps a GetActiveTransactionsRow database row to a TransactionRecord domain model.
// It is intended for use with database rows that contain active transaction records.
// It returns a pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
// CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionQueryRecordMapper) ToTransactionRecordActive(transaction *db.GetActiveTransactionsRow) *record.TransactionRecord {
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

// ToTransactionsRecordActive maps a slice of GetActiveTransactionsRow database rows
// to a slice of TransactionRecord domain models.
//
// Parameters:
//   - transactions: A slice of pointers to GetActiveTransactionsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionQueryRecordMapper) ToTransactionsRecordActive(transactions []*db.GetActiveTransactionsRow) []*record.TransactionRecord {
	var transactionRecords []*record.TransactionRecord
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, s.ToTransactionRecordActive(transaction))
	}
	return transactionRecords
}

// ToTransactionRecordTrashed maps a GetTrashedTransactionsRow database row to a TransactionRecord domain model.
// It is intended for use with database rows that contain trashed transaction records.
// It returns a pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
// CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionQueryRecordMapper) ToTransactionRecordTrashed(transaction *db.GetTrashedTransactionsRow) *record.TransactionRecord {
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

// ToTransactionsRecordTrashed maps a slice of GetTrashedTransactionsRow database rows to a slice of
// TransactionRecord domain models.
//
// Parameters:
//   - transactions: A slice of pointers to GetTrashedTransactionsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionQueryRecordMapper) ToTransactionsRecordTrashed(transactions []*db.GetTrashedTransactionsRow) []*record.TransactionRecord {
	var transactionRecords []*record.TransactionRecord
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, s.ToTransactionRecordTrashed(transaction))
	}
	return transactionRecords
}
