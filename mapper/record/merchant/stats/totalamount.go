package merchantstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantStatisticTotalAmountRecordMapper maps total amount statistics to domain models.
type merchantStatisticTotalAmountRecordMapper struct {
}

// NewMerchantStatisticTotalAmountRecordMapper returns a new instance of
// MerchantStatisticTotalAmountRecordMapper, which provides methods to map
// Merchant database rows to MerchantRecord domain models for statistic operations.
func NewMerchantStatisticTotalAmountRecordMapper() MerchantStatisticTotalAmountRecordMapper {
	return &merchantStatisticTotalAmountRecordMapper{}
}

// ToMerchantMonthlyTotalAmount maps a GetMonthlyTotalAmountMerchantRow to a
// MerchantMonthlyTotalAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetMonthlyTotalAmountMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (m *merchantStatisticTotalAmountRecordMapper) ToMerchantMonthlyTotalAmount(ms *db.GetMonthlyTotalAmountMerchantRow) *record.MerchantMonthlyTotalAmount {
	return &record.MerchantMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyTotalAmounts maps a slice of GetMonthlyTotalAmountMerchantRow to a slice
// of MerchantMonthlyTotalAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetMonthlyTotalAmountMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (m *merchantStatisticTotalAmountRecordMapper) ToMerchantMonthlyTotalAmounts(ms []*db.GetMonthlyTotalAmountMerchantRow) []*record.MerchantMonthlyTotalAmount {
	var records []*record.MerchantMonthlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyTotalAmount(merchant))
	}
	return records
}

// ToMerchantYearlyTotalAmount maps a GetYearlyTotalAmountMerchantRow to a
// MerchantYearlyTotalAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetYearlyTotalAmountMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantStatisticTotalAmountRecordMapper) ToMerchantYearlyTotalAmount(ms *db.GetYearlyTotalAmountMerchantRow) *record.MerchantYearlyTotalAmount {
	return &record.MerchantYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyTotalAmounts maps a slice of GetYearlyTotalAmountMerchantRow to a slice
// of MerchantYearlyTotalAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetYearlyTotalAmountMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantStatisticTotalAmountRecordMapper) ToMerchantYearlyTotalAmounts(ms []*db.GetYearlyTotalAmountMerchantRow) []*record.MerchantYearlyTotalAmount {
	var records []*record.MerchantYearlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyTotalAmount(merchant))
	}
	return records
}
