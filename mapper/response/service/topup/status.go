package topupservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// TopupStatusResponseMapper is an interface that provides methods to map TopupRecord domain models to TopupResponseMonthStatusSuccess, TopupResponseMonthStatusFailed, TopupResponseYearStatusSuccess, and TopupResponseYearStatusFailed API-compatible response types.
type topupStatsStatusResponseMapper struct {
}

// NewTopupStatusResponseMapper returns a pointer to a new topupStatsStatusResponseMapper instance.
// This struct provides methods to map TopupRecord domain models to TopupResponseMonthStatusSuccess, TopupResponseMonthStatusFailed, TopupResponseYearStatusSuccess, and TopupResponseYearStatusFailed API-compatible response types.
func NewTopupStatsStatusResponseMapper() TopupStatsStatusResponseMapper {
	return &topupStatsStatusResponseMapper{}
}

// ToTopupResponseMonthStatusSuccess converts a TopupRecordMonthStatusSuccess domain model
// into a TopupResponseMonthStatusSuccess API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupResponseMonthStatusSuccess containing the mapped data, including
//     Year, Month, TotalSuccess, and TotalAmount.
func (t *topupStatsStatusResponseMapper) ToTopupResponseMonthStatusSuccess(s *record.TopupRecordMonthStatusSuccess) *response.TopupResponseMonthStatusSuccess {
	return &response.TopupResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTopupResponsesMonthStatusSuccess maps a slice of TopupRecordMonthStatusSuccess domain models
// to a slice of TopupResponseMonthStatusSuccess API-compatible response types.
//
// Args:
//   - topups: A slice of TopupRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponseMonthStatusSuccess containing the mapped data, including
//     Year, Month, TotalSuccess, and TotalAmount.
func (t *topupStatsStatusResponseMapper) ToTopupResponsesMonthStatusSuccess(topups []*record.TopupRecordMonthStatusSuccess) []*response.TopupResponseMonthStatusSuccess {
	var topupRecords []*response.TopupResponseMonthStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupResponseMonthStatusSuccess(topup))
	}

	return topupRecords
}

// ToTopupResponseYearStatusSuccess converts a TopupRecordYearStatusSuccess domain model
// into a TopupResponseYearStatusSuccess API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupResponseYearStatusSuccess containing the mapped data, including
//     Year, TotalSuccess, and TotalAmount.
func (t *topupStatsStatusResponseMapper) ToTopupResponseYearStatusSuccess(s *record.TopupRecordYearStatusSuccess) *response.TopupResponseYearStatusSuccess {
	return &response.TopupResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTopupResponsesYearStatusSuccess maps a slice of TopupRecordYearStatusSuccess domain models
// to a slice of TopupResponseYearStatusSuccess API-compatible response types.
//
// Args:
//   - topups: A slice of TopupRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponseYearStatusSuccess containing the mapped data, including
//     Year, TotalSuccess, and TotalAmount.
func (t *topupStatsStatusResponseMapper) ToTopupResponsesYearStatusSuccess(topups []*record.TopupRecordYearStatusSuccess) []*response.TopupResponseYearStatusSuccess {
	var topupRecords []*response.TopupResponseYearStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupResponseYearStatusSuccess(topup))
	}

	return topupRecords
}

// ToTopupResponseMonthStatusFailed converts a TopupRecordMonthStatusFailed domain model
// into a TopupResponseMonthStatusFailed API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupResponseMonthStatusFailed containing the mapped data, including
//     Year, Month, TotalFailed, and TotalAmount.
func (t *topupStatsStatusResponseMapper) ToTopupResponseMonthStatusFailed(s *record.TopupRecordMonthStatusFailed) *response.TopupResponseMonthStatusFailed {
	return &response.TopupResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTopupResponsesMonthStatusFailed maps a slice of TopupRecordMonthStatusFailed domain models
// to a slice of TopupResponseMonthStatusFailed API-compatible response types.
//
// Args:
//   - topups: A slice of TopupRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponseMonthStatusFailed containing the mapped data, including
//     Year, Month, TotalFailed, and TotalAmount.
func (t *topupStatsStatusResponseMapper) ToTopupResponsesMonthStatusFailed(topups []*record.TopupRecordMonthStatusFailed) []*response.TopupResponseMonthStatusFailed {
	var topupRecords []*response.TopupResponseMonthStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupResponseMonthStatusFailed(topup))
	}

	return topupRecords
}

// ToTopupResponseYearStatusFailed maps a TopupRecordYearStatusFailed domain model
// to a TopupResponseYearStatusFailed API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupRecordYearStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *topupStatsStatusResponseMapper) ToTopupResponseYearStatusFailed(s *record.TopupRecordYearStatusFailed) *response.TopupResponseYearStatusFailed {
	return &response.TopupResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTopupResponsesYearStatusFailed maps a slice of TopupRecordYearStatusFailed domain models
// to a slice of TopupResponseYearStatusFailed API-compatible response types.
//
// Args:
//   - topups: A slice of TopupRecordYearStatusFailed containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *topupStatsStatusResponseMapper) ToTopupResponsesYearStatusFailed(topups []*record.TopupRecordYearStatusFailed) []*response.TopupResponseYearStatusFailed {
	var topupRecords []*response.TopupResponseYearStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupResponseYearStatusFailed(topup))
	}

	return topupRecords
}
