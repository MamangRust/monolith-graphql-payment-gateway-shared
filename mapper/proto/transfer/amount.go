package transferprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// TransferStatsAmountProtoMapper is responsible for mapping Transfer responses to their corresponding protobuf representations.
type transferStatsAmountProtoMapper struct {
}

// NewTransferStatsAmountProtoMapper returns a new instance of TransferStatsAmountProtoMapper.
func NewTransferStatsAmountProtoMapper() TransferStatsAmountProtoMapper {
	return &transferStatsAmountProtoMapper{}
}

// ToProtoResponseTransferMonthAmount maps a status, message, and a list of TransferMonthAmountResponse
// to an ApiResponseTransferMonthAmount proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferMonthAmountResponse to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferMonthAmount containing the mapped data.
func (m *transferStatsAmountProtoMapper) ToProtoResponseTransferMonthAmount(status string, message string, pbResponse []*response.TransferMonthAmountResponse) *pb.ApiResponseTransferMonthAmount {
	return &pb.ApiResponseTransferMonthAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransferMonthAmounts(pbResponse),
	}
}

// ToProtoResponseTransferYearAmount maps a status, message, and a list of TransferYearAmountResponse
// to an ApiResponseTransferYearAmount proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferYearAmountResponse to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferYearAmount containing the mapped data.
func (m *transferStatsAmountProtoMapper) ToProtoResponseTransferYearAmount(status string, message string, pbResponse []*response.TransferYearAmountResponse) *pb.ApiResponseTransferYearAmount {
	return &pb.ApiResponseTransferYearAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransferYearAmounts(pbResponse),
	}
}

// mapResponseTransferMonthAmount maps a TransferMonthAmountResponse to its protobuf representation.
//
// Args:
//   - s: The TransferMonthAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to a TransferMonthAmountResponse protobuf message containing the mapped data.
func (m *transferStatsAmountProtoMapper) mapResponseTransferMonthAmount(s *response.TransferMonthAmountResponse) *pb.TransferMonthAmountResponse {
	return &pb.TransferMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTransferMonthAmounts maps a slice of TransferMonthAmountResponse
// to a slice of TransferMonthAmountResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferMonthAmountResponse
// to its corresponding protobuf representation using the mapResponseTransferMonthAmount method.
//
// Args:
//   - s: A slice of TransferMonthAmountResponse to be converted.
//
// Returns:
//   - A slice of TransferMonthAmountResponse containing the mapped protobuf data.
func (m *transferStatsAmountProtoMapper) mapResponseTransferMonthAmounts(s []*response.TransferMonthAmountResponse) []*pb.TransferMonthAmountResponse {
	var protoResponses []*pb.TransferMonthAmountResponse
	for _, transfer := range s {
		protoResponses = append(protoResponses, m.mapResponseTransferMonthAmount(transfer))
	}
	return protoResponses
}

// mapResponseTransferYearAmount maps a TransferYearAmountResponse to its protobuf representation.
//
// Args:
//   - s: The TransferYearAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to a TransferYearAmountResponse protobuf message containing the mapped data.
func (m *transferStatsAmountProtoMapper) mapResponseTransferYearAmount(s *response.TransferYearAmountResponse) *pb.TransferYearAmountResponse {
	return &pb.TransferYearAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTransferYearAmounts maps a slice of TransferYearAmountResponse
// to a slice of TransferYearAmountResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferYearAmountResponse
// to its corresponding protobuf representation using the mapResponseTransferYearAmount method.
//
// Args:
//   - s: A slice of TransferYearAmountResponse to be converted.
//
// Returns:
//   - A slice of TransferYearAmountResponse containing the mapped protobuf data.
func (m *transferStatsAmountProtoMapper) mapResponseTransferYearAmounts(s []*response.TransferYearAmountResponse) []*pb.TransferYearAmountResponse {
	var protoResponses []*pb.TransferYearAmountResponse
	for _, transfer := range s {
		protoResponses = append(protoResponses, m.mapResponseTransferYearAmount(transfer))
	}
	return protoResponses
}
