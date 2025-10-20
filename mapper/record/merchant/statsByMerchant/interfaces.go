package merchantstatbyidrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantStatisticMethodByMerchantMapper maps merchant payment method statistic rows to domain models.
type MerchantStatisticMethodByMerchantMapper interface {
	// ToMerchantMonthlyPaymentMethodByMerchant maps a GetMonthlyPaymentMethodByMerchantsRow to a
	// MerchantMonthlyPaymentMethod domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetMonthlyPaymentMethodByMerchantsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
	//     Month, PaymentMethod, and TotalAmount.
	ToMerchantMonthlyPaymentMethodByMerchant(ms *db.GetMonthlyPaymentMethodByMerchantsRow) *record.MerchantMonthlyPaymentMethod
	// ToMerchantMonthlyPaymentMethodsByMerchant maps a slice of GetMonthlyPaymentMethodByMerchantsRow
	// to a slice of MerchantMonthlyPaymentMethod domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetMonthlyPaymentMethodByMerchantsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
	//     Month, PaymentMethod, and TotalAmount.
	ToMerchantMonthlyPaymentMethodsByMerchant(ms []*db.GetMonthlyPaymentMethodByMerchantsRow) []*record.MerchantMonthlyPaymentMethod

	// ToMerchantYearlyPaymentMethodByMerchant maps a GetYearlyPaymentMethodByMerchantsRow to a
	// MerchantYearlyPaymentMethod domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetYearlyPaymentMethodByMerchantsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
	//     Year, PaymentMethod, and TotalAmount.
	ToMerchantYearlyPaymentMethodByMerchant(ms *db.GetYearlyPaymentMethodByMerchantsRow) *record.MerchantYearlyPaymentMethod
	// ToMerchantYearlyPaymentMethodsByMerchant maps a slice of GetYearlyPaymentMethodByMerchantsRow
	// to a slice of MerchantYearlyPaymentMethod domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetYearlyPaymentMethodByMerchantsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantYearlyPaymentMethod containing the mapped data, including
	//     Year, PaymentMethod, and TotalAmount.
	ToMerchantYearlyPaymentMethodsByMerchant(ms []*db.GetYearlyPaymentMethodByMerchantsRow) []*record.MerchantYearlyPaymentMethod
}

// MerchantStatisticAmountByMerchantMapper maps merchant amount statistic rows to domain models.
type MerchantStatisticAmountByMerchantMapper interface {
	// ToMerchantMonthlyAmountByMerchant maps a GetMonthlyAmountByMerchantsRow to a
	// MerchantMonthlyAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetMonthlyAmountByMerchantsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.
	ToMerchantMonthlyAmountByMerchant(ms *db.GetMonthlyAmountByMerchantsRow) *record.MerchantMonthlyAmount
	// ToMerchantMonthlyAmountsByMerchant maps a slice of GetMonthlyAmountByMerchantsRow
	// to a slice of MerchantMonthlyAmount domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetMonthlyAmountByMerchantsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including
	//     Month and TotalAmount.
	ToMerchantMonthlyAmountsByMerchant(ms []*db.GetMonthlyAmountByMerchantsRow) []*record.MerchantMonthlyAmount

	// ToMerchantYearlyAmountByMerchant maps a GetYearlyAmountByMerchantsRow to a
	// MerchantYearlyAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetYearlyAmountByMerchantsRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantYearlyAmount containing the mapped data, including
	//     Year and TotalAmount.
	ToMerchantYearlyAmountByMerchant(ms *db.GetYearlyAmountByMerchantsRow) *record.MerchantYearlyAmount
	// ToMerchantYearlyAmountsByMerchant maps a slice of GetYearlyAmountByMerchantsRow
	// to a slice of MerchantYearlyAmount domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetYearlyAmountByMerchantsRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
	//     Year and TotalAmount.

	ToMerchantYearlyAmountsByMerchant(ms []*db.GetYearlyAmountByMerchantsRow) []*record.MerchantYearlyAmount
}

// MerchantStatisticTotalAmountByMerchantMapper maps merchant total amount statistic rows to domain models.
type MerchantStatisticTotalAmountByMerchantMapper interface {
	// ToMerchantMonthlyTotalAmountByMerchant maps a GetMonthlyTotalAmountByMerchantRow to a
	// MerchantMonthlyTotalAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetMonthlyTotalAmountByMerchantRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
	//     Month, Year, and TotalAmount.
	ToMerchantMonthlyTotalAmountByMerchant(ms *db.GetMonthlyTotalAmountByMerchantRow) *record.MerchantMonthlyTotalAmount
	// ToMerchantMonthlyTotalAmountsByMerchant maps a slice of GetMonthlyTotalAmountByMerchantRow to a
	// slice of MerchantMonthlyTotalAmount domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetMonthlyTotalAmountByMerchantRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantMonthlyTotalAmount containing the mapped data, including
	//     Month, Year, and TotalAmount.
	ToMerchantMonthlyTotalAmountsByMerchant(ms []*db.GetMonthlyTotalAmountByMerchantRow) []*record.MerchantMonthlyTotalAmount

	// ToMerchantYearlyTotalAmountByMerchant maps a GetYearlyTotalAmountByMerchantRow to a
	// MerchantYearlyTotalAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetYearlyTotalAmountByMerchantRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
	//     Year and TotalAmount.
	ToMerchantYearlyTotalAmountByMerchant(ms *db.GetYearlyTotalAmountByMerchantRow) *record.MerchantYearlyTotalAmount
	// ToMerchantYearlyTotalAmountsByMerchant maps a slice of GetYearlyTotalAmountByMerchantRow to a
	// slice of MerchantYearlyTotalAmount domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetYearlyTotalAmountByMerchantRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantYearlyTotalAmount containing the mapped data, including
	//     Year and TotalAmount.
	ToMerchantYearlyTotalAmountsByMerchant(ms []*db.GetYearlyTotalAmountByMerchantRow) []*record.MerchantYearlyTotalAmount
}
