package merchantprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
)

type merchantStatsAmountProtoMapper struct{}

func NewMerchantStatsAmountProtoMapper() MerchantStatsAmountProtoMapper {
	return &merchantStatsAmountProtoMapper{}
}

// ToProtoResponseMonthlyAmounts maps a status, message and a list of *response.MerchantResponseMonthlyAmount
// to a *pb.ApiResponseMerchantMonthlyAmount proto response.
//
// It is used to generate the response for the MerchantService.GetMonthlyAmount rpc method.
func (m *merchantStatsAmountProtoMapper) ToProtoResponseMonthlyAmounts(status string, message string, ms []*response.MerchantResponseMonthlyAmount) *pb.ApiResponseMerchantMonthlyAmount {
	return &pb.ApiResponseMerchantMonthlyAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesMonthlyAmount(ms),
	}
}

// ToProtoResponseYearlyAmounts maps a status, message and a list of *response.MerchantResponseYearlyAmount
// to a *pb.ApiResponseMerchantYearlyAmount proto response.
//
// It is used to generate the response for the MerchantService.GetYearlyAmount rpc method.
func (m *merchantStatsAmountProtoMapper) ToProtoResponseYearlyAmounts(status string, message string, ms []*response.MerchantResponseYearlyAmount) *pb.ApiResponseMerchantYearlyAmount {
	return &pb.ApiResponseMerchantYearlyAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesYearlyAmount(ms),
	}
}

// mapResponseMonthlyAmount maps a *response.MerchantResponseMonthlyAmount to a *pb.MerchantResponseMonthlyAmount proto message.
//
// It is used to generate the response for the MerchantService.GetMonthlyAmount rpc method.
func (m *merchantStatsAmountProtoMapper) mapResponseMonthlyAmount(ms *response.MerchantResponseMonthlyAmount) *pb.MerchantResponseMonthlyAmount {
	return &pb.MerchantResponseMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int64(ms.TotalAmount),
	}
}

// mapResponsesMonthlyAmount maps a list of *response.MerchantResponseMonthlyAmount to a list of
// *pb.MerchantResponseMonthlyAmount proto responses.
//
// It iterates over each MerchantResponseMonthlyAmount in the input slice, converting
// them to their protobuf equivalent using the mapResponseMonthlyAmount function. This
// function is used to generate the response data for monthly amount RPC methods.
func (s *merchantStatsAmountProtoMapper) mapResponsesMonthlyAmount(roles []*response.MerchantResponseMonthlyAmount) []*pb.MerchantResponseMonthlyAmount {
	var responseRoles []*pb.MerchantResponseMonthlyAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseMonthlyAmount(role))
	}

	return responseRoles
}

// mapResponseYearlyAmount maps a *response.MerchantResponseYearlyAmount to a *pb.MerchantResponseYearlyAmount proto message.
//
// It is used to generate the response for the MerchantService.GetYearlyAmount rpc method.
func (m *merchantStatsAmountProtoMapper) mapResponseYearlyAmount(ms *response.MerchantResponseYearlyAmount) *pb.MerchantResponseYearlyAmount {
	return &pb.MerchantResponseYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int64(ms.TotalAmount),
	}
}

// mapResponsesYearlyAmount maps a list of *response.MerchantResponseYearlyAmount to a list of
// *pb.MerchantResponseYearlyAmount proto responses.
//
// It iterates over each MerchantResponseYearlyAmount in the input slice, converting
// them to their protobuf equivalent using the mapResponseYearlyAmount function.
// This function is used to generate the response data for yearly amount RPC methods.
func (s *merchantStatsAmountProtoMapper) mapResponsesYearlyAmount(roles []*response.MerchantResponseYearlyAmount) []*pb.MerchantResponseYearlyAmount {
	var responseRoles []*pb.MerchantResponseYearlyAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseYearlyAmount(role))
	}

	return responseRoles
}
