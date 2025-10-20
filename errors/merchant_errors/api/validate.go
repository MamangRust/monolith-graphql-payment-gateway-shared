package merchantapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var ErrApiValidateCreateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create merchant request", http.StatusBadRequest)
}

var ErrApiValidateUpdateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant request", http.StatusBadRequest)
}

var ErrApiValidateUpdateMerchantStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant status request", http.StatusBadRequest)
}
