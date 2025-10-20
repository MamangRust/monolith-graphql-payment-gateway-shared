package withdrawprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type withdrawCommandProtoMapper struct {
}

func NewWithdrawCommandProtoMapper() WithdrawCommandProtoMapper {
	return &withdrawCommandProtoMapper{}
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
func (m *withdrawCommandProtoMapper) ToProtoResponseWithdraw(status string, message string, withdraw *response.WithdrawResponse) *pb.ApiResponseWithdraw {
	return &pb.ApiResponseWithdraw{
		Status:  status,
		Message: message,
		Data:    m.mapResponseWithdrawal(withdraw),
	}
}

func (m *withdrawCommandProtoMapper) ToProtoResponseWithdrawDeleteAt(status string, message string, withdraw *response.WithdrawResponseDeleteAt) *pb.ApiResponseWithdrawDeleteAt {
	return &pb.ApiResponseWithdrawDeleteAt{
		Status:  status,
		Message: message,
		Data:    m.mapResponseWithdrawalDeleteAt(withdraw),
	}
}

// ToProtoResponseWithdrawDelete maps a status and message to an ApiResponseWithdrawDelete proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawDelete containing the mapped data.
func (m *withdrawCommandProtoMapper) ToProtoResponseWithdrawDelete(status string, message string) *pb.ApiResponseWithdrawDelete {
	return &pb.ApiResponseWithdrawDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseWithdrawAll maps a status and message to an ApiResponseWithdrawAll proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawAll containing the mapped data.
func (m *withdrawCommandProtoMapper) ToProtoResponseWithdrawAll(status string, message string) *pb.ApiResponseWithdrawAll {
	return &pb.ApiResponseWithdrawAll{
		Status:  status,
		Message: message,
	}
}

// mapResponseWithdrawal maps a single WithdrawResponse to its protobuf representation.
//
// Args:
//   - withdraw: The WithdrawResponse that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponse protobuf message containing the mapped data.
func (w *withdrawCommandProtoMapper) mapResponseWithdrawal(withdraw *response.WithdrawResponse) *pb.WithdrawResponse {
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

// mapResponseWithdrawalDeleteAt maps a single WithdrawResponseDeleteAt to its protobuf representation.
//
// Args:
//   - withdraw: The WithdrawResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseDeleteAt protobuf message containing the mapped data.
func (w *withdrawCommandProtoMapper) mapResponseWithdrawalDeleteAt(withdraw *response.WithdrawResponseDeleteAt) *pb.WithdrawResponseDeleteAt {
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
