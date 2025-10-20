package withdrawresponseservice

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// WithdrawCommandResponseMapper provides methods to map WithdrawRecord domain models to API-compatible WithdrawResponse objects.
type withdrawCommandResponseMapper struct {
}

// NewWithdrawCommandResponseMapper creates and returns a new instance of withdrawCommandResponseMapper.
// This instance provides methods to map WithdrawRecord domain models to
// API-compatible WithdrawResponse objects.
func NewWithdrawCommandResponseMapper() WithdrawCommandResponseMapper {
	return &withdrawCommandResponseMapper{}
}

// ToWithdrawResponse converts a single WithdrawRecord into a WithdrawResponse.
//
// Args:
//   - withdraw: The WithdrawRecord that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponse containing the mapped data, including ID, WithdrawNo,
//     CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.
func (s *withdrawCommandResponseMapper) ToWithdrawResponse(withdraw *record.WithdrawRecord) *response.WithdrawResponse {
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

// ToWithdrawResponseDeleteAt converts a single WithdrawRecord into a WithdrawResponseDeleteAt.
//
// Args:
//   - withdraw: The WithdrawRecord that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseDeleteAt containing the mapped data, including ID, WithdrawNo,
//     CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawCommandResponseMapper) ToWithdrawResponseDeleteAt(withdraw *record.WithdrawRecord) *response.WithdrawResponseDeleteAt {
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
