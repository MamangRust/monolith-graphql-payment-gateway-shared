package cardserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

// ErrCardNotFoundRes is an error response when a requested card was not found.
// Its HTTP status code is 404.
var ErrCardNotFoundRes = response.NewErrorResponse("Card not found", http.StatusNotFound)

// ErrFailedFindAllCards is an error response when retrieving all card records
// fails. Its HTTP status code is 500.
var ErrFailedFindAllCards = response.NewErrorResponse("Failed to fetch Cards", http.StatusInternalServerError)

// ErrFailedFindActiveCards is an error response when retrieving active card
// records fails. Its HTTP status code is 500.
var ErrFailedFindActiveCards = response.NewErrorResponse("Failed to fetch active Cards", http.StatusInternalServerError)

// ErrFailedFindTrashedCards is an error response when retrieving trashed card
// records fails. Its HTTP status code is 500.
var ErrFailedFindTrashedCards = response.NewErrorResponse("Failed to fetch trashed Cards", http.StatusInternalServerError)

// ErrFailedFindById is an error response when finding a card by its ID fails.
// Its HTTP status code is 500.
var ErrFailedFindById = response.NewErrorResponse("Failed to find Card by ID", http.StatusInternalServerError)

// ErrFailedFindByUserID is an error response when finding a card by its user ID
// fails. Its HTTP status code is 500.
var ErrFailedFindByUserID = response.NewErrorResponse("Failed to find Card by User ID", http.StatusInternalServerError)

// ErrFailedFindByCardNumber is an error response when finding a card by its card
// number fails. Its HTTP status code is 500.
var ErrFailedFindByCardNumber = response.NewErrorResponse("Failed to find Card by Card Number", http.StatusInternalServerError)
