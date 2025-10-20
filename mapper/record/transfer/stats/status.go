package transferstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TransferStatisticStatusRecordMapper maps transfer status (success/failed) to domain models.
type transferStatisticStatusRecordMapper struct{}

// NewTransferStatisticStatusRecordMapper creates a new instance of the
// TransferStatisticStatusRecordMapper interface which provides mapping
// functionality from database rows to domain models.
func NewTransferStatisticStatusRecordMapper() TransferStatisticStatusRecordMapper {
	return &transferStatisticStatusRecordMapper{}
}

// ToTransferRecordMonthStatusSuccess maps a GetMonthTransferStatusSuccessRow database row
// to a TransferRecordMonthStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to a GetMonthTransferStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordMonthStatusSuccess containing the mapped data, including
//     Year, Month, TotalSuccess, and TotalAmount.
func (t *transferStatisticStatusRecordMapper) ToTransferRecordMonthStatusSuccess(s *db.GetMonthTransferStatusSuccessRow) *record.TransferRecordMonthStatusSuccess {
	return &record.TransferRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransferRecordsMonthStatusSuccess maps a slice of GetMonthTransferStatusSuccessRow database rows
// to a slice of TransferRecordMonthStatusSuccess domain models.
//
// Parameters:
//   - Transfers: A slice of pointers to GetMonthTransferStatusSuccessRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transferStatisticStatusRecordMapper) ToTransferRecordsMonthStatusSuccess(Transfers []*db.GetMonthTransferStatusSuccessRow) []*record.TransferRecordMonthStatusSuccess {
	var TransferRecords []*record.TransferRecordMonthStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordMonthStatusSuccess(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordYearStatusSuccess maps a GetYearlyTransferStatusSuccessRow database row
// to a TransferRecordYearStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to a GetYearlyTransferStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transferStatisticStatusRecordMapper) ToTransferRecordYearStatusSuccess(s *db.GetYearlyTransferStatusSuccessRow) *record.TransferRecordYearStatusSuccess {
	return &record.TransferRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransferRecordsYearStatusSuccess maps a slice of GetYearlyTransferStatusSuccessRow database rows
// to a slice of TransferRecordYearStatusSuccess domain models.
//
// Parameters:
//   - Transfers: A slice of pointers to GetYearlyTransferStatusSuccessRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transferStatisticStatusRecordMapper) ToTransferRecordsYearStatusSuccess(Transfers []*db.GetYearlyTransferStatusSuccessRow) []*record.TransferRecordYearStatusSuccess {
	var TransferRecords []*record.TransferRecordYearStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordYearStatusSuccess(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordMonthStatusFailed maps a GetMonthTransferStatusFailedRow database row
// to a TransferRecordMonthStatusFailed domain model.
//
// Parameters:
//   - s: A pointer to a GetMonthTransferStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transferStatisticStatusRecordMapper) ToTransferRecordMonthStatusFailed(s *db.GetMonthTransferStatusFailedRow) *record.TransferRecordMonthStatusFailed {
	return &record.TransferRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferRecordsMonthStatusFailed maps a slice of GetMonthTransferStatusFailedRow database rows
// to a slice of TransferRecordMonthStatusFailed domain models.
//
// Parameters:
//   - Transfers: A slice of pointers to GetMonthTransferStatusFailedRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transferStatisticStatusRecordMapper) ToTransferRecordsMonthStatusFailed(Transfers []*db.GetMonthTransferStatusFailedRow) []*record.TransferRecordMonthStatusFailed {
	var TransferRecords []*record.TransferRecordMonthStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordMonthStatusFailed(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordYearStatusFailed maps a GetYearlyTransferStatusFailedRow database row
// to a TransferRecordYearStatusFailed domain model.
//
// Parameters:
//   - s: A pointer to a GetYearlyTransferStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transferStatisticStatusRecordMapper) ToTransferRecordYearStatusFailed(s *db.GetYearlyTransferStatusFailedRow) *record.TransferRecordYearStatusFailed {
	return &record.TransferRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferRecordsYearStatusFailed maps a slice of GetYearlyTransferStatusFailedRow database rows
// to a slice of TransferRecordYearStatusFailed domain models.
//
// Parameters:
//   - Transfers: A slice of pointers to GetYearlyTransferStatusFailedRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transferStatisticStatusRecordMapper) ToTransferRecordsYearStatusFailed(Transfers []*db.GetYearlyTransferStatusFailedRow) []*record.TransferRecordYearStatusFailed {
	var TransferRecords []*record.TransferRecordYearStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordYearStatusFailed(Transfer))
	}

	return TransferRecords
}
