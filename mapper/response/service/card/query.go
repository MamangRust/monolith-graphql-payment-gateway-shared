package cardservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// CardQueryResponseMapper provides methods to map CardRecord domain models to CardResponse
type cardQueryResponseMapper struct {
}

// NewCardQueryResponseMapper returns a new instance of cardQueryResponseMapper, which provides
// methods to map CardRecord domain models to CardResponse domain models for query operations.
func NewCardQueryResponseMapper() CardQueryResponseMapper {
	return &cardQueryResponseMapper{}
}

// ToCardResponse converts a single card record into a CardResponse.
//
// Parameters:
//   - card: A pointer to a CardRecord representing the card record.
//
// Returns:
//   - A pointer to a CardResponse containing the mapped data, including ID, UserID, CardNumber,
//     CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.
func (s *cardQueryResponseMapper) ToCardResponse(card *record.CardRecord) *response.CardResponse {
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

// ToCardsResponse converts a list of card records into a list of CardResponse.
//
// Parameters:
//   - cards: A pointer to a slice of CardRecord representing the card records.
//
// Returns:
//   - A pointer to a slice of CardResponse containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.
func (s *cardQueryResponseMapper) ToCardsResponse(cards []*record.CardRecord) []*response.CardResponse {
	var response []*response.CardResponse

	for _, card := range cards {
		response = append(response, s.ToCardResponse(card))
	}

	return response
}

// ToCardResponseDeleteAt converts a CardRecord into a CardResponseDeleteAt.
//
// Parameters:
//   - card: A pointer to a CardRecord representing the card record.
//
// Returns:
//   - A pointer to a CardResponseDeleteAt containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardQueryResponseMapper) ToCardResponseDeleteAt(card *record.CardRecord) *response.CardResponseDeleteAt {
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

// ToCardsResponseDeleteAt converts a list of soft-deleted card records into a list of
// CardResponseDeleteAt.
//
// Parameters:
//   - cards: A pointer to a slice of CardRecord representing the card records.
//
// Returns:
//   - A pointer to a slice of CardResponseDeleteAt containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardQueryResponseMapper) ToCardsResponseDeleteAt(cards []*record.CardRecord) []*response.CardResponseDeleteAt {
	var response []*response.CardResponseDeleteAt

	for _, card := range cards {
		response = append(response, s.ToCardResponseDeleteAt(card))
	}

	return response
}
