package cardapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyBalance returns an API error response when fetching the monthly balance fails.
var ErrApiFailedFindMonthlyBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly balance", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyBalance returns an API error response when fetching the yearly balance fails.
var ErrApiFailedFindYearlyBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly balance", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyBalanceByCard returns an API error response when fetching the monthly balance by card fails.
var ErrApiFailedFindMonthlyBalanceByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly balance by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyBalanceByCard returns an API error response when fetching the yearly balance by card fails.
var ErrApiFailedFindYearlyBalanceByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly balance by card", http.StatusInternalServerError)
}
