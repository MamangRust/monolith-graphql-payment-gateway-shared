package topupapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyTopupAmounts returns a 500 error when fetching monthly topup amounts fails.
var ErrApiFailedFindMonthlyTopupAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupAmounts returns a 500 error when fetching yearly topup amounts fails.
var ErrApiFailedFindYearlyTopupAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupAmountsByCardNumber returns a 500 error when fetching monthly topup amounts by card number fails.
var ErrApiFailedFindMonthlyTopupAmountsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amounts by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupAmountsByCardNumber returns a 500 error when fetching yearly topup amounts by card number fails.
var ErrApiFailedFindYearlyTopupAmountsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amounts by card number", http.StatusInternalServerError)
}
