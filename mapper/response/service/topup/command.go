package topupservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

type topupCommandResponseMapper struct {
}

func NewTopupCommandResponseMapper() TopupCommandResponseMapper {
	return &topupCommandResponseMapper{}
}

// ToTopupResponse maps a single TopupRecord domain model to a TopupResponse API-compatible response type.
// Args:
//   - topup: A pointer to a TopupRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a TopupResponse containing the mapped data, including
//     ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *topupCommandResponseMapper) ToTopupResponse(topup *record.TopupRecord) *response.TopupResponse {
	return &response.TopupResponse{
		ID:          topup.ID,
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: topup.TopupAmount,
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}

// ToTopupResponseDeleteAt maps a single TopupRecord domain model to a TopupResponseDeleteAt API-compatible response type.
// It includes soft delete information by mapping the DeletedAt field.
//
// Args:
//   - topup: A pointer to a TopupRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a TopupResponseDeleteAt containing the mapped data, including
//     ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *topupCommandResponseMapper) ToTopupResponseDeleteAt(topup *record.TopupRecord) *response.TopupResponseDeleteAt {
	return &response.TopupResponseDeleteAt{
		ID:          topup.ID,
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: topup.TopupAmount,
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
		DeletedAt:   topup.DeletedAt,
	}
}
