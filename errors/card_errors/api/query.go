package cardapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindAllCards returns a 500 Internal Server Error when fetching all cards fails.
var ErrApiFailedFindAllCards = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all cards", http.StatusInternalServerError)
}

// ErrApiFailedFindByIdCard returns a 500 Internal Server Error when fetching a card by ID fails.
var ErrApiFailedFindByIdCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card by ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByUserIdCard represents error when failing to fetch cards by user ID.
var ErrApiFailedFindByUserIdCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card by user ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByActiveCard represents error when failing to fetch active cards.
var ErrApiFailedFindByActiveCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active cards", http.StatusInternalServerError)
}

// ErrApiFailedFindByTrashedCard represents error when failing to fetch trashed cards.
var ErrApiFailedFindByTrashedCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed cards", http.StatusInternalServerError)
}

// ErrApiFailedFindByCardNumber represents error when failing to fetch card by card number.
var ErrApiFailedFindByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card by card number", http.StatusInternalServerError)
}
