package merchantrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantBaseMapper provides methods to map Merchant database rows to MerchantRecord domain models.
type MerchantBaseRecordMapper interface {
	// ToMerchantRecord maps a Merchant to a MerchantRecord domain model.
	//
	// Parameters:
	//   - merchant: A pointer to a Merchant representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantRecord containing the mapped data, including
	//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantRecord(merchant *db.Merchant) *record.MerchantRecord
}

// MerchantQueryMapper provides methods to map Merchant database rows to MerchantRecord domain models for query operations.
type MerchantQueryRecordMapper interface {
	MerchantBaseRecordMapper

	ToMerchantsRecord(merchants []*db.Merchant) []*record.MerchantRecord

	// ToMerchantGetAllRecord maps a GetMerchantsRow to a MerchantRecord domain model.
	//
	// Parameters:
	//   - merchant: A pointer to a GetMerchantsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantRecord containing the mapped data, including
	//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantGetAllRecord(merchant *db.GetMerchantsRow) *record.MerchantRecord

	// ToMerchantsGetAllRecord maps a slice of GetMerchantsRow to a slice of
	// MerchantRecord domain models.
	//
	// Parameters:
	//   - merchants: A slice of pointers to GetMerchantsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantRecord containing the mapped data, including
	//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantsGetAllRecord(merchants []*db.GetMerchantsRow) []*record.MerchantRecord

	// ToMerchantActiveRecord maps a GetActiveMerchantsRow to a MerchantRecord domain model.
	//
	// Parameters:
	//   - merchant: A pointer to a GetActiveMerchantsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantRecord containing the mapped data, including
	//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantActiveRecord(merchant *db.GetActiveMerchantsRow) *record.MerchantRecord

	// ToMerchantsActiveRecord maps a slice of GetActiveMerchantsRow to a slice of
	// MerchantRecord domain models.
	//
	// Parameters:
	//   - merchants: A slice of pointers to GetActiveMerchantsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantRecord containing the mapped data, including
	//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantsActiveRecord(merchants []*db.GetActiveMerchantsRow) []*record.MerchantRecord

	// ToMerchantTrashedRecord maps a GetTrashedMerchantsRow to a MerchantRecord domain model.
	//
	// Parameters:
	//   - merchant: A pointer to a GetTrashedMerchantsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantRecord containing the mapped data, including
	//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantTrashedRecord(merchant *db.GetTrashedMerchantsRow) *record.MerchantRecord

	// ToMerchantsTrashedRecord maps a slice of GetTrashedMerchantsRow to a slice of
	// MerchantRecord domain models.
	//
	// Parameters:
	//   - merchants: A slice of pointers to GetTrashedMerchantsRow representing the trashed database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantRecord containing the mapped data, including
	//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantsTrashedRecord(merchants []*db.GetTrashedMerchantsRow) []*record.MerchantRecord
}

// MerchantCommandMapper provides methods to map Merchant database rows to MerchantRecord domain models for command operations.
type MerchantCommandRecordMapper interface {
	MerchantBaseRecordMapper
}

// MerchantTransactionRecordMapper provides methods to map Merchant database rows to MerchantRecord domain models for transaction operations.
type MerchantTransactionRecordMapper interface {
	// ToMerchantTransactionRecord maps a FindAllTransactionsRow to a MerchantTransactionsRecord domain model.
	//
	// Parameters:
	//   - merchant: A pointer to a FindAllTransactionsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantTransactionsRecord containing the mapped data, including
	//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
	//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantTransactionRecord(merchant *db.FindAllTransactionsRow) *record.MerchantTransactionsRecord

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
	ToMerchantsTransactionRecord(merchants []*db.FindAllTransactionsRow) []*record.MerchantTransactionsRecord

	// ToMerchantTransactionByMerchantRecord maps a FindAllTransactionsByMerchantRow to a MerchantTransactionsRecord domain model.
	//
	// Parameters:
	//   - merchant: A pointer to a FindAllTransactionsByMerchantRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantTransactionsRecord containing the mapped data, including
	//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
	//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantTransactionByMerchantRecord(merchant *db.FindAllTransactionsByMerchantRow) *record.MerchantTransactionsRecord

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
	ToMerchantsTransactionByMerchantRecord(merchants []*db.FindAllTransactionsByMerchantRow) []*record.MerchantTransactionsRecord

	// ToMerchantTransactionByApikeyRecord maps a FindAllTransactionsByApikeyRow to a MerchantTransactionsRecord domain model.
	//
	// Parameters:
	//   - merchant: A pointer to a FindAllTransactionsByApikeyRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantTransactionsRecord containing the mapped data, including
	//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
	//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
	ToMerchantTransactionByApikeyRecord(merchant *db.FindAllTransactionsByApikeyRow) *record.MerchantTransactionsRecord

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
	ToMerchantsTransactionByApikeyRecord(merchants []*db.FindAllTransactionsByApikeyRow) []*record.MerchantTransactionsRecord
}
