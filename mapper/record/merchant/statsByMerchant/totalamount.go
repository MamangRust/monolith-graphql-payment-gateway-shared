package merchantstatbyidrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// merchantStatisticTotalAmountByMerchantMapper provides methods to map Merchant database rows to
type merchantStatisticTotalAmountByMerchantMapper struct {
}

// NewMerchantStatisticTotalAmountByMerchantRecordMapper returns a new instance of
// MerchantStatisticTotalAmountByMerchantMapper, which provides methods to map
// Merchant database rows to MerchantRecord domain models for statistic operations.
func NewMerchantStatisticTotalAmountByMerchantRecordMapper() MerchantStatisticTotalAmountByMerchantMapper {
	return &merchantStatisticTotalAmountByMerchantMapper{}
}

// ToMerchantMonthlyTotalAmountByMerchant maps a GetMonthlyTotalAmountByMerchantRow to a
// MerchantMonthlyTotalAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetMonthlyTotalAmountByMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (m *merchantStatisticTotalAmountByMerchantMapper) ToMerchantMonthlyTotalAmountByMerchant(ms *db.GetMonthlyTotalAmountByMerchantRow) *record.MerchantMonthlyTotalAmount {
	return &record.MerchantMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyTotalAmountsByMerchant maps a slice of GetMonthlyTotalAmountByMerchantRow to a
// slice of MerchantMonthlyTotalAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetMonthlyTotalAmountByMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (m *merchantStatisticTotalAmountByMerchantMapper) ToMerchantMonthlyTotalAmountsByMerchant(ms []*db.GetMonthlyTotalAmountByMerchantRow) []*record.MerchantMonthlyTotalAmount {
	var records []*record.MerchantMonthlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyTotalAmountByMerchant(merchant))
	}
	return records
}

// ToMerchantYearlyTotalAmountByMerchant maps a GetYearlyTotalAmountByMerchantRow to a
// MerchantYearlyTotalAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetYearlyTotalAmountByMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantStatisticTotalAmountByMerchantMapper) ToMerchantYearlyTotalAmountByMerchant(ms *db.GetYearlyTotalAmountByMerchantRow) *record.MerchantYearlyTotalAmount {
	return &record.MerchantYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyTotalAmountsByMerchant maps a slice of GetYearlyTotalAmountByMerchantRow to a
// slice of MerchantYearlyTotalAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetYearlyTotalAmountByMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantStatisticTotalAmountByMerchantMapper) ToMerchantYearlyTotalAmountsByMerchant(ms []*db.GetYearlyTotalAmountByMerchantRow) []*record.MerchantYearlyTotalAmount {
	var records []*record.MerchantYearlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyTotalAmountByMerchant(merchant))
	}
	return records
}
