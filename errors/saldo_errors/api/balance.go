package saldoapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var ErrApiFailedFindMonthlySaldoBalances = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly saldo balances", http.StatusInternalServerError)
}

var ErrApiFailedFindYearlySaldoBalances = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly saldo balances", http.StatusInternalServerError)
}
