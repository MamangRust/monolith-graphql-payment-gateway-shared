package saldoapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var ErrApiFailedFindMonthlyTotalSaldoBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly total saldo balance", http.StatusInternalServerError)
}

var ErrApiFailedFindYearTotalSaldoBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly total saldo balance", http.StatusInternalServerError)
}
