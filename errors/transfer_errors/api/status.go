package transactionapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyTransferStatusSuccess returns an API error response when fetching monthly successful transfer data fails.
var ErrApiFailedFindMonthlyTransferStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferStatusSuccess returns an API error response when fetching yearly successful transfer data fails.
var ErrApiFailedFindYearlyTransferStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferStatusFailed returns an API error response when fetching monthly failed transfer data fails.
var ErrApiFailedFindMonthlyTransferStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferStatusFailed returns an API error response when fetching yearly failed transfer data fails.
var ErrApiFailedFindYearlyTransferStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferStatusSuccessByCardNumber returns an API error response when fetching monthly successful transfers by card number fails.
var ErrApiFailedFindMonthlyTransferStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transfers by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferStatusSuccessByCardNumber returns an API error response when fetching yearly successful transfers by card number fails.
var ErrApiFailedFindYearlyTransferStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transfers by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferStatusFailedByCardNumber returns an API error response when fetching monthly failed transfers by card number fails.
var ErrApiFailedFindMonthlyTransferStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transfers by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferStatusFailedByCardNumber returns an API error response
// when fetching yearly failed transfers by card number fails.
var ErrApiFailedFindYearlyTransferStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transfers by card number", http.StatusInternalServerError)
}
