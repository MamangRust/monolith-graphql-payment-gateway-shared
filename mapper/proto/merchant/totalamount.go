package merchantprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type merchantStatsTotalAmount struct {
}

func NewMerchantStatsTotalAmount() MerchantStatsTotalAmountProtoMapper {
	return &merchantStatsTotalAmount{}
}

// ToProtoResponseMonthlyTotalAmounts maps a status, message and a list of *response.MerchantResponseMonthlyTotalAmount
// to a *pb.ApiResponseMerchantMonthlyTotalAmount proto response.
//
// It is used to generate the response for the MerchantService.GetMonthlyTotalAmount rpc method.
func (m *merchantStatsTotalAmount) ToProtoResponseMonthlyTotalAmounts(status string, message string, ms []*response.MerchantResponseMonthlyTotalAmount) *pb.ApiResponseMerchantMonthlyTotalAmount {
	return &pb.ApiResponseMerchantMonthlyTotalAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesMonthlyTotalAmount(ms),
	}
}

// ToProtoResponseYearlyTotalAmounts maps a status, message, and a list of *response.MerchantResponseYearlyTotalAmount
// to a *pb.ApiResponseMerchantYearlyTotalAmount proto response.
//
// It is used to generate the response for the MerchantService.GetYearlyTotalAmount rpc method.
func (m *merchantStatsTotalAmount) ToProtoResponseYearlyTotalAmounts(status string, message string, ms []*response.MerchantResponseYearlyTotalAmount) *pb.ApiResponseMerchantYearlyTotalAmount {
	return &pb.ApiResponseMerchantYearlyTotalAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesYearlyTotalAmount(ms),
	}
}

// mapResponseMonthlyTotalAmount maps a *response.MerchantResponseMonthlyTotalAmount to a *pb.MerchantResponseMonthlyTotalAmount proto message.
//
// It is used to generate the response for the MerchantService.GetMonthlyTotalAmount rpc method.
func (m *merchantStatsTotalAmount) mapResponseMonthlyTotalAmount(ms *response.MerchantResponseMonthlyTotalAmount) *pb.MerchantResponseMonthlyTotalAmount {
	return &pb.MerchantResponseMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: int64(ms.TotalAmount),
	}
}

// mapResponsesMonthlyTotalAmount maps a list of *response.MerchantResponseMonthlyTotalAmount to a list of
// *pb.MerchantResponseMonthlyTotalAmount proto responses.
//
// It iterates over each MerchantResponseMonthlyTotalAmount in the input slice, converting
// them to their protobuf equivalent using the mapResponseMonthlyTotalAmount function.
// This function is used to generate the response data for monthly total amount RPC methods.
func (s *merchantStatsTotalAmount) mapResponsesMonthlyTotalAmount(roles []*response.MerchantResponseMonthlyTotalAmount) []*pb.MerchantResponseMonthlyTotalAmount {
	var responseRoles []*pb.MerchantResponseMonthlyTotalAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseMonthlyTotalAmount(role))
	}

	return responseRoles
}

// mapResponseYearlyTotalAmount maps a *response.MerchantResponseYearlyTotalAmount to a
// *pb.MerchantResponseYearlyTotalAmount proto message.
//
// It is used to generate the response for the MerchantService.GetYearlyTotalAmount rpc method.
func (m *merchantStatsTotalAmount) mapResponseYearlyTotalAmount(ms *response.MerchantResponseYearlyTotalAmount) *pb.MerchantResponseYearlyTotalAmount {
	return &pb.MerchantResponseYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: int64(ms.TotalAmount),
	}
}

// mapResponsesYearlyTotalAmount maps a list of *response.MerchantResponseYearlyTotalAmount to a list of
// *pb.MerchantResponseYearlyTotalAmount proto responses.
//
// It iterates over each MerchantResponseYearlyTotalAmount in the input slice, converting
// them to their protobuf equivalent using the mapResponseYearlyTotalAmount function.
// This function is used to generate the response data for yearly total amount RPC methods.
func (s *merchantStatsTotalAmount) mapResponsesYearlyTotalAmount(roles []*response.MerchantResponseYearlyTotalAmount) []*pb.MerchantResponseYearlyTotalAmount {
	var responseRoles []*pb.MerchantResponseYearlyTotalAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseYearlyTotalAmount(role))
	}

	return responseRoles
}
