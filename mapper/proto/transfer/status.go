package transferprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer"
)

type transferStatsStatusProtoMapper struct {
}

func NewTransferStatsStatusProtoMapper() TransferStatsStatusProtoMapper {
	return &transferStatsStatusProtoMapper{}
}

// ToProtoResponseTransferMonthStatusSuccess maps a status, message, and a list of TransferResponseMonthStatusSuccess
// to an ApiResponseTransferMonthStatusSuccess proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponseMonthStatusSuccess to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferMonthStatusSuccess containing the mapped data.
func (m *transferStatsStatusProtoMapper) ToProtoResponseTransferMonthStatusSuccess(status string, message string, pbResponse []*response.TransferResponseMonthStatusSuccess) *pb.ApiResponseTransferMonthStatusSuccess {
	return &pb.ApiResponseTransferMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransferMonthStatusSuccess(pbResponse),
	}
}

// ToProtoResponseTransferYearStatusSuccess maps a status, message, and a list of TransferResponseYearStatusSuccess
// to an ApiResponseTransferYearStatusSuccess proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponseYearStatusSuccess to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferYearStatusSuccess containing the mapped data.
func (m *transferStatsStatusProtoMapper) ToProtoResponseTransferYearStatusSuccess(status string, message string, pbResponse []*response.TransferResponseYearStatusSuccess) *pb.ApiResponseTransferYearStatusSuccess {
	return &pb.ApiResponseTransferYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapTransferResponsesYearStatusSuccess(pbResponse),
	}
}

// ToProtoResponseTransferMonthStatusFailed maps a status, message, and a list of TransferResponseMonthStatusFailed
// to an ApiResponseTransferMonthStatusFailed proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponseMonthStatusFailed to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferMonthStatusFailed containing the mapped data.
func (m *transferStatsStatusProtoMapper) ToProtoResponseTransferMonthStatusFailed(status string, message string, pbResponse []*response.TransferResponseMonthStatusFailed) *pb.ApiResponseTransferMonthStatusFailed {
	return &pb.ApiResponseTransferMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransferMonthStatusFailed(pbResponse),
	}
}

// ToProtoResponseTransferYearStatusFailed maps a status, message, and a list of TransferResponseYearStatusFailed
// to an ApiResponseTransferYearStatusFailed proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponseYearStatusFailed to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferYearStatusFailed containing the mapped data.
func (m *transferStatsStatusProtoMapper) ToProtoResponseTransferYearStatusFailed(status string, message string, pbResponse []*response.TransferResponseYearStatusFailed) *pb.ApiResponseTransferYearStatusFailed {
	return &pb.ApiResponseTransferYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapTransferResponsesYearStatusFailed(pbResponse),
	}
}

// mapResponseTransferMonthStatusSuccess maps a TransferResponseMonthStatusSuccess to its protobuf representation.
//
// Args:
//   - s: The TransferResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a TransferMonthStatusSuccessResponse protobuf message containing the mapped data.
func (t *transferStatsStatusProtoMapper) mapResponseTransferMonthStatusSuccess(s *response.TransferResponseMonthStatusSuccess) *pb.TransferMonthStatusSuccessResponse {
	return &pb.TransferMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapResponsesTransferMonthStatusSuccess maps a slice of TransferResponseMonthStatusSuccess
// to a slice of TransferMonthStatusSuccessResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferResponseMonthStatusSuccess
// to its corresponding protobuf representation using the mapResponseTransferMonthStatusSuccess method.
//
// Args:
//   - Transfers: A slice of TransferResponseMonthStatusSuccess to be converted.
//
// Returns:
//   - A slice of TransferMonthStatusSuccessResponse containing the mapped protobuf data.
func (t *transferStatsStatusProtoMapper) mapResponsesTransferMonthStatusSuccess(Transfers []*response.TransferResponseMonthStatusSuccess) []*pb.TransferMonthStatusSuccessResponse {
	var TransferRecords []*pb.TransferMonthStatusSuccessResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapResponseTransferMonthStatusSuccess(Transfer))
	}

	return TransferRecords
}

// mapTransferResponseYearStatusSuccess maps a TransferResponseYearStatusSuccess to its protobuf representation.
//
// Args:
//   - s: The TransferResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a TransferYearStatusSuccessResponse protobuf message containing the mapped data.
func (t *transferStatsStatusProtoMapper) mapTransferResponseYearStatusSuccess(s *response.TransferResponseYearStatusSuccess) *pb.TransferYearStatusSuccessResponse {
	return &pb.TransferYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapTransferResponsesYearStatusSuccess maps a slice of TransferResponseYearStatusSuccess
// to a slice of TransferYearStatusSuccessResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferResponseYearStatusSuccess
// to its corresponding protobuf representation using the mapTransferResponseYearStatusSuccess method.
//
// Args:
//   - Transfers: A slice of TransferResponseYearStatusSuccess to be converted.
//
// Returns:
//   - A slice of TransferYearStatusSuccessResponse containing the mapped protobuf data.
func (t *transferStatsStatusProtoMapper) mapTransferResponsesYearStatusSuccess(Transfers []*response.TransferResponseYearStatusSuccess) []*pb.TransferYearStatusSuccessResponse {
	var TransferRecords []*pb.TransferYearStatusSuccessResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapTransferResponseYearStatusSuccess(Transfer))
	}

	return TransferRecords
}

// mapResponseTransferMonthStatusFailed maps a TransferResponseMonthStatusFailed to its protobuf representation.
//
// Args:
//   - s: The TransferResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a TransferMonthStatusFailedResponse protobuf message containing the mapped data.
func (t *transferStatsStatusProtoMapper) mapResponseTransferMonthStatusFailed(s *response.TransferResponseMonthStatusFailed) *pb.TransferMonthStatusFailedResponse {
	return &pb.TransferMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponsesTransferMonthStatusFailed maps a slice of TransferResponseMonthStatusFailed
// to a slice of TransferMonthStatusFailedResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferResponseMonthStatusFailed
// to its corresponding protobuf representation using the mapResponseTransferMonthStatusFailed method.
//
// Args:
//   - Transfers: A slice of TransferResponseMonthStatusFailed to be converted.
//
// Returns:
//   - A slice of TransferMonthStatusFailedResponse containing the mapped protobuf data.
func (t *transferStatsStatusProtoMapper) mapResponsesTransferMonthStatusFailed(Transfers []*response.TransferResponseMonthStatusFailed) []*pb.TransferMonthStatusFailedResponse {
	var TransferRecords []*pb.TransferMonthStatusFailedResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapResponseTransferMonthStatusFailed(Transfer))
	}

	return TransferRecords
}

// mapTransferResponseYearStatusFailed maps a TransferResponseYearStatusFailed to its protobuf representation.
//
// Args:
//   - s: The TransferResponseYearStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a TransferYearStatusFailedResponse protobuf message containing the mapped data.
func (t *transferStatsStatusProtoMapper) mapTransferResponseYearStatusFailed(s *response.TransferResponseYearStatusFailed) *pb.TransferYearStatusFailedResponse {
	return &pb.TransferYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapTransferResponsesYearStatusFailed maps a slice of TransferResponseYearStatusFailed
// to a slice of TransferYearStatusFailedResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferResponseYearStatusFailed
// to its corresponding protobuf representation using the mapTransferResponseYearStatusFailed method.
//
// Args:
//   - Transfers: A slice of TransferResponseYearStatusFailed to be converted.
//
// Returns:
//   - A slice of TransferYearStatusFailedResponse containing the mapped protobuf data.
func (t *transferStatsStatusProtoMapper) mapTransferResponsesYearStatusFailed(Transfers []*response.TransferResponseYearStatusFailed) []*pb.TransferYearStatusFailedResponse {
	var TransferRecords []*pb.TransferYearStatusFailedResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapTransferResponseYearStatusFailed(Transfer))
	}

	return TransferRecords
}
