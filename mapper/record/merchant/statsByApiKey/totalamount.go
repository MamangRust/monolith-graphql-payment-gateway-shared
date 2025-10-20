package merchantstatbyapikeyrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantStatisticTotalAmountByApiKeyMapper is an interface that provides methods to map Merchant database rows to MerchantRecord domain models for statistic operations.
type merchantStatisticTotalAmountByApiKeyMapper struct{}

// NewMerchantStatisticTotalAmountByApiKeyMapper creates a new instance of
// MerchantStatisticTotalAmountByApiKeyMapper, providing methods to map
// database rows to MerchantMonthlyTotalAmount and MerchantYearlyTotalAmount
// domain models filtered by API key.
func NewMerchantStatisticTotalAmountByApiKeyMapper() MerchantStatisticTotalAmountByApiKeyMapper {
	return &merchantStatisticTotalAmountByApiKeyMapper{}
}

// ToMerchantMonthlyTotalAmountByApikey maps a GetMonthlyTotalAmountByApikeyRow to a
// MerchantMonthlyTotalAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetMonthlyTotalAmountByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (m *merchantStatisticTotalAmountByApiKeyMapper) ToMerchantMonthlyTotalAmountByApikey(ms *db.GetMonthlyTotalAmountByApikeyRow) *record.MerchantMonthlyTotalAmount {
	return &record.MerchantMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyTotalAmountsByApikey maps a slice of GetMonthlyTotalAmountByApikeyRow
// to a slice of MerchantMonthlyTotalAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetMonthlyTotalAmountByApikeyRow representing the
//     database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyTotalAmount containing the mapped data,
//     including Month, Year, and TotalAmount.
func (m *merchantStatisticTotalAmountByApiKeyMapper) ToMerchantMonthlyTotalAmountsByApikey(ms []*db.GetMonthlyTotalAmountByApikeyRow) []*record.MerchantMonthlyTotalAmount {
	var records []*record.MerchantMonthlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyTotalAmountByApikey(merchant))
	}
	return records
}

// ToMerchantYearlyTotalAmountByApikey maps a GetYearlyTotalAmountByApikeyRow to a
// MerchantYearlyTotalAmount domain model.
//
// Parameters:
//   - ms: A pointer to a GetYearlyTotalAmountByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantStatisticTotalAmountByApiKeyMapper) ToMerchantYearlyTotalAmountByApikey(ms *db.GetYearlyTotalAmountByApikeyRow) *record.MerchantYearlyTotalAmount {
	return &record.MerchantYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyTotalAmountsByApikey maps a slice of GetYearlyTotalAmountByApikeyRow
// to a slice of MerchantYearlyTotalAmount domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetYearlyTotalAmountByApikeyRow representing the
//     database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyTotalAmount containing the mapped data,
//     including Year and TotalAmount.
func (m *merchantStatisticTotalAmountByApiKeyMapper) ToMerchantYearlyTotalAmountsByApikey(ms []*db.GetYearlyTotalAmountByApikeyRow) []*record.MerchantYearlyTotalAmount {
	var records []*record.MerchantYearlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyTotalAmountByApikey(merchant))
	}
	return records
}
