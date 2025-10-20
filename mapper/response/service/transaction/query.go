package transactionservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type transactionQueryResponseMapper struct {
}

func NewTransactionQueryResponseMapper() TransactionQueryResponseMapper {
	return &transactionQueryResponseMapper{}
}

// ToTransactionResponse converts a single transaction record into a TransactionResponse.
//
// Args:
//   - transaction: A pointer to a TransactionRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransactionResponse containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, and UpdatedAt.
func (s *transactionQueryResponseMapper) ToTransactionResponse(transaction *record.TransactionRecord) *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:              transaction.ID,
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          transaction.Amount,
		PaymentMethod:   transaction.PaymentMethod,
		MerchantID:      transaction.MerchantID,
		TransactionTime: transaction.TransactionTime,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}
}

// ToTransactionsResponse converts multiple transaction records into a slice of TransactionResponse.
//
// Args:
//   - transactions: A slice of pointers to TransactionRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponse containing the mapped data, including ID, TransactionNo, CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, and UpdatedAt.
func (s *transactionQueryResponseMapper) ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse {
	responses := make([]*response.TransactionResponse, 0, len(transactions))
	for _, transaction := range transactions {
		responses = append(responses, s.ToTransactionResponse(transaction))
	}
	return responses
}

// ToTransactionResponseDeleteAt maps a soft-deleted transaction record into a TransactionResponseDeleteAt.
//
// Args:
//   - transaction: A pointer to a TransactionRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransactionResponseDeleteAt containing the mapped data,
//     including ID, TransactionNo, CardNumber, Amount, PaymentMethod, MerchantID,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionQueryResponseMapper) ToTransactionResponseDeleteAt(transaction *record.TransactionRecord) *response.TransactionResponseDeleteAt {
	return &response.TransactionResponseDeleteAt{
		ID:              transaction.ID,
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          transaction.Amount,
		PaymentMethod:   transaction.PaymentMethod,
		MerchantID:      transaction.MerchantID,
		TransactionTime: transaction.TransactionTime,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
		DeletedAt:       transaction.DeletedAt,
	}
}

// ToTransactionsResponseDeleteAt converts multiple soft-deleted transaction records into a slice of TransactionResponseDeleteAt.
//
// Args:
//   - transactions: A slice of pointers to TransactionRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponseDeleteAt containing the mapped data, including
//     ID, TransactionNo, CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt,
//     UpdatedAt, and DeletedAt.
func (s *transactionQueryResponseMapper) ToTransactionsResponseDeleteAt(transactions []*record.TransactionRecord) []*response.TransactionResponseDeleteAt {
	responses := make([]*response.TransactionResponseDeleteAt, 0, len(transactions))

	for _, transaction := range transactions {
		responses = append(responses, s.ToTransactionResponseDeleteAt(transaction))
	}
	return responses
}
