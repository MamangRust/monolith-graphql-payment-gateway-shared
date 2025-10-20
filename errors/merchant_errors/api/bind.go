package merchantapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var ErrApiBindCreateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create merchant request", http.StatusBadRequest)
}

var ErrApiBindUpdateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant request", http.StatusBadRequest)
}

var ErrApiBindUpdateMerchantStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant status request", http.StatusBadRequest)
}
