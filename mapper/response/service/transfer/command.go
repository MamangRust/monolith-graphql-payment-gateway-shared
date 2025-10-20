package transferresponsemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type transferCommandResponseMapper struct {
}

func NewTransferCommandResponseMapper() TransferCommandResponseMapper {
	return &transferCommandResponseMapper{}
}

// ToTransferResponse converts a single transfer record into a TransferResponse.
//
// Args:
//   - transfer: The transfer record to be converted.
//
// Returns:
//   - A pointer to a response.TransferResponse containing the mapped data.
func (s *transferCommandResponseMapper) ToTransferResponse(transfer *record.TransferRecord) *response.TransferResponse {
	return &response.TransferResponse{
		ID:             transfer.ID,
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: transfer.TransferAmount,
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
	}
}

// ToTransferResponseDeleteAt converts a single transfer record into a TransferResponseDeleteAt.
//
// Args:
//   - transfer: The transfer record to be converted.
//
// Returns:
//   - A pointer to a response.TransferResponseDeleteAt containing the mapped data.
func (s *transferCommandResponseMapper) ToTransferResponseDeleteAt(transfer *record.TransferRecord) *response.TransferResponseDeleteAt {
	return &response.TransferResponseDeleteAt{
		ID:             transfer.ID,
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: transfer.TransferAmount,
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
		DeletedAt:      transfer.DeletedAt,
	}
}
