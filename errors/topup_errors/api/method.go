package topupapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyTopupMethods returns a 500 error when fetching monthly topup methods fails.
var ErrApiFailedFindMonthlyTopupMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup methods", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupMethods returns a 500 error when fetching yearly topup methods fails.
var ErrApiFailedFindYearlyTopupMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup methods", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupMethodsByCardNumber returns a 500 error when fetching monthly topup methods by card number fails.
var ErrApiFailedFindMonthlyTopupMethodsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup methods by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupMethodsByCardNumber returns a 500 error when fetching yearly topup methods by card number fails.
var ErrApiFailedFindYearlyTopupMethodsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup methods by card number", http.StatusInternalServerError)
}
