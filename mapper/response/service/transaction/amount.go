package transactionservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// transactionStatsAmountResponseMapper provides methods to map TransactionMonthAmount and TransactionYearlyAmount
type transactionStatsAmountResponseMapper struct {
}

// NewTransactionStatsAmountResponseMapper returns a new instance of transactionStatsAmountResponseMapper,
// which provides methods to map TransactionMonthAmount and TransactionYearlyAmount domain models to
// API-compatible TransactionMonthAmountResponse and TransactionYearlyAmountResponse types.
func NewTransactionStatsAmountResponseMapper() TransactionStatsAmountResponseMapper {
	return &transactionStatsAmountResponseMapper{}
}

// ToTransactionMonthlyAmountResponse converts a monthly amount record into a TransactionMonthAmountResponse.
//
// Args:
//   - s: A pointer to TransactionMonthAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionMonthAmountResponse containing the mapped data, including Month and TotalAmount.
func (t *transactionStatsAmountResponseMapper) ToTransactionMonthlyAmountResponse(s *record.TransactionMonthAmount) *response.TransactionMonthAmountResponse {
	return &response.TransactionMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionMonthlyAmountResponses maps a slice of TransactionMonthAmount domain models
// into a slice of TransactionMonthAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransactionMonthAmount containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionMonthAmountResponse containing the mapped data, including
//     Month and TotalAmount.
func (t *transactionStatsAmountResponseMapper) ToTransactionMonthlyAmountResponses(s []*record.TransactionMonthAmount) []*response.TransactionMonthAmountResponse {
	var transactionResponses []*response.TransactionMonthAmountResponse
	for _, transaction := range s {
		transactionResponses = append(transactionResponses, t.ToTransactionMonthlyAmountResponse(transaction))
	}
	return transactionResponses
}

// ToTransactionYearlyAmountResponse maps a yearly amount record into a TransactionYearlyAmountResponse.
//
// Args:
//   - s: A pointer to TransactionYearlyAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionYearlyAmountResponse containing the mapped data, including Year and TotalAmount.
func (t *transactionStatsAmountResponseMapper) ToTransactionYearlyAmountResponse(s *record.TransactionYearlyAmount) *response.TransactionYearlyAmountResponse {
	return &response.TransactionYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionYearlyAmountResponses maps a slice of TransactionYearlyAmount domain models
// into a slice of TransactionYearlyAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransactionYearlyAmount containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionYearlyAmountResponse containing the mapped data, including
//     Year and TotalAmount.
func (t *transactionStatsAmountResponseMapper) ToTransactionYearlyAmountResponses(s []*record.TransactionYearlyAmount) []*response.TransactionYearlyAmountResponse {
	var transactionResponses []*response.TransactionYearlyAmountResponse
	for _, transaction := range s {
		transactionResponses = append(transactionResponses, t.ToTransactionYearlyAmountResponse(transaction))
	}
	return transactionResponses
}
