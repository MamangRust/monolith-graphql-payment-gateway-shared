package topupprotomapper

import (
	pbhelpers "github.com/MamangRust/monolith-graphql-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/topup"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	protomapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type topupQueryProtoMapper struct {
}

func NewTopupQueryProtoMapper() TopupQueryProtoMapper {
	return &topupQueryProtoMapper{}
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
func (t *topupQueryProtoMapper) ToProtoResponseTopup(status string, message string, s *response.TopupResponse) *pb.ApiResponseTopup {
	return &pb.ApiResponseTopup{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopup(s),
	}
}

// ToProtoResponsePaginationTopup maps paginated top-up records to a protobuf response.
//
// Args:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponsePaginationTopup containing the status, message, top-up data, and pagination data.
func (t *topupQueryProtoMapper) ToProtoResponsePaginationTopup(pagination *pbhelpers.PaginationMeta, status string, message string, s []*response.TopupResponse) *pb.ApiResponsePaginationTopup {
	return &pb.ApiResponsePaginationTopup{
		Status:         status,
		Message:        message,
		Data:           t.mapResponsesTopup(s),
		PaginationMeta: protomapper.MapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationTopupDeleteAt maps paginated soft-deleted top-up records to a protobuf response.
//
// Args:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponseDeleteAt that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponsePaginationTopupDeleteAt containing the status, message, top-up data, and pagination data.
func (t *topupQueryProtoMapper) ToProtoResponsePaginationTopupDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, s []*response.TopupResponseDeleteAt) *pb.ApiResponsePaginationTopupDeleteAt {
	return &pb.ApiResponsePaginationTopupDeleteAt{
		Status:         status,
		Message:        message,
		Data:           t.mapResponsesTopupDeleteAt(s),
		PaginationMeta: protomapper.MapPaginationMeta(pagination),
	}
}

// mapResponseTopup maps a single TopupResponse to its corresponding protobuf response.
//
// Args:
//   - topup: A pointer to TopupResponse that needs to be converted.
//
// Returns:
//   - A pointer to TopupResponse containing the mapped top-up data.
func (t *topupQueryProtoMapper) mapResponseTopup(topup *response.TopupResponse) *pb.TopupResponse {
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

// mapResponsesTopup maps a slice of TopupResponse to a slice of TopupResponse.
//
// Args:
//   - topups: A slice of TopupResponse that needs to be converted.
//
// Returns:
//   - A slice of TopupResponse containing the mapped top-up data.
func (t *topupQueryProtoMapper) mapResponsesTopup(topups []*response.TopupResponse) []*pb.TopupResponse {
	var responses []*pb.TopupResponse

	for _, response := range topups {
		responses = append(responses, t.mapResponseTopup(response))
	}

	return responses
}

// mapResponseTopupDeleteAt maps a single TopupResponseDeleteAt to its corresponding protobuf response.
//
// Args:
//   - topup: A pointer to TopupResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to TopupResponseDeleteAt containing the mapped top-up data.
func (t *topupQueryProtoMapper) mapResponseTopupDeleteAt(topup *response.TopupResponseDeleteAt) *pb.TopupResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if topup.DeletedAt != nil {
		deletedAt = wrapperspb.String(*topup.DeletedAt)
	}

	return &pb.TopupResponseDeleteAt{
		Id:          int32(topup.ID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int32(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}

// mapResponsesTopupDeleteAt maps a slice of TopupResponseDeleteAt to a slice of TopupResponseDeleteAt.
//
// Args:
//   - topups: A slice of TopupResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A slice of TopupResponseDeleteAt containing the mapped top-up data.
func (t *topupQueryProtoMapper) mapResponsesTopupDeleteAt(topups []*response.TopupResponseDeleteAt) []*pb.TopupResponseDeleteAt {
	var responses []*pb.TopupResponseDeleteAt

	for _, response := range topups {
		responses = append(responses, t.mapResponseTopupDeleteAt(response))
	}

	return responses
}
