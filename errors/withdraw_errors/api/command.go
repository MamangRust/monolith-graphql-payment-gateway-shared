package withdrawapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedCreateWithdraw is an error response for failed to create withdraw.
var ErrApiFailedCreateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create withdraw", http.StatusInternalServerError)
}

// ErrApiFailedUpdateWithdraw is an error response for failed to update withdraw.
var ErrApiFailedUpdateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update withdraw", http.StatusInternalServerError)
}

// ErrApiFailedTrashedWithdraw is an error response for failed to move withdraw to trash.
var ErrApiFailedTrashedWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to move withdraw to trash", http.StatusInternalServerError)
}

// ErrApiFailedRestoreWithdraw is an error response for failed to restore withdraw.
var ErrApiFailedRestoreWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore withdraw", http.StatusInternalServerError)
}

// ErrApiFailedDeleteWithdrawPermanent is an error response for failed to permanently delete withdraw.
var ErrApiFailedDeleteWithdrawPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete withdraw", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllWithdraw is an error response for failed to restore all withdraws.
var ErrApiFailedRestoreAllWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all withdraws", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllWithdrawPermanent is an error response for failed to permanently delete all withdraws.
var ErrApiFailedDeleteAllWithdrawPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all withdraws", http.StatusInternalServerError)
}
