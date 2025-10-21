package transactionprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transaction/stats"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type transactionStatsAmountProtoMapper struct{}

func NewTransactionStatsAmountProtoMapper() TransactionStatsAmountProtoMapper {
	return &transactionStatsAmountProtoMapper{}
}

// ToProtoResponseTransactionMonthAmount converts a status, message, and a list of TransactionMonthAmountResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction month amount data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionMonthAmountResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionMonthAmount containing the status, message, and transaction month amount data.
func (m *transactionStatsAmountProtoMapper) ToProtoResponseTransactionMonthAmount(status string, message string, pbResponse []*response.TransactionMonthAmountResponse) *pb.ApiResponseTransactionMonthAmount {
	return &pb.ApiResponseTransactionMonthAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransactionMonthAmounts(pbResponse),
	}
}

// ToProtoResponseTransactionYearAmount converts a status, message, and a list of TransactionYearlyAmountResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction year amount data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionYearlyAmountResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionYearAmount containing the status, message, and transaction year amount data.
func (m *transactionStatsAmountProtoMapper) ToProtoResponseTransactionYearAmount(status string, message string, pbResponse []*response.TransactionYearlyAmountResponse) *pb.ApiResponseTransactionYearAmount {
	return &pb.ApiResponseTransactionYearAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransactionYearlyAmounts(pbResponse),
	}
}

// mapResponseTransactionMonthAmount maps a *response.TransactionMonthAmountResponse to a *pb.TransactionMonthAmountResponse proto message.
//
// It takes a *response.TransactionMonthAmountResponse as input and returns a pointer to the converted TransactionMonthAmountResponse proto object.
// The mapping includes fields like Month and TotalAmount.
func (m *transactionStatsAmountProtoMapper) mapResponseTransactionMonthAmount(s *response.TransactionMonthAmountResponse) *pb.TransactionMonthAmountResponse {
	return &pb.TransactionMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTransactionMonthAmounts maps a slice of TransactionMonthAmountResponse
// to a slice of protobuf TransactionMonthAmountResponse.
//
// It iterates through each TransactionMonthAmountResponse, converting it to its
// protobuf equivalent using the mapResponseTransactionMonthAmount function.
//
// Args:
//   - s: A slice of TransactionMonthAmountResponse to be converted.
//
// Returns:
//   - A slice of pointers to TransactionMonthAmountResponse containing the mapped data.
func (m *transactionStatsAmountProtoMapper) mapResponseTransactionMonthAmounts(s []*response.TransactionMonthAmountResponse) []*pb.TransactionMonthAmountResponse {
	var protoResponses []*pb.TransactionMonthAmountResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.mapResponseTransactionMonthAmount(transaction))
	}
	return protoResponses
}

// mapResponseTransactionYearlyAmount maps a *response.TransactionYearlyAmountResponse to a *pb.TransactionYearlyAmountResponse proto message.
//
// It takes a *response.TransactionYearlyAmountResponse as input and returns a pointer to the converted TransactionYearlyAmountResponse proto object.
// The mapping includes fields like Year and TotalAmount.
func (m *transactionStatsAmountProtoMapper) mapResponseTransactionYearlyAmount(s *response.TransactionYearlyAmountResponse) *pb.TransactionYearlyAmountResponse {
	return &pb.TransactionYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTransactionYearlyAmounts maps a slice of TransactionYearlyAmountResponse
// to a slice of protobuf TransactionYearlyAmountResponse.
//
// It iterates through each TransactionYearlyAmountResponse, converting it to its
// protobuf equivalent using the mapResponseTransactionYearlyAmount function.
//
// Args:
//   - s: A slice of TransactionYearlyAmountResponse to be converted.
//
// Returns:
//   - A slice of pointers to TransactionYearlyAmountResponse containing the mapped data.
func (m *transactionStatsAmountProtoMapper) mapResponseTransactionYearlyAmounts(s []*response.TransactionYearlyAmountResponse) []*pb.TransactionYearlyAmountResponse {
	var protoResponses []*pb.TransactionYearlyAmountResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.mapResponseTransactionYearlyAmount(transaction))
	}
	return protoResponses
}
