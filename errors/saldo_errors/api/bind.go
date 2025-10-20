package saldoapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiBindCreateSaldo is returned when binding fails for saldo creation.
var ErrApiBindCreateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Saldo request", http.StatusBadRequest)
}

// ErrApiBindUpdateSaldo is returned when binding fails for saldo update.
var ErrApiBindUpdateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Saldo request", http.StatusBadRequest)
}
