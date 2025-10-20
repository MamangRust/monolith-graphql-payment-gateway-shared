package merchantservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// MerchantTotalAmountResponseMapper provides methods to map MerchantMonthlyTotalAmount and MerchantYearlyTotalAmount records to their corresponding API-compatible response types.
type merchantTotalAmountResponseMapper struct{}

// NewMerchantTotalAmountResponseMapper returns a new instance of
// MerchantTotalAmountResponseMapper, which provides methods to map
// MerchantMonthlyTotalAmount and MerchantYearlyTotalAmount domain models
// to their corresponding API-compatible response types.
func NewMerchantTotalAmountResponseMapper() MerchantTotalAmountResponseMapper {
	return &merchantTotalAmountResponseMapper{}
}

// ToMerchantMonthlyTotalAmount maps a single MerchantMonthlyTotalAmount record to a
// MerchantResponseMonthlyTotalAmount API-compatible response.
//
// Args:
//   - ms: A pointer to a MerchantMonthlyTotalAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (s *merchantTotalAmountResponseMapper) ToMerchantMonthlyTotalAmount(ms *record.MerchantMonthlyTotalAmount) *response.MerchantResponseMonthlyTotalAmount {
	return &response.MerchantResponseMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: ms.TotalAmount,
	}
}

// ToMerchantMonthlyTotalAmounts maps multiple monthly total amount records into a slice of
// MerchantResponseMonthlyTotalAmount.
//
// Args:
//   - ms: A slice of pointers to MerchantMonthlyTotalAmount records containing the data
//     to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseMonthlyTotalAmount containing the mapped
//     data, including Month, Year, and TotalAmount.
func (s *merchantTotalAmountResponseMapper) ToMerchantMonthlyTotalAmounts(ms []*record.MerchantMonthlyTotalAmount) []*response.MerchantResponseMonthlyTotalAmount {
	var response []*response.MerchantResponseMonthlyTotalAmount
	for _, merchant := range ms {
		response = append(response, s.ToMerchantMonthlyTotalAmount(merchant))
	}
	return response
}

// ToMerchantYearlyTotalAmount maps a single yearly total amount record to a
// MerchantResponseYearlyTotalAmount API-compatible response.
//
// Args:
//   - ms: A pointer to a MerchantYearlyTotalAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (s *merchantTotalAmountResponseMapper) ToMerchantYearlyTotalAmount(ms *record.MerchantYearlyTotalAmount) *response.MerchantResponseYearlyTotalAmount {
	return &response.MerchantResponseYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: ms.TotalAmount,
	}
}

// ToMerchantYearlyTotalAmounts maps multiple yearly total amount records into a slice of
// MerchantResponseYearlyTotalAmount API-compatible responses.
//
// Args:
//   - ms: A slice of pointers to MerchantYearlyTotalAmount records containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.

func (s *merchantTotalAmountResponseMapper) ToMerchantYearlyTotalAmounts(ms []*record.MerchantYearlyTotalAmount) []*response.MerchantResponseYearlyTotalAmount {
	var response []*response.MerchantResponseYearlyTotalAmount
	for _, merchant := range ms {
		response = append(response, s.ToMerchantYearlyTotalAmount(merchant))
	}
	return response
}
