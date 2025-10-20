package topupservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// TopupQueryResponseMapper provides methods to map TopupRecord domain models to TopupResponse and TopupResponseDeleteAt API-compatible responses for query operations.
type topupQueryResponseMapper struct {
}

// NewTopupQueryResponseMapper returns a pointer to a new topupQueryResponseMapper instance.
// This struct provides methods to map TopupRecord domain models to TopupResponse and TopupResponseDeleteAt API-compatible response types.
func NewTopupQueryResponseMapper() TopupQueryResponseMapper {
	return &topupQueryResponseMapper{}
}

// ToTopupResponse maps a single TopupRecord domain model to a TopupResponse API-compatible response type.
// Args:
//   - topup: A pointer to a TopupRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a TopupResponse containing the mapped data, including
//     ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *topupQueryResponseMapper) ToTopupResponse(topup *record.TopupRecord) *response.TopupResponse {
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

// ToTopupResponses maps a slice of TopupRecord to a slice of TopupResponse API-compatible response types.
//
// Args:
//   - topups: A slice of TopupRecord containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponse containing the mapped data, including
//     ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *topupQueryResponseMapper) ToTopupResponses(topups []*record.TopupRecord) []*response.TopupResponse {
	var responses []*response.TopupResponse

	for _, response := range topups {
		responses = append(responses, s.ToTopupResponse(response))
	}

	return responses
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
func (s *topupQueryResponseMapper) ToTopupResponseDeleteAt(topup *record.TopupRecord) *response.TopupResponseDeleteAt {
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

// ToTopupResponsesDeleteAt maps a slice of TopupRecord domain models to a slice of TopupResponseDeleteAt API-compatible response types.
// It includes soft delete information by mapping the DeletedAt field.
//
// Args:
//   - topups: A slice of TopupRecord containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponseDeleteAt containing the mapped data, including
//     ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *topupQueryResponseMapper) ToTopupResponsesDeleteAt(topups []*record.TopupRecord) []*response.TopupResponseDeleteAt {
	var responses []*response.TopupResponseDeleteAt

	for _, response := range topups {
		responses = append(responses, s.ToTopupResponseDeleteAt(response))
	}

	return responses
}
