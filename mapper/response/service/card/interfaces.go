package cardservicemapper

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// CardBaseResponseMapper provides methods to map CardRecord domain models to CardResponse
type CardBaseResponseMapper interface {
	// ToCardResponse converts a single card record into a CardResponse.
	//
	// Parameters:
	//   - card: A pointer to a CardRecord representing the card record.
	//
	// Returns:
	//   - A pointer to a CardResponse containing the mapped data, including ID, UserID, CardNumber,
	//     CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.
	ToCardResponse(card *record.CardRecord) *response.CardResponse
}

// CardQueryResponseMapper provides methods to map CardRecord domain models to CardResponse and CardResponseDeleteAt
type CardQueryResponseMapper interface {
	CardBaseResponseMapper

	// ToCardsResponse converts a list of card records into a list of CardResponse.
	//
	// Parameters:
	//   - cards: A pointer to a slice of CardRecord representing the card records.
	//
	// Returns:
	//   - A pointer to a slice of CardResponse containing the mapped data, including ID, UserID,
	//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.
	ToCardsResponse(cards []*record.CardRecord) []*response.CardResponse

	// ToCardsResponseDeleteAt converts a list of soft-deleted card records into a list of
	// CardResponseDeleteAt.
	//
	// Parameters:
	//   - cards: A pointer to a slice of CardRecord representing the card records.
	//
	// Returns:
	//   - A pointer to a slice of CardResponseDeleteAt containing the mapped data, including ID, UserID,
	//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
	ToCardsResponseDeleteAt(cards []*record.CardRecord) []*response.CardResponseDeleteAt
}

// CardCommandResponseMapper provides methods to map CardRecord domain models to CardResponse
type CardCommandResponseMapper interface {
	CardBaseResponseMapper

	// ToCardResponseDeleteAt converts a CardRecord into a CardResponseDeleteAt.
	//
	// Parameters:
	//   - card: A pointer to a CardRecord representing the card record.
	//
	// Returns:
	//   - A pointer to a CardResponseDeleteAt containing the mapped data, including ID, UserID,
	//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
	ToCardResponseDeleteAt(card *record.CardRecord) *response.CardResponseDeleteAt
}

// CardStatisticBalanceResponseMapper provides methods to map CardRecord domain models to CardResponseMonthBalance and CardResponseYearlyBalance
type CardStatisticBalanceResponseMapper interface {
	// ToGetMonthlyBalance maps a CardMonthBalance domain model to a CardResponseMonthBalance
	// response object.
	//
	// Parameters:
	//   - card: A pointer to a CardMonthBalance domain model.
	//
	// Returns:
	//   - A pointer to a CardResponseMonthBalance containing the mapped data, including
	//     Month and TotalBalance.
	ToGetMonthlyBalance(card *record.CardMonthBalance) *response.CardResponseMonthBalance

	// ToGetMonthlyBalances converts a slice of CardMonthBalance domain models into a slice
	// of CardResponseMonthBalance response objects.
	//
	// Parameters:
	//   - cards: A slice of pointers to CardMonthBalance domain models.
	//
	// Returns:
	//   - A slice of pointers to CardResponseMonthBalance containing the mapped data for each month.
	ToGetMonthlyBalances(cards []*record.CardMonthBalance) []*response.CardResponseMonthBalance

	// ToGetYearlyBalance maps a CardYearlyBalance domain model to a CardResponseYearlyBalance
	// response object.
	//
	// Parameters:
	//   - card: A pointer to a CardYearlyBalance domain model.
	//
	// Returns:
	//   - A pointer to a CardResponseYearlyBalance containing the mapped data, including
	//     Year and TotalBalance.
	ToGetYearlyBalance(card *record.CardYearlyBalance) *response.CardResponseYearlyBalance

	// ToGetYearlyBalances converts a slice of CardYearlyBalance domain models into a slice
	// of CardResponseYearlyBalance response objects.
	//
	// Parameters:
	//   - cards: A slice of pointers to CardYearlyBalance domain models.
	//
	// Returns:
	//   - A slice of pointers to CardResponseYearlyBalance containing the mapped data for each year.
	ToGetYearlyBalances(cards []*record.CardYearlyBalance) []*response.CardResponseYearlyBalance
}

// CardStatisticAmountResponseMapper provides methods to map CardRecord domain models to CardResponseMonthAmount and CardResponseYearAmount
type CardStatisticAmountResponseMapper interface {
	// ToGetMonthlyAmount maps a CardMonthAmount domain model to a CardResponseMonthAmount
	// domain model.
	//
	// Parameters:
	//   - card: A pointer to a CardMonthAmount representing the domain model.
	//
	// Returns:
	//   - A pointer to a CardResponseMonthAmount containing the mapped data, including Month and TotalAmount.
	ToGetMonthlyAmount(card *record.CardMonthAmount) *response.CardResponseMonthAmount

	// ToGetMonthlyAmounts converts a slice of CardMonthAmount domain models into a slice
	// of CardResponseMonthAmount response objects.
	//
	// Parameters:
	//   - cards: A slice of pointers to CardMonthAmount domain models.
	//
	// Returns:
	//   - A slice of pointers to CardResponseMonthAmount containing the mapped data for each month.
	ToGetMonthlyAmounts(cards []*record.CardMonthAmount) []*response.CardResponseMonthAmount

	// ToGetYearlyAmount maps a CardYearAmount domain model to a CardResponseYearAmount
	// domain model.
	//
	// Parameters:
	//   - card: A pointer to a CardYearAmount representing the domain model.
	//
	// Returns:
	//   - A pointer to a CardResponseYearAmount containing the mapped data, including Year and TotalAmount.
	ToGetYearlyAmount(card *record.CardYearAmount) *response.CardResponseYearAmount

	// ToGetYearlyAmounts converts a slice of CardYearAmount domain models into a slice
	// of CardResponseYearAmount response objects.
	//
	// Parameters:
	//   - cards: A slice of pointers to CardYearAmount domain models.
	//
	// Returns:
	//   - A slice of pointers to CardResponseYearAmount containing the mapped data for each year.
	ToGetYearlyAmounts(cards []*record.CardYearAmount) []*response.CardResponseYearAmount
}
