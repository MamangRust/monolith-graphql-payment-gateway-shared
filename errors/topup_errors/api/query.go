package topupapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindAllTopups returns a 500 Internal Server Error when fetching all top-up records fails.
var ErrApiFailedFindAllTopups = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all topups", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTopupsTrashed returns a 500 error when retrieving all trashed (soft-deleted) top-ups fails.
var ErrApiFailedFindAllTopupsTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all topups trashed", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTopupsActive returns a 500 error when retrieving all active (non-deleted) top-ups fails.
var ErrApiFailedFindAllTopupsActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all topups active", http.StatusInternalServerError)
}

// ErrApiFailedFindAllByCardNumberTopup returns a 500 error when retrieving top-ups by card number fails.
var ErrApiFailedFindAllByCardNumberTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch topups by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindByIdTopup returns a 500 error when retrieving a specific top-up by ID fails.
var ErrApiFailedFindByIdTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch topup by ID", http.StatusInternalServerError)
}
