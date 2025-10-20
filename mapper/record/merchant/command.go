package merchantrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// merchantCommandMapper provides methods to map Merchant database rows to MerchantRecord domain models.
type merchantCommandMapper struct {
}

// NewMerchantCommandMapper returns a new instance of
// merchantCommandMapper which provides methods to map
// Merchant database rows to MerchantRecord domain models.
func NewMerchantCommandMapper() MerchantCommandRecordMapper {
	return &merchantCommandMapper{}
}

// ToMerchantRecord maps a Merchant to a MerchantRecord domain model.
//
// Parameters:
//   - merchant: A pointer to a Merchant representing the database row.
//
// Returns:
//   - A pointer to a MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantCommandMapper) ToMerchantRecord(merchant *db.Merchant) *record.MerchantRecord {
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
