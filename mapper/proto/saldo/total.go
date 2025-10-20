package saldoprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
)

type saldoStatsTotalProtoMapper struct {
}

func NewSaldoStatsTotalProtoMapper() SaldoStatsTotalSaldoProtoMapper {
	return &saldoStatsTotalProtoMapper{}
}

// ToProtoResponseMonthTotalSaldo maps a status, message and a list of *response.SaldoMonthTotalBalanceResponse
// to a *pb.ApiResponseMonthTotalSaldo proto response.
//
// It is used to generate the response for the SaldoService.GetMonthlyTotalSaldoBalance rpc method.
func (s *saldoStatsTotalProtoMapper) ToProtoResponseMonthTotalSaldo(status string, message string, pbResponse []*response.SaldoMonthTotalBalanceResponse) *pb.ApiResponseMonthTotalSaldo {
	return &pb.ApiResponseMonthTotalSaldo{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoMonthTotalBalanceResponses(pbResponse),
	}
}

// ToProtoResponseYearTotalSaldo maps a status, message and a list of *response.SaldoYearTotalBalanceResponse
// to a *pb.ApiResponseYearTotalSaldo proto response.
//
// It is used to generate the response for the SaldoService.GetYearlyTotalSaldoBalance rpc method.
func (s *saldoStatsTotalProtoMapper) ToProtoResponseYearTotalSaldo(status string, message string, pbResponse []*response.SaldoYearTotalBalanceResponse) *pb.ApiResponseYearTotalSaldo {
	return &pb.ApiResponseYearTotalSaldo{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoYearTotalBalanceResponses(pbResponse),
	}
}

// mapSaldoMonthTotalBalanceResponse maps a *response.SaldoMonthTotalBalanceResponse to a *pb.SaldoMonthTotalBalanceResponse.
//
// It maps the fields of the input response to the corresponding fields of the output response.
// The TotalBalance field is converted to int32.
//
// Args:
//
//	ss: The SaldoMonthTotalBalanceResponse to be converted.
//
// Returns:
//
//	A pointer to a SaldoMonthTotalBalanceResponse containing the mapped data.
func (s *saldoStatsTotalProtoMapper) mapSaldoMonthTotalBalanceResponse(ss *response.SaldoMonthTotalBalanceResponse) *pb.SaldoMonthTotalBalanceResponse {
	totalBalance := 0

	if ss.TotalBalance != 0 {
		totalBalance = ss.TotalBalance
	}

	return &pb.SaldoMonthTotalBalanceResponse{
		Month:        ss.Month,
		Year:         ss.Year,
		TotalBalance: int32(totalBalance),
	}
}

// mapSaldoMonthTotalBalanceResponses maps a list of *response.SaldoMonthTotalBalanceResponse
// to a list of *pb.SaldoMonthTotalBalanceResponse proto responses.
//
// It iterates over each SaldoMonthTotalBalanceResponse in the input slice, converting
// them to their protobuf equivalent using the mapSaldoMonthTotalBalanceResponse function.
// This function is used to generate the response data for monthly total saldo balance RPC methods.
//
// Args:
//
//	ss: A slice of SaldoMonthTotalBalanceResponse to be converted.
//
// Returns:
//
//	A slice of SaldoMonthTotalBalanceResponse containing the mapped data.
func (s *saldoStatsTotalProtoMapper) mapSaldoMonthTotalBalanceResponses(ss []*response.SaldoMonthTotalBalanceResponse) []*pb.SaldoMonthTotalBalanceResponse {
	var saldoProtos []*pb.SaldoMonthTotalBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoMonthTotalBalanceResponse(saldo))
	}
	return saldoProtos
}

// mapSaldoYearTotalBalanceResponse maps a *response.SaldoYearTotalBalanceResponse to a *pb.SaldoYearTotalBalanceResponse proto response.
//
// It takes a SaldoYearTotalBalanceResponse from the response domain model and converts it
// into its protobuf representation, ensuring that the year and total balance fields are accurately mapped.
//
// Args:
//
//	ss: The SaldoYearTotalBalanceResponse to be converted.
//
// Returns:
//
//	A pointer to a SaldoYearTotalBalanceResponse containing the mapped data.
func (s *saldoStatsTotalProtoMapper) mapSaldoYearTotalBalanceResponse(ss *response.SaldoYearTotalBalanceResponse) *pb.SaldoYearTotalBalanceResponse {
	totalBalance := 0

	if ss.TotalBalance != 0 {
		totalBalance = ss.TotalBalance
	}

	return &pb.SaldoYearTotalBalanceResponse{
		Year:         ss.Year,
		TotalBalance: int32(totalBalance),
	}
}

// mapSaldoYearTotalBalanceResponses maps a list of *response.SaldoYearTotalBalanceResponse
// to a list of *pb.SaldoYearTotalBalanceResponse proto responses.
//
// It iterates over each SaldoYearTotalBalanceResponse in the input slice, converting
// them to their protobuf equivalent using the mapSaldoYearTotalBalanceResponse function.
// This function is used to generate the response data for yearly total saldo balance RPC methods.
func (s *saldoStatsTotalProtoMapper) mapSaldoYearTotalBalanceResponses(ss []*response.SaldoYearTotalBalanceResponse) []*pb.SaldoYearTotalBalanceResponse {
	var saldoProtos []*pb.SaldoYearTotalBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoYearTotalBalanceResponse(saldo))
	}
	return saldoProtos
}
