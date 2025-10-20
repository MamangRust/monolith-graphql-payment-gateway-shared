package transactionprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
)

type transactionStatsMethodProtoMapper struct {
}

func NewTransactionStatsMethodProtoMapper() TransactionStatsMethodProtoMapper {
	return &transactionStatsMethodProtoMapper{}
}

// ToProtoResponseTransactionMonthMethod converts a status, message, and a list of TransactionMonthMethodResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction method data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionMonthMethodResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionMonthMethod containing the status, message, and transaction method data.
func (m *transactionStatsMethodProtoMapper) ToProtoResponseTransactionMonthMethod(status string, message string, pbResponse []*response.TransactionMonthMethodResponse) *pb.ApiResponseTransactionMonthMethod {
	return &pb.ApiResponseTransactionMonthMethod{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransactionMonthMethods(pbResponse),
	}
}

// ToProtoResponseTransactionYearMethod converts a status, message, and a list of TransactionYearMethodResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction method data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionYearMethodResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionYearMethod containing the status, message, and transaction method data.
func (m *transactionStatsMethodProtoMapper) ToProtoResponseTransactionYearMethod(status string, message string, pbResponse []*response.TransactionYearMethodResponse) *pb.ApiResponseTransactionYearMethod {

	return &pb.ApiResponseTransactionYearMethod{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransactionYearMethods(pbResponse),
	}
}

// mapResponseTransactionMonthMethod maps a *response.TransactionMonthMethodResponse to a *pb.TransactionMonthMethodResponse proto message.
//
// It takes a *response.TransactionMonthMethodResponse as input and returns a pointer to the converted TransactionMonthMethodResponse proto object.
// The mapping includes fields like Month, PaymentMethod, TotalTransactions, and TotalAmount.
func (m *transactionStatsMethodProtoMapper) mapResponseTransactionMonthMethod(s *response.TransactionMonthMethodResponse) *pb.TransactionMonthMethodResponse {
	return &pb.TransactionMonthMethodResponse{
		Month:             s.Month,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int32(s.TotalTransactions),
		TotalAmount:       int32(s.TotalAmount),
	}
}

// mapResponseTransactionMonthMethods maps a slice of TransactionMonthMethodResponse to a slice of protobuf TransactionMonthMethodResponse.
//
// It takes a slice of TransactionMonthMethodResponse as input and returns a slice of corresponding protobuf TransactionMonthMethodResponse.
// The mapping includes fields like Month, PaymentMethod, TotalTransactions, and TotalAmount.
func (m *transactionStatsMethodProtoMapper) mapResponseTransactionMonthMethods(s []*response.TransactionMonthMethodResponse) []*pb.TransactionMonthMethodResponse {
	var protoResponses []*pb.TransactionMonthMethodResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.mapResponseTransactionMonthMethod(transaction))
	}
	return protoResponses
}

// mapResponseTransactionYearMethod maps a *response.TransactionYearMethodResponse to a *pb.TransactionYearMethodResponse proto message.
//
// It takes a *response.TransactionYearMethodResponse as input and returns a pointer to the converted TransactionYearMethodResponse proto object.
// The mapping includes fields like Year, PaymentMethod, TotalTransactions, and TotalAmount.
func (m *transactionStatsMethodProtoMapper) mapResponseTransactionYearMethod(s *response.TransactionYearMethodResponse) *pb.TransactionYearMethodResponse {
	return &pb.TransactionYearMethodResponse{
		Year:              s.Year,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int32(s.TotalTransactions),
		TotalAmount:       int32(s.TotalAmount),
	}
}

// mapResponseTransactionYearMethods maps a slice of TransactionYearMethodResponse to a slice of protobuf TransactionYearMethodResponse.
//
// It takes a slice of TransactionYearMethodResponse as input and returns a slice of corresponding protobuf TransactionYearMethodResponse.
// The mapping includes fields like Year, PaymentMethod, TotalTransactions, and TotalAmount.
func (m *transactionStatsMethodProtoMapper) mapResponseTransactionYearMethods(s []*response.TransactionYearMethodResponse) []*pb.TransactionYearMethodResponse {
	var protoResponses []*pb.TransactionYearMethodResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.mapResponseTransactionYearMethod(transaction))
	}
	return protoResponses
}
