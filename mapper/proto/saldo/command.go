package saldoprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	helpersproto "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto/helpers"
)

type saldoCommandProtoMapper struct {
}

func NewSaldoCommandProtoMapper() SaldoCommandProtoMapper {
	return &saldoCommandProtoMapper{}
}

// ToProtoResponseSaldo maps a SaldoResponse to an ApiResponseSaldo proto message.
//
// Args:
//
//	status: The status of the response.
//	message: The message accompanying the response.
//	pbResponse: The SaldoResponse to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseSaldo containing the mapped data.
func (s *saldoCommandProtoMapper) ToProtoResponseSaldo(status string, message string, pbResponse *response.SaldoResponse) *pb.ApiResponseSaldo {
	return &pb.ApiResponseSaldo{
		Status:  status,
		Message: message,
		Data:    s.mapResponseSaldo(pbResponse),
	}
}

func (s *saldoCommandProtoMapper) ToProtoResponseSaldoDeleteAt(status string, message string, pbResponse *response.SaldoResponseDeleteAt) *pb.ApiResponseSaldoDeleteAt {
	return &pb.ApiResponseSaldoDeleteAt{
		Status:  status,
		Message: message,
		Data:    s.mapResponseSaldoDeleteAt(pbResponse),
	}
}

// ToProtoResponseSaldoDelete maps a status and message to an ApiResponseSaldoDelete proto response.
//
// It is used to generate the response for the SaldoService.DeleteSaldo rpc method.
func (s *saldoCommandProtoMapper) ToProtoResponseSaldoDelete(status string, message string) *pb.ApiResponseSaldoDelete {
	return &pb.ApiResponseSaldoDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseSaldoAll maps a status and message to an ApiResponseSaldoAll proto response.
//
// Args:
//
//	status: The status of the response.
//	message: The message accompanying the response.
//
// Returns:
//
//	A pointer to an ApiResponseSaldoAll containing the mapped data.
func (s *saldoCommandProtoMapper) ToProtoResponseSaldoAll(status string, message string) *pb.ApiResponseSaldoAll {
	return &pb.ApiResponseSaldoAll{
		Status:  status,
		Message: message,
	}
}

// mapResponseSaldo maps a SaldoResponse to its protobuf representation.
//
// Args:
//
//	saldo: The SaldoResponse to be converted.
//
// Returns:
//
//	A pointer to SaldoResponse containing the mapped data.
func (s *saldoCommandProtoMapper) mapResponseSaldo(saldo *response.SaldoResponse) *pb.SaldoResponse {
	return &pb.SaldoResponse{
		SaldoId:        int32(saldo.ID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime,
		WithdrawAmount: int32(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
	}
}

func (s *saldoCommandProtoMapper) mapResponseSaldoDeleteAt(saldo *response.SaldoResponseDeleteAt) *pb.SaldoResponseDeleteAt {
	res := &pb.SaldoResponseDeleteAt{
		SaldoId:        int32(saldo.ID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime,
		WithdrawAmount: int32(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
	}

	if saldo.DeletedAt != nil {
		res.DeletedAt = helpersproto.StringPtrToWrapper(saldo.DeletedAt)
	}

	return res
}
