package transferrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// transferQueryRecordMapper provides methods to map Transfer database rows to TransferRecord domain models.
type transferQueryRecordMapper struct {
}

// NewTransferQueryRecordMapper returns a new instance of transferQueryRecordMapper which provides methods to map Transfer database rows to TransferRecord domain models.
func NewTransferQueryRecordMapper() TransferQueryRecordMapper {
	return &transferQueryRecordMapper{}
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
func (t *transferQueryRecordMapper) ToTransferRecord(transfer *db.Transfer) *record.TransferRecord {
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

func (t *transferQueryRecordMapper) ToTransfersRecord(transfers []*db.Transfer) []*record.TransferRecord {
	var transferRecords []*record.TransferRecord

	for _, transfer := range transfers {
		transferRecords = append(transferRecords, t.ToTransferRecord(transfer))
	}

	return transferRecords
}

// ToTransferRecordAll maps a GetTransfersRow database row to a TransferRecord domain model.
//
// Parameters:
//   - transfer: A pointer to a GetTransfersRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecord containing the mapped data, including
//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
//     CreatedAt, UpdatedAt, and DeletedAt.
func (t *transferQueryRecordMapper) ToTransferRecordAll(transfer *db.GetTransfersRow) *record.TransferRecord {
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

// ToTransfersRecordAll maps a slice of GetTransfersRow database rows to a slice of TransferRecord domain models.
//
// Parameters:
//   - transfers: A slice of pointers to GetTransfersRow structs representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecord structs containing the mapped data.
func (s *transferQueryRecordMapper) ToTransfersRecordAll(transfers []*db.GetTransfersRow) []*record.TransferRecord {
	var transferRecords []*record.TransferRecord

	for _, transfer := range transfers {
		transferRecords = append(transferRecords, s.ToTransferRecordAll(transfer))
	}

	return transferRecords
}

// ToTransferRecordActive maps a GetActiveTransfersRow database row to a TransferRecord domain model.
//
// Parameters:
//   - transfer: A pointer to a GetActiveTransfersRow representing the active transfer database row.
//
// Returns:
//   - A pointer to a TransferRecord containing the mapped data, including
//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
//     CreatedAt, UpdatedAt, and DeletedAt.
func (t *transferQueryRecordMapper) ToTransferRecordActive(transfer *db.GetActiveTransfersRow) *record.TransferRecord {
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

// ToTransfersRecordActive maps a slice of GetActiveTransfersRow database rows to a slice of TransferRecord domain models.
//
// Parameters:
//   - transfers: A slice of pointers to GetActiveTransfersRow structs representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecord structs containing the mapped data.
func (s *transferQueryRecordMapper) ToTransfersRecordActive(transfers []*db.GetActiveTransfersRow) []*record.TransferRecord {
	var transferRecords []*record.TransferRecord

	for _, transfer := range transfers {
		transferRecords = append(transferRecords, s.ToTransferRecordActive(transfer))
	}

	return transferRecords
}

// ToTransferRecordTrashed maps a GetTrashedTransfersRow database row to a TransferRecord domain model.
//
// Parameters:
//   - transfer: A pointer to a GetTrashedTransfersRow representing the active transfer database row.
//
// Returns:
//   - A pointer to a TransferRecord containing the mapped data, including
//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
//     CreatedAt, UpdatedAt, and DeletedAt.
func (t *transferQueryRecordMapper) ToTransferRecordTrashed(transfer *db.GetTrashedTransfersRow) *record.TransferRecord {
	var deletedAt *string

	if transfer.DeletedAt.Valid {
		formatedDeletedAt := transfer.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

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

// ToTransfersRecordTrashed maps a slice of GetTrashedTransfersRow database rows to a slice of TransferRecord domain models.
//
// Parameters:
//   - transfers: A slice of pointers to GetTrashedTransfersRow structs representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecord structs containing the mapped data.
func (s *transferQueryRecordMapper) ToTransfersRecordTrashed(transfers []*db.GetTrashedTransfersRow) []*record.TransferRecord {
	var transferRecords []*record.TransferRecord

	for _, transfer := range transfers {
		transferRecords = append(transferRecords, s.ToTransferRecordTrashed(transfer))
	}

	return transferRecords
}
