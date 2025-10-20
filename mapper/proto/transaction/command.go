package transactionprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	helpersproto "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto/helpers"
)

type transactionCommandProtoMapper struct {
}

func NewTransactionCommandProtoMapper() TransactionCommandProtoMapper {
	return &transactionCommandProtoMapper{}
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
func (m *transactionCommandProtoMapper) ToProtoResponseTransaction(status string, message string, pbResponse *response.TransactionResponse) *pb.ApiResponseTransaction {
	return &pb.ApiResponseTransaction{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransaction(pbResponse),
	}
}

// ToProtoResponseTransactionDeleteAt maps a status, message, and a TransactionResponseDeleteAt
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The TransactionResponseDeleteAt that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionDeleteAt containing the status, message, and transaction data.
func (m *transactionCommandProtoMapper) ToProtoResponseTransactionDeleteAt(status string, message string, pbResponse *response.TransactionResponseDeleteAt) *pb.ApiResponseTransactionDeleteAt {
	return &pb.ApiResponseTransactionDeleteAt{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransactionDeleteAt(pbResponse),
	}
}

// ToProtoResponseTransactionDelete generates a protobuf API response for a transaction delete operation.
//
// It includes the status and message for the API response.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//
// Returns:
//
//	A pointer to ApiResponseTransactionDelete containing the status and message.
func (m *transactionCommandProtoMapper) ToProtoResponseTransactionDelete(status string, message string) *pb.ApiResponseTransactionDelete {
	return &pb.ApiResponseTransactionDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseTransactionAll generates a protobuf API response for all transactions.
//
// It includes the status and message for the API response.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//
// Returns:
//
//	A pointer to ApiResponseTransactionAll containing the status and message.
func (m *transactionCommandProtoMapper) ToProtoResponseTransactionAll(status string, message string) *pb.ApiResponseTransactionAll {
	return &pb.ApiResponseTransactionAll{
		Status:  status,
		Message: message,
	}
}

// mapResponseTransaction maps a single response.TransactionResponse to its protobuf representation.
//
// It takes a response.TransactionResponse as input and returns a pointer to the converted TransactionResponse proto object.
func (m *transactionCommandProtoMapper) mapResponseTransaction(transaction *response.TransactionResponse) *pb.TransactionResponse {
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

// mapResponseTransactionDeleteAt maps a single response.TransactionResponseDeleteAt to its protobuf representation.
//
// It takes a response.TransactionResponseDeleteAt as input and returns a pointer to the converted TransactionResponseDeleteAt proto object.
// The mapping includes fields like ID, TransactionNo, CardNumber, Amount, PaymentMethod, TransactionTime, MerchantId, CreatedAt, UpdatedAt, and DeletedAt.
func (m *transactionCommandProtoMapper) mapResponseTransactionDeleteAt(transaction *response.TransactionResponseDeleteAt) *pb.TransactionResponseDeleteAt {
	res := &pb.TransactionResponseDeleteAt{
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

	if transaction.DeletedAt != nil {
		res.DeletedAt = helpersproto.StringPtrToWrapper(transaction.DeletedAt)
	}

	return res
}
