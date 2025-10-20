package topupstatsbycardrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// TopupStatisticStatusByCardNumberMapper is an interface that provides methods to map Topup database rows to TopupRecord domain models for statistic operations.
type TopupStatisticStatusByCardNumberMapper interface {
	// ToTopupRecordMonthStatusSuccessByCardNumber maps a GetMonthTopupStatusSuccessCardNumberRow database row to a TopupRecordMonthStatusSuccess domain model.
	// It takes a pointer to a GetMonthTopupStatusSuccessCardNumberRow as a parameter and returns a pointer to a TopupRecordMonthStatusSuccess.
	// The function maps the fields of the GetMonthTopupStatusSuccessCardNumberRow to the corresponding fields of the TopupRecordMonthStatusSuccess.
	ToTopupRecordMonthStatusSuccessByCardNumber(s *db.GetMonthTopupStatusSuccessCardNumberRow) *record.TopupRecordMonthStatusSuccess
	// ToTopupRecordsMonthStatusSuccessByCardNumber maps a slice of GetMonthTopupStatusSuccessCardNumberRow database rows to a slice of TopupRecordMonthStatusSuccess domain models.
	// It iterates over the provided slice of GetMonthTopupStatusSuccessCardNumberRow, converting each element
	// using the ToTopupRecordMonthStatusSuccessByCardNumber method and appending the result to a new slice.
	// The function returns a slice of pointers to TopupRecordMonthStatusSuccess.
	ToTopupRecordsMonthStatusSuccessByCardNumber(s []*db.GetMonthTopupStatusSuccessCardNumberRow) []*record.TopupRecordMonthStatusSuccess
	// ToTopupRecordYearStatusSuccessByCardNumber maps a GetYearlyTopupStatusSuccessCardNumberRow database row
	// to a TopupRecordYearStatusSuccess domain model. It takes a pointer to a
	// GetYearlyTopupStatusSuccessCardNumberRow as a parameter and returns a pointer
	// to a TopupRecordYearStatusSuccess. The function maps the fields of
	// GetYearlyTopupStatusSuccessCardNumberRow to the corresponding fields of
	// TopupRecordYearStatusSuccess.
	ToTopupRecordYearStatusSuccessByCardNumber(s *db.GetYearlyTopupStatusSuccessCardNumberRow) *record.TopupRecordYearStatusSuccess
	// ToTopupRecordsYearStatusSuccessByCardNumber maps a slice of GetYearlyTopupStatusSuccessCardNumberRow database rows to a slice of TopupRecordYearStatusSuccess domain models.
	// It iterates over the provided slice of GetYearlyTopupStatusSuccessCardNumberRow, converting each element
	// using the ToTopupRecordYearStatusSuccessByCardNumber method and appending the result to a new slice.
	// The function returns a slice of pointers to TopupRecordYearStatusSuccess.
	ToTopupRecordsYearStatusSuccessByCardNumber(s []*db.GetYearlyTopupStatusSuccessCardNumberRow) []*record.TopupRecordYearStatusSuccess

	// ToTopupRecordMonthStatusFailedByCardNumber maps a GetMonthTopupStatusFailedCardNumberRow database row
	// to a TopupRecordMonthStatusFailed domain model. It takes a pointer to a
	// GetMonthTopupStatusFailedCardNumberRow as a parameter and returns a pointer to a
	// TopupRecordMonthStatusFailed. The function maps the fields of the
	// GetMonthTopupStatusFailedCardNumberRow to the corresponding fields of the
	// TopupRecordMonthStatusFailed.
	ToTopupRecordMonthStatusFailedByCardNumber(s *db.GetMonthTopupStatusFailedCardNumberRow) *record.TopupRecordMonthStatusFailed
	// ToTopupRecordsMonthStatusFailedByCardNumber maps a slice of GetMonthTopupStatusFailedCardNumberRow database rows to a slice of TopupRecordMonthStatusFailed domain models.
	// It iterates over the provided slice of GetMonthTopupStatusFailedCardNumberRow, converting each element
	// using the ToTopupRecordMonthStatusFailedByCardNumber method and appending the result to a new slice.
	// The function returns a slice of pointers to TopupRecordMonthStatusFailed.
	ToTopupRecordsMonthStatusFailedByCardNumber(s []*db.GetMonthTopupStatusFailedCardNumberRow) []*record.TopupRecordMonthStatusFailed
	// ToTopupRecordYearStatusFailedByCardNumber maps a GetYearlyTopupStatusFailedCardNumberRow database row
	// to a TopupRecordYearStatusFailed domain model. It takes a pointer to a
	// GetYearlyTopupStatusFailedCardNumberRow as a parameter and returns a pointer
	// to a TopupRecordYearStatusFailed. The function maps the fields of
	// GetYearlyTopupStatusFailedCardNumberRow to the cor
	ToTopupRecordYearStatusFailedByCardNumber(s *db.GetYearlyTopupStatusFailedCardNumberRow) *record.TopupRecordYearStatusFailed
	// ToTopupRecordsYearStatusFailedByCardNumber maps a slice of GetYearlyTopupStatusFailedCardNumberRow database rows
	// to a slice of TopupRecordYearStatusFailed domain models. It iterates over the provided slice, converting each element
	// using the ToTopupRecordYearStatusFailedByCardNumber method, and appends the result to a new slice.
	// The function returns a slice of pointers to TopupRecordYearStatusFailed.
	ToTopupRecordsYearStatusFailedByCardNumber(s []*db.GetYearlyTopupStatusFailedCardNumberRow) []*record.TopupRecordYearStatusFailed
}

// TopupStatisticMethodByCardNumberMapper is an interface that provides methods to map Topup database rows to TopupRecord domain models for statistic operations.
type TopupStatisticMethodByCardNumberMapper interface {
	ToTopupMonthlyMethodByCardNumber(s *db.GetMonthlyTopupMethodsByCardNumberRow) *record.TopupMonthMethod
	ToTopupMonthlyMethodsByCardNumber(s []*db.GetMonthlyTopupMethodsByCardNumberRow) []*record.TopupMonthMethod
	ToTopupYearlyMethodByCardNumber(s *db.GetYearlyTopupMethodsByCardNumberRow) *record.TopupYearlyMethod
	ToTopupYearlyMethodsByCardNumber(s []*db.GetYearlyTopupMethodsByCardNumberRow) []*record.TopupYearlyMethod
}

// TopupStatisticAmountByCardNumberMapper is an interface that provides methods to map Topup database rows to TopupRecord domain models for statistic operations.
type TopupStatisticAmountByCardNumberMapper interface {
	// ToTopupMonthlyAmountByCardNumber maps a GetMonthlyTopupAmountsByCardNumberRow to a TopupMonthAmount domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetMonthlyTopupAmountsByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TopupMonthAmount containing the mapped data, including Month and TotalAmount.
	ToTopupMonthlyAmountByCardNumber(s *db.GetMonthlyTopupAmountsByCardNumberRow) *record.TopupMonthAmount
	// ToTopupMonthlyAmountsByCardNumber maps a slice of GetMonthlyTopupAmountsByCardNumberRow database rows to a slice of TopupMonthAmount domain models.
	// It iterates over the provided slice, converting each element using the ToTopupMonthlyAmountByCardNumber method,
	// and appends the result to a new slice. The function returns a slice of pointers to TopupMonthAmount.
	ToTopupMonthlyAmountsByCardNumber(s []*db.GetMonthlyTopupAmountsByCardNumberRow) []*record.TopupMonthAmount
	// ToTopupYearlyAmountByCardNumber maps a GetYearlyTopupAmountsByCardNumberRow to a TopupYearlyAmount domain model.
	//
	// Parameters:
	//   - s: A pointer to a GetYearlyTopupAmountsByCardNumberRow representing the database row.
	//
	// Returns:
	//   - A pointer to a TopupYearlyAmount containing the mapped data, including Year and TotalAmount.
	ToTopupYearlyAmountByCardNumber(s *db.GetYearlyTopupAmountsByCardNumberRow) *record.TopupYearlyAmount
	// ToTopupYearlyAmountsByCardNumber maps a slice of GetYearlyTopupAmountsByCardNumberRow to a
	// slice of TopupYearlyAmount domain models.
	//
	// Parameters:
	//   - s: A slice of pointers to GetYearlyTopupAmountsByCardNumberRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to TopupYearlyAmount containing the mapped data, including Year and TotalAmount.
	ToTopupYearlyAmountsByCardNumber(s []*db.GetYearlyTopupAmountsByCardNumberRow) []*record.TopupYearlyAmount
}
