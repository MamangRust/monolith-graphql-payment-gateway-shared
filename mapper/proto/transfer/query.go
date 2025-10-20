package transferprotomapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	protomapper "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto"
	pbhelpers "github.com/MamangRust/monolith-payment-gateway-pb"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type transferQueryProtoMapper struct {
}

func NewTransferQueryProtoMapper() TransferQueryProtoMapper {
	return &transferQueryProtoMapper{}
}

// ToProtoResponseTransfer maps a status, message, and a single TransferResponse
// to an ApiResponseTransfer proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: The TransferResponse to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransfer containing the mapped data.
func (m *transferQueryProtoMapper) ToProtoResponseTransfer(status string, message string, pbResponse *response.TransferResponse) *pb.ApiResponseTransfer {
	return &pb.ApiResponseTransfer{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransfer(pbResponse),
	}
}

// ToProtoResponseTransfers maps a status, message, and a list of TransferResponse
// to an ApiResponseTransfers proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponse to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransfers containing the mapped data.
func (m *transferQueryProtoMapper) ToProtoResponseTransfers(status string, message string, pbResponse []*response.TransferResponse) *pb.ApiResponseTransfers {
	return &pb.ApiResponseTransfers{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransfer(pbResponse),
	}
}

// ToProtoResponsePaginationTransfer maps a pagination meta, status, message, and a list of TransferResponse
// to an ApiResponsePaginationTransfer proto message.
//
// Args:
//   - pagination: The pagination meta that needs to be converted.
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponse to be converted.
//
// Returns:
//   - A pointer to an ApiResponsePaginationTransfer containing the mapped data.
func (m *transferQueryProtoMapper) ToProtoResponsePaginationTransfer(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.TransferResponse) *pb.ApiResponsePaginationTransfer {
	return &pb.ApiResponsePaginationTransfer{
		Status:     status,
		Message:    message,
		Data:       m.mapResponsesTransfer(pbResponse),
		Pagination: protomapper.MapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationTransferDeleteAt maps paginated soft-deleted transfer responses to a protobuf message.
//
// Args:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//   - pbResponse: A slice of TransferResponseDeleteAt objects to be included in the response.
//
// Returns:
//   - A pointer to an ApiResponsePaginationTransferDeleteAt containing the status, message, transfer data, and pagination data.
func (m *transferQueryProtoMapper) ToProtoResponsePaginationTransferDeleteAt(pagination *pbhelpers.PaginationMeta, status string, message string, pbResponse []*response.TransferResponseDeleteAt) *pb.ApiResponsePaginationTransferDeleteAt {
	return &pb.ApiResponsePaginationTransferDeleteAt{
		Status:     status,
		Message:    message,
		Data:       m.mapResponsesTransferDeleteAt(pbResponse),
		Pagination: protomapper.MapPaginationMeta(pagination),
	}
}

// mapResponseTransfer maps a TransferResponse to a TransferResponse proto message.
//
// Args:
//   - transfer: The TransferResponse to be converted.
//
// Returns:
//   - A pointer to a TransferResponse containing the mapped data.
func (t *transferQueryProtoMapper) mapResponseTransfer(transfer *response.TransferResponse) *pb.TransferResponse {
	return &pb.TransferResponse{
		Id:             int32(transfer.ID),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int32(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
	}
}

// mapResponsesTransfer maps a slice of TransferResponse to a slice of TransferResponse proto messages.
//
// Args:
//   - transfers: A slice of TransferResponse to be converted.
//
// Returns:
//   - A slice of pointers to TransferResponse containing the mapped data.
func (t *transferQueryProtoMapper) mapResponsesTransfer(transfers []*response.TransferResponse) []*pb.TransferResponse {
	var responses []*pb.TransferResponse

	for _, response := range transfers {
		responses = append(responses, t.mapResponseTransfer(response))
	}

	return responses
}

// mapResponseTransferDeleteAt maps a TransferResponseDeleteAt to its protobuf representation.
//
// Args:
//   - transfer: The TransferResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to a TransferResponseDeleteAt protobuf message containing the mapped data.
func (t *transferQueryProtoMapper) mapResponseTransferDeleteAt(transfer *response.TransferResponseDeleteAt) *pb.TransferResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if transfer.DeletedAt != nil {
		deletedAt = wrapperspb.String(*transfer.DeletedAt)
	}

	return &pb.TransferResponseDeleteAt{
		Id:             int32(transfer.ID),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int32(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
		DeletedAt:      deletedAt,
	}
}

// mapResponsesTransferDeleteAt maps a slice of TransferResponseDeleteAt to a slice of
// their protobuf representations.
//
// Args:
//   - transfers: A slice of TransferResponseDeleteAt to be converted.
//
// Returns:
//   - A slice of pointers to TransferResponseDeleteAt protobuf messages containing the mapped data.
func (t *transferQueryProtoMapper) mapResponsesTransferDeleteAt(transfers []*response.TransferResponseDeleteAt) []*pb.TransferResponseDeleteAt {
	var responses []*pb.TransferResponseDeleteAt

	for _, response := range transfers {
		responses = append(responses, t.mapResponseTransferDeleteAt(response))
	}

	return responses
}
