package withdrawresponseservice

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type withdrawQueryResponseMapper struct {
}

func NewWithdrawQueryResponseMapper() WithdrawQueryResponseMapper {
	return &withdrawQueryResponseMapper{}
}

// ToWithdrawResponse converts a single WithdrawRecord into a WithdrawResponse.
//
// Args:
//   - withdraw: The WithdrawRecord that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponse containing the mapped data, including ID, WithdrawNo,
//     CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.
func (s *withdrawQueryResponseMapper) ToWithdrawResponse(withdraw *record.WithdrawRecord) *response.WithdrawResponse {
	return &response.WithdrawResponse{
		ID:             withdraw.ID,
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: withdraw.WithdrawAmount,
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
	}
}

// ToWithdrawsResponse converts multiple WithdrawRecords into a slice of WithdrawResponse.
//
// Args:
//   - withdraws: The WithdrawRecords that needs to be converted.
//
// Returns:
//   - A slice of WithdrawResponse containing the mapped data, including ID, WithdrawNo,
//     CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.
func (s *withdrawQueryResponseMapper) ToWithdrawsResponse(withdraws []*record.WithdrawRecord) []*response.WithdrawResponse {
	var withdrawResponses []*response.WithdrawResponse
	for _, withdraw := range withdraws {
		withdrawResponses = append(withdrawResponses, s.ToWithdrawResponse(withdraw))
	}
	return withdrawResponses
}

// ToWithdrawResponseDeleteAt converts a single WithdrawRecord into a WithdrawResponseDeleteAt.
//
// Args:
//   - withdraw: The WithdrawRecord that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseDeleteAt containing the mapped data, including ID, WithdrawNo,
//     CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawQueryResponseMapper) ToWithdrawResponseDeleteAt(withdraw *record.WithdrawRecord) *response.WithdrawResponseDeleteAt {
	return &response.WithdrawResponseDeleteAt{
		ID:             withdraw.ID,
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: withdraw.WithdrawAmount,
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
		DeletedAt:      withdraw.DeletedAt,
	}
}

func (s *withdrawQueryResponseMapper) ToWithdrawsResponseDeleteAt(withdraws []*record.WithdrawRecord) []*response.WithdrawResponseDeleteAt {
	var withdrawResponses []*response.WithdrawResponseDeleteAt

	for _, withdraw := range withdraws {
		withdrawResponses = append(withdrawResponses, s.ToWithdrawResponseDeleteAt(withdraw))
	}
	return withdrawResponses
}
