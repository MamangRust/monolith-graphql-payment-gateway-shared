package withdrawprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
)

type withdrawStatsStatusProtoMapper struct {
}

func NewWithdrawStatsStatusProtoMapper() *withdrawStatsStatusProtoMapper {
	return &withdrawStatsStatusProtoMapper{}
}

// ToProtoResponseWithdrawMonthStatusSuccess maps a status, message, and a list of WithdrawResponseMonthStatusSuccess
// to an ApiResponseWithdrawMonthStatusSuccess proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawMonthStatusSuccess containing the mapped data.
func (m *withdrawStatsStatusProtoMapper) ToProtoResponseWithdrawMonthStatusSuccess(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusSuccess) *pb.ApiResponseWithdrawMonthStatusSuccess {
	return &pb.ApiResponseWithdrawMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesWithdrawMonthStatusSuccess(pbResponse),
	}
}

// ToProtoResponseWithdrawYearStatusSuccess maps a status, message, and a list of WithdrawResponseYearStatusSuccess
// to an ApiResponseWithdrawYearStatusSuccess proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawYearStatusSuccess containing the mapped data.
func (m *withdrawStatsStatusProtoMapper) ToProtoResponseWithdrawYearStatusSuccess(status string, message string, pbResponse []*response.WithdrawResponseYearStatusSuccess) *pb.ApiResponseWithdrawYearStatusSuccess {
	return &pb.ApiResponseWithdrawYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapWithdrawResponsesYearStatusSuccess(pbResponse),
	}
}

// ToProtoResponseWithdrawMonthStatusFailed maps a status, message, and a list of WithdrawResponseMonthStatusFailed
// to an ApiResponseWithdrawMonthStatusFailed proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawMonthStatusFailed containing the mapped data.
func (m *withdrawStatsStatusProtoMapper) ToProtoResponseWithdrawMonthStatusFailed(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusFailed) *pb.ApiResponseWithdrawMonthStatusFailed {
	return &pb.ApiResponseWithdrawMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesWithdrawMonthStatusFailed(pbResponse),
	}
}

// ToProtoResponseWithdrawYearStatusFailed maps a status, message, and a list of WithdrawResponseYearStatusFailed
// to an ApiResponseWithdrawYearStatusFailed proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawResponseYearStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawYearStatusFailed containing the mapped data.
func (m *withdrawStatsStatusProtoMapper) ToProtoResponseWithdrawYearStatusFailed(status string, message string, pbResponse []*response.WithdrawResponseYearStatusFailed) *pb.ApiResponseWithdrawYearStatusFailed {
	return &pb.ApiResponseWithdrawYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapWithdrawResponsesYearStatusFailed(pbResponse),
	}
}

// mapResponseWithdrawMonthStatusSuccess maps a single WithdrawResponseMonthStatusSuccess to its protobuf representation.
//
// Args:
//   - s: The WithdrawResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawMonthStatusSuccessResponse protobuf message containing the mapped data.
func (t *withdrawStatsStatusProtoMapper) mapResponseWithdrawMonthStatusSuccess(s *response.WithdrawResponseMonthStatusSuccess) *pb.WithdrawMonthStatusSuccessResponse {
	return &pb.WithdrawMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapResponsesWithdrawMonthStatusSuccess maps a slice of WithdrawResponseMonthStatusSuccess to a slice of protobuf WithdrawMonthStatusSuccessResponse.
//
// It takes a slice of WithdrawResponseMonthStatusSuccess as input and returns a slice of corresponding protobuf WithdrawMonthStatusSuccessResponse.
// The mapping includes fields like Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawStatsStatusProtoMapper) mapResponsesWithdrawMonthStatusSuccess(Withdraws []*response.WithdrawResponseMonthStatusSuccess) []*pb.WithdrawMonthStatusSuccessResponse {
	var WithdrawRecords []*pb.WithdrawMonthStatusSuccessResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapResponseWithdrawMonthStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// mapWithdrawResponseYearStatusSuccess maps a single WithdrawResponseYearStatusSuccess to a WithdrawYearStatusSuccessResponse proto message.
//
// Args:
//   - s: The WithdrawResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawYearStatusSuccessResponse containing the mapped data, with fields Year, TotalSuccess, and TotalAmount.
func (t *withdrawStatsStatusProtoMapper) mapWithdrawResponseYearStatusSuccess(s *response.WithdrawResponseYearStatusSuccess) *pb.WithdrawYearStatusSuccessResponse {
	return &pb.WithdrawYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapWithdrawResponsesYearStatusSuccess maps a slice of WithdrawResponseYearStatusSuccess to a slice of protobuf WithdrawYearStatusSuccessResponse.
//
// It takes a slice of WithdrawResponseYearStatusSuccess as input and returns a slice of corresponding protobuf WithdrawYearStatusSuccessResponse.
// The mapping includes fields like Year, TotalSuccess, and TotalAmount.
func (t *withdrawStatsStatusProtoMapper) mapWithdrawResponsesYearStatusSuccess(Withdraws []*response.WithdrawResponseYearStatusSuccess) []*pb.WithdrawYearStatusSuccessResponse {
	var WithdrawRecords []*pb.WithdrawYearStatusSuccessResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapWithdrawResponseYearStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// mapResponseWithdrawMonthStatusFailed maps a single WithdrawResponseMonthStatusFailed to its protobuf representation.
//
// Args:
//   - s: The WithdrawResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawMonthStatusFailedResponse containing the mapped data, with fields Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawStatsStatusProtoMapper) mapResponseWithdrawMonthStatusFailed(s *response.WithdrawResponseMonthStatusFailed) *pb.WithdrawMonthStatusFailedResponse {
	return &pb.WithdrawMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponsesWithdrawMonthStatusFailed maps a slice of WithdrawResponseMonthStatusFailed to a slice of protobuf WithdrawMonthStatusFailedResponse.
//
// It takes a slice of WithdrawResponseMonthStatusFailed as input and returns a slice of corresponding protobuf WithdrawMonthStatusFailedResponse.
// The mapping includes fields like Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawStatsStatusProtoMapper) mapResponsesWithdrawMonthStatusFailed(Withdraws []*response.WithdrawResponseMonthStatusFailed) []*pb.WithdrawMonthStatusFailedResponse {
	var WithdrawRecords []*pb.WithdrawMonthStatusFailedResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapResponseWithdrawMonthStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// mapWithdrawResponseYearStatusFailed maps a WithdrawResponseYearStatusFailed to a WithdrawYearStatusFailedResponse proto message.
//
// Args:
//   - s: The WithdrawResponseYearStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawYearStatusFailedResponse containing the mapped data, with fields Year, TotalFailed, and TotalAmount.
func (t *withdrawStatsStatusProtoMapper) mapWithdrawResponseYearStatusFailed(s *response.WithdrawResponseYearStatusFailed) *pb.WithdrawYearStatusFailedResponse {
	return &pb.WithdrawYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapWithdrawResponsesYearStatusFailed maps a slice of WithdrawResponseYearStatusFailed to a slice of protobuf WithdrawYearStatusFailedResponse.
//
// It takes a slice of WithdrawResponseYearStatusFailed as input and returns a slice of corresponding protobuf WithdrawYearStatusFailedResponse.
// The mapping includes fields like Year, TotalFailed, and TotalAmount.
func (t *withdrawStatsStatusProtoMapper) mapWithdrawResponsesYearStatusFailed(Withdraws []*response.WithdrawResponseYearStatusFailed) []*pb.WithdrawYearStatusFailedResponse {
	var WithdrawRecords []*pb.WithdrawYearStatusFailedResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapWithdrawResponseYearStatusFailed(Withdraw))
	}

	return WithdrawRecords
}
