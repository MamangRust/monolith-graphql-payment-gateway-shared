package topupapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedCreateTopup returns a 500 error when creating a new topup fails.
var ErrApiFailedCreateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create topup", http.StatusInternalServerError)
}

// ErrApiFailedUpdateTopup returns a 500 error when updating a topup fails.
var ErrApiFailedUpdateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update topup", http.StatusInternalServerError)
}

// ErrApiFailedTrashTopup returns a 500 error when moving a topup to trash fails.
var ErrApiFailedTrashTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash topup", http.StatusInternalServerError)
}

// ErrApiFailedRestoreTopup returns a 500 error when restoring a trashed topup fails.
var ErrApiFailedRestoreTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore topup", http.StatusInternalServerError)
}

// ErrApiFailedDeletePermanentTopup returns a 500 error when permanently deleting a topup fails.
var ErrApiFailedDeletePermanentTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete topup", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllTopup returns a 500 error when restoring all trashed topups fails.
var ErrApiFailedRestoreAllTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all topups", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllTopupPermanent returns a 500 error when permanently deleting all topups fails.
var ErrApiFailedDeleteAllTopupPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all topups", http.StatusInternalServerError)
}
