package withdrawprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
)

type withdrawStatsAmountProtoMapper struct {
}

func NewWithdrawStatsAmountMapper() WithdrawaStatsAmountProtoMapper {
	return &withdrawStatsAmountProtoMapper{}
}

// ToProtoResponseWithdrawMonthAmount maps a status, message, and a list of WithdrawMonthlyAmountResponse
// to an ApiResponseWithdrawMonthAmount proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawMonthlyAmountResponse that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawMonthAmount containing the mapped data.
func (m *withdrawStatsAmountProtoMapper) ToProtoResponseWithdrawMonthAmount(status string, message string, pbResponse []*response.WithdrawMonthlyAmountResponse) *pb.ApiResponseWithdrawMonthAmount {
	return &pb.ApiResponseWithdrawMonthAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseWithdrawMonthlyAmounts(pbResponse),
	}
}

// ToProtoResponseWithdrawYearAmount maps a status, message, and a list of WithdrawYearlyAmountResponse
// to an ApiResponseWithdrawYearAmount proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawYearlyAmountResponse that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawYearAmount containing the mapped data.
func (m *withdrawStatsAmountProtoMapper) ToProtoResponseWithdrawYearAmount(status string, message string, pbResponse []*response.WithdrawYearlyAmountResponse) *pb.ApiResponseWithdrawYearAmount {
	return &pb.ApiResponseWithdrawYearAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseWithdrawYearlyAmounts(pbResponse),
	}
}

// mapResponseWithdrawMonthlyAmount maps a WithdrawMonthlyAmountResponse to its protobuf representation.
//
// Args:
//   - s: The WithdrawMonthlyAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawMonthlyAmountResponse containing the mapped data, with fields Month and TotalAmount.
func (m *withdrawStatsAmountProtoMapper) mapResponseWithdrawMonthlyAmount(s *response.WithdrawMonthlyAmountResponse) *pb.WithdrawMonthlyAmountResponse {
	return &pb.WithdrawMonthlyAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseWithdrawMonthlyAmounts maps a slice of WithdrawMonthlyAmountResponse to a slice of protobuf WithdrawMonthlyAmountResponse.
//
// It takes a slice of WithdrawMonthlyAmountResponse as input and returns a slice of corresponding protobuf WithdrawMonthlyAmountResponse.
// The mapping includes fields like Month and TotalAmount.
func (m *withdrawStatsAmountProtoMapper) mapResponseWithdrawMonthlyAmounts(s []*response.WithdrawMonthlyAmountResponse) []*pb.WithdrawMonthlyAmountResponse {
	var protoResponses []*pb.WithdrawMonthlyAmountResponse
	for _, withdraw := range s {
		protoResponses = append(protoResponses, m.mapResponseWithdrawMonthlyAmount(withdraw))
	}
	return protoResponses
}

// mapResponseWithdrawYearlyAmount maps a WithdrawYearlyAmountResponse to its protobuf representation.
//
// Args:
//   - s: The WithdrawYearlyAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawYearlyAmountResponse containing the mapped data, with fields Year and TotalAmount.
func (m *withdrawStatsAmountProtoMapper) mapResponseWithdrawYearlyAmount(s *response.WithdrawYearlyAmountResponse) *pb.WithdrawYearlyAmountResponse {
	return &pb.WithdrawYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseWithdrawYearlyAmounts maps a slice of WithdrawYearlyAmountResponse to a slice of protobuf WithdrawYearlyAmountResponse.
//
// It takes a slice of WithdrawYearlyAmountResponse as input and returns a slice of corresponding protobuf WithdrawYearlyAmountResponse.
// The mapping includes fields like Year and TotalAmount.
func (m *withdrawStatsAmountProtoMapper) mapResponseWithdrawYearlyAmounts(s []*response.WithdrawYearlyAmountResponse) []*pb.WithdrawYearlyAmountResponse {
	var protoResponses []*pb.WithdrawYearlyAmountResponse
	for _, withdraw := range s {
		protoResponses = append(protoResponses, m.mapResponseWithdrawYearlyAmount(withdraw))
	}
	return protoResponses
}
