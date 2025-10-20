package transferresponsemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// transferStatsStatusResponseMapper is a struct that implements the TransferStatsStatusResponseMapper interface.
type transferStatsStatusResponseMapper struct {
}

// NewTransferStatsStatusResponseMapper creates a new instance of transferStatsStatusResponseMapper.
// It returns a TransferStatsStatusResponseMapper interface which provides methods for converting
// transfer record statuses (both success and failed) into their corresponding API-compatible response types.
func NewTransferStatsStatusResponseMapper() TransferStatsStatusResponseMapper {
	return &transferStatsStatusResponseMapper{}
}

// ToTransferResponseMonthStatusSuccess converts a monthly success status transfer record
// into a TransferResponseMonthStatusSuccess.
//
// Args:
//   - s: A pointer to a TransferRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferResponseMonthStatusSuccess containing the mapped data, including
//     Year, Month, TotalSuccess, and TotalAmount.
func (t *transferStatsStatusResponseMapper) ToTransferResponseMonthStatusSuccess(s *record.TransferRecordMonthStatusSuccess) *response.TransferResponseMonthStatusSuccess {
	return &response.TransferResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTransferResponsesMonthStatusSuccess converts multiple monthly success status transfer records
// into a slice of TransferResponseMonthStatusSuccess.
//
// Args:
//   - Transfers: A slice of pointers to TransferRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferResponseMonthStatusSuccess containing the mapped data, including
//     Year, Month, TotalSuccess, and TotalAmount.
func (t *transferStatsStatusResponseMapper) ToTransferResponsesMonthStatusSuccess(Transfers []*record.TransferRecordMonthStatusSuccess) []*response.TransferResponseMonthStatusSuccess {
	var TransferRecords []*response.TransferResponseMonthStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferResponseMonthStatusSuccess(Transfer))
	}

	return TransferRecords
}

// ToTransferResponseYearStatusSuccess converts a yearly success status transfer record
// into a TransferResponseYearStatusSuccess.
//
// Args:
//   - s: A pointer to a TransferRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferResponseYearStatusSuccess containing the mapped data, including
//     Year, TotalSuccess, and TotalAmount.
func (t *transferStatsStatusResponseMapper) ToTransferResponseYearStatusSuccess(s *record.TransferRecordYearStatusSuccess) *response.TransferResponseYearStatusSuccess {
	return &response.TransferResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTransferResponsesYearStatusSuccess converts multiple yearly success status transfer records
// into a slice of TransferResponseYearStatusSuccess.
//
// Args:
//   - Transfers: A slice of pointers to TransferRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferResponseYearStatusSuccess containing the mapped data, including
//     Year, TotalSuccess, and TotalAmount.
func (t *transferStatsStatusResponseMapper) ToTransferResponsesYearStatusSuccess(Transfers []*record.TransferRecordYearStatusSuccess) []*response.TransferResponseYearStatusSuccess {
	var TransferRecords []*response.TransferResponseYearStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferResponseYearStatusSuccess(Transfer))
	}

	return TransferRecords
}

// ToTransferResponseMonthStatusFailed converts a monthly failed status transfer record
// into a TransferResponseMonthStatusFailed.
//
// Args:
//   - s: A pointer to a TransferRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferResponseMonthStatusFailed containing the mapped data, including
//     Year, Month, TotalFailed, and TotalAmount.
func (t *transferStatsStatusResponseMapper) ToTransferResponseMonthStatusFailed(s *record.TransferRecordMonthStatusFailed) *response.TransferResponseMonthStatusFailed {
	return &response.TransferResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTransferResponsesMonthStatusFailed converts multiple monthly failed status transfer records
// into a slice of TransferResponseMonthStatusFailed.
//
// Args:
//   - Transfers: A slice of pointers to TransferRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferResponseMonthStatusFailed containing the mapped data, including
//     Year, Month, TotalFailed, and TotalAmount.
func (t *transferStatsStatusResponseMapper) ToTransferResponsesMonthStatusFailed(Transfers []*record.TransferRecordMonthStatusFailed) []*response.TransferResponseMonthStatusFailed {
	var TransferRecords []*response.TransferResponseMonthStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferResponseMonthStatusFailed(Transfer))
	}

	return TransferRecords
}

// ToTransferResponseYearStatusFailed converts a yearly failed status transfer record
// into a TransferResponseYearStatusFailed.
//
// Args:
//   - s: A pointer to a TransferRecordYearStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *transferStatsStatusResponseMapper) ToTransferResponseYearStatusFailed(s *record.TransferRecordYearStatusFailed) *response.TransferResponseYearStatusFailed {
	return &response.TransferResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTransferResponsesYearStatusFailed converts multiple yearly failed status transfer records
// into a slice of TransferResponseYearStatusFailed.
//
// Args:
//   - Transfers: A slice of pointers to TransferRecordYearStatusFailed containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *transferStatsStatusResponseMapper) ToTransferResponsesYearStatusFailed(Transfers []*record.TransferRecordYearStatusFailed) []*response.TransferResponseYearStatusFailed {
	var TransferRecords []*response.TransferResponseYearStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferResponseYearStatusFailed(Transfer))
	}

	return TransferRecords
}
