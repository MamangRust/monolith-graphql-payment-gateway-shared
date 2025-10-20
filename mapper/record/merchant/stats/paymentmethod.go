package merchantstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantStatisticMethodRecordMapper is an interface that provides methods to map Merchant database rows to MerchantRecord domain models for statistic operations.
type merchantStatisticMethodRecordMapper struct{}

// NewMerchantStatisticMethodRecordMapper returns a new instance of
// MerchantStatisticMethodRecordMapper, which provides methods to map payment
// method statistics to domain models.
func NewMerchantStatisticMethodRecordMapper() MerchantStatisticMethodRecordMapper {
	return &merchantStatisticMethodRecordMapper{}
}

// ToMerchantMonthlyPaymentMethod maps a GetMonthlyPaymentMethodsMerchantRow to a
// MerchantMonthlyPaymentMethod domain model.
//
// Parameters:
//   - ms: A pointer to a GetMonthlyPaymentMethodsMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodRecordMapper) ToMerchantMonthlyPaymentMethod(ms *db.GetMonthlyPaymentMethodsMerchantRow) *record.MerchantMonthlyPaymentMethod {
	return &record.MerchantMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyPaymentMethods maps a slice of GetMonthlyPaymentMethodsMerchantRow
// to a slice of MerchantMonthlyPaymentMethod domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetMonthlyPaymentMethodsMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodRecordMapper) ToMerchantMonthlyPaymentMethods(ms []*db.GetMonthlyPaymentMethodsMerchantRow) []*record.MerchantMonthlyPaymentMethod {
	var records []*record.MerchantMonthlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyPaymentMethod(merchant))
	}
	return records
}

// ToMerchantYearlyPaymentMethod maps a GetYearlyPaymentMethodMerchantRow to a
// MerchantYearlyPaymentMethod domain model.
//
// Parameters:
//   - ms: A pointer to a GetYearlyPaymentMethodMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodRecordMapper) ToMerchantYearlyPaymentMethod(ms *db.GetYearlyPaymentMethodMerchantRow) *record.MerchantYearlyPaymentMethod {
	return &record.MerchantYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantYearlyPaymentMethods maps a slice of GetYearlyPaymentMethodMerchantRow to a slice
// of MerchantYearlyPaymentMethod domain models.
//
// Parameters:
//   - ms: A slice of pointers to GetYearlyPaymentMethodMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (m *merchantStatisticMethodRecordMapper) ToMerchantYearlyPaymentMethods(ms []*db.GetYearlyPaymentMethodMerchantRow) []*record.MerchantYearlyPaymentMethod {
	var records []*record.MerchantYearlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyPaymentMethod(merchant))
	}
	return records
}
