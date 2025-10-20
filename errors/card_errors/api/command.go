package cardapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedCreateCard represents error when failing to create a new card.
var ErrApiFailedCreateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create card", http.StatusInternalServerError)
}

// ErrApiFailedUpdateCard represents error when failing to update a card.
var ErrApiFailedUpdateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update card", http.StatusInternalServerError)
}

// ErrApiFailedTrashCard represents error when failing to move a card to trash.
var ErrApiFailedTrashCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash card", http.StatusInternalServerError)
}

// ErrApiFailedRestoreCard represents error when failing to restore a trashed card.
var ErrApiFailedRestoreCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore card", http.StatusInternalServerError)
}

// ErrApiFailedDeleteCardPermanent represents error when failing to permanently delete a card.
var ErrApiFailedDeleteCardPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete card", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllCard represents error when failing to restore all trashed cards.
var ErrApiFailedRestoreAllCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all cards", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllCardPermanent represents error when failing to permanently delete all trashed cards.
var ErrApiFailedDeleteAllCardPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all cards", http.StatusInternalServerError)
}
