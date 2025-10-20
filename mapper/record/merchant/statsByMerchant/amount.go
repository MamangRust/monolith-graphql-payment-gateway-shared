package merchantstatbyidrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantStatisticAmountByMerchantMapper is an interface that provides methods to map Merchant database rows to MerchantRecord domain models for statistic operations.
type merchantStatisticAmountByMerchantMapper struct {
}

// NewMerchantStatisticAmountByMerchantRecordMapper returns a new instance of
// MerchantStatisticAmountByMerchantMapper, which provides methods to map
// Merchant database rows to MerchantRecord domain models for statistic operations.
func NewMerchantStatisticAmountByMerchantRecordMapper() MerchantStatisticAmountByMerchantMapper {
	return &merchantStatisticAmountByMerchantMapper{}
}

// ToMerchantMonthlyAmountByMerchant maps a GetMonthlyAmountByMerchantsRow to a
// MerchantMonthlyAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetMonthlyAmountByMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.
func (m *merchantStatisticAmountByMerchantMapper) ToMerchantMonthlyAmountByMerchant(ms *db.GetMonthlyAmountByMerchantsRow) *record.MerchantMonthlyAmount {
	return &record.MerchantMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyAmountsByMerchant maps a slice of GetMonthlyAmountByMerchantsRow
// to a slice of MerchantMonthlyAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetMonthlyAmountByMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including
//     Month and TotalAmount.
func (m *merchantStatisticAmountByMerchantMapper) ToMerchantMonthlyAmountsByMerchant(ms []*db.GetMonthlyAmountByMerchantsRow) []*record.MerchantMonthlyAmount {
	var records []*record.MerchantMonthlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyAmountByMerchant(merchant))
	}
	return records
}

// ToMerchantYearlyAmountByMerchant maps a GetYearlyAmountByMerchantsRow to a
// MerchantYearlyAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetYearlyAmountByMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantStatisticAmountByMerchantMapper) ToMerchantYearlyAmountByMerchant(ms *db.GetYearlyAmountByMerchantsRow) *record.MerchantYearlyAmount {
	return &record.MerchantYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyAmountsByMerchant maps a slice of GetYearlyAmountByMerchantsRow
// to a slice of MerchantYearlyAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetYearlyAmountByMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantStatisticAmountByMerchantMapper) ToMerchantYearlyAmountsByMerchant(ms []*db.GetYearlyAmountByMerchantsRow) []*record.MerchantYearlyAmount {
	var records []*record.MerchantYearlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyAmountByMerchant(merchant))
	}
	return records
}
