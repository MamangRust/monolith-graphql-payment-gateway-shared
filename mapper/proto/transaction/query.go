package transactionprotomapper

import (
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transaction"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	protomapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type transactonQueryProtoMapper struct {
}

func NewTransactonQueryProtoMapper() TransactionQueryProtoMapper {
	return &transactonQueryProtoMapper{}
}

// ToProtoResponseTransaction converts a status, message, and a TransactionResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The TransactionResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransaction containing the status, message, and transaction data.
func (m *transactonQueryProtoMapper) ToProtoResponseTransaction(status string, message string, pbResponse *response.TransactionResponse) *pb.ApiResponseTransaction {
	return &pb.ApiResponseTransaction{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransaction(pbResponse),
	}
}

// ToProtoResponseTransactions converts a status, message, and a list of TransactionResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactions containing the status, message, and transaction data.
func (m *transactonQueryProtoMapper) ToProtoResponseTransactions(status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponseTransactions {
	return &pb.ApiResponseTransactions{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransaction(pbResponse),
	}
}

// ToProtoResponsePaginationTransaction maps a pagination meta, status, message and a list of *response.TransactionResponse
// to a *pb.ApiResponsePaginationTransaction proto response.
//
// It is used to generate the response for the TransactionService.ListTransaction rpc method.
func (m *transactonQueryProtoMapper) ToProtoResponsePaginationTransaction(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponsePaginationTransaction {

	return &pb.ApiResponsePaginationTransaction{
		Status:         status,
		Message:        message,
		Data:           m.mapResponsesTransaction(pbResponse),
		PaginationMeta: protomapper.MapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationTransactionDeleteAt maps a pagination meta, status, message and a list of *response.TransactionResponseDeleteAt
// to a *pb.ApiResponsePaginationTransactionDeleteAt proto response.
//
// It is used to generate the response for the TransactionService.ListTransactionDeleteAt rpc method.
func (m *transactonQueryProtoMapper) ToProtoResponsePaginationTransactionDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponseDeleteAt) *pb.ApiResponsePaginationTransactionDeleteAt {

	return &pb.ApiResponsePaginationTransactionDeleteAt{
		Status:         status,
		Message:        message,
		Data:           m.mapResponsesTransactionDeleteAt(pbResponse),
		PaginationMeta: protomapper.MapPaginationMeta(pagination),
	}
}

// mapResponseTransaction maps a single response.TransactionResponse to its protobuf representation.
//
// It takes a response.TransactionResponse as input and returns a pointer to the converted TransactionResponse proto object.
func (m *transactonQueryProtoMapper) mapResponseTransaction(transaction *response.TransactionResponse) *pb.TransactionResponse {
	return &pb.TransactionResponse{
		Id:              int32(transaction.ID),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int32(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantId:      int32(transaction.MerchantID),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}

}

// mapResponsesTransaction maps a slice of response.TransactionResponse to a slice of protobuf TransactionResponse.
//
// It takes a slice of response.TransactionResponse as input and returns a slice of corresponding protobuf TransactionResponse.
// The mapping includes fields like ID, TransactionNo, CardNumber, Amount, PaymentMethod, TransactionTime, MerchantId, CreatedAt, and UpdatedAt.
func (m *transactonQueryProtoMapper) mapResponsesTransaction(transactions []*response.TransactionResponse) []*pb.TransactionResponse {
	var result []*pb.TransactionResponse
	for _, transaction := range transactions {
		result = append(result, m.mapResponseTransaction(transaction))
	}
	return result
}

// mapResponseTransactionDeleteAt maps a single response.TransactionResponseDeleteAt to its protobuf representation.
//
// It takes a response.TransactionResponseDeleteAt as input and returns a pointer to the converted TransactionResponseDeleteAt proto object.
// The mapping includes fields like ID, TransactionNo, CardNumber, Amount, PaymentMethod, TransactionTime, MerchantId, CreatedAt, UpdatedAt, and DeletedAt.
func (m *transactonQueryProtoMapper) mapResponseTransactionDeleteAt(transaction *response.TransactionResponseDeleteAt) *pb.TransactionResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if transaction.DeletedAt != nil {
		deletedAt = wrapperspb.String(*transaction.DeletedAt)
	}

	return &pb.TransactionResponseDeleteAt{
		Id:              int32(transaction.ID),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int32(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantId:      int32(transaction.MerchantID),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
		DeletedAt:       deletedAt,
	}

}

// mapResponsesTransactionDeleteAt maps a slice of response.TransactionResponseDeleteAt to a slice of protobuf TransactionResponseDeleteAt.
//
// Args:
//   - transactions: A slice of TransactionResponseDeleteAt to be converted.
//
// Returns:
//
//	A slice of pointers to TransactionResponseDeleteAt containing the mapped data.
func (m *transactonQueryProtoMapper) mapResponsesTransactionDeleteAt(transactions []*response.TransactionResponseDeleteAt) []*pb.TransactionResponseDeleteAt {
	var result []*pb.TransactionResponseDeleteAt

	for _, transaction := range transactions {
		result = append(result, m.mapResponseTransactionDeleteAt(transaction))
	}
	return result
}
