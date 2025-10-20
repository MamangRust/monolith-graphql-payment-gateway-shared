package topupstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TopupStatisticStatusMapper maps topup status success/failed rows to domain models.
type TopupStatisticStatusMapper interface {
	// ToTopupRecordMonthStatusSuccess maps a GetMonthTopupStatusSuccessRow database row to a TopupRecordMonthStatusSuccess domain model.
	// It takes a pointer to a GetMonthTopupStatusSuccessRow as a parameter and returns a pointer to a TopupRecordMonthStatusSuccess.
	// The function maps the fields of the GetMonthTopupStatusSuccessRow to the corresponding fields of the TopupRecordMonthStatusSuccess.
	ToTopupRecordMonthStatusSuccess(s *db.GetMonthTopupStatusSuccessRow) *record.TopupRecordMonthStatusSuccess
	// ToTopupRecordsMonthStatusSuccess maps a slice of GetMonthTopupStatusSuccessRow database rows to a slice of TopupRecordMonthStatusSuccess domain models.
	// It iterates over the provided slice of GetMonthTopupStatusSuccessRow, converting each element
	// using the ToTopupRecordMonthStatusSuccess method and appending the result to a new slice.
	// The function returns a slice of pointers to TopupRecordMonthStatusSuccess.
	ToTopupRecordsMonthStatusSuccess(topups []*db.GetMonthTopupStatusSuccessRow) []*record.TopupRecordMonthStatusSuccess
	// ToTopupRecordYearStatusSuccess maps a GetYearlyTopupStatusSuccessRow database row to a TopupRecordYearStatusSuccess domain model.
	// It takes a pointer to a GetYearlyTopupStatusSuccessRow as a parameter and returns a pointer to a TopupRecordYearStatusSuccess.
	// The function maps the fields of the GetYearlyTopupStatusSuccessRow to the corresponding fields of the TopupRecordYearStatusSuccess.
	ToTopupRecordYearStatusSuccess(s *db.GetYearlyTopupStatusSuccessRow) *record.TopupRecordYearStatusSuccess
	// ToTopupRecordsYearStatusSuccess maps a slice of GetYearlyTopupStatusSuccessRow database rows to a slice of TopupRecordYearStatusSuccess domain models.
	// It iterates over the provided slice of GetYearlyTopupStatusSuccessRow, converting each element
	// using the ToTopupRecordYearStatusSuccess method and appending the result to a new slice.
	// The function returns a slice of pointers to TopupRecordYearStatusSuccess.
	ToTopupRecordsYearStatusSuccess(topups []*db.GetYearlyTopupStatusSuccessRow) []*record.TopupRecordYearStatusSuccess

	// ToTopupRecordMonthStatusFailed maps a GetMonthTopupStatusFailedRow database row
	// to a TopupRecordMonthStatusFailed domain model. It takes a pointer to a
	// GetMonthTopupStatusFailedRow as a parameter and returns a pointer to a
	// TopupRecordMonthStatusFailed. The function maps the fields of the
	// GetMonthTopupStatusFailedRow to the corresponding fields of the
	// TopupRecordMonthStatusFailed.
	ToTopupRecordMonthStatusFailed(s *db.GetMonthTopupStatusFailedRow) *record.TopupRecordMonthStatusFailed
	// ToTopupRecordsMonthStatusFailed maps a slice of GetMonthTopupStatusFailedRow database rows to a slice of TopupRecordMonthStatusFailed domain models.
	// It iterates over the provided slice of GetMonthTopupStatusFailedRow, converting each element
	// using the ToTopupRecordMonthStatusFailed method and appending the result to a new slice.
	// The function returns a slice of pointers to TopupRecordMonthStatusFailed.
	ToTopupRecordsMonthStatusFailed(topups []*db.GetMonthTopupStatusFailedRow) []*record.TopupRecordMonthStatusFailed
	// ToTopupRecordYearStatusFailed maps a GetYearlyTopupStatusFailedRow database row to a TopupRecordYearStatusFailed domain model.
	// It takes a pointer to a GetYearlyTopupStatusFailedRow as a parameter and returns a pointer to a TopupRecordYearStatusFailed.
	// The function maps the fields of the GetYearlyTopupStatusFailedRow to the corresponding fields of the TopupRecordYearStatusFailed.
	ToTopupRecordYearStatusFailed(s *db.GetYearlyTopupStatusFailedRow) *record.TopupRecordYearStatusFailed
	// ToTopupRecordsYearStatusFailed maps a slice of GetYearlyTopupStatusFailedRow database rows to a slice of TopupRecordYearStatusFailed domain models.
	// It iterates over the provided slice of GetYearlyTopupStatusFailedRow, converting each element
	// using the ToTopupRecordYearStatusFailed method and appending the result to a new slice.
	// The function returns a slice of pointers to TopupRecordYearStatusFailed.
	ToTopupRecordsYearStatusFailed(topups []*db.GetYearlyTopupStatusFailedRow) []*record.TopupRecordYearStatusFailed
}

// TopupStatisticMethodMapper maps topup method rows to domain models.
type TopupStatisticMethodMapper interface {
	// ToTopupMonthlyMethod maps a GetMonthlyTopupMethodsRow database row to a TopupMonthMethod domain model.
	// It takes a pointer to a GetMonthlyTopupMethodsRow as a parameter and returns a pointer to a TopupMonthMethod.
	// The function maps the fields of GetMonthlyTopupMethodsRow to the corresponding fields of TopupMonthMethod.
	ToTopupMonthlyMethod(s *db.GetMonthlyTopupMethodsRow) *record.TopupMonthMethod
	// ToTopupMonthlyMethods maps a slice of GetMonthlyTopupMethodsRow database rows to a slice of TopupMonthMethod domain models.
	// It iterates over the provided slice, converting each element using the ToTopupMonthlyMethod method,
	// and appends the result to a new slice. The function returns a slice of pointers to TopupMonthMethod.
	ToTopupMonthlyMethods(s []*db.GetMonthlyTopupMethodsRow) []*record.TopupMonthMethod
	// ToTopupYearlyMethod maps a GetYearlyTopupMethodsRow database row to a TopupYearlyMethod domain model.
	// It takes a pointer to a GetYearlyTopupMethodsRow as a parameter and returns a pointer to a TopupYearlyMethod.
	// The function maps the fields of GetYearlyTopupMethodsRow to the corresponding fields of TopupYearlyMethod.
	ToTopupYearlyMethod(s *db.GetYearlyTopupMethodsRow) *record.TopupYearlyMethod
	// ToTopupYearlyMethods maps a slice of GetYearlyTopupMethodsRow database rows to a slice of TopupYearlyMethod domain models.
	// It iterates over the provided slice, converting each element using the ToTopupYearlyMethod method,
	// and appends the result to a new slice. The function returns a slice of pointers to TopupYearlyMethod.
	ToTopupYearlyMethods(s []*db.GetYearlyTopupMethodsRow) []*record.TopupYearlyMethod
}

// TopupStatisticAmountMapper maps topup amount rows to domain models.
type TopupStatisticAmountMapper interface {
	// ToTopupMonthlyAmount maps a GetMonthlyTopupAmountsRow database row to a TopupMonthAmount domain model.
	// It takes a pointer to a GetMonthlyTopupAmountsRow as a parameter and returns a pointer to a TopupMonthAmount.
	// The function maps the fields of GetMonthlyTopupAmountsRow to the corresponding fields of TopupMonthAmount.
	ToTopupMonthlyAmount(s *db.GetMonthlyTopupAmountsRow) *record.TopupMonthAmount
	// ToTopupMonthlyAmounts maps a slice of GetMonthlyTopupAmountsRow database rows to a slice of TopupMonthAmount domain models.
	// It iterates over the provided slice, converting each element using the ToTopupMonthlyAmount method,
	// and appends the result to a new slice. The function returns a slice of pointers to TopupMonthAmount.
	ToTopupMonthlyAmounts(s []*db.GetMonthlyTopupAmountsRow) []*record.TopupMonthAmount
	// ToTopupYearlyAmount maps a GetYearlyTopupAmountsRow database row to a TopupYearlyAmount domain model.
	// It takes a pointer to a GetYearlyTopupAmountsRow as a parameter and returns a pointer to a TopupYearlyAmount.
	// The function maps the fields of GetYearlyTopupAmountsRow to the corresponding fields of TopupYearlyAmount.
	ToTopupYearlyAmount(s *db.GetYearlyTopupAmountsRow) *record.TopupYearlyAmount
	// ToTopupYearlyAmounts maps a slice of GetYearlyTopupAmountsRow database rows to a slice of TopupYearlyAmount domain models.
	// It iterates over the provided slice, converting each element using the ToTopupYearlyAmount method,
	// and appends the result to a new slice. The function returns a slice of pointers to TopupYearlyAmount.
	ToTopupYearlyAmounts(s []*db.GetYearlyTopupAmountsRow) []*record.TopupYearlyAmount
}
