package withdrawprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	protomapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto"
	pbhelpers "github.com/MamangRust/monolith-payment-gateway-pb"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type withdrawQueryProtoMapper struct {
}

func NewWithdrawQueryProtoMapper() *withdrawQueryProtoMapper {
	return &withdrawQueryProtoMapper{}
}

// ToProtoResponseWithdraw maps a status, message, and a WithdrawResponse
// to an ApiResponseWithdraw proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - withdraw: The WithdrawResponse that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdraw containing the mapped data.
func (m *withdrawQueryProtoMapper) ToProtoResponseWithdraw(status string, message string, withdraw *response.WithdrawResponse) *pb.ApiResponseWithdraw {
	return &pb.ApiResponseWithdraw{
		Status:  status,
		Message: message,
		Data:    m.mapResponseWithdrawal(withdraw),
	}
}

// ToProtoResponsePaginationWithdraw maps a status, message, a pagination meta, and a list of WithdrawResponse
// to an ApiResponsePaginationWithdraw proto message.
//
// Args:
//   - pagination: The pagination meta that needs to be converted.
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - withdraws: The list of WithdrawResponse that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponsePaginationWithdraw containing the mapped data.
func (m *withdrawQueryProtoMapper) ToProtoResponsePaginationWithdraw(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponse) *pb.ApiResponsePaginationWithdraw {
	return &pb.ApiResponsePaginationWithdraw{
		Status:     status,
		Message:    message,
		Data:       m.mapResponsesWithdrawal(pbResponse),
		Pagination: protomapper.MapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationWithdrawDeleteAt maps a pagination meta, status, message, and a list of WithdrawResponseDeleteAt
// to an ApiResponsePaginationWithdrawDeleteAt proto message.
//
// Args:
//   - pagination: The pagination meta that needs to be converted.
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - withdraws: The list of WithdrawResponseDeleteAt that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponsePaginationWithdrawDeleteAt containing the mapped data.
func (m *withdrawQueryProtoMapper) ToProtoResponsePaginationWithdrawDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponseDeleteAt) *pb.ApiResponsePaginationWithdrawDeleteAt {
	return &pb.ApiResponsePaginationWithdrawDeleteAt{
		Status:     status,
		Message:    message,
		Data:       m.mapResponsesWithdrawalDeleteAt(pbResponse),
		Pagination: protomapper.MapPaginationMeta(pagination),
	}
}

// mapResponseWithdrawal maps a single WithdrawResponse to its protobuf representation.
//
// Args:
//   - withdraw: The WithdrawResponse that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponse protobuf message containing the mapped data.
func (w *withdrawQueryProtoMapper) mapResponseWithdrawal(withdraw *response.WithdrawResponse) *pb.WithdrawResponse {
	return &pb.WithdrawResponse{
		WithdrawId:     int32(withdraw.ID),
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int32(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
	}
}

// mapResponsesWithdrawal maps a slice of WithdrawResponse to a slice of protobuf WithdrawResponse.
//
// It takes a slice of WithdrawResponse as input and returns a slice of corresponding protobuf WithdrawResponse.
// The mapping includes fields like WithdrawId, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.
func (w *withdrawQueryProtoMapper) mapResponsesWithdrawal(withdraws []*response.WithdrawResponse) []*pb.WithdrawResponse {
	var responseWithdraws []*pb.WithdrawResponse

	for _, withdraw := range withdraws {
		responseWithdraws = append(responseWithdraws, w.mapResponseWithdrawal(withdraw))
	}

	return responseWithdraws
}

// mapResponseWithdrawalDeleteAt maps a single WithdrawResponseDeleteAt to its protobuf representation.
//
// Args:
//   - withdraw: The WithdrawResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseDeleteAt protobuf message containing the mapped data.
func (w *withdrawQueryProtoMapper) mapResponseWithdrawalDeleteAt(withdraw *response.WithdrawResponseDeleteAt) *pb.WithdrawResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if withdraw.DeletedAt != nil {
		deletedAt = wrapperspb.String(*withdraw.DeletedAt)
	}

	return &pb.WithdrawResponseDeleteAt{
		WithdrawId:     int32(withdraw.ID),
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int32(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
		DeletedAt:      deletedAt,
	}
}

// mapResponsesWithdrawalDeleteAt maps a slice of WithdrawResponseDeleteAt to a slice of protobuf WithdrawResponseDeleteAt.
//
// It takes a slice of WithdrawResponseDeleteAt as input and returns a slice of corresponding protobuf WithdrawResponseDeleteAt.
// The mapping includes fields like WithdrawId, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (w *withdrawQueryProtoMapper) mapResponsesWithdrawalDeleteAt(withdraws []*response.WithdrawResponseDeleteAt) []*pb.WithdrawResponseDeleteAt {
	var responseWithdraws []*pb.WithdrawResponseDeleteAt

	for _, withdraw := range withdraws {
		responseWithdraws = append(responseWithdraws, w.mapResponseWithdrawalDeleteAt(withdraw))
	}

	return responseWithdraws
}
