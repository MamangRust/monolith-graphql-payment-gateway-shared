package transactionprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pbhelpers "github.com/MamangRust/monolith-payment-gateway-pb"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
)

type TransactionBaseProtoMapepr interface {
	// ToProtoResponseTransaction maps a single transaction to a protobuf response.
	ToProtoResponseTransaction(status string, message string, pbResponse *response.TransactionResponse) *pb.ApiResponseTransaction
}

type TransactionQueryProtoMapper interface {
	TransactionBaseProtoMapepr

	// ToProtoResponseTransactions maps a list of transactions to a protobuf response.
	ToProtoResponseTransactions(status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponseTransactions

	// ToProtoResponsePaginationTransaction maps paginated transaction records to a protobuf response.
	ToProtoResponsePaginationTransaction(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponsePaginationTransaction

	// ToProtoResponsePaginationTransactionDeleteAt maps paginated soft-deleted transactions to a protobuf response.
	ToProtoResponsePaginationTransactionDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponseDeleteAt) *pb.ApiResponsePaginationTransactionDeleteAt
}

type TransactionCommandProtoMapper interface {
	TransactionBaseProtoMapepr

	ToProtoResponseTransactionDeleteAt(status string, message string, pbResponse *response.TransactionResponseDeleteAt) *pb.ApiResponseTransactionDeleteAt

	// ToProtoResponseTransactionDelete returns a response indicating a transaction has been deleted.
	ToProtoResponseTransactionDelete(status string, message string) *pb.ApiResponseTransactionDelete

	// ToProtoResponseTransactionAll returns all transactions in a protobuf response.
	ToProtoResponseTransactionAll(status string, message string) *pb.ApiResponseTransactionAll
}

type TransactionStatsStatusProtoMapper interface {
	// ToProtoResponseTransactionMonthStatusSuccess maps monthly successful transactions to a protobuf response.
	ToProtoResponseTransactionMonthStatusSuccess(status string, message string, pbResponse []*response.TransactionResponseMonthStatusSuccess) *pb.ApiResponseTransactionMonthStatusSuccess

	// ToProtoResponseTransactionYearStatusSuccess maps yearly successful transactions to a protobuf response.
	ToProtoResponseTransactionYearStatusSuccess(status string, message string, pbResponse []*response.TransactionResponseYearStatusSuccess) *pb.ApiResponseTransactionYearStatusSuccess

	// ToProtoResponseTransactionMonthStatusFailed maps monthly failed transactions to a protobuf response.
	ToProtoResponseTransactionMonthStatusFailed(status string, message string, pbResponse []*response.TransactionResponseMonthStatusFailed) *pb.ApiResponseTransactionMonthStatusFailed

	// ToProtoResponseTransactionYearStatusFailed maps yearly failed transactions to a protobuf response.
	ToProtoResponseTransactionYearStatusFailed(status string, message string, pbResponse []*response.TransactionResponseYearStatusFailed) *pb.ApiResponseTransactionYearStatusFailed
}

// TransactionStatsAmountProtoMapper maps transaction amounts to a protobuf response.
type TransactionStatsAmountProtoMapper interface {
	// ToProtoResponseTransactionMonthAmount maps monthly transaction amounts to a protobuf response.
	ToProtoResponseTransactionMonthAmount(status string, message string, pbResponse []*response.TransactionMonthAmountResponse) *pb.ApiResponseTransactionMonthAmount

	// ToProtoResponseTransactionYearAmount maps yearly transaction amounts to a protobuf response.
	ToProtoResponseTransactionYearAmount(status string, message string, pbResponse []*response.TransactionYearlyAmountResponse) *pb.ApiResponseTransactionYearAmount
}

// TransactionStatsMethodProtoMapper maps transaction methods to a protobuf response.
type TransactionStatsMethodProtoMapper interface {

	// ToProtoResponseTransactionMonthMethod maps monthly transaction methods to a protobuf response.
	ToProtoResponseTransactionMonthMethod(status string, message string, pbResponse []*response.TransactionMonthMethodResponse) *pb.ApiResponseTransactionMonthMethod

	// ToProtoResponseTransactionYearMethod maps yearly transaction methods to a protobuf response.
	ToProtoResponseTransactionYearMethod(status string, message string, pbResponse []*response.TransactionYearMethodResponse) *pb.ApiResponseTransactionYearMethod
}
