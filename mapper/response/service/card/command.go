package cardservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// CardCommandResponseMapper provides methods to map CardRecord domain models to CardResponse
type cardCommandResponseMapper struct {
}

// NewCardCommandResponseMapper returns a new instance of cardCommandResponseMapper, which provides
// methods to map CardRecord domain models to CardResponse domain models for command operations.
func NewCardCommandResponseMapper() CardCommandResponseMapper {
	return &cardCommandResponseMapper{}
}

// ToCardResponse converts a single card record into a CardResponse.
//
// Args:
//   - card: A pointer to a CardRecord representing the card record.
//
// Returns:
//   - A pointer to a CardResponse containing the mapped data, including ID, UserID, CardNumber,
//     CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.
func (s *cardCommandResponseMapper) ToCardResponse(card *record.CardRecord) *response.CardResponse {
	return &response.CardResponse{
		ID:           card.ID,
		UserID:       card.UserID,
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		CVV:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}
}

// ToCardResponseDeleteAt converts a single card record into a CardResponseDeleteAt.
//
// Parameters:
//   - card: A pointer to a CardRecord representing the card record.
//
// Returns:
//   - A pointer to a CardResponseDeleteAt containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardCommandResponseMapper) ToCardResponseDeleteAt(card *record.CardRecord) *response.CardResponseDeleteAt {
	return &response.CardResponseDeleteAt{
		ID:           card.ID,
		UserID:       card.UserID,
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		CVV:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
		DeletedAt:    card.DeletedAt,
	}
}
