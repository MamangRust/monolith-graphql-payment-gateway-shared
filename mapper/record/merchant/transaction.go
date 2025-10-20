package merchantrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantTransactionRecordMapper is an interface that provides methods to map Merchant database rows to MerchantTransactionsRecord domain models.
type merchantTransactionRecordMapper struct{}

// NewMerchantTransactionRecordMapper returns a new instance of merchantTransactionRecordMapper,
// which provides methods to map Merchant database rows to MerchantTransactionsRecord domain models.
func NewMerchantTransactionRecordMapper() MerchantTransactionRecordMapper {
	return &merchantTransactionRecordMapper{}
}

// ToMerchantTransactionRecord maps a FindAllTransactionsRow to a MerchantTransactionsRecord domain model.
//
// Parameters:
//   - merchant: A pointer to a FindAllTransactionsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantTransactionRecordMapper) ToMerchantTransactionRecord(merchant *db.FindAllTransactionsRow) *record.MerchantTransactionsRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantTransactionsRecord{
		TransactionID:   merchant.TransactionID,
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantID:      merchant.MerchantID,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime,
		CreatedAt:       merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:       merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:       deletedAt,
	}
}

// ToMerchantsTransactionRecord maps multiple FindAllTransactionsRow to multiple
// MerchantTransactionsRecord domain models.
//
// Parameters:
//   - merchants: A slice of pointers to FindAllTransactionsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantTransactionRecordMapper) ToMerchantsTransactionRecord(merchants []*db.FindAllTransactionsRow) []*record.MerchantTransactionsRecord {
	var records []*record.MerchantTransactionsRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantTransactionRecord(merchant))
	}
	return records
}

// ToMerchantTransactionByMerchantRecord maps a FindAllTransactionsByMerchantRow to a MerchantTransactionsRecord domain model.
//
// Parameters:
//   - merchant: A pointer to a FindAllTransactionsByMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantTransactionRecordMapper) ToMerchantTransactionByMerchantRecord(merchant *db.FindAllTransactionsByMerchantRow) *record.MerchantTransactionsRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantTransactionsRecord{
		TransactionID:   merchant.TransactionID,
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantID:      merchant.MerchantID,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime,
		CreatedAt:       merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:       merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:       deletedAt,
	}
}

// ToMerchantsTransactionByMerchantRecord maps a slice of FindAllTransactionsByMerchantRow to a slice
// of MerchantTransactionsRecord domain models.
//
// Parameters:
//   - merchants: A slice of pointers to FindAllTransactionsByMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantTransactionRecordMapper) ToMerchantsTransactionByMerchantRecord(merchants []*db.FindAllTransactionsByMerchantRow) []*record.MerchantTransactionsRecord {
	var records []*record.MerchantTransactionsRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantTransactionByMerchantRecord(merchant))
	}
	return records
}

// ToMerchantTransactionByApikeyRecord maps a FindAllTransactionsByApikeyRow to a MerchantTransactionsRecord domain model.
//
// Parameters:
//   - merchant: A pointer to a FindAllTransactionsByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantTransactionRecordMapper) ToMerchantTransactionByApikeyRecord(merchant *db.FindAllTransactionsByApikeyRow) *record.MerchantTransactionsRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantTransactionsRecord{
		TransactionID:   merchant.TransactionID,
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantID:      merchant.MerchantID,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime,
		CreatedAt:       merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:       merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:       deletedAt,
	}
}

// ToMerchantsTransactionByApikeyRecord maps a slice of FindAllTransactionsByApikeyRow to a slice of
// MerchantTransactionsRecord domain models.
//
// Parameters:
//   - merchants: A slice of pointers to FindAllTransactionsByApikeyRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantTransactionRecordMapper) ToMerchantsTransactionByApikeyRecord(merchants []*db.FindAllTransactionsByApikeyRow) []*record.MerchantTransactionsRecord {
	var records []*record.MerchantTransactionsRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantTransactionByApikeyRecord(merchant))
	}
	return records
}
