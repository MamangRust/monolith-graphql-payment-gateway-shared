package transferstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TransferStatisticStatusRecordMapper maps transfer status (success/failed) to domain models.
type TransferStatisticStatusRecordMapper interface {
	// ToTransferRecordMonthStatusSuccess maps a GetMonthTransferStatusSuccessRow database row
	// to a TransferRecordMonthStatusSuccess domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetMonthTransferStatusSuccessRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferRecordMonthStatusSuccess containing the mapped data, including
	//     Year, Month, TotalSuccess, and TotalAmount.
	ToTransferRecordMonthStatusSuccess(s *db.GetMonthTransferStatusSuccessRow) *record.TransferRecordMonthStatusSuccess
	// ToTransferRecordsMonthStatusSuccess maps a slice of GetMonthTransferStatusSuccessRow database rows
	// to a slice of TransferRecordMonthStatusSuccess domain models.
	//
	// Parameters:
	//   - Transfers: A slice of pointers to GetMonthTransferStatusSuccessRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
	//     Month, TotalSuccess, and TotalAmount.
	ToTransferRecordsMonthStatusSuccess(s []*db.GetMonthTransferStatusSuccessRow) []*record.TransferRecordMonthStatusSuccess

	// ToTransferRecordYearStatusSuccess maps a GetYearlyTransferStatusSuccessRow database row
	// to a TransferRecordYearStatusSuccess domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetYearlyTransferStatusSuccessRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferRecordYearStatusSuccess containing the mapped data, including Year,
	//     TotalSuccess, and TotalAmount.
	ToTransferRecordYearStatusSuccess(s *db.GetYearlyTransferStatusSuccessRow) *record.TransferRecordYearStatusSuccess
	// ToTransferRecordsYearStatusSuccess maps a slice of GetYearlyTransferStatusSuccessRow database rows
	// to a slice of TransferRecordYearStatusSuccess domain models.
	//
	// Parameters:
	//   - Transfers: A slice of pointers to GetYearlyTransferStatusSuccessRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferRecordYearStatusSuccess containing the mapped data, including Year,
	//     TotalSuccess, and TotalAmount.
	ToTransferRecordsYearStatusSuccess(s []*db.GetYearlyTransferStatusSuccessRow) []*record.TransferRecordYearStatusSuccess

	// ToTransferRecordMonthStatusFailed maps a GetMonthTransferStatusFailedRow database row
	// to a TransferRecordMonthStatusFailed domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetMonthTransferStatusFailedRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferRecordMonthStatusFailed containing the mapped data, including Year,
	//     Month, TotalFailed, and TotalAmount.
	ToTransferRecordMonthStatusFailed(s *db.GetMonthTransferStatusFailedRow) *record.TransferRecordMonthStatusFailed
	// ToTransferRecordsMonthStatusFailed maps a slice of GetMonthTransferStatusFailedRow database rows
	// to a slice of TransferRecordMonthStatusFailed domain models.
	//
	// Parameters:
	//   - Transfers: A slice of pointers to GetMonthTransferStatusFailedRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferRecordMonthStatusFailed containing the mapped data, including Year,
	//     Month, TotalFailed, and TotalAmount.
	ToTransferRecordsMonthStatusFailed(s []*db.GetMonthTransferStatusFailedRow) []*record.TransferRecordMonthStatusFailed
	// ToTransferRecordYearStatusFailed maps a GetYearlyTransferStatusFailedRow database row
	// to a TransferRecordYearStatusFailed domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetYearlyTransferStatusFailedRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferRecordYearStatusFailed containing the mapped data, including Year,
	//     TotalFailed, and TotalAmount.
	ToTransferRecordYearStatusFailed(s *db.GetYearlyTransferStatusFailedRow) *record.TransferRecordYearStatusFailed
	// ToTransferRecordsYearStatusFailed maps a slice of GetYearlyTransferStatusFailedRow database rows
	// to a slice of TransferRecordYearStatusFailed domain models.
	//
	// Parameters:
	//   - Transfers: A slice of pointers to GetYearlyTransferStatusFailedRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferRecordYearStatusFailed containing the mapped data, including Year,
	//     TotalFailed, and TotalAmount.
	ToTransferRecordsYearStatusFailed(s []*db.GetYearlyTransferStatusFailedRow) []*record.TransferRecordYearStatusFailed
}

// TransferStatisticAmountRecordMapper maps transfer amount statistics to domain models.
type TransferStatisticAmountRecordMapper interface {
	// ToTransferMonthAmount maps a GetMonthlyTransferAmountsRow database row
	// to a TransferMonthAmount domain model.
	//
	// Parameters:
	//   - ss: A pointer to a GetMonthlyTransferAmountsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferMonthAmount containing the mapped data, including Month
	//     and TotalAmount.
	ToTransferMonthAmount(s *db.GetMonthlyTransferAmountsRow) *record.TransferMonthAmount
	// ToTransferMonthAmounts maps a slice of GetMonthlyTransferAmountsRow database rows
	// to a slice of TransferMonthAmount domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetMonthlyTransferAmountsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferMonthAmount containing the mapped data, including Month
	//     and TotalAmount.
	ToTransferMonthAmounts(s []*db.GetMonthlyTransferAmountsRow) []*record.TransferMonthAmount
	// ToTransferYearAmount maps a GetYearlyTransferAmountsRow database row to a TransferYearAmount domain model.
	//
	// Parameters:
	//   - ss: A pointer to a GetYearlyTransferAmountsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.
	ToTransferYearAmount(s *db.GetYearlyTransferAmountsRow) *record.TransferYearAmount
	// ToTransferYearAmounts maps a slice of GetYearlyTransferAmountsRow database rows
	// to a slice of TransferYearAmount domain models.
	//
	// Parameters:
	//   - ss: A slice of pointers to GetYearlyTransferAmountsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.
	ToTransferYearAmounts(s []*db.GetYearlyTransferAmountsRow) []*record.TransferYearAmount
}
