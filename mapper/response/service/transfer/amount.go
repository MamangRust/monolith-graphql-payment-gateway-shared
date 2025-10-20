package transferresponsemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// transferStatsAmountResponseMapper is a struct that implements the TransferAmountResponseMapper interface.
type transferStatsAmountResponseMapper struct {
}

// NewTransferStatsAmountResponseMapper returns a new instance of TransferAmountResponseMapper
// which provides methods to map TransferMonthAmount and TransferYearAmount domain models
// to TransferMonthAmountResponse and TransferYearAmountResponse API-compatible response types.
func NewTransferStatsAmountResponseMapper() TransferAmountResponseMapper {
	return &transferStatsAmountResponseMapper{}
}

// ToTransferResponseMonthAmount maps a TransferMonthAmount record to a TransferMonthAmountResponse.
//
// Args:
//   - s: A pointer to a TransferMonthAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferMonthAmountResponse containing the mapped data.
func (t *transferStatsAmountResponseMapper) ToTransferResponseMonthAmount(s *record.TransferMonthAmount) *response.TransferMonthAmountResponse {
	return &response.TransferMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferResponsesMonthAmount converts a slice of TransferMonthAmount records
// into a slice of TransferMonthAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransferMonthAmount containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferMonthAmountResponse containing the mapped data.
func (t *transferStatsAmountResponseMapper) ToTransferResponsesMonthAmount(s []*record.TransferMonthAmount) []*response.TransferMonthAmountResponse {
	var transferResponses []*response.TransferMonthAmountResponse
	for _, transfer := range s {
		transferResponses = append(transferResponses, t.ToTransferResponseMonthAmount(transfer))
	}
	return transferResponses
}

// ToTransferResponseYearAmount converts a yearly transfer amount record into a TransferYearAmountResponse.
//
// Args:
//   - s: A pointer to a TransferYearAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferYearAmountResponse containing the mapped data, including Year and TotalAmount.
func (t *transferStatsAmountResponseMapper) ToTransferResponseYearAmount(s *record.TransferYearAmount) *response.TransferYearAmountResponse {
	return &response.TransferYearAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferResponsesYearAmount converts multiple yearly transfer amount records into a slice of TransferYearAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransferYearAmount containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferYearAmountResponse containing the mapped data.
func (t *transferStatsAmountResponseMapper) ToTransferResponsesYearAmount(s []*record.TransferYearAmount) []*response.TransferYearAmountResponse {
	var transferResponses []*response.TransferYearAmountResponse
	for _, transfer := range s {
		transferResponses = append(transferResponses, t.ToTransferResponseYearAmount(transfer))
	}
	return transferResponses
}
