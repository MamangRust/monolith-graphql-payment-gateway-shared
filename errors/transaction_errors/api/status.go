package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedMonthlySuccess returns an API error response
// when retrieving monthly successful transactions fails.
var ErrApiFailedMonthlySuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transactions", http.StatusInternalServerError)
}

// ErrApiFailedYearlySuccess returns an API error response
// when retrieving yearly successful transactions fails.
var ErrApiFailedYearlySuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transactions", http.StatusInternalServerError)
}

// ErrApiFailedMonthlyFailed returns an API error response
// when retrieving monthly failed transactions fails.
var ErrApiFailedMonthlyFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transactions", http.StatusInternalServerError)
}

// ErrApiFailedYearlyFailed returns an API error response
// when retrieving yearly failed transactions fails.
var ErrApiFailedYearlyFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transactions", http.StatusInternalServerError)
}

// ErrApiFailedMonthlySuccessByCard returns an API error response
// when retrieving monthly successful transactions by card number fails.
var ErrApiFailedMonthlySuccessByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transactions by card number", http.StatusInternalServerError)
}

// ErrApiFailedYearlySuccessByCard returns an API error response
// when retrieving yearly successful transactions by card number fails.
var ErrApiFailedYearlySuccessByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transactions by card number", http.StatusInternalServerError)
}

// ErrApiFailedMonthlyFailedByCard returns an API error response
// when retrieving monthly failed transactions by card number fails.
var ErrApiFailedMonthlyFailedByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transactions by card number", http.StatusInternalServerError)
}

// ErrApiFailedYearlyFailedByCard returns an API error response
// when retrieving yearly failed transactions by card number fails.
var ErrApiFailedYearlyFailedByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transactions by card number", http.StatusInternalServerError)
}
