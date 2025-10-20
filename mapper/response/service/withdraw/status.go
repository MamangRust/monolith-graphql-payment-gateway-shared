package withdrawresponseservice

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// WithdrawStatsStatusResponseMapper provides methods to map WithdrawRecord domain models related to withdraw statistics (success/failed) into API-compatible WithdrawResponse objects.
type withdrawStatsStatusResponseMapper struct {
}

// NewWithdrawStatsStatusResponseMapper creates and returns a new instance of
// withdrawStatsStatusResponseMapper, which provides methods to map WithdrawRecord
// domain models related to withdraw statistics (success/failed) into API-compatible
// WithdrawResponse objects.
func NewWithdrawStatsStatusResponseMapper() WithdrawStatsStatusResponseMapper {
	return &withdrawStatsStatusResponseMapper{}
}

// ToWithdrawResponseMonthStatusSuccess converts a single WithdrawRecordMonthStatusSuccess into a WithdrawResponseMonthStatusSuccess.
//
// Args:
//   - s: The WithdrawRecordMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseMonthStatusSuccess containing the mapped data, including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) ToWithdrawResponseMonthStatusSuccess(s *record.WithdrawRecordMonthStatusSuccess) *response.WithdrawResponseMonthStatusSuccess {
	return &response.WithdrawResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToWithdrawResponsesMonthStatusSuccess converts multiple WithdrawRecordMonthStatusSuccess into a slice of WithdrawResponseMonthStatusSuccess.
//
// Args:
//   - Withdraws: The WithdrawRecordMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A slice of WithdrawResponseMonthStatusSuccess containing the mapped data, including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) ToWithdrawResponsesMonthStatusSuccess(Withdraws []*record.WithdrawRecordMonthStatusSuccess) []*response.WithdrawResponseMonthStatusSuccess {
	var WithdrawRecords []*response.WithdrawResponseMonthStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawResponseMonthStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawResponseYearStatusSuccess converts a single WithdrawRecordYearStatusSuccess into a WithdrawResponseYearStatusSuccess.
//
// Args:
//   - s: The WithdrawRecordYearStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseYearStatusSuccess containing the mapped data, including Year, TotalSuccess, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) ToWithdrawResponseYearStatusSuccess(s *record.WithdrawRecordYearStatusSuccess) *response.WithdrawResponseYearStatusSuccess {
	return &response.WithdrawResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToWithdrawResponsesYearStatusSuccess converts multiple yearly success-status WithdrawRecords
// into a slice of WithdrawResponseYearStatusSuccess.
//
// Args:
//   - Withdraws: A slice of pointers to WithdrawRecordYearStatusSuccess that need to be converted.
//
// Returns:
//   - A slice of pointers to WithdrawResponseYearStatusSuccess, each containing the mapped data
//     including Year, TotalSuccess, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) ToWithdrawResponsesYearStatusSuccess(Withdraws []*record.WithdrawRecordYearStatusSuccess) []*response.WithdrawResponseYearStatusSuccess {
	var WithdrawRecords []*response.WithdrawResponseYearStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawResponseYearStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawResponseMonthStatusFailed converts a single WithdrawRecordMonthStatusFailed into a WithdrawResponseMonthStatusFailed.
//
// Args:
//   - s: The WithdrawRecordMonthStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseMonthStatusFailed containing the mapped data, including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) ToWithdrawResponseMonthStatusFailed(s *record.WithdrawRecordMonthStatusFailed) *response.WithdrawResponseMonthStatusFailed {
	return &response.WithdrawResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToWithdrawResponsesMonthStatusFailed converts multiple WithdrawRecordMonthStatusFailed into a slice of WithdrawResponseMonthStatusFailed.
//
// Args:
//   - Withdraws: A slice of pointers to WithdrawRecordMonthStatusFailed that need to be converted.
//
// Returns:
//   - A slice of pointers to WithdrawResponseMonthStatusFailed, each containing the mapped data
//     including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) ToWithdrawResponsesMonthStatusFailed(Withdraws []*record.WithdrawRecordMonthStatusFailed) []*response.WithdrawResponseMonthStatusFailed {
	var WithdrawRecords []*response.WithdrawResponseMonthStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawResponseMonthStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawResponseYearStatusFailed converts a single WithdrawRecordYearStatusFailed into
// a WithdrawResponseYearStatusFailed.
//
// Args:
//   - s: The WithdrawRecordYearStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) ToWithdrawResponseYearStatusFailed(s *record.WithdrawRecordYearStatusFailed) *response.WithdrawResponseYearStatusFailed {
	return &response.WithdrawResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToWithdrawResponsesYearStatusFailed converts multiple yearly failed-status WithdrawRecords
// into a slice of WithdrawResponseYearStatusFailed.
//
// Args:
//   - Withdraws: A slice of pointers to WithdrawRecordYearStatusFailed that need to be converted.
//
// Returns:
//   - A slice of pointers to WithdrawResponseYearStatusFailed, each containing the mapped data
//     including Year, TotalFailed, and TotalAmount.
func (t *withdrawStatsStatusResponseMapper) ToWithdrawResponsesYearStatusFailed(Withdraws []*record.WithdrawRecordYearStatusFailed) []*response.WithdrawResponseYearStatusFailed {
	var WithdrawRecords []*response.WithdrawResponseYearStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawResponseYearStatusFailed(Withdraw))
	}

	return WithdrawRecords
}
