package transferrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// transferCommandRecordMapper provides methods to map Transfer database rows to TransferRecord domain models for command operations.
type transferCommandRecordMapper struct {
}

// NewTransferCommandRecordMapper returns a new instance of transferCommandRecordMapper which provides methods to map Transfer database rows to TransferRecord domain models for command operations.
func NewTransferCommandRecordMapper() TransferCommandRecordMapper {
	return &transferCommandRecordMapper{}
}

// ToTransferRecord maps a Transfer database row to a TransferRecord domain model.
//
// Parameters:
//   - transfer: A pointer to a Transfer representing the database row.
//
// Returns:
//   - A pointer to a TransferRecord containing the mapped data, including
//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
//     CreatedAt, UpdatedAt, and DeletedAt.
func (t *transferCommandRecordMapper) ToTransferRecord(transfer *db.Transfer) *record.TransferRecord {
	var deletedAt *string

	return &record.TransferRecord{
		ID:             int(transfer.TransferID),
		TransferNo:     transfer.TransferNo.String(),
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime.String(),
		CreatedAt:      transfer.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      transfer.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}
