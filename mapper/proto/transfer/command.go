package transferprotomapper

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	helpersproto "github.com/MamangRust/monolith-graphql-payment-gateway-shared/mapper/proto/helpers"
)

type transferCommandProtoMapper struct {
}

func NewTransferCommandProtoMapper() TransferCommandProtoMapper {
	return &transferCommandProtoMapper{}
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
func (m *transferCommandProtoMapper) ToProtoResponseTransfer(status string, message string, pbResponse *response.TransferResponse) *pb.ApiResponseTransfer {
	return &pb.ApiResponseTransfer{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransfer(pbResponse),
	}
}

func (m *transferCommandProtoMapper) ToProtoResponseTransferDeleteAt(status string, message string, pbResponse *response.TransferResponseDeleteAt) *pb.ApiResponseTransferDeleteAt {
	return &pb.ApiResponseTransferDeleteAt{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransferDeleteAt(pbResponse),
	}
}

// ToProtoResponseTransferDelete maps a status, message, and no data
// to an ApiResponseTransferDelete proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//
// Returns:
//   - A pointer to an ApiResponseTransferDelete containing the mapped data.
func (m *transferCommandProtoMapper) ToProtoResponseTransferDelete(status string, message string) *pb.ApiResponseTransferDelete {
	return &pb.ApiResponseTransferDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseTransferAll maps a status, message, and no data
// to an ApiResponseTransferAll proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//
// Returns:
//   - A pointer to an ApiResponseTransferAll containing the mapped data.
func (m *transferCommandProtoMapper) ToProtoResponseTransferAll(status string, message string) *pb.ApiResponseTransferAll {
	return &pb.ApiResponseTransferAll{
		Status:  status,
		Message: message,
	}
}

// mapResponseTransfer maps a TransferResponse to a TransferResponse proto message.
//
// Args:
//   - transfer: The TransferResponse to be converted.
//
// Returns:
//   - A pointer to a TransferResponse containing the mapped data.
func (t *transferCommandProtoMapper) mapResponseTransfer(transfer *response.TransferResponse) *pb.TransferResponse {
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

func (t *transferCommandProtoMapper) mapResponseTransferDeleteAt(transfer *response.TransferResponseDeleteAt) *pb.TransferResponseDeleteAt {
	res := &pb.TransferResponseDeleteAt{
		Id:             int32(transfer.ID),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int32(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
	}

	if transfer.DeletedAt != nil {
		res.DeletedAt = helpersproto.StringPtrToWrapper(transfer.DeletedAt)
	}

	return res
}
