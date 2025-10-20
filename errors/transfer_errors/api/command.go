package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedCreateTransfer returns an API error response
// when creating a new transfer fails.
var ErrApiFailedCreateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create transfer", http.StatusInternalServerError)
}

// ErrApiFailedUpdateTransfer returns an API error response
// when updating a transfer fails.
var ErrApiFailedUpdateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update transfer", http.StatusInternalServerError)
}

// ErrApiFailedTrashedTransfer returns an API error response
// when attempting to trash (soft-delete) a transfer fails.
var ErrApiFailedTrashedTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash transfer", http.StatusInternalServerError)
}

// ErrApiFailedRestoreTransfer returns an API error response
// when attempting to restore a trashed transfer fails.
var ErrApiFailedRestoreTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore transfer", http.StatusInternalServerError)
}

// ErrApiFailedDeleteTransferPermanent returns an API error response
// when attempting to permanently delete a transfer fails.
var ErrApiFailedDeleteTransferPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete transfer", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllTransfer returns an API error response
// when attempting to restore all trashed transfers fails.
var ErrApiFailedRestoreAllTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all transfers", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllTransferPermanent returns an API error response
// when attempting to permanently delete all transfers fails.
var ErrApiFailedDeleteAllTransferPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all transfers", http.StatusInternalServerError)
}
