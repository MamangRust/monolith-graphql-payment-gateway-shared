package transferstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TransferStatisticStatusByCardRecordMapper maps transfer status (success/failed) filtered by card_number to domain models.
type transferStatisticStatusByCardRecordMapper struct {
}

// NewTransferStatisticStatusByCardRecordMapper returns a new instance of
// TransferStatisticStatusByCardRecordMapper which provides methods to map
// transfer status (success/failed) filtered by card_number to domain models.
func NewTransferStatisticStatusByCardRecordMapper() TransferStatisticStatusByCardRecordMapper {
	return &transferStatisticStatusByCardRecordMapper{}
}

// ToTransferRecordMonthStatusSuccessCardNumber maps a GetMonthTransferStatusSuccessCardNumberRow database row
// to a TransferRecordMonthStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to GetMonthTransferStatusSuccessCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transferStatisticStatusByCardRecordMapper) ToTransferRecordMonthStatusSuccessCardNumber(s *db.GetMonthTransferStatusSuccessCardNumberRow) *record.TransferRecordMonthStatusSuccess {
	return &record.TransferRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransferRecordsMonthStatusSuccessCardNumber maps a slice of GetMonthTransferStatusSuccessCardNumberRow database rows
// to a slice of TransferRecordMonthStatusSuccess domain models.
//
// Parameters:
//   - Transfers: A slice of pointers to GetMonthTransferStatusSuccessCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transferStatisticStatusByCardRecordMapper) ToTransferRecordsMonthStatusSuccessCardNumber(Transfers []*db.GetMonthTransferStatusSuccessCardNumberRow) []*record.TransferRecordMonthStatusSuccess {
	var TransferRecords []*record.TransferRecordMonthStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordMonthStatusSuccessCardNumber(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordYearStatusSuccessCardNumber maps a GetYearlyTransferStatusSuccessCardNumberRow database row
// to a TransferRecordYearStatusSuccess domain model.
//
// Parameters:
//   - s: A pointer to a GetYearlyTransferStatusSuccessCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transferStatisticStatusByCardRecordMapper) ToTransferRecordYearStatusSuccessCardNumber(s *db.GetYearlyTransferStatusSuccessCardNumberRow) *record.TransferRecordYearStatusSuccess {
	return &record.TransferRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransferRecordsYearStatusSuccessCardNumber maps a slice of GetYearlyTransferStatusSuccessCardNumberRow database rows
// to a slice of TransferRecordYearStatusSuccess domain models.
//
// Parameters:
//   - Transfers: A slice of pointers to GetYearlyTransferStatusSuccessCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transferStatisticStatusByCardRecordMapper) ToTransferRecordsYearStatusSuccessCardNumber(Transfers []*db.GetYearlyTransferStatusSuccessCardNumberRow) []*record.TransferRecordYearStatusSuccess {
	var TransferRecords []*record.TransferRecordYearStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordYearStatusSuccessCardNumber(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordMonthStatusFailedCardNumber maps a GetMonthTransferStatusFailedCardNumberRow database row
// to a TransferRecordMonthStatusFailed domain model.
//
// Parameters:
//   - s: A pointer to a GetMonthTransferStatusFailedCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transferStatisticStatusByCardRecordMapper) ToTransferRecordMonthStatusFailedCardNumber(s *db.GetMonthTransferStatusFailedCardNumberRow) *record.TransferRecordMonthStatusFailed {
	return &record.TransferRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferRecordsMonthStatusFailedCardNumber maps a slice of GetMonthTransferStatusFailedCardNumberRow database rows
// to a slice of TransferRecordMonthStatusFailed domain models.
//
// Parameters:
//   - Transfers: A slice of pointers to GetMonthTransferStatusFailedCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transferStatisticStatusByCardRecordMapper) ToTransferRecordsMonthStatusFailedCardNumber(Transfers []*db.GetMonthTransferStatusFailedCardNumberRow) []*record.TransferRecordMonthStatusFailed {
	var TransferRecords []*record.TransferRecordMonthStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordMonthStatusFailedCardNumber(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordYearStatusFailedCardNumber maps a GetYearlyTransferStatusFailedCardNumberRow database row
// to a TransferRecordYearStatusFailed domain model.
//
// Parameters:
//   - s: A pointer to a GetYearlyTransferStatusFailedCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transferStatisticStatusByCardRecordMapper) ToTransferRecordYearStatusFailedCardNumber(s *db.GetYearlyTransferStatusFailedCardNumberRow) *record.TransferRecordYearStatusFailed {
	return &record.TransferRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferRecordsYearStatusFailedCardNumber maps a slice of GetYearlyTransferStatusFailedCardNumberRow database rows
// to a slice of TransferRecordYearStatusFailed domain models.
//
// Parameters:
//   - Transfers: A slice of pointers to GetYearlyTransferStatusFailedCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transferStatisticStatusByCardRecordMapper) ToTransferRecordsYearStatusFailedCardNumber(Transfers []*db.GetYearlyTransferStatusFailedCardNumberRow) []*record.TransferRecordYearStatusFailed {
	var TransferRecords []*record.TransferRecordYearStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordYearStatusFailedCardNumber(Transfer))
	}

	return TransferRecords
}
