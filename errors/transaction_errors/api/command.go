package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedCreateTransaction returns an API error response
// when creating a new transaction fails.
var ErrApiFailedCreateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create transaction", http.StatusInternalServerError)
}

// ErrApiFailedUpdateTransaction returns an API error response
// when updating an existing transaction fails.
var ErrApiFailedUpdateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update transaction", http.StatusInternalServerError)
}

// ErrApiFailedRestoreTransaction returns an API error response
// when restoring a soft-deleted transaction fails.
var ErrApiFailedRestoreTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore transaction", http.StatusInternalServerError)
}

// ErrApiFailedTrashTransaction returns an API error response
// when moving a transaction to trash (soft-delete) fails.
var ErrApiFailedTrashTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to move transaction to trash", http.StatusInternalServerError)
}

// ErrApiFailedDeletePermanent returns an API error response
// when permanently deleting a transaction fails.
var ErrApiFailedDeletePermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete transaction", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllTransactions returns an API error response
// when restoring all trashed transactions fails.
var ErrApiFailedRestoreAllTransactions = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all transactions", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllPermanent returns an API error response
// when permanently deleting all trashed transactions fails.
var ErrApiFailedDeleteAllPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all trashed transactions", http.StatusInternalServerError)
}
