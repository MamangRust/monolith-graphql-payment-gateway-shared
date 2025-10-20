package merchantstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantStatisticAmountRecordMapper is an interface that provides methods to map Merchant database rows to MerchantRecord domain models for statistic operations.
type merchantStatisticAmountRecordMapper struct{}

// NewMerchantStatisticAmountRecordMapper returns a new instance of
// MerchantStatisticAmountRecordMapper, which provides methods to map
// Merchant database rows to MerchantRecord domain models for statistic operations.
func NewMerchantStatisticAmountRecordMapper() MerchantStatisticAmountRecordMapper {
	return &merchantStatisticAmountRecordMapper{}
}

// ToMerchantMonthlyAmount maps a GetMonthlyAmountMerchantRow to a MerchantMonthlyAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetMonthlyAmountMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.
func (m *merchantStatisticAmountRecordMapper) ToMerchantMonthlyAmount(ms *db.GetMonthlyAmountMerchantRow) *record.MerchantMonthlyAmount {
	return &record.MerchantMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyAmounts maps a slice of GetMonthlyAmountMerchantRow to a slice of MerchantMonthlyAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetMonthlyAmountMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.
func (m *merchantStatisticAmountRecordMapper) ToMerchantMonthlyAmounts(ms []*db.GetMonthlyAmountMerchantRow) []*record.MerchantMonthlyAmount {
	var records []*record.MerchantMonthlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyAmount(merchant))
	}
	return records
}

// ToMerchantYearlyAmount maps a GetYearlyAmountMerchantRow to a MerchantYearlyAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetYearlyAmountMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyAmount containing the mapped data, including Year and TotalAmount.
func (m *merchantStatisticAmountRecordMapper) ToMerchantYearlyAmount(ms *db.GetYearlyAmountMerchantRow) *record.MerchantYearlyAmount {
	return &record.MerchantYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyAmounts maps a slice of GetYearlyAmountMerchantRow to a slice
// of MerchantYearlyAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetYearlyAmountMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantStatisticAmountRecordMapper) ToMerchantYearlyAmounts(ms []*db.GetYearlyAmountMerchantRow) []*record.MerchantYearlyAmount {
	var records []*record.MerchantYearlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyAmount(merchant))
	}
	return records
}
