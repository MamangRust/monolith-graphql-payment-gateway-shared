package cardapierrors

import (
	"net/http"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiFailedFindMonthlyTopupAmountByCard returns an API error response when fetching the monthly top-up amount by card fails.
var ErrApiFailedFindMonthlyTopupAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupAmountByCard returns an API error response when fetching the yearly top-up amount by card fails.
var ErrApiFailedFindYearlyTopupAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawAmountByCard returns an API error response when fetching the monthly withdrawal amount by card fails.
var ErrApiFailedFindMonthlyWithdrawAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawAmountByCard returns an API error response when fetching the yearly withdrawal amount by card fails.
var ErrApiFailedFindYearlyWithdrawAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransactionAmountByCard returns an API error response when fetching the monthly transaction amount by card fails.
var ErrApiFailedFindMonthlyTransactionAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transaction amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransactionAmountByCard returns an API error response when fetching the yearly transaction amount by card fails.
var ErrApiFailedFindYearlyTransactionAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transaction amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferSenderAmountByCard represents error when failing to fetch monthly transfer amount by sender card.
var ErrApiFailedFindMonthlyTransferSenderAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer sender amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferSenderAmountByCard represents error when failing to fetch yearly transfer amount by sender card.
var ErrApiFailedFindYearlyTransferSenderAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer sender amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferReceiverAmountByCard represents error when failing to fetch monthly transfer amount by receiver card.
var ErrApiFailedFindMonthlyTransferReceiverAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer receiver amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferReceiverAmountByCard represents error when failing to fetch yearly transfer amount by receiver card.
var ErrApiFailedFindYearlyTransferReceiverAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer receiver amount by card", http.StatusInternalServerError)
}
