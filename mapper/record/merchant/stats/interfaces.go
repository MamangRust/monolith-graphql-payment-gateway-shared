package merchantstatsrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantStatisticTotalAmountRecordMapper maps total amount statistics to domain models.
type MerchantStatisticTotalAmountRecordMapper interface {
	// ToMerchantMonthlyTotalAmount maps a GetMonthlyTotalAmountMerchantRow to a
	// MerchantMonthlyTotalAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetMonthlyTotalAmountMerchantRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
	//     Month, Year, and TotalAmount.
	ToMerchantMonthlyTotalAmount(ms *db.GetMonthlyTotalAmountMerchantRow) *record.MerchantMonthlyTotalAmount
	// ToMerchantMonthlyTotalAmounts maps a slice of GetMonthlyTotalAmountMerchantRow to a slice
	// of MerchantMonthlyTotalAmount domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetMonthlyTotalAmountMerchantRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantMonthlyTotalAmount containing the mapped data, including
	//     Month, Year, and TotalAmount.
	ToMerchantMonthlyTotalAmounts(ms []*db.GetMonthlyTotalAmountMerchantRow) []*record.MerchantMonthlyTotalAmount

	// ToMerchantYearlyTotalAmount maps a GetYearlyTotalAmountMerchantRow to a
	// MerchantYearlyTotalAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetYearlyTotalAmountMerchantRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
	//     Year and TotalAmount.
	ToMerchantYearlyTotalAmount(ms *db.GetYearlyTotalAmountMerchantRow) *record.MerchantYearlyTotalAmount
	// ToMerchantYearlyTotalAmounts maps a slice of GetYearlyTotalAmountMerchantRow to a slice
	// of MerchantYearlyTotalAmount domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetYearlyTotalAmountMerchantRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantYearlyTotalAmount containing the mapped data, including
	//     Year and TotalAmount.
	ToMerchantYearlyTotalAmounts(ms []*db.GetYearlyTotalAmountMerchantRow) []*record.MerchantYearlyTotalAmount
}

// MerchantStatisticMethodRecordMapper maps payment method statistics to domain models.
type MerchantStatisticMethodRecordMapper interface {
	// ToMerchantMonthlyPaymentMethod maps a GetMonthlyPaymentMethodsMerchantRow to a
	// MerchantMonthlyPaymentMethod domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetMonthlyPaymentMethodsMerchantRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
	//     Month, PaymentMethod, and TotalAmount.
	ToMerchantMonthlyPaymentMethod(ms *db.GetMonthlyPaymentMethodsMerchantRow) *record.MerchantMonthlyPaymentMethod
	// ToMerchantMonthlyPaymentMethods maps a slice of GetMonthlyPaymentMethodsMerchantRow
	// to a slice of MerchantMonthlyPaymentMethod domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetMonthlyPaymentMethodsMerchantRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
	//     Month, PaymentMethod, and TotalAmount.
	ToMerchantMonthlyPaymentMethods(ms []*db.GetMonthlyPaymentMethodsMerchantRow) []*record.MerchantMonthlyPaymentMethod

	// ToMerchantYearlyPaymentMethod maps a GetYearlyPaymentMethodMerchantRow to a
	// MerchantYearlyPaymentMethod domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetYearlyPaymentMethodMerchantRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
	//     Year, PaymentMethod, and TotalAmount.
	ToMerchantYearlyPaymentMethod(ms *db.GetYearlyPaymentMethodMerchantRow) *record.MerchantYearlyPaymentMethod
	// ToMerchantYearlyPaymentMethods maps a slice of GetYearlyPaymentMethodMerchantRow to a slice
	// of MerchantYearlyPaymentMethod domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetYearlyPaymentMethodMerchantRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantYearlyPaymentMethod containing the mapped data, including
	//     Year, PaymentMethod, and TotalAmount.
	ToMerchantYearlyPaymentMethods(ms []*db.GetYearlyPaymentMethodMerchantRow) []*record.MerchantYearlyPaymentMethod
}

// MerchantStatisticAmountRecordMapper maps amount statistics to domain models.
type MerchantStatisticAmountRecordMapper interface {
	// ToMerchantMonthlyAmount maps a GetMonthlyAmountMerchantRow to a MerchantMonthlyAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetMonthlyAmountMerchantRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantMonthlyAmount
	ToMerchantMonthlyAmount(ms *db.GetMonthlyAmountMerchantRow) *record.MerchantMonthlyAmount
	// ToMerchantMonthlyAmounts maps a slice of GetMonthlyAmountMerchantRow to a slice of MerchantMonthlyAmount domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetMonthlyAmountMerchantRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.
	ToMerchantMonthlyAmounts(ms []*db.GetMonthlyAmountMerchantRow) []*record.MerchantMonthlyAmount

	// ToMerchantYearlyAmount maps a GetYearlyAmountMerchantRow to a MerchantYearlyAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetYearlyAmountMerchantRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantYearlyAmount containing the mapped data, including Year and TotalAmount.
	ToMerchantYearlyAmount(ms *db.GetYearlyAmountMerchantRow) *record.MerchantYearlyAmount
	// ToMerchantYearlyAmounts maps a slice of GetYearlyAmountMerchantRow to a slice
	// of MerchantYearlyAmount domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetYearlyAmountMerchantRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
	//     Year and TotalAmount.
	ToMerchantYearlyAmounts(ms []*db.GetYearlyAmountMerchantRow) []*record.MerchantYearlyAmount
}
