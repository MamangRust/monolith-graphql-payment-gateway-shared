package transactionapimapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
)

type TransactionBaseResponseMapper interface {
	// Converts a single transaction response into an API response.
	ToApiResponseTransaction(pbResponse *pb.ApiResponseTransaction) *response.ApiResponseTransaction
}

type TransactionQueryResponseMapper interface {
	TransactionBaseResponseMapper

	// Converts multiple transaction responses into a grouped API response.
	ToApiResponseTransactions(pbResponse *pb.ApiResponseTransactions) *response.ApiResponseTransactions

	// Converts paginated transaction results into an API response.
	ToApiResponsePaginationTransaction(pbResponse *pb.ApiResponsePaginationTransaction) *response.ApiResponsePaginationTransaction

	// Converts paginated soft-deleted transaction results into an API response.
	ToApiResponsePaginationTransactionDeleteAt(pbResponse *pb.ApiResponsePaginationTransactionDeleteAt) *response.ApiResponsePaginationTransactionDeleteAt
}

type TransactionCommandResponseMapper interface {
	TransactionBaseResponseMapper

	ToApiResponseTransactionDeleteAt(pbResponse *pb.ApiResponseTransactionDeleteAt) *response.ApiResponseTransactionDeleteAt

	// Converts a deleted transaction response into an API response.
	ToApiResponseTransactionDelete(pbResponse *pb.ApiResponseTransactionDelete) *response.ApiResponseTransactionDelete

	// Converts all transaction records into a general API response.
	ToApiResponseTransactionAll(pbResponse *pb.ApiResponseTransactionAll) *response.ApiResponseTransactionAll
}

type TransactionStatsStatusResponseMapper interface {
	// Converts monthly transaction stats with success status into an API response.
	ToApiResponseTransactionMonthStatusSuccess(pbResponse *pb.ApiResponseTransactionMonthStatusSuccess) *response.ApiResponseTransactionMonthStatusSuccess

	// Converts yearly transaction stats with success status into an API response.
	ToApiResponseTransactionYearStatusSuccess(pbResponse *pb.ApiResponseTransactionYearStatusSuccess) *response.ApiResponseTransactionYearStatusSuccess

	// Converts monthly transaction stats with failed status into an API response.
	ToApiResponseTransactionMonthStatusFailed(pbResponse *pb.ApiResponseTransactionMonthStatusFailed) *response.ApiResponseTransactionMonthStatusFailed

	// Converts yearly transaction stats with failed status into an API response.
	ToApiResponseTransactionYearStatusFailed(pbResponse *pb.ApiResponseTransactionYearStatusFailed) *response.ApiResponseTransactionYearStatusFailed
}

type TransactionStatsMethodResponseMapper interface {
	// Converts monthly transaction statistics grouped by payment method into an API response.
	ToApiResponseTransactionMonthMethod(pbResponse *pb.ApiResponseTransactionMonthMethod) *response.ApiResponseTransactionMonthMethod

	// Converts yearly transaction statistics grouped by payment method into an API response.
	ToApiResponseTransactionYearMethod(pbResponse *pb.ApiResponseTransactionYearMethod) *response.ApiResponseTransactionYearMethod
}

type TransactionStatsAmountResponseMapper interface {
	// Converts monthly transaction amount statistics into an API response.
	ToApiResponseTransactionMonthAmount(pbResponse *pb.ApiResponseTransactionMonthAmount) *response.ApiResponseTransactionMonthAmount

	// Converts yearly transaction amount statistics into an API response.
	ToApiResponseTransactionYearAmount(pbResponse *pb.ApiResponseTransactionYearAmount) *response.ApiResponseTransactionYearAmount
}
