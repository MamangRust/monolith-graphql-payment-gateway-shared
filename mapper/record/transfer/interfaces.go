package transferrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TransferBaseRecordMapper maps raw Transfer rows.
type TransferBaseRecordMapper interface {
	// ToTransferRecord maps a Transfer database row to a TransferRecord domain model.
	//
	// Parameters:
	//   - transfer: A pointer to a Transfer representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferRecord containing the mapped data, including
	//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
	//     CreatedAt, UpdatedAt, and DeletedAt.
	ToTransferRecord(transfer *db.Transfer) *record.TransferRecord
}

// TransferCommandRecordMapper maps raw Transfer rows for command operations.
type TransferCommandRecordMapper interface {
	TransferBaseRecordMapper
}

// TransferQueryRecordMapper maps query result rows to Transfer domain models.
type TransferQueryRecordMapper interface {
	TransferBaseRecordMapper

	ToTransfersRecord(transfers []*db.Transfer) []*record.TransferRecord

	// ToTransferRecordAll maps a GetTransfersRow database row to a TransferRecord domain model.
	//
	// Parameters:
	//   - transfer: A pointer to a GetTransfersRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferRecord containing the mapped data, including
	//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
	//     CreatedAt, UpdatedAt, and DeletedAt.
	ToTransferRecordAll(transfer *db.GetTransfersRow) *record.TransferRecord
	// ToTransfersRecordAll maps a slice of GetTransfersRow database rows to a slice of TransferRecord domain models.
	//
	// Parameters:
	//   - transfers: A slice of pointers to GetTransfersRow structs representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferRecord structs containing the mapped data.
	ToTransfersRecordAll(transfers []*db.GetTransfersRow) []*record.TransferRecord
	// ToTransferRecordActive maps a GetActiveTransfersRow database row to a TransferRecord domain model.
	//
	// Parameters:
	//   - transfer: A pointer to a GetActiveTransfersRow representing the active transfer database row.
	//
	// Returns:
	//   - A pointer to a TransferRecord containing the mapped data, including
	//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
	//     CreatedAt, UpdatedAt, and DeletedAt.
	ToTransferRecordActive(transfer *db.GetActiveTransfersRow) *record.TransferRecord
	// ToTransfersRecordActive maps a slice of GetActiveTransfersRow database rows to a slice of TransferRecord domain models.
	//
	// Parameters:
	//   - transfers: A slice of pointers to GetActiveTransfersRow structs representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferRecord structs containing the mapped data.
	ToTransfersRecordActive(transfers []*db.GetActiveTransfersRow) []*record.TransferRecord
	// ToTransferRecordTrashed maps a GetTrashedTransfersRow database row to a TransferRecord domain model.
	//
	// Parameters:
	//   - transfer: A pointer to a GetTrashedTransfersRow representing the active transfer database row.
	//
	// Returns:
	//   - A pointer to a TransferRecord containing the mapped data, including
	//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
	//     CreatedAt, UpdatedAt, and DeletedAt.
	ToTransferRecordTrashed(transfer *db.GetTrashedTransfersRow) *record.TransferRecord
	// ToTransfersRecordTrashed maps a slice of GetTrashedTransfersRow database rows to a slice of TransferRecord domain models.
	//
	// Parameters:
	//   - transfers: A slice of pointers to GetTrashedTransfersRow structs representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferRecord structs containing the mapped data.
	ToTransfersRecordTrashed(transfers []*db.GetTrashedTransfersRow) []*record.TransferRecord
}
