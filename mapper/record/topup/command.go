package topuprecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// topupCommandRecordMapper provides methods to map Topup database rows to TopupRecord domain models for command operations.
type topupCommandRecordMapper struct {
}

// NewTopupCommandRecordMapper returns a new instance of topupCommandRecordMapper which provides methods to map Topup database rows to TopupRecord domain models for command operations.
func NewTopupCommandRecordMapper() TopupCommandRecordMapping {
	return &topupCommandRecordMapper{}
}

// ToTopupRecord maps a Topup database row to a TopupRecord domain model.
// It takes a pointer to a db.Topup as a parameter and returns a pointer to a record.TopupRecord.
// The function maps the fields of the db.Topup to the corresponding fields of the record.TopupRecord.
// If the DeletedAt field of the db.Topup is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
func (t *topupCommandRecordMapper) ToTopupRecord(topup *db.Topup) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}
