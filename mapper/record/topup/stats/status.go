package topupstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

type topupStatisticStatusRecordMapper struct {
}

func NewTopupStatisticStatusRecordMapper() TopupStatisticStatusMapper {
	return &topupStatisticStatusRecordMapper{}
}

// ToTopupRecordMonthStatusSuccess maps a GetMonthTopupStatusSuccessRow database row to a TopupRecordMonthStatusSuccess domain model.
// It takes a pointer to a GetMonthTopupStatusSuccessRow as a parameter and returns a pointer to a TopupRecordMonthStatusSuccess.
// The function maps the fields of the GetMonthTopupStatusSuccessRow to the corresponding fields of the TopupRecordMonthStatusSuccess.
func (t *topupStatisticStatusRecordMapper) ToTopupRecordMonthStatusSuccess(s *db.GetMonthTopupStatusSuccessRow) *record.TopupRecordMonthStatusSuccess {
	return &record.TopupRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTopupRecordsMonthStatusSuccess maps a slice of GetMonthTopupStatusSuccessRow database rows to a slice of TopupRecordMonthStatusSuccess domain models.
// It iterates over the provided slice of GetMonthTopupStatusSuccessRow, converting each element
// using the ToTopupRecordMonthStatusSuccess method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordMonthStatusSuccess.
func (t *topupStatisticStatusRecordMapper) ToTopupRecordsMonthStatusSuccess(topups []*db.GetMonthTopupStatusSuccessRow) []*record.TopupRecordMonthStatusSuccess {
	var topupRecords []*record.TopupRecordMonthStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordMonthStatusSuccess(topup))
	}

	return topupRecords
}

// ToTopupRecordYearStatusSuccess maps a GetYearlyTopupStatusSuccessRow database row to a TopupRecordYearStatusSuccess domain model.
// It takes a pointer to a GetYearlyTopupStatusSuccessRow as a parameter and returns a pointer to a TopupRecordYearStatusSuccess.
// The function maps the fields of the GetYearlyTopupStatusSuccessRow to the corresponding fields of the TopupRecordYearStatusSuccess.
func (t *topupStatisticStatusRecordMapper) ToTopupRecordYearStatusSuccess(s *db.GetYearlyTopupStatusSuccessRow) *record.TopupRecordYearStatusSuccess {
	return &record.TopupRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTopupRecordsYearStatusSuccess maps a slice of GetYearlyTopupStatusSuccessRow database rows to a slice of TopupRecordYearStatusSuccess domain models.
// It iterates over the provided slice of GetYearlyTopupStatusSuccessRow, converting each element
// using the ToTopupRecordYearStatusSuccess method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordYearStatusSuccess.
func (t *topupStatisticStatusRecordMapper) ToTopupRecordsYearStatusSuccess(topups []*db.GetYearlyTopupStatusSuccessRow) []*record.TopupRecordYearStatusSuccess {
	var topupRecords []*record.TopupRecordYearStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordYearStatusSuccess(topup))
	}

	return topupRecords
}

// ToTopupRecordMonthStatusFailed maps a GetMonthTopupStatusFailedRow database row
// to a TopupRecordMonthStatusFailed domain model. It takes a pointer to a
// GetMonthTopupStatusFailedRow as a parameter and returns a pointer to a
// TopupRecordMonthStatusFailed. The function maps the fields of the
// GetMonthTopupStatusFailedRow to the corresponding fields of the
// TopupRecordMonthStatusFailed.
func (t *topupStatisticStatusRecordMapper) ToTopupRecordMonthStatusFailed(s *db.GetMonthTopupStatusFailedRow) *record.TopupRecordMonthStatusFailed {
	return &record.TopupRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupRecordsMonthStatusFailed maps a slice of GetMonthTopupStatusFailedRow database rows to a slice of TopupRecordMonthStatusFailed domain models.
// It iterates over the provided slice of GetMonthTopupStatusFailedRow, converting each element
// using the ToTopupRecordMonthStatusFailed method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordMonthStatusFailed.
func (t *topupStatisticStatusRecordMapper) ToTopupRecordsMonthStatusFailed(topups []*db.GetMonthTopupStatusFailedRow) []*record.TopupRecordMonthStatusFailed {
	var topupRecords []*record.TopupRecordMonthStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordMonthStatusFailed(topup))
	}

	return topupRecords
}

// ToTopupRecordYearStatusFailed maps a GetYearlyTopupStatusFailedRow database row to a TopupRecordYearStatusFailed domain model.
// It takes a pointer to a GetYearlyTopupStatusFailedRow as a parameter and returns a pointer to a TopupRecordYearStatusFailed.
// The function maps the fields of the GetYearlyTopupStatusFailedRow to the corresponding fields of the TopupRecordYearStatusFailed.
func (t *topupStatisticStatusRecordMapper) ToTopupRecordYearStatusFailed(s *db.GetYearlyTopupStatusFailedRow) *record.TopupRecordYearStatusFailed {
	return &record.TopupRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupRecordsYearStatusFailed maps a slice of GetYearlyTopupStatusFailedRow database rows to a slice of TopupRecordYearStatusFailed domain models.
// It iterates over the provided slice of GetYearlyTopupStatusFailedRow, converting each element
// using the ToTopupRecordYearStatusFailed method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordYearStatusFailed.
func (t *topupStatisticStatusRecordMapper) ToTopupRecordsYearStatusFailed(topups []*db.GetYearlyTopupStatusFailedRow) []*record.TopupRecordYearStatusFailed {
	var topupRecords []*record.TopupRecordYearStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordYearStatusFailed(topup))
	}

	return topupRecords
}
