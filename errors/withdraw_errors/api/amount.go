package withdrawapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyWithdraws is an error response for failed to fetch monthly withdraw amounts.
var ErrApiFailedFindMonthlyWithdraws = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdraws is an error response for failed to fetch yearly withdraw amounts.
var ErrApiFailedFindYearlyWithdraws = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawsByCardNumber is an error response for failed to fetch monthly withdraw amounts by card number.
var ErrApiFailedFindMonthlyWithdrawsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amounts by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawsByCardNumber is an error response for failed to fetch yearly withdraw amounts by card number.
var ErrApiFailedFindYearlyWithdrawsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amounts by card number", http.StatusInternalServerError)
}
