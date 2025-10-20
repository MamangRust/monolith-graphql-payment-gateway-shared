package merchantstatbyapikeyrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantStatisticAmountByApiKeyMapper is an interface that provides methods to map Merchant database rows to MerchantRecord domain models for statistic operations.
type merchantStatisticByApiKeyAmountMapper struct{}

// NewMerchantStatisticAmountByApiKeyMapper creates a new instance of
// MerchantStatisticAmountByApiKeyMapper, providing methods to map database rows
// to MerchantMonthlyAmount and MerchantYearlyAmount domain models.
func NewMerchantStatisticAmountByApiKeyMapper() MerchantStatisticAmountByApiKeyMapper {
	return &merchantStatisticByApiKeyAmountMapper{}
}

// ToMerchantMonthlyAmountByApikey maps a GetMonthlyAmountByApikeyRow to a
// MerchantMonthlyAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetMonthlyAmountByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyAmount containing the mapped data, including
//     Month, and TotalAmount.
func (m *merchantStatisticByApiKeyAmountMapper) ToMerchantMonthlyAmountByApikey(ms *db.GetMonthlyAmountByApikeyRow) *record.MerchantMonthlyAmount {
	return &record.MerchantMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyAmountsByApikey maps a slice of GetMonthlyAmountByApikeyRow to a slice of
// MerchantMonthlyAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetMonthlyAmountByApikeyRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.
func (m *merchantStatisticByApiKeyAmountMapper) ToMerchantMonthlyAmountsByApikey(ms []*db.GetMonthlyAmountByApikeyRow) []*record.MerchantMonthlyAmount {
	var records []*record.MerchantMonthlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyAmountByApikey(merchant))
	}
	return records
}

// ToMerchantYearlyAmountByApikey maps a GetYearlyAmountByApikeyRow to a
// MerchantYearlyAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetYearlyAmountByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantStatisticByApiKeyAmountMapper) ToMerchantYearlyAmountByApikey(ms *db.GetYearlyAmountByApikeyRow) *record.MerchantYearlyAmount {
	return &record.MerchantYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyAmountsByApikey maps a slice of GetYearlyAmountByApikeyRow to a slice
// of MerchantYearlyAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetYearlyAmountByApikeyRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantStatisticByApiKeyAmountMapper) ToMerchantYearlyAmountsByApikey(ms []*db.GetYearlyAmountByApikeyRow) []*record.MerchantYearlyAmount {
	var records []*record.MerchantYearlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyAmountByApikey(merchant))
	}
	return records
}
