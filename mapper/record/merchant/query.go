package merchantrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantQueryRecordMapper provides methods to map Merchant database rows to
// MerchantRecord domain models for query operations.
type merchantQueryRecordMapper struct{}

// NewMerchantQueryRecordMapper returns a new instance of
// merchantQueryRecordMapper which provides methods to map Merchant database
// rows to MerchantRecord domain models for query operations.
func NewMerchantQueryRecordMapper() MerchantQueryRecordMapper {
	return &merchantQueryRecordMapper{}
}

// ToMerchantRecord maps a Merchant to a MerchantRecord domain model.
//
// Parameters:
//   - merchant: A pointer to a Merchant representing the database row.
//
// Returns:
//   - A pointer to a MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantQueryRecordMapper) ToMerchantRecord(merchant *db.Merchant) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

func (m *merchantQueryRecordMapper) ToMerchantsRecord(merchants []*db.Merchant) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantRecord(merchant))
	}
	return records
}

// ToMerchantGetAllRecord maps a GetMerchantsRow to a MerchantRecord domain model.
//
// Parameters:
//   - merchant: A pointer to a GetMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantQueryRecordMapper) ToMerchantGetAllRecord(merchant *db.GetMerchantsRow) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

// ToMerchantsGetAllRecord maps a slice of GetMerchantsRow to a slice of
// MerchantRecord domain models.
//
// Parameters:
//   - merchants: A slice of pointers to GetMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantQueryRecordMapper) ToMerchantsGetAllRecord(merchants []*db.GetMerchantsRow) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantGetAllRecord(merchant))
	}
	return records
}

// ToMerchantActiveRecord maps a GetActiveMerchantsRow to a MerchantRecord domain model.
//
// Parameters:
//   - merchant: A pointer to a GetActiveMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantQueryRecordMapper) ToMerchantActiveRecord(merchant *db.GetActiveMerchantsRow) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

// ToMerchantsActiveRecord maps a slice of GetActiveMerchantsRow to a slice of
// MerchantRecord domain models.
//
// Parameters:
//   - merchants: A slice of pointers to GetActiveMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantQueryRecordMapper) ToMerchantsActiveRecord(merchants []*db.GetActiveMerchantsRow) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantActiveRecord(merchant))
	}
	return records
}

// ToMerchantTrashedRecord maps a GetTrashedMerchantsRow to a MerchantRecord domain model.
//
// Parameters:
//   - merchant: A pointer to a GetTrashedMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantQueryRecordMapper) ToMerchantTrashedRecord(merchant *db.GetTrashedMerchantsRow) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

// ToMerchantsTrashedRecord maps a slice of GetTrashedMerchantsRow to a slice of
// MerchantRecord domain models.
//
// Parameters:
//   - merchants: A slice of pointers to GetTrashedMerchantsRow representing the trashed database rows.
//
// Returns:
//   - A slice of pointers to MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantQueryRecordMapper) ToMerchantsTrashedRecord(merchants []*db.GetTrashedMerchantsRow) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantTrashedRecord(merchant))
	}
	return records
}
