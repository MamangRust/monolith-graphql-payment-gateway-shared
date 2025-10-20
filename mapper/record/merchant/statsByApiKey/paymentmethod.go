package merchantstatbyapikeyrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

type merchantStatisticMethodByApiKeyMapper struct{}

// NewMerchantStatisticMethodByApiKeyMapper creates a new instance of
// MerchantStatisticMethodByApiKeyMapper, providing methods to map database rows
// to MerchantMonthlyPaymentMethod and MerchantYearlyPaymentMethod domain models
// filtered by API key.
func NewMerchantStatisticMethodByApiKeyMapper() MerchantStatisticMethodByApiKeyMapper {
	return &merchantStatisticMethodByApiKeyMapper{}
}

// ToMerchantMonthlyPaymentMethodByApikey maps a GetMonthlyPaymentMethodByApikeyRow to a
// MerchantMonthlyPaymentMethod domain model.
//
// Parameters:
//   - ms: A pointer to a GetMonthlyPaymentMethodByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodByApiKeyMapper) ToMerchantMonthlyPaymentMethodByApikey(ms *db.GetMonthlyPaymentMethodByApikeyRow) *record.MerchantMonthlyPaymentMethod {
	return &record.MerchantMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyPaymentMethodsByApikey maps a slice of GetMonthlyPaymentMethodByApikeyRow to a slice
// of MerchantMonthlyPaymentMethod domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetMonthlyPaymentMethodByApikeyRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodByApiKeyMapper) ToMerchantMonthlyPaymentMethodsByApikey(ms []*db.GetMonthlyPaymentMethodByApikeyRow) []*record.MerchantMonthlyPaymentMethod {
	var records []*record.MerchantMonthlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyPaymentMethodByApikey(merchant))
	}
	return records
}

// ToMerchantYearlyPaymentMethodByApikey maps a GetYearlyPaymentMethodByApikeyRow to a
// MerchantYearlyPaymentMethod domain model.
//
// Parameters:
//   - ms: A pointer to a GetYearlyPaymentMethodByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodByApiKeyMapper) ToMerchantYearlyPaymentMethodByApikey(ms *db.GetYearlyPaymentMethodByApikeyRow) *record.MerchantYearlyPaymentMethod {
	return &record.MerchantYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantYearlyPaymentMethodsByApikey maps a slice of GetYearlyPaymentMethodByApikeyRow
// to a slice of MerchantYearlyPaymentMethod domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetYearlyPaymentMethodByApikeyRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodByApiKeyMapper) ToMerchantYearlyPaymentMethodsByApikey(ms []*db.GetYearlyPaymentMethodByApikeyRow) []*record.MerchantYearlyPaymentMethod {
	var records []*record.MerchantYearlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyPaymentMethodByApikey(merchant))
	}
	return records
}
