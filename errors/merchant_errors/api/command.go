package merchantapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var ErrApiFailedCreateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create merchant", http.StatusInternalServerError)
}

var ErrApiFailedUpdateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant", http.StatusInternalServerError)
}

var ErrApiFailedUpdateMerchantStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant status", http.StatusInternalServerError)
}

var ErrApiFailedTrashMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash merchant", http.StatusInternalServerError)
}

var ErrApiFailedRestoreMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore merchant", http.StatusInternalServerError)
}

var ErrApiFailedDeleteMerchantPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete merchant", http.StatusInternalServerError)
}

var ErrApiFailedRestoreAllMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all merchants", http.StatusInternalServerError)
}

var ErrApiFailedDeleteAllMerchantPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all merchants", http.StatusInternalServerError)
}
