package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedMonthlyMethods returns an API error response
// when retrieving monthly payment methods fails.
var ErrApiFailedMonthlyMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods", http.StatusInternalServerError)
}

// ErrApiFailedYearlyMethods returns an API error response
// when retrieving yearly payment methods fails.
var ErrApiFailedYearlyMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods", http.StatusInternalServerError)
}

// ErrApiFailedMonthlyMethodsByCard returns an API error response
// when retrieving monthly payment methods by card number fails.
var ErrApiFailedMonthlyMethodsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by card number", http.StatusInternalServerError)
}

// ErrApiFailedYearlyMethodsByCard returns an API error response
// when retrieving yearly payment methods by card number fails.
var ErrApiFailedYearlyMethodsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by card number", http.StatusInternalServerError)
}
