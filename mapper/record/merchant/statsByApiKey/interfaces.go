package merchantstatbyapikeyrecordmapper

import (
	db "github.com/MamangRust/monolith-graphql-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
)

// MerchantStatisticMethodByApiKeyMapper maps payment method statistics filtered by API key.
type MerchantStatisticMethodByApiKeyMapper interface {
	// ToMerchantMonthlyPaymentMethodByApikey maps a GetMonthlyPaymentMethodByApikeyRow to a
	// MerchantMonthlyPaymentMethod domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetMonthlyPaymentMethodByApikeyRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
	//     Month, PaymentMethod, and TotalAmount.
	ToMerchantMonthlyPaymentMethodByApikey(ms *db.GetMonthlyPaymentMethodByApikeyRow) *record.MerchantMonthlyPaymentMethod
	// ToMerchantMonthlyPaymentMethodsByApikey maps a slice of GetMonthlyPaymentMethodByApikeyRow to a slice
	// of MerchantMonthlyPaymentMethod domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetMonthlyPaymentMethodByApikeyRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
	//     Month, PaymentMethod, and TotalAmount.
	ToMerchantMonthlyPaymentMethodsByApikey(ms []*db.GetMonthlyPaymentMethodByApikeyRow) []*record.MerchantMonthlyPaymentMethod

	// ToMerchantYearlyPaymentMethodByApikey maps a GetYearlyPaymentMethodByApikeyRow to a
	// MerchantYearlyPaymentMethod domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetYearlyPaymentMethodByApikeyRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
	//     Year, PaymentMethod, and TotalAmount.
	ToMerchantYearlyPaymentMethodByApikey(ms *db.GetYearlyPaymentMethodByApikeyRow) *record.MerchantYearlyPaymentMethod
	// ToMerchantYearlyPaymentMethodsByApikey maps a slice of GetYearlyPaymentMethodByApikeyRow
	// to a slice of MerchantYearlyPaymentMethod domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetYearlyPaymentMethodByApikeyRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantYearlyPaymentMethod containing the mapped data, including
	//     Year, PaymentMethod, and TotalAmount.
	ToMerchantYearlyPaymentMethodsByApikey(ms []*db.GetYearlyPaymentMethodByApikeyRow) []*record.MerchantYearlyPaymentMethod
}

// MerchantStatisticAmountByApiKeyMapper maps transaction amount statistics filtered by API key.
type MerchantStatisticAmountByApiKeyMapper interface {
	// ToMerchantMonthlyAmountByApikey maps a GetMonthlyAmountByApikeyRow to a
	// MerchantMonthlyAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetMonthlyAmountByApikeyRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantMonthlyAmount containing the mapped data, including
	//     Month, and TotalAmount.
	ToMerchantMonthlyAmountByApikey(ms *db.GetMonthlyAmountByApikeyRow) *record.MerchantMonthlyAmount
	// ToMerchantMonthlyAmountsByApikey maps a slice of GetMonthlyAmountByApikeyRow to a slice of
	// MerchantMonthlyAmount domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetMonthlyAmountByApikeyRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.
	ToMerchantMonthlyAmountsByApikey(ms []*db.GetMonthlyAmountByApikeyRow) []*record.MerchantMonthlyAmount

	// ToMerchantYearlyAmountByApikey maps a GetYearlyAmountByApikeyRow to a
	// MerchantYearlyAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetYearlyAmountByApikeyRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantYearlyAmount containing the mapped data, including
	//     Year and TotalAmount.
	ToMerchantYearlyAmountByApikey(ms *db.GetYearlyAmountByApikeyRow) *record.MerchantYearlyAmount
	// ToMerchantYearlyAmountsByApikey maps a slice of GetYearlyAmountByApikeyRow to a slice
	// of MerchantYearlyAmount domain models.
	//
	// Parameters:
	//   - ms: A slice of pointers to GetYearlyAmountByApikeyRow representing the database rows.
	//
	// Returns:
	//   - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
	//     Year and TotalAmount.
	ToMerchantYearlyAmountsByApikey(ms []*db.GetYearlyAmountByApikeyRow) []*record.MerchantYearlyAmount
}

// MerchantStatisticTotalAmountByApiKeyMapper maps total amount statistics filtered by API key.
type MerchantStatisticTotalAmountByApiKeyMapper interface {
	// ToMerchantMonthlyTotalAmountByApikey maps a GetMonthlyTotalAmountByApikeyRow to a
	// MerchantMonthlyTotalAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetMonthlyTotalAmountByApikeyRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
	//     Month, Year, and TotalAmount.
	ToMerchantMonthlyTotalAmountByApikey(ms *db.GetMonthlyTotalAmountByApikeyRow) *record.MerchantMonthlyTotalAmount
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
	ToMerchantMonthlyTotalAmountsByApikey(ms []*db.GetMonthlyTotalAmountByApikeyRow) []*record.MerchantMonthlyTotalAmount

	// ToMerchantYearlyTotalAmountByApikey maps a GetYearlyTotalAmountByApikeyRow to a
	// MerchantYearlyTotalAmount domain model.
	//
	// Parameters:
	//   - ms: A pointer to a GetYearlyTotalAmountByApikeyRow representing the database row.
	//
	// Returns:
	//   - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
	//     Year and TotalAmount.
	ToMerchantYearlyTotalAmountByApikey(ms *db.GetYearlyTotalAmountByApikeyRow) *record.MerchantYearlyTotalAmount
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
	ToMerchantYearlyTotalAmountsByApikey(ms []*db.GetYearlyTotalAmountByApikeyRow) []*record.MerchantYearlyTotalAmount
}
