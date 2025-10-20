package saldoapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindAllSaldo is returned when fetching all saldo records fails.
var ErrApiFailedFindAllSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all saldos", http.StatusInternalServerError)
}

// ErrApiFailedFindAllSaldoTrashed is returned when fetching trashed saldo records fails.
var ErrApiFailedFindAllSaldoTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all saldos trashed", http.StatusInternalServerError)
}

// ErrApiFailedFindAllSaldoActive is returned when fetching active saldo records fails.
var ErrApiFailedFindAllSaldoActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all saldos active", http.StatusInternalServerError)
}

// ErrApiFailedFindByIdSaldo is returned when fetching saldo by ID fails.
var ErrApiFailedFindByIdSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch saldo by ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByCardNumberSaldo is returned when fetching saldo by card number fails.
var ErrApiFailedFindByCardNumberSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch saldo by card number", http.StatusInternalServerError)
}
