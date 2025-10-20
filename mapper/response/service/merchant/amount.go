package merchantservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// MerchantAmountResponseMapper provides methods to map MerchantMonthlyAmount and MerchantYearlyAmount records to MerchantResponseMonthlyAmount and MerchantResponseYearlyAmount API-compatible responses.
type merchantAmountResponseMapper struct {
}

// NewMerchantAmountResponseMapper returns a new instance of
// MerchantAmountResponseMapper, which provides methods to map
// MerchantMonthlyAmount and MerchantYearlyAmount records to
// MerchantResponseMonthlyAmount and MerchantResponseYearlyAmount
// API-compatible responses.
func NewMerchantAmountResponseMapper() MerchantAmountResponseMapper {
	return &merchantAmountResponseMapper{}
}

// ToMerchantMonthlyAmount maps a single monthly amount record to a
// MerchantResponseMonthlyAmount API-compatible response.
//
// Args:
//   - ms: A pointer to a MerchantMonthlyAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseMonthlyAmount containing the mapped data, including
//     Month and TotalAmount.
func (s *merchantAmountResponseMapper) ToMerchantMonthlyAmount(ms *record.MerchantMonthlyAmount) *response.MerchantResponseMonthlyAmount {
	return &response.MerchantResponseMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: ms.TotalAmount,
	}
}

// ToMerchantMonthlyAmounts maps multiple monthly amount records into a slice of MerchantResponseMonthlyAmount.
//
// Args:
//   - ms: A slice of pointers to MerchantMonthlyAmount records containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseMonthlyAmount containing the mapped data, including
//     Month and TotalAmount.
func (s *merchantAmountResponseMapper) ToMerchantMonthlyAmounts(ms []*record.MerchantMonthlyAmount) []*response.MerchantResponseMonthlyAmount {
	var response []*response.MerchantResponseMonthlyAmount
	for _, merchant := range ms {
		response = append(response, s.ToMerchantMonthlyAmount(merchant))
	}
	return response
}

// ToMerchantYearlyAmount maps a single yearly amount record to a
// MerchantResponseYearlyAmount API-compatible response.
//
// Args:
//   - ms: A pointer to a MerchantYearlyAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (s *merchantAmountResponseMapper) ToMerchantYearlyAmount(ms *record.MerchantYearlyAmount) *response.MerchantResponseYearlyAmount {
	return &response.MerchantResponseYearlyAmount{
		Year:        ms.Year,
		TotalAmount: ms.TotalAmount,
	}
}

// ToMerchantYearlyAmounts maps multiple yearly amount records into a slice of
// MerchantResponseYearlyAmount.
//
// Args:
//   - ms: A slice of pointers to MerchantYearlyAmount records containing the data
//     to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseYearlyAmount containing the mapped
//     data, including Year and TotalAmount.
func (s *merchantAmountResponseMapper) ToMerchantYearlyAmounts(ms []*record.MerchantYearlyAmount) []*response.MerchantResponseYearlyAmount {
	var response []*response.MerchantResponseYearlyAmount
	for _, merchant := range ms {
		response = append(response, s.ToMerchantYearlyAmount(merchant))
	}
	return response
}
