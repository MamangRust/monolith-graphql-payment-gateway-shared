package transactionservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type transactionStatsMethodResponseMapper struct {
}

func NewTransactionStatsMethodResponseMapper() TransactionStatsMethodResponseMapper {
	return &transactionStatsMethodResponseMapper{}
}

// ToTransactionMonthlyMethodResponse maps a record of monthly transaction methods into a TransactionMonthMethodResponse.
//
// Args:
//   - s: A pointer to TransactionMonthMethod containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionMonthMethodResponse containing the mapped data, including Month,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (t *transactionStatsMethodResponseMapper) ToTransactionMonthlyMethodResponse(s *record.TransactionMonthMethod) *response.TransactionMonthMethodResponse {
	return &response.TransactionMonthMethodResponse{
		Month:             s.Month,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int(s.TotalTransactions),
		TotalAmount:       int(s.TotalAmount),
	}
}

// ToTransactionMonthlyMethodResponses maps a slice of TransactionMonthMethod domain models
// into a slice of TransactionMonthMethodResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransactionMonthMethod containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionMonthMethodResponse containing the mapped data, including
//     Month, PaymentMethod, TotalTransactions, and TotalAmount.
func (t *transactionStatsMethodResponseMapper) ToTransactionMonthlyMethodResponses(s []*record.TransactionMonthMethod) []*response.TransactionMonthMethodResponse {
	var transactionResponses []*response.TransactionMonthMethodResponse
	for _, transaction := range s {
		transactionResponses = append(transactionResponses, t.ToTransactionMonthlyMethodResponse(transaction))
	}
	return transactionResponses
}

// ToTransactionYearlyMethodResponse maps a single yearly transaction method record into a TransactionYearMethodResponse.
//
// Args:
//   - s: A pointer to TransactionYearMethod containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionYearMethodResponse containing the mapped data, including Year,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (t *transactionStatsMethodResponseMapper) ToTransactionYearlyMethodResponse(s *record.TransactionYearMethod) *response.TransactionYearMethodResponse {
	return &response.TransactionYearMethodResponse{
		Year:              s.Year,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int(s.TotalTransactions),
		TotalAmount:       int(s.TotalAmount),
	}
}

// ToTransactionYearlyMethodResponses maps a slice of TransactionYearMethod domain models
// into a slice of TransactionYearMethodResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransactionYearMethod containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionYearMethodResponse containing the mapped data, including
//     Year, PaymentMethod, TotalTransactions, and TotalAmount.
func (t *transactionStatsMethodResponseMapper) ToTransactionYearlyMethodResponses(s []*record.TransactionYearMethod) []*response.TransactionYearMethodResponse {
	var transactionResponses []*response.TransactionYearMethodResponse
	for _, transaction := range s {
		transactionResponses = append(transactionResponses, t.ToTransactionYearlyMethodResponse(transaction))
	}
	return transactionResponses
}
