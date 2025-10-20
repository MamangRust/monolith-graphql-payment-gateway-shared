package topupstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

type topupStatisticStatusByCardNumberMapper struct{}

func NewTopupStatisticStatusByCardNumberMapper() TopupStatisticStatusByCardNumberMapper {
	return &topupStatisticStatusByCardNumberMapper{}
}

// ToTopupRecordMonthStatusSuccessByCardNumber maps a GetMonthTopupStatusSuccessCardNumberRow database row to a TopupRecordMonthStatusSuccess domain model.
// It takes a pointer to a GetMonthTopupStatusSuccessCardNumberRow as a parameter and returns a pointer to a TopupRecordMonthStatusSuccess.
// The function maps the fields of the GetMonthTopupStatusSuccessCardNumberRow to the corresponding fields of the TopupRecordMonthStatusSuccess.
func (t *topupStatisticStatusByCardNumberMapper) ToTopupRecordMonthStatusSuccessByCardNumber(s *db.GetMonthTopupStatusSuccessCardNumberRow) *record.TopupRecordMonthStatusSuccess {
	return &record.TopupRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTopupRecordsMonthStatusSuccessByCardNumber maps a slice of GetMonthTopupStatusSuccessCardNumberRow database rows to a slice of TopupRecordMonthStatusSuccess domain models.
// It iterates over the provided slice of GetMonthTopupStatusSuccessCardNumberRow, converting each element
// using the ToTopupRecordMonthStatusSuccessByCardNumber method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordMonthStatusSuccess.
func (t *topupStatisticStatusByCardNumberMapper) ToTopupRecordsMonthStatusSuccessByCardNumber(topups []*db.GetMonthTopupStatusSuccessCardNumberRow) []*record.TopupRecordMonthStatusSuccess {
	var topupRecords []*record.TopupRecordMonthStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordMonthStatusSuccessByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupRecordYearStatusSuccessByCardNumber maps a GetYearlyTopupStatusSuccessCardNumberRow database row
// to a TopupRecordYearStatusSuccess domain model. It takes a pointer to a
// GetYearlyTopupStatusSuccessCardNumberRow as a parameter and returns a pointer
// to a TopupRecordYearStatusSuccess. The function maps the fields of
// GetYearlyTopupStatusSuccessCardNumberRow to the corresponding fields of
// TopupRecordYearStatusSuccess.
func (t *topupStatisticStatusByCardNumberMapper) ToTopupRecordYearStatusSuccessByCardNumber(s *db.GetYearlyTopupStatusSuccessCardNumberRow) *record.TopupRecordYearStatusSuccess {
	return &record.TopupRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTopupRecordsYearStatusSuccessByCardNumber maps a slice of GetYearlyTopupStatusSuccessCardNumberRow database rows to a slice of TopupRecordYearStatusSuccess domain models.
// It iterates over the provided slice of GetYearlyTopupStatusSuccessCardNumberRow, converting each element
// using the ToTopupRecordYearStatusSuccessByCardNumber method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordYearStatusSuccess.
func (t *topupStatisticStatusByCardNumberMapper) ToTopupRecordsYearStatusSuccessByCardNumber(topups []*db.GetYearlyTopupStatusSuccessCardNumberRow) []*record.TopupRecordYearStatusSuccess {
	var topupRecords []*record.TopupRecordYearStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordYearStatusSuccessByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupRecordMonthStatusFailedByCardNumber maps a GetMonthTopupStatusFailedCardNumberRow database row
// to a TopupRecordMonthStatusFailed domain model. It takes a pointer to a
// GetMonthTopupStatusFailedCardNumberRow as a parameter and returns a pointer to a
// TopupRecordMonthStatusFailed. The function maps the fields of the
// GetMonthTopupStatusFailedCardNumberRow to the corresponding fields of the
// TopupRecordMonthStatusFailed.
func (t *topupStatisticStatusByCardNumberMapper) ToTopupRecordMonthStatusFailedByCardNumber(s *db.GetMonthTopupStatusFailedCardNumberRow) *record.TopupRecordMonthStatusFailed {
	return &record.TopupRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupRecordsMonthStatusFailedByCardNumber maps a slice of GetMonthTopupStatusFailedCardNumberRow database rows to a slice of TopupRecordMonthStatusFailed domain models.
// It iterates over the provided slice of GetMonthTopupStatusFailedCardNumberRow, converting each element
// using the ToTopupRecordMonthStatusFailedByCardNumber method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordMonthStatusFailed.
func (t *topupStatisticStatusByCardNumberMapper) ToTopupRecordsMonthStatusFailedByCardNumber(topups []*db.GetMonthTopupStatusFailedCardNumberRow) []*record.TopupRecordMonthStatusFailed {
	var topupRecords []*record.TopupRecordMonthStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordMonthStatusFailedByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupRecordYearStatusFailedByCardNumber maps a GetYearlyTopupStatusFailedCardNumberRow database row
// to a TopupRecordYearStatusFailed domain model. It takes a pointer to a
// GetYearlyTopupStatusFailedCardNumberRow as a parameter and returns a pointer
// to a TopupRecordYearStatusFailed. The function maps the fields of
// GetYearlyTopupStatusFailedCardNumberRow to the corresponding fields of
// TopupRecordYearStatusFailed.
func (t *topupStatisticStatusByCardNumberMapper) ToTopupRecordYearStatusFailedByCardNumber(s *db.GetYearlyTopupStatusFailedCardNumberRow) *record.TopupRecordYearStatusFailed {
	return &record.TopupRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupRecordsYearStatusFailedByCardNumber maps a slice of GetYearlyTopupStatusFailedCardNumberRow database rows
// to a slice of TopupRecordYearStatusFailed domain models. It iterates over the provided slice, converting each element
// using the ToTopupRecordYearStatusFailedByCardNumber method, and appends the result to a new slice.
// The function returns a slice of pointers to TopupRecordYearStatusFailed.
func (t *topupStatisticStatusByCardNumberMapper) ToTopupRecordsYearStatusFailedByCardNumber(topups []*db.GetYearlyTopupStatusFailedCardNumberRow) []*record.TopupRecordYearStatusFailed {
	var topupRecords []*record.TopupRecordYearStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordYearStatusFailedByCardNumber(topup))
	}

	return topupRecords
}
