package transferresponsemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// transferQueryResponseMapper provides methods to map TransferRecord and TransferResponseDeleteAt domain models to TransferResponse API-compatible response types.
type transferQueryResponseMapper struct {
}

// NewTransferQueryResponseMapper returns a new instance of transferQueryResponseMapper which provides methods to map TransferRecord domain models to TransferResponse API-compatible response types.
func NewTransferQueryResponseMapper() TransferQueryResponseMapper {
	return &transferQueryResponseMapper{}
}

// ToTransferResponse converts a single transfer record into a TransferResponse.
//
// Args:
//   - transfer: The transfer record to be converted.
//
// Returns:
//   - A pointer to a response.TransferResponse containing the mapped data.
func (s *transferQueryResponseMapper) ToTransferResponse(transfer *record.TransferRecord) *response.TransferResponse {
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

// ToTransfersResponse converts multiple transfer records into a slice of TransferResponse.
//
// Args:
//   - transfers: A slice of transfer records to be converted.
//
// Returns:
//   - A slice of pointers to response.TransferResponse containing the mapped data.
func (s *transferQueryResponseMapper) ToTransfersResponse(transfers []*record.TransferRecord) []*response.TransferResponse {
	var responses []*response.TransferResponse

	for _, response := range transfers {
		responses = append(responses, s.ToTransferResponse(response))
	}

	return responses
}

// ToTransferResponseDeleteAt converts a single transfer record into a TransferResponseDeleteAt.
//
// Args:
//   - transfer: The transfer record to be converted.
//
// Returns:
//   - A pointer to a response.TransferResponseDeleteAt containing the mapped data.
func (s *transferQueryResponseMapper) ToTransferResponseDeleteAt(transfer *record.TransferRecord) *response.TransferResponseDeleteAt {
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

// ToTransfersResponseDeleteAt converts multiple soft-deleted transfer records into a slice of TransferResponseDeleteAt.
//
// Args:
//   - transfers: A slice of pointers to TransferRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferResponseDeleteAt containing the mapped data, including
//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime, CreatedAt,
//     UpdatedAt, and DeletedAt.
func (s *transferQueryResponseMapper) ToTransfersResponseDeleteAt(transfers []*record.TransferRecord) []*response.TransferResponseDeleteAt {
	var responses []*response.TransferResponseDeleteAt

	for _, response := range transfers {
		responses = append(responses, s.ToTransferResponseDeleteAt(response))
	}

	return responses
}
