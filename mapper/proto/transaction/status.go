package transactionprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type transactionStatsStatusProtoMapper struct{}

func NewTransactionStatsStatusProtoMapper() TransactionStatsStatusProtoMapper {
	return &transactionStatsStatusProtoMapper{}
}

// ToProtoResponseTransactionMonthStatusSuccess converts a status, message, and a list of TransactionResponseMonthStatusSuccess
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionMonthStatusSuccess containing the status, message, and transaction data.
func (m *transactionStatsStatusProtoMapper) ToProtoResponseTransactionMonthStatusSuccess(status string, message string, pbResponse []*response.TransactionResponseMonthStatusSuccess) *pb.ApiResponseTransactionMonthStatusSuccess {
	return &pb.ApiResponseTransactionMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransactionMonthStatusSuccess(pbResponse),
	}
}

// ToProtoResponseTransactionYearStatusSuccess converts a status, message, and a list of TransactionResponseYearStatusSuccess
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionYearStatusSuccess containing the status, message, and transaction data.
func (m *transactionStatsStatusProtoMapper) ToProtoResponseTransactionYearStatusSuccess(status string, message string, pbResponse []*response.TransactionResponseYearStatusSuccess) *pb.ApiResponseTransactionYearStatusSuccess {
	return &pb.ApiResponseTransactionYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapTransactionResponsesYearStatusSuccess(pbResponse),
	}
}

// ToProtoResponseTransactionMonthStatusFailed converts a status, message, and a list of TransactionResponseMonthStatusFailed
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionMonthStatusFailed containing the status, message, and transaction data.
func (m *transactionStatsStatusProtoMapper) ToProtoResponseTransactionMonthStatusFailed(status string, message string, pbResponse []*response.TransactionResponseMonthStatusFailed) *pb.ApiResponseTransactionMonthStatusFailed {
	return &pb.ApiResponseTransactionMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransactionMonthStatusFailed(pbResponse),
	}
}

// ToProtoResponseTransactionYearStatusFailed converts a status, message, and a list of TransactionResponseYearStatusFailed
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionResponseYearStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionYearStatusFailed containing the status, message, and transaction data.
func (m *transactionStatsStatusProtoMapper) ToProtoResponseTransactionYearStatusFailed(status string, message string, pbResponse []*response.TransactionResponseYearStatusFailed) *pb.ApiResponseTransactionYearStatusFailed {
	return &pb.ApiResponseTransactionYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapTransactionResponsesYearStatusFailed(pbResponse),
	}
}

// mapResponseTransactionMonthStatusSuccess maps a response.TransactionResponseMonthStatusSuccess to its protobuf representation.
//
// It takes a response.TransactionResponseMonthStatusSuccess as input and returns a pointer to the converted TransactionMonthStatusSuccessResponse proto object.
// The mapping includes fields like Year, Month, TotalSuccess, and TotalAmount.
func (t *transactionStatsStatusProtoMapper) mapResponseTransactionMonthStatusSuccess(s *response.TransactionResponseMonthStatusSuccess) *pb.TransactionMonthStatusSuccessResponse {
	return &pb.TransactionMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapResponsesTransactionMonthStatusSuccess converts a slice of TransactionResponseMonthStatusSuccess
// to a slice of protobuf TransactionMonthStatusSuccessResponse.
//
// It iterates through each TransactionResponseMonthStatusSuccess, mapping it to its protobuf equivalent
// using the mapResponseTransactionMonthStatusSuccess function.
//
// Args:
//   - Transactions: A slice of TransactionResponseMonthStatusSuccess to be converted.
//
// Returns:
//
//	A slice of pointers to TransactionMonthStatusSuccessResponse containing the mapped data.
func (t *transactionStatsStatusProtoMapper) mapResponsesTransactionMonthStatusSuccess(Transactions []*response.TransactionResponseMonthStatusSuccess) []*pb.TransactionMonthStatusSuccessResponse {
	var TransactionRecords []*pb.TransactionMonthStatusSuccessResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.mapResponseTransactionMonthStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// mapTransactionResponseYearStatusSuccess maps a response.TransactionResponseYearStatusSuccess to its protobuf representation.
//
// It takes a response.TransactionResponseYearStatusSuccess as input and returns a pointer to the converted TransactionYearStatusSuccessResponse proto object.
// The mapping includes fields like Year, TotalSuccess, and TotalAmount.
func (t *transactionStatsStatusProtoMapper) mapTransactionResponseYearStatusSuccess(s *response.TransactionResponseYearStatusSuccess) *pb.TransactionYearStatusSuccessResponse {
	return &pb.TransactionYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapTransactionResponsesYearStatusSuccess maps a slice of TransactionResponseYearStatusSuccess to a slice of protobuf TransactionYearStatusSuccessResponse.
//
// It takes a slice of TransactionResponseYearStatusSuccess as input and returns a slice of corresponding protobuf TransactionYearStatusSuccessResponse.
// The mapping includes fields like Year, TotalSuccess, and TotalAmount.
func (t *transactionStatsStatusProtoMapper) mapTransactionResponsesYearStatusSuccess(Transactions []*response.TransactionResponseYearStatusSuccess) []*pb.TransactionYearStatusSuccessResponse {
	var TransactionRecords []*pb.TransactionYearStatusSuccessResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.mapTransactionResponseYearStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// mapResponseTransactionMonthStatusFailed maps a response.TransactionResponseMonthStatusFailed to its protobuf representation.
//
// It takes a response.TransactionResponseMonthStatusFailed as input and returns a pointer to the converted TransactionMonthStatusFailedResponse proto object.
// The mapping includes fields like Year, Month, TotalFailed, and TotalAmount.
func (t *transactionStatsStatusProtoMapper) mapResponseTransactionMonthStatusFailed(s *response.TransactionResponseMonthStatusFailed) *pb.TransactionMonthStatusFailedResponse {
	return &pb.TransactionMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponsesTransactionMonthStatusFailed converts a slice of TransactionResponseMonthStatusFailed
// to a slice of protobuf TransactionMonthStatusFailedResponse.
//
// It iterates through each TransactionResponseMonthStatusFailed, mapping it to its protobuf equivalent
// using the mapResponseTransactionMonthStatusFailed function.
//
// Args:
//   - Transactions: A slice of TransactionResponseMonthStatusFailed to be converted.
//
// Returns:
//
//	A slice of pointers to TransactionMonthStatusFailedResponse containing the mapped data.
func (t *transactionStatsStatusProtoMapper) mapResponsesTransactionMonthStatusFailed(Transactions []*response.TransactionResponseMonthStatusFailed) []*pb.TransactionMonthStatusFailedResponse {
	var TransactionRecords []*pb.TransactionMonthStatusFailedResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.mapResponseTransactionMonthStatusFailed(Transaction))
	}

	return TransactionRecords
}

// mapTransactionResponseYearStatusFailed maps a response.TransactionResponseYearStatusFailed to its protobuf representation.
//
// It takes a response.TransactionResponseYearStatusFailed as input and returns a pointer to the converted
// TransactionYearStatusFailedResponse proto object. The mapping includes fields like Year, TotalFailed, and TotalAmount.
//
// Args:
//   - s: A pointer to a TransactionResponseYearStatusFailed to be converted.
//
// Returns:
//   - A pointer to a TransactionYearStatusFailedResponse containing the mapped data.
func (t *transactionStatsStatusProtoMapper) mapTransactionResponseYearStatusFailed(s *response.TransactionResponseYearStatusFailed) *pb.TransactionYearStatusFailedResponse {
	return &pb.TransactionYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapTransactionResponsesYearStatusFailed maps a slice of TransactionResponseYearStatusFailed
// to a slice of protobuf TransactionYearStatusFailedResponse.
//
// It iterates through each TransactionResponseYearStatusFailed, mapping it to its protobuf equivalent
// using the mapTransactionResponseYearStatusFailed function.
//
// Args:
//   - Transactions: A slice of TransactionResponseYearStatusFailed to be converted.
//
// Returns:
//
//	A slice of pointers to TransactionYearStatusFailedResponse containing the mapped data.
func (t *transactionStatsStatusProtoMapper) mapTransactionResponsesYearStatusFailed(Transactions []*response.TransactionResponseYearStatusFailed) []*pb.TransactionYearStatusFailedResponse {
	var TransactionRecords []*pb.TransactionYearStatusFailedResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.mapTransactionResponseYearStatusFailed(Transaction))
	}

	return TransactionRecords
}
