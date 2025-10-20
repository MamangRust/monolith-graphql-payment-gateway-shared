package saldoapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiValidateCreateSaldo is returned when validation fails for saldo creation.
var ErrApiValidateCreateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Saldo request", http.StatusBadRequest)
}

// ErrApiValidateUpdateSaldo is returned when validation fails for saldo update.
var ErrApiValidateUpdateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Saldo request", http.StatusBadRequest)
}
