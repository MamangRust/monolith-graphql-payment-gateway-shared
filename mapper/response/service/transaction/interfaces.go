package transactionservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type TransactionBaseResponseMapper interface {
	// Converts a single transaction record into a TransactionResponse.
	ToTransactionResponse(transaction *record.TransactionRecord) *response.TransactionResponse
}

type TransactionQueryResponseMapper interface {
	TransactionBaseResponseMapper

	// Converts multiple transaction records into a slice of TransactionResponse.
	ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse

	// Converts multiple soft-deleted transaction records into a slice of TransactionResponseDeleteAt.
	ToTransactionsResponseDeleteAt(transactions []*record.TransactionRecord) []*response.TransactionResponseDeleteAt
}

type TransactionCommandResponseMapper interface {
	TransactionBaseResponseMapper

	// Converts a soft-deleted transaction record into a TransactionResponseDeleteAt.
	ToTransactionResponseDeleteAt(transaction *record.TransactionRecord) *response.TransactionResponseDeleteAt
}

type TransactionStatsStatusResponseMapper interface {
	// Converts a monthly success status transaction record into a TransactionResponseMonthStatusSuccess.
	ToTransactionResponseMonthStatusSuccess(s *record.TransactionRecordMonthStatusSuccess) *response.TransactionResponseMonthStatusSuccess

	// Converts multiple monthly success status transaction records into a slice of TransactionResponseMonthStatusSuccess.
	ToTransactionResponsesMonthStatusSuccess(Transactions []*record.TransactionRecordMonthStatusSuccess) []*response.TransactionResponseMonthStatusSuccess

	// Converts a yearly success status transaction record into a TransactionResponseYearStatusSuccess.
	ToTransactionResponseYearStatusSuccess(s *record.TransactionRecordYearStatusSuccess) *response.TransactionResponseYearStatusSuccess

	// Converts multiple yearly success status transaction records into a slice of TransactionResponseYearStatusSuccess.
	ToTransactionResponsesYearStatusSuccess(Transactions []*record.TransactionRecordYearStatusSuccess) []*response.TransactionResponseYearStatusSuccess

	// Converts a monthly failed status transaction record into a TransactionResponseMonthStatusFailed.
	ToTransactionResponseMonthStatusFailed(s *record.TransactionRecordMonthStatusFailed) *response.TransactionResponseMonthStatusFailed

	// Converts multiple monthly failed status transaction records into a slice of TransactionResponseMonthStatusFailed.
	ToTransactionResponsesMonthStatusFailed(Transactions []*record.TransactionRecordMonthStatusFailed) []*response.TransactionResponseMonthStatusFailed

	// Converts a yearly failed status transaction record into a TransactionResponseYearStatusFailed.
	ToTransactionResponseYearStatusFailed(s *record.TransactionRecordYearStatusFailed) *response.TransactionResponseYearStatusFailed

	// Converts multiple yearly failed status transaction records into a slice of TransactionResponseYearStatusFailed.
	ToTransactionResponsesYearStatusFailed(Transactions []*record.TransactionRecordYearStatusFailed) []*response.TransactionResponseYearStatusFailed
}

type TransactionStatsMethodResponseMapper interface {
	// Converts a monthly method record into a TransactionMonthMethodResponse.
	ToTransactionMonthlyMethodResponse(s *record.TransactionMonthMethod) *response.TransactionMonthMethodResponse

	// Converts multiple monthly method records into a slice of TransactionMonthMethodResponse.
	ToTransactionMonthlyMethodResponses(s []*record.TransactionMonthMethod) []*response.TransactionMonthMethodResponse

	// Converts a yearly method record into a TransactionYearMethodResponse.
	ToTransactionYearlyMethodResponse(s *record.TransactionYearMethod) *response.TransactionYearMethodResponse

	// Converts multiple yearly method records into a slice of TransactionYearMethodResponse.
	ToTransactionYearlyMethodResponses(s []*record.TransactionYearMethod) []*response.TransactionYearMethodResponse
}

type TransactionStatsAmountResponseMapper interface {
	// Converts a monthly amount record into a TransactionMonthAmountResponse.
	ToTransactionMonthlyAmountResponse(s *record.TransactionMonthAmount) *response.TransactionMonthAmountResponse

	// Converts multiple monthly amount records into a slice of TransactionMonthAmountResponse.
	ToTransactionMonthlyAmountResponses(s []*record.TransactionMonthAmount) []*response.TransactionMonthAmountResponse

	// Converts a yearly amount record into a TransactionYearlyAmountResponse.
	ToTransactionYearlyAmountResponse(s *record.TransactionYearlyAmount) *response.TransactionYearlyAmountResponse

	// Converts multiple yearly amount records into a slice of TransactionYearlyAmountResponse.
	ToTransactionYearlyAmountResponses(s []*record.TransactionYearlyAmount) []*response.TransactionYearlyAmountResponse
}
