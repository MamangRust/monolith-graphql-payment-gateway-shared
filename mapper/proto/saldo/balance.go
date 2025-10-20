package saldoprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
)

type saldoStatsBalanceProtoMapper struct {
}

func NewSaldoStatsBalanceProtoMapper() SaldoStatsBalanceProtoMapper {
	return &saldoStatsBalanceProtoMapper{}
}

// ToProtoResponseMonthSaldoBalances maps a status, message and a list of *response.SaldoMonthBalanceResponse
// to a *pb.ApiResponseMonthSaldoBalances proto response.
//
// It is used to generate the response for the SaldoService.GetMonthlySaldoBalances rpc method.
func (s *saldoStatsBalanceProtoMapper) ToProtoResponseMonthSaldoBalances(status string, message string, pbResponse []*response.SaldoMonthBalanceResponse) *pb.ApiResponseMonthSaldoBalances {
	return &pb.ApiResponseMonthSaldoBalances{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoMonthBalanceResponses(pbResponse),
	}
}

// ToProtoResponseYearSaldoBalances maps a status, message, and a list of
// *response.SaldoYearBalanceResponse to a *pb.ApiResponseYearSaldoBalances proto response.
//
// It is used to generate the response for the SaldoService.GetYearlySaldoBalances rpc method.
//
// Args:
//
//	status: The status of the response.
//	message: The message accompanying the response.
//	pbResponse: A slice of SaldoYearBalanceResponse to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseYearSaldoBalances containing the mapped data.
func (s *saldoStatsBalanceProtoMapper) ToProtoResponseYearSaldoBalances(status string, message string, pbResponse []*response.SaldoYearBalanceResponse) *pb.ApiResponseYearSaldoBalances {
	return &pb.ApiResponseYearSaldoBalances{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoYearBalanceResponses(pbResponse),
	}
}

// mapSaldoMonthBalanceResponse maps a *response.SaldoMonthBalanceResponse to a *pb.SaldoMonthBalanceResponse proto response.
//
// It takes a SaldoMonthBalanceResponse from the response domain model and converts it
// into its protobuf representation, ensuring that the month and total balance fields are accurately mapped.
//
// Args:
//
//	ss: The SaldoMonthBalanceResponse to be converted.
//
// Returns:
//
//	A pointer to a SaldoMonthBalanceResponse containing the mapped data.
func (s *saldoStatsBalanceProtoMapper) mapSaldoMonthBalanceResponse(ss *response.SaldoMonthBalanceResponse) *pb.SaldoMonthBalanceResponse {
	return &pb.SaldoMonthBalanceResponse{
		Month:        ss.Month,
		TotalBalance: int32(ss.TotalBalance),
	}
}

// mapSaldoMonthBalanceResponses maps a list of *response.SaldoMonthBalanceResponse
// to a list of *pb.SaldoMonthBalanceResponse proto responses.
//
// It iterates over each SaldoMonthBalanceResponse in the input slice, converting
// them to their protobuf equivalent using the mapSaldoMonthBalanceResponse function.
// This function is used to generate the response data for monthly balance RPC methods.
func (s *saldoStatsBalanceProtoMapper) mapSaldoMonthBalanceResponses(ss []*response.SaldoMonthBalanceResponse) []*pb.SaldoMonthBalanceResponse {
	var saldoProtos []*pb.SaldoMonthBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoMonthBalanceResponse(saldo))
	}
	return saldoProtos
}

// mapSaldoYearBalanceResponse maps a *response.SaldoYearBalanceResponse to a *pb.SaldoYearBalanceResponse proto response.
//
// It takes a SaldoYearBalanceResponse from the response domain model and converts it
// into its protobuf representation, ensuring that the year and total balance fields are accurately mapped.
//
// Args:
//
//	ss: The SaldoYearBalanceResponse to be converted.
//
// Returns:
//
//	A pointer to SaldoYearBalanceResponse containing the mapped data.
func (s *saldoStatsBalanceProtoMapper) mapSaldoYearBalanceResponse(ss *response.SaldoYearBalanceResponse) *pb.SaldoYearBalanceResponse {
	return &pb.SaldoYearBalanceResponse{
		Year:         ss.Year,
		TotalBalance: int32(ss.TotalBalance),
	}
}

// mapSaldoYearBalanceResponses maps a list of *response.SaldoYearBalanceResponse
// to a list of *pb.SaldoYearBalanceResponse proto responses.
//
// It iterates over each SaldoYearBalanceResponse in the input slice, converting
// them to their protobuf equivalent using the mapSaldoYearBalanceResponse function.
// This function is used to generate the response data for yearly balance RPC methods.
func (s *saldoStatsBalanceProtoMapper) mapSaldoYearBalanceResponses(ss []*response.SaldoYearBalanceResponse) []*pb.SaldoYearBalanceResponse {
	var saldoProtos []*pb.SaldoYearBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoYearBalanceResponse(saldo))
	}
	return saldoProtos
}
