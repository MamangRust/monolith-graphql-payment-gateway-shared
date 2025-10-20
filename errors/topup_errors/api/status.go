package topupapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyTopupStatusSuccess returns a 500 error when fetching monthly successful top-up stats fails.
var ErrApiFailedFindMonthlyTopupStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful topups", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupStatusSuccess returns a 500 error when fetching yearly successful top-up stats fails.
var ErrApiFailedFindYearlyTopupStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful topups", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupStatusFailed returns a 500 error when fetching monthly failed top-up stats fails.
var ErrApiFailedFindMonthlyTopupStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed topups", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupStatusFailed returns a 500 error when fetching yearly failed top-up stats fails.
var ErrApiFailedFindYearlyTopupStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed topups", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupStatusSuccessByCardNumber returns a 500 error when fetching monthly successful top-ups by card number fails.
var ErrApiFailedFindMonthlyTopupStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful topups by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupStatusSuccessByCardNumber returns a 500 error when fetching yearly successful top-ups by card number fails.
var ErrApiFailedFindYearlyTopupStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful topups by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupStatusFailedByCardNumber returns a 500 error when fetching monthly failed topups by card number fails.
var ErrApiFailedFindMonthlyTopupStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed topups by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupStatusFailedByCardNumber returns a 500 error when fetching yearly failed topups by card number fails.
var ErrApiFailedFindYearlyTopupStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed topups by card number", http.StatusInternalServerError)
}
