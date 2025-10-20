package merchantservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// MerchantPaymentMethodResponseMapper provides methods to map MerchantMonthlyPaymentMethod and MerchantYearlyPaymentMethod
type merchantPaymentMethodResponseMapper struct {
}

// NewMerchantPaymentMethodResponseMapper returns a new instance of
// MerchantPaymentMethodResponseMapper.
func NewMerchantPaymentMethodResponseMapper() MerchantPaymentMethodResponseMapper {
	return &merchantPaymentMethodResponseMapper{}
}

// ToMerchantMonthlyPaymentMethod maps a single monthly payment method record to a
// MerchantResponseMonthlyPaymentMethod API-compatible response.
//
// Args:
//   - ms: A pointer to a MerchantMonthlyPaymentMethod containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (s *merchantPaymentMethodResponseMapper) ToMerchantMonthlyPaymentMethod(ms *record.MerchantMonthlyPaymentMethod) *response.MerchantResponseMonthlyPaymentMethod {
	return &response.MerchantResponseMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   ms.TotalAmount,
	}
}

// ToMerchantMonthlyPaymentMethods maps multiple monthly payment method records into a slice of MerchantResponseMonthlyPaymentMethod.
//
// Args:
//   - ms: A slice of pointers to MerchantMonthlyPaymentMethod records containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (s *merchantPaymentMethodResponseMapper) ToMerchantMonthlyPaymentMethods(ms []*record.MerchantMonthlyPaymentMethod) []*response.MerchantResponseMonthlyPaymentMethod {
	var response []*response.MerchantResponseMonthlyPaymentMethod
	for _, merchant := range ms {
		response = append(response, s.ToMerchantMonthlyPaymentMethod(merchant))
	}
	return response
}

// ToMerchantYearlyPaymentMethod maps a single yearly payment method record to a
// MerchantResponseYearlyPaymentMethod API-compatible response.
//
// Args:
//   - ms: A pointer to a MerchantYearlyPaymentMethod containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (s *merchantPaymentMethodResponseMapper) ToMerchantYearlyPaymentMethod(ms *record.MerchantYearlyPaymentMethod) *response.MerchantResponseYearlyPaymentMethod {
	return &response.MerchantResponseYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   ms.TotalAmount,
	}
}

// ToMerchantYearlyPaymentMethods maps multiple yearly payment method records into a slice of MerchantResponseYearlyPaymentMethod.
//
// Args:
//   - ms: A slice of pointers to MerchantYearlyPaymentMethod records containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (s *merchantPaymentMethodResponseMapper) ToMerchantYearlyPaymentMethods(ms []*record.MerchantYearlyPaymentMethod) []*response.MerchantResponseYearlyPaymentMethod {
	var response []*response.MerchantResponseYearlyPaymentMethod
	for _, merchant := range ms {
		response = append(response, s.ToMerchantYearlyPaymentMethod(merchant))
	}
	return response
}
