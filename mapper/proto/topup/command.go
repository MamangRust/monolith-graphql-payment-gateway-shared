package topupprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	helpersproto "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto/helpers"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup"
)

type topupCommandProtoMapper struct {
}

func NewTopupCommandProtoMapper() TopupCommandProtoMapper {
	return &topupCommandProtoMapper{}
}

// ToProtoResponseTopup maps a single top-up record to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A pointer to TopupResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopup containing the status, message, and top-up data.
func (t *topupCommandProtoMapper) ToProtoResponseTopup(status string, message string, s *response.TopupResponse) *pb.ApiResponseTopup {
	return &pb.ApiResponseTopup{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopup(s),
	}
}

// ToProtoResponseTopupDeleteAt maps a soft-deleted top-up record to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A pointer to TopupResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to ApiResponseTopupDeleteAt containing the status, message, and mapped top-up data.
func (t *topupCommandProtoMapper) ToProtoResponseTopupDeleteAt(status string, message string, s *response.TopupResponseDeleteAt) *pb.ApiResponseTopupDeleteAt {
	return &pb.ApiResponseTopupDeleteAt{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopupDeleteAt(s),
	}
}

// ToProtoResponseTopupDelete returns a response indicating a top-up record has been deleted.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//
// Returns:
//
//	A pointer to ApiResponseTopupDelete containing the status and message.
func (t topupCommandProtoMapper) ToProtoResponseTopupDelete(status string, message string) *pb.ApiResponseTopupDelete {
	return &pb.ApiResponseTopupDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseTopupAll returns all top-up records in a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//
// Returns:
//   - A pointer to ApiResponseTopupAll containing the status and message.
func (t topupCommandProtoMapper) ToProtoResponseTopupAll(status string, message string) *pb.ApiResponseTopupAll {
	return &pb.ApiResponseTopupAll{
		Status:  status,
		Message: message,
	}
}

// mapResponseTopup maps a single TopupResponse to its corresponding protobuf response.
//
// Args:
//   - topup: A pointer to TopupResponse that needs to be converted.
//
// Returns:
//   - A pointer to TopupResponse containing the mapped top-up data.
func (t *topupCommandProtoMapper) mapResponseTopup(topup *response.TopupResponse) *pb.TopupResponse {
	return &pb.TopupResponse{
		Id:          int32(topup.ID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int32(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}

// mapResponseTopupDeleteAt maps a single TopupResponseDeleteAt to its corresponding protobuf response.
//
// Args:
//   - topup: A pointer to TopupResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to TopupResponseDeleteAt containing the mapped top-up data, including a non-nil DeletedAt field if present.
func (t *topupCommandProtoMapper) mapResponseTopupDeleteAt(topup *response.TopupResponseDeleteAt) *pb.TopupResponseDeleteAt {
	res := &pb.TopupResponseDeleteAt{
		Id:          int32(topup.ID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int32(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}

	if topup.DeletedAt != nil {
		res.DeletedAt = helpersproto.StringPtrToWrapper(topup.DeletedAt)
	}

	return res
}
