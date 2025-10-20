package withdrawapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyWithdrawStatusSuccess is an error response for failed to fetch monthly successful withdraws.
var ErrApiFailedFindMonthlyWithdrawStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawStatusSuccess is an error response for failed to fetch yearly successful withdraws.
var ErrApiFailedFindYearlyWithdrawStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawStatusFailed is an error response for failed to fetch monthly failed withdraws.
var ErrApiFailedFindMonthlyWithdrawStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawStatusFailed is an error response for failed to fetch yearly failed withdraws.
var ErrApiFailedFindYearlyWithdrawStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawStatusSuccessCardNumber is an error response for failed to fetch monthly successful withdraws by card number.
var ErrApiFailedFindMonthlyWithdrawStatusSuccessCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful withdraws by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawStatusSuccessCardNumber is an error response for failed to fetch yearly successful withdraws by card number.
var ErrApiFailedFindYearlyWithdrawStatusSuccessCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful withdraws by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawStatusFailedCardNumber is an error response for failed to fetch monthly failed withdraws by card number.
var ErrApiFailedFindMonthlyWithdrawStatusFailedCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed withdraws by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawStatusFailedCardNumber is an error response for failed to fetch yearly failed withdraws by card number.
var ErrApiFailedFindYearlyWithdrawStatusFailedCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed withdraws by card number", http.StatusInternalServerError)
}
