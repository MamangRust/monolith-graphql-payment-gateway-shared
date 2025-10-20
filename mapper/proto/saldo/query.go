package saldoprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	protomapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type saldoQueryProtoMapper struct {
}

func NewSaldoQueryProtoMapper() SaldoQueryProtoMapper {
	return &saldoQueryProtoMapper{}
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
func (s *saldoQueryProtoMapper) ToProtoResponseSaldo(status string, message string, pbResponse *response.SaldoResponse) *pb.ApiResponseSaldo {
	return &pb.ApiResponseSaldo{
		Status:  status,
		Message: message,
		Data:    s.mapResponseSaldo(pbResponse),
	}
}

// ToProtoResponsePaginationSaldo maps a status, message, pagination meta, and a list of *response.SaldoResponse
// to a *pb.ApiResponsePaginationSaldo proto response.
//
// It is used to generate the response for the SaldoService.ListSaldo rpc method.
//
// Args:
//
//	pagination: The pagination data for the paginated response.
//	status: The status of the response.
//	message: The message accompanying the response.
//	pbResponse: A slice of SaldoResponse to be converted.
//
// Returns:
//
//	A pointer to an ApiResponsePaginationSaldo containing the status, message, data, and pagination data.
func (s *saldoQueryProtoMapper) ToProtoResponsePaginationSaldo(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponse) *pb.ApiResponsePaginationSaldo {
	return &pb.ApiResponsePaginationSaldo{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesSaldo(pbResponse),
		Pagination: protomapper.MapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationSaldoDeleteAt creates a new instance of ApiResponsePaginationSaldoDeleteAt.
//
// It maps the provided pagination metadata, status, message, and a list of SaldoResponseDeleteAt
// to its protobuf representation.
//
// Args:
//
//	pagination: The pagination data for the paginated response.
//	status: The status of the API response.
//	message: A descriptive message of the API response.
//	pbResponse: The list of SaldoResponseDeleteAt that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponsePaginationSaldoDeleteAt containing the status, message, saldo data, and pagination data.
func (s *saldoQueryProtoMapper) ToProtoResponsePaginationSaldoDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponseDeleteAt) *pb.ApiResponsePaginationSaldoDeleteAt {
	return &pb.ApiResponsePaginationSaldoDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesSaldoDeleteAt(pbResponse),
		Pagination: protomapper.MapPaginationMeta(pagination),
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
func (s *saldoQueryProtoMapper) mapResponseSaldo(saldo *response.SaldoResponse) *pb.SaldoResponse {
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

// mapResponsesSaldo maps a list of SaldoResponse to a list of *pb.SaldoResponse proto responses.
//
// It iterates over each SaldoResponse in the input slice, converting
// them to their protobuf equivalent using the mapResponseSaldo function.
// This function is used to generate the response data for ListSaldo rpc methods.
func (s *saldoQueryProtoMapper) mapResponsesSaldo(saldos []*response.SaldoResponse) []*pb.SaldoResponse {
	var responseSaldos []*pb.SaldoResponse

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.mapResponseSaldo(saldo))
	}

	return responseSaldos
}

// mapResponseSaldoDeleteAt maps a single response.SaldoResponseDeleteAt to a single *pb.SaldoResponseDeleteAt proto response.
//
// Args:
//
//	saldo: The SaldoResponseDeleteAt to be converted.
//
// Returns:
//
//	A pointer to a SaldoResponseDeleteAt containing the mapped data.
func (s *saldoQueryProtoMapper) mapResponseSaldoDeleteAt(saldo *response.SaldoResponseDeleteAt) *pb.SaldoResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if saldo.DeletedAt != nil {
		deletedAt = wrapperspb.String(*saldo.DeletedAt)
	}

	return &pb.SaldoResponseDeleteAt{
		SaldoId:        int32(saldo.ID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime,
		WithdrawAmount: int32(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
		DeletedAt:      deletedAt,
	}
}

// mapResponsesSaldoDeleteAt maps a list of *response.SaldoResponseDeleteAt to a list of
// *pb.SaldoResponseDeleteAt proto responses.
//
// It iterates over each SaldoResponseDeleteAt in the input slice, converting
// them to their protobuf equivalent using the mapResponseSaldoDeleteAt function.
// This function is used to generate the response data for Saldo-related RPC methods
// involving deleted saldo records.
func (s *saldoQueryProtoMapper) mapResponsesSaldoDeleteAt(saldos []*response.SaldoResponseDeleteAt) []*pb.SaldoResponseDeleteAt {
	var responseSaldos []*pb.SaldoResponseDeleteAt

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.mapResponseSaldoDeleteAt(saldo))
	}

	return responseSaldos
}
