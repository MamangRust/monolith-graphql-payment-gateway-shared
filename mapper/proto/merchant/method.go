package merchantprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type merchantStatsMethodProtoMapper struct {
}

func NewMerchantStatsMethodProtoMapper() MerchantStatsMethodProtoMapper {
	return &merchantStatsMethodProtoMapper{}
}

// ToProtoResponseMonthlyPaymentMethods maps a status, message and a list of *response.MerchantResponseMonthlyPaymentMethod
// to a *pb.ApiResponseMerchantMonthlyPaymentMethod proto response.
//
// It is used to generate the response for the MerchantService.GetMonthlyPaymentMethod rpc method.
func (m *merchantStatsMethodProtoMapper) ToProtoResponseMonthlyPaymentMethods(status string, message string, ms []*response.MerchantResponseMonthlyPaymentMethod) *pb.ApiResponseMerchantMonthlyPaymentMethod {
	return &pb.ApiResponseMerchantMonthlyPaymentMethod{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesMonthlyPaymentMethod(ms),
	}
}

// ToProtoResponseYearlyPaymentMethods maps a status, message and a list of *response.MerchantResponseYearlyPaymentMethod
// to a *pb.ApiResponseMerchantYearlyPaymentMethod proto response.
//
// It is used to generate the response for the MerchantService.GetYearlyPaymentMethod rpc method.
func (m *merchantStatsMethodProtoMapper) ToProtoResponseYearlyPaymentMethods(status string, message string, ms []*response.MerchantResponseYearlyPaymentMethod) *pb.ApiResponseMerchantYearlyPaymentMethod {
	return &pb.ApiResponseMerchantYearlyPaymentMethod{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesYearlyPaymentMethod(ms),
	}
}

// mapResponseMonthlyPaymentMethod maps a *response.MerchantResponseMonthlyPaymentMethod to a *pb.MerchantResponseMonthlyPaymentMethod proto message.
//
// It is used to generate the response for the MerchantService.GetMonthlyPaymentMethod rpc method.
func (m *merchantStatsMethodProtoMapper) mapResponseMonthlyPaymentMethod(ms *response.MerchantResponseMonthlyPaymentMethod) *pb.MerchantResponseMonthlyPaymentMethod {
	return &pb.MerchantResponseMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int64(ms.TotalAmount),
	}
}

// mapResponsesMonthlyPaymentMethod maps a list of *response.MerchantResponseMonthlyPaymentMethod to a list of
// *pb.MerchantResponseMonthlyPaymentMethod proto responses.
//
// It iterates over each MerchantResponseMonthlyPaymentMethod in the input slice, converting
// them to their protobuf equivalent using the mapResponseMonthlyPaymentMethod function. This
// function is used to generate the response data for monthly payment method RPC methods.
func (s *merchantStatsMethodProtoMapper) mapResponsesMonthlyPaymentMethod(roles []*response.MerchantResponseMonthlyPaymentMethod) []*pb.MerchantResponseMonthlyPaymentMethod {
	var responseRoles []*pb.MerchantResponseMonthlyPaymentMethod

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseMonthlyPaymentMethod(role))
	}

	return responseRoles
}

// mapResponseYearlyPaymentMethod maps a *response.MerchantResponseYearlyPaymentMethod to a
// *pb.MerchantResponseYearlyPaymentMethod proto message.
//
// It is used to generate the response for the MerchantService.GetYearlyPaymentMethod rpc method.
func (m *merchantStatsMethodProtoMapper) mapResponseYearlyPaymentMethod(ms *response.MerchantResponseYearlyPaymentMethod) *pb.MerchantResponseYearlyPaymentMethod {
	return &pb.MerchantResponseYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int64(ms.TotalAmount),
	}
}

// mapResponsesYearlyPaymentMethod maps a list of *response.MerchantResponseYearlyPaymentMethod to a list of
// *pb.MerchantResponseYearlyPaymentMethod proto responses.
//
// It iterates over each MerchantResponseYearlyPaymentMethod in the input slice, converting
// them to their protobuf equivalent using the mapResponseYearlyPaymentMethod function. This
// function is used to generate the response data for yearly payment method RPC methods.
func (s *merchantStatsMethodProtoMapper) mapResponsesYearlyPaymentMethod(roles []*response.MerchantResponseYearlyPaymentMethod) []*pb.MerchantResponseYearlyPaymentMethod {
	var responseRoles []*pb.MerchantResponseYearlyPaymentMethod

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseYearlyPaymentMethod(role))
	}

	return responseRoles
}
