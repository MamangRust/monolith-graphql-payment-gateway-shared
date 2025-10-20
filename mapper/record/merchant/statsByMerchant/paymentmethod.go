package merchantstatbyidrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantStatisticMethodByMerchantMapper maps merchant payment method statistic rows to domain models.
type merchantStatisticMethodByMerchantMapper struct {
}

// NewMerchantStatisticMethodByMerchantRecordMapper returns a new instance of
// MerchantStatisticMethodByMerchantMapper, which provides methods to map
// Merchant database rows to MerchantRecord domain models for statistic operations.
func NewMerchantStatisticMethodByMerchantRecordMapper() MerchantStatisticMethodByMerchantMapper {
	return &merchantStatisticMethodByMerchantMapper{}
}

// ToMerchantMonthlyPaymentMethodByMerchant maps a GetMonthlyPaymentMethodByMerchantsRow to a
// MerchantMonthlyPaymentMethod domain model.
//
// Parameters:
//   - ms: A pointer to a GetMonthlyPaymentMethodByMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodByMerchantMapper) ToMerchantMonthlyPaymentMethodByMerchant(ms *db.GetMonthlyPaymentMethodByMerchantsRow) *record.MerchantMonthlyPaymentMethod {
	return &record.MerchantMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyPaymentMethodsByMerchant maps a slice of GetMonthlyPaymentMethodByMerchantsRow
// to a slice of MerchantMonthlyPaymentMethod domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetMonthlyPaymentMethodByMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodByMerchantMapper) ToMerchantMonthlyPaymentMethodsByMerchant(ms []*db.GetMonthlyPaymentMethodByMerchantsRow) []*record.MerchantMonthlyPaymentMethod {
	var records []*record.MerchantMonthlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyPaymentMethodByMerchant(merchant))
	}
	return records
}

// ToMerchantYearlyPaymentMethodByMerchant maps a GetYearlyPaymentMethodByMerchantsRow to a
// MerchantYearlyPaymentMethod domain model.
//
// Parameters:
//   - ms: A pointer to a GetYearlyPaymentMethodByMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodByMerchantMapper) ToMerchantYearlyPaymentMethodByMerchant(ms *db.GetYearlyPaymentMethodByMerchantsRow) *record.MerchantYearlyPaymentMethod {
	return &record.MerchantYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantYearlyPaymentMethodsByMerchant maps a slice of GetYearlyPaymentMethodByMerchantsRow
// to a slice of MerchantYearlyPaymentMethod domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetYearlyPaymentMethodByMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodByMerchantMapper) ToMerchantYearlyPaymentMethodsByMerchant(ms []*db.GetYearlyPaymentMethodByMerchantsRow) []*record.MerchantYearlyPaymentMethod {
	var records []*record.MerchantYearlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyPaymentMethodByMerchant(merchant))
	}
	return records
}
