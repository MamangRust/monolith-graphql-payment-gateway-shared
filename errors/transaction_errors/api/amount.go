package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedMonthlyAmounts returns an API error response
// when retrieving total monthly transaction amounts fails.
var ErrApiFailedMonthlyAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts", http.StatusInternalServerError)
}

// ErrApiFailedYearlyAmounts returns an API error response
// when retrieving total yearly transaction amounts fails.
var ErrApiFailedYearlyAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts", http.StatusInternalServerError)
}

// ErrApiFailedMonthlyAmountsByCard returns an API error response
// when retrieving total monthly transaction amounts by card number fails.
var ErrApiFailedMonthlyAmountsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts by card number", http.StatusInternalServerError)
}

// ErrApiFailedYearlyAmountsByCard returns an API error response
// when retrieving total yearly transaction amounts by card number fails.
var ErrApiFailedYearlyAmountsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts by card number", http.StatusInternalServerError)
}
